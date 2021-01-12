package main

import (
	"os"
	"fmt"
	"sync"
	"time"
	"math/rand"
	. "github.com/pspaces/gospace"
	. "github.com/repligospaces"
)

var uri = make(map[Space]string)
var Sp = make(map[string]*Space)
var rsp Replispace = Replispace{Sp: Sp}

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
	writes_replicated = Getwcount()-writes_local
	reads_success = 0
	reads_insuccess = 0
	rand.Seed(time.Now().UTC().UnixNano())

	wg.Add(32)

	s1 = NewSpace("tcp://localhost:34001/s1")
	Sp["tcp://localhost:34001/s1"] = &s1
	uri[s1] = "tcp://localhost:34001/s1"
	s2 = NewSpace("tcp://localhost:34002/s2")
	Sp["tcp://localhost:34002/s2"] = &s2
	uri[s2] = "tcp://localhost:34002/s2"
	s3 = NewSpace("tcp://localhost:34003/s3")
	Sp["tcp://localhost:34003/s3"] = &s3
	uri[s3] = "tcp://localhost:34003/s3"
	s4 = NewSpace("tcp://localhost:34004/s4")
	Sp["tcp://localhost:34004/s4"] = &s4
	uri[s4] = "tcp://localhost:34004/s4"
	s5 = NewSpace("tcp://localhost:34005/s5")
	Sp["tcp://localhost:34005/s5"] = &s5
	uri[s5] = "tcp://localhost:34005/s5"
	s6 = NewSpace("tcp://localhost:34006/s6")
	Sp["tcp://localhost:34006/s6"] = &s6
	uri[s6] = "tcp://localhost:34006/s6"
	s7 = NewSpace("tcp://localhost:34007/s7")
	Sp["tcp://localhost:34007/s7"] = &s7
	uri[s7] = "tcp://localhost:34007/s7"
	s8 = NewSpace("tcp://localhost:34008/s8")
	Sp["tcp://localhost:34008/s8"] = &s8
	uri[s8] = "tcp://localhost:34008/s8"
	s9 = NewSpace("tcp://localhost:34009/s9")
	Sp["tcp://localhost:34009/s9"] = &s9
	uri[s9] = "tcp://localhost:34009/s9"
	s10 = NewSpace("tcp://localhost:34010/s10")
	Sp["tcp://localhost:34010/s10"] = &s10
	uri[s10] = "tcp://localhost:34010/s10"
	s11 = NewSpace("tcp://localhost:34011/s11")
	Sp["tcp://localhost:34011/s11"] = &s11
	uri[s11] = "tcp://localhost:34011/s11"
	s12 = NewSpace("tcp://localhost:34012/s12")
	Sp["tcp://localhost:34012/s12"] = &s12
	uri[s12] = "tcp://localhost:34012/s12"
	s13 = NewSpace("tcp://localhost:34013/s13")
	Sp["tcp://localhost:34013/s13"] = &s13
	uri[s13] = "tcp://localhost:34013/s13"
	s14 = NewSpace("tcp://localhost:34014/s14")
	Sp["tcp://localhost:34014/s14"] = &s14
	uri[s14] = "tcp://localhost:34014/s14"
	s15 = NewSpace("tcp://localhost:34015/s15")
	Sp["tcp://localhost:34015/s15"] = &s15
	uri[s15] = "tcp://localhost:34015/s15"
	s16 = NewSpace("tcp://localhost:34016/s16")
	Sp["tcp://localhost:34016/s16"] = &s16
	uri[s16] = "tcp://localhost:34016/s16"
	s17 = NewSpace("tcp://localhost:34017/s17")
	Sp["tcp://localhost:34017/s17"] = &s17
	uri[s17] = "tcp://localhost:34017/s17"
	s18 = NewSpace("tcp://localhost:34018/s18")
	Sp["tcp://localhost:34018/s18"] = &s18
	uri[s18] = "tcp://localhost:34018/s18"
	s19 = NewSpace("tcp://localhost:34019/s19")
	Sp["tcp://localhost:34019/s19"] = &s19
	uri[s19] = "tcp://localhost:34019/s19"
	s20 = NewSpace("tcp://localhost:34020/s20")
	Sp["tcp://localhost:34020/s20"] = &s20
	uri[s20] = "tcp://localhost:34020/s20"
	s21 = NewSpace("tcp://localhost:34021/s21")
	Sp["tcp://localhost:34021/s21"] = &s21
	uri[s21] = "tcp://localhost:34021/s21"
	s22 = NewSpace("tcp://localhost:34022/s22")
	Sp["tcp://localhost:34022/s22"] = &s22
	uri[s22] = "tcp://localhost:34022/s22"
	s23 = NewSpace("tcp://localhost:34023/s23")
	Sp["tcp://localhost:34023/s23"] = &s23
	uri[s23] = "tcp://localhost:34023/s23"
	s24 = NewSpace("tcp://localhost:34024/s24")
	Sp["tcp://localhost:34024/s24"] = &s24
	uri[s24] = "tcp://localhost:34024/s24"
	s25 = NewSpace("tcp://localhost:34025/s25")
	Sp["tcp://localhost:34025/s25"] = &s25
	uri[s25] = "tcp://localhost:34025/s25"
	s26 = NewSpace("tcp://localhost:34026/s26")
	Sp["tcp://localhost:34026/s26"] = &s26
	uri[s26] = "tcp://localhost:34026/s26"
	s27 = NewSpace("tcp://localhost:34027/s27")
	Sp["tcp://localhost:34027/s27"] = &s27
	uri[s27] = "tcp://localhost:34027/s27"
	s28 = NewSpace("tcp://localhost:34028/s28")
	Sp["tcp://localhost:34028/s28"] = &s28
	uri[s28] = "tcp://localhost:34028/s28"
	s29 = NewSpace("tcp://localhost:34029/s29")
	Sp["tcp://localhost:34029/s29"] = &s29
	uri[s29] = "tcp://localhost:34029/s29"
	s30 = NewSpace("tcp://localhost:34030/s30")
	Sp["tcp://localhost:34030/s30"] = &s30
	uri[s30] = "tcp://localhost:34030/s30"
	s31 = NewSpace("tcp://localhost:34031/s31")
	Sp["tcp://localhost:34031/s31"] = &s31
	uri[s31] = "tcp://localhost:34031/s31"
	s32 = NewSpace("tcp://localhost:34032/s32")
	Sp["tcp://localhost:34032/s32"] = &s32
	uri[s32] = "tcp://localhost:34032/s32"

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

	writes_replicated = Getwcount()-writes_local	// these include the normal put too

	fmt.Println("       loc w,    rem w,   repl w,    loc r,    rem r,    tot w,   succ r,   fail r")
	fmt.Printf("    %8d, %8d, %8d, %8d, %8d, %8d, %8d, %8d\n", writes_local, writes_remote, writes_replicated, reads_local, reads_remote, writestotal, reads_success, reads_insuccess)
}

func p1() {
	targets3 := make([]string, 2)
	targets3[1] = uri[s24]
	targets3[0] = uri[s10]
	targets2 := make([]string, 0)
	targets1 := make([]string, 0)
	targets0 := make([]string, 0)

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(461, &value), rsp, s1)
	log("p1 r:%d", value)
	countreads(value, 1, 29)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(433, &value), rsp, s1)
	log("p1 r:%d", value)
	countreads(value, 1, 28)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(459, &value), rsp, s1)
	log("p1 r:%d", value)
	countreads(value, 1, 29)
	l.Unlock()

	delay()
	l.Lock()
	value = 9
	Put(CreateTuple(135, value), rsp, targets0)
	log("p1 w:%d", value)
	countwrites(value, 1, 9)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(28, &value), rsp, s1)
	log("p1 r:%d", value)
	countreads(value, 1, 2)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(356, &value), rsp, s1)
	log("p1 r:%d", value)
	countreads(value, 1, 23)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(427, &value), rsp, s1)
	log("p1 r:%d", value)
	countreads(value, 1, 27)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(348, &value), rsp, s1)
	log("p1 r:%d", value)
	countreads(value, 1, 22)
	l.Unlock()

	delay()
	l.Lock()
	value = 15
	Put(CreateTuple(236, value), rsp, targets1)
	log("p1 w:%d", value)
	countwrites(value, 1, 15)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(86, &value), rsp, s1)
	log("p1 r:%d", value)
	countreads(value, 1, 6)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(467, &value), rsp, s1)
	log("p1 r:%d", value)
	countreads(value, 1, 30)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(39, &value), rsp, s1)
	log("p1 r:%d", value)
	countreads(value, 1, 3)
	l.Unlock()

	delay()
	l.Lock()
	value = 29
	Put(CreateTuple(460, value), rsp, targets2)
	log("p1 w:%d", value)
	countwrites(value, 1, 29)
	l.Unlock()

	delay()
	l.Lock()
	value = 7
	Put(CreateTuple(97, value), rsp, targets3)
	log("p1 w:%d", value)
	countwrites(value, 1, 7)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(507, &value), rsp, s1)
	log("p1 r:%d", value)
	countreads(value, 1, 32)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(471, &value), rsp, s1)
	log("p1 r:%d", value)
	countreads(value, 1, 30)
	l.Unlock()
}

func p2() {
	targets5 := make([]string, 0)
	targets4 := make([]string, 2)
	targets4[1] = uri[s20]
	targets4[0] = uri[s5]
	targets3 := make([]string, 3)
	targets3[2] = uri[s32]
	targets3[1] = uri[s8]
	targets3[0] = uri[s1]
	targets2 := make([]string, 2)
	targets2[1] = uri[s18]
	targets2[0] = uri[s17]
	targets1 := make([]string, 0)
	targets0 := make([]string, 0)

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(414, &value), rsp, s2)
	log("p2 r:%d", value)
	countreads(value, 2, 26)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(184, &value), rsp, s2)
	log("p2 r:%d", value)
	countreads(value, 2, 12)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(148, &value), rsp, s2)
	log("p2 r:%d", value)
	countreads(value, 2, 10)
	l.Unlock()

	delay()
	l.Lock()
	value = 15
	Put(CreateTuple(230, value), rsp, targets0)
	log("p2 w:%d", value)
	countwrites(value, 2, 15)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(62, &value), rsp, s2)
	log("p2 r:%d", value)
	countreads(value, 2, 4)
	l.Unlock()

	delay()
	l.Lock()
	value = 31
	Put(CreateTuple(496, value), rsp, targets1)
	log("p2 w:%d", value)
	countwrites(value, 2, 31)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(375, &value), rsp, s2)
	log("p2 r:%d", value)
	countreads(value, 2, 24)
	l.Unlock()

	delay()
	l.Lock()
	value = 13
	Put(CreateTuple(204, value), rsp, targets2)
	log("p2 w:%d", value)
	countwrites(value, 2, 13)
	l.Unlock()

	delay()
	l.Lock()
	value = 22
	Put(CreateTuple(348, value), rsp, targets3)
	log("p2 w:%d", value)
	countwrites(value, 2, 22)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(24, &value), rsp, s2)
	log("p2 r:%d", value)
	countreads(value, 2, 2)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(431, &value), rsp, s2)
	log("p2 r:%d", value)
	countreads(value, 2, 27)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(58, &value), rsp, s2)
	log("p2 r:%d", value)
	countreads(value, 2, 4)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(41, &value), rsp, s2)
	log("p2 r:%d", value)
	countreads(value, 2, 3)
	l.Unlock()

	delay()
	l.Lock()
	value = 29
	Put(CreateTuple(458, value), rsp, targets4)
	log("p2 w:%d", value)
	countwrites(value, 2, 29)
	l.Unlock()

	delay()
	l.Lock()
	value = 24
	Put(CreateTuple(382, value), rsp, targets5)
	log("p2 w:%d", value)
	countwrites(value, 2, 24)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(153, &value), rsp, s2)
	log("p2 r:%d", value)
	countreads(value, 2, 10)
	l.Unlock()
}

