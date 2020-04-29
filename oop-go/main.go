package main

import (
	"fmt"
	"github.com/obengwilliam/oop-go/blog"
	"github.com/obengwilliam/oop-go/employee"
	"github.com/obengwilliam/oop-go/income"
)

func main() {
	fmt.Println("hello")
	e := employee.New("william", "obeng", 20, 5)
	e.LeavesRemaining()

	// blog
	author1 := blog.Author{FirstName: "william", LastName: "obeng", Bio: "fuck you"}

	post1 := blog.Post{
		Author:  author1,
		Title:   "Inheritance in Go",
		Content: "Go supports composition instead of inheritance",
	}

	post1.Details()
	fmt.Println(post1.Author.FullName())

	// PolyMorphism
	project1 := income.FixedBilling{ProjectName: "Project 1", BiddedAmount: 5000}
	project2 := income.FixedBilling{ProjectName: "Project 2", BiddedAmount: 10000}
	project3 := income.TimeAndMaterial{ProjectName: "Project 3", NoOfHours: 160, HourlyRate: 25}
	incomeStreams := []income.Income{project1, project2, project3}
	income.CalculateNetIncome(incomeStreams)
}
