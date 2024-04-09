package slice

import "fmt"

type Task struct {
	Name      string
	Minutes   int
	EventType int
}

func main() {
	tasks := Slice[Task]{}

	task1 := Task{
		Name:      "Task 1",
		Minutes:   30,
		EventType: 1,
	}

	tasks.Push(task1).Push(Task{Name: "Task 2", EventType: 2})
	fmt.Println(tasks.Join(","))

	// tasks.Filter(func(task Task) bool {
	// 	return task.EventType == 1
	// })

	// tasks.Map(func(task Task) Task {
	// 	task.EventType = 5
	// 	return task
	// })

	// tasks.Some(func(task Task) bool {
	// 	return task.EventType == 1
	// })

	Reduce(tasks, 0, func(result int, task Task) int {
		return result + task.EventType
	})

}
