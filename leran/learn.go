package leran

import "fmt"

func Run() {
	array1 := [3]int{1, 2, 3}
	array2 := [...]int{1, 2, 3}
	array3 := [...]int{1: 3, 2: 2}

	slice1 := []int{1, 2, 3}
	map1 := map[int]string{1: "Hello", 2: "world"}

	fmt.Printf("%v+", array1)
	fmt.Printf("%v+", array2)
	fmt.Printf("%v+", array3)
	fmt.Printf("%v+", slice1)
	fmt.Printf("%v+", map1)

}
