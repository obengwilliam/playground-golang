package employee

import (
	"fmt"
)

// Employee doing some wonders
type employee struct {
	firstName   string
	lastName    string
	totalLeaves int
	leavesTaken int
}

// NewEmployee doing awesome work
//constructor
func New(firstName string, lastName string, totalLeaves int, leavesTaken int) employee {
	e := employee{firstName, lastName, totalLeaves, leavesTaken}

	return e
}

//static methods
func (e employee) LeavesRemaining() {
	fmt.Printf("%s %s has %d leaves remaining", e.firstName, e.lastName, e.totalLeaves-e.leavesTaken)

}
