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
	s14.QueryP(224,&value)
	log("p1 r:%d", value)
	countreads(value,1,14)
	l.Unlock()

	delay()
	l.Lock()
	value = 30
	s30.Put(478,value)
	log("p1 w:%d", value)
	countwrites(value,1,30)
	l.Unlock()

	delay()
	l.Lock()
	value = 24
	s24.Put(372,value)
	log("p1 w:%d", value)
	countwrites(value,1,24)
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
	s26.QueryP(407,&value)
	log("p1 r:%d", value)
	countreads(value,1,26)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s23.QueryP(365,&value)
	log("p1 r:%d", value)
	countreads(value,1,23)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s18.QueryP(274,&value)
	log("p1 r:%d", value)
	countreads(value,1,18)
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
	value = -1
	s18.QueryP(281,&value)
	log("p1 r:%d", value)
	countreads(value,1,18)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s26.QueryP(414,&value)
	log("p1 r:%d", value)
	countreads(value,1,26)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s26.QueryP(412,&value)
	log("p1 r:%d", value)
	countreads(value,1,26)
	l.Unlock()

	delay()
	l.Lock()
	value = 13
	s13.Put(202,value)
	log("p1 w:%d", value)
	countwrites(value,1,13)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s8.QueryP(117,&value)
	log("p1 r:%d", value)
	countreads(value,1,8)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s9.QueryP(132,&value)
	log("p1 r:%d", value)
	countreads(value,1,9)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s5.QueryP(75,&value)
	log("p1 r:%d", value)
	countreads(value,1,5)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s29.QueryP(463,&value)
	log("p1 r:%d", value)
	countreads(value,1,29)
	l.Unlock()
}

func p2() {
	defer wg.Done()
	var value int

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
	s18.QueryP(277,&value)
	log("p2 r:%d", value)
	countreads(value,2,18)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s24.QueryP(383,&value)
	log("p2 r:%d", value)
	countreads(value,2,24)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s3.QueryP(36,&value)
	log("p2 r:%d", value)
	countreads(value,2,3)
	l.Unlock()

	delay()
	l.Lock()
	value = 20
	s20.Put(305,value)
	log("p2 w:%d", value)
	countwrites(value,2,20)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s26.QueryP(415,&value)
	log("p2 r:%d", value)
	countreads(value,2,26)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s3.QueryP(40,&value)
	log("p2 r:%d", value)
	countreads(value,2,3)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s9.QueryP(134,&value)
	log("p2 r:%d", value)
	countreads(value,2,9)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s24.QueryP(370,&value)
	log("p2 r:%d", value)
	countreads(value,2,24)
	l.Unlock()

	delay()
	l.Lock()
	value = 19
	s19.Put(292,value)
	log("p2 w:%d", value)
	countwrites(value,2,19)
	l.Unlock()

	delay()
	l.Lock()
	value = 11
	s11.Put(166,value)
	log("p2 w:%d", value)
	countwrites(value,2,11)
	l.Unlock()

	delay()
	l.Lock()
	value = 10
	s10.Put(149,value)
	log("p2 w:%d", value)
	countwrites(value,2,10)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s28.QueryP(434,&value)
	log("p2 r:%d", value)
	countreads(value,2,28)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s23.QueryP(367,&value)
	log("p2 r:%d", value)
	countreads(value,2,23)
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
	value = 1
	s1.Put(7,value)
	log("p2 w:%d", value)
	countwrites(value,2,1)
	l.Unlock()
}

func p3() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	s31.QueryP(482,&value)
	log("p3 r:%d", value)
	countreads(value,3,31)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s31.QueryP(481,&value)
	log("p3 r:%d", value)
	countreads(value,3,31)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s1.QueryP(16,&value)
	log("p3 r:%d", value)
	countreads(value,3,1)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s24.QueryP(369,&value)
	log("p3 r:%d", value)
	countreads(value,3,24)
	l.Unlock()

	delay()
	l.Lock()
	value = 15
	s15.Put(226,value)
	log("p3 w:%d", value)
	countwrites(value,3,15)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s19.QueryP(292,&value)
	log("p3 r:%d", value)
	countreads(value,3,19)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s26.QueryP(416,&value)
	log("p3 r:%d", value)
	countreads(value,3,26)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s3.QueryP(48,&value)
	log("p3 r:%d", value)
	countreads(value,3,3)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s22.QueryP(345,&value)
	log("p3 r:%d", value)
	countreads(value,3,22)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s23.QueryP(359,&value)
	log("p3 r:%d", value)
	countreads(value,3,23)
	l.Unlock()

	delay()
	l.Lock()
	value = 8
	s8.Put(117,value)
	log("p3 w:%d", value)
	countwrites(value,3,8)
	l.Unlock()

	delay()
	l.Lock()
	value = 11
	s11.Put(174,value)
	log("p3 w:%d", value)
	countwrites(value,3,11)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s17.QueryP(269,&value)
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
	value = -1
	s25.QueryP(388,&value)
	log("p3 r:%d", value)
	countreads(value,3,25)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s30.QueryP(465,&value)
	log("p3 r:%d", value)
	countreads(value,3,30)
	l.Unlock()
}

func p4() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	s13.QueryP(206,&value)
	log("p4 r:%d", value)
	countreads(value,4,13)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s29.QueryP(457,&value)
	log("p4 r:%d", value)
	countreads(value,4,29)
	l.Unlock()

	delay()
	l.Lock()
	value = 24
	s24.Put(372,value)
	log("p4 w:%d", value)
	countwrites(value,4,24)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s31.QueryP(481,&value)
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
	s9.QueryP(130,&value)
	log("p4 r:%d", value)
	countreads(value,4,9)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s15.QueryP(234,&value)
	log("p4 r:%d", value)
	countreads(value,4,15)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s1.QueryP(10,&value)
	log("p4 r:%d", value)
	countreads(value,4,1)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s20.QueryP(312,&value)
	log("p4 r:%d", value)
	countreads(value,4,20)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s21.QueryP(322,&value)
	log("p4 r:%d", value)
	countreads(value,4,21)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s27.QueryP(427,&value)
	log("p4 r:%d", value)
	countreads(value,4,27)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s16.QueryP(246,&value)
	log("p4 r:%d", value)
	countreads(value,4,16)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s4.QueryP(51,&value)
	log("p4 r:%d", value)
	countreads(value,4,4)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s4.QueryP(52,&value)
	log("p4 r:%d", value)
	countreads(value,4,4)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s31.QueryP(493,&value)
	log("p4 r:%d", value)
	countreads(value,4,31)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s12.QueryP(179,&value)
	log("p4 r:%d", value)
	countreads(value,4,12)
	l.Unlock()
}

