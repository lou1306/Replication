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
	targets0 := make([]string, 2)
	targets0[1] = uri[s30]
	targets0[0] = uri[s8]

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(226, &value), rsp, s1)
	log("p1 r:%d", value)
	countreads(value, 1, 29)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(249, &value), rsp, s1)
	log("p1 r:%d", value)
	countreads(value, 1, 32)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(176, &value), rsp, s1)
	log("p1 r:%d", value)
	countreads(value, 1, 22)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(71, &value), rsp, s1)
	log("p1 r:%d", value)
	countreads(value, 1, 9)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(131, &value), rsp, s1)
	log("p1 r:%d", value)
	countreads(value, 1, 17)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(242, &value), rsp, s1)
	log("p1 r:%d", value)
	countreads(value, 1, 31)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(16, &value), rsp, s1)
	log("p1 r:%d", value)
	countreads(value, 1, 2)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(45, &value), rsp, s1)
	log("p1 r:%d", value)
	countreads(value, 1, 6)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(186, &value), rsp, s1)
	log("p1 r:%d", value)
	countreads(value, 1, 24)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(187, &value), rsp, s1)
	log("p1 r:%d", value)
	countreads(value, 1, 24)
	l.Unlock()

	delay()
	l.Lock()
	value = 14
	Put(CreateTuple(109, value), rsp, targets0)
	log("p1 w:%d", value)
	countwrites(value, 1, 14)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(221, &value), rsp, s1)
	log("p1 r:%d", value)
	countreads(value, 1, 28)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(155, &value), rsp, s1)
	log("p1 r:%d", value)
	countreads(value, 1, 20)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(207, &value), rsp, s1)
	log("p1 r:%d", value)
	countreads(value, 1, 26)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(3, &value), rsp, s1)
	log("p1 r:%d", value)
	countreads(value, 1, 1)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(203, &value), rsp, s1)
	log("p1 r:%d", value)
	countreads(value, 1, 26)
	l.Unlock()
}

func p2() {
	targets2 := make([]string, 0)
	targets1 := make([]string, 0)
	targets0 := make([]string, 0)
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(237, &value), rsp, s2)
	log("p2 r:%d", value)
	countreads(value, 2, 30)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(186, &value), rsp, s2)
	log("p2 r:%d", value)
	countreads(value, 2, 24)
	l.Unlock()

	delay()
	l.Lock()
	value = 13
	Put(CreateTuple(103, value), rsp, targets0)
	log("p2 w:%d", value)
	countwrites(value, 2, 13)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(57, &value), rsp, s2)
	log("p2 r:%d", value)
	countreads(value, 2, 8)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(31, &value), rsp, s2)
	log("p2 r:%d", value)
	countreads(value, 2, 4)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(127, &value), rsp, s2)
	log("p2 r:%d", value)
	countreads(value, 2, 16)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(62, &value), rsp, s2)
	log("p2 r:%d", value)
	countreads(value, 2, 8)
	l.Unlock()

	delay()
	l.Lock()
	value = 10
	Put(CreateTuple(75, value), rsp, targets1)
	log("p2 w:%d", value)
	countwrites(value, 2, 10)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(183, &value), rsp, s2)
	log("p2 r:%d", value)
	countreads(value, 2, 23)
	l.Unlock()

	delay()
	l.Lock()
	value = 16
	Put(CreateTuple(121, value), rsp, targets2)
	log("p2 w:%d", value)
	countwrites(value, 2, 16)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(10, &value), rsp, s2)
	log("p2 r:%d", value)
	countreads(value, 2, 2)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(148, &value), rsp, s2)
	log("p2 r:%d", value)
	countreads(value, 2, 19)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(96, &value), rsp, s2)
	log("p2 r:%d", value)
	countreads(value, 2, 12)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(110, &value), rsp, s2)
	log("p2 r:%d", value)
	countreads(value, 2, 14)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(3, &value), rsp, s2)
	log("p2 r:%d", value)
	countreads(value, 2, 1)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(192, &value), rsp, s2)
	log("p2 r:%d", value)
	countreads(value, 2, 24)
	l.Unlock()
}

func p3() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(49, &value), rsp, s3)
	log("p3 r:%d", value)
	countreads(value, 3, 7)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(112, &value), rsp, s3)
	log("p3 r:%d", value)
	countreads(value, 3, 14)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(130, &value), rsp, s3)
	log("p3 r:%d", value)
	countreads(value, 3, 17)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(99, &value), rsp, s3)
	log("p3 r:%d", value)
	countreads(value, 3, 13)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(187, &value), rsp, s3)
	log("p3 r:%d", value)
	countreads(value, 3, 24)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(67, &value), rsp, s3)
	log("p3 r:%d", value)
	countreads(value, 3, 9)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(153, &value), rsp, s3)
	log("p3 r:%d", value)
	countreads(value, 3, 20)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(123, &value), rsp, s3)
	log("p3 r:%d", value)
	countreads(value, 3, 16)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(195, &value), rsp, s3)
	log("p3 r:%d", value)
	countreads(value, 3, 25)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(239, &value), rsp, s3)
	log("p3 r:%d", value)
	countreads(value, 3, 30)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(139, &value), rsp, s3)
	log("p3 r:%d", value)
	countreads(value, 3, 18)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(13, &value), rsp, s3)
	log("p3 r:%d", value)
	countreads(value, 3, 2)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(14, &value), rsp, s3)
	log("p3 r:%d", value)
	countreads(value, 3, 2)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(23, &value), rsp, s3)
	log("p3 r:%d", value)
	countreads(value, 3, 3)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(39, &value), rsp, s3)
	log("p3 r:%d", value)
	countreads(value, 3, 5)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(250, &value), rsp, s3)
	log("p3 r:%d", value)
	countreads(value, 3, 32)
	l.Unlock()
}

func p4() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(147, &value), rsp, s4)
	log("p4 r:%d", value)
	countreads(value, 4, 19)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(221, &value), rsp, s4)
	log("p4 r:%d", value)
	countreads(value, 4, 28)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(77, &value), rsp, s4)
	log("p4 r:%d", value)
	countreads(value, 4, 10)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(143, &value), rsp, s4)
	log("p4 r:%d", value)
	countreads(value, 4, 18)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(93, &value), rsp, s4)
	log("p4 r:%d", value)
	countreads(value, 4, 12)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(60, &value), rsp, s4)
	log("p4 r:%d", value)
	countreads(value, 4, 8)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(63, &value), rsp, s4)
	log("p4 r:%d", value)
	countreads(value, 4, 8)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(149, &value), rsp, s4)
	log("p4 r:%d", value)
	countreads(value, 4, 19)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(62, &value), rsp, s4)
	log("p4 r:%d", value)
	countreads(value, 4, 8)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(11, &value), rsp, s4)
	log("p4 r:%d", value)
	countreads(value, 4, 2)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(54, &value), rsp, s4)
	log("p4 r:%d", value)
	countreads(value, 4, 7)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(168, &value), rsp, s4)
	log("p4 r:%d", value)
	countreads(value, 4, 21)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(180, &value), rsp, s4)
	log("p4 r:%d", value)
	countreads(value, 4, 23)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(44, &value), rsp, s4)
	log("p4 r:%d", value)
	countreads(value, 4, 6)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(48, &value), rsp, s4)
	log("p4 r:%d", value)
	countreads(value, 4, 6)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(216, &value), rsp, s4)
	log("p4 r:%d", value)
	countreads(value, 4, 27)
	l.Unlock()
}

