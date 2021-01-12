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
	s15.QueryP(237,&value)
	log("p1 r:%d", value)
	countreads(value,1,15)
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
	value = 10
	s10.Put(148,value)
	log("p1 w:%d", value)
	countwrites(value,1,10)
	l.Unlock()

	delay()
	l.Lock()
	value = 19
	s19.Put(303,value)
	log("p1 w:%d", value)
	countwrites(value,1,19)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s11.QueryP(175,&value)
	log("p1 r:%d", value)
	countreads(value,1,11)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s7.QueryP(102,&value)
	log("p1 r:%d", value)
	countreads(value,1,7)
	l.Unlock()

	delay()
	l.Lock()
	value = 8
	s8.Put(120,value)
	log("p1 w:%d", value)
	countwrites(value,1,8)
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
	value = -1
	s28.QueryP(434,&value)
	log("p1 r:%d", value)
	countreads(value,1,28)
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
	value = -1
	s10.QueryP(159,&value)
	log("p1 r:%d", value)
	countreads(value,1,10)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s32.QueryP(502,&value)
	log("p1 r:%d", value)
	countreads(value,1,32)
	l.Unlock()

	delay()
	l.Lock()
	value = 11
	s11.Put(166,value)
	log("p1 w:%d", value)
	countwrites(value,1,11)
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
	s27.QueryP(431,&value)
	log("p1 r:%d", value)
	countreads(value,1,27)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s22.QueryP(341,&value)
	log("p1 r:%d", value)
	countreads(value,1,22)
	l.Unlock()
}

func p2() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 14
	s14.Put(212,value)
	log("p2 w:%d", value)
	countwrites(value,2,14)
	l.Unlock()

	delay()
	l.Lock()
	value = 19
	s19.Put(302,value)
	log("p2 w:%d", value)
	countwrites(value,2,19)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s22.QueryP(341,&value)
	log("p2 r:%d", value)
	countreads(value,2,22)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s12.QueryP(184,&value)
	log("p2 r:%d", value)
	countreads(value,2,12)
	l.Unlock()

	delay()
	l.Lock()
	value = 29
	s29.Put(464,value)
	log("p2 w:%d", value)
	countwrites(value,2,29)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s24.QueryP(372,&value)
	log("p2 r:%d", value)
	countreads(value,2,24)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s13.QueryP(201,&value)
	log("p2 r:%d", value)
	countreads(value,2,13)
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
	s10.QueryP(158,&value)
	log("p2 r:%d", value)
	countreads(value,2,10)
	l.Unlock()

	delay()
	l.Lock()
	value = 8
	s8.Put(119,value)
	log("p2 w:%d", value)
	countwrites(value,2,8)
	l.Unlock()

	delay()
	l.Lock()
	value = 21
	s21.Put(336,value)
	log("p2 w:%d", value)
	countwrites(value,2,21)
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
	value = -1
	s29.QueryP(451,&value)
	log("p2 r:%d", value)
	countreads(value,2,29)
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
	s22.QueryP(350,&value)
	log("p2 r:%d", value)
	countreads(value,2,22)
	l.Unlock()

	delay()
	l.Lock()
	value = 26
	s26.Put(412,value)
	log("p2 w:%d", value)
	countwrites(value,2,26)
	l.Unlock()
}

func p3() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 13
	s13.Put(197,value)
	log("p3 w:%d", value)
	countwrites(value,3,13)
	l.Unlock()

	delay()
	l.Lock()
	value = 4
	s4.Put(60,value)
	log("p3 w:%d", value)
	countwrites(value,3,4)
	l.Unlock()

	delay()
	l.Lock()
	value = 12
	s12.Put(190,value)
	log("p3 w:%d", value)
	countwrites(value,3,12)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s4.QueryP(59,&value)
	log("p3 r:%d", value)
	countreads(value,3,4)
	l.Unlock()

	delay()
	l.Lock()
	value = 14
	s14.Put(213,value)
	log("p3 w:%d", value)
	countwrites(value,3,14)
	l.Unlock()

	delay()
	l.Lock()
	value = 26
	s26.Put(414,value)
	log("p3 w:%d", value)
	countwrites(value,3,26)
	l.Unlock()

	delay()
	l.Lock()
	value = 20
	s20.Put(308,value)
	log("p3 w:%d", value)
	countwrites(value,3,20)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s4.QueryP(60,&value)
	log("p3 r:%d", value)
	countreads(value,3,4)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s26.QueryP(413,&value)
	log("p3 r:%d", value)
	countreads(value,3,26)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s14.QueryP(213,&value)
	log("p3 r:%d", value)
	countreads(value,3,14)
	l.Unlock()

	delay()
	l.Lock()
	value = 7
	s7.Put(97,value)
	log("p3 w:%d", value)
	countwrites(value,3,7)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s4.QueryP(62,&value)
	log("p3 r:%d", value)
	countreads(value,3,4)
	l.Unlock()

	delay()
	l.Lock()
	value = 11
	s11.Put(175,value)
	log("p3 w:%d", value)
	countwrites(value,3,11)
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
	s10.QueryP(156,&value)
	log("p3 r:%d", value)
	countreads(value,3,10)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s22.QueryP(347,&value)
	log("p3 r:%d", value)
	countreads(value,3,22)
	l.Unlock()
}

func p4() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	s9.QueryP(142,&value)
	log("p4 r:%d", value)
	countreads(value,4,9)
	l.Unlock()

	delay()
	l.Lock()
	value = 22
	s22.Put(345,value)
	log("p4 w:%d", value)
	countwrites(value,4,22)
	l.Unlock()

	delay()
	l.Lock()
	value = 31
	s31.Put(489,value)
	log("p4 w:%d", value)
	countwrites(value,4,31)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s14.QueryP(216,&value)
	log("p4 r:%d", value)
	countreads(value,4,14)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s7.QueryP(100,&value)
	log("p4 r:%d", value)
	countreads(value,4,7)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s31.QueryP(489,&value)
	log("p4 r:%d", value)
	countreads(value,4,31)
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
	value = 28
	s28.Put(440,value)
	log("p4 w:%d", value)
	countwrites(value,4,28)
	l.Unlock()

	delay()
	l.Lock()
	value = 19
	s19.Put(302,value)
	log("p4 w:%d", value)
	countwrites(value,4,19)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s23.QueryP(353,&value)
	log("p4 r:%d", value)
	countreads(value,4,23)
	l.Unlock()

	delay()
	l.Lock()
	value = 5
	s5.Put(70,value)
	log("p4 w:%d", value)
	countwrites(value,4,5)
	l.Unlock()

	delay()
	l.Lock()
	value = 30
	s30.Put(470,value)
	log("p4 w:%d", value)
	countwrites(value,4,30)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s18.QueryP(278,&value)
	log("p4 r:%d", value)
	countreads(value,4,18)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s21.QueryP(323,&value)
	log("p4 r:%d", value)
	countreads(value,4,21)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s5.QueryP(74,&value)
	log("p4 r:%d", value)
	countreads(value,4,5)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s31.QueryP(485,&value)
	log("p4 r:%d", value)
	countreads(value,4,31)
	l.Unlock()
}

