package main

import (
	"fmt"
	"sync"
	// "sync"
	"math/rand"
	"sync/atomic"
	"time"
)

// var mutex = &sync.Mutex{}

// var initial_capital int
// func credit(){
// 	for i :=0; i < 5; i++{
// 		mutex.Lock()
// 		initial_capital +=100
// 		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
// 		fmt.Println("After crediting: ",initial_capital)
// 		mutex.Unlock()
// 	}
// }

// func debit(){
// 	for i :=0; i < 5; i++{
// 		mutex.Lock()
// 		initial_capital -=100
// 		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
// 		fmt.Println("After debiting: ", initial_capital)
// 		mutex.Unlock()
// 	}
// }



// MEHTOD 2 sync/atomic
// var initial_capital int64
// func credit(){
// 	for i :=0; i < 10; i++{
// 		atomic.AddInt64(&initial_capital, 100) //takes initial value and another value
// 		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
// 	}
// }

// func debit(){
// 	for i :=0; i < 5; i++{
// 		atomic.AddInt64(&initial_capital, -100) //takes initital value and value to substract
// 		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
// 	}
// }

// func main(){
// 	// var wg sync.WaitGroup
// 	initial_capital = 200
// 	fmt.Println("Initial balance is", initial_capital)
// 	// wg.Add(1)
// 	go credit()
// 	// wg.Add(1)
// 	go debit()
// 	fmt.Scanln()
// 	fmt.Println(initial_capital)
// }

// using WAITGROUPS method 3
var initial_capital int64
func credit(wg *sync.WaitGroup){
	defer wg.Done()
	for i :=0; i < 10; i++{
		atomic.AddInt64(&initial_capital, 100) //takes initial value and another value
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
	}
}

func debit(wg *sync.WaitGroup){
	defer wg.Done()
	for i :=0; i < 5; i++{
		atomic.AddInt64(&initial_capital, -100) //takes initital value and value to substract
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
	}
}

func main() {
	var wg sync.WaitGroup
	initial_capital = 200
	fmt.Println("Initial balance is", initial_capital)
	wg.Add(1)
	go credit(&wg)
	wg.Add(1)
	go debit(&wg)
	wg.Wait() // use Wait to wait for all goroutines to complete
	fmt.Println(initial_capital)
	ch:= make(chan int)
	go Channels(ch)
	go Add(ch, 2)
	fmt.Scanln()
	fmt.Println("hello wrold")
	c := make(chan int)
	go fib(10, c)
	for j:=range c{
		fmt.Println(j)
	}
	nacci(5)
}

func Add(ch chan int, r int){	
	fmt.Println("Receiving a value from channel")
	time.Sleep(2 * time.Second)

	a := <-ch
	fmt.Println(a+9+r)
}

func Channels(ch chan int){
	fmt.Println("Sending a value into the channel")
	ch<-6
}

func fib(n int, c chan int){
	a, b := 1, 1
	for i:=0; i<n; i++{
		c <- a
		a, b = b, a+b
		time.Sleep(1 * time.Second)
	}
	close(c)
}


func nacci(n int){
	a, b := 1, 1
	for i := 0; i<n; i++{
		a, b = b, a+b
		fmt.Println(a, b)
	}
}

