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
	s14.QueryP(220,&value)
	log("p1 r:%d", value)
	countreads(value,1,14)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s27.QueryP(418,&value)
	log("p1 r:%d", value)
	countreads(value,1,27)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s28.QueryP(435,&value)
	log("p1 r:%d", value)
	countreads(value,1,28)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s11.QueryP(172,&value)
	log("p1 r:%d", value)
	countreads(value,1,11)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s23.QueryP(358,&value)
	log("p1 r:%d", value)
	countreads(value,1,23)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s24.QueryP(379,&value)
	log("p1 r:%d", value)
	countreads(value,1,24)
	l.Unlock()

	delay()
	l.Lock()
	value = 17
	s17.Put(263,value)
	log("p1 w:%d", value)
	countwrites(value,1,17)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s21.QueryP(335,&value)
	log("p1 r:%d", value)
	countreads(value,1,21)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s15.QueryP(227,&value)
	log("p1 r:%d", value)
	countreads(value,1,15)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s18.QueryP(284,&value)
	log("p1 r:%d", value)
	countreads(value,1,18)
	l.Unlock()

	delay()
	l.Lock()
	value = 11
	s11.Put(165,value)
	log("p1 w:%d", value)
	countwrites(value,1,11)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s2.QueryP(30,&value)
	log("p1 r:%d", value)
	countreads(value,1,2)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s24.QueryP(374,&value)
	log("p1 r:%d", value)
	countreads(value,1,24)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s19.QueryP(301,&value)
	log("p1 r:%d", value)
	countreads(value,1,19)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s13.QueryP(204,&value)
	log("p1 r:%d", value)
	countreads(value,1,13)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s26.QueryP(416,&value)
	log("p1 r:%d", value)
	countreads(value,1,26)
	l.Unlock()
}

func p2() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	s16.QueryP(256,&value)
	log("p2 r:%d", value)
	countreads(value,2,16)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s4.QueryP(52,&value)
	log("p2 r:%d", value)
	countreads(value,2,4)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s32.QueryP(505,&value)
	log("p2 r:%d", value)
	countreads(value,2,32)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s27.QueryP(417,&value)
	log("p2 r:%d", value)
	countreads(value,2,27)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s21.QueryP(330,&value)
	log("p2 r:%d", value)
	countreads(value,2,21)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s6.QueryP(96,&value)
	log("p2 r:%d", value)
	countreads(value,2,6)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s28.QueryP(435,&value)
	log("p2 r:%d", value)
	countreads(value,2,28)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s24.QueryP(374,&value)
	log("p2 r:%d", value)
	countreads(value,2,24)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s30.QueryP(475,&value)
	log("p2 r:%d", value)
	countreads(value,2,30)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s1.QueryP(5,&value)
	log("p2 r:%d", value)
	countreads(value,2,1)
	l.Unlock()

	delay()
	l.Lock()
	value = 2
	s2.Put(20,value)
	log("p2 w:%d", value)
	countwrites(value,2,2)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s16.QueryP(243,&value)
	log("p2 r:%d", value)
	countreads(value,2,16)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s30.QueryP(465,&value)
	log("p2 r:%d", value)
	countreads(value,2,30)
	l.Unlock()

	delay()
	l.Lock()
	value = 2
	s2.Put(26,value)
	log("p2 w:%d", value)
	countwrites(value,2,2)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s10.QueryP(152,&value)
	log("p2 r:%d", value)
	countreads(value,2,10)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s26.QueryP(416,&value)
	log("p2 r:%d", value)
	countreads(value,2,26)
	l.Unlock()
}

func p3() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	s27.QueryP(424,&value)
	log("p3 r:%d", value)
	countreads(value,3,27)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s27.QueryP(429,&value)
	log("p3 r:%d", value)
	countreads(value,3,27)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s15.QueryP(230,&value)
	log("p3 r:%d", value)
	countreads(value,3,15)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s14.QueryP(222,&value)
	log("p3 r:%d", value)
	countreads(value,3,14)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s2.QueryP(24,&value)
	log("p3 r:%d", value)
	countreads(value,3,2)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s8.QueryP(127,&value)
	log("p3 r:%d", value)
	countreads(value,3,8)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s17.QueryP(261,&value)
	log("p3 r:%d", value)
	countreads(value,3,17)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s32.QueryP(499,&value)
	log("p3 r:%d", value)
	countreads(value,3,32)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s8.QueryP(118,&value)
	log("p3 r:%d", value)
	countreads(value,3,8)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s17.QueryP(266,&value)
	log("p3 r:%d", value)
	countreads(value,3,17)
	l.Unlock()

	delay()
	l.Lock()
	value = 16
	s16.Put(246,value)
	log("p3 w:%d", value)
	countwrites(value,3,16)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s4.QueryP(49,&value)
	log("p3 r:%d", value)
	countreads(value,3,4)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s1.QueryP(1,&value)
	log("p3 r:%d", value)
	countreads(value,3,1)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s18.QueryP(275,&value)
	log("p3 r:%d", value)
	countreads(value,3,18)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s10.QueryP(149,&value)
	log("p3 r:%d", value)
	countreads(value,3,10)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s11.QueryP(161,&value)
	log("p3 r:%d", value)
	countreads(value,3,11)
	l.Unlock()
}

func p4() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	s31.QueryP(494,&value)
	log("p4 r:%d", value)
	countreads(value,4,31)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s24.QueryP(371,&value)
	log("p4 r:%d", value)
	countreads(value,4,24)
	l.Unlock()

	delay()
	l.Lock()
	value = 28
	s28.Put(442,value)
	log("p4 w:%d", value)
	countwrites(value,4,28)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s10.QueryP(151,&value)
	log("p4 r:%d", value)
	countreads(value,4,10)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s26.QueryP(405,&value)
	log("p4 r:%d", value)
	countreads(value,4,26)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s10.QueryP(156,&value)
	log("p4 r:%d", value)
	countreads(value,4,10)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s12.QueryP(179,&value)
	log("p4 r:%d", value)
	countreads(value,4,12)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s19.QueryP(291,&value)
	log("p4 r:%d", value)
	countreads(value,4,19)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s6.QueryP(83,&value)
	log("p4 r:%d", value)
	countreads(value,4,6)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s10.QueryP(147,&value)
	log("p4 r:%d", value)
	countreads(value,4,10)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s3.QueryP(48,&value)
	log("p4 r:%d", value)
	countreads(value,4,3)
	l.Unlock()

	delay()
	l.Lock()
	value = 32
	s32.Put(508,value)
	log("p4 w:%d", value)
	countwrites(value,4,32)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s15.QueryP(238,&value)
	log("p4 r:%d", value)
	countreads(value,4,15)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s15.QueryP(231,&value)
	log("p4 r:%d", value)
	countreads(value,4,15)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s31.QueryP(494,&value)
	log("p4 r:%d", value)
	countreads(value,4,31)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s22.QueryP(340,&value)
	log("p4 r:%d", value)
	countreads(value,4,22)
	l.Unlock()
}