func p5() {
	defer wg.Done()
	var value int

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
	s31.QueryP(486,&value)
	log("p5 r:%d", value)
	countreads(value,5,31)
	l.Unlock()

	delay()
	l.Lock()
	value = 11
	s11.Put(166,value)
	log("p5 w:%d", value)
	countwrites(value,5,11)
	l.Unlock()

	delay()
	l.Lock()
	value = 25
	s25.Put(398,value)
	log("p5 w:%d", value)
	countwrites(value,5,25)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s26.QueryP(403,&value)
	log("p5 r:%d", value)
	countreads(value,5,26)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s7.QueryP(111,&value)
	log("p5 r:%d", value)
	countreads(value,5,7)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s19.QueryP(291,&value)
	log("p5 r:%d", value)
	countreads(value,5,19)
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
	s5.QueryP(80,&value)
	log("p5 r:%d", value)
	countreads(value,5,5)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s7.QueryP(101,&value)
	log("p5 r:%d", value)
	countreads(value,5,7)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s12.QueryP(181,&value)
	log("p5 r:%d", value)
	countreads(value,5,12)
	l.Unlock()

	delay()
	l.Lock()
	value = 14
	s14.Put(215,value)
	log("p5 w:%d", value)
	countwrites(value,5,14)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s9.QueryP(138,&value)
	log("p5 r:%d", value)
	countreads(value,5,9)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s20.QueryP(313,&value)
	log("p5 r:%d", value)
	countreads(value,5,20)
	l.Unlock()

	delay()
	l.Lock()
	value = 26
	s26.Put(410,value)
	log("p5 w:%d", value)
	countwrites(value,5,26)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s30.QueryP(465,&value)
	log("p5 r:%d", value)
	countreads(value,5,30)
	l.Unlock()
}

func p6() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	s11.QueryP(166,&value)
	log("p6 r:%d", value)
	countreads(value,6,11)
	l.Unlock()

	delay()
	l.Lock()
	value = 23
	s23.Put(367,value)
	log("p6 w:%d", value)
	countwrites(value,6,23)
	l.Unlock()

	delay()
	l.Lock()
	value = 17
	s17.Put(257,value)
	log("p6 w:%d", value)
	countwrites(value,6,17)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s16.QueryP(253,&value)
	log("p6 r:%d", value)
	countreads(value,6,16)
	l.Unlock()

	delay()
	l.Lock()
	value = 3
	s3.Put(47,value)
	log("p6 w:%d", value)
	countwrites(value,6,3)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s32.QueryP(512,&value)
	log("p6 r:%d", value)
	countreads(value,6,32)
	l.Unlock()

	delay()
	l.Lock()
	value = 32
	s32.Put(508,value)
	log("p6 w:%d", value)
	countwrites(value,6,32)
	l.Unlock()

	delay()
	l.Lock()
	value = 6
	s6.Put(94,value)
	log("p6 w:%d", value)
	countwrites(value,6,6)
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
	value = 17
	s17.Put(266,value)
	log("p6 w:%d", value)
	countwrites(value,6,17)
	l.Unlock()

	delay()
	l.Lock()
	value = 31
	s31.Put(481,value)
	log("p6 w:%d", value)
	countwrites(value,6,31)
	l.Unlock()

	delay()
	l.Lock()
	value = 5
	s5.Put(67,value)
	log("p6 w:%d", value)
	countwrites(value,6,5)
	l.Unlock()

	delay()
	l.Lock()
	value = 16
	s16.Put(251,value)
	log("p6 w:%d", value)
	countwrites(value,6,16)
	l.Unlock()

	delay()
	l.Lock()
	value = 14
	s14.Put(222,value)
	log("p6 w:%d", value)
	countwrites(value,6,14)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s18.QueryP(278,&value)
	log("p6 r:%d", value)
	countreads(value,6,18)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s24.QueryP(377,&value)
	log("p6 r:%d", value)
	countreads(value,6,24)
	l.Unlock()
}

func p7() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 18
	s18.Put(278,value)
	log("p7 w:%d", value)
	countwrites(value,7,18)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s6.QueryP(88,&value)
	log("p7 r:%d", value)
	countreads(value,7,6)
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
	s5.QueryP(76,&value)
	log("p7 r:%d", value)
	countreads(value,7,5)
	l.Unlock()

	delay()
	l.Lock()
	value = 26
	s26.Put(413,value)
	log("p7 w:%d", value)
	countwrites(value,7,26)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s24.QueryP(376,&value)
	log("p7 r:%d", value)
	countreads(value,7,24)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s7.QueryP(99,&value)
	log("p7 r:%d", value)
	countreads(value,7,7)
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
	value = -1
	s29.QueryP(450,&value)
	log("p7 r:%d", value)
	countreads(value,7,29)
	l.Unlock()

	delay()
	l.Lock()
	value = 31
	s31.Put(496,value)
	log("p7 w:%d", value)
	countwrites(value,7,31)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s3.QueryP(47,&value)
	log("p7 r:%d", value)
	countreads(value,7,3)
	l.Unlock()

	delay()
	l.Lock()
	value = 15
	s15.Put(227,value)
	log("p7 w:%d", value)
	countwrites(value,7,15)
	l.Unlock()

	delay()
	l.Lock()
	value = 13
	s13.Put(208,value)
	log("p7 w:%d", value)
	countwrites(value,7,13)
	l.Unlock()

	delay()
	l.Lock()
	value = 11
	s11.Put(167,value)
	log("p7 w:%d", value)
	countwrites(value,7,11)
	l.Unlock()

	delay()
	l.Lock()
	value = 31
	s31.Put(494,value)
	log("p7 w:%d", value)
	countwrites(value,7,31)
	l.Unlock()

	delay()
	l.Lock()
	value = 20
	s20.Put(307,value)
	log("p7 w:%d", value)
	countwrites(value,7,20)
	l.Unlock()
}

func p8() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	s9.QueryP(138,&value)
	log("p8 r:%d", value)
	countreads(value,8,9)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s1.QueryP(5,&value)
	log("p8 r:%d", value)
	countreads(value,8,1)
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
	s24.QueryP(369,&value)
	log("p8 r:%d", value)
	countreads(value,8,24)
	l.Unlock()

	delay()
	l.Lock()
	value = 13
	s13.Put(197,value)
	log("p8 w:%d", value)
	countwrites(value,8,13)
	l.Unlock()

	delay()
	l.Lock()
	value = 23
	s23.Put(358,value)
	log("p8 w:%d", value)
	countwrites(value,8,23)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s9.QueryP(144,&value)
	log("p8 r:%d", value)
	countreads(value,8,9)
	l.Unlock()

	delay()
	l.Lock()
	value = 3
	s3.Put(45,value)
	log("p8 w:%d", value)
	countwrites(value,8,3)
	l.Unlock()

	delay()
	l.Lock()
	value = 3
	s3.Put(48,value)
	log("p8 w:%d", value)
	countwrites(value,8,3)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s4.QueryP(56,&value)
	log("p8 r:%d", value)
	countreads(value,8,4)
	l.Unlock()

	delay()
	l.Lock()
	value = 3
	s3.Put(47,value)
	log("p8 w:%d", value)
	countwrites(value,8,3)
	l.Unlock()

	delay()
	l.Lock()
	value = 13
	s13.Put(197,value)
	log("p8 w:%d", value)
	countwrites(value,8,13)
	l.Unlock()

	delay()
	l.Lock()
	value = 18
	s18.Put(286,value)
	log("p8 w:%d", value)
	countwrites(value,8,18)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s11.QueryP(174,&value)
	log("p8 r:%d", value)
	countreads(value,8,11)
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
	value = 19
	s19.Put(303,value)
	log("p8 w:%d", value)
	countwrites(value,8,19)
	l.Unlock()
}

func p9() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 23
	s23.Put(355,value)
	log("p9 w:%d", value)
	countwrites(value,9,23)
	l.Unlock()

	delay()
	l.Lock()
	value = 28
	s28.Put(440,value)
	log("p9 w:%d", value)
	countwrites(value,9,28)
	l.Unlock()

	delay()
	l.Lock()
	value = 19
	s19.Put(293,value)
	log("p9 w:%d", value)
	countwrites(value,9,19)
	l.Unlock()

	delay()
	l.Lock()
	value = 26
	s26.Put(413,value)
	log("p9 w:%d", value)
	countwrites(value,9,26)
	l.Unlock()

	delay()
	l.Lock()
	value = 5
	s5.Put(72,value)
	log("p9 w:%d", value)
	countwrites(value,9,5)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s27.QueryP(423,&value)
	log("p9 r:%d", value)
	countreads(value,9,27)
	l.Unlock()

	delay()
	l.Lock()
	value = 3
	s3.Put(40,value)
	log("p9 w:%d", value)
	countwrites(value,9,3)
	l.Unlock()

	delay()
	l.Lock()
	value = 15
	s15.Put(226,value)
	log("p9 w:%d", value)
	countwrites(value,9,15)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s23.QueryP(364,&value)
	log("p9 r:%d", value)
	countreads(value,9,23)
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
	value = 2
	s2.Put(28,value)
	log("p9 w:%d", value)
	countwrites(value,9,2)
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
	value = -1
	s1.QueryP(6,&value)
	log("p9 r:%d", value)
	countreads(value,9,1)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s32.QueryP(503,&value)
	log("p9 r:%d", value)
	countreads(value,9,32)
	l.Unlock()

	delay()
	l.Lock()
	value = 23
	s23.Put(353,value)
	log("p9 w:%d", value)
	countwrites(value,9,23)
	l.Unlock()

	delay()
	l.Lock()
	value = 30
	s30.Put(479,value)
	log("p9 w:%d", value)
	countwrites(value,9,30)
	l.Unlock()
}


