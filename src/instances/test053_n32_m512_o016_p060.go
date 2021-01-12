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
	s12.QueryP(184,&value)
	log("p1 r:%d", value)
	countreads(value,1,12)
	l.Unlock()

	delay()
	l.Lock()
	value = 18
	s18.Put(286,value)
	log("p1 w:%d", value)
	countwrites(value,1,18)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s26.QueryP(416,&value)
	log("p1 r:%d", value)
	countreads(value,1,26)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s11.QueryP(168,&value)
	log("p1 r:%d", value)
	countreads(value,1,11)
	l.Unlock()

	delay()
	l.Lock()
	value = 26
	s26.Put(408,value)
	log("p1 w:%d", value)
	countwrites(value,1,26)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s1.QueryP(2,&value)
	log("p1 r:%d", value)
	countreads(value,1,1)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s10.QueryP(146,&value)
	log("p1 r:%d", value)
	countreads(value,1,10)
	l.Unlock()

	delay()
	l.Lock()
	value = 7
	s7.Put(107,value)
	log("p1 w:%d", value)
	countwrites(value,1,7)
	l.Unlock()

	delay()
	l.Lock()
	value = 31
	s31.Put(487,value)
	log("p1 w:%d", value)
	countwrites(value,1,31)
	l.Unlock()

	delay()
	l.Lock()
	value = 15
	s15.Put(227,value)
	log("p1 w:%d", value)
	countwrites(value,1,15)
	l.Unlock()

	delay()
	l.Lock()
	value = 11
	s11.Put(169,value)
	log("p1 w:%d", value)
	countwrites(value,1,11)
	l.Unlock()

	delay()
	l.Lock()
	value = 19
	s19.Put(304,value)
	log("p1 w:%d", value)
	countwrites(value,1,19)
	l.Unlock()

	delay()
	l.Lock()
	value = 11
	s11.Put(176,value)
	log("p1 w:%d", value)
	countwrites(value,1,11)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s6.QueryP(94,&value)
	log("p1 r:%d", value)
	countreads(value,1,6)
	l.Unlock()

	delay()
	l.Lock()
	value = 15
	s15.Put(238,value)
	log("p1 w:%d", value)
	countwrites(value,1,15)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s13.QueryP(198,&value)
	log("p1 r:%d", value)
	countreads(value,1,13)
	l.Unlock()
}

func p2() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 28
	s28.Put(434,value)
	log("p2 w:%d", value)
	countwrites(value,2,28)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s21.QueryP(325,&value)
	log("p2 r:%d", value)
	countreads(value,2,21)
	l.Unlock()

	delay()
	l.Lock()
	value = 12
	s12.Put(177,value)
	log("p2 w:%d", value)
	countwrites(value,2,12)
	l.Unlock()

	delay()
	l.Lock()
	value = 25
	s25.Put(394,value)
	log("p2 w:%d", value)
	countwrites(value,2,25)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s2.QueryP(27,&value)
	log("p2 r:%d", value)
	countreads(value,2,2)
	l.Unlock()

	delay()
	l.Lock()
	value = 25
	s25.Put(395,value)
	log("p2 w:%d", value)
	countwrites(value,2,25)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s21.QueryP(325,&value)
	log("p2 r:%d", value)
	countreads(value,2,21)
	l.Unlock()

	delay()
	l.Lock()
	value = 23
	s23.Put(360,value)
	log("p2 w:%d", value)
	countwrites(value,2,23)
	l.Unlock()

	delay()
	l.Lock()
	value = 29
	s29.Put(462,value)
	log("p2 w:%d", value)
	countwrites(value,2,29)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s11.QueryP(161,&value)
	log("p2 r:%d", value)
	countreads(value,2,11)
	l.Unlock()

	delay()
	l.Lock()
	value = 18
	s18.Put(274,value)
	log("p2 w:%d", value)
	countwrites(value,2,18)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s28.QueryP(441,&value)
	log("p2 r:%d", value)
	countreads(value,2,28)
	l.Unlock()

	delay()
	l.Lock()
	value = 13
	s13.Put(194,value)
	log("p2 w:%d", value)
	countwrites(value,2,13)
	l.Unlock()

	delay()
	l.Lock()
	value = 16
	s16.Put(241,value)
	log("p2 w:%d", value)
	countwrites(value,2,16)
	l.Unlock()

	delay()
	l.Lock()
	value = 17
	s17.Put(271,value)
	log("p2 w:%d", value)
	countwrites(value,2,17)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s16.QueryP(256,&value)
	log("p2 r:%d", value)
	countreads(value,2,16)
	l.Unlock()
}

func p3() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	s15.QueryP(231,&value)
	log("p3 r:%d", value)
	countreads(value,3,15)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s9.QueryP(132,&value)
	log("p3 r:%d", value)
	countreads(value,3,9)
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
	value = 26
	s26.Put(407,value)
	log("p3 w:%d", value)
	countwrites(value,3,26)
	l.Unlock()

	delay()
	l.Lock()
	value = 23
	s23.Put(368,value)
	log("p3 w:%d", value)
	countwrites(value,3,23)
	l.Unlock()

	delay()
	l.Lock()
	value = 3
	s3.Put(35,value)
	log("p3 w:%d", value)
	countwrites(value,3,3)
	l.Unlock()

	delay()
	l.Lock()
	value = 11
	s11.Put(165,value)
	log("p3 w:%d", value)
	countwrites(value,3,11)
	l.Unlock()

	delay()
	l.Lock()
	value = 31
	s31.Put(485,value)
	log("p3 w:%d", value)
	countwrites(value,3,31)
	l.Unlock()

	delay()
	l.Lock()
	value = 9
	s9.Put(136,value)
	log("p3 w:%d", value)
	countwrites(value,3,9)
	l.Unlock()

	delay()
	l.Lock()
	value = 11
	s11.Put(172,value)
	log("p3 w:%d", value)
	countwrites(value,3,11)
	l.Unlock()

	delay()
	l.Lock()
	value = 10
	s10.Put(159,value)
	log("p3 w:%d", value)
	countwrites(value,3,10)
	l.Unlock()

	delay()
	l.Lock()
	value = 5
	s5.Put(69,value)
	log("p3 w:%d", value)
	countwrites(value,3,5)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s25.QueryP(400,&value)
	log("p3 r:%d", value)
	countreads(value,3,25)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s7.QueryP(103,&value)
	log("p3 r:%d", value)
	countreads(value,3,7)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s18.QueryP(286,&value)
	log("p3 r:%d", value)
	countreads(value,3,18)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s9.QueryP(141,&value)
	log("p3 r:%d", value)
	countreads(value,3,9)
	l.Unlock()
}

func p4() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	s15.QueryP(240,&value)
	log("p4 r:%d", value)
	countreads(value,4,15)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s32.QueryP(505,&value)
	log("p4 r:%d", value)
	countreads(value,4,32)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s21.QueryP(336,&value)
	log("p4 r:%d", value)
	countreads(value,4,21)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s6.QueryP(81,&value)
	log("p4 r:%d", value)
	countreads(value,4,6)
	l.Unlock()

	delay()
	l.Lock()
	value = 17
	s17.Put(263,value)
	log("p4 w:%d", value)
	countwrites(value,4,17)
	l.Unlock()

	delay()
	l.Lock()
	value = 27
	s27.Put(427,value)
	log("p4 w:%d", value)
	countwrites(value,4,27)
	l.Unlock()

	delay()
	l.Lock()
	value = 3
	s3.Put(37,value)
	log("p4 w:%d", value)
	countwrites(value,4,3)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s18.QueryP(282,&value)
	log("p4 r:%d", value)
	countreads(value,4,18)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s18.QueryP(286,&value)
	log("p4 r:%d", value)
	countreads(value,4,18)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s19.QueryP(290,&value)
	log("p4 r:%d", value)
	countreads(value,4,19)
	l.Unlock()

	delay()
	l.Lock()
	value = 12
	s12.Put(192,value)
	log("p4 w:%d", value)
	countwrites(value,4,12)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s13.QueryP(199,&value)
	log("p4 r:%d", value)
	countreads(value,4,13)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s15.QueryP(226,&value)
	log("p4 r:%d", value)
	countreads(value,4,15)
	l.Unlock()

	delay()
	l.Lock()
	value = 15
	s15.Put(230,value)
	log("p4 w:%d", value)
	countwrites(value,4,15)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s6.QueryP(95,&value)
	log("p4 r:%d", value)
	countreads(value,4,6)
	l.Unlock()

	delay()
	l.Lock()
	value = 2
	s2.Put(28,value)
	log("p4 w:%d", value)
	countwrites(value,4,2)
	l.Unlock()
}

