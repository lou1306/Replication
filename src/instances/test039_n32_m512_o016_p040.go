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
	s1.QueryP(9,&value)
	log("p1 r:%d", value)
	countreads(value,1,1)
	l.Unlock()

	delay()
	l.Lock()
	value = 25
	s25.Put(398,value)
	log("p1 w:%d", value)
	countwrites(value,1,25)
	l.Unlock()

	delay()
	l.Lock()
	value = 10
	s10.Put(159,value)
	log("p1 w:%d", value)
	countwrites(value,1,10)
	l.Unlock()

	delay()
	l.Lock()
	value = 24
	s24.Put(384,value)
	log("p1 w:%d", value)
	countwrites(value,1,24)
	l.Unlock()

	delay()
	l.Lock()
	value = 27
	s27.Put(429,value)
	log("p1 w:%d", value)
	countwrites(value,1,27)
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
	s19.QueryP(303,&value)
	log("p1 r:%d", value)
	countreads(value,1,19)
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
	value = -1
	s30.QueryP(478,&value)
	log("p1 r:%d", value)
	countreads(value,1,30)
	l.Unlock()

	delay()
	l.Lock()
	value = 2
	s2.Put(17,value)
	log("p1 w:%d", value)
	countwrites(value,1,2)
	l.Unlock()

	delay()
	l.Lock()
	value = 6
	s6.Put(87,value)
	log("p1 w:%d", value)
	countwrites(value,1,6)
	l.Unlock()

	delay()
	l.Lock()
	value = 16
	s16.Put(256,value)
	log("p1 w:%d", value)
	countwrites(value,1,16)
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
	value = 10
	s10.Put(145,value)
	log("p1 w:%d", value)
	countwrites(value,1,10)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s23.QueryP(366,&value)
	log("p1 r:%d", value)
	countreads(value,1,23)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s11.QueryP(172,&value)
	log("p1 r:%d", value)
	countreads(value,1,11)
	l.Unlock()
}

func p2() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 29
	s29.Put(449,value)
	log("p2 w:%d", value)
	countwrites(value,2,29)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s5.QueryP(65,&value)
	log("p2 r:%d", value)
	countreads(value,2,5)
	l.Unlock()

	delay()
	l.Lock()
	value = 3
	s3.Put(46,value)
	log("p2 w:%d", value)
	countwrites(value,2,3)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s4.QueryP(54,&value)
	log("p2 r:%d", value)
	countreads(value,2,4)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s32.QueryP(508,&value)
	log("p2 r:%d", value)
	countreads(value,2,32)
	l.Unlock()

	delay()
	l.Lock()
	value = 29
	s29.Put(461,value)
	log("p2 w:%d", value)
	countwrites(value,2,29)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s12.QueryP(178,&value)
	log("p2 r:%d", value)
	countreads(value,2,12)
	l.Unlock()

	delay()
	l.Lock()
	value = 5
	s5.Put(71,value)
	log("p2 w:%d", value)
	countwrites(value,2,5)
	l.Unlock()

	delay()
	l.Lock()
	value = 4
	s4.Put(59,value)
	log("p2 w:%d", value)
	countwrites(value,2,4)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s8.QueryP(122,&value)
	log("p2 r:%d", value)
	countreads(value,2,8)
	l.Unlock()

	delay()
	l.Lock()
	value = 24
	s24.Put(383,value)
	log("p2 w:%d", value)
	countwrites(value,2,24)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s19.QueryP(299,&value)
	log("p2 r:%d", value)
	countreads(value,2,19)
	l.Unlock()

	delay()
	l.Lock()
	value = 20
	s20.Put(316,value)
	log("p2 w:%d", value)
	countwrites(value,2,20)
	l.Unlock()

	delay()
	l.Lock()
	value = 13
	s13.Put(206,value)
	log("p2 w:%d", value)
	countwrites(value,2,13)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s14.QueryP(213,&value)
	log("p2 r:%d", value)
	countreads(value,2,14)
	l.Unlock()

	delay()
	l.Lock()
	value = 15
	s15.Put(240,value)
	log("p2 w:%d", value)
	countwrites(value,2,15)
	l.Unlock()
}

func p3() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	s11.QueryP(169,&value)
	log("p3 r:%d", value)
	countreads(value,3,11)
	l.Unlock()

	delay()
	l.Lock()
	value = 29
	s29.Put(464,value)
	log("p3 w:%d", value)
	countwrites(value,3,29)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s10.QueryP(156,&value)
	log("p3 r:%d", value)
	countreads(value,3,10)
	l.Unlock()

	delay()
	l.Lock()
	value = 3
	s3.Put(45,value)
	log("p3 w:%d", value)
	countwrites(value,3,3)
	l.Unlock()

	delay()
	l.Lock()
	value = 21
	s21.Put(336,value)
	log("p3 w:%d", value)
	countwrites(value,3,21)
	l.Unlock()

	delay()
	l.Lock()
	value = 24
	s24.Put(372,value)
	log("p3 w:%d", value)
	countwrites(value,3,24)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s5.QueryP(80,&value)
	log("p3 r:%d", value)
	countreads(value,3,5)
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
	value = 30
	s30.Put(480,value)
	log("p3 w:%d", value)
	countwrites(value,3,30)
	l.Unlock()

	delay()
	l.Lock()
	value = 2
	s2.Put(27,value)
	log("p3 w:%d", value)
	countwrites(value,3,2)
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
	value = 30
	s30.Put(476,value)
	log("p3 w:%d", value)
	countwrites(value,3,30)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s2.QueryP(19,&value)
	log("p3 r:%d", value)
	countreads(value,3,2)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s17.QueryP(267,&value)
	log("p3 r:%d", value)
	countreads(value,3,17)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s28.QueryP(445,&value)
	log("p3 r:%d", value)
	countreads(value,3,28)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s20.QueryP(312,&value)
	log("p3 r:%d", value)
	countreads(value,3,20)
	l.Unlock()
}

func p4() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	s19.QueryP(301,&value)
	log("p4 r:%d", value)
	countreads(value,4,19)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s23.QueryP(355,&value)
	log("p4 r:%d", value)
	countreads(value,4,23)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s20.QueryP(317,&value)
	log("p4 r:%d", value)
	countreads(value,4,20)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s11.QueryP(163,&value)
	log("p4 r:%d", value)
	countreads(value,4,11)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s24.QueryP(383,&value)
	log("p4 r:%d", value)
	countreads(value,4,24)
	l.Unlock()

	delay()
	l.Lock()
	value = 8
	s8.Put(118,value)
	log("p4 w:%d", value)
	countwrites(value,4,8)
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
	value = 11
	s11.Put(167,value)
	log("p4 w:%d", value)
	countwrites(value,4,11)
	l.Unlock()

	delay()
	l.Lock()
	value = 6
	s6.Put(94,value)
	log("p4 w:%d", value)
	countwrites(value,4,6)
	l.Unlock()

	delay()
	l.Lock()
	value = 20
	s20.Put(306,value)
	log("p4 w:%d", value)
	countwrites(value,4,20)
	l.Unlock()

	delay()
	l.Lock()
	value = 24
	s24.Put(380,value)
	log("p4 w:%d", value)
	countwrites(value,4,24)
	l.Unlock()

	delay()
	l.Lock()
	value = 12
	s12.Put(182,value)
	log("p4 w:%d", value)
	countwrites(value,4,12)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s26.QueryP(407,&value)
	log("p4 r:%d", value)
	countreads(value,4,26)
	l.Unlock()

	delay()
	l.Lock()
	value = 17
	s17.Put(267,value)
	log("p4 w:%d", value)
	countwrites(value,4,17)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s4.QueryP(55,&value)
	log("p4 r:%d", value)
	countreads(value,4,4)
	l.Unlock()

	delay()
	l.Lock()
	value = 19
	s19.Put(299,value)
	log("p4 w:%d", value)
	countwrites(value,4,19)
	l.Unlock()
}

