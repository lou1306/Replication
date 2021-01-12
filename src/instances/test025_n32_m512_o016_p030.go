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
	value = 29
	s29.Put(461,value)
	log("p1 w:%d", value)
	countwrites(value,1,29)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s11.QueryP(167,&value)
	log("p1 r:%d", value)
	countreads(value,1,11)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s20.QueryP(314,&value)
	log("p1 r:%d", value)
	countreads(value,1,20)
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
	s23.QueryP(357,&value)
	log("p1 r:%d", value)
	countreads(value,1,23)
	l.Unlock()

	delay()
	l.Lock()
	value = 4
	s4.Put(52,value)
	log("p1 w:%d", value)
	countwrites(value,1,4)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s27.QueryP(422,&value)
	log("p1 r:%d", value)
	countreads(value,1,27)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s4.QueryP(63,&value)
	log("p1 r:%d", value)
	countreads(value,1,4)
	l.Unlock()

	delay()
	l.Lock()
	value = 14
	s14.Put(223,value)
	log("p1 w:%d", value)
	countwrites(value,1,14)
	l.Unlock()

	delay()
	l.Lock()
	value = 4
	s4.Put(53,value)
	log("p1 w:%d", value)
	countwrites(value,1,4)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s32.QueryP(508,&value)
	log("p1 r:%d", value)
	countreads(value,1,32)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s8.QueryP(114,&value)
	log("p1 r:%d", value)
	countreads(value,1,8)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s21.QueryP(322,&value)
	log("p1 r:%d", value)
	countreads(value,1,21)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s16.QueryP(247,&value)
	log("p1 r:%d", value)
	countreads(value,1,16)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s14.QueryP(221,&value)
	log("p1 r:%d", value)
	countreads(value,1,14)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s25.QueryP(396,&value)
	log("p1 r:%d", value)
	countreads(value,1,25)
	l.Unlock()
}

func p2() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	s10.QueryP(157,&value)
	log("p2 r:%d", value)
	countreads(value,2,10)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s14.QueryP(222,&value)
	log("p2 r:%d", value)
	countreads(value,2,14)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s17.QueryP(262,&value)
	log("p2 r:%d", value)
	countreads(value,2,17)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s14.QueryP(221,&value)
	log("p2 r:%d", value)
	countreads(value,2,14)
	l.Unlock()

	delay()
	l.Lock()
	value = 31
	s31.Put(496,value)
	log("p2 w:%d", value)
	countwrites(value,2,31)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s17.QueryP(259,&value)
	log("p2 r:%d", value)
	countreads(value,2,17)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s6.QueryP(90,&value)
	log("p2 r:%d", value)
	countreads(value,2,6)
	l.Unlock()

	delay()
	l.Lock()
	value = 23
	s23.Put(364,value)
	log("p2 w:%d", value)
	countwrites(value,2,23)
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
	value = -1
	s30.QueryP(475,&value)
	log("p2 r:%d", value)
	countreads(value,2,30)
	l.Unlock()

	delay()
	l.Lock()
	value = 27
	s27.Put(426,value)
	log("p2 w:%d", value)
	countwrites(value,2,27)
	l.Unlock()

	delay()
	l.Lock()
	value = 23
	s23.Put(366,value)
	log("p2 w:%d", value)
	countwrites(value,2,23)
	l.Unlock()

	delay()
	l.Lock()
	value = 8
	s8.Put(125,value)
	log("p2 w:%d", value)
	countwrites(value,2,8)
	l.Unlock()

	delay()
	l.Lock()
	value = 11
	s11.Put(164,value)
	log("p2 w:%d", value)
	countwrites(value,2,11)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s3.QueryP(46,&value)
	log("p2 r:%d", value)
	countreads(value,2,3)
	l.Unlock()

	delay()
	l.Lock()
	value = 9
	s9.Put(142,value)
	log("p2 w:%d", value)
	countwrites(value,2,9)
	l.Unlock()
}

func p3() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	s26.QueryP(402,&value)
	log("p3 r:%d", value)
	countreads(value,3,26)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s7.QueryP(110,&value)
	log("p3 r:%d", value)
	countreads(value,3,7)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s17.QueryP(263,&value)
	log("p3 r:%d", value)
	countreads(value,3,17)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s18.QueryP(282,&value)
	log("p3 r:%d", value)
	countreads(value,3,18)
	l.Unlock()

	delay()
	l.Lock()
	value = 10
	s10.Put(154,value)
	log("p3 w:%d", value)
	countwrites(value,3,10)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s29.QueryP(450,&value)
	log("p3 r:%d", value)
	countreads(value,3,29)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s27.QueryP(424,&value)
	log("p3 r:%d", value)
	countreads(value,3,27)
	l.Unlock()

	delay()
	l.Lock()
	value = 12
	s12.Put(180,value)
	log("p3 w:%d", value)
	countwrites(value,3,12)
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
	s19.QueryP(297,&value)
	log("p3 r:%d", value)
	countreads(value,3,19)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s26.QueryP(410,&value)
	log("p3 r:%d", value)
	countreads(value,3,26)
	l.Unlock()

	delay()
	l.Lock()
	value = 6
	s6.Put(96,value)
	log("p3 w:%d", value)
	countwrites(value,3,6)
	l.Unlock()

	delay()
	l.Lock()
	value = 22
	s22.Put(341,value)
	log("p3 w:%d", value)
	countwrites(value,3,22)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s16.QueryP(243,&value)
	log("p3 r:%d", value)
	countreads(value,3,16)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s6.QueryP(88,&value)
	log("p3 r:%d", value)
	countreads(value,3,6)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s17.QueryP(257,&value)
	log("p3 r:%d", value)
	countreads(value,3,17)
	l.Unlock()
}

func p4() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	s28.QueryP(445,&value)
	log("p4 r:%d", value)
	countreads(value,4,28)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s31.QueryP(487,&value)
	log("p4 r:%d", value)
	countreads(value,4,31)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s18.QueryP(285,&value)
	log("p4 r:%d", value)
	countreads(value,4,18)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s28.QueryP(433,&value)
	log("p4 r:%d", value)
	countreads(value,4,28)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s23.QueryP(366,&value)
	log("p4 r:%d", value)
	countreads(value,4,23)
	l.Unlock()

	delay()
	l.Lock()
	value = 9
	s9.Put(136,value)
	log("p4 w:%d", value)
	countwrites(value,4,9)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s3.QueryP(38,&value)
	log("p4 r:%d", value)
	countreads(value,4,3)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s11.QueryP(168,&value)
	log("p4 r:%d", value)
	countreads(value,4,11)
	l.Unlock()

	delay()
	l.Lock()
	value = 10
	s10.Put(145,value)
	log("p4 w:%d", value)
	countwrites(value,4,10)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s17.QueryP(270,&value)
	log("p4 r:%d", value)
	countreads(value,4,17)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s8.QueryP(126,&value)
	log("p4 r:%d", value)
	countreads(value,4,8)
	l.Unlock()

	delay()
	l.Lock()
	value = 8
	s8.Put(122,value)
	log("p4 w:%d", value)
	countwrites(value,4,8)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s11.QueryP(175,&value)
	log("p4 r:%d", value)
	countreads(value,4,11)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s3.QueryP(37,&value)
	log("p4 r:%d", value)
	countreads(value,4,3)
	l.Unlock()

	delay()
	l.Lock()
	value = 6
	s6.Put(81,value)
	log("p4 w:%d", value)
	countwrites(value,4,6)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s5.QueryP(78,&value)
	log("p4 r:%d", value)
	countreads(value,4,5)
	l.Unlock()
}

