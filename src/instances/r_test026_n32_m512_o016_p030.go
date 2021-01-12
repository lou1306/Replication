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
	targets3 := make([]string, 1)
	targets3[0] = uri[s8]
	targets2 := make([]string, 0)
	targets1 := make([]string, 0)
	targets0 := make([]string, 0)

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(383, &value), rsp, s1)
	log("p1 r:%d", value)
	countreads(value, 1, 24)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(137, &value), rsp, s1)
	log("p1 r:%d", value)
	countreads(value, 1, 9)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(66, &value), rsp, s1)
	log("p1 r:%d", value)
	countreads(value, 1, 5)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(401, &value), rsp, s1)
	log("p1 r:%d", value)
	countreads(value, 1, 26)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(36, &value), rsp, s1)
	log("p1 r:%d", value)
	countreads(value, 1, 3)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(173, &value), rsp, s1)
	log("p1 r:%d", value)
	countreads(value, 1, 11)
	l.Unlock()

	delay()
	l.Lock()
	value = 23
	Put(CreateTuple(357, value), rsp, targets0)
	log("p1 w:%d", value)
	countwrites(value, 1, 23)
	l.Unlock()

	delay()
	l.Lock()
	value = 1
	Put(CreateTuple(4, value), rsp, targets1)
	log("p1 w:%d", value)
	countwrites(value, 1, 1)
	l.Unlock()

	delay()
	l.Lock()
	value = 10
	Put(CreateTuple(148, value), rsp, targets2)
	log("p1 w:%d", value)
	countwrites(value, 1, 10)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(282, &value), rsp, s1)
	log("p1 r:%d", value)
	countreads(value, 1, 18)
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
	QueryP(CreateTuple(469, &value), rsp, s1)
	log("p1 r:%d", value)
	countreads(value, 1, 30)
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
	QueryP(CreateTuple(447, &value), rsp, s1)
	log("p1 r:%d", value)
	countreads(value, 1, 28)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(96, &value), rsp, s1)
	log("p1 r:%d", value)
	countreads(value, 1, 6)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(133, &value), rsp, s1)
	log("p1 r:%d", value)
	countreads(value, 1, 9)
	l.Unlock()
}

func p2() {
	targets5 := make([]string, 4)
	targets5[3] = uri[s28]
	targets5[2] = uri[s21]
	targets5[1] = uri[s20]
	targets5[0] = uri[s5]
	targets4 := make([]string, 2)
	targets4[1] = uri[s18]
	targets4[0] = uri[s10]
	targets3 := make([]string, 0)
	targets2 := make([]string, 1)
	targets2[0] = uri[s3]
	targets1 := make([]string, 0)
	targets0 := make([]string, 1)
	targets0[0] = uri[s19]

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 32
	Put(CreateTuple(509, value), rsp, targets0)
	log("p2 w:%d", value)
	countwrites(value, 2, 32)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(66, &value), rsp, s2)
	log("p2 r:%d", value)
	countreads(value, 2, 5)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(362, &value), rsp, s2)
	log("p2 r:%d", value)
	countreads(value, 2, 23)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(360, &value), rsp, s2)
	log("p2 r:%d", value)
	countreads(value, 2, 23)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(472, &value), rsp, s2)
	log("p2 r:%d", value)
	countreads(value, 2, 30)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(432, &value), rsp, s2)
	log("p2 r:%d", value)
	countreads(value, 2, 27)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(86, &value), rsp, s2)
	log("p2 r:%d", value)
	countreads(value, 2, 6)
	l.Unlock()

	delay()
	l.Lock()
	value = 6
	Put(CreateTuple(92, value), rsp, targets1)
	log("p2 w:%d", value)
	countwrites(value, 2, 6)
	l.Unlock()

	delay()
	l.Lock()
	value = 1
	Put(CreateTuple(15, value), rsp, targets2)
	log("p2 w:%d", value)
	countwrites(value, 2, 1)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(424, &value), rsp, s2)
	log("p2 r:%d", value)
	countreads(value, 2, 27)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(50, &value), rsp, s2)
	log("p2 r:%d", value)
	countreads(value, 2, 4)
	l.Unlock()

	delay()
	l.Lock()
	value = 30
	Put(CreateTuple(480, value), rsp, targets3)
	log("p2 w:%d", value)
	countwrites(value, 2, 30)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(463, &value), rsp, s2)
	log("p2 r:%d", value)
	countreads(value, 2, 29)
	l.Unlock()

	delay()
	l.Lock()
	value = 24
	Put(CreateTuple(375, value), rsp, targets4)
	log("p2 w:%d", value)
	countwrites(value, 2, 24)
	l.Unlock()

	delay()
	l.Lock()
	value = 10
	Put(CreateTuple(155, value), rsp, targets5)
	log("p2 w:%d", value)
	countwrites(value, 2, 10)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(151, &value), rsp, s2)
	log("p2 r:%d", value)
	countreads(value, 2, 10)
	l.Unlock()
}

func p3() {
	targets3 := make([]string, 0)
	targets2 := make([]string, 0)
	targets1 := make([]string, 1)
	targets1[0] = uri[s31]
	targets0 := make([]string, 0)

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 1
	Put(CreateTuple(6, value), rsp, targets0)
	log("p3 w:%d", value)
	countwrites(value, 3, 1)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(327, &value), rsp, s3)
	log("p3 r:%d", value)
	countreads(value, 3, 21)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(380, &value), rsp, s3)
	log("p3 r:%d", value)
	countreads(value, 3, 24)
	l.Unlock()

	delay()
	l.Lock()
	value = 17
	Put(CreateTuple(267, value), rsp, targets1)
	log("p3 w:%d", value)
	countwrites(value, 3, 17)
	l.Unlock()

	delay()
	l.Lock()
	value = 19
	Put(CreateTuple(298, value), rsp, targets2)
	log("p3 w:%d", value)
	countwrites(value, 3, 19)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(187, &value), rsp, s3)
	log("p3 r:%d", value)
	countreads(value, 3, 12)
	l.Unlock()

	delay()
	l.Lock()
	value = 12
	Put(CreateTuple(180, value), rsp, targets3)
	log("p3 w:%d", value)
	countwrites(value, 3, 12)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(502, &value), rsp, s3)
	log("p3 r:%d", value)
	countreads(value, 3, 32)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(22, &value), rsp, s3)
	log("p3 r:%d", value)
	countreads(value, 3, 2)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(264, &value), rsp, s3)
	log("p3 r:%d", value)
	countreads(value, 3, 17)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(380, &value), rsp, s3)
	log("p3 r:%d", value)
	countreads(value, 3, 24)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(198, &value), rsp, s3)
	log("p3 r:%d", value)
	countreads(value, 3, 13)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(30, &value), rsp, s3)
	log("p3 r:%d", value)
	countreads(value, 3, 2)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(323, &value), rsp, s3)
	log("p3 r:%d", value)
	countreads(value, 3, 21)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(15, &value), rsp, s3)
	log("p3 r:%d", value)
	countreads(value, 3, 1)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(197, &value), rsp, s3)
	log("p3 r:%d", value)
	countreads(value, 3, 13)
	l.Unlock()
}

func p4() {
	targets6 := make([]string, 1)
	targets6[0] = uri[s13]
	targets5 := make([]string, 2)
	targets5[1] = uri[s31]
	targets5[0] = uri[s12]
	targets4 := make([]string, 0)
	targets3 := make([]string, 1)
	targets3[0] = uri[s28]
	targets2 := make([]string, 1)
	targets2[0] = uri[s5]
	targets1 := make([]string, 0)
	targets0 := make([]string, 0)

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 12
	Put(CreateTuple(182, value), rsp, targets0)
	log("p4 w:%d", value)
	countwrites(value, 4, 12)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(80, &value), rsp, s4)
	log("p4 r:%d", value)
	countreads(value, 4, 5)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(149, &value), rsp, s4)
	log("p4 r:%d", value)
	countreads(value, 4, 10)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(459, &value), rsp, s4)
	log("p4 r:%d", value)
	countreads(value, 4, 29)
	l.Unlock()

	delay()
	l.Lock()
	value = 14
	Put(CreateTuple(210, value), rsp, targets1)
	log("p4 w:%d", value)
	countwrites(value, 4, 14)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(325, &value), rsp, s4)
	log("p4 r:%d", value)
	countreads(value, 4, 21)
	l.Unlock()

	delay()
	l.Lock()
	value = 5
	Put(CreateTuple(76, value), rsp, targets2)
	log("p4 w:%d", value)
	countwrites(value, 4, 5)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(455, &value), rsp, s4)
	log("p4 r:%d", value)
	countreads(value, 4, 29)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(328, &value), rsp, s4)
	log("p4 r:%d", value)
	countreads(value, 4, 21)
	l.Unlock()

	delay()
	l.Lock()
	value = 13
	Put(CreateTuple(206, value), rsp, targets3)
	log("p4 w:%d", value)
	countwrites(value, 4, 13)
	l.Unlock()

	delay()
	l.Lock()
	value = 11
	Put(CreateTuple(169, value), rsp, targets4)
	log("p4 w:%d", value)
	countwrites(value, 4, 11)
	l.Unlock()

	delay()
	l.Lock()
	value = 6
	Put(CreateTuple(87, value), rsp, targets5)
	log("p4 w:%d", value)
	countwrites(value, 4, 6)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(398, &value), rsp, s4)
	log("p4 r:%d", value)
	countreads(value, 4, 25)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(284, &value), rsp, s4)
	log("p4 r:%d", value)
	countreads(value, 4, 18)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(405, &value), rsp, s4)
	log("p4 r:%d", value)
	countreads(value, 4, 26)
	l.Unlock()

	delay()
	l.Lock()
	value = 6
	Put(CreateTuple(81, value), rsp, targets6)
	log("p4 w:%d", value)
	countwrites(value, 4, 6)
	l.Unlock()
}

