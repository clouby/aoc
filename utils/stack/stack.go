package stack

type Stack struct {
	top    *node
	length int
}

type node struct {
	value interface{}
	prev  *node
}

// Create a new stack
func New() *Stack {
	return &Stack{top: nil, length: 0}
}

// Return the number of iterms in the stack
func (stack *Stack) Len() int {
	return stack.length
}

// View the top item on the stack
func (stack *Stack) Peek() interface{} {
	if stack.length == 0 {
		return nil
	}

	return stack.top.value
}

// Pop the top item of the stack and return it
func (stack *Stack) Pop() interface{} {
	if stack.length == 0 {
		return nil
	}

	n := stack.top
	stack.top = n.prev
	stack.length -= 1

	return n.value
}

// Push a value onto the top of the stack
func (stack *Stack) Push(value interface{}) {
	n := &node{value: value, prev: stack.top}
	stack.top = n
	stack.length += 1
}

// Pop multiples values with a range provided
func (stack *Stack) Pick(latest int) []interface{} {
    latestValues := make([]interface{}, 0, latest)

    for len(latestValues) < latest {
        latestValues = append(latestValues, stack.Pop())
    }

    return latestValues
}

// Bulk of multiples elements to be pushed
func (stack *Stack) Bulk(values ...interface{}) {
    for _, value := range values {
        stack.Push(value)
    }
}