func p5() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	s13.QueryP(204,&value)
	log("p5 r:%d", value)
	countreads(value,5,13)
	l.Unlock()

	delay()
	l.Lock()
	value = 10
	s10.Put(149,value)
	log("p5 w:%d", value)
	countwrites(value,5,10)
	l.Unlock()

	delay()
	l.Lock()
	value = 20
	s20.Put(319,value)
	log("p5 w:%d", value)
	countwrites(value,5,20)
	l.Unlock()

	delay()
	l.Lock()
	value = 7
	s7.Put(110,value)
	log("p5 w:%d", value)
	countwrites(value,5,7)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s4.QueryP(52,&value)
	log("p5 r:%d", value)
	countreads(value,5,4)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s18.QueryP(287,&value)
	log("p5 r:%d", value)
	countreads(value,5,18)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s3.QueryP(36,&value)
	log("p5 r:%d", value)
	countreads(value,5,3)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s24.QueryP(372,&value)
	log("p5 r:%d", value)
	countreads(value,5,24)
	l.Unlock()

	delay()
	l.Lock()
	value = 17
	s17.Put(261,value)
	log("p5 w:%d", value)
	countwrites(value,5,17)
	l.Unlock()

	delay()
	l.Lock()
	value = 15
	s15.Put(233,value)
	log("p5 w:%d", value)
	countwrites(value,5,15)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s27.QueryP(430,&value)
	log("p5 r:%d", value)
	countreads(value,5,27)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s18.QueryP(276,&value)
	log("p5 r:%d", value)
	countreads(value,5,18)
	l.Unlock()

	delay()
	l.Lock()
	value = 15
	s15.Put(229,value)
	log("p5 w:%d", value)
	countwrites(value,5,15)
	l.Unlock()

	delay()
	l.Lock()
	value = 29
	s29.Put(449,value)
	log("p5 w:%d", value)
	countwrites(value,5,29)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s12.QueryP(185,&value)
	log("p5 r:%d", value)
	countreads(value,5,12)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s1.QueryP(12,&value)
	log("p5 r:%d", value)
	countreads(value,5,1)
	l.Unlock()
}

func p6() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 19
	s19.Put(292,value)
	log("p6 w:%d", value)
	countwrites(value,6,19)
	l.Unlock()

	delay()
	l.Lock()
	value = 32
	s32.Put(505,value)
	log("p6 w:%d", value)
	countwrites(value,6,32)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s32.QueryP(497,&value)
	log("p6 r:%d", value)
	countreads(value,6,32)
	l.Unlock()

	delay()
	l.Lock()
	value = 9
	s9.Put(139,value)
	log("p6 w:%d", value)
	countwrites(value,6,9)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s5.QueryP(69,&value)
	log("p6 r:%d", value)
	countreads(value,6,5)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s1.QueryP(3,&value)
	log("p6 r:%d", value)
	countreads(value,6,1)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s22.QueryP(345,&value)
	log("p6 r:%d", value)
	countreads(value,6,22)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s7.QueryP(111,&value)
	log("p6 r:%d", value)
	countreads(value,6,7)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s4.QueryP(57,&value)
	log("p6 r:%d", value)
	countreads(value,6,4)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s19.QueryP(299,&value)
	log("p6 r:%d", value)
	countreads(value,6,19)
	l.Unlock()

	delay()
	l.Lock()
	value = 21
	s21.Put(325,value)
	log("p6 w:%d", value)
	countwrites(value,6,21)
	l.Unlock()

	delay()
	l.Lock()
	value = 10
	s10.Put(149,value)
	log("p6 w:%d", value)
	countwrites(value,6,10)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s14.QueryP(214,&value)
	log("p6 r:%d", value)
	countreads(value,6,14)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s18.QueryP(283,&value)
	log("p6 r:%d", value)
	countreads(value,6,18)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s24.QueryP(383,&value)
	log("p6 r:%d", value)
	countreads(value,6,24)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s3.QueryP(40,&value)
	log("p6 r:%d", value)
	countreads(value,6,3)
	l.Unlock()
}

func p7() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	s3.QueryP(35,&value)
	log("p7 r:%d", value)
	countreads(value,7,3)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s12.QueryP(188,&value)
	log("p7 r:%d", value)
	countreads(value,7,12)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s17.QueryP(263,&value)
	log("p7 r:%d", value)
	countreads(value,7,17)
	l.Unlock()

	delay()
	l.Lock()
	value = 7
	s7.Put(108,value)
	log("p7 w:%d", value)
	countwrites(value,7,7)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s3.QueryP(41,&value)
	log("p7 r:%d", value)
	countreads(value,7,3)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s1.QueryP(7,&value)
	log("p7 r:%d", value)
	countreads(value,7,1)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s8.QueryP(128,&value)
	log("p7 r:%d", value)
	countreads(value,7,8)
	l.Unlock()

	delay()
	l.Lock()
	value = 16
	s16.Put(251,value)
	log("p7 w:%d", value)
	countwrites(value,7,16)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s1.QueryP(16,&value)
	log("p7 r:%d", value)
	countreads(value,7,1)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s2.QueryP(32,&value)
	log("p7 r:%d", value)
	countreads(value,7,2)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s26.QueryP(403,&value)
	log("p7 r:%d", value)
	countreads(value,7,26)
	l.Unlock()

	delay()
	l.Lock()
	value = 29
	s29.Put(459,value)
	log("p7 w:%d", value)
	countwrites(value,7,29)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s21.QueryP(331,&value)
	log("p7 r:%d", value)
	countreads(value,7,21)
	l.Unlock()

	delay()
	l.Lock()
	value = 30
	s30.Put(467,value)
	log("p7 w:%d", value)
	countwrites(value,7,30)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s19.QueryP(295,&value)
	log("p7 r:%d", value)
	countreads(value,7,19)
	l.Unlock()

	delay()
	l.Lock()
	value = 24
	s24.Put(375,value)
	log("p7 w:%d", value)
	countwrites(value,7,24)
	l.Unlock()
}

func p8() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 23
	s23.Put(359,value)
	log("p8 w:%d", value)
	countwrites(value,8,23)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s4.QueryP(61,&value)
	log("p8 r:%d", value)
	countreads(value,8,4)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s7.QueryP(100,&value)
	log("p8 r:%d", value)
	countreads(value,8,7)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s12.QueryP(188,&value)
	log("p8 r:%d", value)
	countreads(value,8,12)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s29.QueryP(462,&value)
	log("p8 r:%d", value)
	countreads(value,8,29)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s16.QueryP(245,&value)
	log("p8 r:%d", value)
	countreads(value,8,16)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s23.QueryP(353,&value)
	log("p8 r:%d", value)
	countreads(value,8,23)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s5.QueryP(65,&value)
	log("p8 r:%d", value)
	countreads(value,8,5)
	l.Unlock()

	delay()
	l.Lock()
	value = 15
	s15.Put(238,value)
	log("p8 w:%d", value)
	countwrites(value,8,15)
	l.Unlock()

	delay()
	l.Lock()
	value = 17
	s17.Put(268,value)
	log("p8 w:%d", value)
	countwrites(value,8,17)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s21.QueryP(332,&value)
	log("p8 r:%d", value)
	countreads(value,8,21)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s18.QueryP(274,&value)
	log("p8 r:%d", value)
	countreads(value,8,18)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s32.QueryP(509,&value)
	log("p8 r:%d", value)
	countreads(value,8,32)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s29.QueryP(450,&value)
	log("p8 r:%d", value)
	countreads(value,8,29)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s6.QueryP(81,&value)
	log("p8 r:%d", value)
	countreads(value,8,6)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s2.QueryP(23,&value)
	log("p8 r:%d", value)
	countreads(value,8,2)
	l.Unlock()
}

