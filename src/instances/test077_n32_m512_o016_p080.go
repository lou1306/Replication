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
	s29.QueryP(454,&value)
	log("p1 r:%d", value)
	countreads(value,1,29)
	l.Unlock()

	delay()
	l.Lock()
	value = 5
	s5.Put(76,value)
	log("p1 w:%d", value)
	countwrites(value,1,5)
	l.Unlock()

	delay()
	l.Lock()
	value = 16
	s16.Put(249,value)
	log("p1 w:%d", value)
	countwrites(value,1,16)
	l.Unlock()

	delay()
	l.Lock()
	value = 32
	s32.Put(501,value)
	log("p1 w:%d", value)
	countwrites(value,1,32)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s24.QueryP(369,&value)
	log("p1 r:%d", value)
	countreads(value,1,24)
	l.Unlock()

	delay()
	l.Lock()
	value = 7
	s7.Put(106,value)
	log("p1 w:%d", value)
	countwrites(value,1,7)
	l.Unlock()

	delay()
	l.Lock()
	value = 3
	s3.Put(41,value)
	log("p1 w:%d", value)
	countwrites(value,1,3)
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
	value = 17
	s17.Put(266,value)
	log("p1 w:%d", value)
	countwrites(value,1,17)
	l.Unlock()

	delay()
	l.Lock()
	value = 13
	s13.Put(200,value)
	log("p1 w:%d", value)
	countwrites(value,1,13)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s18.QueryP(282,&value)
	log("p1 r:%d", value)
	countreads(value,1,18)
	l.Unlock()

	delay()
	l.Lock()
	value = 3
	s3.Put(33,value)
	log("p1 w:%d", value)
	countwrites(value,1,3)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s14.QueryP(224,&value)
	log("p1 r:%d", value)
	countreads(value,1,14)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s11.QueryP(164,&value)
	log("p1 r:%d", value)
	countreads(value,1,11)
	l.Unlock()

	delay()
	l.Lock()
	value = 30
	s30.Put(470,value)
	log("p1 w:%d", value)
	countwrites(value,1,30)
	l.Unlock()

	delay()
	l.Lock()
	value = 27
	s27.Put(430,value)
	log("p1 w:%d", value)
	countwrites(value,1,27)
	l.Unlock()
}

func p2() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 24
	s24.Put(380,value)
	log("p2 w:%d", value)
	countwrites(value,2,24)
	l.Unlock()

	delay()
	l.Lock()
	value = 7
	s7.Put(103,value)
	log("p2 w:%d", value)
	countwrites(value,2,7)
	l.Unlock()

	delay()
	l.Lock()
	value = 18
	s18.Put(281,value)
	log("p2 w:%d", value)
	countwrites(value,2,18)
	l.Unlock()

	delay()
	l.Lock()
	value = 5
	s5.Put(66,value)
	log("p2 w:%d", value)
	countwrites(value,2,5)
	l.Unlock()

	delay()
	l.Lock()
	value = 1
	s1.Put(3,value)
	log("p2 w:%d", value)
	countwrites(value,2,1)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s27.QueryP(421,&value)
	log("p2 r:%d", value)
	countreads(value,2,27)
	l.Unlock()

	delay()
	l.Lock()
	value = 7
	s7.Put(102,value)
	log("p2 w:%d", value)
	countwrites(value,2,7)
	l.Unlock()

	delay()
	l.Lock()
	value = 18
	s18.Put(280,value)
	log("p2 w:%d", value)
	countwrites(value,2,18)
	l.Unlock()

	delay()
	l.Lock()
	value = 31
	s31.Put(487,value)
	log("p2 w:%d", value)
	countwrites(value,2,31)
	l.Unlock()

	delay()
	l.Lock()
	value = 31
	s31.Put(493,value)
	log("p2 w:%d", value)
	countwrites(value,2,31)
	l.Unlock()

	delay()
	l.Lock()
	value = 14
	s14.Put(219,value)
	log("p2 w:%d", value)
	countwrites(value,2,14)
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
	value = 22
	s22.Put(352,value)
	log("p2 w:%d", value)
	countwrites(value,2,22)
	l.Unlock()

	delay()
	l.Lock()
	value = 17
	s17.Put(259,value)
	log("p2 w:%d", value)
	countwrites(value,2,17)
	l.Unlock()

	delay()
	l.Lock()
	value = 21
	s21.Put(324,value)
	log("p2 w:%d", value)
	countwrites(value,2,21)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s27.QueryP(423,&value)
	log("p2 r:%d", value)
	countreads(value,2,27)
	l.Unlock()
}

func p3() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 25
	s25.Put(396,value)
	log("p3 w:%d", value)
	countwrites(value,3,25)
	l.Unlock()

	delay()
	l.Lock()
	value = 8
	s8.Put(121,value)
	log("p3 w:%d", value)
	countwrites(value,3,8)
	l.Unlock()

	delay()
	l.Lock()
	value = 23
	s23.Put(361,value)
	log("p3 w:%d", value)
	countwrites(value,3,23)
	l.Unlock()

	delay()
	l.Lock()
	value = 26
	s26.Put(408,value)
	log("p3 w:%d", value)
	countwrites(value,3,26)
	l.Unlock()

	delay()
	l.Lock()
	value = 32
	s32.Put(501,value)
	log("p3 w:%d", value)
	countwrites(value,3,32)
	l.Unlock()

	delay()
	l.Lock()
	value = 25
	s25.Put(390,value)
	log("p3 w:%d", value)
	countwrites(value,3,25)
	l.Unlock()

	delay()
	l.Lock()
	value = 7
	s7.Put(101,value)
	log("p3 w:%d", value)
	countwrites(value,3,7)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s8.QueryP(123,&value)
	log("p3 r:%d", value)
	countreads(value,3,8)
	l.Unlock()

	delay()
	l.Lock()
	value = 2
	s2.Put(24,value)
	log("p3 w:%d", value)
	countwrites(value,3,2)
	l.Unlock()

	delay()
	l.Lock()
	value = 19
	s19.Put(302,value)
	log("p3 w:%d", value)
	countwrites(value,3,19)
	l.Unlock()

	delay()
	l.Lock()
	value = 32
	s32.Put(500,value)
	log("p3 w:%d", value)
	countwrites(value,3,32)
	l.Unlock()

	delay()
	l.Lock()
	value = 23
	s23.Put(367,value)
	log("p3 w:%d", value)
	countwrites(value,3,23)
	l.Unlock()

	delay()
	l.Lock()
	value = 2
	s2.Put(32,value)
	log("p3 w:%d", value)
	countwrites(value,3,2)
	l.Unlock()

	delay()
	l.Lock()
	value = 17
	s17.Put(262,value)
	log("p3 w:%d", value)
	countwrites(value,3,17)
	l.Unlock()

	delay()
	l.Lock()
	value = 13
	s13.Put(204,value)
	log("p3 w:%d", value)
	countwrites(value,3,13)
	l.Unlock()

	delay()
	l.Lock()
	value = 5
	s5.Put(70,value)
	log("p3 w:%d", value)
	countwrites(value,3,5)
	l.Unlock()
}

func p4() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 32
	s32.Put(508,value)
	log("p4 w:%d", value)
	countwrites(value,4,32)
	l.Unlock()

	delay()
	l.Lock()
	value = 20
	s20.Put(309,value)
	log("p4 w:%d", value)
	countwrites(value,4,20)
	l.Unlock()

	delay()
	l.Lock()
	value = 28
	s28.Put(441,value)
	log("p4 w:%d", value)
	countwrites(value,4,28)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s24.QueryP(373,&value)
	log("p4 r:%d", value)
	countreads(value,4,24)
	l.Unlock()

	delay()
	l.Lock()
	value = 18
	s18.Put(274,value)
	log("p4 w:%d", value)
	countwrites(value,4,18)
	l.Unlock()

	delay()
	l.Lock()
	value = 6
	s6.Put(92,value)
	log("p4 w:%d", value)
	countwrites(value,4,6)
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
	value = 21
	s21.Put(335,value)
	log("p4 w:%d", value)
	countwrites(value,4,21)
	l.Unlock()

	delay()
	l.Lock()
	value = 23
	s23.Put(354,value)
	log("p4 w:%d", value)
	countwrites(value,4,23)
	l.Unlock()

	delay()
	l.Lock()
	value = 22
	s22.Put(344,value)
	log("p4 w:%d", value)
	countwrites(value,4,22)
	l.Unlock()

	delay()
	l.Lock()
	value = 29
	s29.Put(457,value)
	log("p4 w:%d", value)
	countwrites(value,4,29)
	l.Unlock()

	delay()
	l.Lock()
	value = 7
	s7.Put(110,value)
	log("p4 w:%d", value)
	countwrites(value,4,7)
	l.Unlock()

	delay()
	l.Lock()
	value = 4
	s4.Put(63,value)
	log("p4 w:%d", value)
	countwrites(value,4,4)
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
	value = 4
	s4.Put(55,value)
	log("p4 w:%d", value)
	countwrites(value,4,4)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s12.QueryP(177,&value)
	log("p4 r:%d", value)
	countreads(value,4,12)
	l.Unlock()
}