func p5() {
	targets3 := make([]string, 1)
	targets3[0] = uri[s27]
	targets2 := make([]string, 3)
	targets2[2] = uri[s25]
	targets2[1] = uri[s21]
	targets2[0] = uri[s17]
	targets1 := make([]string, 1)
	targets1[0] = uri[s4]
	targets0 := make([]string, 1)
	targets0[0] = uri[s1]

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(316, &value), rsp, s5)
	log("p5 r:%d", value)
	countreads(value, 5, 20)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(491, &value), rsp, s5)
	log("p5 r:%d", value)
	countreads(value, 5, 31)
	l.Unlock()

	delay()
	l.Lock()
	value = 24
	Put(CreateTuple(383, value), rsp, targets0)
	log("p5 w:%d", value)
	countwrites(value, 5, 24)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(258, &value), rsp, s5)
	log("p5 r:%d", value)
	countreads(value, 5, 17)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(499, &value), rsp, s5)
	log("p5 r:%d", value)
	countreads(value, 5, 32)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(297, &value), rsp, s5)
	log("p5 r:%d", value)
	countreads(value, 5, 19)
	l.Unlock()

	delay()
	l.Lock()
	value = 26
	Put(CreateTuple(405, value), rsp, targets1)
	log("p5 w:%d", value)
	countwrites(value, 5, 26)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(76, &value), rsp, s5)
	log("p5 r:%d", value)
	countreads(value, 5, 5)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(155, &value), rsp, s5)
	log("p5 r:%d", value)
	countreads(value, 5, 10)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(365, &value), rsp, s5)
	log("p5 r:%d", value)
	countreads(value, 5, 23)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(90, &value), rsp, s5)
	log("p5 r:%d", value)
	countreads(value, 5, 6)
	l.Unlock()

	delay()
	l.Lock()
	value = 27
	Put(CreateTuple(429, value), rsp, targets2)
	log("p5 w:%d", value)
	countwrites(value, 5, 27)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(331, &value), rsp, s5)
	log("p5 r:%d", value)
	countreads(value, 5, 21)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(30, &value), rsp, s5)
	log("p5 r:%d", value)
	countreads(value, 5, 2)
	l.Unlock()

	delay()
	l.Lock()
	value = 2
	Put(CreateTuple(26, value), rsp, targets3)
	log("p5 w:%d", value)
	countwrites(value, 5, 2)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(215, &value), rsp, s5)
	log("p5 r:%d", value)
	countreads(value, 5, 14)
	l.Unlock()
}

func p6() {
	targets2 := make([]string, 1)
	targets2[0] = uri[s9]
	targets1 := make([]string, 0)
	targets0 := make([]string, 0)

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 30
	Put(CreateTuple(477, value), rsp, targets0)
	log("p6 w:%d", value)
	countwrites(value, 6, 30)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(139, &value), rsp, s6)
	log("p6 r:%d", value)
	countreads(value, 6, 9)
	l.Unlock()

	delay()
	l.Lock()
	value = 24
	Put(CreateTuple(372, value), rsp, targets1)
	log("p6 w:%d", value)
	countwrites(value, 6, 24)
	l.Unlock()

	delay()
	l.Lock()
	value = 22
	Put(CreateTuple(343, value), rsp, targets2)
	log("p6 w:%d", value)
	countwrites(value, 6, 22)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(441, &value), rsp, s6)
	log("p6 r:%d", value)
	countreads(value, 6, 28)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(140, &value), rsp, s6)
	log("p6 r:%d", value)
	countreads(value, 6, 9)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(454, &value), rsp, s6)
	log("p6 r:%d", value)
	countreads(value, 6, 29)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(472, &value), rsp, s6)
	log("p6 r:%d", value)
	countreads(value, 6, 30)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(395, &value), rsp, s6)
	log("p6 r:%d", value)
	countreads(value, 6, 25)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(505, &value), rsp, s6)
	log("p6 r:%d", value)
	countreads(value, 6, 32)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(455, &value), rsp, s6)
	log("p6 r:%d", value)
	countreads(value, 6, 29)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(352, &value), rsp, s6)
	log("p6 r:%d", value)
	countreads(value, 6, 22)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(88, &value), rsp, s6)
	log("p6 r:%d", value)
	countreads(value, 6, 6)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(266, &value), rsp, s6)
	log("p6 r:%d", value)
	countreads(value, 6, 17)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(90, &value), rsp, s6)
	log("p6 r:%d", value)
	countreads(value, 6, 6)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(407, &value), rsp, s6)
	log("p6 r:%d", value)
	countreads(value, 6, 26)
	l.Unlock()
}

func p7() {
	targets6 := make([]string, 0)
	targets5 := make([]string, 1)
	targets5[0] = uri[s6]
	targets4 := make([]string, 0)
	targets3 := make([]string, 1)
	targets3[0] = uri[s4]
	targets2 := make([]string, 0)
	targets1 := make([]string, 2)
	targets1[1] = uri[s23]
	targets1[0] = uri[s6]
	targets0 := make([]string, 1)
	targets0[0] = uri[s14]

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 7
	Put(CreateTuple(112, value), rsp, targets0)
	log("p7 w:%d", value)
	countwrites(value, 7, 7)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(499, &value), rsp, s7)
	log("p7 r:%d", value)
	countreads(value, 7, 32)
	l.Unlock()

	delay()
	l.Lock()
	value = 22
	Put(CreateTuple(352, value), rsp, targets1)
	log("p7 w:%d", value)
	countwrites(value, 7, 22)
	l.Unlock()

	delay()
	l.Lock()
	value = 7
	Put(CreateTuple(100, value), rsp, targets2)
	log("p7 w:%d", value)
	countwrites(value, 7, 7)
	l.Unlock()

	delay()
	l.Lock()
	value = 26
	Put(CreateTuple(405, value), rsp, targets3)
	log("p7 w:%d", value)
	countwrites(value, 7, 26)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(501, &value), rsp, s7)
	log("p7 r:%d", value)
	countreads(value, 7, 32)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(426, &value), rsp, s7)
	log("p7 r:%d", value)
	countreads(value, 7, 27)
	l.Unlock()

	delay()
	l.Lock()
	value = 17
	Put(CreateTuple(271, value), rsp, targets4)
	log("p7 w:%d", value)
	countwrites(value, 7, 17)
	l.Unlock()

	delay()
	l.Lock()
	value = 32
	Put(CreateTuple(505, value), rsp, targets5)
	log("p7 w:%d", value)
	countwrites(value, 7, 32)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(90, &value), rsp, s7)
	log("p7 r:%d", value)
	countreads(value, 7, 6)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(40, &value), rsp, s7)
	log("p7 r:%d", value)
	countreads(value, 7, 3)
	l.Unlock()

	delay()
	l.Lock()
	value = 18
	Put(CreateTuple(277, value), rsp, targets6)
	log("p7 w:%d", value)
	countwrites(value, 7, 18)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(400, &value), rsp, s7)
	log("p7 r:%d", value)
	countreads(value, 7, 25)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(75, &value), rsp, s7)
	log("p7 r:%d", value)
	countreads(value, 7, 5)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(91, &value), rsp, s7)
	log("p7 r:%d", value)
	countreads(value, 7, 6)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(185, &value), rsp, s7)
	log("p7 r:%d", value)
	countreads(value, 7, 12)
	l.Unlock()
}

func p8() {
	targets5 := make([]string, 1)
	targets5[0] = uri[s22]
	targets4 := make([]string, 1)
	targets4[0] = uri[s9]
	targets3 := make([]string, 1)
	targets3[0] = uri[s11]
	targets2 := make([]string, 0)
	targets1 := make([]string, 1)
	targets1[0] = uri[s22]
	targets0 := make([]string, 0)

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(188, &value), rsp, s8)
	log("p8 r:%d", value)
	countreads(value, 8, 12)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(358, &value), rsp, s8)
	log("p8 r:%d", value)
	countreads(value, 8, 23)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(265, &value), rsp, s8)
	log("p8 r:%d", value)
	countreads(value, 8, 17)
	l.Unlock()

	delay()
	l.Lock()
	value = 22
	Put(CreateTuple(350, value), rsp, targets0)
	log("p8 w:%d", value)
	countwrites(value, 8, 22)
	l.Unlock()

	delay()
	l.Lock()
	value = 15
	Put(CreateTuple(234, value), rsp, targets1)
	log("p8 w:%d", value)
	countwrites(value, 8, 15)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(97, &value), rsp, s8)
	log("p8 r:%d", value)
	countreads(value, 8, 7)
	l.Unlock()

	delay()
	l.Lock()
	value = 10
	Put(CreateTuple(156, value), rsp, targets2)
	log("p8 w:%d", value)
	countwrites(value, 8, 10)
	l.Unlock()

	delay()
	l.Lock()
	value = 4
	Put(CreateTuple(60, value), rsp, targets3)
	log("p8 w:%d", value)
	countwrites(value, 8, 4)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(332, &value), rsp, s8)
	log("p8 r:%d", value)
	countreads(value, 8, 21)
	l.Unlock()

	delay()
	l.Lock()
	value = 1
	Put(CreateTuple(13, value), rsp, targets4)
	log("p8 w:%d", value)
	countwrites(value, 8, 1)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(415, &value), rsp, s8)
	log("p8 r:%d", value)
	countreads(value, 8, 26)
	l.Unlock()

	delay()
	l.Lock()
	value = 15
	Put(CreateTuple(234, value), rsp, targets5)
	log("p8 w:%d", value)
	countwrites(value, 8, 15)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(403, &value), rsp, s8)
	log("p8 r:%d", value)
	countreads(value, 8, 26)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(410, &value), rsp, s8)
	log("p8 r:%d", value)
	countreads(value, 8, 26)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(236, &value), rsp, s8)
	log("p8 r:%d", value)
	countreads(value, 8, 15)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(194, &value), rsp, s8)
	log("p8 r:%d", value)
	countreads(value, 8, 13)
	l.Unlock()
}

