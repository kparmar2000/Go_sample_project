package math

import (
	"testing"
)

func TestOlder(t *testing.T) {
	cases := []struct {
		name     string
		age1     int
		age2     int
		expected bool
	}{
		{
			name:     "FirstOlderThanSecond",
			age1:     1,
			age2:     2,
			expected: false,
		},
		{
			name:     "SecondOlderThanFirst",
			age1:     2,
			age2:     1,
			expected: true,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			p1 := NewPerson(c.age1)
			p2 := NewPerson(c.age2)

			got := p1.older(p2)

			if (got.Age > p2.Age) != c.expected {
				t.Errorf("Expected %v > %v to be %v, but got %v", p1.Age, p2.Age, c.expected, got.Age)
			}
		})
	}
}
