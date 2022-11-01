package main

import (
	"fmt"

)

func main() {

	// s := []int{1, 2, 3, 4, 5}
	// fmt.Println(s)

	
	// // i :=  0
	// // fmt.Println(s[i])
	// // s = append(s[:i], s[i+1:]...)

	// s = SliceRemoveByIndex(s,2)
	// fmt.Println(s)

	// s2 := []string{"1","2","3","4","5"}
	// s2 = SliceRemoveByIndex(s2,3)
	// fmt.Println(s2)

	// m := make(map[string]int)
	// m["0"] = 0
	// m["1"] = 1
	// fmt.Println(m)


	// for i, v := range s {
	// 	fmt.Printf("%d %d\n", i, v)
	// 	if v == 2 {
	// 		s = SliceRemoveByIndex(s,i)
	// 	}
	
	// }
	// fmt.Println(s)
	
	// sum := 0
	// for i := 0; i < 10; i++ {
	// 	sum += i
	// }
	// fmt.Println(sum)

	s := []int{1, 2, 3, 4, 5, 6}

	c := make(chan int)
	fmt.Println(s[:len(s)/2])
	fmt.Println(s[len(s)/2:])

	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // receive from c

	fmt.Println(x, y, x+y)

	

}

// func SliceRemoveByIndex(slice []int, index int) []int {
// 	return append(slice[:index], slice[index+1:]...)
// }


func SliceRemoveByIndex[Type int|string](slice []Type, index int) []Type{
    return append(slice[:index], slice[index+1:]...)
}

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