func p9() {
	targets4 := make([]string, 2)
	targets4[1] = uri[s25]
	targets4[0] = uri[s6]
	targets3 := make([]string, 0)
	targets2 := make([]string, 0)
	targets1 := make([]string, 3)
	targets1[2] = uri[s25]
	targets1[1] = uri[s21]
	targets1[0] = uri[s20]
	targets0 := make([]string, 0)

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(93, &value), rsp, s9)
	log("p9 r:%d", value)
	countreads(value, 9, 6)
	l.Unlock()

	delay()
	l.Lock()
	value = 32
	Put(CreateTuple(504, value), rsp, targets0)
	log("p9 w:%d", value)
	countwrites(value, 9, 32)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(343, &value), rsp, s9)
	log("p9 r:%d", value)
	countreads(value, 9, 22)
	l.Unlock()

	delay()
	l.Lock()
	value = 2
	Put(CreateTuple(31, value), rsp, targets1)
	log("p9 w:%d", value)
	countwrites(value, 9, 2)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(109, &value), rsp, s9)
	log("p9 r:%d", value)
	countreads(value, 9, 7)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(232, &value), rsp, s9)
	log("p9 r:%d", value)
	countreads(value, 9, 15)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(78, &value), rsp, s9)
	log("p9 r:%d", value)
	countreads(value, 9, 5)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(461, &value), rsp, s9)
	log("p9 r:%d", value)
	countreads(value, 9, 29)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(474, &value), rsp, s9)
	log("p9 r:%d", value)
	countreads(value, 9, 30)
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
	value = -1
	QueryP(CreateTuple(327, &value), rsp, s9)
	log("p9 r:%d", value)
	countreads(value, 9, 21)
	l.Unlock()

	delay()
	l.Lock()
	value = 11
	Put(CreateTuple(164, value), rsp, targets2)
	log("p9 w:%d", value)
	countwrites(value, 9, 11)
	l.Unlock()

	delay()
	l.Lock()
	value = 21
	Put(CreateTuple(329, value), rsp, targets3)
	log("p9 w:%d", value)
	countwrites(value, 9, 21)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(72, &value), rsp, s9)
	log("p9 r:%d", value)
	countreads(value, 9, 5)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(13, &value), rsp, s9)
	log("p9 r:%d", value)
	countreads(value, 9, 1)
	l.Unlock()

	delay()
	l.Lock()
	value = 26
	Put(CreateTuple(407, value), rsp, targets4)
	log("p9 w:%d", value)
	countwrites(value, 9, 26)
	l.Unlock()
}

func p10() {
	targets5 := make([]string, 0)
	targets4 := make([]string, 0)
	targets3 := make([]string, 1)
	targets3[0] = uri[s18]
	targets2 := make([]string, 0)
	targets1 := make([]string, 1)
	targets1[0] = uri[s22]
	targets0 := make([]string, 1)
	targets0[0] = uri[s5]

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 19
	Put(CreateTuple(297, value), rsp, targets0)
	log("p10 w:%d", value)
	countwrites(value, 10, 19)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(375, &value), rsp, s10)
	log("p10 r:%d", value)
	countreads(value, 10, 24)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(495, &value), rsp, s10)
	log("p10 r:%d", value)
	countreads(value, 10, 31)
	l.Unlock()

	delay()
	l.Lock()
	value = 20
	Put(CreateTuple(311, value), rsp, targets1)
	log("p10 w:%d", value)
	countwrites(value, 10, 20)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(175, &value), rsp, s10)
	log("p10 r:%d", value)
	countreads(value, 10, 11)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(455, &value), rsp, s10)
	log("p10 r:%d", value)
	countreads(value, 10, 29)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(232, &value), rsp, s10)
	log("p10 r:%d", value)
	countreads(value, 10, 15)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(389, &value), rsp, s10)
	log("p10 r:%d", value)
	countreads(value, 10, 25)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(497, &value), rsp, s10)
	log("p10 r:%d", value)
	countreads(value, 10, 32)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(323, &value), rsp, s10)
	log("p10 r:%d", value)
	countreads(value, 10, 21)
	l.Unlock()

	delay()
	l.Lock()
	value = 25
	Put(CreateTuple(396, value), rsp, targets2)
	log("p10 w:%d", value)
	countwrites(value, 10, 25)
	l.Unlock()

	delay()
	l.Lock()
	value = 13
	Put(CreateTuple(205, value), rsp, targets3)
	log("p10 w:%d", value)
	countwrites(value, 10, 13)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(422, &value), rsp, s10)
	log("p10 r:%d", value)
	countreads(value, 10, 27)
	l.Unlock()

	delay()
	l.Lock()
	value = 28
	Put(CreateTuple(445, value), rsp, targets4)
	log("p10 w:%d", value)
	countwrites(value, 10, 28)
	l.Unlock()

	delay()
	l.Lock()
	value = 22
	Put(CreateTuple(346, value), rsp, targets5)
	log("p10 w:%d", value)
	countwrites(value, 10, 22)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(167, &value), rsp, s10)
	log("p10 r:%d", value)
	countreads(value, 10, 11)
	l.Unlock()
}

func p11() {
	targets5 := make([]string, 1)
	targets5[0] = uri[s2]
	targets4 := make([]string, 1)
	targets4[0] = uri[s19]
	targets3 := make([]string, 3)
	targets3[2] = uri[s21]
	targets3[1] = uri[s16]
	targets3[0] = uri[s6]
	targets2 := make([]string, 1)
	targets2[0] = uri[s4]
	targets1 := make([]string, 1)
	targets1[0] = uri[s8]
	targets0 := make([]string, 0)

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 16
	Put(CreateTuple(250, value), rsp, targets0)
	log("p11 w:%d", value)
	countwrites(value, 11, 16)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(114, &value), rsp, s11)
	log("p11 r:%d", value)
	countreads(value, 11, 8)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(265, &value), rsp, s11)
	log("p11 r:%d", value)
	countreads(value, 11, 17)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(364, &value), rsp, s11)
	log("p11 r:%d", value)
	countreads(value, 11, 23)
	l.Unlock()

	delay()
	l.Lock()
	value = 26
	Put(CreateTuple(410, value), rsp, targets1)
	log("p11 w:%d", value)
	countwrites(value, 11, 26)
	l.Unlock()

	delay()
	l.Lock()
	value = 21
	Put(CreateTuple(325, value), rsp, targets2)
	log("p11 w:%d", value)
	countwrites(value, 11, 21)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(367, &value), rsp, s11)
	log("p11 r:%d", value)
	countreads(value, 11, 23)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(138, &value), rsp, s11)
	log("p11 r:%d", value)
	countreads(value, 11, 9)
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
	QueryP(CreateTuple(334, &value), rsp, s11)
	log("p11 r:%d", value)
	countreads(value, 11, 21)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(465, &value), rsp, s11)
	log("p11 r:%d", value)
	countreads(value, 11, 30)
	l.Unlock()

	delay()
	l.Lock()
	value = 25
	Put(CreateTuple(395, value), rsp, targets3)
	log("p11 w:%d", value)
	countwrites(value, 11, 25)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(173, &value), rsp, s11)
	log("p11 r:%d", value)
	countreads(value, 11, 11)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(60, &value), rsp, s11)
	log("p11 r:%d", value)
	countreads(value, 11, 4)
	l.Unlock()

	delay()
	l.Lock()
	value = 27
	Put(CreateTuple(417, value), rsp, targets4)
	log("p11 w:%d", value)
	countwrites(value, 11, 27)
	l.Unlock()

	delay()
	l.Lock()
	value = 4
	Put(CreateTuple(50, value), rsp, targets5)
	log("p11 w:%d", value)
	countwrites(value, 11, 4)
	l.Unlock()
}

func p12() {
	targets4 := make([]string, 0)
	targets3 := make([]string, 0)
	targets2 := make([]string, 2)
	targets2[1] = uri[s23]
	targets2[0] = uri[s2]
	targets1 := make([]string, 1)
	targets1[0] = uri[s6]
	targets0 := make([]string, 3)
	targets0[2] = uri[s19]
	targets0[1] = uri[s9]
	targets0[0] = uri[s3]

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(368, &value), rsp, s12)
	log("p12 r:%d", value)
	countreads(value, 12, 23)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(263, &value), rsp, s12)
	log("p12 r:%d", value)
	countreads(value, 12, 17)
	l.Unlock()

	delay()
	l.Lock()
	value = 21
	Put(CreateTuple(327, value), rsp, targets0)
	log("p12 w:%d", value)
	countwrites(value, 12, 21)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(484, &value), rsp, s12)
	log("p12 r:%d", value)
	countreads(value, 12, 31)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(437, &value), rsp, s12)
	log("p12 r:%d", value)
	countreads(value, 12, 28)
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
	QueryP(CreateTuple(1, &value), rsp, s12)
	log("p12 r:%d", value)
	countreads(value, 12, 1)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(87, &value), rsp, s12)
	log("p12 r:%d", value)
	countreads(value, 12, 6)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(124, &value), rsp, s12)
	log("p12 r:%d", value)
	countreads(value, 12, 8)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(292, &value), rsp, s12)
	log("p12 r:%d", value)
	countreads(value, 12, 19)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(368, &value), rsp, s12)
	log("p12 r:%d", value)
	countreads(value, 12, 23)
	l.Unlock()

	delay()
	l.Lock()
	value = 28
	Put(CreateTuple(441, value), rsp, targets1)
	log("p12 w:%d", value)
	countwrites(value, 12, 28)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(454, &value), rsp, s12)
	log("p12 r:%d", value)
	countreads(value, 12, 29)
	l.Unlock()

	delay()
	l.Lock()
	value = 27
	Put(CreateTuple(424, value), rsp, targets2)
	log("p12 w:%d", value)
	countwrites(value, 12, 27)
	l.Unlock()

	delay()
	l.Lock()
	value = 2
	Put(CreateTuple(23, value), rsp, targets3)
	log("p12 w:%d", value)
	countwrites(value, 12, 2)
	l.Unlock()

	delay()
	l.Lock()
	value = 12
	Put(CreateTuple(180, value), rsp, targets4)
	log("p12 w:%d", value)
	countwrites(value, 12, 12)
	l.Unlock()
}

