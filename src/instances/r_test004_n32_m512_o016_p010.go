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
	targets0 := make([]string, 1)
	targets0[0] = uri[s32]

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(228, &value), rsp, s1)
	log("p1 r:%d", value)
	countreads(value, 1, 15)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(489, &value), rsp, s1)
	log("p1 r:%d", value)
	countreads(value, 1, 31)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(319, &value), rsp, s1)
	log("p1 r:%d", value)
	countreads(value, 1, 20)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(471, &value), rsp, s1)
	log("p1 r:%d", value)
	countreads(value, 1, 30)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(197, &value), rsp, s1)
	log("p1 r:%d", value)
	countreads(value, 1, 13)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(53, &value), rsp, s1)
	log("p1 r:%d", value)
	countreads(value, 1, 4)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(434, &value), rsp, s1)
	log("p1 r:%d", value)
	countreads(value, 1, 28)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(203, &value), rsp, s1)
	log("p1 r:%d", value)
	countreads(value, 1, 13)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(508, &value), rsp, s1)
	log("p1 r:%d", value)
	countreads(value, 1, 32)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(392, &value), rsp, s1)
	log("p1 r:%d", value)
	countreads(value, 1, 25)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(114, &value), rsp, s1)
	log("p1 r:%d", value)
	countreads(value, 1, 8)
	l.Unlock()

	delay()
	l.Lock()
	value = 24
	Put(CreateTuple(374, value), rsp, targets0)
	log("p1 w:%d", value)
	countwrites(value, 1, 24)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(135, &value), rsp, s1)
	log("p1 r:%d", value)
	countreads(value, 1, 9)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(483, &value), rsp, s1)
	log("p1 r:%d", value)
	countreads(value, 1, 31)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(309, &value), rsp, s1)
	log("p1 r:%d", value)
	countreads(value, 1, 20)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(85, &value), rsp, s1)
	log("p1 r:%d", value)
	countreads(value, 1, 6)
	l.Unlock()
}

func p2() {
	targets1 := make([]string, 1)
	targets1[0] = uri[s29]
	targets0 := make([]string, 0)

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 30
	Put(CreateTuple(467, value), rsp, targets0)
	log("p2 w:%d", value)
	countwrites(value, 2, 30)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(257, &value), rsp, s2)
	log("p2 r:%d", value)
	countreads(value, 2, 17)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(479, &value), rsp, s2)
	log("p2 r:%d", value)
	countreads(value, 2, 30)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(49, &value), rsp, s2)
	log("p2 r:%d", value)
	countreads(value, 2, 4)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(390, &value), rsp, s2)
	log("p2 r:%d", value)
	countreads(value, 2, 25)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(247, &value), rsp, s2)
	log("p2 r:%d", value)
	countreads(value, 2, 16)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(293, &value), rsp, s2)
	log("p2 r:%d", value)
	countreads(value, 2, 19)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(223, &value), rsp, s2)
	log("p2 r:%d", value)
	countreads(value, 2, 14)
	l.Unlock()

	delay()
	l.Lock()
	value = 2
	Put(CreateTuple(30, value), rsp, targets1)
	log("p2 w:%d", value)
	countwrites(value, 2, 2)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(465, &value), rsp, s2)
	log("p2 r:%d", value)
	countreads(value, 2, 30)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(319, &value), rsp, s2)
	log("p2 r:%d", value)
	countreads(value, 2, 20)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(376, &value), rsp, s2)
	log("p2 r:%d", value)
	countreads(value, 2, 24)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(478, &value), rsp, s2)
	log("p2 r:%d", value)
	countreads(value, 2, 30)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(314, &value), rsp, s2)
	log("p2 r:%d", value)
	countreads(value, 2, 20)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(443, &value), rsp, s2)
	log("p2 r:%d", value)
	countreads(value, 2, 28)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(144, &value), rsp, s2)
	log("p2 r:%d", value)
	countreads(value, 2, 9)
	l.Unlock()
}

func p3() {
	targets1 := make([]string, 1)
	targets1[0] = uri[s21]
	targets0 := make([]string, 0)

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(66, &value), rsp, s3)
	log("p3 r:%d", value)
	countreads(value, 3, 5)
	l.Unlock()

	delay()
	l.Lock()
	value = 25
	Put(CreateTuple(400, value), rsp, targets0)
	log("p3 w:%d", value)
	countwrites(value, 3, 25)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(497, &value), rsp, s3)
	log("p3 r:%d", value)
	countreads(value, 3, 32)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(7, &value), rsp, s3)
	log("p3 r:%d", value)
	countreads(value, 3, 1)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(104, &value), rsp, s3)
	log("p3 r:%d", value)
	countreads(value, 3, 7)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(221, &value), rsp, s3)
	log("p3 r:%d", value)
	countreads(value, 3, 14)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(45, &value), rsp, s3)
	log("p3 r:%d", value)
	countreads(value, 3, 3)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(509, &value), rsp, s3)
	log("p3 r:%d", value)
	countreads(value, 3, 32)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(126, &value), rsp, s3)
	log("p3 r:%d", value)
	countreads(value, 3, 8)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(384, &value), rsp, s3)
	log("p3 r:%d", value)
	countreads(value, 3, 24)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(362, &value), rsp, s3)
	log("p3 r:%d", value)
	countreads(value, 3, 23)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(355, &value), rsp, s3)
	log("p3 r:%d", value)
	countreads(value, 3, 23)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(477, &value), rsp, s3)
	log("p3 r:%d", value)
	countreads(value, 3, 30)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(246, &value), rsp, s3)
	log("p3 r:%d", value)
	countreads(value, 3, 16)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(287, &value), rsp, s3)
	log("p3 r:%d", value)
	countreads(value, 3, 18)
	l.Unlock()

	delay()
	l.Lock()
	value = 29
	Put(CreateTuple(462, value), rsp, targets1)
	log("p3 w:%d", value)
	countwrites(value, 3, 29)
	l.Unlock()
}

func p4() {
	targets0 := make([]string, 2)
	targets0[1] = uri[s17]
	targets0[0] = uri[s11]

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(442, &value), rsp, s4)
	log("p4 r:%d", value)
	countreads(value, 4, 28)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(172, &value), rsp, s4)
	log("p4 r:%d", value)
	countreads(value, 4, 11)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(38, &value), rsp, s4)
	log("p4 r:%d", value)
	countreads(value, 4, 3)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(439, &value), rsp, s4)
	log("p4 r:%d", value)
	countreads(value, 4, 28)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(311, &value), rsp, s4)
	log("p4 r:%d", value)
	countreads(value, 4, 20)
	l.Unlock()

	delay()
	l.Lock()
	value = 15
	Put(CreateTuple(238, value), rsp, targets0)
	log("p4 w:%d", value)
	countwrites(value, 4, 15)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(124, &value), rsp, s4)
	log("p4 r:%d", value)
	countreads(value, 4, 8)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(442, &value), rsp, s4)
	log("p4 r:%d", value)
	countreads(value, 4, 28)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(187, &value), rsp, s4)
	log("p4 r:%d", value)
	countreads(value, 4, 12)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(358, &value), rsp, s4)
	log("p4 r:%d", value)
	countreads(value, 4, 23)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(28, &value), rsp, s4)
	log("p4 r:%d", value)
	countreads(value, 4, 2)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(268, &value), rsp, s4)
	log("p4 r:%d", value)
	countreads(value, 4, 17)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(276, &value), rsp, s4)
	log("p4 r:%d", value)
	countreads(value, 4, 18)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(274, &value), rsp, s4)
	log("p4 r:%d", value)
	countreads(value, 4, 18)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(320, &value), rsp, s4)
	log("p4 r:%d", value)
	countreads(value, 4, 20)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(429, &value), rsp, s4)
	log("p4 r:%d", value)
	countreads(value, 4, 27)
	l.Unlock()
}

