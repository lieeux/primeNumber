package main

import (
	"fmt"
)

func putNum(intChan chan int) {
	for i := 2; i <= 1000; i++ {
		intChan <- i
	}
	close(intChan)
}

func primeNum(intChan chan int, primeChan chan int, exitChan chan bool) {
	var flag bool
	for {
		num, ok := <-intChan
		flag = true //假设是素数
		if !ok {    //取不出数
			break
		}
		for i := 2; i < num; i++ {
			if num%i == 0 { //不是素数
				flag = false
				break
			}
		}
		if flag { //是素数则放入管道
			primeChan <- num
		}
	}
	exitChan <- true //取不出数
	fmt.Println("取不出数，协程关闭")
}

func main() {
	intChan := make(chan int, 100)
	primeChan := make(chan int, 500)
	exitChan := make(chan bool, 4)

	go putNum(intChan)
	for i := 0; i < 4; i++ {
		go primeNum(intChan, primeChan, exitChan)
	}

	go func() {
		for i := 0; i < 4; i++ {
			<-exitChan
		}
		close(primeChan)
	}() //匿名函数

	primeNum := 0
	for {
		res, ok := <-primeChan
		if !ok {
			break
		}
		fmt.Println(res)
		primeNum++
	}
	fmt.Println("素数个数为", primeNum)
}