func p10() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 7
	s7.Put(99,value)
	log("p10 w:%d", value)
	countwrites(value,10,7)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s14.QueryP(210,&value)
	log("p10 r:%d", value)
	countreads(value,10,14)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s2.QueryP(25,&value)
	log("p10 r:%d", value)
	countreads(value,10,2)
	l.Unlock()

	delay()
	l.Lock()
	value = 15
	s15.Put(227,value)
	log("p10 w:%d", value)
	countwrites(value,10,15)
	l.Unlock()

	delay()
	l.Lock()
	value = 28
	s28.Put(446,value)
	log("p10 w:%d", value)
	countwrites(value,10,28)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s18.QueryP(287,&value)
	log("p10 r:%d", value)
	countreads(value,10,18)
	l.Unlock()

	delay()
	l.Lock()
	value = 7
	s7.Put(106,value)
	log("p10 w:%d", value)
	countwrites(value,10,7)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s17.QueryP(264,&value)
	log("p10 r:%d", value)
	countreads(value,10,17)
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
	s31.QueryP(490,&value)
	log("p10 r:%d", value)
	countreads(value,10,31)
	l.Unlock()

	delay()
	l.Lock()
	value = 14
	s14.Put(211,value)
	log("p10 w:%d", value)
	countwrites(value,10,14)
	l.Unlock()

	delay()
	l.Lock()
	value = 28
	s28.Put(440,value)
	log("p10 w:%d", value)
	countwrites(value,10,28)
	l.Unlock()

	delay()
	l.Lock()
	value = 14
	s14.Put(219,value)
	log("p10 w:%d", value)
	countwrites(value,10,14)
	l.Unlock()

	delay()
	l.Lock()
	value = 6
	s6.Put(81,value)
	log("p10 w:%d", value)
	countwrites(value,10,6)
	l.Unlock()

	delay()
	l.Lock()
	value = 11
	s11.Put(171,value)
	log("p10 w:%d", value)
	countwrites(value,10,11)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s7.QueryP(99,&value)
	log("p10 r:%d", value)
	countreads(value,10,7)
	l.Unlock()
}


func p11() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	s27.QueryP(423,&value)
	log("p11 r:%d", value)
	countreads(value,11,27)
	l.Unlock()

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
	s24.QueryP(372,&value)
	log("p11 r:%d", value)
	countreads(value,11,24)
	l.Unlock()

	delay()
	l.Lock()
	value = 20
	s20.Put(312,value)
	log("p11 w:%d", value)
	countwrites(value,11,20)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s11.QueryP(162,&value)
	log("p11 r:%d", value)
	countreads(value,11,11)
	l.Unlock()

	delay()
	l.Lock()
	value = 20
	s20.Put(317,value)
	log("p11 w:%d", value)
	countwrites(value,11,20)
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
	value = 23
	s23.Put(366,value)
	log("p11 w:%d", value)
	countwrites(value,11,23)
	l.Unlock()

	delay()
	l.Lock()
	value = 13
	s13.Put(206,value)
	log("p11 w:%d", value)
	countwrites(value,11,13)
	l.Unlock()

	delay()
	l.Lock()
	value = 1
	s1.Put(7,value)
	log("p11 w:%d", value)
	countwrites(value,11,1)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s8.QueryP(122,&value)
	log("p11 r:%d", value)
	countreads(value,11,8)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s7.QueryP(111,&value)
	log("p11 r:%d", value)
	countreads(value,11,7)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s32.QueryP(505,&value)
	log("p11 r:%d", value)
	countreads(value,11,32)
	l.Unlock()

	delay()
	l.Lock()
	value = 22
	s22.Put(342,value)
	log("p11 w:%d", value)
	countwrites(value,11,22)
	l.Unlock()

	delay()
	l.Lock()
	value = 24
	s24.Put(375,value)
	log("p11 w:%d", value)
	countwrites(value,11,24)
	l.Unlock()

	delay()
	l.Lock()
	value = 17
	s17.Put(259,value)
	log("p11 w:%d", value)
	countwrites(value,11,17)
	l.Unlock()
}


func p12() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 29
	s29.Put(455,value)
	log("p12 w:%d", value)
	countwrites(value,12,29)
	l.Unlock()

	delay()
	l.Lock()
	value = 9
	s9.Put(130,value)
	log("p12 w:%d", value)
	countwrites(value,12,9)
	l.Unlock()

	delay()
	l.Lock()
	value = 22
	s22.Put(347,value)
	log("p12 w:%d", value)
	countwrites(value,12,22)
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
	value = -1
	s31.QueryP(489,&value)
	log("p12 r:%d", value)
	countreads(value,12,31)
	l.Unlock()

	delay()
	l.Lock()
	value = 28
	s28.Put(440,value)
	log("p12 w:%d", value)
	countwrites(value,12,28)
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
	value = 28
	s28.Put(438,value)
	log("p12 w:%d", value)
	countwrites(value,12,28)
	l.Unlock()

	delay()
	l.Lock()
	value = 3
	s3.Put(33,value)
	log("p12 w:%d", value)
	countwrites(value,12,3)
	l.Unlock()

	delay()
	l.Lock()
	value = 14
	s14.Put(224,value)
	log("p12 w:%d", value)
	countwrites(value,12,14)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s26.QueryP(412,&value)
	log("p12 r:%d", value)
	countreads(value,12,26)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s5.QueryP(79,&value)
	log("p12 r:%d", value)
	countreads(value,12,5)
	l.Unlock()

	delay()
	l.Lock()
	value = 11
	s11.Put(162,value)
	log("p12 w:%d", value)
	countwrites(value,12,11)
	l.Unlock()

	delay()
	l.Lock()
	value = 27
	s27.Put(417,value)
	log("p12 w:%d", value)
	countwrites(value,12,27)
	l.Unlock()

	delay()
	l.Lock()
	value = 2
	s2.Put(31,value)
	log("p12 w:%d", value)
	countwrites(value,12,2)
	l.Unlock()

	delay()
	l.Lock()
	value = 11
	s11.Put(167,value)
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
	s27.QueryP(423,&value)
	log("p13 r:%d", value)
	countreads(value,13,27)
	l.Unlock()

	delay()
	l.Lock()
	value = 13
	s13.Put(195,value)
	log("p13 w:%d", value)
	countwrites(value,13,13)
	l.Unlock()

	delay()
	l.Lock()
	value = 14
	s14.Put(224,value)
	log("p13 w:%d", value)
	countwrites(value,13,14)
	l.Unlock()

	delay()
	l.Lock()
	value = 1
	s1.Put(16,value)
	log("p13 w:%d", value)
	countwrites(value,13,1)
	l.Unlock()

	delay()
	l.Lock()
	value = 22
	s22.Put(346,value)
	log("p13 w:%d", value)
	countwrites(value,13,22)
	l.Unlock()

	delay()
	l.Lock()
	value = 12
	s12.Put(186,value)
	log("p13 w:%d", value)
	countwrites(value,13,12)
	l.Unlock()

	delay()
	l.Lock()
	value = 27
	s27.Put(422,value)
	log("p13 w:%d", value)
	countwrites(value,13,27)
	l.Unlock()

	delay()
	l.Lock()
	value = 10
	s10.Put(158,value)
	log("p13 w:%d", value)
	countwrites(value,13,10)
	l.Unlock()

	delay()
	l.Lock()
	value = 8
	s8.Put(124,value)
	log("p13 w:%d", value)
	countwrites(value,13,8)
	l.Unlock()

	delay()
	l.Lock()
	value = 16
	s16.Put(245,value)
	log("p13 w:%d", value)
	countwrites(value,13,16)
	l.Unlock()

	delay()
	l.Lock()
	value = 12
	s12.Put(181,value)
	log("p13 w:%d", value)
	countwrites(value,13,12)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s25.QueryP(393,&value)
	log("p13 r:%d", value)
	countreads(value,13,25)
	l.Unlock()

	delay()
	l.Lock()
	value = 29
	s29.Put(459,value)
	log("p13 w:%d", value)
	countwrites(value,13,29)
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
	value = 3
	s3.Put(46,value)
	log("p13 w:%d", value)
	countwrites(value,13,3)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s10.QueryP(145,&value)
	log("p13 r:%d", value)
	countreads(value,13,10)
	l.Unlock()
}

