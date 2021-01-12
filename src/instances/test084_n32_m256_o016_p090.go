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
	value = 13
	s13.Put(99,value)
	log("p1 w:%d", value)
	countwrites(value,1,13)
	l.Unlock()

	delay()
	l.Lock()
	value = 12
	s12.Put(89,value)
	log("p1 w:%d", value)
	countwrites(value,1,12)
	l.Unlock()

	delay()
	l.Lock()
	value = 16
	s16.Put(126,value)
	log("p1 w:%d", value)
	countwrites(value,1,16)
	l.Unlock()

	delay()
	l.Lock()
	value = 9
	s9.Put(68,value)
	log("p1 w:%d", value)
	countwrites(value,1,9)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s6.QueryP(46,&value)
	log("p1 r:%d", value)
	countreads(value,1,6)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s10.QueryP(76,&value)
	log("p1 r:%d", value)
	countreads(value,1,10)
	l.Unlock()

	delay()
	l.Lock()
	value = 27
	s27.Put(214,value)
	log("p1 w:%d", value)
	countwrites(value,1,27)
	l.Unlock()

	delay()
	l.Lock()
	value = 18
	s18.Put(141,value)
	log("p1 w:%d", value)
	countwrites(value,1,18)
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
	value = 3
	s3.Put(22,value)
	log("p1 w:%d", value)
	countwrites(value,1,3)
	l.Unlock()

	delay()
	l.Lock()
	value = 17
	s17.Put(130,value)
	log("p1 w:%d", value)
	countwrites(value,1,17)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s30.QueryP(233,&value)
	log("p1 r:%d", value)
	countreads(value,1,30)
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
	value = 28
	s28.Put(223,value)
	log("p1 w:%d", value)
	countwrites(value,1,28)
	l.Unlock()

	delay()
	l.Lock()
	value = 20
	s20.Put(157,value)
	log("p1 w:%d", value)
	countwrites(value,1,20)
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
	value = 9
	s9.Put(67,value)
	log("p2 w:%d", value)
	countwrites(value,2,9)
	l.Unlock()

	delay()
	l.Lock()
	value = 23
	s23.Put(181,value)
	log("p2 w:%d", value)
	countwrites(value,2,23)
	l.Unlock()

	delay()
	l.Lock()
	value = 20
	s20.Put(154,value)
	log("p2 w:%d", value)
	countwrites(value,2,20)
	l.Unlock()

	delay()
	l.Lock()
	value = 25
	s25.Put(195,value)
	log("p2 w:%d", value)
	countwrites(value,2,25)
	l.Unlock()

	delay()
	l.Lock()
	value = 28
	s28.Put(218,value)
	log("p2 w:%d", value)
	countwrites(value,2,28)
	l.Unlock()

	delay()
	l.Lock()
	value = 4
	s4.Put(30,value)
	log("p2 w:%d", value)
	countwrites(value,2,4)
	l.Unlock()

	delay()
	l.Lock()
	value = 18
	s18.Put(142,value)
	log("p2 w:%d", value)
	countwrites(value,2,18)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s18.QueryP(137,&value)
	log("p2 r:%d", value)
	countreads(value,2,18)
	l.Unlock()

	delay()
	l.Lock()
	value = 14
	s14.Put(112,value)
	log("p2 w:%d", value)
	countwrites(value,2,14)
	l.Unlock()

	delay()
	l.Lock()
	value = 27
	s27.Put(216,value)
	log("p2 w:%d", value)
	countwrites(value,2,27)
	l.Unlock()

	delay()
	l.Lock()
	value = 19
	s19.Put(150,value)
	log("p2 w:%d", value)
	countwrites(value,2,19)
	l.Unlock()

	delay()
	l.Lock()
	value = 13
	s13.Put(103,value)
	log("p2 w:%d", value)
	countwrites(value,2,13)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s10.QueryP(80,&value)
	log("p2 r:%d", value)
	countreads(value,2,10)
	l.Unlock()

	delay()
	l.Lock()
	value = 1
	s1.Put(5,value)
	log("p2 w:%d", value)
	countwrites(value,2,1)
	l.Unlock()

	delay()
	l.Lock()
	value = 30
	s30.Put(238,value)
	log("p2 w:%d", value)
	countwrites(value,2,30)
	l.Unlock()

	delay()
	l.Lock()
	value = 14
	s14.Put(107,value)
	log("p2 w:%d", value)
	countwrites(value,2,14)
	l.Unlock()
}

func p3() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 15
	s15.Put(120,value)
	log("p3 w:%d", value)
	countwrites(value,3,15)
	l.Unlock()

	delay()
	l.Lock()
	value = 6
	s6.Put(44,value)
	log("p3 w:%d", value)
	countwrites(value,3,6)
	l.Unlock()

	delay()
	l.Lock()
	value = 21
	s21.Put(162,value)
	log("p3 w:%d", value)
	countwrites(value,3,21)
	l.Unlock()

	delay()
	l.Lock()
	value = 26
	s26.Put(205,value)
	log("p3 w:%d", value)
	countwrites(value,3,26)
	l.Unlock()

	delay()
	l.Lock()
	value = 7
	s7.Put(50,value)
	log("p3 w:%d", value)
	countwrites(value,3,7)
	l.Unlock()

	delay()
	l.Lock()
	value = 10
	s10.Put(80,value)
	log("p3 w:%d", value)
	countwrites(value,3,10)
	l.Unlock()

	delay()
	l.Lock()
	value = 2
	s2.Put(15,value)
	log("p3 w:%d", value)
	countwrites(value,3,2)
	l.Unlock()

	delay()
	l.Lock()
	value = 22
	s22.Put(174,value)
	log("p3 w:%d", value)
	countwrites(value,3,22)
	l.Unlock()

	delay()
	l.Lock()
	value = 1
	s1.Put(5,value)
	log("p3 w:%d", value)
	countwrites(value,3,1)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s1.QueryP(2,&value)
	log("p3 r:%d", value)
	countreads(value,3,1)
	l.Unlock()

	delay()
	l.Lock()
	value = 19
	s19.Put(152,value)
	log("p3 w:%d", value)
	countwrites(value,3,19)
	l.Unlock()

	delay()
	l.Lock()
	value = 16
	s16.Put(124,value)
	log("p3 w:%d", value)
	countwrites(value,3,16)
	l.Unlock()

	delay()
	l.Lock()
	value = 9
	s9.Put(67,value)
	log("p3 w:%d", value)
	countwrites(value,3,9)
	l.Unlock()

	delay()
	l.Lock()
	value = 1
	s1.Put(6,value)
	log("p3 w:%d", value)
	countwrites(value,3,1)
	l.Unlock()

	delay()
	l.Lock()
	value = 7
	s7.Put(55,value)
	log("p3 w:%d", value)
	countwrites(value,3,7)
	l.Unlock()

	delay()
	l.Lock()
	value = 26
	s26.Put(207,value)
	log("p3 w:%d", value)
	countwrites(value,3,26)
	l.Unlock()
}

func p4() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	s18.QueryP(140,&value)
	log("p4 r:%d", value)
	countreads(value,4,18)
	l.Unlock()

	delay()
	l.Lock()
	value = 15
	s15.Put(113,value)
	log("p4 w:%d", value)
	countwrites(value,4,15)
	l.Unlock()

	delay()
	l.Lock()
	value = 13
	s13.Put(100,value)
	log("p4 w:%d", value)
	countwrites(value,4,13)
	l.Unlock()

	delay()
	l.Lock()
	value = 11
	s11.Put(87,value)
	log("p4 w:%d", value)
	countwrites(value,4,11)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s16.QueryP(124,&value)
	log("p4 r:%d", value)
	countreads(value,4,16)
	l.Unlock()

	delay()
	l.Lock()
	value = 23
	s23.Put(178,value)
	log("p4 w:%d", value)
	countwrites(value,4,23)
	l.Unlock()

	delay()
	l.Lock()
	value = 26
	s26.Put(205,value)
	log("p4 w:%d", value)
	countwrites(value,4,26)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s32.QueryP(253,&value)
	log("p4 r:%d", value)
	countreads(value,4,32)
	l.Unlock()

	delay()
	l.Lock()
	value = 8
	s8.Put(59,value)
	log("p4 w:%d", value)
	countwrites(value,4,8)
	l.Unlock()

	delay()
	l.Lock()
	value = 32
	s32.Put(253,value)
	log("p4 w:%d", value)
	countwrites(value,4,32)
	l.Unlock()

	delay()
	l.Lock()
	value = 23
	s23.Put(184,value)
	log("p4 w:%d", value)
	countwrites(value,4,23)
	l.Unlock()

	delay()
	l.Lock()
	value = 10
	s10.Put(78,value)
	log("p4 w:%d", value)
	countwrites(value,4,10)
	l.Unlock()

	delay()
	l.Lock()
	value = 29
	s29.Put(225,value)
	log("p4 w:%d", value)
	countwrites(value,4,29)
	l.Unlock()

	delay()
	l.Lock()
	value = 1
	s1.Put(8,value)
	log("p4 w:%d", value)
	countwrites(value,4,1)
	l.Unlock()

	delay()
	l.Lock()
	value = 27
	s27.Put(209,value)
	log("p4 w:%d", value)
	countwrites(value,4,27)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s10.QueryP(78,&value)
	log("p4 r:%d", value)
	countreads(value,4,10)
	l.Unlock()
}