func p5() {
	defer wg.Done()
	var value int

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
	s27.QueryP(431,&value)
	log("p5 r:%d", value)
	countreads(value,5,27)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s5.QueryP(75,&value)
	log("p5 r:%d", value)
	countreads(value,5,5)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s12.QueryP(183,&value)
	log("p5 r:%d", value)
	countreads(value,5,12)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s21.QueryP(327,&value)
	log("p5 r:%d", value)
	countreads(value,5,21)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s13.QueryP(193,&value)
	log("p5 r:%d", value)
	countreads(value,5,13)
	l.Unlock()

	delay()
	l.Lock()
	value = 14
	s14.Put(210,value)
	log("p5 w:%d", value)
	countwrites(value,5,14)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s9.QueryP(137,&value)
	log("p5 r:%d", value)
	countreads(value,5,9)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s28.QueryP(436,&value)
	log("p5 r:%d", value)
	countreads(value,5,28)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s29.QueryP(459,&value)
	log("p5 r:%d", value)
	countreads(value,5,29)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s6.QueryP(87,&value)
	log("p5 r:%d", value)
	countreads(value,5,6)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s12.QueryP(187,&value)
	log("p5 r:%d", value)
	countreads(value,5,12)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s6.QueryP(94,&value)
	log("p5 r:%d", value)
	countreads(value,5,6)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s21.QueryP(324,&value)
	log("p5 r:%d", value)
	countreads(value,5,21)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s18.QueryP(284,&value)
	log("p5 r:%d", value)
	countreads(value,5,18)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s8.QueryP(118,&value)
	log("p5 r:%d", value)
	countreads(value,5,8)
	l.Unlock()
}

func p6() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	s29.QueryP(451,&value)
	log("p6 r:%d", value)
	countreads(value,6,29)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s29.QueryP(449,&value)
	log("p6 r:%d", value)
	countreads(value,6,29)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s9.QueryP(130,&value)
	log("p6 r:%d", value)
	countreads(value,6,9)
	l.Unlock()

	delay()
	l.Lock()
	value = 4
	s4.Put(62,value)
	log("p6 w:%d", value)
	countwrites(value,6,4)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s14.QueryP(223,&value)
	log("p6 r:%d", value)
	countreads(value,6,14)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s4.QueryP(56,&value)
	log("p6 r:%d", value)
	countreads(value,6,4)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s18.QueryP(288,&value)
	log("p6 r:%d", value)
	countreads(value,6,18)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s12.QueryP(177,&value)
	log("p6 r:%d", value)
	countreads(value,6,12)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s14.QueryP(215,&value)
	log("p6 r:%d", value)
	countreads(value,6,14)
	l.Unlock()

	delay()
	l.Lock()
	value = 14
	s14.Put(219,value)
	log("p6 w:%d", value)
	countwrites(value,6,14)
	l.Unlock()

	delay()
	l.Lock()
	value = 22
	s22.Put(347,value)
	log("p6 w:%d", value)
	countwrites(value,6,22)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s3.QueryP(36,&value)
	log("p6 r:%d", value)
	countreads(value,6,3)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s32.QueryP(509,&value)
	log("p6 r:%d", value)
	countreads(value,6,32)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s23.QueryP(359,&value)
	log("p6 r:%d", value)
	countreads(value,6,23)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s7.QueryP(105,&value)
	log("p6 r:%d", value)
	countreads(value,6,7)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s15.QueryP(231,&value)
	log("p6 r:%d", value)
	countreads(value,6,15)
	l.Unlock()
}

func p7() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	s27.QueryP(419,&value)
	log("p7 r:%d", value)
	countreads(value,7,27)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s16.QueryP(250,&value)
	log("p7 r:%d", value)
	countreads(value,7,16)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s5.QueryP(80,&value)
	log("p7 r:%d", value)
	countreads(value,7,5)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s31.QueryP(489,&value)
	log("p7 r:%d", value)
	countreads(value,7,31)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s20.QueryP(317,&value)
	log("p7 r:%d", value)
	countreads(value,7,20)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s22.QueryP(341,&value)
	log("p7 r:%d", value)
	countreads(value,7,22)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s2.QueryP(19,&value)
	log("p7 r:%d", value)
	countreads(value,7,2)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s11.QueryP(171,&value)
	log("p7 r:%d", value)
	countreads(value,7,11)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s20.QueryP(312,&value)
	log("p7 r:%d", value)
	countreads(value,7,20)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s21.QueryP(326,&value)
	log("p7 r:%d", value)
	countreads(value,7,21)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s9.QueryP(131,&value)
	log("p7 r:%d", value)
	countreads(value,7,9)
	l.Unlock()

	delay()
	l.Lock()
	value = 13
	s13.Put(197,value)
	log("p7 w:%d", value)
	countwrites(value,7,13)
	l.Unlock()

	delay()
	l.Lock()
	value = 10
	s10.Put(148,value)
	log("p7 w:%d", value)
	countwrites(value,7,10)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s32.QueryP(497,&value)
	log("p7 r:%d", value)
	countreads(value,7,32)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s13.QueryP(199,&value)
	log("p7 r:%d", value)
	countreads(value,7,13)
	l.Unlock()

	delay()
	l.Lock()
	value = 16
	s16.Put(241,value)
	log("p7 w:%d", value)
	countwrites(value,7,16)
	l.Unlock()
}

func p8() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	s16.QueryP(243,&value)
	log("p8 r:%d", value)
	countreads(value,8,16)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s22.QueryP(338,&value)
	log("p8 r:%d", value)
	countreads(value,8,22)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s30.QueryP(466,&value)
	log("p8 r:%d", value)
	countreads(value,8,30)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s24.QueryP(371,&value)
	log("p8 r:%d", value)
	countreads(value,8,24)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s20.QueryP(308,&value)
	log("p8 r:%d", value)
	countreads(value,8,20)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s28.QueryP(437,&value)
	log("p8 r:%d", value)
	countreads(value,8,28)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s10.QueryP(145,&value)
	log("p8 r:%d", value)
	countreads(value,8,10)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s19.QueryP(289,&value)
	log("p8 r:%d", value)
	countreads(value,8,19)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s2.QueryP(17,&value)
	log("p8 r:%d", value)
	countreads(value,8,2)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s21.QueryP(335,&value)
	log("p8 r:%d", value)
	countreads(value,8,21)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s23.QueryP(355,&value)
	log("p8 r:%d", value)
	countreads(value,8,23)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s2.QueryP(31,&value)
	log("p8 r:%d", value)
	countreads(value,8,2)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s13.QueryP(197,&value)
	log("p8 r:%d", value)
	countreads(value,8,13)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s6.QueryP(82,&value)
	log("p8 r:%d", value)
	countreads(value,8,6)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s22.QueryP(345,&value)
	log("p8 r:%d", value)
	countreads(value,8,22)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s25.QueryP(393,&value)
	log("p8 r:%d", value)
	countreads(value,8,25)
	l.Unlock()
}

