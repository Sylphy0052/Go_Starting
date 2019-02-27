package main

import "fmt"

// func main() {
// 	var i int
// 	p := &i
// 	fmt.Printf("%d\n", p)
// 	fmt.Println(p)
// 	fmt.Println(*p)
// 	fmt.Println(&p)
// 	pp := &p
// 	fmt.Printf("%d\n", pp)
// }

func main() {
	var i int
	p := &i
	i = 5
	fmt.Println(*p)
	*p = 10
	fmt.Println(i)
}
