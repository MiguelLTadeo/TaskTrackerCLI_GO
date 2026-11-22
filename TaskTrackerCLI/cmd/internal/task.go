package internal

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
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
	StatusDone
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

	if err != nil {
		return err
	}

	return nil
}

func DeleteTask(id int) error {
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

	var iDel int = -1

	for i := range tasks {
		if tasks[i].ID == id {
			iDel = i
			break
		}
	}

	if iDel != -1 {
		tasks = append(tasks[:iDel], tasks[iDel+1:]...)
	} else {
		return fmt.Errorf("Task not found!")
	}

	err = WriteTasks(tasks, path)

	if err != nil {
		return err
	}

	return nil
}

func MarkInProgres(id int) error {
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
			tasks[i].Status = int(StatusInProgress)
		}
	}

	err = WriteTasks(tasks, path)

	if err != nil {
		return err
	}

	return nil
}

func MarkDone(id int) error {
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
			tasks[i].Status = int(StatusDone)
		}
	}

	err = WriteTasks(tasks, path)

	if err != nil {
		return err
	}

	return nil
}

func ListTasks(status string) {
	switch status {
	case "todo":
		List(0)
	case "in-progress":
		List(1)
	case "done":
		List(2)
	default:
		fmt.Println("Invalid Status")
	}
}

func List(status int) {
	home, err := os.UserHomeDir()
	if err != nil {
		return
	}

	path := filepath.Join(home, "tasks.json")
	//criação do arquivo json
	file, err := utils.CreateOrGetFile(path)

	if err != nil {
		return
	}
	tasks, _, err := GenID(file)

	fmt.Println("ID" + "   " + "Status" + "   " + "Description" + "   " + "Created At" + "   " + "Updated At")

	for i := range tasks {
		if tasks[i].Status == status {
			fmt.Println(strconv.Itoa(tasks[i].ID) + "   " + StatusString(Status(tasks[i].Status)) + "   " + tasks[i].Description + "   " + tasks[i].CreatedAt.String() + "   " + tasks[i].UpdatedAt.String())
		}
	}
}

func StatusString(status Status) string {
	switch status {
	case StatusTodo:
		return "To Do"
	case StatusInProgress:
		return "In Progress"
	case StatusDone:
		return "Done"
	default:
		return "Unknown"
	}
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
