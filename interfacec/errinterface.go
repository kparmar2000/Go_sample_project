package main

import "fmt"

type usererr struct {
	name string
}

func (e usererr) Error() string {
	return fmt.Sprintf("%v user is facing problem\n", e.name)
}