func p5() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	s29.QueryP(452,&value)
	log("p5 r:%d", value)
	countreads(value,5,29)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s14.QueryP(217,&value)
	log("p5 r:%d", value)
	countreads(value,5,14)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s16.QueryP(247,&value)
	log("p5 r:%d", value)
	countreads(value,5,16)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s5.QueryP(69,&value)
	log("p5 r:%d", value)
	countreads(value,5,5)
	l.Unlock()

	delay()
	l.Lock()
	value = 2
	s2.Put(31,value)
	log("p5 w:%d", value)
	countwrites(value,5,2)
	l.Unlock()

	delay()
	l.Lock()
	value = 32
	s32.Put(509,value)
	log("p5 w:%d", value)
	countwrites(value,5,32)
	l.Unlock()

	delay()
	l.Lock()
	value = 29
	s29.Put(454,value)
	log("p5 w:%d", value)
	countwrites(value,5,29)
	l.Unlock()

	delay()
	l.Lock()
	value = 3
	s3.Put(40,value)
	log("p5 w:%d", value)
	countwrites(value,5,3)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s16.QueryP(245,&value)
	log("p5 r:%d", value)
	countreads(value,5,16)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s19.QueryP(292,&value)
	log("p5 r:%d", value)
	countreads(value,5,19)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s25.QueryP(392,&value)
	log("p5 r:%d", value)
	countreads(value,5,25)
	l.Unlock()

	delay()
	l.Lock()
	value = 24
	s24.Put(377,value)
	log("p5 w:%d", value)
	countwrites(value,5,24)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s8.QueryP(123,&value)
	log("p5 r:%d", value)
	countreads(value,5,8)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s9.QueryP(139,&value)
	log("p5 r:%d", value)
	countreads(value,5,9)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s10.QueryP(159,&value)
	log("p5 r:%d", value)
	countreads(value,5,10)
	l.Unlock()

	delay()
	l.Lock()
	value = 20
	s20.Put(311,value)
	log("p5 w:%d", value)
	countwrites(value,5,20)
	l.Unlock()
}

func p6() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	s8.QueryP(125,&value)
	log("p6 r:%d", value)
	countreads(value,6,8)
	l.Unlock()

	delay()
	l.Lock()
	value = 6
	s6.Put(84,value)
	log("p6 w:%d", value)
	countwrites(value,6,6)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s13.QueryP(201,&value)
	log("p6 r:%d", value)
	countreads(value,6,13)
	l.Unlock()

	delay()
	l.Lock()
	value = 13
	s13.Put(195,value)
	log("p6 w:%d", value)
	countwrites(value,6,13)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s29.QueryP(454,&value)
	log("p6 r:%d", value)
	countreads(value,6,29)
	l.Unlock()

	delay()
	l.Lock()
	value = 3
	s3.Put(40,value)
	log("p6 w:%d", value)
	countwrites(value,6,3)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s5.QueryP(75,&value)
	log("p6 r:%d", value)
	countreads(value,6,5)
	l.Unlock()

	delay()
	l.Lock()
	value = 12
	s12.Put(186,value)
	log("p6 w:%d", value)
	countwrites(value,6,12)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s17.QueryP(272,&value)
	log("p6 r:%d", value)
	countreads(value,6,17)
	l.Unlock()

	delay()
	l.Lock()
	value = 30
	s30.Put(467,value)
	log("p6 w:%d", value)
	countwrites(value,6,30)
	l.Unlock()

	delay()
	l.Lock()
	value = 14
	s14.Put(218,value)
	log("p6 w:%d", value)
	countwrites(value,6,14)
	l.Unlock()

	delay()
	l.Lock()
	value = 21
	s21.Put(331,value)
	log("p6 w:%d", value)
	countwrites(value,6,21)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s2.QueryP(21,&value)
	log("p6 r:%d", value)
	countreads(value,6,2)
	l.Unlock()

	delay()
	l.Lock()
	value = 28
	s28.Put(436,value)
	log("p6 w:%d", value)
	countwrites(value,6,28)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s3.QueryP(39,&value)
	log("p6 r:%d", value)
	countreads(value,6,3)
	l.Unlock()

	delay()
	l.Lock()
	value = 17
	s17.Put(258,value)
	log("p6 w:%d", value)
	countwrites(value,6,17)
	l.Unlock()
}

func p7() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 31
	s31.Put(488,value)
	log("p7 w:%d", value)
	countwrites(value,7,31)
	l.Unlock()

	delay()
	l.Lock()
	value = 19
	s19.Put(296,value)
	log("p7 w:%d", value)
	countwrites(value,7,19)
	l.Unlock()

	delay()
	l.Lock()
	value = 26
	s26.Put(402,value)
	log("p7 w:%d", value)
	countwrites(value,7,26)
	l.Unlock()

	delay()
	l.Lock()
	value = 24
	s24.Put(378,value)
	log("p7 w:%d", value)
	countwrites(value,7,24)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s3.QueryP(38,&value)
	log("p7 r:%d", value)
	countreads(value,7,3)
	l.Unlock()

	delay()
	l.Lock()
	value = 5
	s5.Put(69,value)
	log("p7 w:%d", value)
	countwrites(value,7,5)
	l.Unlock()

	delay()
	l.Lock()
	value = 31
	s31.Put(491,value)
	log("p7 w:%d", value)
	countwrites(value,7,31)
	l.Unlock()

	delay()
	l.Lock()
	value = 28
	s28.Put(444,value)
	log("p7 w:%d", value)
	countwrites(value,7,28)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s24.QueryP(380,&value)
	log("p7 r:%d", value)
	countreads(value,7,24)
	l.Unlock()

	delay()
	l.Lock()
	value = 23
	s23.Put(368,value)
	log("p7 w:%d", value)
	countwrites(value,7,23)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s25.QueryP(389,&value)
	log("p7 r:%d", value)
	countreads(value,7,25)
	l.Unlock()

	delay()
	l.Lock()
	value = 3
	s3.Put(45,value)
	log("p7 w:%d", value)
	countwrites(value,7,3)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s15.QueryP(232,&value)
	log("p7 r:%d", value)
	countreads(value,7,15)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s18.QueryP(275,&value)
	log("p7 r:%d", value)
	countreads(value,7,18)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s17.QueryP(258,&value)
	log("p7 r:%d", value)
	countreads(value,7,17)
	l.Unlock()

	delay()
	l.Lock()
	value = 21
	s21.Put(333,value)
	log("p7 w:%d", value)
	countwrites(value,7,21)
	l.Unlock()
}

func p8() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 15
	s15.Put(232,value)
	log("p8 w:%d", value)
	countwrites(value,8,15)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s17.QueryP(263,&value)
	log("p8 r:%d", value)
	countreads(value,8,17)
	l.Unlock()

	delay()
	l.Lock()
	value = 2
	s2.Put(27,value)
	log("p8 w:%d", value)
	countwrites(value,8,2)
	l.Unlock()

	delay()
	l.Lock()
	value = 11
	s11.Put(168,value)
	log("p8 w:%d", value)
	countwrites(value,8,11)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s32.QueryP(504,&value)
	log("p8 r:%d", value)
	countreads(value,8,32)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s13.QueryP(198,&value)
	log("p8 r:%d", value)
	countreads(value,8,13)
	l.Unlock()

	delay()
	l.Lock()
	value = 12
	s12.Put(190,value)
	log("p8 w:%d", value)
	countwrites(value,8,12)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s14.QueryP(215,&value)
	log("p8 r:%d", value)
	countreads(value,8,14)
	l.Unlock()

	delay()
	l.Lock()
	value = 8
	s8.Put(115,value)
	log("p8 w:%d", value)
	countwrites(value,8,8)
	l.Unlock()

	delay()
	l.Lock()
	value = 21
	s21.Put(335,value)
	log("p8 w:%d", value)
	countwrites(value,8,21)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s14.QueryP(222,&value)
	log("p8 r:%d", value)
	countreads(value,8,14)
	l.Unlock()

	delay()
	l.Lock()
	value = 1
	s1.Put(12,value)
	log("p8 w:%d", value)
	countwrites(value,8,1)
	l.Unlock()

	delay()
	l.Lock()
	value = 27
	s27.Put(431,value)
	log("p8 w:%d", value)
	countwrites(value,8,27)
	l.Unlock()

	delay()
	l.Lock()
	value = 28
	s28.Put(442,value)
	log("p8 w:%d", value)
	countwrites(value,8,28)
	l.Unlock()

	delay()
	l.Lock()
	value = 15
	s15.Put(226,value)
	log("p8 w:%d", value)
	countwrites(value,8,15)
	l.Unlock()

	delay()
	l.Lock()
	value = 3
	s3.Put(35,value)
	log("p8 w:%d", value)
	countwrites(value,8,3)
	l.Unlock()
}

