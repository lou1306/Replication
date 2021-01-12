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
	targets6 := make([]string, 0)
	targets5 := make([]string, 0)
	targets4 := make([]string, 1)
	targets4[0] = uri[s12]
	targets3 := make([]string, 0)
	targets2 := make([]string, 0)
	targets1 := make([]string, 1)
	targets1[0] = uri[s12]
	targets0 := make([]string, 0)

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(92, &value), rsp, s1)
	log("p1 r:%d", value)
	countreads(value, 1, 6)
	l.Unlock()

	delay()
	l.Lock()
	value = 16
	Put(CreateTuple(256, value), rsp, targets0)
	log("p1 w:%d", value)
	countwrites(value, 1, 16)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(180, &value), rsp, s1)
	log("p1 r:%d", value)
	countreads(value, 1, 12)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(253, &value), rsp, s1)
	log("p1 r:%d", value)
	countreads(value, 1, 16)
	l.Unlock()

	delay()
	l.Lock()
	value = 17
	Put(CreateTuple(261, value), rsp, targets1)
	log("p1 w:%d", value)
	countwrites(value, 1, 17)
	l.Unlock()

	delay()
	l.Lock()
	value = 20
	Put(CreateTuple(316, value), rsp, targets2)
	log("p1 w:%d", value)
	countwrites(value, 1, 20)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(237, &value), rsp, s1)
	log("p1 r:%d", value)
	countreads(value, 1, 15)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(338, &value), rsp, s1)
	log("p1 r:%d", value)
	countreads(value, 1, 22)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(291, &value), rsp, s1)
	log("p1 r:%d", value)
	countreads(value, 1, 19)
	l.Unlock()

	delay()
	l.Lock()
	value = 9
	Put(CreateTuple(143, value), rsp, targets3)
	log("p1 w:%d", value)
	countwrites(value, 1, 9)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(303, &value), rsp, s1)
	log("p1 r:%d", value)
	countreads(value, 1, 19)
	l.Unlock()

	delay()
	l.Lock()
	value = 17
	Put(CreateTuple(261, value), rsp, targets4)
	log("p1 w:%d", value)
	countwrites(value, 1, 17)
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
	value = 1
	Put(CreateTuple(13, value), rsp, targets5)
	log("p1 w:%d", value)
	countwrites(value, 1, 1)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(55, &value), rsp, s1)
	log("p1 r:%d", value)
	countreads(value, 1, 4)
	l.Unlock()

	delay()
	l.Lock()
	value = 3
	Put(CreateTuple(48, value), rsp, targets6)
	log("p1 w:%d", value)
	countwrites(value, 1, 3)
	l.Unlock()
}

func p2() {
	targets6 := make([]string, 1)
	targets6[0] = uri[s18]
	targets5 := make([]string, 2)
	targets5[1] = uri[s31]
	targets5[0] = uri[s3]
	targets4 := make([]string, 1)
	targets4[0] = uri[s13]
	targets3 := make([]string, 0)
	targets2 := make([]string, 1)
	targets2[0] = uri[s29]
	targets1 := make([]string, 0)
	targets0 := make([]string, 0)

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(470, &value), rsp, s2)
	log("p2 r:%d", value)
	countreads(value, 2, 30)
	l.Unlock()

	delay()
	l.Lock()
	value = 2
	Put(CreateTuple(24, value), rsp, targets0)
	log("p2 w:%d", value)
	countwrites(value, 2, 2)
	l.Unlock()

	delay()
	l.Lock()
	value = 16
	Put(CreateTuple(241, value), rsp, targets1)
	log("p2 w:%d", value)
	countwrites(value, 2, 16)
	l.Unlock()

	delay()
	l.Lock()
	value = 16
	Put(CreateTuple(248, value), rsp, targets2)
	log("p2 w:%d", value)
	countwrites(value, 2, 16)
	l.Unlock()

	delay()
	l.Lock()
	value = 16
	Put(CreateTuple(250, value), rsp, targets3)
	log("p2 w:%d", value)
	countwrites(value, 2, 16)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(506, &value), rsp, s2)
	log("p2 r:%d", value)
	countreads(value, 2, 32)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(451, &value), rsp, s2)
	log("p2 r:%d", value)
	countreads(value, 2, 29)
	l.Unlock()

	delay()
	l.Lock()
	value = 31
	Put(CreateTuple(495, value), rsp, targets4)
	log("p2 w:%d", value)
	countwrites(value, 2, 31)
	l.Unlock()

	delay()
	l.Lock()
	value = 23
	Put(CreateTuple(353, value), rsp, targets5)
	log("p2 w:%d", value)
	countwrites(value, 2, 23)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(464, &value), rsp, s2)
	log("p2 r:%d", value)
	countreads(value, 2, 29)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(181, &value), rsp, s2)
	log("p2 r:%d", value)
	countreads(value, 2, 12)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(324, &value), rsp, s2)
	log("p2 r:%d", value)
	countreads(value, 2, 21)
	l.Unlock()

	delay()
	l.Lock()
	value = 31
	Put(CreateTuple(490, value), rsp, targets6)
	log("p2 w:%d", value)
	countwrites(value, 2, 31)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(372, &value), rsp, s2)
	log("p2 r:%d", value)
	countreads(value, 2, 24)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(100, &value), rsp, s2)
	log("p2 r:%d", value)
	countreads(value, 2, 7)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(132, &value), rsp, s2)
	log("p2 r:%d", value)
	countreads(value, 2, 9)
	l.Unlock()
}

func p3() {
	targets9 := make([]string, 0)
	targets8 := make([]string, 1)
	targets8[0] = uri[s15]
	targets7 := make([]string, 0)
	targets6 := make([]string, 1)
	targets6[0] = uri[s14]
	targets5 := make([]string, 0)
	targets4 := make([]string, 0)
	targets3 := make([]string, 0)
	targets2 := make([]string, 0)
	targets1 := make([]string, 0)
	targets0 := make([]string, 0)

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(402, &value), rsp, s3)
	log("p3 r:%d", value)
	countreads(value, 3, 26)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(271, &value), rsp, s3)
	log("p3 r:%d", value)
	countreads(value, 3, 17)
	l.Unlock()

	delay()
	l.Lock()
	value = 23
	Put(CreateTuple(366, value), rsp, targets0)
	log("p3 w:%d", value)
	countwrites(value, 3, 23)
	l.Unlock()

	delay()
	l.Lock()
	value = 27
	Put(CreateTuple(424, value), rsp, targets1)
	log("p3 w:%d", value)
	countwrites(value, 3, 27)
	l.Unlock()

	delay()
	l.Lock()
	value = 22
	Put(CreateTuple(341, value), rsp, targets2)
	log("p3 w:%d", value)
	countwrites(value, 3, 22)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(204, &value), rsp, s3)
	log("p3 r:%d", value)
	countreads(value, 3, 13)
	l.Unlock()

	delay()
	l.Lock()
	value = 10
	Put(CreateTuple(153, value), rsp, targets3)
	log("p3 w:%d", value)
	countwrites(value, 3, 10)
	l.Unlock()

	delay()
	l.Lock()
	value = 19
	Put(CreateTuple(297, value), rsp, targets4)
	log("p3 w:%d", value)
	countwrites(value, 3, 19)
	l.Unlock()

	delay()
	l.Lock()
	value = 2
	Put(CreateTuple(17, value), rsp, targets5)
	log("p3 w:%d", value)
	countwrites(value, 3, 2)
	l.Unlock()

	delay()
	l.Lock()
	value = 4
	Put(CreateTuple(58, value), rsp, targets6)
	log("p3 w:%d", value)
	countwrites(value, 3, 4)
	l.Unlock()

	delay()
	l.Lock()
	value = 30
	Put(CreateTuple(475, value), rsp, targets7)
	log("p3 w:%d", value)
	countwrites(value, 3, 30)
	l.Unlock()

	delay()
	l.Lock()
	value = 21
	Put(CreateTuple(334, value), rsp, targets8)
	log("p3 w:%d", value)
	countwrites(value, 3, 21)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(353, &value), rsp, s3)
	log("p3 r:%d", value)
	countreads(value, 3, 23)
	l.Unlock()

	delay()
	l.Lock()
	value = 20
	Put(CreateTuple(319, value), rsp, targets9)
	log("p3 w:%d", value)
	countwrites(value, 3, 20)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(411, &value), rsp, s3)
	log("p3 r:%d", value)
	countreads(value, 3, 26)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(242, &value), rsp, s3)
	log("p3 r:%d", value)
	countreads(value, 3, 16)
	l.Unlock()
}

func p4() {
	targets9 := make([]string, 0)
	targets8 := make([]string, 2)
	targets8[1] = uri[s21]
	targets8[0] = uri[s18]
	targets7 := make([]string, 0)
	targets6 := make([]string, 0)
	targets5 := make([]string, 1)
	targets5[0] = uri[s30]
	targets4 := make([]string, 0)
	targets3 := make([]string, 0)
	targets2 := make([]string, 0)
	targets1 := make([]string, 0)
	targets0 := make([]string, 0)

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(115, &value), rsp, s4)
	log("p4 r:%d", value)
	countreads(value, 4, 8)
	l.Unlock()

	delay()
	l.Lock()
	value = 20
	Put(CreateTuple(313, value), rsp, targets0)
	log("p4 w:%d", value)
	countwrites(value, 4, 20)
	l.Unlock()

	delay()
	l.Lock()
	value = 10
	Put(CreateTuple(151, value), rsp, targets1)
	log("p4 w:%d", value)
	countwrites(value, 4, 10)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(348, &value), rsp, s4)
	log("p4 r:%d", value)
	countreads(value, 4, 22)
	l.Unlock()

	delay()
	l.Lock()
	value = 30
	Put(CreateTuple(465, value), rsp, targets2)
	log("p4 w:%d", value)
	countwrites(value, 4, 30)
	l.Unlock()

	delay()
	l.Lock()
	value = 22
	Put(CreateTuple(341, value), rsp, targets3)
	log("p4 w:%d", value)
	countwrites(value, 4, 22)
	l.Unlock()

	delay()
	l.Lock()
	value = 19
	Put(CreateTuple(299, value), rsp, targets4)
	log("p4 w:%d", value)
	countwrites(value, 4, 19)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(63, &value), rsp, s4)
	log("p4 r:%d", value)
	countreads(value, 4, 4)
	l.Unlock()

	delay()
	l.Lock()
	value = 24
	Put(CreateTuple(371, value), rsp, targets5)
	log("p4 w:%d", value)
	countwrites(value, 4, 24)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(501, &value), rsp, s4)
	log("p4 r:%d", value)
	countreads(value, 4, 32)
	l.Unlock()

	delay()
	l.Lock()
	value = 26
	Put(CreateTuple(408, value), rsp, targets6)
	log("p4 w:%d", value)
	countwrites(value, 4, 26)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(9, &value), rsp, s4)
	log("p4 r:%d", value)
	countreads(value, 4, 1)
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
	value = 13
	Put(CreateTuple(194, value), rsp, targets7)
	log("p4 w:%d", value)
	countwrites(value, 4, 13)
	l.Unlock()

	delay()
	l.Lock()
	value = 10
	Put(CreateTuple(148, value), rsp, targets8)
	log("p4 w:%d", value)
	countwrites(value, 4, 10)
	l.Unlock()

	delay()
	l.Lock()
	value = 27
	Put(CreateTuple(432, value), rsp, targets9)
	log("p4 w:%d", value)
	countwrites(value, 4, 27)
	l.Unlock()
}