func p5() {
	targets1 := make([]string, 3)
	targets1[2] = uri[s24]
	targets1[1] = uri[s11]
	targets1[0] = uri[s6]
	targets0 := make([]string, 0)

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(84, &value), rsp, s5)
	log("p5 r:%d", value)
	countreads(value, 5, 11)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(127, &value), rsp, s5)
	log("p5 r:%d", value)
	countreads(value, 5, 16)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(117, &value), rsp, s5)
	log("p5 r:%d", value)
	countreads(value, 5, 15)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(65, &value), rsp, s5)
	log("p5 r:%d", value)
	countreads(value, 5, 9)
	l.Unlock()

	delay()
	l.Lock()
	value = 6
	Put(CreateTuple(43, value), rsp, targets0)
	log("p5 w:%d", value)
	countwrites(value, 5, 6)
	l.Unlock()

	delay()
	l.Lock()
	value = 32
	Put(CreateTuple(256, value), rsp, targets1)
	log("p5 w:%d", value)
	countwrites(value, 5, 32)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(49, &value), rsp, s5)
	log("p5 r:%d", value)
	countreads(value, 5, 7)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(112, &value), rsp, s5)
	log("p5 r:%d", value)
	countreads(value, 5, 14)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(166, &value), rsp, s5)
	log("p5 r:%d", value)
	countreads(value, 5, 21)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(253, &value), rsp, s5)
	log("p5 r:%d", value)
	countreads(value, 5, 32)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(236, &value), rsp, s5)
	log("p5 r:%d", value)
	countreads(value, 5, 30)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(160, &value), rsp, s5)
	log("p5 r:%d", value)
	countreads(value, 5, 20)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(7, &value), rsp, s5)
	log("p5 r:%d", value)
	countreads(value, 5, 1)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(179, &value), rsp, s5)
	log("p5 r:%d", value)
	countreads(value, 5, 23)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(184, &value), rsp, s5)
	log("p5 r:%d", value)
	countreads(value, 5, 23)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(12, &value), rsp, s5)
	log("p5 r:%d", value)
	countreads(value, 5, 2)
	l.Unlock()
}

func p6() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(199, &value), rsp, s6)
	log("p6 r:%d", value)
	countreads(value, 6, 25)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(135, &value), rsp, s6)
	log("p6 r:%d", value)
	countreads(value, 6, 17)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(158, &value), rsp, s6)
	log("p6 r:%d", value)
	countreads(value, 6, 20)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(154, &value), rsp, s6)
	log("p6 r:%d", value)
	countreads(value, 6, 20)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(8, &value), rsp, s6)
	log("p6 r:%d", value)
	countreads(value, 6, 1)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(256, &value), rsp, s6)
	log("p6 r:%d", value)
	countreads(value, 6, 32)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(68, &value), rsp, s6)
	log("p6 r:%d", value)
	countreads(value, 6, 9)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(221, &value), rsp, s6)
	log("p6 r:%d", value)
	countreads(value, 6, 28)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(165, &value), rsp, s6)
	log("p6 r:%d", value)
	countreads(value, 6, 21)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(73, &value), rsp, s6)
	log("p6 r:%d", value)
	countreads(value, 6, 10)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(229, &value), rsp, s6)
	log("p6 r:%d", value)
	countreads(value, 6, 29)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(189, &value), rsp, s6)
	log("p6 r:%d", value)
	countreads(value, 6, 24)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(138, &value), rsp, s6)
	log("p6 r:%d", value)
	countreads(value, 6, 18)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(245, &value), rsp, s6)
	log("p6 r:%d", value)
	countreads(value, 6, 31)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(29, &value), rsp, s6)
	log("p6 r:%d", value)
	countreads(value, 6, 4)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(205, &value), rsp, s6)
	log("p6 r:%d", value)
	countreads(value, 6, 26)
	l.Unlock()
}

func p7() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(77, &value), rsp, s7)
	log("p7 r:%d", value)
	countreads(value, 7, 10)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(102, &value), rsp, s7)
	log("p7 r:%d", value)
	countreads(value, 7, 13)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(46, &value), rsp, s7)
	log("p7 r:%d", value)
	countreads(value, 7, 6)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(88, &value), rsp, s7)
	log("p7 r:%d", value)
	countreads(value, 7, 11)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(31, &value), rsp, s7)
	log("p7 r:%d", value)
	countreads(value, 7, 4)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(139, &value), rsp, s7)
	log("p7 r:%d", value)
	countreads(value, 7, 18)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(172, &value), rsp, s7)
	log("p7 r:%d", value)
	countreads(value, 7, 22)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(174, &value), rsp, s7)
	log("p7 r:%d", value)
	countreads(value, 7, 22)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(200, &value), rsp, s7)
	log("p7 r:%d", value)
	countreads(value, 7, 25)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(185, &value), rsp, s7)
	log("p7 r:%d", value)
	countreads(value, 7, 24)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(138, &value), rsp, s7)
	log("p7 r:%d", value)
	countreads(value, 7, 18)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(32, &value), rsp, s7)
	log("p7 r:%d", value)
	countreads(value, 7, 4)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(135, &value), rsp, s7)
	log("p7 r:%d", value)
	countreads(value, 7, 17)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(205, &value), rsp, s7)
	log("p7 r:%d", value)
	countreads(value, 7, 26)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(26, &value), rsp, s7)
	log("p7 r:%d", value)
	countreads(value, 7, 4)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(187, &value), rsp, s7)
	log("p7 r:%d", value)
	countreads(value, 7, 24)
	l.Unlock()
}

func p8() {
	targets0 := make([]string, 1)
	targets0[0] = uri[s8]

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(204, &value), rsp, s8)
	log("p8 r:%d", value)
	countreads(value, 8, 26)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(35, &value), rsp, s8)
	log("p8 r:%d", value)
	countreads(value, 8, 5)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(237, &value), rsp, s8)
	log("p8 r:%d", value)
	countreads(value, 8, 30)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(163, &value), rsp, s8)
	log("p8 r:%d", value)
	countreads(value, 8, 21)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(253, &value), rsp, s8)
	log("p8 r:%d", value)
	countreads(value, 8, 32)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(220, &value), rsp, s8)
	log("p8 r:%d", value)
	countreads(value, 8, 28)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(37, &value), rsp, s8)
	log("p8 r:%d", value)
	countreads(value, 8, 5)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(135, &value), rsp, s8)
	log("p8 r:%d", value)
	countreads(value, 8, 17)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(106, &value), rsp, s8)
	log("p8 r:%d", value)
	countreads(value, 8, 14)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(22, &value), rsp, s8)
	log("p8 r:%d", value)
	countreads(value, 8, 3)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(85, &value), rsp, s8)
	log("p8 r:%d", value)
	countreads(value, 8, 11)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(109, &value), rsp, s8)
	log("p8 r:%d", value)
	countreads(value, 8, 14)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(132, &value), rsp, s8)
	log("p8 r:%d", value)
	countreads(value, 8, 17)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(149, &value), rsp, s8)
	log("p8 r:%d", value)
	countreads(value, 8, 19)
	l.Unlock()

	delay()
	l.Lock()
	value = 3
	Put(CreateTuple(22, value), rsp, targets0)
	log("p8 w:%d", value)
	countwrites(value, 8, 3)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(93, &value), rsp, s8)
	log("p8 r:%d", value)
	countreads(value, 8, 12)
	l.Unlock()
}