func p9() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 18
	s18.Put(278,value)
	log("p9 w:%d", value)
	countwrites(value,9,18)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s17.QueryP(266,&value)
	log("p9 r:%d", value)
	countreads(value,9,17)
	l.Unlock()

	delay()
	l.Lock()
	value = 28
	s28.Put(447,value)
	log("p9 w:%d", value)
	countwrites(value,9,28)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s12.QueryP(181,&value)
	log("p9 r:%d", value)
	countreads(value,9,12)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s12.QueryP(192,&value)
	log("p9 r:%d", value)
	countreads(value,9,12)
	l.Unlock()

	delay()
	l.Lock()
	value = 10
	s10.Put(160,value)
	log("p9 w:%d", value)
	countwrites(value,9,10)
	l.Unlock()

	delay()
	l.Lock()
	value = 23
	s23.Put(368,value)
	log("p9 w:%d", value)
	countwrites(value,9,23)
	l.Unlock()

	delay()
	l.Lock()
	value = 32
	s32.Put(506,value)
	log("p9 w:%d", value)
	countwrites(value,9,32)
	l.Unlock()

	delay()
	l.Lock()
	value = 26
	s26.Put(407,value)
	log("p9 w:%d", value)
	countwrites(value,9,26)
	l.Unlock()

	delay()
	l.Lock()
	value = 17
	s17.Put(262,value)
	log("p9 w:%d", value)
	countwrites(value,9,17)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s14.QueryP(224,&value)
	log("p9 r:%d", value)
	countreads(value,9,14)
	l.Unlock()

	delay()
	l.Lock()
	value = 12
	s12.Put(183,value)
	log("p9 w:%d", value)
	countwrites(value,9,12)
	l.Unlock()

	delay()
	l.Lock()
	value = 30
	s30.Put(480,value)
	log("p9 w:%d", value)
	countwrites(value,9,30)
	l.Unlock()

	delay()
	l.Lock()
	value = 31
	s31.Put(484,value)
	log("p9 w:%d", value)
	countwrites(value,9,31)
	l.Unlock()

	delay()
	l.Lock()
	value = 21
	s21.Put(323,value)
	log("p9 w:%d", value)
	countwrites(value,9,21)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s4.QueryP(62,&value)
	log("p9 r:%d", value)
	countreads(value,9,4)
	l.Unlock()
}


func p10() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 26
	s26.Put(403,value)
	log("p10 w:%d", value)
	countwrites(value,10,26)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s8.QueryP(121,&value)
	log("p10 r:%d", value)
	countreads(value,10,8)
	l.Unlock()

	delay()
	l.Lock()
	value = 14
	s14.Put(213,value)
	log("p10 w:%d", value)
	countwrites(value,10,14)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s32.QueryP(498,&value)
	log("p10 r:%d", value)
	countreads(value,10,32)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s18.QueryP(279,&value)
	log("p10 r:%d", value)
	countreads(value,10,18)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s30.QueryP(473,&value)
	log("p10 r:%d", value)
	countreads(value,10,30)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s26.QueryP(415,&value)
	log("p10 r:%d", value)
	countreads(value,10,26)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s6.QueryP(88,&value)
	log("p10 r:%d", value)
	countreads(value,10,6)
	l.Unlock()

	delay()
	l.Lock()
	value = 27
	s27.Put(418,value)
	log("p10 w:%d", value)
	countwrites(value,10,27)
	l.Unlock()

	delay()
	l.Lock()
	value = 12
	s12.Put(181,value)
	log("p10 w:%d", value)
	countwrites(value,10,12)
	l.Unlock()

	delay()
	l.Lock()
	value = 23
	s23.Put(368,value)
	log("p10 w:%d", value)
	countwrites(value,10,23)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s10.QueryP(152,&value)
	log("p10 r:%d", value)
	countreads(value,10,10)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s21.QueryP(332,&value)
	log("p10 r:%d", value)
	countreads(value,10,21)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s24.QueryP(369,&value)
	log("p10 r:%d", value)
	countreads(value,10,24)
	l.Unlock()

	delay()
	l.Lock()
	value = 30
	s30.Put(472,value)
	log("p10 w:%d", value)
	countwrites(value,10,30)
	l.Unlock()

	delay()
	l.Lock()
	value = 7
	s7.Put(107,value)
	log("p10 w:%d", value)
	countwrites(value,10,7)
	l.Unlock()
}


func p11() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 23
	s23.Put(362,value)
	log("p11 w:%d", value)
	countwrites(value,11,23)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s7.QueryP(99,&value)
	log("p11 r:%d", value)
	countreads(value,11,7)
	l.Unlock()

	delay()
	l.Lock()
	value = 12
	s12.Put(190,value)
	log("p11 w:%d", value)
	countwrites(value,11,12)
	l.Unlock()

	delay()
	l.Lock()
	value = 19
	s19.Put(304,value)
	log("p11 w:%d", value)
	countwrites(value,11,19)
	l.Unlock()

	delay()
	l.Lock()
	value = 15
	s15.Put(233,value)
	log("p11 w:%d", value)
	countwrites(value,11,15)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s22.QueryP(349,&value)
	log("p11 r:%d", value)
	countreads(value,11,22)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s21.QueryP(329,&value)
	log("p11 r:%d", value)
	countreads(value,11,21)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s11.QueryP(170,&value)
	log("p11 r:%d", value)
	countreads(value,11,11)
	l.Unlock()

	delay()
	l.Lock()
	value = 22
	s22.Put(351,value)
	log("p11 w:%d", value)
	countwrites(value,11,22)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s18.QueryP(284,&value)
	log("p11 r:%d", value)
	countreads(value,11,18)
	l.Unlock()

	delay()
	l.Lock()
	value = 3
	s3.Put(46,value)
	log("p11 w:%d", value)
	countwrites(value,11,3)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s1.QueryP(11,&value)
	log("p11 r:%d", value)
	countreads(value,11,1)
	l.Unlock()

	delay()
	l.Lock()
	value = 17
	s17.Put(268,value)
	log("p11 w:%d", value)
	countwrites(value,11,17)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s16.QueryP(247,&value)
	log("p11 r:%d", value)
	countreads(value,11,16)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s1.QueryP(13,&value)
	log("p11 r:%d", value)
	countreads(value,11,1)
	l.Unlock()

	delay()
	l.Lock()
	value = 14
	s14.Put(219,value)
	log("p11 w:%d", value)
	countwrites(value,11,14)
	l.Unlock()
}


func p12() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	s23.QueryP(367,&value)
	log("p12 r:%d", value)
	countreads(value,12,23)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s4.QueryP(54,&value)
	log("p12 r:%d", value)
	countreads(value,12,4)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s23.QueryP(362,&value)
	log("p12 r:%d", value)
	countreads(value,12,23)
	l.Unlock()

	delay()
	l.Lock()
	value = 9
	s9.Put(138,value)
	log("p12 w:%d", value)
	countwrites(value,12,9)
	l.Unlock()

	delay()
	l.Lock()
	value = 17
	s17.Put(271,value)
	log("p12 w:%d", value)
	countwrites(value,12,17)
	l.Unlock()

	delay()
	l.Lock()
	value = 4
	s4.Put(50,value)
	log("p12 w:%d", value)
	countwrites(value,12,4)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s4.QueryP(62,&value)
	log("p12 r:%d", value)
	countreads(value,12,4)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s1.QueryP(14,&value)
	log("p12 r:%d", value)
	countreads(value,12,1)
	l.Unlock()

	delay()
	l.Lock()
	value = 1
	s1.Put(10,value)
	log("p12 w:%d", value)
	countwrites(value,12,1)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s18.QueryP(286,&value)
	log("p12 r:%d", value)
	countreads(value,12,18)
	l.Unlock()

	delay()
	l.Lock()
	value = 13
	s13.Put(203,value)
	log("p12 w:%d", value)
	countwrites(value,12,13)
	l.Unlock()

	delay()
	l.Lock()
	value = 8
	s8.Put(123,value)
	log("p12 w:%d", value)
	countwrites(value,12,8)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s32.QueryP(499,&value)
	log("p12 r:%d", value)
	countreads(value,12,32)
	l.Unlock()

	delay()
	l.Lock()
	value = 4
	s4.Put(54,value)
	log("p12 w:%d", value)
	countwrites(value,12,4)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s30.QueryP(467,&value)
	log("p12 r:%d", value)
	countreads(value,12,30)
	l.Unlock()

	delay()
	l.Lock()
	value = 11
	s11.Put(168,value)
	log("p12 w:%d", value)
	countwrites(value,12,11)
	l.Unlock()
}