func p5() {
	targets12 := make([]string, 3)
	targets12[2] = uri[s28]
	targets12[1] = uri[s12]
	targets12[0] = uri[s7]
	targets11 := make([]string, 1)
	targets11[0] = uri[s20]
	targets10 := make([]string, 0)
	targets9 := make([]string, 0)
	targets8 := make([]string, 0)
	targets7 := make([]string, 0)
	targets6 := make([]string, 0)
	targets5 := make([]string, 0)
	targets4 := make([]string, 0)
	targets3 := make([]string, 2)
	targets3[1] = uri[s19]
	targets3[0] = uri[s15]
	targets2 := make([]string, 0)
	targets1 := make([]string, 2)
	targets1[1] = uri[s19]
	targets1[0] = uri[s13]
	targets0 := make([]string, 1)
	targets0[0] = uri[s5]

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 20
	Put(CreateTuple(314, value), rsp, targets0)
	log("p5 w:%d", value)
	countwrites(value, 5, 20)
	l.Unlock()

	delay()
	l.Lock()
	value = 8
	Put(CreateTuple(124, value), rsp, targets1)
	log("p5 w:%d", value)
	countwrites(value, 5, 8)
	l.Unlock()

	delay()
	l.Lock()
	value = 24
	Put(CreateTuple(373, value), rsp, targets2)
	log("p5 w:%d", value)
	countwrites(value, 5, 24)
	l.Unlock()

	delay()
	l.Lock()
	value = 17
	Put(CreateTuple(269, value), rsp, targets3)
	log("p5 w:%d", value)
	countwrites(value, 5, 17)
	l.Unlock()

	delay()
	l.Lock()
	value = 18
	Put(CreateTuple(286, value), rsp, targets4)
	log("p5 w:%d", value)
	countwrites(value, 5, 18)
	l.Unlock()

	delay()
	l.Lock()
	value = 25
	Put(CreateTuple(389, value), rsp, targets5)
	log("p5 w:%d", value)
	countwrites(value, 5, 25)
	l.Unlock()

	delay()
	l.Lock()
	value = 18
	Put(CreateTuple(280, value), rsp, targets6)
	log("p5 w:%d", value)
	countwrites(value, 5, 18)
	l.Unlock()

	delay()
	l.Lock()
	value = 26
	Put(CreateTuple(409, value), rsp, targets7)
	log("p5 w:%d", value)
	countwrites(value, 5, 26)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(477, &value), rsp, s5)
	log("p5 r:%d", value)
	countreads(value, 5, 30)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(314, &value), rsp, s5)
	log("p5 r:%d", value)
	countreads(value, 5, 20)
	l.Unlock()

	delay()
	l.Lock()
	value = 19
	Put(CreateTuple(290, value), rsp, targets8)
	log("p5 w:%d", value)
	countwrites(value, 5, 19)
	l.Unlock()

	delay()
	l.Lock()
	value = 25
	Put(CreateTuple(387, value), rsp, targets9)
	log("p5 w:%d", value)
	countwrites(value, 5, 25)
	l.Unlock()

	delay()
	l.Lock()
	value = 28
	Put(CreateTuple(438, value), rsp, targets10)
	log("p5 w:%d", value)
	countwrites(value, 5, 28)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(249, &value), rsp, s5)
	log("p5 r:%d", value)
	countreads(value, 5, 16)
	l.Unlock()

	delay()
	l.Lock()
	value = 21
	Put(CreateTuple(328, value), rsp, targets11)
	log("p5 w:%d", value)
	countwrites(value, 5, 21)
	l.Unlock()

	delay()
	l.Lock()
	value = 6
	Put(CreateTuple(88, value), rsp, targets12)
	log("p5 w:%d", value)
	countwrites(value, 5, 6)
	l.Unlock()
}

func p6() {
	targets8 := make([]string, 0)
	targets7 := make([]string, 0)
	targets6 := make([]string, 1)
	targets6[0] = uri[s30]
	targets5 := make([]string, 1)
	targets5[0] = uri[s8]
	targets4 := make([]string, 0)
	targets3 := make([]string, 1)
	targets3[0] = uri[s6]
	targets2 := make([]string, 0)
	targets1 := make([]string, 0)
	targets0 := make([]string, 0)

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 16
	Put(CreateTuple(254, value), rsp, targets0)
	log("p6 w:%d", value)
	countwrites(value, 6, 16)
	l.Unlock()

	delay()
	l.Lock()
	value = 15
	Put(CreateTuple(232, value), rsp, targets1)
	log("p6 w:%d", value)
	countwrites(value, 6, 15)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(388, &value), rsp, s6)
	log("p6 r:%d", value)
	countreads(value, 6, 25)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(52, &value), rsp, s6)
	log("p6 r:%d", value)
	countreads(value, 6, 4)
	l.Unlock()

	delay()
	l.Lock()
	value = 8
	Put(CreateTuple(116, value), rsp, targets2)
	log("p6 w:%d", value)
	countwrites(value, 6, 8)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(141, &value), rsp, s6)
	log("p6 r:%d", value)
	countreads(value, 6, 9)
	l.Unlock()

	delay()
	l.Lock()
	value = 3
	Put(CreateTuple(45, value), rsp, targets3)
	log("p6 w:%d", value)
	countwrites(value, 6, 3)
	l.Unlock()

	delay()
	l.Lock()
	value = 29
	Put(CreateTuple(456, value), rsp, targets4)
	log("p6 w:%d", value)
	countwrites(value, 6, 29)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(240, &value), rsp, s6)
	log("p6 r:%d", value)
	countreads(value, 6, 15)
	l.Unlock()

	delay()
	l.Lock()
	value = 31
	Put(CreateTuple(496, value), rsp, targets5)
	log("p6 w:%d", value)
	countwrites(value, 6, 31)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(45, &value), rsp, s6)
	log("p6 r:%d", value)
	countreads(value, 6, 3)
	l.Unlock()

	delay()
	l.Lock()
	value = 24
	Put(CreateTuple(371, value), rsp, targets6)
	log("p6 w:%d", value)
	countwrites(value, 6, 24)
	l.Unlock()

	delay()
	l.Lock()
	value = 20
	Put(CreateTuple(310, value), rsp, targets7)
	log("p6 w:%d", value)
	countwrites(value, 6, 20)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(135, &value), rsp, s6)
	log("p6 r:%d", value)
	countreads(value, 6, 9)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(459, &value), rsp, s6)
	log("p6 r:%d", value)
	countreads(value, 6, 29)
	l.Unlock()

	delay()
	l.Lock()
	value = 29
	Put(CreateTuple(455, value), rsp, targets8)
	log("p6 w:%d", value)
	countwrites(value, 6, 29)
	l.Unlock()
}

func p7() {
	targets7 := make([]string, 0)
	targets6 := make([]string, 0)
	targets5 := make([]string, 0)
	targets4 := make([]string, 0)
	targets3 := make([]string, 0)
	targets2 := make([]string, 2)
	targets2[1] = uri[s26]
	targets2[0] = uri[s7]
	targets1 := make([]string, 0)
	targets0 := make([]string, 0)

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(88, &value), rsp, s7)
	log("p7 r:%d", value)
	countreads(value, 7, 6)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(415, &value), rsp, s7)
	log("p7 r:%d", value)
	countreads(value, 7, 26)
	l.Unlock()

	delay()
	l.Lock()
	value = 8
	Put(CreateTuple(114, value), rsp, targets0)
	log("p7 w:%d", value)
	countwrites(value, 7, 8)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(106, &value), rsp, s7)
	log("p7 r:%d", value)
	countreads(value, 7, 7)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(410, &value), rsp, s7)
	log("p7 r:%d", value)
	countreads(value, 7, 26)
	l.Unlock()

	delay()
	l.Lock()
	value = 23
	Put(CreateTuple(358, value), rsp, targets1)
	log("p7 w:%d", value)
	countwrites(value, 7, 23)
	l.Unlock()

	delay()
	l.Lock()
	value = 26
	Put(CreateTuple(415, value), rsp, targets2)
	log("p7 w:%d", value)
	countwrites(value, 7, 26)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(511, &value), rsp, s7)
	log("p7 r:%d", value)
	countreads(value, 7, 32)
	l.Unlock()

	delay()
	l.Lock()
	value = 9
	Put(CreateTuple(138, value), rsp, targets3)
	log("p7 w:%d", value)
	countwrites(value, 7, 9)
	l.Unlock()

	delay()
	l.Lock()
	value = 3
	Put(CreateTuple(41, value), rsp, targets4)
	log("p7 w:%d", value)
	countwrites(value, 7, 3)
	l.Unlock()

	delay()
	l.Lock()
	value = 14
	Put(CreateTuple(211, value), rsp, targets5)
	log("p7 w:%d", value)
	countwrites(value, 7, 14)
	l.Unlock()

	delay()
	l.Lock()
	value = 6
	Put(CreateTuple(83, value), rsp, targets6)
	log("p7 w:%d", value)
	countwrites(value, 7, 6)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(410, &value), rsp, s7)
	log("p7 r:%d", value)
	countreads(value, 7, 26)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(336, &value), rsp, s7)
	log("p7 r:%d", value)
	countreads(value, 7, 21)
	l.Unlock()

	delay()
	l.Lock()
	value = 27
	Put(CreateTuple(425, value), rsp, targets7)
	log("p7 w:%d", value)
	countwrites(value, 7, 27)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(227, &value), rsp, s7)
	log("p7 r:%d", value)
	countreads(value, 7, 15)
	l.Unlock()
}

func p8() {
	targets7 := make([]string, 0)
	targets6 := make([]string, 0)
	targets5 := make([]string, 0)
	targets4 := make([]string, 1)
	targets4[0] = uri[s21]
	targets3 := make([]string, 0)
	targets2 := make([]string, 1)
	targets2[0] = uri[s8]
	targets1 := make([]string, 1)
	targets1[0] = uri[s21]
	targets0 := make([]string, 1)
	targets0[0] = uri[s15]

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(442, &value), rsp, s8)
	log("p8 r:%d", value)
	countreads(value, 8, 28)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(36, &value), rsp, s8)
	log("p8 r:%d", value)
	countreads(value, 8, 3)
	l.Unlock()

	delay()
	l.Lock()
	value = 6
	Put(CreateTuple(84, value), rsp, targets0)
	log("p8 w:%d", value)
	countwrites(value, 8, 6)
	l.Unlock()

	delay()
	l.Lock()
	value = 26
	Put(CreateTuple(406, value), rsp, targets1)
	log("p8 w:%d", value)
	countwrites(value, 8, 26)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(26, &value), rsp, s8)
	log("p8 r:%d", value)
	countreads(value, 8, 2)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(77, &value), rsp, s8)
	log("p8 r:%d", value)
	countreads(value, 8, 5)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(103, &value), rsp, s8)
	log("p8 r:%d", value)
	countreads(value, 8, 7)
	l.Unlock()

	delay()
	l.Lock()
	value = 31
	Put(CreateTuple(496, value), rsp, targets2)
	log("p8 w:%d", value)
	countwrites(value, 8, 31)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(496, &value), rsp, s8)
	log("p8 r:%d", value)
	countreads(value, 8, 31)
	l.Unlock()

	delay()
	l.Lock()
	value = 2
	Put(CreateTuple(28, value), rsp, targets3)
	log("p8 w:%d", value)
	countwrites(value, 8, 2)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(181, &value), rsp, s8)
	log("p8 r:%d", value)
	countreads(value, 8, 12)
	l.Unlock()

	delay()
	l.Lock()
	value = 20
	Put(CreateTuple(305, value), rsp, targets4)
	log("p8 w:%d", value)
	countwrites(value, 8, 20)
	l.Unlock()

	delay()
	l.Lock()
	value = 9
	Put(CreateTuple(133, value), rsp, targets5)
	log("p8 w:%d", value)
	countwrites(value, 8, 9)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(421, &value), rsp, s8)
	log("p8 r:%d", value)
	countreads(value, 8, 27)
	l.Unlock()

	delay()
	l.Lock()
	value = 13
	Put(CreateTuple(194, value), rsp, targets6)
	log("p8 w:%d", value)
	countwrites(value, 8, 13)
	l.Unlock()

	delay()
	l.Lock()
	value = 15
	Put(CreateTuple(235, value), rsp, targets7)
	log("p8 w:%d", value)
	countwrites(value, 8, 15)
	l.Unlock()
}

