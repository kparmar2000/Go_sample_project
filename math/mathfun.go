package math

func Add(a int, b int) int {
	return a + b
}
func Sub(a int, b int) int {
	return a - b
}
func Mult(a int, b int) int {
	return a * b
}

/*
	 type Person struct {
		Name string
		Age  int
	}

// older method compares the ages of two Person instances and returns the older person.

	func (p Person) older(other Person) Person {
		if p.Age > other.Age {
			return p
		}
		return other
	}

	func NewPerson(age int) Person {
		return Person{
			Name: "", // You can set the name to an empty string or any default value you prefer
			Age:  age,
		}
	}
*/
type Person struct {
	Name string
	Age  int
}

// NewPerson creates and returns a new Person instance with the given age.
func NewPerson(age int) Person {
	return Person{
		Name: "", // You can set the name to an empty string or any default value you prefer
		Age:  age,
	}
}

// older method compares the ages of two Person instances and returns true if the first person is older than the second person.
func (p Person) older(other Person) Person {
	//return p.Age > other.Age
	if p.Age > other.Age {
		return p
	}
	return other
}
