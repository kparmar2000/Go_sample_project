package math

import "testing"

func TestAdd(t *testing.T) {
	if Add(2, 3) != 5 {
		t.Errorf("Error: Got ,%q Want ,%q ", Add(2, 3), 5)
	}
}