func p5() {
	targets3 := make([]string, 1)
	targets3[0] = uri[s13]
	targets2 := make([]string, 0)
	targets1 := make([]string, 0)
	targets0 := make([]string, 2)
	targets0[1] = uri[s20]
	targets0[0] = uri[s7]

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(237, &value), rsp, s5)
	log("p5 r:%d", value)
	countreads(value, 5, 15)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(10, &value), rsp, s5)
	log("p5 r:%d", value)
	countreads(value, 5, 1)
	l.Unlock()

	delay()
	l.Lock()
	value = 1
	Put(CreateTuple(15, value), rsp, targets0)
	log("p5 w:%d", value)
	countwrites(value, 5, 1)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(79, &value), rsp, s5)
	log("p5 r:%d", value)
	countreads(value, 5, 5)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(383, &value), rsp, s5)
	log("p5 r:%d", value)
	countreads(value, 5, 24)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(9, &value), rsp, s5)
	log("p5 r:%d", value)
	countreads(value, 5, 1)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(288, &value), rsp, s5)
	log("p5 r:%d", value)
	countreads(value, 5, 18)
	l.Unlock()

	delay()
	l.Lock()
	value = 32
	Put(CreateTuple(511, value), rsp, targets1)
	log("p5 w:%d", value)
	countwrites(value, 5, 32)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(81, &value), rsp, s5)
	log("p5 r:%d", value)
	countreads(value, 5, 6)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(204, &value), rsp, s5)
	log("p5 r:%d", value)
	countreads(value, 5, 13)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(317, &value), rsp, s5)
	log("p5 r:%d", value)
	countreads(value, 5, 20)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(99, &value), rsp, s5)
	log("p5 r:%d", value)
	countreads(value, 5, 7)
	l.Unlock()

	delay()
	l.Lock()
	value = 30
	Put(CreateTuple(467, value), rsp, targets2)
	log("p5 w:%d", value)
	countwrites(value, 5, 30)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(257, &value), rsp, s5)
	log("p5 r:%d", value)
	countreads(value, 5, 17)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(137, &value), rsp, s5)
	log("p5 r:%d", value)
	countreads(value, 5, 9)
	l.Unlock()

	delay()
	l.Lock()
	value = 13
	Put(CreateTuple(195, value), rsp, targets3)
	log("p5 w:%d", value)
	countwrites(value, 5, 13)
	l.Unlock()
}

func p6() {
	targets0 := make([]string, 1)
	targets0[0] = uri[s7]

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(175, &value), rsp, s6)
	log("p6 r:%d", value)
	countreads(value, 6, 11)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(43, &value), rsp, s6)
	log("p6 r:%d", value)
	countreads(value, 6, 3)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(295, &value), rsp, s6)
	log("p6 r:%d", value)
	countreads(value, 6, 19)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(74, &value), rsp, s6)
	log("p6 r:%d", value)
	countreads(value, 6, 5)
	l.Unlock()

	delay()
	l.Lock()
	value = 7
	Put(CreateTuple(108, value), rsp, targets0)
	log("p6 w:%d", value)
	countwrites(value, 6, 7)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(1, &value), rsp, s6)
	log("p6 r:%d", value)
	countreads(value, 6, 1)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(331, &value), rsp, s6)
	log("p6 r:%d", value)
	countreads(value, 6, 21)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(248, &value), rsp, s6)
	log("p6 r:%d", value)
	countreads(value, 6, 16)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(410, &value), rsp, s6)
	log("p6 r:%d", value)
	countreads(value, 6, 26)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(91, &value), rsp, s6)
	log("p6 r:%d", value)
	countreads(value, 6, 6)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(227, &value), rsp, s6)
	log("p6 r:%d", value)
	countreads(value, 6, 15)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(165, &value), rsp, s6)
	log("p6 r:%d", value)
	countreads(value, 6, 11)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(433, &value), rsp, s6)
	log("p6 r:%d", value)
	countreads(value, 6, 28)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(486, &value), rsp, s6)
	log("p6 r:%d", value)
	countreads(value, 6, 31)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(109, &value), rsp, s6)
	log("p6 r:%d", value)
	countreads(value, 6, 7)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(60, &value), rsp, s6)
	log("p6 r:%d", value)
	countreads(value, 6, 4)
	l.Unlock()
}

func p7() {
	targets0 := make([]string, 0)
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 18
	Put(CreateTuple(280, value), rsp, targets0)
	log("p7 w:%d", value)
	countwrites(value, 7, 18)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(278, &value), rsp, s7)
	log("p7 r:%d", value)
	countreads(value, 7, 18)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(108, &value), rsp, s7)
	log("p7 r:%d", value)
	countreads(value, 7, 7)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(101, &value), rsp, s7)
	log("p7 r:%d", value)
	countreads(value, 7, 7)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(483, &value), rsp, s7)
	log("p7 r:%d", value)
	countreads(value, 7, 31)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(233, &value), rsp, s7)
	log("p7 r:%d", value)
	countreads(value, 7, 15)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(61, &value), rsp, s7)
	log("p7 r:%d", value)
	countreads(value, 7, 4)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(20, &value), rsp, s7)
	log("p7 r:%d", value)
	countreads(value, 7, 2)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(1, &value), rsp, s7)
	log("p7 r:%d", value)
	countreads(value, 7, 1)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(201, &value), rsp, s7)
	log("p7 r:%d", value)
	countreads(value, 7, 13)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(173, &value), rsp, s7)
	log("p7 r:%d", value)
	countreads(value, 7, 11)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(439, &value), rsp, s7)
	log("p7 r:%d", value)
	countreads(value, 7, 28)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(43, &value), rsp, s7)
	log("p7 r:%d", value)
	countreads(value, 7, 3)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(15, &value), rsp, s7)
	log("p7 r:%d", value)
	countreads(value, 7, 1)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(509, &value), rsp, s7)
	log("p7 r:%d", value)
	countreads(value, 7, 32)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(255, &value), rsp, s7)
	log("p7 r:%d", value)
	countreads(value, 7, 16)
	l.Unlock()
}

func p8() {
	targets0 := make([]string, 2)
	targets0[1] = uri[s18]
	targets0[0] = uri[s3]

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(167, &value), rsp, s8)
	log("p8 r:%d", value)
	countreads(value, 8, 11)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(481, &value), rsp, s8)
	log("p8 r:%d", value)
	countreads(value, 8, 31)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(334, &value), rsp, s8)
	log("p8 r:%d", value)
	countreads(value, 8, 21)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(344, &value), rsp, s8)
	log("p8 r:%d", value)
	countreads(value, 8, 22)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(423, &value), rsp, s8)
	log("p8 r:%d", value)
	countreads(value, 8, 27)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(210, &value), rsp, s8)
	log("p8 r:%d", value)
	countreads(value, 8, 14)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(500, &value), rsp, s8)
	log("p8 r:%d", value)
	countreads(value, 8, 32)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(464, &value), rsp, s8)
	log("p8 r:%d", value)
	countreads(value, 8, 29)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(291, &value), rsp, s8)
	log("p8 r:%d", value)
	countreads(value, 8, 19)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(9, &value), rsp, s8)
	log("p8 r:%d", value)
	countreads(value, 8, 1)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(263, &value), rsp, s8)
	log("p8 r:%d", value)
	countreads(value, 8, 17)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(204, &value), rsp, s8)
	log("p8 r:%d", value)
	countreads(value, 8, 13)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(439, &value), rsp, s8)
	log("p8 r:%d", value)
	countreads(value, 8, 28)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(443, &value), rsp, s8)
	log("p8 r:%d", value)
	countreads(value, 8, 28)
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
	value = 5
	Put(CreateTuple(66, value), rsp, targets0)
	log("p8 w:%d", value)
	countwrites(value, 8, 5)
	l.Unlock()
}