func p3() {
	targets7 := make([]string, 0)
	targets6 := make([]string, 0)
	targets5 := make([]string, 2)
	targets5[1] = uri[s25]
	targets5[0] = uri[s16]
	targets4 := make([]string, 1)
	targets4[0] = uri[s16]
	targets3 := make([]string, 0)
	targets2 := make([]string, 0)
	targets1 := make([]string, 1)
	targets1[0] = uri[s19]
	targets0 := make([]string, 1)
	targets0[0] = uri[s6]

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(124, &value), rsp, s3)
	log("p3 r:%d", value)
	countreads(value, 3, 8)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(179, &value), rsp, s3)
	log("p3 r:%d", value)
	countreads(value, 3, 12)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(363, &value), rsp, s3)
	log("p3 r:%d", value)
	countreads(value, 3, 23)
	l.Unlock()

	delay()
	l.Lock()
	value = 13
	Put(CreateTuple(199, value), rsp, targets0)
	log("p3 w:%d", value)
	countwrites(value, 3, 13)
	l.Unlock()

	delay()
	l.Lock()
	value = 19
	Put(CreateTuple(304, value), rsp, targets1)
	log("p3 w:%d", value)
	countwrites(value, 3, 19)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(446, &value), rsp, s3)
	log("p3 r:%d", value)
	countreads(value, 3, 28)
	l.Unlock()

	delay()
	l.Lock()
	value = 2
	Put(CreateTuple(18, value), rsp, targets2)
	log("p3 w:%d", value)
	countwrites(value, 3, 2)
	l.Unlock()

	delay()
	l.Lock()
	value = 31
	Put(CreateTuple(496, value), rsp, targets3)
	log("p3 w:%d", value)
	countwrites(value, 3, 31)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(297, &value), rsp, s3)
	log("p3 r:%d", value)
	countreads(value, 3, 19)
	l.Unlock()

	delay()
	l.Lock()
	value = 23
	Put(CreateTuple(359, value), rsp, targets4)
	log("p3 w:%d", value)
	countwrites(value, 3, 23)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(312, &value), rsp, s3)
	log("p3 r:%d", value)
	countreads(value, 3, 20)
	l.Unlock()

	delay()
	l.Lock()
	value = 15
	Put(CreateTuple(228, value), rsp, targets5)
	log("p3 w:%d", value)
	countwrites(value, 3, 15)
	l.Unlock()

	delay()
	l.Lock()
	value = 4
	Put(CreateTuple(52, value), rsp, targets6)
	log("p3 w:%d", value)
	countwrites(value, 3, 4)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(461, &value), rsp, s3)
	log("p3 r:%d", value)
	countreads(value, 3, 29)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(473, &value), rsp, s3)
	log("p3 r:%d", value)
	countreads(value, 3, 30)
	l.Unlock()

	delay()
	l.Lock()
	value = 21
	Put(CreateTuple(336, value), rsp, targets7)
	log("p3 w:%d", value)
	countwrites(value, 3, 21)
	l.Unlock()
}

func p4() {
	targets5 := make([]string, 0)
	targets4 := make([]string, 0)
	targets3 := make([]string, 1)
	targets3[0] = uri[s25]
	targets2 := make([]string, 2)
	targets2[1] = uri[s20]
	targets2[0] = uri[s15]
	targets1 := make([]string, 1)
	targets1[0] = uri[s17]
	targets0 := make([]string, 0)

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(151, &value), rsp, s4)
	log("p4 r:%d", value)
	countreads(value, 4, 10)
	l.Unlock()

	delay()
	l.Lock()
	value = 22
	Put(CreateTuple(342, value), rsp, targets0)
	log("p4 w:%d", value)
	countwrites(value, 4, 22)
	l.Unlock()

	delay()
	l.Lock()
	value = 27
	Put(CreateTuple(428, value), rsp, targets1)
	log("p4 w:%d", value)
	countwrites(value, 4, 27)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(448, &value), rsp, s4)
	log("p4 r:%d", value)
	countreads(value, 4, 28)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(390, &value), rsp, s4)
	log("p4 r:%d", value)
	countreads(value, 4, 25)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(285, &value), rsp, s4)
	log("p4 r:%d", value)
	countreads(value, 4, 18)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(323, &value), rsp, s4)
	log("p4 r:%d", value)
	countreads(value, 4, 21)
	l.Unlock()

	delay()
	l.Lock()
	value = 27
	Put(CreateTuple(432, value), rsp, targets2)
	log("p4 w:%d", value)
	countwrites(value, 4, 27)
	l.Unlock()

	delay()
	l.Lock()
	value = 13
	Put(CreateTuple(198, value), rsp, targets3)
	log("p4 w:%d", value)
	countwrites(value, 4, 13)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(59, &value), rsp, s4)
	log("p4 r:%d", value)
	countreads(value, 4, 4)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(154, &value), rsp, s4)
	log("p4 r:%d", value)
	countreads(value, 4, 10)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(1, &value), rsp, s4)
	log("p4 r:%d", value)
	countreads(value, 4, 1)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(464, &value), rsp, s4)
	log("p4 r:%d", value)
	countreads(value, 4, 29)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(212, &value), rsp, s4)
	log("p4 r:%d", value)
	countreads(value, 4, 14)
	l.Unlock()

	delay()
	l.Lock()
	value = 18
	Put(CreateTuple(275, value), rsp, targets4)
	log("p4 w:%d", value)
	countwrites(value, 4, 18)
	l.Unlock()

	delay()
	l.Lock()
	value = 7
	Put(CreateTuple(108, value), rsp, targets5)
	log("p4 w:%d", value)
	countwrites(value, 4, 7)
	l.Unlock()
}

func p5() {
	targets4 := make([]string, 1)
	targets4[0] = uri[s19]
	targets3 := make([]string, 0)
	targets2 := make([]string, 2)
	targets2[1] = uri[s30]
	targets2[0] = uri[s18]
	targets1 := make([]string, 0)
	targets0 := make([]string, 1)
	targets0[0] = uri[s14]

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(81, &value), rsp, s5)
	log("p5 r:%d", value)
	countreads(value, 5, 6)
	l.Unlock()

	delay()
	l.Lock()
	value = 25
	Put(CreateTuple(386, value), rsp, targets0)
	log("p5 w:%d", value)
	countwrites(value, 5, 25)
	l.Unlock()

	delay()
	l.Lock()
	value = 18
	Put(CreateTuple(275, value), rsp, targets1)
	log("p5 w:%d", value)
	countwrites(value, 5, 18)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(458, &value), rsp, s5)
	log("p5 r:%d", value)
	countreads(value, 5, 29)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(64, &value), rsp, s5)
	log("p5 r:%d", value)
	countreads(value, 5, 4)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(486, &value), rsp, s5)
	log("p5 r:%d", value)
	countreads(value, 5, 31)
	l.Unlock()

	delay()
	l.Lock()
	value = 9
	Put(CreateTuple(129, value), rsp, targets2)
	log("p5 w:%d", value)
	countwrites(value, 5, 9)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(492, &value), rsp, s5)
	log("p5 r:%d", value)
	countreads(value, 5, 31)
	l.Unlock()

	delay()
	l.Lock()
	value = 19
	Put(CreateTuple(301, value), rsp, targets3)
	log("p5 w:%d", value)
	countwrites(value, 5, 19)
	l.Unlock()

	delay()
	l.Lock()
	value = 16
	Put(CreateTuple(245, value), rsp, targets4)
	log("p5 w:%d", value)
	countwrites(value, 5, 16)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(509, &value), rsp, s5)
	log("p5 r:%d", value)
	countreads(value, 5, 32)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(315, &value), rsp, s5)
	log("p5 r:%d", value)
	countreads(value, 5, 20)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(146, &value), rsp, s5)
	log("p5 r:%d", value)
	countreads(value, 5, 10)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(251, &value), rsp, s5)
	log("p5 r:%d", value)
	countreads(value, 5, 16)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(235, &value), rsp, s5)
	log("p5 r:%d", value)
	countreads(value, 5, 15)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(16, &value), rsp, s5)
	log("p5 r:%d", value)
	countreads(value, 5, 1)
	l.Unlock()
}

func p6() {
	targets2 := make([]string, 0)
	targets1 := make([]string, 0)
	targets0 := make([]string, 0)
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 5
	Put(CreateTuple(69, value), rsp, targets0)
	log("p6 w:%d", value)
	countwrites(value, 6, 5)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(175, &value), rsp, s6)
	log("p6 r:%d", value)
	countreads(value, 6, 11)
	l.Unlock()

	delay()
	l.Lock()
	value = 11
	Put(CreateTuple(170, value), rsp, targets1)
	log("p6 w:%d", value)
	countwrites(value, 6, 11)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(453, &value), rsp, s6)
	log("p6 r:%d", value)
	countreads(value, 6, 29)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(234, &value), rsp, s6)
	log("p6 r:%d", value)
	countreads(value, 6, 15)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(92, &value), rsp, s6)
	log("p6 r:%d", value)
	countreads(value, 6, 6)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(478, &value), rsp, s6)
	log("p6 r:%d", value)
	countreads(value, 6, 30)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(199, &value), rsp, s6)
	log("p6 r:%d", value)
	countreads(value, 6, 13)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(28, &value), rsp, s6)
	log("p6 r:%d", value)
	countreads(value, 6, 2)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(218, &value), rsp, s6)
	log("p6 r:%d", value)
	countreads(value, 6, 14)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(10, &value), rsp, s6)
	log("p6 r:%d", value)
	countreads(value, 6, 1)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(498, &value), rsp, s6)
	log("p6 r:%d", value)
	countreads(value, 6, 32)
	l.Unlock()

	delay()
	l.Lock()
	value = 8
	Put(CreateTuple(123, value), rsp, targets2)
	log("p6 w:%d", value)
	countwrites(value, 6, 8)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(280, &value), rsp, s6)
	log("p6 r:%d", value)
	countreads(value, 6, 18)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(420, &value), rsp, s6)
	log("p6 r:%d", value)
	countreads(value, 6, 27)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(363, &value), rsp, s6)
	log("p6 r:%d", value)
	countreads(value, 6, 23)
	l.Unlock()
}

