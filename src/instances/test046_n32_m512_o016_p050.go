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
	s28.QueryP(447,&value)
	log("p1 r:%d", value)
	countreads(value,1,28)
	l.Unlock()

	delay()
	l.Lock()
	value = 21
	s21.Put(329,value)
	log("p1 w:%d", value)
	countwrites(value,1,21)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s18.QueryP(288,&value)
	log("p1 r:%d", value)
	countreads(value,1,18)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s27.QueryP(428,&value)
	log("p1 r:%d", value)
	countreads(value,1,27)
	l.Unlock()

	delay()
	l.Lock()
	value = 27
	s27.Put(431,value)
	log("p1 w:%d", value)
	countwrites(value,1,27)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s3.QueryP(42,&value)
	log("p1 r:%d", value)
	countreads(value,1,3)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s27.QueryP(424,&value)
	log("p1 r:%d", value)
	countreads(value,1,27)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s27.QueryP(431,&value)
	log("p1 r:%d", value)
	countreads(value,1,27)
	l.Unlock()

	delay()
	l.Lock()
	value = 29
	s29.Put(464,value)
	log("p1 w:%d", value)
	countwrites(value,1,29)
	l.Unlock()

	delay()
	l.Lock()
	value = 27
	s27.Put(424,value)
	log("p1 w:%d", value)
	countwrites(value,1,27)
	l.Unlock()

	delay()
	l.Lock()
	value = 12
	s12.Put(192,value)
	log("p1 w:%d", value)
	countwrites(value,1,12)
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
	value = 16
	s16.Put(252,value)
	log("p1 w:%d", value)
	countwrites(value,1,16)
	l.Unlock()

	delay()
	l.Lock()
	value = 2
	s2.Put(27,value)
	log("p1 w:%d", value)
	countwrites(value,1,2)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s28.QueryP(448,&value)
	log("p1 r:%d", value)
	countreads(value,1,28)
	l.Unlock()

	delay()
	l.Lock()
	value = 17
	s17.Put(267,value)
	log("p1 w:%d", value)
	countwrites(value,1,17)
	l.Unlock()
}

func p2() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	s2.QueryP(19,&value)
	log("p2 r:%d", value)
	countreads(value,2,2)
	l.Unlock()

	delay()
	l.Lock()
	value = 22
	s22.Put(343,value)
	log("p2 w:%d", value)
	countwrites(value,2,22)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s31.QueryP(496,&value)
	log("p2 r:%d", value)
	countreads(value,2,31)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s24.QueryP(373,&value)
	log("p2 r:%d", value)
	countreads(value,2,24)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s5.QueryP(69,&value)
	log("p2 r:%d", value)
	countreads(value,2,5)
	l.Unlock()

	delay()
	l.Lock()
	value = 21
	s21.Put(325,value)
	log("p2 w:%d", value)
	countwrites(value,2,21)
	l.Unlock()

	delay()
	l.Lock()
	value = 12
	s12.Put(182,value)
	log("p2 w:%d", value)
	countwrites(value,2,12)
	l.Unlock()

	delay()
	l.Lock()
	value = 4
	s4.Put(63,value)
	log("p2 w:%d", value)
	countwrites(value,2,4)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s17.QueryP(261,&value)
	log("p2 r:%d", value)
	countreads(value,2,17)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s1.QueryP(12,&value)
	log("p2 r:%d", value)
	countreads(value,2,1)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s6.QueryP(95,&value)
	log("p2 r:%d", value)
	countreads(value,2,6)
	l.Unlock()

	delay()
	l.Lock()
	value = 27
	s27.Put(432,value)
	log("p2 w:%d", value)
	countwrites(value,2,27)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s15.QueryP(230,&value)
	log("p2 r:%d", value)
	countreads(value,2,15)
	l.Unlock()

	delay()
	l.Lock()
	value = 27
	s27.Put(431,value)
	log("p2 w:%d", value)
	countwrites(value,2,27)
	l.Unlock()

	delay()
	l.Lock()
	value = 16
	s16.Put(255,value)
	log("p2 w:%d", value)
	countwrites(value,2,16)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s3.QueryP(48,&value)
	log("p2 r:%d", value)
	countreads(value,2,3)
	l.Unlock()
}

func p3() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 9
	s9.Put(142,value)
	log("p3 w:%d", value)
	countwrites(value,3,9)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s18.QueryP(281,&value)
	log("p3 r:%d", value)
	countreads(value,3,18)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s22.QueryP(340,&value)
	log("p3 r:%d", value)
	countreads(value,3,22)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s3.QueryP(46,&value)
	log("p3 r:%d", value)
	countreads(value,3,3)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s21.QueryP(321,&value)
	log("p3 r:%d", value)
	countreads(value,3,21)
	l.Unlock()

	delay()
	l.Lock()
	value = 31
	s31.Put(484,value)
	log("p3 w:%d", value)
	countwrites(value,3,31)
	l.Unlock()

	delay()
	l.Lock()
	value = 7
	s7.Put(100,value)
	log("p3 w:%d", value)
	countwrites(value,3,7)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s13.QueryP(208,&value)
	log("p3 r:%d", value)
	countreads(value,3,13)
	l.Unlock()

	delay()
	l.Lock()
	value = 16
	s16.Put(241,value)
	log("p3 w:%d", value)
	countwrites(value,3,16)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s27.QueryP(417,&value)
	log("p3 r:%d", value)
	countreads(value,3,27)
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
	value = 2
	s2.Put(25,value)
	log("p3 w:%d", value)
	countwrites(value,3,2)
	l.Unlock()

	delay()
	l.Lock()
	value = 23
	s23.Put(359,value)
	log("p3 w:%d", value)
	countwrites(value,3,23)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s14.QueryP(219,&value)
	log("p3 r:%d", value)
	countreads(value,3,14)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s4.QueryP(57,&value)
	log("p3 r:%d", value)
	countreads(value,3,4)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s9.QueryP(129,&value)
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
	s12.QueryP(185,&value)
	log("p4 r:%d", value)
	countreads(value,4,12)
	l.Unlock()

	delay()
	l.Lock()
	value = 2
	s2.Put(17,value)
	log("p4 w:%d", value)
	countwrites(value,4,2)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s4.QueryP(63,&value)
	log("p4 r:%d", value)
	countreads(value,4,4)
	l.Unlock()

	delay()
	l.Lock()
	value = 19
	s19.Put(289,value)
	log("p4 w:%d", value)
	countwrites(value,4,19)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s29.QueryP(461,&value)
	log("p4 r:%d", value)
	countreads(value,4,29)
	l.Unlock()

	delay()
	l.Lock()
	value = 2
	s2.Put(28,value)
	log("p4 w:%d", value)
	countwrites(value,4,2)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s25.QueryP(396,&value)
	log("p4 r:%d", value)
	countreads(value,4,25)
	l.Unlock()

	delay()
	l.Lock()
	value = 2
	s2.Put(20,value)
	log("p4 w:%d", value)
	countwrites(value,4,2)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s22.QueryP(344,&value)
	log("p4 r:%d", value)
	countreads(value,4,22)
	l.Unlock()

	delay()
	l.Lock()
	value = 7
	s7.Put(111,value)
	log("p4 w:%d", value)
	countwrites(value,4,7)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s27.QueryP(419,&value)
	log("p4 r:%d", value)
	countreads(value,4,27)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s19.QueryP(299,&value)
	log("p4 r:%d", value)
	countreads(value,4,19)
	l.Unlock()

	delay()
	l.Lock()
	value = 7
	s7.Put(98,value)
	log("p4 w:%d", value)
	countwrites(value,4,7)
	l.Unlock()

	delay()
	l.Lock()
	value = 8
	s8.Put(117,value)
	log("p4 w:%d", value)
	countwrites(value,4,8)
	l.Unlock()

	delay()
	l.Lock()
	value = 26
	s26.Put(402,value)
	log("p4 w:%d", value)
	countwrites(value,4,26)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s30.QueryP(468,&value)
	log("p4 r:%d", value)
	countreads(value,4,30)
	l.Unlock()
}

