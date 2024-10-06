package main

import (
	"testing"
)

func TestStack_Push(t *testing.T) {
	stack := newStack("google")
	stack.push("yahoo")
	stack.push("udemy")

	if stack.Top.Data != "udemy" {
		t.Errorf("Expected top of stack to be 'udemy', but got %s", stack.Top.Data)
	}

	if stack.Length != 3 {
		t.Errorf("Expected stack length to be 3, but got %d", stack.Length)
	}
}

func TestStack_Pop(t *testing.T) {
	stack := newStack("google")
	stack.push("yahoo")
	stack.push("udemy")

	stack.pop()

	if stack.Top.Data != "yahoo" {
		t.Errorf("Expected top of stack to be 'yahoo', but got %s", stack.Top.Data)
	}

	if stack.Length != 2 {
		t.Errorf("Expected stack length to be 2, but got %d", stack.Length)
	}

	// Test pop on empty stack
	stack.pop()
	stack.pop()
	stack.pop() // trying to pop when stack is empty

	if stack.Length != 0 {
		t.Errorf("Expected stack length to be 0 after popping all elements, but got %d", stack.Length)
	}

	if stack.Top != nil {
		t.Error("Expected top of stack to be nil after popping all elements")
	}
}

func TestStack_Peek(t *testing.T) {
	stack := newStack("google")

	if stack.peek() != "google" {
		t.Errorf("Expected peek to return 'google', but got %s", stack.peek())
	}

	stack.push("yahoo")
	if stack.peek() != "yahoo" {
		t.Errorf("Expected peek to return 'yahoo', but got %s", stack.peek())
	}

	// Test peek on empty stack
	stack.pop()
	stack.pop()
	if stack.peek() != "Stack is empty" {
		t.Errorf("Expected peek on empty stack to return 'Stack is empty', but got %s", stack.peek())
	}
}

func TestStack_Show(t *testing.T) {
	stack := newStack("google")
	stack.push("yahoo")
	stack.push("udemy")

	expectedOrder := []string{"udemy", "yahoo", "google"}
	tmpTopNode := stack.Top

	for i := 0; i < len(expectedOrder); i++ {
		if tmpTopNode == nil || tmpTopNode.Data != expectedOrder[i] {
			t.Errorf("Expected %s at position %d, but got %s", expectedOrder[i], i, tmpTopNode.Data)
		}
		tmpTopNode = tmpTopNode.Next
	}

	if stack.Length != 3 {
		t.Errorf("Expected stack length to be 3, but got %d", stack.Length)
	}
}
