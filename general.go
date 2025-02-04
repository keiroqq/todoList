package main

import (
	"fmt"
	"time"
)

const todo, inProgress, done = 0, 1, 2

type TaskArr struct {
	Tasks     []task
	IdCounter int
}

type task struct {
	Id          int
	Description string
	Status      status
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func NewTaskList() *TaskArr {
	return &TaskArr{Tasks: make([]task, 0), IdCounter: 0}
}

type status string

func (s *status) updateStatus(newStatus int) {
	switch newStatus {
	case inProgress:
		*s = "In progress"
	case done:
		*s = "Done"
	default:
		*s = "To do"
	}
}

func (taskArr *TaskArr) NewTask(description string) (int, error) {
	if len(description) == 0 {
		return -1, fmt.Errorf("task is empty")
	}
	stat := new(status)
	stat.updateStatus(0)
	newtask := task{
		Id:          taskArr.IdCounter,
		Description: description,
		Status:      *stat,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now()}
	taskArr.Tasks = append(taskArr.Tasks, newtask)
	taskArr.IdCounter++
	return newtask.Id, nil
}

func (taskArr *TaskArr) Update(id int, description string) (int, error) {
	if id == len(taskArr.Tasks) {
		return -1, fmt.Errorf("task with ID: %d doesn't exist", id)
	}

	if description == "" {
		return -1, fmt.Errorf("task is empty")
	}
	taskArr.Tasks[id].Description = description
	taskArr.Tasks[id].UpdatedAt = time.Now()

	return id, nil
}

func (taskArr *TaskArr) Delete(id int) (int, error) {
	if id == len(taskArr.Tasks) {
		return -1, fmt.Errorf("task with ID: %d doesn't exist", id)
	}
	deleted := append(taskArr.Tasks[:id], ordering(taskArr.Tasks[id+1:])...)
	taskArr.Tasks = deleted
	taskArr.IdCounter--
	return id, nil
}

func ordering(arr []task) []task {
	for i := 0; i < len(arr); i++ {
		arr[i].Id--
	}
	return arr
}

func (taskArr *TaskArr) SetInProgress(id int) (int, error) {
	if id == len(taskArr.Tasks) {
		return -1, fmt.Errorf("task with ID: %d doesn't exist", id)
	}
	taskArr.Tasks[id].Status.updateStatus(inProgress)
	taskArr.Tasks[id].UpdatedAt = time.Now()

	return id, nil
}

func (taskArr *TaskArr) SetDone(id int) (int, error) {
	if id == len(taskArr.Tasks) {
		return -1, fmt.Errorf("task with ID: %d doesn't exist", id)
	}
	taskArr.Tasks[id].Status.updateStatus(done)
	taskArr.Tasks[id].UpdatedAt = time.Now()

	return id, nil
}

func (task task) String() string {
	return fmt.Sprintf("ID: %d\tStatus: %s\n%s\nCreated: %v\tUpdated: %v\n", task.Id, task.Status, task.Description, task.CreatedAt.Format(time.ANSIC), task.UpdatedAt.Format(time.ANSIC))
}

func (taskArr TaskArr) list(status status) {
	if status == "" {
		for _, task := range taskArr.Tasks {
			fmt.Println(task)
		}
		return
	}
	for _, task := range taskArr.Tasks {
		if task.Status == status {
			fmt.Println(task)
		}
	}
}