func p13() {
	targets3 := make([]string, 0)
	targets2 := make([]string, 1)
	targets2[0] = uri[s22]
	targets1 := make([]string, 0)
	targets0 := make([]string, 2)
	targets0[1] = uri[s29]
	targets0[0] = uri[s29]

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(359, &value), rsp, s13)
	log("p13 r:%d", value)
	countreads(value, 13, 23)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(95, &value), rsp, s13)
	log("p13 r:%d", value)
	countreads(value, 13, 6)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(263, &value), rsp, s13)
	log("p13 r:%d", value)
	countreads(value, 13, 17)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(248, &value), rsp, s13)
	log("p13 r:%d", value)
	countreads(value, 13, 16)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(476, &value), rsp, s13)
	log("p13 r:%d", value)
	countreads(value, 13, 30)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(102, &value), rsp, s13)
	log("p13 r:%d", value)
	countreads(value, 13, 7)
	l.Unlock()

	delay()
	l.Lock()
	value = 7
	Put(CreateTuple(104, value), rsp, targets0)
	log("p13 w:%d", value)
	countwrites(value, 13, 7)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(457, &value), rsp, s13)
	log("p13 r:%d", value)
	countreads(value, 13, 29)
	l.Unlock()

	delay()
	l.Lock()
	value = 25
	Put(CreateTuple(390, value), rsp, targets1)
	log("p13 w:%d", value)
	countwrites(value, 13, 25)
	l.Unlock()

	delay()
	l.Lock()
	value = 22
	Put(CreateTuple(344, value), rsp, targets2)
	log("p13 w:%d", value)
	countwrites(value, 13, 22)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(412, &value), rsp, s13)
	log("p13 r:%d", value)
	countreads(value, 13, 26)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(376, &value), rsp, s13)
	log("p13 r:%d", value)
	countreads(value, 13, 24)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(85, &value), rsp, s13)
	log("p13 r:%d", value)
	countreads(value, 13, 6)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(486, &value), rsp, s13)
	log("p13 r:%d", value)
	countreads(value, 13, 31)
	l.Unlock()

	delay()
	l.Lock()
	value = 15
	Put(CreateTuple(226, value), rsp, targets3)
	log("p13 w:%d", value)
	countwrites(value, 13, 15)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(81, &value), rsp, s13)
	log("p13 r:%d", value)
	countreads(value, 13, 6)
	l.Unlock()
}

func p14() {
	targets5 := make([]string, 1)
	targets5[0] = uri[s19]
	targets4 := make([]string, 0)
	targets3 := make([]string, 0)
	targets2 := make([]string, 2)
	targets2[1] = uri[s27]
	targets2[0] = uri[s19]
	targets1 := make([]string, 0)
	targets0 := make([]string, 0)

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(117, &value), rsp, s14)
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
	value = 7
	Put(CreateTuple(108, value), rsp, targets0)
	log("p14 w:%d", value)
	countwrites(value, 14, 7)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(16, &value), rsp, s14)
	log("p14 r:%d", value)
	countreads(value, 14, 1)
	l.Unlock()

	delay()
	l.Lock()
	value = 18
	Put(CreateTuple(280, value), rsp, targets1)
	log("p14 w:%d", value)
	countwrites(value, 14, 18)
	l.Unlock()

	delay()
	l.Lock()
	value = 19
	Put(CreateTuple(303, value), rsp, targets2)
	log("p14 w:%d", value)
	countwrites(value, 14, 19)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(330, &value), rsp, s14)
	log("p14 r:%d", value)
	countreads(value, 14, 21)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(508, &value), rsp, s14)
	log("p14 r:%d", value)
	countreads(value, 14, 32)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(458, &value), rsp, s14)
	log("p14 r:%d", value)
	countreads(value, 14, 29)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(377, &value), rsp, s14)
	log("p14 r:%d", value)
	countreads(value, 14, 24)
	l.Unlock()

	delay()
	l.Lock()
	value = 23
	Put(CreateTuple(357, value), rsp, targets3)
	log("p14 w:%d", value)
	countwrites(value, 14, 23)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(219, &value), rsp, s14)
	log("p14 r:%d", value)
	countreads(value, 14, 14)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(330, &value), rsp, s14)
	log("p14 r:%d", value)
	countreads(value, 14, 21)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(287, &value), rsp, s14)
	log("p14 r:%d", value)
	countreads(value, 14, 18)
	l.Unlock()

	delay()
	l.Lock()
	value = 13
	Put(CreateTuple(196, value), rsp, targets4)
	log("p14 w:%d", value)
	countwrites(value, 14, 13)
	l.Unlock()

	delay()
	l.Lock()
	value = 4
	Put(CreateTuple(63, value), rsp, targets5)
	log("p14 w:%d", value)
	countwrites(value, 14, 4)
	l.Unlock()
}

func p15() {
	targets5 := make([]string, 0)
	targets4 := make([]string, 1)
	targets4[0] = uri[s4]
	targets3 := make([]string, 0)
	targets2 := make([]string, 1)
	targets2[0] = uri[s4]
	targets1 := make([]string, 2)
	targets1[1] = uri[s27]
	targets1[0] = uri[s17]
	targets0 := make([]string, 0)

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(168, &value), rsp, s15)
	log("p15 r:%d", value)
	countreads(value, 15, 11)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(224, &value), rsp, s15)
	log("p15 r:%d", value)
	countreads(value, 15, 14)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(222, &value), rsp, s15)
	log("p15 r:%d", value)
	countreads(value, 15, 14)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(98, &value), rsp, s15)
	log("p15 r:%d", value)
	countreads(value, 15, 7)
	l.Unlock()

	delay()
	l.Lock()
	value = 11
	Put(CreateTuple(169, value), rsp, targets0)
	log("p15 w:%d", value)
	countwrites(value, 15, 11)
	l.Unlock()

	delay()
	l.Lock()
	value = 3
	Put(CreateTuple(34, value), rsp, targets1)
	log("p15 w:%d", value)
	countwrites(value, 15, 3)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(247, &value), rsp, s15)
	log("p15 r:%d", value)
	countreads(value, 15, 16)
	l.Unlock()

	delay()
	l.Lock()
	value = 18
	Put(CreateTuple(284, value), rsp, targets2)
	log("p15 w:%d", value)
	countwrites(value, 15, 18)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(386, &value), rsp, s15)
	log("p15 r:%d", value)
	countreads(value, 15, 25)
	l.Unlock()

	delay()
	l.Lock()
	value = 20
	Put(CreateTuple(307, value), rsp, targets3)
	log("p15 w:%d", value)
	countwrites(value, 15, 20)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(426, &value), rsp, s15)
	log("p15 r:%d", value)
	countreads(value, 15, 27)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(484, &value), rsp, s15)
	log("p15 r:%d", value)
	countreads(value, 15, 31)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(272, &value), rsp, s15)
	log("p15 r:%d", value)
	countreads(value, 15, 17)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(127, &value), rsp, s15)
	log("p15 r:%d", value)
	countreads(value, 15, 8)
	l.Unlock()

	delay()
	l.Lock()
	value = 21
	Put(CreateTuple(328, value), rsp, targets4)
	log("p15 w:%d", value)
	countwrites(value, 15, 21)
	l.Unlock()

	delay()
	l.Lock()
	value = 8
	Put(CreateTuple(115, value), rsp, targets5)
	log("p15 w:%d", value)
	countwrites(value, 15, 8)
	l.Unlock()
}

func p16() {
	targets3 := make([]string, 0)
	targets2 := make([]string, 1)
	targets2[0] = uri[s10]
	targets1 := make([]string, 0)
	targets0 := make([]string, 0)

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(191, &value), rsp, s16)
	log("p16 r:%d", value)
	countreads(value, 16, 12)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(406, &value), rsp, s16)
	log("p16 r:%d", value)
	countreads(value, 16, 26)
	l.Unlock()

	delay()
	l.Lock()
	value = 8
	Put(CreateTuple(122, value), rsp, targets0)
	log("p16 w:%d", value)
	countwrites(value, 16, 8)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(395, &value), rsp, s16)
	log("p16 r:%d", value)
	countreads(value, 16, 25)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(370, &value), rsp, s16)
	log("p16 r:%d", value)
	countreads(value, 16, 24)
	l.Unlock()

	delay()
	l.Lock()
	value = 32
	Put(CreateTuple(498, value), rsp, targets1)
	log("p16 w:%d", value)
	countwrites(value, 16, 32)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(292, &value), rsp, s16)
	log("p16 r:%d", value)
	countreads(value, 16, 19)
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
	QueryP(CreateTuple(19, &value), rsp, s16)
	log("p16 r:%d", value)
	countreads(value, 16, 2)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(454, &value), rsp, s16)
	log("p16 r:%d", value)
	countreads(value, 16, 29)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(497, &value), rsp, s16)
	log("p16 r:%d", value)
	countreads(value, 16, 32)
	l.Unlock()

	delay()
	l.Lock()
	value = 31
	Put(CreateTuple(495, value), rsp, targets2)
	log("p16 w:%d", value)
	countwrites(value, 16, 31)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(336, &value), rsp, s16)
	log("p16 r:%d", value)
	countreads(value, 16, 21)
	l.Unlock()

	delay()
	l.Lock()
	value = 21
	Put(CreateTuple(324, value), rsp, targets3)
	log("p16 w:%d", value)
	countwrites(value, 16, 21)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(459, &value), rsp, s16)
	log("p16 r:%d", value)
	countreads(value, 16, 29)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(510, &value), rsp, s16)
	log("p16 r:%d", value)
	countreads(value, 16, 32)
	l.Unlock()
}