func p9() {
	targets2 := make([]string, 1)
	targets2[0] = uri[s32]
	targets1 := make([]string, 0)
	targets0 := make([]string, 1)
	targets0[0] = uri[s5]

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(18, &value), rsp, s9)
	log("p9 r:%d", value)
	countreads(value, 9, 2)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(215, &value), rsp, s9)
	log("p9 r:%d", value)
	countreads(value, 9, 14)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(315, &value), rsp, s9)
	log("p9 r:%d", value)
	countreads(value, 9, 20)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(147, &value), rsp, s9)
	log("p9 r:%d", value)
	countreads(value, 9, 10)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(436, &value), rsp, s9)
	log("p9 r:%d", value)
	countreads(value, 9, 28)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(191, &value), rsp, s9)
	log("p9 r:%d", value)
	countreads(value, 9, 12)
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
	QueryP(CreateTuple(319, &value), rsp, s9)
	log("p9 r:%d", value)
	countreads(value, 9, 20)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(230, &value), rsp, s9)
	log("p9 r:%d", value)
	countreads(value, 9, 15)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(172, &value), rsp, s9)
	log("p9 r:%d", value)
	countreads(value, 9, 11)
	l.Unlock()

	delay()
	l.Lock()
	value = 1
	Put(CreateTuple(10, value), rsp, targets0)
	log("p9 w:%d", value)
	countwrites(value, 9, 1)
	l.Unlock()

	delay()
	l.Lock()
	value = 29
	Put(CreateTuple(453, value), rsp, targets1)
	log("p9 w:%d", value)
	countwrites(value, 9, 29)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(150, &value), rsp, s9)
	log("p9 r:%d", value)
	countreads(value, 9, 10)
	l.Unlock()

	delay()
	l.Lock()
	value = 8
	Put(CreateTuple(115, value), rsp, targets2)
	log("p9 w:%d", value)
	countwrites(value, 9, 8)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(20, &value), rsp, s9)
	log("p9 r:%d", value)
	countreads(value, 9, 2)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(241, &value), rsp, s9)
	log("p9 r:%d", value)
	countreads(value, 9, 16)
	l.Unlock()
}

func p10() {
	targets3 := make([]string, 1)
	targets3[0] = uri[s27]
	targets2 := make([]string, 1)
	targets2[0] = uri[s17]
	targets1 := make([]string, 0)
	targets0 := make([]string, 2)
	targets0[1] = uri[s18]
	targets0[0] = uri[s10]

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 19
	Put(CreateTuple(302, value), rsp, targets0)
	log("p10 w:%d", value)
	countwrites(value, 10, 19)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(98, &value), rsp, s10)
	log("p10 r:%d", value)
	countreads(value, 10, 7)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(454, &value), rsp, s10)
	log("p10 r:%d", value)
	countreads(value, 10, 29)
	l.Unlock()

	delay()
	l.Lock()
	value = 20
	Put(CreateTuple(305, value), rsp, targets1)
	log("p10 w:%d", value)
	countwrites(value, 10, 20)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(116, &value), rsp, s10)
	log("p10 r:%d", value)
	countreads(value, 10, 8)
	l.Unlock()

	delay()
	l.Lock()
	value = 24
	Put(CreateTuple(379, value), rsp, targets2)
	log("p10 w:%d", value)
	countwrites(value, 10, 24)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(417, &value), rsp, s10)
	log("p10 r:%d", value)
	countreads(value, 10, 27)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(346, &value), rsp, s10)
	log("p10 r:%d", value)
	countreads(value, 10, 22)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(393, &value), rsp, s10)
	log("p10 r:%d", value)
	countreads(value, 10, 25)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(401, &value), rsp, s10)
	log("p10 r:%d", value)
	countreads(value, 10, 26)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(435, &value), rsp, s10)
	log("p10 r:%d", value)
	countreads(value, 10, 28)
	l.Unlock()

	delay()
	l.Lock()
	value = 6
	Put(CreateTuple(82, value), rsp, targets3)
	log("p10 w:%d", value)
	countwrites(value, 10, 6)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(7, &value), rsp, s10)
	log("p10 r:%d", value)
	countreads(value, 10, 1)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(302, &value), rsp, s10)
	log("p10 r:%d", value)
	countreads(value, 10, 19)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(406, &value), rsp, s10)
	log("p10 r:%d", value)
	countreads(value, 10, 26)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(289, &value), rsp, s10)
	log("p10 r:%d", value)
	countreads(value, 10, 19)
	l.Unlock()
}

func p11() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(364, &value), rsp, s11)
	log("p11 r:%d", value)
	countreads(value, 11, 23)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(238, &value), rsp, s11)
	log("p11 r:%d", value)
	countreads(value, 11, 15)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(463, &value), rsp, s11)
	log("p11 r:%d", value)
	countreads(value, 11, 29)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(366, &value), rsp, s11)
	log("p11 r:%d", value)
	countreads(value, 11, 23)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(386, &value), rsp, s11)
	log("p11 r:%d", value)
	countreads(value, 11, 25)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(414, &value), rsp, s11)
	log("p11 r:%d", value)
	countreads(value, 11, 26)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(38, &value), rsp, s11)
	log("p11 r:%d", value)
	countreads(value, 11, 3)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(263, &value), rsp, s11)
	log("p11 r:%d", value)
	countreads(value, 11, 17)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(397, &value), rsp, s11)
	log("p11 r:%d", value)
	countreads(value, 11, 25)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(251, &value), rsp, s11)
	log("p11 r:%d", value)
	countreads(value, 11, 16)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(165, &value), rsp, s11)
	log("p11 r:%d", value)
	countreads(value, 11, 11)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(375, &value), rsp, s11)
	log("p11 r:%d", value)
	countreads(value, 11, 24)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(510, &value), rsp, s11)
	log("p11 r:%d", value)
	countreads(value, 11, 32)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(117, &value), rsp, s11)
	log("p11 r:%d", value)
	countreads(value, 11, 8)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(203, &value), rsp, s11)
	log("p11 r:%d", value)
	countreads(value, 11, 13)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(350, &value), rsp, s11)
	log("p11 r:%d", value)
	countreads(value, 11, 22)
	l.Unlock()
}

func p12() {
	targets1 := make([]string, 1)
	targets1[0] = uri[s4]
	targets0 := make([]string, 1)
	targets0[0] = uri[s17]

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(291, &value), rsp, s12)
	log("p12 r:%d", value)
	countreads(value, 12, 19)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(359, &value), rsp, s12)
	log("p12 r:%d", value)
	countreads(value, 12, 23)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(178, &value), rsp, s12)
	log("p12 r:%d", value)
	countreads(value, 12, 12)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(168, &value), rsp, s12)
	log("p12 r:%d", value)
	countreads(value, 12, 11)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(159, &value), rsp, s12)
	log("p12 r:%d", value)
	countreads(value, 12, 10)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(58, &value), rsp, s12)
	log("p12 r:%d", value)
	countreads(value, 12, 4)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(485, &value), rsp, s12)
	log("p12 r:%d", value)
	countreads(value, 12, 31)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(141, &value), rsp, s12)
	log("p12 r:%d", value)
	countreads(value, 12, 9)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(143, &value), rsp, s12)
	log("p12 r:%d", value)
	countreads(value, 12, 9)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(224, &value), rsp, s12)
	log("p12 r:%d", value)
	countreads(value, 12, 14)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(488, &value), rsp, s12)
	log("p12 r:%d", value)
	countreads(value, 12, 31)
	l.Unlock()

	delay()
	l.Lock()
	value = 31
	Put(CreateTuple(482, value), rsp, targets0)
	log("p12 w:%d", value)
	countwrites(value, 12, 31)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(310, &value), rsp, s12)
	log("p12 r:%d", value)
	countreads(value, 12, 20)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(57, &value), rsp, s12)
	log("p12 r:%d", value)
	countreads(value, 12, 4)
	l.Unlock()

	delay()
	l.Lock()
	value = 8
	Put(CreateTuple(124, value), rsp, targets1)
	log("p12 w:%d", value)
	countwrites(value, 12, 8)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(7, &value), rsp, s12)
	log("p12 r:%d", value)
	countreads(value, 12, 1)
	l.Unlock()
}