func p9() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	s21.QueryP(334,&value)
	log("p9 r:%d", value)
	countreads(value,9,21)
	l.Unlock()

	delay()
	l.Lock()
	value = 30
	s30.Put(475,value)
	log("p9 w:%d", value)
	countwrites(value,9,30)
	l.Unlock()

	delay()
	l.Lock()
	value = 4
	s4.Put(63,value)
	log("p9 w:%d", value)
	countwrites(value,9,4)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s8.QueryP(115,&value)
	log("p9 r:%d", value)
	countreads(value,9,8)
	l.Unlock()

	delay()
	l.Lock()
	value = 24
	s24.Put(375,value)
	log("p9 w:%d", value)
	countwrites(value,9,24)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s5.QueryP(73,&value)
	log("p9 r:%d", value)
	countreads(value,9,5)
	l.Unlock()

	delay()
	l.Lock()
	value = 8
	s8.Put(124,value)
	log("p9 w:%d", value)
	countwrites(value,9,8)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s9.QueryP(141,&value)
	log("p9 r:%d", value)
	countreads(value,9,9)
	l.Unlock()

	delay()
	l.Lock()
	value = 4
	s4.Put(51,value)
	log("p9 w:%d", value)
	countwrites(value,9,4)
	l.Unlock()

	delay()
	l.Lock()
	value = 16
	s16.Put(255,value)
	log("p9 w:%d", value)
	countwrites(value,9,16)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s2.QueryP(19,&value)
	log("p9 r:%d", value)
	countreads(value,9,2)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s1.QueryP(12,&value)
	log("p9 r:%d", value)
	countreads(value,9,1)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s8.QueryP(118,&value)
	log("p9 r:%d", value)
	countreads(value,9,8)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s12.QueryP(183,&value)
	log("p9 r:%d", value)
	countreads(value,9,12)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s29.QueryP(458,&value)
	log("p9 r:%d", value)
	countreads(value,9,29)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s15.QueryP(229,&value)
	log("p9 r:%d", value)
	countreads(value,9,15)
	l.Unlock()
}


func p10() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 5
	s5.Put(68,value)
	log("p10 w:%d", value)
	countwrites(value,10,5)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s31.QueryP(491,&value)
	log("p10 r:%d", value)
	countreads(value,10,31)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s17.QueryP(265,&value)
	log("p10 r:%d", value)
	countreads(value,10,17)
	l.Unlock()

	delay()
	l.Lock()
	value = 11
	s11.Put(162,value)
	log("p10 w:%d", value)
	countwrites(value,10,11)
	l.Unlock()

	delay()
	l.Lock()
	value = 31
	s31.Put(484,value)
	log("p10 w:%d", value)
	countwrites(value,10,31)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s23.QueryP(357,&value)
	log("p10 r:%d", value)
	countreads(value,10,23)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s25.QueryP(393,&value)
	log("p10 r:%d", value)
	countreads(value,10,25)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s21.QueryP(325,&value)
	log("p10 r:%d", value)
	countreads(value,10,21)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s8.QueryP(126,&value)
	log("p10 r:%d", value)
	countreads(value,10,8)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s1.QueryP(5,&value)
	log("p10 r:%d", value)
	countreads(value,10,1)
	l.Unlock()

	delay()
	l.Lock()
	value = 27
	s27.Put(420,value)
	log("p10 w:%d", value)
	countwrites(value,10,27)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s32.QueryP(500,&value)
	log("p10 r:%d", value)
	countreads(value,10,32)
	l.Unlock()

	delay()
	l.Lock()
	value = 1
	s1.Put(7,value)
	log("p10 w:%d", value)
	countwrites(value,10,1)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s9.QueryP(139,&value)
	log("p10 r:%d", value)
	countreads(value,10,9)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s2.QueryP(22,&value)
	log("p10 r:%d", value)
	countreads(value,10,2)
	l.Unlock()

	delay()
	l.Lock()
	value = 7
	s7.Put(110,value)
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
	s31.QueryP(496,&value)
	log("p11 r:%d", value)
	countreads(value,11,31)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s27.QueryP(425,&value)
	log("p11 r:%d", value)
	countreads(value,11,27)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s21.QueryP(323,&value)
	log("p11 r:%d", value)
	countreads(value,11,21)
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
	value = 5
	s5.Put(79,value)
	log("p11 w:%d", value)
	countwrites(value,11,5)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s25.QueryP(387,&value)
	log("p11 r:%d", value)
	countreads(value,11,25)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s1.QueryP(12,&value)
	log("p11 r:%d", value)
	countreads(value,11,1)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s1.QueryP(5,&value)
	log("p11 r:%d", value)
	countreads(value,11,1)
	l.Unlock()

	delay()
	l.Lock()
	value = 1
	s1.Put(16,value)
	log("p11 w:%d", value)
	countwrites(value,11,1)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s19.QueryP(297,&value)
	log("p11 r:%d", value)
	countreads(value,11,19)
	l.Unlock()

	delay()
	l.Lock()
	value = 32
	s32.Put(505,value)
	log("p11 w:%d", value)
	countwrites(value,11,32)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s29.QueryP(463,&value)
	log("p11 r:%d", value)
	countreads(value,11,29)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s23.QueryP(357,&value)
	log("p11 r:%d", value)
	countreads(value,11,23)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s17.QueryP(263,&value)
	log("p11 r:%d", value)
	countreads(value,11,17)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s17.QueryP(267,&value)
	log("p11 r:%d", value)
	countreads(value,11,17)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s17.QueryP(270,&value)
	log("p11 r:%d", value)
	countreads(value,11,17)
	l.Unlock()
}


func p12() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	s13.QueryP(199,&value)
	log("p12 r:%d", value)
	countreads(value,12,13)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s31.QueryP(486,&value)
	log("p12 r:%d", value)
	countreads(value,12,31)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s16.QueryP(247,&value)
	log("p12 r:%d", value)
	countreads(value,12,16)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s29.QueryP(454,&value)
	log("p12 r:%d", value)
	countreads(value,12,29)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s15.QueryP(226,&value)
	log("p12 r:%d", value)
	countreads(value,12,15)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s26.QueryP(413,&value)
	log("p12 r:%d", value)
	countreads(value,12,26)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s13.QueryP(208,&value)
	log("p12 r:%d", value)
	countreads(value,12,13)
	l.Unlock()

	delay()
	l.Lock()
	value = 21
	s21.Put(332,value)
	log("p12 w:%d", value)
	countwrites(value,12,21)
	l.Unlock()

	delay()
	l.Lock()
	value = 1
	s1.Put(3,value)
	log("p12 w:%d", value)
	countwrites(value,12,1)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s3.QueryP(48,&value)
	log("p12 r:%d", value)
	countreads(value,12,3)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s15.QueryP(232,&value)
	log("p12 r:%d", value)
	countreads(value,12,15)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s9.QueryP(132,&value)
	log("p12 r:%d", value)
	countreads(value,12,9)
	l.Unlock()

	delay()
	l.Lock()
	value = 14
	s14.Put(212,value)
	log("p12 w:%d", value)
	countwrites(value,12,14)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s22.QueryP(338,&value)
	log("p12 r:%d", value)
	countreads(value,12,22)
	l.Unlock()

	delay()
	l.Lock()
	value = 18
	s18.Put(282,value)
	log("p12 w:%d", value)
	countwrites(value,12,18)
	l.Unlock()

	delay()
	l.Lock()
	value = 16
	s16.Put(245,value)
	log("p12 w:%d", value)
	countwrites(value,12,16)
	l.Unlock()
}


