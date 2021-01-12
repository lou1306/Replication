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

	delay()
	l.Lock()
	value = -1
	s1.QueryP(3,&value)
	log("p1 r:%d", value)
	countreads(value,1,1)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s20.QueryP(159,&value)
	log("p1 r:%d", value)
	countreads(value,1,20)
	l.Unlock()

	delay()
	l.Lock()
	value = 19
	s19.Put(151,value)
	log("p1 w:%d", value)
	countwrites(value,1,19)
	l.Unlock()

	delay()
	l.Lock()
	value = 9
	s9.Put(69,value)
	log("p1 w:%d", value)
	countwrites(value,1,9)
	l.Unlock()

	delay()
	l.Lock()
	value = 22
	s22.Put(172,value)
	log("p1 w:%d", value)
	countwrites(value,1,22)
	l.Unlock()

	delay()
	l.Lock()
	value = 32
	s32.Put(252,value)
	log("p1 w:%d", value)
	countwrites(value,1,32)
	l.Unlock()

	delay()
	l.Lock()
	value = 26
	s26.Put(207,value)
	log("p1 w:%d", value)
	countwrites(value,1,26)
	l.Unlock()

	delay()
	l.Lock()
	value = 20
	s20.Put(159,value)
	log("p1 w:%d", value)
	countwrites(value,1,20)
	l.Unlock()

	delay()
	l.Lock()
	value = 17
	s17.Put(129,value)
	log("p1 w:%d", value)
	countwrites(value,1,17)
	l.Unlock()

	delay()
	l.Lock()
	value = 21
	s21.Put(168,value)
	log("p1 w:%d", value)
	countwrites(value,1,21)
	l.Unlock()

	delay()
	l.Lock()
	value = 32
	s32.Put(249,value)
	log("p1 w:%d", value)
	countwrites(value,1,32)
	l.Unlock()

	delay()
	l.Lock()
	value = 30
	s30.Put(236,value)
	log("p1 w:%d", value)
	countwrites(value,1,30)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s9.QueryP(69,&value)
	log("p1 r:%d", value)
	countreads(value,1,9)
	l.Unlock()

	delay()
	l.Lock()
	value = 30
	s30.Put(237,value)
	log("p1 w:%d", value)
	countwrites(value,1,30)
	l.Unlock()

	delay()
	l.Lock()
	value = 10
	s10.Put(75,value)
	log("p1 w:%d", value)
	countwrites(value,1,10)
	l.Unlock()

	delay()
	l.Lock()
	value = 24
	s24.Put(189,value)
	log("p1 w:%d", value)
	countwrites(value,1,24)
	l.Unlock()
}

func p2() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 27
	s27.Put(213,value)
	log("p2 w:%d", value)
	countwrites(value,2,27)
	l.Unlock()

	delay()
	l.Lock()
	value = 23
	s23.Put(177,value)
	log("p2 w:%d", value)
	countwrites(value,2,23)
	l.Unlock()

	delay()
	l.Lock()
	value = 18
	s18.Put(139,value)
	log("p2 w:%d", value)
	countwrites(value,2,18)
	l.Unlock()

	delay()
	l.Lock()
	value = 4
	s4.Put(29,value)
	log("p2 w:%d", value)
	countwrites(value,2,4)
	l.Unlock()

	delay()
	l.Lock()
	value = 7
	s7.Put(54,value)
	log("p2 w:%d", value)
	countwrites(value,2,7)
	l.Unlock()

	delay()
	l.Lock()
	value = 7
	s7.Put(53,value)
	log("p2 w:%d", value)
	countwrites(value,2,7)
	l.Unlock()

	delay()
	l.Lock()
	value = 21
	s21.Put(168,value)
	log("p2 w:%d", value)
	countwrites(value,2,21)
	l.Unlock()

	delay()
	l.Lock()
	value = 28
	s28.Put(220,value)
	log("p2 w:%d", value)
	countwrites(value,2,28)
	l.Unlock()

	delay()
	l.Lock()
	value = 24
	s24.Put(191,value)
	log("p2 w:%d", value)
	countwrites(value,2,24)
	l.Unlock()

	delay()
	l.Lock()
	value = 2
	s2.Put(14,value)
	log("p2 w:%d", value)
	countwrites(value,2,2)
	l.Unlock()

	delay()
	l.Lock()
	value = 3
	s3.Put(21,value)
	log("p2 w:%d", value)
	countwrites(value,2,3)
	l.Unlock()

	delay()
	l.Lock()
	value = 10
	s10.Put(80,value)
	log("p2 w:%d", value)
	countwrites(value,2,10)
	l.Unlock()

	delay()
	l.Lock()
	value = 24
	s24.Put(191,value)
	log("p2 w:%d", value)
	countwrites(value,2,24)
	l.Unlock()

	delay()
	l.Lock()
	value = 30
	s30.Put(233,value)
	log("p2 w:%d", value)
	countwrites(value,2,30)
	l.Unlock()

	delay()
	l.Lock()
	value = 9
	s9.Put(67,value)
	log("p2 w:%d", value)
	countwrites(value,2,9)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s1.QueryP(2,&value)
	log("p2 r:%d", value)
	countreads(value,2,1)
	l.Unlock()
}

func p3() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 13
	s13.Put(98,value)
	log("p3 w:%d", value)
	countwrites(value,3,13)
	l.Unlock()

	delay()
	l.Lock()
	value = 17
	s17.Put(131,value)
	log("p3 w:%d", value)
	countwrites(value,3,17)
	l.Unlock()

	delay()
	l.Lock()
	value = 10
	s10.Put(77,value)
	log("p3 w:%d", value)
	countwrites(value,3,10)
	l.Unlock()

	delay()
	l.Lock()
	value = 2
	s2.Put(9,value)
	log("p3 w:%d", value)
	countwrites(value,3,2)
	l.Unlock()

	delay()
	l.Lock()
	value = 8
	s8.Put(58,value)
	log("p3 w:%d", value)
	countwrites(value,3,8)
	l.Unlock()

	delay()
	l.Lock()
	value = 23
	s23.Put(180,value)
	log("p3 w:%d", value)
	countwrites(value,3,23)
	l.Unlock()

	delay()
	l.Lock()
	value = 6
	s6.Put(46,value)
	log("p3 w:%d", value)
	countwrites(value,3,6)
	l.Unlock()

	delay()
	l.Lock()
	value = 18
	s18.Put(141,value)
	log("p3 w:%d", value)
	countwrites(value,3,18)
	l.Unlock()

	delay()
	l.Lock()
	value = 28
	s28.Put(218,value)
	log("p3 w:%d", value)
	countwrites(value,3,28)
	l.Unlock()

	delay()
	l.Lock()
	value = 12
	s12.Put(93,value)
	log("p3 w:%d", value)
	countwrites(value,3,12)
	l.Unlock()

	delay()
	l.Lock()
	value = 4
	s4.Put(31,value)
	log("p3 w:%d", value)
	countwrites(value,3,4)
	l.Unlock()

	delay()
	l.Lock()
	value = 29
	s29.Put(227,value)
	log("p3 w:%d", value)
	countwrites(value,3,29)
	l.Unlock()

	delay()
	l.Lock()
	value = 19
	s19.Put(151,value)
	log("p3 w:%d", value)
	countwrites(value,3,19)
	l.Unlock()

	delay()
	l.Lock()
	value = 30
	s30.Put(240,value)
	log("p3 w:%d", value)
	countwrites(value,3,30)
	l.Unlock()

	delay()
	l.Lock()
	value = 14
	s14.Put(105,value)
	log("p3 w:%d", value)
	countwrites(value,3,14)
	l.Unlock()

	delay()
	l.Lock()
	value = 17
	s17.Put(133,value)
	log("p3 w:%d", value)
	countwrites(value,3,17)
	l.Unlock()
}

func p4() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 18
	s18.Put(144,value)
	log("p4 w:%d", value)
	countwrites(value,4,18)
	l.Unlock()

	delay()
	l.Lock()
	value = 19
	s19.Put(148,value)
	log("p4 w:%d", value)
	countwrites(value,4,19)
	l.Unlock()

	delay()
	l.Lock()
	value = 16
	s16.Put(125,value)
	log("p4 w:%d", value)
	countwrites(value,4,16)
	l.Unlock()

	delay()
	l.Lock()
	value = 4
	s4.Put(31,value)
	log("p4 w:%d", value)
	countwrites(value,4,4)
	l.Unlock()

	delay()
	l.Lock()
	value = 20
	s20.Put(158,value)
	log("p4 w:%d", value)
	countwrites(value,4,20)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s27.QueryP(216,&value)
	log("p4 r:%d", value)
	countreads(value,4,27)
	l.Unlock()

	delay()
	l.Lock()
	value = 25
	s25.Put(194,value)
	log("p4 w:%d", value)
	countwrites(value,4,25)
	l.Unlock()

	delay()
	l.Lock()
	value = 24
	s24.Put(189,value)
	log("p4 w:%d", value)
	countwrites(value,4,24)
	l.Unlock()

	delay()
	l.Lock()
	value = 10
	s10.Put(80,value)
	log("p4 w:%d", value)
	countwrites(value,4,10)
	l.Unlock()

	delay()
	l.Lock()
	value = 12
	s12.Put(91,value)
	log("p4 w:%d", value)
	countwrites(value,4,12)
	l.Unlock()

	delay()
	l.Lock()
	value = 14
	s14.Put(109,value)
	log("p4 w:%d", value)
	countwrites(value,4,14)
	l.Unlock()

	delay()
	l.Lock()
	value = 11
	s11.Put(84,value)
	log("p4 w:%d", value)
	countwrites(value,4,11)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s10.QueryP(79,&value)
	log("p4 r:%d", value)
	countreads(value,4,10)
	l.Unlock()

	delay()
	l.Lock()
	value = 7
	s7.Put(55,value)
	log("p4 w:%d", value)
	countwrites(value,4,7)
	l.Unlock()

	delay()
	l.Lock()
	value = 4
	s4.Put(25,value)
	log("p4 w:%d", value)
	countwrites(value,4,4)
	l.Unlock()

	delay()
	l.Lock()
	value = 31
	s31.Put(243,value)
	log("p4 w:%d", value)
	countwrites(value,4,31)
	l.Unlock()
}