func p9() {
	targets1 := make([]string, 2)
	targets1[1] = uri[s10]
	targets1[0] = uri[s7]
	targets0 := make([]string, 2)
	targets0[1] = uri[s29]
	targets0[0] = uri[s6]

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(16, &value), rsp, s9)
	log("p9 r:%d", value)
	countreads(value, 9, 2)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(241, &value), rsp, s9)
	log("p9 r:%d", value)
	countreads(value, 9, 31)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(19, &value), rsp, s9)
	log("p9 r:%d", value)
	countreads(value, 9, 3)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(108, &value), rsp, s9)
	log("p9 r:%d", value)
	countreads(value, 9, 14)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(191, &value), rsp, s9)
	log("p9 r:%d", value)
	countreads(value, 9, 24)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(229, &value), rsp, s9)
	log("p9 r:%d", value)
	countreads(value, 9, 29)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(129, &value), rsp, s9)
	log("p9 r:%d", value)
	countreads(value, 9, 17)
	l.Unlock()

	delay()
	l.Lock()
	value = 1
	Put(CreateTuple(8, value), rsp, targets0)
	log("p9 w:%d", value)
	countwrites(value, 9, 1)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(47, &value), rsp, s9)
	log("p9 r:%d", value)
	countreads(value, 9, 6)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(24, &value), rsp, s9)
	log("p9 r:%d", value)
	countreads(value, 9, 3)
	l.Unlock()

	delay()
	l.Lock()
	value = 13
	Put(CreateTuple(102, value), rsp, targets1)
	log("p9 w:%d", value)
	countwrites(value, 9, 13)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(178, &value), rsp, s9)
	log("p9 r:%d", value)
	countreads(value, 9, 23)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(170, &value), rsp, s9)
	log("p9 r:%d", value)
	countreads(value, 9, 22)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(117, &value), rsp, s9)
	log("p9 r:%d", value)
	countreads(value, 9, 15)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(86, &value), rsp, s9)
	log("p9 r:%d", value)
	countreads(value, 9, 11)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(47, &value), rsp, s9)
	log("p9 r:%d", value)
	countreads(value, 9, 6)
	l.Unlock()
}

func p10() {
	targets1 := make([]string, 1)
	targets1[0] = uri[s3]
	targets0 := make([]string, 3)
	targets0[2] = uri[s12]
	targets0[1] = uri[s7]
	targets0[0] = uri[s3]

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(245, &value), rsp, s10)
	log("p10 r:%d", value)
	countreads(value, 10, 31)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(163, &value), rsp, s10)
	log("p10 r:%d", value)
	countreads(value, 10, 21)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(30, &value), rsp, s10)
	log("p10 r:%d", value)
	countreads(value, 10, 4)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(102, &value), rsp, s10)
	log("p10 r:%d", value)
	countreads(value, 10, 13)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(225, &value), rsp, s10)
	log("p10 r:%d", value)
	countreads(value, 10, 29)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(255, &value), rsp, s10)
	log("p10 r:%d", value)
	countreads(value, 10, 32)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(135, &value), rsp, s10)
	log("p10 r:%d", value)
	countreads(value, 10, 17)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(183, &value), rsp, s10)
	log("p10 r:%d", value)
	countreads(value, 10, 23)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(185, &value), rsp, s10)
	log("p10 r:%d", value)
	countreads(value, 10, 24)
	l.Unlock()

	delay()
	l.Lock()
	value = 18
	Put(CreateTuple(139, value), rsp, targets0)
	log("p10 w:%d", value)
	countwrites(value, 10, 18)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(143, &value), rsp, s10)
	log("p10 r:%d", value)
	countreads(value, 10, 18)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(236, &value), rsp, s10)
	log("p10 r:%d", value)
	countreads(value, 10, 30)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(182, &value), rsp, s10)
	log("p10 r:%d", value)
	countreads(value, 10, 23)
	l.Unlock()

	delay()
	l.Lock()
	value = 30
	Put(CreateTuple(239, value), rsp, targets1)
	log("p10 w:%d", value)
	countwrites(value, 10, 30)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(13, &value), rsp, s10)
	log("p10 r:%d", value)
	countreads(value, 10, 2)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(216, &value), rsp, s10)
	log("p10 r:%d", value)
	countreads(value, 10, 27)
	l.Unlock()
}

func p11() {
	targets1 := make([]string, 4)
	targets1[3] = uri[s28]
	targets1[2] = uri[s22]
	targets1[1] = uri[s14]
	targets1[0] = uri[s6]
	targets0 := make([]string, 3)
	targets0[2] = uri[s26]
	targets0[1] = uri[s9]
	targets0[0] = uri[s9]

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(163, &value), rsp, s11)
	log("p11 r:%d", value)
	countreads(value, 11, 21)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(119, &value), rsp, s11)
	log("p11 r:%d", value)
	countreads(value, 11, 15)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(30, &value), rsp, s11)
	log("p11 r:%d", value)
	countreads(value, 11, 4)
	l.Unlock()

	delay()
	l.Lock()
	value = 6
	Put(CreateTuple(47, value), rsp, targets0)
	log("p11 w:%d", value)
	countwrites(value, 11, 6)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(122, &value), rsp, s11)
	log("p11 r:%d", value)
	countreads(value, 11, 16)
	l.Unlock()

	delay()
	l.Lock()
	value = 9
	Put(CreateTuple(68, value), rsp, targets1)
	log("p11 w:%d", value)
	countwrites(value, 11, 9)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(256, &value), rsp, s11)
	log("p11 r:%d", value)
	countreads(value, 11, 32)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(128, &value), rsp, s11)
	log("p11 r:%d", value)
	countreads(value, 11, 16)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(29, &value), rsp, s11)
	log("p11 r:%d", value)
	countreads(value, 11, 4)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(190, &value), rsp, s11)
	log("p11 r:%d", value)
	countreads(value, 11, 24)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(230, &value), rsp, s11)
	log("p11 r:%d", value)
	countreads(value, 11, 29)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(119, &value), rsp, s11)
	log("p11 r:%d", value)
	countreads(value, 11, 15)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(60, &value), rsp, s11)
	log("p11 r:%d", value)
	countreads(value, 11, 8)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(140, &value), rsp, s11)
	log("p11 r:%d", value)
	countreads(value, 11, 18)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(29, &value), rsp, s11)
	log("p11 r:%d", value)
	countreads(value, 11, 4)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(88, &value), rsp, s11)
	log("p11 r:%d", value)
	countreads(value, 11, 11)
	l.Unlock()
}

func p12() {
	targets0 := make([]string, 1)
	targets0[0] = uri[s22]

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(172, &value), rsp, s12)
	log("p12 r:%d", value)
	countreads(value, 12, 22)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(150, &value), rsp, s12)
	log("p12 r:%d", value)
	countreads(value, 12, 19)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(231, &value), rsp, s12)
	log("p12 r:%d", value)
	countreads(value, 12, 29)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(59, &value), rsp, s12)
	log("p12 r:%d", value)
	countreads(value, 12, 8)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(71, &value), rsp, s12)
	log("p12 r:%d", value)
	countreads(value, 12, 9)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(101, &value), rsp, s12)
	log("p12 r:%d", value)
	countreads(value, 12, 13)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(223, &value), rsp, s12)
	log("p12 r:%d", value)
	countreads(value, 12, 28)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(144, &value), rsp, s12)
	log("p12 r:%d", value)
	countreads(value, 12, 18)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(151, &value), rsp, s12)
	log("p12 r:%d", value)
	countreads(value, 12, 19)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(25, &value), rsp, s12)
	log("p12 r:%d", value)
	countreads(value, 12, 4)
	l.Unlock()

	delay()
	l.Lock()
	value = 11
	Put(CreateTuple(81, value), rsp, targets0)
	log("p12 w:%d", value)
	countwrites(value, 12, 11)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(139, &value), rsp, s12)
	log("p12 r:%d", value)
	countreads(value, 12, 18)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(143, &value), rsp, s12)
	log("p12 r:%d", value)
	countreads(value, 12, 18)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(14, &value), rsp, s12)
	log("p12 r:%d", value)
	countreads(value, 12, 2)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(153, &value), rsp, s12)
	log("p12 r:%d", value)
	countreads(value, 12, 20)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(79, &value), rsp, s12)
	log("p12 r:%d", value)
	countreads(value, 12, 10)
	l.Unlock()
}

