# Taskky CLI

Taskky is an CLI application used to track and manage your tasks. With Tasky you can track what you need to do, what you have done, and what you are currently working on. This is a solution to a challenge from [Roadmap.sh](https://roadmap.sh/projects/task-tracker)

# Requirements
The application should run from the command line, accept user actions and inputs as arguments, and store the tasks in a JSON file. The user should be able to:

- Add, Update, and Delete tasks
- Mark a task as in progress or done
- List all tasks
- List all tasks that are done
- List all tasks that are not done
- List all tasks that are in progress

# Usage
Taskky accepts commands and arguments to manage tasks. Below are the supported commands:

Add a Task
```bash
./taskky add "Task description"
```
Update a Task

```bash

./taskky update <task-id> "Updated description"
```

Delete a Task
```bash
./taskky delete <task-id>
```
Mark a Task as "In Progress"

```bash
./taskky progress <task-id>
```

Mark a Task as "Done"
```bash
./taskky done <task-id>
```

List All Tasks
```bash
./taskky list
```

List Tasks by Status Done
```bash
./taskky list done
```
List Tasks by In Progress
```bash
./taskky list in-progress
```

## Storage

Taskky uses a tasks.json file in the project directory to store and manage tasks. Ensure the file has appropriate read/write permissions.