func p13() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 11
	s11.Put(168,value)
	log("p13 w:%d", value)
	countwrites(value,13,11)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s30.QueryP(474,&value)
	log("p13 r:%d", value)
	countreads(value,13,30)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s29.QueryP(459,&value)
	log("p13 r:%d", value)
	countreads(value,13,29)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s8.QueryP(123,&value)
	log("p13 r:%d", value)
	countreads(value,13,8)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s6.QueryP(87,&value)
	log("p13 r:%d", value)
	countreads(value,13,6)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s7.QueryP(107,&value)
	log("p13 r:%d", value)
	countreads(value,13,7)
	l.Unlock()

	delay()
	l.Lock()
	value = 23
	s23.Put(363,value)
	log("p13 w:%d", value)
	countwrites(value,13,23)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s1.QueryP(7,&value)
	log("p13 r:%d", value)
	countreads(value,13,1)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s9.QueryP(143,&value)
	log("p13 r:%d", value)
	countreads(value,13,9)
	l.Unlock()

	delay()
	l.Lock()
	value = 30
	s30.Put(465,value)
	log("p13 w:%d", value)
	countwrites(value,13,30)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s3.QueryP(37,&value)
	log("p13 r:%d", value)
	countreads(value,13,3)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s29.QueryP(451,&value)
	log("p13 r:%d", value)
	countreads(value,13,29)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s22.QueryP(344,&value)
	log("p13 r:%d", value)
	countreads(value,13,22)
	l.Unlock()

	delay()
	l.Lock()
	value = 9
	s9.Put(131,value)
	log("p13 w:%d", value)
	countwrites(value,13,9)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s26.QueryP(403,&value)
	log("p13 r:%d", value)
	countreads(value,13,26)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s14.QueryP(213,&value)
	log("p13 r:%d", value)
	countreads(value,13,14)
	l.Unlock()
}

func p14() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	s13.QueryP(193,&value)
	log("p14 r:%d", value)
	countreads(value,14,13)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s22.QueryP(345,&value)
	log("p14 r:%d", value)
	countreads(value,14,22)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s2.QueryP(27,&value)
	log("p14 r:%d", value)
	countreads(value,14,2)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s24.QueryP(381,&value)
	log("p14 r:%d", value)
	countreads(value,14,24)
	l.Unlock()

	delay()
	l.Lock()
	value = 32
	s32.Put(498,value)
	log("p14 w:%d", value)
	countwrites(value,14,32)
	l.Unlock()

	delay()
	l.Lock()
	value = 17
	s17.Put(270,value)
	log("p14 w:%d", value)
	countwrites(value,14,17)
	l.Unlock()

	delay()
	l.Lock()
	value = 13
	s13.Put(197,value)
	log("p14 w:%d", value)
	countwrites(value,14,13)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s4.QueryP(50,&value)
	log("p14 r:%d", value)
	countreads(value,14,4)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s13.QueryP(199,&value)
	log("p14 r:%d", value)
	countreads(value,14,13)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s1.QueryP(11,&value)
	log("p14 r:%d", value)
	countreads(value,14,1)
	l.Unlock()

	delay()
	l.Lock()
	value = 31
	s31.Put(486,value)
	log("p14 w:%d", value)
	countwrites(value,14,31)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s12.QueryP(177,&value)
	log("p14 r:%d", value)
	countreads(value,14,12)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s16.QueryP(242,&value)
	log("p14 r:%d", value)
	countreads(value,14,16)
	l.Unlock()

	delay()
	l.Lock()
	value = 21
	s21.Put(328,value)
	log("p14 w:%d", value)
	countwrites(value,14,21)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s13.QueryP(202,&value)
	log("p14 r:%d", value)
	countreads(value,14,13)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s23.QueryP(357,&value)
	log("p14 r:%d", value)
	countreads(value,14,23)
	l.Unlock()
}


func p15() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 13
	s13.Put(196,value)
	log("p15 w:%d", value)
	countwrites(value,15,13)
	l.Unlock()

	delay()
	l.Lock()
	value = 20
	s20.Put(313,value)
	log("p15 w:%d", value)
	countwrites(value,15,20)
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
	s24.QueryP(380,&value)
	log("p15 r:%d", value)
	countreads(value,15,24)
	l.Unlock()

	delay()
	l.Lock()
	value = 1
	s1.Put(7,value)
	log("p15 w:%d", value)
	countwrites(value,15,1)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s15.QueryP(225,&value)
	log("p15 r:%d", value)
	countreads(value,15,15)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s31.QueryP(481,&value)
	log("p15 r:%d", value)
	countreads(value,15,31)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s13.QueryP(196,&value)
	log("p15 r:%d", value)
	countreads(value,15,13)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s3.QueryP(37,&value)
	log("p15 r:%d", value)
	countreads(value,15,3)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s28.QueryP(435,&value)
	log("p15 r:%d", value)
	countreads(value,15,28)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s13.QueryP(193,&value)
	log("p15 r:%d", value)
	countreads(value,15,13)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s7.QueryP(109,&value)
	log("p15 r:%d", value)
	countreads(value,15,7)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s21.QueryP(328,&value)
	log("p15 r:%d", value)
	countreads(value,15,21)
	l.Unlock()

	delay()
	l.Lock()
	value = 25
	s25.Put(399,value)
	log("p15 w:%d", value)
	countwrites(value,15,25)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s14.QueryP(210,&value)
	log("p15 r:%d", value)
	countreads(value,15,14)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s5.QueryP(65,&value)
	log("p15 r:%d", value)
	countreads(value,15,5)
	l.Unlock()
}


func p16() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	s18.QueryP(288,&value)
	log("p16 r:%d", value)
	countreads(value,16,18)
	l.Unlock()

	delay()
	l.Lock()
	value = 2
	s2.Put(28,value)
	log("p16 w:%d", value)
	countwrites(value,16,2)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s26.QueryP(414,&value)
	log("p16 r:%d", value)
	countreads(value,16,26)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s26.QueryP(401,&value)
	log("p16 r:%d", value)
	countreads(value,16,26)
	l.Unlock()

	delay()
	l.Lock()
	value = 4
	s4.Put(53,value)
	log("p16 w:%d", value)
	countwrites(value,16,4)
	l.Unlock()

	delay()
	l.Lock()
	value = 11
	s11.Put(168,value)
	log("p16 w:%d", value)
	countwrites(value,16,11)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s13.QueryP(198,&value)
	log("p16 r:%d", value)
	countreads(value,16,13)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s28.QueryP(440,&value)
	log("p16 r:%d", value)
	countreads(value,16,28)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s13.QueryP(206,&value)
	log("p16 r:%d", value)
	countreads(value,16,13)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s6.QueryP(85,&value)
	log("p16 r:%d", value)
	countreads(value,16,6)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s15.QueryP(236,&value)
	log("p16 r:%d", value)
	countreads(value,16,15)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s17.QueryP(262,&value)
	log("p16 r:%d", value)
	countreads(value,16,17)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s24.QueryP(382,&value)
	log("p16 r:%d", value)
	countreads(value,16,24)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s20.QueryP(320,&value)
	log("p16 r:%d", value)
	countreads(value,16,20)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s9.QueryP(137,&value)
	log("p16 r:%d", value)
	countreads(value,16,9)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s17.QueryP(258,&value)
	log("p16 r:%d", value)
	countreads(value,16,17)
	l.Unlock()
}


