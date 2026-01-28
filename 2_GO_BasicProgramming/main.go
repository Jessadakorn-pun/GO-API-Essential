package main

import (
	"errors"
	"fmt"
)

type Student struct {
	Name   string
	Height int
	Weight int
	Grade  float32
}

type Address struct {
	Street  string
	City    string
	Zipcode int
}

type Student2 struct {
	FirstName string
	LastName  string
}

// Speaker Interface
type Speaker interface {
	Speak() string
}

type Dog struct {
	Name string
}

func (d Dog) Speak() string {
	return "Woof"
}

type Person struct {
	Name    string
	Address Address
}

func (p Person) Speak() string {
	return "Hello"
}

// function that accept Speaker interface
func makeSound(s Speaker) {
	fmt.Println(s.Speak())
}

type Shape interface {
	area() float32
}

type triangle struct {
	base float32
	high float32
}

func (t triangle) area() float32 {
	return 0.5 * t.base * t.high
}

type regtangle struct {
	length float32
}

func (r regtangle) area() float32 {
	return r.length * r.length
}

func calulateArea(s Shape) float32 {
	return s.area()
}

// pointer with struct
type Employee struct {
	Name   string
	Salary int
}

// custom error
type LoginError struct {
	Username string
	Message  string
}

func (e *LoginError) Error() string {
	return fmt.Sprintf("Login Error For User '%s': %s", e.Username, e.Message)
}

func login(username, password string) error {
	if username != "admin" || password != "admin" {
		return &LoginError{Username: username, Message: "invalid credentials"}
	}
	return nil
}