func p5() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 2
	s2.Put(24,value)
	log("p5 w:%d", value)
	countwrites(value,5,2)
	l.Unlock()

	delay()
	l.Lock()
	value = 8
	s8.Put(118,value)
	log("p5 w:%d", value)
	countwrites(value,5,8)
	l.Unlock()

	delay()
	l.Lock()
	value = 17
	s17.Put(272,value)
	log("p5 w:%d", value)
	countwrites(value,5,17)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s11.QueryP(167,&value)
	log("p5 r:%d", value)
	countreads(value,5,11)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s1.QueryP(4,&value)
	log("p5 r:%d", value)
	countreads(value,5,1)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s18.QueryP(286,&value)
	log("p5 r:%d", value)
	countreads(value,5,18)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s16.QueryP(253,&value)
	log("p5 r:%d", value)
	countreads(value,5,16)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s14.QueryP(222,&value)
	log("p5 r:%d", value)
	countreads(value,5,14)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s16.QueryP(243,&value)
	log("p5 r:%d", value)
	countreads(value,5,16)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s22.QueryP(345,&value)
	log("p5 r:%d", value)
	countreads(value,5,22)
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
	s32.QueryP(503,&value)
	log("p5 r:%d", value)
	countreads(value,5,32)
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
	s1.QueryP(9,&value)
	log("p5 r:%d", value)
	countreads(value,5,1)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s3.QueryP(34,&value)
	log("p5 r:%d", value)
	countreads(value,5,3)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s28.QueryP(436,&value)
	log("p5 r:%d", value)
	countreads(value,5,28)
	l.Unlock()
}

func p6() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	s30.QueryP(467,&value)
	log("p6 r:%d", value)
	countreads(value,6,30)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s15.QueryP(236,&value)
	log("p6 r:%d", value)
	countreads(value,6,15)
	l.Unlock()

	delay()
	l.Lock()
	value = 8
	s8.Put(116,value)
	log("p6 w:%d", value)
	countwrites(value,6,8)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s2.QueryP(20,&value)
	log("p6 r:%d", value)
	countreads(value,6,2)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s5.QueryP(79,&value)
	log("p6 r:%d", value)
	countreads(value,6,5)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s23.QueryP(357,&value)
	log("p6 r:%d", value)
	countreads(value,6,23)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s31.QueryP(493,&value)
	log("p6 r:%d", value)
	countreads(value,6,31)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s25.QueryP(389,&value)
	log("p6 r:%d", value)
	countreads(value,6,25)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s3.QueryP(34,&value)
	log("p6 r:%d", value)
	countreads(value,6,3)
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
	value = -1
	s2.QueryP(26,&value)
	log("p6 r:%d", value)
	countreads(value,6,2)
	l.Unlock()

	delay()
	l.Lock()
	value = 32
	s32.Put(504,value)
	log("p6 w:%d", value)
	countwrites(value,6,32)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s7.QueryP(102,&value)
	log("p6 r:%d", value)
	countreads(value,6,7)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s28.QueryP(442,&value)
	log("p6 r:%d", value)
	countreads(value,6,28)
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
	value = -1
	s14.QueryP(211,&value)
	log("p6 r:%d", value)
	countreads(value,6,14)
	l.Unlock()
}

func p7() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	s6.QueryP(92,&value)
	log("p7 r:%d", value)
	countreads(value,7,6)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s10.QueryP(158,&value)
	log("p7 r:%d", value)
	countreads(value,7,10)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s25.QueryP(393,&value)
	log("p7 r:%d", value)
	countreads(value,7,25)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s24.QueryP(377,&value)
	log("p7 r:%d", value)
	countreads(value,7,24)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s20.QueryP(305,&value)
	log("p7 r:%d", value)
	countreads(value,7,20)
	l.Unlock()

	delay()
	l.Lock()
	value = 28
	s28.Put(438,value)
	log("p7 w:%d", value)
	countwrites(value,7,28)
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
	value = 29
	s29.Put(462,value)
	log("p7 w:%d", value)
	countwrites(value,7,29)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s31.QueryP(487,&value)
	log("p7 r:%d", value)
	countreads(value,7,31)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s9.QueryP(138,&value)
	log("p7 r:%d", value)
	countreads(value,7,9)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s18.QueryP(273,&value)
	log("p7 r:%d", value)
	countreads(value,7,18)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s32.QueryP(505,&value)
	log("p7 r:%d", value)
	countreads(value,7,32)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s25.QueryP(400,&value)
	log("p7 r:%d", value)
	countreads(value,7,25)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s10.QueryP(150,&value)
	log("p7 r:%d", value)
	countreads(value,7,10)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s8.QueryP(123,&value)
	log("p7 r:%d", value)
	countreads(value,7,8)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s26.QueryP(404,&value)
	log("p7 r:%d", value)
	countreads(value,7,26)
	l.Unlock()
}

func p8() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	s9.QueryP(131,&value)
	log("p8 r:%d", value)
	countreads(value,8,9)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s20.QueryP(307,&value)
	log("p8 r:%d", value)
	countreads(value,8,20)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s31.QueryP(491,&value)
	log("p8 r:%d", value)
	countreads(value,8,31)
	l.Unlock()

	delay()
	l.Lock()
	value = 4
	s4.Put(58,value)
	log("p8 w:%d", value)
	countwrites(value,8,4)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s19.QueryP(299,&value)
	log("p8 r:%d", value)
	countreads(value,8,19)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s18.QueryP(287,&value)
	log("p8 r:%d", value)
	countreads(value,8,18)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s3.QueryP(39,&value)
	log("p8 r:%d", value)
	countreads(value,8,3)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s11.QueryP(161,&value)
	log("p8 r:%d", value)
	countreads(value,8,11)
	l.Unlock()

	delay()
	l.Lock()
	value = 29
	s29.Put(462,value)
	log("p8 w:%d", value)
	countwrites(value,8,29)
	l.Unlock()

	delay()
	l.Lock()
	value = 5
	s5.Put(76,value)
	log("p8 w:%d", value)
	countwrites(value,8,5)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s21.QueryP(328,&value)
	log("p8 r:%d", value)
	countreads(value,8,21)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s2.QueryP(25,&value)
	log("p8 r:%d", value)
	countreads(value,8,2)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s23.QueryP(368,&value)
	log("p8 r:%d", value)
	countreads(value,8,23)
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
	s21.QueryP(329,&value)
	log("p8 r:%d", value)
	countreads(value,8,21)
	l.Unlock()

	delay()
	l.Lock()
	value = 31
	s31.Put(485,value)
	log("p8 w:%d", value)
	countwrites(value,8,31)
	l.Unlock()
}

