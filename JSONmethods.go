package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

func (taskArr *TaskArr) LoadData() error {
	file, err := os.Open("tasks.json")
	defer file.Close()
	if err != nil {
		return fmt.Errorf("cannot open file: %v", err)
	}
	if err = json.NewDecoder(file).Decode(taskArr); err != nil {
		return fmt.Errorf("cannot decode file: %v", err)
	}
	return nil
}

func (taskArr *TaskArr) StoreData() error {
	file, err := os.OpenFile("tasks.json", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	defer file.Close()
	if err != nil {
		return fmt.Errorf("cannot open file: %v", err)
	}
	if err = json.NewEncoder(file).Encode(taskArr); err != nil {
		return fmt.Errorf("cannot encode data: %v", err)
	}
	return nil
}

func (taskArr *TaskArr) UnmarshalJSON(b []byte) error {
	var tmp struct {
		Tasks     []task
		IdCounter int
	}

	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}

	taskArr.Tasks = tmp.Tasks
	taskArr.IdCounter = tmp.IdCounter
	return nil
}

func (t *task) UnmarshalJSON(b []byte) error {
	var tmp struct {
		Id          int       `json:"Id"`
		Description string    `json:"Description"`
		Status      string    `json:"Status"`
		CreatedAt   time.Time `json:"CreatedAt"`
		UpdatedAt   time.Time `json:"UpdatedAt"`
	}

	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}

	t.Id = tmp.Id
	t.Description = tmp.Description
	t.Status = status(tmp.Status)
	t.CreatedAt = tmp.CreatedAt
	t.UpdatedAt = tmp.UpdatedAt

	return nil
}