func p5() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 12
	s12.Put(90,value)
	log("p5 w:%d", value)
	countwrites(value,5,12)
	l.Unlock()

	delay()
	l.Lock()
	value = 21
	s21.Put(161,value)
	log("p5 w:%d", value)
	countwrites(value,5,21)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s3.QueryP(24,&value)
	log("p5 r:%d", value)
	countreads(value,5,3)
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
	value = 3
	s3.Put(17,value)
	log("p5 w:%d", value)
	countwrites(value,5,3)
	l.Unlock()

	delay()
	l.Lock()
	value = 17
	s17.Put(136,value)
	log("p5 w:%d", value)
	countwrites(value,5,17)
	l.Unlock()

	delay()
	l.Lock()
	value = 21
	s21.Put(161,value)
	log("p5 w:%d", value)
	countwrites(value,5,21)
	l.Unlock()

	delay()
	l.Lock()
	value = 14
	s14.Put(108,value)
	log("p5 w:%d", value)
	countwrites(value,5,14)
	l.Unlock()

	delay()
	l.Lock()
	value = 20
	s20.Put(160,value)
	log("p5 w:%d", value)
	countwrites(value,5,20)
	l.Unlock()

	delay()
	l.Lock()
	value = 5
	s5.Put(37,value)
	log("p5 w:%d", value)
	countwrites(value,5,5)
	l.Unlock()

	delay()
	l.Lock()
	value = 7
	s7.Put(50,value)
	log("p5 w:%d", value)
	countwrites(value,5,7)
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
	value = 10
	s10.Put(77,value)
	log("p5 w:%d", value)
	countwrites(value,5,10)
	l.Unlock()

	delay()
	l.Lock()
	value = 19
	s19.Put(149,value)
	log("p5 w:%d", value)
	countwrites(value,5,19)
	l.Unlock()

	delay()
	l.Lock()
	value = 18
	s18.Put(141,value)
	log("p5 w:%d", value)
	countwrites(value,5,18)
	l.Unlock()

	delay()
	l.Lock()
	value = 30
	s30.Put(240,value)
	log("p5 w:%d", value)
	countwrites(value,5,30)
	l.Unlock()
}

func p6() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 14
	s14.Put(110,value)
	log("p6 w:%d", value)
	countwrites(value,6,14)
	l.Unlock()

	delay()
	l.Lock()
	value = 23
	s23.Put(183,value)
	log("p6 w:%d", value)
	countwrites(value,6,23)
	l.Unlock()

	delay()
	l.Lock()
	value = 7
	s7.Put(54,value)
	log("p6 w:%d", value)
	countwrites(value,6,7)
	l.Unlock()

	delay()
	l.Lock()
	value = 5
	s5.Put(38,value)
	log("p6 w:%d", value)
	countwrites(value,6,5)
	l.Unlock()

	delay()
	l.Lock()
	value = 17
	s17.Put(131,value)
	log("p6 w:%d", value)
	countwrites(value,6,17)
	l.Unlock()

	delay()
	l.Lock()
	value = 29
	s29.Put(232,value)
	log("p6 w:%d", value)
	countwrites(value,6,29)
	l.Unlock()

	delay()
	l.Lock()
	value = 5
	s5.Put(39,value)
	log("p6 w:%d", value)
	countwrites(value,6,5)
	l.Unlock()

	delay()
	l.Lock()
	value = 4
	s4.Put(29,value)
	log("p6 w:%d", value)
	countwrites(value,6,4)
	l.Unlock()

	delay()
	l.Lock()
	value = 3
	s3.Put(24,value)
	log("p6 w:%d", value)
	countwrites(value,6,3)
	l.Unlock()

	delay()
	l.Lock()
	value = 18
	s18.Put(138,value)
	log("p6 w:%d", value)
	countwrites(value,6,18)
	l.Unlock()

	delay()
	l.Lock()
	value = 5
	s5.Put(37,value)
	log("p6 w:%d", value)
	countwrites(value,6,5)
	l.Unlock()

	delay()
	l.Lock()
	value = 8
	s8.Put(64,value)
	log("p6 w:%d", value)
	countwrites(value,6,8)
	l.Unlock()

	delay()
	l.Lock()
	value = 31
	s31.Put(243,value)
	log("p6 w:%d", value)
	countwrites(value,6,31)
	l.Unlock()

	delay()
	l.Lock()
	value = 32
	s32.Put(249,value)
	log("p6 w:%d", value)
	countwrites(value,6,32)
	l.Unlock()

	delay()
	l.Lock()
	value = 31
	s31.Put(242,value)
	log("p6 w:%d", value)
	countwrites(value,6,31)
	l.Unlock()

	delay()
	l.Lock()
	value = 24
	s24.Put(190,value)
	log("p6 w:%d", value)
	countwrites(value,6,24)
	l.Unlock()
}

func p7() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 20
	s20.Put(159,value)
	log("p7 w:%d", value)
	countwrites(value,7,20)
	l.Unlock()

	delay()
	l.Lock()
	value = 32
	s32.Put(256,value)
	log("p7 w:%d", value)
	countwrites(value,7,32)
	l.Unlock()

	delay()
	l.Lock()
	value = 1
	s1.Put(6,value)
	log("p7 w:%d", value)
	countwrites(value,7,1)
	l.Unlock()

	delay()
	l.Lock()
	value = 4
	s4.Put(26,value)
	log("p7 w:%d", value)
	countwrites(value,7,4)
	l.Unlock()

	delay()
	l.Lock()
	value = 5
	s5.Put(39,value)
	log("p7 w:%d", value)
	countwrites(value,7,5)
	l.Unlock()

	delay()
	l.Lock()
	value = 26
	s26.Put(206,value)
	log("p7 w:%d", value)
	countwrites(value,7,26)
	l.Unlock()

	delay()
	l.Lock()
	value = 28
	s28.Put(218,value)
	log("p7 w:%d", value)
	countwrites(value,7,28)
	l.Unlock()

	delay()
	l.Lock()
	value = 11
	s11.Put(82,value)
	log("p7 w:%d", value)
	countwrites(value,7,11)
	l.Unlock()

	delay()
	l.Lock()
	value = 6
	s6.Put(46,value)
	log("p7 w:%d", value)
	countwrites(value,7,6)
	l.Unlock()

	delay()
	l.Lock()
	value = 31
	s31.Put(247,value)
	log("p7 w:%d", value)
	countwrites(value,7,31)
	l.Unlock()

	delay()
	l.Lock()
	value = 4
	s4.Put(32,value)
	log("p7 w:%d", value)
	countwrites(value,7,4)
	l.Unlock()

	delay()
	l.Lock()
	value = 18
	s18.Put(144,value)
	log("p7 w:%d", value)
	countwrites(value,7,18)
	l.Unlock()

	delay()
	l.Lock()
	value = 7
	s7.Put(55,value)
	log("p7 w:%d", value)
	countwrites(value,7,7)
	l.Unlock()

	delay()
	l.Lock()
	value = 10
	s10.Put(79,value)
	log("p7 w:%d", value)
	countwrites(value,7,10)
	l.Unlock()

	delay()
	l.Lock()
	value = 9
	s9.Put(69,value)
	log("p7 w:%d", value)
	countwrites(value,7,9)
	l.Unlock()

	delay()
	l.Lock()
	value = 2
	s2.Put(13,value)
	log("p7 w:%d", value)
	countwrites(value,7,2)
	l.Unlock()
}

func p8() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 25
	s25.Put(199,value)
	log("p8 w:%d", value)
	countwrites(value,8,25)
	l.Unlock()

	delay()
	l.Lock()
	value = 31
	s31.Put(242,value)
	log("p8 w:%d", value)
	countwrites(value,8,31)
	l.Unlock()

	delay()
	l.Lock()
	value = 13
	s13.Put(104,value)
	log("p8 w:%d", value)
	countwrites(value,8,13)
	l.Unlock()

	delay()
	l.Lock()
	value = 8
	s8.Put(61,value)
	log("p8 w:%d", value)
	countwrites(value,8,8)
	l.Unlock()

	delay()
	l.Lock()
	value = 12
	s12.Put(95,value)
	log("p8 w:%d", value)
	countwrites(value,8,12)
	l.Unlock()

	delay()
	l.Lock()
	value = 12
	s12.Put(94,value)
	log("p8 w:%d", value)
	countwrites(value,8,12)
	l.Unlock()

	delay()
	l.Lock()
	value = 27
	s27.Put(216,value)
	log("p8 w:%d", value)
	countwrites(value,8,27)
	l.Unlock()

	delay()
	l.Lock()
	value = 32
	s32.Put(254,value)
	log("p8 w:%d", value)
	countwrites(value,8,32)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s21.QueryP(164,&value)
	log("p8 r:%d", value)
	countreads(value,8,21)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s24.QueryP(186,&value)
	log("p8 r:%d", value)
	countreads(value,8,24)
	l.Unlock()

	delay()
	l.Lock()
	value = 5
	s5.Put(34,value)
	log("p8 w:%d", value)
	countwrites(value,8,5)
	l.Unlock()

	delay()
	l.Lock()
	value = 25
	s25.Put(198,value)
	log("p8 w:%d", value)
	countwrites(value,8,25)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s16.QueryP(128,&value)
	log("p8 r:%d", value)
	countreads(value,8,16)
	l.Unlock()

	delay()
	l.Lock()
	value = 5
	s5.Put(37,value)
	log("p8 w:%d", value)
	countwrites(value,8,5)
	l.Unlock()

	delay()
	l.Lock()
	value = 19
	s19.Put(149,value)
	log("p8 w:%d", value)
	countwrites(value,8,19)
	l.Unlock()

	delay()
	l.Lock()
	value = 3
	s3.Put(21,value)
	log("p8 w:%d", value)
	countwrites(value,8,3)
	l.Unlock()
}