func p9() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 13
	s13.Put(194,value)
	log("p9 w:%d", value)
	countwrites(value,9,13)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s24.QueryP(379,&value)
	log("p9 r:%d", value)
	countreads(value,9,24)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s29.QueryP(451,&value)
	log("p9 r:%d", value)
	countreads(value,9,29)
	l.Unlock()

	delay()
	l.Lock()
	value = 8
	s8.Put(123,value)
	log("p9 w:%d", value)
	countwrites(value,9,8)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s15.QueryP(236,&value)
	log("p9 r:%d", value)
	countreads(value,9,15)
	l.Unlock()

	delay()
	l.Lock()
	value = 6
	s6.Put(86,value)
	log("p9 w:%d", value)
	countwrites(value,9,6)
	l.Unlock()

	delay()
	l.Lock()
	value = 27
	s27.Put(425,value)
	log("p9 w:%d", value)
	countwrites(value,9,27)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s3.QueryP(47,&value)
	log("p9 r:%d", value)
	countreads(value,9,3)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s3.QueryP(38,&value)
	log("p9 r:%d", value)
	countreads(value,9,3)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s20.QueryP(306,&value)
	log("p9 r:%d", value)
	countreads(value,9,20)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s31.QueryP(482,&value)
	log("p9 r:%d", value)
	countreads(value,9,31)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s25.QueryP(387,&value)
	log("p9 r:%d", value)
	countreads(value,9,25)
	l.Unlock()

	delay()
	l.Lock()
	value = 9
	s9.Put(142,value)
	log("p9 w:%d", value)
	countwrites(value,9,9)
	l.Unlock()

	delay()
	l.Lock()
	value = 26
	s26.Put(404,value)
	log("p9 w:%d", value)
	countwrites(value,9,26)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s22.QueryP(349,&value)
	log("p9 r:%d", value)
	countreads(value,9,22)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s11.QueryP(165,&value)
	log("p9 r:%d", value)
	countreads(value,9,11)
	l.Unlock()
}


func p10() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	s4.QueryP(57,&value)
	log("p10 r:%d", value)
	countreads(value,10,4)
	l.Unlock()

	delay()
	l.Lock()
	value = 32
	s32.Put(504,value)
	log("p10 w:%d", value)
	countwrites(value,10,32)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s13.QueryP(198,&value)
	log("p10 r:%d", value)
	countreads(value,10,13)
	l.Unlock()

	delay()
	l.Lock()
	value = 1
	s1.Put(2,value)
	log("p10 w:%d", value)
	countwrites(value,10,1)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s3.QueryP(46,&value)
	log("p10 r:%d", value)
	countreads(value,10,3)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s17.QueryP(259,&value)
	log("p10 r:%d", value)
	countreads(value,10,17)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s27.QueryP(430,&value)
	log("p10 r:%d", value)
	countreads(value,10,27)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s29.QueryP(453,&value)
	log("p10 r:%d", value)
	countreads(value,10,29)
	l.Unlock()

	delay()
	l.Lock()
	value = 20
	s20.Put(314,value)
	log("p10 w:%d", value)
	countwrites(value,10,20)
	l.Unlock()

	delay()
	l.Lock()
	value = 13
	s13.Put(204,value)
	log("p10 w:%d", value)
	countwrites(value,10,13)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s24.QueryP(373,&value)
	log("p10 r:%d", value)
	countreads(value,10,24)
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
	s12.QueryP(183,&value)
	log("p10 r:%d", value)
	countreads(value,10,12)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s12.QueryP(189,&value)
	log("p10 r:%d", value)
	countreads(value,10,12)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s22.QueryP(344,&value)
	log("p10 r:%d", value)
	countreads(value,10,22)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s10.QueryP(157,&value)
	log("p10 r:%d", value)
	countreads(value,10,10)
	l.Unlock()
}


func p11() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	s31.QueryP(489,&value)
	log("p11 r:%d", value)
	countreads(value,11,31)
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
	value = -1
	s21.QueryP(332,&value)
	log("p11 r:%d", value)
	countreads(value,11,21)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s23.QueryP(354,&value)
	log("p11 r:%d", value)
	countreads(value,11,23)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s23.QueryP(360,&value)
	log("p11 r:%d", value)
	countreads(value,11,23)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s14.QueryP(216,&value)
	log("p11 r:%d", value)
	countreads(value,11,14)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s24.QueryP(382,&value)
	log("p11 r:%d", value)
	countreads(value,11,24)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s18.QueryP(287,&value)
	log("p11 r:%d", value)
	countreads(value,11,18)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s4.QueryP(60,&value)
	log("p11 r:%d", value)
	countreads(value,11,4)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s2.QueryP(18,&value)
	log("p11 r:%d", value)
	countreads(value,11,2)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s32.QueryP(503,&value)
	log("p11 r:%d", value)
	countreads(value,11,32)
	l.Unlock()

	delay()
	l.Lock()
	value = 7
	s7.Put(100,value)
	log("p11 w:%d", value)
	countwrites(value,11,7)
	l.Unlock()

	delay()
	l.Lock()
	value = 12
	s12.Put(185,value)
	log("p11 w:%d", value)
	countwrites(value,11,12)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s21.QueryP(327,&value)
	log("p11 r:%d", value)
	countreads(value,11,21)
	l.Unlock()

	delay()
	l.Lock()
	value = 8
	s8.Put(125,value)
	log("p11 w:%d", value)
	countwrites(value,11,8)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s24.QueryP(381,&value)
	log("p11 r:%d", value)
	countreads(value,11,24)
	l.Unlock()
}


func p12() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	s29.QueryP(464,&value)
	log("p12 r:%d", value)
	countreads(value,12,29)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s15.QueryP(234,&value)
	log("p12 r:%d", value)
	countreads(value,12,15)
	l.Unlock()

	delay()
	l.Lock()
	value = 21
	s21.Put(321,value)
	log("p12 w:%d", value)
	countwrites(value,12,21)
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
	value = 17
	s17.Put(265,value)
	log("p12 w:%d", value)
	countwrites(value,12,17)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s28.QueryP(445,&value)
	log("p12 r:%d", value)
	countreads(value,12,28)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s14.QueryP(211,&value)
	log("p12 r:%d", value)
	countreads(value,12,14)
	l.Unlock()

	delay()
	l.Lock()
	value = 28
	s28.Put(444,value)
	log("p12 w:%d", value)
	countwrites(value,12,28)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s14.QueryP(212,&value)
	log("p12 r:%d", value)
	countreads(value,12,14)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s28.QueryP(448,&value)
	log("p12 r:%d", value)
	countreads(value,12,28)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s16.QueryP(253,&value)
	log("p12 r:%d", value)
	countreads(value,12,16)
	l.Unlock()

	delay()
	l.Lock()
	value = 10
	s10.Put(147,value)
	log("p12 w:%d", value)
	countwrites(value,12,10)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s8.QueryP(114,&value)
	log("p12 r:%d", value)
	countreads(value,12,8)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s10.QueryP(160,&value)
	log("p12 r:%d", value)
	countreads(value,12,10)
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
	s12.QueryP(192,&value)
	log("p12 r:%d", value)
	countreads(value,12,12)
	l.Unlock()
}