func p9() {
	targets9 := make([]string, 0)
	targets8 := make([]string, 1)
	targets8[0] = uri[s23]
	targets7 := make([]string, 0)
	targets6 := make([]string, 0)
	targets5 := make([]string, 0)
	targets4 := make([]string, 0)
	targets3 := make([]string, 0)
	targets2 := make([]string, 0)
	targets1 := make([]string, 1)
	targets1[0] = uri[s17]
	targets0 := make([]string, 0)

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 18
	Put(CreateTuple(288, value), rsp, targets0)
	log("p9 w:%d", value)
	countwrites(value, 9, 18)
	l.Unlock()

	delay()
	l.Lock()
	value = 23
	Put(CreateTuple(356, value), rsp, targets1)
	log("p9 w:%d", value)
	countwrites(value, 9, 23)
	l.Unlock()

	delay()
	l.Lock()
	value = 32
	Put(CreateTuple(507, value), rsp, targets2)
	log("p9 w:%d", value)
	countwrites(value, 9, 32)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(87, &value), rsp, s9)
	log("p9 r:%d", value)
	countreads(value, 9, 6)
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
	value = 1
	Put(CreateTuple(5, value), rsp, targets3)
	log("p9 w:%d", value)
	countwrites(value, 9, 1)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(27, &value), rsp, s9)
	log("p9 r:%d", value)
	countreads(value, 9, 2)
	l.Unlock()

	delay()
	l.Lock()
	value = 9
	Put(CreateTuple(139, value), rsp, targets4)
	log("p9 w:%d", value)
	countwrites(value, 9, 9)
	l.Unlock()

	delay()
	l.Lock()
	value = 23
	Put(CreateTuple(357, value), rsp, targets5)
	log("p9 w:%d", value)
	countwrites(value, 9, 23)
	l.Unlock()

	delay()
	l.Lock()
	value = 15
	Put(CreateTuple(233, value), rsp, targets6)
	log("p9 w:%d", value)
	countwrites(value, 9, 15)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(70, &value), rsp, s9)
	log("p9 r:%d", value)
	countreads(value, 9, 5)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(67, &value), rsp, s9)
	log("p9 r:%d", value)
	countreads(value, 9, 5)
	l.Unlock()

	delay()
	l.Lock()
	value = 18
	Put(CreateTuple(285, value), rsp, targets7)
	log("p9 w:%d", value)
	countwrites(value, 9, 18)
	l.Unlock()

	delay()
	l.Lock()
	value = 30
	Put(CreateTuple(472, value), rsp, targets8)
	log("p9 w:%d", value)
	countwrites(value, 9, 30)
	l.Unlock()

	delay()
	l.Lock()
	value = 5
	Put(CreateTuple(68, value), rsp, targets9)
	log("p9 w:%d", value)
	countwrites(value, 9, 5)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(264, &value), rsp, s9)
	log("p9 r:%d", value)
	countreads(value, 9, 17)
	l.Unlock()
}

func p10() {
	targets9 := make([]string, 0)
	targets8 := make([]string, 1)
	targets8[0] = uri[s24]
	targets7 := make([]string, 1)
	targets7[0] = uri[s15]
	targets6 := make([]string, 0)
	targets5 := make([]string, 0)
	targets4 := make([]string, 0)
	targets3 := make([]string, 0)
	targets2 := make([]string, 0)
	targets1 := make([]string, 0)
	targets0 := make([]string, 1)
	targets0[0] = uri[s17]

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 11
	Put(CreateTuple(168, value), rsp, targets0)
	log("p10 w:%d", value)
	countwrites(value, 10, 11)
	l.Unlock()

	delay()
	l.Lock()
	value = 17
	Put(CreateTuple(257, value), rsp, targets1)
	log("p10 w:%d", value)
	countwrites(value, 10, 17)
	l.Unlock()

	delay()
	l.Lock()
	value = 14
	Put(CreateTuple(216, value), rsp, targets2)
	log("p10 w:%d", value)
	countwrites(value, 10, 14)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(105, &value), rsp, s10)
	log("p10 r:%d", value)
	countreads(value, 10, 7)
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
	value = 5
	Put(CreateTuple(69, value), rsp, targets3)
	log("p10 w:%d", value)
	countwrites(value, 10, 5)
	l.Unlock()

	delay()
	l.Lock()
	value = 2
	Put(CreateTuple(20, value), rsp, targets4)
	log("p10 w:%d", value)
	countwrites(value, 10, 2)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(271, &value), rsp, s10)
	log("p10 r:%d", value)
	countreads(value, 10, 17)
	l.Unlock()

	delay()
	l.Lock()
	value = 7
	Put(CreateTuple(111, value), rsp, targets5)
	log("p10 w:%d", value)
	countwrites(value, 10, 7)
	l.Unlock()

	delay()
	l.Lock()
	value = 2
	Put(CreateTuple(25, value), rsp, targets6)
	log("p10 w:%d", value)
	countwrites(value, 10, 2)
	l.Unlock()

	delay()
	l.Lock()
	value = 17
	Put(CreateTuple(265, value), rsp, targets7)
	log("p10 w:%d", value)
	countwrites(value, 10, 17)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(333, &value), rsp, s10)
	log("p10 r:%d", value)
	countreads(value, 10, 21)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(332, &value), rsp, s10)
	log("p10 r:%d", value)
	countreads(value, 10, 21)
	l.Unlock()

	delay()
	l.Lock()
	value = 18
	Put(CreateTuple(277, value), rsp, targets8)
	log("p10 w:%d", value)
	countwrites(value, 10, 18)
	l.Unlock()

	delay()
	l.Lock()
	value = 31
	Put(CreateTuple(491, value), rsp, targets9)
	log("p10 w:%d", value)
	countwrites(value, 10, 31)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(463, &value), rsp, s10)
	log("p10 r:%d", value)
	countreads(value, 10, 29)
	l.Unlock()
}

func p11() {
	targets13 := make([]string, 1)
	targets13[0] = uri[s22]
	targets12 := make([]string, 0)
	targets11 := make([]string, 1)
	targets11[0] = uri[s15]
	targets10 := make([]string, 2)
	targets10[1] = uri[s24]
	targets10[0] = uri[s20]
	targets9 := make([]string, 0)
	targets8 := make([]string, 0)
	targets7 := make([]string, 2)
	targets7[1] = uri[s13]
	targets7[0] = uri[s5]
	targets6 := make([]string, 0)
	targets5 := make([]string, 0)
	targets4 := make([]string, 1)
	targets4[0] = uri[s30]
	targets3 := make([]string, 0)
	targets2 := make([]string, 0)
	targets1 := make([]string, 1)
	targets1[0] = uri[s29]
	targets0 := make([]string, 2)
	targets0[1] = uri[s19]
	targets0[0] = uri[s13]

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 8
	Put(CreateTuple(124, value), rsp, targets0)
	log("p11 w:%d", value)
	countwrites(value, 11, 8)
	l.Unlock()

	delay()
	l.Lock()
	value = 16
	Put(CreateTuple(248, value), rsp, targets1)
	log("p11 w:%d", value)
	countwrites(value, 11, 16)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(110, &value), rsp, s11)
	log("p11 r:%d", value)
	countreads(value, 11, 7)
	l.Unlock()

	delay()
	l.Lock()
	value = 12
	Put(CreateTuple(188, value), rsp, targets2)
	log("p11 w:%d", value)
	countwrites(value, 11, 12)
	l.Unlock()

	delay()
	l.Lock()
	value = 28
	Put(CreateTuple(447, value), rsp, targets3)
	log("p11 w:%d", value)
	countwrites(value, 11, 28)
	l.Unlock()

	delay()
	l.Lock()
	value = 32
	Put(CreateTuple(498, value), rsp, targets4)
	log("p11 w:%d", value)
	countwrites(value, 11, 32)
	l.Unlock()

	delay()
	l.Lock()
	value = 14
	Put(CreateTuple(217, value), rsp, targets5)
	log("p11 w:%d", value)
	countwrites(value, 11, 14)
	l.Unlock()

	delay()
	l.Lock()
	value = 3
	Put(CreateTuple(42, value), rsp, targets6)
	log("p11 w:%d", value)
	countwrites(value, 11, 3)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(106, &value), rsp, s11)
	log("p11 r:%d", value)
	countreads(value, 11, 7)
	l.Unlock()

	delay()
	l.Lock()
	value = 30
	Put(CreateTuple(477, value), rsp, targets7)
	log("p11 w:%d", value)
	countwrites(value, 11, 30)
	l.Unlock()

	delay()
	l.Lock()
	value = 5
	Put(CreateTuple(78, value), rsp, targets8)
	log("p11 w:%d", value)
	countwrites(value, 11, 5)
	l.Unlock()

	delay()
	l.Lock()
	value = 18
	Put(CreateTuple(287, value), rsp, targets9)
	log("p11 w:%d", value)
	countwrites(value, 11, 18)
	l.Unlock()

	delay()
	l.Lock()
	value = 5
	Put(CreateTuple(66, value), rsp, targets10)
	log("p11 w:%d", value)
	countwrites(value, 11, 5)
	l.Unlock()

	delay()
	l.Lock()
	value = 13
	Put(CreateTuple(200, value), rsp, targets11)
	log("p11 w:%d", value)
	countwrites(value, 11, 13)
	l.Unlock()

	delay()
	l.Lock()
	value = 14
	Put(CreateTuple(218, value), rsp, targets12)
	log("p11 w:%d", value)
	countwrites(value, 11, 14)
	l.Unlock()

	delay()
	l.Lock()
	value = 19
	Put(CreateTuple(293, value), rsp, targets13)
	log("p11 w:%d", value)
	countwrites(value, 11, 19)
	l.Unlock()
}

func p12() {
	targets10 := make([]string, 0)
	targets9 := make([]string, 0)
	targets8 := make([]string, 0)
	targets7 := make([]string, 0)
	targets6 := make([]string, 0)
	targets5 := make([]string, 0)
	targets4 := make([]string, 0)
	targets3 := make([]string, 0)
	targets2 := make([]string, 2)
	targets2[1] = uri[s20]
	targets2[0] = uri[s12]
	targets1 := make([]string, 1)
	targets1[0] = uri[s28]
	targets0 := make([]string, 0)

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 15
	Put(CreateTuple(233, value), rsp, targets0)
	log("p12 w:%d", value)
	countwrites(value, 12, 15)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(300, &value), rsp, s12)
	log("p12 r:%d", value)
	countreads(value, 12, 19)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(88, &value), rsp, s12)
	log("p12 r:%d", value)
	countreads(value, 12, 6)
	l.Unlock()

	delay()
	l.Lock()
	value = 12
	Put(CreateTuple(177, value), rsp, targets1)
	log("p12 w:%d", value)
	countwrites(value, 12, 12)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(72, &value), rsp, s12)
	log("p12 r:%d", value)
	countreads(value, 12, 5)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(302, &value), rsp, s12)
	log("p12 r:%d", value)
	countreads(value, 12, 19)
	l.Unlock()

	delay()
	l.Lock()
	value = 19
	Put(CreateTuple(300, value), rsp, targets2)
	log("p12 w:%d", value)
	countwrites(value, 12, 19)
	l.Unlock()

	delay()
	l.Lock()
	value = 30
	Put(CreateTuple(480, value), rsp, targets3)
	log("p12 w:%d", value)
	countwrites(value, 12, 30)
	l.Unlock()

	delay()
	l.Lock()
	value = 13
	Put(CreateTuple(202, value), rsp, targets4)
	log("p12 w:%d", value)
	countwrites(value, 12, 13)
	l.Unlock()

	delay()
	l.Lock()
	value = 1
	Put(CreateTuple(3, value), rsp, targets5)
	log("p12 w:%d", value)
	countwrites(value, 12, 1)
	l.Unlock()

	delay()
	l.Lock()
	value = 23
	Put(CreateTuple(368, value), rsp, targets6)
	log("p12 w:%d", value)
	countwrites(value, 12, 23)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(261, &value), rsp, s12)
	log("p12 r:%d", value)
	countreads(value, 12, 17)
	l.Unlock()

	delay()
	l.Lock()
	value = 30
	Put(CreateTuple(467, value), rsp, targets7)
	log("p12 w:%d", value)
	countwrites(value, 12, 30)
	l.Unlock()

	delay()
	l.Lock()
	value = 16
	Put(CreateTuple(247, value), rsp, targets8)
	log("p12 w:%d", value)
	countwrites(value, 12, 16)
	l.Unlock()

	delay()
	l.Lock()
	value = 10
	Put(CreateTuple(158, value), rsp, targets9)
	log("p12 w:%d", value)
	countwrites(value, 12, 10)
	l.Unlock()

	delay()
	l.Lock()
	value = 25
	Put(CreateTuple(396, value), rsp, targets10)
	log("p12 w:%d", value)
	countwrites(value, 12, 25)
	l.Unlock()
}