func p9() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 25
	s25.Put(196,value)
	log("p9 w:%d", value)
	countwrites(value,9,25)
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
	value = 14
	s14.Put(111,value)
	log("p9 w:%d", value)
	countwrites(value,9,14)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s21.QueryP(165,&value)
	log("p9 r:%d", value)
	countreads(value,9,21)
	l.Unlock()

	delay()
	l.Lock()
	value = 31
	s31.Put(246,value)
	log("p9 w:%d", value)
	countwrites(value,9,31)
	l.Unlock()

	delay()
	l.Lock()
	value = 3
	s3.Put(21,value)
	log("p9 w:%d", value)
	countwrites(value,9,3)
	l.Unlock()

	delay()
	l.Lock()
	value = 26
	s26.Put(208,value)
	log("p9 w:%d", value)
	countwrites(value,9,26)
	l.Unlock()

	delay()
	l.Lock()
	value = 18
	s18.Put(143,value)
	log("p9 w:%d", value)
	countwrites(value,9,18)
	l.Unlock()

	delay()
	l.Lock()
	value = 8
	s8.Put(64,value)
	log("p9 w:%d", value)
	countwrites(value,9,8)
	l.Unlock()

	delay()
	l.Lock()
	value = 13
	s13.Put(102,value)
	log("p9 w:%d", value)
	countwrites(value,9,13)
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
	value = 6
	s6.Put(47,value)
	log("p9 w:%d", value)
	countwrites(value,9,6)
	l.Unlock()

	delay()
	l.Lock()
	value = 13
	s13.Put(99,value)
	log("p9 w:%d", value)
	countwrites(value,9,13)
	l.Unlock()

	delay()
	l.Lock()
	value = 23
	s23.Put(181,value)
	log("p9 w:%d", value)
	countwrites(value,9,23)
	l.Unlock()

	delay()
	l.Lock()
	value = 9
	s9.Put(71,value)
	log("p9 w:%d", value)
	countwrites(value,9,9)
	l.Unlock()

	delay()
	l.Lock()
	value = 21
	s21.Put(163,value)
	log("p9 w:%d", value)
	countwrites(value,9,21)
	l.Unlock()
}


func p10() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 2
	s2.Put(15,value)
	log("p10 w:%d", value)
	countwrites(value,10,2)
	l.Unlock()

	delay()
	l.Lock()
	value = 14
	s14.Put(110,value)
	log("p10 w:%d", value)
	countwrites(value,10,14)
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
	value = 2
	s2.Put(16,value)
	log("p10 w:%d", value)
	countwrites(value,10,2)
	l.Unlock()

	delay()
	l.Lock()
	value = 11
	s11.Put(82,value)
	log("p10 w:%d", value)
	countwrites(value,10,11)
	l.Unlock()

	delay()
	l.Lock()
	value = 26
	s26.Put(203,value)
	log("p10 w:%d", value)
	countwrites(value,10,26)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s31.QueryP(241,&value)
	log("p10 r:%d", value)
	countreads(value,10,31)
	l.Unlock()

	delay()
	l.Lock()
	value = 25
	s25.Put(199,value)
	log("p10 w:%d", value)
	countwrites(value,10,25)
	l.Unlock()

	delay()
	l.Lock()
	value = 6
	s6.Put(41,value)
	log("p10 w:%d", value)
	countwrites(value,10,6)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s13.QueryP(102,&value)
	log("p10 r:%d", value)
	countreads(value,10,13)
	l.Unlock()

	delay()
	l.Lock()
	value = 5
	s5.Put(38,value)
	log("p10 w:%d", value)
	countwrites(value,10,5)
	l.Unlock()

	delay()
	l.Lock()
	value = 17
	s17.Put(131,value)
	log("p10 w:%d", value)
	countwrites(value,10,17)
	l.Unlock()

	delay()
	l.Lock()
	value = 22
	s22.Put(170,value)
	log("p10 w:%d", value)
	countwrites(value,10,22)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s6.QueryP(43,&value)
	log("p10 r:%d", value)
	countreads(value,10,6)
	l.Unlock()

	delay()
	l.Lock()
	value = 31
	s31.Put(242,value)
	log("p10 w:%d", value)
	countwrites(value,10,31)
	l.Unlock()

	delay()
	l.Lock()
	value = 7
	s7.Put(52,value)
	log("p10 w:%d", value)
	countwrites(value,10,7)
	l.Unlock()
}


func p11() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	s3.QueryP(23,&value)
	log("p11 r:%d", value)
	countreads(value,11,3)
	l.Unlock()

	delay()
	l.Lock()
	value = 4
	s4.Put(30,value)
	log("p11 w:%d", value)
	countwrites(value,11,4)
	l.Unlock()

	delay()
	l.Lock()
	value = 8
	s8.Put(59,value)
	log("p11 w:%d", value)
	countwrites(value,11,8)
	l.Unlock()

	delay()
	l.Lock()
	value = 22
	s22.Put(175,value)
	log("p11 w:%d", value)
	countwrites(value,11,22)
	l.Unlock()

	delay()
	l.Lock()
	value = 17
	s17.Put(130,value)
	log("p11 w:%d", value)
	countwrites(value,11,17)
	l.Unlock()

	delay()
	l.Lock()
	value = 26
	s26.Put(206,value)
	log("p11 w:%d", value)
	countwrites(value,11,26)
	l.Unlock()

	delay()
	l.Lock()
	value = 15
	s15.Put(115,value)
	log("p11 w:%d", value)
	countwrites(value,11,15)
	l.Unlock()

	delay()
	l.Lock()
	value = 22
	s22.Put(169,value)
	log("p11 w:%d", value)
	countwrites(value,11,22)
	l.Unlock()

	delay()
	l.Lock()
	value = 17
	s17.Put(130,value)
	log("p11 w:%d", value)
	countwrites(value,11,17)
	l.Unlock()

	delay()
	l.Lock()
	value = 24
	s24.Put(188,value)
	log("p11 w:%d", value)
	countwrites(value,11,24)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s8.QueryP(59,&value)
	log("p11 r:%d", value)
	countreads(value,11,8)
	l.Unlock()

	delay()
	l.Lock()
	value = 5
	s5.Put(37,value)
	log("p11 w:%d", value)
	countwrites(value,11,5)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s18.QueryP(144,&value)
	log("p11 r:%d", value)
	countreads(value,11,18)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s20.QueryP(158,&value)
	log("p11 r:%d", value)
	countreads(value,11,20)
	l.Unlock()

	delay()
	l.Lock()
	value = 7
	s7.Put(49,value)
	log("p11 w:%d", value)
	countwrites(value,11,7)
	l.Unlock()

	delay()
	l.Lock()
	value = 25
	s25.Put(198,value)
	log("p11 w:%d", value)
	countwrites(value,11,25)
	l.Unlock()
}


func p12() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 13
	s13.Put(100,value)
	log("p12 w:%d", value)
	countwrites(value,12,13)
	l.Unlock()

	delay()
	l.Lock()
	value = 26
	s26.Put(202,value)
	log("p12 w:%d", value)
	countwrites(value,12,26)
	l.Unlock()

	delay()
	l.Lock()
	value = 2
	s2.Put(16,value)
	log("p12 w:%d", value)
	countwrites(value,12,2)
	l.Unlock()

	delay()
	l.Lock()
	value = 14
	s14.Put(111,value)
	log("p12 w:%d", value)
	countwrites(value,12,14)
	l.Unlock()

	delay()
	l.Lock()
	value = 28
	s28.Put(221,value)
	log("p12 w:%d", value)
	countwrites(value,12,28)
	l.Unlock()

	delay()
	l.Lock()
	value = 18
	s18.Put(138,value)
	log("p12 w:%d", value)
	countwrites(value,12,18)
	l.Unlock()

	delay()
	l.Lock()
	value = 10
	s10.Put(77,value)
	log("p12 w:%d", value)
	countwrites(value,12,10)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s7.QueryP(49,&value)
	log("p12 r:%d", value)
	countreads(value,12,7)
	l.Unlock()

	delay()
	l.Lock()
	value = 26
	s26.Put(204,value)
	log("p12 w:%d", value)
	countwrites(value,12,26)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s11.QueryP(87,&value)
	log("p12 r:%d", value)
	countreads(value,12,11)
	l.Unlock()

	delay()
	l.Lock()
	value = 1
	s1.Put(4,value)
	log("p12 w:%d", value)
	countwrites(value,12,1)
	l.Unlock()

	delay()
	l.Lock()
	value = 21
	s21.Put(163,value)
	log("p12 w:%d", value)
	countwrites(value,12,21)
	l.Unlock()

	delay()
	l.Lock()
	value = 11
	s11.Put(82,value)
	log("p12 w:%d", value)
	countwrites(value,12,11)
	l.Unlock()

	delay()
	l.Lock()
	value = 26
	s26.Put(203,value)
	log("p12 w:%d", value)
	countwrites(value,12,26)
	l.Unlock()

	delay()
	l.Lock()
	value = 17
	s17.Put(136,value)
	log("p12 w:%d", value)
	countwrites(value,12,17)
	l.Unlock()

	delay()
	l.Lock()
	value = 31
	s31.Put(241,value)
	log("p12 w:%d", value)
	countwrites(value,12,31)
	l.Unlock()
}