func p13() {
	targets2 := make([]string, 1)
	targets2[0] = uri[s9]
	targets1 := make([]string, 0)
	targets0 := make([]string, 0)

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(32, &value), rsp, s13)
	log("p13 r:%d", value)
	countreads(value, 13, 2)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(496, &value), rsp, s13)
	log("p13 r:%d", value)
	countreads(value, 13, 31)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(174, &value), rsp, s13)
	log("p13 r:%d", value)
	countreads(value, 13, 11)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(182, &value), rsp, s13)
	log("p13 r:%d", value)
	countreads(value, 13, 12)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(420, &value), rsp, s13)
	log("p13 r:%d", value)
	countreads(value, 13, 27)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(94, &value), rsp, s13)
	log("p13 r:%d", value)
	countreads(value, 13, 6)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(241, &value), rsp, s13)
	log("p13 r:%d", value)
	countreads(value, 13, 16)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(195, &value), rsp, s13)
	log("p13 r:%d", value)
	countreads(value, 13, 13)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(299, &value), rsp, s13)
	log("p13 r:%d", value)
	countreads(value, 13, 19)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(416, &value), rsp, s13)
	log("p13 r:%d", value)
	countreads(value, 13, 26)
	l.Unlock()

	delay()
	l.Lock()
	value = 16
	Put(CreateTuple(242, value), rsp, targets0)
	log("p13 w:%d", value)
	countwrites(value, 13, 16)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(14, &value), rsp, s13)
	log("p13 r:%d", value)
	countreads(value, 13, 1)
	l.Unlock()

	delay()
	l.Lock()
	value = 5
	Put(CreateTuple(75, value), rsp, targets1)
	log("p13 w:%d", value)
	countwrites(value, 13, 5)
	l.Unlock()

	delay()
	l.Lock()
	value = 2
	Put(CreateTuple(18, value), rsp, targets2)
	log("p13 w:%d", value)
	countwrites(value, 13, 2)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(161, &value), rsp, s13)
	log("p13 r:%d", value)
	countreads(value, 13, 11)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(185, &value), rsp, s13)
	log("p13 r:%d", value)
	countreads(value, 13, 12)
	l.Unlock()
}

func p14() {
	targets1 := make([]string, 1)
	targets1[0] = uri[s25]
	targets0 := make([]string, 0)

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(301, &value), rsp, s14)
	log("p14 r:%d", value)
	countreads(value, 14, 19)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(317, &value), rsp, s14)
	log("p14 r:%d", value)
	countreads(value, 14, 20)
	l.Unlock()

	delay()
	l.Lock()
	value = 23
	Put(CreateTuple(354, value), rsp, targets0)
	log("p14 w:%d", value)
	countwrites(value, 14, 23)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(341, &value), rsp, s14)
	log("p14 r:%d", value)
	countreads(value, 14, 22)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(119, &value), rsp, s14)
	log("p14 r:%d", value)
	countreads(value, 14, 8)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(112, &value), rsp, s14)
	log("p14 r:%d", value)
	countreads(value, 14, 7)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(319, &value), rsp, s14)
	log("p14 r:%d", value)
	countreads(value, 14, 20)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(348, &value), rsp, s14)
	log("p14 r:%d", value)
	countreads(value, 14, 22)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(28, &value), rsp, s14)
	log("p14 r:%d", value)
	countreads(value, 14, 2)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(388, &value), rsp, s14)
	log("p14 r:%d", value)
	countreads(value, 14, 25)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(447, &value), rsp, s14)
	log("p14 r:%d", value)
	countreads(value, 14, 28)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(499, &value), rsp, s14)
	log("p14 r:%d", value)
	countreads(value, 14, 32)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(61, &value), rsp, s14)
	log("p14 r:%d", value)
	countreads(value, 14, 4)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(41, &value), rsp, s14)
	log("p14 r:%d", value)
	countreads(value, 14, 3)
	l.Unlock()

	delay()
	l.Lock()
	value = 23
	Put(CreateTuple(361, value), rsp, targets1)
	log("p14 w:%d", value)
	countwrites(value, 14, 23)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(270, &value), rsp, s14)
	log("p14 r:%d", value)
	countreads(value, 14, 17)
	l.Unlock()
}

func p15() {
	targets2 := make([]string, 0)
	targets1 := make([]string, 0)
	targets0 := make([]string, 1)
	targets0[0] = uri[s11]

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(265, &value), rsp, s15)
	log("p15 r:%d", value)
	countreads(value, 15, 17)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(182, &value), rsp, s15)
	log("p15 r:%d", value)
	countreads(value, 15, 12)
	l.Unlock()

	delay()
	l.Lock()
	value = 8
	Put(CreateTuple(117, value), rsp, targets0)
	log("p15 w:%d", value)
	countwrites(value, 15, 8)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(326, &value), rsp, s15)
	log("p15 r:%d", value)
	countreads(value, 15, 21)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(396, &value), rsp, s15)
	log("p15 r:%d", value)
	countreads(value, 15, 25)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(264, &value), rsp, s15)
	log("p15 r:%d", value)
	countreads(value, 15, 17)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(149, &value), rsp, s15)
	log("p15 r:%d", value)
	countreads(value, 15, 10)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(495, &value), rsp, s15)
	log("p15 r:%d", value)
	countreads(value, 15, 31)
	l.Unlock()

	delay()
	l.Lock()
	value = 16
	Put(CreateTuple(245, value), rsp, targets1)
	log("p15 w:%d", value)
	countwrites(value, 15, 16)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(292, &value), rsp, s15)
	log("p15 r:%d", value)
	countreads(value, 15, 19)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(120, &value), rsp, s15)
	log("p15 r:%d", value)
	countreads(value, 15, 8)
	l.Unlock()

	delay()
	l.Lock()
	value = 6
	Put(CreateTuple(95, value), rsp, targets2)
	log("p15 w:%d", value)
	countwrites(value, 15, 6)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(318, &value), rsp, s15)
	log("p15 r:%d", value)
	countreads(value, 15, 20)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(55, &value), rsp, s15)
	log("p15 r:%d", value)
	countreads(value, 15, 4)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(49, &value), rsp, s15)
	log("p15 r:%d", value)
	countreads(value, 15, 4)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(296, &value), rsp, s15)
	log("p15 r:%d", value)
	countreads(value, 15, 19)
	l.Unlock()
}

func p16() {
	targets0 := make([]string, 0)
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(224, &value), rsp, s16)
	log("p16 r:%d", value)
	countreads(value, 16, 14)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(168, &value), rsp, s16)
	log("p16 r:%d", value)
	countreads(value, 16, 11)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(432, &value), rsp, s16)
	log("p16 r:%d", value)
	countreads(value, 16, 27)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(173, &value), rsp, s16)
	log("p16 r:%d", value)
	countreads(value, 16, 11)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(438, &value), rsp, s16)
	log("p16 r:%d", value)
	countreads(value, 16, 28)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(420, &value), rsp, s16)
	log("p16 r:%d", value)
	countreads(value, 16, 27)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(201, &value), rsp, s16)
	log("p16 r:%d", value)
	countreads(value, 16, 13)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(158, &value), rsp, s16)
	log("p16 r:%d", value)
	countreads(value, 16, 10)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(440, &value), rsp, s16)
	log("p16 r:%d", value)
	countreads(value, 16, 28)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(383, &value), rsp, s16)
	log("p16 r:%d", value)
	countreads(value, 16, 24)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(143, &value), rsp, s16)
	log("p16 r:%d", value)
	countreads(value, 16, 9)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(240, &value), rsp, s16)
	log("p16 r:%d", value)
	countreads(value, 16, 15)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(287, &value), rsp, s16)
	log("p16 r:%d", value)
	countreads(value, 16, 18)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(176, &value), rsp, s16)
	log("p16 r:%d", value)
	countreads(value, 16, 11)
	l.Unlock()

	delay()
	l.Lock()
	value = 11
	Put(CreateTuple(170, value), rsp, targets0)
	log("p16 w:%d", value)
	countwrites(value, 16, 11)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(412, &value), rsp, s16)
	log("p16 r:%d", value)
	countreads(value, 16, 26)
	l.Unlock()
}