func p13() {
	targets7 := make([]string, 0)
	targets6 := make([]string, 1)
	targets6[0] = uri[s8]
	targets5 := make([]string, 1)
	targets5[0] = uri[s28]
	targets4 := make([]string, 0)
	targets3 := make([]string, 0)
	targets2 := make([]string, 1)
	targets2[0] = uri[s3]
	targets1 := make([]string, 2)
	targets1[1] = uri[s21]
	targets1[0] = uri[s17]
	targets0 := make([]string, 0)

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 9
	Put(CreateTuple(138, value), rsp, targets0)
	log("p13 w:%d", value)
	countwrites(value, 13, 9)
	l.Unlock()

	delay()
	l.Lock()
	value = 9
	Put(CreateTuple(142, value), rsp, targets1)
	log("p13 w:%d", value)
	countwrites(value, 13, 9)
	l.Unlock()

	delay()
	l.Lock()
	value = 13
	Put(CreateTuple(204, value), rsp, targets2)
	log("p13 w:%d", value)
	countwrites(value, 13, 13)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(477, &value), rsp, s13)
	log("p13 r:%d", value)
	countreads(value, 13, 30)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(53, &value), rsp, s13)
	log("p13 r:%d", value)
	countreads(value, 13, 4)
	l.Unlock()

	delay()
	l.Lock()
	value = 32
	Put(CreateTuple(503, value), rsp, targets3)
	log("p13 w:%d", value)
	countwrites(value, 13, 32)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(419, &value), rsp, s13)
	log("p13 r:%d", value)
	countreads(value, 13, 27)
	l.Unlock()

	delay()
	l.Lock()
	value = 5
	Put(CreateTuple(76, value), rsp, targets4)
	log("p13 w:%d", value)
	countwrites(value, 13, 5)
	l.Unlock()

	delay()
	l.Lock()
	value = 12
	Put(CreateTuple(177, value), rsp, targets5)
	log("p13 w:%d", value)
	countwrites(value, 13, 12)
	l.Unlock()

	delay()
	l.Lock()
	value = 31
	Put(CreateTuple(496, value), rsp, targets6)
	log("p13 w:%d", value)
	countwrites(value, 13, 31)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(292, &value), rsp, s13)
	log("p13 r:%d", value)
	countreads(value, 13, 19)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(124, &value), rsp, s13)
	log("p13 r:%d", value)
	countreads(value, 13, 8)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(466, &value), rsp, s13)
	log("p13 r:%d", value)
	countreads(value, 13, 30)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(309, &value), rsp, s13)
	log("p13 r:%d", value)
	countreads(value, 13, 20)
	l.Unlock()

	delay()
	l.Lock()
	value = 27
	Put(CreateTuple(423, value), rsp, targets7)
	log("p13 w:%d", value)
	countwrites(value, 13, 27)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(495, &value), rsp, s13)
	log("p13 r:%d", value)
	countreads(value, 13, 31)
	l.Unlock()
}

func p14() {
	targets7 := make([]string, 0)
	targets6 := make([]string, 0)
	targets5 := make([]string, 0)
	targets4 := make([]string, 1)
	targets4[0] = uri[s8]
	targets3 := make([]string, 1)
	targets3[0] = uri[s6]
	targets2 := make([]string, 0)
	targets1 := make([]string, 0)
	targets0 := make([]string, 0)

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 25
	Put(CreateTuple(389, value), rsp, targets0)
	log("p14 w:%d", value)
	countwrites(value, 14, 25)
	l.Unlock()

	delay()
	l.Lock()
	value = 18
	Put(CreateTuple(274, value), rsp, targets1)
	log("p14 w:%d", value)
	countwrites(value, 14, 18)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(439, &value), rsp, s14)
	log("p14 r:%d", value)
	countreads(value, 14, 28)
	l.Unlock()

	delay()
	l.Lock()
	value = 13
	Put(CreateTuple(196, value), rsp, targets2)
	log("p14 w:%d", value)
	countwrites(value, 14, 13)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(207, &value), rsp, s14)
	log("p14 r:%d", value)
	countreads(value, 14, 13)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(445, &value), rsp, s14)
	log("p14 r:%d", value)
	countreads(value, 14, 28)
	l.Unlock()

	delay()
	l.Lock()
	value = 15
	Put(CreateTuple(240, value), rsp, targets3)
	log("p14 w:%d", value)
	countwrites(value, 14, 15)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(38, &value), rsp, s14)
	log("p14 r:%d", value)
	countreads(value, 14, 3)
	l.Unlock()

	delay()
	l.Lock()
	value = 3
	Put(CreateTuple(36, value), rsp, targets4)
	log("p14 w:%d", value)
	countwrites(value, 14, 3)
	l.Unlock()

	delay()
	l.Lock()
	value = 16
	Put(CreateTuple(247, value), rsp, targets5)
	log("p14 w:%d", value)
	countwrites(value, 14, 16)
	l.Unlock()

	delay()
	l.Lock()
	value = 9
	Put(CreateTuple(133, value), rsp, targets6)
	log("p14 w:%d", value)
	countwrites(value, 14, 9)
	l.Unlock()

	delay()
	l.Lock()
	value = 16
	Put(CreateTuple(241, value), rsp, targets7)
	log("p14 w:%d", value)
	countwrites(value, 14, 16)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(246, &value), rsp, s14)
	log("p14 r:%d", value)
	countreads(value, 14, 16)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(58, &value), rsp, s14)
	log("p14 r:%d", value)
	countreads(value, 14, 4)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(104, &value), rsp, s14)
	log("p14 r:%d", value)
	countreads(value, 14, 7)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(263, &value), rsp, s14)
	log("p14 r:%d", value)
	countreads(value, 14, 17)
	l.Unlock()
}

func p15() {
	targets8 := make([]string, 1)
	targets8[0] = uri[s14]
	targets7 := make([]string, 0)
	targets6 := make([]string, 0)
	targets5 := make([]string, 1)
	targets5[0] = uri[s6]
	targets4 := make([]string, 0)
	targets3 := make([]string, 0)
	targets2 := make([]string, 0)
	targets1 := make([]string, 0)
	targets0 := make([]string, 0)

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 22
	Put(CreateTuple(340, value), rsp, targets0)
	log("p15 w:%d", value)
	countwrites(value, 15, 22)
	l.Unlock()

	delay()
	l.Lock()
	value = 29
	Put(CreateTuple(458, value), rsp, targets1)
	log("p15 w:%d", value)
	countwrites(value, 15, 29)
	l.Unlock()

	delay()
	l.Lock()
	value = 12
	Put(CreateTuple(189, value), rsp, targets2)
	log("p15 w:%d", value)
	countwrites(value, 15, 12)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(200, &value), rsp, s15)
	log("p15 r:%d", value)
	countreads(value, 15, 13)
	l.Unlock()

	delay()
	l.Lock()
	value = 11
	Put(CreateTuple(163, value), rsp, targets3)
	log("p15 w:%d", value)
	countwrites(value, 15, 11)
	l.Unlock()

	delay()
	l.Lock()
	value = 26
	Put(CreateTuple(401, value), rsp, targets4)
	log("p15 w:%d", value)
	countwrites(value, 15, 26)
	l.Unlock()

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
	QueryP(CreateTuple(243, &value), rsp, s15)
	log("p15 r:%d", value)
	countreads(value, 15, 16)
	l.Unlock()

	delay()
	l.Lock()
	value = 4
	Put(CreateTuple(52, value), rsp, targets5)
	log("p15 w:%d", value)
	countwrites(value, 15, 4)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(269, &value), rsp, s15)
	log("p15 r:%d", value)
	countreads(value, 15, 17)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(334, &value), rsp, s15)
	log("p15 r:%d", value)
	countreads(value, 15, 21)
	l.Unlock()

	delay()
	l.Lock()
	value = 32
	Put(CreateTuple(504, value), rsp, targets6)
	log("p15 w:%d", value)
	countwrites(value, 15, 32)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(84, &value), rsp, s15)
	log("p15 r:%d", value)
	countreads(value, 15, 6)
	l.Unlock()

	delay()
	l.Lock()
	value = 6
	Put(CreateTuple(96, value), rsp, targets7)
	log("p15 w:%d", value)
	countwrites(value, 15, 6)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(463, &value), rsp, s15)
	log("p15 r:%d", value)
	countreads(value, 15, 29)
	l.Unlock()

	delay()
	l.Lock()
	value = 28
	Put(CreateTuple(439, value), rsp, targets8)
	log("p15 w:%d", value)
	countwrites(value, 15, 28)
	l.Unlock()
}

func p16() {
	targets7 := make([]string, 1)
	targets7[0] = uri[s11]
	targets6 := make([]string, 1)
	targets6[0] = uri[s20]
	targets5 := make([]string, 1)
	targets5[0] = uri[s15]
	targets4 := make([]string, 0)
	targets3 := make([]string, 0)
	targets2 := make([]string, 0)
	targets1 := make([]string, 0)
	targets0 := make([]string, 0)

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(199, &value), rsp, s16)
	log("p16 r:%d", value)
	countreads(value, 16, 13)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(44, &value), rsp, s16)
	log("p16 r:%d", value)
	countreads(value, 16, 3)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(73, &value), rsp, s16)
	log("p16 r:%d", value)
	countreads(value, 16, 5)
	l.Unlock()

	delay()
	l.Lock()
	value = 14
	Put(CreateTuple(217, value), rsp, targets0)
	log("p16 w:%d", value)
	countwrites(value, 16, 14)
	l.Unlock()

	delay()
	l.Lock()
	value = 8
	Put(CreateTuple(125, value), rsp, targets1)
	log("p16 w:%d", value)
	countwrites(value, 16, 8)
	l.Unlock()

	delay()
	l.Lock()
	value = 22
	Put(CreateTuple(347, value), rsp, targets2)
	log("p16 w:%d", value)
	countwrites(value, 16, 22)
	l.Unlock()

	delay()
	l.Lock()
	value = 1
	Put(CreateTuple(11, value), rsp, targets3)
	log("p16 w:%d", value)
	countwrites(value, 16, 1)
	l.Unlock()

	delay()
	l.Lock()
	value = 14
	Put(CreateTuple(224, value), rsp, targets4)
	log("p16 w:%d", value)
	countwrites(value, 16, 14)
	l.Unlock()

	delay()
	l.Lock()
	value = 17
	Put(CreateTuple(265, value), rsp, targets5)
	log("p16 w:%d", value)
	countwrites(value, 16, 17)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(155, &value), rsp, s16)
	log("p16 r:%d", value)
	countreads(value, 16, 10)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(365, &value), rsp, s16)
	log("p16 r:%d", value)
	countreads(value, 16, 23)
	l.Unlock()

	delay()
	l.Lock()
	value = 13
	Put(CreateTuple(197, value), rsp, targets6)
	log("p16 w:%d", value)
	countwrites(value, 16, 13)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(2, &value), rsp, s16)
	log("p16 r:%d", value)
	countreads(value, 16, 1)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(80, &value), rsp, s16)
	log("p16 r:%d", value)
	countreads(value, 16, 5)
	l.Unlock()

	delay()
	l.Lock()
	value = 7
	Put(CreateTuple(110, value), rsp, targets7)
	log("p16 w:%d", value)
	countwrites(value, 16, 7)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(231, &value), rsp, s16)
	log("p16 r:%d", value)
	countreads(value, 16, 15)
	l.Unlock()
}