func p13() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 1
	s1.Put(2,value)
	log("p13 w:%d", value)
	countwrites(value,13,1)
	l.Unlock()

	delay()
	l.Lock()
	value = 10
	s10.Put(78,value)
	log("p13 w:%d", value)
	countwrites(value,13,10)
	l.Unlock()

	delay()
	l.Lock()
	value = 21
	s21.Put(161,value)
	log("p13 w:%d", value)
	countwrites(value,13,21)
	l.Unlock()

	delay()
	l.Lock()
	value = 29
	s29.Put(226,value)
	log("p13 w:%d", value)
	countwrites(value,13,29)
	l.Unlock()

	delay()
	l.Lock()
	value = 6
	s6.Put(42,value)
	log("p13 w:%d", value)
	countwrites(value,13,6)
	l.Unlock()

	delay()
	l.Lock()
	value = 6
	s6.Put(47,value)
	log("p13 w:%d", value)
	countwrites(value,13,6)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s15.QueryP(113,&value)
	log("p13 r:%d", value)
	countreads(value,13,15)
	l.Unlock()

	delay()
	l.Lock()
	value = 14
	s14.Put(108,value)
	log("p13 w:%d", value)
	countwrites(value,13,14)
	l.Unlock()

	delay()
	l.Lock()
	value = 29
	s29.Put(232,value)
	log("p13 w:%d", value)
	countwrites(value,13,29)
	l.Unlock()

	delay()
	l.Lock()
	value = 23
	s23.Put(182,value)
	log("p13 w:%d", value)
	countwrites(value,13,23)
	l.Unlock()

	delay()
	l.Lock()
	value = 6
	s6.Put(45,value)
	log("p13 w:%d", value)
	countwrites(value,13,6)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s11.QueryP(85,&value)
	log("p13 r:%d", value)
	countreads(value,13,11)
	l.Unlock()

	delay()
	l.Lock()
	value = 22
	s22.Put(172,value)
	log("p13 w:%d", value)
	countwrites(value,13,22)
	l.Unlock()

	delay()
	l.Lock()
	value = 9
	s9.Put(67,value)
	log("p13 w:%d", value)
	countwrites(value,13,9)
	l.Unlock()

	delay()
	l.Lock()
	value = 3
	s3.Put(18,value)
	log("p13 w:%d", value)
	countwrites(value,13,3)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s23.QueryP(182,&value)
	log("p13 r:%d", value)
	countreads(value,13,23)
	l.Unlock()
}

func p14() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 29
	s29.Put(225,value)
	log("p14 w:%d", value)
	countwrites(value,14,29)
	l.Unlock()

	delay()
	l.Lock()
	value = 10
	s10.Put(78,value)
	log("p14 w:%d", value)
	countwrites(value,14,10)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s17.QueryP(134,&value)
	log("p14 r:%d", value)
	countreads(value,14,17)
	l.Unlock()

	delay()
	l.Lock()
	value = 22
	s22.Put(169,value)
	log("p14 w:%d", value)
	countwrites(value,14,22)
	l.Unlock()

	delay()
	l.Lock()
	value = 32
	s32.Put(251,value)
	log("p14 w:%d", value)
	countwrites(value,14,32)
	l.Unlock()

	delay()
	l.Lock()
	value = 28
	s28.Put(224,value)
	log("p14 w:%d", value)
	countwrites(value,14,28)
	l.Unlock()

	delay()
	l.Lock()
	value = 17
	s17.Put(135,value)
	log("p14 w:%d", value)
	countwrites(value,14,17)
	l.Unlock()

	delay()
	l.Lock()
	value = 25
	s25.Put(198,value)
	log("p14 w:%d", value)
	countwrites(value,14,25)
	l.Unlock()

	delay()
	l.Lock()
	value = 14
	s14.Put(111,value)
	log("p14 w:%d", value)
	countwrites(value,14,14)
	l.Unlock()

	delay()
	l.Lock()
	value = 21
	s21.Put(166,value)
	log("p14 w:%d", value)
	countwrites(value,14,21)
	l.Unlock()

	delay()
	l.Lock()
	value = 16
	s16.Put(125,value)
	log("p14 w:%d", value)
	countwrites(value,14,16)
	l.Unlock()

	delay()
	l.Lock()
	value = 5
	s5.Put(34,value)
	log("p14 w:%d", value)
	countwrites(value,14,5)
	l.Unlock()

	delay()
	l.Lock()
	value = 6
	s6.Put(46,value)
	log("p14 w:%d", value)
	countwrites(value,14,6)
	l.Unlock()

	delay()
	l.Lock()
	value = 4
	s4.Put(25,value)
	log("p14 w:%d", value)
	countwrites(value,14,4)
	l.Unlock()

	delay()
	l.Lock()
	value = 29
	s29.Put(228,value)
	log("p14 w:%d", value)
	countwrites(value,14,29)
	l.Unlock()

	delay()
	l.Lock()
	value = 23
	s23.Put(178,value)
	log("p14 w:%d", value)
	countwrites(value,14,23)
	l.Unlock()
}


func p15() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 8
	s8.Put(60,value)
	log("p15 w:%d", value)
	countwrites(value,15,8)
	l.Unlock()

	delay()
	l.Lock()
	value = 5
	s5.Put(37,value)
	log("p15 w:%d", value)
	countwrites(value,15,5)
	l.Unlock()

	delay()
	l.Lock()
	value = 25
	s25.Put(196,value)
	log("p15 w:%d", value)
	countwrites(value,15,25)
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
	value = 16
	s16.Put(122,value)
	log("p15 w:%d", value)
	countwrites(value,15,16)
	l.Unlock()

	delay()
	l.Lock()
	value = 6
	s6.Put(47,value)
	log("p15 w:%d", value)
	countwrites(value,15,6)
	l.Unlock()

	delay()
	l.Lock()
	value = 29
	s29.Put(227,value)
	log("p15 w:%d", value)
	countwrites(value,15,29)
	l.Unlock()

	delay()
	l.Lock()
	value = 4
	s4.Put(32,value)
	log("p15 w:%d", value)
	countwrites(value,15,4)
	l.Unlock()

	delay()
	l.Lock()
	value = 12
	s12.Put(94,value)
	log("p15 w:%d", value)
	countwrites(value,15,12)
	l.Unlock()

	delay()
	l.Lock()
	value = 17
	s17.Put(136,value)
	log("p15 w:%d", value)
	countwrites(value,15,17)
	l.Unlock()

	delay()
	l.Lock()
	value = 14
	s14.Put(112,value)
	log("p15 w:%d", value)
	countwrites(value,15,14)
	l.Unlock()

	delay()
	l.Lock()
	value = 14
	s14.Put(108,value)
	log("p15 w:%d", value)
	countwrites(value,15,14)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s8.QueryP(63,&value)
	log("p15 r:%d", value)
	countreads(value,15,8)
	l.Unlock()

	delay()
	l.Lock()
	value = 1
	s1.Put(5,value)
	log("p15 w:%d", value)
	countwrites(value,15,1)
	l.Unlock()

	delay()
	l.Lock()
	value = 23
	s23.Put(180,value)
	log("p15 w:%d", value)
	countwrites(value,15,23)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s18.QueryP(144,&value)
	log("p15 r:%d", value)
	countreads(value,15,18)
	l.Unlock()
}


func p16() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 8
	s8.Put(61,value)
	log("p16 w:%d", value)
	countwrites(value,16,8)
	l.Unlock()

	delay()
	l.Lock()
	value = 3
	s3.Put(22,value)
	log("p16 w:%d", value)
	countwrites(value,16,3)
	l.Unlock()

	delay()
	l.Lock()
	value = 10
	s10.Put(76,value)
	log("p16 w:%d", value)
	countwrites(value,16,10)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s6.QueryP(48,&value)
	log("p16 r:%d", value)
	countreads(value,16,6)
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
	s29.Put(226,value)
	log("p16 w:%d", value)
	countwrites(value,16,29)
	l.Unlock()

	delay()
	l.Lock()
	value = 27
	s27.Put(214,value)
	log("p16 w:%d", value)
	countwrites(value,16,27)
	l.Unlock()

	delay()
	l.Lock()
	value = 25
	s25.Put(197,value)
	log("p16 w:%d", value)
	countwrites(value,16,25)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s4.QueryP(28,&value)
	log("p16 r:%d", value)
	countreads(value,16,4)
	l.Unlock()

	delay()
	l.Lock()
	value = 27
	s27.Put(213,value)
	log("p16 w:%d", value)
	countwrites(value,16,27)
	l.Unlock()

	delay()
	l.Lock()
	value = 27
	s27.Put(215,value)
	log("p16 w:%d", value)
	countwrites(value,16,27)
	l.Unlock()

	delay()
	l.Lock()
	value = 3
	s3.Put(19,value)
	log("p16 w:%d", value)
	countwrites(value,16,3)
	l.Unlock()

	delay()
	l.Lock()
	value = 9
	s9.Put(70,value)
	log("p16 w:%d", value)
	countwrites(value,16,9)
	l.Unlock()

	delay()
	l.Lock()
	value = 14
	s14.Put(107,value)
	log("p16 w:%d", value)
	countwrites(value,16,14)
	l.Unlock()

	delay()
	l.Lock()
	value = 25
	s25.Put(199,value)
	log("p16 w:%d", value)
	countwrites(value,16,25)
	l.Unlock()

	delay()
	l.Lock()
	value = 12
	s12.Put(94,value)
	log("p16 w:%d", value)
	countwrites(value,16,12)
	l.Unlock()
}