func p14() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	s15.QueryP(229,&value)
	log("p14 r:%d", value)
	countreads(value,14,15)
	l.Unlock()

	delay()
	l.Lock()
	value = 26
	s26.Put(413,value)
	log("p14 w:%d", value)
	countwrites(value,14,26)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s15.QueryP(231,&value)
	log("p14 r:%d", value)
	countreads(value,14,15)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s10.QueryP(146,&value)
	log("p14 r:%d", value)
	countreads(value,14,10)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s22.QueryP(343,&value)
	log("p14 r:%d", value)
	countreads(value,14,22)
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
	value = 26
	s26.Put(401,value)
	log("p14 w:%d", value)
	countwrites(value,14,26)
	l.Unlock()

	delay()
	l.Lock()
	value = 28
	s28.Put(438,value)
	log("p14 w:%d", value)
	countwrites(value,14,28)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s21.QueryP(325,&value)
	log("p14 r:%d", value)
	countreads(value,14,21)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s7.QueryP(107,&value)
	log("p14 r:%d", value)
	countreads(value,14,7)
	l.Unlock()

	delay()
	l.Lock()
	value = 6
	s6.Put(91,value)
	log("p14 w:%d", value)
	countwrites(value,14,6)
	l.Unlock()

	delay()
	l.Lock()
	value = 14
	s14.Put(209,value)
	log("p14 w:%d", value)
	countwrites(value,14,14)
	l.Unlock()

	delay()
	l.Lock()
	value = 22
	s22.Put(352,value)
	log("p14 w:%d", value)
	countwrites(value,14,22)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s22.QueryP(337,&value)
	log("p14 r:%d", value)
	countreads(value,14,22)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s21.QueryP(325,&value)
	log("p14 r:%d", value)
	countreads(value,14,21)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s2.QueryP(24,&value)
	log("p14 r:%d", value)
	countreads(value,14,2)
	l.Unlock()
}


func p15() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 2
	s2.Put(25,value)
	log("p15 w:%d", value)
	countwrites(value,15,2)
	l.Unlock()

	delay()
	l.Lock()
	value = 7
	s7.Put(100,value)
	log("p15 w:%d", value)
	countwrites(value,15,7)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s27.QueryP(418,&value)
	log("p15 r:%d", value)
	countreads(value,15,27)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s5.QueryP(69,&value)
	log("p15 r:%d", value)
	countreads(value,15,5)
	l.Unlock()

	delay()
	l.Lock()
	value = 7
	s7.Put(98,value)
	log("p15 w:%d", value)
	countwrites(value,15,7)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s4.QueryP(51,&value)
	log("p15 r:%d", value)
	countreads(value,15,4)
	l.Unlock()

	delay()
	l.Lock()
	value = 14
	s14.Put(223,value)
	log("p15 w:%d", value)
	countwrites(value,15,14)
	l.Unlock()

	delay()
	l.Lock()
	value = 2
	s2.Put(18,value)
	log("p15 w:%d", value)
	countwrites(value,15,2)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s19.QueryP(295,&value)
	log("p15 r:%d", value)
	countreads(value,15,19)
	l.Unlock()

	delay()
	l.Lock()
	value = 20
	s20.Put(315,value)
	log("p15 w:%d", value)
	countwrites(value,15,20)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s14.QueryP(211,&value)
	log("p15 r:%d", value)
	countreads(value,15,14)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s1.QueryP(10,&value)
	log("p15 r:%d", value)
	countreads(value,15,1)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s19.QueryP(304,&value)
	log("p15 r:%d", value)
	countreads(value,15,19)
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
	value = 4
	s4.Put(49,value)
	log("p15 w:%d", value)
	countwrites(value,15,4)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s9.QueryP(133,&value)
	log("p15 r:%d", value)
	countreads(value,15,9)
	l.Unlock()
}


func p16() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 21
	s21.Put(326,value)
	log("p16 w:%d", value)
	countwrites(value,16,21)
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
	value = 10
	s10.Put(145,value)
	log("p16 w:%d", value)
	countwrites(value,16,10)
	l.Unlock()

	delay()
	l.Lock()
	value = 32
	s32.Put(502,value)
	log("p16 w:%d", value)
	countwrites(value,16,32)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s4.QueryP(56,&value)
	log("p16 r:%d", value)
	countreads(value,16,4)
	l.Unlock()

	delay()
	l.Lock()
	value = 27
	s27.Put(432,value)
	log("p16 w:%d", value)
	countwrites(value,16,27)
	l.Unlock()

	delay()
	l.Lock()
	value = 31
	s31.Put(491,value)
	log("p16 w:%d", value)
	countwrites(value,16,31)
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
	value = 6
	s6.Put(90,value)
	log("p16 w:%d", value)
	countwrites(value,16,6)
	l.Unlock()

	delay()
	l.Lock()
	value = 25
	s25.Put(394,value)
	log("p16 w:%d", value)
	countwrites(value,16,25)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s18.QueryP(284,&value)
	log("p16 r:%d", value)
	countreads(value,16,18)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s10.QueryP(147,&value)
	log("p16 r:%d", value)
	countreads(value,16,10)
	l.Unlock()

	delay()
	l.Lock()
	value = 32
	s32.Put(502,value)
	log("p16 w:%d", value)
	countwrites(value,16,32)
	l.Unlock()

	delay()
	l.Lock()
	value = 3
	s3.Put(39,value)
	log("p16 w:%d", value)
	countwrites(value,16,3)
	l.Unlock()

	delay()
	l.Lock()
	value = 24
	s24.Put(380,value)
	log("p16 w:%d", value)
	countwrites(value,16,24)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s17.QueryP(257,&value)
	log("p16 r:%d", value)
	countreads(value,16,17)
	l.Unlock()
}