func p17() {
	targets7 := make([]string, 0)
	targets6 := make([]string, 0)
	targets5 := make([]string, 2)
	targets5[1] = uri[s19]
	targets5[0] = uri[s13]
	targets4 := make([]string, 0)
	targets3 := make([]string, 0)
	targets2 := make([]string, 1)
	targets2[0] = uri[s8]
	targets1 := make([]string, 0)
	targets0 := make([]string, 1)
	targets0[0] = uri[s21]

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(481, &value), rsp, s17)
	log("p17 r:%d", value)
	countreads(value, 17, 31)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(142, &value), rsp, s17)
	log("p17 r:%d", value)
	countreads(value, 17, 9)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(295, &value), rsp, s17)
	log("p17 r:%d", value)
	countreads(value, 17, 19)
	l.Unlock()

	delay()
	l.Lock()
	value = 20
	Put(CreateTuple(305, value), rsp, targets0)
	log("p17 w:%d", value)
	countwrites(value, 17, 20)
	l.Unlock()

	delay()
	l.Lock()
	value = 30
	Put(CreateTuple(476, value), rsp, targets1)
	log("p17 w:%d", value)
	countwrites(value, 17, 30)
	l.Unlock()

	delay()
	l.Lock()
	value = 3
	Put(CreateTuple(36, value), rsp, targets2)
	log("p17 w:%d", value)
	countwrites(value, 17, 3)
	l.Unlock()

	delay()
	l.Lock()
	value = 7
	Put(CreateTuple(101, value), rsp, targets3)
	log("p17 w:%d", value)
	countwrites(value, 17, 7)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(356, &value), rsp, s17)
	log("p17 r:%d", value)
	countreads(value, 17, 23)
	l.Unlock()

	delay()
	l.Lock()
	value = 30
	Put(CreateTuple(478, value), rsp, targets4)
	log("p17 w:%d", value)
	countwrites(value, 17, 30)
	l.Unlock()

	delay()
	l.Lock()
	value = 8
	Put(CreateTuple(124, value), rsp, targets5)
	log("p17 w:%d", value)
	countwrites(value, 17, 8)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(10, &value), rsp, s17)
	log("p17 r:%d", value)
	countreads(value, 17, 1)
	l.Unlock()

	delay()
	l.Lock()
	value = 22
	Put(CreateTuple(344, value), rsp, targets6)
	log("p17 w:%d", value)
	countwrites(value, 17, 22)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(168, &value), rsp, s17)
	log("p17 r:%d", value)
	countreads(value, 17, 11)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(343, &value), rsp, s17)
	log("p17 r:%d", value)
	countreads(value, 17, 22)
	l.Unlock()

	delay()
	l.Lock()
	value = 9
	Put(CreateTuple(133, value), rsp, targets7)
	log("p17 w:%d", value)
	countwrites(value, 17, 9)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(444, &value), rsp, s17)
	log("p17 r:%d", value)
	countreads(value, 17, 28)
	l.Unlock()
}

func p18() {
	targets8 := make([]string, 0)
	targets7 := make([]string, 2)
	targets7[1] = uri[s26]
	targets7[0] = uri[s8]
	targets6 := make([]string, 0)
	targets5 := make([]string, 1)
	targets5[0] = uri[s30]
	targets4 := make([]string, 1)
	targets4[0] = uri[s1]
	targets3 := make([]string, 0)
	targets2 := make([]string, 0)
	targets1 := make([]string, 0)
	targets0 := make([]string, 1)
	targets0[0] = uri[s24]

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(148, &value), rsp, s18)
	log("p18 r:%d", value)
	countreads(value, 18, 10)
	l.Unlock()

	delay()
	l.Lock()
	value = 8
	Put(CreateTuple(117, value), rsp, targets0)
	log("p18 w:%d", value)
	countwrites(value, 18, 8)
	l.Unlock()

	delay()
	l.Lock()
	value = 15
	Put(CreateTuple(232, value), rsp, targets1)
	log("p18 w:%d", value)
	countwrites(value, 18, 15)
	l.Unlock()

	delay()
	l.Lock()
	value = 16
	Put(CreateTuple(247, value), rsp, targets2)
	log("p18 w:%d", value)
	countwrites(value, 18, 16)
	l.Unlock()

	delay()
	l.Lock()
	value = 9
	Put(CreateTuple(144, value), rsp, targets3)
	log("p18 w:%d", value)
	countwrites(value, 18, 9)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(54, &value), rsp, s18)
	log("p18 r:%d", value)
	countreads(value, 18, 4)
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
	value = 4
	Put(CreateTuple(55, value), rsp, targets4)
	log("p18 w:%d", value)
	countwrites(value, 18, 4)
	l.Unlock()

	delay()
	l.Lock()
	value = 32
	Put(CreateTuple(498, value), rsp, targets5)
	log("p18 w:%d", value)
	countwrites(value, 18, 32)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(420, &value), rsp, s18)
	log("p18 r:%d", value)
	countreads(value, 18, 27)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(298, &value), rsp, s18)
	log("p18 r:%d", value)
	countreads(value, 18, 19)
	l.Unlock()

	delay()
	l.Lock()
	value = 19
	Put(CreateTuple(297, value), rsp, targets6)
	log("p18 w:%d", value)
	countwrites(value, 18, 19)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(345, &value), rsp, s18)
	log("p18 r:%d", value)
	countreads(value, 18, 22)
	l.Unlock()

	delay()
	l.Lock()
	value = 27
	Put(CreateTuple(421, value), rsp, targets7)
	log("p18 w:%d", value)
	countwrites(value, 18, 27)
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
	value = 25
	Put(CreateTuple(389, value), rsp, targets8)
	log("p18 w:%d", value)
	countwrites(value, 18, 25)
	l.Unlock()
}

func p19() {
	targets6 := make([]string, 0)
	targets5 := make([]string, 0)
	targets4 := make([]string, 0)
	targets3 := make([]string, 0)
	targets2 := make([]string, 1)
	targets2[0] = uri[s28]
	targets1 := make([]string, 0)
	targets0 := make([]string, 0)

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 20
	Put(CreateTuple(313, value), rsp, targets0)
	log("p19 w:%d", value)
	countwrites(value, 19, 20)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(276, &value), rsp, s19)
	log("p19 r:%d", value)
	countreads(value, 19, 18)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(500, &value), rsp, s19)
	log("p19 r:%d", value)
	countreads(value, 19, 32)
	l.Unlock()

	delay()
	l.Lock()
	value = 29
	Put(CreateTuple(457, value), rsp, targets1)
	log("p19 w:%d", value)
	countwrites(value, 19, 29)
	l.Unlock()

	delay()
	l.Lock()
	value = 12
	Put(CreateTuple(177, value), rsp, targets2)
	log("p19 w:%d", value)
	countwrites(value, 19, 12)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(208, &value), rsp, s19)
	log("p19 r:%d", value)
	countreads(value, 19, 13)
	l.Unlock()

	delay()
	l.Lock()
	value = 24
	Put(CreateTuple(375, value), rsp, targets3)
	log("p19 w:%d", value)
	countwrites(value, 19, 24)
	l.Unlock()

	delay()
	l.Lock()
	value = 12
	Put(CreateTuple(182, value), rsp, targets4)
	log("p19 w:%d", value)
	countwrites(value, 19, 12)
	l.Unlock()

	delay()
	l.Lock()
	value = 25
	Put(CreateTuple(390, value), rsp, targets5)
	log("p19 w:%d", value)
	countwrites(value, 19, 25)
	l.Unlock()

	delay()
	l.Lock()
	value = 29
	Put(CreateTuple(449, value), rsp, targets6)
	log("p19 w:%d", value)
	countwrites(value, 19, 29)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(337, &value), rsp, s19)
	log("p19 r:%d", value)
	countreads(value, 19, 22)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(370, &value), rsp, s19)
	log("p19 r:%d", value)
	countreads(value, 19, 24)
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
	QueryP(CreateTuple(124, &value), rsp, s19)
	log("p19 r:%d", value)
	countreads(value, 19, 8)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(60, &value), rsp, s19)
	log("p19 r:%d", value)
	countreads(value, 19, 4)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(170, &value), rsp, s19)
	log("p19 r:%d", value)
	countreads(value, 19, 11)
	l.Unlock()
}

func p20() {
	targets7 := make([]string, 0)
	targets6 := make([]string, 1)
	targets6[0] = uri[s8]
	targets5 := make([]string, 1)
	targets5[0] = uri[s27]
	targets4 := make([]string, 0)
	targets3 := make([]string, 0)
	targets2 := make([]string, 0)
	targets1 := make([]string, 3)
	targets1[2] = uri[s30]
	targets1[1] = uri[s20]
	targets1[0] = uri[s3]
	targets0 := make([]string, 0)

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(164, &value), rsp, s20)
	log("p20 r:%d", value)
	countreads(value, 20, 11)
	l.Unlock()

	delay()
	l.Lock()
	value = 9
	Put(CreateTuple(134, value), rsp, targets0)
	log("p20 w:%d", value)
	countwrites(value, 20, 9)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(300, &value), rsp, s20)
	log("p20 r:%d", value)
	countreads(value, 20, 19)
	l.Unlock()

	delay()
	l.Lock()
	value = 26
	Put(CreateTuple(402, value), rsp, targets1)
	log("p20 w:%d", value)
	countwrites(value, 20, 26)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(484, &value), rsp, s20)
	log("p20 r:%d", value)
	countreads(value, 20, 31)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(392, &value), rsp, s20)
	log("p20 r:%d", value)
	countreads(value, 20, 25)
	l.Unlock()

	delay()
	l.Lock()
	value = 12
	Put(CreateTuple(182, value), rsp, targets2)
	log("p20 w:%d", value)
	countwrites(value, 20, 12)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(66, &value), rsp, s20)
	log("p20 r:%d", value)
	countreads(value, 20, 5)
	l.Unlock()

	delay()
	l.Lock()
	value = 20
	Put(CreateTuple(319, value), rsp, targets3)
	log("p20 w:%d", value)
	countwrites(value, 20, 20)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(328, &value), rsp, s20)
	log("p20 r:%d", value)
	countreads(value, 20, 21)
	l.Unlock()

	delay()
	l.Lock()
	value = 2
	Put(CreateTuple(28, value), rsp, targets4)
	log("p20 w:%d", value)
	countwrites(value, 20, 2)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(197, &value), rsp, s20)
	log("p20 r:%d", value)
	countreads(value, 20, 13)
	l.Unlock()

	delay()
	l.Lock()
	value = 14
	Put(CreateTuple(209, value), rsp, targets5)
	log("p20 w:%d", value)
	countwrites(value, 20, 14)
	l.Unlock()

	delay()
	l.Lock()
	value = 31
	Put(CreateTuple(496, value), rsp, targets6)
	log("p20 w:%d", value)
	countwrites(value, 20, 31)
	l.Unlock()

	delay()
	l.Lock()
	value = 29
	Put(CreateTuple(455, value), rsp, targets7)
	log("p20 w:%d", value)
	countwrites(value, 20, 29)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(402, &value), rsp, s20)
	log("p20 r:%d", value)
	countreads(value, 20, 26)
	l.Unlock()
}