func p17() {
	targets5 := make([]string, 0)
	targets4 := make([]string, 2)
	targets4[1] = uri[s15]
	targets4[0] = uri[s12]
	targets3 := make([]string, 3)
	targets3[2] = uri[s22]
	targets3[1] = uri[s18]
	targets3[0] = uri[s11]
	targets2 := make([]string, 1)
	targets2[0] = uri[s29]
	targets1 := make([]string, 0)
	targets0 := make([]string, 1)
	targets0[0] = uri[s19]

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(314, &value), rsp, s17)
	log("p17 r:%d", value)
	countreads(value, 17, 20)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(418, &value), rsp, s17)
	log("p17 r:%d", value)
	countreads(value, 17, 27)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(230, &value), rsp, s17)
	log("p17 r:%d", value)
	countreads(value, 17, 15)
	l.Unlock()

	delay()
	l.Lock()
	value = 22
	Put(CreateTuple(338, value), rsp, targets0)
	log("p17 w:%d", value)
	countwrites(value, 17, 22)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(444, &value), rsp, s17)
	log("p17 r:%d", value)
	countreads(value, 17, 28)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(34, &value), rsp, s17)
	log("p17 r:%d", value)
	countreads(value, 17, 3)
	l.Unlock()

	delay()
	l.Lock()
	value = 25
	Put(CreateTuple(390, value), rsp, targets1)
	log("p17 w:%d", value)
	countwrites(value, 17, 25)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(135, &value), rsp, s17)
	log("p17 r:%d", value)
	countreads(value, 17, 9)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(173, &value), rsp, s17)
	log("p17 r:%d", value)
	countreads(value, 17, 11)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(485, &value), rsp, s17)
	log("p17 r:%d", value)
	countreads(value, 17, 31)
	l.Unlock()

	delay()
	l.Lock()
	value = 6
	Put(CreateTuple(82, value), rsp, targets2)
	log("p17 w:%d", value)
	countwrites(value, 17, 6)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(429, &value), rsp, s17)
	log("p17 r:%d", value)
	countreads(value, 17, 27)
	l.Unlock()

	delay()
	l.Lock()
	value = 21
	Put(CreateTuple(334, value), rsp, targets3)
	log("p17 w:%d", value)
	countwrites(value, 17, 21)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(64, &value), rsp, s17)
	log("p17 r:%d", value)
	countreads(value, 17, 4)
	l.Unlock()

	delay()
	l.Lock()
	value = 31
	Put(CreateTuple(484, value), rsp, targets4)
	log("p17 w:%d", value)
	countwrites(value, 17, 31)
	l.Unlock()

	delay()
	l.Lock()
	value = 2
	Put(CreateTuple(23, value), rsp, targets5)
	log("p17 w:%d", value)
	countwrites(value, 17, 2)
	l.Unlock()
}

func p18() {
	targets4 := make([]string, 0)
	targets3 := make([]string, 3)
	targets3[2] = uri[s24]
	targets3[1] = uri[s18]
	targets3[0] = uri[s14]
	targets2 := make([]string, 0)
	targets1 := make([]string, 1)
	targets1[0] = uri[s30]
	targets0 := make([]string, 0)

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 8
	Put(CreateTuple(118, value), rsp, targets0)
	log("p18 w:%d", value)
	countwrites(value, 18, 8)
	l.Unlock()

	delay()
	l.Lock()
	value = 29
	Put(CreateTuple(453, value), rsp, targets1)
	log("p18 w:%d", value)
	countwrites(value, 18, 29)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(458, &value), rsp, s18)
	log("p18 r:%d", value)
	countreads(value, 18, 29)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(490, &value), rsp, s18)
	log("p18 r:%d", value)
	countreads(value, 18, 31)
	l.Unlock()

	delay()
	l.Lock()
	value = 30
	Put(CreateTuple(477, value), rsp, targets2)
	log("p18 w:%d", value)
	countwrites(value, 18, 30)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(172, &value), rsp, s18)
	log("p18 r:%d", value)
	countreads(value, 18, 11)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(334, &value), rsp, s18)
	log("p18 r:%d", value)
	countreads(value, 18, 21)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(289, &value), rsp, s18)
	log("p18 r:%d", value)
	countreads(value, 18, 19)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(375, &value), rsp, s18)
	log("p18 r:%d", value)
	countreads(value, 18, 24)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(205, &value), rsp, s18)
	log("p18 r:%d", value)
	countreads(value, 18, 13)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(470, &value), rsp, s18)
	log("p18 r:%d", value)
	countreads(value, 18, 30)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(415, &value), rsp, s18)
	log("p18 r:%d", value)
	countreads(value, 18, 26)
	l.Unlock()

	delay()
	l.Lock()
	value = 29
	Put(CreateTuple(458, value), rsp, targets3)
	log("p18 w:%d", value)
	countwrites(value, 18, 29)
	l.Unlock()

	delay()
	l.Lock()
	value = 16
	Put(CreateTuple(249, value), rsp, targets4)
	log("p18 w:%d", value)
	countwrites(value, 18, 16)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(512, &value), rsp, s18)
	log("p18 r:%d", value)
	countreads(value, 18, 32)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(103, &value), rsp, s18)
	log("p18 r:%d", value)
	countreads(value, 18, 7)
	l.Unlock()
}

func p19() {
	targets1 := make([]string, 0)
	targets0 := make([]string, 2)
	targets0[1] = uri[s23]
	targets0[0] = uri[s14]

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(417, &value), rsp, s19)
	log("p19 r:%d", value)
	countreads(value, 19, 27)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(283, &value), rsp, s19)
	log("p19 r:%d", value)
	countreads(value, 19, 18)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(327, &value), rsp, s19)
	log("p19 r:%d", value)
	countreads(value, 19, 21)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(71, &value), rsp, s19)
	log("p19 r:%d", value)
	countreads(value, 19, 5)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(63, &value), rsp, s19)
	log("p19 r:%d", value)
	countreads(value, 19, 4)
	l.Unlock()

	delay()
	l.Lock()
	value = 18
	Put(CreateTuple(287, value), rsp, targets0)
	log("p19 w:%d", value)
	countwrites(value, 19, 18)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(223, &value), rsp, s19)
	log("p19 r:%d", value)
	countreads(value, 19, 14)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(264, &value), rsp, s19)
	log("p19 r:%d", value)
	countreads(value, 19, 17)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(306, &value), rsp, s19)
	log("p19 r:%d", value)
	countreads(value, 19, 20)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(509, &value), rsp, s19)
	log("p19 r:%d", value)
	countreads(value, 19, 32)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(338, &value), rsp, s19)
	log("p19 r:%d", value)
	countreads(value, 19, 22)
	l.Unlock()

	delay()
	l.Lock()
	value = 13
	Put(CreateTuple(201, value), rsp, targets1)
	log("p19 w:%d", value)
	countwrites(value, 19, 13)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(303, &value), rsp, s19)
	log("p19 r:%d", value)
	countreads(value, 19, 19)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(388, &value), rsp, s19)
	log("p19 r:%d", value)
	countreads(value, 19, 25)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(134, &value), rsp, s19)
	log("p19 r:%d", value)
	countreads(value, 19, 9)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(362, &value), rsp, s19)
	log("p19 r:%d", value)
	countreads(value, 19, 23)
	l.Unlock()
}

func p20() {
	targets5 := make([]string, 2)
	targets5[1] = uri[s21]
	targets5[0] = uri[s13]
	targets4 := make([]string, 0)
	targets3 := make([]string, 1)
	targets3[0] = uri[s23]
	targets2 := make([]string, 0)
	targets1 := make([]string, 1)
	targets1[0] = uri[s19]
	targets0 := make([]string, 0)

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(493, &value), rsp, s20)
	log("p20 r:%d", value)
	countreads(value, 20, 31)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(31, &value), rsp, s20)
	log("p20 r:%d", value)
	countreads(value, 20, 2)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(132, &value), rsp, s20)
	log("p20 r:%d", value)
	countreads(value, 20, 9)
	l.Unlock()

	delay()
	l.Lock()
	value = 11
	Put(CreateTuple(163, value), rsp, targets0)
	log("p20 w:%d", value)
	countwrites(value, 20, 11)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(345, &value), rsp, s20)
	log("p20 r:%d", value)
	countreads(value, 20, 22)
	l.Unlock()

	delay()
	l.Lock()
	value = 22
	Put(CreateTuple(338, value), rsp, targets1)
	log("p20 w:%d", value)
	countwrites(value, 20, 22)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(155, &value), rsp, s20)
	log("p20 r:%d", value)
	countreads(value, 20, 10)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(125, &value), rsp, s20)
	log("p20 r:%d", value)
	countreads(value, 20, 8)
	l.Unlock()

	delay()
	l.Lock()
	value = 9
	Put(CreateTuple(141, value), rsp, targets2)
	log("p20 w:%d", value)
	countwrites(value, 20, 9)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(293, &value), rsp, s20)
	log("p20 r:%d", value)
	countreads(value, 20, 19)
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
	value = 25
	Put(CreateTuple(399, value), rsp, targets3)
	log("p20 w:%d", value)
	countwrites(value, 20, 25)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(331, &value), rsp, s20)
	log("p20 r:%d", value)
	countreads(value, 20, 21)
	l.Unlock()

	delay()
	l.Lock()
	value = 21
	Put(CreateTuple(322, value), rsp, targets4)
	log("p20 w:%d", value)
	countwrites(value, 20, 21)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(437, &value), rsp, s20)
	log("p20 r:%d", value)
	countreads(value, 20, 28)
	l.Unlock()

	delay()
	l.Lock()
	value = 6
	Put(CreateTuple(85, value), rsp, targets5)
	log("p20 w:%d", value)
	countwrites(value, 20, 6)
	l.Unlock()
}