func p9() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	s2.QueryP(17,&value)
	log("p9 r:%d", value)
	countreads(value,9,2)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s10.QueryP(159,&value)
	log("p9 r:%d", value)
	countreads(value,9,10)
	l.Unlock()

	delay()
	l.Lock()
	value = 24
	s24.Put(370,value)
	log("p9 w:%d", value)
	countwrites(value,9,24)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s19.QueryP(300,&value)
	log("p9 r:%d", value)
	countreads(value,9,19)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s7.QueryP(99,&value)
	log("p9 r:%d", value)
	countreads(value,9,7)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s29.QueryP(455,&value)
	log("p9 r:%d", value)
	countreads(value,9,29)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s7.QueryP(108,&value)
	log("p9 r:%d", value)
	countreads(value,9,7)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s18.QueryP(285,&value)
	log("p9 r:%d", value)
	countreads(value,9,18)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s29.QueryP(461,&value)
	log("p9 r:%d", value)
	countreads(value,9,29)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s8.QueryP(116,&value)
	log("p9 r:%d", value)
	countreads(value,9,8)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s3.QueryP(48,&value)
	log("p9 r:%d", value)
	countreads(value,9,3)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s29.QueryP(459,&value)
	log("p9 r:%d", value)
	countreads(value,9,29)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s3.QueryP(40,&value)
	log("p9 r:%d", value)
	countreads(value,9,3)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s12.QueryP(182,&value)
	log("p9 r:%d", value)
	countreads(value,9,12)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s17.QueryP(263,&value)
	log("p9 r:%d", value)
	countreads(value,9,17)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s27.QueryP(428,&value)
	log("p9 r:%d", value)
	countreads(value,9,27)
	l.Unlock()
}


func p10() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	s10.QueryP(151,&value)
	log("p10 r:%d", value)
	countreads(value,10,10)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s9.QueryP(135,&value)
	log("p10 r:%d", value)
	countreads(value,10,9)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s27.QueryP(424,&value)
	log("p10 r:%d", value)
	countreads(value,10,27)
	l.Unlock()

	delay()
	l.Lock()
	value = 17
	s17.Put(267,value)
	log("p10 w:%d", value)
	countwrites(value,10,17)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s27.QueryP(424,&value)
	log("p10 r:%d", value)
	countreads(value,10,27)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s13.QueryP(207,&value)
	log("p10 r:%d", value)
	countreads(value,10,13)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s29.QueryP(456,&value)
	log("p10 r:%d", value)
	countreads(value,10,29)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s30.QueryP(470,&value)
	log("p10 r:%d", value)
	countreads(value,10,30)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s1.QueryP(6,&value)
	log("p10 r:%d", value)
	countreads(value,10,1)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s22.QueryP(337,&value)
	log("p10 r:%d", value)
	countreads(value,10,22)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s15.QueryP(238,&value)
	log("p10 r:%d", value)
	countreads(value,10,15)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s23.QueryP(362,&value)
	log("p10 r:%d", value)
	countreads(value,10,23)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s19.QueryP(295,&value)
	log("p10 r:%d", value)
	countreads(value,10,19)
	l.Unlock()

	delay()
	l.Lock()
	value = 31
	s31.Put(486,value)
	log("p10 w:%d", value)
	countwrites(value,10,31)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s1.QueryP(11,&value)
	log("p10 r:%d", value)
	countreads(value,10,1)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s8.QueryP(113,&value)
	log("p10 r:%d", value)
	countreads(value,10,8)
	l.Unlock()
}


func p11() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	s25.QueryP(399,&value)
	log("p11 r:%d", value)
	countreads(value,11,25)
	l.Unlock()

	delay()
	l.Lock()
	value = 26
	s26.Put(408,value)
	log("p11 w:%d", value)
	countwrites(value,11,26)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s21.QueryP(322,&value)
	log("p11 r:%d", value)
	countreads(value,11,21)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s15.QueryP(240,&value)
	log("p11 r:%d", value)
	countreads(value,11,15)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s22.QueryP(346,&value)
	log("p11 r:%d", value)
	countreads(value,11,22)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s18.QueryP(281,&value)
	log("p11 r:%d", value)
	countreads(value,11,18)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s7.QueryP(104,&value)
	log("p11 r:%d", value)
	countreads(value,11,7)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s1.QueryP(10,&value)
	log("p11 r:%d", value)
	countreads(value,11,1)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s10.QueryP(153,&value)
	log("p11 r:%d", value)
	countreads(value,11,10)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s31.QueryP(494,&value)
	log("p11 r:%d", value)
	countreads(value,11,31)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s29.QueryP(449,&value)
	log("p11 r:%d", value)
	countreads(value,11,29)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s7.QueryP(100,&value)
	log("p11 r:%d", value)
	countreads(value,11,7)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s1.QueryP(7,&value)
	log("p11 r:%d", value)
	countreads(value,11,1)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s12.QueryP(179,&value)
	log("p11 r:%d", value)
	countreads(value,11,12)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s20.QueryP(310,&value)
	log("p11 r:%d", value)
	countreads(value,11,20)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s27.QueryP(430,&value)
	log("p11 r:%d", value)
	countreads(value,11,27)
	l.Unlock()
}


func p12() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 9
	s9.Put(133,value)
	log("p12 w:%d", value)
	countwrites(value,12,9)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s9.QueryP(133,&value)
	log("p12 r:%d", value)
	countreads(value,12,9)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s19.QueryP(304,&value)
	log("p12 r:%d", value)
	countreads(value,12,19)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s22.QueryP(339,&value)
	log("p12 r:%d", value)
	countreads(value,12,22)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s28.QueryP(433,&value)
	log("p12 r:%d", value)
	countreads(value,12,28)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s5.QueryP(77,&value)
	log("p12 r:%d", value)
	countreads(value,12,5)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s1.QueryP(3,&value)
	log("p12 r:%d", value)
	countreads(value,12,1)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s1.QueryP(6,&value)
	log("p12 r:%d", value)
	countreads(value,12,1)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s12.QueryP(191,&value)
	log("p12 r:%d", value)
	countreads(value,12,12)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s10.QueryP(145,&value)
	log("p12 r:%d", value)
	countreads(value,12,10)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s30.QueryP(469,&value)
	log("p12 r:%d", value)
	countreads(value,12,30)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s1.QueryP(13,&value)
	log("p12 r:%d", value)
	countreads(value,12,1)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s15.QueryP(239,&value)
	log("p12 r:%d", value)
	countreads(value,12,15)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s23.QueryP(359,&value)
	log("p12 r:%d", value)
	countreads(value,12,23)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s20.QueryP(312,&value)
	log("p12 r:%d", value)
	countreads(value,12,20)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s11.QueryP(162,&value)
	log("p12 r:%d", value)
	countreads(value,12,11)
	l.Unlock()
}