func p13() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	s22.QueryP(344,&value)
	log("p13 r:%d", value)
	countreads(value,13,22)
	l.Unlock()

	delay()
	l.Lock()
	value = 8
	s8.Put(115,value)
	log("p13 w:%d", value)
	countwrites(value,13,8)
	l.Unlock()

	delay()
	l.Lock()
	value = 31
	s31.Put(494,value)
	log("p13 w:%d", value)
	countwrites(value,13,31)
	l.Unlock()

	delay()
	l.Lock()
	value = 6
	s6.Put(88,value)
	log("p13 w:%d", value)
	countwrites(value,13,6)
	l.Unlock()

	delay()
	l.Lock()
	value = 10
	s10.Put(150,value)
	log("p13 w:%d", value)
	countwrites(value,13,10)
	l.Unlock()

	delay()
	l.Lock()
	value = 32
	s32.Put(504,value)
	log("p13 w:%d", value)
	countwrites(value,13,32)
	l.Unlock()

	delay()
	l.Lock()
	value = 4
	s4.Put(52,value)
	log("p13 w:%d", value)
	countwrites(value,13,4)
	l.Unlock()

	delay()
	l.Lock()
	value = 29
	s29.Put(450,value)
	log("p13 w:%d", value)
	countwrites(value,13,29)
	l.Unlock()

	delay()
	l.Lock()
	value = 25
	s25.Put(396,value)
	log("p13 w:%d", value)
	countwrites(value,13,25)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s11.QueryP(163,&value)
	log("p13 r:%d", value)
	countreads(value,13,11)
	l.Unlock()

	delay()
	l.Lock()
	value = 28
	s28.Put(434,value)
	log("p13 w:%d", value)
	countwrites(value,13,28)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s17.QueryP(269,&value)
	log("p13 r:%d", value)
	countreads(value,13,17)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s1.QueryP(14,&value)
	log("p13 r:%d", value)
	countreads(value,13,1)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s31.QueryP(484,&value)
	log("p13 r:%d", value)
	countreads(value,13,31)
	l.Unlock()

	delay()
	l.Lock()
	value = 11
	s11.Put(169,value)
	log("p13 w:%d", value)
	countwrites(value,13,11)
	l.Unlock()

	delay()
	l.Lock()
	value = 22
	s22.Put(349,value)
	log("p13 w:%d", value)
	countwrites(value,13,22)
	l.Unlock()
}

func p14() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	s30.QueryP(475,&value)
	log("p14 r:%d", value)
	countreads(value,14,30)
	l.Unlock()

	delay()
	l.Lock()
	value = 21
	s21.Put(333,value)
	log("p14 w:%d", value)
	countwrites(value,14,21)
	l.Unlock()

	delay()
	l.Lock()
	value = 6
	s6.Put(95,value)
	log("p14 w:%d", value)
	countwrites(value,14,6)
	l.Unlock()

	delay()
	l.Lock()
	value = 14
	s14.Put(215,value)
	log("p14 w:%d", value)
	countwrites(value,14,14)
	l.Unlock()

	delay()
	l.Lock()
	value = 1
	s1.Put(11,value)
	log("p14 w:%d", value)
	countwrites(value,14,1)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s9.QueryP(131,&value)
	log("p14 r:%d", value)
	countreads(value,14,9)
	l.Unlock()

	delay()
	l.Lock()
	value = 7
	s7.Put(100,value)
	log("p14 w:%d", value)
	countwrites(value,14,7)
	l.Unlock()

	delay()
	l.Lock()
	value = 14
	s14.Put(210,value)
	log("p14 w:%d", value)
	countwrites(value,14,14)
	l.Unlock()

	delay()
	l.Lock()
	value = 20
	s20.Put(309,value)
	log("p14 w:%d", value)
	countwrites(value,14,20)
	l.Unlock()

	delay()
	l.Lock()
	value = 30
	s30.Put(472,value)
	log("p14 w:%d", value)
	countwrites(value,14,30)
	l.Unlock()

	delay()
	l.Lock()
	value = 24
	s24.Put(384,value)
	log("p14 w:%d", value)
	countwrites(value,14,24)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s29.QueryP(453,&value)
	log("p14 r:%d", value)
	countreads(value,14,29)
	l.Unlock()

	delay()
	l.Lock()
	value = 15
	s15.Put(239,value)
	log("p14 w:%d", value)
	countwrites(value,14,15)
	l.Unlock()

	delay()
	l.Lock()
	value = 29
	s29.Put(454,value)
	log("p14 w:%d", value)
	countwrites(value,14,29)
	l.Unlock()

	delay()
	l.Lock()
	value = 1
	s1.Put(7,value)
	log("p14 w:%d", value)
	countwrites(value,14,1)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s15.QueryP(234,&value)
	log("p14 r:%d", value)
	countreads(value,14,15)
	l.Unlock()
}


func p15() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 21
	s21.Put(333,value)
	log("p15 w:%d", value)
	countwrites(value,15,21)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s30.QueryP(476,&value)
	log("p15 r:%d", value)
	countreads(value,15,30)
	l.Unlock()

	delay()
	l.Lock()
	value = 28
	s28.Put(437,value)
	log("p15 w:%d", value)
	countwrites(value,15,28)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s25.QueryP(394,&value)
	log("p15 r:%d", value)
	countreads(value,15,25)
	l.Unlock()

	delay()
	l.Lock()
	value = 12
	s12.Put(186,value)
	log("p15 w:%d", value)
	countwrites(value,15,12)
	l.Unlock()

	delay()
	l.Lock()
	value = 29
	s29.Put(457,value)
	log("p15 w:%d", value)
	countwrites(value,15,29)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s16.QueryP(252,&value)
	log("p15 r:%d", value)
	countreads(value,15,16)
	l.Unlock()

	delay()
	l.Lock()
	value = 5
	s5.Put(78,value)
	log("p15 w:%d", value)
	countwrites(value,15,5)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s28.QueryP(436,&value)
	log("p15 r:%d", value)
	countreads(value,15,28)
	l.Unlock()

	delay()
	l.Lock()
	value = 26
	s26.Put(413,value)
	log("p15 w:%d", value)
	countwrites(value,15,26)
	l.Unlock()

	delay()
	l.Lock()
	value = 23
	s23.Put(356,value)
	log("p15 w:%d", value)
	countwrites(value,15,23)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s26.QueryP(416,&value)
	log("p15 r:%d", value)
	countreads(value,15,26)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s18.QueryP(273,&value)
	log("p15 r:%d", value)
	countreads(value,15,18)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s16.QueryP(256,&value)
	log("p15 r:%d", value)
	countreads(value,15,16)
	l.Unlock()

	delay()
	l.Lock()
	value = 14
	s14.Put(209,value)
	log("p15 w:%d", value)
	countwrites(value,15,14)
	l.Unlock()

	delay()
	l.Lock()
	value = 24
	s24.Put(382,value)
	log("p15 w:%d", value)
	countwrites(value,15,24)
	l.Unlock()
}


func p16() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 15
	s15.Put(237,value)
	log("p16 w:%d", value)
	countwrites(value,16,15)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s2.QueryP(23,&value)
	log("p16 r:%d", value)
	countreads(value,16,2)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s3.QueryP(45,&value)
	log("p16 r:%d", value)
	countreads(value,16,3)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s28.QueryP(444,&value)
	log("p16 r:%d", value)
	countreads(value,16,28)
	l.Unlock()

	delay()
	l.Lock()
	value = 12
	s12.Put(185,value)
	log("p16 w:%d", value)
	countwrites(value,16,12)
	l.Unlock()

	delay()
	l.Lock()
	value = 14
	s14.Put(212,value)
	log("p16 w:%d", value)
	countwrites(value,16,14)
	l.Unlock()

	delay()
	l.Lock()
	value = 24
	s24.Put(384,value)
	log("p16 w:%d", value)
	countwrites(value,16,24)
	l.Unlock()

	delay()
	l.Lock()
	value = 13
	s13.Put(199,value)
	log("p16 w:%d", value)
	countwrites(value,16,13)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s30.QueryP(472,&value)
	log("p16 r:%d", value)
	countreads(value,16,30)
	l.Unlock()

	delay()
	l.Lock()
	value = 4
	s4.Put(49,value)
	log("p16 w:%d", value)
	countwrites(value,16,4)
	l.Unlock()

	delay()
	l.Lock()
	value = 7
	s7.Put(105,value)
	log("p16 w:%d", value)
	countwrites(value,16,7)
	l.Unlock()

	delay()
	l.Lock()
	value = 17
	s17.Put(259,value)
	log("p16 w:%d", value)
	countwrites(value,16,17)
	l.Unlock()

	delay()
	l.Lock()
	value = 12
	s12.Put(189,value)
	log("p16 w:%d", value)
	countwrites(value,16,12)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s31.QueryP(494,&value)
	log("p16 r:%d", value)
	countreads(value,16,31)
	l.Unlock()

	delay()
	l.Lock()
	value = 11
	s11.Put(175,value)
	log("p16 w:%d", value)
	countwrites(value,16,11)
	l.Unlock()

	delay()
	l.Lock()
	value = 8
	s8.Put(128,value)
	log("p16 w:%d", value)
	countwrites(value,16,8)
	l.Unlock()
}