func p13() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 21
	s21.Put(323,value)
	log("p13 w:%d", value)
	countwrites(value,13,21)
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
	s20.QueryP(308,&value)
	log("p13 r:%d", value)
	countreads(value,13,20)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s10.QueryP(154,&value)
	log("p13 r:%d", value)
	countreads(value,13,10)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s13.QueryP(203,&value)
	log("p13 r:%d", value)
	countreads(value,13,13)
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
	value = -1
	s32.QueryP(501,&value)
	log("p13 r:%d", value)
	countreads(value,13,32)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s23.QueryP(357,&value)
	log("p13 r:%d", value)
	countreads(value,13,23)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s18.QueryP(282,&value)
	log("p13 r:%d", value)
	countreads(value,13,18)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s12.QueryP(179,&value)
	log("p13 r:%d", value)
	countreads(value,13,12)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s29.QueryP(457,&value)
	log("p13 r:%d", value)
	countreads(value,13,29)
	l.Unlock()

	delay()
	l.Lock()
	value = 7
	s7.Put(103,value)
	log("p13 w:%d", value)
	countwrites(value,13,7)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s32.QueryP(511,&value)
	log("p13 r:%d", value)
	countreads(value,13,32)
	l.Unlock()

	delay()
	l.Lock()
	value = 1
	s1.Put(10,value)
	log("p13 w:%d", value)
	countwrites(value,13,1)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s27.QueryP(417,&value)
	log("p13 r:%d", value)
	countreads(value,13,27)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s22.QueryP(347,&value)
	log("p13 r:%d", value)
	countreads(value,13,22)
	l.Unlock()
}

func p14() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	s25.QueryP(400,&value)
	log("p14 r:%d", value)
	countreads(value,14,25)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s17.QueryP(259,&value)
	log("p14 r:%d", value)
	countreads(value,14,17)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s4.QueryP(59,&value)
	log("p14 r:%d", value)
	countreads(value,14,4)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s20.QueryP(313,&value)
	log("p14 r:%d", value)
	countreads(value,14,20)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s2.QueryP(20,&value)
	log("p14 r:%d", value)
	countreads(value,14,2)
	l.Unlock()

	delay()
	l.Lock()
	value = 21
	s21.Put(327,value)
	log("p14 w:%d", value)
	countwrites(value,14,21)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s5.QueryP(74,&value)
	log("p14 r:%d", value)
	countreads(value,14,5)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s10.QueryP(147,&value)
	log("p14 r:%d", value)
	countreads(value,14,10)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s4.QueryP(49,&value)
	log("p14 r:%d", value)
	countreads(value,14,4)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s11.QueryP(170,&value)
	log("p14 r:%d", value)
	countreads(value,14,11)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s18.QueryP(276,&value)
	log("p14 r:%d", value)
	countreads(value,14,18)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s5.QueryP(66,&value)
	log("p14 r:%d", value)
	countreads(value,14,5)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s22.QueryP(349,&value)
	log("p14 r:%d", value)
	countreads(value,14,22)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s17.QueryP(272,&value)
	log("p14 r:%d", value)
	countreads(value,14,17)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s14.QueryP(209,&value)
	log("p14 r:%d", value)
	countreads(value,14,14)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s18.QueryP(285,&value)
	log("p14 r:%d", value)
	countreads(value,14,18)
	l.Unlock()
}


func p15() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	s30.QueryP(472,&value)
	log("p15 r:%d", value)
	countreads(value,15,30)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s12.QueryP(192,&value)
	log("p15 r:%d", value)
	countreads(value,15,12)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s9.QueryP(132,&value)
	log("p15 r:%d", value)
	countreads(value,15,9)
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
	s10.QueryP(154,&value)
	log("p15 r:%d", value)
	countreads(value,15,10)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s9.QueryP(135,&value)
	log("p15 r:%d", value)
	countreads(value,15,9)
	l.Unlock()

	delay()
	l.Lock()
	value = 23
	s23.Put(362,value)
	log("p15 w:%d", value)
	countwrites(value,15,23)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s23.QueryP(359,&value)
	log("p15 r:%d", value)
	countreads(value,15,23)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s4.QueryP(58,&value)
	log("p15 r:%d", value)
	countreads(value,15,4)
	l.Unlock()

	delay()
	l.Lock()
	value = 21
	s21.Put(335,value)
	log("p15 w:%d", value)
	countwrites(value,15,21)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s12.QueryP(182,&value)
	log("p15 r:%d", value)
	countreads(value,15,12)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s4.QueryP(56,&value)
	log("p15 r:%d", value)
	countreads(value,15,4)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s31.QueryP(493,&value)
	log("p15 r:%d", value)
	countreads(value,15,31)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s26.QueryP(408,&value)
	log("p15 r:%d", value)
	countreads(value,15,26)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s28.QueryP(448,&value)
	log("p15 r:%d", value)
	countreads(value,15,28)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s6.QueryP(84,&value)
	log("p15 r:%d", value)
	countreads(value,15,6)
	l.Unlock()
}


func p16() {
	defer wg.Done()
	var value int

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
	s13.QueryP(196,&value)
	log("p16 r:%d", value)
	countreads(value,16,13)
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
	s5.QueryP(80,&value)
	log("p16 r:%d", value)
	countreads(value,16,5)
	l.Unlock()

	delay()
	l.Lock()
	value = 13
	s13.Put(206,value)
	log("p16 w:%d", value)
	countwrites(value,16,13)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s22.QueryP(346,&value)
	log("p16 r:%d", value)
	countreads(value,16,22)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s20.QueryP(314,&value)
	log("p16 r:%d", value)
	countreads(value,16,20)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s5.QueryP(77,&value)
	log("p16 r:%d", value)
	countreads(value,16,5)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s7.QueryP(98,&value)
	log("p16 r:%d", value)
	countreads(value,16,7)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s27.QueryP(424,&value)
	log("p16 r:%d", value)
	countreads(value,16,27)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s18.QueryP(274,&value)
	log("p16 r:%d", value)
	countreads(value,16,18)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s30.QueryP(473,&value)
	log("p16 r:%d", value)
	countreads(value,16,30)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s21.QueryP(327,&value)
	log("p16 r:%d", value)
	countreads(value,16,21)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s15.QueryP(226,&value)
	log("p16 r:%d", value)
	countreads(value,16,15)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s5.QueryP(69,&value)
	log("p16 r:%d", value)
	countreads(value,16,5)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s30.QueryP(472,&value)
	log("p16 r:%d", value)
	countreads(value,16,30)
	l.Unlock()
}