func p17() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 2
	s2.Put(14,value)
	log("p17 w:%d", value)
	countwrites(value,17,2)
	l.Unlock()

	delay()
	l.Lock()
	value = 5
	s5.Put(33,value)
	log("p17 w:%d", value)
	countwrites(value,17,5)
	l.Unlock()

	delay()
	l.Lock()
	value = 27
	s27.Put(214,value)
	log("p17 w:%d", value)
	countwrites(value,17,27)
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
	value = 15
	s15.Put(118,value)
	log("p17 w:%d", value)
	countwrites(value,17,15)
	l.Unlock()

	delay()
	l.Lock()
	value = 23
	s23.Put(182,value)
	log("p17 w:%d", value)
	countwrites(value,17,23)
	l.Unlock()

	delay()
	l.Lock()
	value = 25
	s25.Put(198,value)
	log("p17 w:%d", value)
	countwrites(value,17,25)
	l.Unlock()

	delay()
	l.Lock()
	value = 26
	s26.Put(201,value)
	log("p17 w:%d", value)
	countwrites(value,17,26)
	l.Unlock()

	delay()
	l.Lock()
	value = 24
	s24.Put(190,value)
	log("p17 w:%d", value)
	countwrites(value,17,24)
	l.Unlock()

	delay()
	l.Lock()
	value = 25
	s25.Put(196,value)
	log("p17 w:%d", value)
	countwrites(value,17,25)
	l.Unlock()

	delay()
	l.Lock()
	value = 28
	s28.Put(217,value)
	log("p17 w:%d", value)
	countwrites(value,17,28)
	l.Unlock()

	delay()
	l.Lock()
	value = 8
	s8.Put(64,value)
	log("p17 w:%d", value)
	countwrites(value,17,8)
	l.Unlock()

	delay()
	l.Lock()
	value = 29
	s29.Put(232,value)
	log("p17 w:%d", value)
	countwrites(value,17,29)
	l.Unlock()

	delay()
	l.Lock()
	value = 12
	s12.Put(89,value)
	log("p17 w:%d", value)
	countwrites(value,17,12)
	l.Unlock()

	delay()
	l.Lock()
	value = 24
	s24.Put(191,value)
	log("p17 w:%d", value)
	countwrites(value,17,24)
	l.Unlock()

	delay()
	l.Lock()
	value = 29
	s29.Put(229,value)
	log("p17 w:%d", value)
	countwrites(value,17,29)
	l.Unlock()
}


func p18() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 26
	s26.Put(204,value)
	log("p18 w:%d", value)
	countwrites(value,18,26)
	l.Unlock()

	delay()
	l.Lock()
	value = 19
	s19.Put(150,value)
	log("p18 w:%d", value)
	countwrites(value,18,19)
	l.Unlock()

	delay()
	l.Lock()
	value = 15
	s15.Put(117,value)
	log("p18 w:%d", value)
	countwrites(value,18,15)
	l.Unlock()

	delay()
	l.Lock()
	value = 1
	s1.Put(3,value)
	log("p18 w:%d", value)
	countwrites(value,18,1)
	l.Unlock()

	delay()
	l.Lock()
	value = 18
	s18.Put(140,value)
	log("p18 w:%d", value)
	countwrites(value,18,18)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s2.QueryP(12,&value)
	log("p18 r:%d", value)
	countreads(value,18,2)
	l.Unlock()

	delay()
	l.Lock()
	value = 2
	s2.Put(11,value)
	log("p18 w:%d", value)
	countwrites(value,18,2)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s21.QueryP(166,&value)
	log("p18 r:%d", value)
	countreads(value,18,21)
	l.Unlock()

	delay()
	l.Lock()
	value = 23
	s23.Put(184,value)
	log("p18 w:%d", value)
	countwrites(value,18,23)
	l.Unlock()

	delay()
	l.Lock()
	value = 28
	s28.Put(218,value)
	log("p18 w:%d", value)
	countwrites(value,18,28)
	l.Unlock()

	delay()
	l.Lock()
	value = 23
	s23.Put(178,value)
	log("p18 w:%d", value)
	countwrites(value,18,23)
	l.Unlock()

	delay()
	l.Lock()
	value = 28
	s28.Put(221,value)
	log("p18 w:%d", value)
	countwrites(value,18,28)
	l.Unlock()

	delay()
	l.Lock()
	value = 30
	s30.Put(239,value)
	log("p18 w:%d", value)
	countwrites(value,18,30)
	l.Unlock()

	delay()
	l.Lock()
	value = 21
	s21.Put(165,value)
	log("p18 w:%d", value)
	countwrites(value,18,21)
	l.Unlock()

	delay()
	l.Lock()
	value = 24
	s24.Put(185,value)
	log("p18 w:%d", value)
	countwrites(value,18,24)
	l.Unlock()

	delay()
	l.Lock()
	value = 7
	s7.Put(50,value)
	log("p18 w:%d", value)
	countwrites(value,18,7)
	l.Unlock()
}


func p19() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 15
	s15.Put(115,value)
	log("p19 w:%d", value)
	countwrites(value,19,15)
	l.Unlock()

	delay()
	l.Lock()
	value = 23
	s23.Put(177,value)
	log("p19 w:%d", value)
	countwrites(value,19,23)
	l.Unlock()

	delay()
	l.Lock()
	value = 13
	s13.Put(103,value)
	log("p19 w:%d", value)
	countwrites(value,19,13)
	l.Unlock()

	delay()
	l.Lock()
	value = 29
	s29.Put(232,value)
	log("p19 w:%d", value)
	countwrites(value,19,29)
	l.Unlock()

	delay()
	l.Lock()
	value = 21
	s21.Put(168,value)
	log("p19 w:%d", value)
	countwrites(value,19,21)
	l.Unlock()

	delay()
	l.Lock()
	value = 19
	s19.Put(147,value)
	log("p19 w:%d", value)
	countwrites(value,19,19)
	l.Unlock()

	delay()
	l.Lock()
	value = 22
	s22.Put(176,value)
	log("p19 w:%d", value)
	countwrites(value,19,22)
	l.Unlock()

	delay()
	l.Lock()
	value = 26
	s26.Put(202,value)
	log("p19 w:%d", value)
	countwrites(value,19,26)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s22.QueryP(172,&value)
	log("p19 r:%d", value)
	countreads(value,19,22)
	l.Unlock()

	delay()
	l.Lock()
	value = 28
	s28.Put(224,value)
	log("p19 w:%d", value)
	countwrites(value,19,28)
	l.Unlock()

	delay()
	l.Lock()
	value = 26
	s26.Put(206,value)
	log("p19 w:%d", value)
	countwrites(value,19,26)
	l.Unlock()

	delay()
	l.Lock()
	value = 14
	s14.Put(107,value)
	log("p19 w:%d", value)
	countwrites(value,19,14)
	l.Unlock()

	delay()
	l.Lock()
	value = 25
	s25.Put(195,value)
	log("p19 w:%d", value)
	countwrites(value,19,25)
	l.Unlock()

	delay()
	l.Lock()
	value = 11
	s11.Put(83,value)
	log("p19 w:%d", value)
	countwrites(value,19,11)
	l.Unlock()

	delay()
	l.Lock()
	value = 21
	s21.Put(167,value)
	log("p19 w:%d", value)
	countwrites(value,19,21)
	l.Unlock()

	delay()
	l.Lock()
	value = 25
	s25.Put(195,value)
	log("p19 w:%d", value)
	countwrites(value,19,25)
	l.Unlock()
}


func p20() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 28
	s28.Put(219,value)
	log("p20 w:%d", value)
	countwrites(value,20,28)
	l.Unlock()

	delay()
	l.Lock()
	value = 30
	s30.Put(238,value)
	log("p20 w:%d", value)
	countwrites(value,20,30)
	l.Unlock()

	delay()
	l.Lock()
	value = 22
	s22.Put(174,value)
	log("p20 w:%d", value)
	countwrites(value,20,22)
	l.Unlock()

	delay()
	l.Lock()
	value = 27
	s27.Put(214,value)
	log("p20 w:%d", value)
	countwrites(value,20,27)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s19.QueryP(152,&value)
	log("p20 r:%d", value)
	countreads(value,20,19)
	l.Unlock()

	delay()
	l.Lock()
	value = 14
	s14.Put(105,value)
	log("p20 w:%d", value)
	countwrites(value,20,14)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s3.QueryP(18,&value)
	log("p20 r:%d", value)
	countreads(value,20,3)
	l.Unlock()

	delay()
	l.Lock()
	value = 17
	s17.Put(131,value)
	log("p20 w:%d", value)
	countwrites(value,20,17)
	l.Unlock()

	delay()
	l.Lock()
	value = 2
	s2.Put(11,value)
	log("p20 w:%d", value)
	countwrites(value,20,2)
	l.Unlock()

	delay()
	l.Lock()
	value = 22
	s22.Put(174,value)
	log("p20 w:%d", value)
	countwrites(value,20,22)
	l.Unlock()

	delay()
	l.Lock()
	value = 27
	s27.Put(216,value)
	log("p20 w:%d", value)
	countwrites(value,20,27)
	l.Unlock()

	delay()
	l.Lock()
	value = 25
	s25.Put(195,value)
	log("p20 w:%d", value)
	countwrites(value,20,25)
	l.Unlock()

	delay()
	l.Lock()
	value = 22
	s22.Put(176,value)
	log("p20 w:%d", value)
	countwrites(value,20,22)
	l.Unlock()

	delay()
	l.Lock()
	value = 23
	s23.Put(181,value)
	log("p20 w:%d", value)
	countwrites(value,20,23)
	l.Unlock()

	delay()
	l.Lock()
	value = 2
	s2.Put(12,value)
	log("p20 w:%d", value)
	countwrites(value,20,2)
	l.Unlock()

	delay()
	l.Lock()
	value = 17
	s17.Put(133,value)
	log("p20 w:%d", value)
	countwrites(value,20,17)
	l.Unlock()
}


