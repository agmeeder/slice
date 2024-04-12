# slice
A Go slice package with a generic Slice[T] type and Javascript like methods to manipulate the slice

Add the package to your Go project

````go
go get github.com/agmeeder/slice
````

Example code:

````go
// Example struct
type Task struct {
  EventType int
  Minutes   int
  Name      string
}

// Define a slice 'tasks' of type 'Task'
tasks := Slice[Task]{}

// Create a new task
task1 := Task{
  Name:      "Task 1",
  Minutes:   30,
  EventType: 1,
}

// Add the task to the slice and add a second one 'inline'
tasks.Push(task1).Push(Task{EventType: 1, Minutes: 15, Name: "Task 2"})

// Print the slice as a comma separated string
fmt.Println(tasks.Join(","))

// Search for a specific task where the EventType = 1
t, found := tasks.Find(func(task Task) bool {
  return task.EventType == 1
})

// Print the found task and if it was found
fmt.Println(t.Name, found)
````

# Methods

The Slice packages contains the following methods:

## Basic methods
- At
- Concat
- Join
- Length
- Merge
- Pop
- Push
- Shift
- Slice
- Splice
- ToSliced
- ToSpliced
- ToString
- Unshift

## Sort methods
- Reverse
- Sort
- ToReversed
- ToSorted

## Iteration methods
- Every
- Filter
- Map
- Some
- Reduce (not a method but a function)

## Search methods
- Find
- FindLast
- Includes
- IndexOf
- LastIndexOf