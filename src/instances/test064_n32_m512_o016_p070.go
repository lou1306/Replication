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
	value = 32
	s32.Put(506,value)
	log("p1 w:%d", value)
	countwrites(value,1,32)
	l.Unlock()

	delay()
	l.Lock()
	value = 2
	s2.Put(25,value)
	log("p1 w:%d", value)
	countwrites(value,1,2)
	l.Unlock()

	delay()
	l.Lock()
	value = 16
	s16.Put(248,value)
	log("p1 w:%d", value)
	countwrites(value,1,16)
	l.Unlock()

	delay()
	l.Lock()
	value = 27
	s27.Put(426,value)
	log("p1 w:%d", value)
	countwrites(value,1,27)
	l.Unlock()

	delay()
	l.Lock()
	value = 3
	s3.Put(34,value)
	log("p1 w:%d", value)
	countwrites(value,1,3)
	l.Unlock()

	delay()
	l.Lock()
	value = 1
	s1.Put(6,value)
	log("p1 w:%d", value)
	countwrites(value,1,1)
	l.Unlock()

	delay()
	l.Lock()
	value = 19
	s19.Put(299,value)
	log("p1 w:%d", value)
	countwrites(value,1,19)
	l.Unlock()

	delay()
	l.Lock()
	value = 24
	s24.Put(382,value)
	log("p1 w:%d", value)
	countwrites(value,1,24)
	l.Unlock()

	delay()
	l.Lock()
	value = 16
	s16.Put(254,value)
	log("p1 w:%d", value)
	countwrites(value,1,16)
	l.Unlock()

	delay()
	l.Lock()
	value = 22
	s22.Put(344,value)
	log("p1 w:%d", value)
	countwrites(value,1,22)
	l.Unlock()

	delay()
	l.Lock()
	value = 4
	s4.Put(58,value)
	log("p1 w:%d", value)
	countwrites(value,1,4)
	l.Unlock()

	delay()
	l.Lock()
	value = 25
	s25.Put(387,value)
	log("p1 w:%d", value)
	countwrites(value,1,25)
	l.Unlock()

	delay()
	l.Lock()
	value = 15
	s15.Put(233,value)
	log("p1 w:%d", value)
	countwrites(value,1,15)
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
	value = 23
	s23.Put(361,value)
	log("p1 w:%d", value)
	countwrites(value,1,23)
	l.Unlock()

	delay()
	l.Lock()
	value = 28
	s28.Put(444,value)
	log("p1 w:%d", value)
	countwrites(value,1,28)
	l.Unlock()
}

func p2() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 20
	s20.Put(311,value)
	log("p2 w:%d", value)
	countwrites(value,2,20)
	l.Unlock()

	delay()
	l.Lock()
	value = 22
	s22.Put(342,value)
	log("p2 w:%d", value)
	countwrites(value,2,22)
	l.Unlock()

	delay()
	l.Lock()
	value = 6
	s6.Put(95,value)
	log("p2 w:%d", value)
	countwrites(value,2,6)
	l.Unlock()

	delay()
	l.Lock()
	value = 6
	s6.Put(81,value)
	log("p2 w:%d", value)
	countwrites(value,2,6)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s4.QueryP(59,&value)
	log("p2 r:%d", value)
	countreads(value,2,4)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s23.QueryP(362,&value)
	log("p2 r:%d", value)
	countreads(value,2,23)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s10.QueryP(150,&value)
	log("p2 r:%d", value)
	countreads(value,2,10)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s18.QueryP(275,&value)
	log("p2 r:%d", value)
	countreads(value,2,18)
	l.Unlock()

	delay()
	l.Lock()
	value = 1
	s1.Put(4,value)
	log("p2 w:%d", value)
	countwrites(value,2,1)
	l.Unlock()

	delay()
	l.Lock()
	value = 22
	s22.Put(340,value)
	log("p2 w:%d", value)
	countwrites(value,2,22)
	l.Unlock()

	delay()
	l.Lock()
	value = 31
	s31.Put(481,value)
	log("p2 w:%d", value)
	countwrites(value,2,31)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s28.QueryP(442,&value)
	log("p2 r:%d", value)
	countreads(value,2,28)
	l.Unlock()

	delay()
	l.Lock()
	value = 32
	s32.Put(497,value)
	log("p2 w:%d", value)
	countwrites(value,2,32)
	l.Unlock()

	delay()
	l.Lock()
	value = 32
	s32.Put(510,value)
	log("p2 w:%d", value)
	countwrites(value,2,32)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s21.QueryP(334,&value)
	log("p2 r:%d", value)
	countreads(value,2,21)
	l.Unlock()

	delay()
	l.Lock()
	value = 21
	s21.Put(325,value)
	log("p2 w:%d", value)
	countwrites(value,2,21)
	l.Unlock()
}

func p3() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 9
	s9.Put(130,value)
	log("p3 w:%d", value)
	countwrites(value,3,9)
	l.Unlock()

	delay()
	l.Lock()
	value = 32
	s32.Put(508,value)
	log("p3 w:%d", value)
	countwrites(value,3,32)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s7.QueryP(107,&value)
	log("p3 r:%d", value)
	countreads(value,3,7)
	l.Unlock()

	delay()
	l.Lock()
	value = 24
	s24.Put(369,value)
	log("p3 w:%d", value)
	countwrites(value,3,24)
	l.Unlock()

	delay()
	l.Lock()
	value = 14
	s14.Put(223,value)
	log("p3 w:%d", value)
	countwrites(value,3,14)
	l.Unlock()

	delay()
	l.Lock()
	value = 28
	s28.Put(447,value)
	log("p3 w:%d", value)
	countwrites(value,3,28)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s20.QueryP(311,&value)
	log("p3 r:%d", value)
	countreads(value,3,20)
	l.Unlock()

	delay()
	l.Lock()
	value = 13
	s13.Put(193,value)
	log("p3 w:%d", value)
	countwrites(value,3,13)
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
	value = 23
	s23.Put(354,value)
	log("p3 w:%d", value)
	countwrites(value,3,23)
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
	value = -1
	s9.QueryP(140,&value)
	log("p3 r:%d", value)
	countreads(value,3,9)
	l.Unlock()

	delay()
	l.Lock()
	value = 23
	s23.Put(365,value)
	log("p3 w:%d", value)
	countwrites(value,3,23)
	l.Unlock()

	delay()
	l.Lock()
	value = 32
	s32.Put(503,value)
	log("p3 w:%d", value)
	countwrites(value,3,32)
	l.Unlock()

	delay()
	l.Lock()
	value = 14
	s14.Put(212,value)
	log("p3 w:%d", value)
	countwrites(value,3,14)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s6.QueryP(86,&value)
	log("p3 r:%d", value)
	countreads(value,3,6)
	l.Unlock()
}

func p4() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 14
	s14.Put(221,value)
	log("p4 w:%d", value)
	countwrites(value,4,14)
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
	value = -1
	s23.QueryP(359,&value)
	log("p4 r:%d", value)
	countreads(value,4,23)
	l.Unlock()

	delay()
	l.Lock()
	value = 23
	s23.Put(353,value)
	log("p4 w:%d", value)
	countwrites(value,4,23)
	l.Unlock()

	delay()
	l.Lock()
	value = 30
	s30.Put(474,value)
	log("p4 w:%d", value)
	countwrites(value,4,30)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s32.QueryP(502,&value)
	log("p4 r:%d", value)
	countreads(value,4,32)
	l.Unlock()

	delay()
	l.Lock()
	value = 9
	s9.Put(143,value)
	log("p4 w:%d", value)
	countwrites(value,4,9)
	l.Unlock()

	delay()
	l.Lock()
	value = 24
	s24.Put(382,value)
	log("p4 w:%d", value)
	countwrites(value,4,24)
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
	s19.QueryP(292,&value)
	log("p4 r:%d", value)
	countreads(value,4,19)
	l.Unlock()

	delay()
	l.Lock()
	value = 19
	s19.Put(293,value)
	log("p4 w:%d", value)
	countwrites(value,4,19)
	l.Unlock()

	delay()
	l.Lock()
	value = 20
	s20.Put(312,value)
	log("p4 w:%d", value)
	countwrites(value,4,20)
	l.Unlock()

	delay()
	l.Lock()
	value = 14
	s14.Put(218,value)
	log("p4 w:%d", value)
	countwrites(value,4,14)
	l.Unlock()

	delay()
	l.Lock()
	value = 10
	s10.Put(148,value)
	log("p4 w:%d", value)
	countwrites(value,4,10)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s6.QueryP(96,&value)
	log("p4 r:%d", value)
	countreads(value,4,6)
	l.Unlock()

	delay()
	l.Lock()
	value = 15
	s15.Put(230,value)
	log("p4 w:%d", value)
	countwrites(value,4,15)
	l.Unlock()
}