func p17() {
	targets1 := make([]string, 2)
	targets1[1] = uri[s24]
	targets1[0] = uri[s3]
	targets0 := make([]string, 1)
	targets0[0] = uri[s3]

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(295, &value), rsp, s17)
	log("p17 r:%d", value)
	countreads(value, 17, 19)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(224, &value), rsp, s17)
	log("p17 r:%d", value)
	countreads(value, 17, 14)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(482, &value), rsp, s17)
	log("p17 r:%d", value)
	countreads(value, 17, 31)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(76, &value), rsp, s17)
	log("p17 r:%d", value)
	countreads(value, 17, 5)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(238, &value), rsp, s17)
	log("p17 r:%d", value)
	countreads(value, 17, 15)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(512, &value), rsp, s17)
	log("p17 r:%d", value)
	countreads(value, 17, 32)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(503, &value), rsp, s17)
	log("p17 r:%d", value)
	countreads(value, 17, 32)
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
	value = 23
	Put(CreateTuple(355, value), rsp, targets0)
	log("p17 w:%d", value)
	countwrites(value, 17, 23)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(202, &value), rsp, s17)
	log("p17 r:%d", value)
	countreads(value, 17, 13)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(379, &value), rsp, s17)
	log("p17 r:%d", value)
	countreads(value, 17, 24)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(313, &value), rsp, s17)
	log("p17 r:%d", value)
	countreads(value, 17, 20)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(436, &value), rsp, s17)
	log("p17 r:%d", value)
	countreads(value, 17, 28)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(265, &value), rsp, s17)
	log("p17 r:%d", value)
	countreads(value, 17, 17)
	l.Unlock()

	delay()
	l.Lock()
	value = 24
	Put(CreateTuple(384, value), rsp, targets1)
	log("p17 w:%d", value)
	countwrites(value, 17, 24)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(57, &value), rsp, s17)
	log("p17 r:%d", value)
	countreads(value, 17, 4)
	l.Unlock()
}

func p18() {
	targets0 := make([]string, 1)
	targets0[0] = uri[s8]

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(365, &value), rsp, s18)
	log("p18 r:%d", value)
	countreads(value, 18, 23)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(137, &value), rsp, s18)
	log("p18 r:%d", value)
	countreads(value, 18, 9)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(475, &value), rsp, s18)
	log("p18 r:%d", value)
	countreads(value, 18, 30)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(154, &value), rsp, s18)
	log("p18 r:%d", value)
	countreads(value, 18, 10)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(100, &value), rsp, s18)
	log("p18 r:%d", value)
	countreads(value, 18, 7)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(233, &value), rsp, s18)
	log("p18 r:%d", value)
	countreads(value, 18, 15)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(66, &value), rsp, s18)
	log("p18 r:%d", value)
	countreads(value, 18, 5)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(446, &value), rsp, s18)
	log("p18 r:%d", value)
	countreads(value, 18, 28)
	l.Unlock()

	delay()
	l.Lock()
	value = 11
	Put(CreateTuple(167, value), rsp, targets0)
	log("p18 w:%d", value)
	countwrites(value, 18, 11)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(72, &value), rsp, s18)
	log("p18 r:%d", value)
	countreads(value, 18, 5)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(302, &value), rsp, s18)
	log("p18 r:%d", value)
	countreads(value, 18, 19)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(29, &value), rsp, s18)
	log("p18 r:%d", value)
	countreads(value, 18, 2)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(179, &value), rsp, s18)
	log("p18 r:%d", value)
	countreads(value, 18, 12)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(7, &value), rsp, s18)
	log("p18 r:%d", value)
	countreads(value, 18, 1)
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
	value = -1
	QueryP(CreateTuple(449, &value), rsp, s18)
	log("p18 r:%d", value)
	countreads(value, 18, 29)
	l.Unlock()
}

func p19() {
	targets1 := make([]string, 2)
	targets1[1] = uri[s17]
	targets1[0] = uri[s9]
	targets0 := make([]string, 1)
	targets0[0] = uri[s11]

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(114, &value), rsp, s19)
	log("p19 r:%d", value)
	countreads(value, 19, 8)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(369, &value), rsp, s19)
	log("p19 r:%d", value)
	countreads(value, 19, 24)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(351, &value), rsp, s19)
	log("p19 r:%d", value)
	countreads(value, 19, 22)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(269, &value), rsp, s19)
	log("p19 r:%d", value)
	countreads(value, 19, 17)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(339, &value), rsp, s19)
	log("p19 r:%d", value)
	countreads(value, 19, 22)
	l.Unlock()

	delay()
	l.Lock()
	value = 8
	Put(CreateTuple(117, value), rsp, targets0)
	log("p19 w:%d", value)
	countwrites(value, 19, 8)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(333, &value), rsp, s19)
	log("p19 r:%d", value)
	countreads(value, 19, 21)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(112, &value), rsp, s19)
	log("p19 r:%d", value)
	countreads(value, 19, 7)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(309, &value), rsp, s19)
	log("p19 r:%d", value)
	countreads(value, 19, 20)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(86, &value), rsp, s19)
	log("p19 r:%d", value)
	countreads(value, 19, 6)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(344, &value), rsp, s19)
	log("p19 r:%d", value)
	countreads(value, 19, 22)
	l.Unlock()

	delay()
	l.Lock()
	value = 28
	Put(CreateTuple(436, value), rsp, targets1)
	log("p19 w:%d", value)
	countwrites(value, 19, 28)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(29, &value), rsp, s19)
	log("p19 r:%d", value)
	countreads(value, 19, 2)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(183, &value), rsp, s19)
	log("p19 r:%d", value)
	countreads(value, 19, 12)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(28, &value), rsp, s19)
	log("p19 r:%d", value)
	countreads(value, 19, 2)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(397, &value), rsp, s19)
	log("p19 r:%d", value)
	countreads(value, 19, 25)
	l.Unlock()
}

func p20() {
	targets1 := make([]string, 2)
	targets1[1] = uri[s32]
	targets1[0] = uri[s28]
	targets0 := make([]string, 2)
	targets0[1] = uri[s17]
	targets0[0] = uri[s11]

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(370, &value), rsp, s20)
	log("p20 r:%d", value)
	countreads(value, 20, 24)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(247, &value), rsp, s20)
	log("p20 r:%d", value)
	countreads(value, 20, 16)
	l.Unlock()

	delay()
	l.Lock()
	value = 15
	Put(CreateTuple(238, value), rsp, targets0)
	log("p20 w:%d", value)
	countwrites(value, 20, 15)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(445, &value), rsp, s20)
	log("p20 r:%d", value)
	countreads(value, 20, 28)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(59, &value), rsp, s20)
	log("p20 r:%d", value)
	countreads(value, 20, 4)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(102, &value), rsp, s20)
	log("p20 r:%d", value)
	countreads(value, 20, 7)
	l.Unlock()

	delay()
	l.Lock()
	value = 28
	Put(CreateTuple(444, value), rsp, targets1)
	log("p20 w:%d", value)
	countwrites(value, 20, 28)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(190, &value), rsp, s20)
	log("p20 r:%d", value)
	countreads(value, 20, 12)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(296, &value), rsp, s20)
	log("p20 r:%d", value)
	countreads(value, 20, 19)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(154, &value), rsp, s20)
	log("p20 r:%d", value)
	countreads(value, 20, 10)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(15, &value), rsp, s20)
	log("p20 r:%d", value)
	countreads(value, 20, 1)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(163, &value), rsp, s20)
	log("p20 r:%d", value)
	countreads(value, 20, 11)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(39, &value), rsp, s20)
	log("p20 r:%d", value)
	countreads(value, 20, 3)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(260, &value), rsp, s20)
	log("p20 r:%d", value)
	countreads(value, 20, 17)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(103, &value), rsp, s20)
	log("p20 r:%d", value)
	countreads(value, 20, 7)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(472, &value), rsp, s20)
	log("p20 r:%d", value)
	countreads(value, 20, 30)
	l.Unlock()
}