func p17() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	s8.QueryP(121,&value)
	log("p17 r:%d", value)
	countreads(value,17,8)
	l.Unlock()

	delay()
	l.Lock()
	value = 10
	s10.Put(151,value)
	log("p17 w:%d", value)
	countwrites(value,17,10)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s21.QueryP(329,&value)
	log("p17 r:%d", value)
	countreads(value,17,21)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s31.QueryP(489,&value)
	log("p17 r:%d", value)
	countreads(value,17,31)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s20.QueryP(306,&value)
	log("p17 r:%d", value)
	countreads(value,17,20)
	l.Unlock()

	delay()
	l.Lock()
	value = 1
	s1.Put(5,value)
	log("p17 w:%d", value)
	countwrites(value,17,1)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s14.QueryP(211,&value)
	log("p17 r:%d", value)
	countreads(value,17,14)
	l.Unlock()

	delay()
	l.Lock()
	value = 15
	s15.Put(232,value)
	log("p17 w:%d", value)
	countwrites(value,17,15)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s4.QueryP(55,&value)
	log("p17 r:%d", value)
	countreads(value,17,4)
	l.Unlock()

	delay()
	l.Lock()
	value = 14
	s14.Put(213,value)
	log("p17 w:%d", value)
	countwrites(value,17,14)
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
	s28.QueryP(442,&value)
	log("p17 r:%d", value)
	countreads(value,17,28)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s12.QueryP(185,&value)
	log("p17 r:%d", value)
	countreads(value,17,12)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s6.QueryP(92,&value)
	log("p17 r:%d", value)
	countreads(value,17,6)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s19.QueryP(303,&value)
	log("p17 r:%d", value)
	countreads(value,17,19)
	l.Unlock()

	delay()
	l.Lock()
	value = 30
	s30.Put(479,value)
	log("p17 w:%d", value)
	countwrites(value,17,30)
	l.Unlock()
}


func p18() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	s26.QueryP(414,&value)
	log("p18 r:%d", value)
	countreads(value,18,26)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s9.QueryP(130,&value)
	log("p18 r:%d", value)
	countreads(value,18,9)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s18.QueryP(276,&value)
	log("p18 r:%d", value)
	countreads(value,18,18)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s7.QueryP(108,&value)
	log("p18 r:%d", value)
	countreads(value,18,7)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s26.QueryP(405,&value)
	log("p18 r:%d", value)
	countreads(value,18,26)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s16.QueryP(248,&value)
	log("p18 r:%d", value)
	countreads(value,18,16)
	l.Unlock()

	delay()
	l.Lock()
	value = 10
	s10.Put(159,value)
	log("p18 w:%d", value)
	countwrites(value,18,10)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s20.QueryP(308,&value)
	log("p18 r:%d", value)
	countreads(value,18,20)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s27.QueryP(431,&value)
	log("p18 r:%d", value)
	countreads(value,18,27)
	l.Unlock()

	delay()
	l.Lock()
	value = 17
	s17.Put(267,value)
	log("p18 w:%d", value)
	countwrites(value,18,17)
	l.Unlock()

	delay()
	l.Lock()
	value = 3
	s3.Put(44,value)
	log("p18 w:%d", value)
	countwrites(value,18,3)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s12.QueryP(182,&value)
	log("p18 r:%d", value)
	countreads(value,18,12)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s17.QueryP(271,&value)
	log("p18 r:%d", value)
	countreads(value,18,17)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s19.QueryP(296,&value)
	log("p18 r:%d", value)
	countreads(value,18,19)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s24.QueryP(379,&value)
	log("p18 r:%d", value)
	countreads(value,18,24)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s9.QueryP(129,&value)
	log("p18 r:%d", value)
	countreads(value,18,9)
	l.Unlock()
}


func p19() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 24
	s24.Put(374,value)
	log("p19 w:%d", value)
	countwrites(value,19,24)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s26.QueryP(413,&value)
	log("p19 r:%d", value)
	countreads(value,19,26)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s3.QueryP(47,&value)
	log("p19 r:%d", value)
	countreads(value,19,3)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s7.QueryP(108,&value)
	log("p19 r:%d", value)
	countreads(value,19,7)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s19.QueryP(297,&value)
	log("p19 r:%d", value)
	countreads(value,19,19)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s15.QueryP(233,&value)
	log("p19 r:%d", value)
	countreads(value,19,15)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s10.QueryP(154,&value)
	log("p19 r:%d", value)
	countreads(value,19,10)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s29.QueryP(463,&value)
	log("p19 r:%d", value)
	countreads(value,19,29)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s18.QueryP(278,&value)
	log("p19 r:%d", value)
	countreads(value,19,18)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s25.QueryP(393,&value)
	log("p19 r:%d", value)
	countreads(value,19,25)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s4.QueryP(53,&value)
	log("p19 r:%d", value)
	countreads(value,19,4)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s5.QueryP(71,&value)
	log("p19 r:%d", value)
	countreads(value,19,5)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s19.QueryP(300,&value)
	log("p19 r:%d", value)
	countreads(value,19,19)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s2.QueryP(24,&value)
	log("p19 r:%d", value)
	countreads(value,19,2)
	l.Unlock()

	delay()
	l.Lock()
	value = 4
	s4.Put(61,value)
	log("p19 w:%d", value)
	countwrites(value,19,4)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s18.QueryP(281,&value)
	log("p19 r:%d", value)
	countreads(value,19,18)
	l.Unlock()
}


func p20() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	s16.QueryP(251,&value)
	log("p20 r:%d", value)
	countreads(value,20,16)
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
	value = 32
	s32.Put(508,value)
	log("p20 w:%d", value)
	countwrites(value,20,32)
	l.Unlock()

	delay()
	l.Lock()
	value = 22
	s22.Put(345,value)
	log("p20 w:%d", value)
	countwrites(value,20,22)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s14.QueryP(220,&value)
	log("p20 r:%d", value)
	countreads(value,20,14)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s7.QueryP(109,&value)
	log("p20 r:%d", value)
	countreads(value,20,7)
	l.Unlock()

	delay()
	l.Lock()
	value = 13
	s13.Put(207,value)
	log("p20 w:%d", value)
	countwrites(value,20,13)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s5.QueryP(74,&value)
	log("p20 r:%d", value)
	countreads(value,20,5)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s2.QueryP(29,&value)
	log("p20 r:%d", value)
	countreads(value,20,2)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s28.QueryP(440,&value)
	log("p20 r:%d", value)
	countreads(value,20,28)
	l.Unlock()

	delay()
	l.Lock()
	value = 4
	s4.Put(63,value)
	log("p20 w:%d", value)
	countwrites(value,20,4)
	l.Unlock()

	delay()
	l.Lock()
	value = 19
	s19.Put(298,value)
	log("p20 w:%d", value)
	countwrites(value,20,19)
	l.Unlock()

	delay()
	l.Lock()
	value = 13
	s13.Put(198,value)
	log("p20 w:%d", value)
	countwrites(value,20,13)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s19.QueryP(302,&value)
	log("p20 r:%d", value)
	countreads(value,20,19)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s12.QueryP(190,&value)
	log("p20 r:%d", value)
	countreads(value,20,12)
	l.Unlock()

	delay()
	l.Lock()
	value = 20
	s20.Put(305,value)
	log("p20 w:%d", value)
	countwrites(value,20,20)
	l.Unlock()
}