func p5() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 31
	s31.Put(489,value)
	log("p5 w:%d", value)
	countwrites(value,5,31)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s26.QueryP(410,&value)
	log("p5 r:%d", value)
	countreads(value,5,26)
	l.Unlock()

	delay()
	l.Lock()
	value = 27
	s27.Put(432,value)
	log("p5 w:%d", value)
	countwrites(value,5,27)
	l.Unlock()

	delay()
	l.Lock()
	value = 24
	s24.Put(384,value)
	log("p5 w:%d", value)
	countwrites(value,5,24)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s18.QueryP(275,&value)
	log("p5 r:%d", value)
	countreads(value,5,18)
	l.Unlock()

	delay()
	l.Lock()
	value = 6
	s6.Put(88,value)
	log("p5 w:%d", value)
	countwrites(value,5,6)
	l.Unlock()

	delay()
	l.Lock()
	value = 12
	s12.Put(187,value)
	log("p5 w:%d", value)
	countwrites(value,5,12)
	l.Unlock()

	delay()
	l.Lock()
	value = 22
	s22.Put(337,value)
	log("p5 w:%d", value)
	countwrites(value,5,22)
	l.Unlock()

	delay()
	l.Lock()
	value = 18
	s18.Put(285,value)
	log("p5 w:%d", value)
	countwrites(value,5,18)
	l.Unlock()

	delay()
	l.Lock()
	value = 24
	s24.Put(371,value)
	log("p5 w:%d", value)
	countwrites(value,5,24)
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
	value = 32
	s32.Put(505,value)
	log("p5 w:%d", value)
	countwrites(value,5,32)
	l.Unlock()

	delay()
	l.Lock()
	value = 23
	s23.Put(362,value)
	log("p5 w:%d", value)
	countwrites(value,5,23)
	l.Unlock()

	delay()
	l.Lock()
	value = 31
	s31.Put(494,value)
	log("p5 w:%d", value)
	countwrites(value,5,31)
	l.Unlock()

	delay()
	l.Lock()
	value = 5
	s5.Put(67,value)
	log("p5 w:%d", value)
	countwrites(value,5,5)
	l.Unlock()

	delay()
	l.Lock()
	value = 9
	s9.Put(138,value)
	log("p5 w:%d", value)
	countwrites(value,5,9)
	l.Unlock()
}

func p6() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	s28.QueryP(439,&value)
	log("p6 r:%d", value)
	countreads(value,6,28)
	l.Unlock()

	delay()
	l.Lock()
	value = 10
	s10.Put(160,value)
	log("p6 w:%d", value)
	countwrites(value,6,10)
	l.Unlock()

	delay()
	l.Lock()
	value = 22
	s22.Put(350,value)
	log("p6 w:%d", value)
	countwrites(value,6,22)
	l.Unlock()

	delay()
	l.Lock()
	value = 4
	s4.Put(57,value)
	log("p6 w:%d", value)
	countwrites(value,6,4)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s10.QueryP(151,&value)
	log("p6 r:%d", value)
	countreads(value,6,10)
	l.Unlock()

	delay()
	l.Lock()
	value = 12
	s12.Put(188,value)
	log("p6 w:%d", value)
	countwrites(value,6,12)
	l.Unlock()

	delay()
	l.Lock()
	value = 24
	s24.Put(384,value)
	log("p6 w:%d", value)
	countwrites(value,6,24)
	l.Unlock()

	delay()
	l.Lock()
	value = 8
	s8.Put(120,value)
	log("p6 w:%d", value)
	countwrites(value,6,8)
	l.Unlock()

	delay()
	l.Lock()
	value = 18
	s18.Put(288,value)
	log("p6 w:%d", value)
	countwrites(value,6,18)
	l.Unlock()

	delay()
	l.Lock()
	value = 23
	s23.Put(360,value)
	log("p6 w:%d", value)
	countwrites(value,6,23)
	l.Unlock()

	delay()
	l.Lock()
	value = 28
	s28.Put(440,value)
	log("p6 w:%d", value)
	countwrites(value,6,28)
	l.Unlock()

	delay()
	l.Lock()
	value = 12
	s12.Put(183,value)
	log("p6 w:%d", value)
	countwrites(value,6,12)
	l.Unlock()

	delay()
	l.Lock()
	value = 28
	s28.Put(447,value)
	log("p6 w:%d", value)
	countwrites(value,6,28)
	l.Unlock()

	delay()
	l.Lock()
	value = 21
	s21.Put(323,value)
	log("p6 w:%d", value)
	countwrites(value,6,21)
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
	value = 21
	s21.Put(329,value)
	log("p6 w:%d", value)
	countwrites(value,6,21)
	l.Unlock()
}

func p7() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	s25.QueryP(397,&value)
	log("p7 r:%d", value)
	countreads(value,7,25)
	l.Unlock()

	delay()
	l.Lock()
	value = 2
	s2.Put(17,value)
	log("p7 w:%d", value)
	countwrites(value,7,2)
	l.Unlock()

	delay()
	l.Lock()
	value = 13
	s13.Put(206,value)
	log("p7 w:%d", value)
	countwrites(value,7,13)
	l.Unlock()

	delay()
	l.Lock()
	value = 25
	s25.Put(399,value)
	log("p7 w:%d", value)
	countwrites(value,7,25)
	l.Unlock()

	delay()
	l.Lock()
	value = 23
	s23.Put(356,value)
	log("p7 w:%d", value)
	countwrites(value,7,23)
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
	value = 2
	s2.Put(22,value)
	log("p7 w:%d", value)
	countwrites(value,7,2)
	l.Unlock()

	delay()
	l.Lock()
	value = 6
	s6.Put(88,value)
	log("p7 w:%d", value)
	countwrites(value,7,6)
	l.Unlock()

	delay()
	l.Lock()
	value = 29
	s29.Put(458,value)
	log("p7 w:%d", value)
	countwrites(value,7,29)
	l.Unlock()

	delay()
	l.Lock()
	value = 14
	s14.Put(224,value)
	log("p7 w:%d", value)
	countwrites(value,7,14)
	l.Unlock()

	delay()
	l.Lock()
	value = 30
	s30.Put(477,value)
	log("p7 w:%d", value)
	countwrites(value,7,30)
	l.Unlock()

	delay()
	l.Lock()
	value = 25
	s25.Put(387,value)
	log("p7 w:%d", value)
	countwrites(value,7,25)
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
	value = 19
	s19.Put(299,value)
	log("p7 w:%d", value)
	countwrites(value,7,19)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s22.QueryP(347,&value)
	log("p7 r:%d", value)
	countreads(value,7,22)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s13.QueryP(204,&value)
	log("p7 r:%d", value)
	countreads(value,7,13)
	l.Unlock()
}