func p13() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	s22.QueryP(338,&value)
	log("p13 r:%d", value)
	countreads(value,13,22)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s24.QueryP(383,&value)
	log("p13 r:%d", value)
	countreads(value,13,24)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s3.QueryP(33,&value)
	log("p13 r:%d", value)
	countreads(value,13,3)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s1.QueryP(2,&value)
	log("p13 r:%d", value)
	countreads(value,13,1)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s2.QueryP(24,&value)
	log("p13 r:%d", value)
	countreads(value,13,2)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s20.QueryP(316,&value)
	log("p13 r:%d", value)
	countreads(value,13,20)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s20.QueryP(311,&value)
	log("p13 r:%d", value)
	countreads(value,13,20)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s24.QueryP(371,&value)
	log("p13 r:%d", value)
	countreads(value,13,24)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s10.QueryP(155,&value)
	log("p13 r:%d", value)
	countreads(value,13,10)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s19.QueryP(293,&value)
	log("p13 r:%d", value)
	countreads(value,13,19)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s19.QueryP(296,&value)
	log("p13 r:%d", value)
	countreads(value,13,19)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s31.QueryP(485,&value)
	log("p13 r:%d", value)
	countreads(value,13,31)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s9.QueryP(130,&value)
	log("p13 r:%d", value)
	countreads(value,13,9)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s6.QueryP(88,&value)
	log("p13 r:%d", value)
	countreads(value,13,6)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s26.QueryP(407,&value)
	log("p13 r:%d", value)
	countreads(value,13,26)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s4.QueryP(53,&value)
	log("p13 r:%d", value)
	countreads(value,13,4)
	l.Unlock()
}

func p14() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	s30.QueryP(470,&value)
	log("p14 r:%d", value)
	countreads(value,14,30)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s3.QueryP(44,&value)
	log("p14 r:%d", value)
	countreads(value,14,3)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s13.QueryP(194,&value)
	log("p14 r:%d", value)
	countreads(value,14,13)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s16.QueryP(245,&value)
	log("p14 r:%d", value)
	countreads(value,14,16)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s24.QueryP(376,&value)
	log("p14 r:%d", value)
	countreads(value,14,24)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s9.QueryP(141,&value)
	log("p14 r:%d", value)
	countreads(value,14,9)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s30.QueryP(470,&value)
	log("p14 r:%d", value)
	countreads(value,14,30)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s20.QueryP(320,&value)
	log("p14 r:%d", value)
	countreads(value,14,20)
	l.Unlock()

	delay()
	l.Lock()
	value = 7
	s7.Put(107,value)
	log("p14 w:%d", value)
	countwrites(value,14,7)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s13.QueryP(201,&value)
	log("p14 r:%d", value)
	countreads(value,14,13)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s11.QueryP(176,&value)
	log("p14 r:%d", value)
	countreads(value,14,11)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s29.QueryP(455,&value)
	log("p14 r:%d", value)
	countreads(value,14,29)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s28.QueryP(434,&value)
	log("p14 r:%d", value)
	countreads(value,14,28)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s6.QueryP(83,&value)
	log("p14 r:%d", value)
	countreads(value,14,6)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s20.QueryP(316,&value)
	log("p14 r:%d", value)
	countreads(value,14,20)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s13.QueryP(194,&value)
	log("p14 r:%d", value)
	countreads(value,14,13)
	l.Unlock()
}


func p15() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	s19.QueryP(296,&value)
	log("p15 r:%d", value)
	countreads(value,15,19)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s18.QueryP(288,&value)
	log("p15 r:%d", value)
	countreads(value,15,18)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s9.QueryP(129,&value)
	log("p15 r:%d", value)
	countreads(value,15,9)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s18.QueryP(275,&value)
	log("p15 r:%d", value)
	countreads(value,15,18)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s10.QueryP(160,&value)
	log("p15 r:%d", value)
	countreads(value,15,10)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s8.QueryP(114,&value)
	log("p15 r:%d", value)
	countreads(value,15,8)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s13.QueryP(208,&value)
	log("p15 r:%d", value)
	countreads(value,15,13)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s2.QueryP(28,&value)
	log("p15 r:%d", value)
	countreads(value,15,2)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s7.QueryP(110,&value)
	log("p15 r:%d", value)
	countreads(value,15,7)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s6.QueryP(86,&value)
	log("p15 r:%d", value)
	countreads(value,15,6)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s15.QueryP(236,&value)
	log("p15 r:%d", value)
	countreads(value,15,15)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s10.QueryP(151,&value)
	log("p15 r:%d", value)
	countreads(value,15,10)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s11.QueryP(172,&value)
	log("p15 r:%d", value)
	countreads(value,15,11)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s19.QueryP(297,&value)
	log("p15 r:%d", value)
	countreads(value,15,19)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s2.QueryP(20,&value)
	log("p15 r:%d", value)
	countreads(value,15,2)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s26.QueryP(407,&value)
	log("p15 r:%d", value)
	countreads(value,15,26)
	l.Unlock()
}


func p16() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 29
	s29.Put(451,value)
	log("p16 w:%d", value)
	countwrites(value,16,29)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s7.QueryP(112,&value)
	log("p16 r:%d", value)
	countreads(value,16,7)
	l.Unlock()

	delay()
	l.Lock()
	value = 27
	s27.Put(429,value)
	log("p16 w:%d", value)
	countwrites(value,16,27)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s16.QueryP(249,&value)
	log("p16 r:%d", value)
	countreads(value,16,16)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s29.QueryP(459,&value)
	log("p16 r:%d", value)
	countreads(value,16,29)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s10.QueryP(151,&value)
	log("p16 r:%d", value)
	countreads(value,16,10)
	l.Unlock()

	delay()
	l.Lock()
	value = 15
	s15.Put(236,value)
	log("p16 w:%d", value)
	countwrites(value,16,15)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s1.QueryP(7,&value)
	log("p16 r:%d", value)
	countreads(value,16,1)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s21.QueryP(330,&value)
	log("p16 r:%d", value)
	countreads(value,16,21)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s31.QueryP(490,&value)
	log("p16 r:%d", value)
	countreads(value,16,31)
	l.Unlock()

	delay()
	l.Lock()
	value = 15
	s15.Put(225,value)
	log("p16 w:%d", value)
	countwrites(value,16,15)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s11.QueryP(175,&value)
	log("p16 r:%d", value)
	countreads(value,16,11)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s24.QueryP(371,&value)
	log("p16 r:%d", value)
	countreads(value,16,24)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s3.QueryP(44,&value)
	log("p16 r:%d", value)
	countreads(value,16,3)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s14.QueryP(214,&value)
	log("p16 r:%d", value)
	countreads(value,16,14)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s27.QueryP(432,&value)
	log("p16 r:%d", value)
	countreads(value,16,27)
	l.Unlock()
}