func p5() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 25
	s25.Put(196,value)
	log("p5 w:%d", value)
	countwrites(value,5,25)
	l.Unlock()

	delay()
	l.Lock()
	value = 8
	s8.Put(58,value)
	log("p5 w:%d", value)
	countwrites(value,5,8)
	l.Unlock()

	delay()
	l.Lock()
	value = 16
	s16.Put(121,value)
	log("p5 w:%d", value)
	countwrites(value,5,16)
	l.Unlock()

	delay()
	l.Lock()
	value = 8
	s8.Put(63,value)
	log("p5 w:%d", value)
	countwrites(value,5,8)
	l.Unlock()

	delay()
	l.Lock()
	value = 31
	s31.Put(245,value)
	log("p5 w:%d", value)
	countwrites(value,5,31)
	l.Unlock()

	delay()
	l.Lock()
	value = 23
	s23.Put(184,value)
	log("p5 w:%d", value)
	countwrites(value,5,23)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s16.QueryP(121,&value)
	log("p5 r:%d", value)
	countreads(value,5,16)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s14.QueryP(106,&value)
	log("p5 r:%d", value)
	countreads(value,5,14)
	l.Unlock()

	delay()
	l.Lock()
	value = 17
	s17.Put(129,value)
	log("p5 w:%d", value)
	countwrites(value,5,17)
	l.Unlock()

	delay()
	l.Lock()
	value = 6
	s6.Put(44,value)
	log("p5 w:%d", value)
	countwrites(value,5,6)
	l.Unlock()

	delay()
	l.Lock()
	value = 9
	s9.Put(67,value)
	log("p5 w:%d", value)
	countwrites(value,5,9)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s8.QueryP(59,&value)
	log("p5 r:%d", value)
	countreads(value,5,8)
	l.Unlock()

	delay()
	l.Lock()
	value = 26
	s26.Put(208,value)
	log("p5 w:%d", value)
	countwrites(value,5,26)
	l.Unlock()

	delay()
	l.Lock()
	value = 31
	s31.Put(245,value)
	log("p5 w:%d", value)
	countwrites(value,5,31)
	l.Unlock()

	delay()
	l.Lock()
	value = 16
	s16.Put(123,value)
	log("p5 w:%d", value)
	countwrites(value,5,16)
	l.Unlock()

	delay()
	l.Lock()
	value = 29
	s29.Put(231,value)
	log("p5 w:%d", value)
	countwrites(value,5,29)
	l.Unlock()
}

func p6() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 26
	s26.Put(206,value)
	log("p6 w:%d", value)
	countwrites(value,6,26)
	l.Unlock()

	delay()
	l.Lock()
	value = 8
	s8.Put(59,value)
	log("p6 w:%d", value)
	countwrites(value,6,8)
	l.Unlock()

	delay()
	l.Lock()
	value = 1
	s1.Put(2,value)
	log("p6 w:%d", value)
	countwrites(value,6,1)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s1.QueryP(1,&value)
	log("p6 r:%d", value)
	countreads(value,6,1)
	l.Unlock()

	delay()
	l.Lock()
	value = 27
	s27.Put(213,value)
	log("p6 w:%d", value)
	countwrites(value,6,27)
	l.Unlock()

	delay()
	l.Lock()
	value = 25
	s25.Put(197,value)
	log("p6 w:%d", value)
	countwrites(value,6,25)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s22.QueryP(173,&value)
	log("p6 r:%d", value)
	countreads(value,6,22)
	l.Unlock()

	delay()
	l.Lock()
	value = 29
	s29.Put(229,value)
	log("p6 w:%d", value)
	countwrites(value,6,29)
	l.Unlock()

	delay()
	l.Lock()
	value = 1
	s1.Put(3,value)
	log("p6 w:%d", value)
	countwrites(value,6,1)
	l.Unlock()

	delay()
	l.Lock()
	value = 7
	s7.Put(51,value)
	log("p6 w:%d", value)
	countwrites(value,6,7)
	l.Unlock()

	delay()
	l.Lock()
	value = 14
	s14.Put(108,value)
	log("p6 w:%d", value)
	countwrites(value,6,14)
	l.Unlock()

	delay()
	l.Lock()
	value = 9
	s9.Put(72,value)
	log("p6 w:%d", value)
	countwrites(value,6,9)
	l.Unlock()

	delay()
	l.Lock()
	value = 18
	s18.Put(143,value)
	log("p6 w:%d", value)
	countwrites(value,6,18)
	l.Unlock()

	delay()
	l.Lock()
	value = 19
	s19.Put(149,value)
	log("p6 w:%d", value)
	countwrites(value,6,19)
	l.Unlock()

	delay()
	l.Lock()
	value = 7
	s7.Put(52,value)
	log("p6 w:%d", value)
	countwrites(value,6,7)
	l.Unlock()

	delay()
	l.Lock()
	value = 19
	s19.Put(146,value)
	log("p6 w:%d", value)
	countwrites(value,6,19)
	l.Unlock()
}

func p7() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 28
	s28.Put(222,value)
	log("p7 w:%d", value)
	countwrites(value,7,28)
	l.Unlock()

	delay()
	l.Lock()
	value = 21
	s21.Put(166,value)
	log("p7 w:%d", value)
	countwrites(value,7,21)
	l.Unlock()

	delay()
	l.Lock()
	value = 24
	s24.Put(188,value)
	log("p7 w:%d", value)
	countwrites(value,7,24)
	l.Unlock()

	delay()
	l.Lock()
	value = 28
	s28.Put(219,value)
	log("p7 w:%d", value)
	countwrites(value,7,28)
	l.Unlock()

	delay()
	l.Lock()
	value = 1
	s1.Put(7,value)
	log("p7 w:%d", value)
	countwrites(value,7,1)
	l.Unlock()

	delay()
	l.Lock()
	value = 18
	s18.Put(139,value)
	log("p7 w:%d", value)
	countwrites(value,7,18)
	l.Unlock()

	delay()
	l.Lock()
	value = 22
	s22.Put(175,value)
	log("p7 w:%d", value)
	countwrites(value,7,22)
	l.Unlock()

	delay()
	l.Lock()
	value = 19
	s19.Put(152,value)
	log("p7 w:%d", value)
	countwrites(value,7,19)
	l.Unlock()

	delay()
	l.Lock()
	value = 1
	s1.Put(8,value)
	log("p7 w:%d", value)
	countwrites(value,7,1)
	l.Unlock()

	delay()
	l.Lock()
	value = 31
	s31.Put(248,value)
	log("p7 w:%d", value)
	countwrites(value,7,31)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s8.QueryP(57,&value)
	log("p7 r:%d", value)
	countreads(value,7,8)
	l.Unlock()

	delay()
	l.Lock()
	value = 27
	s27.Put(216,value)
	log("p7 w:%d", value)
	countwrites(value,7,27)
	l.Unlock()

	delay()
	l.Lock()
	value = 6
	s6.Put(45,value)
	log("p7 w:%d", value)
	countwrites(value,7,6)
	l.Unlock()

	delay()
	l.Lock()
	value = 23
	s23.Put(184,value)
	log("p7 w:%d", value)
	countwrites(value,7,23)
	l.Unlock()

	delay()
	l.Lock()
	value = 20
	s20.Put(156,value)
	log("p7 w:%d", value)
	countwrites(value,7,20)
	l.Unlock()

	delay()
	l.Lock()
	value = 24
	s24.Put(185,value)
	log("p7 w:%d", value)
	countwrites(value,7,24)
	l.Unlock()
}

func p8() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 26
	s26.Put(204,value)
	log("p8 w:%d", value)
	countwrites(value,8,26)
	l.Unlock()

	delay()
	l.Lock()
	value = 18
	s18.Put(143,value)
	log("p8 w:%d", value)
	countwrites(value,8,18)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s19.QueryP(151,&value)
	log("p8 r:%d", value)
	countreads(value,8,19)
	l.Unlock()

	delay()
	l.Lock()
	value = 27
	s27.Put(215,value)
	log("p8 w:%d", value)
	countwrites(value,8,27)
	l.Unlock()

	delay()
	l.Lock()
	value = 22
	s22.Put(169,value)
	log("p8 w:%d", value)
	countwrites(value,8,22)
	l.Unlock()

	delay()
	l.Lock()
	value = 25
	s25.Put(195,value)
	log("p8 w:%d", value)
	countwrites(value,8,25)
	l.Unlock()

	delay()
	l.Lock()
	value = 24
	s24.Put(191,value)
	log("p8 w:%d", value)
	countwrites(value,8,24)
	l.Unlock()

	delay()
	l.Lock()
	value = 9
	s9.Put(71,value)
	log("p8 w:%d", value)
	countwrites(value,8,9)
	l.Unlock()

	delay()
	l.Lock()
	value = 22
	s22.Put(175,value)
	log("p8 w:%d", value)
	countwrites(value,8,22)
	l.Unlock()

	delay()
	l.Lock()
	value = 9
	s9.Put(65,value)
	log("p8 w:%d", value)
	countwrites(value,8,9)
	l.Unlock()

	delay()
	l.Lock()
	value = 9
	s9.Put(72,value)
	log("p8 w:%d", value)
	countwrites(value,8,9)
	l.Unlock()

	delay()
	l.Lock()
	value = 31
	s31.Put(241,value)
	log("p8 w:%d", value)
	countwrites(value,8,31)
	l.Unlock()

	delay()
	l.Lock()
	value = 25
	s25.Put(195,value)
	log("p8 w:%d", value)
	countwrites(value,8,25)
	l.Unlock()

	delay()
	l.Lock()
	value = 32
	s32.Put(250,value)
	log("p8 w:%d", value)
	countwrites(value,8,32)
	l.Unlock()

	delay()
	l.Lock()
	value = 30
	s30.Put(239,value)
	log("p8 w:%d", value)
	countwrites(value,8,30)
	l.Unlock()

	delay()
	l.Lock()
	value = 6
	s6.Put(47,value)
	log("p8 w:%d", value)
	countwrites(value,8,6)
	l.Unlock()
}