func p17() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	s14.QueryP(215,&value)
	log("p17 r:%d", value)
	countreads(value,17,14)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s24.QueryP(370,&value)
	log("p17 r:%d", value)
	countreads(value,17,24)
	l.Unlock()

	delay()
	l.Lock()
	value = 22
	s22.Put(346,value)
	log("p17 w:%d", value)
	countwrites(value,17,22)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s15.QueryP(229,&value)
	log("p17 r:%d", value)
	countreads(value,17,15)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s20.QueryP(308,&value)
	log("p17 r:%d", value)
	countreads(value,17,20)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s22.QueryP(347,&value)
	log("p17 r:%d", value)
	countreads(value,17,22)
	l.Unlock()

	delay()
	l.Lock()
	value = 27
	s27.Put(426,value)
	log("p17 w:%d", value)
	countwrites(value,17,27)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s32.QueryP(502,&value)
	log("p17 r:%d", value)
	countreads(value,17,32)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s8.QueryP(123,&value)
	log("p17 r:%d", value)
	countreads(value,17,8)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s5.QueryP(77,&value)
	log("p17 r:%d", value)
	countreads(value,17,5)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s17.QueryP(260,&value)
	log("p17 r:%d", value)
	countreads(value,17,17)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s28.QueryP(437,&value)
	log("p17 r:%d", value)
	countreads(value,17,28)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s20.QueryP(319,&value)
	log("p17 r:%d", value)
	countreads(value,17,20)
	l.Unlock()

	delay()
	l.Lock()
	value = 6
	s6.Put(94,value)
	log("p17 w:%d", value)
	countwrites(value,17,6)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s8.QueryP(126,&value)
	log("p17 r:%d", value)
	countreads(value,17,8)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s17.QueryP(264,&value)
	log("p17 r:%d", value)
	countreads(value,17,17)
	l.Unlock()
}


func p18() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	s25.QueryP(385,&value)
	log("p18 r:%d", value)
	countreads(value,18,25)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s9.QueryP(144,&value)
	log("p18 r:%d", value)
	countreads(value,18,9)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s26.QueryP(401,&value)
	log("p18 r:%d", value)
	countreads(value,18,26)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s2.QueryP(18,&value)
	log("p18 r:%d", value)
	countreads(value,18,2)
	l.Unlock()

	delay()
	l.Lock()
	value = 25
	s25.Put(394,value)
	log("p18 w:%d", value)
	countwrites(value,18,25)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s22.QueryP(339,&value)
	log("p18 r:%d", value)
	countreads(value,18,22)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s23.QueryP(368,&value)
	log("p18 r:%d", value)
	countreads(value,18,23)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s5.QueryP(66,&value)
	log("p18 r:%d", value)
	countreads(value,18,5)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s5.QueryP(68,&value)
	log("p18 r:%d", value)
	countreads(value,18,5)
	l.Unlock()

	delay()
	l.Lock()
	value = 2
	s2.Put(18,value)
	log("p18 w:%d", value)
	countwrites(value,18,2)
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
	value = 2
	s2.Put(26,value)
	log("p18 w:%d", value)
	countwrites(value,18,2)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s15.QueryP(239,&value)
	log("p18 r:%d", value)
	countreads(value,18,15)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s9.QueryP(143,&value)
	log("p18 r:%d", value)
	countreads(value,18,9)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s19.QueryP(293,&value)
	log("p18 r:%d", value)
	countreads(value,18,19)
	l.Unlock()

	delay()
	l.Lock()
	value = 15
	s15.Put(239,value)
	log("p18 w:%d", value)
	countwrites(value,18,15)
	l.Unlock()
}


func p19() {
	defer wg.Done()
	var value int

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
	s24.QueryP(378,&value)
	log("p19 r:%d", value)
	countreads(value,19,24)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s32.QueryP(504,&value)
	log("p19 r:%d", value)
	countreads(value,19,32)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s2.QueryP(28,&value)
	log("p19 r:%d", value)
	countreads(value,19,2)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s17.QueryP(259,&value)
	log("p19 r:%d", value)
	countreads(value,19,17)
	l.Unlock()

	delay()
	l.Lock()
	value = 3
	s3.Put(39,value)
	log("p19 w:%d", value)
	countwrites(value,19,3)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s31.QueryP(493,&value)
	log("p19 r:%d", value)
	countreads(value,19,31)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s19.QueryP(294,&value)
	log("p19 r:%d", value)
	countreads(value,19,19)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s23.QueryP(358,&value)
	log("p19 r:%d", value)
	countreads(value,19,23)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s3.QueryP(33,&value)
	log("p19 r:%d", value)
	countreads(value,19,3)
	l.Unlock()

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
	s1.QueryP(13,&value)
	log("p19 r:%d", value)
	countreads(value,19,1)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s22.QueryP(342,&value)
	log("p19 r:%d", value)
	countreads(value,19,22)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s12.QueryP(184,&value)
	log("p19 r:%d", value)
	countreads(value,19,12)
	l.Unlock()

	delay()
	l.Lock()
	value = 26
	s26.Put(413,value)
	log("p19 w:%d", value)
	countwrites(value,19,26)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s14.QueryP(209,&value)
	log("p19 r:%d", value)
	countreads(value,19,14)
	l.Unlock()
}


func p20() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	s17.QueryP(262,&value)
	log("p20 r:%d", value)
	countreads(value,20,17)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s4.QueryP(52,&value)
	log("p20 r:%d", value)
	countreads(value,20,4)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s18.QueryP(285,&value)
	log("p20 r:%d", value)
	countreads(value,20,18)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s3.QueryP(42,&value)
	log("p20 r:%d", value)
	countreads(value,20,3)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s20.QueryP(314,&value)
	log("p20 r:%d", value)
	countreads(value,20,20)
	l.Unlock()

	delay()
	l.Lock()
	value = 12
	s12.Put(187,value)
	log("p20 w:%d", value)
	countwrites(value,20,12)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s21.QueryP(321,&value)
	log("p20 r:%d", value)
	countreads(value,20,21)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s20.QueryP(313,&value)
	log("p20 r:%d", value)
	countreads(value,20,20)
	l.Unlock()

	delay()
	l.Lock()
	value = 6
	s6.Put(88,value)
	log("p20 w:%d", value)
	countwrites(value,20,6)
	l.Unlock()

	delay()
	l.Lock()
	value = 9
	s9.Put(133,value)
	log("p20 w:%d", value)
	countwrites(value,20,9)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s3.QueryP(35,&value)
	log("p20 r:%d", value)
	countreads(value,20,3)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s17.QueryP(266,&value)
	log("p20 r:%d", value)
	countreads(value,20,17)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s18.QueryP(283,&value)
	log("p20 r:%d", value)
	countreads(value,20,18)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s4.QueryP(51,&value)
	log("p20 r:%d", value)
	countreads(value,20,4)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s11.QueryP(170,&value)
	log("p20 r:%d", value)
	countreads(value,20,11)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s26.QueryP(403,&value)
	log("p20 r:%d", value)
	countreads(value,20,26)
	l.Unlock()
}