func p5() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 5
	s5.Put(79,value)
	log("p5 w:%d", value)
	countwrites(value,5,5)
	l.Unlock()

	delay()
	l.Lock()
	value = 29
	s29.Put(453,value)
	log("p5 w:%d", value)
	countwrites(value,5,29)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s13.QueryP(202,&value)
	log("p5 r:%d", value)
	countreads(value,5,13)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s17.QueryP(266,&value)
	log("p5 r:%d", value)
	countreads(value,5,17)
	l.Unlock()

	delay()
	l.Lock()
	value = 32
	s32.Put(507,value)
	log("p5 w:%d", value)
	countwrites(value,5,32)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s23.QueryP(362,&value)
	log("p5 r:%d", value)
	countreads(value,5,23)
	l.Unlock()

	delay()
	l.Lock()
	value = 22
	s22.Put(338,value)
	log("p5 w:%d", value)
	countwrites(value,5,22)
	l.Unlock()

	delay()
	l.Lock()
	value = 16
	s16.Put(247,value)
	log("p5 w:%d", value)
	countwrites(value,5,16)
	l.Unlock()

	delay()
	l.Lock()
	value = 8
	s8.Put(127,value)
	log("p5 w:%d", value)
	countwrites(value,5,8)
	l.Unlock()

	delay()
	l.Lock()
	value = 27
	s27.Put(423,value)
	log("p5 w:%d", value)
	countwrites(value,5,27)
	l.Unlock()

	delay()
	l.Lock()
	value = 16
	s16.Put(251,value)
	log("p5 w:%d", value)
	countwrites(value,5,16)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s13.QueryP(203,&value)
	log("p5 r:%d", value)
	countreads(value,5,13)
	l.Unlock()

	delay()
	l.Lock()
	value = 7
	s7.Put(107,value)
	log("p5 w:%d", value)
	countwrites(value,5,7)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s18.QueryP(277,&value)
	log("p5 r:%d", value)
	countreads(value,5,18)
	l.Unlock()

	delay()
	l.Lock()
	value = 28
	s28.Put(437,value)
	log("p5 w:%d", value)
	countwrites(value,5,28)
	l.Unlock()

	delay()
	l.Lock()
	value = 27
	s27.Put(430,value)
	log("p5 w:%d", value)
	countwrites(value,5,27)
	l.Unlock()
}

func p6() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	s32.QueryP(510,&value)
	log("p6 r:%d", value)
	countreads(value,6,32)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s30.QueryP(469,&value)
	log("p6 r:%d", value)
	countreads(value,6,30)
	l.Unlock()

	delay()
	l.Lock()
	value = 29
	s29.Put(462,value)
	log("p6 w:%d", value)
	countwrites(value,6,29)
	l.Unlock()

	delay()
	l.Lock()
	value = 30
	s30.Put(469,value)
	log("p6 w:%d", value)
	countwrites(value,6,30)
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
	value = 18
	s18.Put(275,value)
	log("p6 w:%d", value)
	countwrites(value,6,18)
	l.Unlock()

	delay()
	l.Lock()
	value = 27
	s27.Put(420,value)
	log("p6 w:%d", value)
	countwrites(value,6,27)
	l.Unlock()

	delay()
	l.Lock()
	value = 22
	s22.Put(343,value)
	log("p6 w:%d", value)
	countwrites(value,6,22)
	l.Unlock()

	delay()
	l.Lock()
	value = 29
	s29.Put(459,value)
	log("p6 w:%d", value)
	countwrites(value,6,29)
	l.Unlock()

	delay()
	l.Lock()
	value = 17
	s17.Put(259,value)
	log("p6 w:%d", value)
	countwrites(value,6,17)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s7.QueryP(110,&value)
	log("p6 r:%d", value)
	countreads(value,6,7)
	l.Unlock()

	delay()
	l.Lock()
	value = 22
	s22.Put(344,value)
	log("p6 w:%d", value)
	countwrites(value,6,22)
	l.Unlock()

	delay()
	l.Lock()
	value = 27
	s27.Put(425,value)
	log("p6 w:%d", value)
	countwrites(value,6,27)
	l.Unlock()

	delay()
	l.Lock()
	value = 1
	s1.Put(13,value)
	log("p6 w:%d", value)
	countwrites(value,6,1)
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
	value = 21
	s21.Put(322,value)
	log("p6 w:%d", value)
	countwrites(value,6,21)
	l.Unlock()
}

func p7() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 15
	s15.Put(237,value)
	log("p7 w:%d", value)
	countwrites(value,7,15)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s31.QueryP(495,&value)
	log("p7 r:%d", value)
	countreads(value,7,31)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s9.QueryP(130,&value)
	log("p7 r:%d", value)
	countreads(value,7,9)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s31.QueryP(482,&value)
	log("p7 r:%d", value)
	countreads(value,7,31)
	l.Unlock()

	delay()
	l.Lock()
	value = 18
	s18.Put(282,value)
	log("p7 w:%d", value)
	countwrites(value,7,18)
	l.Unlock()

	delay()
	l.Lock()
	value = 30
	s30.Put(466,value)
	log("p7 w:%d", value)
	countwrites(value,7,30)
	l.Unlock()

	delay()
	l.Lock()
	value = 28
	s28.Put(442,value)
	log("p7 w:%d", value)
	countwrites(value,7,28)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s10.QueryP(147,&value)
	log("p7 r:%d", value)
	countreads(value,7,10)
	l.Unlock()

	delay()
	l.Lock()
	value = 30
	s30.Put(470,value)
	log("p7 w:%d", value)
	countwrites(value,7,30)
	l.Unlock()

	delay()
	l.Lock()
	value = 5
	s5.Put(74,value)
	log("p7 w:%d", value)
	countwrites(value,7,5)
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
	value = 22
	s22.Put(349,value)
	log("p7 w:%d", value)
	countwrites(value,7,22)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s32.QueryP(508,&value)
	log("p7 r:%d", value)
	countreads(value,7,32)
	l.Unlock()

	delay()
	l.Lock()
	value = 14
	s14.Put(221,value)
	log("p7 w:%d", value)
	countwrites(value,7,14)
	l.Unlock()

	delay()
	l.Lock()
	value = 31
	s31.Put(485,value)
	log("p7 w:%d", value)
	countwrites(value,7,31)
	l.Unlock()

	delay()
	l.Lock()
	value = 21
	s21.Put(324,value)
	log("p7 w:%d", value)
	countwrites(value,7,21)
	l.Unlock()
}

func p8() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 7
	s7.Put(112,value)
	log("p8 w:%d", value)
	countwrites(value,8,7)
	l.Unlock()

	delay()
	l.Lock()
	value = 1
	s1.Put(3,value)
	log("p8 w:%d", value)
	countwrites(value,8,1)
	l.Unlock()

	delay()
	l.Lock()
	value = 16
	s16.Put(254,value)
	log("p8 w:%d", value)
	countwrites(value,8,16)
	l.Unlock()

	delay()
	l.Lock()
	value = 22
	s22.Put(338,value)
	log("p8 w:%d", value)
	countwrites(value,8,22)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s23.QueryP(363,&value)
	log("p8 r:%d", value)
	countreads(value,8,23)
	l.Unlock()

	delay()
	l.Lock()
	value = 24
	s24.Put(375,value)
	log("p8 w:%d", value)
	countwrites(value,8,24)
	l.Unlock()

	delay()
	l.Lock()
	value = 29
	s29.Put(461,value)
	log("p8 w:%d", value)
	countwrites(value,8,29)
	l.Unlock()

	delay()
	l.Lock()
	value = 13
	s13.Put(202,value)
	log("p8 w:%d", value)
	countwrites(value,8,13)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s14.QueryP(213,&value)
	log("p8 r:%d", value)
	countreads(value,8,14)
	l.Unlock()

	delay()
	l.Lock()
	value = 6
	s6.Put(91,value)
	log("p8 w:%d", value)
	countwrites(value,8,6)
	l.Unlock()

	delay()
	l.Lock()
	value = 18
	s18.Put(273,value)
	log("p8 w:%d", value)
	countwrites(value,8,18)
	l.Unlock()

	delay()
	l.Lock()
	value = 2
	s2.Put(32,value)
	log("p8 w:%d", value)
	countwrites(value,8,2)
	l.Unlock()

	delay()
	l.Lock()
	value = 22
	s22.Put(349,value)
	log("p8 w:%d", value)
	countwrites(value,8,22)
	l.Unlock()

	delay()
	l.Lock()
	value = 8
	s8.Put(128,value)
	log("p8 w:%d", value)
	countwrites(value,8,8)
	l.Unlock()

	delay()
	l.Lock()
	value = 1
	s1.Put(8,value)
	log("p8 w:%d", value)
	countwrites(value,8,1)
	l.Unlock()

	delay()
	l.Lock()
	value = 4
	s4.Put(62,value)
	log("p8 w:%d", value)
	countwrites(value,8,4)
	l.Unlock()
}