func p9() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 3
	s3.Put(22,value)
	log("p9 w:%d", value)
	countwrites(value,9,3)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s7.QueryP(52,&value)
	log("p9 r:%d", value)
	countreads(value,9,7)
	l.Unlock()

	delay()
	l.Lock()
	value = 24
	s24.Put(190,value)
	log("p9 w:%d", value)
	countwrites(value,9,24)
	l.Unlock()

	delay()
	l.Lock()
	value = 13
	s13.Put(97,value)
	log("p9 w:%d", value)
	countwrites(value,9,13)
	l.Unlock()

	delay()
	l.Lock()
	value = 10
	s10.Put(73,value)
	log("p9 w:%d", value)
	countwrites(value,9,10)
	l.Unlock()

	delay()
	l.Lock()
	value = 22
	s22.Put(169,value)
	log("p9 w:%d", value)
	countwrites(value,9,22)
	l.Unlock()

	delay()
	l.Lock()
	value = 9
	s9.Put(67,value)
	log("p9 w:%d", value)
	countwrites(value,9,9)
	l.Unlock()

	delay()
	l.Lock()
	value = 23
	s23.Put(180,value)
	log("p9 w:%d", value)
	countwrites(value,9,23)
	l.Unlock()

	delay()
	l.Lock()
	value = 11
	s11.Put(83,value)
	log("p9 w:%d", value)
	countwrites(value,9,11)
	l.Unlock()

	delay()
	l.Lock()
	value = 19
	s19.Put(149,value)
	log("p9 w:%d", value)
	countwrites(value,9,19)
	l.Unlock()

	delay()
	l.Lock()
	value = 19
	s19.Put(152,value)
	log("p9 w:%d", value)
	countwrites(value,9,19)
	l.Unlock()

	delay()
	l.Lock()
	value = 2
	s2.Put(16,value)
	log("p9 w:%d", value)
	countwrites(value,9,2)
	l.Unlock()

	delay()
	l.Lock()
	value = 22
	s22.Put(175,value)
	log("p9 w:%d", value)
	countwrites(value,9,22)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s12.QueryP(89,&value)
	log("p9 r:%d", value)
	countreads(value,9,12)
	l.Unlock()

	delay()
	l.Lock()
	value = 32
	s32.Put(250,value)
	log("p9 w:%d", value)
	countwrites(value,9,32)
	l.Unlock()

	delay()
	l.Lock()
	value = 6
	s6.Put(42,value)
	log("p9 w:%d", value)
	countwrites(value,9,6)
	l.Unlock()
}


func p10() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 26
	s26.Put(205,value)
	log("p10 w:%d", value)
	countwrites(value,10,26)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s21.QueryP(167,&value)
	log("p10 r:%d", value)
	countreads(value,10,21)
	l.Unlock()

	delay()
	l.Lock()
	value = 22
	s22.Put(169,value)
	log("p10 w:%d", value)
	countwrites(value,10,22)
	l.Unlock()

	delay()
	l.Lock()
	value = 26
	s26.Put(204,value)
	log("p10 w:%d", value)
	countwrites(value,10,26)
	l.Unlock()

	delay()
	l.Lock()
	value = 22
	s22.Put(173,value)
	log("p10 w:%d", value)
	countwrites(value,10,22)
	l.Unlock()

	delay()
	l.Lock()
	value = 19
	s19.Put(147,value)
	log("p10 w:%d", value)
	countwrites(value,10,19)
	l.Unlock()

	delay()
	l.Lock()
	value = 27
	s27.Put(214,value)
	log("p10 w:%d", value)
	countwrites(value,10,27)
	l.Unlock()

	delay()
	l.Lock()
	value = 14
	s14.Put(108,value)
	log("p10 w:%d", value)
	countwrites(value,10,14)
	l.Unlock()

	delay()
	l.Lock()
	value = 11
	s11.Put(81,value)
	log("p10 w:%d", value)
	countwrites(value,10,11)
	l.Unlock()

	delay()
	l.Lock()
	value = 4
	s4.Put(32,value)
	log("p10 w:%d", value)
	countwrites(value,10,4)
	l.Unlock()

	delay()
	l.Lock()
	value = 32
	s32.Put(249,value)
	log("p10 w:%d", value)
	countwrites(value,10,32)
	l.Unlock()

	delay()
	l.Lock()
	value = 16
	s16.Put(122,value)
	log("p10 w:%d", value)
	countwrites(value,10,16)
	l.Unlock()

	delay()
	l.Lock()
	value = 15
	s15.Put(113,value)
	log("p10 w:%d", value)
	countwrites(value,10,15)
	l.Unlock()

	delay()
	l.Lock()
	value = 32
	s32.Put(252,value)
	log("p10 w:%d", value)
	countwrites(value,10,32)
	l.Unlock()

	delay()
	l.Lock()
	value = 2
	s2.Put(10,value)
	log("p10 w:%d", value)
	countwrites(value,10,2)
	l.Unlock()

	delay()
	l.Lock()
	value = 29
	s29.Put(229,value)
	log("p10 w:%d", value)
	countwrites(value,10,29)
	l.Unlock()
}


func p11() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	s7.QueryP(50,&value)
	log("p11 r:%d", value)
	countreads(value,11,7)
	l.Unlock()

	delay()
	l.Lock()
	value = 21
	s21.Put(166,value)
	log("p11 w:%d", value)
	countwrites(value,11,21)
	l.Unlock()

	delay()
	l.Lock()
	value = 1
	s1.Put(6,value)
	log("p11 w:%d", value)
	countwrites(value,11,1)
	l.Unlock()

	delay()
	l.Lock()
	value = 12
	s12.Put(95,value)
	log("p11 w:%d", value)
	countwrites(value,11,12)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s23.QueryP(177,&value)
	log("p11 r:%d", value)
	countreads(value,11,23)
	l.Unlock()

	delay()
	l.Lock()
	value = 20
	s20.Put(155,value)
	log("p11 w:%d", value)
	countwrites(value,11,20)
	l.Unlock()

	delay()
	l.Lock()
	value = 20
	s20.Put(154,value)
	log("p11 w:%d", value)
	countwrites(value,11,20)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s26.QueryP(203,&value)
	log("p11 r:%d", value)
	countreads(value,11,26)
	l.Unlock()

	delay()
	l.Lock()
	value = 9
	s9.Put(65,value)
	log("p11 w:%d", value)
	countwrites(value,11,9)
	l.Unlock()

	delay()
	l.Lock()
	value = 21
	s21.Put(167,value)
	log("p11 w:%d", value)
	countwrites(value,11,21)
	l.Unlock()

	delay()
	l.Lock()
	value = 17
	s17.Put(134,value)
	log("p11 w:%d", value)
	countwrites(value,11,17)
	l.Unlock()

	delay()
	l.Lock()
	value = 17
	s17.Put(133,value)
	log("p11 w:%d", value)
	countwrites(value,11,17)
	l.Unlock()

	delay()
	l.Lock()
	value = 26
	s26.Put(203,value)
	log("p11 w:%d", value)
	countwrites(value,11,26)
	l.Unlock()

	delay()
	l.Lock()
	value = 17
	s17.Put(131,value)
	log("p11 w:%d", value)
	countwrites(value,11,17)
	l.Unlock()

	delay()
	l.Lock()
	value = 26
	s26.Put(208,value)
	log("p11 w:%d", value)
	countwrites(value,11,26)
	l.Unlock()

	delay()
	l.Lock()
	value = 7
	s7.Put(52,value)
	log("p11 w:%d", value)
	countwrites(value,11,7)
	l.Unlock()
}


func p12() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 26
	s26.Put(208,value)
	log("p12 w:%d", value)
	countwrites(value,12,26)
	l.Unlock()

	delay()
	l.Lock()
	value = 31
	s31.Put(244,value)
	log("p12 w:%d", value)
	countwrites(value,12,31)
	l.Unlock()

	delay()
	l.Lock()
	value = 23
	s23.Put(184,value)
	log("p12 w:%d", value)
	countwrites(value,12,23)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s4.QueryP(29,&value)
	log("p12 r:%d", value)
	countreads(value,12,4)
	l.Unlock()

	delay()
	l.Lock()
	value = 26
	s26.Put(205,value)
	log("p12 w:%d", value)
	countwrites(value,12,26)
	l.Unlock()

	delay()
	l.Lock()
	value = 25
	s25.Put(200,value)
	log("p12 w:%d", value)
	countwrites(value,12,25)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s23.QueryP(183,&value)
	log("p12 r:%d", value)
	countreads(value,12,23)
	l.Unlock()

	delay()
	l.Lock()
	value = 23
	s23.Put(181,value)
	log("p12 w:%d", value)
	countwrites(value,12,23)
	l.Unlock()

	delay()
	l.Lock()
	value = 11
	s11.Put(84,value)
	log("p12 w:%d", value)
	countwrites(value,12,11)
	l.Unlock()

	delay()
	l.Lock()
	value = 31
	s31.Put(245,value)
	log("p12 w:%d", value)
	countwrites(value,12,31)
	l.Unlock()

	delay()
	l.Lock()
	value = 6
	s6.Put(46,value)
	log("p12 w:%d", value)
	countwrites(value,12,6)
	l.Unlock()

	delay()
	l.Lock()
	value = 28
	s28.Put(219,value)
	log("p12 w:%d", value)
	countwrites(value,12,28)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s29.QueryP(226,&value)
	log("p12 r:%d", value)
	countreads(value,12,29)
	l.Unlock()

	delay()
	l.Lock()
	value = 13
	s13.Put(99,value)
	log("p12 w:%d", value)
	countwrites(value,12,13)
	l.Unlock()

	delay()
	l.Lock()
	value = 1
	s1.Put(5,value)
	log("p12 w:%d", value)
	countwrites(value,12,1)
	l.Unlock()

	delay()
	l.Lock()
	value = 20
	s20.Put(156,value)
	log("p12 w:%d", value)
	countwrites(value,12,20)
	l.Unlock()
}