func p21() {
	targets7 := make([]string, 0)
	targets6 := make([]string, 0)
	targets5 := make([]string, 1)
	targets5[0] = uri[s32]
	targets4 := make([]string, 0)
	targets3 := make([]string, 0)
	targets2 := make([]string, 0)
	targets1 := make([]string, 0)
	targets0 := make([]string, 3)
	targets0[2] = uri[s28]
	targets0[1] = uri[s12]
	targets0[0] = uri[s7]

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(305, &value), rsp, s21)
	log("p21 r:%d", value)
	countreads(value, 21, 20)
	l.Unlock()

	delay()
	l.Lock()
	value = 6
	Put(CreateTuple(88, value), rsp, targets0)
	log("p21 w:%d", value)
	countwrites(value, 21, 6)
	l.Unlock()

	delay()
	l.Lock()
	value = 13
	Put(CreateTuple(198, value), rsp, targets1)
	log("p21 w:%d", value)
	countwrites(value, 21, 13)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(176, &value), rsp, s21)
	log("p21 r:%d", value)
	countreads(value, 21, 11)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(406, &value), rsp, s21)
	log("p21 r:%d", value)
	countreads(value, 21, 26)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(148, &value), rsp, s21)
	log("p21 r:%d", value)
	countreads(value, 21, 10)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(142, &value), rsp, s21)
	log("p21 r:%d", value)
	countreads(value, 21, 9)
	l.Unlock()

	delay()
	l.Lock()
	value = 12
	Put(CreateTuple(191, value), rsp, targets2)
	log("p21 w:%d", value)
	countwrites(value, 21, 12)
	l.Unlock()

	delay()
	l.Lock()
	value = 16
	Put(CreateTuple(241, value), rsp, targets3)
	log("p21 w:%d", value)
	countwrites(value, 21, 16)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(223, &value), rsp, s21)
	log("p21 r:%d", value)
	countreads(value, 21, 14)
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
	QueryP(CreateTuple(395, &value), rsp, s21)
	log("p21 r:%d", value)
	countreads(value, 21, 25)
	l.Unlock()

	delay()
	l.Lock()
	value = 28
	Put(CreateTuple(443, value), rsp, targets4)
	log("p21 w:%d", value)
	countwrites(value, 21, 28)
	l.Unlock()

	delay()
	l.Lock()
	value = 3
	Put(CreateTuple(40, value), rsp, targets5)
	log("p21 w:%d", value)
	countwrites(value, 21, 3)
	l.Unlock()

	delay()
	l.Lock()
	value = 28
	Put(CreateTuple(438, value), rsp, targets6)
	log("p21 w:%d", value)
	countwrites(value, 21, 28)
	l.Unlock()

	delay()
	l.Lock()
	value = 21
	Put(CreateTuple(329, value), rsp, targets7)
	log("p21 w:%d", value)
	countwrites(value, 21, 21)
	l.Unlock()
}

func p22() {
	targets7 := make([]string, 2)
	targets7[1] = uri[s26]
	targets7[0] = uri[s16]
	targets6 := make([]string, 0)
	targets5 := make([]string, 0)
	targets4 := make([]string, 0)
	targets3 := make([]string, 1)
	targets3[0] = uri[s2]
	targets2 := make([]string, 0)
	targets1 := make([]string, 0)
	targets0 := make([]string, 3)
	targets0[2] = uri[s28]
	targets0[1] = uri[s12]
	targets0[0] = uri[s7]

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(112, &value), rsp, s22)
	log("p22 r:%d", value)
	countreads(value, 22, 7)
	l.Unlock()

	delay()
	l.Lock()
	value = 6
	Put(CreateTuple(88, value), rsp, targets0)
	log("p22 w:%d", value)
	countwrites(value, 22, 6)
	l.Unlock()

	delay()
	l.Lock()
	value = 9
	Put(CreateTuple(134, value), rsp, targets1)
	log("p22 w:%d", value)
	countwrites(value, 22, 9)
	l.Unlock()

	delay()
	l.Lock()
	value = 11
	Put(CreateTuple(171, value), rsp, targets2)
	log("p22 w:%d", value)
	countwrites(value, 22, 11)
	l.Unlock()

	delay()
	l.Lock()
	value = 9
	Put(CreateTuple(132, value), rsp, targets3)
	log("p22 w:%d", value)
	countwrites(value, 22, 9)
	l.Unlock()

	delay()
	l.Lock()
	value = 32
	Put(CreateTuple(502, value), rsp, targets4)
	log("p22 w:%d", value)
	countwrites(value, 22, 32)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(437, &value), rsp, s22)
	log("p22 r:%d", value)
	countreads(value, 22, 28)
	l.Unlock()

	delay()
	l.Lock()
	value = 6
	Put(CreateTuple(93, value), rsp, targets5)
	log("p22 w:%d", value)
	countwrites(value, 22, 6)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(293, &value), rsp, s22)
	log("p22 r:%d", value)
	countreads(value, 22, 19)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(1, &value), rsp, s22)
	log("p22 r:%d", value)
	countreads(value, 22, 1)
	l.Unlock()

	delay()
	l.Lock()
	value = 14
	Put(CreateTuple(221, value), rsp, targets6)
	log("p22 w:%d", value)
	countwrites(value, 22, 14)
	l.Unlock()

	delay()
	l.Lock()
	value = 15
	Put(CreateTuple(231, value), rsp, targets7)
	log("p22 w:%d", value)
	countwrites(value, 22, 15)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(326, &value), rsp, s22)
	log("p22 r:%d", value)
	countreads(value, 22, 21)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(260, &value), rsp, s22)
	log("p22 r:%d", value)
	countreads(value, 22, 17)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(361, &value), rsp, s22)
	log("p22 r:%d", value)
	countreads(value, 22, 23)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(511, &value), rsp, s22)
	log("p22 r:%d", value)
	countreads(value, 22, 32)
	l.Unlock()
}

func p23() {
	targets10 := make([]string, 0)
	targets9 := make([]string, 0)
	targets8 := make([]string, 1)
	targets8[0] = uri[s5]
	targets7 := make([]string, 2)
	targets7[1] = uri[s23]
	targets7[0] = uri[s4]
	targets6 := make([]string, 0)
	targets5 := make([]string, 0)
	targets4 := make([]string, 0)
	targets3 := make([]string, 1)
	targets3[0] = uri[s5]
	targets2 := make([]string, 0)
	targets1 := make([]string, 2)
	targets1[1] = uri[s20]
	targets1[0] = uri[s12]
	targets0 := make([]string, 1)
	targets0[0] = uri[s7]

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(172, &value), rsp, s23)
	log("p23 r:%d", value)
	countreads(value, 23, 11)
	l.Unlock()

	delay()
	l.Lock()
	value = 21
	Put(CreateTuple(336, value), rsp, targets0)
	log("p23 w:%d", value)
	countwrites(value, 23, 21)
	l.Unlock()

	delay()
	l.Lock()
	value = 19
	Put(CreateTuple(300, value), rsp, targets1)
	log("p23 w:%d", value)
	countwrites(value, 23, 19)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(472, &value), rsp, s23)
	log("p23 r:%d", value)
	countreads(value, 23, 30)
	l.Unlock()

	delay()
	l.Lock()
	value = 31
	Put(CreateTuple(491, value), rsp, targets2)
	log("p23 w:%d", value)
	countwrites(value, 23, 31)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(453, &value), rsp, s23)
	log("p23 r:%d", value)
	countreads(value, 23, 29)
	l.Unlock()

	delay()
	l.Lock()
	value = 16
	Put(CreateTuple(249, value), rsp, targets3)
	log("p23 w:%d", value)
	countwrites(value, 23, 16)
	l.Unlock()

	delay()
	l.Lock()
	value = 19
	Put(CreateTuple(290, value), rsp, targets4)
	log("p23 w:%d", value)
	countwrites(value, 23, 19)
	l.Unlock()

	delay()
	l.Lock()
	value = 6
	Put(CreateTuple(95, value), rsp, targets5)
	log("p23 w:%d", value)
	countwrites(value, 23, 6)
	l.Unlock()

	delay()
	l.Lock()
	value = 11
	Put(CreateTuple(166, value), rsp, targets6)
	log("p23 w:%d", value)
	countwrites(value, 23, 11)
	l.Unlock()

	delay()
	l.Lock()
	value = 11
	Put(CreateTuple(172, value), rsp, targets7)
	log("p23 w:%d", value)
	countwrites(value, 23, 11)
	l.Unlock()

	delay()
	l.Lock()
	value = 16
	Put(CreateTuple(249, value), rsp, targets8)
	log("p23 w:%d", value)
	countwrites(value, 23, 16)
	l.Unlock()

	delay()
	l.Lock()
	value = 18
	Put(CreateTuple(278, value), rsp, targets9)
	log("p23 w:%d", value)
	countwrites(value, 23, 18)
	l.Unlock()

	delay()
	l.Lock()
	value = 3
	Put(CreateTuple(37, value), rsp, targets10)
	log("p23 w:%d", value)
	countwrites(value, 23, 3)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(220, &value), rsp, s23)
	log("p23 r:%d", value)
	countreads(value, 23, 14)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(90, &value), rsp, s23)
	log("p23 r:%d", value)
	countreads(value, 23, 6)
	l.Unlock()
}

func p24() {
	targets6 := make([]string, 1)
	targets6[0] = uri[s9]
	targets5 := make([]string, 1)
	targets5[0] = uri[s18]
	targets4 := make([]string, 1)
	targets4[0] = uri[s1]
	targets3 := make([]string, 1)
	targets3[0] = uri[s22]
	targets2 := make([]string, 1)
	targets2[0] = uri[s14]
	targets1 := make([]string, 1)
	targets1[0] = uri[s4]
	targets0 := make([]string, 0)

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(89, &value), rsp, s24)
	log("p24 r:%d", value)
	countreads(value, 24, 6)
	l.Unlock()

	delay()
	l.Lock()
	value = 25
	Put(CreateTuple(393, value), rsp, targets0)
	log("p24 w:%d", value)
	countwrites(value, 24, 25)
	l.Unlock()

	delay()
	l.Lock()
	value = 22
	Put(CreateTuple(348, value), rsp, targets1)
	log("p24 w:%d", value)
	countwrites(value, 24, 22)
	l.Unlock()

	delay()
	l.Lock()
	value = 28
	Put(CreateTuple(439, value), rsp, targets2)
	log("p24 w:%d", value)
	countwrites(value, 24, 28)
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
	QueryP(CreateTuple(66, &value), rsp, s24)
	log("p24 r:%d", value)
	countreads(value, 24, 5)
	l.Unlock()

	delay()
	l.Lock()
	value = 28
	Put(CreateTuple(437, value), rsp, targets3)
	log("p24 w:%d", value)
	countwrites(value, 24, 28)
	l.Unlock()

	delay()
	l.Lock()
	value = 4
	Put(CreateTuple(55, value), rsp, targets4)
	log("p24 w:%d", value)
	countwrites(value, 24, 4)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(117, &value), rsp, s24)
	log("p24 r:%d", value)
	countreads(value, 24, 8)
	l.Unlock()

	delay()
	l.Lock()
	value = 30
	Put(CreateTuple(473, value), rsp, targets5)
	log("p24 w:%d", value)
	countwrites(value, 24, 30)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(436, &value), rsp, s24)
	log("p24 r:%d", value)
	countreads(value, 24, 28)
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
	QueryP(CreateTuple(277, &value), rsp, s24)
	log("p24 r:%d", value)
	countreads(value, 24, 18)
	l.Unlock()

	delay()
	l.Lock()
	value = 2
	Put(CreateTuple(27, value), rsp, targets6)
	log("p24 w:%d", value)
	countwrites(value, 24, 2)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(183, &value), rsp, s24)
	log("p24 r:%d", value)
	countreads(value, 24, 12)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(292, &value), rsp, s24)
	log("p24 r:%d", value)
	countreads(value, 24, 19)
	l.Unlock()
}