func p13() {
	targets1 := make([]string, 2)
	targets1[1] = uri[s31]
	targets1[0] = uri[s27]
	targets0 := make([]string, 3)
	targets0[2] = uri[s30]
	targets0[1] = uri[s23]
	targets0[0] = uri[s1]

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 26
	Put(CreateTuple(203, value), rsp, targets0)
	log("p13 w:%d", value)
	countwrites(value, 13, 26)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(26, &value), rsp, s13)
	log("p13 r:%d", value)
	countreads(value, 13, 4)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(145, &value), rsp, s13)
	log("p13 r:%d", value)
	countreads(value, 13, 19)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(65, &value), rsp, s13)
	log("p13 r:%d", value)
	countreads(value, 13, 9)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(58, &value), rsp, s13)
	log("p13 r:%d", value)
	countreads(value, 13, 8)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(145, &value), rsp, s13)
	log("p13 r:%d", value)
	countreads(value, 13, 19)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(233, &value), rsp, s13)
	log("p13 r:%d", value)
	countreads(value, 13, 30)
	l.Unlock()

	delay()
	l.Lock()
	value = 15
	Put(CreateTuple(116, value), rsp, targets1)
	log("p13 w:%d", value)
	countwrites(value, 13, 15)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(91, &value), rsp, s13)
	log("p13 r:%d", value)
	countreads(value, 13, 12)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(212, &value), rsp, s13)
	log("p13 r:%d", value)
	countreads(value, 13, 27)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(252, &value), rsp, s13)
	log("p13 r:%d", value)
	countreads(value, 13, 32)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(60, &value), rsp, s13)
	log("p13 r:%d", value)
	countreads(value, 13, 8)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(31, &value), rsp, s13)
	log("p13 r:%d", value)
	countreads(value, 13, 4)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(162, &value), rsp, s13)
	log("p13 r:%d", value)
	countreads(value, 13, 21)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(243, &value), rsp, s13)
	log("p13 r:%d", value)
	countreads(value, 13, 31)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(66, &value), rsp, s13)
	log("p13 r:%d", value)
	countreads(value, 13, 9)
	l.Unlock()
}

func p14() {
	targets1 := make([]string, 3)
	targets1[2] = uri[s22]
	targets1[1] = uri[s10]
	targets1[0] = uri[s5]
	targets0 := make([]string, 0)

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(68, &value), rsp, s14)
	log("p14 r:%d", value)
	countreads(value, 14, 9)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(105, &value), rsp, s14)
	log("p14 r:%d", value)
	countreads(value, 14, 14)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(63, &value), rsp, s14)
	log("p14 r:%d", value)
	countreads(value, 14, 8)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(123, &value), rsp, s14)
	log("p14 r:%d", value)
	countreads(value, 14, 16)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(97, &value), rsp, s14)
	log("p14 r:%d", value)
	countreads(value, 14, 13)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(159, &value), rsp, s14)
	log("p14 r:%d", value)
	countreads(value, 14, 20)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(134, &value), rsp, s14)
	log("p14 r:%d", value)
	countreads(value, 14, 17)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(73, &value), rsp, s14)
	log("p14 r:%d", value)
	countreads(value, 14, 10)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(87, &value), rsp, s14)
	log("p14 r:%d", value)
	countreads(value, 14, 11)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(72, &value), rsp, s14)
	log("p14 r:%d", value)
	countreads(value, 14, 9)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(36, &value), rsp, s14)
	log("p14 r:%d", value)
	countreads(value, 14, 5)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(126, &value), rsp, s14)
	log("p14 r:%d", value)
	countreads(value, 14, 16)
	l.Unlock()

	delay()
	l.Lock()
	value = 28
	Put(CreateTuple(224, value), rsp, targets0)
	log("p14 w:%d", value)
	countwrites(value, 14, 28)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(45, &value), rsp, s14)
	log("p14 r:%d", value)
	countreads(value, 14, 6)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(65, &value), rsp, s14)
	log("p14 r:%d", value)
	countreads(value, 14, 9)
	l.Unlock()

	delay()
	l.Lock()
	value = 30
	Put(CreateTuple(236, value), rsp, targets1)
	log("p14 w:%d", value)
	countwrites(value, 14, 30)
	l.Unlock()
}

func p15() {
	targets1 := make([]string, 0)
	targets0 := make([]string, 1)
	targets0[0] = uri[s25]

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(173, &value), rsp, s15)
	log("p15 r:%d", value)
	countreads(value, 15, 22)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(101, &value), rsp, s15)
	log("p15 r:%d", value)
	countreads(value, 15, 13)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(141, &value), rsp, s15)
	log("p15 r:%d", value)
	countreads(value, 15, 18)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(46, &value), rsp, s15)
	log("p15 r:%d", value)
	countreads(value, 15, 6)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(70, &value), rsp, s15)
	log("p15 r:%d", value)
	countreads(value, 15, 9)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(62, &value), rsp, s15)
	log("p15 r:%d", value)
	countreads(value, 15, 8)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(141, &value), rsp, s15)
	log("p15 r:%d", value)
	countreads(value, 15, 18)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(210, &value), rsp, s15)
	log("p15 r:%d", value)
	countreads(value, 15, 27)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(96, &value), rsp, s15)
	log("p15 r:%d", value)
	countreads(value, 15, 12)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(137, &value), rsp, s15)
	log("p15 r:%d", value)
	countreads(value, 15, 18)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(132, &value), rsp, s15)
	log("p15 r:%d", value)
	countreads(value, 15, 17)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(53, &value), rsp, s15)
	log("p15 r:%d", value)
	countreads(value, 15, 7)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(150, &value), rsp, s15)
	log("p15 r:%d", value)
	countreads(value, 15, 19)
	l.Unlock()

	delay()
	l.Lock()
	value = 24
	Put(CreateTuple(188, value), rsp, targets0)
	log("p15 w:%d", value)
	countwrites(value, 15, 24)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(177, &value), rsp, s15)
	log("p15 r:%d", value)
	countreads(value, 15, 23)
	l.Unlock()

	delay()
	l.Lock()
	value = 27
	Put(CreateTuple(215, value), rsp, targets1)
	log("p15 w:%d", value)
	countwrites(value, 15, 27)
	l.Unlock()
}

func p16() {
	targets1 := make([]string, 0)
	targets0 := make([]string, 3)
	targets0[2] = uri[s23]
	targets0[1] = uri[s19]
	targets0[0] = uri[s5]

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(174, &value), rsp, s16)
	log("p16 r:%d", value)
	countreads(value, 16, 22)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(206, &value), rsp, s16)
	log("p16 r:%d", value)
	countreads(value, 16, 26)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(41, &value), rsp, s16)
	log("p16 r:%d", value)
	countreads(value, 16, 6)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(79, &value), rsp, s16)
	log("p16 r:%d", value)
	countreads(value, 16, 10)
	l.Unlock()

	delay()
	l.Lock()
	value = 11
	Put(CreateTuple(84, value), rsp, targets0)
	log("p16 w:%d", value)
	countwrites(value, 16, 11)
	l.Unlock()

	delay()
	l.Lock()
	value = 20
	Put(CreateTuple(156, value), rsp, targets1)
	log("p16 w:%d", value)
	countwrites(value, 16, 20)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(128, &value), rsp, s16)
	log("p16 r:%d", value)
	countreads(value, 16, 16)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(3, &value), rsp, s16)
	log("p16 r:%d", value)
	countreads(value, 16, 1)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(143, &value), rsp, s16)
	log("p16 r:%d", value)
	countreads(value, 16, 18)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(122, &value), rsp, s16)
	log("p16 r:%d", value)
	countreads(value, 16, 16)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(59, &value), rsp, s16)
	log("p16 r:%d", value)
	countreads(value, 16, 8)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(220, &value), rsp, s16)
	log("p16 r:%d", value)
	countreads(value, 16, 28)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(235, &value), rsp, s16)
	log("p16 r:%d", value)
	countreads(value, 16, 30)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(72, &value), rsp, s16)
	log("p16 r:%d", value)
	countreads(value, 16, 9)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(52, &value), rsp, s16)
	log("p16 r:%d", value)
	countreads(value, 16, 7)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(123, &value), rsp, s16)
	log("p16 r:%d", value)
	countreads(value, 16, 16)
	l.Unlock()
}

