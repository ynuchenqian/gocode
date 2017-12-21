package main

import (
	"math"
	"math/cmplx"
	"fmt"
)
var c,python,java bool
var i,j int =1,2
var (
	ToBe bool = false
	MaxInt uint=1<<64-1
	z complex128=cmplx.Sqrt(-5+12i)
)
func main() {
	fmt.Printf("你的年龄是 %g岁",math.Sqrt(100))
	fmt.Println(math.Pi)
	fmt.Println(add(23,33))
	fmt.Println(divide(20,10))
	a,b:=swap("world","hello")
	fmt.Println(a,b)
	//分割调用
	fmt.Println(split(17))

	var c,python,java=true,false,"no"
	fmt.Println(i,j,c,python,java)
	const f ="%T(%v)\n"
	fmt.Printf(f, ToBe, ToBe)
	fmt.Printf(f, MaxInt, MaxInt)
	fmt.Printf(f, z, z)
}

func add(x int,y int)  int{
	return x+y
}

func divide(x, y int) int {
	return  x-y
}
/**
  函数位置交换
 */
func swap(x,y string) (string,string) {
	return y,x
}

func split(sum int) (x, y int) {
	x=sum*4/9
	y=sum-x
	return
}