func p21() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 6
	s6.Put(41,value)
	log("p21 w:%d", value)
	countwrites(value,21,6)
	l.Unlock()

	delay()
	l.Lock()
	value = 9
	s9.Put(65,value)
	log("p21 w:%d", value)
	countwrites(value,21,9)
	l.Unlock()

	delay()
	l.Lock()
	value = 14
	s14.Put(111,value)
	log("p21 w:%d", value)
	countwrites(value,21,14)
	l.Unlock()

	delay()
	l.Lock()
	value = 12
	s12.Put(93,value)
	log("p21 w:%d", value)
	countwrites(value,21,12)
	l.Unlock()

	delay()
	l.Lock()
	value = 26
	s26.Put(204,value)
	log("p21 w:%d", value)
	countwrites(value,21,26)
	l.Unlock()

	delay()
	l.Lock()
	value = 5
	s5.Put(37,value)
	log("p21 w:%d", value)
	countwrites(value,21,5)
	l.Unlock()

	delay()
	l.Lock()
	value = 19
	s19.Put(149,value)
	log("p21 w:%d", value)
	countwrites(value,21,19)
	l.Unlock()

	delay()
	l.Lock()
	value = 23
	s23.Put(183,value)
	log("p21 w:%d", value)
	countwrites(value,21,23)
	l.Unlock()

	delay()
	l.Lock()
	value = 7
	s7.Put(53,value)
	log("p21 w:%d", value)
	countwrites(value,21,7)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s4.QueryP(25,&value)
	log("p21 r:%d", value)
	countreads(value,21,4)
	l.Unlock()

	delay()
	l.Lock()
	value = 15
	s15.Put(114,value)
	log("p21 w:%d", value)
	countwrites(value,21,15)
	l.Unlock()

	delay()
	l.Lock()
	value = 6
	s6.Put(46,value)
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
	value = 8
	s8.Put(63,value)
	log("p21 w:%d", value)
	countwrites(value,21,8)
	l.Unlock()

	delay()
	l.Lock()
	value = 3
	s3.Put(21,value)
	log("p21 w:%d", value)
	countwrites(value,21,3)
	l.Unlock()

	delay()
	l.Lock()
	value = 7
	s7.Put(55,value)
	log("p21 w:%d", value)
	countwrites(value,21,7)
	l.Unlock()
}


func p22() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 2
	s2.Put(13,value)
	log("p22 w:%d", value)
	countwrites(value,22,2)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s27.QueryP(213,&value)
	log("p22 r:%d", value)
	countreads(value,22,27)
	l.Unlock()

	delay()
	l.Lock()
	value = 9
	s9.Put(67,value)
	log("p22 w:%d", value)
	countwrites(value,22,9)
	l.Unlock()

	delay()
	l.Lock()
	value = 6
	s6.Put(41,value)
	log("p22 w:%d", value)
	countwrites(value,22,6)
	l.Unlock()

	delay()
	l.Lock()
	value = 17
	s17.Put(135,value)
	log("p22 w:%d", value)
	countwrites(value,22,17)
	l.Unlock()

	delay()
	l.Lock()
	value = 25
	s25.Put(194,value)
	log("p22 w:%d", value)
	countwrites(value,22,25)
	l.Unlock()

	delay()
	l.Lock()
	value = 9
	s9.Put(70,value)
	log("p22 w:%d", value)
	countwrites(value,22,9)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s29.QueryP(231,&value)
	log("p22 r:%d", value)
	countreads(value,22,29)
	l.Unlock()

	delay()
	l.Lock()
	value = 4
	s4.Put(28,value)
	log("p22 w:%d", value)
	countwrites(value,22,4)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s5.QueryP(33,&value)
	log("p22 r:%d", value)
	countreads(value,22,5)
	l.Unlock()

	delay()
	l.Lock()
	value = 21
	s21.Put(164,value)
	log("p22 w:%d", value)
	countwrites(value,22,21)
	l.Unlock()

	delay()
	l.Lock()
	value = 22
	s22.Put(173,value)
	log("p22 w:%d", value)
	countwrites(value,22,22)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s1.QueryP(6,&value)
	log("p22 r:%d", value)
	countreads(value,22,1)
	l.Unlock()

	delay()
	l.Lock()
	value = 12
	s12.Put(93,value)
	log("p22 w:%d", value)
	countwrites(value,22,12)
	l.Unlock()

	delay()
	l.Lock()
	value = 22
	s22.Put(171,value)
	log("p22 w:%d", value)
	countwrites(value,22,22)
	l.Unlock()

	delay()
	l.Lock()
	value = 3
	s3.Put(17,value)
	log("p22 w:%d", value)
	countwrites(value,22,3)
	l.Unlock()
}


func p23() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 8
	s8.Put(61,value)
	log("p23 w:%d", value)
	countwrites(value,23,8)
	l.Unlock()

	delay()
	l.Lock()
	value = 6
	s6.Put(44,value)
	log("p23 w:%d", value)
	countwrites(value,23,6)
	l.Unlock()

	delay()
	l.Lock()
	value = 1
	s1.Put(6,value)
	log("p23 w:%d", value)
	countwrites(value,23,1)
	l.Unlock()

	delay()
	l.Lock()
	value = 21
	s21.Put(161,value)
	log("p23 w:%d", value)
	countwrites(value,23,21)
	l.Unlock()

	delay()
	l.Lock()
	value = 11
	s11.Put(81,value)
	log("p23 w:%d", value)
	countwrites(value,23,11)
	l.Unlock()

	delay()
	l.Lock()
	value = 20
	s20.Put(158,value)
	log("p23 w:%d", value)
	countwrites(value,23,20)
	l.Unlock()

	delay()
	l.Lock()
	value = 2
	s2.Put(11,value)
	log("p23 w:%d", value)
	countwrites(value,23,2)
	l.Unlock()

	delay()
	l.Lock()
	value = 20
	s20.Put(155,value)
	log("p23 w:%d", value)
	countwrites(value,23,20)
	l.Unlock()

	delay()
	l.Lock()
	value = 22
	s22.Put(171,value)
	log("p23 w:%d", value)
	countwrites(value,23,22)
	l.Unlock()

	delay()
	l.Lock()
	value = 31
	s31.Put(244,value)
	log("p23 w:%d", value)
	countwrites(value,23,31)
	l.Unlock()

	delay()
	l.Lock()
	value = 4
	s4.Put(27,value)
	log("p23 w:%d", value)
	countwrites(value,23,4)
	l.Unlock()

	delay()
	l.Lock()
	value = 11
	s11.Put(87,value)
	log("p23 w:%d", value)
	countwrites(value,23,11)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s20.QueryP(158,&value)
	log("p23 r:%d", value)
	countreads(value,23,20)
	l.Unlock()

	delay()
	l.Lock()
	value = 13
	s13.Put(100,value)
	log("p23 w:%d", value)
	countwrites(value,23,13)
	l.Unlock()

	delay()
	l.Lock()
	value = 2
	s2.Put(10,value)
	log("p23 w:%d", value)
	countwrites(value,23,2)
	l.Unlock()

	delay()
	l.Lock()
	value = 28
	s28.Put(218,value)
	log("p23 w:%d", value)
	countwrites(value,23,28)
	l.Unlock()
}


func p24() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 5
	s5.Put(35,value)
	log("p24 w:%d", value)
	countwrites(value,24,5)
	l.Unlock()

	delay()
	l.Lock()
	value = 3
	s3.Put(19,value)
	log("p24 w:%d", value)
	countwrites(value,24,3)
	l.Unlock()

	delay()
	l.Lock()
	value = 29
	s29.Put(230,value)
	log("p24 w:%d", value)
	countwrites(value,24,29)
	l.Unlock()

	delay()
	l.Lock()
	value = 4
	s4.Put(32,value)
	log("p24 w:%d", value)
	countwrites(value,24,4)
	l.Unlock()

	delay()
	l.Lock()
	value = 5
	s5.Put(35,value)
	log("p24 w:%d", value)
	countwrites(value,24,5)
	l.Unlock()

	delay()
	l.Lock()
	value = 16
	s16.Put(122,value)
	log("p24 w:%d", value)
	countwrites(value,24,16)
	l.Unlock()

	delay()
	l.Lock()
	value = 8
	s8.Put(60,value)
	log("p24 w:%d", value)
	countwrites(value,24,8)
	l.Unlock()

	delay()
	l.Lock()
	value = 29
	s29.Put(230,value)
	log("p24 w:%d", value)
	countwrites(value,24,29)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s25.QueryP(197,&value)
	log("p24 r:%d", value)
	countreads(value,24,25)
	l.Unlock()

	delay()
	l.Lock()
	value = 23
	s23.Put(182,value)
	log("p24 w:%d", value)
	countwrites(value,24,23)
	l.Unlock()

	delay()
	l.Lock()
	value = 30
	s30.Put(234,value)
	log("p24 w:%d", value)
	countwrites(value,24,30)
	l.Unlock()

	delay()
	l.Lock()
	value = 3
	s3.Put(18,value)
	log("p24 w:%d", value)
	countwrites(value,24,3)
	l.Unlock()

	delay()
	l.Lock()
	value = 1
	s1.Put(1,value)
	log("p24 w:%d", value)
	countwrites(value,24,1)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s30.QueryP(235,&value)
	log("p24 r:%d", value)
	countreads(value,24,30)
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
	s9.Put(66,value)
	log("p24 w:%d", value)
	countwrites(value,24,9)
	l.Unlock()
}