func p17() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	s13.QueryP(203,&value)
	log("p17 r:%d", value)
	countreads(value,17,13)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s6.QueryP(85,&value)
	log("p17 r:%d", value)
	countreads(value,17,6)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s5.QueryP(65,&value)
	log("p17 r:%d", value)
	countreads(value,17,5)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s12.QueryP(179,&value)
	log("p17 r:%d", value)
	countreads(value,17,12)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s13.QueryP(204,&value)
	log("p17 r:%d", value)
	countreads(value,17,13)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s17.QueryP(267,&value)
	log("p17 r:%d", value)
	countreads(value,17,17)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s27.QueryP(430,&value)
	log("p17 r:%d", value)
	countreads(value,17,27)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s25.QueryP(388,&value)
	log("p17 r:%d", value)
	countreads(value,17,25)
	l.Unlock()

	delay()
	l.Lock()
	value = 21
	s21.Put(335,value)
	log("p17 w:%d", value)
	countwrites(value,17,21)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s20.QueryP(313,&value)
	log("p17 r:%d", value)
	countreads(value,17,20)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s19.QueryP(295,&value)
	log("p17 r:%d", value)
	countreads(value,17,19)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s24.QueryP(374,&value)
	log("p17 r:%d", value)
	countreads(value,17,24)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s20.QueryP(315,&value)
	log("p17 r:%d", value)
	countreads(value,17,20)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s29.QueryP(452,&value)
	log("p17 r:%d", value)
	countreads(value,17,29)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s31.QueryP(496,&value)
	log("p17 r:%d", value)
	countreads(value,17,31)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s1.QueryP(14,&value)
	log("p17 r:%d", value)
	countreads(value,17,1)
	l.Unlock()
}


func p18() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	s5.QueryP(75,&value)
	log("p18 r:%d", value)
	countreads(value,18,5)
	l.Unlock()

	delay()
	l.Lock()
	value = 14
	s14.Put(214,value)
	log("p18 w:%d", value)
	countwrites(value,18,14)
	l.Unlock()

	delay()
	l.Lock()
	value = 20
	s20.Put(310,value)
	log("p18 w:%d", value)
	countwrites(value,18,20)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s5.QueryP(79,&value)
	log("p18 r:%d", value)
	countreads(value,18,5)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s23.QueryP(362,&value)
	log("p18 r:%d", value)
	countreads(value,18,23)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s13.QueryP(206,&value)
	log("p18 r:%d", value)
	countreads(value,18,13)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s5.QueryP(67,&value)
	log("p18 r:%d", value)
	countreads(value,18,5)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s4.QueryP(59,&value)
	log("p18 r:%d", value)
	countreads(value,18,4)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s28.QueryP(437,&value)
	log("p18 r:%d", value)
	countreads(value,18,28)
	l.Unlock()

	delay()
	l.Lock()
	value = 2
	s2.Put(29,value)
	log("p18 w:%d", value)
	countwrites(value,18,2)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s28.QueryP(438,&value)
	log("p18 r:%d", value)
	countreads(value,18,28)
	l.Unlock()

	delay()
	l.Lock()
	value = 26
	s26.Put(405,value)
	log("p18 w:%d", value)
	countwrites(value,18,26)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s23.QueryP(363,&value)
	log("p18 r:%d", value)
	countreads(value,18,23)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s20.QueryP(317,&value)
	log("p18 r:%d", value)
	countreads(value,18,20)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s29.QueryP(458,&value)
	log("p18 r:%d", value)
	countreads(value,18,29)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s18.QueryP(276,&value)
	log("p18 r:%d", value)
	countreads(value,18,18)
	l.Unlock()
}


func p19() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	s31.QueryP(481,&value)
	log("p19 r:%d", value)
	countreads(value,19,31)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s5.QueryP(65,&value)
	log("p19 r:%d", value)
	countreads(value,19,5)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s2.QueryP(19,&value)
	log("p19 r:%d", value)
	countreads(value,19,2)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s30.QueryP(468,&value)
	log("p19 r:%d", value)
	countreads(value,19,30)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s18.QueryP(283,&value)
	log("p19 r:%d", value)
	countreads(value,19,18)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s13.QueryP(206,&value)
	log("p19 r:%d", value)
	countreads(value,19,13)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s8.QueryP(127,&value)
	log("p19 r:%d", value)
	countreads(value,19,8)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s28.QueryP(445,&value)
	log("p19 r:%d", value)
	countreads(value,19,28)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s29.QueryP(449,&value)
	log("p19 r:%d", value)
	countreads(value,19,29)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s22.QueryP(343,&value)
	log("p19 r:%d", value)
	countreads(value,19,22)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s29.QueryP(462,&value)
	log("p19 r:%d", value)
	countreads(value,19,29)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s6.QueryP(87,&value)
	log("p19 r:%d", value)
	countreads(value,19,6)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s30.QueryP(468,&value)
	log("p19 r:%d", value)
	countreads(value,19,30)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s13.QueryP(193,&value)
	log("p19 r:%d", value)
	countreads(value,19,13)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s5.QueryP(65,&value)
	log("p19 r:%d", value)
	countreads(value,19,5)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s19.QueryP(294,&value)
	log("p19 r:%d", value)
	countreads(value,19,19)
	l.Unlock()
}


func p20() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	s6.QueryP(90,&value)
	log("p20 r:%d", value)
	countreads(value,20,6)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s25.QueryP(398,&value)
	log("p20 r:%d", value)
	countreads(value,20,25)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s32.QueryP(505,&value)
	log("p20 r:%d", value)
	countreads(value,20,32)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s22.QueryP(348,&value)
	log("p20 r:%d", value)
	countreads(value,20,22)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s32.QueryP(506,&value)
	log("p20 r:%d", value)
	countreads(value,20,32)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s25.QueryP(386,&value)
	log("p20 r:%d", value)
	countreads(value,20,25)
	l.Unlock()

	delay()
	l.Lock()
	value = 25
	s25.Put(396,value)
	log("p20 w:%d", value)
	countwrites(value,20,25)
	l.Unlock()

	delay()
	l.Lock()
	value = 6
	s6.Put(92,value)
	log("p20 w:%d", value)
	countwrites(value,20,6)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s23.QueryP(358,&value)
	log("p20 r:%d", value)
	countreads(value,20,23)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s22.QueryP(343,&value)
	log("p20 r:%d", value)
	countreads(value,20,22)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s5.QueryP(69,&value)
	log("p20 r:%d", value)
	countreads(value,20,5)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s29.QueryP(452,&value)
	log("p20 r:%d", value)
	countreads(value,20,29)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s15.QueryP(227,&value)
	log("p20 r:%d", value)
	countreads(value,20,15)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s27.QueryP(424,&value)
	log("p20 r:%d", value)
	countreads(value,20,27)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s23.QueryP(355,&value)
	log("p20 r:%d", value)
	countreads(value,20,23)
	l.Unlock()

	delay()
	l.Lock()
	value = 6
	s6.Put(84,value)
	log("p20 w:%d", value)
	countwrites(value,20,6)
	l.Unlock()
}


