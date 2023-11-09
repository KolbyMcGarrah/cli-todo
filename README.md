# cli-todo
A cli tool todo app for managing tasks without leaving the terminal. This tool was written in Go using the charm libraries, bubbletea, lipgloss and bubbles so that I could experience building a Terminal User Interface (TUI) while also building something that I wanted to have.

## Usage
run the `todo` or `cli-todo` command to bring up the terminal interface

to add a new task, press the `a` button while in the list view. This will bring up a prompt to enter the task details. After inputing the task details, press the `enter` key and it will bring you back to the `list` view

in the `list` view, navigate the tasks with the arrow keys or the `k` `j` keys. select tasks by pressing the `space` key or `enter` key. Delete selected tasks by pressing the `d` key.

## Installation
This is available to install via homebrew, or can be run as a binary by cloning the repo and using the 
```Go
go build -o <name-that-you-want>
```

### Installing with Homebrew
After installing homebrew, you will need to connect to the tap with the following command
```shell
brew tap kolbymcgarrah/kolbymcgarrah
```

After that succeeds, you can install the tool with
```shell
brew install kolbymcgarrah/kolbymcgarrah/todo
```

Verify the installation by running 
```shell
which todo
```
if that shows up empty, try
```shell
which cli-todo
```

whichever shows a result will be your command for using the tool. (It should show up with `todo` but didn't on my first installation on my mac.)

## Future Plans
I plan on allowing different projects and statuses to be added to the tasks down the line and will allow a `kanban`-esk table to view your progress on your tasks.