func p21() {
	targets5 := make([]string, 0)
	targets4 := make([]string, 1)
	targets4[0] = uri[s29]
	targets3 := make([]string, 0)
	targets2 := make([]string, 0)
	targets1 := make([]string, 0)
	targets0 := make([]string, 0)

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(229, &value), rsp, s21)
	log("p21 r:%d", value)
	countreads(value, 21, 15)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(31, &value), rsp, s21)
	log("p21 r:%d", value)
	countreads(value, 21, 2)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(289, &value), rsp, s21)
	log("p21 r:%d", value)
	countreads(value, 21, 19)
	l.Unlock()

	delay()
	l.Lock()
	value = 26
	Put(CreateTuple(409, value), rsp, targets0)
	log("p21 w:%d", value)
	countwrites(value, 21, 26)
	l.Unlock()

	delay()
	l.Lock()
	value = 3
	Put(CreateTuple(41, value), rsp, targets1)
	log("p21 w:%d", value)
	countwrites(value, 21, 3)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(274, &value), rsp, s21)
	log("p21 r:%d", value)
	countreads(value, 21, 18)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(429, &value), rsp, s21)
	log("p21 r:%d", value)
	countreads(value, 21, 27)
	l.Unlock()

	delay()
	l.Lock()
	value = 10
	Put(CreateTuple(148, value), rsp, targets2)
	log("p21 w:%d", value)
	countwrites(value, 21, 10)
	l.Unlock()

	delay()
	l.Lock()
	value = 26
	Put(CreateTuple(414, value), rsp, targets3)
	log("p21 w:%d", value)
	countwrites(value, 21, 26)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(155, &value), rsp, s21)
	log("p21 r:%d", value)
	countreads(value, 21, 10)
	l.Unlock()

	delay()
	l.Lock()
	value = 10
	Put(CreateTuple(150, value), rsp, targets4)
	log("p21 w:%d", value)
	countwrites(value, 21, 10)
	l.Unlock()

	delay()
	l.Lock()
	value = 20
	Put(CreateTuple(308, value), rsp, targets5)
	log("p21 w:%d", value)
	countwrites(value, 21, 20)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(274, &value), rsp, s21)
	log("p21 r:%d", value)
	countreads(value, 21, 18)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(395, &value), rsp, s21)
	log("p21 r:%d", value)
	countreads(value, 21, 25)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(85, &value), rsp, s21)
	log("p21 r:%d", value)
	countreads(value, 21, 6)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(66, &value), rsp, s21)
	log("p21 r:%d", value)
	countreads(value, 21, 5)
	l.Unlock()
}

func p22() {
	targets5 := make([]string, 0)
	targets4 := make([]string, 0)
	targets3 := make([]string, 0)
	targets2 := make([]string, 0)
	targets1 := make([]string, 0)
	targets0 := make([]string, 2)
	targets0[1] = uri[s3]
	targets0[0] = uri[s1]

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(311, &value), rsp, s22)
	log("p22 r:%d", value)
	countreads(value, 22, 20)
	l.Unlock()

	delay()
	l.Lock()
	value = 13
	Put(CreateTuple(197, value), rsp, targets0)
	log("p22 w:%d", value)
	countwrites(value, 22, 13)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(344, &value), rsp, s22)
	log("p22 r:%d", value)
	countreads(value, 22, 22)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(62, &value), rsp, s22)
	log("p22 r:%d", value)
	countreads(value, 22, 4)
	l.Unlock()

	delay()
	l.Lock()
	value = 12
	Put(CreateTuple(177, value), rsp, targets1)
	log("p22 w:%d", value)
	countwrites(value, 22, 12)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(234, &value), rsp, s22)
	log("p22 r:%d", value)
	countreads(value, 22, 15)
	l.Unlock()

	delay()
	l.Lock()
	value = 25
	Put(CreateTuple(394, value), rsp, targets2)
	log("p22 w:%d", value)
	countwrites(value, 22, 25)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(295, &value), rsp, s22)
	log("p22 r:%d", value)
	countreads(value, 22, 19)
	l.Unlock()

	delay()
	l.Lock()
	value = 31
	Put(CreateTuple(481, value), rsp, targets3)
	log("p22 w:%d", value)
	countwrites(value, 22, 31)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(419, &value), rsp, s22)
	log("p22 r:%d", value)
	countreads(value, 22, 27)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(254, &value), rsp, s22)
	log("p22 r:%d", value)
	countreads(value, 22, 16)
	l.Unlock()

	delay()
	l.Lock()
	value = 18
	Put(CreateTuple(281, value), rsp, targets4)
	log("p22 w:%d", value)
	countwrites(value, 22, 18)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(55, &value), rsp, s22)
	log("p22 r:%d", value)
	countreads(value, 22, 4)
	l.Unlock()

	delay()
	l.Lock()
	value = 4
	Put(CreateTuple(53, value), rsp, targets5)
	log("p22 w:%d", value)
	countwrites(value, 22, 4)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(124, &value), rsp, s22)
	log("p22 r:%d", value)
	countreads(value, 22, 8)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(334, &value), rsp, s22)
	log("p22 r:%d", value)
	countreads(value, 22, 21)
	l.Unlock()
}

func p23() {
	targets5 := make([]string, 1)
	targets5[0] = uri[s13]
	targets4 := make([]string, 0)
	targets3 := make([]string, 0)
	targets2 := make([]string, 1)
	targets2[0] = uri[s28]
	targets1 := make([]string, 2)
	targets1[1] = uri[s32]
	targets1[0] = uri[s8]
	targets0 := make([]string, 2)
	targets0[1] = uri[s16]
	targets0[0] = uri[s10]

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 32
	Put(CreateTuple(497, value), rsp, targets0)
	log("p23 w:%d", value)
	countwrites(value, 23, 32)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(399, &value), rsp, s23)
	log("p23 r:%d", value)
	countreads(value, 23, 25)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(364, &value), rsp, s23)
	log("p23 r:%d", value)
	countreads(value, 23, 23)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(95, &value), rsp, s23)
	log("p23 r:%d", value)
	countreads(value, 23, 6)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(287, &value), rsp, s23)
	log("p23 r:%d", value)
	countreads(value, 23, 18)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(483, &value), rsp, s23)
	log("p23 r:%d", value)
	countreads(value, 23, 31)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(424, &value), rsp, s23)
	log("p23 r:%d", value)
	countreads(value, 23, 27)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(175, &value), rsp, s23)
	log("p23 r:%d", value)
	countreads(value, 23, 11)
	l.Unlock()

	delay()
	l.Lock()
	value = 23
	Put(CreateTuple(358, value), rsp, targets1)
	log("p23 w:%d", value)
	countwrites(value, 23, 23)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(352, &value), rsp, s23)
	log("p23 r:%d", value)
	countreads(value, 23, 22)
	l.Unlock()

	delay()
	l.Lock()
	value = 17
	Put(CreateTuple(259, value), rsp, targets2)
	log("p23 w:%d", value)
	countwrites(value, 23, 17)
	l.Unlock()

	delay()
	l.Lock()
	value = 19
	Put(CreateTuple(300, value), rsp, targets3)
	log("p23 w:%d", value)
	countwrites(value, 23, 19)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(125, &value), rsp, s23)
	log("p23 r:%d", value)
	countreads(value, 23, 8)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(29, &value), rsp, s23)
	log("p23 r:%d", value)
	countreads(value, 23, 2)
	l.Unlock()

	delay()
	l.Lock()
	value = 4
	Put(CreateTuple(58, value), rsp, targets4)
	log("p23 w:%d", value)
	countwrites(value, 23, 4)
	l.Unlock()

	delay()
	l.Lock()
	value = 26
	Put(CreateTuple(412, value), rsp, targets5)
	log("p23 w:%d", value)
	countwrites(value, 23, 26)
	l.Unlock()
}

func p24() {
	targets6 := make([]string, 2)
	targets6[1] = uri[s27]
	targets6[0] = uri[s17]
	targets5 := make([]string, 0)
	targets4 := make([]string, 1)
	targets4[0] = uri[s8]
	targets3 := make([]string, 0)
	targets2 := make([]string, 0)
	targets1 := make([]string, 0)
	targets0 := make([]string, 0)

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 5
	Put(CreateTuple(74, value), rsp, targets0)
	log("p24 w:%d", value)
	countwrites(value, 24, 5)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(223, &value), rsp, s24)
	log("p24 r:%d", value)
	countreads(value, 24, 14)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(16, &value), rsp, s24)
	log("p24 r:%d", value)
	countreads(value, 24, 1)
	l.Unlock()

	delay()
	l.Lock()
	value = 3
	Put(CreateTuple(33, value), rsp, targets1)
	log("p24 w:%d", value)
	countwrites(value, 24, 3)
	l.Unlock()

	delay()
	l.Lock()
	value = 4
	Put(CreateTuple(49, value), rsp, targets2)
	log("p24 w:%d", value)
	countwrites(value, 24, 4)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(362, &value), rsp, s24)
	log("p24 r:%d", value)
	countreads(value, 24, 23)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(129, &value), rsp, s24)
	log("p24 r:%d", value)
	countreads(value, 24, 9)
	l.Unlock()

	delay()
	l.Lock()
	value = 19
	Put(CreateTuple(300, value), rsp, targets3)
	log("p24 w:%d", value)
	countwrites(value, 24, 19)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(458, &value), rsp, s24)
	log("p24 r:%d", value)
	countreads(value, 24, 29)
	l.Unlock()

	delay()
	l.Lock()
	value = 7
	Put(CreateTuple(97, value), rsp, targets4)
	log("p24 w:%d", value)
	countwrites(value, 24, 7)
	l.Unlock()

	delay()
	l.Lock()
	value = 23
	Put(CreateTuple(361, value), rsp, targets5)
	log("p24 w:%d", value)
	countwrites(value, 24, 23)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(125, &value), rsp, s24)
	log("p24 r:%d", value)
	countreads(value, 24, 8)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(30, &value), rsp, s24)
	log("p24 r:%d", value)
	countreads(value, 24, 2)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(65, &value), rsp, s24)
	log("p24 r:%d", value)
	countreads(value, 24, 5)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(216, &value), rsp, s24)
	log("p24 r:%d", value)
	countreads(value, 24, 14)
	l.Unlock()

	delay()
	l.Lock()
	value = 3
	Put(CreateTuple(34, value), rsp, targets6)
	log("p24 w:%d", value)
	countwrites(value, 24, 3)
	l.Unlock()
}

