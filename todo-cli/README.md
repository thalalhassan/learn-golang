# Todo CLI Application

This is a simple command-line interface (CLI) application for managing a to-do list. The application is built using Go and leverages the local flag library for handling command-line arguments.

## Features

- Add tasks to your to-do list.
- View all tasks .
- Delete tasks from the list.

## Prerequisites

- Go 1.18 or later installed on your system.

## Installation

1. Clone the repository:
    ```bash
    git clone https://github.com/your-username/todo-cli.git
    cd todo-cli
    ```

2. Build the application:
    ```bash
    go build -o todo
    ```

3. Run the application:
    ```bash
    ./todo
    ```

## Usage

Below are some example commands:

- Add a new task:
  ```bash
  ./todo --add "Buy groceries"
  ```

- List all tasks:
  ```bash
  ./todo --list
  ```

- Delete a task:
  ```bash
  ./todo --delete 1
  ```
