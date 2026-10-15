package main

import "fmt"

type Student struct {
	name string
	age  int
	sem  string
}

var studentData []Student

// Adding Student
func Addstudent() {

	var name, sem string
	var age int
	for {
		fmt.Print("Enter you name: ")
		_, err := fmt.Scanln(&name)
		if err == nil && name != "" {
			break
		}
		fmt.Println("Please enter a your name")
	}
	for {
		fmt.Print("\nEnter you age: ")
		_, err := fmt.Scanln(&age)

		if err == nil && age > 0 {
			break
		}
		fmt.Println("Please enter a valid age")
	}
	for {
		fmt.Print("Enter you sem: ")
		_, err := fmt.Scanln(&sem)
		if err == nil && sem != "" {
			break
		}
		fmt.Println("please enter a valid semester")
	}
	student := Student{
		name: name,
		age:  age,
		sem:  sem,
	}

	studentData = append(studentData, student)
	fmt.Println("Student Added  successfully")

}

//Delete student

func DeleteStudent() {
	var name string
	fmt.Print("Enter you name: ")
	fmt.Scanln(&name)

	for index, student := range studentData {
		if student.name == name {
			studentData = append(studentData[:index], studentData[index+1:]...)
			fmt.Println("student deleted successfully")
			return
		}
	}

	fmt.Println("Student  not found\n\n ")
}

// Show all student
func ShowStudent() {
	if len(studentData) == 0 {
		fmt.Println("No student Found..")
		return
	}
	for index, value := range studentData {
		fmt.Println("All student")
		fmt.Printf("The data  %d student\n", index+1)
		fmt.Println("name", value.name)
		fmt.Println("age", value.age)
		fmt.Println("sem", value.sem)

	}
}

func main() {
	fmt.Println("Student Management System...")

	for true {
		fmt.Println("\n\nEnter 1 for Add Student\n Enter 2 for Delete student\n Enter 3 for show all student\n Enter 4 for exit")
		fmt.Println("Enter your choice")
		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			Addstudent()

		case 2:
			DeleteStudent()

		case 3:
			ShowStudent()

		case 4:
			fmt.Println("Exiting program")
			return

		default:
			fmt.Println("Invalid choices..")

		}
	}
}