func p25() {
	targets10 := make([]string, 0)
	targets9 := make([]string, 1)
	targets9[0] = uri[s2]
	targets8 := make([]string, 1)
	targets8[0] = uri[s29]
	targets7 := make([]string, 1)
	targets7[0] = uri[s28]
	targets6 := make([]string, 0)
	targets5 := make([]string, 0)
	targets4 := make([]string, 0)
	targets3 := make([]string, 0)
	targets2 := make([]string, 0)
	targets1 := make([]string, 0)
	targets0 := make([]string, 0)

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 11
	Put(CreateTuple(161, value), rsp, targets0)
	log("p25 w:%d", value)
	countwrites(value, 25, 11)
	l.Unlock()

	delay()
	l.Lock()
	value = 8
	Put(CreateTuple(126, value), rsp, targets1)
	log("p25 w:%d", value)
	countwrites(value, 25, 8)
	l.Unlock()

	delay()
	l.Lock()
	value = 11
	Put(CreateTuple(166, value), rsp, targets2)
	log("p25 w:%d", value)
	countwrites(value, 25, 11)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(75, &value), rsp, s25)
	log("p25 r:%d", value)
	countreads(value, 25, 5)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(264, &value), rsp, s25)
	log("p25 r:%d", value)
	countreads(value, 25, 17)
	l.Unlock()

	delay()
	l.Lock()
	value = 28
	Put(CreateTuple(435, value), rsp, targets3)
	log("p25 w:%d", value)
	countwrites(value, 25, 28)
	l.Unlock()

	delay()
	l.Lock()
	value = 30
	Put(CreateTuple(468, value), rsp, targets4)
	log("p25 w:%d", value)
	countwrites(value, 25, 30)
	l.Unlock()

	delay()
	l.Lock()
	value = 18
	Put(CreateTuple(284, value), rsp, targets5)
	log("p25 w:%d", value)
	countwrites(value, 25, 18)
	l.Unlock()

	delay()
	l.Lock()
	value = 20
	Put(CreateTuple(308, value), rsp, targets6)
	log("p25 w:%d", value)
	countwrites(value, 25, 20)
	l.Unlock()

	delay()
	l.Lock()
	value = 4
	Put(CreateTuple(51, value), rsp, targets7)
	log("p25 w:%d", value)
	countwrites(value, 25, 4)
	l.Unlock()

	delay()
	l.Lock()
	value = 21
	Put(CreateTuple(321, value), rsp, targets8)
	log("p25 w:%d", value)
	countwrites(value, 25, 21)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(146, &value), rsp, s25)
	log("p25 r:%d", value)
	countreads(value, 25, 10)
	l.Unlock()

	delay()
	l.Lock()
	value = 21
	Put(CreateTuple(324, value), rsp, targets9)
	log("p25 w:%d", value)
	countwrites(value, 25, 21)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(145, &value), rsp, s25)
	log("p25 r:%d", value)
	countreads(value, 25, 10)
	l.Unlock()

	delay()
	l.Lock()
	value = 26
	Put(CreateTuple(407, value), rsp, targets10)
	log("p25 w:%d", value)
	countwrites(value, 25, 26)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(419, &value), rsp, s25)
	log("p25 r:%d", value)
	countreads(value, 25, 27)
	l.Unlock()
}

func p26() {
	targets6 := make([]string, 0)
	targets5 := make([]string, 1)
	targets5[0] = uri[s2]
	targets4 := make([]string, 0)
	targets3 := make([]string, 1)
	targets3[0] = uri[s1]
	targets2 := make([]string, 0)
	targets1 := make([]string, 1)
	targets1[0] = uri[s12]
	targets0 := make([]string, 0)

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(9, &value), rsp, s26)
	log("p26 r:%d", value)
	countreads(value, 26, 1)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(231, &value), rsp, s26)
	log("p26 r:%d", value)
	countreads(value, 26, 15)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(362, &value), rsp, s26)
	log("p26 r:%d", value)
	countreads(value, 26, 23)
	l.Unlock()

	delay()
	l.Lock()
	value = 14
	Put(CreateTuple(214, value), rsp, targets0)
	log("p26 w:%d", value)
	countwrites(value, 26, 14)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(156, &value), rsp, s26)
	log("p26 r:%d", value)
	countreads(value, 26, 10)
	l.Unlock()

	delay()
	l.Lock()
	value = 17
	Put(CreateTuple(261, value), rsp, targets1)
	log("p26 w:%d", value)
	countwrites(value, 26, 17)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(415, &value), rsp, s26)
	log("p26 r:%d", value)
	countreads(value, 26, 26)
	l.Unlock()

	delay()
	l.Lock()
	value = 29
	Put(CreateTuple(454, value), rsp, targets2)
	log("p26 w:%d", value)
	countwrites(value, 26, 29)
	l.Unlock()

	delay()
	l.Lock()
	value = 12
	Put(CreateTuple(180, value), rsp, targets3)
	log("p26 w:%d", value)
	countwrites(value, 26, 12)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(380, &value), rsp, s26)
	log("p26 r:%d", value)
	countreads(value, 26, 24)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(421, &value), rsp, s26)
	log("p26 r:%d", value)
	countreads(value, 26, 27)
	l.Unlock()

	delay()
	l.Lock()
	value = 9
	Put(CreateTuple(138, value), rsp, targets4)
	log("p26 w:%d", value)
	countwrites(value, 26, 9)
	l.Unlock()

	delay()
	l.Lock()
	value = 29
	Put(CreateTuple(451, value), rsp, targets5)
	log("p26 w:%d", value)
	countwrites(value, 26, 29)
	l.Unlock()

	delay()
	l.Lock()
	value = 3
	Put(CreateTuple(41, value), rsp, targets6)
	log("p26 w:%d", value)
	countwrites(value, 26, 3)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(311, &value), rsp, s26)
	log("p26 r:%d", value)
	countreads(value, 26, 20)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(295, &value), rsp, s26)
	log("p26 r:%d", value)
	countreads(value, 26, 19)
	l.Unlock()
}

func p27() {
	targets12 := make([]string, 0)
	targets11 := make([]string, 0)
	targets10 := make([]string, 1)
	targets10[0] = uri[s1]
	targets9 := make([]string, 0)
	targets8 := make([]string, 0)
	targets7 := make([]string, 0)
	targets6 := make([]string, 0)
	targets5 := make([]string, 0)
	targets4 := make([]string, 0)
	targets3 := make([]string, 1)
	targets3[0] = uri[s6]
	targets2 := make([]string, 0)
	targets1 := make([]string, 1)
	targets1[0] = uri[s21]
	targets0 := make([]string, 0)

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(270, &value), rsp, s27)
	log("p27 r:%d", value)
	countreads(value, 27, 17)
	l.Unlock()

	delay()
	l.Lock()
	value = 14
	Put(CreateTuple(214, value), rsp, targets0)
	log("p27 w:%d", value)
	countwrites(value, 27, 14)
	l.Unlock()

	delay()
	l.Lock()
	value = 11
	Put(CreateTuple(176, value), rsp, targets1)
	log("p27 w:%d", value)
	countwrites(value, 27, 11)
	l.Unlock()

	delay()
	l.Lock()
	value = 11
	Put(CreateTuple(166, value), rsp, targets2)
	log("p27 w:%d", value)
	countwrites(value, 27, 11)
	l.Unlock()

	delay()
	l.Lock()
	value = 29
	Put(CreateTuple(459, value), rsp, targets3)
	log("p27 w:%d", value)
	countwrites(value, 27, 29)
	l.Unlock()

	delay()
	l.Lock()
	value = 23
	Put(CreateTuple(363, value), rsp, targets4)
	log("p27 w:%d", value)
	countwrites(value, 27, 23)
	l.Unlock()

	delay()
	l.Lock()
	value = 28
	Put(CreateTuple(443, value), rsp, targets5)
	log("p27 w:%d", value)
	countwrites(value, 27, 28)
	l.Unlock()

	delay()
	l.Lock()
	value = 5
	Put(CreateTuple(74, value), rsp, targets6)
	log("p27 w:%d", value)
	countwrites(value, 27, 5)
	l.Unlock()

	delay()
	l.Lock()
	value = 8
	Put(CreateTuple(123, value), rsp, targets7)
	log("p27 w:%d", value)
	countwrites(value, 27, 8)
	l.Unlock()

	delay()
	l.Lock()
	value = 2
	Put(CreateTuple(23, value), rsp, targets8)
	log("p27 w:%d", value)
	countwrites(value, 27, 2)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(209, &value), rsp, s27)
	log("p27 r:%d", value)
	countreads(value, 27, 14)
	l.Unlock()

	delay()
	l.Lock()
	value = 24
	Put(CreateTuple(381, value), rsp, targets9)
	log("p27 w:%d", value)
	countwrites(value, 27, 24)
	l.Unlock()

	delay()
	l.Lock()
	value = 12
	Put(CreateTuple(180, value), rsp, targets10)
	log("p27 w:%d", value)
	countwrites(value, 27, 12)
	l.Unlock()

	delay()
	l.Lock()
	value = 12
	Put(CreateTuple(187, value), rsp, targets11)
	log("p27 w:%d", value)
	countwrites(value, 27, 12)
	l.Unlock()

	delay()
	l.Lock()
	value = 27
	Put(CreateTuple(425, value), rsp, targets12)
	log("p27 w:%d", value)
	countwrites(value, 27, 27)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(149, &value), rsp, s27)
	log("p27 r:%d", value)
	countreads(value, 27, 10)
	l.Unlock()
}

func p28() {
	targets8 := make([]string, 1)
	targets8[0] = uri[s20]
	targets7 := make([]string, 1)
	targets7[0] = uri[s15]
	targets6 := make([]string, 0)
	targets5 := make([]string, 0)
	targets4 := make([]string, 0)
	targets3 := make([]string, 0)
	targets2 := make([]string, 2)
	targets2[1] = uri[s25]
	targets2[0] = uri[s9]
	targets1 := make([]string, 1)
	targets1[0] = uri[s4]
	targets0 := make([]string, 2)
	targets0[1] = uri[s8]
	targets0[0] = uri[s2]

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 12
	Put(CreateTuple(181, value), rsp, targets0)
	log("p28 w:%d", value)
	countwrites(value, 28, 12)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(51, &value), rsp, s28)
	log("p28 r:%d", value)
	countreads(value, 28, 4)
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
	value = 8
	Put(CreateTuple(115, value), rsp, targets1)
	log("p28 w:%d", value)
	countwrites(value, 28, 8)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(85, &value), rsp, s28)
	log("p28 r:%d", value)
	countreads(value, 28, 6)
	l.Unlock()

	delay()
	l.Lock()
	value = 17
	Put(CreateTuple(264, value), rsp, targets2)
	log("p28 w:%d", value)
	countwrites(value, 28, 17)
	l.Unlock()

	delay()
	l.Lock()
	value = 1
	Put(CreateTuple(7, value), rsp, targets3)
	log("p28 w:%d", value)
	countwrites(value, 28, 1)
	l.Unlock()

	delay()
	l.Lock()
	value = 22
	Put(CreateTuple(341, value), rsp, targets4)
	log("p28 w:%d", value)
	countwrites(value, 28, 22)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(54, &value), rsp, s28)
	log("p28 r:%d", value)
	countreads(value, 28, 4)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(294, &value), rsp, s28)
	log("p28 r:%d", value)
	countreads(value, 28, 19)
	l.Unlock()

	delay()
	l.Lock()
	value = 32
	Put(CreateTuple(507, value), rsp, targets5)
	log("p28 w:%d", value)
	countwrites(value, 28, 32)
	l.Unlock()

	delay()
	l.Lock()
	value = 31
	Put(CreateTuple(483, value), rsp, targets6)
	log("p28 w:%d", value)
	countwrites(value, 28, 31)
	l.Unlock()

	delay()
	l.Lock()
	value = 13
	Put(CreateTuple(200, value), rsp, targets7)
	log("p28 w:%d", value)
	countwrites(value, 28, 13)
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
	value = 13
	Put(CreateTuple(197, value), rsp, targets8)
	log("p28 w:%d", value)
	countwrites(value, 28, 13)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(177, &value), rsp, s28)
	log("p28 r:%d", value)
	countreads(value, 28, 12)
	l.Unlock()
}

