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
var s5 Space
var s6 Space
var s7 Space
var s8 Space
var s9 Space
var s10 Space
var s11 Space
var s12 Space
var s13 Space
var s14 Space
var s15 Space
var s16 Space
var s17 Space
var s18 Space
var s19 Space
var s20 Space
var s21 Space
var s22 Space
var s23 Space
var s24 Space
var s25 Space
var s26 Space
var s27 Space
var s28 Space
var s29 Space
var s30 Space
var s31 Space
var s32 Space


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

	wg.Add(32)

	s1 = NewSpace("tcp://localhost:34001/s1")
	s2 = NewSpace("tcp://localhost:34002/s2")
	s3 = NewSpace("tcp://localhost:34003/s3")
	s4 = NewSpace("tcp://localhost:34004/s4")
    s5 = NewSpace("tcp://localhost:34005/s5")
    s6 = NewSpace("tcp://localhost:34006/s6")
    s7 = NewSpace("tcp://localhost:34007/s7")
    s8 = NewSpace("tcp://localhost:34008/s8")
    s9 = NewSpace("tcp://localhost:34009/s9")
    s10 = NewSpace("tcp://localhost:34010/s10")
    s11 = NewSpace("tcp://localhost:34011/s11")
    s12 = NewSpace("tcp://localhost:34012/s12")
    s13 = NewSpace("tcp://localhost:34013/s13")
    s14 = NewSpace("tcp://localhost:34014/s14")
    s15 = NewSpace("tcp://localhost:34015/s15")
  	s16 = NewSpace("tcp://localhost:34016/s16")
	s17 = NewSpace("tcp://localhost:34017/s17")
	s18 = NewSpace("tcp://localhost:34018/s18")
    s19 = NewSpace("tcp://localhost:34019/s19")
    s20 = NewSpace("tcp://localhost:34020/s20")
	s21 = NewSpace("tcp://localhost:34021/s21")
	s22 = NewSpace("tcp://localhost:34022/s22")
	s23 = NewSpace("tcp://localhost:34023/s23")
	s24 = NewSpace("tcp://localhost:34024/s24")
	s25 = NewSpace("tcp://localhost:34025/s25")
	s26 = NewSpace("tcp://localhost:34026/s26")
	s27 = NewSpace("tcp://localhost:34027/s27")
	s28 = NewSpace("tcp://localhost:34028/s28")
	s29 = NewSpace("tcp://localhost:34029/s29")
	s30 = NewSpace("tcp://localhost:34030/s30")
	s31 = NewSpace("tcp://localhost:34031/s31")
	s32 = NewSpace("tcp://localhost:34032/s32")



	go p1()
	go p2()
	go p3()
	go p4()
    go p5()
    go p6()
    go p7()
    go p8()
    go p9()
    go p10()
    go p11()
    go p12()
    go p13()
    go p14()
    go p15()
    go p16()
    go p17()
    go p18()
    go p19()
    go p20()
	go p21()
	go p22()
	go p23()
	go p24()
	go p25()
	go p26()
	go p27()
	go p28()
	go p29()
	go p30()
	go p31()
	go p32()

	wg.Wait()

	writes_replicated = 0   // these include the normal put too

	fmt.Println("       loc w,    rem w,   repl w,    loc r,    rem r,    tot w,   succ r,   fail r")
	fmt.Printf ("    %8d, %8d, %8d, %8d, %8d, %8d, %8d, %8d\n", writes_local, writes_remote, writes_replicated, reads_local, reads_remote, writestotal, reads_success, reads_insuccess)
}

func p1() {
	defer wg.Done()
	var value int

	<P1>
}

func p2() {
	defer wg.Done()
	var value int

	<P2>
}

func p3() {
	defer wg.Done()
	var value int

	<P3>
}

func p4() {
	defer wg.Done()
	var value int

	<P4>
}

func p5() {
	defer wg.Done()
	var value int

	<P5>
}

func p6() {
	defer wg.Done()
	var value int

	<P6>
}

func p7() {
	defer wg.Done()
	var value int

	<P7>
}

func p8() {
	defer wg.Done()
	var value int

	<P8>
}

func p9() {
	defer wg.Done()
	var value int

	<P9>
}


func p10() {
	defer wg.Done()
	var value int

	<P10>
}


func p11() {
	defer wg.Done()
	var value int

	<P11>
}


func p12() {
	defer wg.Done()
	var value int

	<P12>
}


func p13() {
	defer wg.Done()
	var value int

	<P13>
}

func p14() {
	defer wg.Done()
	var value int

	<P14>
}


func p15() {
	defer wg.Done()
	var value int

	<P15>
}


func p16() {
	defer wg.Done()
	var value int

	<P16>
}


func p17() {
	defer wg.Done()
	var value int

	<P17>
}


func p18() {
	defer wg.Done()
	var value int

	<P18>
}


func p19() {
	defer wg.Done()
	var value int

	<P19>
}


func p20() {
	defer wg.Done()
	var value int

	<P20>
}


func p21() {
	defer wg.Done()
	var value int

	<P21>
}


func p22() {
	defer wg.Done()
	var value int

	<P22>
}


func p23() {
	defer wg.Done()
	var value int

	<P23>
}


func p24() {
	defer wg.Done()
	var value int

	<P24>
}


func p25() {
	defer wg.Done()
	var value int

	<P25>
}


func p26() {
	defer wg.Done()
	var value int

	<P26>
}


func p27() {
	defer wg.Done()
	var value int

	<P27>
}


func p28() {
	defer wg.Done()
	var value int

	<P28>
}


func p29() {
	defer wg.Done()
	var value int

	<P29>
}


func p30() {
	defer wg.Done()
	var value int

	<P30>
}


func p31() {
	defer wg.Done()
	var value int

	<P31>
}


func p32() {
	defer wg.Done()
	var value int

	<P32>
}


func log(format string, a ...interface{}) {
	if debug {
    	fmt.Fprintf(os.Stdout, format+"\n", a ...)
    }
}

func delay() {
	time.Sleep(time.Duration(rand.Int63n(75))*time.Millisecond)
}

func countreads(value int, localspace int, targetspace int) {
	if value == -1 {
		reads_insuccess+=1
	} else {
		reads_success+=1
	}

	if localspace == targetspace {
		reads_local+=1
	} else {
		reads_remote+=1
	}
}

func countwrites(value int, localspace int, targetspace int) {
	writestotal+=1

	if localspace == targetspace {
		writes_local+=1
	} else {
		writes_remote+=1
	}
}