func p17() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 19
	s19.Put(290,value)
	log("p17 w:%d", value)
	countwrites(value,17,19)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s5.QueryP(70,&value)
	log("p17 r:%d", value)
	countreads(value,17,5)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s19.QueryP(293,&value)
	log("p17 r:%d", value)
	countreads(value,17,19)
	l.Unlock()

	delay()
	l.Lock()
	value = 23
	s23.Put(356,value)
	log("p17 w:%d", value)
	countwrites(value,17,23)
	l.Unlock()

	delay()
	l.Lock()
	value = 19
	s19.Put(304,value)
	log("p17 w:%d", value)
	countwrites(value,17,19)
	l.Unlock()

	delay()
	l.Lock()
	value = 14
	s14.Put(212,value)
	log("p17 w:%d", value)
	countwrites(value,17,14)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s21.QueryP(333,&value)
	log("p17 r:%d", value)
	countreads(value,17,21)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s16.QueryP(253,&value)
	log("p17 r:%d", value)
	countreads(value,17,16)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s32.QueryP(507,&value)
	log("p17 r:%d", value)
	countreads(value,17,32)
	l.Unlock()

	delay()
	l.Lock()
	value = 1
	s1.Put(14,value)
	log("p17 w:%d", value)
	countwrites(value,17,1)
	l.Unlock()

	delay()
	l.Lock()
	value = 12
	s12.Put(182,value)
	log("p17 w:%d", value)
	countwrites(value,17,12)
	l.Unlock()

	delay()
	l.Lock()
	value = 9
	s9.Put(132,value)
	log("p17 w:%d", value)
	countwrites(value,17,9)
	l.Unlock()

	delay()
	l.Lock()
	value = 24
	s24.Put(373,value)
	log("p17 w:%d", value)
	countwrites(value,17,24)
	l.Unlock()

	delay()
	l.Lock()
	value = 25
	s25.Put(400,value)
	log("p17 w:%d", value)
	countwrites(value,17,25)
	l.Unlock()

	delay()
	l.Lock()
	value = 3
	s3.Put(37,value)
	log("p17 w:%d", value)
	countwrites(value,17,3)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s5.QueryP(80,&value)
	log("p17 r:%d", value)
	countreads(value,17,5)
	l.Unlock()
}


func p18() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	s9.QueryP(141,&value)
	log("p18 r:%d", value)
	countreads(value,18,9)
	l.Unlock()

	delay()
	l.Lock()
	value = 1
	s1.Put(16,value)
	log("p18 w:%d", value)
	countwrites(value,18,1)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s5.QueryP(75,&value)
	log("p18 r:%d", value)
	countreads(value,18,5)
	l.Unlock()

	delay()
	l.Lock()
	value = 19
	s19.Put(302,value)
	log("p18 w:%d", value)
	countwrites(value,18,19)
	l.Unlock()

	delay()
	l.Lock()
	value = 18
	s18.Put(275,value)
	log("p18 w:%d", value)
	countwrites(value,18,18)
	l.Unlock()

	delay()
	l.Lock()
	value = 28
	s28.Put(444,value)
	log("p18 w:%d", value)
	countwrites(value,18,28)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s9.QueryP(137,&value)
	log("p18 r:%d", value)
	countreads(value,18,9)
	l.Unlock()

	delay()
	l.Lock()
	value = 13
	s13.Put(199,value)
	log("p18 w:%d", value)
	countwrites(value,18,13)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s3.QueryP(47,&value)
	log("p18 r:%d", value)
	countreads(value,18,3)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s24.QueryP(372,&value)
	log("p18 r:%d", value)
	countreads(value,18,24)
	l.Unlock()

	delay()
	l.Lock()
	value = 24
	s24.Put(372,value)
	log("p18 w:%d", value)
	countwrites(value,18,24)
	l.Unlock()

	delay()
	l.Lock()
	value = 26
	s26.Put(411,value)
	log("p18 w:%d", value)
	countwrites(value,18,26)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s21.QueryP(327,&value)
	log("p18 r:%d", value)
	countreads(value,18,21)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s12.QueryP(181,&value)
	log("p18 r:%d", value)
	countreads(value,18,12)
	l.Unlock()

	delay()
	l.Lock()
	value = 25
	s25.Put(397,value)
	log("p18 w:%d", value)
	countwrites(value,18,25)
	l.Unlock()

	delay()
	l.Lock()
	value = 26
	s26.Put(403,value)
	log("p18 w:%d", value)
	countwrites(value,18,26)
	l.Unlock()
}


func p19() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	s30.QueryP(474,&value)
	log("p19 r:%d", value)
	countreads(value,19,30)
	l.Unlock()

	delay()
	l.Lock()
	value = 32
	s32.Put(500,value)
	log("p19 w:%d", value)
	countwrites(value,19,32)
	l.Unlock()

	delay()
	l.Lock()
	value = 15
	s15.Put(239,value)
	log("p19 w:%d", value)
	countwrites(value,19,15)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s3.QueryP(40,&value)
	log("p19 r:%d", value)
	countreads(value,19,3)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s29.QueryP(455,&value)
	log("p19 r:%d", value)
	countreads(value,19,29)
	l.Unlock()

	delay()
	l.Lock()
	value = 21
	s21.Put(331,value)
	log("p19 w:%d", value)
	countwrites(value,19,21)
	l.Unlock()

	delay()
	l.Lock()
	value = 31
	s31.Put(488,value)
	log("p19 w:%d", value)
	countwrites(value,19,31)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s32.QueryP(508,&value)
	log("p19 r:%d", value)
	countreads(value,19,32)
	l.Unlock()

	delay()
	l.Lock()
	value = 12
	s12.Put(192,value)
	log("p19 w:%d", value)
	countwrites(value,19,12)
	l.Unlock()

	delay()
	l.Lock()
	value = 20
	s20.Put(307,value)
	log("p19 w:%d", value)
	countwrites(value,19,20)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s22.QueryP(350,&value)
	log("p19 r:%d", value)
	countreads(value,19,22)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s4.QueryP(60,&value)
	log("p19 r:%d", value)
	countreads(value,19,4)
	l.Unlock()

	delay()
	l.Lock()
	value = 13
	s13.Put(204,value)
	log("p19 w:%d", value)
	countwrites(value,19,13)
	l.Unlock()

	delay()
	l.Lock()
	value = 15
	s15.Put(238,value)
	log("p19 w:%d", value)
	countwrites(value,19,15)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s17.QueryP(258,&value)
	log("p19 r:%d", value)
	countreads(value,19,17)
	l.Unlock()

	delay()
	l.Lock()
	value = 23
	s23.Put(367,value)
	log("p19 w:%d", value)
	countwrites(value,19,23)
	l.Unlock()
}


func p20() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 1
	s1.Put(9,value)
	log("p20 w:%d", value)
	countwrites(value,20,1)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s7.QueryP(97,&value)
	log("p20 r:%d", value)
	countreads(value,20,7)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s14.QueryP(212,&value)
	log("p20 r:%d", value)
	countreads(value,20,14)
	l.Unlock()

	delay()
	l.Lock()
	value = 21
	s21.Put(329,value)
	log("p20 w:%d", value)
	countwrites(value,20,21)
	l.Unlock()

	delay()
	l.Lock()
	value = 18
	s18.Put(277,value)
	log("p20 w:%d", value)
	countwrites(value,20,18)
	l.Unlock()

	delay()
	l.Lock()
	value = 1
	s1.Put(8,value)
	log("p20 w:%d", value)
	countwrites(value,20,1)
	l.Unlock()

	delay()
	l.Lock()
	value = 31
	s31.Put(482,value)
	log("p20 w:%d", value)
	countwrites(value,20,31)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s29.QueryP(457,&value)
	log("p20 r:%d", value)
	countreads(value,20,29)
	l.Unlock()

	delay()
	l.Lock()
	value = 29
	s29.Put(456,value)
	log("p20 w:%d", value)
	countwrites(value,20,29)
	l.Unlock()

	delay()
	l.Lock()
	value = 20
	s20.Put(305,value)
	log("p20 w:%d", value)
	countwrites(value,20,20)
	l.Unlock()

	delay()
	l.Lock()
	value = 21
	s21.Put(331,value)
	log("p20 w:%d", value)
	countwrites(value,20,21)
	l.Unlock()

	delay()
	l.Lock()
	value = 2
	s2.Put(31,value)
	log("p20 w:%d", value)
	countwrites(value,20,2)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s26.QueryP(402,&value)
	log("p20 r:%d", value)
	countreads(value,20,26)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s3.QueryP(48,&value)
	log("p20 r:%d", value)
	countreads(value,20,3)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s31.QueryP(492,&value)
	log("p20 r:%d", value)
	countreads(value,20,31)
	l.Unlock()

	delay()
	l.Lock()
	value = 26
	s26.Put(403,value)
	log("p20 w:%d", value)
	countwrites(value,20,26)
	l.Unlock()
}