func p21() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 29
	s29.Put(463,value)
	log("p21 w:%d", value)
	countwrites(value,21,29)
	l.Unlock()

	delay()
	l.Lock()
	value = 19
	s19.Put(297,value)
	log("p21 w:%d", value)
	countwrites(value,21,19)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s9.QueryP(131,&value)
	log("p21 r:%d", value)
	countreads(value,21,9)
	l.Unlock()

	delay()
	l.Lock()
	value = 4
	s4.Put(52,value)
	log("p21 w:%d", value)
	countwrites(value,21,4)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s22.QueryP(351,&value)
	log("p21 r:%d", value)
	countreads(value,21,22)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s26.QueryP(409,&value)
	log("p21 r:%d", value)
	countreads(value,21,26)
	l.Unlock()

	delay()
	l.Lock()
	value = 31
	s31.Put(490,value)
	log("p21 w:%d", value)
	countwrites(value,21,31)
	l.Unlock()

	delay()
	l.Lock()
	value = 23
	s23.Put(357,value)
	log("p21 w:%d", value)
	countwrites(value,21,23)
	l.Unlock()

	delay()
	l.Lock()
	value = 11
	s11.Put(170,value)
	log("p21 w:%d", value)
	countwrites(value,21,11)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s2.QueryP(21,&value)
	log("p21 r:%d", value)
	countreads(value,21,2)
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
	value = 14
	s14.Put(211,value)
	log("p21 w:%d", value)
	countwrites(value,21,14)
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
	value = -1
	s1.QueryP(3,&value)
	log("p21 r:%d", value)
	countreads(value,21,1)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s30.QueryP(479,&value)
	log("p21 r:%d", value)
	countreads(value,21,30)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s28.QueryP(439,&value)
	log("p21 r:%d", value)
	countreads(value,21,28)
	l.Unlock()
}


func p22() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	s19.QueryP(295,&value)
	log("p22 r:%d", value)
	countreads(value,22,19)
	l.Unlock()

	delay()
	l.Lock()
	value = 10
	s10.Put(148,value)
	log("p22 w:%d", value)
	countwrites(value,22,10)
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
	value = -1
	s7.QueryP(105,&value)
	log("p22 r:%d", value)
	countreads(value,22,7)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s32.QueryP(500,&value)
	log("p22 r:%d", value)
	countreads(value,22,32)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s24.QueryP(372,&value)
	log("p22 r:%d", value)
	countreads(value,22,24)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s7.QueryP(98,&value)
	log("p22 r:%d", value)
	countreads(value,22,7)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s26.QueryP(411,&value)
	log("p22 r:%d", value)
	countreads(value,22,26)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s16.QueryP(252,&value)
	log("p22 r:%d", value)
	countreads(value,22,16)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s26.QueryP(413,&value)
	log("p22 r:%d", value)
	countreads(value,22,26)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s26.QueryP(406,&value)
	log("p22 r:%d", value)
	countreads(value,22,26)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s1.QueryP(2,&value)
	log("p22 r:%d", value)
	countreads(value,22,1)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s15.QueryP(235,&value)
	log("p22 r:%d", value)
	countreads(value,22,15)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s2.QueryP(17,&value)
	log("p22 r:%d", value)
	countreads(value,22,2)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s7.QueryP(100,&value)
	log("p22 r:%d", value)
	countreads(value,22,7)
	l.Unlock()

	delay()
	l.Lock()
	value = 27
	s27.Put(422,value)
	log("p22 w:%d", value)
	countwrites(value,22,27)
	l.Unlock()
}


func p23() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 10
	s10.Put(153,value)
	log("p23 w:%d", value)
	countwrites(value,23,10)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s14.QueryP(209,&value)
	log("p23 r:%d", value)
	countreads(value,23,14)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s12.QueryP(189,&value)
	log("p23 r:%d", value)
	countreads(value,23,12)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s17.QueryP(270,&value)
	log("p23 r:%d", value)
	countreads(value,23,17)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s22.QueryP(351,&value)
	log("p23 r:%d", value)
	countreads(value,23,22)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s2.QueryP(26,&value)
	log("p23 r:%d", value)
	countreads(value,23,2)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s17.QueryP(259,&value)
	log("p23 r:%d", value)
	countreads(value,23,17)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s15.QueryP(231,&value)
	log("p23 r:%d", value)
	countreads(value,23,15)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s14.QueryP(223,&value)
	log("p23 r:%d", value)
	countreads(value,23,14)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s6.QueryP(87,&value)
	log("p23 r:%d", value)
	countreads(value,23,6)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s18.QueryP(286,&value)
	log("p23 r:%d", value)
	countreads(value,23,18)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s6.QueryP(87,&value)
	log("p23 r:%d", value)
	countreads(value,23,6)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s19.QueryP(299,&value)
	log("p23 r:%d", value)
	countreads(value,23,19)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s12.QueryP(187,&value)
	log("p23 r:%d", value)
	countreads(value,23,12)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s1.QueryP(6,&value)
	log("p23 r:%d", value)
	countreads(value,23,1)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s8.QueryP(126,&value)
	log("p23 r:%d", value)
	countreads(value,23,8)
	l.Unlock()
}


func p24() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	s1.QueryP(1,&value)
	log("p24 r:%d", value)
	countreads(value,24,1)
	l.Unlock()

	delay()
	l.Lock()
	value = 23
	s23.Put(366,value)
	log("p24 w:%d", value)
	countwrites(value,24,23)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s16.QueryP(246,&value)
	log("p24 r:%d", value)
	countreads(value,24,16)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s12.QueryP(192,&value)
	log("p24 r:%d", value)
	countreads(value,24,12)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s16.QueryP(246,&value)
	log("p24 r:%d", value)
	countreads(value,24,16)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s28.QueryP(433,&value)
	log("p24 r:%d", value)
	countreads(value,24,28)
	l.Unlock()

	delay()
	l.Lock()
	value = 21
	s21.Put(329,value)
	log("p24 w:%d", value)
	countwrites(value,24,21)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s18.QueryP(274,&value)
	log("p24 r:%d", value)
	countreads(value,24,18)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s28.QueryP(445,&value)
	log("p24 r:%d", value)
	countreads(value,24,28)
	l.Unlock()

	delay()
	l.Lock()
	value = 23
	s23.Put(367,value)
	log("p24 w:%d", value)
	countwrites(value,24,23)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s11.QueryP(169,&value)
	log("p24 r:%d", value)
	countreads(value,24,11)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s29.QueryP(456,&value)
	log("p24 r:%d", value)
	countreads(value,24,29)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s1.QueryP(12,&value)
	log("p24 r:%d", value)
	countreads(value,24,1)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s1.QueryP(12,&value)
	log("p24 r:%d", value)
	countreads(value,24,1)
	l.Unlock()

	delay()
	l.Lock()
	value = 6
	s6.Put(90,value)
	log("p24 w:%d", value)
	countwrites(value,24,6)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s29.QueryP(462,&value)
	log("p24 r:%d", value)
	countreads(value,24,29)
	l.Unlock()
}