func p7() {
	targets3 := make([]string, 1)
	targets3[0] = uri[s32]
	targets2 := make([]string, 0)
	targets1 := make([]string, 0)
	targets0 := make([]string, 0)

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 30
	Put(CreateTuple(465, value), rsp, targets0)
	log("p7 w:%d", value)
	countwrites(value, 7, 30)
	l.Unlock()

	delay()
	l.Lock()
	value = 25
	Put(CreateTuple(396, value), rsp, targets1)
	log("p7 w:%d", value)
	countwrites(value, 7, 25)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(77, &value), rsp, s7)
	log("p7 r:%d", value)
	countreads(value, 7, 5)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(110, &value), rsp, s7)
	log("p7 r:%d", value)
	countreads(value, 7, 7)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(358, &value), rsp, s7)
	log("p7 r:%d", value)
	countreads(value, 7, 23)
	l.Unlock()

	delay()
	l.Lock()
	value = 31
	Put(CreateTuple(496, value), rsp, targets2)
	log("p7 w:%d", value)
	countwrites(value, 7, 31)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(417, &value), rsp, s7)
	log("p7 r:%d", value)
	countreads(value, 7, 27)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(379, &value), rsp, s7)
	log("p7 r:%d", value)
	countreads(value, 7, 24)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(39, &value), rsp, s7)
	log("p7 r:%d", value)
	countreads(value, 7, 3)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(180, &value), rsp, s7)
	log("p7 r:%d", value)
	countreads(value, 7, 12)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(188, &value), rsp, s7)
	log("p7 r:%d", value)
	countreads(value, 7, 12)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(195, &value), rsp, s7)
	log("p7 r:%d", value)
	countreads(value, 7, 13)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(500, &value), rsp, s7)
	log("p7 r:%d", value)
	countreads(value, 7, 32)
	l.Unlock()

	delay()
	l.Lock()
	value = 26
	Put(CreateTuple(404, value), rsp, targets3)
	log("p7 w:%d", value)
	countwrites(value, 7, 26)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(487, &value), rsp, s7)
	log("p7 r:%d", value)
	countreads(value, 7, 31)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(330, &value), rsp, s7)
	log("p7 r:%d", value)
	countreads(value, 7, 21)
	l.Unlock()
}

func p8() {
	targets2 := make([]string, 0)
	targets1 := make([]string, 0)
	targets0 := make([]string, 0)
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(241, &value), rsp, s8)
	log("p8 r:%d", value)
	countreads(value, 8, 16)
	l.Unlock()

	delay()
	l.Lock()
	value = 19
	Put(CreateTuple(299, value), rsp, targets0)
	log("p8 w:%d", value)
	countwrites(value, 8, 19)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(367, &value), rsp, s8)
	log("p8 r:%d", value)
	countreads(value, 8, 23)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(55, &value), rsp, s8)
	log("p8 r:%d", value)
	countreads(value, 8, 4)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(330, &value), rsp, s8)
	log("p8 r:%d", value)
	countreads(value, 8, 21)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(462, &value), rsp, s8)
	log("p8 r:%d", value)
	countreads(value, 8, 29)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(32, &value), rsp, s8)
	log("p8 r:%d", value)
	countreads(value, 8, 2)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(292, &value), rsp, s8)
	log("p8 r:%d", value)
	countreads(value, 8, 19)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(348, &value), rsp, s8)
	log("p8 r:%d", value)
	countreads(value, 8, 22)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(303, &value), rsp, s8)
	log("p8 r:%d", value)
	countreads(value, 8, 19)
	l.Unlock()

	delay()
	l.Lock()
	value = 17
	Put(CreateTuple(270, value), rsp, targets1)
	log("p8 w:%d", value)
	countwrites(value, 8, 17)
	l.Unlock()

	delay()
	l.Lock()
	value = 4
	Put(CreateTuple(49, value), rsp, targets2)
	log("p8 w:%d", value)
	countwrites(value, 8, 4)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(341, &value), rsp, s8)
	log("p8 r:%d", value)
	countreads(value, 8, 22)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(501, &value), rsp, s8)
	log("p8 r:%d", value)
	countreads(value, 8, 32)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(308, &value), rsp, s8)
	log("p8 r:%d", value)
	countreads(value, 8, 20)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(334, &value), rsp, s8)
	log("p8 r:%d", value)
	countreads(value, 8, 21)
	l.Unlock()
}

func p9() {
	targets3 := make([]string, 0)
	targets2 := make([]string, 0)
	targets1 := make([]string, 2)
	targets1[1] = uri[s20]
	targets1[0] = uri[s15]
	targets0 := make([]string, 0)

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(72, &value), rsp, s9)
	log("p9 r:%d", value)
	countreads(value, 9, 5)
	l.Unlock()

	delay()
	l.Lock()
	value = 22
	Put(CreateTuple(349, value), rsp, targets0)
	log("p9 w:%d", value)
	countwrites(value, 9, 22)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(39, &value), rsp, s9)
	log("p9 r:%d", value)
	countreads(value, 9, 3)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(31, &value), rsp, s9)
	log("p9 r:%d", value)
	countreads(value, 9, 2)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(492, &value), rsp, s9)
	log("p9 r:%d", value)
	countreads(value, 9, 31)
	l.Unlock()

	delay()
	l.Lock()
	value = 27
	Put(CreateTuple(432, value), rsp, targets1)
	log("p9 w:%d", value)
	countwrites(value, 9, 27)
	l.Unlock()

	delay()
	l.Lock()
	value = 16
	Put(CreateTuple(242, value), rsp, targets2)
	log("p9 w:%d", value)
	countwrites(value, 9, 16)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(422, &value), rsp, s9)
	log("p9 r:%d", value)
	countreads(value, 9, 27)
	l.Unlock()

	delay()
	l.Lock()
	value = 2
	Put(CreateTuple(27, value), rsp, targets3)
	log("p9 w:%d", value)
	countwrites(value, 9, 2)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(43, &value), rsp, s9)
	log("p9 r:%d", value)
	countreads(value, 9, 3)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(224, &value), rsp, s9)
	log("p9 r:%d", value)
	countreads(value, 9, 14)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(505, &value), rsp, s9)
	log("p9 r:%d", value)
	countreads(value, 9, 32)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(377, &value), rsp, s9)
	log("p9 r:%d", value)
	countreads(value, 9, 24)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(414, &value), rsp, s9)
	log("p9 r:%d", value)
	countreads(value, 9, 26)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(500, &value), rsp, s9)
	log("p9 r:%d", value)
	countreads(value, 9, 32)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(17, &value), rsp, s9)
	log("p9 r:%d", value)
	countreads(value, 9, 2)
	l.Unlock()
}

func p10() {
	targets4 := make([]string, 0)
	targets3 := make([]string, 1)
	targets3[0] = uri[s32]
	targets2 := make([]string, 1)
	targets2[0] = uri[s19]
	targets1 := make([]string, 0)
	targets0 := make([]string, 2)
	targets0[1] = uri[s27]
	targets0[0] = uri[s19]

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(344, &value), rsp, s10)
	log("p10 r:%d", value)
	countreads(value, 10, 22)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(97, &value), rsp, s10)
	log("p10 r:%d", value)
	countreads(value, 10, 7)
	l.Unlock()

	delay()
	l.Lock()
	value = 30
	Put(CreateTuple(472, value), rsp, targets0)
	log("p10 w:%d", value)
	countwrites(value, 10, 30)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(207, &value), rsp, s10)
	log("p10 r:%d", value)
	countreads(value, 10, 13)
	l.Unlock()

	delay()
	l.Lock()
	value = 8
	Put(CreateTuple(127, value), rsp, targets1)
	log("p10 w:%d", value)
	countwrites(value, 10, 8)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(273, &value), rsp, s10)
	log("p10 r:%d", value)
	countreads(value, 10, 18)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(509, &value), rsp, s10)
	log("p10 r:%d", value)
	countreads(value, 10, 32)
	l.Unlock()

	delay()
	l.Lock()
	value = 18
	Put(CreateTuple(282, value), rsp, targets2)
	log("p10 w:%d", value)
	countwrites(value, 10, 18)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(398, &value), rsp, s10)
	log("p10 r:%d", value)
	countreads(value, 10, 25)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(57, &value), rsp, s10)
	log("p10 r:%d", value)
	countreads(value, 10, 4)
	l.Unlock()

	delay()
	l.Lock()
	value = 23
	Put(CreateTuple(368, value), rsp, targets3)
	log("p10 w:%d", value)
	countwrites(value, 10, 23)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(367, &value), rsp, s10)
	log("p10 r:%d", value)
	countreads(value, 10, 23)
	l.Unlock()

	delay()
	l.Lock()
	value = 4
	Put(CreateTuple(56, value), rsp, targets4)
	log("p10 w:%d", value)
	countwrites(value, 10, 4)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(352, &value), rsp, s10)
	log("p10 r:%d", value)
	countreads(value, 10, 22)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(188, &value), rsp, s10)
	log("p10 r:%d", value)
	countreads(value, 10, 12)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(434, &value), rsp, s10)
	log("p10 r:%d", value)
	countreads(value, 10, 28)
	l.Unlock()
}