func p21() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 19
	s19.Put(298,value)
	log("p21 w:%d", value)
	countwrites(value,21,19)
	l.Unlock()

	delay()
	l.Lock()
	value = 10
	s10.Put(150,value)
	log("p21 w:%d", value)
	countwrites(value,21,10)
	l.Unlock()

	delay()
	l.Lock()
	value = 13
	s13.Put(194,value)
	log("p21 w:%d", value)
	countwrites(value,21,13)
	l.Unlock()

	delay()
	l.Lock()
	value = 12
	s12.Put(181,value)
	log("p21 w:%d", value)
	countwrites(value,21,12)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s19.QueryP(293,&value)
	log("p21 r:%d", value)
	countreads(value,21,19)
	l.Unlock()

	delay()
	l.Lock()
	value = 4
	s4.Put(64,value)
	log("p21 w:%d", value)
	countwrites(value,21,4)
	l.Unlock()

	delay()
	l.Lock()
	value = 19
	s19.Put(298,value)
	log("p21 w:%d", value)
	countwrites(value,21,19)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s13.QueryP(205,&value)
	log("p21 r:%d", value)
	countreads(value,21,13)
	l.Unlock()

	delay()
	l.Lock()
	value = 27
	s27.Put(428,value)
	log("p21 w:%d", value)
	countwrites(value,21,27)
	l.Unlock()

	delay()
	l.Lock()
	value = 21
	s21.Put(333,value)
	log("p21 w:%d", value)
	countwrites(value,21,21)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s20.QueryP(306,&value)
	log("p21 r:%d", value)
	countreads(value,21,20)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s5.QueryP(74,&value)
	log("p21 r:%d", value)
	countreads(value,21,5)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s32.QueryP(510,&value)
	log("p21 r:%d", value)
	countreads(value,21,32)
	l.Unlock()

	delay()
	l.Lock()
	value = 22
	s22.Put(340,value)
	log("p21 w:%d", value)
	countwrites(value,21,22)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s31.QueryP(494,&value)
	log("p21 r:%d", value)
	countreads(value,21,31)
	l.Unlock()

	delay()
	l.Lock()
	value = 19
	s19.Put(299,value)
	log("p21 w:%d", value)
	countwrites(value,21,19)
	l.Unlock()
}


func p22() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	s17.QueryP(271,&value)
	log("p22 r:%d", value)
	countreads(value,22,17)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s2.QueryP(26,&value)
	log("p22 r:%d", value)
	countreads(value,22,2)
	l.Unlock()

	delay()
	l.Lock()
	value = 3
	s3.Put(48,value)
	log("p22 w:%d", value)
	countwrites(value,22,3)
	l.Unlock()

	delay()
	l.Lock()
	value = 23
	s23.Put(368,value)
	log("p22 w:%d", value)
	countwrites(value,22,23)
	l.Unlock()

	delay()
	l.Lock()
	value = 7
	s7.Put(102,value)
	log("p22 w:%d", value)
	countwrites(value,22,7)
	l.Unlock()

	delay()
	l.Lock()
	value = 19
	s19.Put(292,value)
	log("p22 w:%d", value)
	countwrites(value,22,19)
	l.Unlock()

	delay()
	l.Lock()
	value = 16
	s16.Put(249,value)
	log("p22 w:%d", value)
	countwrites(value,22,16)
	l.Unlock()

	delay()
	l.Lock()
	value = 16
	s16.Put(251,value)
	log("p22 w:%d", value)
	countwrites(value,22,16)
	l.Unlock()

	delay()
	l.Lock()
	value = 22
	s22.Put(352,value)
	log("p22 w:%d", value)
	countwrites(value,22,22)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s22.QueryP(340,&value)
	log("p22 r:%d", value)
	countreads(value,22,22)
	l.Unlock()

	delay()
	l.Lock()
	value = 30
	s30.Put(466,value)
	log("p22 w:%d", value)
	countwrites(value,22,30)
	l.Unlock()

	delay()
	l.Lock()
	value = 17
	s17.Put(272,value)
	log("p22 w:%d", value)
	countwrites(value,22,17)
	l.Unlock()

	delay()
	l.Lock()
	value = 14
	s14.Put(224,value)
	log("p22 w:%d", value)
	countwrites(value,22,14)
	l.Unlock()

	delay()
	l.Lock()
	value = 13
	s13.Put(201,value)
	log("p22 w:%d", value)
	countwrites(value,22,13)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s25.QueryP(386,&value)
	log("p22 r:%d", value)
	countreads(value,22,25)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s24.QueryP(370,&value)
	log("p22 r:%d", value)
	countreads(value,22,24)
	l.Unlock()
}


func p23() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 2
	s2.Put(26,value)
	log("p23 w:%d", value)
	countwrites(value,23,2)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s4.QueryP(54,&value)
	log("p23 r:%d", value)
	countreads(value,23,4)
	l.Unlock()

	delay()
	l.Lock()
	value = 9
	s9.Put(130,value)
	log("p23 w:%d", value)
	countwrites(value,23,9)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s16.QueryP(252,&value)
	log("p23 r:%d", value)
	countreads(value,23,16)
	l.Unlock()

	delay()
	l.Lock()
	value = 21
	s21.Put(336,value)
	log("p23 w:%d", value)
	countwrites(value,23,21)
	l.Unlock()

	delay()
	l.Lock()
	value = 4
	s4.Put(55,value)
	log("p23 w:%d", value)
	countwrites(value,23,4)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s15.QueryP(237,&value)
	log("p23 r:%d", value)
	countreads(value,23,15)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s11.QueryP(162,&value)
	log("p23 r:%d", value)
	countreads(value,23,11)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s22.QueryP(340,&value)
	log("p23 r:%d", value)
	countreads(value,23,22)
	l.Unlock()

	delay()
	l.Lock()
	value = 3
	s3.Put(44,value)
	log("p23 w:%d", value)
	countwrites(value,23,3)
	l.Unlock()

	delay()
	l.Lock()
	value = 17
	s17.Put(264,value)
	log("p23 w:%d", value)
	countwrites(value,23,17)
	l.Unlock()

	delay()
	l.Lock()
	value = 12
	s12.Put(181,value)
	log("p23 w:%d", value)
	countwrites(value,23,12)
	l.Unlock()

	delay()
	l.Lock()
	value = 28
	s28.Put(448,value)
	log("p23 w:%d", value)
	countwrites(value,23,28)
	l.Unlock()

	delay()
	l.Lock()
	value = 29
	s29.Put(464,value)
	log("p23 w:%d", value)
	countwrites(value,23,29)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s8.QueryP(117,&value)
	log("p23 r:%d", value)
	countreads(value,23,8)
	l.Unlock()

	delay()
	l.Lock()
	value = 14
	s14.Put(222,value)
	log("p23 w:%d", value)
	countwrites(value,23,14)
	l.Unlock()
}


func p24() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 26
	s26.Put(416,value)
	log("p24 w:%d", value)
	countwrites(value,24,26)
	l.Unlock()

	delay()
	l.Lock()
	value = 25
	s25.Put(395,value)
	log("p24 w:%d", value)
	countwrites(value,24,25)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s20.QueryP(308,&value)
	log("p24 r:%d", value)
	countreads(value,24,20)
	l.Unlock()

	delay()
	l.Lock()
	value = 30
	s30.Put(473,value)
	log("p24 w:%d", value)
	countwrites(value,24,30)
	l.Unlock()

	delay()
	l.Lock()
	value = 4
	s4.Put(52,value)
	log("p24 w:%d", value)
	countwrites(value,24,4)
	l.Unlock()

	delay()
	l.Lock()
	value = 27
	s27.Put(422,value)
	log("p24 w:%d", value)
	countwrites(value,24,27)
	l.Unlock()

	delay()
	l.Lock()
	value = 10
	s10.Put(150,value)
	log("p24 w:%d", value)
	countwrites(value,24,10)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s31.QueryP(489,&value)
	log("p24 r:%d", value)
	countreads(value,24,31)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s14.QueryP(223,&value)
	log("p24 r:%d", value)
	countreads(value,24,14)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s15.QueryP(240,&value)
	log("p24 r:%d", value)
	countreads(value,24,15)
	l.Unlock()

	delay()
	l.Lock()
	value = 17
	s17.Put(267,value)
	log("p24 w:%d", value)
	countwrites(value,24,17)
	l.Unlock()

	delay()
	l.Lock()
	value = 25
	s25.Put(400,value)
	log("p24 w:%d", value)
	countwrites(value,24,25)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s15.QueryP(240,&value)
	log("p24 r:%d", value)
	countreads(value,24,15)
	l.Unlock()

	delay()
	l.Lock()
	value = 10
	s10.Put(156,value)
	log("p24 w:%d", value)
	countwrites(value,24,10)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s21.QueryP(323,&value)
	log("p24 r:%d", value)
	countreads(value,24,21)
	l.Unlock()

	delay()
	l.Lock()
	value = 1
	s1.Put(10,value)
	log("p24 w:%d", value)
	countwrites(value,24,1)
	l.Unlock()
}