func p13() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 24
	s24.Put(185,value)
	log("p13 w:%d", value)
	countwrites(value,13,24)
	l.Unlock()

	delay()
	l.Lock()
	value = 10
	s10.Put(73,value)
	log("p13 w:%d", value)
	countwrites(value,13,10)
	l.Unlock()

	delay()
	l.Lock()
	value = 2
	s2.Put(13,value)
	log("p13 w:%d", value)
	countwrites(value,13,2)
	l.Unlock()

	delay()
	l.Lock()
	value = 15
	s15.Put(117,value)
	log("p13 w:%d", value)
	countwrites(value,13,15)
	l.Unlock()

	delay()
	l.Lock()
	value = 19
	s19.Put(146,value)
	log("p13 w:%d", value)
	countwrites(value,13,19)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s32.QueryP(253,&value)
	log("p13 r:%d", value)
	countreads(value,13,32)
	l.Unlock()

	delay()
	l.Lock()
	value = 9
	s9.Put(66,value)
	log("p13 w:%d", value)
	countwrites(value,13,9)
	l.Unlock()

	delay()
	l.Lock()
	value = 10
	s10.Put(75,value)
	log("p13 w:%d", value)
	countwrites(value,13,10)
	l.Unlock()

	delay()
	l.Lock()
	value = 14
	s14.Put(106,value)
	log("p13 w:%d", value)
	countwrites(value,13,14)
	l.Unlock()

	delay()
	l.Lock()
	value = 32
	s32.Put(250,value)
	log("p13 w:%d", value)
	countwrites(value,13,32)
	l.Unlock()

	delay()
	l.Lock()
	value = 12
	s12.Put(94,value)
	log("p13 w:%d", value)
	countwrites(value,13,12)
	l.Unlock()

	delay()
	l.Lock()
	value = 28
	s28.Put(224,value)
	log("p13 w:%d", value)
	countwrites(value,13,28)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s15.QueryP(117,&value)
	log("p13 r:%d", value)
	countreads(value,13,15)
	l.Unlock()

	delay()
	l.Lock()
	value = 13
	s13.Put(104,value)
	log("p13 w:%d", value)
	countwrites(value,13,13)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s22.QueryP(176,&value)
	log("p13 r:%d", value)
	countreads(value,13,22)
	l.Unlock()

	delay()
	l.Lock()
	value = 31
	s31.Put(247,value)
	log("p13 w:%d", value)
	countwrites(value,13,31)
	l.Unlock()
}

func p14() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 21
	s21.Put(167,value)
	log("p14 w:%d", value)
	countwrites(value,14,21)
	l.Unlock()

	delay()
	l.Lock()
	value = 25
	s25.Put(195,value)
	log("p14 w:%d", value)
	countwrites(value,14,25)
	l.Unlock()

	delay()
	l.Lock()
	value = 32
	s32.Put(252,value)
	log("p14 w:%d", value)
	countwrites(value,14,32)
	l.Unlock()

	delay()
	l.Lock()
	value = 29
	s29.Put(225,value)
	log("p14 w:%d", value)
	countwrites(value,14,29)
	l.Unlock()

	delay()
	l.Lock()
	value = 2
	s2.Put(15,value)
	log("p14 w:%d", value)
	countwrites(value,14,2)
	l.Unlock()

	delay()
	l.Lock()
	value = 11
	s11.Put(83,value)
	log("p14 w:%d", value)
	countwrites(value,14,11)
	l.Unlock()

	delay()
	l.Lock()
	value = 15
	s15.Put(118,value)
	log("p14 w:%d", value)
	countwrites(value,14,15)
	l.Unlock()

	delay()
	l.Lock()
	value = 3
	s3.Put(23,value)
	log("p14 w:%d", value)
	countwrites(value,14,3)
	l.Unlock()

	delay()
	l.Lock()
	value = 3
	s3.Put(18,value)
	log("p14 w:%d", value)
	countwrites(value,14,3)
	l.Unlock()

	delay()
	l.Lock()
	value = 21
	s21.Put(161,value)
	log("p14 w:%d", value)
	countwrites(value,14,21)
	l.Unlock()

	delay()
	l.Lock()
	value = 23
	s23.Put(180,value)
	log("p14 w:%d", value)
	countwrites(value,14,23)
	l.Unlock()

	delay()
	l.Lock()
	value = 23
	s23.Put(177,value)
	log("p14 w:%d", value)
	countwrites(value,14,23)
	l.Unlock()

	delay()
	l.Lock()
	value = 31
	s31.Put(242,value)
	log("p14 w:%d", value)
	countwrites(value,14,31)
	l.Unlock()

	delay()
	l.Lock()
	value = 31
	s31.Put(243,value)
	log("p14 w:%d", value)
	countwrites(value,14,31)
	l.Unlock()

	delay()
	l.Lock()
	value = 30
	s30.Put(233,value)
	log("p14 w:%d", value)
	countwrites(value,14,30)
	l.Unlock()

	delay()
	l.Lock()
	value = 15
	s15.Put(115,value)
	log("p14 w:%d", value)
	countwrites(value,14,15)
	l.Unlock()
}


func p15() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 27
	s27.Put(216,value)
	log("p15 w:%d", value)
	countwrites(value,15,27)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s3.QueryP(21,&value)
	log("p15 r:%d", value)
	countreads(value,15,3)
	l.Unlock()

	delay()
	l.Lock()
	value = 1
	s1.Put(4,value)
	log("p15 w:%d", value)
	countwrites(value,15,1)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s4.QueryP(25,&value)
	log("p15 r:%d", value)
	countreads(value,15,4)
	l.Unlock()

	delay()
	l.Lock()
	value = 1
	s1.Put(6,value)
	log("p15 w:%d", value)
	countwrites(value,15,1)
	l.Unlock()

	delay()
	l.Lock()
	value = 23
	s23.Put(184,value)
	log("p15 w:%d", value)
	countwrites(value,15,23)
	l.Unlock()

	delay()
	l.Lock()
	value = 3
	s3.Put(18,value)
	log("p15 w:%d", value)
	countwrites(value,15,3)
	l.Unlock()

	delay()
	l.Lock()
	value = 25
	s25.Put(194,value)
	log("p15 w:%d", value)
	countwrites(value,15,25)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s21.QueryP(163,&value)
	log("p15 r:%d", value)
	countreads(value,15,21)
	l.Unlock()

	delay()
	l.Lock()
	value = 9
	s9.Put(65,value)
	log("p15 w:%d", value)
	countwrites(value,15,9)
	l.Unlock()

	delay()
	l.Lock()
	value = 20
	s20.Put(155,value)
	log("p15 w:%d", value)
	countwrites(value,15,20)
	l.Unlock()

	delay()
	l.Lock()
	value = 19
	s19.Put(146,value)
	log("p15 w:%d", value)
	countwrites(value,15,19)
	l.Unlock()

	delay()
	l.Lock()
	value = 15
	s15.Put(118,value)
	log("p15 w:%d", value)
	countwrites(value,15,15)
	l.Unlock()

	delay()
	l.Lock()
	value = 1
	s1.Put(8,value)
	log("p15 w:%d", value)
	countwrites(value,15,1)
	l.Unlock()

	delay()
	l.Lock()
	value = 14
	s14.Put(107,value)
	log("p15 w:%d", value)
	countwrites(value,15,14)
	l.Unlock()

	delay()
	l.Lock()
	value = 5
	s5.Put(35,value)
	log("p15 w:%d", value)
	countwrites(value,15,5)
	l.Unlock()
}


func p16() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 10
	s10.Put(76,value)
	log("p16 w:%d", value)
	countwrites(value,16,10)
	l.Unlock()

	delay()
	l.Lock()
	value = 26
	s26.Put(206,value)
	log("p16 w:%d", value)
	countwrites(value,16,26)
	l.Unlock()

	delay()
	l.Lock()
	value = 12
	s12.Put(89,value)
	log("p16 w:%d", value)
	countwrites(value,16,12)
	l.Unlock()

	delay()
	l.Lock()
	value = 32
	s32.Put(255,value)
	log("p16 w:%d", value)
	countwrites(value,16,32)
	l.Unlock()

	delay()
	l.Lock()
	value = 6
	s6.Put(47,value)
	log("p16 w:%d", value)
	countwrites(value,16,6)
	l.Unlock()

	delay()
	l.Lock()
	value = 15
	s15.Put(115,value)
	log("p16 w:%d", value)
	countwrites(value,16,15)
	l.Unlock()

	delay()
	l.Lock()
	value = 14
	s14.Put(109,value)
	log("p16 w:%d", value)
	countwrites(value,16,14)
	l.Unlock()

	delay()
	l.Lock()
	value = 20
	s20.Put(155,value)
	log("p16 w:%d", value)
	countwrites(value,16,20)
	l.Unlock()

	delay()
	l.Lock()
	value = 25
	s25.Put(193,value)
	log("p16 w:%d", value)
	countwrites(value,16,25)
	l.Unlock()

	delay()
	l.Lock()
	value = 10
	s10.Put(77,value)
	log("p16 w:%d", value)
	countwrites(value,16,10)
	l.Unlock()

	delay()
	l.Lock()
	value = 29
	s29.Put(225,value)
	log("p16 w:%d", value)
	countwrites(value,16,29)
	l.Unlock()

	delay()
	l.Lock()
	value = 7
	s7.Put(53,value)
	log("p16 w:%d", value)
	countwrites(value,16,7)
	l.Unlock()

	delay()
	l.Lock()
	value = 15
	s15.Put(117,value)
	log("p16 w:%d", value)
	countwrites(value,16,15)
	l.Unlock()

	delay()
	l.Lock()
	value = 5
	s5.Put(37,value)
	log("p16 w:%d", value)
	countwrites(value,16,5)
	l.Unlock()

	delay()
	l.Lock()
	value = 9
	s9.Put(68,value)
	log("p16 w:%d", value)
	countwrites(value,16,9)
	l.Unlock()

	delay()
	l.Lock()
	value = 21
	s21.Put(164,value)
	log("p16 w:%d", value)
	countwrites(value,16,21)
	l.Unlock()
}