func p25() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 12
	s12.Put(96,value)
	log("p25 w:%d", value)
	countwrites(value,25,12)
	l.Unlock()

	delay()
	l.Lock()
	value = 18
	s18.Put(138,value)
	log("p25 w:%d", value)
	countwrites(value,25,18)
	l.Unlock()

	delay()
	l.Lock()
	value = 19
	s19.Put(148,value)
	log("p25 w:%d", value)
	countwrites(value,25,19)
	l.Unlock()

	delay()
	l.Lock()
	value = 9
	s9.Put(69,value)
	log("p25 w:%d", value)
	countwrites(value,25,9)
	l.Unlock()

	delay()
	l.Lock()
	value = 20
	s20.Put(160,value)
	log("p25 w:%d", value)
	countwrites(value,25,20)
	l.Unlock()

	delay()
	l.Lock()
	value = 25
	s25.Put(198,value)
	log("p25 w:%d", value)
	countwrites(value,25,25)
	l.Unlock()

	delay()
	l.Lock()
	value = 3
	s3.Put(20,value)
	log("p25 w:%d", value)
	countwrites(value,25,3)
	l.Unlock()

	delay()
	l.Lock()
	value = 32
	s32.Put(251,value)
	log("p25 w:%d", value)
	countwrites(value,25,32)
	l.Unlock()

	delay()
	l.Lock()
	value = 26
	s26.Put(208,value)
	log("p25 w:%d", value)
	countwrites(value,25,26)
	l.Unlock()

	delay()
	l.Lock()
	value = 6
	s6.Put(42,value)
	log("p25 w:%d", value)
	countwrites(value,25,6)
	l.Unlock()

	delay()
	l.Lock()
	value = 3
	s3.Put(19,value)
	log("p25 w:%d", value)
	countwrites(value,25,3)
	l.Unlock()

	delay()
	l.Lock()
	value = 5
	s5.Put(40,value)
	log("p25 w:%d", value)
	countwrites(value,25,5)
	l.Unlock()

	delay()
	l.Lock()
	value = 24
	s24.Put(190,value)
	log("p25 w:%d", value)
	countwrites(value,25,24)
	l.Unlock()

	delay()
	l.Lock()
	value = 4
	s4.Put(29,value)
	log("p25 w:%d", value)
	countwrites(value,25,4)
	l.Unlock()

	delay()
	l.Lock()
	value = 15
	s15.Put(113,value)
	log("p25 w:%d", value)
	countwrites(value,25,15)
	l.Unlock()

	delay()
	l.Lock()
	value = 10
	s10.Put(73,value)
	log("p25 w:%d", value)
	countwrites(value,25,10)
	l.Unlock()
}


func p26() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 17
	s17.Put(131,value)
	log("p26 w:%d", value)
	countwrites(value,26,17)
	l.Unlock()

	delay()
	l.Lock()
	value = 30
	s30.Put(236,value)
	log("p26 w:%d", value)
	countwrites(value,26,30)
	l.Unlock()

	delay()
	l.Lock()
	value = 10
	s10.Put(79,value)
	log("p26 w:%d", value)
	countwrites(value,26,10)
	l.Unlock()

	delay()
	l.Lock()
	value = 31
	s31.Put(242,value)
	log("p26 w:%d", value)
	countwrites(value,26,31)
	l.Unlock()

	delay()
	l.Lock()
	value = 24
	s24.Put(186,value)
	log("p26 w:%d", value)
	countwrites(value,26,24)
	l.Unlock()

	delay()
	l.Lock()
	value = 14
	s14.Put(110,value)
	log("p26 w:%d", value)
	countwrites(value,26,14)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s14.QueryP(111,&value)
	log("p26 r:%d", value)
	countreads(value,26,14)
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
	value = 2
	s2.Put(10,value)
	log("p26 w:%d", value)
	countwrites(value,26,2)
	l.Unlock()

	delay()
	l.Lock()
	value = 20
	s20.Put(159,value)
	log("p26 w:%d", value)
	countwrites(value,26,20)
	l.Unlock()

	delay()
	l.Lock()
	value = 31
	s31.Put(248,value)
	log("p26 w:%d", value)
	countwrites(value,26,31)
	l.Unlock()

	delay()
	l.Lock()
	value = 23
	s23.Put(182,value)
	log("p26 w:%d", value)
	countwrites(value,26,23)
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
	s27.Put(209,value)
	log("p26 w:%d", value)
	countwrites(value,26,27)
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
	value = 18
	s18.Put(139,value)
	log("p26 w:%d", value)
	countwrites(value,26,18)
	l.Unlock()
}


func p27() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 20
	s20.Put(159,value)
	log("p27 w:%d", value)
	countwrites(value,27,20)
	l.Unlock()

	delay()
	l.Lock()
	value = 24
	s24.Put(185,value)
	log("p27 w:%d", value)
	countwrites(value,27,24)
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
	s32.Put(254,value)
	log("p27 w:%d", value)
	countwrites(value,27,32)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s9.QueryP(70,&value)
	log("p27 r:%d", value)
	countreads(value,27,9)
	l.Unlock()

	delay()
	l.Lock()
	value = 27
	s27.Put(209,value)
	log("p27 w:%d", value)
	countwrites(value,27,27)
	l.Unlock()

	delay()
	l.Lock()
	value = 30
	s30.Put(233,value)
	log("p27 w:%d", value)
	countwrites(value,27,30)
	l.Unlock()

	delay()
	l.Lock()
	value = 3
	s3.Put(20,value)
	log("p27 w:%d", value)
	countwrites(value,27,3)
	l.Unlock()

	delay()
	l.Lock()
	value = 27
	s27.Put(210,value)
	log("p27 w:%d", value)
	countwrites(value,27,27)
	l.Unlock()

	delay()
	l.Lock()
	value = 8
	s8.Put(64,value)
	log("p27 w:%d", value)
	countwrites(value,27,8)
	l.Unlock()

	delay()
	l.Lock()
	value = 17
	s17.Put(130,value)
	log("p27 w:%d", value)
	countwrites(value,27,17)
	l.Unlock()

	delay()
	l.Lock()
	value = 14
	s14.Put(106,value)
	log("p27 w:%d", value)
	countwrites(value,27,14)
	l.Unlock()

	delay()
	l.Lock()
	value = 20
	s20.Put(156,value)
	log("p27 w:%d", value)
	countwrites(value,27,20)
	l.Unlock()

	delay()
	l.Lock()
	value = 19
	s19.Put(150,value)
	log("p27 w:%d", value)
	countwrites(value,27,19)
	l.Unlock()

	delay()
	l.Lock()
	value = 31
	s31.Put(243,value)
	log("p27 w:%d", value)
	countwrites(value,27,31)
	l.Unlock()

	delay()
	l.Lock()
	value = 20
	s20.Put(155,value)
	log("p27 w:%d", value)
	countwrites(value,27,20)
	l.Unlock()
}


func p28() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 7
	s7.Put(55,value)
	log("p28 w:%d", value)
	countwrites(value,28,7)
	l.Unlock()

	delay()
	l.Lock()
	value = 21
	s21.Put(167,value)
	log("p28 w:%d", value)
	countwrites(value,28,21)
	l.Unlock()

	delay()
	l.Lock()
	value = 3
	s3.Put(21,value)
	log("p28 w:%d", value)
	countwrites(value,28,3)
	l.Unlock()

	delay()
	l.Lock()
	value = 29
	s29.Put(226,value)
	log("p28 w:%d", value)
	countwrites(value,28,29)
	l.Unlock()

	delay()
	l.Lock()
	value = 27
	s27.Put(210,value)
	log("p28 w:%d", value)
	countwrites(value,28,27)
	l.Unlock()

	delay()
	l.Lock()
	value = 29
	s29.Put(229,value)
	log("p28 w:%d", value)
	countwrites(value,28,29)
	l.Unlock()

	delay()
	l.Lock()
	value = 18
	s18.Put(138,value)
	log("p28 w:%d", value)
	countwrites(value,28,18)
	l.Unlock()

	delay()
	l.Lock()
	value = 12
	s12.Put(95,value)
	log("p28 w:%d", value)
	countwrites(value,28,12)
	l.Unlock()

	delay()
	l.Lock()
	value = 18
	s18.Put(137,value)
	log("p28 w:%d", value)
	countwrites(value,28,18)
	l.Unlock()

	delay()
	l.Lock()
	value = 14
	s14.Put(111,value)
	log("p28 w:%d", value)
	countwrites(value,28,14)
	l.Unlock()

	delay()
	l.Lock()
	value = 8
	s8.Put(61,value)
	log("p28 w:%d", value)
	countwrites(value,28,8)
	l.Unlock()

	delay()
	l.Lock()
	value = 23
	s23.Put(178,value)
	log("p28 w:%d", value)
	countwrites(value,28,23)
	l.Unlock()

	delay()
	l.Lock()
	value = 28
	s28.Put(219,value)
	log("p28 w:%d", value)
	countwrites(value,28,28)
	l.Unlock()

	delay()
	l.Lock()
	value = 5
	s5.Put(33,value)
	log("p28 w:%d", value)
	countwrites(value,28,5)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s31.QueryP(241,&value)
	log("p28 r:%d", value)
	countreads(value,28,31)
	l.Unlock()

	delay()
	l.Lock()
	value = 6
	s6.Put(46,value)
	log("p28 w:%d", value)
	countwrites(value,28,6)
	l.Unlock()
}