func p9() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 1
	s1.Put(12,value)
	log("p9 w:%d", value)
	countwrites(value,9,1)
	l.Unlock()

	delay()
	l.Lock()
	value = 32
	s32.Put(499,value)
	log("p9 w:%d", value)
	countwrites(value,9,32)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s13.QueryP(199,&value)
	log("p9 r:%d", value)
	countreads(value,9,13)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s25.QueryP(388,&value)
	log("p9 r:%d", value)
	countreads(value,9,25)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s26.QueryP(402,&value)
	log("p9 r:%d", value)
	countreads(value,9,26)
	l.Unlock()

	delay()
	l.Lock()
	value = 10
	s10.Put(155,value)
	log("p9 w:%d", value)
	countwrites(value,9,10)
	l.Unlock()

	delay()
	l.Lock()
	value = 16
	s16.Put(250,value)
	log("p9 w:%d", value)
	countwrites(value,9,16)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s23.QueryP(368,&value)
	log("p9 r:%d", value)
	countreads(value,9,23)
	l.Unlock()

	delay()
	l.Lock()
	value = 27
	s27.Put(424,value)
	log("p9 w:%d", value)
	countwrites(value,9,27)
	l.Unlock()

	delay()
	l.Lock()
	value = 24
	s24.Put(373,value)
	log("p9 w:%d", value)
	countwrites(value,9,24)
	l.Unlock()

	delay()
	l.Lock()
	value = 11
	s11.Put(163,value)
	log("p9 w:%d", value)
	countwrites(value,9,11)
	l.Unlock()

	delay()
	l.Lock()
	value = 27
	s27.Put(424,value)
	log("p9 w:%d", value)
	countwrites(value,9,27)
	l.Unlock()

	delay()
	l.Lock()
	value = 4
	s4.Put(53,value)
	log("p9 w:%d", value)
	countwrites(value,9,4)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s19.QueryP(298,&value)
	log("p9 r:%d", value)
	countreads(value,9,19)
	l.Unlock()

	delay()
	l.Lock()
	value = 22
	s22.Put(349,value)
	log("p9 w:%d", value)
	countwrites(value,9,22)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s17.QueryP(261,&value)
	log("p9 r:%d", value)
	countreads(value,9,17)
	l.Unlock()
}


func p10() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 7
	s7.Put(103,value)
	log("p10 w:%d", value)
	countwrites(value,10,7)
	l.Unlock()

	delay()
	l.Lock()
	value = 12
	s12.Put(183,value)
	log("p10 w:%d", value)
	countwrites(value,10,12)
	l.Unlock()

	delay()
	l.Lock()
	value = 24
	s24.Put(377,value)
	log("p10 w:%d", value)
	countwrites(value,10,24)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s28.QueryP(439,&value)
	log("p10 r:%d", value)
	countreads(value,10,28)
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
	value = -1
	s16.QueryP(247,&value)
	log("p10 r:%d", value)
	countreads(value,10,16)
	l.Unlock()

	delay()
	l.Lock()
	value = 11
	s11.Put(166,value)
	log("p10 w:%d", value)
	countwrites(value,10,11)
	l.Unlock()

	delay()
	l.Lock()
	value = 21
	s21.Put(333,value)
	log("p10 w:%d", value)
	countwrites(value,10,21)
	l.Unlock()

	delay()
	l.Lock()
	value = 18
	s18.Put(283,value)
	log("p10 w:%d", value)
	countwrites(value,10,18)
	l.Unlock()

	delay()
	l.Lock()
	value = 31
	s31.Put(490,value)
	log("p10 w:%d", value)
	countwrites(value,10,31)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s12.QueryP(183,&value)
	log("p10 r:%d", value)
	countreads(value,10,12)
	l.Unlock()

	delay()
	l.Lock()
	value = 26
	s26.Put(413,value)
	log("p10 w:%d", value)
	countwrites(value,10,26)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s31.QueryP(490,&value)
	log("p10 r:%d", value)
	countreads(value,10,31)
	l.Unlock()

	delay()
	l.Lock()
	value = 22
	s22.Put(346,value)
	log("p10 w:%d", value)
	countwrites(value,10,22)
	l.Unlock()

	delay()
	l.Lock()
	value = 31
	s31.Put(489,value)
	log("p10 w:%d", value)
	countwrites(value,10,31)
	l.Unlock()

	delay()
	l.Lock()
	value = 20
	s20.Put(316,value)
	log("p10 w:%d", value)
	countwrites(value,10,20)
	l.Unlock()
}


func p11() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 28
	s28.Put(445,value)
	log("p11 w:%d", value)
	countwrites(value,11,28)
	l.Unlock()

	delay()
	l.Lock()
	value = 28
	s28.Put(442,value)
	log("p11 w:%d", value)
	countwrites(value,11,28)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s2.QueryP(29,&value)
	log("p11 r:%d", value)
	countreads(value,11,2)
	l.Unlock()

	delay()
	l.Lock()
	value = 21
	s21.Put(329,value)
	log("p11 w:%d", value)
	countwrites(value,11,21)
	l.Unlock()

	delay()
	l.Lock()
	value = 25
	s25.Put(393,value)
	log("p11 w:%d", value)
	countwrites(value,11,25)
	l.Unlock()

	delay()
	l.Lock()
	value = 19
	s19.Put(299,value)
	log("p11 w:%d", value)
	countwrites(value,11,19)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s9.QueryP(141,&value)
	log("p11 r:%d", value)
	countreads(value,11,9)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s8.QueryP(128,&value)
	log("p11 r:%d", value)
	countreads(value,11,8)
	l.Unlock()

	delay()
	l.Lock()
	value = 15
	s15.Put(234,value)
	log("p11 w:%d", value)
	countwrites(value,11,15)
	l.Unlock()

	delay()
	l.Lock()
	value = 9
	s9.Put(133,value)
	log("p11 w:%d", value)
	countwrites(value,11,9)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s20.QueryP(314,&value)
	log("p11 r:%d", value)
	countreads(value,11,20)
	l.Unlock()

	delay()
	l.Lock()
	value = 8
	s8.Put(114,value)
	log("p11 w:%d", value)
	countwrites(value,11,8)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s25.QueryP(388,&value)
	log("p11 r:%d", value)
	countreads(value,11,25)
	l.Unlock()

	delay()
	l.Lock()
	value = 18
	s18.Put(285,value)
	log("p11 w:%d", value)
	countwrites(value,11,18)
	l.Unlock()

	delay()
	l.Lock()
	value = 14
	s14.Put(217,value)
	log("p11 w:%d", value)
	countwrites(value,11,14)
	l.Unlock()

	delay()
	l.Lock()
	value = 9
	s9.Put(142,value)
	log("p11 w:%d", value)
	countwrites(value,11,9)
	l.Unlock()
}


func p12() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	s21.QueryP(325,&value)
	log("p12 r:%d", value)
	countreads(value,12,21)
	l.Unlock()

	delay()
	l.Lock()
	value = 8
	s8.Put(120,value)
	log("p12 w:%d", value)
	countwrites(value,12,8)
	l.Unlock()

	delay()
	l.Lock()
	value = 10
	s10.Put(146,value)
	log("p12 w:%d", value)
	countwrites(value,12,10)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s27.QueryP(426,&value)
	log("p12 r:%d", value)
	countreads(value,12,27)
	l.Unlock()

	delay()
	l.Lock()
	value = 28
	s28.Put(435,value)
	log("p12 w:%d", value)
	countwrites(value,12,28)
	l.Unlock()

	delay()
	l.Lock()
	value = 15
	s15.Put(227,value)
	log("p12 w:%d", value)
	countwrites(value,12,15)
	l.Unlock()

	delay()
	l.Lock()
	value = 12
	s12.Put(192,value)
	log("p12 w:%d", value)
	countwrites(value,12,12)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s17.QueryP(266,&value)
	log("p12 r:%d", value)
	countreads(value,12,17)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s28.QueryP(446,&value)
	log("p12 r:%d", value)
	countreads(value,12,28)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s15.QueryP(233,&value)
	log("p12 r:%d", value)
	countreads(value,12,15)
	l.Unlock()

	delay()
	l.Lock()
	value = 4
	s4.Put(58,value)
	log("p12 w:%d", value)
	countwrites(value,12,4)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s6.QueryP(95,&value)
	log("p12 r:%d", value)
	countreads(value,12,6)
	l.Unlock()

	delay()
	l.Lock()
	value = 25
	s25.Put(391,value)
	log("p12 w:%d", value)
	countwrites(value,12,25)
	l.Unlock()

	delay()
	l.Lock()
	value = 32
	s32.Put(508,value)
	log("p12 w:%d", value)
	countwrites(value,12,32)
	l.Unlock()

	delay()
	l.Lock()
	value = 22
	s22.Put(337,value)
	log("p12 w:%d", value)
	countwrites(value,12,22)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s31.QueryP(485,&value)
	log("p12 r:%d", value)
	countreads(value,12,31)
	l.Unlock()
}