func p17() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 1
	s1.Put(3,value)
	log("p17 w:%d", value)
	countwrites(value,17,1)
	l.Unlock()

	delay()
	l.Lock()
	value = 13
	s13.Put(99,value)
	log("p17 w:%d", value)
	countwrites(value,17,13)
	l.Unlock()

	delay()
	l.Lock()
	value = 5
	s5.Put(37,value)
	log("p17 w:%d", value)
	countwrites(value,17,5)
	l.Unlock()

	delay()
	l.Lock()
	value = 31
	s31.Put(242,value)
	log("p17 w:%d", value)
	countwrites(value,17,31)
	l.Unlock()

	delay()
	l.Lock()
	value = 11
	s11.Put(86,value)
	log("p17 w:%d", value)
	countwrites(value,17,11)
	l.Unlock()

	delay()
	l.Lock()
	value = 5
	s5.Put(40,value)
	log("p17 w:%d", value)
	countwrites(value,17,5)
	l.Unlock()

	delay()
	l.Lock()
	value = 22
	s22.Put(176,value)
	log("p17 w:%d", value)
	countwrites(value,17,22)
	l.Unlock()

	delay()
	l.Lock()
	value = 13
	s13.Put(98,value)
	log("p17 w:%d", value)
	countwrites(value,17,13)
	l.Unlock()

	delay()
	l.Lock()
	value = 21
	s21.Put(164,value)
	log("p17 w:%d", value)
	countwrites(value,17,21)
	l.Unlock()

	delay()
	l.Lock()
	value = 10
	s10.Put(74,value)
	log("p17 w:%d", value)
	countwrites(value,17,10)
	l.Unlock()

	delay()
	l.Lock()
	value = 11
	s11.Put(81,value)
	log("p17 w:%d", value)
	countwrites(value,17,11)
	l.Unlock()

	delay()
	l.Lock()
	value = 6
	s6.Put(48,value)
	log("p17 w:%d", value)
	countwrites(value,17,6)
	l.Unlock()

	delay()
	l.Lock()
	value = 3
	s3.Put(21,value)
	log("p17 w:%d", value)
	countwrites(value,17,3)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s20.QueryP(159,&value)
	log("p17 r:%d", value)
	countreads(value,17,20)
	l.Unlock()

	delay()
	l.Lock()
	value = 28
	s28.Put(220,value)
	log("p17 w:%d", value)
	countwrites(value,17,28)
	l.Unlock()

	delay()
	l.Lock()
	value = 8
	s8.Put(59,value)
	log("p17 w:%d", value)
	countwrites(value,17,8)
	l.Unlock()
}


func p18() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 3
	s3.Put(24,value)
	log("p18 w:%d", value)
	countwrites(value,18,3)
	l.Unlock()

	delay()
	l.Lock()
	value = 8
	s8.Put(58,value)
	log("p18 w:%d", value)
	countwrites(value,18,8)
	l.Unlock()

	delay()
	l.Lock()
	value = 26
	s26.Put(202,value)
	log("p18 w:%d", value)
	countwrites(value,18,26)
	l.Unlock()

	delay()
	l.Lock()
	value = 15
	s15.Put(120,value)
	log("p18 w:%d", value)
	countwrites(value,18,15)
	l.Unlock()

	delay()
	l.Lock()
	value = 32
	s32.Put(250,value)
	log("p18 w:%d", value)
	countwrites(value,18,32)
	l.Unlock()

	delay()
	l.Lock()
	value = 24
	s24.Put(191,value)
	log("p18 w:%d", value)
	countwrites(value,18,24)
	l.Unlock()

	delay()
	l.Lock()
	value = 28
	s28.Put(219,value)
	log("p18 w:%d", value)
	countwrites(value,18,28)
	l.Unlock()

	delay()
	l.Lock()
	value = 26
	s26.Put(203,value)
	log("p18 w:%d", value)
	countwrites(value,18,26)
	l.Unlock()

	delay()
	l.Lock()
	value = 14
	s14.Put(108,value)
	log("p18 w:%d", value)
	countwrites(value,18,14)
	l.Unlock()

	delay()
	l.Lock()
	value = 18
	s18.Put(143,value)
	log("p18 w:%d", value)
	countwrites(value,18,18)
	l.Unlock()

	delay()
	l.Lock()
	value = 18
	s18.Put(143,value)
	log("p18 w:%d", value)
	countwrites(value,18,18)
	l.Unlock()

	delay()
	l.Lock()
	value = 4
	s4.Put(27,value)
	log("p18 w:%d", value)
	countwrites(value,18,4)
	l.Unlock()

	delay()
	l.Lock()
	value = 30
	s30.Put(234,value)
	log("p18 w:%d", value)
	countwrites(value,18,30)
	l.Unlock()

	delay()
	l.Lock()
	value = 32
	s32.Put(253,value)
	log("p18 w:%d", value)
	countwrites(value,18,32)
	l.Unlock()

	delay()
	l.Lock()
	value = 31
	s31.Put(241,value)
	log("p18 w:%d", value)
	countwrites(value,18,31)
	l.Unlock()

	delay()
	l.Lock()
	value = 5
	s5.Put(40,value)
	log("p18 w:%d", value)
	countwrites(value,18,5)
	l.Unlock()
}


func p19() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 23
	s23.Put(182,value)
	log("p19 w:%d", value)
	countwrites(value,19,23)
	l.Unlock()

	delay()
	l.Lock()
	value = 17
	s17.Put(129,value)
	log("p19 w:%d", value)
	countwrites(value,19,17)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s18.QueryP(141,&value)
	log("p19 r:%d", value)
	countreads(value,19,18)
	l.Unlock()

	delay()
	l.Lock()
	value = 9
	s9.Put(72,value)
	log("p19 w:%d", value)
	countwrites(value,19,9)
	l.Unlock()

	delay()
	l.Lock()
	value = 25
	s25.Put(197,value)
	log("p19 w:%d", value)
	countwrites(value,19,25)
	l.Unlock()

	delay()
	l.Lock()
	value = 8
	s8.Put(59,value)
	log("p19 w:%d", value)
	countwrites(value,19,8)
	l.Unlock()

	delay()
	l.Lock()
	value = 32
	s32.Put(254,value)
	log("p19 w:%d", value)
	countwrites(value,19,32)
	l.Unlock()

	delay()
	l.Lock()
	value = 6
	s6.Put(42,value)
	log("p19 w:%d", value)
	countwrites(value,19,6)
	l.Unlock()

	delay()
	l.Lock()
	value = 27
	s27.Put(213,value)
	log("p19 w:%d", value)
	countwrites(value,19,27)
	l.Unlock()

	delay()
	l.Lock()
	value = 11
	s11.Put(82,value)
	log("p19 w:%d", value)
	countwrites(value,19,11)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s26.QueryP(204,&value)
	log("p19 r:%d", value)
	countreads(value,19,26)
	l.Unlock()

	delay()
	l.Lock()
	value = 26
	s26.Put(208,value)
	log("p19 w:%d", value)
	countwrites(value,19,26)
	l.Unlock()

	delay()
	l.Lock()
	value = 16
	s16.Put(128,value)
	log("p19 w:%d", value)
	countwrites(value,19,16)
	l.Unlock()

	delay()
	l.Lock()
	value = 24
	s24.Put(186,value)
	log("p19 w:%d", value)
	countwrites(value,19,24)
	l.Unlock()

	delay()
	l.Lock()
	value = 21
	s21.Put(161,value)
	log("p19 w:%d", value)
	countwrites(value,19,21)
	l.Unlock()

	delay()
	l.Lock()
	value = 5
	s5.Put(39,value)
	log("p19 w:%d", value)
	countwrites(value,19,5)
	l.Unlock()
}


func p20() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	s19.QueryP(147,&value)
	log("p20 r:%d", value)
	countreads(value,20,19)
	l.Unlock()

	delay()
	l.Lock()
	value = 13
	s13.Put(101,value)
	log("p20 w:%d", value)
	countwrites(value,20,13)
	l.Unlock()

	delay()
	l.Lock()
	value = 13
	s13.Put(100,value)
	log("p20 w:%d", value)
	countwrites(value,20,13)
	l.Unlock()

	delay()
	l.Lock()
	value = 25
	s25.Put(200,value)
	log("p20 w:%d", value)
	countwrites(value,20,25)
	l.Unlock()

	delay()
	l.Lock()
	value = 3
	s3.Put(24,value)
	log("p20 w:%d", value)
	countwrites(value,20,3)
	l.Unlock()

	delay()
	l.Lock()
	value = 14
	s14.Put(109,value)
	log("p20 w:%d", value)
	countwrites(value,20,14)
	l.Unlock()

	delay()
	l.Lock()
	value = 20
	s20.Put(157,value)
	log("p20 w:%d", value)
	countwrites(value,20,20)
	l.Unlock()

	delay()
	l.Lock()
	value = 9
	s9.Put(71,value)
	log("p20 w:%d", value)
	countwrites(value,20,9)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s22.QueryP(175,&value)
	log("p20 r:%d", value)
	countreads(value,20,22)
	l.Unlock()

	delay()
	l.Lock()
	value = 22
	s22.Put(170,value)
	log("p20 w:%d", value)
	countwrites(value,20,22)
	l.Unlock()

	delay()
	l.Lock()
	value = 21
	s21.Put(162,value)
	log("p20 w:%d", value)
	countwrites(value,20,21)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s18.QueryP(144,&value)
	log("p20 r:%d", value)
	countreads(value,20,18)
	l.Unlock()

	delay()
	l.Lock()
	value = 23
	s23.Put(183,value)
	log("p20 w:%d", value)
	countwrites(value,20,23)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s30.QueryP(235,&value)
	log("p20 r:%d", value)
	countreads(value,20,30)
	l.Unlock()

	delay()
	l.Lock()
	value = 10
	s10.Put(77,value)
	log("p20 w:%d", value)
	countwrites(value,20,10)
	l.Unlock()

	delay()
	l.Lock()
	value = 5
	s5.Put(37,value)
	log("p20 w:%d", value)
	countwrites(value,20,5)
	l.Unlock()
}