func p5() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	s10.QueryP(146,&value)
	log("p5 r:%d", value)
	countreads(value,5,10)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s1.QueryP(2,&value)
	log("p5 r:%d", value)
	countreads(value,5,1)
	l.Unlock()

	delay()
	l.Lock()
	value = 9
	s9.Put(136,value)
	log("p5 w:%d", value)
	countwrites(value,5,9)
	l.Unlock()

	delay()
	l.Lock()
	value = 5
	s5.Put(75,value)
	log("p5 w:%d", value)
	countwrites(value,5,5)
	l.Unlock()

	delay()
	l.Lock()
	value = 25
	s25.Put(387,value)
	log("p5 w:%d", value)
	countwrites(value,5,25)
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
	s27.QueryP(431,&value)
	log("p5 r:%d", value)
	countreads(value,5,27)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s21.QueryP(326,&value)
	log("p5 r:%d", value)
	countreads(value,5,21)
	l.Unlock()

	delay()
	l.Lock()
	value = 25
	s25.Put(400,value)
	log("p5 w:%d", value)
	countwrites(value,5,25)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s1.QueryP(14,&value)
	log("p5 r:%d", value)
	countreads(value,5,1)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s10.QueryP(152,&value)
	log("p5 r:%d", value)
	countreads(value,5,10)
	l.Unlock()

	delay()
	l.Lock()
	value = 13
	s13.Put(202,value)
	log("p5 w:%d", value)
	countwrites(value,5,13)
	l.Unlock()

	delay()
	l.Lock()
	value = 7
	s7.Put(98,value)
	log("p5 w:%d", value)
	countwrites(value,5,7)
	l.Unlock()

	delay()
	l.Lock()
	value = 29
	s29.Put(462,value)
	log("p5 w:%d", value)
	countwrites(value,5,29)
	l.Unlock()

	delay()
	l.Lock()
	value = 1
	s1.Put(10,value)
	log("p5 w:%d", value)
	countwrites(value,5,1)
	l.Unlock()

	delay()
	l.Lock()
	value = 21
	s21.Put(325,value)
	log("p5 w:%d", value)
	countwrites(value,5,21)
	l.Unlock()
}

func p6() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 22
	s22.Put(339,value)
	log("p6 w:%d", value)
	countwrites(value,6,22)
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
	value = 12
	s12.Put(178,value)
	log("p6 w:%d", value)
	countwrites(value,6,12)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s9.QueryP(129,&value)
	log("p6 r:%d", value)
	countreads(value,6,9)
	l.Unlock()

	delay()
	l.Lock()
	value = 2
	s2.Put(21,value)
	log("p6 w:%d", value)
	countwrites(value,6,2)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s17.QueryP(257,&value)
	log("p6 r:%d", value)
	countreads(value,6,17)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s30.QueryP(471,&value)
	log("p6 r:%d", value)
	countreads(value,6,30)
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
	s28.QueryP(446,&value)
	log("p6 r:%d", value)
	countreads(value,6,28)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s6.QueryP(89,&value)
	log("p6 r:%d", value)
	countreads(value,6,6)
	l.Unlock()

	delay()
	l.Lock()
	value = 11
	s11.Put(162,value)
	log("p6 w:%d", value)
	countwrites(value,6,11)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s14.QueryP(220,&value)
	log("p6 r:%d", value)
	countreads(value,6,14)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s14.QueryP(216,&value)
	log("p6 r:%d", value)
	countreads(value,6,14)
	l.Unlock()

	delay()
	l.Lock()
	value = 3
	s3.Put(38,value)
	log("p6 w:%d", value)
	countwrites(value,6,3)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s10.QueryP(145,&value)
	log("p6 r:%d", value)
	countreads(value,6,10)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s20.QueryP(311,&value)
	log("p6 r:%d", value)
	countreads(value,6,20)
	l.Unlock()
}

func p7() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	s24.QueryP(378,&value)
	log("p7 r:%d", value)
	countreads(value,7,24)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s23.QueryP(367,&value)
	log("p7 r:%d", value)
	countreads(value,7,23)
	l.Unlock()

	delay()
	l.Lock()
	value = 9
	s9.Put(136,value)
	log("p7 w:%d", value)
	countwrites(value,7,9)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s29.QueryP(456,&value)
	log("p7 r:%d", value)
	countreads(value,7,29)
	l.Unlock()

	delay()
	l.Lock()
	value = 14
	s14.Put(219,value)
	log("p7 w:%d", value)
	countwrites(value,7,14)
	l.Unlock()

	delay()
	l.Lock()
	value = 13
	s13.Put(200,value)
	log("p7 w:%d", value)
	countwrites(value,7,13)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s17.QueryP(268,&value)
	log("p7 r:%d", value)
	countreads(value,7,17)
	l.Unlock()

	delay()
	l.Lock()
	value = 20
	s20.Put(309,value)
	log("p7 w:%d", value)
	countwrites(value,7,20)
	l.Unlock()

	delay()
	l.Lock()
	value = 8
	s8.Put(126,value)
	log("p7 w:%d", value)
	countwrites(value,7,8)
	l.Unlock()

	delay()
	l.Lock()
	value = 7
	s7.Put(107,value)
	log("p7 w:%d", value)
	countwrites(value,7,7)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s30.QueryP(465,&value)
	log("p7 r:%d", value)
	countreads(value,7,30)
	l.Unlock()

	delay()
	l.Lock()
	value = 7
	s7.Put(102,value)
	log("p7 w:%d", value)
	countwrites(value,7,7)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s12.QueryP(182,&value)
	log("p7 r:%d", value)
	countreads(value,7,12)
	l.Unlock()

	delay()
	l.Lock()
	value = 27
	s27.Put(426,value)
	log("p7 w:%d", value)
	countwrites(value,7,27)
	l.Unlock()

	delay()
	l.Lock()
	value = 2
	s2.Put(25,value)
	log("p7 w:%d", value)
	countwrites(value,7,2)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s26.QueryP(415,&value)
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
	s24.QueryP(373,&value)
	log("p8 r:%d", value)
	countreads(value,8,24)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s23.QueryP(364,&value)
	log("p8 r:%d", value)
	countreads(value,8,23)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s30.QueryP(465,&value)
	log("p8 r:%d", value)
	countreads(value,8,30)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s29.QueryP(461,&value)
	log("p8 r:%d", value)
	countreads(value,8,29)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s5.QueryP(80,&value)
	log("p8 r:%d", value)
	countreads(value,8,5)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s6.QueryP(85,&value)
	log("p8 r:%d", value)
	countreads(value,8,6)
	l.Unlock()

	delay()
	l.Lock()
	value = 17
	s17.Put(267,value)
	log("p8 w:%d", value)
	countwrites(value,8,17)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s1.QueryP(6,&value)
	log("p8 r:%d", value)
	countreads(value,8,1)
	l.Unlock()

	delay()
	l.Lock()
	value = 24
	s24.Put(379,value)
	log("p8 w:%d", value)
	countwrites(value,8,24)
	l.Unlock()

	delay()
	l.Lock()
	value = 7
	s7.Put(99,value)
	log("p8 w:%d", value)
	countwrites(value,8,7)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s14.QueryP(214,&value)
	log("p8 r:%d", value)
	countreads(value,8,14)
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
	s19.QueryP(292,&value)
	log("p8 r:%d", value)
	countreads(value,8,19)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s11.QueryP(172,&value)
	log("p8 r:%d", value)
	countreads(value,8,11)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s21.QueryP(334,&value)
	log("p8 r:%d", value)
	countreads(value,8,21)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s27.QueryP(426,&value)
	log("p8 r:%d", value)
	countreads(value,8,27)
	l.Unlock()
}