func p8() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 13
	s13.Put(205,value)
	log("p8 w:%d", value)
	countwrites(value,8,13)
	l.Unlock()

	delay()
	l.Lock()
	value = 1
	s1.Put(14,value)
	log("p8 w:%d", value)
	countwrites(value,8,1)
	l.Unlock()

	delay()
	l.Lock()
	value = 13
	s13.Put(196,value)
	log("p8 w:%d", value)
	countwrites(value,8,13)
	l.Unlock()

	delay()
	l.Lock()
	value = 8
	s8.Put(120,value)
	log("p8 w:%d", value)
	countwrites(value,8,8)
	l.Unlock()

	delay()
	l.Lock()
	value = 4
	s4.Put(60,value)
	log("p8 w:%d", value)
	countwrites(value,8,4)
	l.Unlock()

	delay()
	l.Lock()
	value = 6
	s6.Put(95,value)
	log("p8 w:%d", value)
	countwrites(value,8,6)
	l.Unlock()

	delay()
	l.Lock()
	value = 23
	s23.Put(364,value)
	log("p8 w:%d", value)
	countwrites(value,8,23)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s31.QueryP(485,&value)
	log("p8 r:%d", value)
	countreads(value,8,31)
	l.Unlock()

	delay()
	l.Lock()
	value = 8
	s8.Put(127,value)
	log("p8 w:%d", value)
	countwrites(value,8,8)
	l.Unlock()

	delay()
	l.Lock()
	value = 3
	s3.Put(35,value)
	log("p8 w:%d", value)
	countwrites(value,8,3)
	l.Unlock()

	delay()
	l.Lock()
	value = 27
	s27.Put(417,value)
	log("p8 w:%d", value)
	countwrites(value,8,27)
	l.Unlock()

	delay()
	l.Lock()
	value = 6
	s6.Put(82,value)
	log("p8 w:%d", value)
	countwrites(value,8,6)
	l.Unlock()

	delay()
	l.Lock()
	value = 1
	s1.Put(16,value)
	log("p8 w:%d", value)
	countwrites(value,8,1)
	l.Unlock()

	delay()
	l.Lock()
	value = 27
	s27.Put(423,value)
	log("p8 w:%d", value)
	countwrites(value,8,27)
	l.Unlock()

	delay()
	l.Lock()
	value = 25
	s25.Put(393,value)
	log("p8 w:%d", value)
	countwrites(value,8,25)
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
	value = 19
	s19.Put(292,value)
	log("p9 w:%d", value)
	countwrites(value,9,19)
	l.Unlock()

	delay()
	l.Lock()
	value = 13
	s13.Put(207,value)
	log("p9 w:%d", value)
	countwrites(value,9,13)
	l.Unlock()

	delay()
	l.Lock()
	value = 8
	s8.Put(119,value)
	log("p9 w:%d", value)
	countwrites(value,9,8)
	l.Unlock()

	delay()
	l.Lock()
	value = 19
	s19.Put(289,value)
	log("p9 w:%d", value)
	countwrites(value,9,19)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s23.QueryP(360,&value)
	log("p9 r:%d", value)
	countreads(value,9,23)
	l.Unlock()

	delay()
	l.Lock()
	value = 32
	s32.Put(510,value)
	log("p9 w:%d", value)
	countwrites(value,9,32)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s21.QueryP(326,&value)
	log("p9 r:%d", value)
	countreads(value,9,21)
	l.Unlock()

	delay()
	l.Lock()
	value = 16
	s16.Put(254,value)
	log("p9 w:%d", value)
	countwrites(value,9,16)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s13.QueryP(202,&value)
	log("p9 r:%d", value)
	countreads(value,9,13)
	l.Unlock()

	delay()
	l.Lock()
	value = 1
	s1.Put(7,value)
	log("p9 w:%d", value)
	countwrites(value,9,1)
	l.Unlock()

	delay()
	l.Lock()
	value = 13
	s13.Put(199,value)
	log("p9 w:%d", value)
	countwrites(value,9,13)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s10.QueryP(151,&value)
	log("p9 r:%d", value)
	countreads(value,9,10)
	l.Unlock()

	delay()
	l.Lock()
	value = 14
	s14.Put(223,value)
	log("p9 w:%d", value)
	countwrites(value,9,14)
	l.Unlock()

	delay()
	l.Lock()
	value = 8
	s8.Put(125,value)
	log("p9 w:%d", value)
	countwrites(value,9,8)
	l.Unlock()

	delay()
	l.Lock()
	value = 1
	s1.Put(15,value)
	log("p9 w:%d", value)
	countwrites(value,9,1)
	l.Unlock()

	delay()
	l.Lock()
	value = 17
	s17.Put(259,value)
	log("p9 w:%d", value)
	countwrites(value,9,17)
	l.Unlock()
}


func p10() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 19
	s19.Put(291,value)
	log("p10 w:%d", value)
	countwrites(value,10,19)
	l.Unlock()

	delay()
	l.Lock()
	value = 29
	s29.Put(449,value)
	log("p10 w:%d", value)
	countwrites(value,10,29)
	l.Unlock()

	delay()
	l.Lock()
	value = 16
	s16.Put(244,value)
	log("p10 w:%d", value)
	countwrites(value,10,16)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s29.QueryP(464,&value)
	log("p10 r:%d", value)
	countreads(value,10,29)
	l.Unlock()

	delay()
	l.Lock()
	value = 3
	s3.Put(48,value)
	log("p10 w:%d", value)
	countwrites(value,10,3)
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
	value = 19
	s19.Put(298,value)
	log("p10 w:%d", value)
	countwrites(value,10,19)
	l.Unlock()

	delay()
	l.Lock()
	value = 5
	s5.Put(75,value)
	log("p10 w:%d", value)
	countwrites(value,10,5)
	l.Unlock()

	delay()
	l.Lock()
	value = 17
	s17.Put(266,value)
	log("p10 w:%d", value)
	countwrites(value,10,17)
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
	value = 17
	s17.Put(268,value)
	log("p10 w:%d", value)
	countwrites(value,10,17)
	l.Unlock()

	delay()
	l.Lock()
	value = 13
	s13.Put(196,value)
	log("p10 w:%d", value)
	countwrites(value,10,13)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s8.QueryP(118,&value)
	log("p10 r:%d", value)
	countreads(value,10,8)
	l.Unlock()

	delay()
	l.Lock()
	value = 13
	s13.Put(207,value)
	log("p10 w:%d", value)
	countwrites(value,10,13)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s7.QueryP(100,&value)
	log("p10 r:%d", value)
	countreads(value,10,7)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s15.QueryP(238,&value)
	log("p10 r:%d", value)
	countreads(value,10,15)
	l.Unlock()
}


func p11() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 22
	s22.Put(352,value)
	log("p11 w:%d", value)
	countwrites(value,11,22)
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
	value = 29
	s29.Put(462,value)
	log("p11 w:%d", value)
	countwrites(value,11,29)
	l.Unlock()

	delay()
	l.Lock()
	value = 21
	s21.Put(322,value)
	log("p11 w:%d", value)
	countwrites(value,11,21)
	l.Unlock()

	delay()
	l.Lock()
	value = 26
	s26.Put(411,value)
	log("p11 w:%d", value)
	countwrites(value,11,26)
	l.Unlock()

	delay()
	l.Lock()
	value = 27
	s27.Put(418,value)
	log("p11 w:%d", value)
	countwrites(value,11,27)
	l.Unlock()

	delay()
	l.Lock()
	value = 16
	s16.Put(243,value)
	log("p11 w:%d", value)
	countwrites(value,11,16)
	l.Unlock()

	delay()
	l.Lock()
	value = 30
	s30.Put(466,value)
	log("p11 w:%d", value)
	countwrites(value,11,30)
	l.Unlock()

	delay()
	l.Lock()
	value = 14
	s14.Put(212,value)
	log("p11 w:%d", value)
	countwrites(value,11,14)
	l.Unlock()

	delay()
	l.Lock()
	value = 18
	s18.Put(283,value)
	log("p11 w:%d", value)
	countwrites(value,11,18)
	l.Unlock()

	delay()
	l.Lock()
	value = 3
	s3.Put(36,value)
	log("p11 w:%d", value)
	countwrites(value,11,3)
	l.Unlock()

	delay()
	l.Lock()
	value = 6
	s6.Put(89,value)
	log("p11 w:%d", value)
	countwrites(value,11,6)
	l.Unlock()

	delay()
	l.Lock()
	value = 5
	s5.Put(65,value)
	log("p11 w:%d", value)
	countwrites(value,11,5)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s29.QueryP(450,&value)
	log("p11 r:%d", value)
	countreads(value,11,29)
	l.Unlock()

	delay()
	l.Lock()
	value = 31
	s31.Put(493,value)
	log("p11 w:%d", value)
	countwrites(value,11,31)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s31.QueryP(491,&value)
	log("p11 r:%d", value)
	countreads(value,11,31)
	l.Unlock()
}


func p12() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 7
	s7.Put(103,value)
	log("p12 w:%d", value)
	countwrites(value,12,7)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s27.QueryP(425,&value)
	log("p12 r:%d", value)
	countreads(value,12,27)
	l.Unlock()

	delay()
	l.Lock()
	value = 23
	s23.Put(362,value)
	log("p12 w:%d", value)
	countwrites(value,12,23)
	l.Unlock()

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
	s23.QueryP(359,&value)
	log("p12 r:%d", value)
	countreads(value,12,23)
	l.Unlock()

	delay()
	l.Lock()
	value = 19
	s19.Put(304,value)
	log("p12 w:%d", value)
	countwrites(value,12,19)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s14.QueryP(210,&value)
	log("p12 r:%d", value)
	countreads(value,12,14)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s10.QueryP(158,&value)
	log("p12 r:%d", value)
	countreads(value,12,10)
	l.Unlock()

	delay()
	l.Lock()
	value = 12
	s12.Put(177,value)
	log("p12 w:%d", value)
	countwrites(value,12,12)
	l.Unlock()

	delay()
	l.Lock()
	value = 25
	s25.Put(393,value)
	log("p12 w:%d", value)
	countwrites(value,12,25)
	l.Unlock()

	delay()
	l.Lock()
	value = 3
	s3.Put(41,value)
	log("p12 w:%d", value)
	countwrites(value,12,3)
	l.Unlock()

	delay()
	l.Lock()
	value = 6
	s6.Put(92,value)
	log("p12 w:%d", value)
	countwrites(value,12,6)
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
	value = 2
	s2.Put(20,value)
	log("p12 w:%d", value)
	countwrites(value,12,2)
	l.Unlock()

	delay()
	l.Lock()
	value = 4
	s4.Put(55,value)
	log("p12 w:%d", value)
	countwrites(value,12,4)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s9.QueryP(144,&value)
	log("p12 r:%d", value)
	countreads(value,12,9)
	l.Unlock()
}


