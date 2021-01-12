package main

import (
	"os"
	"fmt"
	"sync"
	"time"
	"math/rand"
	. "github.com/pspaces/gospace"
)

var debug bool
var l sync.Mutex
var wg sync.WaitGroup
var nowrites int
var noreads int
var noinconclusivereads int

var s1 Space
var s2 Space
var s3 Space
var s4 Space

func main() {
	debug = false
	nowrites = 0
	noreads = 0
	noinconclusivereads = 0
	rand.Seed(time.Now().UTC().UnixNano())

	wg.Add(4)

	s1 = NewSpace("tcp://localhost:34001/s1")
	s2 = NewSpace("tcp://localhost:34002/s2")
	s3 = NewSpace("tcp://localhost:34003/s3")
	s4 = NewSpace("tcp://localhost:34004/s4")

	go p1()
	go p2()
	go p3()
	go p4()

	wg.Wait()

	fmt.Println("writes, successful reads, inconclusive reads")
	fmt.Printf ("%d,      %d,                %d\n", nowrites, noreads, noinconclusivereads)
}

func p1() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 33
	s1.Put(3,value)
	log("p1 w:%d", value)
	nowrites+=1
	l.Unlock()

	delay()
	l.Lock()
	value = 33
	s1.Put(3,value)
	nowrites+=1
	log("p1 w:%d", value)
	l.Unlock()
}

func p2() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	s1.QueryP(3,&value)
	log("p2 r:%d", value)
	count(value)
	l.Unlock()
}

func p3() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	s1.QueryP(3,&value)
	log("p3 r:%d", value)
	count(value)
	l.Unlock()
}

func p4() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	s1.QueryP(3,&value)
	log("p4 r:%d", value)
	count(value)
	l.Unlock()
}

func log(format string, a ...interface{}) {
	if debug {
    	fmt.Fprintf(os.Stdout, format+"\n", a ...)
    }
}

func delay() {
	time.Sleep(time.Duration(rand.Int63n(75))*time.Millisecond)
}

func count(value int) {
	if value == -1 {
		noinconclusivereads+=1
	} else {
		noreads+=1
	}
}
