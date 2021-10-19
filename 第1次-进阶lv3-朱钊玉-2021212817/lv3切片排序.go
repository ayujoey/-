package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main(){
		var num [] int
		t:=int64(time.Now().Nanosecond())
		rand.Seed(t)

		for i:=0; i < 100; i++{
			num = append(num,rand.Int())
		}
		px(num)

		for _,j:= range num{
			fmt.Println(j)
		}
}

func px(num []int){

	for i:=0;i<100;i++ {
		for j := i; j < 100; j++ {
			if num[i] > num[j] {
				num[i], num[j] = num[j], num[i]
			}
		}
	}
}