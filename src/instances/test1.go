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

var writes_local int
var writes_remote int
var reads_local int
var reads_remote int
var writes_replicated int
var writestotal int
var reads_success int
var reads_insuccess int
var writes_extra int
var wg sync.WaitGroup

var s1 Space
var s2 Space
var s3 Space
var s4 Space

func main() {
	debug = false
	writes_local = 0
	writes_remote = 0
	reads_local = 0
	reads_remote = 0
	writes_replicated = 0
	reads_success = 0
	reads_insuccess = 0
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

	writes_replicated = 0   // these include the normal put too

	fmt.Println("   loc w,    rem w,   repl w,    loc r,    rem r,    tot w,   succ r,   fail r")
	fmt.Printf ("%8d, %8d, %8d, %8d, %8d, %8d, %8d, %8d\n", writes_local, writes_remote, writes_replicated, reads_local, reads_remote, writestotal, reads_success, reads_insuccess)
}

func p1() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	s4.QueryP(14,&value)
	log("p1 r:%d", value)
	count(value)
	l.Unlock()

	delay()
	l.Lock()
	value = 4
	s4.Put(15,value)
	log("p1 w:%d", value)
	writestotal+=1
	l.Unlock()
}

func p2() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	s3.QueryP(9,&value)
	log("p2 r:%d", value)
	count(value)
	l.Unlock()

	delay()
	l.Lock()
	value = 4
	s4.Put(13,value)
	log("p2 w:%d", value)
	writestotal+=1
	l.Unlock()
}

func p3() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 4
	s4.Put(16,value)
	log("p3 w:%d", value)
	writestotal+=1
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s4.QueryP(16,&value)
	log("p3 r:%d", value)
	count(value)
	l.Unlock()
}

func p4() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 1
	s1.Put(3,value)
	log("p4 w:%d", value)
	writestotal+=1
	l.Unlock()

	delay()
	l.Lock()
	value = 1
	s1.Put(2,value)
	log("p4 w:%d", value)
	writestotal+=1
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
		reads_insuccess+=1
	} else {
		reads_success+=1
	}
}