func p11() {
	targets5 := make([]string, 0)
	targets4 := make([]string, 1)
	targets4[0] = uri[s1]
	targets3 := make([]string, 0)
	targets2 := make([]string, 1)
	targets2[0] = uri[s26]
	targets1 := make([]string, 3)
	targets1[2] = uri[s27]
	targets1[1] = uri[s14]
	targets1[0] = uri[s4]
	targets0 := make([]string, 0)

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(101, &value), rsp, s11)
	log("p11 r:%d", value)
	countreads(value, 11, 7)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(281, &value), rsp, s11)
	log("p11 r:%d", value)
	countreads(value, 11, 18)
	l.Unlock()

	delay()
	l.Lock()
	value = 12
	Put(CreateTuple(192, value), rsp, targets0)
	log("p11 w:%d", value)
	countwrites(value, 11, 12)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(4, &value), rsp, s11)
	log("p11 r:%d", value)
	countreads(value, 11, 1)
	l.Unlock()

	delay()
	l.Lock()
	value = 1
	Put(CreateTuple(1, value), rsp, targets1)
	log("p11 w:%d", value)
	countwrites(value, 11, 1)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(486, &value), rsp, s11)
	log("p11 r:%d", value)
	countreads(value, 11, 31)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(73, &value), rsp, s11)
	log("p11 r:%d", value)
	countreads(value, 11, 5)
	l.Unlock()

	delay()
	l.Lock()
	value = 2
	Put(CreateTuple(29, value), rsp, targets2)
	log("p11 w:%d", value)
	countwrites(value, 11, 2)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(389, &value), rsp, s11)
	log("p11 r:%d", value)
	countreads(value, 11, 25)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(5, &value), rsp, s11)
	log("p11 r:%d", value)
	countreads(value, 11, 1)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(443, &value), rsp, s11)
	log("p11 r:%d", value)
	countreads(value, 11, 28)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(124, &value), rsp, s11)
	log("p11 r:%d", value)
	countreads(value, 11, 8)
	l.Unlock()

	delay()
	l.Lock()
	value = 7
	Put(CreateTuple(99, value), rsp, targets3)
	log("p11 w:%d", value)
	countwrites(value, 11, 7)
	l.Unlock()

	delay()
	l.Lock()
	value = 29
	Put(CreateTuple(459, value), rsp, targets4)
	log("p11 w:%d", value)
	countwrites(value, 11, 29)
	l.Unlock()

	delay()
	l.Lock()
	value = 27
	Put(CreateTuple(429, value), rsp, targets5)
	log("p11 w:%d", value)
	countwrites(value, 11, 27)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(137, &value), rsp, s11)
	log("p11 r:%d", value)
	countreads(value, 11, 9)
	l.Unlock()
}

func p12() {
	targets7 := make([]string, 0)
	targets6 := make([]string, 0)
	targets5 := make([]string, 0)
	targets4 := make([]string, 0)
	targets3 := make([]string, 1)
	targets3[0] = uri[s26]
	targets2 := make([]string, 0)
	targets1 := make([]string, 1)
	targets1[0] = uri[s23]
	targets0 := make([]string, 2)
	targets0[1] = uri[s26]
	targets0[0] = uri[s23]

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 1
	Put(CreateTuple(3, value), rsp, targets0)
	log("p12 w:%d", value)
	countwrites(value, 12, 1)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(51, &value), rsp, s12)
	log("p12 r:%d", value)
	countreads(value, 12, 4)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(187, &value), rsp, s12)
	log("p12 r:%d", value)
	countreads(value, 12, 12)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(471, &value), rsp, s12)
	log("p12 r:%d", value)
	countreads(value, 12, 30)
	l.Unlock()

	delay()
	l.Lock()
	value = 20
	Put(CreateTuple(318, value), rsp, targets1)
	log("p12 w:%d", value)
	countwrites(value, 12, 20)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(476, &value), rsp, s12)
	log("p12 r:%d", value)
	countreads(value, 12, 30)
	l.Unlock()

	delay()
	l.Lock()
	value = 3
	Put(CreateTuple(38, value), rsp, targets2)
	log("p12 w:%d", value)
	countwrites(value, 12, 3)
	l.Unlock()

	delay()
	l.Lock()
	value = 21
	Put(CreateTuple(328, value), rsp, targets3)
	log("p12 w:%d", value)
	countwrites(value, 12, 21)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(243, &value), rsp, s12)
	log("p12 r:%d", value)
	countreads(value, 12, 16)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(144, &value), rsp, s12)
	log("p12 r:%d", value)
	countreads(value, 12, 9)
	l.Unlock()

	delay()
	l.Lock()
	value = 32
	Put(CreateTuple(510, value), rsp, targets4)
	log("p12 w:%d", value)
	countwrites(value, 12, 32)
	l.Unlock()

	delay()
	l.Lock()
	value = 16
	Put(CreateTuple(253, value), rsp, targets5)
	log("p12 w:%d", value)
	countwrites(value, 12, 16)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(223, &value), rsp, s12)
	log("p12 r:%d", value)
	countreads(value, 12, 14)
	l.Unlock()

	delay()
	l.Lock()
	value = 3
	Put(CreateTuple(46, value), rsp, targets6)
	log("p12 w:%d", value)
	countwrites(value, 12, 3)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(392, &value), rsp, s12)
	log("p12 r:%d", value)
	countreads(value, 12, 25)
	l.Unlock()

	delay()
	l.Lock()
	value = 11
	Put(CreateTuple(173, value), rsp, targets7)
	log("p12 w:%d", value)
	countwrites(value, 12, 11)
	l.Unlock()
}

func p13() {
	targets3 := make([]string, 1)
	targets3[0] = uri[s16]
	targets2 := make([]string, 1)
	targets2[0] = uri[s6]
	targets1 := make([]string, 0)
	targets0 := make([]string, 1)
	targets0[0] = uri[s32]

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 27
	Put(CreateTuple(426, value), rsp, targets0)
	log("p13 w:%d", value)
	countwrites(value, 13, 27)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(511, &value), rsp, s13)
	log("p13 r:%d", value)
	countreads(value, 13, 32)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(243, &value), rsp, s13)
	log("p13 r:%d", value)
	countreads(value, 13, 16)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(491, &value), rsp, s13)
	log("p13 r:%d", value)
	countreads(value, 13, 31)
	l.Unlock()

	delay()
	l.Lock()
	value = 21
	Put(CreateTuple(325, value), rsp, targets1)
	log("p13 w:%d", value)
	countwrites(value, 13, 21)
	l.Unlock()

	delay()
	l.Lock()
	value = 14
	Put(CreateTuple(218, value), rsp, targets2)
	log("p13 w:%d", value)
	countwrites(value, 13, 14)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(330, &value), rsp, s13)
	log("p13 r:%d", value)
	countreads(value, 13, 21)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(115, &value), rsp, s13)
	log("p13 r:%d", value)
	countreads(value, 13, 8)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(406, &value), rsp, s13)
	log("p13 r:%d", value)
	countreads(value, 13, 26)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(450, &value), rsp, s13)
	log("p13 r:%d", value)
	countreads(value, 13, 29)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(178, &value), rsp, s13)
	log("p13 r:%d", value)
	countreads(value, 13, 12)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(4, &value), rsp, s13)
	log("p13 r:%d", value)
	countreads(value, 13, 1)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(5, &value), rsp, s13)
	log("p13 r:%d", value)
	countreads(value, 13, 1)
	l.Unlock()

	delay()
	l.Lock()
	value = 7
	Put(CreateTuple(103, value), rsp, targets3)
	log("p13 w:%d", value)
	countwrites(value, 13, 7)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(475, &value), rsp, s13)
	log("p13 r:%d", value)
	countreads(value, 13, 30)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(468, &value), rsp, s13)
	log("p13 r:%d", value)
	countreads(value, 13, 30)
	l.Unlock()
}

func p14() {
	targets2 := make([]string, 0)
	targets1 := make([]string, 0)
	targets0 := make([]string, 0)
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(269, &value), rsp, s14)
	log("p14 r:%d", value)
	countreads(value, 14, 17)
	l.Unlock()

	delay()
	l.Lock()
	value = 10
	Put(CreateTuple(155, value), rsp, targets0)
	log("p14 w:%d", value)
	countwrites(value, 14, 10)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(408, &value), rsp, s14)
	log("p14 r:%d", value)
	countreads(value, 14, 26)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(307, &value), rsp, s14)
	log("p14 r:%d", value)
	countreads(value, 14, 20)
	l.Unlock()

	delay()
	l.Lock()
	value = 27
	Put(CreateTuple(429, value), rsp, targets1)
	log("p14 w:%d", value)
	countwrites(value, 14, 27)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(140, &value), rsp, s14)
	log("p14 r:%d", value)
	countreads(value, 14, 9)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(1, &value), rsp, s14)
	log("p14 r:%d", value)
	countreads(value, 14, 1)
	l.Unlock()

	delay()
	l.Lock()
	value = 28
	Put(CreateTuple(436, value), rsp, targets2)
	log("p14 w:%d", value)
	countwrites(value, 14, 28)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(274, &value), rsp, s14)
	log("p14 r:%d", value)
	countreads(value, 14, 18)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(214, &value), rsp, s14)
	log("p14 r:%d", value)
	countreads(value, 14, 14)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(66, &value), rsp, s14)
	log("p14 r:%d", value)
	countreads(value, 14, 5)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(176, &value), rsp, s14)
	log("p14 r:%d", value)
	countreads(value, 14, 11)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(386, &value), rsp, s14)
	log("p14 r:%d", value)
	countreads(value, 14, 25)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(302, &value), rsp, s14)
	log("p14 r:%d", value)
	countreads(value, 14, 19)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(139, &value), rsp, s14)
	log("p14 r:%d", value)
	countreads(value, 14, 9)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(117, &value), rsp, s14)
	log("p14 r:%d", value)
	countreads(value, 14, 8)
	l.Unlock()
}

func p15() {
	targets3 := make([]string, 1)
	targets3[0] = uri[s12]
	targets2 := make([]string, 0)
	targets1 := make([]string, 0)
	targets0 := make([]string, 0)

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 24
	Put(CreateTuple(374, value), rsp, targets0)
	log("p15 w:%d", value)
	countwrites(value, 15, 24)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(195, &value), rsp, s15)
	log("p15 r:%d", value)
	countreads(value, 15, 13)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(160, &value), rsp, s15)
	log("p15 r:%d", value)
	countreads(value, 15, 10)
	l.Unlock()

	delay()
	l.Lock()
	value = 23
	Put(CreateTuple(360, value), rsp, targets1)
	log("p15 w:%d", value)
	countwrites(value, 15, 23)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(357, &value), rsp, s15)
	log("p15 r:%d", value)
	countreads(value, 15, 23)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(310, &value), rsp, s15)
	log("p15 r:%d", value)
	countreads(value, 15, 20)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(20, &value), rsp, s15)
	log("p15 r:%d", value)
	countreads(value, 15, 2)
	l.Unlock()

	delay()
	l.Lock()
	value = 28
	Put(CreateTuple(439, value), rsp, targets2)
	log("p15 w:%d", value)
	countwrites(value, 15, 28)
	l.Unlock()

	delay()
	l.Lock()
	value = 9
	Put(CreateTuple(144, value), rsp, targets3)
	log("p15 w:%d", value)
	countwrites(value, 15, 9)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(128, &value), rsp, s15)
	log("p15 r:%d", value)
	countreads(value, 15, 8)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(432, &value), rsp, s15)
	log("p15 r:%d", value)
	countreads(value, 15, 27)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(435, &value), rsp, s15)
	log("p15 r:%d", value)
	countreads(value, 15, 28)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(298, &value), rsp, s15)
	log("p15 r:%d", value)
	countreads(value, 15, 19)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(128, &value), rsp, s15)
	log("p15 r:%d", value)
	countreads(value, 15, 8)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(221, &value), rsp, s15)
	log("p15 r:%d", value)
	countreads(value, 15, 14)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(337, &value), rsp, s15)
	log("p15 r:%d", value)
	countreads(value, 15, 22)
	l.Unlock()
}