func p9() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 14
	s14.Put(209,value)
	log("p9 w:%d", value)
	countwrites(value,9,14)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s21.QueryP(333,&value)
	log("p9 r:%d", value)
	countreads(value,9,21)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s32.QueryP(509,&value)
	log("p9 r:%d", value)
	countreads(value,9,32)
	l.Unlock()

	delay()
	l.Lock()
	value = 22
	s22.Put(350,value)
	log("p9 w:%d", value)
	countwrites(value,9,22)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s7.QueryP(111,&value)
	log("p9 r:%d", value)
	countreads(value,9,7)
	l.Unlock()

	delay()
	l.Lock()
	value = 23
	s23.Put(355,value)
	log("p9 w:%d", value)
	countwrites(value,9,23)
	l.Unlock()

	delay()
	l.Lock()
	value = 31
	s31.Put(485,value)
	log("p9 w:%d", value)
	countwrites(value,9,31)
	l.Unlock()

	delay()
	l.Lock()
	value = 14
	s14.Put(214,value)
	log("p9 w:%d", value)
	countwrites(value,9,14)
	l.Unlock()

	delay()
	l.Lock()
	value = 3
	s3.Put(37,value)
	log("p9 w:%d", value)
	countwrites(value,9,3)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s19.QueryP(302,&value)
	log("p9 r:%d", value)
	countreads(value,9,19)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s14.QueryP(212,&value)
	log("p9 r:%d", value)
	countreads(value,9,14)
	l.Unlock()

	delay()
	l.Lock()
	value = 19
	s19.Put(297,value)
	log("p9 w:%d", value)
	countwrites(value,9,19)
	l.Unlock()

	delay()
	l.Lock()
	value = 1
	s1.Put(16,value)
	log("p9 w:%d", value)
	countwrites(value,9,1)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s19.QueryP(290,&value)
	log("p9 r:%d", value)
	countreads(value,9,19)
	l.Unlock()

	delay()
	l.Lock()
	value = 17
	s17.Put(268,value)
	log("p9 w:%d", value)
	countwrites(value,9,17)
	l.Unlock()

	delay()
	l.Lock()
	value = 32
	s32.Put(500,value)
	log("p9 w:%d", value)
	countwrites(value,9,32)
	l.Unlock()
}


func p10() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	s26.QueryP(402,&value)
	log("p10 r:%d", value)
	countreads(value,10,26)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s17.QueryP(268,&value)
	log("p10 r:%d", value)
	countreads(value,10,17)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s12.QueryP(186,&value)
	log("p10 r:%d", value)
	countreads(value,10,12)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s9.QueryP(131,&value)
	log("p10 r:%d", value)
	countreads(value,10,9)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s8.QueryP(122,&value)
	log("p10 r:%d", value)
	countreads(value,10,8)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s19.QueryP(293,&value)
	log("p10 r:%d", value)
	countreads(value,10,19)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s14.QueryP(216,&value)
	log("p10 r:%d", value)
	countreads(value,10,14)
	l.Unlock()

	delay()
	l.Lock()
	value = 7
	s7.Put(97,value)
	log("p10 w:%d", value)
	countwrites(value,10,7)
	l.Unlock()

	delay()
	l.Lock()
	value = 10
	s10.Put(147,value)
	log("p10 w:%d", value)
	countwrites(value,10,10)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s2.QueryP(23,&value)
	log("p10 r:%d", value)
	countreads(value,10,2)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s19.QueryP(291,&value)
	log("p10 r:%d", value)
	countreads(value,10,19)
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
	value = 6
	s6.Put(90,value)
	log("p10 w:%d", value)
	countwrites(value,10,6)
	l.Unlock()

	delay()
	l.Lock()
	value = 10
	s10.Put(151,value)
	log("p10 w:%d", value)
	countwrites(value,10,10)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s9.QueryP(133,&value)
	log("p10 r:%d", value)
	countreads(value,10,9)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s26.QueryP(401,&value)
	log("p10 r:%d", value)
	countreads(value,10,26)
	l.Unlock()
}


func p11() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 5
	s5.Put(72,value)
	log("p11 w:%d", value)
	countwrites(value,11,5)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s2.QueryP(17,&value)
	log("p11 r:%d", value)
	countreads(value,11,2)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s30.QueryP(468,&value)
	log("p11 r:%d", value)
	countreads(value,11,30)
	l.Unlock()

	delay()
	l.Lock()
	value = 24
	s24.Put(374,value)
	log("p11 w:%d", value)
	countwrites(value,11,24)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s12.QueryP(177,&value)
	log("p11 r:%d", value)
	countreads(value,11,12)
	l.Unlock()

	delay()
	l.Lock()
	value = 22
	s22.Put(350,value)
	log("p11 w:%d", value)
	countwrites(value,11,22)
	l.Unlock()

	delay()
	l.Lock()
	value = 20
	s20.Put(319,value)
	log("p11 w:%d", value)
	countwrites(value,11,20)
	l.Unlock()

	delay()
	l.Lock()
	value = 4
	s4.Put(56,value)
	log("p11 w:%d", value)
	countwrites(value,11,4)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s3.QueryP(46,&value)
	log("p11 r:%d", value)
	countreads(value,11,3)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s16.QueryP(254,&value)
	log("p11 r:%d", value)
	countreads(value,11,16)
	l.Unlock()

	delay()
	l.Lock()
	value = 2
	s2.Put(17,value)
	log("p11 w:%d", value)
	countwrites(value,11,2)
	l.Unlock()

	delay()
	l.Lock()
	value = 31
	s31.Put(489,value)
	log("p11 w:%d", value)
	countwrites(value,11,31)
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
	s2.QueryP(24,&value)
	log("p11 r:%d", value)
	countreads(value,11,2)
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
	value = 6
	s6.Put(94,value)
	log("p11 w:%d", value)
	countwrites(value,11,6)
	l.Unlock()
}


func p12() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 17
	s17.Put(271,value)
	log("p12 w:%d", value)
	countwrites(value,12,17)
	l.Unlock()

	delay()
	l.Lock()
	value = 8
	s8.Put(126,value)
	log("p12 w:%d", value)
	countwrites(value,12,8)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s12.QueryP(179,&value)
	log("p12 r:%d", value)
	countreads(value,12,12)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s21.QueryP(329,&value)
	log("p12 r:%d", value)
	countreads(value,12,21)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s11.QueryP(168,&value)
	log("p12 r:%d", value)
	countreads(value,12,11)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s15.QueryP(227,&value)
	log("p12 r:%d", value)
	countreads(value,12,15)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s11.QueryP(163,&value)
	log("p12 r:%d", value)
	countreads(value,12,11)
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
	s22.QueryP(340,&value)
	log("p12 r:%d", value)
	countreads(value,12,22)
	l.Unlock()

	delay()
	l.Lock()
	value = 2
	s2.Put(23,value)
	log("p12 w:%d", value)
	countwrites(value,12,2)
	l.Unlock()

	delay()
	l.Lock()
	value = 18
	s18.Put(283,value)
	log("p12 w:%d", value)
	countwrites(value,12,18)
	l.Unlock()

	delay()
	l.Lock()
	value = 28
	s28.Put(445,value)
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
	s4.QueryP(55,&value)
	log("p12 r:%d", value)
	countreads(value,12,4)
	l.Unlock()

	delay()
	l.Lock()
	value = 27
	s27.Put(425,value)
	log("p12 w:%d", value)
	countwrites(value,12,27)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s8.QueryP(123,&value)
	log("p12 r:%d", value)
	countreads(value,12,8)
	l.Unlock()
}


