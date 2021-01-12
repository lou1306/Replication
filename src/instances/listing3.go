package main

import (
	"fmt"
	. "github.com/pspaces/gospace"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(3)

	s1 := NewSpace("tcp://localhost:34115/s1")
	s2 := NewSpace("tcp://localhost:34117/s2")
	s3 := NewSpace("tcp://localhost:34118/s3")

	fmt.Println("Start")
	go Process1(&s1,&s2,&s3,&wg)
	go Process2(&s1,&s2,&s3,&wg)
	go Process3(&s1,&s2,&s3,&wg)

	wg.Wait()
	fmt.Println("Stop")
}

func Process1(s1 *Space, s2 *Space, s3 *Space, wg *sync.WaitGroup) { 
	defer wg.Done()
	var name string = "A"
	var quantity int = 10
	s1.Put(name, quantity)
	fmt.Println(" Tuple stored to s1 is:", name, quantity)
}

func Process2(s1 *Space, s2 *Space, s3 *Space, wg *sync.WaitGroup) {
	defer wg.Done()
	var check string
	var quantity int
	s1.GetP(&check, &quantity)
	fmt.Println(" Tuple retrieved from a space 2:", check, quantity)
}

func Process3(s1 *Space, s2 *Space, s3 *Space, wg *sync.WaitGroup) {  
	defer wg.Done()
	var description string
	var d int
	s1.GetP(&description, &d)
	fmt.Println("---> abcc:", description, d)
}