func p16() {
	targets3 := make([]string, 0)
	targets2 := make([]string, 1)
	targets2[0] = uri[s2]
	targets1 := make([]string, 0)
	targets0 := make([]string, 0)

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 6
	Put(CreateTuple(84, value), rsp, targets0)
	log("p16 w:%d", value)
	countwrites(value, 16, 6)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(154, &value), rsp, s16)
	log("p16 r:%d", value)
	countreads(value, 16, 10)
	l.Unlock()

	delay()
	l.Lock()
	value = 27
	Put(CreateTuple(423, value), rsp, targets1)
	log("p16 w:%d", value)
	countwrites(value, 16, 27)
	l.Unlock()

	delay()
	l.Lock()
	value = 4
	Put(CreateTuple(58, value), rsp, targets2)
	log("p16 w:%d", value)
	countwrites(value, 16, 4)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(340, &value), rsp, s16)
	log("p16 r:%d", value)
	countreads(value, 16, 22)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(103, &value), rsp, s16)
	log("p16 r:%d", value)
	countreads(value, 16, 7)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(111, &value), rsp, s16)
	log("p16 r:%d", value)
	countreads(value, 16, 7)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(228, &value), rsp, s16)
	log("p16 r:%d", value)
	countreads(value, 16, 15)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(394, &value), rsp, s16)
	log("p16 r:%d", value)
	countreads(value, 16, 25)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(91, &value), rsp, s16)
	log("p16 r:%d", value)
	countreads(value, 16, 6)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(174, &value), rsp, s16)
	log("p16 r:%d", value)
	countreads(value, 16, 11)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(359, &value), rsp, s16)
	log("p16 r:%d", value)
	countreads(value, 16, 23)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(183, &value), rsp, s16)
	log("p16 r:%d", value)
	countreads(value, 16, 12)
	l.Unlock()

	delay()
	l.Lock()
	value = 17
	Put(CreateTuple(263, value), rsp, targets3)
	log("p16 w:%d", value)
	countwrites(value, 16, 17)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(306, &value), rsp, s16)
	log("p16 r:%d", value)
	countreads(value, 16, 20)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(288, &value), rsp, s16)
	log("p16 r:%d", value)
	countreads(value, 16, 18)
	l.Unlock()
}

func p17() {
	targets2 := make([]string, 1)
	targets2[0] = uri[s3]
	targets1 := make([]string, 1)
	targets1[0] = uri[s29]
	targets0 := make([]string, 1)
	targets0[0] = uri[s29]

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(204, &value), rsp, s17)
	log("p17 r:%d", value)
	countreads(value, 17, 13)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(428, &value), rsp, s17)
	log("p17 r:%d", value)
	countreads(value, 17, 27)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(186, &value), rsp, s17)
	log("p17 r:%d", value)
	countreads(value, 17, 12)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(464, &value), rsp, s17)
	log("p17 r:%d", value)
	countreads(value, 17, 29)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(93, &value), rsp, s17)
	log("p17 r:%d", value)
	countreads(value, 17, 6)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(157, &value), rsp, s17)
	log("p17 r:%d", value)
	countreads(value, 17, 10)
	l.Unlock()

	delay()
	l.Lock()
	value = 31
	Put(CreateTuple(484, value), rsp, targets0)
	log("p17 w:%d", value)
	countwrites(value, 17, 31)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(441, &value), rsp, s17)
	log("p17 r:%d", value)
	countreads(value, 17, 28)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(278, &value), rsp, s17)
	log("p17 r:%d", value)
	countreads(value, 17, 18)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(130, &value), rsp, s17)
	log("p17 r:%d", value)
	countreads(value, 17, 9)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(77, &value), rsp, s17)
	log("p17 r:%d", value)
	countreads(value, 17, 5)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(273, &value), rsp, s17)
	log("p17 r:%d", value)
	countreads(value, 17, 18)
	l.Unlock()

	delay()
	l.Lock()
	value = 28
	Put(CreateTuple(447, value), rsp, targets1)
	log("p17 w:%d", value)
	countwrites(value, 17, 28)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(239, &value), rsp, s17)
	log("p17 r:%d", value)
	countreads(value, 17, 15)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(405, &value), rsp, s17)
	log("p17 r:%d", value)
	countreads(value, 17, 26)
	l.Unlock()

	delay()
	l.Lock()
	value = 19
	Put(CreateTuple(297, value), rsp, targets2)
	log("p17 w:%d", value)
	countwrites(value, 17, 19)
	l.Unlock()
}

func p18() {
	targets4 := make([]string, 2)
	targets4[1] = uri[s27]
	targets4[0] = uri[s19]
	targets3 := make([]string, 2)
	targets3[1] = uri[s25]
	targets3[0] = uri[s5]
	targets2 := make([]string, 0)
	targets1 := make([]string, 2)
	targets1[1] = uri[s25]
	targets1[0] = uri[s5]
	targets0 := make([]string, 0)

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 19
	Put(CreateTuple(289, value), rsp, targets0)
	log("p18 w:%d", value)
	countwrites(value, 18, 19)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(256, &value), rsp, s18)
	log("p18 r:%d", value)
	countreads(value, 18, 16)
	l.Unlock()

	delay()
	l.Lock()
	value = 1
	Put(CreateTuple(16, value), rsp, targets1)
	log("p18 w:%d", value)
	countwrites(value, 18, 1)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(176, &value), rsp, s18)
	log("p18 r:%d", value)
	countreads(value, 18, 11)
	l.Unlock()

	delay()
	l.Lock()
	value = 18
	Put(CreateTuple(279, value), rsp, targets2)
	log("p18 w:%d", value)
	countwrites(value, 18, 18)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(411, &value), rsp, s18)
	log("p18 r:%d", value)
	countreads(value, 18, 26)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(332, &value), rsp, s18)
	log("p18 r:%d", value)
	countreads(value, 18, 21)
	l.Unlock()

	delay()
	l.Lock()
	value = 1
	Put(CreateTuple(16, value), rsp, targets3)
	log("p18 w:%d", value)
	countwrites(value, 18, 1)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(339, &value), rsp, s18)
	log("p18 r:%d", value)
	countreads(value, 18, 22)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(204, &value), rsp, s18)
	log("p18 r:%d", value)
	countreads(value, 18, 13)
	l.Unlock()

	delay()
	l.Lock()
	value = 30
	Put(CreateTuple(472, value), rsp, targets4)
	log("p18 w:%d", value)
	countwrites(value, 18, 30)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(473, &value), rsp, s18)
	log("p18 r:%d", value)
	countreads(value, 18, 30)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(212, &value), rsp, s18)
	log("p18 r:%d", value)
	countreads(value, 18, 14)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(129, &value), rsp, s18)
	log("p18 r:%d", value)
	countreads(value, 18, 9)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(35, &value), rsp, s18)
	log("p18 r:%d", value)
	countreads(value, 18, 3)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(367, &value), rsp, s18)
	log("p18 r:%d", value)
	countreads(value, 18, 23)
	l.Unlock()
}

func p19() {
	targets4 := make([]string, 1)
	targets4[0] = uri[s6]
	targets3 := make([]string, 1)
	targets3[0] = uri[s17]
	targets2 := make([]string, 0)
	targets1 := make([]string, 2)
	targets1[1] = uri[s9]
	targets1[0] = uri[s2]
	targets0 := make([]string, 0)

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(508, &value), rsp, s19)
	log("p19 r:%d", value)
	countreads(value, 19, 32)
	l.Unlock()

	delay()
	l.Lock()
	value = 16
	Put(CreateTuple(255, value), rsp, targets0)
	log("p19 w:%d", value)
	countwrites(value, 19, 16)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(407, &value), rsp, s19)
	log("p19 r:%d", value)
	countreads(value, 19, 26)
	l.Unlock()

	delay()
	l.Lock()
	value = 26
	Put(CreateTuple(414, value), rsp, targets1)
	log("p19 w:%d", value)
	countwrites(value, 19, 26)
	l.Unlock()

	delay()
	l.Lock()
	value = 9
	Put(CreateTuple(142, value), rsp, targets2)
	log("p19 w:%d", value)
	countwrites(value, 19, 9)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(304, &value), rsp, s19)
	log("p19 r:%d", value)
	countreads(value, 19, 19)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(33, &value), rsp, s19)
	log("p19 r:%d", value)
	countreads(value, 19, 3)
	l.Unlock()

	delay()
	l.Lock()
	value = 12
	Put(CreateTuple(186, value), rsp, targets3)
	log("p19 w:%d", value)
	countwrites(value, 19, 12)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(434, &value), rsp, s19)
	log("p19 r:%d", value)
	countreads(value, 19, 28)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(472, &value), rsp, s19)
	log("p19 r:%d", value)
	countreads(value, 19, 30)
	l.Unlock()

	delay()
	l.Lock()
	value = 18
	Put(CreateTuple(280, value), rsp, targets4)
	log("p19 w:%d", value)
	countwrites(value, 19, 18)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(245, &value), rsp, s19)
	log("p19 r:%d", value)
	countreads(value, 19, 16)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(332, &value), rsp, s19)
	log("p19 r:%d", value)
	countreads(value, 19, 21)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(469, &value), rsp, s19)
	log("p19 r:%d", value)
	countreads(value, 19, 30)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(91, &value), rsp, s19)
	log("p19 r:%d", value)
	countreads(value, 19, 6)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(282, &value), rsp, s19)
	log("p19 r:%d", value)
	countreads(value, 19, 18)
	l.Unlock()
}