func p13() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 32
	s32.Put(503,value)
	log("p13 w:%d", value)
	countwrites(value,13,32)
	l.Unlock()

	delay()
	l.Lock()
	value = 18
	s18.Put(283,value)
	log("p13 w:%d", value)
	countwrites(value,13,18)
	l.Unlock()

	delay()
	l.Lock()
	value = 14
	s14.Put(212,value)
	log("p13 w:%d", value)
	countwrites(value,13,14)
	l.Unlock()

	delay()
	l.Lock()
	value = 14
	s14.Put(220,value)
	log("p13 w:%d", value)
	countwrites(value,13,14)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s11.QueryP(174,&value)
	log("p13 r:%d", value)
	countreads(value,13,11)
	l.Unlock()

	delay()
	l.Lock()
	value = 26
	s26.Put(410,value)
	log("p13 w:%d", value)
	countwrites(value,13,26)
	l.Unlock()

	delay()
	l.Lock()
	value = 5
	s5.Put(71,value)
	log("p13 w:%d", value)
	countwrites(value,13,5)
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
	s10.QueryP(145,&value)
	log("p13 r:%d", value)
	countreads(value,13,10)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s17.QueryP(263,&value)
	log("p13 r:%d", value)
	countreads(value,13,17)
	l.Unlock()

	delay()
	l.Lock()
	value = 1
	s1.Put(3,value)
	log("p13 w:%d", value)
	countwrites(value,13,1)
	l.Unlock()

	delay()
	l.Lock()
	value = 19
	s19.Put(304,value)
	log("p13 w:%d", value)
	countwrites(value,13,19)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s2.QueryP(19,&value)
	log("p13 r:%d", value)
	countreads(value,13,2)
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
	value = -1
	s15.QueryP(229,&value)
	log("p13 r:%d", value)
	countreads(value,13,15)
	l.Unlock()

	delay()
	l.Lock()
	value = 22
	s22.Put(346,value)
	log("p13 w:%d", value)
	countwrites(value,13,22)
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
	value = 1
	s1.Put(13,value)
	log("p14 w:%d", value)
	countwrites(value,14,1)
	l.Unlock()

	delay()
	l.Lock()
	value = 8
	s8.Put(124,value)
	log("p14 w:%d", value)
	countwrites(value,14,8)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s6.QueryP(82,&value)
	log("p14 r:%d", value)
	countreads(value,14,6)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s18.QueryP(288,&value)
	log("p14 r:%d", value)
	countreads(value,14,18)
	l.Unlock()

	delay()
	l.Lock()
	value = 2
	s2.Put(28,value)
	log("p14 w:%d", value)
	countwrites(value,14,2)
	l.Unlock()

	delay()
	l.Lock()
	value = 13
	s13.Put(199,value)
	log("p14 w:%d", value)
	countwrites(value,14,13)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s10.QueryP(153,&value)
	log("p14 r:%d", value)
	countreads(value,14,10)
	l.Unlock()

	delay()
	l.Lock()
	value = 11
	s11.Put(174,value)
	log("p14 w:%d", value)
	countwrites(value,14,11)
	l.Unlock()

	delay()
	l.Lock()
	value = 15
	s15.Put(238,value)
	log("p14 w:%d", value)
	countwrites(value,14,15)
	l.Unlock()

	delay()
	l.Lock()
	value = 1
	s1.Put(15,value)
	log("p14 w:%d", value)
	countwrites(value,14,1)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s9.QueryP(142,&value)
	log("p14 r:%d", value)
	countreads(value,14,9)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s31.QueryP(491,&value)
	log("p14 r:%d", value)
	countreads(value,14,31)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s31.QueryP(496,&value)
	log("p14 r:%d", value)
	countreads(value,14,31)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s22.QueryP(341,&value)
	log("p14 r:%d", value)
	countreads(value,14,22)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s14.QueryP(209,&value)
	log("p14 r:%d", value)
	countreads(value,14,14)
	l.Unlock()
}


func p15() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	s12.QueryP(180,&value)
	log("p15 r:%d", value)
	countreads(value,15,12)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s12.QueryP(184,&value)
	log("p15 r:%d", value)
	countreads(value,15,12)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s20.QueryP(319,&value)
	log("p15 r:%d", value)
	countreads(value,15,20)
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
	s24.QueryP(369,&value)
	log("p15 r:%d", value)
	countreads(value,15,24)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s20.QueryP(305,&value)
	log("p15 r:%d", value)
	countreads(value,15,20)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s16.QueryP(242,&value)
	log("p15 r:%d", value)
	countreads(value,15,16)
	l.Unlock()

	delay()
	l.Lock()
	value = 22
	s22.Put(339,value)
	log("p15 w:%d", value)
	countwrites(value,15,22)
	l.Unlock()

	delay()
	l.Lock()
	value = 27
	s27.Put(430,value)
	log("p15 w:%d", value)
	countwrites(value,15,27)
	l.Unlock()

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
	s12.QueryP(190,&value)
	log("p15 r:%d", value)
	countreads(value,15,12)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s14.QueryP(214,&value)
	log("p15 r:%d", value)
	countreads(value,15,14)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s22.QueryP(345,&value)
	log("p15 r:%d", value)
	countreads(value,15,22)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s16.QueryP(255,&value)
	log("p15 r:%d", value)
	countreads(value,15,16)
	l.Unlock()

	delay()
	l.Lock()
	value = 20
	s20.Put(307,value)
	log("p15 w:%d", value)
	countwrites(value,15,20)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s6.QueryP(85,&value)
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
	s8.QueryP(125,&value)
	log("p16 r:%d", value)
	countreads(value,16,8)
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
	s11.QueryP(174,&value)
	log("p16 r:%d", value)
	countreads(value,16,11)
	l.Unlock()

	delay()
	l.Lock()
	value = 11
	s11.Put(173,value)
	log("p16 w:%d", value)
	countwrites(value,16,11)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s31.QueryP(496,&value)
	log("p16 r:%d", value)
	countreads(value,16,31)
	l.Unlock()

	delay()
	l.Lock()
	value = 29
	s29.Put(450,value)
	log("p16 w:%d", value)
	countwrites(value,16,29)
	l.Unlock()

	delay()
	l.Lock()
	value = 11
	s11.Put(174,value)
	log("p16 w:%d", value)
	countwrites(value,16,11)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s18.QueryP(285,&value)
	log("p16 r:%d", value)
	countreads(value,16,18)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s26.QueryP(411,&value)
	log("p16 r:%d", value)
	countreads(value,16,26)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s8.QueryP(121,&value)
	log("p16 r:%d", value)
	countreads(value,16,8)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s19.QueryP(298,&value)
	log("p16 r:%d", value)
	countreads(value,16,19)
	l.Unlock()

	delay()
	l.Lock()
	value = 3
	s3.Put(41,value)
	log("p16 w:%d", value)
	countwrites(value,16,3)
	l.Unlock()

	delay()
	l.Lock()
	value = 30
	s30.Put(476,value)
	log("p16 w:%d", value)
	countwrites(value,16,30)
	l.Unlock()

	delay()
	l.Lock()
	value = 5
	s5.Put(75,value)
	log("p16 w:%d", value)
	countwrites(value,16,5)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s18.QueryP(275,&value)
	log("p16 r:%d", value)
	countreads(value,16,18)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s9.QueryP(144,&value)
	log("p16 r:%d", value)
	countreads(value,16,9)
	l.Unlock()
}


