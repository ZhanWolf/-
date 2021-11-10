package main

import (
	"fmt"
	"time"
)

var ch1 chan int

func main() {
	go fuhu()
    go twooclock()
	<-ch1
}


func fuhu(){

	ticker := time.NewTicker(time.Second*30)

	for {
		 <-ticker.C
		 fmt.Println("芜湖！起飞！")
	}
}

func twooclock(){
	for {
		now := time.Now()
		if now.Equal(time.Date(now.Year(), now.Month(), now.Day(), 02, 00, 00,now.Nanosecond(), now.Location())){
         fmt.Println("谁能比我卷！")
		 time.Sleep(time.Second*3)
		}
		if now.Equal(time.Date(now.Year(), now.Month(), now.Day(), 06, 00, 00,now.Nanosecond(), now.Location())){
			fmt.Println("早八算什么，早六才是吾辈应起之时")
			time.Sleep(time.Second*3)
		}
	}
}