func p21() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 29
	s29.Put(226,value)
	log("p21 w:%d", value)
	countwrites(value,21,29)
	l.Unlock()

	delay()
	l.Lock()
	value = 7
	s7.Put(51,value)
	log("p21 w:%d", value)
	countwrites(value,21,7)
	l.Unlock()

	delay()
	l.Lock()
	value = 13
	s13.Put(104,value)
	log("p21 w:%d", value)
	countwrites(value,21,13)
	l.Unlock()

	delay()
	l.Lock()
	value = 8
	s8.Put(58,value)
	log("p21 w:%d", value)
	countwrites(value,21,8)
	l.Unlock()

	delay()
	l.Lock()
	value = 2
	s2.Put(10,value)
	log("p21 w:%d", value)
	countwrites(value,21,2)
	l.Unlock()

	delay()
	l.Lock()
	value = 3
	s3.Put(23,value)
	log("p21 w:%d", value)
	countwrites(value,21,3)
	l.Unlock()

	delay()
	l.Lock()
	value = 4
	s4.Put(31,value)
	log("p21 w:%d", value)
	countwrites(value,21,4)
	l.Unlock()

	delay()
	l.Lock()
	value = 22
	s22.Put(175,value)
	log("p21 w:%d", value)
	countwrites(value,21,22)
	l.Unlock()

	delay()
	l.Lock()
	value = 18
	s18.Put(139,value)
	log("p21 w:%d", value)
	countwrites(value,21,18)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s19.QueryP(149,&value)
	log("p21 r:%d", value)
	countreads(value,21,19)
	l.Unlock()

	delay()
	l.Lock()
	value = 31
	s31.Put(241,value)
	log("p21 w:%d", value)
	countwrites(value,21,31)
	l.Unlock()

	delay()
	l.Lock()
	value = 10
	s10.Put(76,value)
	log("p21 w:%d", value)
	countwrites(value,21,10)
	l.Unlock()

	delay()
	l.Lock()
	value = 4
	s4.Put(26,value)
	log("p21 w:%d", value)
	countwrites(value,21,4)
	l.Unlock()

	delay()
	l.Lock()
	value = 6
	s6.Put(45,value)
	log("p21 w:%d", value)
	countwrites(value,21,6)
	l.Unlock()

	delay()
	l.Lock()
	value = 18
	s18.Put(138,value)
	log("p21 w:%d", value)
	countwrites(value,21,18)
	l.Unlock()

	delay()
	l.Lock()
	value = 18
	s18.Put(138,value)
	log("p21 w:%d", value)
	countwrites(value,21,18)
	l.Unlock()
}


func p22() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 18
	s18.Put(139,value)
	log("p22 w:%d", value)
	countwrites(value,22,18)
	l.Unlock()

	delay()
	l.Lock()
	value = 10
	s10.Put(77,value)
	log("p22 w:%d", value)
	countwrites(value,22,10)
	l.Unlock()

	delay()
	l.Lock()
	value = 28
	s28.Put(223,value)
	log("p22 w:%d", value)
	countwrites(value,22,28)
	l.Unlock()

	delay()
	l.Lock()
	value = 5
	s5.Put(34,value)
	log("p22 w:%d", value)
	countwrites(value,22,5)
	l.Unlock()

	delay()
	l.Lock()
	value = 28
	s28.Put(218,value)
	log("p22 w:%d", value)
	countwrites(value,22,28)
	l.Unlock()

	delay()
	l.Lock()
	value = 4
	s4.Put(31,value)
	log("p22 w:%d", value)
	countwrites(value,22,4)
	l.Unlock()

	delay()
	l.Lock()
	value = 21
	s21.Put(163,value)
	log("p22 w:%d", value)
	countwrites(value,22,21)
	l.Unlock()

	delay()
	l.Lock()
	value = 22
	s22.Put(172,value)
	log("p22 w:%d", value)
	countwrites(value,22,22)
	l.Unlock()

	delay()
	l.Lock()
	value = 30
	s30.Put(240,value)
	log("p22 w:%d", value)
	countwrites(value,22,30)
	l.Unlock()

	delay()
	l.Lock()
	value = 12
	s12.Put(92,value)
	log("p22 w:%d", value)
	countwrites(value,22,12)
	l.Unlock()

	delay()
	l.Lock()
	value = 32
	s32.Put(249,value)
	log("p22 w:%d", value)
	countwrites(value,22,32)
	l.Unlock()

	delay()
	l.Lock()
	value = 8
	s8.Put(57,value)
	log("p22 w:%d", value)
	countwrites(value,22,8)
	l.Unlock()

	delay()
	l.Lock()
	value = 30
	s30.Put(240,value)
	log("p22 w:%d", value)
	countwrites(value,22,30)
	l.Unlock()

	delay()
	l.Lock()
	value = 11
	s11.Put(83,value)
	log("p22 w:%d", value)
	countwrites(value,22,11)
	l.Unlock()

	delay()
	l.Lock()
	value = 19
	s19.Put(147,value)
	log("p22 w:%d", value)
	countwrites(value,22,19)
	l.Unlock()

	delay()
	l.Lock()
	value = 15
	s15.Put(117,value)
	log("p22 w:%d", value)
	countwrites(value,22,15)
	l.Unlock()
}


func p23() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	s19.QueryP(146,&value)
	log("p23 r:%d", value)
	countreads(value,23,19)
	l.Unlock()

	delay()
	l.Lock()
	value = 24
	s24.Put(187,value)
	log("p23 w:%d", value)
	countwrites(value,23,24)
	l.Unlock()

	delay()
	l.Lock()
	value = 1
	s1.Put(5,value)
	log("p23 w:%d", value)
	countwrites(value,23,1)
	l.Unlock()

	delay()
	l.Lock()
	value = 21
	s21.Put(168,value)
	log("p23 w:%d", value)
	countwrites(value,23,21)
	l.Unlock()

	delay()
	l.Lock()
	value = 28
	s28.Put(220,value)
	log("p23 w:%d", value)
	countwrites(value,23,28)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s8.QueryP(64,&value)
	log("p23 r:%d", value)
	countreads(value,23,8)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s28.QueryP(220,&value)
	log("p23 r:%d", value)
	countreads(value,23,28)
	l.Unlock()

	delay()
	l.Lock()
	value = 9
	s9.Put(69,value)
	log("p23 w:%d", value)
	countwrites(value,23,9)
	l.Unlock()

	delay()
	l.Lock()
	value = 17
	s17.Put(132,value)
	log("p23 w:%d", value)
	countwrites(value,23,17)
	l.Unlock()

	delay()
	l.Lock()
	value = 32
	s32.Put(253,value)
	log("p23 w:%d", value)
	countwrites(value,23,32)
	l.Unlock()

	delay()
	l.Lock()
	value = 13
	s13.Put(102,value)
	log("p23 w:%d", value)
	countwrites(value,23,13)
	l.Unlock()

	delay()
	l.Lock()
	value = 18
	s18.Put(142,value)
	log("p23 w:%d", value)
	countwrites(value,23,18)
	l.Unlock()

	delay()
	l.Lock()
	value = 10
	s10.Put(78,value)
	log("p23 w:%d", value)
	countwrites(value,23,10)
	l.Unlock()

	delay()
	l.Lock()
	value = 28
	s28.Put(219,value)
	log("p23 w:%d", value)
	countwrites(value,23,28)
	l.Unlock()

	delay()
	l.Lock()
	value = 10
	s10.Put(75,value)
	log("p23 w:%d", value)
	countwrites(value,23,10)
	l.Unlock()

	delay()
	l.Lock()
	value = 8
	s8.Put(58,value)
	log("p23 w:%d", value)
	countwrites(value,23,8)
	l.Unlock()
}


func p24() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 15
	s15.Put(116,value)
	log("p24 w:%d", value)
	countwrites(value,24,15)
	l.Unlock()

	delay()
	l.Lock()
	value = 2
	s2.Put(9,value)
	log("p24 w:%d", value)
	countwrites(value,24,2)
	l.Unlock()

	delay()
	l.Lock()
	value = 24
	s24.Put(187,value)
	log("p24 w:%d", value)
	countwrites(value,24,24)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s15.QueryP(114,&value)
	log("p24 r:%d", value)
	countreads(value,24,15)
	l.Unlock()

	delay()
	l.Lock()
	value = 17
	s17.Put(133,value)
	log("p24 w:%d", value)
	countwrites(value,24,17)
	l.Unlock()

	delay()
	l.Lock()
	value = 24
	s24.Put(191,value)
	log("p24 w:%d", value)
	countwrites(value,24,24)
	l.Unlock()

	delay()
	l.Lock()
	value = 26
	s26.Put(208,value)
	log("p24 w:%d", value)
	countwrites(value,24,26)
	l.Unlock()

	delay()
	l.Lock()
	value = 9
	s9.Put(70,value)
	log("p24 w:%d", value)
	countwrites(value,24,9)
	l.Unlock()

	delay()
	l.Lock()
	value = 29
	s29.Put(228,value)
	log("p24 w:%d", value)
	countwrites(value,24,29)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s14.QueryP(106,&value)
	log("p24 r:%d", value)
	countreads(value,24,14)
	l.Unlock()

	delay()
	l.Lock()
	value = 24
	s24.Put(190,value)
	log("p24 w:%d", value)
	countwrites(value,24,24)
	l.Unlock()

	delay()
	l.Lock()
	value = 17
	s17.Put(130,value)
	log("p24 w:%d", value)
	countwrites(value,24,17)
	l.Unlock()

	delay()
	l.Lock()
	value = 14
	s14.Put(106,value)
	log("p24 w:%d", value)
	countwrites(value,24,14)
	l.Unlock()

	delay()
	l.Lock()
	value = 31
	s31.Put(243,value)
	log("p24 w:%d", value)
	countwrites(value,24,31)
	l.Unlock()

	delay()
	l.Lock()
	value = 31
	s31.Put(243,value)
	log("p24 w:%d", value)
	countwrites(value,24,31)
	l.Unlock()

	delay()
	l.Lock()
	value = 24
	s24.Put(192,value)
	log("p24 w:%d", value)
	countwrites(value,24,24)
	l.Unlock()
}