func p17() {
	targets0 := make([]string, 2)
	targets0[1] = uri[s26]
	targets0[0] = uri[s5]

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 2
	Put(CreateTuple(12, value), rsp, targets0)
	log("p17 w:%d", value)
	countwrites(value, 17, 2)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(106, &value), rsp, s17)
	log("p17 r:%d", value)
	countreads(value, 17, 14)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(200, &value), rsp, s17)
	log("p17 r:%d", value)
	countreads(value, 17, 25)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(118, &value), rsp, s17)
	log("p17 r:%d", value)
	countreads(value, 17, 15)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(180, &value), rsp, s17)
	log("p17 r:%d", value)
	countreads(value, 17, 23)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(70, &value), rsp, s17)
	log("p17 r:%d", value)
	countreads(value, 17, 9)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(209, &value), rsp, s17)
	log("p17 r:%d", value)
	countreads(value, 17, 27)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(146, &value), rsp, s17)
	log("p17 r:%d", value)
	countreads(value, 17, 19)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(255, &value), rsp, s17)
	log("p17 r:%d", value)
	countreads(value, 17, 32)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(191, &value), rsp, s17)
	log("p17 r:%d", value)
	countreads(value, 17, 24)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(124, &value), rsp, s17)
	log("p17 r:%d", value)
	countreads(value, 17, 16)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(46, &value), rsp, s17)
	log("p17 r:%d", value)
	countreads(value, 17, 6)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(36, &value), rsp, s17)
	log("p17 r:%d", value)
	countreads(value, 17, 5)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(56, &value), rsp, s17)
	log("p17 r:%d", value)
	countreads(value, 17, 7)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(120, &value), rsp, s17)
	log("p17 r:%d", value)
	countreads(value, 17, 15)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(112, &value), rsp, s17)
	log("p17 r:%d", value)
	countreads(value, 17, 14)
	l.Unlock()
}

func p18() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(60, &value), rsp, s18)
	log("p18 r:%d", value)
	countreads(value, 18, 8)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(117, &value), rsp, s18)
	log("p18 r:%d", value)
	countreads(value, 18, 15)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(160, &value), rsp, s18)
	log("p18 r:%d", value)
	countreads(value, 18, 20)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(170, &value), rsp, s18)
	log("p18 r:%d", value)
	countreads(value, 18, 22)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(25, &value), rsp, s18)
	log("p18 r:%d", value)
	countreads(value, 18, 4)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(50, &value), rsp, s18)
	log("p18 r:%d", value)
	countreads(value, 18, 7)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(228, &value), rsp, s18)
	log("p18 r:%d", value)
	countreads(value, 18, 29)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(119, &value), rsp, s18)
	log("p18 r:%d", value)
	countreads(value, 18, 15)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(164, &value), rsp, s18)
	log("p18 r:%d", value)
	countreads(value, 18, 21)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(158, &value), rsp, s18)
	log("p18 r:%d", value)
	countreads(value, 18, 20)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(149, &value), rsp, s18)
	log("p18 r:%d", value)
	countreads(value, 18, 19)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(238, &value), rsp, s18)
	log("p18 r:%d", value)
	countreads(value, 18, 30)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(169, &value), rsp, s18)
	log("p18 r:%d", value)
	countreads(value, 18, 22)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(62, &value), rsp, s18)
	log("p18 r:%d", value)
	countreads(value, 18, 8)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(252, &value), rsp, s18)
	log("p18 r:%d", value)
	countreads(value, 18, 32)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(20, &value), rsp, s18)
	log("p18 r:%d", value)
	countreads(value, 18, 3)
	l.Unlock()
}

func p19() {
	targets1 := make([]string, 3)
	targets1[2] = uri[s23]
	targets1[1] = uri[s17]
	targets1[0] = uri[s8]
	targets0 := make([]string, 1)
	targets0[0] = uri[s3]

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(31, &value), rsp, s19)
	log("p19 r:%d", value)
	countreads(value, 19, 4)
	l.Unlock()

	delay()
	l.Lock()
	value = 17
	Put(CreateTuple(130, value), rsp, targets0)
	log("p19 w:%d", value)
	countwrites(value, 19, 17)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(101, &value), rsp, s19)
	log("p19 r:%d", value)
	countreads(value, 19, 13)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(195, &value), rsp, s19)
	log("p19 r:%d", value)
	countreads(value, 19, 25)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(167, &value), rsp, s19)
	log("p19 r:%d", value)
	countreads(value, 19, 21)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(71, &value), rsp, s19)
	log("p19 r:%d", value)
	countreads(value, 19, 9)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(52, &value), rsp, s19)
	log("p19 r:%d", value)
	countreads(value, 19, 7)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(92, &value), rsp, s19)
	log("p19 r:%d", value)
	countreads(value, 19, 12)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(151, &value), rsp, s19)
	log("p19 r:%d", value)
	countreads(value, 19, 19)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(124, &value), rsp, s19)
	log("p19 r:%d", value)
	countreads(value, 19, 16)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(84, &value), rsp, s19)
	log("p19 r:%d", value)
	countreads(value, 19, 11)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(37, &value), rsp, s19)
	log("p19 r:%d", value)
	countreads(value, 19, 5)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(178, &value), rsp, s19)
	log("p19 r:%d", value)
	countreads(value, 19, 23)
	l.Unlock()

	delay()
	l.Lock()
	value = 14
	Put(CreateTuple(106, value), rsp, targets1)
	log("p19 w:%d", value)
	countwrites(value, 19, 14)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(225, &value), rsp, s19)
	log("p19 r:%d", value)
	countreads(value, 19, 29)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(243, &value), rsp, s19)
	log("p19 r:%d", value)
	countreads(value, 19, 31)
	l.Unlock()
}

func p20() {
	targets1 := make([]string, 0)
	targets0 := make([]string, 0)
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(69, &value), rsp, s20)
	log("p20 r:%d", value)
	countreads(value, 20, 9)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(89, &value), rsp, s20)
	log("p20 r:%d", value)
	countreads(value, 20, 12)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(234, &value), rsp, s20)
	log("p20 r:%d", value)
	countreads(value, 20, 30)
	l.Unlock()

	delay()
	l.Lock()
	value = 15
	Put(CreateTuple(113, value), rsp, targets0)
	log("p20 w:%d", value)
	countwrites(value, 20, 15)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(88, &value), rsp, s20)
	log("p20 r:%d", value)
	countreads(value, 20, 11)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(178, &value), rsp, s20)
	log("p20 r:%d", value)
	countreads(value, 20, 23)
	l.Unlock()

	delay()
	l.Lock()
	value = 10
	Put(CreateTuple(76, value), rsp, targets1)
	log("p20 w:%d", value)
	countwrites(value, 20, 10)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(225, &value), rsp, s20)
	log("p20 r:%d", value)
	countreads(value, 20, 29)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(72, &value), rsp, s20)
	log("p20 r:%d", value)
	countreads(value, 20, 9)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(183, &value), rsp, s20)
	log("p20 r:%d", value)
	countreads(value, 20, 23)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(93, &value), rsp, s20)
	log("p20 r:%d", value)
	countreads(value, 20, 12)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(167, &value), rsp, s20)
	log("p20 r:%d", value)
	countreads(value, 20, 21)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(63, &value), rsp, s20)
	log("p20 r:%d", value)
	countreads(value, 20, 8)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(129, &value), rsp, s20)
	log("p20 r:%d", value)
	countreads(value, 20, 17)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(19, &value), rsp, s20)
	log("p20 r:%d", value)
	countreads(value, 20, 3)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(39, &value), rsp, s20)
	log("p20 r:%d", value)
	countreads(value, 20, 5)
	l.Unlock()
}