func p5() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	s32.QueryP(497,&value)
	log("p5 r:%d", value)
	countreads(value,5,32)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s27.QueryP(419,&value)
	log("p5 r:%d", value)
	countreads(value,5,27)
	l.Unlock()

	delay()
	l.Lock()
	value = 11
	s11.Put(173,value)
	log("p5 w:%d", value)
	countwrites(value,5,11)
	l.Unlock()

	delay()
	l.Lock()
	value = 19
	s19.Put(293,value)
	log("p5 w:%d", value)
	countwrites(value,5,19)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s21.QueryP(329,&value)
	log("p5 r:%d", value)
	countreads(value,5,21)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s21.QueryP(322,&value)
	log("p5 r:%d", value)
	countreads(value,5,21)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s25.QueryP(391,&value)
	log("p5 r:%d", value)
	countreads(value,5,25)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s6.QueryP(91,&value)
	log("p5 r:%d", value)
	countreads(value,5,6)
	l.Unlock()

	delay()
	l.Lock()
	value = 21
	s21.Put(335,value)
	log("p5 w:%d", value)
	countwrites(value,5,21)
	l.Unlock()

	delay()
	l.Lock()
	value = 2
	s2.Put(25,value)
	log("p5 w:%d", value)
	countwrites(value,5,2)
	l.Unlock()

	delay()
	l.Lock()
	value = 17
	s17.Put(271,value)
	log("p5 w:%d", value)
	countwrites(value,5,17)
	l.Unlock()

	delay()
	l.Lock()
	value = 26
	s26.Put(415,value)
	log("p5 w:%d", value)
	countwrites(value,5,26)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s7.QueryP(100,&value)
	log("p5 r:%d", value)
	countreads(value,5,7)
	l.Unlock()

	delay()
	l.Lock()
	value = 11
	s11.Put(164,value)
	log("p5 w:%d", value)
	countwrites(value,5,11)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s22.QueryP(337,&value)
	log("p5 r:%d", value)
	countreads(value,5,22)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s6.QueryP(89,&value)
	log("p5 r:%d", value)
	countreads(value,5,6)
	l.Unlock()
}

func p6() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 30
	s30.Put(476,value)
	log("p6 w:%d", value)
	countwrites(value,6,30)
	l.Unlock()

	delay()
	l.Lock()
	value = 3
	s3.Put(46,value)
	log("p6 w:%d", value)
	countwrites(value,6,3)
	l.Unlock()

	delay()
	l.Lock()
	value = 14
	s14.Put(210,value)
	log("p6 w:%d", value)
	countwrites(value,6,14)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s8.QueryP(127,&value)
	log("p6 r:%d", value)
	countreads(value,6,8)
	l.Unlock()

	delay()
	l.Lock()
	value = 15
	s15.Put(237,value)
	log("p6 w:%d", value)
	countwrites(value,6,15)
	l.Unlock()

	delay()
	l.Lock()
	value = 3
	s3.Put(41,value)
	log("p6 w:%d", value)
	countwrites(value,6,3)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s22.QueryP(340,&value)
	log("p6 r:%d", value)
	countreads(value,6,22)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s26.QueryP(405,&value)
	log("p6 r:%d", value)
	countreads(value,6,26)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s8.QueryP(126,&value)
	log("p6 r:%d", value)
	countreads(value,6,8)
	l.Unlock()

	delay()
	l.Lock()
	value = 17
	s17.Put(264,value)
	log("p6 w:%d", value)
	countwrites(value,6,17)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s14.QueryP(217,&value)
	log("p6 r:%d", value)
	countreads(value,6,14)
	l.Unlock()

	delay()
	l.Lock()
	value = 32
	s32.Put(503,value)
	log("p6 w:%d", value)
	countwrites(value,6,32)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s18.QueryP(275,&value)
	log("p6 r:%d", value)
	countreads(value,6,18)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s2.QueryP(30,&value)
	log("p6 r:%d", value)
	countreads(value,6,2)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s27.QueryP(424,&value)
	log("p6 r:%d", value)
	countreads(value,6,27)
	l.Unlock()

	delay()
	l.Lock()
	value = 1
	s1.Put(9,value)
	log("p6 w:%d", value)
	countwrites(value,6,1)
	l.Unlock()
}

func p7() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	s15.QueryP(231,&value)
	log("p7 r:%d", value)
	countreads(value,7,15)
	l.Unlock()

	delay()
	l.Lock()
	value = 16
	s16.Put(245,value)
	log("p7 w:%d", value)
	countwrites(value,7,16)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s1.QueryP(14,&value)
	log("p7 r:%d", value)
	countreads(value,7,1)
	l.Unlock()

	delay()
	l.Lock()
	value = 4
	s4.Put(61,value)
	log("p7 w:%d", value)
	countwrites(value,7,4)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s11.QueryP(166,&value)
	log("p7 r:%d", value)
	countreads(value,7,11)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s20.QueryP(318,&value)
	log("p7 r:%d", value)
	countreads(value,7,20)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s28.QueryP(435,&value)
	log("p7 r:%d", value)
	countreads(value,7,28)
	l.Unlock()

	delay()
	l.Lock()
	value = 16
	s16.Put(243,value)
	log("p7 w:%d", value)
	countwrites(value,7,16)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s17.QueryP(270,&value)
	log("p7 r:%d", value)
	countreads(value,7,17)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s26.QueryP(415,&value)
	log("p7 r:%d", value)
	countreads(value,7,26)
	l.Unlock()

	delay()
	l.Lock()
	value = 13
	s13.Put(203,value)
	log("p7 w:%d", value)
	countwrites(value,7,13)
	l.Unlock()

	delay()
	l.Lock()
	value = 14
	s14.Put(220,value)
	log("p7 w:%d", value)
	countwrites(value,7,14)
	l.Unlock()

	delay()
	l.Lock()
	value = 23
	s23.Put(359,value)
	log("p7 w:%d", value)
	countwrites(value,7,23)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s12.QueryP(184,&value)
	log("p7 r:%d", value)
	countreads(value,7,12)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s5.QueryP(76,&value)
	log("p7 r:%d", value)
	countreads(value,7,5)
	l.Unlock()

	delay()
	l.Lock()
	value = 13
	s13.Put(197,value)
	log("p7 w:%d", value)
	countwrites(value,7,13)
	l.Unlock()
}

func p8() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 24
	s24.Put(384,value)
	log("p8 w:%d", value)
	countwrites(value,8,24)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s11.QueryP(162,&value)
	log("p8 r:%d", value)
	countreads(value,8,11)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s13.QueryP(206,&value)
	log("p8 r:%d", value)
	countreads(value,8,13)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s13.QueryP(207,&value)
	log("p8 r:%d", value)
	countreads(value,8,13)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s3.QueryP(38,&value)
	log("p8 r:%d", value)
	countreads(value,8,3)
	l.Unlock()

	delay()
	l.Lock()
	value = 16
	s16.Put(246,value)
	log("p8 w:%d", value)
	countwrites(value,8,16)
	l.Unlock()

	delay()
	l.Lock()
	value = 17
	s17.Put(259,value)
	log("p8 w:%d", value)
	countwrites(value,8,17)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s12.QueryP(186,&value)
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
	value = 15
	s15.Put(239,value)
	log("p8 w:%d", value)
	countwrites(value,8,15)
	l.Unlock()

	delay()
	l.Lock()
	value = 4
	s4.Put(50,value)
	log("p8 w:%d", value)
	countwrites(value,8,4)
	l.Unlock()

	delay()
	l.Lock()
	value = 13
	s13.Put(194,value)
	log("p8 w:%d", value)
	countwrites(value,8,13)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s21.QueryP(330,&value)
	log("p8 r:%d", value)
	countreads(value,8,21)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s30.QueryP(469,&value)
	log("p8 r:%d", value)
	countreads(value,8,30)
	l.Unlock()

	delay()
	l.Lock()
	value = 30
	s30.Put(468,value)
	log("p8 w:%d", value)
	countwrites(value,8,30)
	l.Unlock()

	delay()
	l.Lock()
	value = 12
	s12.Put(184,value)
	log("p8 w:%d", value)
	countwrites(value,8,12)
	l.Unlock()
}

