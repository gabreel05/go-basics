package main

import (
	"fmt"
	"runtime"
)

func add(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main2() {
	k := 10
	fmt.Println(k)
	var c, python, java bool
	fmt.Println(c, python, java)
	var one, two, three = 1, 2, "Three"

	fmt.Println(one, two, three)

	a, b := swap("Hello", "World")
	fmt.Println(split(17))
	fmt.Println(add(15, 10))
	fmt.Println(a, b)
	fmt.Println("Hello, World!")

	sum := 0
	for i := 0; i <= 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s.\n", os)
	}
	var p *int

	i := 10
	p = &i
	*p = 20
	fmt.Println(*p)
	fmt.Println(i)

	fmt.Println(Vertex{1, 2})
	vertex := Vertex{10, 20}
	fmt.Println(vertex.X)
	fmt.Println(vertex.Y)

	vertex2 := Vertex{15, 30}
	pVertex := &vertex2
	fmt.Println(pVertex)

	vertex3 := Vertex{X: 1, Y: 2}
	fmt.Println(vertex3)

	var arrMatrix [2]string
	arrMatrix[0] = "Hello"
	arrMatrix[1] = "World"
	fmt.Println(arrMatrix[0], arrMatrix[1])
	fmt.Println(arrMatrix)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	var s []int = primes[0:4]
	fmt.Println(s)

	sliceWithMake := make([]int, 5)
	fmt.Println(sliceWithMake)

	sliceWithMake = append(sliceWithMake, 10)
	fmt.Println(sliceWithMake)

	pow := []int{1, 2, 4, 8, 16, 32, 64}

	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	newMap := make(map[string]Vertex)
	newMap["Bell Labs"] = Vertex{
		40, -74,
	}
	fmt.Println(newMap["Bell Labs"])
}

type Vertex struct {
	X int
	Y int
}
