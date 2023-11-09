package db

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"os"

	"github.com/boltdb/bolt"
	"github.com/kolbymcgarrah/cli-todo/internal/task"
)

type TaskDB struct {
	db *bolt.DB
}

func OpenDB() (*TaskDB, error) {
	err := os.MkdirAll("~/.cli-todo", 0777)
	if err != nil {
		fmt.Println(err)
	}
	db, err := bolt.Open("~/.cli-todo/task.db", 0777, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}
	return &TaskDB{
		db: db,
	}, nil
}

func (t *TaskDB) Close() error {
	return t.db.Close()
}

func (t *TaskDB) CreateTaskBucket() error {
	return t.db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("tasks"))
		if err != nil {
			return fmt.Errorf("error creating bucket: %w", err)
		}
		return nil
	})
}

func (t *TaskDB) AddTask(tsk task.Task) error {
	return t.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("tasks"))
		id, _ := b.NextSequence()
		tsk.ID = id
		data, err := tsk.ToBytes()
		if err != nil {
			return fmt.Errorf("error converting task to bytes: %w", err)
		}
		err = b.Put(itob(id), data)
		if err != nil {
			return fmt.Errorf("error putting data in bucket: %w", err)
		}
		return nil
	})
}

func itob(v uint64) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, v)
	return b
}

func bytesToTask(data []byte) (task.Task, error) {
	var tsk task.Task
	err := json.Unmarshal(data, &tsk)
	if err != nil {
		return tsk, fmt.Errorf("error unmarshaling data to task: %w", err)
	}
	return tsk, nil
}

func (t *TaskDB) GetTasks() ([]task.Task, error) {
	var tasks []task.Task
	err := t.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("tasks"))
		return b.ForEach(func(k, v []byte) error {
			tsk, err := bytesToTask(v)
			if err != nil {
				return fmt.Errorf("error parsing task: %w", err)
			}
			id := binary.BigEndian.Uint64(k)
			tsk.ID = id
			tasks = append(tasks, tsk)
			return nil
		})
	})
	return tasks, err
}

func (t *TaskDB) DeleteTask(id uint64) error {
	return t.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("tasks"))
		return b.Delete(itob(id))
	})
}