func p20() {
	targets4 := make([]string, 2)
	targets4[1] = uri[s21]
	targets4[0] = uri[s6]
	targets3 := make([]string, 0)
	targets2 := make([]string, 0)
	targets1 := make([]string, 1)
	targets1[0] = uri[s19]
	targets0 := make([]string, 0)

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(322, &value), rsp, s20)
	log("p20 r:%d", value)
	countreads(value, 20, 21)
	l.Unlock()

	delay()
	l.Lock()
	value = 22
	Put(CreateTuple(350, value), rsp, targets0)
	log("p20 w:%d", value)
	countwrites(value, 20, 22)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(70, &value), rsp, s20)
	log("p20 r:%d", value)
	countreads(value, 20, 5)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(458, &value), rsp, s20)
	log("p20 r:%d", value)
	countreads(value, 20, 29)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(413, &value), rsp, s20)
	log("p20 r:%d", value)
	countreads(value, 20, 26)
	l.Unlock()

	delay()
	l.Lock()
	value = 19
	Put(CreateTuple(304, value), rsp, targets1)
	log("p20 w:%d", value)
	countwrites(value, 20, 19)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(432, &value), rsp, s20)
	log("p20 r:%d", value)
	countreads(value, 20, 27)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(252, &value), rsp, s20)
	log("p20 r:%d", value)
	countreads(value, 20, 16)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(424, &value), rsp, s20)
	log("p20 r:%d", value)
	countreads(value, 20, 27)
	l.Unlock()

	delay()
	l.Lock()
	value = 13
	Put(CreateTuple(205, value), rsp, targets2)
	log("p20 w:%d", value)
	countwrites(value, 20, 13)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(54, &value), rsp, s20)
	log("p20 r:%d", value)
	countreads(value, 20, 4)
	l.Unlock()

	delay()
	l.Lock()
	value = 15
	Put(CreateTuple(230, value), rsp, targets3)
	log("p20 w:%d", value)
	countwrites(value, 20, 15)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(206, &value), rsp, s20)
	log("p20 r:%d", value)
	countreads(value, 20, 13)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(313, &value), rsp, s20)
	log("p20 r:%d", value)
	countreads(value, 20, 20)
	l.Unlock()

	delay()
	l.Lock()
	value = 6
	Put(CreateTuple(92, value), rsp, targets4)
	log("p20 w:%d", value)
	countwrites(value, 20, 6)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(288, &value), rsp, s20)
	log("p20 r:%d", value)
	countreads(value, 20, 18)
	l.Unlock()
}

func p21() {
	targets1 := make([]string, 0)
	targets0 := make([]string, 0)
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(247, &value), rsp, s21)
	log("p21 r:%d", value)
	countreads(value, 21, 16)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(296, &value), rsp, s21)
	log("p21 r:%d", value)
	countreads(value, 21, 19)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(8, &value), rsp, s21)
	log("p21 r:%d", value)
	countreads(value, 21, 1)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(481, &value), rsp, s21)
	log("p21 r:%d", value)
	countreads(value, 21, 31)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(475, &value), rsp, s21)
	log("p21 r:%d", value)
	countreads(value, 21, 30)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(44, &value), rsp, s21)
	log("p21 r:%d", value)
	countreads(value, 21, 3)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(92, &value), rsp, s21)
	log("p21 r:%d", value)
	countreads(value, 21, 6)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(288, &value), rsp, s21)
	log("p21 r:%d", value)
	countreads(value, 21, 18)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(83, &value), rsp, s21)
	log("p21 r:%d", value)
	countreads(value, 21, 6)
	l.Unlock()

	delay()
	l.Lock()
	value = 11
	Put(CreateTuple(172, value), rsp, targets0)
	log("p21 w:%d", value)
	countwrites(value, 21, 11)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(237, &value), rsp, s21)
	log("p21 r:%d", value)
	countreads(value, 21, 15)
	l.Unlock()

	delay()
	l.Lock()
	value = 13
	Put(CreateTuple(203, value), rsp, targets1)
	log("p21 w:%d", value)
	countwrites(value, 21, 13)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(6, &value), rsp, s21)
	log("p21 r:%d", value)
	countreads(value, 21, 1)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(117, &value), rsp, s21)
	log("p21 r:%d", value)
	countreads(value, 21, 8)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(303, &value), rsp, s21)
	log("p21 r:%d", value)
	countreads(value, 21, 19)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(55, &value), rsp, s21)
	log("p21 r:%d", value)
	countreads(value, 21, 4)
	l.Unlock()
}

func p22() {
	targets7 := make([]string, 0)
	targets6 := make([]string, 0)
	targets5 := make([]string, 0)
	targets4 := make([]string, 1)
	targets4[0] = uri[s17]
	targets3 := make([]string, 0)
	targets2 := make([]string, 0)
	targets1 := make([]string, 1)
	targets1[0] = uri[s5]
	targets0 := make([]string, 0)

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 30
	Put(CreateTuple(480, value), rsp, targets0)
	log("p22 w:%d", value)
	countwrites(value, 22, 30)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(68, &value), rsp, s22)
	log("p22 r:%d", value)
	countreads(value, 22, 5)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(395, &value), rsp, s22)
	log("p22 r:%d", value)
	countreads(value, 22, 25)
	l.Unlock()

	delay()
	l.Lock()
	value = 6
	Put(CreateTuple(81, value), rsp, targets1)
	log("p22 w:%d", value)
	countwrites(value, 22, 6)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(93, &value), rsp, s22)
	log("p22 r:%d", value)
	countreads(value, 22, 6)
	l.Unlock()

	delay()
	l.Lock()
	value = 32
	Put(CreateTuple(502, value), rsp, targets2)
	log("p22 w:%d", value)
	countwrites(value, 22, 32)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(440, &value), rsp, s22)
	log("p22 r:%d", value)
	countreads(value, 22, 28)
	l.Unlock()

	delay()
	l.Lock()
	value = 15
	Put(CreateTuple(230, value), rsp, targets3)
	log("p22 w:%d", value)
	countwrites(value, 22, 15)
	l.Unlock()

	delay()
	l.Lock()
	value = 18
	Put(CreateTuple(278, value), rsp, targets4)
	log("p22 w:%d", value)
	countwrites(value, 22, 18)
	l.Unlock()

	delay()
	l.Lock()
	value = 21
	Put(CreateTuple(335, value), rsp, targets5)
	log("p22 w:%d", value)
	countwrites(value, 22, 21)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(364, &value), rsp, s22)
	log("p22 r:%d", value)
	countreads(value, 22, 23)
	l.Unlock()

	delay()
	l.Lock()
	value = 31
	Put(CreateTuple(482, value), rsp, targets6)
	log("p22 w:%d", value)
	countwrites(value, 22, 31)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(152, &value), rsp, s22)
	log("p22 r:%d", value)
	countreads(value, 22, 10)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(339, &value), rsp, s22)
	log("p22 r:%d", value)
	countreads(value, 22, 22)
	l.Unlock()

	delay()
	l.Lock()
	value = 25
	Put(CreateTuple(396, value), rsp, targets7)
	log("p22 w:%d", value)
	countwrites(value, 22, 25)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(435, &value), rsp, s22)
	log("p22 r:%d", value)
	countreads(value, 22, 28)
	l.Unlock()
}

func p23() {
	targets4 := make([]string, 1)
	targets4[0] = uri[s32]
	targets3 := make([]string, 2)
	targets3[1] = uri[s24]
	targets3[0] = uri[s6]
	targets2 := make([]string, 0)
	targets1 := make([]string, 2)
	targets1[1] = uri[s30]
	targets1[0] = uri[s21]
	targets0 := make([]string, 2)
	targets0[1] = uri[s27]
	targets0[0] = uri[s2]

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(318, &value), rsp, s23)
	log("p23 r:%d", value)
	countreads(value, 23, 20)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(347, &value), rsp, s23)
	log("p23 r:%d", value)
	countreads(value, 23, 22)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(504, &value), rsp, s23)
	log("p23 r:%d", value)
	countreads(value, 23, 32)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(3, &value), rsp, s23)
	log("p23 r:%d", value)
	countreads(value, 23, 1)
	l.Unlock()

	delay()
	l.Lock()
	value = 24
	Put(CreateTuple(375, value), rsp, targets0)
	log("p23 w:%d", value)
	countwrites(value, 23, 24)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(116, &value), rsp, s23)
	log("p23 r:%d", value)
	countreads(value, 23, 8)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(23, &value), rsp, s23)
	log("p23 r:%d", value)
	countreads(value, 23, 2)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(252, &value), rsp, s23)
	log("p23 r:%d", value)
	countreads(value, 23, 16)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(397, &value), rsp, s23)
	log("p23 r:%d", value)
	countreads(value, 23, 25)
	l.Unlock()

	delay()
	l.Lock()
	value = 31
	Put(CreateTuple(481, value), rsp, targets1)
	log("p23 w:%d", value)
	countwrites(value, 23, 31)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(196, &value), rsp, s23)
	log("p23 r:%d", value)
	countreads(value, 23, 13)
	l.Unlock()

	delay()
	l.Lock()
	value = 30
	Put(CreateTuple(466, value), rsp, targets2)
	log("p23 w:%d", value)
	countwrites(value, 23, 30)
	l.Unlock()

	delay()
	l.Lock()
	value = 15
	Put(CreateTuple(234, value), rsp, targets3)
	log("p23 w:%d", value)
	countwrites(value, 23, 15)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(314, &value), rsp, s23)
	log("p23 r:%d", value)
	countreads(value, 23, 20)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(213, &value), rsp, s23)
	log("p23 r:%d", value)
	countreads(value, 23, 14)
	l.Unlock()

	delay()
	l.Lock()
	value = 21
	Put(CreateTuple(327, value), rsp, targets4)
	log("p23 w:%d", value)
	countwrites(value, 23, 21)
	l.Unlock()
}