func p17() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	s14.QueryP(209,&value)
	log("p17 r:%d", value)
	countreads(value,17,14)
	l.Unlock()

	delay()
	l.Lock()
	value = 28
	s28.Put(446,value)
	log("p17 w:%d", value)
	countwrites(value,17,28)
	l.Unlock()

	delay()
	l.Lock()
	value = 15
	s15.Put(238,value)
	log("p17 w:%d", value)
	countwrites(value,17,15)
	l.Unlock()

	delay()
	l.Lock()
	value = 29
	s29.Put(464,value)
	log("p17 w:%d", value)
	countwrites(value,17,29)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s18.QueryP(284,&value)
	log("p17 r:%d", value)
	countreads(value,17,18)
	l.Unlock()

	delay()
	l.Lock()
	value = 5
	s5.Put(80,value)
	log("p17 w:%d", value)
	countwrites(value,17,5)
	l.Unlock()

	delay()
	l.Lock()
	value = 11
	s11.Put(176,value)
	log("p17 w:%d", value)
	countwrites(value,17,11)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s11.QueryP(163,&value)
	log("p17 r:%d", value)
	countreads(value,17,11)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s21.QueryP(328,&value)
	log("p17 r:%d", value)
	countreads(value,17,21)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s22.QueryP(340,&value)
	log("p17 r:%d", value)
	countreads(value,17,22)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s3.QueryP(43,&value)
	log("p17 r:%d", value)
	countreads(value,17,3)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s1.QueryP(8,&value)
	log("p17 r:%d", value)
	countreads(value,17,1)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s18.QueryP(287,&value)
	log("p17 r:%d", value)
	countreads(value,17,18)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s2.QueryP(19,&value)
	log("p17 r:%d", value)
	countreads(value,17,2)
	l.Unlock()

	delay()
	l.Lock()
	value = 29
	s29.Put(452,value)
	log("p17 w:%d", value)
	countwrites(value,17,29)
	l.Unlock()

	delay()
	l.Lock()
	value = 22
	s22.Put(352,value)
	log("p17 w:%d", value)
	countwrites(value,17,22)
	l.Unlock()
}


func p18() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	s19.QueryP(295,&value)
	log("p18 r:%d", value)
	countreads(value,18,19)
	l.Unlock()

	delay()
	l.Lock()
	value = 12
	s12.Put(185,value)
	log("p18 w:%d", value)
	countwrites(value,18,12)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s4.QueryP(61,&value)
	log("p18 r:%d", value)
	countreads(value,18,4)
	l.Unlock()

	delay()
	l.Lock()
	value = 29
	s29.Put(457,value)
	log("p18 w:%d", value)
	countwrites(value,18,29)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s6.QueryP(93,&value)
	log("p18 r:%d", value)
	countreads(value,18,6)
	l.Unlock()

	delay()
	l.Lock()
	value = 27
	s27.Put(423,value)
	log("p18 w:%d", value)
	countwrites(value,18,27)
	l.Unlock()

	delay()
	l.Lock()
	value = 13
	s13.Put(195,value)
	log("p18 w:%d", value)
	countwrites(value,18,13)
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
	s24.QueryP(370,&value)
	log("p18 r:%d", value)
	countreads(value,18,24)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s23.QueryP(357,&value)
	log("p18 r:%d", value)
	countreads(value,18,23)
	l.Unlock()

	delay()
	l.Lock()
	value = 14
	s14.Put(220,value)
	log("p18 w:%d", value)
	countwrites(value,18,14)
	l.Unlock()

	delay()
	l.Lock()
	value = 10
	s10.Put(154,value)
	log("p18 w:%d", value)
	countwrites(value,18,10)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s10.QueryP(156,&value)
	log("p18 r:%d", value)
	countreads(value,18,10)
	l.Unlock()

	delay()
	l.Lock()
	value = 27
	s27.Put(424,value)
	log("p18 w:%d", value)
	countwrites(value,18,27)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s18.QueryP(274,&value)
	log("p18 r:%d", value)
	countreads(value,18,18)
	l.Unlock()

	delay()
	l.Lock()
	value = 18
	s18.Put(276,value)
	log("p18 w:%d", value)
	countwrites(value,18,18)
	l.Unlock()
}


func p19() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 24
	s24.Put(377,value)
	log("p19 w:%d", value)
	countwrites(value,19,24)
	l.Unlock()

	delay()
	l.Lock()
	value = 27
	s27.Put(430,value)
	log("p19 w:%d", value)
	countwrites(value,19,27)
	l.Unlock()

	delay()
	l.Lock()
	value = 22
	s22.Put(349,value)
	log("p19 w:%d", value)
	countwrites(value,19,22)
	l.Unlock()

	delay()
	l.Lock()
	value = 29
	s29.Put(451,value)
	log("p19 w:%d", value)
	countwrites(value,19,29)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s9.QueryP(140,&value)
	log("p19 r:%d", value)
	countreads(value,19,9)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s1.QueryP(11,&value)
	log("p19 r:%d", value)
	countreads(value,19,1)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s24.QueryP(375,&value)
	log("p19 r:%d", value)
	countreads(value,19,24)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s11.QueryP(166,&value)
	log("p19 r:%d", value)
	countreads(value,19,11)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s26.QueryP(407,&value)
	log("p19 r:%d", value)
	countreads(value,19,26)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s19.QueryP(304,&value)
	log("p19 r:%d", value)
	countreads(value,19,19)
	l.Unlock()

	delay()
	l.Lock()
	value = 28
	s28.Put(442,value)
	log("p19 w:%d", value)
	countwrites(value,19,28)
	l.Unlock()

	delay()
	l.Lock()
	value = 9
	s9.Put(135,value)
	log("p19 w:%d", value)
	countwrites(value,19,9)
	l.Unlock()

	delay()
	l.Lock()
	value = 18
	s18.Put(284,value)
	log("p19 w:%d", value)
	countwrites(value,19,18)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s10.QueryP(145,&value)
	log("p19 r:%d", value)
	countreads(value,19,10)
	l.Unlock()

	delay()
	l.Lock()
	value = 14
	s14.Put(219,value)
	log("p19 w:%d", value)
	countwrites(value,19,14)
	l.Unlock()

	delay()
	l.Lock()
	value = 8
	s8.Put(123,value)
	log("p19 w:%d", value)
	countwrites(value,19,8)
	l.Unlock()
}


func p20() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 4
	s4.Put(55,value)
	log("p20 w:%d", value)
	countwrites(value,20,4)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s26.QueryP(415,&value)
	log("p20 r:%d", value)
	countreads(value,20,26)
	l.Unlock()

	delay()
	l.Lock()
	value = 16
	s16.Put(251,value)
	log("p20 w:%d", value)
	countwrites(value,20,16)
	l.Unlock()

	delay()
	l.Lock()
	value = 9
	s9.Put(140,value)
	log("p20 w:%d", value)
	countwrites(value,20,9)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s25.QueryP(392,&value)
	log("p20 r:%d", value)
	countreads(value,20,25)
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
	value = 26
	s26.Put(415,value)
	log("p20 w:%d", value)
	countwrites(value,20,26)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s25.QueryP(394,&value)
	log("p20 r:%d", value)
	countreads(value,20,25)
	l.Unlock()

	delay()
	l.Lock()
	value = 19
	s19.Put(289,value)
	log("p20 w:%d", value)
	countwrites(value,20,19)
	l.Unlock()

	delay()
	l.Lock()
	value = 21
	s21.Put(333,value)
	log("p20 w:%d", value)
	countwrites(value,20,21)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s19.QueryP(297,&value)
	log("p20 r:%d", value)
	countreads(value,20,19)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s28.QueryP(448,&value)
	log("p20 r:%d", value)
	countreads(value,20,28)
	l.Unlock()

	delay()
	l.Lock()
	value = 3
	s3.Put(44,value)
	log("p20 w:%d", value)
	countwrites(value,20,3)
	l.Unlock()

	delay()
	l.Lock()
	value = 28
	s28.Put(439,value)
	log("p20 w:%d", value)
	countwrites(value,20,28)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s4.QueryP(59,&value)
	log("p20 r:%d", value)
	countreads(value,20,4)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s15.QueryP(234,&value)
	log("p20 r:%d", value)
	countreads(value,20,15)
	l.Unlock()
}


