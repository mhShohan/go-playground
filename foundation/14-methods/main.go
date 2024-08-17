package main

import "fmt"

type User struct {
	Name    string
	Age     int
	isAdmin bool
}

func (user User) GetInfo() string {
	res := fmt.Sprintf("Name: %v, Age: %v, isAdmin: %v", user.Name, user.Age, user.isAdmin)
	return res
}

func (user *User) SetName(name string) {
	user.Name = name
}

func main() {
	fmt.Println("--------Methods---------")
	user := User{"mhShohan", 25, true}

	// getter
	result := user.GetInfo()
	fmt.Println(result)

	// setter
	user.SetName("MH Shohan")

	result = user.GetInfo()
	fmt.Println(result)
}
