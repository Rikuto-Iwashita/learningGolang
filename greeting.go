package main

import "fmt"

type User struct {
	gender string
	age    int
}

func main() {
	// var a User
	// a.gender = "male"
	// a.age = 20

	b := User{gender: "male", age: 20}

	fmt.Println(b)
}