func p21() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	s29.QueryP(450,&value)
	log("p21 r:%d", value)
	countreads(value,21,29)
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
	s20.QueryP(313,&value)
	log("p21 r:%d", value)
	countreads(value,21,20)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s28.QueryP(441,&value)
	log("p21 r:%d", value)
	countreads(value,21,28)
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
	value = 6
	s6.Put(82,value)
	log("p21 w:%d", value)
	countwrites(value,21,6)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s14.QueryP(221,&value)
	log("p21 r:%d", value)
	countreads(value,21,14)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s12.QueryP(192,&value)
	log("p21 r:%d", value)
	countreads(value,21,12)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s25.QueryP(398,&value)
	log("p21 r:%d", value)
	countreads(value,21,25)
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
	value = -1
	s32.QueryP(510,&value)
	log("p21 r:%d", value)
	countreads(value,21,32)
	l.Unlock()

	delay()
	l.Lock()
	value = 31
	s31.Put(486,value)
	log("p21 w:%d", value)
	countwrites(value,21,31)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s22.QueryP(352,&value)
	log("p21 r:%d", value)
	countreads(value,21,22)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s2.QueryP(31,&value)
	log("p21 r:%d", value)
	countreads(value,21,2)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s18.QueryP(280,&value)
	log("p21 r:%d", value)
	countreads(value,21,18)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s30.QueryP(473,&value)
	log("p21 r:%d", value)
	countreads(value,21,30)
	l.Unlock()
}


func p22() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	s11.QueryP(174,&value)
	log("p22 r:%d", value)
	countreads(value,22,11)
	l.Unlock()

	delay()
	l.Lock()
	value = 4
	s4.Put(54,value)
	log("p22 w:%d", value)
	countwrites(value,22,4)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s11.QueryP(176,&value)
	log("p22 r:%d", value)
	countreads(value,22,11)
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
	value = 9
	s9.Put(134,value)
	log("p22 w:%d", value)
	countwrites(value,22,9)
	l.Unlock()

	delay()
	l.Lock()
	value = 11
	s11.Put(172,value)
	log("p22 w:%d", value)
	countwrites(value,22,11)
	l.Unlock()

	delay()
	l.Lock()
	value = 17
	s17.Put(258,value)
	log("p22 w:%d", value)
	countwrites(value,22,17)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s29.QueryP(456,&value)
	log("p22 r:%d", value)
	countreads(value,22,29)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s12.QueryP(190,&value)
	log("p22 r:%d", value)
	countreads(value,22,12)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s19.QueryP(297,&value)
	log("p22 r:%d", value)
	countreads(value,22,19)
	l.Unlock()

	delay()
	l.Lock()
	value = 18
	s18.Put(282,value)
	log("p22 w:%d", value)
	countwrites(value,22,18)
	l.Unlock()

	delay()
	l.Lock()
	value = 1
	s1.Put(3,value)
	log("p22 w:%d", value)
	countwrites(value,22,1)
	l.Unlock()

	delay()
	l.Lock()
	value = 20
	s20.Put(306,value)
	log("p22 w:%d", value)
	countwrites(value,22,20)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s6.QueryP(94,&value)
	log("p22 r:%d", value)
	countreads(value,22,6)
	l.Unlock()

	delay()
	l.Lock()
	value = 4
	s4.Put(61,value)
	log("p22 w:%d", value)
	countwrites(value,22,4)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s25.QueryP(390,&value)
	log("p22 r:%d", value)
	countreads(value,22,25)
	l.Unlock()
}


func p23() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	s13.QueryP(204,&value)
	log("p23 r:%d", value)
	countreads(value,23,13)
	l.Unlock()

	delay()
	l.Lock()
	value = 29
	s29.Put(451,value)
	log("p23 w:%d", value)
	countwrites(value,23,29)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s26.QueryP(414,&value)
	log("p23 r:%d", value)
	countreads(value,23,26)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s21.QueryP(328,&value)
	log("p23 r:%d", value)
	countreads(value,23,21)
	l.Unlock()

	delay()
	l.Lock()
	value = 17
	s17.Put(258,value)
	log("p23 w:%d", value)
	countwrites(value,23,17)
	l.Unlock()

	delay()
	l.Lock()
	value = 19
	s19.Put(299,value)
	log("p23 w:%d", value)
	countwrites(value,23,19)
	l.Unlock()

	delay()
	l.Lock()
	value = 32
	s32.Put(511,value)
	log("p23 w:%d", value)
	countwrites(value,23,32)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s12.QueryP(186,&value)
	log("p23 r:%d", value)
	countreads(value,23,12)
	l.Unlock()

	delay()
	l.Lock()
	value = 19
	s19.Put(291,value)
	log("p23 w:%d", value)
	countwrites(value,23,19)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s21.QueryP(325,&value)
	log("p23 r:%d", value)
	countreads(value,23,21)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s24.QueryP(372,&value)
	log("p23 r:%d", value)
	countreads(value,23,24)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s10.QueryP(160,&value)
	log("p23 r:%d", value)
	countreads(value,23,10)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s12.QueryP(185,&value)
	log("p23 r:%d", value)
	countreads(value,23,12)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s13.QueryP(204,&value)
	log("p23 r:%d", value)
	countreads(value,23,13)
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
	s14.QueryP(217,&value)
	log("p23 r:%d", value)
	countreads(value,23,14)
	l.Unlock()
}


func p24() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	s20.QueryP(318,&value)
	log("p24 r:%d", value)
	countreads(value,24,20)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s2.QueryP(17,&value)
	log("p24 r:%d", value)
	countreads(value,24,2)
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
	value = -1
	s8.QueryP(115,&value)
	log("p24 r:%d", value)
	countreads(value,24,8)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s15.QueryP(235,&value)
	log("p24 r:%d", value)
	countreads(value,24,15)
	l.Unlock()

	delay()
	l.Lock()
	value = 10
	s10.Put(151,value)
	log("p24 w:%d", value)
	countwrites(value,24,10)
	l.Unlock()

	delay()
	l.Lock()
	value = 4
	s4.Put(51,value)
	log("p24 w:%d", value)
	countwrites(value,24,4)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s12.QueryP(184,&value)
	log("p24 r:%d", value)
	countreads(value,24,12)
	l.Unlock()

	delay()
	l.Lock()
	value = 1
	s1.Put(14,value)
	log("p24 w:%d", value)
	countwrites(value,24,1)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s7.QueryP(97,&value)
	log("p24 r:%d", value)
	countreads(value,24,7)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s16.QueryP(241,&value)
	log("p24 r:%d", value)
	countreads(value,24,16)
	l.Unlock()

	delay()
	l.Lock()
	value = 9
	s9.Put(131,value)
	log("p24 w:%d", value)
	countwrites(value,24,9)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s26.QueryP(406,&value)
	log("p24 r:%d", value)
	countreads(value,24,26)
	l.Unlock()

	delay()
	l.Lock()
	value = 23
	s23.Put(360,value)
	log("p24 w:%d", value)
	countwrites(value,24,23)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s7.QueryP(108,&value)
	log("p24 r:%d", value)
	countreads(value,24,7)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s8.QueryP(115,&value)
	log("p24 r:%d", value)
	countreads(value,24,8)
	l.Unlock()
}