func p24() {
	targets3 := make([]string, 0)
	targets2 := make([]string, 2)
	targets2[1] = uri[s30]
	targets2[0] = uri[s18]
	targets1 := make([]string, 1)
	targets1[0] = uri[s2]
	targets0 := make([]string, 0)

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(33, &value), rsp, s24)
	log("p24 r:%d", value)
	countreads(value, 24, 3)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(32, &value), rsp, s24)
	log("p24 r:%d", value)
	countreads(value, 24, 2)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(453, &value), rsp, s24)
	log("p24 r:%d", value)
	countreads(value, 24, 29)
	l.Unlock()

	delay()
	l.Lock()
	value = 10
	Put(CreateTuple(145, value), rsp, targets0)
	log("p24 w:%d", value)
	countwrites(value, 24, 10)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(249, &value), rsp, s24)
	log("p24 r:%d", value)
	countreads(value, 24, 16)
	l.Unlock()

	delay()
	l.Lock()
	value = 10
	Put(CreateTuple(148, value), rsp, targets1)
	log("p24 w:%d", value)
	countwrites(value, 24, 10)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(292, &value), rsp, s24)
	log("p24 r:%d", value)
	countreads(value, 24, 19)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(95, &value), rsp, s24)
	log("p24 r:%d", value)
	countreads(value, 24, 6)
	l.Unlock()

	delay()
	l.Lock()
	value = 9
	Put(CreateTuple(129, value), rsp, targets2)
	log("p24 w:%d", value)
	countwrites(value, 24, 9)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(412, &value), rsp, s24)
	log("p24 r:%d", value)
	countreads(value, 24, 26)
	l.Unlock()

	delay()
	l.Lock()
	value = 25
	Put(CreateTuple(400, value), rsp, targets3)
	log("p24 w:%d", value)
	countwrites(value, 24, 25)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(234, &value), rsp, s24)
	log("p24 r:%d", value)
	countreads(value, 24, 15)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(363, &value), rsp, s24)
	log("p24 r:%d", value)
	countreads(value, 24, 23)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(97, &value), rsp, s24)
	log("p24 r:%d", value)
	countreads(value, 24, 7)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(53, &value), rsp, s24)
	log("p24 r:%d", value)
	countreads(value, 24, 4)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(4, &value), rsp, s24)
	log("p24 r:%d", value)
	countreads(value, 24, 1)
	l.Unlock()
}

func p25() {
	targets3 := make([]string, 1)
	targets3[0] = uri[s23]
	targets2 := make([]string, 1)
	targets2[0] = uri[s26]
	targets1 := make([]string, 1)
	targets1[0] = uri[s10]
	targets0 := make([]string, 2)
	targets0[1] = uri[s21]
	targets0[0] = uri[s13]

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 30
	Put(CreateTuple(475, value), rsp, targets0)
	log("p25 w:%d", value)
	countwrites(value, 25, 30)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(262, &value), rsp, s25)
	log("p25 r:%d", value)
	countreads(value, 25, 17)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(112, &value), rsp, s25)
	log("p25 r:%d", value)
	countreads(value, 25, 7)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(185, &value), rsp, s25)
	log("p25 r:%d", value)
	countreads(value, 25, 12)
	l.Unlock()

	delay()
	l.Lock()
	value = 25
	Put(CreateTuple(398, value), rsp, targets1)
	log("p25 w:%d", value)
	countwrites(value, 25, 25)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(28, &value), rsp, s25)
	log("p25 r:%d", value)
	countreads(value, 25, 2)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(228, &value), rsp, s25)
	log("p25 r:%d", value)
	countreads(value, 25, 15)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(198, &value), rsp, s25)
	log("p25 r:%d", value)
	countreads(value, 25, 13)
	l.Unlock()

	delay()
	l.Lock()
	value = 19
	Put(CreateTuple(294, value), rsp, targets2)
	log("p25 w:%d", value)
	countwrites(value, 25, 19)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(16, &value), rsp, s25)
	log("p25 r:%d", value)
	countreads(value, 25, 1)
	l.Unlock()

	delay()
	l.Lock()
	value = 13
	Put(CreateTuple(196, value), rsp, targets3)
	log("p25 w:%d", value)
	countwrites(value, 25, 13)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(243, &value), rsp, s25)
	log("p25 r:%d", value)
	countreads(value, 25, 16)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(320, &value), rsp, s25)
	log("p25 r:%d", value)
	countreads(value, 25, 20)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(201, &value), rsp, s25)
	log("p25 r:%d", value)
	countreads(value, 25, 13)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(435, &value), rsp, s25)
	log("p25 r:%d", value)
	countreads(value, 25, 28)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(315, &value), rsp, s25)
	log("p25 r:%d", value)
	countreads(value, 25, 20)
	l.Unlock()
}

func p26() {
	targets2 := make([]string, 1)
	targets2[0] = uri[s10]
	targets1 := make([]string, 2)
	targets1[1] = uri[s15]
	targets1[0] = uri[s15]
	targets0 := make([]string, 1)
	targets0[0] = uri[s28]

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(29, &value), rsp, s26)
	log("p26 r:%d", value)
	countreads(value, 26, 2)
	l.Unlock()

	delay()
	l.Lock()
	value = 5
	Put(CreateTuple(65, value), rsp, targets0)
	log("p26 w:%d", value)
	countwrites(value, 26, 5)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(294, &value), rsp, s26)
	log("p26 r:%d", value)
	countreads(value, 26, 19)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(197, &value), rsp, s26)
	log("p26 r:%d", value)
	countreads(value, 26, 13)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(188, &value), rsp, s26)
	log("p26 r:%d", value)
	countreads(value, 26, 12)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(463, &value), rsp, s26)
	log("p26 r:%d", value)
	countreads(value, 26, 29)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(133, &value), rsp, s26)
	log("p26 r:%d", value)
	countreads(value, 26, 9)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(258, &value), rsp, s26)
	log("p26 r:%d", value)
	countreads(value, 26, 17)
	l.Unlock()

	delay()
	l.Lock()
	value = 8
	Put(CreateTuple(128, value), rsp, targets1)
	log("p26 w:%d", value)
	countwrites(value, 26, 8)
	l.Unlock()

	delay()
	l.Lock()
	value = 13
	Put(CreateTuple(207, value), rsp, targets2)
	log("p26 w:%d", value)
	countwrites(value, 26, 13)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(377, &value), rsp, s26)
	log("p26 r:%d", value)
	countreads(value, 26, 24)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(3, &value), rsp, s26)
	log("p26 r:%d", value)
	countreads(value, 26, 1)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(328, &value), rsp, s26)
	log("p26 r:%d", value)
	countreads(value, 26, 21)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(351, &value), rsp, s26)
	log("p26 r:%d", value)
	countreads(value, 26, 22)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(256, &value), rsp, s26)
	log("p26 r:%d", value)
	countreads(value, 26, 16)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(377, &value), rsp, s26)
	log("p26 r:%d", value)
	countreads(value, 26, 24)
	l.Unlock()
}

func p27() {
	targets1 := make([]string, 0)
	targets0 := make([]string, 0)
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 13
	Put(CreateTuple(193, value), rsp, targets0)
	log("p27 w:%d", value)
	countwrites(value, 27, 13)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(224, &value), rsp, s27)
	log("p27 r:%d", value)
	countreads(value, 27, 14)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(315, &value), rsp, s27)
	log("p27 r:%d", value)
	countreads(value, 27, 20)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(329, &value), rsp, s27)
	log("p27 r:%d", value)
	countreads(value, 27, 21)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(433, &value), rsp, s27)
	log("p27 r:%d", value)
	countreads(value, 27, 28)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(501, &value), rsp, s27)
	log("p27 r:%d", value)
	countreads(value, 27, 32)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(472, &value), rsp, s27)
	log("p27 r:%d", value)
	countreads(value, 27, 30)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(224, &value), rsp, s27)
	log("p27 r:%d", value)
	countreads(value, 27, 14)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(101, &value), rsp, s27)
	log("p27 r:%d", value)
	countreads(value, 27, 7)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(375, &value), rsp, s27)
	log("p27 r:%d", value)
	countreads(value, 27, 24)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(296, &value), rsp, s27)
	log("p27 r:%d", value)
	countreads(value, 27, 19)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(346, &value), rsp, s27)
	log("p27 r:%d", value)
	countreads(value, 27, 22)
	l.Unlock()

	delay()
	l.Lock()
	value = 19
	Put(CreateTuple(291, value), rsp, targets1)
	log("p27 w:%d", value)
	countwrites(value, 27, 19)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(469, &value), rsp, s27)
	log("p27 r:%d", value)
	countreads(value, 27, 30)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(1, &value), rsp, s27)
	log("p27 r:%d", value)
	countreads(value, 27, 1)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(161, &value), rsp, s27)
	log("p27 r:%d", value)
	countreads(value, 27, 11)
	l.Unlock()
}

func p28() {
	targets2 := make([]string, 3)
	targets2[2] = uri[s28]
	targets2[1] = uri[s21]
	targets2[0] = uri[s8]
	targets1 := make([]string, 0)
	targets0 := make([]string, 0)

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 6
	Put(CreateTuple(90, value), rsp, targets0)
	log("p28 w:%d", value)
	countwrites(value, 28, 6)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(479, &value), rsp, s28)
	log("p28 r:%d", value)
	countreads(value, 28, 30)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(118, &value), rsp, s28)
	log("p28 r:%d", value)
	countreads(value, 28, 8)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(65, &value), rsp, s28)
	log("p28 r:%d", value)
	countreads(value, 28, 5)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(358, &value), rsp, s28)
	log("p28 r:%d", value)
	countreads(value, 28, 23)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(248, &value), rsp, s28)
	log("p28 r:%d", value)
	countreads(value, 28, 16)
	l.Unlock()

	delay()
	l.Lock()
	value = 32
	Put(CreateTuple(497, value), rsp, targets1)
	log("p28 w:%d", value)
	countwrites(value, 28, 32)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(252, &value), rsp, s28)
	log("p28 r:%d", value)
	countreads(value, 28, 16)
	l.Unlock()

	delay()
	l.Lock()
	value = 4
	Put(CreateTuple(55, value), rsp, targets2)
	log("p28 w:%d", value)
	countwrites(value, 28, 4)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(146, &value), rsp, s28)
	log("p28 r:%d", value)
	countreads(value, 28, 10)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(285, &value), rsp, s28)
	log("p28 r:%d", value)
	countreads(value, 28, 18)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(106, &value), rsp, s28)
	log("p28 r:%d", value)
	countreads(value, 28, 7)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(471, &value), rsp, s28)
	log("p28 r:%d", value)
	countreads(value, 28, 30)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(48, &value), rsp, s28)
	log("p28 r:%d", value)
	countreads(value, 28, 3)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(55, &value), rsp, s28)
	log("p28 r:%d", value)
	countreads(value, 28, 4)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(488, &value), rsp, s28)
	log("p28 r:%d", value)
	countreads(value, 28, 31)
	l.Unlock()
}

