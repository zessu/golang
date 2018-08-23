package main

import "fmt"

func main() {
	s := make([]int, 3)
	s[0] = 1
	s[0] = 2
	// s[0] = "asdad" this would error
	fmt.Println(s[0])
	// we can also define this as
	z := []int{1, 2, 3, 4} // size if inferred
	// first index is where to start, second is where to stop
	fmt.Println(z[0:1]) // does not change slice
	fmt.Println(z[:2])  // start at 0 and end at index 2
	fmt.Println(z[3:])  // start at 3 and get everything to the end
	fmt.Println(z)
	//appening to slice
	z = append(z, 500)
	fmt.Println(z)
	// deleting from slice, lets delete 3 from [1,2,3,4,500]
	z = append(z[:2], z[3:]...)
	fmt.Println(z)
}