func p25() {
	targets3 := make([]string, 1)
	targets3[0] = uri[s28]
	targets2 := make([]string, 1)
	targets2[0] = uri[s19]
	targets1 := make([]string, 0)
	targets0 := make([]string, 1)
	targets0[0] = uri[s3]

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(242, &value), rsp, s25)
	log("p25 r:%d", value)
	countreads(value, 25, 16)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(31, &value), rsp, s25)
	log("p25 r:%d", value)
	countreads(value, 25, 2)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(200, &value), rsp, s25)
	log("p25 r:%d", value)
	countreads(value, 25, 13)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(333, &value), rsp, s25)
	log("p25 r:%d", value)
	countreads(value, 25, 21)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(363, &value), rsp, s25)
	log("p25 r:%d", value)
	countreads(value, 25, 23)
	l.Unlock()

	delay()
	l.Lock()
	value = 2
	Put(CreateTuple(22, value), rsp, targets0)
	log("p25 w:%d", value)
	countwrites(value, 25, 2)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(335, &value), rsp, s25)
	log("p25 r:%d", value)
	countreads(value, 25, 21)
	l.Unlock()

	delay()
	l.Lock()
	value = 24
	Put(CreateTuple(372, value), rsp, targets1)
	log("p25 w:%d", value)
	countwrites(value, 25, 24)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(66, &value), rsp, s25)
	log("p25 r:%d", value)
	countreads(value, 25, 5)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(429, &value), rsp, s25)
	log("p25 r:%d", value)
	countreads(value, 25, 27)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(499, &value), rsp, s25)
	log("p25 r:%d", value)
	countreads(value, 25, 32)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(428, &value), rsp, s25)
	log("p25 r:%d", value)
	countreads(value, 25, 27)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(407, &value), rsp, s25)
	log("p25 r:%d", value)
	countreads(value, 25, 26)
	l.Unlock()

	delay()
	l.Lock()
	value = 4
	Put(CreateTuple(63, value), rsp, targets2)
	log("p25 w:%d", value)
	countwrites(value, 25, 4)
	l.Unlock()

	delay()
	l.Lock()
	value = 13
	Put(CreateTuple(204, value), rsp, targets3)
	log("p25 w:%d", value)
	countwrites(value, 25, 13)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(400, &value), rsp, s25)
	log("p25 r:%d", value)
	countreads(value, 25, 25)
	l.Unlock()
}

func p26() {
	targets5 := make([]string, 1)
	targets5[0] = uri[s16]
	targets4 := make([]string, 1)
	targets4[0] = uri[s15]
	targets3 := make([]string, 0)
	targets2 := make([]string, 0)
	targets1 := make([]string, 0)
	targets0 := make([]string, 0)

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(162, &value), rsp, s26)
	log("p26 r:%d", value)
	countreads(value, 26, 11)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(130, &value), rsp, s26)
	log("p26 r:%d", value)
	countreads(value, 26, 9)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(227, &value), rsp, s26)
	log("p26 r:%d", value)
	countreads(value, 26, 15)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(195, &value), rsp, s26)
	log("p26 r:%d", value)
	countreads(value, 26, 13)
	l.Unlock()

	delay()
	l.Lock()
	value = 12
	Put(CreateTuple(183, value), rsp, targets0)
	log("p26 w:%d", value)
	countwrites(value, 26, 12)
	l.Unlock()

	delay()
	l.Lock()
	value = 29
	Put(CreateTuple(456, value), rsp, targets1)
	log("p26 w:%d", value)
	countwrites(value, 26, 29)
	l.Unlock()

	delay()
	l.Lock()
	value = 1
	Put(CreateTuple(7, value), rsp, targets2)
	log("p26 w:%d", value)
	countwrites(value, 26, 1)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(379, &value), rsp, s26)
	log("p26 r:%d", value)
	countreads(value, 26, 24)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(109, &value), rsp, s26)
	log("p26 r:%d", value)
	countreads(value, 26, 7)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(96, &value), rsp, s26)
	log("p26 r:%d", value)
	countreads(value, 26, 6)
	l.Unlock()

	delay()
	l.Lock()
	value = 5
	Put(CreateTuple(79, value), rsp, targets3)
	log("p26 w:%d", value)
	countwrites(value, 26, 5)
	l.Unlock()

	delay()
	l.Lock()
	value = 16
	Put(CreateTuple(247, value), rsp, targets4)
	log("p26 w:%d", value)
	countwrites(value, 26, 16)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(419, &value), rsp, s26)
	log("p26 r:%d", value)
	countreads(value, 26, 27)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(223, &value), rsp, s26)
	log("p26 r:%d", value)
	countreads(value, 26, 14)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(90, &value), rsp, s26)
	log("p26 r:%d", value)
	countreads(value, 26, 6)
	l.Unlock()

	delay()
	l.Lock()
	value = 2
	Put(CreateTuple(19, value), rsp, targets5)
	log("p26 w:%d", value)
	countwrites(value, 26, 2)
	l.Unlock()
}

func p27() {
	targets4 := make([]string, 0)
	targets3 := make([]string, 0)
	targets2 := make([]string, 0)
	targets1 := make([]string, 1)
	targets1[0] = uri[s23]
	targets0 := make([]string, 1)
	targets0[0] = uri[s31]

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(26, &value), rsp, s27)
	log("p27 r:%d", value)
	countreads(value, 27, 2)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(379, &value), rsp, s27)
	log("p27 r:%d", value)
	countreads(value, 27, 24)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(232, &value), rsp, s27)
	log("p27 r:%d", value)
	countreads(value, 27, 15)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(303, &value), rsp, s27)
	log("p27 r:%d", value)
	countreads(value, 27, 19)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(34, &value), rsp, s27)
	log("p27 r:%d", value)
	countreads(value, 27, 3)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(393, &value), rsp, s27)
	log("p27 r:%d", value)
	countreads(value, 27, 25)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(135, &value), rsp, s27)
	log("p27 r:%d", value)
	countreads(value, 27, 9)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(333, &value), rsp, s27)
	log("p27 r:%d", value)
	countreads(value, 27, 21)
	l.Unlock()

	delay()
	l.Lock()
	value = 20
	Put(CreateTuple(310, value), rsp, targets0)
	log("p27 w:%d", value)
	countwrites(value, 27, 20)
	l.Unlock()

	delay()
	l.Lock()
	value = 31
	Put(CreateTuple(483, value), rsp, targets1)
	log("p27 w:%d", value)
	countwrites(value, 27, 31)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(217, &value), rsp, s27)
	log("p27 r:%d", value)
	countreads(value, 27, 14)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(117, &value), rsp, s27)
	log("p27 r:%d", value)
	countreads(value, 27, 8)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(290, &value), rsp, s27)
	log("p27 r:%d", value)
	countreads(value, 27, 19)
	l.Unlock()

	delay()
	l.Lock()
	value = 28
	Put(CreateTuple(440, value), rsp, targets2)
	log("p27 w:%d", value)
	countwrites(value, 27, 28)
	l.Unlock()

	delay()
	l.Lock()
	value = 20
	Put(CreateTuple(317, value), rsp, targets3)
	log("p27 w:%d", value)
	countwrites(value, 27, 20)
	l.Unlock()

	delay()
	l.Lock()
	value = 20
	Put(CreateTuple(313, value), rsp, targets4)
	log("p27 w:%d", value)
	countwrites(value, 27, 20)
	l.Unlock()
}

func p28() {
	targets6 := make([]string, 1)
	targets6[0] = uri[s8]
	targets5 := make([]string, 3)
	targets5[2] = uri[s17]
	targets5[1] = uri[s11]
	targets5[0] = uri[s1]
	targets4 := make([]string, 2)
	targets4[1] = uri[s31]
	targets4[0] = uri[s12]
	targets3 := make([]string, 0)
	targets2 := make([]string, 2)
	targets2[1] = uri[s25]
	targets2[0] = uri[s6]
	targets1 := make([]string, 0)
	targets0 := make([]string, 0)

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(206, &value), rsp, s28)
	log("p28 r:%d", value)
	countreads(value, 28, 13)
	l.Unlock()

	delay()
	l.Lock()
	value = 31
	Put(CreateTuple(489, value), rsp, targets0)
	log("p28 w:%d", value)
	countwrites(value, 28, 31)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(137, &value), rsp, s28)
	log("p28 r:%d", value)
	countreads(value, 28, 9)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(259, &value), rsp, s28)
	log("p28 r:%d", value)
	countreads(value, 28, 17)
	l.Unlock()

	delay()
	l.Lock()
	value = 20
	Put(CreateTuple(320, value), rsp, targets1)
	log("p28 w:%d", value)
	countwrites(value, 28, 20)
	l.Unlock()

	delay()
	l.Lock()
	value = 26
	Put(CreateTuple(407, value), rsp, targets2)
	log("p28 w:%d", value)
	countwrites(value, 28, 26)
	l.Unlock()

	delay()
	l.Lock()
	value = 28
	Put(CreateTuple(435, value), rsp, targets3)
	log("p28 w:%d", value)
	countwrites(value, 28, 28)
	l.Unlock()

	delay()
	l.Lock()
	value = 6
	Put(CreateTuple(87, value), rsp, targets4)
	log("p28 w:%d", value)
	countwrites(value, 28, 6)
	l.Unlock()

	delay()
	l.Lock()
	value = 11
	Put(CreateTuple(173, value), rsp, targets5)
	log("p28 w:%d", value)
	countwrites(value, 28, 11)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(80, &value), rsp, s28)
	log("p28 r:%d", value)
	countreads(value, 28, 5)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(216, &value), rsp, s28)
	log("p28 r:%d", value)
	countreads(value, 28, 14)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(186, &value), rsp, s28)
	log("p28 r:%d", value)
	countreads(value, 28, 12)
	l.Unlock()

	delay()
	l.Lock()
	value = 26
	Put(CreateTuple(403, value), rsp, targets6)
	log("p28 w:%d", value)
	countwrites(value, 28, 26)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(404, &value), rsp, s28)
	log("p28 r:%d", value)
	countreads(value, 28, 26)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(204, &value), rsp, s28)
	log("p28 r:%d", value)
	countreads(value, 28, 13)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(155, &value), rsp, s28)
	log("p28 r:%d", value)
	countreads(value, 28, 10)
	l.Unlock()
}