func p13() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	s4.QueryP(57,&value)
	log("p13 r:%d", value)
	countreads(value,13,4)
	l.Unlock()

	delay()
	l.Lock()
	value = 15
	s15.Put(226,value)
	log("p13 w:%d", value)
	countwrites(value,13,15)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s19.QueryP(298,&value)
	log("p13 r:%d", value)
	countreads(value,13,19)
	l.Unlock()

	delay()
	l.Lock()
	value = 14
	s14.Put(219,value)
	log("p13 w:%d", value)
	countwrites(value,13,14)
	l.Unlock()

	delay()
	l.Lock()
	value = 2
	s2.Put(29,value)
	log("p13 w:%d", value)
	countwrites(value,13,2)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s10.QueryP(147,&value)
	log("p13 r:%d", value)
	countreads(value,13,10)
	l.Unlock()

	delay()
	l.Lock()
	value = 2
	s2.Put(31,value)
	log("p13 w:%d", value)
	countwrites(value,13,2)
	l.Unlock()

	delay()
	l.Lock()
	value = 2
	s2.Put(31,value)
	log("p13 w:%d", value)
	countwrites(value,13,2)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s19.QueryP(297,&value)
	log("p13 r:%d", value)
	countreads(value,13,19)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s23.QueryP(356,&value)
	log("p13 r:%d", value)
	countreads(value,13,23)
	l.Unlock()

	delay()
	l.Lock()
	value = 32
	s32.Put(500,value)
	log("p13 w:%d", value)
	countwrites(value,13,32)
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
	value = 22
	s22.Put(349,value)
	log("p13 w:%d", value)
	countwrites(value,13,22)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s27.QueryP(421,&value)
	log("p13 r:%d", value)
	countreads(value,13,27)
	l.Unlock()

	delay()
	l.Lock()
	value = 19
	s19.Put(290,value)
	log("p13 w:%d", value)
	countwrites(value,13,19)
	l.Unlock()

	delay()
	l.Lock()
	value = 6
	s6.Put(85,value)
	log("p13 w:%d", value)
	countwrites(value,13,6)
	l.Unlock()
}

func p14() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 10
	s10.Put(158,value)
	log("p14 w:%d", value)
	countwrites(value,14,10)
	l.Unlock()

	delay()
	l.Lock()
	value = 29
	s29.Put(451,value)
	log("p14 w:%d", value)
	countwrites(value,14,29)
	l.Unlock()

	delay()
	l.Lock()
	value = 21
	s21.Put(332,value)
	log("p14 w:%d", value)
	countwrites(value,14,21)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s24.QueryP(373,&value)
	log("p14 r:%d", value)
	countreads(value,14,24)
	l.Unlock()

	delay()
	l.Lock()
	value = 14
	s14.Put(213,value)
	log("p14 w:%d", value)
	countwrites(value,14,14)
	l.Unlock()

	delay()
	l.Lock()
	value = 27
	s27.Put(422,value)
	log("p14 w:%d", value)
	countwrites(value,14,27)
	l.Unlock()

	delay()
	l.Lock()
	value = 3
	s3.Put(46,value)
	log("p14 w:%d", value)
	countwrites(value,14,3)
	l.Unlock()

	delay()
	l.Lock()
	value = 4
	s4.Put(61,value)
	log("p14 w:%d", value)
	countwrites(value,14,4)
	l.Unlock()

	delay()
	l.Lock()
	value = 26
	s26.Put(415,value)
	log("p14 w:%d", value)
	countwrites(value,14,26)
	l.Unlock()

	delay()
	l.Lock()
	value = 14
	s14.Put(213,value)
	log("p14 w:%d", value)
	countwrites(value,14,14)
	l.Unlock()

	delay()
	l.Lock()
	value = 20
	s20.Put(317,value)
	log("p14 w:%d", value)
	countwrites(value,14,20)
	l.Unlock()

	delay()
	l.Lock()
	value = 13
	s13.Put(193,value)
	log("p14 w:%d", value)
	countwrites(value,14,13)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s21.QueryP(326,&value)
	log("p14 r:%d", value)
	countreads(value,14,21)
	l.Unlock()

	delay()
	l.Lock()
	value = 27
	s27.Put(426,value)
	log("p14 w:%d", value)
	countwrites(value,14,27)
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
	value = 21
	s21.Put(325,value)
	log("p14 w:%d", value)
	countwrites(value,14,21)
	l.Unlock()
}


func p15() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 7
	s7.Put(97,value)
	log("p15 w:%d", value)
	countwrites(value,15,7)
	l.Unlock()

	delay()
	l.Lock()
	value = 23
	s23.Put(353,value)
	log("p15 w:%d", value)
	countwrites(value,15,23)
	l.Unlock()

	delay()
	l.Lock()
	value = 8
	s8.Put(120,value)
	log("p15 w:%d", value)
	countwrites(value,15,8)
	l.Unlock()

	delay()
	l.Lock()
	value = 21
	s21.Put(336,value)
	log("p15 w:%d", value)
	countwrites(value,15,21)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s9.QueryP(142,&value)
	log("p15 r:%d", value)
	countreads(value,15,9)
	l.Unlock()

	delay()
	l.Lock()
	value = 29
	s29.Put(458,value)
	log("p15 w:%d", value)
	countwrites(value,15,29)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s17.QueryP(270,&value)
	log("p15 r:%d", value)
	countreads(value,15,17)
	l.Unlock()

	delay()
	l.Lock()
	value = 22
	s22.Put(337,value)
	log("p15 w:%d", value)
	countwrites(value,15,22)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s10.QueryP(157,&value)
	log("p15 r:%d", value)
	countreads(value,15,10)
	l.Unlock()

	delay()
	l.Lock()
	value = 4
	s4.Put(58,value)
	log("p15 w:%d", value)
	countwrites(value,15,4)
	l.Unlock()

	delay()
	l.Lock()
	value = 9
	s9.Put(139,value)
	log("p15 w:%d", value)
	countwrites(value,15,9)
	l.Unlock()

	delay()
	l.Lock()
	value = 29
	s29.Put(459,value)
	log("p15 w:%d", value)
	countwrites(value,15,29)
	l.Unlock()

	delay()
	l.Lock()
	value = 9
	s9.Put(135,value)
	log("p15 w:%d", value)
	countwrites(value,15,9)
	l.Unlock()

	delay()
	l.Lock()
	value = 16
	s16.Put(247,value)
	log("p15 w:%d", value)
	countwrites(value,15,16)
	l.Unlock()

	delay()
	l.Lock()
	value = 20
	s20.Put(319,value)
	log("p15 w:%d", value)
	countwrites(value,15,20)
	l.Unlock()

	delay()
	l.Lock()
	value = 21
	s21.Put(328,value)
	log("p15 w:%d", value)
	countwrites(value,15,21)
	l.Unlock()
}


func p16() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	s22.QueryP(339,&value)
	log("p16 r:%d", value)
	countreads(value,16,22)
	l.Unlock()

	delay()
	l.Lock()
	value = 15
	s15.Put(238,value)
	log("p16 w:%d", value)
	countwrites(value,16,15)
	l.Unlock()

	delay()
	l.Lock()
	value = 24
	s24.Put(379,value)
	log("p16 w:%d", value)
	countwrites(value,16,24)
	l.Unlock()

	delay()
	l.Lock()
	value = 23
	s23.Put(356,value)
	log("p16 w:%d", value)
	countwrites(value,16,23)
	l.Unlock()

	delay()
	l.Lock()
	value = 32
	s32.Put(500,value)
	log("p16 w:%d", value)
	countwrites(value,16,32)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s32.QueryP(499,&value)
	log("p16 r:%d", value)
	countreads(value,16,32)
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
	value = 19
	s19.Put(290,value)
	log("p16 w:%d", value)
	countwrites(value,16,19)
	l.Unlock()

	delay()
	l.Lock()
	value = 25
	s25.Put(393,value)
	log("p16 w:%d", value)
	countwrites(value,16,25)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s28.QueryP(442,&value)
	log("p16 r:%d", value)
	countreads(value,16,28)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s1.QueryP(8,&value)
	log("p16 r:%d", value)
	countreads(value,16,1)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s1.QueryP(5,&value)
	log("p16 r:%d", value)
	countreads(value,16,1)
	l.Unlock()

	delay()
	l.Lock()
	value = 23
	s23.Put(368,value)
	log("p16 w:%d", value)
	countwrites(value,16,23)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s2.QueryP(21,&value)
	log("p16 r:%d", value)
	countreads(value,16,2)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s23.QueryP(361,&value)
	log("p16 r:%d", value)
	countreads(value,16,23)
	l.Unlock()

	delay()
	l.Lock()
	value = 27
	s27.Put(423,value)
	log("p16 w:%d", value)
	countwrites(value,16,27)
	l.Unlock()
}