func p17() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 11
	s11.Put(174,value)
	log("p17 w:%d", value)
	countwrites(value,17,11)
	l.Unlock()

	delay()
	l.Lock()
	value = 23
	s23.Put(357,value)
	log("p17 w:%d", value)
	countwrites(value,17,23)
	l.Unlock()

	delay()
	l.Lock()
	value = 10
	s10.Put(148,value)
	log("p17 w:%d", value)
	countwrites(value,17,10)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s9.QueryP(142,&value)
	log("p17 r:%d", value)
	countreads(value,17,9)
	l.Unlock()

	delay()
	l.Lock()
	value = 26
	s26.Put(415,value)
	log("p17 w:%d", value)
	countwrites(value,17,26)
	l.Unlock()

	delay()
	l.Lock()
	value = 10
	s10.Put(145,value)
	log("p17 w:%d", value)
	countwrites(value,17,10)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s26.QueryP(415,&value)
	log("p17 r:%d", value)
	countreads(value,17,26)
	l.Unlock()

	delay()
	l.Lock()
	value = 2
	s2.Put(25,value)
	log("p17 w:%d", value)
	countwrites(value,17,2)
	l.Unlock()

	delay()
	l.Lock()
	value = 18
	s18.Put(279,value)
	log("p17 w:%d", value)
	countwrites(value,17,18)
	l.Unlock()

	delay()
	l.Lock()
	value = 8
	s8.Put(125,value)
	log("p17 w:%d", value)
	countwrites(value,17,8)
	l.Unlock()

	delay()
	l.Lock()
	value = 32
	s32.Put(512,value)
	log("p17 w:%d", value)
	countwrites(value,17,32)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s14.QueryP(223,&value)
	log("p17 r:%d", value)
	countreads(value,17,14)
	l.Unlock()

	delay()
	l.Lock()
	value = 14
	s14.Put(218,value)
	log("p17 w:%d", value)
	countwrites(value,17,14)
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
	value = 30
	s30.Put(473,value)
	log("p17 w:%d", value)
	countwrites(value,17,30)
	l.Unlock()

	delay()
	l.Lock()
	value = 3
	s3.Put(35,value)
	log("p17 w:%d", value)
	countwrites(value,17,3)
	l.Unlock()
}


func p18() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 17
	s17.Put(263,value)
	log("p18 w:%d", value)
	countwrites(value,18,17)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s7.QueryP(104,&value)
	log("p18 r:%d", value)
	countreads(value,18,7)
	l.Unlock()

	delay()
	l.Lock()
	value = 24
	s24.Put(384,value)
	log("p18 w:%d", value)
	countwrites(value,18,24)
	l.Unlock()

	delay()
	l.Lock()
	value = 2
	s2.Put(23,value)
	log("p18 w:%d", value)
	countwrites(value,18,2)
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
	value = -1
	s22.QueryP(350,&value)
	log("p18 r:%d", value)
	countreads(value,18,22)
	l.Unlock()

	delay()
	l.Lock()
	value = 13
	s13.Put(201,value)
	log("p18 w:%d", value)
	countwrites(value,18,13)
	l.Unlock()

	delay()
	l.Lock()
	value = 26
	s26.Put(406,value)
	log("p18 w:%d", value)
	countwrites(value,18,26)
	l.Unlock()

	delay()
	l.Lock()
	value = 10
	s10.Put(145,value)
	log("p18 w:%d", value)
	countwrites(value,18,10)
	l.Unlock()

	delay()
	l.Lock()
	value = 2
	s2.Put(23,value)
	log("p18 w:%d", value)
	countwrites(value,18,2)
	l.Unlock()

	delay()
	l.Lock()
	value = 15
	s15.Put(239,value)
	log("p18 w:%d", value)
	countwrites(value,18,15)
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
	value = 4
	s4.Put(60,value)
	log("p18 w:%d", value)
	countwrites(value,18,4)
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
	value = 18
	s18.Put(277,value)
	log("p18 w:%d", value)
	countwrites(value,18,18)
	l.Unlock()

	delay()
	l.Lock()
	value = 32
	s32.Put(511,value)
	log("p18 w:%d", value)
	countwrites(value,18,32)
	l.Unlock()
}


func p19() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	s1.QueryP(14,&value)
	log("p19 r:%d", value)
	countreads(value,19,1)
	l.Unlock()

	delay()
	l.Lock()
	value = 31
	s31.Put(485,value)
	log("p19 w:%d", value)
	countwrites(value,19,31)
	l.Unlock()

	delay()
	l.Lock()
	value = 10
	s10.Put(145,value)
	log("p19 w:%d", value)
	countwrites(value,19,10)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s10.QueryP(146,&value)
	log("p19 r:%d", value)
	countreads(value,19,10)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s9.QueryP(141,&value)
	log("p19 r:%d", value)
	countreads(value,19,9)
	l.Unlock()

	delay()
	l.Lock()
	value = 22
	s22.Put(339,value)
	log("p19 w:%d", value)
	countwrites(value,19,22)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s29.QueryP(464,&value)
	log("p19 r:%d", value)
	countreads(value,19,29)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s10.QueryP(159,&value)
	log("p19 r:%d", value)
	countreads(value,19,10)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s2.QueryP(22,&value)
	log("p19 r:%d", value)
	countreads(value,19,2)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s17.QueryP(267,&value)
	log("p19 r:%d", value)
	countreads(value,19,17)
	l.Unlock()

	delay()
	l.Lock()
	value = 24
	s24.Put(375,value)
	log("p19 w:%d", value)
	countwrites(value,19,24)
	l.Unlock()

	delay()
	l.Lock()
	value = 22
	s22.Put(350,value)
	log("p19 w:%d", value)
	countwrites(value,19,22)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s9.QueryP(144,&value)
	log("p19 r:%d", value)
	countreads(value,19,9)
	l.Unlock()

	delay()
	l.Lock()
	value = 26
	s26.Put(407,value)
	log("p19 w:%d", value)
	countwrites(value,19,26)
	l.Unlock()

	delay()
	l.Lock()
	value = 1
	s1.Put(2,value)
	log("p19 w:%d", value)
	countwrites(value,19,1)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s10.QueryP(153,&value)
	log("p19 r:%d", value)
	countreads(value,19,10)
	l.Unlock()
}


func p20() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	s20.QueryP(314,&value)
	log("p20 r:%d", value)
	countreads(value,20,20)
	l.Unlock()

	delay()
	l.Lock()
	value = 2
	s2.Put(24,value)
	log("p20 w:%d", value)
	countwrites(value,20,2)
	l.Unlock()

	delay()
	l.Lock()
	value = 13
	s13.Put(197,value)
	log("p20 w:%d", value)
	countwrites(value,20,13)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s6.QueryP(96,&value)
	log("p20 r:%d", value)
	countreads(value,20,6)
	l.Unlock()

	delay()
	l.Lock()
	value = 6
	s6.Put(81,value)
	log("p20 w:%d", value)
	countwrites(value,20,6)
	l.Unlock()

	delay()
	l.Lock()
	value = 9
	s9.Put(138,value)
	log("p20 w:%d", value)
	countwrites(value,20,9)
	l.Unlock()

	delay()
	l.Lock()
	value = 24
	s24.Put(374,value)
	log("p20 w:%d", value)
	countwrites(value,20,24)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s6.QueryP(94,&value)
	log("p20 r:%d", value)
	countreads(value,20,6)
	l.Unlock()

	delay()
	l.Lock()
	value = 28
	s28.Put(442,value)
	log("p20 w:%d", value)
	countwrites(value,20,28)
	l.Unlock()

	delay()
	l.Lock()
	value = 24
	s24.Put(371,value)
	log("p20 w:%d", value)
	countwrites(value,20,24)
	l.Unlock()

	delay()
	l.Lock()
	value = 23
	s23.Put(356,value)
	log("p20 w:%d", value)
	countwrites(value,20,23)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s3.QueryP(37,&value)
	log("p20 r:%d", value)
	countreads(value,20,3)
	l.Unlock()

	delay()
	l.Lock()
	value = 22
	s22.Put(340,value)
	log("p20 w:%d", value)
	countwrites(value,20,22)
	l.Unlock()

	delay()
	l.Lock()
	value = 23
	s23.Put(355,value)
	log("p20 w:%d", value)
	countwrites(value,20,23)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s13.QueryP(201,&value)
	log("p20 r:%d", value)
	countreads(value,20,13)
	l.Unlock()

	delay()
	l.Lock()
	value = 8
	s8.Put(116,value)
	log("p20 w:%d", value)
	countwrites(value,20,8)
	l.Unlock()
}