func p29() {
	targets1 := make([]string, 0)
	targets0 := make([]string, 0)
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(132, &value), rsp, s29)
	log("p29 r:%d", value)
	countreads(value, 29, 9)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(69, &value), rsp, s29)
	log("p29 r:%d", value)
	countreads(value, 29, 5)
	l.Unlock()

	delay()
	l.Lock()
	value = 29
	Put(CreateTuple(462, value), rsp, targets0)
	log("p29 w:%d", value)
	countwrites(value, 29, 29)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(82, &value), rsp, s29)
	log("p29 r:%d", value)
	countreads(value, 29, 6)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(285, &value), rsp, s29)
	log("p29 r:%d", value)
	countreads(value, 29, 18)
	l.Unlock()

	delay()
	l.Lock()
	value = 23
	Put(CreateTuple(354, value), rsp, targets1)
	log("p29 w:%d", value)
	countwrites(value, 29, 23)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(376, &value), rsp, s29)
	log("p29 r:%d", value)
	countreads(value, 29, 24)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(39, &value), rsp, s29)
	log("p29 r:%d", value)
	countreads(value, 29, 3)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(104, &value), rsp, s29)
	log("p29 r:%d", value)
	countreads(value, 29, 7)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(139, &value), rsp, s29)
	log("p29 r:%d", value)
	countreads(value, 29, 9)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(106, &value), rsp, s29)
	log("p29 r:%d", value)
	countreads(value, 29, 7)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(56, &value), rsp, s29)
	log("p29 r:%d", value)
	countreads(value, 29, 4)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(150, &value), rsp, s29)
	log("p29 r:%d", value)
	countreads(value, 29, 10)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(244, &value), rsp, s29)
	log("p29 r:%d", value)
	countreads(value, 29, 16)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(376, &value), rsp, s29)
	log("p29 r:%d", value)
	countreads(value, 29, 24)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(104, &value), rsp, s29)
	log("p29 r:%d", value)
	countreads(value, 29, 7)
	l.Unlock()
}

func p30() {
	targets4 := make([]string, 1)
	targets4[0] = uri[s5]
	targets3 := make([]string, 0)
	targets2 := make([]string, 1)
	targets2[0] = uri[s28]
	targets1 := make([]string, 1)
	targets1[0] = uri[s9]
	targets0 := make([]string, 1)
	targets0[0] = uri[s9]

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 29
	Put(CreateTuple(461, value), rsp, targets0)
	log("p30 w:%d", value)
	countwrites(value, 30, 29)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(24, &value), rsp, s30)
	log("p30 r:%d", value)
	countreads(value, 30, 2)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(35, &value), rsp, s30)
	log("p30 r:%d", value)
	countreads(value, 30, 3)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(38, &value), rsp, s30)
	log("p30 r:%d", value)
	countreads(value, 30, 3)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(466, &value), rsp, s30)
	log("p30 r:%d", value)
	countreads(value, 30, 30)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(507, &value), rsp, s30)
	log("p30 r:%d", value)
	countreads(value, 30, 32)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(27, &value), rsp, s30)
	log("p30 r:%d", value)
	countreads(value, 30, 2)
	l.Unlock()

	delay()
	l.Lock()
	value = 6
	Put(CreateTuple(93, value), rsp, targets1)
	log("p30 w:%d", value)
	countwrites(value, 30, 6)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(453, &value), rsp, s30)
	log("p30 r:%d", value)
	countreads(value, 30, 29)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(469, &value), rsp, s30)
	log("p30 r:%d", value)
	countreads(value, 30, 30)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(154, &value), rsp, s30)
	log("p30 r:%d", value)
	countreads(value, 30, 10)
	l.Unlock()

	delay()
	l.Lock()
	value = 13
	Put(CreateTuple(206, value), rsp, targets2)
	log("p30 w:%d", value)
	countwrites(value, 30, 13)
	l.Unlock()

	delay()
	l.Lock()
	value = 4
	Put(CreateTuple(51, value), rsp, targets3)
	log("p30 w:%d", value)
	countwrites(value, 30, 4)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(340, &value), rsp, s30)
	log("p30 r:%d", value)
	countreads(value, 30, 22)
	l.Unlock()

	delay()
	l.Lock()
	value = 14
	Put(CreateTuple(215, value), rsp, targets4)
	log("p30 w:%d", value)
	countwrites(value, 30, 14)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(162, &value), rsp, s30)
	log("p30 r:%d", value)
	countreads(value, 30, 11)
	l.Unlock()
}

func p31() {
	targets2 := make([]string, 0)
	targets1 := make([]string, 1)
	targets1[0] = uri[s15]
	targets0 := make([]string, 0)

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(288, &value), rsp, s31)
	log("p31 r:%d", value)
	countreads(value, 31, 18)
	l.Unlock()

	delay()
	l.Lock()
	value = 8
	Put(CreateTuple(128, value), rsp, targets0)
	log("p31 w:%d", value)
	countwrites(value, 31, 8)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(267, &value), rsp, s31)
	log("p31 r:%d", value)
	countreads(value, 31, 17)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(83, &value), rsp, s31)
	log("p31 r:%d", value)
	countreads(value, 31, 6)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(140, &value), rsp, s31)
	log("p31 r:%d", value)
	countreads(value, 31, 9)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(310, &value), rsp, s31)
	log("p31 r:%d", value)
	countreads(value, 31, 20)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(408, &value), rsp, s31)
	log("p31 r:%d", value)
	countreads(value, 31, 26)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(507, &value), rsp, s31)
	log("p31 r:%d", value)
	countreads(value, 31, 32)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(421, &value), rsp, s31)
	log("p31 r:%d", value)
	countreads(value, 31, 27)
	l.Unlock()

	delay()
	l.Lock()
	value = 11
	Put(CreateTuple(168, value), rsp, targets1)
	log("p31 w:%d", value)
	countwrites(value, 31, 11)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(87, &value), rsp, s31)
	log("p31 r:%d", value)
	countreads(value, 31, 6)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(253, &value), rsp, s31)
	log("p31 r:%d", value)
	countreads(value, 31, 16)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(62, &value), rsp, s31)
	log("p31 r:%d", value)
	countreads(value, 31, 4)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(69, &value), rsp, s31)
	log("p31 r:%d", value)
	countreads(value, 31, 5)
	l.Unlock()

	delay()
	l.Lock()
	value = 15
	Put(CreateTuple(225, value), rsp, targets2)
	log("p31 w:%d", value)
	countwrites(value, 31, 15)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(454, &value), rsp, s31)
	log("p31 r:%d", value)
	countreads(value, 31, 29)
	l.Unlock()
}

func p32() {
	targets1 := make([]string, 0)
	targets0 := make([]string, 1)
	targets0[0] = uri[s5]

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(52, &value), rsp, s32)
	log("p32 r:%d", value)
	countreads(value, 32, 4)
	l.Unlock()

	delay()
	l.Lock()
	value = 31
	Put(CreateTuple(491, value), rsp, targets0)
	log("p32 w:%d", value)
	countwrites(value, 32, 31)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(312, &value), rsp, s32)
	log("p32 r:%d", value)
	countreads(value, 32, 20)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(86, &value), rsp, s32)
	log("p32 r:%d", value)
	countreads(value, 32, 6)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(475, &value), rsp, s32)
	log("p32 r:%d", value)
	countreads(value, 32, 30)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(321, &value), rsp, s32)
	log("p32 r:%d", value)
	countreads(value, 32, 21)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(84, &value), rsp, s32)
	log("p32 r:%d", value)
	countreads(value, 32, 6)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(222, &value), rsp, s32)
	log("p32 r:%d", value)
	countreads(value, 32, 14)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(154, &value), rsp, s32)
	log("p32 r:%d", value)
	countreads(value, 32, 10)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(365, &value), rsp, s32)
	log("p32 r:%d", value)
	countreads(value, 32, 23)
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
	QueryP(CreateTuple(358, &value), rsp, s32)
	log("p32 r:%d", value)
	countreads(value, 32, 23)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(415, &value), rsp, s32)
	log("p32 r:%d", value)
	countreads(value, 32, 26)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(263, &value), rsp, s32)
	log("p32 r:%d", value)
	countreads(value, 32, 17)
	l.Unlock()

	delay()
	l.Lock()
	value = 17
	Put(CreateTuple(262, value), rsp, targets1)
	log("p32 w:%d", value)
	countwrites(value, 32, 17)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(470, &value), rsp, s32)
	log("p32 r:%d", value)
	countreads(value, 32, 30)
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
