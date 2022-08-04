package main

import (
	"fmt"
	"os"
	"time"
)

func Launch() {
	fmt.Println("The Rocket is launched")
}

func Abort(abort chan struct{}) {
	os.Stdin.Read(make([]byte, 1))
	abort <- struct{}{}
}

func RocketExample() {
	abort := make(chan struct{})
	fmt.Println("Commencing countdown. Press return to abort.")
	go Abort(abort)
	select {
	case <-time.After(10 * time.Second):
		fmt.Println("Starting Launch")
	case <-abort:
		return
	}
	Launch()
}

func RocketTickerExample() {
	abort := make(chan struct{})
	fmt.Println("Commencing countdown. Press return to abort.")
	go Abort(abort)
	ticker := time.NewTicker(1 * time.Second)
	for countdown := 10; countdown > 0; countdown-- {
		select {
		case <-ticker.C:
			fmt.Println(countdown)
		case <-abort:
			return
		}
	}
	ticker.Stop()
	Launch()
}

func ExemploForChannel() {
	ch := make(chan int, 1)
	for i := 0; i < 10; i++ {
		select {
		case x := <-ch:
			fmt.Println("Recebendo", x)
		case ch <- i:
			fmt.Println("Enviando", i)
		}
	}
}

func main() {
	RocketTickerExample()
}