func p21() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 28
	s28.Put(440,value)
	log("p21 w:%d", value)
	countwrites(value,21,28)
	l.Unlock()

	delay()
	l.Lock()
	value = 29
	s29.Put(449,value)
	log("p21 w:%d", value)
	countwrites(value,21,29)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s10.QueryP(149,&value)
	log("p21 r:%d", value)
	countreads(value,21,10)
	l.Unlock()

	delay()
	l.Lock()
	value = 8
	s8.Put(128,value)
	log("p21 w:%d", value)
	countwrites(value,21,8)
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
	value = 19
	s19.Put(292,value)
	log("p21 w:%d", value)
	countwrites(value,21,19)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s29.QueryP(464,&value)
	log("p21 r:%d", value)
	countreads(value,21,29)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s24.QueryP(384,&value)
	log("p21 r:%d", value)
	countreads(value,21,24)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s32.QueryP(501,&value)
	log("p21 r:%d", value)
	countreads(value,21,32)
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
	value = 2
	s2.Put(20,value)
	log("p21 w:%d", value)
	countwrites(value,21,2)
	l.Unlock()

	delay()
	l.Lock()
	value = 9
	s9.Put(138,value)
	log("p21 w:%d", value)
	countwrites(value,21,9)
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
	value = 32
	s32.Put(500,value)
	log("p21 w:%d", value)
	countwrites(value,21,32)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s32.QueryP(498,&value)
	log("p21 r:%d", value)
	countreads(value,21,32)
	l.Unlock()

	delay()
	l.Lock()
	value = 9
	s9.Put(131,value)
	log("p21 w:%d", value)
	countwrites(value,21,9)
	l.Unlock()
}


func p22() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	s15.QueryP(230,&value)
	log("p22 r:%d", value)
	countreads(value,22,15)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s23.QueryP(368,&value)
	log("p22 r:%d", value)
	countreads(value,22,23)
	l.Unlock()

	delay()
	l.Lock()
	value = 26
	s26.Put(407,value)
	log("p22 w:%d", value)
	countwrites(value,22,26)
	l.Unlock()

	delay()
	l.Lock()
	value = 31
	s31.Put(484,value)
	log("p22 w:%d", value)
	countwrites(value,22,31)
	l.Unlock()

	delay()
	l.Lock()
	value = 18
	s18.Put(280,value)
	log("p22 w:%d", value)
	countwrites(value,22,18)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s19.QueryP(291,&value)
	log("p22 r:%d", value)
	countreads(value,22,19)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s8.QueryP(116,&value)
	log("p22 r:%d", value)
	countreads(value,22,8)
	l.Unlock()

	delay()
	l.Lock()
	value = 20
	s20.Put(308,value)
	log("p22 w:%d", value)
	countwrites(value,22,20)
	l.Unlock()

	delay()
	l.Lock()
	value = 8
	s8.Put(126,value)
	log("p22 w:%d", value)
	countwrites(value,22,8)
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
	value = -1
	s31.QueryP(487,&value)
	log("p22 r:%d", value)
	countreads(value,22,31)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s8.QueryP(115,&value)
	log("p22 r:%d", value)
	countreads(value,22,8)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s24.QueryP(379,&value)
	log("p22 r:%d", value)
	countreads(value,22,24)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s24.QueryP(369,&value)
	log("p22 r:%d", value)
	countreads(value,22,24)
	l.Unlock()

	delay()
	l.Lock()
	value = 9
	s9.Put(132,value)
	log("p22 w:%d", value)
	countwrites(value,22,9)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s2.QueryP(23,&value)
	log("p22 r:%d", value)
	countreads(value,22,2)
	l.Unlock()
}


func p23() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	s24.QueryP(375,&value)
	log("p23 r:%d", value)
	countreads(value,23,24)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s18.QueryP(280,&value)
	log("p23 r:%d", value)
	countreads(value,23,18)
	l.Unlock()

	delay()
	l.Lock()
	value = 30
	s30.Put(474,value)
	log("p23 w:%d", value)
	countwrites(value,23,30)
	l.Unlock()

	delay()
	l.Lock()
	value = 30
	s30.Put(474,value)
	log("p23 w:%d", value)
	countwrites(value,23,30)
	l.Unlock()

	delay()
	l.Lock()
	value = 20
	s20.Put(306,value)
	log("p23 w:%d", value)
	countwrites(value,23,20)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s29.QueryP(458,&value)
	log("p23 r:%d", value)
	countreads(value,23,29)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s16.QueryP(249,&value)
	log("p23 r:%d", value)
	countreads(value,23,16)
	l.Unlock()

	delay()
	l.Lock()
	value = 20
	s20.Put(310,value)
	log("p23 w:%d", value)
	countwrites(value,23,20)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s3.QueryP(43,&value)
	log("p23 r:%d", value)
	countreads(value,23,3)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s24.QueryP(378,&value)
	log("p23 r:%d", value)
	countreads(value,23,24)
	l.Unlock()

	delay()
	l.Lock()
	value = 13
	s13.Put(194,value)
	log("p23 w:%d", value)
	countwrites(value,23,13)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s18.QueryP(273,&value)
	log("p23 r:%d", value)
	countreads(value,23,18)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s31.QueryP(493,&value)
	log("p23 r:%d", value)
	countreads(value,23,31)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s11.QueryP(169,&value)
	log("p23 r:%d", value)
	countreads(value,23,11)
	l.Unlock()

	delay()
	l.Lock()
	value = 8
	s8.Put(120,value)
	log("p23 w:%d", value)
	countwrites(value,23,8)
	l.Unlock()

	delay()
	l.Lock()
	value = 22
	s22.Put(344,value)
	log("p23 w:%d", value)
	countwrites(value,23,22)
	l.Unlock()
}


func p24() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	s28.QueryP(448,&value)
	log("p24 r:%d", value)
	countreads(value,24,28)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s26.QueryP(411,&value)
	log("p24 r:%d", value)
	countreads(value,24,26)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s5.QueryP(73,&value)
	log("p24 r:%d", value)
	countreads(value,24,5)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s7.QueryP(102,&value)
	log("p24 r:%d", value)
	countreads(value,24,7)
	l.Unlock()

	delay()
	l.Lock()
	value = 31
	s31.Put(490,value)
	log("p24 w:%d", value)
	countwrites(value,24,31)
	l.Unlock()

	delay()
	l.Lock()
	value = 9
	s9.Put(144,value)
	log("p24 w:%d", value)
	countwrites(value,24,9)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s15.QueryP(233,&value)
	log("p24 r:%d", value)
	countreads(value,24,15)
	l.Unlock()

	delay()
	l.Lock()
	value = 8
	s8.Put(115,value)
	log("p24 w:%d", value)
	countwrites(value,24,8)
	l.Unlock()

	delay()
	l.Lock()
	value = 21
	s21.Put(322,value)
	log("p24 w:%d", value)
	countwrites(value,24,21)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s27.QueryP(428,&value)
	log("p24 r:%d", value)
	countreads(value,24,27)
	l.Unlock()

	delay()
	l.Lock()
	value = 29
	s29.Put(452,value)
	log("p24 w:%d", value)
	countwrites(value,24,29)
	l.Unlock()

	delay()
	l.Lock()
	value = 32
	s32.Put(502,value)
	log("p24 w:%d", value)
	countwrites(value,24,32)
	l.Unlock()

	delay()
	l.Lock()
	value = 2
	s2.Put(23,value)
	log("p24 w:%d", value)
	countwrites(value,24,2)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s30.QueryP(472,&value)
	log("p24 r:%d", value)
	countreads(value,24,30)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s10.QueryP(148,&value)
	log("p24 r:%d", value)
	countreads(value,24,10)
	l.Unlock()

	delay()
	l.Lock()
	value = 1
	s1.Put(15,value)
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
	s30.QueryP(476,&value)
	log("p25 r:%d", value)
	countreads(value,25,30)
	l.Unlock()

	delay()
	l.Lock()
	value = 4
	s4.Put(63,value)
	log("p25 w:%d", value)
	countwrites(value,25,4)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s14.QueryP(220,&value)
	log("p25 r:%d", value)
	countreads(value,25,14)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s28.QueryP(443,&value)
	log("p25 r:%d", value)
	countreads(value,25,28)
	l.Unlock()

	delay()
	l.Lock()
	value = 9
	s9.Put(143,value)
	log("p25 w:%d", value)
	countwrites(value,25,9)
	l.Unlock()

	delay()
	l.Lock()
	value = 10
	s10.Put(149,value)
	log("p25 w:%d", value)
	countwrites(value,25,10)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s17.QueryP(260,&value)
	log("p25 r:%d", value)
	countreads(value,25,17)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s27.QueryP(421,&value)
	log("p25 r:%d", value)
	countreads(value,25,27)
	l.Unlock()

	delay()
	l.Lock()
	value = 25
	s25.Put(393,value)
	log("p25 w:%d", value)
	countwrites(value,25,25)
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
	s15.QueryP(235,&value)
	log("p25 r:%d", value)
	countreads(value,25,15)
	l.Unlock()

	delay()
	l.Lock()
	value = 28
	s28.Put(448,value)
	log("p25 w:%d", value)
	countwrites(value,25,28)
	l.Unlock()

	delay()
	l.Lock()
	value = 23
	s23.Put(358,value)
	log("p25 w:%d", value)
	countwrites(value,25,23)
	l.Unlock()

	delay()
	l.Lock()
	value = 32
	s32.Put(498,value)
	log("p25 w:%d", value)
	countwrites(value,25,32)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s31.QueryP(486,&value)
	log("p25 r:%d", value)
	countreads(value,25,31)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s5.QueryP(79,&value)
	log("p25 r:%d", value)
	countreads(value,25,5)
	l.Unlock()
}