func p21() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	s28.QueryP(448,&value)
	log("p21 r:%d", value)
	countreads(value,21,28)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s23.QueryP(363,&value)
	log("p21 r:%d", value)
	countreads(value,21,23)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s11.QueryP(164,&value)
	log("p21 r:%d", value)
	countreads(value,21,11)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s32.QueryP(503,&value)
	log("p21 r:%d", value)
	countreads(value,21,32)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s14.QueryP(211,&value)
	log("p21 r:%d", value)
	countreads(value,21,14)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s16.QueryP(241,&value)
	log("p21 r:%d", value)
	countreads(value,21,16)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s28.QueryP(444,&value)
	log("p21 r:%d", value)
	countreads(value,21,28)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s11.QueryP(171,&value)
	log("p21 r:%d", value)
	countreads(value,21,11)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s21.QueryP(327,&value)
	log("p21 r:%d", value)
	countreads(value,21,21)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s9.QueryP(137,&value)
	log("p21 r:%d", value)
	countreads(value,21,9)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s24.QueryP(373,&value)
	log("p21 r:%d", value)
	countreads(value,21,24)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s31.QueryP(496,&value)
	log("p21 r:%d", value)
	countreads(value,21,31)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s26.QueryP(411,&value)
	log("p21 r:%d", value)
	countreads(value,21,26)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s6.QueryP(96,&value)
	log("p21 r:%d", value)
	countreads(value,21,6)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s16.QueryP(252,&value)
	log("p21 r:%d", value)
	countreads(value,21,16)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s22.QueryP(345,&value)
	log("p21 r:%d", value)
	countreads(value,21,22)
	l.Unlock()
}


func p22() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	s28.QueryP(438,&value)
	log("p22 r:%d", value)
	countreads(value,22,28)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s30.QueryP(475,&value)
	log("p22 r:%d", value)
	countreads(value,22,30)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s7.QueryP(109,&value)
	log("p22 r:%d", value)
	countreads(value,22,7)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s5.QueryP(71,&value)
	log("p22 r:%d", value)
	countreads(value,22,5)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s5.QueryP(73,&value)
	log("p22 r:%d", value)
	countreads(value,22,5)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s2.QueryP(27,&value)
	log("p22 r:%d", value)
	countreads(value,22,2)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s15.QueryP(232,&value)
	log("p22 r:%d", value)
	countreads(value,22,15)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s11.QueryP(161,&value)
	log("p22 r:%d", value)
	countreads(value,22,11)
	l.Unlock()

	delay()
	l.Lock()
	value = 30
	s30.Put(479,value)
	log("p22 w:%d", value)
	countwrites(value,22,30)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s32.QueryP(504,&value)
	log("p22 r:%d", value)
	countreads(value,22,32)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s15.QueryP(233,&value)
	log("p22 r:%d", value)
	countreads(value,22,15)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s22.QueryP(349,&value)
	log("p22 r:%d", value)
	countreads(value,22,22)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s17.QueryP(263,&value)
	log("p22 r:%d", value)
	countreads(value,22,17)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s4.QueryP(56,&value)
	log("p22 r:%d", value)
	countreads(value,22,4)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s24.QueryP(384,&value)
	log("p22 r:%d", value)
	countreads(value,22,24)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s31.QueryP(496,&value)
	log("p22 r:%d", value)
	countreads(value,22,31)
	l.Unlock()
}


func p23() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	s21.QueryP(335,&value)
	log("p23 r:%d", value)
	countreads(value,23,21)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s19.QueryP(296,&value)
	log("p23 r:%d", value)
	countreads(value,23,19)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s24.QueryP(383,&value)
	log("p23 r:%d", value)
	countreads(value,23,24)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s29.QueryP(457,&value)
	log("p23 r:%d", value)
	countreads(value,23,29)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s28.QueryP(448,&value)
	log("p23 r:%d", value)
	countreads(value,23,28)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s26.QueryP(403,&value)
	log("p23 r:%d", value)
	countreads(value,23,26)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s9.QueryP(143,&value)
	log("p23 r:%d", value)
	countreads(value,23,9)
	l.Unlock()

	delay()
	l.Lock()
	value = 23
	s23.Put(367,value)
	log("p23 w:%d", value)
	countwrites(value,23,23)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s31.QueryP(482,&value)
	log("p23 r:%d", value)
	countreads(value,23,31)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s3.QueryP(36,&value)
	log("p23 r:%d", value)
	countreads(value,23,3)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s21.QueryP(327,&value)
	log("p23 r:%d", value)
	countreads(value,23,21)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s17.QueryP(271,&value)
	log("p23 r:%d", value)
	countreads(value,23,17)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s23.QueryP(357,&value)
	log("p23 r:%d", value)
	countreads(value,23,23)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s26.QueryP(403,&value)
	log("p23 r:%d", value)
	countreads(value,23,26)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s20.QueryP(316,&value)
	log("p23 r:%d", value)
	countreads(value,23,20)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s11.QueryP(172,&value)
	log("p23 r:%d", value)
	countreads(value,23,11)
	l.Unlock()
}


func p24() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	s10.QueryP(151,&value)
	log("p24 r:%d", value)
	countreads(value,24,10)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s25.QueryP(393,&value)
	log("p24 r:%d", value)
	countreads(value,24,25)
	l.Unlock()

	delay()
	l.Lock()
	value = 3
	s3.Put(48,value)
	log("p24 w:%d", value)
	countwrites(value,24,3)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s22.QueryP(343,&value)
	log("p24 r:%d", value)
	countreads(value,24,22)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s2.QueryP(18,&value)
	log("p24 r:%d", value)
	countreads(value,24,2)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s8.QueryP(124,&value)
	log("p24 r:%d", value)
	countreads(value,24,8)
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
	s3.QueryP(40,&value)
	log("p24 r:%d", value)
	countreads(value,24,3)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s1.QueryP(10,&value)
	log("p24 r:%d", value)
	countreads(value,24,1)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s25.QueryP(387,&value)
	log("p24 r:%d", value)
	countreads(value,24,25)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s23.QueryP(355,&value)
	log("p24 r:%d", value)
	countreads(value,24,23)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s23.QueryP(354,&value)
	log("p24 r:%d", value)
	countreads(value,24,23)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s27.QueryP(429,&value)
	log("p24 r:%d", value)
	countreads(value,24,27)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s12.QueryP(181,&value)
	log("p24 r:%d", value)
	countreads(value,24,12)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s1.QueryP(2,&value)
	log("p24 r:%d", value)
	countreads(value,24,1)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s32.QueryP(505,&value)
	log("p24 r:%d", value)
	countreads(value,24,32)
	l.Unlock()
}