func p25() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	s21.QueryP(330,&value)
	log("p25 r:%d", value)
	countreads(value,25,21)
	l.Unlock()

	delay()
	l.Lock()
	value = 3
	s3.Put(46,value)
	log("p25 w:%d", value)
	countwrites(value,25,3)
	l.Unlock()

	delay()
	l.Lock()
	value = 18
	s18.Put(277,value)
	log("p25 w:%d", value)
	countwrites(value,25,18)
	l.Unlock()

	delay()
	l.Lock()
	value = 18
	s18.Put(282,value)
	log("p25 w:%d", value)
	countwrites(value,25,18)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s9.QueryP(130,&value)
	log("p25 r:%d", value)
	countreads(value,25,9)
	l.Unlock()

	delay()
	l.Lock()
	value = 24
	s24.Put(384,value)
	log("p25 w:%d", value)
	countwrites(value,25,24)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s31.QueryP(488,&value)
	log("p25 r:%d", value)
	countreads(value,25,31)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s20.QueryP(313,&value)
	log("p25 r:%d", value)
	countreads(value,25,20)
	l.Unlock()

	delay()
	l.Lock()
	value = 6
	s6.Put(82,value)
	log("p25 w:%d", value)
	countwrites(value,25,6)
	l.Unlock()

	delay()
	l.Lock()
	value = 17
	s17.Put(268,value)
	log("p25 w:%d", value)
	countwrites(value,25,17)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s19.QueryP(304,&value)
	log("p25 r:%d", value)
	countreads(value,25,19)
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
	s5.QueryP(77,&value)
	log("p25 r:%d", value)
	countreads(value,25,5)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s3.QueryP(45,&value)
	log("p25 r:%d", value)
	countreads(value,25,3)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s15.QueryP(231,&value)
	log("p25 r:%d", value)
	countreads(value,25,15)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s31.QueryP(482,&value)
	log("p25 r:%d", value)
	countreads(value,25,31)
	l.Unlock()
}


func p26() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	s3.QueryP(35,&value)
	log("p26 r:%d", value)
	countreads(value,26,3)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s4.QueryP(58,&value)
	log("p26 r:%d", value)
	countreads(value,26,4)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s25.QueryP(394,&value)
	log("p26 r:%d", value)
	countreads(value,26,25)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s28.QueryP(437,&value)
	log("p26 r:%d", value)
	countreads(value,26,28)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s10.QueryP(149,&value)
	log("p26 r:%d", value)
	countreads(value,26,10)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s17.QueryP(269,&value)
	log("p26 r:%d", value)
	countreads(value,26,17)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s11.QueryP(169,&value)
	log("p26 r:%d", value)
	countreads(value,26,11)
	l.Unlock()

	delay()
	l.Lock()
	value = 19
	s19.Put(290,value)
	log("p26 w:%d", value)
	countwrites(value,26,19)
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
	s30.QueryP(470,&value)
	log("p26 r:%d", value)
	countreads(value,26,30)
	l.Unlock()

	delay()
	l.Lock()
	value = 7
	s7.Put(101,value)
	log("p26 w:%d", value)
	countwrites(value,26,7)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s15.QueryP(227,&value)
	log("p26 r:%d", value)
	countreads(value,26,15)
	l.Unlock()

	delay()
	l.Lock()
	value = 23
	s23.Put(365,value)
	log("p26 w:%d", value)
	countwrites(value,26,23)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s12.QueryP(180,&value)
	log("p26 r:%d", value)
	countreads(value,26,12)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s16.QueryP(249,&value)
	log("p26 r:%d", value)
	countreads(value,26,16)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s1.QueryP(1,&value)
	log("p26 r:%d", value)
	countreads(value,26,1)
	l.Unlock()
}


func p27() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	s7.QueryP(108,&value)
	log("p27 r:%d", value)
	countreads(value,27,7)
	l.Unlock()

	delay()
	l.Lock()
	value = 16
	s16.Put(247,value)
	log("p27 w:%d", value)
	countwrites(value,27,16)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s2.QueryP(30,&value)
	log("p27 r:%d", value)
	countreads(value,27,2)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s27.QueryP(430,&value)
	log("p27 r:%d", value)
	countreads(value,27,27)
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
	s9.QueryP(133,&value)
	log("p27 r:%d", value)
	countreads(value,27,9)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s28.QueryP(437,&value)
	log("p27 r:%d", value)
	countreads(value,27,28)
	l.Unlock()

	delay()
	l.Lock()
	value = 13
	s13.Put(199,value)
	log("p27 w:%d", value)
	countwrites(value,27,13)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s10.QueryP(154,&value)
	log("p27 r:%d", value)
	countreads(value,27,10)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s3.QueryP(40,&value)
	log("p27 r:%d", value)
	countreads(value,27,3)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s28.QueryP(448,&value)
	log("p27 r:%d", value)
	countreads(value,27,28)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s20.QueryP(313,&value)
	log("p27 r:%d", value)
	countreads(value,27,20)
	l.Unlock()

	delay()
	l.Lock()
	value = 24
	s24.Put(378,value)
	log("p27 w:%d", value)
	countwrites(value,27,24)
	l.Unlock()

	delay()
	l.Lock()
	value = 6
	s6.Put(95,value)
	log("p27 w:%d", value)
	countwrites(value,27,6)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s19.QueryP(291,&value)
	log("p27 r:%d", value)
	countreads(value,27,19)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s17.QueryP(263,&value)
	log("p27 r:%d", value)
	countreads(value,27,17)
	l.Unlock()
}


func p28() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	s5.QueryP(78,&value)
	log("p28 r:%d", value)
	countreads(value,28,5)
	l.Unlock()

	delay()
	l.Lock()
	value = 24
	s24.Put(374,value)
	log("p28 w:%d", value)
	countwrites(value,28,24)
	l.Unlock()

	delay()
	l.Lock()
	value = 1
	s1.Put(6,value)
	log("p28 w:%d", value)
	countwrites(value,28,1)
	l.Unlock()

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
	s32.QueryP(505,&value)
	log("p28 r:%d", value)
	countreads(value,28,32)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s16.QueryP(252,&value)
	log("p28 r:%d", value)
	countreads(value,28,16)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s22.QueryP(337,&value)
	log("p28 r:%d", value)
	countreads(value,28,22)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s17.QueryP(269,&value)
	log("p28 r:%d", value)
	countreads(value,28,17)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s20.QueryP(306,&value)
	log("p28 r:%d", value)
	countreads(value,28,20)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s9.QueryP(140,&value)
	log("p28 r:%d", value)
	countreads(value,28,9)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s14.QueryP(211,&value)
	log("p28 r:%d", value)
	countreads(value,28,14)
	l.Unlock()

	delay()
	l.Lock()
	value = 22
	s22.Put(339,value)
	log("p28 w:%d", value)
	countwrites(value,28,22)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s5.QueryP(69,&value)
	log("p28 r:%d", value)
	countreads(value,28,5)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s11.QueryP(175,&value)
	log("p28 r:%d", value)
	countreads(value,28,11)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s23.QueryP(356,&value)
	log("p28 r:%d", value)
	countreads(value,28,23)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s3.QueryP(39,&value)
	log("p28 r:%d", value)
	countreads(value,28,3)
	l.Unlock()
}


