# Kanban CLI

An interactive command-line Kanban task manager written in Go. The program allows users to create, view, update, filter, and delete tasks. Task data is saved in `tasks.txt` and loaded again when the program starts.

## How to Run

Run the interactive menu:


go run main.go

Show help:

go run main.go --help

You can also use:

go run main.go -h


The program uses only the Go standard library.

## Menu

The interactive menu provides the following actions:

1. **Create task** — adds a new task with `TODO` status.
2. **Change status** — moves a task through `TODO → IN_PROGRESS → DONE`.
3. **Delete task** — removes a task by ID.
4. **All tasks** — displays the task board as an ASCII table.
5. **Exit** — saves the current state and closes the program.
6. **Filter by status** — displays tasks with a selected status.
7. **Edit title** — changes the title of a task by ID.

The program validates user input and does not crash on invalid input.

## Task Statuses

Tasks can have three statuses:

* `TODO` — task has not been started.
* `IN_PROGRESS` — task is currently being worked on.
* `DONE` — task is completed.

The status changes step by step:

```text
TODO → IN_PROGRESS → DONE
```

## File Format

Tasks are stored in `tasks.txt`, one task per line:

```text
id;status;title
```

Example:

```text
1;DONE;Write tests
2;IN_PROGRESS;Fix bug
3;TODO;Deploy
```

Task IDs are unique and new IDs are generated as `max(existing IDs) + 1`.

## Example

After creating three tasks and selecting **All tasks**:

```text
+----+----------------------+-------------+
| ID | Title                | Status      |
+----+----------------------+-------------+
|  1 | Write tests          | TODO        |
|  2 | Fix bug              | TODO        |
|  3 | Deploy               | TODO        |
+----+----------------------+-------------+
```

After changing task `1` to `DONE`:

```text
+----+----------------------+-------------+
| ID | Title                | Status      |
+----+----------------------+-------------+
|  1 | Write tests          | DONE        |
|  2 | Fix bug              | TODO        |
|  3 | Deploy               | TODO        |
+----+----------------------+-------------+
```

## Team 33

| Name                | GitHub       |
| ------------------- | ------------ |
| Bakytgul Bolatkhan  | `@babolat`   |
| Kamila Kengesbekova | `@kaken`     |
| Sati Kabylbek       | `@skabylbek` |