func p17() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 8
	s8.Put(117,value)
	log("p17 w:%d", value)
	countwrites(value,17,8)
	l.Unlock()

	delay()
	l.Lock()
	value = 5
	s5.Put(78,value)
	log("p17 w:%d", value)
	countwrites(value,17,5)
	l.Unlock()

	delay()
	l.Lock()
	value = 11
	s11.Put(173,value)
	log("p17 w:%d", value)
	countwrites(value,17,11)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s25.QueryP(396,&value)
	log("p17 r:%d", value)
	countreads(value,17,25)
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
	value = 10
	s10.Put(148,value)
	log("p17 w:%d", value)
	countwrites(value,17,10)
	l.Unlock()

	delay()
	l.Lock()
	value = 6
	s6.Put(81,value)
	log("p17 w:%d", value)
	countwrites(value,17,6)
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
	value = 5
	s5.Put(65,value)
	log("p17 w:%d", value)
	countwrites(value,17,5)
	l.Unlock()

	delay()
	l.Lock()
	value = 14
	s14.Put(215,value)
	log("p17 w:%d", value)
	countwrites(value,17,14)
	l.Unlock()

	delay()
	l.Lock()
	value = 18
	s18.Put(280,value)
	log("p17 w:%d", value)
	countwrites(value,17,18)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s2.QueryP(24,&value)
	log("p17 r:%d", value)
	countreads(value,17,2)
	l.Unlock()

	delay()
	l.Lock()
	value = 2
	s2.Put(26,value)
	log("p17 w:%d", value)
	countwrites(value,17,2)
	l.Unlock()

	delay()
	l.Lock()
	value = 17
	s17.Put(257,value)
	log("p17 w:%d", value)
	countwrites(value,17,17)
	l.Unlock()

	delay()
	l.Lock()
	value = 17
	s17.Put(262,value)
	log("p17 w:%d", value)
	countwrites(value,17,17)
	l.Unlock()

	delay()
	l.Lock()
	value = 10
	s10.Put(147,value)
	log("p17 w:%d", value)
	countwrites(value,17,10)
	l.Unlock()
}


func p18() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 29
	s29.Put(458,value)
	log("p18 w:%d", value)
	countwrites(value,18,29)
	l.Unlock()

	delay()
	l.Lock()
	value = 19
	s19.Put(297,value)
	log("p18 w:%d", value)
	countwrites(value,18,19)
	l.Unlock()

	delay()
	l.Lock()
	value = 20
	s20.Put(311,value)
	log("p18 w:%d", value)
	countwrites(value,18,20)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s17.QueryP(262,&value)
	log("p18 r:%d", value)
	countreads(value,18,17)
	l.Unlock()

	delay()
	l.Lock()
	value = 10
	s10.Put(151,value)
	log("p18 w:%d", value)
	countwrites(value,18,10)
	l.Unlock()

	delay()
	l.Lock()
	value = 11
	s11.Put(168,value)
	log("p18 w:%d", value)
	countwrites(value,18,11)
	l.Unlock()

	delay()
	l.Lock()
	value = 13
	s13.Put(196,value)
	log("p18 w:%d", value)
	countwrites(value,18,13)
	l.Unlock()

	delay()
	l.Lock()
	value = 23
	s23.Put(365,value)
	log("p18 w:%d", value)
	countwrites(value,18,23)
	l.Unlock()

	delay()
	l.Lock()
	value = 6
	s6.Put(87,value)
	log("p18 w:%d", value)
	countwrites(value,18,6)
	l.Unlock()

	delay()
	l.Lock()
	value = 16
	s16.Put(243,value)
	log("p18 w:%d", value)
	countwrites(value,18,16)
	l.Unlock()

	delay()
	l.Lock()
	value = 2
	s2.Put(30,value)
	log("p18 w:%d", value)
	countwrites(value,18,2)
	l.Unlock()

	delay()
	l.Lock()
	value = 22
	s22.Put(348,value)
	log("p18 w:%d", value)
	countwrites(value,18,22)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s12.QueryP(180,&value)
	log("p18 r:%d", value)
	countreads(value,18,12)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s9.QueryP(132,&value)
	log("p18 r:%d", value)
	countreads(value,18,9)
	l.Unlock()

	delay()
	l.Lock()
	value = 16
	s16.Put(241,value)
	log("p18 w:%d", value)
	countwrites(value,18,16)
	l.Unlock()

	delay()
	l.Lock()
	value = 13
	s13.Put(206,value)
	log("p18 w:%d", value)
	countwrites(value,18,13)
	l.Unlock()
}


func p19() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 10
	s10.Put(153,value)
	log("p19 w:%d", value)
	countwrites(value,19,10)
	l.Unlock()

	delay()
	l.Lock()
	value = 3
	s3.Put(40,value)
	log("p19 w:%d", value)
	countwrites(value,19,3)
	l.Unlock()

	delay()
	l.Lock()
	value = 1
	s1.Put(1,value)
	log("p19 w:%d", value)
	countwrites(value,19,1)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s14.QueryP(224,&value)
	log("p19 r:%d", value)
	countreads(value,19,14)
	l.Unlock()

	delay()
	l.Lock()
	value = 32
	s32.Put(498,value)
	log("p19 w:%d", value)
	countwrites(value,19,32)
	l.Unlock()

	delay()
	l.Lock()
	value = 28
	s28.Put(433,value)
	log("p19 w:%d", value)
	countwrites(value,19,28)
	l.Unlock()

	delay()
	l.Lock()
	value = 12
	s12.Put(190,value)
	log("p19 w:%d", value)
	countwrites(value,19,12)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s31.QueryP(488,&value)
	log("p19 r:%d", value)
	countreads(value,19,31)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s5.QueryP(74,&value)
	log("p19 r:%d", value)
	countreads(value,19,5)
	l.Unlock()

	delay()
	l.Lock()
	value = 27
	s27.Put(425,value)
	log("p19 w:%d", value)
	countwrites(value,19,27)
	l.Unlock()

	delay()
	l.Lock()
	value = 10
	s10.Put(156,value)
	log("p19 w:%d", value)
	countwrites(value,19,10)
	l.Unlock()

	delay()
	l.Lock()
	value = 7
	s7.Put(112,value)
	log("p19 w:%d", value)
	countwrites(value,19,7)
	l.Unlock()

	delay()
	l.Lock()
	value = 12
	s12.Put(191,value)
	log("p19 w:%d", value)
	countwrites(value,19,12)
	l.Unlock()

	delay()
	l.Lock()
	value = 30
	s30.Put(472,value)
	log("p19 w:%d", value)
	countwrites(value,19,30)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s7.QueryP(109,&value)
	log("p19 r:%d", value)
	countreads(value,19,7)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s7.QueryP(107,&value)
	log("p19 r:%d", value)
	countreads(value,19,7)
	l.Unlock()
}


func p20() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 18
	s18.Put(288,value)
	log("p20 w:%d", value)
	countwrites(value,20,18)
	l.Unlock()

	delay()
	l.Lock()
	value = 7
	s7.Put(104,value)
	log("p20 w:%d", value)
	countwrites(value,20,7)
	l.Unlock()

	delay()
	l.Lock()
	value = 6
	s6.Put(95,value)
	log("p20 w:%d", value)
	countwrites(value,20,6)
	l.Unlock()

	delay()
	l.Lock()
	value = 12
	s12.Put(190,value)
	log("p20 w:%d", value)
	countwrites(value,20,12)
	l.Unlock()

	delay()
	l.Lock()
	value = 27
	s27.Put(431,value)
	log("p20 w:%d", value)
	countwrites(value,20,27)
	l.Unlock()

	delay()
	l.Lock()
	value = 11
	s11.Put(175,value)
	log("p20 w:%d", value)
	countwrites(value,20,11)
	l.Unlock()

	delay()
	l.Lock()
	value = 20
	s20.Put(309,value)
	log("p20 w:%d", value)
	countwrites(value,20,20)
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
	value = 15
	s15.Put(234,value)
	log("p20 w:%d", value)
	countwrites(value,20,15)
	l.Unlock()

	delay()
	l.Lock()
	value = 5
	s5.Put(66,value)
	log("p20 w:%d", value)
	countwrites(value,20,5)
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
	value = 20
	s20.Put(314,value)
	log("p20 w:%d", value)
	countwrites(value,20,20)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s11.QueryP(166,&value)
	log("p20 r:%d", value)
	countreads(value,20,11)
	l.Unlock()

	delay()
	l.Lock()
	value = 22
	s22.Put(352,value)
	log("p20 w:%d", value)
	countwrites(value,20,22)
	l.Unlock()

	delay()
	l.Lock()
	value = 9
	s9.Put(129,value)
	log("p20 w:%d", value)
	countwrites(value,20,9)
	l.Unlock()

	delay()
	l.Lock()
	value = 28
	s28.Put(438,value)
	log("p20 w:%d", value)
	countwrites(value,20,28)
	l.Unlock()
}