func p9() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 10
	s10.Put(145,value)
	log("p9 w:%d", value)
	countwrites(value,9,10)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s5.QueryP(74,&value)
	log("p9 r:%d", value)
	countreads(value,9,5)
	l.Unlock()

	delay()
	l.Lock()
	value = 8
	s8.Put(127,value)
	log("p9 w:%d", value)
	countwrites(value,9,8)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s9.QueryP(144,&value)
	log("p9 r:%d", value)
	countreads(value,9,9)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s22.QueryP(341,&value)
	log("p9 r:%d", value)
	countreads(value,9,22)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s15.QueryP(228,&value)
	log("p9 r:%d", value)
	countreads(value,9,15)
	l.Unlock()

	delay()
	l.Lock()
	value = 12
	s12.Put(185,value)
	log("p9 w:%d", value)
	countwrites(value,9,12)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s9.QueryP(134,&value)
	log("p9 r:%d", value)
	countreads(value,9,9)
	l.Unlock()

	delay()
	l.Lock()
	value = 10
	s10.Put(156,value)
	log("p9 w:%d", value)
	countwrites(value,9,10)
	l.Unlock()

	delay()
	l.Lock()
	value = 22
	s22.Put(343,value)
	log("p9 w:%d", value)
	countwrites(value,9,22)
	l.Unlock()

	delay()
	l.Lock()
	value = 3
	s3.Put(45,value)
	log("p9 w:%d", value)
	countwrites(value,9,3)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s20.QueryP(308,&value)
	log("p9 r:%d", value)
	countreads(value,9,20)
	l.Unlock()

	delay()
	l.Lock()
	value = 5
	s5.Put(66,value)
	log("p9 w:%d", value)
	countwrites(value,9,5)
	l.Unlock()

	delay()
	l.Lock()
	value = 22
	s22.Put(346,value)
	log("p9 w:%d", value)
	countwrites(value,9,22)
	l.Unlock()

	delay()
	l.Lock()
	value = 16
	s16.Put(251,value)
	log("p9 w:%d", value)
	countwrites(value,9,16)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s20.QueryP(315,&value)
	log("p9 r:%d", value)
	countreads(value,9,20)
	l.Unlock()
}


func p10() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	s2.QueryP(18,&value)
	log("p10 r:%d", value)
	countreads(value,10,2)
	l.Unlock()

	delay()
	l.Lock()
	value = 26
	s26.Put(415,value)
	log("p10 w:%d", value)
	countwrites(value,10,26)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s29.QueryP(460,&value)
	log("p10 r:%d", value)
	countreads(value,10,29)
	l.Unlock()

	delay()
	l.Lock()
	value = 23
	s23.Put(357,value)
	log("p10 w:%d", value)
	countwrites(value,10,23)
	l.Unlock()

	delay()
	l.Lock()
	value = 2
	s2.Put(24,value)
	log("p10 w:%d", value)
	countwrites(value,10,2)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s24.QueryP(376,&value)
	log("p10 r:%d", value)
	countreads(value,10,24)
	l.Unlock()

	delay()
	l.Lock()
	value = 29
	s29.Put(457,value)
	log("p10 w:%d", value)
	countwrites(value,10,29)
	l.Unlock()

	delay()
	l.Lock()
	value = 14
	s14.Put(218,value)
	log("p10 w:%d", value)
	countwrites(value,10,14)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s14.QueryP(221,&value)
	log("p10 r:%d", value)
	countreads(value,10,14)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s13.QueryP(199,&value)
	log("p10 r:%d", value)
	countreads(value,10,13)
	l.Unlock()

	delay()
	l.Lock()
	value = 7
	s7.Put(108,value)
	log("p10 w:%d", value)
	countwrites(value,10,7)
	l.Unlock()

	delay()
	l.Lock()
	value = 14
	s14.Put(216,value)
	log("p10 w:%d", value)
	countwrites(value,10,14)
	l.Unlock()

	delay()
	l.Lock()
	value = 15
	s15.Put(240,value)
	log("p10 w:%d", value)
	countwrites(value,10,15)
	l.Unlock()

	delay()
	l.Lock()
	value = 9
	s9.Put(132,value)
	log("p10 w:%d", value)
	countwrites(value,10,9)
	l.Unlock()

	delay()
	l.Lock()
	value = 6
	s6.Put(89,value)
	log("p10 w:%d", value)
	countwrites(value,10,6)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s23.QueryP(353,&value)
	log("p10 r:%d", value)
	countreads(value,10,23)
	l.Unlock()
}


func p11() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	s5.QueryP(76,&value)
	log("p11 r:%d", value)
	countreads(value,11,5)
	l.Unlock()

	delay()
	l.Lock()
	value = 32
	s32.Put(512,value)
	log("p11 w:%d", value)
	countwrites(value,11,32)
	l.Unlock()

	delay()
	l.Lock()
	value = 12
	s12.Put(183,value)
	log("p11 w:%d", value)
	countwrites(value,11,12)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s24.QueryP(372,&value)
	log("p11 r:%d", value)
	countreads(value,11,24)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s11.QueryP(176,&value)
	log("p11 r:%d", value)
	countreads(value,11,11)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s14.QueryP(218,&value)
	log("p11 r:%d", value)
	countreads(value,11,14)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s1.QueryP(4,&value)
	log("p11 r:%d", value)
	countreads(value,11,1)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s8.QueryP(120,&value)
	log("p11 r:%d", value)
	countreads(value,11,8)
	l.Unlock()

	delay()
	l.Lock()
	value = 25
	s25.Put(396,value)
	log("p11 w:%d", value)
	countwrites(value,11,25)
	l.Unlock()

	delay()
	l.Lock()
	value = 20
	s20.Put(310,value)
	log("p11 w:%d", value)
	countwrites(value,11,20)
	l.Unlock()

	delay()
	l.Lock()
	value = 19
	s19.Put(289,value)
	log("p11 w:%d", value)
	countwrites(value,11,19)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s28.QueryP(436,&value)
	log("p11 r:%d", value)
	countreads(value,11,28)
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
	s16.QueryP(253,&value)
	log("p11 r:%d", value)
	countreads(value,11,16)
	l.Unlock()

	delay()
	l.Lock()
	value = 17
	s17.Put(263,value)
	log("p11 w:%d", value)
	countwrites(value,11,17)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s27.QueryP(424,&value)
	log("p11 r:%d", value)
	countreads(value,11,27)
	l.Unlock()
}


func p12() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	s6.QueryP(89,&value)
	log("p12 r:%d", value)
	countreads(value,12,6)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s1.QueryP(4,&value)
	log("p12 r:%d", value)
	countreads(value,12,1)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s18.QueryP(288,&value)
	log("p12 r:%d", value)
	countreads(value,12,18)
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
	s24.QueryP(381,&value)
	log("p12 r:%d", value)
	countreads(value,12,24)
	l.Unlock()

	delay()
	l.Lock()
	value = 25
	s25.Put(386,value)
	log("p12 w:%d", value)
	countwrites(value,12,25)
	l.Unlock()

	delay()
	l.Lock()
	value = 14
	s14.Put(213,value)
	log("p12 w:%d", value)
	countwrites(value,12,14)
	l.Unlock()

	delay()
	l.Lock()
	value = 12
	s12.Put(191,value)
	log("p12 w:%d", value)
	countwrites(value,12,12)
	l.Unlock()

	delay()
	l.Lock()
	value = 22
	s22.Put(352,value)
	log("p12 w:%d", value)
	countwrites(value,12,22)
	l.Unlock()

	delay()
	l.Lock()
	value = 18
	s18.Put(276,value)
	log("p12 w:%d", value)
	countwrites(value,12,18)
	l.Unlock()

	delay()
	l.Lock()
	value = 20
	s20.Put(305,value)
	log("p12 w:%d", value)
	countwrites(value,12,20)
	l.Unlock()

	delay()
	l.Lock()
	value = 15
	s15.Put(226,value)
	log("p12 w:%d", value)
	countwrites(value,12,15)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s15.QueryP(240,&value)
	log("p12 r:%d", value)
	countreads(value,12,15)
	l.Unlock()

	delay()
	l.Lock()
	value = 13
	s13.Put(196,value)
	log("p12 w:%d", value)
	countwrites(value,12,13)
	l.Unlock()

	delay()
	l.Lock()
	value = 25
	s25.Put(387,value)
	log("p12 w:%d", value)
	countwrites(value,12,25)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s29.QueryP(459,&value)
	log("p12 r:%d", value)
	countreads(value,12,29)
	l.Unlock()
}


