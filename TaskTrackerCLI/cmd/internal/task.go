package internal

import (
	"encoding/json"
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/MiguelLTadeo/TaskTrackerCLI_GO.git/cmd/utils"
)

type Task struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	Status      int       `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Status int

const (
	StatusTodo Status = iota
	StatusInProgress
	StausDone
)

func CreateTask(description string) error {
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	path := filepath.Join(home, "tasks.json")
	//criação do arquivo json
	file, err := utils.CreateOrGetFile(path)

	if err != nil {
		return err
	}

	tasks, newId, err := GenID(file)

	if err != nil {
		return err
	}

	newTask := Task{
		ID:          *newId,
		Description: description,
		Status:      int(StatusTodo),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	tasks = append(tasks, newTask)

	err = WriteTasks(tasks, path)
	if err != nil {
		return err
	}
	return nil
}

func UpdateTask(id int, description string) error {
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	path := filepath.Join(home, "tasks.json")
	//criação do arquivo json
	file, err := utils.CreateOrGetFile(path)

	if err != nil {
		return err
	}
	tasks, _, err := GenID(file)

	for i := range tasks {
		if tasks[i].ID == id {
			tasks[i].Description = description
			tasks[i].UpdatedAt = time.Now()
		}
	}

	err = WriteTasks(tasks, path)

	return nil
}

func GenID(file *os.File) ([]Task, *int, error) {
	data, err := io.ReadAll(file)
	if err != nil {
		return nil, nil, err
	}

	if len(data) == 0 {
		nextID := 1
		return nil, &nextID, nil
	}

	tasks := []Task{}

	err = json.Unmarshal(data, &tasks)
	if err != nil {
		return nil, nil, err
	}

	if len(tasks) == 0 {
		num := 1
		return nil, &num, nil
	}

	maxID := tasks[0].ID
	for _, task := range tasks[1:] {
		if task.ID > maxID {
			maxID = task.ID
		}
	}

	nextID := maxID + 1
	return tasks, &nextID, nil
}

func WriteTasks(tasks []Task, path string) error {
	data, err := json.Marshal(&tasks)
	if err != nil {
		return err
	}
	err = os.WriteFile(path, data, 0644)
	if err != nil {
		return err
	}
	return nil
}