func p21() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 21
	s21.Put(326,value)
	log("p21 w:%d", value)
	countwrites(value,21,21)
	l.Unlock()

	delay()
	l.Lock()
	value = 23
	s23.Put(362,value)
	log("p21 w:%d", value)
	countwrites(value,21,23)
	l.Unlock()

	delay()
	l.Lock()
	value = 9
	s9.Put(142,value)
	log("p21 w:%d", value)
	countwrites(value,21,9)
	l.Unlock()

	delay()
	l.Lock()
	value = 11
	s11.Put(163,value)
	log("p21 w:%d", value)
	countwrites(value,21,11)
	l.Unlock()

	delay()
	l.Lock()
	value = 28
	s28.Put(433,value)
	log("p21 w:%d", value)
	countwrites(value,21,28)
	l.Unlock()

	delay()
	l.Lock()
	value = 1
	s1.Put(2,value)
	log("p21 w:%d", value)
	countwrites(value,21,1)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s28.QueryP(445,&value)
	log("p21 r:%d", value)
	countreads(value,21,28)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s29.QueryP(454,&value)
	log("p21 r:%d", value)
	countreads(value,21,29)
	l.Unlock()

	delay()
	l.Lock()
	value = 3
	s3.Put(33,value)
	log("p21 w:%d", value)
	countwrites(value,21,3)
	l.Unlock()

	delay()
	l.Lock()
	value = 31
	s31.Put(496,value)
	log("p21 w:%d", value)
	countwrites(value,21,31)
	l.Unlock()

	delay()
	l.Lock()
	value = 12
	s12.Put(190,value)
	log("p21 w:%d", value)
	countwrites(value,21,12)
	l.Unlock()

	delay()
	l.Lock()
	value = 22
	s22.Put(350,value)
	log("p21 w:%d", value)
	countwrites(value,21,22)
	l.Unlock()

	delay()
	l.Lock()
	value = 29
	s29.Put(452,value)
	log("p21 w:%d", value)
	countwrites(value,21,29)
	l.Unlock()

	delay()
	l.Lock()
	value = 8
	s8.Put(127,value)
	log("p21 w:%d", value)
	countwrites(value,21,8)
	l.Unlock()

	delay()
	l.Lock()
	value = 9
	s9.Put(142,value)
	log("p21 w:%d", value)
	countwrites(value,21,9)
	l.Unlock()

	delay()
	l.Lock()
	value = 2
	s2.Put(26,value)
	log("p21 w:%d", value)
	countwrites(value,21,2)
	l.Unlock()
}


func p22() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 27
	s27.Put(420,value)
	log("p22 w:%d", value)
	countwrites(value,22,27)
	l.Unlock()

	delay()
	l.Lock()
	value = 14
	s14.Put(223,value)
	log("p22 w:%d", value)
	countwrites(value,22,14)
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
	value = 8
	s8.Put(125,value)
	log("p22 w:%d", value)
	countwrites(value,22,8)
	l.Unlock()

	delay()
	l.Lock()
	value = 1
	s1.Put(11,value)
	log("p22 w:%d", value)
	countwrites(value,22,1)
	l.Unlock()

	delay()
	l.Lock()
	value = 32
	s32.Put(510,value)
	log("p22 w:%d", value)
	countwrites(value,22,32)
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
	value = 28
	s28.Put(446,value)
	log("p22 w:%d", value)
	countwrites(value,22,28)
	l.Unlock()

	delay()
	l.Lock()
	value = 12
	s12.Put(177,value)
	log("p22 w:%d", value)
	countwrites(value,22,12)
	l.Unlock()

	delay()
	l.Lock()
	value = 15
	s15.Put(229,value)
	log("p22 w:%d", value)
	countwrites(value,22,15)
	l.Unlock()

	delay()
	l.Lock()
	value = 25
	s25.Put(397,value)
	log("p22 w:%d", value)
	countwrites(value,22,25)
	l.Unlock()

	delay()
	l.Lock()
	value = 22
	s22.Put(349,value)
	log("p22 w:%d", value)
	countwrites(value,22,22)
	l.Unlock()

	delay()
	l.Lock()
	value = 3
	s3.Put(36,value)
	log("p22 w:%d", value)
	countwrites(value,22,3)
	l.Unlock()

	delay()
	l.Lock()
	value = 2
	s2.Put(28,value)
	log("p22 w:%d", value)
	countwrites(value,22,2)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s5.QueryP(70,&value)
	log("p22 r:%d", value)
	countreads(value,22,5)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s6.QueryP(95,&value)
	log("p22 r:%d", value)
	countreads(value,22,6)
	l.Unlock()
}


func p23() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 7
	s7.Put(107,value)
	log("p23 w:%d", value)
	countwrites(value,23,7)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s21.QueryP(330,&value)
	log("p23 r:%d", value)
	countreads(value,23,21)
	l.Unlock()

	delay()
	l.Lock()
	value = 15
	s15.Put(233,value)
	log("p23 w:%d", value)
	countwrites(value,23,15)
	l.Unlock()

	delay()
	l.Lock()
	value = 4
	s4.Put(53,value)
	log("p23 w:%d", value)
	countwrites(value,23,4)
	l.Unlock()

	delay()
	l.Lock()
	value = 29
	s29.Put(461,value)
	log("p23 w:%d", value)
	countwrites(value,23,29)
	l.Unlock()

	delay()
	l.Lock()
	value = 14
	s14.Put(217,value)
	log("p23 w:%d", value)
	countwrites(value,23,14)
	l.Unlock()

	delay()
	l.Lock()
	value = 12
	s12.Put(188,value)
	log("p23 w:%d", value)
	countwrites(value,23,12)
	l.Unlock()

	delay()
	l.Lock()
	value = 31
	s31.Put(487,value)
	log("p23 w:%d", value)
	countwrites(value,23,31)
	l.Unlock()

	delay()
	l.Lock()
	value = 7
	s7.Put(111,value)
	log("p23 w:%d", value)
	countwrites(value,23,7)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s6.QueryP(96,&value)
	log("p23 r:%d", value)
	countreads(value,23,6)
	l.Unlock()

	delay()
	l.Lock()
	value = 18
	s18.Put(274,value)
	log("p23 w:%d", value)
	countwrites(value,23,18)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s21.QueryP(326,&value)
	log("p23 r:%d", value)
	countreads(value,23,21)
	l.Unlock()

	delay()
	l.Lock()
	value = 9
	s9.Put(138,value)
	log("p23 w:%d", value)
	countwrites(value,23,9)
	l.Unlock()

	delay()
	l.Lock()
	value = 1
	s1.Put(2,value)
	log("p23 w:%d", value)
	countwrites(value,23,1)
	l.Unlock()

	delay()
	l.Lock()
	value = 5
	s5.Put(76,value)
	log("p23 w:%d", value)
	countwrites(value,23,5)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s22.QueryP(341,&value)
	log("p23 r:%d", value)
	countreads(value,23,22)
	l.Unlock()
}


func p24() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 3
	s3.Put(45,value)
	log("p24 w:%d", value)
	countwrites(value,24,3)
	l.Unlock()

	delay()
	l.Lock()
	value = 11
	s11.Put(168,value)
	log("p24 w:%d", value)
	countwrites(value,24,11)
	l.Unlock()

	delay()
	l.Lock()
	value = 11
	s11.Put(176,value)
	log("p24 w:%d", value)
	countwrites(value,24,11)
	l.Unlock()

	delay()
	l.Lock()
	value = 4
	s4.Put(57,value)
	log("p24 w:%d", value)
	countwrites(value,24,4)
	l.Unlock()

	delay()
	l.Lock()
	value = 16
	s16.Put(242,value)
	log("p24 w:%d", value)
	countwrites(value,24,16)
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
	s1.Put(4,value)
	log("p24 w:%d", value)
	countwrites(value,24,1)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s3.QueryP(34,&value)
	log("p24 r:%d", value)
	countreads(value,24,3)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s32.QueryP(512,&value)
	log("p24 r:%d", value)
	countreads(value,24,32)
	l.Unlock()

	delay()
	l.Lock()
	value = 24
	s24.Put(381,value)
	log("p24 w:%d", value)
	countwrites(value,24,24)
	l.Unlock()

	delay()
	l.Lock()
	value = 25
	s25.Put(392,value)
	log("p24 w:%d", value)
	countwrites(value,24,25)
	l.Unlock()

	delay()
	l.Lock()
	value = 3
	s3.Put(41,value)
	log("p24 w:%d", value)
	countwrites(value,24,3)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s14.QueryP(217,&value)
	log("p24 r:%d", value)
	countreads(value,24,14)
	l.Unlock()

	delay()
	l.Lock()
	value = 17
	s17.Put(269,value)
	log("p24 w:%d", value)
	countwrites(value,24,17)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s24.QueryP(384,&value)
	log("p24 r:%d", value)
	countreads(value,24,24)
	l.Unlock()

	delay()
	l.Lock()
	value = 18
	s18.Put(285,value)
	log("p24 w:%d", value)
	countwrites(value,24,18)
	l.Unlock()
}