func p13() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 8
	s8.Put(117,value)
	log("p13 w:%d", value)
	countwrites(value,13,8)
	l.Unlock()

	delay()
	l.Lock()
	value = 14
	s14.Put(223,value)
	log("p13 w:%d", value)
	countwrites(value,13,14)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s27.QueryP(425,&value)
	log("p13 r:%d", value)
	countreads(value,13,27)
	l.Unlock()

	delay()
	l.Lock()
	value = 17
	s17.Put(270,value)
	log("p13 w:%d", value)
	countwrites(value,13,17)
	l.Unlock()

	delay()
	l.Lock()
	value = 1
	s1.Put(7,value)
	log("p13 w:%d", value)
	countwrites(value,13,1)
	l.Unlock()

	delay()
	l.Lock()
	value = 29
	s29.Put(462,value)
	log("p13 w:%d", value)
	countwrites(value,13,29)
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
	value = 11
	s11.Put(175,value)
	log("p13 w:%d", value)
	countwrites(value,13,11)
	l.Unlock()

	delay()
	l.Lock()
	value = 25
	s25.Put(387,value)
	log("p13 w:%d", value)
	countwrites(value,13,25)
	l.Unlock()

	delay()
	l.Lock()
	value = 6
	s6.Put(84,value)
	log("p13 w:%d", value)
	countwrites(value,13,6)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s28.QueryP(446,&value)
	log("p13 r:%d", value)
	countreads(value,13,28)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s24.QueryP(380,&value)
	log("p13 r:%d", value)
	countreads(value,13,24)
	l.Unlock()

	delay()
	l.Lock()
	value = 31
	s31.Put(481,value)
	log("p13 w:%d", value)
	countwrites(value,13,31)
	l.Unlock()

	delay()
	l.Lock()
	value = 27
	s27.Put(423,value)
	log("p13 w:%d", value)
	countwrites(value,13,27)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s22.QueryP(339,&value)
	log("p13 r:%d", value)
	countreads(value,13,22)
	l.Unlock()

	delay()
	l.Lock()
	value = 1
	s1.Put(1,value)
	log("p13 w:%d", value)
	countwrites(value,13,1)
	l.Unlock()
}

func p14() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 15
	s15.Put(239,value)
	log("p14 w:%d", value)
	countwrites(value,14,15)
	l.Unlock()

	delay()
	l.Lock()
	value = 21
	s21.Put(323,value)
	log("p14 w:%d", value)
	countwrites(value,14,21)
	l.Unlock()

	delay()
	l.Lock()
	value = 8
	s8.Put(120,value)
	log("p14 w:%d", value)
	countwrites(value,14,8)
	l.Unlock()

	delay()
	l.Lock()
	value = 19
	s19.Put(303,value)
	log("p14 w:%d", value)
	countwrites(value,14,19)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s28.QueryP(437,&value)
	log("p14 r:%d", value)
	countreads(value,14,28)
	l.Unlock()

	delay()
	l.Lock()
	value = 14
	s14.Put(223,value)
	log("p14 w:%d", value)
	countwrites(value,14,14)
	l.Unlock()

	delay()
	l.Lock()
	value = 27
	s27.Put(430,value)
	log("p14 w:%d", value)
	countwrites(value,14,27)
	l.Unlock()

	delay()
	l.Lock()
	value = 31
	s31.Put(496,value)
	log("p14 w:%d", value)
	countwrites(value,14,31)
	l.Unlock()

	delay()
	l.Lock()
	value = 20
	s20.Put(319,value)
	log("p14 w:%d", value)
	countwrites(value,14,20)
	l.Unlock()

	delay()
	l.Lock()
	value = 28
	s28.Put(435,value)
	log("p14 w:%d", value)
	countwrites(value,14,28)
	l.Unlock()

	delay()
	l.Lock()
	value = 30
	s30.Put(471,value)
	log("p14 w:%d", value)
	countwrites(value,14,30)
	l.Unlock()

	delay()
	l.Lock()
	value = 17
	s17.Put(267,value)
	log("p14 w:%d", value)
	countwrites(value,14,17)
	l.Unlock()

	delay()
	l.Lock()
	value = 32
	s32.Put(501,value)
	log("p14 w:%d", value)
	countwrites(value,14,32)
	l.Unlock()

	delay()
	l.Lock()
	value = 6
	s6.Put(90,value)
	log("p14 w:%d", value)
	countwrites(value,14,6)
	l.Unlock()

	delay()
	l.Lock()
	value = 19
	s19.Put(304,value)
	log("p14 w:%d", value)
	countwrites(value,14,19)
	l.Unlock()

	delay()
	l.Lock()
	value = 4
	s4.Put(49,value)
	log("p14 w:%d", value)
	countwrites(value,14,4)
	l.Unlock()
}


func p15() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	s7.QueryP(112,&value)
	log("p15 r:%d", value)
	countreads(value,15,7)
	l.Unlock()

	delay()
	l.Lock()
	value = 32
	s32.Put(506,value)
	log("p15 w:%d", value)
	countwrites(value,15,32)
	l.Unlock()

	delay()
	l.Lock()
	value = 3
	s3.Put(41,value)
	log("p15 w:%d", value)
	countwrites(value,15,3)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s15.QueryP(226,&value)
	log("p15 r:%d", value)
	countreads(value,15,15)
	l.Unlock()

	delay()
	l.Lock()
	value = 25
	s25.Put(391,value)
	log("p15 w:%d", value)
	countwrites(value,15,25)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s22.QueryP(344,&value)
	log("p15 r:%d", value)
	countreads(value,15,22)
	l.Unlock()

	delay()
	l.Lock()
	value = 24
	s24.Put(381,value)
	log("p15 w:%d", value)
	countwrites(value,15,24)
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
	value = 16
	s16.Put(245,value)
	log("p15 w:%d", value)
	countwrites(value,15,16)
	l.Unlock()

	delay()
	l.Lock()
	value = 17
	s17.Put(270,value)
	log("p15 w:%d", value)
	countwrites(value,15,17)
	l.Unlock()

	delay()
	l.Lock()
	value = 17
	s17.Put(259,value)
	log("p15 w:%d", value)
	countwrites(value,15,17)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s27.QueryP(426,&value)
	log("p15 r:%d", value)
	countreads(value,15,27)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s5.QueryP(75,&value)
	log("p15 r:%d", value)
	countreads(value,15,5)
	l.Unlock()

	delay()
	l.Lock()
	value = 5
	s5.Put(70,value)
	log("p15 w:%d", value)
	countwrites(value,15,5)
	l.Unlock()

	delay()
	l.Lock()
	value = 31
	s31.Put(489,value)
	log("p15 w:%d", value)
	countwrites(value,15,31)
	l.Unlock()

	delay()
	l.Lock()
	value = 1
	s1.Put(14,value)
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
	s27.QueryP(421,&value)
	log("p16 r:%d", value)
	countreads(value,16,27)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s14.QueryP(210,&value)
	log("p16 r:%d", value)
	countreads(value,16,14)
	l.Unlock()

	delay()
	l.Lock()
	value = 31
	s31.Put(482,value)
	log("p16 w:%d", value)
	countwrites(value,16,31)
	l.Unlock()

	delay()
	l.Lock()
	value = 5
	s5.Put(80,value)
	log("p16 w:%d", value)
	countwrites(value,16,5)
	l.Unlock()

	delay()
	l.Lock()
	value = 23
	s23.Put(361,value)
	log("p16 w:%d", value)
	countwrites(value,16,23)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s16.QueryP(253,&value)
	log("p16 r:%d", value)
	countreads(value,16,16)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s14.QueryP(210,&value)
	log("p16 r:%d", value)
	countreads(value,16,14)
	l.Unlock()

	delay()
	l.Lock()
	value = 2
	s2.Put(20,value)
	log("p16 w:%d", value)
	countwrites(value,16,2)
	l.Unlock()

	delay()
	l.Lock()
	value = 12
	s12.Put(184,value)
	log("p16 w:%d", value)
	countwrites(value,16,12)
	l.Unlock()

	delay()
	l.Lock()
	value = 6
	s6.Put(95,value)
	log("p16 w:%d", value)
	countwrites(value,16,6)
	l.Unlock()

	delay()
	l.Lock()
	value = 28
	s28.Put(437,value)
	log("p16 w:%d", value)
	countwrites(value,16,28)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s30.QueryP(469,&value)
	log("p16 r:%d", value)
	countreads(value,16,30)
	l.Unlock()

	delay()
	l.Lock()
	value = 14
	s14.Put(220,value)
	log("p16 w:%d", value)
	countwrites(value,16,14)
	l.Unlock()

	delay()
	l.Lock()
	value = 5
	s5.Put(76,value)
	log("p16 w:%d", value)
	countwrites(value,16,5)
	l.Unlock()

	delay()
	l.Lock()
	value = 7
	s7.Put(101,value)
	log("p16 w:%d", value)
	countwrites(value,16,7)
	l.Unlock()

	delay()
	l.Lock()
	value = 32
	s32.Put(510,value)
	log("p16 w:%d", value)
	countwrites(value,16,32)
	l.Unlock()
}