func p25() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	s25.QueryP(393,&value)
	log("p25 r:%d", value)
	countreads(value,25,25)
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
	value = -1
	s19.QueryP(294,&value)
	log("p25 r:%d", value)
	countreads(value,25,19)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s9.QueryP(132,&value)
	log("p25 r:%d", value)
	countreads(value,25,9)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s27.QueryP(432,&value)
	log("p25 r:%d", value)
	countreads(value,25,27)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s2.QueryP(20,&value)
	log("p25 r:%d", value)
	countreads(value,25,2)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s13.QueryP(193,&value)
	log("p25 r:%d", value)
	countreads(value,25,13)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s9.QueryP(141,&value)
	log("p25 r:%d", value)
	countreads(value,25,9)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s24.QueryP(382,&value)
	log("p25 r:%d", value)
	countreads(value,25,24)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s9.QueryP(137,&value)
	log("p25 r:%d", value)
	countreads(value,25,9)
	l.Unlock()

	delay()
	l.Lock()
	value = 11
	s11.Put(174,value)
	log("p25 w:%d", value)
	countwrites(value,25,11)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s21.QueryP(335,&value)
	log("p25 r:%d", value)
	countreads(value,25,21)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s18.QueryP(282,&value)
	log("p25 r:%d", value)
	countreads(value,25,18)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s16.QueryP(255,&value)
	log("p25 r:%d", value)
	countreads(value,25,16)
	l.Unlock()

	delay()
	l.Lock()
	value = 10
	s10.Put(152,value)
	log("p25 w:%d", value)
	countwrites(value,25,10)
	l.Unlock()

	delay()
	l.Lock()
	value = 2
	s2.Put(17,value)
	log("p25 w:%d", value)
	countwrites(value,25,2)
	l.Unlock()
}


func p26() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	s31.QueryP(486,&value)
	log("p26 r:%d", value)
	countreads(value,26,31)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s31.QueryP(486,&value)
	log("p26 r:%d", value)
	countreads(value,26,31)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s24.QueryP(383,&value)
	log("p26 r:%d", value)
	countreads(value,26,24)
	l.Unlock()

	delay()
	l.Lock()
	value = 23
	s23.Put(361,value)
	log("p26 w:%d", value)
	countwrites(value,26,23)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s26.QueryP(412,&value)
	log("p26 r:%d", value)
	countreads(value,26,26)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s7.QueryP(100,&value)
	log("p26 r:%d", value)
	countreads(value,26,7)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s10.QueryP(148,&value)
	log("p26 r:%d", value)
	countreads(value,26,10)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s19.QueryP(303,&value)
	log("p26 r:%d", value)
	countreads(value,26,19)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s15.QueryP(229,&value)
	log("p26 r:%d", value)
	countreads(value,26,15)
	l.Unlock()

	delay()
	l.Lock()
	value = 2
	s2.Put(25,value)
	log("p26 w:%d", value)
	countwrites(value,26,2)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s17.QueryP(267,&value)
	log("p26 r:%d", value)
	countreads(value,26,17)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s20.QueryP(318,&value)
	log("p26 r:%d", value)
	countreads(value,26,20)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s22.QueryP(338,&value)
	log("p26 r:%d", value)
	countreads(value,26,22)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s9.QueryP(131,&value)
	log("p26 r:%d", value)
	countreads(value,26,9)
	l.Unlock()

	delay()
	l.Lock()
	value = 12
	s12.Put(181,value)
	log("p26 w:%d", value)
	countwrites(value,26,12)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s31.QueryP(493,&value)
	log("p26 r:%d", value)
	countreads(value,26,31)
	l.Unlock()
}


func p27() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	s30.QueryP(465,&value)
	log("p27 r:%d", value)
	countreads(value,27,30)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s30.QueryP(467,&value)
	log("p27 r:%d", value)
	countreads(value,27,30)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s28.QueryP(442,&value)
	log("p27 r:%d", value)
	countreads(value,27,28)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s22.QueryP(347,&value)
	log("p27 r:%d", value)
	countreads(value,27,22)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s18.QueryP(287,&value)
	log("p27 r:%d", value)
	countreads(value,27,18)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s22.QueryP(341,&value)
	log("p27 r:%d", value)
	countreads(value,27,22)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s17.QueryP(258,&value)
	log("p27 r:%d", value)
	countreads(value,27,17)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s22.QueryP(345,&value)
	log("p27 r:%d", value)
	countreads(value,27,22)
	l.Unlock()

	delay()
	l.Lock()
	value = 32
	s32.Put(501,value)
	log("p27 w:%d", value)
	countwrites(value,27,32)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s1.QueryP(7,&value)
	log("p27 r:%d", value)
	countreads(value,27,1)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s17.QueryP(260,&value)
	log("p27 r:%d", value)
	countreads(value,27,17)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s13.QueryP(200,&value)
	log("p27 r:%d", value)
	countreads(value,27,13)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s25.QueryP(396,&value)
	log("p27 r:%d", value)
	countreads(value,27,25)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s31.QueryP(494,&value)
	log("p27 r:%d", value)
	countreads(value,27,31)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s13.QueryP(201,&value)
	log("p27 r:%d", value)
	countreads(value,27,13)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s5.QueryP(74,&value)
	log("p27 r:%d", value)
	countreads(value,27,5)
	l.Unlock()
}


func p28() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	s11.QueryP(172,&value)
	log("p28 r:%d", value)
	countreads(value,28,11)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s8.QueryP(126,&value)
	log("p28 r:%d", value)
	countreads(value,28,8)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s20.QueryP(305,&value)
	log("p28 r:%d", value)
	countreads(value,28,20)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s12.QueryP(183,&value)
	log("p28 r:%d", value)
	countreads(value,28,12)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s13.QueryP(198,&value)
	log("p28 r:%d", value)
	countreads(value,28,13)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s5.QueryP(75,&value)
	log("p28 r:%d", value)
	countreads(value,28,5)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s31.QueryP(485,&value)
	log("p28 r:%d", value)
	countreads(value,28,31)
	l.Unlock()

	delay()
	l.Lock()
	value = 2
	s2.Put(19,value)
	log("p28 w:%d", value)
	countwrites(value,28,2)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s8.QueryP(119,&value)
	log("p28 r:%d", value)
	countreads(value,28,8)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s3.QueryP(36,&value)
	log("p28 r:%d", value)
	countreads(value,28,3)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s5.QueryP(80,&value)
	log("p28 r:%d", value)
	countreads(value,28,5)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s13.QueryP(203,&value)
	log("p28 r:%d", value)
	countreads(value,28,13)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s8.QueryP(122,&value)
	log("p28 r:%d", value)
	countreads(value,28,8)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s4.QueryP(59,&value)
	log("p28 r:%d", value)
	countreads(value,28,4)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s26.QueryP(405,&value)
	log("p28 r:%d", value)
	countreads(value,28,26)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s29.QueryP(450,&value)
	log("p28 r:%d", value)
	countreads(value,28,29)
	l.Unlock()
}