func p25() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 15
	s15.Put(225,value)
	log("p25 w:%d", value)
	countwrites(value,25,15)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s1.QueryP(13,&value)
	log("p25 r:%d", value)
	countreads(value,25,1)
	l.Unlock()

	delay()
	l.Lock()
	value = 24
	s24.Put(377,value)
	log("p25 w:%d", value)
	countwrites(value,25,24)
	l.Unlock()

	delay()
	l.Lock()
	value = 21
	s21.Put(336,value)
	log("p25 w:%d", value)
	countwrites(value,25,21)
	l.Unlock()

	delay()
	l.Lock()
	value = 21
	s21.Put(333,value)
	log("p25 w:%d", value)
	countwrites(value,25,21)
	l.Unlock()

	delay()
	l.Lock()
	value = 22
	s22.Put(337,value)
	log("p25 w:%d", value)
	countwrites(value,25,22)
	l.Unlock()

	delay()
	l.Lock()
	value = 9
	s9.Put(142,value)
	log("p25 w:%d", value)
	countwrites(value,25,9)
	l.Unlock()

	delay()
	l.Lock()
	value = 14
	s14.Put(224,value)
	log("p25 w:%d", value)
	countwrites(value,25,14)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s2.QueryP(26,&value)
	log("p25 r:%d", value)
	countreads(value,25,2)
	l.Unlock()

	delay()
	l.Lock()
	value = 19
	s19.Put(291,value)
	log("p25 w:%d", value)
	countwrites(value,25,19)
	l.Unlock()

	delay()
	l.Lock()
	value = 15
	s15.Put(228,value)
	log("p25 w:%d", value)
	countwrites(value,25,15)
	l.Unlock()

	delay()
	l.Lock()
	value = 30
	s30.Put(467,value)
	log("p25 w:%d", value)
	countwrites(value,25,30)
	l.Unlock()

	delay()
	l.Lock()
	value = 17
	s17.Put(259,value)
	log("p25 w:%d", value)
	countwrites(value,25,17)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s15.QueryP(233,&value)
	log("p25 r:%d", value)
	countreads(value,25,15)
	l.Unlock()

	delay()
	l.Lock()
	value = 9
	s9.Put(138,value)
	log("p25 w:%d", value)
	countwrites(value,25,9)
	l.Unlock()

	delay()
	l.Lock()
	value = 12
	s12.Put(192,value)
	log("p25 w:%d", value)
	countwrites(value,25,12)
	l.Unlock()
}


func p26() {
	defer wg.Done()
	var value int

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
	s19.QueryP(300,&value)
	log("p26 r:%d", value)
	countreads(value,26,19)
	l.Unlock()

	delay()
	l.Lock()
	value = 21
	s21.Put(323,value)
	log("p26 w:%d", value)
	countwrites(value,26,21)
	l.Unlock()

	delay()
	l.Lock()
	value = 11
	s11.Put(168,value)
	log("p26 w:%d", value)
	countwrites(value,26,11)
	l.Unlock()

	delay()
	l.Lock()
	value = 3
	s3.Put(35,value)
	log("p26 w:%d", value)
	countwrites(value,26,3)
	l.Unlock()

	delay()
	l.Lock()
	value = 21
	s21.Put(321,value)
	log("p26 w:%d", value)
	countwrites(value,26,21)
	l.Unlock()

	delay()
	l.Lock()
	value = 4
	s4.Put(56,value)
	log("p26 w:%d", value)
	countwrites(value,26,4)
	l.Unlock()

	delay()
	l.Lock()
	value = 4
	s4.Put(62,value)
	log("p26 w:%d", value)
	countwrites(value,26,4)
	l.Unlock()

	delay()
	l.Lock()
	value = 1
	s1.Put(3,value)
	log("p26 w:%d", value)
	countwrites(value,26,1)
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
	value = 2
	s2.Put(28,value)
	log("p26 w:%d", value)
	countwrites(value,26,2)
	l.Unlock()

	delay()
	l.Lock()
	value = 5
	s5.Put(79,value)
	log("p26 w:%d", value)
	countwrites(value,26,5)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s32.QueryP(506,&value)
	log("p26 r:%d", value)
	countreads(value,26,32)
	l.Unlock()

	delay()
	l.Lock()
	value = 20
	s20.Put(318,value)
	log("p26 w:%d", value)
	countwrites(value,26,20)
	l.Unlock()

	delay()
	l.Lock()
	value = 9
	s9.Put(144,value)
	log("p26 w:%d", value)
	countwrites(value,26,9)
	l.Unlock()

	delay()
	l.Lock()
	value = 19
	s19.Put(290,value)
	log("p26 w:%d", value)
	countwrites(value,26,19)
	l.Unlock()
}


func p27() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 11
	s11.Put(164,value)
	log("p27 w:%d", value)
	countwrites(value,27,11)
	l.Unlock()

	delay()
	l.Lock()
	value = 10
	s10.Put(156,value)
	log("p27 w:%d", value)
	countwrites(value,27,10)
	l.Unlock()

	delay()
	l.Lock()
	value = 15
	s15.Put(232,value)
	log("p27 w:%d", value)
	countwrites(value,27,15)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s16.QueryP(247,&value)
	log("p27 r:%d", value)
	countreads(value,27,16)
	l.Unlock()

	delay()
	l.Lock()
	value = 26
	s26.Put(405,value)
	log("p27 w:%d", value)
	countwrites(value,27,26)
	l.Unlock()

	delay()
	l.Lock()
	value = 12
	s12.Put(189,value)
	log("p27 w:%d", value)
	countwrites(value,27,12)
	l.Unlock()

	delay()
	l.Lock()
	value = 15
	s15.Put(226,value)
	log("p27 w:%d", value)
	countwrites(value,27,15)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s23.QueryP(361,&value)
	log("p27 r:%d", value)
	countreads(value,27,23)
	l.Unlock()

	delay()
	l.Lock()
	value = 6
	s6.Put(83,value)
	log("p27 w:%d", value)
	countwrites(value,27,6)
	l.Unlock()

	delay()
	l.Lock()
	value = 9
	s9.Put(142,value)
	log("p27 w:%d", value)
	countwrites(value,27,9)
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
	value = 6
	s6.Put(96,value)
	log("p27 w:%d", value)
	countwrites(value,27,6)
	l.Unlock()

	delay()
	l.Lock()
	value = 26
	s26.Put(404,value)
	log("p27 w:%d", value)
	countwrites(value,27,26)
	l.Unlock()

	delay()
	l.Lock()
	value = 26
	s26.Put(412,value)
	log("p27 w:%d", value)
	countwrites(value,27,26)
	l.Unlock()

	delay()
	l.Lock()
	value = 14
	s14.Put(209,value)
	log("p27 w:%d", value)
	countwrites(value,27,14)
	l.Unlock()

	delay()
	l.Lock()
	value = 6
	s6.Put(95,value)
	log("p27 w:%d", value)
	countwrites(value,27,6)
	l.Unlock()
}


func p28() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	s23.QueryP(361,&value)
	log("p28 r:%d", value)
	countreads(value,28,23)
	l.Unlock()

	delay()
	l.Lock()
	value = 18
	s18.Put(279,value)
	log("p28 w:%d", value)
	countwrites(value,28,18)
	l.Unlock()

	delay()
	l.Lock()
	value = 32
	s32.Put(512,value)
	log("p28 w:%d", value)
	countwrites(value,28,32)
	l.Unlock()

	delay()
	l.Lock()
	value = 25
	s25.Put(399,value)
	log("p28 w:%d", value)
	countwrites(value,28,25)
	l.Unlock()

	delay()
	l.Lock()
	value = 19
	s19.Put(290,value)
	log("p28 w:%d", value)
	countwrites(value,28,19)
	l.Unlock()

	delay()
	l.Lock()
	value = 19
	s19.Put(290,value)
	log("p28 w:%d", value)
	countwrites(value,28,19)
	l.Unlock()

	delay()
	l.Lock()
	value = 25
	s25.Put(399,value)
	log("p28 w:%d", value)
	countwrites(value,28,25)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s27.QueryP(430,&value)
	log("p28 r:%d", value)
	countreads(value,28,27)
	l.Unlock()

	delay()
	l.Lock()
	value = 14
	s14.Put(222,value)
	log("p28 w:%d", value)
	countwrites(value,28,14)
	l.Unlock()

	delay()
	l.Lock()
	value = 6
	s6.Put(82,value)
	log("p28 w:%d", value)
	countwrites(value,28,6)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s23.QueryP(355,&value)
	log("p28 r:%d", value)
	countreads(value,28,23)
	l.Unlock()

	delay()
	l.Lock()
	value = 25
	s25.Put(394,value)
	log("p28 w:%d", value)
	countwrites(value,28,25)
	l.Unlock()

	delay()
	l.Lock()
	value = 20
	s20.Put(309,value)
	log("p28 w:%d", value)
	countwrites(value,28,20)
	l.Unlock()

	delay()
	l.Lock()
	value = 5
	s5.Put(70,value)
	log("p28 w:%d", value)
	countwrites(value,28,5)
	l.Unlock()

	delay()
	l.Lock()
	value = 28
	s28.Put(442,value)
	log("p28 w:%d", value)
	countwrites(value,28,28)
	l.Unlock()

	delay()
	l.Lock()
	value = 20
	s20.Put(306,value)
	log("p28 w:%d", value)
	countwrites(value,28,20)
	l.Unlock()
}