func p26() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 25
	s25.Put(398,value)
	log("p26 w:%d", value)
	countwrites(value,26,25)
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
	value = 7
	s7.Put(97,value)
	log("p26 w:%d", value)
	countwrites(value,26,7)
	l.Unlock()

	delay()
	l.Lock()
	value = 20
	s20.Put(309,value)
	log("p26 w:%d", value)
	countwrites(value,26,20)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s24.QueryP(369,&value)
	log("p26 r:%d", value)
	countreads(value,26,24)
	l.Unlock()

	delay()
	l.Lock()
	value = 27
	s27.Put(428,value)
	log("p26 w:%d", value)
	countwrites(value,26,27)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s25.QueryP(396,&value)
	log("p26 r:%d", value)
	countreads(value,26,25)
	l.Unlock()

	delay()
	l.Lock()
	value = 14
	s14.Put(210,value)
	log("p26 w:%d", value)
	countwrites(value,26,14)
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
	value = 5
	s5.Put(66,value)
	log("p26 w:%d", value)
	countwrites(value,26,5)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s3.QueryP(45,&value)
	log("p26 r:%d", value)
	countreads(value,26,3)
	l.Unlock()

	delay()
	l.Lock()
	value = 30
	s30.Put(470,value)
	log("p26 w:%d", value)
	countwrites(value,26,30)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s24.QueryP(379,&value)
	log("p26 r:%d", value)
	countreads(value,26,24)
	l.Unlock()

	delay()
	l.Lock()
	value = 16
	s16.Put(256,value)
	log("p26 w:%d", value)
	countwrites(value,26,16)
	l.Unlock()

	delay()
	l.Lock()
	value = 4
	s4.Put(59,value)
	log("p26 w:%d", value)
	countwrites(value,26,4)
	l.Unlock()

	delay()
	l.Lock()
	value = 10
	s10.Put(146,value)
	log("p26 w:%d", value)
	countwrites(value,26,10)
	l.Unlock()
}


func p27() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	s32.QueryP(511,&value)
	log("p27 r:%d", value)
	countreads(value,27,32)
	l.Unlock()

	delay()
	l.Lock()
	value = 30
	s30.Put(473,value)
	log("p27 w:%d", value)
	countwrites(value,27,30)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s24.QueryP(369,&value)
	log("p27 r:%d", value)
	countreads(value,27,24)
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
	value = -1
	s22.QueryP(348,&value)
	log("p27 r:%d", value)
	countreads(value,27,22)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s23.QueryP(360,&value)
	log("p27 r:%d", value)
	countreads(value,27,23)
	l.Unlock()

	delay()
	l.Lock()
	value = 10
	s10.Put(152,value)
	log("p27 w:%d", value)
	countwrites(value,27,10)
	l.Unlock()

	delay()
	l.Lock()
	value = 6
	s6.Put(82,value)
	log("p27 w:%d", value)
	countwrites(value,27,6)
	l.Unlock()

	delay()
	l.Lock()
	value = 4
	s4.Put(60,value)
	log("p27 w:%d", value)
	countwrites(value,27,4)
	l.Unlock()

	delay()
	l.Lock()
	value = 31
	s31.Put(484,value)
	log("p27 w:%d", value)
	countwrites(value,27,31)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s10.QueryP(147,&value)
	log("p27 r:%d", value)
	countreads(value,27,10)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s10.QueryP(147,&value)
	log("p27 r:%d", value)
	countreads(value,27,10)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s7.QueryP(107,&value)
	log("p27 r:%d", value)
	countreads(value,27,7)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s3.QueryP(44,&value)
	log("p27 r:%d", value)
	countreads(value,27,3)
	l.Unlock()

	delay()
	l.Lock()
	value = 30
	s30.Put(473,value)
	log("p27 w:%d", value)
	countwrites(value,27,30)
	l.Unlock()

	delay()
	l.Lock()
	value = 32
	s32.Put(503,value)
	log("p27 w:%d", value)
	countwrites(value,27,32)
	l.Unlock()
}


func p28() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 8
	s8.Put(122,value)
	log("p28 w:%d", value)
	countwrites(value,28,8)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s22.QueryP(341,&value)
	log("p28 r:%d", value)
	countreads(value,28,22)
	l.Unlock()

	delay()
	l.Lock()
	value = 19
	s19.Put(297,value)
	log("p28 w:%d", value)
	countwrites(value,28,19)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s10.QueryP(159,&value)
	log("p28 r:%d", value)
	countreads(value,28,10)
	l.Unlock()

	delay()
	l.Lock()
	value = 18
	s18.Put(277,value)
	log("p28 w:%d", value)
	countwrites(value,28,18)
	l.Unlock()

	delay()
	l.Lock()
	value = 4
	s4.Put(57,value)
	log("p28 w:%d", value)
	countwrites(value,28,4)
	l.Unlock()

	delay()
	l.Lock()
	value = 12
	s12.Put(186,value)
	log("p28 w:%d", value)
	countwrites(value,28,12)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s31.QueryP(491,&value)
	log("p28 r:%d", value)
	countreads(value,28,31)
	l.Unlock()

	delay()
	l.Lock()
	value = 31
	s31.Put(481,value)
	log("p28 w:%d", value)
	countwrites(value,28,31)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s7.QueryP(104,&value)
	log("p28 r:%d", value)
	countreads(value,28,7)
	l.Unlock()

	delay()
	l.Lock()
	value = 7
	s7.Put(97,value)
	log("p28 w:%d", value)
	countwrites(value,28,7)
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
	value = -1
	s14.QueryP(210,&value)
	log("p28 r:%d", value)
	countreads(value,28,14)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s1.QueryP(13,&value)
	log("p28 r:%d", value)
	countreads(value,28,1)
	l.Unlock()

	delay()
	l.Lock()
	value = 13
	s13.Put(207,value)
	log("p28 w:%d", value)
	countwrites(value,28,13)
	l.Unlock()

	delay()
	l.Lock()
	value = 24
	s24.Put(383,value)
	log("p28 w:%d", value)
	countwrites(value,28,24)
	l.Unlock()
}


