package queue

import (
	"testing"
)

func Test_all(t *testing.T) {

	v0 := NewQueue()
	v0.Push(1)
	v0.Push("dlrjehlsk?")
	v0.Push(3.2)

	value, err := v0.Front()

	if err != nil {
		t.Fatal("expect Not Empty, but got Empty")
	}
	if value != 1 {
		t.Fatal("Front is Error.")

	}

	value, err = v0.Back()

	if err != nil {
		t.Fatal("expect Not Empty, but got Empty")
	}
	if value != 3.2 {
		t.Fatal("Back is Error.")
	}

	err = v0.Pop()

	if err != nil {
		t.Fatal("expect Not Empty, but got Empty")
	}

	value, err = v0.Front()

	if err != nil {
		t.Fatal("expect Not Empty, but got Empty")
	}
	if value != "dlrjehlsk?" {
		t.Fatal("Front is Error.")
	}

	value, err = v0.Back()

	if err != nil {
		t.Fatal("expect Not Empty, but got Empty")
	}
	if value != 3.2 {
		t.Fatal("Back is Error.")
	}

	err = v0.Pop()

	if err != nil {
		t.Fatal("expect Not Empty, but got Empty")
	}

	value, err = v0.Front()

	if err != nil {
		t.Fatal("expect Not Empty, but got Empty")
	}
	if value != 3.2 {
		t.Fatal("Front is Error.")
	}

	value, err = v0.Back()

	if err != nil {
		t.Fatal("expect Not Empty, but got Empty")
	}
	if value != 3.2 {
		t.Fatal("Back is Error.")
	}

	err = v0.Pop()

	if err != nil {
		t.Fatal("expect Not Empty, but got Empty")
	}

}
