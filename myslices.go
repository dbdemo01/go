package main

import "fmt"

func main() {
	s := make([]string, 3)
	fmt.Println("emp: ", s)

	s[0] = "this"
	s[1] = "is"
	s[2] = "my"
	fmt.Println("set: ", s)
	fmt.Println("get: ", s[2])
	fmt.Println("len: ", len(s))

	s = append(s, "slice1")
	s = append(s, "slice2", "slice3")
	fmt.Println("appd: ", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy: ", c)

	l := s[2:5]
	fmt.Println("sl1: ", l)

	l = s[:5]
	fmt.Println("sl2: ", l)

	l = s[2:]
	fmt.Println("sl3: ", l)

	t := []string{"slice4", "slice5", "slice6"}
	fmt.Println("dcl: ", t)

	twoDimension := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoDimension[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoDimension[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoDimension)
}

$ go run myslices.go 
emp:  [  ]
set:  [this is my]
get:  my
len:  3
appd:  [this is my slice1 slice2 slice3]
cpy:  [this is my slice1 slice2 slice3]
sl1:  [my slice1 slice2]
sl2:  [this is my slice1 slice2]
sl3:  [my slice1 slice2 slice3]
dcl:  [slice4 slice5 slice6]
2d:  [[0] [1 2] [2 3 4]]

$ go build myslices.go
$ myslices 
emp:  [  ]
set:  [this is my]
get:  my
len:  3
appd:  [this is my slice1 slice2 slice3]
cpy:  [this is my slice1 slice2 slice3]
sl1:  [my slice1 slice2]
sl2:  [this is my slice1 slice2]
sl3:  [my slice1 slice2 slice3]
dcl:  [slice4 slice5 slice6]
2d:  [[0] [1 2] [2 3 4]]

