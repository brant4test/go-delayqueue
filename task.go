package godelayqueue

type Task struct {
	Id string
	// the number of cycles of the task on the time wheel,
	// when it is equal to 0, the task is executed
	CycleCount int
	// the position of the task on the time wheel
	WheelPosition int
	// the task type,
	// which is used by the factory method to determine which implementation object to use
	TaskType string
	// task method parameters
	TaskParams string

	Next *Task
}