func p21() {
	targets0 := make([]string, 1)
	targets0[0] = uri[s18]

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(54, &value), rsp, s21)
	log("p21 r:%d", value)
	countreads(value, 21, 7)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(11, &value), rsp, s21)
	log("p21 r:%d", value)
	countreads(value, 21, 2)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(51, &value), rsp, s21)
	log("p21 r:%d", value)
	countreads(value, 21, 7)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(136, &value), rsp, s21)
	log("p21 r:%d", value)
	countreads(value, 21, 17)
	l.Unlock()

	delay()
	l.Lock()
	value = 7
	Put(CreateTuple(50, value), rsp, targets0)
	log("p21 w:%d", value)
	countwrites(value, 21, 7)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(52, &value), rsp, s21)
	log("p21 r:%d", value)
	countreads(value, 21, 7)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(30, &value), rsp, s21)
	log("p21 r:%d", value)
	countreads(value, 21, 4)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(96, &value), rsp, s21)
	log("p21 r:%d", value)
	countreads(value, 21, 12)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(40, &value), rsp, s21)
	log("p21 r:%d", value)
	countreads(value, 21, 5)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(34, &value), rsp, s21)
	log("p21 r:%d", value)
	countreads(value, 21, 5)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(111, &value), rsp, s21)
	log("p21 r:%d", value)
	countreads(value, 21, 14)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(125, &value), rsp, s21)
	log("p21 r:%d", value)
	countreads(value, 21, 16)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(246, &value), rsp, s21)
	log("p21 r:%d", value)
	countreads(value, 21, 31)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(237, &value), rsp, s21)
	log("p21 r:%d", value)
	countreads(value, 21, 30)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(95, &value), rsp, s21)
	log("p21 r:%d", value)
	countreads(value, 21, 12)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(64, &value), rsp, s21)
	log("p21 r:%d", value)
	countreads(value, 21, 8)
	l.Unlock()
}

func p22() {
	targets0 := make([]string, 3)
	targets0[2] = uri[s21]
	targets0[1] = uri[s11]
	targets0[0] = uri[s10]

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(222, &value), rsp, s22)
	log("p22 r:%d", value)
	countreads(value, 22, 28)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(39, &value), rsp, s22)
	log("p22 r:%d", value)
	countreads(value, 22, 5)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(91, &value), rsp, s22)
	log("p22 r:%d", value)
	countreads(value, 22, 12)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(14, &value), rsp, s22)
	log("p22 r:%d", value)
	countreads(value, 22, 2)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(68, &value), rsp, s22)
	log("p22 r:%d", value)
	countreads(value, 22, 9)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(127, &value), rsp, s22)
	log("p22 r:%d", value)
	countreads(value, 22, 16)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(33, &value), rsp, s22)
	log("p22 r:%d", value)
	countreads(value, 22, 5)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(236, &value), rsp, s22)
	log("p22 r:%d", value)
	countreads(value, 22, 30)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(81, &value), rsp, s22)
	log("p22 r:%d", value)
	countreads(value, 22, 11)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(216, &value), rsp, s22)
	log("p22 r:%d", value)
	countreads(value, 22, 27)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(254, &value), rsp, s22)
	log("p22 r:%d", value)
	countreads(value, 22, 32)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(127, &value), rsp, s22)
	log("p22 r:%d", value)
	countreads(value, 22, 16)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(52, &value), rsp, s22)
	log("p22 r:%d", value)
	countreads(value, 22, 7)
	l.Unlock()

	delay()
	l.Lock()
	value = 4
	Put(CreateTuple(30, value), rsp, targets0)
	log("p22 w:%d", value)
	countwrites(value, 22, 4)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(20, &value), rsp, s22)
	log("p22 r:%d", value)
	countreads(value, 22, 3)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(208, &value), rsp, s22)
	log("p22 r:%d", value)
	countreads(value, 22, 26)
	l.Unlock()
}

func p23() {
	targets3 := make([]string, 0)
	targets2 := make([]string, 4)
	targets2[3] = uri[s22]
	targets2[2] = uri[s22]
	targets2[1] = uri[s5]
	targets2[0] = uri[s2]
	targets1 := make([]string, 1)
	targets1[0] = uri[s20]
	targets0 := make([]string, 1)
	targets0[0] = uri[s26]

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 12
	Put(CreateTuple(90, value), rsp, targets0)
	log("p23 w:%d", value)
	countwrites(value, 23, 12)
	l.Unlock()

	delay()
	l.Lock()
	value = 12
	Put(CreateTuple(89, value), rsp, targets1)
	log("p23 w:%d", value)
	countwrites(value, 23, 12)
	l.Unlock()

	delay()
	l.Lock()
	value = 16
	Put(CreateTuple(127, value), rsp, targets2)
	log("p23 w:%d", value)
	countwrites(value, 23, 16)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(98, &value), rsp, s23)
	log("p23 r:%d", value)
	countreads(value, 23, 13)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(33, &value), rsp, s23)
	log("p23 r:%d", value)
	countreads(value, 23, 5)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(73, &value), rsp, s23)
	log("p23 r:%d", value)
	countreads(value, 23, 10)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(134, &value), rsp, s23)
	log("p23 r:%d", value)
	countreads(value, 23, 17)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(171, &value), rsp, s23)
	log("p23 r:%d", value)
	countreads(value, 23, 22)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(179, &value), rsp, s23)
	log("p23 r:%d", value)
	countreads(value, 23, 23)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(44, &value), rsp, s23)
	log("p23 r:%d", value)
	countreads(value, 23, 6)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(252, &value), rsp, s23)
	log("p23 r:%d", value)
	countreads(value, 23, 32)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(33, &value), rsp, s23)
	log("p23 r:%d", value)
	countreads(value, 23, 5)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(106, &value), rsp, s23)
	log("p23 r:%d", value)
	countreads(value, 23, 14)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(203, &value), rsp, s23)
	log("p23 r:%d", value)
	countreads(value, 23, 26)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(84, &value), rsp, s23)
	log("p23 r:%d", value)
	countreads(value, 23, 11)
	l.Unlock()

	delay()
	l.Lock()
	value = 29
	Put(CreateTuple(227, value), rsp, targets3)
	log("p23 w:%d", value)
	countwrites(value, 23, 29)
	l.Unlock()
}

func p24() {
	targets1 := make([]string, 0)
	targets0 := make([]string, 0)
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(114, &value), rsp, s24)
	log("p24 r:%d", value)
	countreads(value, 24, 15)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(212, &value), rsp, s24)
	log("p24 r:%d", value)
	countreads(value, 24, 27)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(250, &value), rsp, s24)
	log("p24 r:%d", value)
	countreads(value, 24, 32)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(255, &value), rsp, s24)
	log("p24 r:%d", value)
	countreads(value, 24, 32)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(41, &value), rsp, s24)
	log("p24 r:%d", value)
	countreads(value, 24, 6)
	l.Unlock()

	delay()
	l.Lock()
	value = 27
	Put(CreateTuple(214, value), rsp, targets0)
	log("p24 w:%d", value)
	countwrites(value, 24, 27)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(1, &value), rsp, s24)
	log("p24 r:%d", value)
	countreads(value, 24, 1)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(115, &value), rsp, s24)
	log("p24 r:%d", value)
	countreads(value, 24, 15)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(158, &value), rsp, s24)
	log("p24 r:%d", value)
	countreads(value, 24, 20)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(21, &value), rsp, s24)
	log("p24 r:%d", value)
	countreads(value, 24, 3)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(256, &value), rsp, s24)
	log("p24 r:%d", value)
	countreads(value, 24, 32)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(208, &value), rsp, s24)
	log("p24 r:%d", value)
	countreads(value, 24, 26)
	l.Unlock()

	delay()
	l.Lock()
	value = 28
	Put(CreateTuple(218, value), rsp, targets1)
	log("p24 w:%d", value)
	countwrites(value, 24, 28)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(23, &value), rsp, s24)
	log("p24 r:%d", value)
	countreads(value, 24, 3)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(228, &value), rsp, s24)
	log("p24 r:%d", value)
	countreads(value, 24, 29)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(97, &value), rsp, s24)
	log("p24 r:%d", value)
	countreads(value, 24, 13)
	l.Unlock()
}

