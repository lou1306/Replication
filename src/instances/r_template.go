package main

import (
	"os"
	"fmt"
	"sync"
	"time"
	"math/rand"
	. "github.com/pspaces/gospace"
	. "github.com/repligospaces"
)

var uri = make(map[Space]string)
var Sp = make(map[string]*Space)
var rsp Replispace = Replispace{Sp: Sp}

var debug bool
var l sync.Mutex

var writesloc int
var writesrem int
var readsloc int
var readsrem int
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
	writesloc = 0
	writesrem = 0
	readsloc = 0
	readsrem = 0
	writestotal = 0
	reads_success = 0
	reads_insuccess = 0
	rand.Seed(time.Now().UTC().UnixNano())

	wg.Add(4)

	s1 = NewSpace("tcp://localhost:34001/s1")
	Sp["tcp://localhost:34001/s1"] = &s1
	uri[s1] = "tcp://localhost:34001/s1"
	s2 = NewSpace("tcp://localhost:34002/s2")
	Sp["tcp://localhost:34002/s2"] = &s2
	uri[s2] = "tcp://localhost:34002/s2"
	s3 = NewSpace("tcp://localhost:34003/s3")
	Sp["tcp://localhost:34003/s3"] = &s3
	uri[s3] = "tcp://localhost:34003/s3"
	s4 = NewSpace("tcp://localhost:34004/s4")
	Sp["tcp://localhost:34004/s4"] = &s4
	uri[s4] = "tcp://localhost:34004/s4"

	go p1()
	go p2()
	go p3()
	go p4()

	wg.Wait()

	writes_extra = 0	// change here

	fmt.Println("local writes, remote writes, local reads, remote reads, total writes, successful reads, insuccessful reads")
	fmt.Printf("%d, %d, %d, %d, %d, %d, %d\n", writesloc, writesrem, readsloc, readsrem, writestotal, reads_success, reads_insuccess)
}

func p1() {
	targets1 := make([]string, 3)
	targets1[2] = uri[s4]
	targets1[1] = uri[s3]
	targets1[0] = uri[s2]
	targets0 := make([]string, 3)
	targets0[2] = uri[s4]
	targets0[1] = uri[s3]
	targets0[0] = uri[s2]

	defer wg.Done()
	var value int
	value = 33

	delay()
	l.Lock()
	Put(CreateTuple(3, value), rsp, targets0)
	log("p1 w:%d", value)
	writestotal += 1
	l.Unlock()

	delay()
	l.Lock()
	Put(CreateTuple(3, value), rsp, targets1)
	writestotal += 1
	log("p1 w:%d", value)
	l.Unlock()
}

func p2() {
	defer wg.Done()

	var value int
	value = -1

	delay()
	l.Lock()
	QueryP(CreateTuple(3, &value), rsp, s2)
	log("p2 r:%d", value)
	count(value)
	l.Unlock()
}

func p3() {
	defer wg.Done()

	var value int
	value = -1

	delay()
	l.Lock()
	QueryP(CreateTuple(3, &value), rsp, s3)
	log("p3 r:%d", value)
	count(value)
	l.Unlock()
}

func p4() {
	defer wg.Done()

	var value int
	value = -1

	delay()
	l.Lock()
	QueryP(CreateTuple(3, &value), rsp, s4)
	log("p4 r:%d", value)
	count(value)
	l.Unlock()
}

func log(format string, a ...interface{}) {
	if debug {
		fmt.Fprintf(os.Stdout, format+"\n", a...)
	}
}

func delay() {
	time.Sleep(time.Duration(rand.Int63n(75)) * time.Millisecond)
}

func count(value int) {
	if value == -1 {
		reads_insuccess += 1
	} else {
		reads_success += 1
	}
}