func p21() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(376, &value), rsp, s21)
	log("p21 r:%d", value)
	countreads(value, 21, 24)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(17, &value), rsp, s21)
	log("p21 r:%d", value)
	countreads(value, 21, 2)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(501, &value), rsp, s21)
	log("p21 r:%d", value)
	countreads(value, 21, 32)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(333, &value), rsp, s21)
	log("p21 r:%d", value)
	countreads(value, 21, 21)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(112, &value), rsp, s21)
	log("p21 r:%d", value)
	countreads(value, 21, 7)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(329, &value), rsp, s21)
	log("p21 r:%d", value)
	countreads(value, 21, 21)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(47, &value), rsp, s21)
	log("p21 r:%d", value)
	countreads(value, 21, 3)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(462, &value), rsp, s21)
	log("p21 r:%d", value)
	countreads(value, 21, 29)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(421, &value), rsp, s21)
	log("p21 r:%d", value)
	countreads(value, 21, 27)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(14, &value), rsp, s21)
	log("p21 r:%d", value)
	countreads(value, 21, 1)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(81, &value), rsp, s21)
	log("p21 r:%d", value)
	countreads(value, 21, 6)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(171, &value), rsp, s21)
	log("p21 r:%d", value)
	countreads(value, 21, 11)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(270, &value), rsp, s21)
	log("p21 r:%d", value)
	countreads(value, 21, 17)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(50, &value), rsp, s21)
	log("p21 r:%d", value)
	countreads(value, 21, 4)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(383, &value), rsp, s21)
	log("p21 r:%d", value)
	countreads(value, 21, 24)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(190, &value), rsp, s21)
	log("p21 r:%d", value)
	countreads(value, 21, 12)
	l.Unlock()
}

func p22() {
	targets1 := make([]string, 2)
	targets1[1] = uri[s24]
	targets1[0] = uri[s21]
	targets0 := make([]string, 0)

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(27, &value), rsp, s22)
	log("p22 r:%d", value)
	countreads(value, 22, 2)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(127, &value), rsp, s22)
	log("p22 r:%d", value)
	countreads(value, 22, 8)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(410, &value), rsp, s22)
	log("p22 r:%d", value)
	countreads(value, 22, 26)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(241, &value), rsp, s22)
	log("p22 r:%d", value)
	countreads(value, 22, 16)
	l.Unlock()

	delay()
	l.Lock()
	value = 9
	Put(CreateTuple(136, value), rsp, targets0)
	log("p22 w:%d", value)
	countwrites(value, 22, 9)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(54, &value), rsp, s22)
	log("p22 r:%d", value)
	countreads(value, 22, 4)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(271, &value), rsp, s22)
	log("p22 r:%d", value)
	countreads(value, 22, 17)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(448, &value), rsp, s22)
	log("p22 r:%d", value)
	countreads(value, 22, 28)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(79, &value), rsp, s22)
	log("p22 r:%d", value)
	countreads(value, 22, 5)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(71, &value), rsp, s22)
	log("p22 r:%d", value)
	countreads(value, 22, 5)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(187, &value), rsp, s22)
	log("p22 r:%d", value)
	countreads(value, 22, 12)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(72, &value), rsp, s22)
	log("p22 r:%d", value)
	countreads(value, 22, 5)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(391, &value), rsp, s22)
	log("p22 r:%d", value)
	countreads(value, 22, 25)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(64, &value), rsp, s22)
	log("p22 r:%d", value)
	countreads(value, 22, 4)
	l.Unlock()

	delay()
	l.Lock()
	value = 11
	Put(CreateTuple(171, value), rsp, targets1)
	log("p22 w:%d", value)
	countwrites(value, 22, 11)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(174, &value), rsp, s22)
	log("p22 r:%d", value)
	countreads(value, 22, 11)
	l.Unlock()
}

func p23() {
	targets0 := make([]string, 0)
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(119, &value), rsp, s23)
	log("p23 r:%d", value)
	countreads(value, 23, 8)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(425, &value), rsp, s23)
	log("p23 r:%d", value)
	countreads(value, 23, 27)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(331, &value), rsp, s23)
	log("p23 r:%d", value)
	countreads(value, 23, 21)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(73, &value), rsp, s23)
	log("p23 r:%d", value)
	countreads(value, 23, 5)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(382, &value), rsp, s23)
	log("p23 r:%d", value)
	countreads(value, 23, 24)
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
	value = 26
	Put(CreateTuple(404, value), rsp, targets0)
	log("p23 w:%d", value)
	countwrites(value, 23, 26)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(326, &value), rsp, s23)
	log("p23 r:%d", value)
	countreads(value, 23, 21)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(217, &value), rsp, s23)
	log("p23 r:%d", value)
	countreads(value, 23, 14)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(188, &value), rsp, s23)
	log("p23 r:%d", value)
	countreads(value, 23, 12)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(365, &value), rsp, s23)
	log("p23 r:%d", value)
	countreads(value, 23, 23)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(310, &value), rsp, s23)
	log("p23 r:%d", value)
	countreads(value, 23, 20)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(412, &value), rsp, s23)
	log("p23 r:%d", value)
	countreads(value, 23, 26)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(481, &value), rsp, s23)
	log("p23 r:%d", value)
	countreads(value, 23, 31)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(62, &value), rsp, s23)
	log("p23 r:%d", value)
	countreads(value, 23, 4)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(282, &value), rsp, s23)
	log("p23 r:%d", value)
	countreads(value, 23, 18)
	l.Unlock()
}

func p24() {
	targets1 := make([]string, 1)
	targets1[0] = uri[s12]
	targets0 := make([]string, 1)
	targets0[0] = uri[s5]

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(319, &value), rsp, s24)
	log("p24 r:%d", value)
	countreads(value, 24, 20)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(281, &value), rsp, s24)
	log("p24 r:%d", value)
	countreads(value, 24, 18)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(254, &value), rsp, s24)
	log("p24 r:%d", value)
	countreads(value, 24, 16)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(13, &value), rsp, s24)
	log("p24 r:%d", value)
	countreads(value, 24, 1)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(384, &value), rsp, s24)
	log("p24 r:%d", value)
	countreads(value, 24, 24)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(342, &value), rsp, s24)
	log("p24 r:%d", value)
	countreads(value, 24, 22)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(509, &value), rsp, s24)
	log("p24 r:%d", value)
	countreads(value, 24, 32)
	l.Unlock()

	delay()
	l.Lock()
	value = 7
	Put(CreateTuple(99, value), rsp, targets0)
	log("p24 w:%d", value)
	countwrites(value, 24, 7)
	l.Unlock()

	delay()
	l.Lock()
	value = 10
	Put(CreateTuple(159, value), rsp, targets1)
	log("p24 w:%d", value)
	countwrites(value, 24, 10)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(454, &value), rsp, s24)
	log("p24 r:%d", value)
	countreads(value, 24, 29)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(466, &value), rsp, s24)
	log("p24 r:%d", value)
	countreads(value, 24, 30)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(145, &value), rsp, s24)
	log("p24 r:%d", value)
	countreads(value, 24, 10)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(171, &value), rsp, s24)
	log("p24 r:%d", value)
	countreads(value, 24, 11)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(377, &value), rsp, s24)
	log("p24 r:%d", value)
	countreads(value, 24, 24)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(35, &value), rsp, s24)
	log("p24 r:%d", value)
	countreads(value, 24, 3)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(272, &value), rsp, s24)
	log("p24 r:%d", value)
	countreads(value, 24, 17)
	l.Unlock()
}

