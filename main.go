package main

import (
	"fmt"
)

var (
	a    string
	b    string
	c    string
	d    string
	str1 string
)

func main() {
	a = "a"
	b = "b"
	c = "c"
	d = "d"
	str1 = "牛逼"
	a, b = b, a
	pointer := &b
	*pointer = c
	pointer2 := new(a)
	*pointer2 = d
	fmt.Println(pointer, b, c, d, pointer2, len(str1))}
