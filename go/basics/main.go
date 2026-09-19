// Go ke basics: variables, loops, slices, maps, defer.
package main

import "fmt"

func add(a, b int) int { return a + b }

// Go me multiple return values normal baat hai
func divide(a, b int) (int, bool) {
	if b == 0 {
		return 0, false
	}
	return a / b, true
}

func main() {
	// := short declaration hai, type Go khud samajh leta hai
	name := "tanishk"
	var age int = 25
	fmt.Println(name, age, add(2, 3))

	if q, ok := divide(10, 2); ok {
		fmt.Println("quotient:", q)
	}

	// Go me sirf ek hi loop hai: for. while bhi for hi hai
	for i := 0; i < 3; i++ {
		fmt.Println("i =", i)
	}

	// slice = dynamic array. append naya slice return karta hai, isliye reassign karo
	nums := []int{1, 2, 3}
	nums = append(nums, 4)
	for idx, v := range nums {
		fmt.Println(idx, v)
	}

	// map = dictionary. missing key pe zero value milti hai, error nahi
	ages := map[string]int{"tanishk": 25, "shreya": 24}
	if v, ok := ages["omi"]; !ok {
		fmt.Println("omi nahi mila, zero value:", v)
	}

	// defer function return hone ke baad chalta hai (LIFO order me), cleanup ke liye best
	defer fmt.Println("defer 1")
	defer fmt.Println("defer 2") // ye pehle print hoga
	fmt.Println("main khatam")
}