func p29() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	s3.QueryP(48,&value)
	log("p29 r:%d", value)
	countreads(value,29,3)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s18.QueryP(283,&value)
	log("p29 r:%d", value)
	countreads(value,29,18)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s10.QueryP(146,&value)
	log("p29 r:%d", value)
	countreads(value,29,10)
	l.Unlock()

	delay()
	l.Lock()
	value = 26
	s26.Put(414,value)
	log("p29 w:%d", value)
	countwrites(value,29,26)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s10.QueryP(159,&value)
	log("p29 r:%d", value)
	countreads(value,29,10)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s30.QueryP(470,&value)
	log("p29 r:%d", value)
	countreads(value,29,30)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s24.QueryP(377,&value)
	log("p29 r:%d", value)
	countreads(value,29,24)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s26.QueryP(411,&value)
	log("p29 r:%d", value)
	countreads(value,29,26)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s16.QueryP(255,&value)
	log("p29 r:%d", value)
	countreads(value,29,16)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s27.QueryP(422,&value)
	log("p29 r:%d", value)
	countreads(value,29,27)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s3.QueryP(43,&value)
	log("p29 r:%d", value)
	countreads(value,29,3)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s3.QueryP(45,&value)
	log("p29 r:%d", value)
	countreads(value,29,3)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s20.QueryP(308,&value)
	log("p29 r:%d", value)
	countreads(value,29,20)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s29.QueryP(450,&value)
	log("p29 r:%d", value)
	countreads(value,29,29)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s13.QueryP(203,&value)
	log("p29 r:%d", value)
	countreads(value,29,13)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s29.QueryP(453,&value)
	log("p29 r:%d", value)
	countreads(value,29,29)
	l.Unlock()
}


func p30() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	s28.QueryP(447,&value)
	log("p30 r:%d", value)
	countreads(value,30,28)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s24.QueryP(379,&value)
	log("p30 r:%d", value)
	countreads(value,30,24)
	l.Unlock()

	delay()
	l.Lock()
	value = 22
	s22.Put(346,value)
	log("p30 w:%d", value)
	countwrites(value,30,22)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s21.QueryP(321,&value)
	log("p30 r:%d", value)
	countreads(value,30,21)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s13.QueryP(197,&value)
	log("p30 r:%d", value)
	countreads(value,30,13)
	l.Unlock()

	delay()
	l.Lock()
	value = 2
	s2.Put(18,value)
	log("p30 w:%d", value)
	countwrites(value,30,2)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s8.QueryP(117,&value)
	log("p30 r:%d", value)
	countreads(value,30,8)
	l.Unlock()

	delay()
	l.Lock()
	value = 7
	s7.Put(100,value)
	log("p30 w:%d", value)
	countwrites(value,30,7)
	l.Unlock()

	delay()
	l.Lock()
	value = 10
	s10.Put(148,value)
	log("p30 w:%d", value)
	countwrites(value,30,10)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s23.QueryP(361,&value)
	log("p30 r:%d", value)
	countreads(value,30,23)
	l.Unlock()

	delay()
	l.Lock()
	value = 4
	s4.Put(63,value)
	log("p30 w:%d", value)
	countwrites(value,30,4)
	l.Unlock()

	delay()
	l.Lock()
	value = 28
	s28.Put(437,value)
	log("p30 w:%d", value)
	countwrites(value,30,28)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s13.QueryP(193,&value)
	log("p30 r:%d", value)
	countreads(value,30,13)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s22.QueryP(349,&value)
	log("p30 r:%d", value)
	countreads(value,30,22)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s28.QueryP(439,&value)
	log("p30 r:%d", value)
	countreads(value,30,28)
	l.Unlock()

	delay()
	l.Lock()
	value = 27
	s27.Put(421,value)
	log("p30 w:%d", value)
	countwrites(value,30,27)
	l.Unlock()
}


func p31() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 13
	s13.Put(193,value)
	log("p31 w:%d", value)
	countwrites(value,31,13)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s25.QueryP(398,&value)
	log("p31 r:%d", value)
	countreads(value,31,25)
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
	s28.QueryP(446,&value)
	log("p31 r:%d", value)
	countreads(value,31,28)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s4.QueryP(63,&value)
	log("p31 r:%d", value)
	countreads(value,31,4)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s10.QueryP(158,&value)
	log("p31 r:%d", value)
	countreads(value,31,10)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s15.QueryP(231,&value)
	log("p31 r:%d", value)
	countreads(value,31,15)
	l.Unlock()

	delay()
	l.Lock()
	value = 15
	s15.Put(225,value)
	log("p31 w:%d", value)
	countwrites(value,31,15)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s10.QueryP(157,&value)
	log("p31 r:%d", value)
	countreads(value,31,10)
	l.Unlock()

	delay()
	l.Lock()
	value = 16
	s16.Put(244,value)
	log("p31 w:%d", value)
	countwrites(value,31,16)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s12.QueryP(181,&value)
	log("p31 r:%d", value)
	countreads(value,31,12)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s7.QueryP(111,&value)
	log("p31 r:%d", value)
	countreads(value,31,7)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s26.QueryP(415,&value)
	log("p31 r:%d", value)
	countreads(value,31,26)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s1.QueryP(5,&value)
	log("p31 r:%d", value)
	countreads(value,31,1)
	l.Unlock()

	delay()
	l.Lock()
	value = 27
	s27.Put(429,value)
	log("p31 w:%d", value)
	countwrites(value,31,27)
	l.Unlock()

	delay()
	l.Lock()
	value = 7
	s7.Put(105,value)
	log("p31 w:%d", value)
	countwrites(value,31,7)
	l.Unlock()
}


func p32() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	s25.QueryP(390,&value)
	log("p32 r:%d", value)
	countreads(value,32,25)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s28.QueryP(435,&value)
	log("p32 r:%d", value)
	countreads(value,32,28)
	l.Unlock()

	delay()
	l.Lock()
	value = 3
	s3.Put(41,value)
	log("p32 w:%d", value)
	countwrites(value,32,3)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s26.QueryP(403,&value)
	log("p32 r:%d", value)
	countreads(value,32,26)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s9.QueryP(134,&value)
	log("p32 r:%d", value)
	countreads(value,32,9)
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
	s26.QueryP(405,&value)
	log("p32 r:%d", value)
	countreads(value,32,26)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s15.QueryP(225,&value)
	log("p32 r:%d", value)
	countreads(value,32,15)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s12.QueryP(188,&value)
	log("p32 r:%d", value)
	countreads(value,32,12)
	l.Unlock()

	delay()
	l.Lock()
	value = 13
	s13.Put(208,value)
	log("p32 w:%d", value)
	countwrites(value,32,13)
	l.Unlock()

	delay()
	l.Lock()
	value = 31
	s31.Put(490,value)
	log("p32 w:%d", value)
	countwrites(value,32,31)
	l.Unlock()

	delay()
	l.Lock()
	value = 20
	s20.Put(312,value)
	log("p32 w:%d", value)
	countwrites(value,32,20)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s1.QueryP(11,&value)
	log("p32 r:%d", value)
	countreads(value,32,1)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s13.QueryP(206,&value)
	log("p32 r:%d", value)
	countreads(value,32,13)
	l.Unlock()

	delay()
	l.Lock()
	value = 8
	s8.Put(119,value)
	log("p32 w:%d", value)
	countwrites(value,32,8)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s26.QueryP(412,&value)
	log("p32 r:%d", value)
	countreads(value,32,26)
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
