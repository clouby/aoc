package stack

import "testing"

var (
	want interface{}
	got  interface{}
)

func TestStack(t *testing.T) {
	stack := New()

	want = 0
	if stack.Len() != want {
		t.Errorf("got %v, want %v", stack.Len(), want)
	}

	stack.Push("Doe")

	want = 1
	if stack.Len() != want {
		t.Errorf("got %v, want %v", stack.Len(), want)
	}

	stack.Push("Foo")

	got = stack.Peek()
	want = "Foo"
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}

	stack.Pick(2)

	got = stack.Len()
	want = 0
	if got != want.(int) {
		t.Errorf("got %v, want %v", got, want)
	}
}