func p21() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	s11.QueryP(169,&value)
	log("p21 r:%d", value)
	countreads(value,21,11)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s21.QueryP(324,&value)
	log("p21 r:%d", value)
	countreads(value,21,21)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s10.QueryP(158,&value)
	log("p21 r:%d", value)
	countreads(value,21,10)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s4.QueryP(60,&value)
	log("p21 r:%d", value)
	countreads(value,21,4)
	l.Unlock()

	delay()
	l.Lock()
	value = 12
	s12.Put(179,value)
	log("p21 w:%d", value)
	countwrites(value,21,12)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s20.QueryP(313,&value)
	log("p21 r:%d", value)
	countreads(value,21,20)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s28.QueryP(443,&value)
	log("p21 r:%d", value)
	countreads(value,21,28)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s30.QueryP(471,&value)
	log("p21 r:%d", value)
	countreads(value,21,30)
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
	value = 10
	s10.Put(148,value)
	log("p21 w:%d", value)
	countwrites(value,21,10)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s29.QueryP(462,&value)
	log("p21 r:%d", value)
	countreads(value,21,29)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s3.QueryP(42,&value)
	log("p21 r:%d", value)
	countreads(value,21,3)
	l.Unlock()

	delay()
	l.Lock()
	value = 24
	s24.Put(371,value)
	log("p21 w:%d", value)
	countwrites(value,21,24)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s5.QueryP(73,&value)
	log("p21 r:%d", value)
	countreads(value,21,5)
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
	s12.QueryP(177,&value)
	log("p21 r:%d", value)
	countreads(value,21,12)
	l.Unlock()
}


func p22() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 9
	s9.Put(142,value)
	log("p22 w:%d", value)
	countwrites(value,22,9)
	l.Unlock()

	delay()
	l.Lock()
	value = 12
	s12.Put(185,value)
	log("p22 w:%d", value)
	countwrites(value,22,12)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s3.QueryP(39,&value)
	log("p22 r:%d", value)
	countreads(value,22,3)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s32.QueryP(506,&value)
	log("p22 r:%d", value)
	countreads(value,22,32)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s3.QueryP(48,&value)
	log("p22 r:%d", value)
	countreads(value,22,3)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s2.QueryP(29,&value)
	log("p22 r:%d", value)
	countreads(value,22,2)
	l.Unlock()

	delay()
	l.Lock()
	value = 29
	s29.Put(451,value)
	log("p22 w:%d", value)
	countwrites(value,22,29)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s25.QueryP(396,&value)
	log("p22 r:%d", value)
	countreads(value,22,25)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s5.QueryP(78,&value)
	log("p22 r:%d", value)
	countreads(value,22,5)
	l.Unlock()

	delay()
	l.Lock()
	value = 2
	s2.Put(31,value)
	log("p22 w:%d", value)
	countwrites(value,22,2)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s29.QueryP(451,&value)
	log("p22 r:%d", value)
	countreads(value,22,29)
	l.Unlock()

	delay()
	l.Lock()
	value = 12
	s12.Put(190,value)
	log("p22 w:%d", value)
	countwrites(value,22,12)
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
	s23.QueryP(357,&value)
	log("p22 r:%d", value)
	countreads(value,22,23)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s17.QueryP(262,&value)
	log("p22 r:%d", value)
	countreads(value,22,17)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s31.QueryP(484,&value)
	log("p22 r:%d", value)
	countreads(value,22,31)
	l.Unlock()
}


func p23() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 27
	s27.Put(430,value)
	log("p23 w:%d", value)
	countwrites(value,23,27)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s3.QueryP(48,&value)
	log("p23 r:%d", value)
	countreads(value,23,3)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s17.QueryP(261,&value)
	log("p23 r:%d", value)
	countreads(value,23,17)
	l.Unlock()

	delay()
	l.Lock()
	value = 13
	s13.Put(193,value)
	log("p23 w:%d", value)
	countwrites(value,23,13)
	l.Unlock()

	delay()
	l.Lock()
	value = 14
	s14.Put(224,value)
	log("p23 w:%d", value)
	countwrites(value,23,14)
	l.Unlock()

	delay()
	l.Lock()
	value = 8
	s8.Put(114,value)
	log("p23 w:%d", value)
	countwrites(value,23,8)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s7.QueryP(112,&value)
	log("p23 r:%d", value)
	countreads(value,23,7)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s11.QueryP(170,&value)
	log("p23 r:%d", value)
	countreads(value,23,11)
	l.Unlock()

	delay()
	l.Lock()
	value = 14
	s14.Put(214,value)
	log("p23 w:%d", value)
	countwrites(value,23,14)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s2.QueryP(20,&value)
	log("p23 r:%d", value)
	countreads(value,23,2)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s20.QueryP(317,&value)
	log("p23 r:%d", value)
	countreads(value,23,20)
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
	value = 22
	s22.Put(344,value)
	log("p23 w:%d", value)
	countwrites(value,23,22)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s13.QueryP(201,&value)
	log("p23 r:%d", value)
	countreads(value,23,13)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s3.QueryP(44,&value)
	log("p23 r:%d", value)
	countreads(value,23,3)
	l.Unlock()

	delay()
	l.Lock()
	value = 13
	s13.Put(200,value)
	log("p23 w:%d", value)
	countwrites(value,23,13)
	l.Unlock()
}


func p24() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	s17.QueryP(263,&value)
	log("p24 r:%d", value)
	countreads(value,24,17)
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
	value = -1
	s9.QueryP(140,&value)
	log("p24 r:%d", value)
	countreads(value,24,9)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s24.QueryP(379,&value)
	log("p24 r:%d", value)
	countreads(value,24,24)
	l.Unlock()

	delay()
	l.Lock()
	value = 1
	s1.Put(13,value)
	log("p24 w:%d", value)
	countwrites(value,24,1)
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
	s32.QueryP(508,&value)
	log("p24 r:%d", value)
	countreads(value,24,32)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s29.QueryP(455,&value)
	log("p24 r:%d", value)
	countreads(value,24,29)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s29.QueryP(454,&value)
	log("p24 r:%d", value)
	countreads(value,24,29)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s32.QueryP(499,&value)
	log("p24 r:%d", value)
	countreads(value,24,32)
	l.Unlock()

	delay()
	l.Lock()
	value = 20
	s20.Put(314,value)
	log("p24 w:%d", value)
	countwrites(value,24,20)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s23.QueryP(357,&value)
	log("p24 r:%d", value)
	countreads(value,24,23)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s16.QueryP(247,&value)
	log("p24 r:%d", value)
	countreads(value,24,16)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s8.QueryP(117,&value)
	log("p24 r:%d", value)
	countreads(value,24,8)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s5.QueryP(69,&value)
	log("p24 r:%d", value)
	countreads(value,24,5)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s4.QueryP(59,&value)
	log("p24 r:%d", value)
	countreads(value,24,4)
	l.Unlock()
}


