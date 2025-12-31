package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"
)

// FUNCTION THAT CHECKS IF THE FILE EXISTS OR NOT:
func fileExists(filePath string) bool {
	_, err := os.Stat(filePath)
	return !errors.Is(err, os.ErrNotExist)
}

// THE TASK STRUCT:
type Task struct {
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Status    string `json:"status"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func main() {

	fmt.Println(os.Args)
	command := os.Args[1]
	fileName := "tasks.json"
	if fileExists(fileName) {
		f, err := os.Open(fileName)
		if err != nil {
			fmt.Println("Error opening file:", err)
			return
		}
		defer f.Close()

	} else {
		f, err := os.Create(fileName)
		if err != nil {
			fmt.Println("Error creating file:", err)
			return
		}
		defer f.Close()

	}
	content, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	var taskSlice []Task

	if len(content) != 0 {
		err = json.Unmarshal(content, &taskSlice)
		if err != nil {
			fmt.Println("Error unmarshaling content:", err)
			return
		}

	} else {
		taskSlice = []Task{}
	}
	switch command {

	// THE ADD TASK COMMAND'S CODE:
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Correct way: task-cli add <title>")
			return
		}
		title := os.Args[2]
		var maxId int = 0
		for _, t := range taskSlice {
			if t.Id > maxId {
				maxId = t.Id
			}
		}
		maxId += 1
		newTask := Task{
			Id:        maxId,
			Title:     title,
			Status:    "todo",
			CreatedAt: time.Now().Format(time.RFC3339),
		}
		taskSlice = append(taskSlice, newTask)
		updatedJSONdata, err := json.MarshalIndent(taskSlice, "", "  ")
		if err != nil {
			fmt.Println("Error marshalling JSON:", err)
			return
		}
		err = os.WriteFile(fileName, updatedJSONdata, 0644)
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}
		fmt.Println("Task added successfully (ID:", newTask.Id, ")")

	// THE UPDATE TASK COMMAND'S CODE:
	case "update":
		if len(os.Args) < 4 {
			fmt.Println("Correct way: task-cli update <ID> <new title>")
			return
		}
		title := os.Args[3]
		number, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Error: invalid number provided")
			return
		}
		found := false
		for i, aim := range taskSlice {
			if aim.Id == number {
				found = true
				taskSlice[i].Title = title
				taskSlice[i].UpdatedAt = time.Now().Format(time.RFC3339)
			}
		}
		if found != true {
			fmt.Println("The Id that you provided doesn't exist <try again>")
		}

		updatedJSONdata, err := json.MarshalIndent(taskSlice, "", "  ")
		if err != nil {
			fmt.Println("Error marshalling JSON:", err)
			return
		}
		err = os.WriteFile(fileName, updatedJSONdata, 0644)
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}
		fmt.Println("Task updated successfully (Id:", number, ")")

	// THE DELETE COMMAND'S CODE:
	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("Correct way: task-cli delete <ID>")
			return
		}
		number, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Error: invalid number provided")
			return
		}
		found := false
		newSlice := []Task{}
		for _, aim := range taskSlice {
			if aim.Id == number {
				found = true
			} else {
				newSlice = append(newSlice, aim)
			}
		}
		if found != true {
			fmt.Println("The Id that you provided doesn't exist <try again>")
		}
		if found == true {
			taskSlice = newSlice
			updatedJSONdata, err := json.MarshalIndent(taskSlice, "", "  ")
			if err != nil {
				fmt.Println("Error marshalling JSON:", err)
				return
			}
			err = os.WriteFile(fileName, updatedJSONdata, 0644)
			if err != nil {
				fmt.Println("Error writing to file:", err)
				return
			}
			fmt.Println("Task deleted successfully (Id:", number, ")")
		}

	// THE MARK-DONE COMMAND'S CODE:
	case "mark-done":
		if len(os.Args) < 3 {
			fmt.Println("Correct way: task-cli mark-done <ID>")
			return
		}
		number, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Error: invalid number provided")
			return
		}
		found := false
		for i, aim := range taskSlice {
			if aim.Id == number {
				found = true
				taskSlice[i].Status = "done"
				taskSlice[i].UpdatedAt = time.Now().Format(time.RFC3339)
			}
		}
		if found != true {
			fmt.Println("The Id that you provided doesn't exist <try again>")
		}
		updatedJSONdata, err := json.MarshalIndent(taskSlice, "", "  ")
		if err != nil {
			fmt.Println("Error marshalling JSON:", err)
			return
		}
		err = os.WriteFile(fileName, updatedJSONdata, 0644)
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}
		fmt.Println("Task marked-done successfully")

	// THE MARK-IN-PROGRESS COMMAND'S CODE
	case "mark-in-progress":
		if len(os.Args) < 3 {
			fmt.Println("Correct way: task-cli mark-in-progress <ID>")
			return
		}
		number, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Error: invalid number provided")
			return
		}
		found := false
		for i, aim := range taskSlice {
			if aim.Id == number {
				found = true
				taskSlice[i].Status = "in-progress"
				taskSlice[i].UpdatedAt = time.Now().Format(time.RFC3339)
			}
		}
		if found != true {
			fmt.Println("The Id that you provided doesn't exist <try again>")
		}
		updatedJSONdata, err := json.MarshalIndent(taskSlice, "", "  ")
		if err != nil {
			fmt.Println("Error marshalling JSON:", err)
			return
		}
		err = os.WriteFile(fileName, updatedJSONdata, 0644)
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}
		fmt.Println("Task marked-in-progress successfully")

	// THE LIST COMMAND'S CODE:
	case "list":
		var filter string
		if len(os.Args) >= 3 {
			filter = os.Args[2]
		}
		if filter == "" {
			fmt.Println(taskSlice)
		} else if filter == "in-progress" {
			for i, aim := range taskSlice {
				if taskSlice[i].Status == "in-progress" {
					fmt.Println(aim)
				}
			}
		} else if filter == "done" {
			for i, aim := range taskSlice {
				if taskSlice[i].Status == "done" {
					fmt.Println(aim)
				}
			}
		} else if filter == "todo" {
			for i, aim := range taskSlice {
				if taskSlice[i].Status == "todo" {
					fmt.Println(aim)
				}
			}
		}

	//	THE HELP COMMAND'S CODE:
	case "help":
		if len(os.Args) == 2 {
			fmt.Println("\n# Adding a new task\ntask-cli add Buy groceries ")
			fmt.Println("\n# Updating and deleting tasks\ntask-cli update 1 Buy groceries and cook dinner\ntask-cli delete 1")
			fmt.Println("\n# Marking a task as in progress or done\ntask-cli mark-in-progress 1\ntask-cli mark-done 1")
			fmt.Println("\n# Listing all tasks\ntask-cli list")
			fmt.Println("\n# Listing tasks by status\ntask-cli list done\ntask-cli list todo\ntask-cli list in-progress")
			fmt.Println("\n# Help with all commands\n task-cli help")
			return
		}
	// IF COMMAND DOESN'T EXIST:
	default:
		fmt.Println("Unknown command:", command)
	}
}
