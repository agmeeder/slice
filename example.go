package slice

import "fmt"

type Task struct {
	Name      string
	EventType int
}

func main() {
	tasks := Slice[Task]{}

	task1 := Task{
		Name:      "Task 1",
		EventType: 1,
	}

	tasks.Push(task1).Push(Task{Name: "Task 2", EventType: 2})
	fmt.Println(tasks.Join(","))
}