func p13() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 31
	s31.Put(484,value)
	log("p13 w:%d", value)
	countwrites(value,13,31)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s6.QueryP(89,&value)
	log("p13 r:%d", value)
	countreads(value,13,6)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s8.QueryP(114,&value)
	log("p13 r:%d", value)
	countreads(value,13,8)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s8.QueryP(122,&value)
	log("p13 r:%d", value)
	countreads(value,13,8)
	l.Unlock()

	delay()
	l.Lock()
	value = 32
	s32.Put(508,value)
	log("p13 w:%d", value)
	countwrites(value,13,32)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s23.QueryP(364,&value)
	log("p13 r:%d", value)
	countreads(value,13,23)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s21.QueryP(323,&value)
	log("p13 r:%d", value)
	countreads(value,13,21)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s25.QueryP(386,&value)
	log("p13 r:%d", value)
	countreads(value,13,25)
	l.Unlock()

	delay()
	l.Lock()
	value = 19
	s19.Put(294,value)
	log("p13 w:%d", value)
	countwrites(value,13,19)
	l.Unlock()

	delay()
	l.Lock()
	value = 5
	s5.Put(67,value)
	log("p13 w:%d", value)
	countwrites(value,13,5)
	l.Unlock()

	delay()
	l.Lock()
	value = 16
	s16.Put(244,value)
	log("p13 w:%d", value)
	countwrites(value,13,16)
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
	s28.QueryP(443,&value)
	log("p13 r:%d", value)
	countreads(value,13,28)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s28.QueryP(439,&value)
	log("p13 r:%d", value)
	countreads(value,13,28)
	l.Unlock()

	delay()
	l.Lock()
	value = 28
	s28.Put(445,value)
	log("p13 w:%d", value)
	countwrites(value,13,28)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s29.QueryP(456,&value)
	log("p13 r:%d", value)
	countreads(value,13,29)
	l.Unlock()
}

func p14() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	s29.QueryP(454,&value)
	log("p14 r:%d", value)
	countreads(value,14,29)
	l.Unlock()

	delay()
	l.Lock()
	value = 27
	s27.Put(423,value)
	log("p14 w:%d", value)
	countwrites(value,14,27)
	l.Unlock()

	delay()
	l.Lock()
	value = 9
	s9.Put(136,value)
	log("p14 w:%d", value)
	countwrites(value,14,9)
	l.Unlock()

	delay()
	l.Lock()
	value = 19
	s19.Put(291,value)
	log("p14 w:%d", value)
	countwrites(value,14,19)
	l.Unlock()

	delay()
	l.Lock()
	value = 7
	s7.Put(99,value)
	log("p14 w:%d", value)
	countwrites(value,14,7)
	l.Unlock()

	delay()
	l.Lock()
	value = 22
	s22.Put(348,value)
	log("p14 w:%d", value)
	countwrites(value,14,22)
	l.Unlock()

	delay()
	l.Lock()
	value = 30
	s30.Put(467,value)
	log("p14 w:%d", value)
	countwrites(value,14,30)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s13.QueryP(206,&value)
	log("p14 r:%d", value)
	countreads(value,14,13)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s20.QueryP(312,&value)
	log("p14 r:%d", value)
	countreads(value,14,20)
	l.Unlock()

	delay()
	l.Lock()
	value = 5
	s5.Put(78,value)
	log("p14 w:%d", value)
	countwrites(value,14,5)
	l.Unlock()

	delay()
	l.Lock()
	value = 14
	s14.Put(212,value)
	log("p14 w:%d", value)
	countwrites(value,14,14)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s17.QueryP(265,&value)
	log("p14 r:%d", value)
	countreads(value,14,17)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s2.QueryP(25,&value)
	log("p14 r:%d", value)
	countreads(value,14,2)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s13.QueryP(196,&value)
	log("p14 r:%d", value)
	countreads(value,14,13)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s16.QueryP(249,&value)
	log("p14 r:%d", value)
	countreads(value,14,16)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s27.QueryP(421,&value)
	log("p14 r:%d", value)
	countreads(value,14,27)
	l.Unlock()
}


func p15() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 25
	s25.Put(394,value)
	log("p15 w:%d", value)
	countwrites(value,15,25)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s2.QueryP(24,&value)
	log("p15 r:%d", value)
	countreads(value,15,2)
	l.Unlock()

	delay()
	l.Lock()
	value = 30
	s30.Put(477,value)
	log("p15 w:%d", value)
	countwrites(value,15,30)
	l.Unlock()

	delay()
	l.Lock()
	value = 29
	s29.Put(462,value)
	log("p15 w:%d", value)
	countwrites(value,15,29)
	l.Unlock()

	delay()
	l.Lock()
	value = 5
	s5.Put(66,value)
	log("p15 w:%d", value)
	countwrites(value,15,5)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s4.QueryP(60,&value)
	log("p15 r:%d", value)
	countreads(value,15,4)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s6.QueryP(81,&value)
	log("p15 r:%d", value)
	countreads(value,15,6)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s6.QueryP(92,&value)
	log("p15 r:%d", value)
	countreads(value,15,6)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s20.QueryP(316,&value)
	log("p15 r:%d", value)
	countreads(value,15,20)
	l.Unlock()

	delay()
	l.Lock()
	value = 23
	s23.Put(357,value)
	log("p15 w:%d", value)
	countwrites(value,15,23)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s25.QueryP(396,&value)
	log("p15 r:%d", value)
	countreads(value,15,25)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s10.QueryP(148,&value)
	log("p15 r:%d", value)
	countreads(value,15,10)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s29.QueryP(464,&value)
	log("p15 r:%d", value)
	countreads(value,15,29)
	l.Unlock()

	delay()
	l.Lock()
	value = 3
	s3.Put(40,value)
	log("p15 w:%d", value)
	countwrites(value,15,3)
	l.Unlock()

	delay()
	l.Lock()
	value = 26
	s26.Put(407,value)
	log("p15 w:%d", value)
	countwrites(value,15,26)
	l.Unlock()

	delay()
	l.Lock()
	value = 1
	s1.Put(1,value)
	log("p15 w:%d", value)
	countwrites(value,15,1)
	l.Unlock()
}


func p16() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	s11.QueryP(170,&value)
	log("p16 r:%d", value)
	countreads(value,16,11)
	l.Unlock()

	delay()
	l.Lock()
	value = 24
	s24.Put(371,value)
	log("p16 w:%d", value)
	countwrites(value,16,24)
	l.Unlock()

	delay()
	l.Lock()
	value = 17
	s17.Put(266,value)
	log("p16 w:%d", value)
	countwrites(value,16,17)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s22.QueryP(351,&value)
	log("p16 r:%d", value)
	countreads(value,16,22)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s28.QueryP(438,&value)
	log("p16 r:%d", value)
	countreads(value,16,28)
	l.Unlock()

	delay()
	l.Lock()
	value = 27
	s27.Put(421,value)
	log("p16 w:%d", value)
	countwrites(value,16,27)
	l.Unlock()

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
	s18.QueryP(280,&value)
	log("p16 r:%d", value)
	countreads(value,16,18)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s3.QueryP(40,&value)
	log("p16 r:%d", value)
	countreads(value,16,3)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s3.QueryP(40,&value)
	log("p16 r:%d", value)
	countreads(value,16,3)
	l.Unlock()

	delay()
	l.Lock()
	value = 22
	s22.Put(342,value)
	log("p16 w:%d", value)
	countwrites(value,16,22)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s5.QueryP(70,&value)
	log("p16 r:%d", value)
	countreads(value,16,5)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s6.QueryP(93,&value)
	log("p16 r:%d", value)
	countreads(value,16,6)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s20.QueryP(317,&value)
	log("p16 r:%d", value)
	countreads(value,16,20)
	l.Unlock()

	delay()
	l.Lock()
	value = 30
	s30.Put(478,value)
	log("p16 w:%d", value)
	countwrites(value,16,30)
	l.Unlock()

	delay()
	l.Lock()
	value = 11
	s11.Put(167,value)
	log("p16 w:%d", value)
	countwrites(value,16,11)
	l.Unlock()
}