func p25() {
	targets0 := make([]string, 2)
	targets0[1] = uri[s8]
	targets0[0] = uri[s5]

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(380, &value), rsp, s25)
	log("p25 r:%d", value)
	countreads(value, 25, 24)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(356, &value), rsp, s25)
	log("p25 r:%d", value)
	countreads(value, 25, 23)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(406, &value), rsp, s25)
	log("p25 r:%d", value)
	countreads(value, 25, 26)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(361, &value), rsp, s25)
	log("p25 r:%d", value)
	countreads(value, 25, 23)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(26, &value), rsp, s25)
	log("p25 r:%d", value)
	countreads(value, 25, 2)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(329, &value), rsp, s25)
	log("p25 r:%d", value)
	countreads(value, 25, 21)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(266, &value), rsp, s25)
	log("p25 r:%d", value)
	countreads(value, 25, 17)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(17, &value), rsp, s25)
	log("p25 r:%d", value)
	countreads(value, 25, 2)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(345, &value), rsp, s25)
	log("p25 r:%d", value)
	countreads(value, 25, 22)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(479, &value), rsp, s25)
	log("p25 r:%d", value)
	countreads(value, 25, 30)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(411, &value), rsp, s25)
	log("p25 r:%d", value)
	countreads(value, 25, 26)
	l.Unlock()

	delay()
	l.Lock()
	value = 1
	Put(CreateTuple(9, value), rsp, targets0)
	log("p25 w:%d", value)
	countwrites(value, 25, 1)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(483, &value), rsp, s25)
	log("p25 r:%d", value)
	countreads(value, 25, 31)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(144, &value), rsp, s25)
	log("p25 r:%d", value)
	countreads(value, 25, 9)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(160, &value), rsp, s25)
	log("p25 r:%d", value)
	countreads(value, 25, 10)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(191, &value), rsp, s25)
	log("p25 r:%d", value)
	countreads(value, 25, 12)
	l.Unlock()
}

func p26() {
	targets1 := make([]string, 1)
	targets1[0] = uri[s28]
	targets0 := make([]string, 0)

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(155, &value), rsp, s26)
	log("p26 r:%d", value)
	countreads(value, 26, 10)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(146, &value), rsp, s26)
	log("p26 r:%d", value)
	countreads(value, 26, 10)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(426, &value), rsp, s26)
	log("p26 r:%d", value)
	countreads(value, 26, 27)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(306, &value), rsp, s26)
	log("p26 r:%d", value)
	countreads(value, 26, 20)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(125, &value), rsp, s26)
	log("p26 r:%d", value)
	countreads(value, 26, 8)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(495, &value), rsp, s26)
	log("p26 r:%d", value)
	countreads(value, 26, 31)
	l.Unlock()

	delay()
	l.Lock()
	value = 1
	Put(CreateTuple(3, value), rsp, targets0)
	log("p26 w:%d", value)
	countwrites(value, 26, 1)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(469, &value), rsp, s26)
	log("p26 r:%d", value)
	countreads(value, 26, 30)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(367, &value), rsp, s26)
	log("p26 r:%d", value)
	countreads(value, 26, 23)
	l.Unlock()

	delay()
	l.Lock()
	value = 11
	Put(CreateTuple(162, value), rsp, targets1)
	log("p26 w:%d", value)
	countwrites(value, 26, 11)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(191, &value), rsp, s26)
	log("p26 r:%d", value)
	countreads(value, 26, 12)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(29, &value), rsp, s26)
	log("p26 r:%d", value)
	countreads(value, 26, 2)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(79, &value), rsp, s26)
	log("p26 r:%d", value)
	countreads(value, 26, 5)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(496, &value), rsp, s26)
	log("p26 r:%d", value)
	countreads(value, 26, 31)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(435, &value), rsp, s26)
	log("p26 r:%d", value)
	countreads(value, 26, 28)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(63, &value), rsp, s26)
	log("p26 r:%d", value)
	countreads(value, 26, 4)
	l.Unlock()
}

func p27() {
	targets2 := make([]string, 2)
	targets2[1] = uri[s27]
	targets2[0] = uri[s3]
	targets1 := make([]string, 1)
	targets1[0] = uri[s16]
	targets0 := make([]string, 0)

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 30
	Put(CreateTuple(467, value), rsp, targets0)
	log("p27 w:%d", value)
	countwrites(value, 27, 30)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(164, &value), rsp, s27)
	log("p27 r:%d", value)
	countreads(value, 27, 11)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(27, &value), rsp, s27)
	log("p27 r:%d", value)
	countreads(value, 27, 2)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(253, &value), rsp, s27)
	log("p27 r:%d", value)
	countreads(value, 27, 16)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(38, &value), rsp, s27)
	log("p27 r:%d", value)
	countreads(value, 27, 3)
	l.Unlock()

	delay()
	l.Lock()
	value = 28
	Put(CreateTuple(440, value), rsp, targets1)
	log("p27 w:%d", value)
	countwrites(value, 27, 28)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(176, &value), rsp, s27)
	log("p27 r:%d", value)
	countreads(value, 27, 11)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(61, &value), rsp, s27)
	log("p27 r:%d", value)
	countreads(value, 27, 4)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(246, &value), rsp, s27)
	log("p27 r:%d", value)
	countreads(value, 27, 16)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(407, &value), rsp, s27)
	log("p27 r:%d", value)
	countreads(value, 27, 26)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(96, &value), rsp, s27)
	log("p27 r:%d", value)
	countreads(value, 27, 6)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(143, &value), rsp, s27)
	log("p27 r:%d", value)
	countreads(value, 27, 9)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(60, &value), rsp, s27)
	log("p27 r:%d", value)
	countreads(value, 27, 4)
	l.Unlock()

	delay()
	l.Lock()
	value = 16
	Put(CreateTuple(246, value), rsp, targets2)
	log("p27 w:%d", value)
	countwrites(value, 27, 16)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(82, &value), rsp, s27)
	log("p27 r:%d", value)
	countreads(value, 27, 6)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(398, &value), rsp, s27)
	log("p27 r:%d", value)
	countreads(value, 27, 25)
	l.Unlock()
}

func p28() {
	targets0 := make([]string, 1)
	targets0[0] = uri[s21]

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(87, &value), rsp, s28)
	log("p28 r:%d", value)
	countreads(value, 28, 6)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(324, &value), rsp, s28)
	log("p28 r:%d", value)
	countreads(value, 28, 21)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(402, &value), rsp, s28)
	log("p28 r:%d", value)
	countreads(value, 28, 26)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(162, &value), rsp, s28)
	log("p28 r:%d", value)
	countreads(value, 28, 11)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(447, &value), rsp, s28)
	log("p28 r:%d", value)
	countreads(value, 28, 28)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(342, &value), rsp, s28)
	log("p28 r:%d", value)
	countreads(value, 28, 22)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(202, &value), rsp, s28)
	log("p28 r:%d", value)
	countreads(value, 28, 13)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(372, &value), rsp, s28)
	log("p28 r:%d", value)
	countreads(value, 28, 24)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(266, &value), rsp, s28)
	log("p28 r:%d", value)
	countreads(value, 28, 17)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(201, &value), rsp, s28)
	log("p28 r:%d", value)
	countreads(value, 28, 13)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(387, &value), rsp, s28)
	log("p28 r:%d", value)
	countreads(value, 28, 25)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(444, &value), rsp, s28)
	log("p28 r:%d", value)
	countreads(value, 28, 28)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(466, &value), rsp, s28)
	log("p28 r:%d", value)
	countreads(value, 28, 30)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(88, &value), rsp, s28)
	log("p28 r:%d", value)
	countreads(value, 28, 6)
	l.Unlock()

	delay()
	l.Lock()
	value = 29
	Put(CreateTuple(462, value), rsp, targets0)
	log("p28 w:%d", value)
	countwrites(value, 28, 29)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(182, &value), rsp, s28)
	log("p28 r:%d", value)
	countreads(value, 28, 12)
	l.Unlock()
}