func p29() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 19
	s19.Put(150,value)
	log("p29 w:%d", value)
	countwrites(value,29,19)
	l.Unlock()

	delay()
	l.Lock()
	value = 6
	s6.Put(41,value)
	log("p29 w:%d", value)
	countwrites(value,29,6)
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
	value = 8
	s8.Put(62,value)
	log("p29 w:%d", value)
	countwrites(value,29,8)
	l.Unlock()

	delay()
	l.Lock()
	value = 10
	s10.Put(73,value)
	log("p29 w:%d", value)
	countwrites(value,29,10)
	l.Unlock()

	delay()
	l.Lock()
	value = 5
	s5.Put(39,value)
	log("p29 w:%d", value)
	countwrites(value,29,5)
	l.Unlock()

	delay()
	l.Lock()
	value = 25
	s25.Put(197,value)
	log("p29 w:%d", value)
	countwrites(value,29,25)
	l.Unlock()

	delay()
	l.Lock()
	value = 27
	s27.Put(216,value)
	log("p29 w:%d", value)
	countwrites(value,29,27)
	l.Unlock()

	delay()
	l.Lock()
	value = 5
	s5.Put(37,value)
	log("p29 w:%d", value)
	countwrites(value,29,5)
	l.Unlock()

	delay()
	l.Lock()
	value = 19
	s19.Put(149,value)
	log("p29 w:%d", value)
	countwrites(value,29,19)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s8.QueryP(64,&value)
	log("p29 r:%d", value)
	countreads(value,29,8)
	l.Unlock()

	delay()
	l.Lock()
	value = 18
	s18.Put(142,value)
	log("p29 w:%d", value)
	countwrites(value,29,18)
	l.Unlock()

	delay()
	l.Lock()
	value = 10
	s10.Put(78,value)
	log("p29 w:%d", value)
	countwrites(value,29,10)
	l.Unlock()

	delay()
	l.Lock()
	value = 6
	s6.Put(46,value)
	log("p29 w:%d", value)
	countwrites(value,29,6)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s7.QueryP(51,&value)
	log("p29 r:%d", value)
	countreads(value,29,7)
	l.Unlock()

	delay()
	l.Lock()
	value = 4
	s4.Put(26,value)
	log("p29 w:%d", value)
	countwrites(value,29,4)
	l.Unlock()
}


func p30() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 23
	s23.Put(177,value)
	log("p30 w:%d", value)
	countwrites(value,30,23)
	l.Unlock()

	delay()
	l.Lock()
	value = 20
	s20.Put(157,value)
	log("p30 w:%d", value)
	countwrites(value,30,20)
	l.Unlock()

	delay()
	l.Lock()
	value = 7
	s7.Put(51,value)
	log("p30 w:%d", value)
	countwrites(value,30,7)
	l.Unlock()

	delay()
	l.Lock()
	value = 3
	s3.Put(17,value)
	log("p30 w:%d", value)
	countwrites(value,30,3)
	l.Unlock()

	delay()
	l.Lock()
	value = 10
	s10.Put(77,value)
	log("p30 w:%d", value)
	countwrites(value,30,10)
	l.Unlock()

	delay()
	l.Lock()
	value = 28
	s28.Put(222,value)
	log("p30 w:%d", value)
	countwrites(value,30,28)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s25.QueryP(200,&value)
	log("p30 r:%d", value)
	countreads(value,30,25)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s30.QueryP(233,&value)
	log("p30 r:%d", value)
	countreads(value,30,30)
	l.Unlock()

	delay()
	l.Lock()
	value = 13
	s13.Put(100,value)
	log("p30 w:%d", value)
	countwrites(value,30,13)
	l.Unlock()

	delay()
	l.Lock()
	value = 7
	s7.Put(50,value)
	log("p30 w:%d", value)
	countwrites(value,30,7)
	l.Unlock()

	delay()
	l.Lock()
	value = 2
	s2.Put(12,value)
	log("p30 w:%d", value)
	countwrites(value,30,2)
	l.Unlock()

	delay()
	l.Lock()
	value = 14
	s14.Put(110,value)
	log("p30 w:%d", value)
	countwrites(value,30,14)
	l.Unlock()

	delay()
	l.Lock()
	value = 20
	s20.Put(156,value)
	log("p30 w:%d", value)
	countwrites(value,30,20)
	l.Unlock()

	delay()
	l.Lock()
	value = 24
	s24.Put(188,value)
	log("p30 w:%d", value)
	countwrites(value,30,24)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s29.QueryP(230,&value)
	log("p30 r:%d", value)
	countreads(value,30,29)
	l.Unlock()

	delay()
	l.Lock()
	value = 8
	s8.Put(60,value)
	log("p30 w:%d", value)
	countwrites(value,30,8)
	l.Unlock()
}


func p31() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 22
	s22.Put(176,value)
	log("p31 w:%d", value)
	countwrites(value,31,22)
	l.Unlock()

	delay()
	l.Lock()
	value = 30
	s30.Put(235,value)
	log("p31 w:%d", value)
	countwrites(value,31,30)
	l.Unlock()

	delay()
	l.Lock()
	value = 25
	s25.Put(197,value)
	log("p31 w:%d", value)
	countwrites(value,31,25)
	l.Unlock()

	delay()
	l.Lock()
	value = 16
	s16.Put(128,value)
	log("p31 w:%d", value)
	countwrites(value,31,16)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s23.QueryP(183,&value)
	log("p31 r:%d", value)
	countreads(value,31,23)
	l.Unlock()

	delay()
	l.Lock()
	value = 12
	s12.Put(91,value)
	log("p31 w:%d", value)
	countwrites(value,31,12)
	l.Unlock()

	delay()
	l.Lock()
	value = 14
	s14.Put(111,value)
	log("p31 w:%d", value)
	countwrites(value,31,14)
	l.Unlock()

	delay()
	l.Lock()
	value = 22
	s22.Put(170,value)
	log("p31 w:%d", value)
	countwrites(value,31,22)
	l.Unlock()

	delay()
	l.Lock()
	value = 10
	s10.Put(75,value)
	log("p31 w:%d", value)
	countwrites(value,31,10)
	l.Unlock()

	delay()
	l.Lock()
	value = 14
	s14.Put(112,value)
	log("p31 w:%d", value)
	countwrites(value,31,14)
	l.Unlock()

	delay()
	l.Lock()
	value = 23
	s23.Put(181,value)
	log("p31 w:%d", value)
	countwrites(value,31,23)
	l.Unlock()

	delay()
	l.Lock()
	value = 29
	s29.Put(229,value)
	log("p31 w:%d", value)
	countwrites(value,31,29)
	l.Unlock()

	delay()
	l.Lock()
	value = 12
	s12.Put(92,value)
	log("p31 w:%d", value)
	countwrites(value,31,12)
	l.Unlock()

	delay()
	l.Lock()
	value = 5
	s5.Put(37,value)
	log("p31 w:%d", value)
	countwrites(value,31,5)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s1.QueryP(7,&value)
	log("p31 r:%d", value)
	countreads(value,31,1)
	l.Unlock()

	delay()
	l.Lock()
	value = 4
	s4.Put(26,value)
	log("p31 w:%d", value)
	countwrites(value,31,4)
	l.Unlock()
}


func p32() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 4
	s4.Put(29,value)
	log("p32 w:%d", value)
	countwrites(value,32,4)
	l.Unlock()

	delay()
	l.Lock()
	value = 12
	s12.Put(93,value)
	log("p32 w:%d", value)
	countwrites(value,32,12)
	l.Unlock()

	delay()
	l.Lock()
	value = 15
	s15.Put(114,value)
	log("p32 w:%d", value)
	countwrites(value,32,15)
	l.Unlock()

	delay()
	l.Lock()
	value = 19
	s19.Put(148,value)
	log("p32 w:%d", value)
	countwrites(value,32,19)
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
	value = 26
	s26.Put(204,value)
	log("p32 w:%d", value)
	countwrites(value,32,26)
	l.Unlock()

	delay()
	l.Lock()
	value = 5
	s5.Put(38,value)
	log("p32 w:%d", value)
	countwrites(value,32,5)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s19.QueryP(146,&value)
	log("p32 r:%d", value)
	countreads(value,32,19)
	l.Unlock()

	delay()
	l.Lock()
	value = 3
	s3.Put(19,value)
	log("p32 w:%d", value)
	countwrites(value,32,3)
	l.Unlock()

	delay()
	l.Lock()
	value = 32
	s32.Put(249,value)
	log("p32 w:%d", value)
	countwrites(value,32,32)
	l.Unlock()

	delay()
	l.Lock()
	value = 1
	s1.Put(7,value)
	log("p32 w:%d", value)
	countwrites(value,32,1)
	l.Unlock()

	delay()
	l.Lock()
	value = 5
	s5.Put(35,value)
	log("p32 w:%d", value)
	countwrites(value,32,5)
	l.Unlock()

	delay()
	l.Lock()
	value = 8
	s8.Put(63,value)
	log("p32 w:%d", value)
	countwrites(value,32,8)
	l.Unlock()

	delay()
	l.Lock()
	value = 12
	s12.Put(92,value)
	log("p32 w:%d", value)
	countwrites(value,32,12)
	l.Unlock()

	delay()
	l.Lock()
	value = 22
	s22.Put(176,value)
	log("p32 w:%d", value)
	countwrites(value,32,22)
	l.Unlock()

	delay()
	l.Lock()
	value = 14
	s14.Put(109,value)
	log("p32 w:%d", value)
	countwrites(value,32,14)
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
