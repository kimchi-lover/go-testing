package validate

import "testing"

func TestValidateName(t *testing.T) {
	if got := ValidateName("Hexlet"); got != nil {
		t.Errorf("got %v instead of nil", got)
	}
}

func TestValidateName_Empty(t *testing.T) {
	err := ValidateName("")

	if err == nil {
		t.Fatal("got nil instead of error")
	}

	want := ErrEmptyName
	if err.Error() != want.Error() {
		t.Errorf("got %q want %q", err.Error(), want.Error())
	}
}