func p25() {
	targets0 := make([]string, 1)
	targets0[0] = uri[s1]

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(204, &value), rsp, s25)
	log("p25 r:%d", value)
	countreads(value, 25, 26)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(61, &value), rsp, s25)
	log("p25 r:%d", value)
	countreads(value, 25, 8)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(210, &value), rsp, s25)
	log("p25 r:%d", value)
	countreads(value, 25, 27)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(188, &value), rsp, s25)
	log("p25 r:%d", value)
	countreads(value, 25, 24)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(60, &value), rsp, s25)
	log("p25 r:%d", value)
	countreads(value, 25, 8)
	l.Unlock()

	delay()
	l.Lock()
	value = 26
	Put(CreateTuple(207, value), rsp, targets0)
	log("p25 w:%d", value)
	countwrites(value, 25, 26)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(26, &value), rsp, s25)
	log("p25 r:%d", value)
	countreads(value, 25, 4)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(114, &value), rsp, s25)
	log("p25 r:%d", value)
	countreads(value, 25, 15)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(253, &value), rsp, s25)
	log("p25 r:%d", value)
	countreads(value, 25, 32)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(165, &value), rsp, s25)
	log("p25 r:%d", value)
	countreads(value, 25, 21)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(168, &value), rsp, s25)
	log("p25 r:%d", value)
	countreads(value, 25, 21)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(23, &value), rsp, s25)
	log("p25 r:%d", value)
	countreads(value, 25, 3)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(231, &value), rsp, s25)
	log("p25 r:%d", value)
	countreads(value, 25, 29)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(58, &value), rsp, s25)
	log("p25 r:%d", value)
	countreads(value, 25, 8)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(31, &value), rsp, s25)
	log("p25 r:%d", value)
	countreads(value, 25, 4)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(38, &value), rsp, s25)
	log("p25 r:%d", value)
	countreads(value, 25, 5)
	l.Unlock()
}

func p26() {
	targets0 := make([]string, 1)
	targets0[0] = uri[s3]

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(232, &value), rsp, s26)
	log("p26 r:%d", value)
	countreads(value, 26, 29)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(201, &value), rsp, s26)
	log("p26 r:%d", value)
	countreads(value, 26, 26)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(163, &value), rsp, s26)
	log("p26 r:%d", value)
	countreads(value, 26, 21)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(78, &value), rsp, s26)
	log("p26 r:%d", value)
	countreads(value, 26, 10)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(90, &value), rsp, s26)
	log("p26 r:%d", value)
	countreads(value, 26, 12)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(12, &value), rsp, s26)
	log("p26 r:%d", value)
	countreads(value, 26, 2)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(36, &value), rsp, s26)
	log("p26 r:%d", value)
	countreads(value, 26, 5)
	l.Unlock()

	delay()
	l.Lock()
	value = 30
	Put(CreateTuple(239, value), rsp, targets0)
	log("p26 w:%d", value)
	countwrites(value, 26, 30)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(225, &value), rsp, s26)
	log("p26 r:%d", value)
	countreads(value, 26, 29)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(17, &value), rsp, s26)
	log("p26 r:%d", value)
	countreads(value, 26, 3)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(52, &value), rsp, s26)
	log("p26 r:%d", value)
	countreads(value, 26, 7)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(85, &value), rsp, s26)
	log("p26 r:%d", value)
	countreads(value, 26, 11)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(179, &value), rsp, s26)
	log("p26 r:%d", value)
	countreads(value, 26, 23)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(246, &value), rsp, s26)
	log("p26 r:%d", value)
	countreads(value, 26, 31)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(47, &value), rsp, s26)
	log("p26 r:%d", value)
	countreads(value, 26, 6)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(195, &value), rsp, s26)
	log("p26 r:%d", value)
	countreads(value, 26, 25)
	l.Unlock()
}

func p27() {
	targets0 := make([]string, 2)
	targets0[1] = uri[s5]
	targets0[0] = uri[s3]

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(252, &value), rsp, s27)
	log("p27 r:%d", value)
	countreads(value, 27, 32)
	l.Unlock()

	delay()
	l.Lock()
	value = 7
	Put(CreateTuple(49, value), rsp, targets0)
	log("p27 w:%d", value)
	countwrites(value, 27, 7)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(122, &value), rsp, s27)
	log("p27 r:%d", value)
	countreads(value, 27, 16)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(116, &value), rsp, s27)
	log("p27 r:%d", value)
	countreads(value, 27, 15)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(212, &value), rsp, s27)
	log("p27 r:%d", value)
	countreads(value, 27, 27)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(252, &value), rsp, s27)
	log("p27 r:%d", value)
	countreads(value, 27, 32)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(233, &value), rsp, s27)
	log("p27 r:%d", value)
	countreads(value, 27, 30)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(202, &value), rsp, s27)
	log("p27 r:%d", value)
	countreads(value, 27, 26)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(114, &value), rsp, s27)
	log("p27 r:%d", value)
	countreads(value, 27, 15)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(55, &value), rsp, s27)
	log("p27 r:%d", value)
	countreads(value, 27, 7)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(129, &value), rsp, s27)
	log("p27 r:%d", value)
	countreads(value, 27, 17)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(20, &value), rsp, s27)
	log("p27 r:%d", value)
	countreads(value, 27, 3)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(35, &value), rsp, s27)
	log("p27 r:%d", value)
	countreads(value, 27, 5)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(213, &value), rsp, s27)
	log("p27 r:%d", value)
	countreads(value, 27, 27)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(244, &value), rsp, s27)
	log("p27 r:%d", value)
	countreads(value, 27, 31)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(28, &value), rsp, s27)
	log("p27 r:%d", value)
	countreads(value, 27, 4)
	l.Unlock()
}

func p28() {
	targets0 := make([]string, 2)
	targets0[1] = uri[s10]
	targets0[0] = uri[s6]

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(229, &value), rsp, s28)
	log("p28 r:%d", value)
	countreads(value, 28, 29)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(237, &value), rsp, s28)
	log("p28 r:%d", value)
	countreads(value, 28, 30)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(247, &value), rsp, s28)
	log("p28 r:%d", value)
	countreads(value, 28, 31)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(216, &value), rsp, s28)
	log("p28 r:%d", value)
	countreads(value, 28, 27)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(196, &value), rsp, s28)
	log("p28 r:%d", value)
	countreads(value, 28, 25)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(143, &value), rsp, s28)
	log("p28 r:%d", value)
	countreads(value, 28, 18)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(68, &value), rsp, s28)
	log("p28 r:%d", value)
	countreads(value, 28, 9)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(190, &value), rsp, s28)
	log("p28 r:%d", value)
	countreads(value, 28, 24)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(143, &value), rsp, s28)
	log("p28 r:%d", value)
	countreads(value, 28, 18)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(223, &value), rsp, s28)
	log("p28 r:%d", value)
	countreads(value, 28, 28)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(86, &value), rsp, s28)
	log("p28 r:%d", value)
	countreads(value, 28, 11)
	l.Unlock()

	delay()
	l.Lock()
	value = 31
	Put(CreateTuple(245, value), rsp, targets0)
	log("p28 w:%d", value)
	countwrites(value, 28, 31)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(36, &value), rsp, s28)
	log("p28 r:%d", value)
	countreads(value, 28, 5)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(158, &value), rsp, s28)
	log("p28 r:%d", value)
	countreads(value, 28, 20)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(129, &value), rsp, s28)
	log("p28 r:%d", value)
	countreads(value, 28, 17)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(155, &value), rsp, s28)
	log("p28 r:%d", value)
	countreads(value, 28, 20)
	l.Unlock()
}

