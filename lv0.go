package main

import (
	"fmt"
	"time"
)

func main(){
	go spinner(20*time.Millisecond)
	const n = 66
	go jiechen(n)
    fibN:=fib(n)
	fmt.Println(n,fibN)

}

func spinner(delay time.Duration){
	for{
		for _,r :=range `~-\|/`{
			fmt.Printf("\r%c",r)
			time.Sleep(delay)
		}
	}
}

func fib(x int64) int64{
	if x<2{
		return x
	}
	return fib(x-1)+fib(x-2)
}

func jiechen(x int) {
	for i:=1;i<x;i++ {
		x*=i
	}
	fmt.Println(x)
}