func p25() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	s29.QueryP(452,&value)
	log("p25 r:%d", value)
	countreads(value,25,29)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s4.QueryP(57,&value)
	log("p25 r:%d", value)
	countreads(value,25,4)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s24.QueryP(374,&value)
	log("p25 r:%d", value)
	countreads(value,25,24)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s32.QueryP(497,&value)
	log("p25 r:%d", value)
	countreads(value,25,32)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s14.QueryP(222,&value)
	log("p25 r:%d", value)
	countreads(value,25,14)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s8.QueryP(119,&value)
	log("p25 r:%d", value)
	countreads(value,25,8)
	l.Unlock()

	delay()
	l.Lock()
	value = 26
	s26.Put(403,value)
	log("p25 w:%d", value)
	countwrites(value,25,26)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s29.QueryP(453,&value)
	log("p25 r:%d", value)
	countreads(value,25,29)
	l.Unlock()

	delay()
	l.Lock()
	value = 7
	s7.Put(100,value)
	log("p25 w:%d", value)
	countwrites(value,25,7)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s24.QueryP(372,&value)
	log("p25 r:%d", value)
	countreads(value,25,24)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s31.QueryP(487,&value)
	log("p25 r:%d", value)
	countreads(value,25,31)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s14.QueryP(210,&value)
	log("p25 r:%d", value)
	countreads(value,25,14)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s7.QueryP(102,&value)
	log("p25 r:%d", value)
	countreads(value,25,7)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s30.QueryP(473,&value)
	log("p25 r:%d", value)
	countreads(value,25,30)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s21.QueryP(326,&value)
	log("p25 r:%d", value)
	countreads(value,25,21)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s10.QueryP(145,&value)
	log("p25 r:%d", value)
	countreads(value,25,10)
	l.Unlock()
}


func p26() {
	defer wg.Done()
	var value int

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
	s29.QueryP(458,&value)
	log("p26 r:%d", value)
	countreads(value,26,29)
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
	s8.QueryP(123,&value)
	log("p26 r:%d", value)
	countreads(value,26,8)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s23.QueryP(356,&value)
	log("p26 r:%d", value)
	countreads(value,26,23)
	l.Unlock()

	delay()
	l.Lock()
	value = 7
	s7.Put(99,value)
	log("p26 w:%d", value)
	countwrites(value,26,7)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s30.QueryP(466,&value)
	log("p26 r:%d", value)
	countreads(value,26,30)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s17.QueryP(258,&value)
	log("p26 r:%d", value)
	countreads(value,26,17)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s28.QueryP(435,&value)
	log("p26 r:%d", value)
	countreads(value,26,28)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s13.QueryP(208,&value)
	log("p26 r:%d", value)
	countreads(value,26,13)
	l.Unlock()

	delay()
	l.Lock()
	value = 1
	s1.Put(5,value)
	log("p26 w:%d", value)
	countwrites(value,26,1)
	l.Unlock()

	delay()
	l.Lock()
	value = 16
	s16.Put(246,value)
	log("p26 w:%d", value)
	countwrites(value,26,16)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s23.QueryP(358,&value)
	log("p26 r:%d", value)
	countreads(value,26,23)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s1.QueryP(13,&value)
	log("p26 r:%d", value)
	countreads(value,26,1)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s22.QueryP(350,&value)
	log("p26 r:%d", value)
	countreads(value,26,22)
	l.Unlock()

	delay()
	l.Lock()
	value = 1
	s1.Put(10,value)
	log("p26 w:%d", value)
	countwrites(value,26,1)
	l.Unlock()
}


func p27() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 13
	s13.Put(200,value)
	log("p27 w:%d", value)
	countwrites(value,27,13)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s21.QueryP(326,&value)
	log("p27 r:%d", value)
	countreads(value,27,21)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s31.QueryP(486,&value)
	log("p27 r:%d", value)
	countreads(value,27,31)
	l.Unlock()

	delay()
	l.Lock()
	value = 2
	s2.Put(24,value)
	log("p27 w:%d", value)
	countwrites(value,27,2)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s14.QueryP(215,&value)
	log("p27 r:%d", value)
	countreads(value,27,14)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s11.QueryP(169,&value)
	log("p27 r:%d", value)
	countreads(value,27,11)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s30.QueryP(478,&value)
	log("p27 r:%d", value)
	countreads(value,27,30)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s15.QueryP(229,&value)
	log("p27 r:%d", value)
	countreads(value,27,15)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s30.QueryP(474,&value)
	log("p27 r:%d", value)
	countreads(value,27,30)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s20.QueryP(310,&value)
	log("p27 r:%d", value)
	countreads(value,27,20)
	l.Unlock()

	delay()
	l.Lock()
	value = 9
	s9.Put(141,value)
	log("p27 w:%d", value)
	countwrites(value,27,9)
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
	s15.QueryP(225,&value)
	log("p27 r:%d", value)
	countreads(value,27,15)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s31.QueryP(481,&value)
	log("p27 r:%d", value)
	countreads(value,27,31)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s12.QueryP(186,&value)
	log("p27 r:%d", value)
	countreads(value,27,12)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s29.QueryP(463,&value)
	log("p27 r:%d", value)
	countreads(value,27,29)
	l.Unlock()
}


func p28() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 30
	s30.Put(477,value)
	log("p28 w:%d", value)
	countwrites(value,28,30)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s18.QueryP(278,&value)
	log("p28 r:%d", value)
	countreads(value,28,18)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s4.QueryP(58,&value)
	log("p28 r:%d", value)
	countreads(value,28,4)
	l.Unlock()

	delay()
	l.Lock()
	value = 20
	s20.Put(311,value)
	log("p28 w:%d", value)
	countwrites(value,28,20)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s19.QueryP(293,&value)
	log("p28 r:%d", value)
	countreads(value,28,19)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s17.QueryP(266,&value)
	log("p28 r:%d", value)
	countreads(value,28,17)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s18.QueryP(277,&value)
	log("p28 r:%d", value)
	countreads(value,28,18)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s2.QueryP(28,&value)
	log("p28 r:%d", value)
	countreads(value,28,2)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s7.QueryP(97,&value)
	log("p28 r:%d", value)
	countreads(value,28,7)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s4.QueryP(55,&value)
	log("p28 r:%d", value)
	countreads(value,28,4)
	l.Unlock()

	delay()
	l.Lock()
	value = 4
	s4.Put(58,value)
	log("p28 w:%d", value)
	countwrites(value,28,4)
	l.Unlock()

	delay()
	l.Lock()
	value = 19
	s19.Put(300,value)
	log("p28 w:%d", value)
	countwrites(value,28,19)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s2.QueryP(29,&value)
	log("p28 r:%d", value)
	countreads(value,28,2)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s4.QueryP(63,&value)
	log("p28 r:%d", value)
	countreads(value,28,4)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s28.QueryP(435,&value)
	log("p28 r:%d", value)
	countreads(value,28,28)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s7.QueryP(100,&value)
	log("p28 r:%d", value)
	countreads(value,28,7)
	l.Unlock()
}