func p17() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	s25.QueryP(393,&value)
	log("p17 r:%d", value)
	countreads(value,17,25)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s21.QueryP(334,&value)
	log("p17 r:%d", value)
	countreads(value,17,21)
	l.Unlock()

	delay()
	l.Lock()
	value = 19
	s19.Put(300,value)
	log("p17 w:%d", value)
	countwrites(value,17,19)
	l.Unlock()

	delay()
	l.Lock()
	value = 18
	s18.Put(277,value)
	log("p17 w:%d", value)
	countwrites(value,17,18)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s1.QueryP(16,&value)
	log("p17 r:%d", value)
	countreads(value,17,1)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s9.QueryP(144,&value)
	log("p17 r:%d", value)
	countreads(value,17,9)
	l.Unlock()

	delay()
	l.Lock()
	value = 11
	s11.Put(169,value)
	log("p17 w:%d", value)
	countwrites(value,17,11)
	l.Unlock()

	delay()
	l.Lock()
	value = 16
	s16.Put(252,value)
	log("p17 w:%d", value)
	countwrites(value,17,16)
	l.Unlock()

	delay()
	l.Lock()
	value = 12
	s12.Put(190,value)
	log("p17 w:%d", value)
	countwrites(value,17,12)
	l.Unlock()

	delay()
	l.Lock()
	value = 1
	s1.Put(15,value)
	log("p17 w:%d", value)
	countwrites(value,17,1)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s5.QueryP(78,&value)
	log("p17 r:%d", value)
	countreads(value,17,5)
	l.Unlock()

	delay()
	l.Lock()
	value = 26
	s26.Put(409,value)
	log("p17 w:%d", value)
	countwrites(value,17,26)
	l.Unlock()

	delay()
	l.Lock()
	value = 20
	s20.Put(313,value)
	log("p17 w:%d", value)
	countwrites(value,17,20)
	l.Unlock()

	delay()
	l.Lock()
	value = 13
	s13.Put(194,value)
	log("p17 w:%d", value)
	countwrites(value,17,13)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s11.QueryP(171,&value)
	log("p17 r:%d", value)
	countreads(value,17,11)
	l.Unlock()

	delay()
	l.Lock()
	value = 27
	s27.Put(424,value)
	log("p17 w:%d", value)
	countwrites(value,17,27)
	l.Unlock()
}


func p18() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	s5.QueryP(79,&value)
	log("p18 r:%d", value)
	countreads(value,18,5)
	l.Unlock()

	delay()
	l.Lock()
	value = 18
	s18.Put(273,value)
	log("p18 w:%d", value)
	countwrites(value,18,18)
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
	s17.QueryP(259,&value)
	log("p18 r:%d", value)
	countreads(value,18,17)
	l.Unlock()

	delay()
	l.Lock()
	value = 31
	s31.Put(496,value)
	log("p18 w:%d", value)
	countwrites(value,18,31)
	l.Unlock()

	delay()
	l.Lock()
	value = 19
	s19.Put(304,value)
	log("p18 w:%d", value)
	countwrites(value,18,19)
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
	value = -1
	s27.QueryP(424,&value)
	log("p18 r:%d", value)
	countreads(value,18,27)
	l.Unlock()

	delay()
	l.Lock()
	value = 20
	s20.Put(316,value)
	log("p18 w:%d", value)
	countwrites(value,18,20)
	l.Unlock()

	delay()
	l.Lock()
	value = 12
	s12.Put(181,value)
	log("p18 w:%d", value)
	countwrites(value,18,12)
	l.Unlock()

	delay()
	l.Lock()
	value = 29
	s29.Put(461,value)
	log("p18 w:%d", value)
	countwrites(value,18,29)
	l.Unlock()

	delay()
	l.Lock()
	value = 27
	s27.Put(430,value)
	log("p18 w:%d", value)
	countwrites(value,18,27)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s30.QueryP(468,&value)
	log("p18 r:%d", value)
	countreads(value,18,30)
	l.Unlock()

	delay()
	l.Lock()
	value = 29
	s29.Put(461,value)
	log("p18 w:%d", value)
	countwrites(value,18,29)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s2.QueryP(32,&value)
	log("p18 r:%d", value)
	countreads(value,18,2)
	l.Unlock()

	delay()
	l.Lock()
	value = 12
	s12.Put(192,value)
	log("p18 w:%d", value)
	countwrites(value,18,12)
	l.Unlock()
}


func p19() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	s30.QueryP(477,&value)
	log("p19 r:%d", value)
	countreads(value,19,30)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s30.QueryP(472,&value)
	log("p19 r:%d", value)
	countreads(value,19,30)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s16.QueryP(255,&value)
	log("p19 r:%d", value)
	countreads(value,19,16)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s29.QueryP(450,&value)
	log("p19 r:%d", value)
	countreads(value,19,29)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s21.QueryP(330,&value)
	log("p19 r:%d", value)
	countreads(value,19,21)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s21.QueryP(324,&value)
	log("p19 r:%d", value)
	countreads(value,19,21)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s12.QueryP(180,&value)
	log("p19 r:%d", value)
	countreads(value,19,12)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s27.QueryP(427,&value)
	log("p19 r:%d", value)
	countreads(value,19,27)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s23.QueryP(367,&value)
	log("p19 r:%d", value)
	countreads(value,19,23)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s28.QueryP(433,&value)
	log("p19 r:%d", value)
	countreads(value,19,28)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s15.QueryP(234,&value)
	log("p19 r:%d", value)
	countreads(value,19,15)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s25.QueryP(391,&value)
	log("p19 r:%d", value)
	countreads(value,19,25)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s12.QueryP(188,&value)
	log("p19 r:%d", value)
	countreads(value,19,12)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s30.QueryP(469,&value)
	log("p19 r:%d", value)
	countreads(value,19,30)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s24.QueryP(383,&value)
	log("p19 r:%d", value)
	countreads(value,19,24)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s3.QueryP(34,&value)
	log("p19 r:%d", value)
	countreads(value,19,3)
	l.Unlock()
}


func p20() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 17
	s17.Put(261,value)
	log("p20 w:%d", value)
	countwrites(value,20,17)
	l.Unlock()

	delay()
	l.Lock()
	value = 2
	s2.Put(23,value)
	log("p20 w:%d", value)
	countwrites(value,20,2)
	l.Unlock()

	delay()
	l.Lock()
	value = 6
	s6.Put(87,value)
	log("p20 w:%d", value)
	countwrites(value,20,6)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s13.QueryP(203,&value)
	log("p20 r:%d", value)
	countreads(value,20,13)
	l.Unlock()

	delay()
	l.Lock()
	value = 2
	s2.Put(29,value)
	log("p20 w:%d", value)
	countwrites(value,20,2)
	l.Unlock()

	delay()
	l.Lock()
	value = 30
	s30.Put(468,value)
	log("p20 w:%d", value)
	countwrites(value,20,30)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s19.QueryP(290,&value)
	log("p20 r:%d", value)
	countreads(value,20,19)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s16.QueryP(248,&value)
	log("p20 r:%d", value)
	countreads(value,20,16)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s9.QueryP(137,&value)
	log("p20 r:%d", value)
	countreads(value,20,9)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s24.QueryP(372,&value)
	log("p20 r:%d", value)
	countreads(value,20,24)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s21.QueryP(323,&value)
	log("p20 r:%d", value)
	countreads(value,20,21)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s6.QueryP(83,&value)
	log("p20 r:%d", value)
	countreads(value,20,6)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s7.QueryP(108,&value)
	log("p20 r:%d", value)
	countreads(value,20,7)
	l.Unlock()

	delay()
	l.Lock()
	value = 22
	s22.Put(350,value)
	log("p20 w:%d", value)
	countwrites(value,20,22)
	l.Unlock()

	delay()
	l.Lock()
	value = 2
	s2.Put(25,value)
	log("p20 w:%d", value)
	countwrites(value,20,2)
	l.Unlock()

	delay()
	l.Lock()
	value = 21
	s21.Put(332,value)
	log("p20 w:%d", value)
	countwrites(value,20,21)
	l.Unlock()
}


