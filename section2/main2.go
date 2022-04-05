package main

import(
	"fmt"
)

//arrays
// func todo() {
// 	//var arr []int
// 	arr := []int{1, 2, 3, 4}
// 	arr2 := []string{"hi", "my", "name"}

// 	arr2 = append(arr2, "is", "angad", "!")
// 	fmt.Println(arr, arr2)
// }

//structure encapsulation (encapsulacion de datos)

type Car struct {
	Name    string `json: "name"`
	Age     int
	ModelNo int
}

func (c Car) Print() {
	fmt.Println(c)
}

func (c Car) Drive() {
	fmt.Println("driiving..")
}

func (c Car) GetAge() int {
	return c.Age
}

func main() {
	//todo()
	c := Car{
		Name:    "chevy",
		Age:     5,
		ModelNo: 2,
	}
	c.Print()
	c.Drive()
	//fmt.Println(c)
	fmt.Println(c.GetAge())
}