func p25() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 21
	s21.Put(165,value)
	log("p25 w:%d", value)
	countwrites(value,25,21)
	l.Unlock()

	delay()
	l.Lock()
	value = 29
	s29.Put(231,value)
	log("p25 w:%d", value)
	countwrites(value,25,29)
	l.Unlock()

	delay()
	l.Lock()
	value = 25
	s25.Put(199,value)
	log("p25 w:%d", value)
	countwrites(value,25,25)
	l.Unlock()

	delay()
	l.Lock()
	value = 6
	s6.Put(46,value)
	log("p25 w:%d", value)
	countwrites(value,25,6)
	l.Unlock()

	delay()
	l.Lock()
	value = 25
	s25.Put(194,value)
	log("p25 w:%d", value)
	countwrites(value,25,25)
	l.Unlock()

	delay()
	l.Lock()
	value = 18
	s18.Put(142,value)
	log("p25 w:%d", value)
	countwrites(value,25,18)
	l.Unlock()

	delay()
	l.Lock()
	value = 27
	s27.Put(209,value)
	log("p25 w:%d", value)
	countwrites(value,25,27)
	l.Unlock()

	delay()
	l.Lock()
	value = 23
	s23.Put(183,value)
	log("p25 w:%d", value)
	countwrites(value,25,23)
	l.Unlock()

	delay()
	l.Lock()
	value = 4
	s4.Put(25,value)
	log("p25 w:%d", value)
	countwrites(value,25,4)
	l.Unlock()

	delay()
	l.Lock()
	value = 18
	s18.Put(141,value)
	log("p25 w:%d", value)
	countwrites(value,25,18)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s16.QueryP(121,&value)
	log("p25 r:%d", value)
	countreads(value,25,16)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s27.QueryP(210,&value)
	log("p25 r:%d", value)
	countreads(value,25,27)
	l.Unlock()

	delay()
	l.Lock()
	value = 29
	s29.Put(227,value)
	log("p25 w:%d", value)
	countwrites(value,25,29)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s14.QueryP(106,&value)
	log("p25 r:%d", value)
	countreads(value,25,14)
	l.Unlock()

	delay()
	l.Lock()
	value = 13
	s13.Put(100,value)
	log("p25 w:%d", value)
	countwrites(value,25,13)
	l.Unlock()

	delay()
	l.Lock()
	value = 21
	s21.Put(167,value)
	log("p25 w:%d", value)
	countwrites(value,25,21)
	l.Unlock()
}


func p26() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 5
	s5.Put(33,value)
	log("p26 w:%d", value)
	countwrites(value,26,5)
	l.Unlock()

	delay()
	l.Lock()
	value = 16
	s16.Put(126,value)
	log("p26 w:%d", value)
	countwrites(value,26,16)
	l.Unlock()

	delay()
	l.Lock()
	value = 28
	s28.Put(219,value)
	log("p26 w:%d", value)
	countwrites(value,26,28)
	l.Unlock()

	delay()
	l.Lock()
	value = 5
	s5.Put(39,value)
	log("p26 w:%d", value)
	countwrites(value,26,5)
	l.Unlock()

	delay()
	l.Lock()
	value = 8
	s8.Put(64,value)
	log("p26 w:%d", value)
	countwrites(value,26,8)
	l.Unlock()

	delay()
	l.Lock()
	value = 27
	s27.Put(214,value)
	log("p26 w:%d", value)
	countwrites(value,26,27)
	l.Unlock()

	delay()
	l.Lock()
	value = 1
	s1.Put(2,value)
	log("p26 w:%d", value)
	countwrites(value,26,1)
	l.Unlock()

	delay()
	l.Lock()
	value = 15
	s15.Put(115,value)
	log("p26 w:%d", value)
	countwrites(value,26,15)
	l.Unlock()

	delay()
	l.Lock()
	value = 15
	s15.Put(118,value)
	log("p26 w:%d", value)
	countwrites(value,26,15)
	l.Unlock()

	delay()
	l.Lock()
	value = 9
	s9.Put(68,value)
	log("p26 w:%d", value)
	countwrites(value,26,9)
	l.Unlock()

	delay()
	l.Lock()
	value = 20
	s20.Put(157,value)
	log("p26 w:%d", value)
	countwrites(value,26,20)
	l.Unlock()

	delay()
	l.Lock()
	value = 15
	s15.Put(114,value)
	log("p26 w:%d", value)
	countwrites(value,26,15)
	l.Unlock()

	delay()
	l.Lock()
	value = 29
	s29.Put(227,value)
	log("p26 w:%d", value)
	countwrites(value,26,29)
	l.Unlock()

	delay()
	l.Lock()
	value = 9
	s9.Put(65,value)
	log("p26 w:%d", value)
	countwrites(value,26,9)
	l.Unlock()

	delay()
	l.Lock()
	value = 27
	s27.Put(212,value)
	log("p26 w:%d", value)
	countwrites(value,26,27)
	l.Unlock()

	delay()
	l.Lock()
	value = 25
	s25.Put(197,value)
	log("p26 w:%d", value)
	countwrites(value,26,25)
	l.Unlock()
}


func p27() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 22
	s22.Put(169,value)
	log("p27 w:%d", value)
	countwrites(value,27,22)
	l.Unlock()

	delay()
	l.Lock()
	value = 5
	s5.Put(35,value)
	log("p27 w:%d", value)
	countwrites(value,27,5)
	l.Unlock()

	delay()
	l.Lock()
	value = 28
	s28.Put(220,value)
	log("p27 w:%d", value)
	countwrites(value,27,28)
	l.Unlock()

	delay()
	l.Lock()
	value = 22
	s22.Put(170,value)
	log("p27 w:%d", value)
	countwrites(value,27,22)
	l.Unlock()

	delay()
	l.Lock()
	value = 28
	s28.Put(224,value)
	log("p27 w:%d", value)
	countwrites(value,27,28)
	l.Unlock()

	delay()
	l.Lock()
	value = 22
	s22.Put(172,value)
	log("p27 w:%d", value)
	countwrites(value,27,22)
	l.Unlock()

	delay()
	l.Lock()
	value = 12
	s12.Put(95,value)
	log("p27 w:%d", value)
	countwrites(value,27,12)
	l.Unlock()

	delay()
	l.Lock()
	value = 24
	s24.Put(191,value)
	log("p27 w:%d", value)
	countwrites(value,27,24)
	l.Unlock()

	delay()
	l.Lock()
	value = 10
	s10.Put(78,value)
	log("p27 w:%d", value)
	countwrites(value,27,10)
	l.Unlock()

	delay()
	l.Lock()
	value = 9
	s9.Put(66,value)
	log("p27 w:%d", value)
	countwrites(value,27,9)
	l.Unlock()

	delay()
	l.Lock()
	value = 15
	s15.Put(115,value)
	log("p27 w:%d", value)
	countwrites(value,27,15)
	l.Unlock()

	delay()
	l.Lock()
	value = 32
	s32.Put(252,value)
	log("p27 w:%d", value)
	countwrites(value,27,32)
	l.Unlock()

	delay()
	l.Lock()
	value = 23
	s23.Put(184,value)
	log("p27 w:%d", value)
	countwrites(value,27,23)
	l.Unlock()

	delay()
	l.Lock()
	value = 6
	s6.Put(41,value)
	log("p27 w:%d", value)
	countwrites(value,27,6)
	l.Unlock()

	delay()
	l.Lock()
	value = 18
	s18.Put(144,value)
	log("p27 w:%d", value)
	countwrites(value,27,18)
	l.Unlock()

	delay()
	l.Lock()
	value = 23
	s23.Put(177,value)
	log("p27 w:%d", value)
	countwrites(value,27,23)
	l.Unlock()
}


func p28() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 14
	s14.Put(110,value)
	log("p28 w:%d", value)
	countwrites(value,28,14)
	l.Unlock()

	delay()
	l.Lock()
	value = 32
	s32.Put(250,value)
	log("p28 w:%d", value)
	countwrites(value,28,32)
	l.Unlock()

	delay()
	l.Lock()
	value = 8
	s8.Put(64,value)
	log("p28 w:%d", value)
	countwrites(value,28,8)
	l.Unlock()

	delay()
	l.Lock()
	value = 18
	s18.Put(143,value)
	log("p28 w:%d", value)
	countwrites(value,28,18)
	l.Unlock()

	delay()
	l.Lock()
	value = 23
	s23.Put(183,value)
	log("p28 w:%d", value)
	countwrites(value,28,23)
	l.Unlock()

	delay()
	l.Lock()
	value = 12
	s12.Put(90,value)
	log("p28 w:%d", value)
	countwrites(value,28,12)
	l.Unlock()

	delay()
	l.Lock()
	value = 4
	s4.Put(31,value)
	log("p28 w:%d", value)
	countwrites(value,28,4)
	l.Unlock()

	delay()
	l.Lock()
	value = 10
	s10.Put(76,value)
	log("p28 w:%d", value)
	countwrites(value,28,10)
	l.Unlock()

	delay()
	l.Lock()
	value = 15
	s15.Put(117,value)
	log("p28 w:%d", value)
	countwrites(value,28,15)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s15.QueryP(115,&value)
	log("p28 r:%d", value)
	countreads(value,28,15)
	l.Unlock()

	delay()
	l.Lock()
	value = 10
	s10.Put(75,value)
	log("p28 w:%d", value)
	countwrites(value,28,10)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s5.QueryP(34,&value)
	log("p28 r:%d", value)
	countreads(value,28,5)
	l.Unlock()

	delay()
	l.Lock()
	value = 15
	s15.Put(118,value)
	log("p28 w:%d", value)
	countwrites(value,28,15)
	l.Unlock()

	delay()
	l.Lock()
	value = 31
	s31.Put(247,value)
	log("p28 w:%d", value)
	countwrites(value,28,31)
	l.Unlock()

	delay()
	l.Lock()
	value = 21
	s21.Put(163,value)
	log("p28 w:%d", value)
	countwrites(value,28,21)
	l.Unlock()

	delay()
	l.Lock()
	value = 29
	s29.Put(229,value)
	log("p28 w:%d", value)
	countwrites(value,28,29)
	l.Unlock()
}