func p21() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 30
	s30.Put(474,value)
	log("p21 w:%d", value)
	countwrites(value,21,30)
	l.Unlock()

	delay()
	l.Lock()
	value = 19
	s19.Put(304,value)
	log("p21 w:%d", value)
	countwrites(value,21,19)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s1.QueryP(13,&value)
	log("p21 r:%d", value)
	countreads(value,21,1)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s18.QueryP(282,&value)
	log("p21 r:%d", value)
	countreads(value,21,18)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s3.QueryP(39,&value)
	log("p21 r:%d", value)
	countreads(value,21,3)
	l.Unlock()

	delay()
	l.Lock()
	value = 16
	s16.Put(243,value)
	log("p21 w:%d", value)
	countwrites(value,21,16)
	l.Unlock()

	delay()
	l.Lock()
	value = 16
	s16.Put(248,value)
	log("p21 w:%d", value)
	countwrites(value,21,16)
	l.Unlock()

	delay()
	l.Lock()
	value = 27
	s27.Put(419,value)
	log("p21 w:%d", value)
	countwrites(value,21,27)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s12.QueryP(184,&value)
	log("p21 r:%d", value)
	countreads(value,21,12)
	l.Unlock()

	delay()
	l.Lock()
	value = 25
	s25.Put(399,value)
	log("p21 w:%d", value)
	countwrites(value,21,25)
	l.Unlock()

	delay()
	l.Lock()
	value = 12
	s12.Put(178,value)
	log("p21 w:%d", value)
	countwrites(value,21,12)
	l.Unlock()

	delay()
	l.Lock()
	value = 30
	s30.Put(468,value)
	log("p21 w:%d", value)
	countwrites(value,21,30)
	l.Unlock()

	delay()
	l.Lock()
	value = 27
	s27.Put(419,value)
	log("p21 w:%d", value)
	countwrites(value,21,27)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s2.QueryP(20,&value)
	log("p21 r:%d", value)
	countreads(value,21,2)
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
	s7.QueryP(99,&value)
	log("p21 r:%d", value)
	countreads(value,21,7)
	l.Unlock()
}


func p22() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 20
	s20.Put(312,value)
	log("p22 w:%d", value)
	countwrites(value,22,20)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s3.QueryP(35,&value)
	log("p22 r:%d", value)
	countreads(value,22,3)
	l.Unlock()

	delay()
	l.Lock()
	value = 19
	s19.Put(291,value)
	log("p22 w:%d", value)
	countwrites(value,22,19)
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
	value = -1
	s2.QueryP(28,&value)
	log("p22 r:%d", value)
	countreads(value,22,2)
	l.Unlock()

	delay()
	l.Lock()
	value = 30
	s30.Put(475,value)
	log("p22 w:%d", value)
	countwrites(value,22,30)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s8.QueryP(122,&value)
	log("p22 r:%d", value)
	countreads(value,22,8)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s13.QueryP(201,&value)
	log("p22 r:%d", value)
	countreads(value,22,13)
	l.Unlock()

	delay()
	l.Lock()
	value = 29
	s29.Put(458,value)
	log("p22 w:%d", value)
	countwrites(value,22,29)
	l.Unlock()

	delay()
	l.Lock()
	value = 21
	s21.Put(334,value)
	log("p22 w:%d", value)
	countwrites(value,22,21)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s24.QueryP(373,&value)
	log("p22 r:%d", value)
	countreads(value,22,24)
	l.Unlock()

	delay()
	l.Lock()
	value = 26
	s26.Put(401,value)
	log("p22 w:%d", value)
	countwrites(value,22,26)
	l.Unlock()

	delay()
	l.Lock()
	value = 27
	s27.Put(432,value)
	log("p22 w:%d", value)
	countwrites(value,22,27)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s5.QueryP(80,&value)
	log("p22 r:%d", value)
	countreads(value,22,5)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s6.QueryP(85,&value)
	log("p22 r:%d", value)
	countreads(value,22,6)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s1.QueryP(1,&value)
	log("p22 r:%d", value)
	countreads(value,22,1)
	l.Unlock()
}


func p23() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	s18.QueryP(277,&value)
	log("p23 r:%d", value)
	countreads(value,23,18)
	l.Unlock()

	delay()
	l.Lock()
	value = 13
	s13.Put(196,value)
	log("p23 w:%d", value)
	countwrites(value,23,13)
	l.Unlock()

	delay()
	l.Lock()
	value = 24
	s24.Put(371,value)
	log("p23 w:%d", value)
	countwrites(value,23,24)
	l.Unlock()

	delay()
	l.Lock()
	value = 32
	s32.Put(500,value)
	log("p23 w:%d", value)
	countwrites(value,23,32)
	l.Unlock()

	delay()
	l.Lock()
	value = 28
	s28.Put(440,value)
	log("p23 w:%d", value)
	countwrites(value,23,28)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s15.QueryP(240,&value)
	log("p23 r:%d", value)
	countreads(value,23,15)
	l.Unlock()

	delay()
	l.Lock()
	value = 1
	s1.Put(10,value)
	log("p23 w:%d", value)
	countwrites(value,23,1)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s1.QueryP(16,&value)
	log("p23 r:%d", value)
	countreads(value,23,1)
	l.Unlock()

	delay()
	l.Lock()
	value = 24
	s24.Put(373,value)
	log("p23 w:%d", value)
	countwrites(value,23,24)
	l.Unlock()

	delay()
	l.Lock()
	value = 21
	s21.Put(323,value)
	log("p23 w:%d", value)
	countwrites(value,23,21)
	l.Unlock()

	delay()
	l.Lock()
	value = 15
	s15.Put(237,value)
	log("p23 w:%d", value)
	countwrites(value,23,15)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s12.QueryP(188,&value)
	log("p23 r:%d", value)
	countreads(value,23,12)
	l.Unlock()

	delay()
	l.Lock()
	value = 16
	s16.Put(248,value)
	log("p23 w:%d", value)
	countwrites(value,23,16)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s11.QueryP(164,&value)
	log("p23 r:%d", value)
	countreads(value,23,11)
	l.Unlock()

	delay()
	l.Lock()
	value = 24
	s24.Put(379,value)
	log("p23 w:%d", value)
	countwrites(value,23,24)
	l.Unlock()

	delay()
	l.Lock()
	value = 21
	s21.Put(334,value)
	log("p23 w:%d", value)
	countwrites(value,23,21)
	l.Unlock()
}


func p24() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	s24.QueryP(370,&value)
	log("p24 r:%d", value)
	countreads(value,24,24)
	l.Unlock()

	delay()
	l.Lock()
	value = 8
	s8.Put(123,value)
	log("p24 w:%d", value)
	countwrites(value,24,8)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s28.QueryP(439,&value)
	log("p24 r:%d", value)
	countreads(value,24,28)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s21.QueryP(321,&value)
	log("p24 r:%d", value)
	countreads(value,24,21)
	l.Unlock()

	delay()
	l.Lock()
	value = 17
	s17.Put(263,value)
	log("p24 w:%d", value)
	countwrites(value,24,17)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s4.QueryP(64,&value)
	log("p24 r:%d", value)
	countreads(value,24,4)
	l.Unlock()

	delay()
	l.Lock()
	value = 3
	s3.Put(40,value)
	log("p24 w:%d", value)
	countwrites(value,24,3)
	l.Unlock()

	delay()
	l.Lock()
	value = 26
	s26.Put(416,value)
	log("p24 w:%d", value)
	countwrites(value,24,26)
	l.Unlock()

	delay()
	l.Lock()
	value = 11
	s11.Put(163,value)
	log("p24 w:%d", value)
	countwrites(value,24,11)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s4.QueryP(61,&value)
	log("p24 r:%d", value)
	countreads(value,24,4)
	l.Unlock()

	delay()
	l.Lock()
	value = 4
	s4.Put(61,value)
	log("p24 w:%d", value)
	countwrites(value,24,4)
	l.Unlock()

	delay()
	l.Lock()
	value = 15
	s15.Put(237,value)
	log("p24 w:%d", value)
	countwrites(value,24,15)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s6.QueryP(87,&value)
	log("p24 r:%d", value)
	countreads(value,24,6)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s11.QueryP(166,&value)
	log("p24 r:%d", value)
	countreads(value,24,11)
	l.Unlock()

	delay()
	l.Lock()
	value = 7
	s7.Put(102,value)
	log("p24 w:%d", value)
	countwrites(value,24,7)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s10.QueryP(150,&value)
	log("p24 r:%d", value)
	countreads(value,24,10)
	l.Unlock()
}


