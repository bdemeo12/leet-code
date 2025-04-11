package main

import (
	"fmt"
	"sort"
)

// Design a Todo List Where users can add tasks, mark them as complete, or get a list of pending tasks.
//  Users can also add tags to tasks and can filter the tasks by certain tags.

// Implement the TodoList class:

// TodoList() Initializes the object.

// int addTask(int userId, String taskDescription, int dueDate, List<String> tags)
//  Adds a task for the user with the ID userId with a due date equal to dueDate and a list
// of tags attached to the task. The return value is the ID of the task.
//  This ID starts at 1 and is sequentially increasing.
//  That is, the first task's id should be 1, the second task's id should be 2, and so on.

// List<String> getAllTasks(int userId) Returns a list of all the tasks not marked as complete for the user
//  with ID userId, ordered by the due date. You should return an empty list if the user has no uncompleted tasks.

// List<String> getTasksForTag(int userId, String tag) Returns a list of all the tasks that are not marked as
//  complete for the user with the ID userId and have tag as one of their tags, ordered by their due date.
// Return an empty list if no such task exists.

// void completeTask(int userId, int taskId) Marks the task with the ID taskId as completed only
// if the task exists and the user with the ID userId has this task, and it is uncompleted.

type TodoList struct {
	todoList map[int]Tasks // map of userID to tasks
	nextId   int
}

type Tasks struct {
	taskList []TaskData // map of taskID to each task details
}

type TaskData struct {
	taskId      int
	description string
	dueDate     int
	tags        []string
	complete    bool
}

func Constructor() TodoList {
	todoList := make(map[int]Tasks)
	return TodoList{
		todoList: todoList,
		nextId:   1,
	}
}

func (this *TodoList) AddTask(userId int, taskDescription string, dueDate int, tags []string) int {

	tasks := this.todoList[userId]
	taskId := this.nextId
	this.nextId++

	tasks.taskList = append(tasks.taskList, TaskData{
		taskId:      taskId,
		description: taskDescription,
		dueDate:     dueDate,
		tags:        tags,
	})

	this.todoList[userId] = tasks // save back!

	return taskId
}

func (this *TodoList) GetAllTasks(userId int) []string {
	uncompletedTasks := []string{}

	tasks := this.todoList[userId]

	// sort tasks by due date so that when we append descriptions, they are in order
	sort.Slice(tasks.taskList, func(a, b int) bool {
		return tasks.taskList[a].dueDate < tasks.taskList[b].dueDate
	})

	for _, task := range tasks.taskList {
		if !task.complete {
			uncompletedTasks = append(uncompletedTasks, task.description)
		}
	}

	return uncompletedTasks
}

func (this *TodoList) GetTasksForTag(userId int, tag string) []string {
	uncompletedTasks := []string{}

	tasks := this.todoList[userId]

	// sort tasks by due date so that when we append descriptions, they are in order
	sort.Slice(tasks.taskList, func(a, b int) bool {
		return tasks.taskList[a].dueDate < tasks.taskList[b].dueDate
	})

	for _, task := range tasks.taskList {
		if !task.complete {
			for _, t := range task.tags {
				if tag == t {
					uncompletedTasks = append(uncompletedTasks, task.description)
					break
				}
			}
		}
	}

	return uncompletedTasks
}

func (this *TodoList) CompleteTask(userId int, taskId int) {

	tasks := this.todoList[userId] // get the struct value

	for i := range tasks.taskList {
		if tasks.taskList[i].taskId == taskId {
			tasks.taskList[i].complete = true
			break
		}
	}

	this.todoList[userId] = tasks // save back!

}

func main() {
	obj := Constructor()

	obj.AddTask(1, "Task1", 50, []string{})
	obj.AddTask(1, "Task2", 100, []string{"P1"})

	fmt.Println(obj.GetAllTasks(1))
}

/**
 * Your TodoList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.AddTask(userId,taskDescription,dueDate,tags);
 * param_2 := obj.GetAllTasks(userId);
 * param_3 := obj.GetTasksForTag(userId,tag);
 * obj.CompleteTask(userId,taskId);
 */