func p29() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 24
	s24.Put(384,value)
	log("p29 w:%d", value)
	countwrites(value,29,24)
	l.Unlock()

	delay()
	l.Lock()
	value = 19
	s19.Put(298,value)
	log("p29 w:%d", value)
	countwrites(value,29,19)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s20.QueryP(312,&value)
	log("p29 r:%d", value)
	countreads(value,29,20)
	l.Unlock()

	delay()
	l.Lock()
	value = 10
	s10.Put(146,value)
	log("p29 w:%d", value)
	countwrites(value,29,10)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s22.QueryP(338,&value)
	log("p29 r:%d", value)
	countreads(value,29,22)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s16.QueryP(252,&value)
	log("p29 r:%d", value)
	countreads(value,29,16)
	l.Unlock()

	delay()
	l.Lock()
	value = 27
	s27.Put(422,value)
	log("p29 w:%d", value)
	countwrites(value,29,27)
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
	value = -1
	s8.QueryP(114,&value)
	log("p29 r:%d", value)
	countreads(value,29,8)
	l.Unlock()

	delay()
	l.Lock()
	value = 9
	s9.Put(129,value)
	log("p29 w:%d", value)
	countwrites(value,29,9)
	l.Unlock()

	delay()
	l.Lock()
	value = 29
	s29.Put(457,value)
	log("p29 w:%d", value)
	countwrites(value,29,29)
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
	value = 22
	s22.Put(340,value)
	log("p29 w:%d", value)
	countwrites(value,29,22)
	l.Unlock()

	delay()
	l.Lock()
	value = 32
	s32.Put(502,value)
	log("p29 w:%d", value)
	countwrites(value,29,32)
	l.Unlock()

	delay()
	l.Lock()
	value = 18
	s18.Put(280,value)
	log("p29 w:%d", value)
	countwrites(value,29,18)
	l.Unlock()

	delay()
	l.Lock()
	value = 32
	s32.Put(501,value)
	log("p29 w:%d", value)
	countwrites(value,29,32)
	l.Unlock()
}


func p30() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 20
	s20.Put(318,value)
	log("p30 w:%d", value)
	countwrites(value,30,20)
	l.Unlock()

	delay()
	l.Lock()
	value = 3
	s3.Put(33,value)
	log("p30 w:%d", value)
	countwrites(value,30,3)
	l.Unlock()

	delay()
	l.Lock()
	value = 2
	s2.Put(29,value)
	log("p30 w:%d", value)
	countwrites(value,30,2)
	l.Unlock()

	delay()
	l.Lock()
	value = 18
	s18.Put(286,value)
	log("p30 w:%d", value)
	countwrites(value,30,18)
	l.Unlock()

	delay()
	l.Lock()
	value = 29
	s29.Put(463,value)
	log("p30 w:%d", value)
	countwrites(value,30,29)
	l.Unlock()

	delay()
	l.Lock()
	value = 24
	s24.Put(384,value)
	log("p30 w:%d", value)
	countwrites(value,30,24)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s2.QueryP(26,&value)
	log("p30 r:%d", value)
	countreads(value,30,2)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s6.QueryP(93,&value)
	log("p30 r:%d", value)
	countreads(value,30,6)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s8.QueryP(121,&value)
	log("p30 r:%d", value)
	countreads(value,30,8)
	l.Unlock()

	delay()
	l.Lock()
	value = 28
	s28.Put(442,value)
	log("p30 w:%d", value)
	countwrites(value,30,28)
	l.Unlock()

	delay()
	l.Lock()
	value = 5
	s5.Put(72,value)
	log("p30 w:%d", value)
	countwrites(value,30,5)
	l.Unlock()

	delay()
	l.Lock()
	value = 12
	s12.Put(185,value)
	log("p30 w:%d", value)
	countwrites(value,30,12)
	l.Unlock()

	delay()
	l.Lock()
	value = 18
	s18.Put(280,value)
	log("p30 w:%d", value)
	countwrites(value,30,18)
	l.Unlock()

	delay()
	l.Lock()
	value = 24
	s24.Put(382,value)
	log("p30 w:%d", value)
	countwrites(value,30,24)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s24.QueryP(371,&value)
	log("p30 r:%d", value)
	countreads(value,30,24)
	l.Unlock()

	delay()
	l.Lock()
	value = 11
	s11.Put(173,value)
	log("p30 w:%d", value)
	countwrites(value,30,11)
	l.Unlock()
}


func p31() {
	defer wg.Done()
	var value int

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
	s22.QueryP(339,&value)
	log("p31 r:%d", value)
	countreads(value,31,22)
	l.Unlock()

	delay()
	l.Lock()
	value = 29
	s29.Put(459,value)
	log("p31 w:%d", value)
	countwrites(value,31,29)
	l.Unlock()

	delay()
	l.Lock()
	value = 5
	s5.Put(67,value)
	log("p31 w:%d", value)
	countwrites(value,31,5)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s12.QueryP(190,&value)
	log("p31 r:%d", value)
	countreads(value,31,12)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s6.QueryP(95,&value)
	log("p31 r:%d", value)
	countreads(value,31,6)
	l.Unlock()

	delay()
	l.Lock()
	value = 11
	s11.Put(162,value)
	log("p31 w:%d", value)
	countwrites(value,31,11)
	l.Unlock()

	delay()
	l.Lock()
	value = 17
	s17.Put(271,value)
	log("p31 w:%d", value)
	countwrites(value,31,17)
	l.Unlock()

	delay()
	l.Lock()
	value = 2
	s2.Put(30,value)
	log("p31 w:%d", value)
	countwrites(value,31,2)
	l.Unlock()

	delay()
	l.Lock()
	value = 5
	s5.Put(68,value)
	log("p31 w:%d", value)
	countwrites(value,31,5)
	l.Unlock()

	delay()
	l.Lock()
	value = 27
	s27.Put(421,value)
	log("p31 w:%d", value)
	countwrites(value,31,27)
	l.Unlock()

	delay()
	l.Lock()
	value = 29
	s29.Put(461,value)
	log("p31 w:%d", value)
	countwrites(value,31,29)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s29.QueryP(456,&value)
	log("p31 r:%d", value)
	countreads(value,31,29)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s4.QueryP(56,&value)
	log("p31 r:%d", value)
	countreads(value,31,4)
	l.Unlock()

	delay()
	l.Lock()
	value = 27
	s27.Put(424,value)
	log("p31 w:%d", value)
	countwrites(value,31,27)
	l.Unlock()

	delay()
	l.Lock()
	value = 6
	s6.Put(83,value)
	log("p31 w:%d", value)
	countwrites(value,31,6)
	l.Unlock()
}


func p32() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 12
	s12.Put(188,value)
	log("p32 w:%d", value)
	countwrites(value,32,12)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s25.QueryP(386,&value)
	log("p32 r:%d", value)
	countreads(value,32,25)
	l.Unlock()

	delay()
	l.Lock()
	value = 8
	s8.Put(123,value)
	log("p32 w:%d", value)
	countwrites(value,32,8)
	l.Unlock()

	delay()
	l.Lock()
	value = 23
	s23.Put(364,value)
	log("p32 w:%d", value)
	countwrites(value,32,23)
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
	value = 5
	s5.Put(70,value)
	log("p32 w:%d", value)
	countwrites(value,32,5)
	l.Unlock()

	delay()
	l.Lock()
	value = 3
	s3.Put(48,value)
	log("p32 w:%d", value)
	countwrites(value,32,3)
	l.Unlock()

	delay()
	l.Lock()
	value = 8
	s8.Put(123,value)
	log("p32 w:%d", value)
	countwrites(value,32,8)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s29.QueryP(453,&value)
	log("p32 r:%d", value)
	countreads(value,32,29)
	l.Unlock()

	delay()
	l.Lock()
	value = 15
	s15.Put(233,value)
	log("p32 w:%d", value)
	countwrites(value,32,15)
	l.Unlock()

	delay()
	l.Lock()
	value = 18
	s18.Put(286,value)
	log("p32 w:%d", value)
	countwrites(value,32,18)
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
	value = 19
	s19.Put(295,value)
	log("p32 w:%d", value)
	countwrites(value,32,19)
	l.Unlock()

	delay()
	l.Lock()
	value = 9
	s9.Put(137,value)
	log("p32 w:%d", value)
	countwrites(value,32,9)
	l.Unlock()

	delay()
	l.Lock()
	value = 8
	s8.Put(114,value)
	log("p32 w:%d", value)
	countwrites(value,32,8)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s29.QueryP(464,&value)
	log("p32 r:%d", value)
	countreads(value,32,29)
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
