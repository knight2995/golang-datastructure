package stack

import (
	"testing"
)

func Test_all(t *testing.T) {

	v0 := NewStack()
	v0.Push(1)
	v0.Push("dlrjehlsk?")
	v0.Push(3.2)

	value, err := v0.Top()

	if err != nil {
		t.Fatal("expect Not Empty, but got Empty")
	}
	if value != 3.2 {
		t.Fatal("Top is Error.")
	}

	err = v0.Pop()

	if err != nil {
		t.Fatal("expect Not Empty, but got Empty")
	}

	value, err = v0.Top() //expect "dlrjehlsk?"

	if err != nil {
		t.Fatal("expect Not Empty, but got Empty")
	}
	if value != "dlrjehlsk?" {
		t.Fatal("Pop is Error.")
	}

	err = v0.Pop() //expect "1"

	if err != nil {
		t.Fatal("expect Not Empty, but got Empty")
	}

	value, err = v0.Top() //expect "dlrjehlsk?"

	if err != nil {
		t.Fatal("expect Not Empty, but got Empty")
	}
	if value != 1 {
		t.Fatal("Top is Error.")
	}

	err = v0.Pop()

	if err != nil {
		t.Fatal("expect Not Empty, but got Empty")
	}

}