func p25() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	s26.QueryP(407,&value)
	log("p25 r:%d", value)
	countreads(value,25,26)
	l.Unlock()

	delay()
	l.Lock()
	value = 15
	s15.Put(237,value)
	log("p25 w:%d", value)
	countwrites(value,25,15)
	l.Unlock()

	delay()
	l.Lock()
	value = 27
	s27.Put(418,value)
	log("p25 w:%d", value)
	countwrites(value,25,27)
	l.Unlock()

	delay()
	l.Lock()
	value = 5
	s5.Put(71,value)
	log("p25 w:%d", value)
	countwrites(value,25,5)
	l.Unlock()

	delay()
	l.Lock()
	value = 31
	s31.Put(489,value)
	log("p25 w:%d", value)
	countwrites(value,25,31)
	l.Unlock()

	delay()
	l.Lock()
	value = 27
	s27.Put(424,value)
	log("p25 w:%d", value)
	countwrites(value,25,27)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s18.QueryP(277,&value)
	log("p25 r:%d", value)
	countreads(value,25,18)
	l.Unlock()

	delay()
	l.Lock()
	value = 8
	s8.Put(121,value)
	log("p25 w:%d", value)
	countwrites(value,25,8)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s11.QueryP(166,&value)
	log("p25 r:%d", value)
	countreads(value,25,11)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s19.QueryP(301,&value)
	log("p25 r:%d", value)
	countreads(value,25,19)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s15.QueryP(226,&value)
	log("p25 r:%d", value)
	countreads(value,25,15)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s8.QueryP(128,&value)
	log("p25 r:%d", value)
	countreads(value,25,8)
	l.Unlock()

	delay()
	l.Lock()
	value = 23
	s23.Put(359,value)
	log("p25 w:%d", value)
	countwrites(value,25,23)
	l.Unlock()

	delay()
	l.Lock()
	value = 21
	s21.Put(324,value)
	log("p25 w:%d", value)
	countwrites(value,25,21)
	l.Unlock()

	delay()
	l.Lock()
	value = 20
	s20.Put(315,value)
	log("p25 w:%d", value)
	countwrites(value,25,20)
	l.Unlock()

	delay()
	l.Lock()
	value = 15
	s15.Put(226,value)
	log("p25 w:%d", value)
	countwrites(value,25,15)
	l.Unlock()
}


func p26() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 12
	s12.Put(184,value)
	log("p26 w:%d", value)
	countwrites(value,26,12)
	l.Unlock()

	delay()
	l.Lock()
	value = 12
	s12.Put(186,value)
	log("p26 w:%d", value)
	countwrites(value,26,12)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s14.QueryP(215,&value)
	log("p26 r:%d", value)
	countreads(value,26,14)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s21.QueryP(324,&value)
	log("p26 r:%d", value)
	countreads(value,26,21)
	l.Unlock()

	delay()
	l.Lock()
	value = 14
	s14.Put(217,value)
	log("p26 w:%d", value)
	countwrites(value,26,14)
	l.Unlock()

	delay()
	l.Lock()
	value = 6
	s6.Put(94,value)
	log("p26 w:%d", value)
	countwrites(value,26,6)
	l.Unlock()

	delay()
	l.Lock()
	value = 8
	s8.Put(126,value)
	log("p26 w:%d", value)
	countwrites(value,26,8)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s27.QueryP(424,&value)
	log("p26 r:%d", value)
	countreads(value,26,27)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s11.QueryP(162,&value)
	log("p26 r:%d", value)
	countreads(value,26,11)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s7.QueryP(106,&value)
	log("p26 r:%d", value)
	countreads(value,26,7)
	l.Unlock()

	delay()
	l.Lock()
	value = 15
	s15.Put(233,value)
	log("p26 w:%d", value)
	countwrites(value,26,15)
	l.Unlock()

	delay()
	l.Lock()
	value = 18
	s18.Put(277,value)
	log("p26 w:%d", value)
	countwrites(value,26,18)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s9.QueryP(132,&value)
	log("p26 r:%d", value)
	countreads(value,26,9)
	l.Unlock()

	delay()
	l.Lock()
	value = 9
	s9.Put(131,value)
	log("p26 w:%d", value)
	countwrites(value,26,9)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s28.QueryP(444,&value)
	log("p26 r:%d", value)
	countreads(value,26,28)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s19.QueryP(291,&value)
	log("p26 r:%d", value)
	countreads(value,26,19)
	l.Unlock()
}


func p27() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 30
	s30.Put(469,value)
	log("p27 w:%d", value)
	countwrites(value,27,30)
	l.Unlock()

	delay()
	l.Lock()
	value = 28
	s28.Put(437,value)
	log("p27 w:%d", value)
	countwrites(value,27,28)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s6.QueryP(87,&value)
	log("p27 r:%d", value)
	countreads(value,27,6)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s12.QueryP(188,&value)
	log("p27 r:%d", value)
	countreads(value,27,12)
	l.Unlock()

	delay()
	l.Lock()
	value = 31
	s31.Put(492,value)
	log("p27 w:%d", value)
	countwrites(value,27,31)
	l.Unlock()

	delay()
	l.Lock()
	value = 26
	s26.Put(406,value)
	log("p27 w:%d", value)
	countwrites(value,27,26)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s10.QueryP(152,&value)
	log("p27 r:%d", value)
	countreads(value,27,10)
	l.Unlock()

	delay()
	l.Lock()
	value = 1
	s1.Put(7,value)
	log("p27 w:%d", value)
	countwrites(value,27,1)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s13.QueryP(199,&value)
	log("p27 r:%d", value)
	countreads(value,27,13)
	l.Unlock()

	delay()
	l.Lock()
	value = 11
	s11.Put(170,value)
	log("p27 w:%d", value)
	countwrites(value,27,11)
	l.Unlock()

	delay()
	l.Lock()
	value = 3
	s3.Put(46,value)
	log("p27 w:%d", value)
	countwrites(value,27,3)
	l.Unlock()

	delay()
	l.Lock()
	value = 18
	s18.Put(280,value)
	log("p27 w:%d", value)
	countwrites(value,27,18)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s3.QueryP(36,&value)
	log("p27 r:%d", value)
	countreads(value,27,3)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s8.QueryP(115,&value)
	log("p27 r:%d", value)
	countreads(value,27,8)
	l.Unlock()

	delay()
	l.Lock()
	value = 22
	s22.Put(348,value)
	log("p27 w:%d", value)
	countwrites(value,27,22)
	l.Unlock()

	delay()
	l.Lock()
	value = 32
	s32.Put(497,value)
	log("p27 w:%d", value)
	countwrites(value,27,32)
	l.Unlock()
}


func p28() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 15
	s15.Put(230,value)
	log("p28 w:%d", value)
	countwrites(value,28,15)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s29.QueryP(461,&value)
	log("p28 r:%d", value)
	countreads(value,28,29)
	l.Unlock()

	delay()
	l.Lock()
	value = 2
	s2.Put(23,value)
	log("p28 w:%d", value)
	countwrites(value,28,2)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s20.QueryP(317,&value)
	log("p28 r:%d", value)
	countreads(value,28,20)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s21.QueryP(323,&value)
	log("p28 r:%d", value)
	countreads(value,28,21)
	l.Unlock()

	delay()
	l.Lock()
	value = 12
	s12.Put(183,value)
	log("p28 w:%d", value)
	countwrites(value,28,12)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s8.QueryP(118,&value)
	log("p28 r:%d", value)
	countreads(value,28,8)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s14.QueryP(221,&value)
	log("p28 r:%d", value)
	countreads(value,28,14)
	l.Unlock()

	delay()
	l.Lock()
	value = 14
	s14.Put(210,value)
	log("p28 w:%d", value)
	countwrites(value,28,14)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s17.QueryP(259,&value)
	log("p28 r:%d", value)
	countreads(value,28,17)
	l.Unlock()

	delay()
	l.Lock()
	value = 24
	s24.Put(370,value)
	log("p28 w:%d", value)
	countwrites(value,28,24)
	l.Unlock()

	delay()
	l.Lock()
	value = 10
	s10.Put(147,value)
	log("p28 w:%d", value)
	countwrites(value,28,10)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s12.QueryP(191,&value)
	log("p28 r:%d", value)
	countreads(value,28,12)
	l.Unlock()

	delay()
	l.Lock()
	value = 31
	s31.Put(494,value)
	log("p28 w:%d", value)
	countwrites(value,28,31)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s16.QueryP(247,&value)
	log("p28 r:%d", value)
	countreads(value,28,16)
	l.Unlock()

	delay()
	l.Lock()
	value = 9
	s9.Put(130,value)
	log("p28 w:%d", value)
	countwrites(value,28,9)
	l.Unlock()
}


