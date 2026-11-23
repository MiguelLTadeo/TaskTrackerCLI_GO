package internal

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/MiguelLTadeo/TaskTrackerCLI_GO.git/utils"
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
	// create or open the JSON file
	file, err := utils.CreateOrGetFile(path)

	if err != nil {
		return err
	}
	defer file.Close()

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
	// create or open the JSON file
	file, err := utils.CreateOrGetFile(path)

	if err != nil {
		return err
	}
	defer file.Close()

	tasks, _, err := GenID(file)

	if err != nil {
		return err
	}

	found := false

	for i := range tasks {
		if tasks[i].ID == id {
			tasks[i].Description = description
			tasks[i].UpdatedAt = time.Now()
			found = true
		}
	}

	if found == false {
		return fmt.Errorf("Task not found!")
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
	// create or open the JSON file
	file, err := utils.CreateOrGetFile(path)

	if err != nil {
		return err
	}

	defer file.Close()

	tasks, _, err := GenID(file)

	if err != nil {
		return err
	}

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

func MarkInProgress(id int) error {
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	path := filepath.Join(home, "tasks.json")
	// create or open the JSON file
	file, err := utils.CreateOrGetFile(path)

	if err != nil {
		return err
	}

	defer file.Close()

	tasks, _, err := GenID(file)

	if err != nil {
		return err
	}

	found := false

	for i := range tasks {
		if tasks[i].ID == id {
			tasks[i].Status = int(StatusInProgress)
			tasks[i].UpdatedAt = time.Now()
			found = true
		}
	}

	if found == false {
		return fmt.Errorf("Task not found!")
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
	// create or open the JSON file
	file, err := utils.CreateOrGetFile(path)

	if err != nil {
		return err
	}

	defer file.Close()

	tasks, _, err := GenID(file)

	if err != nil {
		return err
	}

	found := false

	for i := range tasks {
		if tasks[i].ID == id {
			tasks[i].Status = int(StatusDone)
			tasks[i].UpdatedAt = time.Now()
			found = true
		}
	}

	if found == false {
		return fmt.Errorf("Task not found!")
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
		fmt.Println("Invalid status!")
	}
}

func List(status int) {
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	path := filepath.Join(home, "tasks.json")
	// create or open the JSON file
	file, err := utils.CreateOrGetFile(path)

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer file.Close()
	tasks, _, err := GenID(file)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for i := range tasks {
		if tasks[i].Status == status {
			fmt.Println("ID:" + strconv.Itoa(tasks[i].ID) + "   STATUS:" + StatusString(Status(tasks[i].Status)) + "   DESCRIPTION:" + tasks[i].Description + "   CREATED AT:" + tasks[i].CreatedAt.String() + "   UPDATED AT:" + tasks[i].UpdatedAt.String())
		}
	}
}

func ListAll() {
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	path := filepath.Join(home, "tasks.json")
	// create or open the JSON file
	file, err := utils.CreateOrGetFile(path)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer file.Close()

	tasks, _, err := GenID(file)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for i := range tasks {

		fmt.Println("ID:" + strconv.Itoa(tasks[i].ID) + "   STATUS:" + StatusString(Status(tasks[i].Status)) + "   DESCRIPTION:" + tasks[i].Description + "   CREATED AT:" + tasks[i].CreatedAt.String() + "   UPDATED AT:" + tasks[i].UpdatedAt.String())

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
		return nil, nil, fmt.Errorf("could not parse tasks file: %w", err)
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