func p29() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	s6.QueryP(86,&value)
	log("p29 r:%d", value)
	countreads(value,29,6)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s21.QueryP(336,&value)
	log("p29 r:%d", value)
	countreads(value,29,21)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s20.QueryP(316,&value)
	log("p29 r:%d", value)
	countreads(value,29,20)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s22.QueryP(346,&value)
	log("p29 r:%d", value)
	countreads(value,29,22)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s31.QueryP(493,&value)
	log("p29 r:%d", value)
	countreads(value,29,31)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s25.QueryP(397,&value)
	log("p29 r:%d", value)
	countreads(value,29,25)
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
	s5.QueryP(79,&value)
	log("p29 r:%d", value)
	countreads(value,29,5)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s2.QueryP(31,&value)
	log("p29 r:%d", value)
	countreads(value,29,2)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s18.QueryP(277,&value)
	log("p29 r:%d", value)
	countreads(value,29,18)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s24.QueryP(384,&value)
	log("p29 r:%d", value)
	countreads(value,29,24)
	l.Unlock()

	delay()
	l.Lock()
	value = 26
	s26.Put(403,value)
	log("p29 w:%d", value)
	countwrites(value,29,26)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s19.QueryP(293,&value)
	log("p29 r:%d", value)
	countreads(value,29,19)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s17.QueryP(262,&value)
	log("p29 r:%d", value)
	countreads(value,29,17)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s3.QueryP(46,&value)
	log("p29 r:%d", value)
	countreads(value,29,3)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s31.QueryP(494,&value)
	log("p29 r:%d", value)
	countreads(value,29,31)
	l.Unlock()
}


func p30() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	s2.QueryP(22,&value)
	log("p30 r:%d", value)
	countreads(value,30,2)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s5.QueryP(71,&value)
	log("p30 r:%d", value)
	countreads(value,30,5)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s7.QueryP(107,&value)
	log("p30 r:%d", value)
	countreads(value,30,7)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s13.QueryP(207,&value)
	log("p30 r:%d", value)
	countreads(value,30,13)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s27.QueryP(429,&value)
	log("p30 r:%d", value)
	countreads(value,30,27)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s13.QueryP(205,&value)
	log("p30 r:%d", value)
	countreads(value,30,13)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s21.QueryP(336,&value)
	log("p30 r:%d", value)
	countreads(value,30,21)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s6.QueryP(87,&value)
	log("p30 r:%d", value)
	countreads(value,30,6)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s31.QueryP(486,&value)
	log("p30 r:%d", value)
	countreads(value,30,31)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s3.QueryP(42,&value)
	log("p30 r:%d", value)
	countreads(value,30,3)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s7.QueryP(97,&value)
	log("p30 r:%d", value)
	countreads(value,30,7)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s15.QueryP(229,&value)
	log("p30 r:%d", value)
	countreads(value,30,15)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s30.QueryP(472,&value)
	log("p30 r:%d", value)
	countreads(value,30,30)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s30.QueryP(474,&value)
	log("p30 r:%d", value)
	countreads(value,30,30)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s3.QueryP(41,&value)
	log("p30 r:%d", value)
	countreads(value,30,3)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s2.QueryP(31,&value)
	log("p30 r:%d", value)
	countreads(value,30,2)
	l.Unlock()
}


func p31() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 23
	s23.Put(360,value)
	log("p31 w:%d", value)
	countwrites(value,31,23)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s15.QueryP(232,&value)
	log("p31 r:%d", value)
	countreads(value,31,15)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s5.QueryP(73,&value)
	log("p31 r:%d", value)
	countreads(value,31,5)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s21.QueryP(323,&value)
	log("p31 r:%d", value)
	countreads(value,31,21)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s14.QueryP(211,&value)
	log("p31 r:%d", value)
	countreads(value,31,14)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s10.QueryP(148,&value)
	log("p31 r:%d", value)
	countreads(value,31,10)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s30.QueryP(476,&value)
	log("p31 r:%d", value)
	countreads(value,31,30)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s13.QueryP(198,&value)
	log("p31 r:%d", value)
	countreads(value,31,13)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s3.QueryP(43,&value)
	log("p31 r:%d", value)
	countreads(value,31,3)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s27.QueryP(418,&value)
	log("p31 r:%d", value)
	countreads(value,31,27)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s15.QueryP(225,&value)
	log("p31 r:%d", value)
	countreads(value,31,15)
	l.Unlock()

	delay()
	l.Lock()
	value = 19
	s19.Put(297,value)
	log("p31 w:%d", value)
	countwrites(value,31,19)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s16.QueryP(246,&value)
	log("p31 r:%d", value)
	countreads(value,31,16)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s6.QueryP(87,&value)
	log("p31 r:%d", value)
	countreads(value,31,6)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s18.QueryP(273,&value)
	log("p31 r:%d", value)
	countreads(value,31,18)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s22.QueryP(351,&value)
	log("p31 r:%d", value)
	countreads(value,31,22)
	l.Unlock()
}


func p32() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	s23.QueryP(363,&value)
	log("p32 r:%d", value)
	countreads(value,32,23)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s6.QueryP(90,&value)
	log("p32 r:%d", value)
	countreads(value,32,6)
	l.Unlock()

	delay()
	l.Lock()
	value = 12
	s12.Put(181,value)
	log("p32 w:%d", value)
	countwrites(value,32,12)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s31.QueryP(484,&value)
	log("p32 r:%d", value)
	countreads(value,32,31)
	l.Unlock()

	delay()
	l.Lock()
	value = 22
	s22.Put(348,value)
	log("p32 w:%d", value)
	countwrites(value,32,22)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s24.QueryP(383,&value)
	log("p32 r:%d", value)
	countreads(value,32,24)
	l.Unlock()

	delay()
	l.Lock()
	value = 14
	s14.Put(215,value)
	log("p32 w:%d", value)
	countwrites(value,32,14)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s2.QueryP(21,&value)
	log("p32 r:%d", value)
	countreads(value,32,2)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s32.QueryP(506,&value)
	log("p32 r:%d", value)
	countreads(value,32,32)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s27.QueryP(424,&value)
	log("p32 r:%d", value)
	countreads(value,32,27)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s18.QueryP(286,&value)
	log("p32 r:%d", value)
	countreads(value,32,18)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s1.QueryP(16,&value)
	log("p32 r:%d", value)
	countreads(value,32,1)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s24.QueryP(384,&value)
	log("p32 r:%d", value)
	countreads(value,32,24)
	l.Unlock()

	delay()
	l.Lock()
	value = 19
	s19.Put(291,value)
	log("p32 w:%d", value)
	countwrites(value,32,19)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s25.QueryP(398,&value)
	log("p32 r:%d", value)
	countreads(value,32,25)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s15.QueryP(229,&value)
	log("p32 r:%d", value)
	countreads(value,32,15)
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