func p17() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 14
	s14.Put(214,value)
	log("p17 w:%d", value)
	countwrites(value,17,14)
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
	value = 28
	s28.Put(438,value)
	log("p17 w:%d", value)
	countwrites(value,17,28)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s11.QueryP(176,&value)
	log("p17 r:%d", value)
	countreads(value,17,11)
	l.Unlock()

	delay()
	l.Lock()
	value = 4
	s4.Put(53,value)
	log("p17 w:%d", value)
	countwrites(value,17,4)
	l.Unlock()

	delay()
	l.Lock()
	value = 28
	s28.Put(439,value)
	log("p17 w:%d", value)
	countwrites(value,17,28)
	l.Unlock()

	delay()
	l.Lock()
	value = 13
	s13.Put(201,value)
	log("p17 w:%d", value)
	countwrites(value,17,13)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s31.QueryP(484,&value)
	log("p17 r:%d", value)
	countreads(value,17,31)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s24.QueryP(382,&value)
	log("p17 r:%d", value)
	countreads(value,17,24)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s14.QueryP(217,&value)
	log("p17 r:%d", value)
	countreads(value,17,14)
	l.Unlock()

	delay()
	l.Lock()
	value = 30
	s30.Put(478,value)
	log("p17 w:%d", value)
	countwrites(value,17,30)
	l.Unlock()

	delay()
	l.Lock()
	value = 13
	s13.Put(205,value)
	log("p17 w:%d", value)
	countwrites(value,17,13)
	l.Unlock()

	delay()
	l.Lock()
	value = 29
	s29.Put(459,value)
	log("p17 w:%d", value)
	countwrites(value,17,29)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s21.QueryP(321,&value)
	log("p17 r:%d", value)
	countreads(value,17,21)
	l.Unlock()

	delay()
	l.Lock()
	value = 19
	s19.Put(292,value)
	log("p17 w:%d", value)
	countwrites(value,17,19)
	l.Unlock()

	delay()
	l.Lock()
	value = 2
	s2.Put(32,value)
	log("p17 w:%d", value)
	countwrites(value,17,2)
	l.Unlock()
}


func p18() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	s12.QueryP(187,&value)
	log("p18 r:%d", value)
	countreads(value,18,12)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s2.QueryP(20,&value)
	log("p18 r:%d", value)
	countreads(value,18,2)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s8.QueryP(116,&value)
	log("p18 r:%d", value)
	countreads(value,18,8)
	l.Unlock()

	delay()
	l.Lock()
	value = 6
	s6.Put(94,value)
	log("p18 w:%d", value)
	countwrites(value,18,6)
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
	value = 32
	s32.Put(504,value)
	log("p18 w:%d", value)
	countwrites(value,18,32)
	l.Unlock()

	delay()
	l.Lock()
	value = 17
	s17.Put(260,value)
	log("p18 w:%d", value)
	countwrites(value,18,17)
	l.Unlock()

	delay()
	l.Lock()
	value = 5
	s5.Put(76,value)
	log("p18 w:%d", value)
	countwrites(value,18,5)
	l.Unlock()

	delay()
	l.Lock()
	value = 5
	s5.Put(72,value)
	log("p18 w:%d", value)
	countwrites(value,18,5)
	l.Unlock()

	delay()
	l.Lock()
	value = 8
	s8.Put(115,value)
	log("p18 w:%d", value)
	countwrites(value,18,8)
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
	value = -1
	s1.QueryP(9,&value)
	log("p18 r:%d", value)
	countreads(value,18,1)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s24.QueryP(377,&value)
	log("p18 r:%d", value)
	countreads(value,18,24)
	l.Unlock()

	delay()
	l.Lock()
	value = 32
	s32.Put(507,value)
	log("p18 w:%d", value)
	countwrites(value,18,32)
	l.Unlock()

	delay()
	l.Lock()
	value = 1
	s1.Put(9,value)
	log("p18 w:%d", value)
	countwrites(value,18,1)
	l.Unlock()

	delay()
	l.Lock()
	value = 32
	s32.Put(500,value)
	log("p18 w:%d", value)
	countwrites(value,18,32)
	l.Unlock()
}


func p19() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 6
	s6.Put(94,value)
	log("p19 w:%d", value)
	countwrites(value,19,6)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s11.QueryP(163,&value)
	log("p19 r:%d", value)
	countreads(value,19,11)
	l.Unlock()

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
	s18.QueryP(273,&value)
	log("p19 r:%d", value)
	countreads(value,19,18)
	l.Unlock()

	delay()
	l.Lock()
	value = 24
	s24.Put(369,value)
	log("p19 w:%d", value)
	countwrites(value,19,24)
	l.Unlock()

	delay()
	l.Lock()
	value = 11
	s11.Put(167,value)
	log("p19 w:%d", value)
	countwrites(value,19,11)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s21.QueryP(321,&value)
	log("p19 r:%d", value)
	countreads(value,19,21)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s12.QueryP(187,&value)
	log("p19 r:%d", value)
	countreads(value,19,12)
	l.Unlock()

	delay()
	l.Lock()
	value = 10
	s10.Put(154,value)
	log("p19 w:%d", value)
	countwrites(value,19,10)
	l.Unlock()

	delay()
	l.Lock()
	value = 11
	s11.Put(162,value)
	log("p19 w:%d", value)
	countwrites(value,19,11)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s11.QueryP(164,&value)
	log("p19 r:%d", value)
	countreads(value,19,11)
	l.Unlock()

	delay()
	l.Lock()
	value = 28
	s28.Put(441,value)
	log("p19 w:%d", value)
	countwrites(value,19,28)
	l.Unlock()

	delay()
	l.Lock()
	value = 10
	s10.Put(158,value)
	log("p19 w:%d", value)
	countwrites(value,19,10)
	l.Unlock()

	delay()
	l.Lock()
	value = 18
	s18.Put(276,value)
	log("p19 w:%d", value)
	countwrites(value,19,18)
	l.Unlock()

	delay()
	l.Lock()
	value = 6
	s6.Put(81,value)
	log("p19 w:%d", value)
	countwrites(value,19,6)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s24.QueryP(374,&value)
	log("p19 r:%d", value)
	countreads(value,19,24)
	l.Unlock()
}


func p20() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 2
	s2.Put(18,value)
	log("p20 w:%d", value)
	countwrites(value,20,2)
	l.Unlock()

	delay()
	l.Lock()
	value = 9
	s9.Put(144,value)
	log("p20 w:%d", value)
	countwrites(value,20,9)
	l.Unlock()

	delay()
	l.Lock()
	value = 19
	s19.Put(304,value)
	log("p20 w:%d", value)
	countwrites(value,20,19)
	l.Unlock()

	delay()
	l.Lock()
	value = 29
	s29.Put(463,value)
	log("p20 w:%d", value)
	countwrites(value,20,29)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s13.QueryP(205,&value)
	log("p20 r:%d", value)
	countreads(value,20,13)
	l.Unlock()

	delay()
	l.Lock()
	value = 1
	s1.Put(14,value)
	log("p20 w:%d", value)
	countwrites(value,20,1)
	l.Unlock()

	delay()
	l.Lock()
	value = 6
	s6.Put(94,value)
	log("p20 w:%d", value)
	countwrites(value,20,6)
	l.Unlock()

	delay()
	l.Lock()
	value = 23
	s23.Put(368,value)
	log("p20 w:%d", value)
	countwrites(value,20,23)
	l.Unlock()

	delay()
	l.Lock()
	value = 12
	s12.Put(179,value)
	log("p20 w:%d", value)
	countwrites(value,20,12)
	l.Unlock()

	delay()
	l.Lock()
	value = 3
	s3.Put(48,value)
	log("p20 w:%d", value)
	countwrites(value,20,3)
	l.Unlock()

	delay()
	l.Lock()
	value = 3
	s3.Put(48,value)
	log("p20 w:%d", value)
	countwrites(value,20,3)
	l.Unlock()

	delay()
	l.Lock()
	value = 29
	s29.Put(457,value)
	log("p20 w:%d", value)
	countwrites(value,20,29)
	l.Unlock()

	delay()
	l.Lock()
	value = 21
	s21.Put(332,value)
	log("p20 w:%d", value)
	countwrites(value,20,21)
	l.Unlock()

	delay()
	l.Lock()
	value = 7
	s7.Put(99,value)
	log("p20 w:%d", value)
	countwrites(value,20,7)
	l.Unlock()

	delay()
	l.Lock()
	value = 24
	s24.Put(379,value)
	log("p20 w:%d", value)
	countwrites(value,20,24)
	l.Unlock()

	delay()
	l.Lock()
	value = 22
	s22.Put(345,value)
	log("p20 w:%d", value)
	countwrites(value,20,22)
	l.Unlock()
}


