package simplemath 

import (
	"math"
)

/*
加法运算操作
*/
func Addx(n1 int,n2 int) (returnI int){
	return n1+n2;
}

/*
开根号运算
*/
func Sqrtx(n1 int)(returnf float64){
	return math.Sqrt(float64(n1));
}