func p25() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 24
	s24.Put(382,value)
	log("p25 w:%d", value)
	countwrites(value,25,24)
	l.Unlock()

	delay()
	l.Lock()
	value = 20
	s20.Put(317,value)
	log("p25 w:%d", value)
	countwrites(value,25,20)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s25.QueryP(397,&value)
	log("p25 r:%d", value)
	countreads(value,25,25)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s4.QueryP(49,&value)
	log("p25 r:%d", value)
	countreads(value,25,4)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s16.QueryP(246,&value)
	log("p25 r:%d", value)
	countreads(value,25,16)
	l.Unlock()

	delay()
	l.Lock()
	value = 24
	s24.Put(369,value)
	log("p25 w:%d", value)
	countwrites(value,25,24)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s4.QueryP(50,&value)
	log("p25 r:%d", value)
	countreads(value,25,4)
	l.Unlock()

	delay()
	l.Lock()
	value = 12
	s12.Put(191,value)
	log("p25 w:%d", value)
	countwrites(value,25,12)
	l.Unlock()

	delay()
	l.Lock()
	value = 21
	s21.Put(332,value)
	log("p25 w:%d", value)
	countwrites(value,25,21)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s9.QueryP(144,&value)
	log("p25 r:%d", value)
	countreads(value,25,9)
	l.Unlock()

	delay()
	l.Lock()
	value = 1
	s1.Put(3,value)
	log("p25 w:%d", value)
	countwrites(value,25,1)
	l.Unlock()

	delay()
	l.Lock()
	value = 31
	s31.Put(485,value)
	log("p25 w:%d", value)
	countwrites(value,25,31)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s25.QueryP(396,&value)
	log("p25 r:%d", value)
	countreads(value,25,25)
	l.Unlock()

	delay()
	l.Lock()
	value = 6
	s6.Put(83,value)
	log("p25 w:%d", value)
	countwrites(value,25,6)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s8.QueryP(114,&value)
	log("p25 r:%d", value)
	countreads(value,25,8)
	l.Unlock()

	delay()
	l.Lock()
	value = 1
	s1.Put(14,value)
	log("p25 w:%d", value)
	countwrites(value,25,1)
	l.Unlock()
}


func p26() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	s9.QueryP(130,&value)
	log("p26 r:%d", value)
	countreads(value,26,9)
	l.Unlock()

	delay()
	l.Lock()
	value = 25
	s25.Put(385,value)
	log("p26 w:%d", value)
	countwrites(value,26,25)
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
	s4.QueryP(56,&value)
	log("p26 r:%d", value)
	countreads(value,26,4)
	l.Unlock()

	delay()
	l.Lock()
	value = 27
	s27.Put(419,value)
	log("p26 w:%d", value)
	countwrites(value,26,27)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s14.QueryP(214,&value)
	log("p26 r:%d", value)
	countreads(value,26,14)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s10.QueryP(157,&value)
	log("p26 r:%d", value)
	countreads(value,26,10)
	l.Unlock()

	delay()
	l.Lock()
	value = 30
	s30.Put(475,value)
	log("p26 w:%d", value)
	countwrites(value,26,30)
	l.Unlock()

	delay()
	l.Lock()
	value = 14
	s14.Put(211,value)
	log("p26 w:%d", value)
	countwrites(value,26,14)
	l.Unlock()

	delay()
	l.Lock()
	value = 25
	s25.Put(386,value)
	log("p26 w:%d", value)
	countwrites(value,26,25)
	l.Unlock()

	delay()
	l.Lock()
	value = 1
	s1.Put(9,value)
	log("p26 w:%d", value)
	countwrites(value,26,1)
	l.Unlock()

	delay()
	l.Lock()
	value = 6
	s6.Put(93,value)
	log("p26 w:%d", value)
	countwrites(value,26,6)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s6.QueryP(96,&value)
	log("p26 r:%d", value)
	countreads(value,26,6)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s31.QueryP(494,&value)
	log("p26 r:%d", value)
	countreads(value,26,31)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s20.QueryP(317,&value)
	log("p26 r:%d", value)
	countreads(value,26,20)
	l.Unlock()

	delay()
	l.Lock()
	value = 22
	s22.Put(352,value)
	log("p26 w:%d", value)
	countwrites(value,26,22)
	l.Unlock()
}


func p27() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 5
	s5.Put(65,value)
	log("p27 w:%d", value)
	countwrites(value,27,5)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s11.QueryP(163,&value)
	log("p27 r:%d", value)
	countreads(value,27,11)
	l.Unlock()

	delay()
	l.Lock()
	value = 12
	s12.Put(178,value)
	log("p27 w:%d", value)
	countwrites(value,27,12)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s17.QueryP(270,&value)
	log("p27 r:%d", value)
	countreads(value,27,17)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s23.QueryP(368,&value)
	log("p27 r:%d", value)
	countreads(value,27,23)
	l.Unlock()

	delay()
	l.Lock()
	value = 4
	s4.Put(58,value)
	log("p27 w:%d", value)
	countwrites(value,27,4)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s7.QueryP(111,&value)
	log("p27 r:%d", value)
	countreads(value,27,7)
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
	value = 10
	s10.Put(154,value)
	log("p27 w:%d", value)
	countwrites(value,27,10)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s13.QueryP(197,&value)
	log("p27 r:%d", value)
	countreads(value,27,13)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s19.QueryP(297,&value)
	log("p27 r:%d", value)
	countreads(value,27,19)
	l.Unlock()

	delay()
	l.Lock()
	value = 28
	s28.Put(442,value)
	log("p27 w:%d", value)
	countwrites(value,27,28)
	l.Unlock()

	delay()
	l.Lock()
	value = 21
	s21.Put(333,value)
	log("p27 w:%d", value)
	countwrites(value,27,21)
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
	value = 22
	s22.Put(339,value)
	log("p27 w:%d", value)
	countwrites(value,27,22)
	l.Unlock()

	delay()
	l.Lock()
	value = 1
	s1.Put(13,value)
	log("p27 w:%d", value)
	countwrites(value,27,1)
	l.Unlock()
}


func p28() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 24
	s24.Put(375,value)
	log("p28 w:%d", value)
	countwrites(value,28,24)
	l.Unlock()

	delay()
	l.Lock()
	value = 31
	s31.Put(495,value)
	log("p28 w:%d", value)
	countwrites(value,28,31)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s14.QueryP(218,&value)
	log("p28 r:%d", value)
	countreads(value,28,14)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s19.QueryP(300,&value)
	log("p28 r:%d", value)
	countreads(value,28,19)
	l.Unlock()

	delay()
	l.Lock()
	value = 8
	s8.Put(119,value)
	log("p28 w:%d", value)
	countwrites(value,28,8)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s5.QueryP(65,&value)
	log("p28 r:%d", value)
	countreads(value,28,5)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s11.QueryP(167,&value)
	log("p28 r:%d", value)
	countreads(value,28,11)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s15.QueryP(239,&value)
	log("p28 r:%d", value)
	countreads(value,28,15)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s7.QueryP(106,&value)
	log("p28 r:%d", value)
	countreads(value,28,7)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s29.QueryP(455,&value)
	log("p28 r:%d", value)
	countreads(value,28,29)
	l.Unlock()

	delay()
	l.Lock()
	value = 18
	s18.Put(283,value)
	log("p28 w:%d", value)
	countwrites(value,28,18)
	l.Unlock()

	delay()
	l.Lock()
	value = 8
	s8.Put(125,value)
	log("p28 w:%d", value)
	countwrites(value,28,8)
	l.Unlock()

	delay()
	l.Lock()
	value = 3
	s3.Put(33,value)
	log("p28 w:%d", value)
	countwrites(value,28,3)
	l.Unlock()

	delay()
	l.Lock()
	value = 20
	s20.Put(307,value)
	log("p28 w:%d", value)
	countwrites(value,28,20)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s13.QueryP(195,&value)
	log("p28 r:%d", value)
	countreads(value,28,13)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s28.QueryP(445,&value)
	log("p28 r:%d", value)
	countreads(value,28,28)
	l.Unlock()
}