func p29() {
	targets5 := make([]string, 0)
	targets4 := make([]string, 0)
	targets3 := make([]string, 0)
	targets2 := make([]string, 0)
	targets1 := make([]string, 0)
	targets0 := make([]string, 1)
	targets0[0] = uri[s15]

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 17
	Put(CreateTuple(265, value), rsp, targets0)
	log("p29 w:%d", value)
	countwrites(value, 29, 17)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(248, &value), rsp, s29)
	log("p29 r:%d", value)
	countreads(value, 29, 16)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(317, &value), rsp, s29)
	log("p29 r:%d", value)
	countreads(value, 29, 20)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(226, &value), rsp, s29)
	log("p29 r:%d", value)
	countreads(value, 29, 15)
	l.Unlock()

	delay()
	l.Lock()
	value = 1
	Put(CreateTuple(4, value), rsp, targets1)
	log("p29 w:%d", value)
	countwrites(value, 29, 1)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(94, &value), rsp, s29)
	log("p29 r:%d", value)
	countreads(value, 29, 6)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(186, &value), rsp, s29)
	log("p29 r:%d", value)
	countreads(value, 29, 12)
	l.Unlock()

	delay()
	l.Lock()
	value = 3
	Put(CreateTuple(48, value), rsp, targets2)
	log("p29 w:%d", value)
	countwrites(value, 29, 3)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(53, &value), rsp, s29)
	log("p29 r:%d", value)
	countreads(value, 29, 4)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(203, &value), rsp, s29)
	log("p29 r:%d", value)
	countreads(value, 29, 13)
	l.Unlock()

	delay()
	l.Lock()
	value = 23
	Put(CreateTuple(358, value), rsp, targets3)
	log("p29 w:%d", value)
	countwrites(value, 29, 23)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(369, &value), rsp, s29)
	log("p29 r:%d", value)
	countreads(value, 29, 24)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(301, &value), rsp, s29)
	log("p29 r:%d", value)
	countreads(value, 29, 19)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(321, &value), rsp, s29)
	log("p29 r:%d", value)
	countreads(value, 29, 21)
	l.Unlock()

	delay()
	l.Lock()
	value = 3
	Put(CreateTuple(41, value), rsp, targets4)
	log("p29 w:%d", value)
	countwrites(value, 29, 3)
	l.Unlock()

	delay()
	l.Lock()
	value = 20
	Put(CreateTuple(320, value), rsp, targets5)
	log("p29 w:%d", value)
	countwrites(value, 29, 20)
	l.Unlock()
}

func p30() {
	targets5 := make([]string, 0)
	targets4 := make([]string, 0)
	targets3 := make([]string, 0)
	targets2 := make([]string, 0)
	targets1 := make([]string, 1)
	targets1[0] = uri[s24]
	targets0 := make([]string, 0)

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 32
	Put(CreateTuple(502, value), rsp, targets0)
	log("p30 w:%d", value)
	countwrites(value, 30, 32)
	l.Unlock()

	delay()
	l.Lock()
	value = 24
	Put(CreateTuple(377, value), rsp, targets1)
	log("p30 w:%d", value)
	countwrites(value, 30, 24)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(371, &value), rsp, s30)
	log("p30 r:%d", value)
	countreads(value, 30, 24)
	l.Unlock()

	delay()
	l.Lock()
	value = 22
	Put(CreateTuple(339, value), rsp, targets2)
	log("p30 w:%d", value)
	countwrites(value, 30, 22)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(402, &value), rsp, s30)
	log("p30 r:%d", value)
	countreads(value, 30, 26)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(85, &value), rsp, s30)
	log("p30 r:%d", value)
	countreads(value, 30, 6)
	l.Unlock()

	delay()
	l.Lock()
	value = 16
	Put(CreateTuple(256, value), rsp, targets3)
	log("p30 w:%d", value)
	countwrites(value, 30, 16)
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
	QueryP(CreateTuple(382, &value), rsp, s30)
	log("p30 r:%d", value)
	countreads(value, 30, 24)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(440, &value), rsp, s30)
	log("p30 r:%d", value)
	countreads(value, 30, 28)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(186, &value), rsp, s30)
	log("p30 r:%d", value)
	countreads(value, 30, 12)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(317, &value), rsp, s30)
	log("p30 r:%d", value)
	countreads(value, 30, 20)
	l.Unlock()

	delay()
	l.Lock()
	value = 23
	Put(CreateTuple(363, value), rsp, targets4)
	log("p30 w:%d", value)
	countwrites(value, 30, 23)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(498, &value), rsp, s30)
	log("p30 r:%d", value)
	countreads(value, 30, 32)
	l.Unlock()

	delay()
	l.Lock()
	value = 16
	Put(CreateTuple(254, value), rsp, targets5)
	log("p30 w:%d", value)
	countwrites(value, 30, 16)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(97, &value), rsp, s30)
	log("p30 r:%d", value)
	countreads(value, 30, 7)
	l.Unlock()
}

func p31() {
	targets13 := make([]string, 1)
	targets13[0] = uri[s16]
	targets12 := make([]string, 0)
	targets11 := make([]string, 0)
	targets10 := make([]string, 0)
	targets9 := make([]string, 0)
	targets8 := make([]string, 1)
	targets8[0] = uri[s17]
	targets7 := make([]string, 0)
	targets6 := make([]string, 0)
	targets5 := make([]string, 1)
	targets5[0] = uri[s16]
	targets4 := make([]string, 0)
	targets3 := make([]string, 0)
	targets2 := make([]string, 2)
	targets2[1] = uri[s25]
	targets2[0] = uri[s9]
	targets1 := make([]string, 0)
	targets0 := make([]string, 1)
	targets0[0] = uri[s18]

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 19
	Put(CreateTuple(298, value), rsp, targets0)
	log("p31 w:%d", value)
	countwrites(value, 31, 19)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(353, &value), rsp, s31)
	log("p31 r:%d", value)
	countreads(value, 31, 23)
	l.Unlock()

	delay()
	l.Lock()
	value = 31
	Put(CreateTuple(486, value), rsp, targets1)
	log("p31 w:%d", value)
	countwrites(value, 31, 31)
	l.Unlock()

	delay()
	l.Lock()
	value = 17
	Put(CreateTuple(264, value), rsp, targets2)
	log("p31 w:%d", value)
	countwrites(value, 31, 17)
	l.Unlock()

	delay()
	l.Lock()
	value = 26
	Put(CreateTuple(403, value), rsp, targets3)
	log("p31 w:%d", value)
	countwrites(value, 31, 26)
	l.Unlock()

	delay()
	l.Lock()
	value = 14
	Put(CreateTuple(216, value), rsp, targets4)
	log("p31 w:%d", value)
	countwrites(value, 31, 14)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(230, &value), rsp, s31)
	log("p31 r:%d", value)
	countreads(value, 31, 15)
	l.Unlock()

	delay()
	l.Lock()
	value = 3
	Put(CreateTuple(44, value), rsp, targets5)
	log("p31 w:%d", value)
	countwrites(value, 31, 3)
	l.Unlock()

	delay()
	l.Lock()
	value = 3
	Put(CreateTuple(33, value), rsp, targets6)
	log("p31 w:%d", value)
	countwrites(value, 31, 3)
	l.Unlock()

	delay()
	l.Lock()
	value = 27
	Put(CreateTuple(422, value), rsp, targets7)
	log("p31 w:%d", value)
	countwrites(value, 31, 27)
	l.Unlock()

	delay()
	l.Lock()
	value = 23
	Put(CreateTuple(356, value), rsp, targets8)
	log("p31 w:%d", value)
	countwrites(value, 31, 23)
	l.Unlock()

	delay()
	l.Lock()
	value = 28
	Put(CreateTuple(434, value), rsp, targets9)
	log("p31 w:%d", value)
	countwrites(value, 31, 28)
	l.Unlock()

	delay()
	l.Lock()
	value = 2
	Put(CreateTuple(28, value), rsp, targets10)
	log("p31 w:%d", value)
	countwrites(value, 31, 2)
	l.Unlock()

	delay()
	l.Lock()
	value = 24
	Put(CreateTuple(376, value), rsp, targets11)
	log("p31 w:%d", value)
	countwrites(value, 31, 24)
	l.Unlock()

	delay()
	l.Lock()
	value = 20
	Put(CreateTuple(318, value), rsp, targets12)
	log("p31 w:%d", value)
	countwrites(value, 31, 20)
	l.Unlock()

	delay()
	l.Lock()
	value = 10
	Put(CreateTuple(155, value), rsp, targets13)
	log("p31 w:%d", value)
	countwrites(value, 31, 10)
	l.Unlock()
}

func p32() {
	targets8 := make([]string, 0)
	targets7 := make([]string, 0)
	targets6 := make([]string, 0)
	targets5 := make([]string, 0)
	targets4 := make([]string, 0)
	targets3 := make([]string, 1)
	targets3[0] = uri[s20]
	targets2 := make([]string, 0)
	targets1 := make([]string, 0)
	targets0 := make([]string, 0)

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 28
	Put(CreateTuple(441, value), rsp, targets0)
	log("p32 w:%d", value)
	countwrites(value, 32, 28)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(317, &value), rsp, s32)
	log("p32 r:%d", value)
	countreads(value, 32, 20)
	l.Unlock()

	delay()
	l.Lock()
	value = 31
	Put(CreateTuple(494, value), rsp, targets1)
	log("p32 w:%d", value)
	countwrites(value, 32, 31)
	l.Unlock()

	delay()
	l.Lock()
	value = 31
	Put(CreateTuple(483, value), rsp, targets2)
	log("p32 w:%d", value)
	countwrites(value, 32, 31)
	l.Unlock()

	delay()
	l.Lock()
	value = 11
	Put(CreateTuple(164, value), rsp, targets3)
	log("p32 w:%d", value)
	countwrites(value, 32, 11)
	l.Unlock()

	delay()
	l.Lock()
	value = 18
	Put(CreateTuple(282, value), rsp, targets4)
	log("p32 w:%d", value)
	countwrites(value, 32, 18)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(230, &value), rsp, s32)
	log("p32 r:%d", value)
	countreads(value, 32, 15)
	l.Unlock()

	delay()
	l.Lock()
	value = 22
	Put(CreateTuple(346, value), rsp, targets5)
	log("p32 w:%d", value)
	countwrites(value, 32, 22)
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
	value = 29
	Put(CreateTuple(456, value), rsp, targets6)
	log("p32 w:%d", value)
	countwrites(value, 32, 29)
	l.Unlock()

	delay()
	l.Lock()
	value = 23
	Put(CreateTuple(360, value), rsp, targets7)
	log("p32 w:%d", value)
	countwrites(value, 32, 23)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(80, &value), rsp, s32)
	log("p32 r:%d", value)
	countreads(value, 32, 5)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(40, &value), rsp, s32)
	log("p32 r:%d", value)
	countreads(value, 32, 3)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(12, &value), rsp, s32)
	log("p32 r:%d", value)
	countreads(value, 32, 1)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(266, &value), rsp, s32)
	log("p32 r:%d", value)
	countreads(value, 32, 17)
	l.Unlock()

	delay()
	l.Lock()
	value = 8
	Put(CreateTuple(114, value), rsp, targets8)
	log("p32 w:%d", value)
	countwrites(value, 32, 8)
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