func p29() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	s12.QueryP(187,&value)
	log("p29 r:%d", value)
	countreads(value,29,12)
	l.Unlock()

	delay()
	l.Lock()
	value = 24
	s24.Put(379,value)
	log("p29 w:%d", value)
	countwrites(value,29,24)
	l.Unlock()

	delay()
	l.Lock()
	value = 11
	s11.Put(170,value)
	log("p29 w:%d", value)
	countwrites(value,29,11)
	l.Unlock()

	delay()
	l.Lock()
	value = 32
	s32.Put(508,value)
	log("p29 w:%d", value)
	countwrites(value,29,32)
	l.Unlock()

	delay()
	l.Lock()
	value = 20
	s20.Put(311,value)
	log("p29 w:%d", value)
	countwrites(value,29,20)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s22.QueryP(350,&value)
	log("p29 r:%d", value)
	countreads(value,29,22)
	l.Unlock()

	delay()
	l.Lock()
	value = 29
	s29.Put(458,value)
	log("p29 w:%d", value)
	countwrites(value,29,29)
	l.Unlock()

	delay()
	l.Lock()
	value = 22
	s22.Put(346,value)
	log("p29 w:%d", value)
	countwrites(value,29,22)
	l.Unlock()

	delay()
	l.Lock()
	value = 11
	s11.Put(161,value)
	log("p29 w:%d", value)
	countwrites(value,29,11)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s19.QueryP(291,&value)
	log("p29 r:%d", value)
	countreads(value,29,19)
	l.Unlock()

	delay()
	l.Lock()
	value = 22
	s22.Put(351,value)
	log("p29 w:%d", value)
	countwrites(value,29,22)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s15.QueryP(240,&value)
	log("p29 r:%d", value)
	countreads(value,29,15)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s30.QueryP(473,&value)
	log("p29 r:%d", value)
	countreads(value,29,30)
	l.Unlock()

	delay()
	l.Lock()
	value = 9
	s9.Put(143,value)
	log("p29 w:%d", value)
	countwrites(value,29,9)
	l.Unlock()

	delay()
	l.Lock()
	value = 4
	s4.Put(50,value)
	log("p29 w:%d", value)
	countwrites(value,29,4)
	l.Unlock()

	delay()
	l.Lock()
	value = 29
	s29.Put(452,value)
	log("p29 w:%d", value)
	countwrites(value,29,29)
	l.Unlock()
}


func p30() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 6
	s6.Put(91,value)
	log("p30 w:%d", value)
	countwrites(value,30,6)
	l.Unlock()

	delay()
	l.Lock()
	value = 1
	s1.Put(2,value)
	log("p30 w:%d", value)
	countwrites(value,30,1)
	l.Unlock()

	delay()
	l.Lock()
	value = 13
	s13.Put(196,value)
	log("p30 w:%d", value)
	countwrites(value,30,13)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s8.QueryP(126,&value)
	log("p30 r:%d", value)
	countreads(value,30,8)
	l.Unlock()

	delay()
	l.Lock()
	value = 9
	s9.Put(133,value)
	log("p30 w:%d", value)
	countwrites(value,30,9)
	l.Unlock()

	delay()
	l.Lock()
	value = 22
	s22.Put(337,value)
	log("p30 w:%d", value)
	countwrites(value,30,22)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s29.QueryP(450,&value)
	log("p30 r:%d", value)
	countreads(value,30,29)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s21.QueryP(332,&value)
	log("p30 r:%d", value)
	countreads(value,30,21)
	l.Unlock()

	delay()
	l.Lock()
	value = 4
	s4.Put(53,value)
	log("p30 w:%d", value)
	countwrites(value,30,4)
	l.Unlock()

	delay()
	l.Lock()
	value = 22
	s22.Put(350,value)
	log("p30 w:%d", value)
	countwrites(value,30,22)
	l.Unlock()

	delay()
	l.Lock()
	value = 1
	s1.Put(11,value)
	log("p30 w:%d", value)
	countwrites(value,30,1)
	l.Unlock()

	delay()
	l.Lock()
	value = 11
	s11.Put(176,value)
	log("p30 w:%d", value)
	countwrites(value,30,11)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s7.QueryP(105,&value)
	log("p30 r:%d", value)
	countreads(value,30,7)
	l.Unlock()

	delay()
	l.Lock()
	value = 13
	s13.Put(202,value)
	log("p30 w:%d", value)
	countwrites(value,30,13)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s23.QueryP(368,&value)
	log("p30 r:%d", value)
	countreads(value,30,23)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s6.QueryP(84,&value)
	log("p30 r:%d", value)
	countreads(value,30,6)
	l.Unlock()
}


func p31() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	s28.QueryP(440,&value)
	log("p31 r:%d", value)
	countreads(value,31,28)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s6.QueryP(84,&value)
	log("p31 r:%d", value)
	countreads(value,31,6)
	l.Unlock()

	delay()
	l.Lock()
	value = 3
	s3.Put(41,value)
	log("p31 w:%d", value)
	countwrites(value,31,3)
	l.Unlock()

	delay()
	l.Lock()
	value = 4
	s4.Put(52,value)
	log("p31 w:%d", value)
	countwrites(value,31,4)
	l.Unlock()

	delay()
	l.Lock()
	value = 27
	s27.Put(426,value)
	log("p31 w:%d", value)
	countwrites(value,31,27)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s15.QueryP(235,&value)
	log("p31 r:%d", value)
	countreads(value,31,15)
	l.Unlock()

	delay()
	l.Lock()
	value = 30
	s30.Put(469,value)
	log("p31 w:%d", value)
	countwrites(value,31,30)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s6.QueryP(94,&value)
	log("p31 r:%d", value)
	countreads(value,31,6)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s11.QueryP(166,&value)
	log("p31 r:%d", value)
	countreads(value,31,11)
	l.Unlock()

	delay()
	l.Lock()
	value = 17
	s17.Put(267,value)
	log("p31 w:%d", value)
	countwrites(value,31,17)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s6.QueryP(89,&value)
	log("p31 r:%d", value)
	countreads(value,31,6)
	l.Unlock()

	delay()
	l.Lock()
	value = 20
	s20.Put(320,value)
	log("p31 w:%d", value)
	countwrites(value,31,20)
	l.Unlock()

	delay()
	l.Lock()
	value = 5
	s5.Put(77,value)
	log("p31 w:%d", value)
	countwrites(value,31,5)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s26.QueryP(412,&value)
	log("p31 r:%d", value)
	countreads(value,31,26)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s15.QueryP(234,&value)
	log("p31 r:%d", value)
	countreads(value,31,15)
	l.Unlock()

	delay()
	l.Lock()
	value = 23
	s23.Put(368,value)
	log("p31 w:%d", value)
	countwrites(value,31,23)
	l.Unlock()
}


func p32() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 22
	s22.Put(350,value)
	log("p32 w:%d", value)
	countwrites(value,32,22)
	l.Unlock()

	delay()
	l.Lock()
	value = 28
	s28.Put(446,value)
	log("p32 w:%d", value)
	countwrites(value,32,28)
	l.Unlock()

	delay()
	l.Lock()
	value = 9
	s9.Put(138,value)
	log("p32 w:%d", value)
	countwrites(value,32,9)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s19.QueryP(302,&value)
	log("p32 r:%d", value)
	countreads(value,32,19)
	l.Unlock()

	delay()
	l.Lock()
	value = 19
	s19.Put(293,value)
	log("p32 w:%d", value)
	countwrites(value,32,19)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s1.QueryP(5,&value)
	log("p32 r:%d", value)
	countreads(value,32,1)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s27.QueryP(430,&value)
	log("p32 r:%d", value)
	countreads(value,32,27)
	l.Unlock()

	delay()
	l.Lock()
	value = 25
	s25.Put(385,value)
	log("p32 w:%d", value)
	countwrites(value,32,25)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s22.QueryP(347,&value)
	log("p32 r:%d", value)
	countreads(value,32,22)
	l.Unlock()

	delay()
	l.Lock()
	value = 12
	s12.Put(179,value)
	log("p32 w:%d", value)
	countwrites(value,32,12)
	l.Unlock()

	delay()
	l.Lock()
	value = 17
	s17.Put(261,value)
	log("p32 w:%d", value)
	countwrites(value,32,17)
	l.Unlock()

	delay()
	l.Lock()
	value = 2
	s2.Put(20,value)
	log("p32 w:%d", value)
	countwrites(value,32,2)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s25.QueryP(391,&value)
	log("p32 r:%d", value)
	countreads(value,32,25)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s29.QueryP(449,&value)
	log("p32 r:%d", value)
	countreads(value,32,29)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s25.QueryP(393,&value)
	log("p32 r:%d", value)
	countreads(value,32,25)
	l.Unlock()

	delay()
	l.Lock()
	value = 13
	s13.Put(201,value)
	log("p32 w:%d", value)
	countwrites(value,32,13)
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