func p29() {
	targets5 := make([]string, 2)
	targets5[1] = uri[s10]
	targets5[0] = uri[s5]
	targets4 := make([]string, 2)
	targets4[1] = uri[s13]
	targets4[0] = uri[s11]
	targets3 := make([]string, 0)
	targets2 := make([]string, 2)
	targets2[1] = uri[s30]
	targets2[0] = uri[s21]
	targets1 := make([]string, 3)
	targets1[2] = uri[s9]
	targets1[1] = uri[s7]
	targets1[0] = uri[s1]
	targets0 := make([]string, 1)
	targets0[0] = uri[s25]

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 13
	Put(CreateTuple(201, value), rsp, targets0)
	log("p29 w:%d", value)
	countwrites(value, 29, 13)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(447, &value), rsp, s29)
	log("p29 r:%d", value)
	countreads(value, 29, 28)
	l.Unlock()

	delay()
	l.Lock()
	value = 3
	Put(CreateTuple(39, value), rsp, targets1)
	log("p29 w:%d", value)
	countwrites(value, 29, 3)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(28, &value), rsp, s29)
	log("p29 r:%d", value)
	countreads(value, 29, 2)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(484, &value), rsp, s29)
	log("p29 r:%d", value)
	countreads(value, 29, 31)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(10, &value), rsp, s29)
	log("p29 r:%d", value)
	countreads(value, 29, 1)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(91, &value), rsp, s29)
	log("p29 r:%d", value)
	countreads(value, 29, 6)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(356, &value), rsp, s29)
	log("p29 r:%d", value)
	countreads(value, 29, 23)
	l.Unlock()

	delay()
	l.Lock()
	value = 31
	Put(CreateTuple(481, value), rsp, targets2)
	log("p29 w:%d", value)
	countwrites(value, 29, 31)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(62, &value), rsp, s29)
	log("p29 r:%d", value)
	countreads(value, 29, 4)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(183, &value), rsp, s29)
	log("p29 r:%d", value)
	countreads(value, 29, 12)
	l.Unlock()

	delay()
	l.Lock()
	value = 21
	Put(CreateTuple(326, value), rsp, targets3)
	log("p29 w:%d", value)
	countwrites(value, 29, 21)
	l.Unlock()

	delay()
	l.Lock()
	value = 1
	Put(CreateTuple(5, value), rsp, targets4)
	log("p29 w:%d", value)
	countwrites(value, 29, 1)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(473, &value), rsp, s29)
	log("p29 r:%d", value)
	countreads(value, 29, 30)
	l.Unlock()

	delay()
	l.Lock()
	value = 32
	Put(CreateTuple(509, value), rsp, targets5)
	log("p29 w:%d", value)
	countwrites(value, 29, 32)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(130, &value), rsp, s29)
	log("p29 r:%d", value)
	countreads(value, 29, 9)
	l.Unlock()
}

func p30() {
	targets3 := make([]string, 1)
	targets3[0] = uri[s26]
	targets2 := make([]string, 1)
	targets2[0] = uri[s12]
	targets1 := make([]string, 0)
	targets0 := make([]string, 3)
	targets0[2] = uri[s26]
	targets0[1] = uri[s10]
	targets0[0] = uri[s7]

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(143, &value), rsp, s30)
	log("p30 r:%d", value)
	countreads(value, 30, 9)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(373, &value), rsp, s30)
	log("p30 r:%d", value)
	countreads(value, 30, 24)
	l.Unlock()

	delay()
	l.Lock()
	value = 12
	Put(CreateTuple(188, value), rsp, targets0)
	log("p30 w:%d", value)
	countwrites(value, 30, 12)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(149, &value), rsp, s30)
	log("p30 r:%d", value)
	countreads(value, 30, 10)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(153, &value), rsp, s30)
	log("p30 r:%d", value)
	countreads(value, 30, 10)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(130, &value), rsp, s30)
	log("p30 r:%d", value)
	countreads(value, 30, 9)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(481, &value), rsp, s30)
	log("p30 r:%d", value)
	countreads(value, 30, 31)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(215, &value), rsp, s30)
	log("p30 r:%d", value)
	countreads(value, 30, 14)
	l.Unlock()

	delay()
	l.Lock()
	value = 9
	Put(CreateTuple(134, value), rsp, targets1)
	log("p30 w:%d", value)
	countwrites(value, 30, 9)
	l.Unlock()

	delay()
	l.Lock()
	value = 25
	Put(CreateTuple(392, value), rsp, targets2)
	log("p30 w:%d", value)
	countwrites(value, 30, 25)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(129, &value), rsp, s30)
	log("p30 r:%d", value)
	countreads(value, 30, 9)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(223, &value), rsp, s30)
	log("p30 r:%d", value)
	countreads(value, 30, 14)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(420, &value), rsp, s30)
	log("p30 r:%d", value)
	countreads(value, 30, 27)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(138, &value), rsp, s30)
	log("p30 r:%d", value)
	countreads(value, 30, 9)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(6, &value), rsp, s30)
	log("p30 r:%d", value)
	countreads(value, 30, 1)
	l.Unlock()

	delay()
	l.Lock()
	value = 29
	Put(CreateTuple(463, value), rsp, targets3)
	log("p30 w:%d", value)
	countwrites(value, 30, 29)
	l.Unlock()
}

func p31() {
	targets5 := make([]string, 0)
	targets4 := make([]string, 0)
	targets3 := make([]string, 0)
	targets2 := make([]string, 3)
	targets2[2] = uri[s27]
	targets2[1] = uri[s25]
	targets2[0] = uri[s5]
	targets1 := make([]string, 0)
	targets0 := make([]string, 0)

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(369, &value), rsp, s31)
	log("p31 r:%d", value)
	countreads(value, 31, 24)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(286, &value), rsp, s31)
	log("p31 r:%d", value)
	countreads(value, 31, 18)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(471, &value), rsp, s31)
	log("p31 r:%d", value)
	countreads(value, 31, 30)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(437, &value), rsp, s31)
	log("p31 r:%d", value)
	countreads(value, 31, 28)
	l.Unlock()

	delay()
	l.Lock()
	value = 17
	Put(CreateTuple(271, value), rsp, targets0)
	log("p31 w:%d", value)
	countwrites(value, 31, 17)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(37, &value), rsp, s31)
	log("p31 r:%d", value)
	countreads(value, 31, 3)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(389, &value), rsp, s31)
	log("p31 r:%d", value)
	countreads(value, 31, 25)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(448, &value), rsp, s31)
	log("p31 r:%d", value)
	countreads(value, 31, 28)
	l.Unlock()

	delay()
	l.Lock()
	value = 16
	Put(CreateTuple(246, value), rsp, targets1)
	log("p31 w:%d", value)
	countwrites(value, 31, 16)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(133, &value), rsp, s31)
	log("p31 r:%d", value)
	countreads(value, 31, 9)
	l.Unlock()

	delay()
	l.Lock()
	value = 20
	Put(CreateTuple(315, value), rsp, targets2)
	log("p31 w:%d", value)
	countwrites(value, 31, 20)
	l.Unlock()

	delay()
	l.Lock()
	value = 17
	Put(CreateTuple(271, value), rsp, targets3)
	log("p31 w:%d", value)
	countwrites(value, 31, 17)
	l.Unlock()

	delay()
	l.Lock()
	value = 15
	Put(CreateTuple(230, value), rsp, targets4)
	log("p31 w:%d", value)
	countwrites(value, 31, 15)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(177, &value), rsp, s31)
	log("p31 r:%d", value)
	countreads(value, 31, 12)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(251, &value), rsp, s31)
	log("p31 r:%d", value)
	countreads(value, 31, 16)
	l.Unlock()

	delay()
	l.Lock()
	value = 31
	Put(CreateTuple(494, value), rsp, targets5)
	log("p31 w:%d", value)
	countwrites(value, 31, 31)
	l.Unlock()
}

func p32() {
	targets4 := make([]string, 0)
	targets3 := make([]string, 0)
	targets2 := make([]string, 1)
	targets2[0] = uri[s27]
	targets1 := make([]string, 1)
	targets1[0] = uri[s25]
	targets0 := make([]string, 3)
	targets0[2] = uri[s26]
	targets0[1] = uri[s10]
	targets0[0] = uri[s7]

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(48, &value), rsp, s32)
	log("p32 r:%d", value)
	countreads(value, 32, 3)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(348, &value), rsp, s32)
	log("p32 r:%d", value)
	countreads(value, 32, 22)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(327, &value), rsp, s32)
	log("p32 r:%d", value)
	countreads(value, 32, 21)
	l.Unlock()

	delay()
	l.Lock()
	value = 12
	Put(CreateTuple(188, value), rsp, targets0)
	log("p32 w:%d", value)
	countwrites(value, 32, 12)
	l.Unlock()

	delay()
	l.Lock()
	value = 13
	Put(CreateTuple(201, value), rsp, targets1)
	log("p32 w:%d", value)
	countwrites(value, 32, 13)
	l.Unlock()

	delay()
	l.Lock()
	value = 21
	Put(CreateTuple(329, value), rsp, targets2)
	log("p32 w:%d", value)
	countwrites(value, 32, 21)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(28, &value), rsp, s32)
	log("p32 r:%d", value)
	countreads(value, 32, 2)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(404, &value), rsp, s32)
	log("p32 r:%d", value)
	countreads(value, 32, 26)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(308, &value), rsp, s32)
	log("p32 r:%d", value)
	countreads(value, 32, 20)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(331, &value), rsp, s32)
	log("p32 r:%d", value)
	countreads(value, 32, 21)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(368, &value), rsp, s32)
	log("p32 r:%d", value)
	countreads(value, 32, 23)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(425, &value), rsp, s32)
	log("p32 r:%d", value)
	countreads(value, 32, 27)
	l.Unlock()

	delay()
	l.Lock()
	value = 8
	Put(CreateTuple(126, value), rsp, targets3)
	log("p32 w:%d", value)
	countwrites(value, 32, 8)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(426, &value), rsp, s32)
	log("p32 r:%d", value)
	countreads(value, 32, 27)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(51, &value), rsp, s32)
	log("p32 r:%d", value)
	countreads(value, 32, 4)
	l.Unlock()

	delay()
	l.Lock()
	value = 31
	Put(CreateTuple(493, value), rsp, targets4)
	log("p32 w:%d", value)
	countwrites(value, 32, 31)
	l.Unlock()
}

func log(format string, a ...interface{}) {
	if debug {
		fmt.Fprintf(os.Stdout, format+"\n", a...)
	}
}

func delay() {
	time.Sleep(time.Duration(rand.Int63n(75)) * time.Millisecond)
}

func countreads(value int, localspace int, targetspace int) {
	if value == -1 {
		reads_insuccess += 1
	} else {
		reads_success += 1
	}

	if localspace == targetspace {
		reads_local += 1
	} else {
		reads_remote += 1
	}
}

func countwrites(value int, localspace int, targetspace int) {
	writestotal += 1

	if localspace == targetspace {
		writes_local += 1
	} else {
		writes_remote += 1
	}
}