func p21() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	s9.QueryP(136,&value)
	log("p21 r:%d", value)
	countreads(value,21,9)
	l.Unlock()

	delay()
	l.Lock()
	value = 27
	s27.Put(431,value)
	log("p21 w:%d", value)
	countwrites(value,21,27)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s19.QueryP(292,&value)
	log("p21 r:%d", value)
	countreads(value,21,19)
	l.Unlock()

	delay()
	l.Lock()
	value = 16
	s16.Put(245,value)
	log("p21 w:%d", value)
	countwrites(value,21,16)
	l.Unlock()

	delay()
	l.Lock()
	value = 21
	s21.Put(322,value)
	log("p21 w:%d", value)
	countwrites(value,21,21)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s8.QueryP(128,&value)
	log("p21 r:%d", value)
	countreads(value,21,8)
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
	value = 26
	s26.Put(407,value)
	log("p21 w:%d", value)
	countwrites(value,21,26)
	l.Unlock()

	delay()
	l.Lock()
	value = 19
	s19.Put(302,value)
	log("p21 w:%d", value)
	countwrites(value,21,19)
	l.Unlock()

	delay()
	l.Lock()
	value = 15
	s15.Put(236,value)
	log("p21 w:%d", value)
	countwrites(value,21,15)
	l.Unlock()

	delay()
	l.Lock()
	value = 32
	s32.Put(509,value)
	log("p21 w:%d", value)
	countwrites(value,21,32)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s21.QueryP(336,&value)
	log("p21 r:%d", value)
	countreads(value,21,21)
	l.Unlock()

	delay()
	l.Lock()
	value = 27
	s27.Put(430,value)
	log("p21 w:%d", value)
	countwrites(value,21,27)
	l.Unlock()

	delay()
	l.Lock()
	value = 30
	s30.Put(477,value)
	log("p21 w:%d", value)
	countwrites(value,21,30)
	l.Unlock()

	delay()
	l.Lock()
	value = 26
	s26.Put(403,value)
	log("p21 w:%d", value)
	countwrites(value,21,26)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s20.QueryP(316,&value)
	log("p21 r:%d", value)
	countreads(value,21,20)
	l.Unlock()
}


func p22() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	s12.QueryP(187,&value)
	log("p22 r:%d", value)
	countreads(value,22,12)
	l.Unlock()

	delay()
	l.Lock()
	value = 11
	s11.Put(167,value)
	log("p22 w:%d", value)
	countwrites(value,22,11)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s28.QueryP(435,&value)
	log("p22 r:%d", value)
	countreads(value,22,28)
	l.Unlock()

	delay()
	l.Lock()
	value = 22
	s22.Put(345,value)
	log("p22 w:%d", value)
	countwrites(value,22,22)
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
	value = 3
	s3.Put(45,value)
	log("p22 w:%d", value)
	countwrites(value,22,3)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s24.QueryP(383,&value)
	log("p22 r:%d", value)
	countreads(value,22,24)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s26.QueryP(409,&value)
	log("p22 r:%d", value)
	countreads(value,22,26)
	l.Unlock()

	delay()
	l.Lock()
	value = 13
	s13.Put(205,value)
	log("p22 w:%d", value)
	countwrites(value,22,13)
	l.Unlock()

	delay()
	l.Lock()
	value = 15
	s15.Put(228,value)
	log("p22 w:%d", value)
	countwrites(value,22,15)
	l.Unlock()

	delay()
	l.Lock()
	value = 8
	s8.Put(113,value)
	log("p22 w:%d", value)
	countwrites(value,22,8)
	l.Unlock()

	delay()
	l.Lock()
	value = 27
	s27.Put(423,value)
	log("p22 w:%d", value)
	countwrites(value,22,27)
	l.Unlock()

	delay()
	l.Lock()
	value = 21
	s21.Put(326,value)
	log("p22 w:%d", value)
	countwrites(value,22,21)
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
	value = 23
	s23.Put(358,value)
	log("p22 w:%d", value)
	countwrites(value,22,23)
	l.Unlock()

	delay()
	l.Lock()
	value = 17
	s17.Put(269,value)
	log("p22 w:%d", value)
	countwrites(value,22,17)
	l.Unlock()
}


func p23() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 22
	s22.Put(348,value)
	log("p23 w:%d", value)
	countwrites(value,23,22)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s7.QueryP(102,&value)
	log("p23 r:%d", value)
	countreads(value,23,7)
	l.Unlock()

	delay()
	l.Lock()
	value = 18
	s18.Put(282,value)
	log("p23 w:%d", value)
	countwrites(value,23,18)
	l.Unlock()

	delay()
	l.Lock()
	value = 18
	s18.Put(285,value)
	log("p23 w:%d", value)
	countwrites(value,23,18)
	l.Unlock()

	delay()
	l.Lock()
	value = 2
	s2.Put(25,value)
	log("p23 w:%d", value)
	countwrites(value,23,2)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s6.QueryP(90,&value)
	log("p23 r:%d", value)
	countreads(value,23,6)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s7.QueryP(97,&value)
	log("p23 r:%d", value)
	countreads(value,23,7)
	l.Unlock()

	delay()
	l.Lock()
	value = 13
	s13.Put(208,value)
	log("p23 w:%d", value)
	countwrites(value,23,13)
	l.Unlock()

	delay()
	l.Lock()
	value = 26
	s26.Put(402,value)
	log("p23 w:%d", value)
	countwrites(value,23,26)
	l.Unlock()

	delay()
	l.Lock()
	value = 24
	s24.Put(380,value)
	log("p23 w:%d", value)
	countwrites(value,23,24)
	l.Unlock()

	delay()
	l.Lock()
	value = 29
	s29.Put(457,value)
	log("p23 w:%d", value)
	countwrites(value,23,29)
	l.Unlock()

	delay()
	l.Lock()
	value = 11
	s11.Put(173,value)
	log("p23 w:%d", value)
	countwrites(value,23,11)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s7.QueryP(110,&value)
	log("p23 r:%d", value)
	countreads(value,23,7)
	l.Unlock()

	delay()
	l.Lock()
	value = 9
	s9.Put(139,value)
	log("p23 w:%d", value)
	countwrites(value,23,9)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s10.QueryP(148,&value)
	log("p23 r:%d", value)
	countreads(value,23,10)
	l.Unlock()

	delay()
	l.Lock()
	value = 30
	s30.Put(465,value)
	log("p23 w:%d", value)
	countwrites(value,23,30)
	l.Unlock()
}


func p24() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 7
	s7.Put(98,value)
	log("p24 w:%d", value)
	countwrites(value,24,7)
	l.Unlock()

	delay()
	l.Lock()
	value = 28
	s28.Put(440,value)
	log("p24 w:%d", value)
	countwrites(value,24,28)
	l.Unlock()

	delay()
	l.Lock()
	value = 5
	s5.Put(77,value)
	log("p24 w:%d", value)
	countwrites(value,24,5)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s5.QueryP(78,&value)
	log("p24 r:%d", value)
	countreads(value,24,5)
	l.Unlock()

	delay()
	l.Lock()
	value = 4
	s4.Put(56,value)
	log("p24 w:%d", value)
	countwrites(value,24,4)
	l.Unlock()

	delay()
	l.Lock()
	value = 10
	s10.Put(149,value)
	log("p24 w:%d", value)
	countwrites(value,24,10)
	l.Unlock()

	delay()
	l.Lock()
	value = 30
	s30.Put(476,value)
	log("p24 w:%d", value)
	countwrites(value,24,30)
	l.Unlock()

	delay()
	l.Lock()
	value = 10
	s10.Put(152,value)
	log("p24 w:%d", value)
	countwrites(value,24,10)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s28.QueryP(444,&value)
	log("p24 r:%d", value)
	countreads(value,24,28)
	l.Unlock()

	delay()
	l.Lock()
	value = 17
	s17.Put(258,value)
	log("p24 w:%d", value)
	countwrites(value,24,17)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s32.QueryP(498,&value)
	log("p24 r:%d", value)
	countreads(value,24,32)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s22.QueryP(338,&value)
	log("p24 r:%d", value)
	countreads(value,24,22)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s10.QueryP(147,&value)
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
	value = 31
	s31.Put(484,value)
	log("p24 w:%d", value)
	countwrites(value,24,31)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s28.QueryP(443,&value)
	log("p24 r:%d", value)
	countreads(value,24,28)
	l.Unlock()
}