func main() {
	// define variable
	var fullName string
	fullName = "mikelopster"

	var age int = 12
	major := "CS"

	fmt.Printf("Hello %s, Age: %d, Major: %s \n", fullName, age, major)

	// operator : + - * /
	a := 10
	b := 3
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)

	fullName += "Channel"
	fmt.Println(fullName)

	fmt.Println(a == b)
	fmt.Println(a != b)
	fmt.Println(a < b)
	fmt.Println(a > b)
	fmt.Println(a <= b)
	fmt.Println(a >= b)

	// and or not
	c := true
	d := false
	fmt.Println(c && d)
	fmt.Println(c || d)
	fmt.Println(!c)

	// constant variable
	const constant = "ABC"
	fmt.Println(constant)

	// control structure
	// if-else
	score := 70
	if score >= 70 {
		fmt.Printf("Pass")
	} else {
		fmt.Printf("Fail")
	}

	if score >= 80 {
		fmt.Printf("A")
	} else if score >= 70 {
		fmt.Printf("B")
	} else {
		fmt.Printf("Fail")
	}

	// switch case
	dayOfWeek := 6
	switch dayOfWeek {
	case 6, 7:
		fmt.Println("WeekEnd")
	default:
		fmt.Println("WeekDay")
	}

	// if short form
	num1 := 5
	num2 := 10

	if sumNum := num1 + num2; sumNum >= 10 {
		fmt.Println("sumNum more than 10")
	}

	// Loop
	// For Loop
	for i := 1; i <= 10; i++ {
		if i > 5 {
			break
		}
		fmt.Println(i)
	}

	// do while loop
	i := 1
	for {
		fmt.Println(i)
		i++
		if i >= 10 {
			break
		}
	}

	// while loop
	cnt := 1
	for cnt < 10 {
		fmt.Println(cnt)
		i++
	}

	// DataStructure
	// Array
	numArray := [3]int{1, 2, 3}

	numArray[0] = 2 // re-assign value

	fmt.Println(numArray[0], numArray[1], numArray[2], numArray)

	for _, value := range numArray {
		fmt.Println(value)
	}

	for i := 0; i < len(numArray); i++ {
		fmt.Println(numArray[i])
	}

	// Slice
	mySlice := []int{1, 2, 3, 4, 5}
	fmt.Println(mySlice)
	fmt.Println(len(mySlice)) // length of slice
	fmt.Println(cap(mySlice)) // capacity of slice

	subSlice := mySlice[:2] // slicing
	fmt.Println(subSlice)

	mySlice = append(mySlice, 10, 20, 30) // append slice
	mySlice2 := numArray[:]               // convert array to slice
	fmt.Println(mySlice2)

	// Map
	myMap := make(map[string]int) // make = reserve memory for map

	myMap["apple"] = 5 // create key value pair
	myMap["orange"] = 10
	myMap["banana"] = 8

	for key, value := range myMap {
		fmt.Printf("%s -> %d /n", key, value) // loop through map
	}

	delete(myMap, "banana") // remove key

	val, ok := myMap["pear"] // checking of a key exist
	if ok {
		fmt.Println("Pear value :", val)
	} else {
		fmt.Println("Pear not found")
	}

	// Struct
	var student1 Student
	student1.Name = "Mikelopster"
	student1.Weight = 60
	student1.Height = 180
	student1.Grade = 3.14

	fmt.Println(student1, student1.Name, student1.Weight, student1.Height, student1.Grade)

	var student [2]Student // combination of struct and array
	student[0].Name = "Mikelopster"
	student[0].Weight = 60
	student[0].Height = 180
	student[0].Grade = 3.14

	student[1].Name = "Mikelopster"
	student[1].Weight = 60
	student[1].Height = 180
	student[1].Grade = 3.14

	for i := 0; i < len(student); i++ {
		fmt.Println("Student", (i + 1))
		fmt.Println(student[i].Name)
		fmt.Println(student[i].Weight)
		fmt.Println(student[i].Height)
		fmt.Println(student[i].Grade)
	}

	students := make(map[string]Student)

	students["st01"] = Student{Name: "Mikelopster", Weight: 60, Height: 180, Grade: 3.14}

	bob := Person{
		Name: "Bob",
		Address: Address{
			Street:  "ABC",
			City:    "DEF",
			Zipcode: 123,
		},
	}

	fmt.Println(bob)

	// function
	sayHello("pun", 10)
	add(a, b)

	// Method
	stu := Student2{
		FirstName: "Jessadakorn",
		LastName:  "Pun",
	}

	fullName2 := stu.FullName()
	fmt.Printf("Full Name of Student: %s", fullName2)

	// interface
	dog := Dog{Name: "ABC"}
	person := Person{Name: "DEF"}

	makeSound(dog)
	makeSound(person)

	regtangle := regtangle{
		length: 32,
	}

	fmt.Println(calulateArea(regtangle))

	// pointer
	x := 10
	var pt *int = &x
	fmt.Printf("Value of x: %d \n", x)
	fmt.Printf("Value at p: %d \n", *pt)

	*pt = 100 // modify value at address x

	// pass by refernec
	emp := Employee{
		Name:   "Pun",
		Salary: 300,
	}

	giveRaise(&emp, 3000)
	fmt.Println(emp)

	// Handling Error
	num3 := 5
	num4 := 0

	val, err := divied(num3, num4)
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println("Result: ", val)
	}

	err2 := login("user", "pass")
	if err2 != nil {
		switch e := err2.(type) {
		case *LoginError:
			fmt.Println("Custom Error Occured: ", e)
		default:
			fmt.Println("Generic Error Occured: ", e)
		}
	}

}

// function
func sayHello(name string, round int) {
	for i := 0; i < round; i++ {
		fmt.Printf("Hello %s \n", name)
	}

}

func add(a int, b int) int {
	return a + b
}

// Method with  a receuver of type student
// This method return the full name of the student
func (s Student2) FullName() string {
	return s.FirstName + " " + s.LastName
}

func giveRaise(e *Employee, raise int) {
	e.Salary += raise
}

// Handling Error
func divied(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("cannot divied by zero")
	}

	return a / b, nil
}