func p29() {
	targets0 := make([]string, 2)
	targets0[1] = uri[s18]
	targets0[0] = uri[s7]

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 15
	Put(CreateTuple(233, value), rsp, targets0)
	log("p29 w:%d", value)
	countwrites(value, 29, 15)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(408, &value), rsp, s29)
	log("p29 r:%d", value)
	countreads(value, 29, 26)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(22, &value), rsp, s29)
	log("p29 r:%d", value)
	countreads(value, 29, 2)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(466, &value), rsp, s29)
	log("p29 r:%d", value)
	countreads(value, 29, 30)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(300, &value), rsp, s29)
	log("p29 r:%d", value)
	countreads(value, 29, 19)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(20, &value), rsp, s29)
	log("p29 r:%d", value)
	countreads(value, 29, 2)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(30, &value), rsp, s29)
	log("p29 r:%d", value)
	countreads(value, 29, 2)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(332, &value), rsp, s29)
	log("p29 r:%d", value)
	countreads(value, 29, 21)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(41, &value), rsp, s29)
	log("p29 r:%d", value)
	countreads(value, 29, 3)
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
	value = -1
	QueryP(CreateTuple(472, &value), rsp, s29)
	log("p29 r:%d", value)
	countreads(value, 29, 30)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(368, &value), rsp, s29)
	log("p29 r:%d", value)
	countreads(value, 29, 23)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(151, &value), rsp, s29)
	log("p29 r:%d", value)
	countreads(value, 29, 10)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(134, &value), rsp, s29)
	log("p29 r:%d", value)
	countreads(value, 29, 9)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(163, &value), rsp, s29)
	log("p29 r:%d", value)
	countreads(value, 29, 11)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(206, &value), rsp, s29)
	log("p29 r:%d", value)
	countreads(value, 29, 13)
	l.Unlock()
}

func p30() {
	targets0 := make([]string, 1)
	targets0[0] = uri[s31]

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(144, &value), rsp, s30)
	log("p30 r:%d", value)
	countreads(value, 30, 9)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(419, &value), rsp, s30)
	log("p30 r:%d", value)
	countreads(value, 30, 27)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(23, &value), rsp, s30)
	log("p30 r:%d", value)
	countreads(value, 30, 2)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(125, &value), rsp, s30)
	log("p30 r:%d", value)
	countreads(value, 30, 8)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(116, &value), rsp, s30)
	log("p30 r:%d", value)
	countreads(value, 30, 8)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(277, &value), rsp, s30)
	log("p30 r:%d", value)
	countreads(value, 30, 18)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(479, &value), rsp, s30)
	log("p30 r:%d", value)
	countreads(value, 30, 30)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(393, &value), rsp, s30)
	log("p30 r:%d", value)
	countreads(value, 30, 25)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(188, &value), rsp, s30)
	log("p30 r:%d", value)
	countreads(value, 30, 12)
	l.Unlock()

	delay()
	l.Lock()
	value = 14
	Put(CreateTuple(213, value), rsp, targets0)
	log("p30 w:%d", value)
	countwrites(value, 30, 14)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(287, &value), rsp, s30)
	log("p30 r:%d", value)
	countreads(value, 30, 18)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(58, &value), rsp, s30)
	log("p30 r:%d", value)
	countreads(value, 30, 4)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(20, &value), rsp, s30)
	log("p30 r:%d", value)
	countreads(value, 30, 2)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(273, &value), rsp, s30)
	log("p30 r:%d", value)
	countreads(value, 30, 18)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(150, &value), rsp, s30)
	log("p30 r:%d", value)
	countreads(value, 30, 10)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(307, &value), rsp, s30)
	log("p30 r:%d", value)
	countreads(value, 30, 20)
	l.Unlock()
}

func p31() {
	targets2 := make([]string, 0)
	targets1 := make([]string, 1)
	targets1[0] = uri[s29]
	targets0 := make([]string, 1)
	targets0[0] = uri[s5]

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(229, &value), rsp, s31)
	log("p31 r:%d", value)
	countreads(value, 31, 15)
	l.Unlock()

	delay()
	l.Lock()
	value = 1
	Put(CreateTuple(10, value), rsp, targets0)
	log("p31 w:%d", value)
	countwrites(value, 31, 1)
	l.Unlock()

	delay()
	l.Lock()
	value = 21
	Put(CreateTuple(332, value), rsp, targets1)
	log("p31 w:%d", value)
	countwrites(value, 31, 21)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(244, &value), rsp, s31)
	log("p31 r:%d", value)
	countreads(value, 31, 16)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(213, &value), rsp, s31)
	log("p31 r:%d", value)
	countreads(value, 31, 14)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(154, &value), rsp, s31)
	log("p31 r:%d", value)
	countreads(value, 31, 10)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(172, &value), rsp, s31)
	log("p31 r:%d", value)
	countreads(value, 31, 11)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(146, &value), rsp, s31)
	log("p31 r:%d", value)
	countreads(value, 31, 10)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(287, &value), rsp, s31)
	log("p31 r:%d", value)
	countreads(value, 31, 18)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(443, &value), rsp, s31)
	log("p31 r:%d", value)
	countreads(value, 31, 28)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(248, &value), rsp, s31)
	log("p31 r:%d", value)
	countreads(value, 31, 16)
	l.Unlock()

	delay()
	l.Lock()
	value = 14
	Put(CreateTuple(218, value), rsp, targets2)
	log("p31 w:%d", value)
	countwrites(value, 31, 14)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(210, &value), rsp, s31)
	log("p31 r:%d", value)
	countreads(value, 31, 14)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(446, &value), rsp, s31)
	log("p31 r:%d", value)
	countreads(value, 31, 28)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(137, &value), rsp, s31)
	log("p31 r:%d", value)
	countreads(value, 31, 9)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(445, &value), rsp, s31)
	log("p31 r:%d", value)
	countreads(value, 31, 28)
	l.Unlock()
}

func p32() {
	targets3 := make([]string, 1)
	targets3[0] = uri[s30]
	targets2 := make([]string, 0)
	targets1 := make([]string, 3)
	targets1[2] = uri[s19]
	targets1[1] = uri[s14]
	targets1[0] = uri[s4]
	targets0 := make([]string, 0)

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(444, &value), rsp, s32)
	log("p32 r:%d", value)
	countreads(value, 32, 28)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(374, &value), rsp, s32)
	log("p32 r:%d", value)
	countreads(value, 32, 24)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(309, &value), rsp, s32)
	log("p32 r:%d", value)
	countreads(value, 32, 20)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(219, &value), rsp, s32)
	log("p32 r:%d", value)
	countreads(value, 32, 14)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(61, &value), rsp, s32)
	log("p32 r:%d", value)
	countreads(value, 32, 4)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(176, &value), rsp, s32)
	log("p32 r:%d", value)
	countreads(value, 32, 11)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(65, &value), rsp, s32)
	log("p32 r:%d", value)
	countreads(value, 32, 5)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(134, &value), rsp, s32)
	log("p32 r:%d", value)
	countreads(value, 32, 9)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(289, &value), rsp, s32)
	log("p32 r:%d", value)
	countreads(value, 32, 19)
	l.Unlock()

	delay()
	l.Lock()
	value = 1
	Put(CreateTuple(8, value), rsp, targets0)
	log("p32 w:%d", value)
	countwrites(value, 32, 1)
	l.Unlock()

	delay()
	l.Lock()
	value = 2
	Put(CreateTuple(28, value), rsp, targets1)
	log("p32 w:%d", value)
	countwrites(value, 32, 2)
	l.Unlock()

	delay()
	l.Lock()
	value = 21
	Put(CreateTuple(330, value), rsp, targets2)
	log("p32 w:%d", value)
	countwrites(value, 32, 21)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(158, &value), rsp, s32)
	log("p32 r:%d", value)
	countreads(value, 32, 10)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(441, &value), rsp, s32)
	log("p32 r:%d", value)
	countreads(value, 32, 28)
	l.Unlock()

	delay()
	l.Lock()
	value = 20
	Put(CreateTuple(307, value), rsp, targets3)
	log("p32 w:%d", value)
	countwrites(value, 32, 20)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(115, &value), rsp, s32)
	log("p32 r:%d", value)
	countreads(value, 32, 8)
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