func p25() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 6
	s6.Put(96,value)
	log("p25 w:%d", value)
	countwrites(value,25,6)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s16.QueryP(244,&value)
	log("p25 r:%d", value)
	countreads(value,25,16)
	l.Unlock()

	delay()
	l.Lock()
	value = 32
	s32.Put(501,value)
	log("p25 w:%d", value)
	countwrites(value,25,32)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s28.QueryP(437,&value)
	log("p25 r:%d", value)
	countreads(value,25,28)
	l.Unlock()

	delay()
	l.Lock()
	value = 24
	s24.Put(378,value)
	log("p25 w:%d", value)
	countwrites(value,25,24)
	l.Unlock()

	delay()
	l.Lock()
	value = 29
	s29.Put(454,value)
	log("p25 w:%d", value)
	countwrites(value,25,29)
	l.Unlock()

	delay()
	l.Lock()
	value = 25
	s25.Put(398,value)
	log("p25 w:%d", value)
	countwrites(value,25,25)
	l.Unlock()

	delay()
	l.Lock()
	value = 32
	s32.Put(505,value)
	log("p25 w:%d", value)
	countwrites(value,25,32)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s6.QueryP(90,&value)
	log("p25 r:%d", value)
	countreads(value,25,6)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s22.QueryP(343,&value)
	log("p25 r:%d", value)
	countreads(value,25,22)
	l.Unlock()

	delay()
	l.Lock()
	value = 1
	s1.Put(1,value)
	log("p25 w:%d", value)
	countwrites(value,25,1)
	l.Unlock()

	delay()
	l.Lock()
	value = 31
	s31.Put(482,value)
	log("p25 w:%d", value)
	countwrites(value,25,31)
	l.Unlock()

	delay()
	l.Lock()
	value = 20
	s20.Put(309,value)
	log("p25 w:%d", value)
	countwrites(value,25,20)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s16.QueryP(244,&value)
	log("p25 r:%d", value)
	countreads(value,25,16)
	l.Unlock()

	delay()
	l.Lock()
	value = 18
	s18.Put(276,value)
	log("p25 w:%d", value)
	countwrites(value,25,18)
	l.Unlock()

	delay()
	l.Lock()
	value = 17
	s17.Put(267,value)
	log("p25 w:%d", value)
	countwrites(value,25,17)
	l.Unlock()
}


func p26() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 24
	s24.Put(379,value)
	log("p26 w:%d", value)
	countwrites(value,26,24)
	l.Unlock()

	delay()
	l.Lock()
	value = 32
	s32.Put(498,value)
	log("p26 w:%d", value)
	countwrites(value,26,32)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s24.QueryP(376,&value)
	log("p26 r:%d", value)
	countreads(value,26,24)
	l.Unlock()

	delay()
	l.Lock()
	value = 8
	s8.Put(125,value)
	log("p26 w:%d", value)
	countwrites(value,26,8)
	l.Unlock()

	delay()
	l.Lock()
	value = 21
	s21.Put(335,value)
	log("p26 w:%d", value)
	countwrites(value,26,21)
	l.Unlock()

	delay()
	l.Lock()
	value = 11
	s11.Put(169,value)
	log("p26 w:%d", value)
	countwrites(value,26,11)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s18.QueryP(287,&value)
	log("p26 r:%d", value)
	countreads(value,26,18)
	l.Unlock()

	delay()
	l.Lock()
	value = 16
	s16.Put(254,value)
	log("p26 w:%d", value)
	countwrites(value,26,16)
	l.Unlock()

	delay()
	l.Lock()
	value = 10
	s10.Put(149,value)
	log("p26 w:%d", value)
	countwrites(value,26,10)
	l.Unlock()

	delay()
	l.Lock()
	value = 31
	s31.Put(489,value)
	log("p26 w:%d", value)
	countwrites(value,26,31)
	l.Unlock()

	delay()
	l.Lock()
	value = 16
	s16.Put(249,value)
	log("p26 w:%d", value)
	countwrites(value,26,16)
	l.Unlock()

	delay()
	l.Lock()
	value = 28
	s28.Put(447,value)
	log("p26 w:%d", value)
	countwrites(value,26,28)
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
	value = 18
	s18.Put(281,value)
	log("p26 w:%d", value)
	countwrites(value,26,18)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s5.QueryP(73,&value)
	log("p26 r:%d", value)
	countreads(value,26,5)
	l.Unlock()

	delay()
	l.Lock()
	value = 25
	s25.Put(387,value)
	log("p26 w:%d", value)
	countwrites(value,26,25)
	l.Unlock()
}


func p27() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 12
	s12.Put(178,value)
	log("p27 w:%d", value)
	countwrites(value,27,12)
	l.Unlock()

	delay()
	l.Lock()
	value = 25
	s25.Put(388,value)
	log("p27 w:%d", value)
	countwrites(value,27,25)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s15.QueryP(239,&value)
	log("p27 r:%d", value)
	countreads(value,27,15)
	l.Unlock()

	delay()
	l.Lock()
	value = 15
	s15.Put(230,value)
	log("p27 w:%d", value)
	countwrites(value,27,15)
	l.Unlock()

	delay()
	l.Lock()
	value = 10
	s10.Put(151,value)
	log("p27 w:%d", value)
	countwrites(value,27,10)
	l.Unlock()

	delay()
	l.Lock()
	value = 4
	s4.Put(56,value)
	log("p27 w:%d", value)
	countwrites(value,27,4)
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
	value = 24
	s24.Put(371,value)
	log("p27 w:%d", value)
	countwrites(value,27,24)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s28.QueryP(435,&value)
	log("p27 r:%d", value)
	countreads(value,27,28)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s29.QueryP(450,&value)
	log("p27 r:%d", value)
	countreads(value,27,29)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s1.QueryP(6,&value)
	log("p27 r:%d", value)
	countreads(value,27,1)
	l.Unlock()

	delay()
	l.Lock()
	value = 9
	s9.Put(135,value)
	log("p27 w:%d", value)
	countwrites(value,27,9)
	l.Unlock()

	delay()
	l.Lock()
	value = 2
	s2.Put(26,value)
	log("p27 w:%d", value)
	countwrites(value,27,2)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s26.QueryP(412,&value)
	log("p27 r:%d", value)
	countreads(value,27,26)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s12.QueryP(179,&value)
	log("p27 r:%d", value)
	countreads(value,27,12)
	l.Unlock()

	delay()
	l.Lock()
	value = 27
	s27.Put(423,value)
	log("p27 w:%d", value)
	countwrites(value,27,27)
	l.Unlock()
}


func p28() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 24
	s24.Put(371,value)
	log("p28 w:%d", value)
	countwrites(value,28,24)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s10.QueryP(147,&value)
	log("p28 r:%d", value)
	countreads(value,28,10)
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
	value = 20
	s20.Put(310,value)
	log("p28 w:%d", value)
	countwrites(value,28,20)
	l.Unlock()

	delay()
	l.Lock()
	value = 29
	s29.Put(464,value)
	log("p28 w:%d", value)
	countwrites(value,28,29)
	l.Unlock()

	delay()
	l.Lock()
	value = 30
	s30.Put(468,value)
	log("p28 w:%d", value)
	countwrites(value,28,30)
	l.Unlock()

	delay()
	l.Lock()
	value = 18
	s18.Put(288,value)
	log("p28 w:%d", value)
	countwrites(value,28,18)
	l.Unlock()

	delay()
	l.Lock()
	value = 4
	s4.Put(50,value)
	log("p28 w:%d", value)
	countwrites(value,28,4)
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
	s4.QueryP(50,&value)
	log("p28 r:%d", value)
	countreads(value,28,4)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s7.QueryP(101,&value)
	log("p28 r:%d", value)
	countreads(value,28,7)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s21.QueryP(327,&value)
	log("p28 r:%d", value)
	countreads(value,28,21)
	l.Unlock()

	delay()
	l.Lock()
	value = 4
	s4.Put(54,value)
	log("p28 w:%d", value)
	countwrites(value,28,4)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s26.QueryP(416,&value)
	log("p28 r:%d", value)
	countreads(value,28,26)
	l.Unlock()

	delay()
	l.Lock()
	value = 12
	s12.Put(191,value)
	log("p28 w:%d", value)
	countwrites(value,28,12)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s10.QueryP(156,&value)
	log("p28 r:%d", value)
	countreads(value,28,10)
	l.Unlock()
}