func p29() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	s11.QueryP(166,&value)
	log("p29 r:%d", value)
	countreads(value,29,11)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s26.QueryP(405,&value)
	log("p29 r:%d", value)
	countreads(value,29,26)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s24.QueryP(373,&value)
	log("p29 r:%d", value)
	countreads(value,29,24)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s28.QueryP(439,&value)
	log("p29 r:%d", value)
	countreads(value,29,28)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s25.QueryP(393,&value)
	log("p29 r:%d", value)
	countreads(value,29,25)
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
	s23.QueryP(359,&value)
	log("p29 r:%d", value)
	countreads(value,29,23)
	l.Unlock()

	delay()
	l.Lock()
	value = 1
	s1.Put(7,value)
	log("p29 w:%d", value)
	countwrites(value,29,1)
	l.Unlock()

	delay()
	l.Lock()
	value = 6
	s6.Put(89,value)
	log("p29 w:%d", value)
	countwrites(value,29,6)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s3.QueryP(39,&value)
	log("p29 r:%d", value)
	countreads(value,29,3)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s1.QueryP(5,&value)
	log("p29 r:%d", value)
	countreads(value,29,1)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s5.QueryP(70,&value)
	log("p29 r:%d", value)
	countreads(value,29,5)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s26.QueryP(415,&value)
	log("p29 r:%d", value)
	countreads(value,29,26)
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
	s24.QueryP(371,&value)
	log("p29 r:%d", value)
	countreads(value,29,24)
	l.Unlock()

	delay()
	l.Lock()
	value = 18
	s18.Put(277,value)
	log("p29 w:%d", value)
	countwrites(value,29,18)
	l.Unlock()
}


func p30() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	s28.QueryP(443,&value)
	log("p30 r:%d", value)
	countreads(value,30,28)
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
	value = 20
	s20.Put(316,value)
	log("p30 w:%d", value)
	countwrites(value,30,20)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s25.QueryP(385,&value)
	log("p30 r:%d", value)
	countreads(value,30,25)
	l.Unlock()

	delay()
	l.Lock()
	value = 30
	s30.Put(474,value)
	log("p30 w:%d", value)
	countwrites(value,30,30)
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
	s17.QueryP(264,&value)
	log("p30 r:%d", value)
	countreads(value,30,17)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s25.QueryP(397,&value)
	log("p30 r:%d", value)
	countreads(value,30,25)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s5.QueryP(66,&value)
	log("p30 r:%d", value)
	countreads(value,30,5)
	l.Unlock()

	delay()
	l.Lock()
	value = 14
	s14.Put(215,value)
	log("p30 w:%d", value)
	countwrites(value,30,14)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s13.QueryP(195,&value)
	log("p30 r:%d", value)
	countreads(value,30,13)
	l.Unlock()

	delay()
	l.Lock()
	value = 22
	s22.Put(339,value)
	log("p30 w:%d", value)
	countwrites(value,30,22)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s21.QueryP(322,&value)
	log("p30 r:%d", value)
	countreads(value,30,21)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s8.QueryP(123,&value)
	log("p30 r:%d", value)
	countreads(value,30,8)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s30.QueryP(479,&value)
	log("p30 r:%d", value)
	countreads(value,30,30)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s21.QueryP(327,&value)
	log("p30 r:%d", value)
	countreads(value,30,21)
	l.Unlock()
}


func p31() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	s9.QueryP(141,&value)
	log("p31 r:%d", value)
	countreads(value,31,9)
	l.Unlock()

	delay()
	l.Lock()
	value = 19
	s19.Put(298,value)
	log("p31 w:%d", value)
	countwrites(value,31,19)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s25.QueryP(393,&value)
	log("p31 r:%d", value)
	countreads(value,31,25)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s29.QueryP(458,&value)
	log("p31 r:%d", value)
	countreads(value,31,29)
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
	value = -1
	s23.QueryP(361,&value)
	log("p31 r:%d", value)
	countreads(value,31,23)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s1.QueryP(4,&value)
	log("p31 r:%d", value)
	countreads(value,31,1)
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
	s6.QueryP(91,&value)
	log("p31 r:%d", value)
	countreads(value,31,6)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s8.QueryP(125,&value)
	log("p31 r:%d", value)
	countreads(value,31,8)
	l.Unlock()

	delay()
	l.Lock()
	value = 22
	s22.Put(347,value)
	log("p31 w:%d", value)
	countwrites(value,31,22)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s18.QueryP(280,&value)
	log("p31 r:%d", value)
	countreads(value,31,18)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s3.QueryP(37,&value)
	log("p31 r:%d", value)
	countreads(value,31,3)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s21.QueryP(324,&value)
	log("p31 r:%d", value)
	countreads(value,31,21)
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
	s26.QueryP(416,&value)
	log("p31 r:%d", value)
	countreads(value,31,26)
	l.Unlock()
}


func p32() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 26
	s26.Put(407,value)
	log("p32 w:%d", value)
	countwrites(value,32,26)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s23.QueryP(365,&value)
	log("p32 r:%d", value)
	countreads(value,32,23)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s4.QueryP(50,&value)
	log("p32 r:%d", value)
	countreads(value,32,4)
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
	value = -1
	s20.QueryP(305,&value)
	log("p32 r:%d", value)
	countreads(value,32,20)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s30.QueryP(470,&value)
	log("p32 r:%d", value)
	countreads(value,32,30)
	l.Unlock()

	delay()
	l.Lock()
	value = 23
	s23.Put(353,value)
	log("p32 w:%d", value)
	countwrites(value,32,23)
	l.Unlock()

	delay()
	l.Lock()
	value = 32
	s32.Put(502,value)
	log("p32 w:%d", value)
	countwrites(value,32,32)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s9.QueryP(143,&value)
	log("p32 r:%d", value)
	countreads(value,32,9)
	l.Unlock()

	delay()
	l.Lock()
	value = 29
	s29.Put(457,value)
	log("p32 w:%d", value)
	countwrites(value,32,29)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s5.QueryP(71,&value)
	log("p32 r:%d", value)
	countreads(value,32,5)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s18.QueryP(277,&value)
	log("p32 r:%d", value)
	countreads(value,32,18)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s27.QueryP(417,&value)
	log("p32 r:%d", value)
	countreads(value,32,27)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s21.QueryP(326,&value)
	log("p32 r:%d", value)
	countreads(value,32,21)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s13.QueryP(198,&value)
	log("p32 r:%d", value)
	countreads(value,32,13)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s7.QueryP(109,&value)
	log("p32 r:%d", value)
	countreads(value,32,7)
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