func p25() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 30
	s30.Put(470,value)
	log("p25 w:%d", value)
	countwrites(value,25,30)
	l.Unlock()

	delay()
	l.Lock()
	value = 11
	s11.Put(167,value)
	log("p25 w:%d", value)
	countwrites(value,25,11)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s20.QueryP(318,&value)
	log("p25 r:%d", value)
	countreads(value,25,20)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s25.QueryP(394,&value)
	log("p25 r:%d", value)
	countreads(value,25,25)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s5.QueryP(68,&value)
	log("p25 r:%d", value)
	countreads(value,25,5)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s16.QueryP(252,&value)
	log("p25 r:%d", value)
	countreads(value,25,16)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s17.QueryP(259,&value)
	log("p25 r:%d", value)
	countreads(value,25,17)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s6.QueryP(81,&value)
	log("p25 r:%d", value)
	countreads(value,25,6)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s22.QueryP(345,&value)
	log("p25 r:%d", value)
	countreads(value,25,22)
	l.Unlock()

	delay()
	l.Lock()
	value = 11
	s11.Put(167,value)
	log("p25 w:%d", value)
	countwrites(value,25,11)
	l.Unlock()

	delay()
	l.Lock()
	value = 21
	s21.Put(323,value)
	log("p25 w:%d", value)
	countwrites(value,25,21)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s17.QueryP(264,&value)
	log("p25 r:%d", value)
	countreads(value,25,17)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s25.QueryP(400,&value)
	log("p25 r:%d", value)
	countreads(value,25,25)
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
	value = 16
	s16.Put(256,value)
	log("p25 w:%d", value)
	countwrites(value,25,16)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s11.QueryP(171,&value)
	log("p25 r:%d", value)
	countreads(value,25,11)
	l.Unlock()
}


func p26() {
	defer wg.Done()
	var value int

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
	s18.QueryP(279,&value)
	log("p26 r:%d", value)
	countreads(value,26,18)
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
	value = 25
	s25.Put(394,value)
	log("p26 w:%d", value)
	countwrites(value,26,25)
	l.Unlock()

	delay()
	l.Lock()
	value = 30
	s30.Put(466,value)
	log("p26 w:%d", value)
	countwrites(value,26,30)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s7.QueryP(99,&value)
	log("p26 r:%d", value)
	countreads(value,26,7)
	l.Unlock()

	delay()
	l.Lock()
	value = 16
	s16.Put(244,value)
	log("p26 w:%d", value)
	countwrites(value,26,16)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s32.QueryP(500,&value)
	log("p26 r:%d", value)
	countreads(value,26,32)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s29.QueryP(455,&value)
	log("p26 r:%d", value)
	countreads(value,26,29)
	l.Unlock()

	delay()
	l.Lock()
	value = 32
	s32.Put(506,value)
	log("p26 w:%d", value)
	countwrites(value,26,32)
	l.Unlock()

	delay()
	l.Lock()
	value = 30
	s30.Put(476,value)
	log("p26 w:%d", value)
	countwrites(value,26,30)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s5.QueryP(65,&value)
	log("p26 r:%d", value)
	countreads(value,26,5)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s16.QueryP(256,&value)
	log("p26 r:%d", value)
	countreads(value,26,16)
	l.Unlock()

	delay()
	l.Lock()
	value = 16
	s16.Put(244,value)
	log("p26 w:%d", value)
	countwrites(value,26,16)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s19.QueryP(294,&value)
	log("p26 r:%d", value)
	countreads(value,26,19)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s23.QueryP(368,&value)
	log("p26 r:%d", value)
	countreads(value,26,23)
	l.Unlock()
}


func p27() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	s11.QueryP(165,&value)
	log("p27 r:%d", value)
	countreads(value,27,11)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s6.QueryP(94,&value)
	log("p27 r:%d", value)
	countreads(value,27,6)
	l.Unlock()

	delay()
	l.Lock()
	value = 21
	s21.Put(328,value)
	log("p27 w:%d", value)
	countwrites(value,27,21)
	l.Unlock()

	delay()
	l.Lock()
	value = 27
	s27.Put(420,value)
	log("p27 w:%d", value)
	countwrites(value,27,27)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s17.QueryP(269,&value)
	log("p27 r:%d", value)
	countreads(value,27,17)
	l.Unlock()

	delay()
	l.Lock()
	value = 26
	s26.Put(401,value)
	log("p27 w:%d", value)
	countwrites(value,27,26)
	l.Unlock()

	delay()
	l.Lock()
	value = 26
	s26.Put(414,value)
	log("p27 w:%d", value)
	countwrites(value,27,26)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s9.QueryP(144,&value)
	log("p27 r:%d", value)
	countreads(value,27,9)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s16.QueryP(255,&value)
	log("p27 r:%d", value)
	countreads(value,27,16)
	l.Unlock()

	delay()
	l.Lock()
	value = 27
	s27.Put(430,value)
	log("p27 w:%d", value)
	countwrites(value,27,27)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s22.QueryP(338,&value)
	log("p27 r:%d", value)
	countreads(value,27,22)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s23.QueryP(364,&value)
	log("p27 r:%d", value)
	countreads(value,27,23)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s17.QueryP(264,&value)
	log("p27 r:%d", value)
	countreads(value,27,17)
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
	s3.QueryP(48,&value)
	log("p27 r:%d", value)
	countreads(value,27,3)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s10.QueryP(151,&value)
	log("p27 r:%d", value)
	countreads(value,27,10)
	l.Unlock()
}


func p28() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	s22.QueryP(352,&value)
	log("p28 r:%d", value)
	countreads(value,28,22)
	l.Unlock()

	delay()
	l.Lock()
	value = 1
	s1.Put(15,value)
	log("p28 w:%d", value)
	countwrites(value,28,1)
	l.Unlock()

	delay()
	l.Lock()
	value = 22
	s22.Put(342,value)
	log("p28 w:%d", value)
	countwrites(value,28,22)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s7.QueryP(98,&value)
	log("p28 r:%d", value)
	countreads(value,28,7)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s6.QueryP(87,&value)
	log("p28 r:%d", value)
	countreads(value,28,6)
	l.Unlock()

	delay()
	l.Lock()
	value = 28
	s28.Put(436,value)
	log("p28 w:%d", value)
	countwrites(value,28,28)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s9.QueryP(142,&value)
	log("p28 r:%d", value)
	countreads(value,28,9)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s9.QueryP(141,&value)
	log("p28 r:%d", value)
	countreads(value,28,9)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s9.QueryP(129,&value)
	log("p28 r:%d", value)
	countreads(value,28,9)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s2.QueryP(17,&value)
	log("p28 r:%d", value)
	countreads(value,28,2)
	l.Unlock()

	delay()
	l.Lock()
	value = 26
	s26.Put(416,value)
	log("p28 w:%d", value)
	countwrites(value,28,26)
	l.Unlock()

	delay()
	l.Lock()
	value = 27
	s27.Put(432,value)
	log("p28 w:%d", value)
	countwrites(value,28,27)
	l.Unlock()

	delay()
	l.Lock()
	value = 21
	s21.Put(326,value)
	log("p28 w:%d", value)
	countwrites(value,28,21)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s6.QueryP(88,&value)
	log("p28 r:%d", value)
	countreads(value,28,6)
	l.Unlock()

	delay()
	l.Lock()
	value = 8
	s8.Put(123,value)
	log("p28 w:%d", value)
	countwrites(value,28,8)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s11.QueryP(169,&value)
	log("p28 r:%d", value)
	countreads(value,28,11)
	l.Unlock()
}