func p29() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 11
	s11.Put(82,value)
	log("p29 w:%d", value)
	countwrites(value,29,11)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s11.QueryP(87,&value)
	log("p29 r:%d", value)
	countreads(value,29,11)
	l.Unlock()

	delay()
	l.Lock()
	value = 13
	s13.Put(101,value)
	log("p29 w:%d", value)
	countwrites(value,29,13)
	l.Unlock()

	delay()
	l.Lock()
	value = 14
	s14.Put(106,value)
	log("p29 w:%d", value)
	countwrites(value,29,14)
	l.Unlock()

	delay()
	l.Lock()
	value = 3
	s3.Put(22,value)
	log("p29 w:%d", value)
	countwrites(value,29,3)
	l.Unlock()

	delay()
	l.Lock()
	value = 2
	s2.Put(9,value)
	log("p29 w:%d", value)
	countwrites(value,29,2)
	l.Unlock()

	delay()
	l.Lock()
	value = 27
	s27.Put(212,value)
	log("p29 w:%d", value)
	countwrites(value,29,27)
	l.Unlock()

	delay()
	l.Lock()
	value = 5
	s5.Put(34,value)
	log("p29 w:%d", value)
	countwrites(value,29,5)
	l.Unlock()

	delay()
	l.Lock()
	value = 13
	s13.Put(102,value)
	log("p29 w:%d", value)
	countwrites(value,29,13)
	l.Unlock()

	delay()
	l.Lock()
	value = 14
	s14.Put(105,value)
	log("p29 w:%d", value)
	countwrites(value,29,14)
	l.Unlock()

	delay()
	l.Lock()
	value = 31
	s31.Put(248,value)
	log("p29 w:%d", value)
	countwrites(value,29,31)
	l.Unlock()

	delay()
	l.Lock()
	value = 7
	s7.Put(53,value)
	log("p29 w:%d", value)
	countwrites(value,29,7)
	l.Unlock()

	delay()
	l.Lock()
	value = 20
	s20.Put(159,value)
	log("p29 w:%d", value)
	countwrites(value,29,20)
	l.Unlock()

	delay()
	l.Lock()
	value = 5
	s5.Put(34,value)
	log("p29 w:%d", value)
	countwrites(value,29,5)
	l.Unlock()

	delay()
	l.Lock()
	value = 27
	s27.Put(210,value)
	log("p29 w:%d", value)
	countwrites(value,29,27)
	l.Unlock()

	delay()
	l.Lock()
	value = 6
	s6.Put(43,value)
	log("p29 w:%d", value)
	countwrites(value,29,6)
	l.Unlock()
}


func p30() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 6
	s6.Put(46,value)
	log("p30 w:%d", value)
	countwrites(value,30,6)
	l.Unlock()

	delay()
	l.Lock()
	value = 3
	s3.Put(22,value)
	log("p30 w:%d", value)
	countwrites(value,30,3)
	l.Unlock()

	delay()
	l.Lock()
	value = 8
	s8.Put(57,value)
	log("p30 w:%d", value)
	countwrites(value,30,8)
	l.Unlock()

	delay()
	l.Lock()
	value = 3
	s3.Put(19,value)
	log("p30 w:%d", value)
	countwrites(value,30,3)
	l.Unlock()

	delay()
	l.Lock()
	value = 5
	s5.Put(39,value)
	log("p30 w:%d", value)
	countwrites(value,30,5)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s17.QueryP(135,&value)
	log("p30 r:%d", value)
	countreads(value,30,17)
	l.Unlock()

	delay()
	l.Lock()
	value = 16
	s16.Put(125,value)
	log("p30 w:%d", value)
	countwrites(value,30,16)
	l.Unlock()

	delay()
	l.Lock()
	value = 16
	s16.Put(122,value)
	log("p30 w:%d", value)
	countwrites(value,30,16)
	l.Unlock()

	delay()
	l.Lock()
	value = 12
	s12.Put(93,value)
	log("p30 w:%d", value)
	countwrites(value,30,12)
	l.Unlock()

	delay()
	l.Lock()
	value = 15
	s15.Put(117,value)
	log("p30 w:%d", value)
	countwrites(value,30,15)
	l.Unlock()

	delay()
	l.Lock()
	value = 29
	s29.Put(227,value)
	log("p30 w:%d", value)
	countwrites(value,30,29)
	l.Unlock()

	delay()
	l.Lock()
	value = 16
	s16.Put(124,value)
	log("p30 w:%d", value)
	countwrites(value,30,16)
	l.Unlock()

	delay()
	l.Lock()
	value = 3
	s3.Put(24,value)
	log("p30 w:%d", value)
	countwrites(value,30,3)
	l.Unlock()

	delay()
	l.Lock()
	value = 21
	s21.Put(161,value)
	log("p30 w:%d", value)
	countwrites(value,30,21)
	l.Unlock()

	delay()
	l.Lock()
	value = 15
	s15.Put(113,value)
	log("p30 w:%d", value)
	countwrites(value,30,15)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s32.QueryP(251,&value)
	log("p30 r:%d", value)
	countreads(value,30,32)
	l.Unlock()
}


func p31() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 7
	s7.Put(51,value)
	log("p31 w:%d", value)
	countwrites(value,31,7)
	l.Unlock()

	delay()
	l.Lock()
	value = 21
	s21.Put(163,value)
	log("p31 w:%d", value)
	countwrites(value,31,21)
	l.Unlock()

	delay()
	l.Lock()
	value = 21
	s21.Put(166,value)
	log("p31 w:%d", value)
	countwrites(value,31,21)
	l.Unlock()

	delay()
	l.Lock()
	value = 19
	s19.Put(150,value)
	log("p31 w:%d", value)
	countwrites(value,31,19)
	l.Unlock()

	delay()
	l.Lock()
	value = 21
	s21.Put(161,value)
	log("p31 w:%d", value)
	countwrites(value,31,21)
	l.Unlock()

	delay()
	l.Lock()
	value = 16
	s16.Put(127,value)
	log("p31 w:%d", value)
	countwrites(value,31,16)
	l.Unlock()

	delay()
	l.Lock()
	value = 26
	s26.Put(207,value)
	log("p31 w:%d", value)
	countwrites(value,31,26)
	l.Unlock()

	delay()
	l.Lock()
	value = 17
	s17.Put(133,value)
	log("p31 w:%d", value)
	countwrites(value,31,17)
	l.Unlock()

	delay()
	l.Lock()
	value = 14
	s14.Put(106,value)
	log("p31 w:%d", value)
	countwrites(value,31,14)
	l.Unlock()

	delay()
	l.Lock()
	value = 10
	s10.Put(77,value)
	log("p31 w:%d", value)
	countwrites(value,31,10)
	l.Unlock()

	delay()
	l.Lock()
	value = 2
	s2.Put(13,value)
	log("p31 w:%d", value)
	countwrites(value,31,2)
	l.Unlock()

	delay()
	l.Lock()
	value = 16
	s16.Put(121,value)
	log("p31 w:%d", value)
	countwrites(value,31,16)
	l.Unlock()

	delay()
	l.Lock()
	value = 15
	s15.Put(116,value)
	log("p31 w:%d", value)
	countwrites(value,31,15)
	l.Unlock()

	delay()
	l.Lock()
	value = 30
	s30.Put(234,value)
	log("p31 w:%d", value)
	countwrites(value,31,30)
	l.Unlock()

	delay()
	l.Lock()
	value = 5
	s5.Put(40,value)
	log("p31 w:%d", value)
	countwrites(value,31,5)
	l.Unlock()

	delay()
	l.Lock()
	value = 7
	s7.Put(50,value)
	log("p31 w:%d", value)
	countwrites(value,31,7)
	l.Unlock()
}


func p32() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 2
	s2.Put(15,value)
	log("p32 w:%d", value)
	countwrites(value,32,2)
	l.Unlock()

	delay()
	l.Lock()
	value = 28
	s28.Put(224,value)
	log("p32 w:%d", value)
	countwrites(value,32,28)
	l.Unlock()

	delay()
	l.Lock()
	value = 13
	s13.Put(101,value)
	log("p32 w:%d", value)
	countwrites(value,32,13)
	l.Unlock()

	delay()
	l.Lock()
	value = 4
	s4.Put(29,value)
	log("p32 w:%d", value)
	countwrites(value,32,4)
	l.Unlock()

	delay()
	l.Lock()
	value = 13
	s13.Put(103,value)
	log("p32 w:%d", value)
	countwrites(value,32,13)
	l.Unlock()

	delay()
	l.Lock()
	value = 4
	s4.Put(29,value)
	log("p32 w:%d", value)
	countwrites(value,32,4)
	l.Unlock()

	delay()
	l.Lock()
	value = 22
	s22.Put(175,value)
	log("p32 w:%d", value)
	countwrites(value,32,22)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s4.QueryP(30,&value)
	log("p32 r:%d", value)
	countreads(value,32,4)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s12.QueryP(92,&value)
	log("p32 r:%d", value)
	countreads(value,32,12)
	l.Unlock()

	delay()
	l.Lock()
	value = 30
	s30.Put(235,value)
	log("p32 w:%d", value)
	countwrites(value,32,30)
	l.Unlock()

	delay()
	l.Lock()
	value = 29
	s29.Put(228,value)
	log("p32 w:%d", value)
	countwrites(value,32,29)
	l.Unlock()

	delay()
	l.Lock()
	value = 30
	s30.Put(237,value)
	log("p32 w:%d", value)
	countwrites(value,32,30)
	l.Unlock()

	delay()
	l.Lock()
	value = 19
	s19.Put(146,value)
	log("p32 w:%d", value)
	countwrites(value,32,19)
	l.Unlock()

	delay()
	l.Lock()
	value = 31
	s31.Put(245,value)
	log("p32 w:%d", value)
	countwrites(value,32,31)
	l.Unlock()

	delay()
	l.Lock()
	value = 27
	s27.Put(214,value)
	log("p32 w:%d", value)
	countwrites(value,32,27)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s1.QueryP(7,&value)
	log("p32 r:%d", value)
	countreads(value,32,1)
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
