package safe

import "testing"

func TestMustAtOutOfRange(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatal("got nil instead of panic")
		} else if r != "index out of range" {
			t.Errorf("got %q instead of 'index out of range'", r)
		}
	}()

	_ = MustAt([]int{1, 2, 3}, 3)
}

func TestMustAtLessThanZero(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatal("got nil instead of panic")
		} else if r != "index out of range" {
			t.Errorf("got %q instead of 'index out of range'", r)
		}
	}()

	_ = MustAt([]int{1, 2, 3}, -3)
}

func TestMustAtInRange(t *testing.T) {
	got := MustAt([]int{1, 2, 3}, 1)
	want := 2

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