func p29() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	s27.QueryP(428,&value)
	log("p29 r:%d", value)
	countreads(value,29,27)
	l.Unlock()

	delay()
	l.Lock()
	value = 27
	s27.Put(431,value)
	log("p29 w:%d", value)
	countwrites(value,29,27)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s30.QueryP(478,&value)
	log("p29 r:%d", value)
	countreads(value,29,30)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s2.QueryP(27,&value)
	log("p29 r:%d", value)
	countreads(value,29,2)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s27.QueryP(423,&value)
	log("p29 r:%d", value)
	countreads(value,29,27)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s3.QueryP(42,&value)
	log("p29 r:%d", value)
	countreads(value,29,3)
	l.Unlock()

	delay()
	l.Lock()
	value = 8
	s8.Put(124,value)
	log("p29 w:%d", value)
	countwrites(value,29,8)
	l.Unlock()

	delay()
	l.Lock()
	value = 11
	s11.Put(176,value)
	log("p29 w:%d", value)
	countwrites(value,29,11)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s16.QueryP(249,&value)
	log("p29 r:%d", value)
	countreads(value,29,16)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s2.QueryP(17,&value)
	log("p29 r:%d", value)
	countreads(value,29,2)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s26.QueryP(412,&value)
	log("p29 r:%d", value)
	countreads(value,29,26)
	l.Unlock()

	delay()
	l.Lock()
	value = 16
	s16.Put(245,value)
	log("p29 w:%d", value)
	countwrites(value,29,16)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s21.QueryP(334,&value)
	log("p29 r:%d", value)
	countreads(value,29,21)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s9.QueryP(135,&value)
	log("p29 r:%d", value)
	countreads(value,29,9)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s15.QueryP(233,&value)
	log("p29 r:%d", value)
	countreads(value,29,15)
	l.Unlock()

	delay()
	l.Lock()
	value = 29
	s29.Put(461,value)
	log("p29 w:%d", value)
	countwrites(value,29,29)
	l.Unlock()
}


func p30() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	s29.QueryP(453,&value)
	log("p30 r:%d", value)
	countreads(value,30,29)
	l.Unlock()

	delay()
	l.Lock()
	value = 20
	s20.Put(307,value)
	log("p30 w:%d", value)
	countwrites(value,30,20)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s14.QueryP(217,&value)
	log("p30 r:%d", value)
	countreads(value,30,14)
	l.Unlock()

	delay()
	l.Lock()
	value = 13
	s13.Put(205,value)
	log("p30 w:%d", value)
	countwrites(value,30,13)
	l.Unlock()

	delay()
	l.Lock()
	value = 25
	s25.Put(395,value)
	log("p30 w:%d", value)
	countwrites(value,30,25)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s27.QueryP(418,&value)
	log("p30 r:%d", value)
	countreads(value,30,27)
	l.Unlock()

	delay()
	l.Lock()
	value = 22
	s22.Put(352,value)
	log("p30 w:%d", value)
	countwrites(value,30,22)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s23.QueryP(363,&value)
	log("p30 r:%d", value)
	countreads(value,30,23)
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
	value = 18
	s18.Put(281,value)
	log("p30 w:%d", value)
	countwrites(value,30,18)
	l.Unlock()

	delay()
	l.Lock()
	value = 27
	s27.Put(425,value)
	log("p30 w:%d", value)
	countwrites(value,30,27)
	l.Unlock()

	delay()
	l.Lock()
	value = 20
	s20.Put(318,value)
	log("p30 w:%d", value)
	countwrites(value,30,20)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s25.QueryP(396,&value)
	log("p30 r:%d", value)
	countreads(value,30,25)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s6.QueryP(86,&value)
	log("p30 r:%d", value)
	countreads(value,30,6)
	l.Unlock()

	delay()
	l.Lock()
	value = 29
	s29.Put(460,value)
	log("p30 w:%d", value)
	countwrites(value,30,29)
	l.Unlock()

	delay()
	l.Lock()
	value = 3
	s3.Put(34,value)
	log("p30 w:%d", value)
	countwrites(value,30,3)
	l.Unlock()
}


func p31() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 4
	s4.Put(58,value)
	log("p31 w:%d", value)
	countwrites(value,31,4)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s11.QueryP(164,&value)
	log("p31 r:%d", value)
	countreads(value,31,11)
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
	s6.QueryP(91,&value)
	log("p31 r:%d", value)
	countreads(value,31,6)
	l.Unlock()

	delay()
	l.Lock()
	value = 11
	s11.Put(176,value)
	log("p31 w:%d", value)
	countwrites(value,31,11)
	l.Unlock()

	delay()
	l.Lock()
	value = 21
	s21.Put(333,value)
	log("p31 w:%d", value)
	countwrites(value,31,21)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s22.QueryP(342,&value)
	log("p31 r:%d", value)
	countreads(value,31,22)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s7.QueryP(106,&value)
	log("p31 r:%d", value)
	countreads(value,31,7)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s15.QueryP(233,&value)
	log("p31 r:%d", value)
	countreads(value,31,15)
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
	value = 18
	s18.Put(283,value)
	log("p31 w:%d", value)
	countwrites(value,31,18)
	l.Unlock()

	delay()
	l.Lock()
	value = 1
	s1.Put(6,value)
	log("p31 w:%d", value)
	countwrites(value,31,1)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s29.QueryP(461,&value)
	log("p31 r:%d", value)
	countreads(value,31,29)
	l.Unlock()

	delay()
	l.Lock()
	value = 14
	s14.Put(218,value)
	log("p31 w:%d", value)
	countwrites(value,31,14)
	l.Unlock()

	delay()
	l.Lock()
	value = 29
	s29.Put(458,value)
	log("p31 w:%d", value)
	countwrites(value,31,29)
	l.Unlock()

	delay()
	l.Lock()
	value = 15
	s15.Put(239,value)
	log("p31 w:%d", value)
	countwrites(value,31,15)
	l.Unlock()
}


func p32() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	s24.QueryP(372,&value)
	log("p32 r:%d", value)
	countreads(value,32,24)
	l.Unlock()

	delay()
	l.Lock()
	value = 10
	s10.Put(153,value)
	log("p32 w:%d", value)
	countwrites(value,32,10)
	l.Unlock()

	delay()
	l.Lock()
	value = 30
	s30.Put(465,value)
	log("p32 w:%d", value)
	countwrites(value,32,30)
	l.Unlock()

	delay()
	l.Lock()
	value = 32
	s32.Put(503,value)
	log("p32 w:%d", value)
	countwrites(value,32,32)
	l.Unlock()

	delay()
	l.Lock()
	value = 20
	s20.Put(316,value)
	log("p32 w:%d", value)
	countwrites(value,32,20)
	l.Unlock()

	delay()
	l.Lock()
	value = 31
	s31.Put(496,value)
	log("p32 w:%d", value)
	countwrites(value,32,31)
	l.Unlock()

	delay()
	l.Lock()
	value = 4
	s4.Put(58,value)
	log("p32 w:%d", value)
	countwrites(value,32,4)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s24.QueryP(378,&value)
	log("p32 r:%d", value)
	countreads(value,32,24)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s20.QueryP(309,&value)
	log("p32 r:%d", value)
	countreads(value,32,20)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s25.QueryP(396,&value)
	log("p32 r:%d", value)
	countreads(value,32,25)
	l.Unlock()

	delay()
	l.Lock()
	value = 17
	s17.Put(270,value)
	log("p32 w:%d", value)
	countwrites(value,32,17)
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
	value = 7
	s7.Put(101,value)
	log("p32 w:%d", value)
	countwrites(value,32,7)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s31.QueryP(491,&value)
	log("p32 r:%d", value)
	countreads(value,32,31)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s3.QueryP(47,&value)
	log("p32 r:%d", value)
	countreads(value,32,3)
	l.Unlock()

	delay()
	l.Lock()
	value = 9
	s9.Put(131,value)
	log("p32 w:%d", value)
	countwrites(value,32,9)
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