func p29() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 13
	s13.Put(196,value)
	log("p29 w:%d", value)
	countwrites(value,29,13)
	l.Unlock()

	delay()
	l.Lock()
	value = 26
	s26.Put(407,value)
	log("p29 w:%d", value)
	countwrites(value,29,26)
	l.Unlock()

	delay()
	l.Lock()
	value = 8
	s8.Put(115,value)
	log("p29 w:%d", value)
	countwrites(value,29,8)
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
	s5.QueryP(79,&value)
	log("p29 r:%d", value)
	countreads(value,29,5)
	l.Unlock()

	delay()
	l.Lock()
	value = 26
	s26.Put(413,value)
	log("p29 w:%d", value)
	countwrites(value,29,26)
	l.Unlock()

	delay()
	l.Lock()
	value = 11
	s11.Put(169,value)
	log("p29 w:%d", value)
	countwrites(value,29,11)
	l.Unlock()

	delay()
	l.Lock()
	value = 2
	s2.Put(23,value)
	log("p29 w:%d", value)
	countwrites(value,29,2)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s27.QueryP(432,&value)
	log("p29 r:%d", value)
	countreads(value,29,27)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s9.QueryP(144,&value)
	log("p29 r:%d", value)
	countreads(value,29,9)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s29.QueryP(463,&value)
	log("p29 r:%d", value)
	countreads(value,29,29)
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
	value = 5
	s5.Put(80,value)
	log("p29 w:%d", value)
	countwrites(value,29,5)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s7.QueryP(101,&value)
	log("p29 r:%d", value)
	countreads(value,29,7)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s29.QueryP(455,&value)
	log("p29 r:%d", value)
	countreads(value,29,29)
	l.Unlock()

	delay()
	l.Lock()
	value = 24
	s24.Put(373,value)
	log("p29 w:%d", value)
	countwrites(value,29,24)
	l.Unlock()
}


func p30() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	s3.QueryP(44,&value)
	log("p30 r:%d", value)
	countreads(value,30,3)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s20.QueryP(305,&value)
	log("p30 r:%d", value)
	countreads(value,30,20)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s31.QueryP(493,&value)
	log("p30 r:%d", value)
	countreads(value,30,31)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s30.QueryP(467,&value)
	log("p30 r:%d", value)
	countreads(value,30,30)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s30.QueryP(475,&value)
	log("p30 r:%d", value)
	countreads(value,30,30)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s7.QueryP(110,&value)
	log("p30 r:%d", value)
	countreads(value,30,7)
	l.Unlock()

	delay()
	l.Lock()
	value = 6
	s6.Put(83,value)
	log("p30 w:%d", value)
	countwrites(value,30,6)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s20.QueryP(317,&value)
	log("p30 r:%d", value)
	countreads(value,30,20)
	l.Unlock()

	delay()
	l.Lock()
	value = 29
	s29.Put(462,value)
	log("p30 w:%d", value)
	countwrites(value,30,29)
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
	value = 5
	s5.Put(77,value)
	log("p30 w:%d", value)
	countwrites(value,30,5)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s9.QueryP(138,&value)
	log("p30 r:%d", value)
	countreads(value,30,9)
	l.Unlock()

	delay()
	l.Lock()
	value = 17
	s17.Put(272,value)
	log("p30 w:%d", value)
	countwrites(value,30,17)
	l.Unlock()

	delay()
	l.Lock()
	value = 17
	s17.Put(260,value)
	log("p30 w:%d", value)
	countwrites(value,30,17)
	l.Unlock()

	delay()
	l.Lock()
	value = 13
	s13.Put(201,value)
	log("p30 w:%d", value)
	countwrites(value,30,13)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s18.QueryP(287,&value)
	log("p30 r:%d", value)
	countreads(value,30,18)
	l.Unlock()
}


func p31() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	s2.QueryP(26,&value)
	log("p31 r:%d", value)
	countreads(value,31,2)
	l.Unlock()

	delay()
	l.Lock()
	value = 12
	s12.Put(186,value)
	log("p31 w:%d", value)
	countwrites(value,31,12)
	l.Unlock()

	delay()
	l.Lock()
	value = 26
	s26.Put(407,value)
	log("p31 w:%d", value)
	countwrites(value,31,26)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s6.QueryP(93,&value)
	log("p31 r:%d", value)
	countreads(value,31,6)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s27.QueryP(419,&value)
	log("p31 r:%d", value)
	countreads(value,31,27)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s16.QueryP(244,&value)
	log("p31 r:%d", value)
	countreads(value,31,16)
	l.Unlock()

	delay()
	l.Lock()
	value = 22
	s22.Put(338,value)
	log("p31 w:%d", value)
	countwrites(value,31,22)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s24.QueryP(375,&value)
	log("p31 r:%d", value)
	countreads(value,31,24)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s24.QueryP(369,&value)
	log("p31 r:%d", value)
	countreads(value,31,24)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s16.QueryP(247,&value)
	log("p31 r:%d", value)
	countreads(value,31,16)
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
	s9.QueryP(129,&value)
	log("p31 r:%d", value)
	countreads(value,31,9)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s14.QueryP(217,&value)
	log("p31 r:%d", value)
	countreads(value,31,14)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s6.QueryP(81,&value)
	log("p31 r:%d", value)
	countreads(value,31,6)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s30.QueryP(473,&value)
	log("p31 r:%d", value)
	countreads(value,31,30)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s13.QueryP(207,&value)
	log("p31 r:%d", value)
	countreads(value,31,13)
	l.Unlock()
}


func p32() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 5
	s5.Put(77,value)
	log("p32 w:%d", value)
	countwrites(value,32,5)
	l.Unlock()

	delay()
	l.Lock()
	value = 28
	s28.Put(433,value)
	log("p32 w:%d", value)
	countwrites(value,32,28)
	l.Unlock()

	delay()
	l.Lock()
	value = 19
	s19.Put(297,value)
	log("p32 w:%d", value)
	countwrites(value,32,19)
	l.Unlock()

	delay()
	l.Lock()
	value = 24
	s24.Put(372,value)
	log("p32 w:%d", value)
	countwrites(value,32,24)
	l.Unlock()

	delay()
	l.Lock()
	value = 18
	s18.Put(288,value)
	log("p32 w:%d", value)
	countwrites(value,32,18)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s14.QueryP(221,&value)
	log("p32 r:%d", value)
	countreads(value,32,14)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s27.QueryP(423,&value)
	log("p32 r:%d", value)
	countreads(value,32,27)
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
	s9.QueryP(134,&value)
	log("p32 r:%d", value)
	countreads(value,32,9)
	l.Unlock()

	delay()
	l.Lock()
	value = 15
	s15.Put(236,value)
	log("p32 w:%d", value)
	countwrites(value,32,15)
	l.Unlock()

	delay()
	l.Lock()
	value = 9
	s9.Put(143,value)
	log("p32 w:%d", value)
	countwrites(value,32,9)
	l.Unlock()

	delay()
	l.Lock()
	value = 11
	s11.Put(169,value)
	log("p32 w:%d", value)
	countwrites(value,32,11)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s29.QueryP(457,&value)
	log("p32 r:%d", value)
	countreads(value,32,29)
	l.Unlock()

	delay()
	l.Lock()
	value = 25
	s25.Put(390,value)
	log("p32 w:%d", value)
	countwrites(value,32,25)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s20.QueryP(317,&value)
	log("p32 r:%d", value)
	countreads(value,32,20)
	l.Unlock()

	delay()
	l.Lock()
	value = 30
	s30.Put(475,value)
	log("p32 w:%d", value)
	countwrites(value,32,30)
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
