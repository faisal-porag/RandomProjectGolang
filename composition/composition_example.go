package composition

import "fmt"

type author struct {
	firstName string
	lastName  string
	bio       string
}

func (a author) fullName() string {
	return fmt.Sprintf("%s %s", a.firstName, a.lastName)
}

func (a author) getBio() string {
	return fmt.Sprintf("%s. Thank u for visiting.", a.bio)
}

type blogPost struct {
	title   string
	content string
	author
}

func (b blogPost) details() {
	fmt.Println("Title: ", b.title)
	fmt.Println("Content: ", b.content)
	fmt.Println("Author: ", b.fullName())
	fmt.Println("Bio: ", b.getBio())
}

func CompositionFunc() {
	author1 := author{
		"Faisal",
		"Porag",
		"Golang Learner",
	}
	blogPost1 := blogPost{
		"Inheritance in Go",
		"Go supports composition instead of inheritance",
		author1,
	}
	blogPost1.details()
}
