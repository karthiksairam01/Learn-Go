# It's project time!

Based on everything learned so far, it's important to gather them and put it to test! Here's a project to test my understanding so far!

# Project: Simple Task Management System

## Question:

Create a simple task management system in Go to practice your knowledge of primitive types, composite types, blocks, shadowing, and control structures. The system should allow you to add tasks, view tasks, and mark tasks as completed.

## Specifications:

### Task Struct:

Define a struct named `Task` with the following fields:
- **Name** (string): The name of the task.
- **Description** (string): A brief description of the task.
- **Completed** (bool): A flag indicating whether the task is completed.

### Task List:

Use a slice to store multiple tasks.

### Functions:

Implement the following functions:
- `addTask(name, description string)`: Adds a new task to the task list.
- `viewTasks()`: Prints all tasks with their details (name, description, completion status).
- `completeTask(index int)`: Marks a task as completed based on its index in the task list.

### Control Structures:

Use if-else statements, for loops, and switch statements to control the flow of the program in the main function.

### Main Function:

Implement a main function that:
- Displays a menu with options to add a task, view all tasks, complete a task, and exit the program.