func p29() {
	targets1 := make([]string, 0)
	targets0 := make([]string, 1)
	targets0[0] = uri[s3]

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(164, &value), rsp, s29)
	log("p29 r:%d", value)
	countreads(value, 29, 21)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(200, &value), rsp, s29)
	log("p29 r:%d", value)
	countreads(value, 29, 25)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(145, &value), rsp, s29)
	log("p29 r:%d", value)
	countreads(value, 29, 19)
	l.Unlock()

	delay()
	l.Lock()
	value = 9
	Put(CreateTuple(67, value), rsp, targets0)
	log("p29 w:%d", value)
	countwrites(value, 29, 9)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(8, &value), rsp, s29)
	log("p29 r:%d", value)
	countreads(value, 29, 1)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(250, &value), rsp, s29)
	log("p29 r:%d", value)
	countreads(value, 29, 32)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(110, &value), rsp, s29)
	log("p29 r:%d", value)
	countreads(value, 29, 14)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(125, &value), rsp, s29)
	log("p29 r:%d", value)
	countreads(value, 29, 16)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(98, &value), rsp, s29)
	log("p29 r:%d", value)
	countreads(value, 29, 13)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(9, &value), rsp, s29)
	log("p29 r:%d", value)
	countreads(value, 29, 2)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(223, &value), rsp, s29)
	log("p29 r:%d", value)
	countreads(value, 29, 28)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(184, &value), rsp, s29)
	log("p29 r:%d", value)
	countreads(value, 29, 23)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(115, &value), rsp, s29)
	log("p29 r:%d", value)
	countreads(value, 29, 15)
	l.Unlock()

	delay()
	l.Lock()
	value = 30
	Put(CreateTuple(240, value), rsp, targets1)
	log("p29 w:%d", value)
	countwrites(value, 29, 30)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(189, &value), rsp, s29)
	log("p29 r:%d", value)
	countreads(value, 29, 24)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(145, &value), rsp, s29)
	log("p29 r:%d", value)
	countreads(value, 29, 19)
	l.Unlock()
}

func p30() {
	targets1 := make([]string, 1)
	targets1[0] = uri[s29]
	targets0 := make([]string, 1)
	targets0[0] = uri[s24]

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(203, &value), rsp, s30)
	log("p30 r:%d", value)
	countreads(value, 30, 26)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(229, &value), rsp, s30)
	log("p30 r:%d", value)
	countreads(value, 30, 29)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(173, &value), rsp, s30)
	log("p30 r:%d", value)
	countreads(value, 30, 22)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(28, &value), rsp, s30)
	log("p30 r:%d", value)
	countreads(value, 30, 4)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(64, &value), rsp, s30)
	log("p30 r:%d", value)
	countreads(value, 30, 8)
	l.Unlock()

	delay()
	l.Lock()
	value = 1
	Put(CreateTuple(1, value), rsp, targets0)
	log("p30 w:%d", value)
	countwrites(value, 30, 1)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(197, &value), rsp, s30)
	log("p30 r:%d", value)
	countreads(value, 30, 25)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(198, &value), rsp, s30)
	log("p30 r:%d", value)
	countreads(value, 30, 25)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(147, &value), rsp, s30)
	log("p30 r:%d", value)
	countreads(value, 30, 19)
	l.Unlock()

	delay()
	l.Lock()
	value = 2
	Put(CreateTuple(9, value), rsp, targets1)
	log("p30 w:%d", value)
	countwrites(value, 30, 2)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(201, &value), rsp, s30)
	log("p30 r:%d", value)
	countreads(value, 30, 26)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(65, &value), rsp, s30)
	log("p30 r:%d", value)
	countreads(value, 30, 9)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(109, &value), rsp, s30)
	log("p30 r:%d", value)
	countreads(value, 30, 14)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(119, &value), rsp, s30)
	log("p30 r:%d", value)
	countreads(value, 30, 15)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(254, &value), rsp, s30)
	log("p30 r:%d", value)
	countreads(value, 30, 32)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(167, &value), rsp, s30)
	log("p30 r:%d", value)
	countreads(value, 30, 21)
	l.Unlock()
}

func p31() {
	targets0 := make([]string, 1)
	targets0[0] = uri[s1]

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(230, &value), rsp, s31)
	log("p31 r:%d", value)
	countreads(value, 31, 29)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(66, &value), rsp, s31)
	log("p31 r:%d", value)
	countreads(value, 31, 9)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(194, &value), rsp, s31)
	log("p31 r:%d", value)
	countreads(value, 31, 25)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(247, &value), rsp, s31)
	log("p31 r:%d", value)
	countreads(value, 31, 31)
	l.Unlock()

	delay()
	l.Lock()
	value = 17
	Put(CreateTuple(131, value), rsp, targets0)
	log("p31 w:%d", value)
	countwrites(value, 31, 17)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(232, &value), rsp, s31)
	log("p31 r:%d", value)
	countreads(value, 31, 29)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(6, &value), rsp, s31)
	log("p31 r:%d", value)
	countreads(value, 31, 1)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(191, &value), rsp, s31)
	log("p31 r:%d", value)
	countreads(value, 31, 24)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(86, &value), rsp, s31)
	log("p31 r:%d", value)
	countreads(value, 31, 11)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(54, &value), rsp, s31)
	log("p31 r:%d", value)
	countreads(value, 31, 7)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(166, &value), rsp, s31)
	log("p31 r:%d", value)
	countreads(value, 31, 21)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(141, &value), rsp, s31)
	log("p31 r:%d", value)
	countreads(value, 31, 18)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(116, &value), rsp, s31)
	log("p31 r:%d", value)
	countreads(value, 31, 15)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(71, &value), rsp, s31)
	log("p31 r:%d", value)
	countreads(value, 31, 9)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(71, &value), rsp, s31)
	log("p31 r:%d", value)
	countreads(value, 31, 9)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(97, &value), rsp, s31)
	log("p31 r:%d", value)
	countreads(value, 31, 13)
	l.Unlock()
}

func p32() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(235, &value), rsp, s32)
	log("p32 r:%d", value)
	countreads(value, 32, 30)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(142, &value), rsp, s32)
	log("p32 r:%d", value)
	countreads(value, 32, 18)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(180, &value), rsp, s32)
	log("p32 r:%d", value)
	countreads(value, 32, 23)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(87, &value), rsp, s32)
	log("p32 r:%d", value)
	countreads(value, 32, 11)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(230, &value), rsp, s32)
	log("p32 r:%d", value)
	countreads(value, 32, 29)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(33, &value), rsp, s32)
	log("p32 r:%d", value)
	countreads(value, 32, 5)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(10, &value), rsp, s32)
	log("p32 r:%d", value)
	countreads(value, 32, 2)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(251, &value), rsp, s32)
	log("p32 r:%d", value)
	countreads(value, 32, 32)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(184, &value), rsp, s32)
	log("p32 r:%d", value)
	countreads(value, 32, 23)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(192, &value), rsp, s32)
	log("p32 r:%d", value)
	countreads(value, 32, 24)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(249, &value), rsp, s32)
	log("p32 r:%d", value)
	countreads(value, 32, 32)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(95, &value), rsp, s32)
	log("p32 r:%d", value)
	countreads(value, 32, 12)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(223, &value), rsp, s32)
	log("p32 r:%d", value)
	countreads(value, 32, 28)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(172, &value), rsp, s32)
	log("p32 r:%d", value)
	countreads(value, 32, 22)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(112, &value), rsp, s32)
	log("p32 r:%d", value)
	countreads(value, 32, 14)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(59, &value), rsp, s32)
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