func p29() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 20
	s20.Put(310,value)
	log("p29 w:%d", value)
	countwrites(value,29,20)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s18.QueryP(279,&value)
	log("p29 r:%d", value)
	countreads(value,29,18)
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
	value = 30
	s30.Put(467,value)
	log("p29 w:%d", value)
	countwrites(value,29,30)
	l.Unlock()

	delay()
	l.Lock()
	value = 26
	s26.Put(404,value)
	log("p29 w:%d", value)
	countwrites(value,29,26)
	l.Unlock()

	delay()
	l.Lock()
	value = 28
	s28.Put(444,value)
	log("p29 w:%d", value)
	countwrites(value,29,28)
	l.Unlock()

	delay()
	l.Lock()
	value = 3
	s3.Put(47,value)
	log("p29 w:%d", value)
	countwrites(value,29,3)
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
	s27.QueryP(423,&value)
	log("p29 r:%d", value)
	countreads(value,29,27)
	l.Unlock()

	delay()
	l.Lock()
	value = 21
	s21.Put(328,value)
	log("p29 w:%d", value)
	countwrites(value,29,21)
	l.Unlock()

	delay()
	l.Lock()
	value = 25
	s25.Put(385,value)
	log("p29 w:%d", value)
	countwrites(value,29,25)
	l.Unlock()

	delay()
	l.Lock()
	value = 3
	s3.Put(46,value)
	log("p29 w:%d", value)
	countwrites(value,29,3)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s22.QueryP(341,&value)
	log("p29 r:%d", value)
	countreads(value,29,22)
	l.Unlock()

	delay()
	l.Lock()
	value = 29
	s29.Put(461,value)
	log("p29 w:%d", value)
	countwrites(value,29,29)
	l.Unlock()

	delay()
	l.Lock()
	value = 30
	s30.Put(476,value)
	log("p29 w:%d", value)
	countwrites(value,29,30)
	l.Unlock()

	delay()
	l.Lock()
	value = 15
	s15.Put(239,value)
	log("p29 w:%d", value)
	countwrites(value,29,15)
	l.Unlock()
}


func p30() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	s15.QueryP(234,&value)
	log("p30 r:%d", value)
	countreads(value,30,15)
	l.Unlock()

	delay()
	l.Lock()
	value = 21
	s21.Put(335,value)
	log("p30 w:%d", value)
	countwrites(value,30,21)
	l.Unlock()

	delay()
	l.Lock()
	value = 6
	s6.Put(90,value)
	log("p30 w:%d", value)
	countwrites(value,30,6)
	l.Unlock()

	delay()
	l.Lock()
	value = 27
	s27.Put(429,value)
	log("p30 w:%d", value)
	countwrites(value,30,27)
	l.Unlock()

	delay()
	l.Lock()
	value = 30
	s30.Put(469,value)
	log("p30 w:%d", value)
	countwrites(value,30,30)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s17.QueryP(265,&value)
	log("p30 r:%d", value)
	countreads(value,30,17)
	l.Unlock()

	delay()
	l.Lock()
	value = 32
	s32.Put(512,value)
	log("p30 w:%d", value)
	countwrites(value,30,32)
	l.Unlock()

	delay()
	l.Lock()
	value = 24
	s24.Put(373,value)
	log("p30 w:%d", value)
	countwrites(value,30,24)
	l.Unlock()

	delay()
	l.Lock()
	value = 10
	s10.Put(149,value)
	log("p30 w:%d", value)
	countwrites(value,30,10)
	l.Unlock()

	delay()
	l.Lock()
	value = 7
	s7.Put(101,value)
	log("p30 w:%d", value)
	countwrites(value,30,7)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s25.QueryP(389,&value)
	log("p30 r:%d", value)
	countreads(value,30,25)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s18.QueryP(278,&value)
	log("p30 r:%d", value)
	countreads(value,30,18)
	l.Unlock()

	delay()
	l.Lock()
	value = 32
	s32.Put(507,value)
	log("p30 w:%d", value)
	countwrites(value,30,32)
	l.Unlock()

	delay()
	l.Lock()
	value = 26
	s26.Put(411,value)
	log("p30 w:%d", value)
	countwrites(value,30,26)
	l.Unlock()

	delay()
	l.Lock()
	value = 9
	s9.Put(131,value)
	log("p30 w:%d", value)
	countwrites(value,30,9)
	l.Unlock()

	delay()
	l.Lock()
	value = 14
	s14.Put(217,value)
	log("p30 w:%d", value)
	countwrites(value,30,14)
	l.Unlock()
}


func p31() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	s12.QueryP(188,&value)
	log("p31 r:%d", value)
	countreads(value,31,12)
	l.Unlock()

	delay()
	l.Lock()
	value = 4
	s4.Put(51,value)
	log("p31 w:%d", value)
	countwrites(value,31,4)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s9.QueryP(142,&value)
	log("p31 r:%d", value)
	countreads(value,31,9)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s15.QueryP(238,&value)
	log("p31 r:%d", value)
	countreads(value,31,15)
	l.Unlock()

	delay()
	l.Lock()
	value = 2
	s2.Put(17,value)
	log("p31 w:%d", value)
	countwrites(value,31,2)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s8.QueryP(119,&value)
	log("p31 r:%d", value)
	countreads(value,31,8)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s20.QueryP(317,&value)
	log("p31 r:%d", value)
	countreads(value,31,20)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s24.QueryP(374,&value)
	log("p31 r:%d", value)
	countreads(value,31,24)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s11.QueryP(167,&value)
	log("p31 r:%d", value)
	countreads(value,31,11)
	l.Unlock()

	delay()
	l.Lock()
	value = 24
	s24.Put(373,value)
	log("p31 w:%d", value)
	countwrites(value,31,24)
	l.Unlock()

	delay()
	l.Lock()
	value = 24
	s24.Put(374,value)
	log("p31 w:%d", value)
	countwrites(value,31,24)
	l.Unlock()

	delay()
	l.Lock()
	value = 13
	s13.Put(202,value)
	log("p31 w:%d", value)
	countwrites(value,31,13)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s2.QueryP(17,&value)
	log("p31 r:%d", value)
	countreads(value,31,2)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s25.QueryP(397,&value)
	log("p31 r:%d", value)
	countreads(value,31,25)
	l.Unlock()

	delay()
	l.Lock()
	value = 27
	s27.Put(425,value)
	log("p31 w:%d", value)
	countwrites(value,31,27)
	l.Unlock()

	delay()
	l.Lock()
	value = 11
	s11.Put(167,value)
	log("p31 w:%d", value)
	countwrites(value,31,11)
	l.Unlock()
}


func p32() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 26
	s26.Put(415,value)
	log("p32 w:%d", value)
	countwrites(value,32,26)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s27.QueryP(422,&value)
	log("p32 r:%d", value)
	countreads(value,32,27)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s9.QueryP(138,&value)
	log("p32 r:%d", value)
	countreads(value,32,9)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s32.QueryP(503,&value)
	log("p32 r:%d", value)
	countreads(value,32,32)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s11.QueryP(174,&value)
	log("p32 r:%d", value)
	countreads(value,32,11)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s29.QueryP(458,&value)
	log("p32 r:%d", value)
	countreads(value,32,29)
	l.Unlock()

	delay()
	l.Lock()
	value = 30
	s30.Put(471,value)
	log("p32 w:%d", value)
	countwrites(value,32,30)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s31.QueryP(493,&value)
	log("p32 r:%d", value)
	countreads(value,32,31)
	l.Unlock()

	delay()
	l.Lock()
	value = 15
	s15.Put(239,value)
	log("p32 w:%d", value)
	countwrites(value,32,15)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s18.QueryP(285,&value)
	log("p32 r:%d", value)
	countreads(value,32,18)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s29.QueryP(452,&value)
	log("p32 r:%d", value)
	countreads(value,32,29)
	l.Unlock()

	delay()
	l.Lock()
	value = 11
	s11.Put(173,value)
	log("p32 w:%d", value)
	countwrites(value,32,11)
	l.Unlock()

	delay()
	l.Lock()
	value = 27
	s27.Put(420,value)
	log("p32 w:%d", value)
	countwrites(value,32,27)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s23.QueryP(359,&value)
	log("p32 r:%d", value)
	countreads(value,32,23)
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
	value = 8
	s8.Put(118,value)
	log("p32 w:%d", value)
	countwrites(value,32,8)
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