func p29() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	s2.QueryP(20,&value)
	log("p29 r:%d", value)
	countreads(value,29,2)
	l.Unlock()

	delay()
	l.Lock()
	value = 12
	s12.Put(189,value)
	log("p29 w:%d", value)
	countwrites(value,29,12)
	l.Unlock()

	delay()
	l.Lock()
	value = 30
	s30.Put(466,value)
	log("p29 w:%d", value)
	countwrites(value,29,30)
	l.Unlock()

	delay()
	l.Lock()
	value = 16
	s16.Put(241,value)
	log("p29 w:%d", value)
	countwrites(value,29,16)
	l.Unlock()

	delay()
	l.Lock()
	value = 14
	s14.Put(218,value)
	log("p29 w:%d", value)
	countwrites(value,29,14)
	l.Unlock()

	delay()
	l.Lock()
	value = 17
	s17.Put(260,value)
	log("p29 w:%d", value)
	countwrites(value,29,17)
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
	value = 26
	s26.Put(416,value)
	log("p29 w:%d", value)
	countwrites(value,29,26)
	l.Unlock()

	delay()
	l.Lock()
	value = 23
	s23.Put(356,value)
	log("p29 w:%d", value)
	countwrites(value,29,23)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s23.QueryP(365,&value)
	log("p29 r:%d", value)
	countreads(value,29,23)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s31.QueryP(485,&value)
	log("p29 r:%d", value)
	countreads(value,29,31)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s1.QueryP(14,&value)
	log("p29 r:%d", value)
	countreads(value,29,1)
	l.Unlock()

	delay()
	l.Lock()
	value = 32
	s32.Put(512,value)
	log("p29 w:%d", value)
	countwrites(value,29,32)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s3.QueryP(44,&value)
	log("p29 r:%d", value)
	countreads(value,29,3)
	l.Unlock()

	delay()
	l.Lock()
	value = 32
	s32.Put(506,value)
	log("p29 w:%d", value)
	countwrites(value,29,32)
	l.Unlock()

	delay()
	l.Lock()
	value = 19
	s19.Put(300,value)
	log("p29 w:%d", value)
	countwrites(value,29,19)
	l.Unlock()
}


func p30() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	s8.QueryP(113,&value)
	log("p30 r:%d", value)
	countreads(value,30,8)
	l.Unlock()

	delay()
	l.Lock()
	value = 31
	s31.Put(489,value)
	log("p30 w:%d", value)
	countwrites(value,30,31)
	l.Unlock()

	delay()
	l.Lock()
	value = 2
	s2.Put(23,value)
	log("p30 w:%d", value)
	countwrites(value,30,2)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s1.QueryP(4,&value)
	log("p30 r:%d", value)
	countreads(value,30,1)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s15.QueryP(236,&value)
	log("p30 r:%d", value)
	countreads(value,30,15)
	l.Unlock()

	delay()
	l.Lock()
	value = 29
	s29.Put(457,value)
	log("p30 w:%d", value)
	countwrites(value,30,29)
	l.Unlock()

	delay()
	l.Lock()
	value = 23
	s23.Put(363,value)
	log("p30 w:%d", value)
	countwrites(value,30,23)
	l.Unlock()

	delay()
	l.Lock()
	value = 26
	s26.Put(409,value)
	log("p30 w:%d", value)
	countwrites(value,30,26)
	l.Unlock()

	delay()
	l.Lock()
	value = 12
	s12.Put(191,value)
	log("p30 w:%d", value)
	countwrites(value,30,12)
	l.Unlock()

	delay()
	l.Lock()
	value = 25
	s25.Put(391,value)
	log("p30 w:%d", value)
	countwrites(value,30,25)
	l.Unlock()

	delay()
	l.Lock()
	value = 4
	s4.Put(56,value)
	log("p30 w:%d", value)
	countwrites(value,30,4)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s32.QueryP(512,&value)
	log("p30 r:%d", value)
	countreads(value,30,32)
	l.Unlock()

	delay()
	l.Lock()
	value = 10
	s10.Put(155,value)
	log("p30 w:%d", value)
	countwrites(value,30,10)
	l.Unlock()

	delay()
	l.Lock()
	value = 7
	s7.Put(107,value)
	log("p30 w:%d", value)
	countwrites(value,30,7)
	l.Unlock()

	delay()
	l.Lock()
	value = 26
	s26.Put(401,value)
	log("p30 w:%d", value)
	countwrites(value,30,26)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s7.QueryP(102,&value)
	log("p30 r:%d", value)
	countreads(value,30,7)
	l.Unlock()
}


func p31() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 21
	s21.Put(326,value)
	log("p31 w:%d", value)
	countwrites(value,31,21)
	l.Unlock()

	delay()
	l.Lock()
	value = 8
	s8.Put(125,value)
	log("p31 w:%d", value)
	countwrites(value,31,8)
	l.Unlock()

	delay()
	l.Lock()
	value = 23
	s23.Put(359,value)
	log("p31 w:%d", value)
	countwrites(value,31,23)
	l.Unlock()

	delay()
	l.Lock()
	value = 23
	s23.Put(366,value)
	log("p31 w:%d", value)
	countwrites(value,31,23)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s9.QueryP(140,&value)
	log("p31 r:%d", value)
	countreads(value,31,9)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s23.QueryP(354,&value)
	log("p31 r:%d", value)
	countreads(value,31,23)
	l.Unlock()

	delay()
	l.Lock()
	value = 13
	s13.Put(199,value)
	log("p31 w:%d", value)
	countwrites(value,31,13)
	l.Unlock()

	delay()
	l.Lock()
	value = 30
	s30.Put(467,value)
	log("p31 w:%d", value)
	countwrites(value,31,30)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s14.QueryP(220,&value)
	log("p31 r:%d", value)
	countreads(value,31,14)
	l.Unlock()

	delay()
	l.Lock()
	value = 25
	s25.Put(387,value)
	log("p31 w:%d", value)
	countwrites(value,31,25)
	l.Unlock()

	delay()
	l.Lock()
	value = 2
	s2.Put(31,value)
	log("p31 w:%d", value)
	countwrites(value,31,2)
	l.Unlock()

	delay()
	l.Lock()
	value = 26
	s26.Put(409,value)
	log("p31 w:%d", value)
	countwrites(value,31,26)
	l.Unlock()

	delay()
	l.Lock()
	value = 13
	s13.Put(204,value)
	log("p31 w:%d", value)
	countwrites(value,31,13)
	l.Unlock()

	delay()
	l.Lock()
	value = 11
	s11.Put(166,value)
	log("p31 w:%d", value)
	countwrites(value,31,11)
	l.Unlock()

	delay()
	l.Lock()
	value = 4
	s4.Put(53,value)
	log("p31 w:%d", value)
	countwrites(value,31,4)
	l.Unlock()

	delay()
	l.Lock()
	value = 20
	s20.Put(311,value)
	log("p31 w:%d", value)
	countwrites(value,31,20)
	l.Unlock()
}


func p32() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 9
	s9.Put(140,value)
	log("p32 w:%d", value)
	countwrites(value,32,9)
	l.Unlock()

	delay()
	l.Lock()
	value = 5
	s5.Put(80,value)
	log("p32 w:%d", value)
	countwrites(value,32,5)
	l.Unlock()

	delay()
	l.Lock()
	value = 8
	s8.Put(113,value)
	log("p32 w:%d", value)
	countwrites(value,32,8)
	l.Unlock()

	delay()
	l.Lock()
	value = 19
	s19.Put(299,value)
	log("p32 w:%d", value)
	countwrites(value,32,19)
	l.Unlock()

	delay()
	l.Lock()
	value = 31
	s31.Put(489,value)
	log("p32 w:%d", value)
	countwrites(value,32,31)
	l.Unlock()

	delay()
	l.Lock()
	value = 9
	s9.Put(132,value)
	log("p32 w:%d", value)
	countwrites(value,32,9)
	l.Unlock()

	delay()
	l.Lock()
	value = 28
	s28.Put(440,value)
	log("p32 w:%d", value)
	countwrites(value,32,28)
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
	value = 17
	s17.Put(270,value)
	log("p32 w:%d", value)
	countwrites(value,32,17)
	l.Unlock()

	delay()
	l.Lock()
	value = 9
	s9.Put(130,value)
	log("p32 w:%d", value)
	countwrites(value,32,9)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s25.QueryP(387,&value)
	log("p32 r:%d", value)
	countreads(value,32,25)
	l.Unlock()

	delay()
	l.Lock()
	value = 24
	s24.Put(380,value)
	log("p32 w:%d", value)
	countwrites(value,32,24)
	l.Unlock()

	delay()
	l.Lock()
	value = 24
	s24.Put(379,value)
	log("p32 w:%d", value)
	countwrites(value,32,24)
	l.Unlock()

	delay()
	l.Lock()
	value = 22
	s22.Put(343,value)
	log("p32 w:%d", value)
	countwrites(value,32,22)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s5.QueryP(66,&value)
	log("p32 r:%d", value)
	countreads(value,32,5)
	l.Unlock()

	delay()
	l.Lock()
	value = 17
	s17.Put(266,value)
	log("p32 w:%d", value)
	countwrites(value,32,17)
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
