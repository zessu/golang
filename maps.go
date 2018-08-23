package main

import (
	"fmt"
)

func main() {
	grades := make(map[string]int32)
	grades["andrew"] = 12
	grades["jane"] = 45
	grades["michael"] = 99
	andrewsGrade := grades["andrew"]
	fmt.Println(andrewsGrade)
	// delete something from map
	delete(grades, "andrew")
	fmt.Println(grades["andrew"])
	// range it
	for name, value := range grades {
		fmt.Println(name, value)
	}
	// we can also define maps faster like this
	otherGrades := map[string]int32{
		"andrew":  12,
		"jane":    45,
		"michael": 999,
	}
	fmt.Println(otherGrades["michael"])
	// update map
	otherGrades["sylvia"] = 70
	fmt.Println(otherGrades["sylvia"])

	// checking for existance in a really cool way
	if _, exists := grades["jane"]; exists {
		fmt.Println("value exists")
	}

	fmt.Println(grades["whatsndadsadasd"]) // returns 0, as second arg
}
