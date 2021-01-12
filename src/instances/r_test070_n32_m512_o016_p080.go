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
	targets11 := make([]string, 0)
	targets10 := make([]string, 1)
	targets10[0] = uri[s4]
	targets9 := make([]string, 0)
	targets8 := make([]string, 0)
	targets7 := make([]string, 0)
	targets6 := make([]string, 0)
	targets5 := make([]string, 2)
	targets5[1] = uri[s32]
	targets5[0] = uri[s11]
	targets4 := make([]string, 0)
	targets3 := make([]string, 0)
	targets2 := make([]string, 0)
	targets1 := make([]string, 0)
	targets0 := make([]string, 0)

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 15
	Put(CreateTuple(240, value), rsp, targets0)
	log("p1 w:%d", value)
	countwrites(value, 1, 15)
	l.Unlock()

	delay()
	l.Lock()
	value = 28
	Put(CreateTuple(439, value), rsp, targets1)
	log("p1 w:%d", value)
	countwrites(value, 1, 28)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(409, &value), rsp, s1)
	log("p1 r:%d", value)
	countreads(value, 1, 26)
	l.Unlock()

	delay()
	l.Lock()
	value = 3
	Put(CreateTuple(44, value), rsp, targets2)
	log("p1 w:%d", value)
	countwrites(value, 1, 3)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(344, &value), rsp, s1)
	log("p1 r:%d", value)
	countreads(value, 1, 22)
	l.Unlock()

	delay()
	l.Lock()
	value = 25
	Put(CreateTuple(387, value), rsp, targets3)
	log("p1 w:%d", value)
	countwrites(value, 1, 25)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(491, &value), rsp, s1)
	log("p1 r:%d", value)
	countreads(value, 1, 31)
	l.Unlock()

	delay()
	l.Lock()
	value = 27
	Put(CreateTuple(430, value), rsp, targets4)
	log("p1 w:%d", value)
	countwrites(value, 1, 27)
	l.Unlock()

	delay()
	l.Lock()
	value = 31
	Put(CreateTuple(482, value), rsp, targets5)
	log("p1 w:%d", value)
	countwrites(value, 1, 31)
	l.Unlock()

	delay()
	l.Lock()
	value = 31
	Put(CreateTuple(489, value), rsp, targets6)
	log("p1 w:%d", value)
	countwrites(value, 1, 31)
	l.Unlock()

	delay()
	l.Lock()
	value = 18
	Put(CreateTuple(276, value), rsp, targets7)
	log("p1 w:%d", value)
	countwrites(value, 1, 18)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(193, &value), rsp, s1)
	log("p1 r:%d", value)
	countreads(value, 1, 13)
	l.Unlock()

	delay()
	l.Lock()
	value = 2
	Put(CreateTuple(20, value), rsp, targets8)
	log("p1 w:%d", value)
	countwrites(value, 1, 2)
	l.Unlock()

	delay()
	l.Lock()
	value = 5
	Put(CreateTuple(77, value), rsp, targets9)
	log("p1 w:%d", value)
	countwrites(value, 1, 5)
	l.Unlock()

	delay()
	l.Lock()
	value = 13
	Put(CreateTuple(195, value), rsp, targets10)
	log("p1 w:%d", value)
	countwrites(value, 1, 13)
	l.Unlock()

	delay()
	l.Lock()
	value = 10
	Put(CreateTuple(156, value), rsp, targets11)
	log("p1 w:%d", value)
	countwrites(value, 1, 10)
	l.Unlock()
}

func p2() {
	targets11 := make([]string, 0)
	targets10 := make([]string, 0)
	targets9 := make([]string, 0)
	targets8 := make([]string, 0)
	targets7 := make([]string, 1)
	targets7[0] = uri[s2]
	targets6 := make([]string, 0)
	targets5 := make([]string, 0)
	targets4 := make([]string, 0)
	targets3 := make([]string, 1)
	targets3[0] = uri[s19]
	targets2 := make([]string, 0)
	targets1 := make([]string, 0)
	targets0 := make([]string, 0)

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(216, &value), rsp, s2)
	log("p2 r:%d", value)
	countreads(value, 2, 14)
	l.Unlock()

	delay()
	l.Lock()
	value = 27
	Put(CreateTuple(428, value), rsp, targets0)
	log("p2 w:%d", value)
	countwrites(value, 2, 27)
	l.Unlock()

	delay()
	l.Lock()
	value = 6
	Put(CreateTuple(94, value), rsp, targets1)
	log("p2 w:%d", value)
	countwrites(value, 2, 6)
	l.Unlock()

	delay()
	l.Lock()
	value = 22
	Put(CreateTuple(347, value), rsp, targets2)
	log("p2 w:%d", value)
	countwrites(value, 2, 22)
	l.Unlock()

	delay()
	l.Lock()
	value = 29
	Put(CreateTuple(450, value), rsp, targets3)
	log("p2 w:%d", value)
	countwrites(value, 2, 29)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(497, &value), rsp, s2)
	log("p2 r:%d", value)
	countreads(value, 2, 32)
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
	value = 22
	Put(CreateTuple(348, value), rsp, targets4)
	log("p2 w:%d", value)
	countwrites(value, 2, 22)
	l.Unlock()

	delay()
	l.Lock()
	value = 13
	Put(CreateTuple(196, value), rsp, targets5)
	log("p2 w:%d", value)
	countwrites(value, 2, 13)
	l.Unlock()

	delay()
	l.Lock()
	value = 17
	Put(CreateTuple(261, value), rsp, targets6)
	log("p2 w:%d", value)
	countwrites(value, 2, 17)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(494, &value), rsp, s2)
	log("p2 r:%d", value)
	countreads(value, 2, 31)
	l.Unlock()

	delay()
	l.Lock()
	value = 31
	Put(CreateTuple(494, value), rsp, targets7)
	log("p2 w:%d", value)
	countwrites(value, 2, 31)
	l.Unlock()

	delay()
	l.Lock()
	value = 7
	Put(CreateTuple(109, value), rsp, targets8)
	log("p2 w:%d", value)
	countwrites(value, 2, 7)
	l.Unlock()

	delay()
	l.Lock()
	value = 29
	Put(CreateTuple(458, value), rsp, targets9)
	log("p2 w:%d", value)
	countwrites(value, 2, 29)
	l.Unlock()

	delay()
	l.Lock()
	value = 8
	Put(CreateTuple(127, value), rsp, targets10)
	log("p2 w:%d", value)
	countwrites(value, 2, 8)
	l.Unlock()

	delay()
	l.Lock()
	value = 23
	Put(CreateTuple(357, value), rsp, targets11)
	log("p2 w:%d", value)
	countwrites(value, 2, 23)
	l.Unlock()
}

func p3() {
	targets9 := make([]string, 0)
	targets8 := make([]string, 0)
	targets7 := make([]string, 0)
	targets6 := make([]string, 0)
	targets5 := make([]string, 0)
	targets4 := make([]string, 0)
	targets3 := make([]string, 0)
	targets2 := make([]string, 1)
	targets2[0] = uri[s1]
	targets1 := make([]string, 0)
	targets0 := make([]string, 1)
	targets0[0] = uri[s13]

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(386, &value), rsp, s3)
	log("p3 r:%d", value)
	countreads(value, 3, 25)
	l.Unlock()

	delay()
	l.Lock()
	value = 2
	Put(CreateTuple(23, value), rsp, targets0)
	log("p3 w:%d", value)
	countwrites(value, 3, 2)
	l.Unlock()

	delay()
	l.Lock()
	value = 3
	Put(CreateTuple(34, value), rsp, targets1)
	log("p3 w:%d", value)
	countwrites(value, 3, 3)
	l.Unlock()

	delay()
	l.Lock()
	value = 13
	Put(CreateTuple(193, value), rsp, targets2)
	log("p3 w:%d", value)
	countwrites(value, 3, 13)
	l.Unlock()

	delay()
	l.Lock()
	value = 21
	Put(CreateTuple(336, value), rsp, targets3)
	log("p3 w:%d", value)
	countwrites(value, 3, 21)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(426, &value), rsp, s3)
	log("p3 r:%d", value)
	countreads(value, 3, 27)
	l.Unlock()

	delay()
	l.Lock()
	value = 18
	Put(CreateTuple(288, value), rsp, targets4)
	log("p3 w:%d", value)
	countwrites(value, 3, 18)
	l.Unlock()

	delay()
	l.Lock()
	value = 20
	Put(CreateTuple(307, value), rsp, targets5)
	log("p3 w:%d", value)
	countwrites(value, 3, 20)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(395, &value), rsp, s3)
	log("p3 r:%d", value)
	countreads(value, 3, 25)
	l.Unlock()

	delay()
	l.Lock()
	value = 1
	Put(CreateTuple(4, value), rsp, targets6)
	log("p3 w:%d", value)
	countwrites(value, 3, 1)
	l.Unlock()

	delay()
	l.Lock()
	value = 29
	Put(CreateTuple(453, value), rsp, targets7)
	log("p3 w:%d", value)
	countwrites(value, 3, 29)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(155, &value), rsp, s3)
	log("p3 r:%d", value)
	countreads(value, 3, 10)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(231, &value), rsp, s3)
	log("p3 r:%d", value)
	countreads(value, 3, 15)
	l.Unlock()

	delay()
	l.Lock()
	value = 13
	Put(CreateTuple(198, value), rsp, targets8)
	log("p3 w:%d", value)
	countwrites(value, 3, 13)
	l.Unlock()

	delay()
	l.Lock()
	value = 14
	Put(CreateTuple(209, value), rsp, targets9)
	log("p3 w:%d", value)
	countwrites(value, 3, 14)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(460, &value), rsp, s3)
	log("p3 r:%d", value)
	countreads(value, 3, 29)
	l.Unlock()
}

func p4() {
	targets13 := make([]string, 0)
	targets12 := make([]string, 1)
	targets12[0] = uri[s11]
	targets11 := make([]string, 0)
	targets10 := make([]string, 0)
	targets9 := make([]string, 0)
	targets8 := make([]string, 0)
	targets7 := make([]string, 0)
	targets6 := make([]string, 0)
	targets5 := make([]string, 0)
	targets4 := make([]string, 1)
	targets4[0] = uri[s23]
	targets3 := make([]string, 0)
	targets2 := make([]string, 0)
	targets1 := make([]string, 0)
	targets0 := make([]string, 1)
	targets0[0] = uri[s3]

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 10
	Put(CreateTuple(155, value), rsp, targets0)
	log("p4 w:%d", value)
	countwrites(value, 4, 10)
	l.Unlock()

	delay()
	l.Lock()
	value = 6
	Put(CreateTuple(93, value), rsp, targets1)
	log("p4 w:%d", value)
	countwrites(value, 4, 6)
	l.Unlock()

	delay()
	l.Lock()
	value = 19
	Put(CreateTuple(300, value), rsp, targets2)
	log("p4 w:%d", value)
	countwrites(value, 4, 19)
	l.Unlock()

	delay()
	l.Lock()
	value = 15
	Put(CreateTuple(236, value), rsp, targets3)
	log("p4 w:%d", value)
	countwrites(value, 4, 15)
	l.Unlock()

	delay()
	l.Lock()
	value = 12
	Put(CreateTuple(178, value), rsp, targets4)
	log("p4 w:%d", value)
	countwrites(value, 4, 12)
	l.Unlock()

	delay()
	l.Lock()
	value = 23
	Put(CreateTuple(356, value), rsp, targets5)
	log("p4 w:%d", value)
	countwrites(value, 4, 23)
	l.Unlock()

	delay()
	l.Lock()
	value = 20
	Put(CreateTuple(319, value), rsp, targets6)
	log("p4 w:%d", value)
	countwrites(value, 4, 20)
	l.Unlock()

	delay()
	l.Lock()
	value = 9
	Put(CreateTuple(140, value), rsp, targets7)
	log("p4 w:%d", value)
	countwrites(value, 4, 9)
	l.Unlock()

	delay()
	l.Lock()
	value = 11
	Put(CreateTuple(164, value), rsp, targets8)
	log("p4 w:%d", value)
	countwrites(value, 4, 11)
	l.Unlock()

	delay()
	l.Lock()
	value = 7
	Put(CreateTuple(109, value), rsp, targets9)
	log("p4 w:%d", value)
	countwrites(value, 4, 7)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(200, &value), rsp, s4)
	log("p4 r:%d", value)
	countreads(value, 4, 13)
	l.Unlock()

	delay()
	l.Lock()
	value = 24
	Put(CreateTuple(374, value), rsp, targets10)
	log("p4 w:%d", value)
	countwrites(value, 4, 24)
	l.Unlock()

	delay()
	l.Lock()
	value = 24
	Put(CreateTuple(379, value), rsp, targets11)
	log("p4 w:%d", value)
	countwrites(value, 4, 24)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(195, &value), rsp, s4)
	log("p4 r:%d", value)
	countreads(value, 4, 13)
	l.Unlock()

	delay()
	l.Lock()
	value = 5
	Put(CreateTuple(80, value), rsp, targets12)
	log("p4 w:%d", value)
	countwrites(value, 4, 5)
	l.Unlock()

	delay()
	l.Lock()
	value = 25
	Put(CreateTuple(399, value), rsp, targets13)
	log("p4 w:%d", value)
	countwrites(value, 4, 25)
	l.Unlock()
}

func p5() {
	targets12 := make([]string, 0)
	targets11 := make([]string, 0)
	targets10 := make([]string, 0)
	targets9 := make([]string, 0)
	targets8 := make([]string, 0)
	targets7 := make([]string, 0)
	targets6 := make([]string, 0)
	targets5 := make([]string, 0)
	targets4 := make([]string, 1)
	targets4[0] = uri[s20]
	targets3 := make([]string, 0)
	targets2 := make([]string, 0)
	targets1 := make([]string, 0)
	targets0 := make([]string, 0)

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 14
	Put(CreateTuple(212, value), rsp, targets0)
	log("p5 w:%d", value)
	countwrites(value, 5, 14)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(47, &value), rsp, s5)
	log("p5 r:%d", value)
	countreads(value, 5, 3)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(22, &value), rsp, s5)
	log("p5 r:%d", value)
	countreads(value, 5, 2)
	l.Unlock()

	delay()
	l.Lock()
	value = 25
	Put(CreateTuple(398, value), rsp, targets1)
	log("p5 w:%d", value)
	countwrites(value, 5, 25)
	l.Unlock()

	delay()
	l.Lock()
	value = 25
	Put(CreateTuple(388, value), rsp, targets2)
	log("p5 w:%d", value)
	countwrites(value, 5, 25)
	l.Unlock()

	delay()
	l.Lock()
	value = 6
	Put(CreateTuple(86, value), rsp, targets3)
	log("p5 w:%d", value)
	countwrites(value, 5, 6)
	l.Unlock()

	delay()
	l.Lock()
	value = 4
	Put(CreateTuple(52, value), rsp, targets4)
	log("p5 w:%d", value)
	countwrites(value, 5, 4)
	l.Unlock()

	delay()
	l.Lock()
	value = 4
	Put(CreateTuple(63, value), rsp, targets5)
	log("p5 w:%d", value)
	countwrites(value, 5, 4)
	l.Unlock()

	delay()
	l.Lock()
	value = 20
	Put(CreateTuple(317, value), rsp, targets6)
	log("p5 w:%d", value)
	countwrites(value, 5, 20)
	l.Unlock()

	delay()
	l.Lock()
	value = 14
	Put(CreateTuple(215, value), rsp, targets7)
	log("p5 w:%d", value)
	countwrites(value, 5, 14)
	l.Unlock()

	delay()
	l.Lock()
	value = 15
	Put(CreateTuple(238, value), rsp, targets8)
	log("p5 w:%d", value)
	countwrites(value, 5, 15)
	l.Unlock()

	delay()
	l.Lock()
	value = 19
	Put(CreateTuple(293, value), rsp, targets9)
	log("p5 w:%d", value)
	countwrites(value, 5, 19)
	l.Unlock()

	delay()
	l.Lock()
	value = 30
	Put(CreateTuple(480, value), rsp, targets10)
	log("p5 w:%d", value)
	countwrites(value, 5, 30)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(43, &value), rsp, s5)
	log("p5 r:%d", value)
	countreads(value, 5, 3)
	l.Unlock()

	delay()
	l.Lock()
	value = 16
	Put(CreateTuple(254, value), rsp, targets11)
	log("p5 w:%d", value)
	countwrites(value, 5, 16)
	l.Unlock()

	delay()
	l.Lock()
	value = 32
	Put(CreateTuple(508, value), rsp, targets12)
	log("p5 w:%d", value)
	countwrites(value, 5, 32)
	l.Unlock()
}

func p6() {
	targets12 := make([]string, 0)
	targets11 := make([]string, 1)
	targets11[0] = uri[s25]
	targets10 := make([]string, 0)
	targets9 := make([]string, 0)
	targets8 := make([]string, 0)
	targets7 := make([]string, 0)
	targets6 := make([]string, 0)
	targets5 := make([]string, 2)
	targets5[1] = uri[s11]
	targets5[0] = uri[s9]
	targets4 := make([]string, 0)
	targets3 := make([]string, 0)
	targets2 := make([]string, 0)
	targets1 := make([]string, 1)
	targets1[0] = uri[s23]
	targets0 := make([]string, 0)

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(334, &value), rsp, s6)
	log("p6 r:%d", value)
	countreads(value, 6, 21)
	l.Unlock()

	delay()
	l.Lock()
	value = 5
	Put(CreateTuple(77, value), rsp, targets0)
	log("p6 w:%d", value)
	countwrites(value, 6, 5)
	l.Unlock()

	delay()
	l.Lock()
	value = 30
	Put(CreateTuple(465, value), rsp, targets1)
	log("p6 w:%d", value)
	countwrites(value, 6, 30)
	l.Unlock()

	delay()
	l.Lock()
	value = 21
	Put(CreateTuple(330, value), rsp, targets2)
	log("p6 w:%d", value)
	countwrites(value, 6, 21)
	l.Unlock()

	delay()
	l.Lock()
	value = 27
	Put(CreateTuple(427, value), rsp, targets3)
	log("p6 w:%d", value)
	countwrites(value, 6, 27)
	l.Unlock()

	delay()
	l.Lock()
	value = 20
	Put(CreateTuple(307, value), rsp, targets4)
	log("p6 w:%d", value)
	countwrites(value, 6, 20)
	l.Unlock()

	delay()
	l.Lock()
	value = 8
	Put(CreateTuple(117, value), rsp, targets5)
	log("p6 w:%d", value)
	countwrites(value, 6, 8)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(268, &value), rsp, s6)
	log("p6 r:%d", value)
	countreads(value, 6, 17)
	l.Unlock()

	delay()
	l.Lock()
	value = 10
	Put(CreateTuple(157, value), rsp, targets6)
	log("p6 w:%d", value)
	countwrites(value, 6, 10)
	l.Unlock()

	delay()
	l.Lock()
	value = 5
	Put(CreateTuple(65, value), rsp, targets7)
	log("p6 w:%d", value)
	countwrites(value, 6, 5)
	l.Unlock()

	delay()
	l.Lock()
	value = 11
	Put(CreateTuple(163, value), rsp, targets8)
	log("p6 w:%d", value)
	countwrites(value, 6, 11)
	l.Unlock()

	delay()
	l.Lock()
	value = 2
	Put(CreateTuple(27, value), rsp, targets9)
	log("p6 w:%d", value)
	countwrites(value, 6, 2)
	l.Unlock()

	delay()
	l.Lock()
	value = 20
	Put(CreateTuple(315, value), rsp, targets10)
	log("p6 w:%d", value)
	countwrites(value, 6, 20)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(448, &value), rsp, s6)
	log("p6 r:%d", value)
	countreads(value, 6, 28)
	l.Unlock()

	delay()
	l.Lock()
	value = 26
	Put(CreateTuple(411, value), rsp, targets11)
	log("p6 w:%d", value)
	countwrites(value, 6, 26)
	l.Unlock()

	delay()
	l.Lock()
	value = 3
	Put(CreateTuple(40, value), rsp, targets12)
	log("p6 w:%d", value)
	countwrites(value, 6, 3)
	l.Unlock()
}

func p7() {
	targets13 := make([]string, 0)
	targets12 := make([]string, 1)
	targets12[0] = uri[s23]
	targets11 := make([]string, 0)
	targets10 := make([]string, 1)
	targets10[0] = uri[s23]
	targets9 := make([]string, 1)
	targets9[0] = uri[s27]
	targets8 := make([]string, 0)
	targets7 := make([]string, 0)
	targets6 := make([]string, 1)
	targets6[0] = uri[s3]
	targets5 := make([]string, 2)
	targets5[1] = uri[s23]
	targets5[0] = uri[s17]
	targets4 := make([]string, 0)
	targets3 := make([]string, 0)
	targets2 := make([]string, 0)
	targets1 := make([]string, 0)
	targets0 := make([]string, 1)
	targets0[0] = uri[s19]

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 29
	Put(CreateTuple(450, value), rsp, targets0)
	log("p7 w:%d", value)
	countwrites(value, 7, 29)
	l.Unlock()

	delay()
	l.Lock()
	value = 15
	Put(CreateTuple(232, value), rsp, targets1)
	log("p7 w:%d", value)
	countwrites(value, 7, 15)
	l.Unlock()

	delay()
	l.Lock()
	value = 4
	Put(CreateTuple(49, value), rsp, targets2)
	log("p7 w:%d", value)
	countwrites(value, 7, 4)
	l.Unlock()

	delay()
	l.Lock()
	value = 27
	Put(CreateTuple(421, value), rsp, targets3)
	log("p7 w:%d", value)
	countwrites(value, 7, 27)
	l.Unlock()

	delay()
	l.Lock()
	value = 30
	Put(CreateTuple(469, value), rsp, targets4)
	log("p7 w:%d", value)
	countwrites(value, 7, 30)
	l.Unlock()

	delay()
	l.Lock()
	value = 7
	Put(CreateTuple(99, value), rsp, targets5)
	log("p7 w:%d", value)
	countwrites(value, 7, 7)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(324, &value), rsp, s7)
	log("p7 r:%d", value)
	countreads(value, 7, 21)
	l.Unlock()

	delay()
	l.Lock()
	value = 25
	Put(CreateTuple(386, value), rsp, targets6)
	log("p7 w:%d", value)
	countwrites(value, 7, 25)
	l.Unlock()

	delay()
	l.Lock()
	value = 4
	Put(CreateTuple(63, value), rsp, targets7)
	log("p7 w:%d", value)
	countwrites(value, 7, 4)
	l.Unlock()

	delay()
	l.Lock()
	value = 10
	Put(CreateTuple(160, value), rsp, targets8)
	log("p7 w:%d", value)
	countwrites(value, 7, 10)
	l.Unlock()

	delay()
	l.Lock()
	value = 3
	Put(CreateTuple(36, value), rsp, targets9)
	log("p7 w:%d", value)
	countwrites(value, 7, 3)
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
	value = 30
	Put(CreateTuple(465, value), rsp, targets10)
	log("p7 w:%d", value)
	countwrites(value, 7, 30)
	l.Unlock()

	delay()
	l.Lock()
	value = 4
	Put(CreateTuple(56, value), rsp, targets11)
	log("p7 w:%d", value)
	countwrites(value, 7, 4)
	l.Unlock()

	delay()
	l.Lock()
	value = 7
	Put(CreateTuple(108, value), rsp, targets12)
	log("p7 w:%d", value)
	countwrites(value, 7, 7)
	l.Unlock()

	delay()
	l.Lock()
	value = 30
	Put(CreateTuple(480, value), rsp, targets13)
	log("p7 w:%d", value)
	countwrites(value, 7, 30)
	l.Unlock()
}

func p8() {
	targets13 := make([]string, 0)
	targets12 := make([]string, 0)
	targets11 := make([]string, 1)
	targets11[0] = uri[s21]
	targets10 := make([]string, 0)
	targets9 := make([]string, 0)
	targets8 := make([]string, 0)
	targets7 := make([]string, 0)
	targets6 := make([]string, 0)
	targets5 := make([]string, 0)
	targets4 := make([]string, 2)
	targets4[1] = uri[s11]
	targets4[0] = uri[s9]
	targets3 := make([]string, 1)
	targets3[0] = uri[s23]
	targets2 := make([]string, 0)
	targets1 := make([]string, 0)
	targets0 := make([]string, 0)

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 9
	Put(CreateTuple(144, value), rsp, targets0)
	log("p8 w:%d", value)
	countwrites(value, 8, 9)
	l.Unlock()

	delay()
	l.Lock()
	value = 9
	Put(CreateTuple(132, value), rsp, targets1)
	log("p8 w:%d", value)
	countwrites(value, 8, 9)
	l.Unlock()

	delay()
	l.Lock()
	value = 15
	Put(CreateTuple(239, value), rsp, targets2)
	log("p8 w:%d", value)
	countwrites(value, 8, 15)
	l.Unlock()

	delay()
	l.Lock()
	value = 30
	Put(CreateTuple(465, value), rsp, targets3)
	log("p8 w:%d", value)
	countwrites(value, 8, 30)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(373, &value), rsp, s8)
	log("p8 r:%d", value)
	countreads(value, 8, 24)
	l.Unlock()

	delay()
	l.Lock()
	value = 8
	Put(CreateTuple(117, value), rsp, targets4)
	log("p8 w:%d", value)
	countwrites(value, 8, 8)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(111, &value), rsp, s8)
	log("p8 r:%d", value)
	countreads(value, 8, 7)
	l.Unlock()

	delay()
	l.Lock()
	value = 9
	Put(CreateTuple(134, value), rsp, targets5)
	log("p8 w:%d", value)
	countwrites(value, 8, 9)
	l.Unlock()

	delay()
	l.Lock()
	value = 28
	Put(CreateTuple(441, value), rsp, targets6)
	log("p8 w:%d", value)
	countwrites(value, 8, 28)
	l.Unlock()

	delay()
	l.Lock()
	value = 2
	Put(CreateTuple(32, value), rsp, targets7)
	log("p8 w:%d", value)
	countwrites(value, 8, 2)
	l.Unlock()

	delay()
	l.Lock()
	value = 20
	Put(CreateTuple(307, value), rsp, targets8)
	log("p8 w:%d", value)
	countwrites(value, 8, 20)
	l.Unlock()

	delay()
	l.Lock()
	value = 25
	Put(CreateTuple(392, value), rsp, targets9)
	log("p8 w:%d", value)
	countwrites(value, 8, 25)
	l.Unlock()

	delay()
	l.Lock()
	value = 17
	Put(CreateTuple(261, value), rsp, targets10)
	log("p8 w:%d", value)
	countwrites(value, 8, 17)
	l.Unlock()

	delay()
	l.Lock()
	value = 17
	Put(CreateTuple(266, value), rsp, targets11)
	log("p8 w:%d", value)
	countwrites(value, 8, 17)
	l.Unlock()

	delay()
	l.Lock()
	value = 31
	Put(CreateTuple(493, value), rsp, targets12)
	log("p8 w:%d", value)
	countwrites(value, 8, 31)
	l.Unlock()

	delay()
	l.Lock()
	value = 8
	Put(CreateTuple(126, value), rsp, targets13)
	log("p8 w:%d", value)
	countwrites(value, 8, 8)
	l.Unlock()
}

func p9() {
	targets11 := make([]string, 0)
	targets10 := make([]string, 0)
	targets9 := make([]string, 0)
	targets8 := make([]string, 1)
	targets8[0] = uri[s16]
	targets7 := make([]string, 0)
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
	value = -1
	QueryP(CreateTuple(117, &value), rsp, s9)
	log("p9 r:%d", value)
	countreads(value, 9, 8)
	l.Unlock()

	delay()
	l.Lock()
	value = 29
	Put(CreateTuple(458, value), rsp, targets0)
	log("p9 w:%d", value)
	countwrites(value, 9, 29)
	l.Unlock()

	delay()
	l.Lock()
	value = 26
	Put(CreateTuple(415, value), rsp, targets1)
	log("p9 w:%d", value)
	countwrites(value, 9, 26)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(488, &value), rsp, s9)
	log("p9 r:%d", value)
	countreads(value, 9, 31)
	l.Unlock()

	delay()
	l.Lock()
	value = 3
	Put(CreateTuple(35, value), rsp, targets2)
	log("p9 w:%d", value)
	countwrites(value, 9, 3)
	l.Unlock()

	delay()
	l.Lock()
	value = 30
	Put(CreateTuple(475, value), rsp, targets3)
	log("p9 w:%d", value)
	countwrites(value, 9, 30)
	l.Unlock()

	delay()
	l.Lock()
	value = 23
	Put(CreateTuple(359, value), rsp, targets4)
	log("p9 w:%d", value)
	countwrites(value, 9, 23)
	l.Unlock()

	delay()
	l.Lock()
	value = 1
	Put(CreateTuple(15, value), rsp, targets5)
	log("p9 w:%d", value)
	countwrites(value, 9, 1)
	l.Unlock()

	delay()
	l.Lock()
	value = 17
	Put(CreateTuple(269, value), rsp, targets6)
	log("p9 w:%d", value)
	countwrites(value, 9, 17)
	l.Unlock()

	delay()
	l.Lock()
	value = 17
	Put(CreateTuple(267, value), rsp, targets7)
	log("p9 w:%d", value)
	countwrites(value, 9, 17)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(504, &value), rsp, s9)
	log("p9 r:%d", value)
	countreads(value, 9, 32)
	l.Unlock()

	delay()
	l.Lock()
	value = 7
	Put(CreateTuple(104, value), rsp, targets8)
	log("p9 w:%d", value)
	countwrites(value, 9, 7)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(186, &value), rsp, s9)
	log("p9 r:%d", value)
	countreads(value, 9, 12)
	l.Unlock()

	delay()
	l.Lock()
	value = 17
	Put(CreateTuple(261, value), rsp, targets9)
	log("p9 w:%d", value)
	countwrites(value, 9, 17)
	l.Unlock()

	delay()
	l.Lock()
	value = 10
	Put(CreateTuple(150, value), rsp, targets10)
	log("p9 w:%d", value)
	countwrites(value, 9, 10)
	l.Unlock()

	delay()
	l.Lock()
	value = 19
	Put(CreateTuple(293, value), rsp, targets11)
	log("p9 w:%d", value)
	countwrites(value, 9, 19)
	l.Unlock()
}

func p10() {
	targets12 := make([]string, 1)
	targets12[0] = uri[s9]
	targets11 := make([]string, 0)
	targets10 := make([]string, 1)
	targets10[0] = uri[s19]
	targets9 := make([]string, 0)
	targets8 := make([]string, 0)
	targets7 := make([]string, 2)
	targets7[1] = uri[s21]
	targets7[0] = uri[s3]
	targets6 := make([]string, 0)
	targets5 := make([]string, 0)
	targets4 := make([]string, 1)
	targets4[0] = uri[s16]
	targets3 := make([]string, 0)
	targets2 := make([]string, 0)
	targets1 := make([]string, 0)
	targets0 := make([]string, 0)

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(26, &value), rsp, s10)
	log("p10 r:%d", value)
	countreads(value, 10, 2)
	l.Unlock()

	delay()
	l.Lock()
	value = 6
	Put(CreateTuple(95, value), rsp, targets0)
	log("p10 w:%d", value)
	countwrites(value, 10, 6)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(213, &value), rsp, s10)
	log("p10 r:%d", value)
	countreads(value, 10, 14)
	l.Unlock()

	delay()
	l.Lock()
	value = 31
	Put(CreateTuple(496, value), rsp, targets1)
	log("p10 w:%d", value)
	countwrites(value, 10, 31)
	l.Unlock()

	delay()
	l.Lock()
	value = 16
	Put(CreateTuple(243, value), rsp, targets2)
	log("p10 w:%d", value)
	countwrites(value, 10, 16)
	l.Unlock()

	delay()
	l.Lock()
	value = 25
	Put(CreateTuple(390, value), rsp, targets3)
	log("p10 w:%d", value)
	countwrites(value, 10, 25)
	l.Unlock()

	delay()
	l.Lock()
	value = 17
	Put(CreateTuple(260, value), rsp, targets4)
	log("p10 w:%d", value)
	countwrites(value, 10, 17)
	l.Unlock()

	delay()
	l.Lock()
	value = 24
	Put(CreateTuple(379, value), rsp, targets5)
	log("p10 w:%d", value)
	countwrites(value, 10, 24)
	l.Unlock()

	delay()
	l.Lock()
	value = 5
	Put(CreateTuple(79, value), rsp, targets6)
	log("p10 w:%d", value)
	countwrites(value, 10, 5)
	l.Unlock()

	delay()
	l.Lock()
	value = 29
	Put(CreateTuple(460, value), rsp, targets7)
	log("p10 w:%d", value)
	countwrites(value, 10, 29)
	l.Unlock()

	delay()
	l.Lock()
	value = 22
	Put(CreateTuple(338, value), rsp, targets8)
	log("p10 w:%d", value)
	countwrites(value, 10, 22)
	l.Unlock()

	delay()
	l.Lock()
	value = 25
	Put(CreateTuple(387, value), rsp, targets9)
	log("p10 w:%d", value)
	countwrites(value, 10, 25)
	l.Unlock()

	delay()
	l.Lock()
	value = 26
	Put(CreateTuple(413, value), rsp, targets10)
	log("p10 w:%d", value)
	countwrites(value, 10, 26)
	l.Unlock()

	delay()
	l.Lock()
	value = 1
	Put(CreateTuple(16, value), rsp, targets11)
	log("p10 w:%d", value)
	countwrites(value, 10, 1)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(477, &value), rsp, s10)
	log("p10 r:%d", value)
	countreads(value, 10, 30)
	l.Unlock()

	delay()
	l.Lock()
	value = 12
	Put(CreateTuple(186, value), rsp, targets12)
	log("p10 w:%d", value)
	countwrites(value, 10, 12)
	l.Unlock()
}

func p11() {
	targets9 := make([]string, 0)
	targets8 := make([]string, 0)
	targets7 := make([]string, 0)
	targets6 := make([]string, 0)
	targets5 := make([]string, 1)
	targets5[0] = uri[s1]
	targets4 := make([]string, 0)
	targets3 := make([]string, 0)
	targets2 := make([]string, 0)
	targets1 := make([]string, 0)
	targets0 := make([]string, 0)

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 27
	Put(CreateTuple(421, value), rsp, targets0)
	log("p11 w:%d", value)
	countwrites(value, 11, 27)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(482, &value), rsp, s11)
	log("p11 r:%d", value)
	countreads(value, 11, 31)
	l.Unlock()

	delay()
	l.Lock()
	value = 11
	Put(CreateTuple(173, value), rsp, targets1)
	log("p11 w:%d", value)
	countwrites(value, 11, 11)
	l.Unlock()

	delay()
	l.Lock()
	value = 10
	Put(CreateTuple(154, value), rsp, targets2)
	log("p11 w:%d", value)
	countwrites(value, 11, 10)
	l.Unlock()

	delay()
	l.Lock()
	value = 30
	Put(CreateTuple(469, value), rsp, targets3)
	log("p11 w:%d", value)
	countwrites(value, 11, 30)
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
	QueryP(CreateTuple(80, &value), rsp, s11)
	log("p11 r:%d", value)
	countreads(value, 11, 5)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(280, &value), rsp, s11)
	log("p11 r:%d", value)
	countreads(value, 11, 18)
	l.Unlock()

	delay()
	l.Lock()
	value = 17
	Put(CreateTuple(272, value), rsp, targets4)
	log("p11 w:%d", value)
	countwrites(value, 11, 17)
	l.Unlock()

	delay()
	l.Lock()
	value = 13
	Put(CreateTuple(193, value), rsp, targets5)
	log("p11 w:%d", value)
	countwrites(value, 11, 13)
	l.Unlock()

	delay()
	l.Lock()
	value = 13
	Put(CreateTuple(202, value), rsp, targets6)
	log("p11 w:%d", value)
	countwrites(value, 11, 13)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(113, &value), rsp, s11)
	log("p11 r:%d", value)
	countreads(value, 11, 8)
	l.Unlock()

	delay()
	l.Lock()
	value = 31
	Put(CreateTuple(493, value), rsp, targets7)
	log("p11 w:%d", value)
	countwrites(value, 11, 31)
	l.Unlock()

	delay()
	l.Lock()
	value = 23
	Put(CreateTuple(358, value), rsp, targets8)
	log("p11 w:%d", value)
	countwrites(value, 11, 23)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(420, &value), rsp, s11)
	log("p11 r:%d", value)
	countreads(value, 11, 27)
	l.Unlock()

	delay()
	l.Lock()
	value = 18
	Put(CreateTuple(281, value), rsp, targets9)
	log("p11 w:%d", value)
	countwrites(value, 11, 18)
	l.Unlock()
}

func p12() {
	targets14 := make([]string, 0)
	targets13 := make([]string, 0)
	targets12 := make([]string, 0)
	targets11 := make([]string, 0)
	targets10 := make([]string, 0)
	targets9 := make([]string, 0)
	targets8 := make([]string, 0)
	targets7 := make([]string, 1)
	targets7[0] = uri[s4]
	targets6 := make([]string, 0)
	targets5 := make([]string, 0)
	targets4 := make([]string, 1)
	targets4[0] = uri[s23]
	targets3 := make([]string, 0)
	targets2 := make([]string, 0)
	targets1 := make([]string, 1)
	targets1[0] = uri[s10]
	targets0 := make([]string, 1)
	targets0[0] = uri[s32]

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 11
	Put(CreateTuple(169, value), rsp, targets0)
	log("p12 w:%d", value)
	countwrites(value, 12, 11)
	l.Unlock()

	delay()
	l.Lock()
	value = 30
	Put(CreateTuple(477, value), rsp, targets1)
	log("p12 w:%d", value)
	countwrites(value, 12, 30)
	l.Unlock()

	delay()
	l.Lock()
	value = 2
	Put(CreateTuple(29, value), rsp, targets2)
	log("p12 w:%d", value)
	countwrites(value, 12, 2)
	l.Unlock()

	delay()
	l.Lock()
	value = 11
	Put(CreateTuple(174, value), rsp, targets3)
	log("p12 w:%d", value)
	countwrites(value, 12, 11)
	l.Unlock()

	delay()
	l.Lock()
	value = 7
	Put(CreateTuple(108, value), rsp, targets4)
	log("p12 w:%d", value)
	countwrites(value, 12, 7)
	l.Unlock()

	delay()
	l.Lock()
	value = 12
	Put(CreateTuple(179, value), rsp, targets5)
	log("p12 w:%d", value)
	countwrites(value, 12, 12)
	l.Unlock()

	delay()
	l.Lock()
	value = 22
	Put(CreateTuple(338, value), rsp, targets6)
	log("p12 w:%d", value)
	countwrites(value, 12, 22)
	l.Unlock()

	delay()
	l.Lock()
	value = 13
	Put(CreateTuple(200, value), rsp, targets7)
	log("p12 w:%d", value)
	countwrites(value, 12, 13)
	l.Unlock()

	delay()
	l.Lock()
	value = 23
	Put(CreateTuple(368, value), rsp, targets8)
	log("p12 w:%d", value)
	countwrites(value, 12, 23)
	l.Unlock()

	delay()
	l.Lock()
	value = 26
	Put(CreateTuple(405, value), rsp, targets9)
	log("p12 w:%d", value)
	countwrites(value, 12, 26)
	l.Unlock()

	delay()
	l.Lock()
	value = 29
	Put(CreateTuple(451, value), rsp, targets10)
	log("p12 w:%d", value)
	countwrites(value, 12, 29)
	l.Unlock()

	delay()
	l.Lock()
	value = 23
	Put(CreateTuple(365, value), rsp, targets11)
	log("p12 w:%d", value)
	countwrites(value, 12, 23)
	l.Unlock()

	delay()
	l.Lock()
	value = 24
	Put(CreateTuple(377, value), rsp, targets12)
	log("p12 w:%d", value)
	countwrites(value, 12, 24)
	l.Unlock()

	delay()
	l.Lock()
	value = 28
	Put(CreateTuple(446, value), rsp, targets13)
	log("p12 w:%d", value)
	countwrites(value, 12, 28)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(376, &value), rsp, s12)
	log("p12 r:%d", value)
	countreads(value, 12, 24)
	l.Unlock()

	delay()
	l.Lock()
	value = 21
	Put(CreateTuple(331, value), rsp, targets14)
	log("p12 w:%d", value)
	countwrites(value, 12, 21)
	l.Unlock()
}

func p13() {
	targets10 := make([]string, 0)
	targets9 := make([]string, 0)
	targets8 := make([]string, 0)
	targets7 := make([]string, 1)
	targets7[0] = uri[s27]
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
	value = -1
	QueryP(CreateTuple(135, &value), rsp, s13)
	log("p13 r:%d", value)
	countreads(value, 13, 9)
	l.Unlock()

	delay()
	l.Lock()
	value = 22
	Put(CreateTuple(342, value), rsp, targets0)
	log("p13 w:%d", value)
	countwrites(value, 13, 22)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(135, &value), rsp, s13)
	log("p13 r:%d", value)
	countreads(value, 13, 9)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(314, &value), rsp, s13)
	log("p13 r:%d", value)
	countreads(value, 13, 20)
	l.Unlock()

	delay()
	l.Lock()
	value = 6
	Put(CreateTuple(87, value), rsp, targets1)
	log("p13 w:%d", value)
	countwrites(value, 13, 6)
	l.Unlock()

	delay()
	l.Lock()
	value = 20
	Put(CreateTuple(316, value), rsp, targets2)
	log("p13 w:%d", value)
	countwrites(value, 13, 20)
	l.Unlock()

	delay()
	l.Lock()
	value = 20
	Put(CreateTuple(313, value), rsp, targets3)
	log("p13 w:%d", value)
	countwrites(value, 13, 20)
	l.Unlock()

	delay()
	l.Lock()
	value = 26
	Put(CreateTuple(402, value), rsp, targets4)
	log("p13 w:%d", value)
	countwrites(value, 13, 26)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(170, &value), rsp, s13)
	log("p13 r:%d", value)
	countreads(value, 13, 11)
	l.Unlock()

	delay()
	l.Lock()
	value = 17
	Put(CreateTuple(258, value), rsp, targets5)
	log("p13 w:%d", value)
	countwrites(value, 13, 17)
	l.Unlock()

	delay()
	l.Lock()
	value = 23
	Put(CreateTuple(367, value), rsp, targets6)
	log("p13 w:%d", value)
	countwrites(value, 13, 23)
	l.Unlock()

	delay()
	l.Lock()
	value = 18
	Put(CreateTuple(287, value), rsp, targets7)
	log("p13 w:%d", value)
	countwrites(value, 13, 18)
	l.Unlock()

	delay()
	l.Lock()
	value = 21
	Put(CreateTuple(321, value), rsp, targets8)
	log("p13 w:%d", value)
	countwrites(value, 13, 21)
	l.Unlock()

	delay()
	l.Lock()
	value = 10
	Put(CreateTuple(148, value), rsp, targets9)
	log("p13 w:%d", value)
	countwrites(value, 13, 10)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(23, &value), rsp, s13)
	log("p13 r:%d", value)
	countreads(value, 13, 2)
	l.Unlock()

	delay()
	l.Lock()
	value = 32
	Put(CreateTuple(512, value), rsp, targets10)
	log("p13 w:%d", value)
	countwrites(value, 13, 32)
	l.Unlock()
}

func p14() {
	targets13 := make([]string, 0)
	targets12 := make([]string, 0)
	targets11 := make([]string, 2)
	targets11[1] = uri[s13]
	targets11[0] = uri[s13]
	targets10 := make([]string, 1)
	targets10[0] = uri[s27]
	targets9 := make([]string, 1)
	targets9[0] = uri[s26]
	targets8 := make([]string, 0)
	targets7 := make([]string, 1)
	targets7[0] = uri[s13]
	targets6 := make([]string, 1)
	targets6[0] = uri[s21]
	targets5 := make([]string, 2)
	targets5[1] = uri[s32]
	targets5[0] = uri[s11]
	targets4 := make([]string, 0)
	targets3 := make([]string, 0)
	targets2 := make([]string, 1)
	targets2[0] = uri[s27]
	targets1 := make([]string, 0)
	targets0 := make([]string, 1)
	targets0[0] = uri[s6]

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 17
	Put(CreateTuple(268, value), rsp, targets0)
	log("p14 w:%d", value)
	countwrites(value, 14, 17)
	l.Unlock()

	delay()
	l.Lock()
	value = 15
	Put(CreateTuple(232, value), rsp, targets1)
	log("p14 w:%d", value)
	countwrites(value, 14, 15)
	l.Unlock()

	delay()
	l.Lock()
	value = 1
	Put(CreateTuple(13, value), rsp, targets2)
	log("p14 w:%d", value)
	countwrites(value, 14, 1)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(384, &value), rsp, s14)
	log("p14 r:%d", value)
	countreads(value, 14, 24)
	l.Unlock()

	delay()
	l.Lock()
	value = 29
	Put(CreateTuple(458, value), rsp, targets3)
	log("p14 w:%d", value)
	countwrites(value, 14, 29)
	l.Unlock()

	delay()
	l.Lock()
	value = 32
	Put(CreateTuple(505, value), rsp, targets4)
	log("p14 w:%d", value)
	countwrites(value, 14, 32)
	l.Unlock()

	delay()
	l.Lock()
	value = 31
	Put(CreateTuple(482, value), rsp, targets5)
	log("p14 w:%d", value)
	countwrites(value, 14, 31)
	l.Unlock()

	delay()
	l.Lock()
	value = 17
	Put(CreateTuple(266, value), rsp, targets6)
	log("p14 w:%d", value)
	countwrites(value, 14, 17)
	l.Unlock()

	delay()
	l.Lock()
	value = 11
	Put(CreateTuple(170, value), rsp, targets7)
	log("p14 w:%d", value)
	countwrites(value, 14, 11)
	l.Unlock()

	delay()
	l.Lock()
	value = 14
	Put(CreateTuple(211, value), rsp, targets8)
	log("p14 w:%d", value)
	countwrites(value, 14, 14)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(162, &value), rsp, s14)
	log("p14 r:%d", value)
	countreads(value, 14, 11)
	l.Unlock()

	delay()
	l.Lock()
	value = 27
	Put(CreateTuple(418, value), rsp, targets9)
	log("p14 w:%d", value)
	countwrites(value, 14, 27)
	l.Unlock()

	delay()
	l.Lock()
	value = 18
	Put(CreateTuple(287, value), rsp, targets10)
	log("p14 w:%d", value)
	countwrites(value, 14, 18)
	l.Unlock()

	delay()
	l.Lock()
	value = 9
	Put(CreateTuple(135, value), rsp, targets11)
	log("p14 w:%d", value)
	countwrites(value, 14, 9)
	l.Unlock()

	delay()
	l.Lock()
	value = 20
	Put(CreateTuple(319, value), rsp, targets12)
	log("p14 w:%d", value)
	countwrites(value, 14, 20)
	l.Unlock()

	delay()
	l.Lock()
	value = 5
	Put(CreateTuple(67, value), rsp, targets13)
	log("p14 w:%d", value)
	countwrites(value, 14, 5)
	l.Unlock()
}

func p15() {
	targets11 := make([]string, 0)
	targets10 := make([]string, 0)
	targets9 := make([]string, 0)
	targets8 := make([]string, 0)
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
	Put(CreateTuple(165, value), rsp, targets0)
	log("p15 w:%d", value)
	countwrites(value, 15, 11)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(448, &value), rsp, s15)
	log("p15 r:%d", value)
	countreads(value, 15, 28)
	l.Unlock()

	delay()
	l.Lock()
	value = 26
	Put(CreateTuple(410, value), rsp, targets1)
	log("p15 w:%d", value)
	countwrites(value, 15, 26)
	l.Unlock()

	delay()
	l.Lock()
	value = 4
	Put(CreateTuple(55, value), rsp, targets2)
	log("p15 w:%d", value)
	countwrites(value, 15, 4)
	l.Unlock()

	delay()
	l.Lock()
	value = 9
	Put(CreateTuple(141, value), rsp, targets3)
	log("p15 w:%d", value)
	countwrites(value, 15, 9)
	l.Unlock()

	delay()
	l.Lock()
	value = 5
	Put(CreateTuple(75, value), rsp, targets4)
	log("p15 w:%d", value)
	countwrites(value, 15, 5)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(71, &value), rsp, s15)
	log("p15 r:%d", value)
	countreads(value, 15, 5)
	l.Unlock()

	delay()
	l.Lock()
	value = 14
	Put(CreateTuple(219, value), rsp, targets5)
	log("p15 w:%d", value)
	countwrites(value, 15, 14)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(270, &value), rsp, s15)
	log("p15 r:%d", value)
	countreads(value, 15, 17)
	l.Unlock()

	delay()
	l.Lock()
	value = 20
	Put(CreateTuple(307, value), rsp, targets6)
	log("p15 w:%d", value)
	countwrites(value, 15, 20)
	l.Unlock()

	delay()
	l.Lock()
	value = 16
	Put(CreateTuple(255, value), rsp, targets7)
	log("p15 w:%d", value)
	countwrites(value, 15, 16)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(125, &value), rsp, s15)
	log("p15 r:%d", value)
	countreads(value, 15, 8)
	l.Unlock()

	delay()
	l.Lock()
	value = 6
	Put(CreateTuple(86, value), rsp, targets8)
	log("p15 w:%d", value)
	countwrites(value, 15, 6)
	l.Unlock()

	delay()
	l.Lock()
	value = 4
	Put(CreateTuple(62, value), rsp, targets9)
	log("p15 w:%d", value)
	countwrites(value, 15, 4)
	l.Unlock()

	delay()
	l.Lock()
	value = 18
	Put(CreateTuple(281, value), rsp, targets10)
	log("p15 w:%d", value)
	countwrites(value, 15, 18)
	l.Unlock()

	delay()
	l.Lock()
	value = 23
	Put(CreateTuple(355, value), rsp, targets11)
	log("p15 w:%d", value)
	countwrites(value, 15, 23)
	l.Unlock()
}

func p16() {
	targets12 := make([]string, 0)
	targets11 := make([]string, 1)
	targets11[0] = uri[s13]
	targets10 := make([]string, 0)
	targets9 := make([]string, 0)
	targets8 := make([]string, 1)
	targets8[0] = uri[s10]
	targets7 := make([]string, 0)
	targets6 := make([]string, 0)
	targets5 := make([]string, 1)
	targets5[0] = uri[s25]
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
	QueryP(CreateTuple(104, &value), rsp, s16)
	log("p16 r:%d", value)
	countreads(value, 16, 7)
	l.Unlock()

	delay()
	l.Lock()
	value = 29
	Put(CreateTuple(464, value), rsp, targets0)
	log("p16 w:%d", value)
	countwrites(value, 16, 29)
	l.Unlock()

	delay()
	l.Lock()
	value = 11
	Put(CreateTuple(174, value), rsp, targets1)
	log("p16 w:%d", value)
	countwrites(value, 16, 11)
	l.Unlock()

	delay()
	l.Lock()
	value = 2
	Put(CreateTuple(19, value), rsp, targets2)
	log("p16 w:%d", value)
	countwrites(value, 16, 2)
	l.Unlock()

	delay()
	l.Lock()
	value = 29
	Put(CreateTuple(459, value), rsp, targets3)
	log("p16 w:%d", value)
	countwrites(value, 16, 29)
	l.Unlock()

	delay()
	l.Lock()
	value = 19
	Put(CreateTuple(295, value), rsp, targets4)
	log("p16 w:%d", value)
	countwrites(value, 16, 19)
	l.Unlock()

	delay()
	l.Lock()
	value = 29
	Put(CreateTuple(456, value), rsp, targets5)
	log("p16 w:%d", value)
	countwrites(value, 16, 29)
	l.Unlock()

	delay()
	l.Lock()
	value = 2
	Put(CreateTuple(20, value), rsp, targets6)
	log("p16 w:%d", value)
	countwrites(value, 16, 2)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(417, &value), rsp, s16)
	log("p16 r:%d", value)
	countreads(value, 16, 27)
	l.Unlock()

	delay()
	l.Lock()
	value = 28
	Put(CreateTuple(439, value), rsp, targets7)
	log("p16 w:%d", value)
	countwrites(value, 16, 28)
	l.Unlock()

	delay()
	l.Lock()
	value = 30
	Put(CreateTuple(477, value), rsp, targets8)
	log("p16 w:%d", value)
	countwrites(value, 16, 30)
	l.Unlock()

	delay()
	l.Lock()
	value = 24
	Put(CreateTuple(378, value), rsp, targets9)
	log("p16 w:%d", value)
	countwrites(value, 16, 24)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(260, &value), rsp, s16)
	log("p16 r:%d", value)
	countreads(value, 16, 17)
	l.Unlock()

	delay()
	l.Lock()
	value = 19
	Put(CreateTuple(292, value), rsp, targets10)
	log("p16 w:%d", value)
	countwrites(value, 16, 19)
	l.Unlock()

	delay()
	l.Lock()
	value = 11
	Put(CreateTuple(170, value), rsp, targets11)
	log("p16 w:%d", value)
	countwrites(value, 16, 11)
	l.Unlock()

	delay()
	l.Lock()
	value = 22
	Put(CreateTuple(349, value), rsp, targets12)
	log("p16 w:%d", value)
	countwrites(value, 16, 22)
	l.Unlock()
}

func p17() {
	targets13 := make([]string, 0)
	targets12 := make([]string, 1)
	targets12[0] = uri[s2]
	targets11 := make([]string, 0)
	targets10 := make([]string, 0)
	targets9 := make([]string, 0)
	targets8 := make([]string, 0)
	targets7 := make([]string, 1)
	targets7[0] = uri[s2]
	targets6 := make([]string, 0)
	targets5 := make([]string, 0)
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
	value = 20
	Put(CreateTuple(317, value), rsp, targets0)
	log("p17 w:%d", value)
	countwrites(value, 17, 20)
	l.Unlock()

	delay()
	l.Lock()
	value = 2
	Put(CreateTuple(29, value), rsp, targets1)
	log("p17 w:%d", value)
	countwrites(value, 17, 2)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(99, &value), rsp, s17)
	log("p17 r:%d", value)
	countreads(value, 17, 7)
	l.Unlock()

	delay()
	l.Lock()
	value = 4
	Put(CreateTuple(55, value), rsp, targets2)
	log("p17 w:%d", value)
	countwrites(value, 17, 4)
	l.Unlock()

	delay()
	l.Lock()
	value = 25
	Put(CreateTuple(391, value), rsp, targets3)
	log("p17 w:%d", value)
	countwrites(value, 17, 25)
	l.Unlock()

	delay()
	l.Lock()
	value = 5
	Put(CreateTuple(71, value), rsp, targets4)
	log("p17 w:%d", value)
	countwrites(value, 17, 5)
	l.Unlock()

	delay()
	l.Lock()
	value = 27
	Put(CreateTuple(427, value), rsp, targets5)
	log("p17 w:%d", value)
	countwrites(value, 17, 27)
	l.Unlock()

	delay()
	l.Lock()
	value = 2
	Put(CreateTuple(30, value), rsp, targets6)
	log("p17 w:%d", value)
	countwrites(value, 17, 2)
	l.Unlock()

	delay()
	l.Lock()
	value = 31
	Put(CreateTuple(494, value), rsp, targets7)
	log("p17 w:%d", value)
	countwrites(value, 17, 31)
	l.Unlock()

	delay()
	l.Lock()
	value = 19
	Put(CreateTuple(293, value), rsp, targets8)
	log("p17 w:%d", value)
	countwrites(value, 17, 19)
	l.Unlock()

	delay()
	l.Lock()
	value = 18
	Put(CreateTuple(286, value), rsp, targets9)
	log("p17 w:%d", value)
	countwrites(value, 17, 18)
	l.Unlock()

	delay()
	l.Lock()
	value = 26
	Put(CreateTuple(401, value), rsp, targets10)
	log("p17 w:%d", value)
	countwrites(value, 17, 26)
	l.Unlock()

	delay()
	l.Lock()
	value = 25
	Put(CreateTuple(393, value), rsp, targets11)
	log("p17 w:%d", value)
	countwrites(value, 17, 25)
	l.Unlock()

	delay()
	l.Lock()
	value = 14
	Put(CreateTuple(216, value), rsp, targets12)
	log("p17 w:%d", value)
	countwrites(value, 17, 14)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(488, &value), rsp, s17)
	log("p17 r:%d", value)
	countreads(value, 17, 31)
	l.Unlock()

	delay()
	l.Lock()
	value = 16
	Put(CreateTuple(252, value), rsp, targets13)
	log("p17 w:%d", value)
	countwrites(value, 17, 16)
	l.Unlock()
}

func p18() {
	targets13 := make([]string, 0)
	targets12 := make([]string, 0)
	targets11 := make([]string, 1)
	targets11[0] = uri[s10]
	targets10 := make([]string, 0)
	targets9 := make([]string, 0)
	targets8 := make([]string, 0)
	targets7 := make([]string, 0)
	targets6 := make([]string, 1)
	targets6[0] = uri[s25]
	targets5 := make([]string, 0)
	targets4 := make([]string, 0)
	targets3 := make([]string, 1)
	targets3[0] = uri[s3]
	targets2 := make([]string, 0)
	targets1 := make([]string, 1)
	targets1[0] = uri[s1]
	targets0 := make([]string, 0)

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 11
	Put(CreateTuple(174, value), rsp, targets0)
	log("p18 w:%d", value)
	countwrites(value, 18, 11)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(283, &value), rsp, s18)
	log("p18 r:%d", value)
	countreads(value, 18, 18)
	l.Unlock()

	delay()
	l.Lock()
	value = 26
	Put(CreateTuple(409, value), rsp, targets1)
	log("p18 w:%d", value)
	countwrites(value, 18, 26)
	l.Unlock()

	delay()
	l.Lock()
	value = 17
	Put(CreateTuple(261, value), rsp, targets2)
	log("p18 w:%d", value)
	countwrites(value, 18, 17)
	l.Unlock()

	delay()
	l.Lock()
	value = 25
	Put(CreateTuple(395, value), rsp, targets3)
	log("p18 w:%d", value)
	countwrites(value, 18, 25)
	l.Unlock()

	delay()
	l.Lock()
	value = 26
	Put(CreateTuple(403, value), rsp, targets4)
	log("p18 w:%d", value)
	countwrites(value, 18, 26)
	l.Unlock()

	delay()
	l.Lock()
	value = 18
	Put(CreateTuple(281, value), rsp, targets5)
	log("p18 w:%d", value)
	countwrites(value, 18, 18)
	l.Unlock()

	delay()
	l.Lock()
	value = 29
	Put(CreateTuple(456, value), rsp, targets6)
	log("p18 w:%d", value)
	countwrites(value, 18, 29)
	l.Unlock()

	delay()
	l.Lock()
	value = 3
	Put(CreateTuple(38, value), rsp, targets7)
	log("p18 w:%d", value)
	countwrites(value, 18, 3)
	l.Unlock()

	delay()
	l.Lock()
	value = 22
	Put(CreateTuple(348, value), rsp, targets8)
	log("p18 w:%d", value)
	countwrites(value, 18, 22)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(208, &value), rsp, s18)
	log("p18 r:%d", value)
	countreads(value, 18, 13)
	l.Unlock()

	delay()
	l.Lock()
	value = 19
	Put(CreateTuple(296, value), rsp, targets9)
	log("p18 w:%d", value)
	countwrites(value, 18, 19)
	l.Unlock()

	delay()
	l.Lock()
	value = 32
	Put(CreateTuple(499, value), rsp, targets10)
	log("p18 w:%d", value)
	countwrites(value, 18, 32)
	l.Unlock()

	delay()
	l.Lock()
	value = 2
	Put(CreateTuple(26, value), rsp, targets11)
	log("p18 w:%d", value)
	countwrites(value, 18, 2)
	l.Unlock()

	delay()
	l.Lock()
	value = 19
	Put(CreateTuple(297, value), rsp, targets12)
	log("p18 w:%d", value)
	countwrites(value, 18, 19)
	l.Unlock()

	delay()
	l.Lock()
	value = 22
	Put(CreateTuple(338, value), rsp, targets13)
	log("p18 w:%d", value)
	countwrites(value, 18, 22)
	l.Unlock()
}

func p19() {
	targets11 := make([]string, 0)
	targets10 := make([]string, 0)
	targets9 := make([]string, 1)
	targets9[0] = uri[s29]
	targets8 := make([]string, 0)
	targets7 := make([]string, 0)
	targets6 := make([]string, 0)
	targets5 := make([]string, 0)
	targets4 := make([]string, 0)
	targets3 := make([]string, 0)
	targets2 := make([]string, 1)
	targets2[0] = uri[s8]
	targets1 := make([]string, 0)
	targets0 := make([]string, 0)

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 25
	Put(CreateTuple(397, value), rsp, targets0)
	log("p19 w:%d", value)
	countwrites(value, 19, 25)
	l.Unlock()

	delay()
	l.Lock()
	value = 29
	Put(CreateTuple(463, value), rsp, targets1)
	log("p19 w:%d", value)
	countwrites(value, 19, 29)
	l.Unlock()

	delay()
	l.Lock()
	value = 7
	Put(CreateTuple(111, value), rsp, targets2)
	log("p19 w:%d", value)
	countwrites(value, 19, 7)
	l.Unlock()

	delay()
	l.Lock()
	value = 8
	Put(CreateTuple(114, value), rsp, targets3)
	log("p19 w:%d", value)
	countwrites(value, 19, 8)
	l.Unlock()

	delay()
	l.Lock()
	value = 21
	Put(CreateTuple(321, value), rsp, targets4)
	log("p19 w:%d", value)
	countwrites(value, 19, 21)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(450, &value), rsp, s19)
	log("p19 r:%d", value)
	countreads(value, 19, 29)
	l.Unlock()

	delay()
	l.Lock()
	value = 15
	Put(CreateTuple(228, value), rsp, targets5)
	log("p19 w:%d", value)
	countwrites(value, 19, 15)
	l.Unlock()

	delay()
	l.Lock()
	value = 15
	Put(CreateTuple(239, value), rsp, targets6)
	log("p19 w:%d", value)
	countwrites(value, 19, 15)
	l.Unlock()

	delay()
	l.Lock()
	value = 12
	Put(CreateTuple(189, value), rsp, targets7)
	log("p19 w:%d", value)
	countwrites(value, 19, 12)
	l.Unlock()

	delay()
	l.Lock()
	value = 7
	Put(CreateTuple(106, value), rsp, targets8)
	log("p19 w:%d", value)
	countwrites(value, 19, 7)
	l.Unlock()

	delay()
	l.Lock()
	value = 31
	Put(CreateTuple(490, value), rsp, targets9)
	log("p19 w:%d", value)
	countwrites(value, 19, 31)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(123, &value), rsp, s19)
	log("p19 r:%d", value)
	countreads(value, 19, 8)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(145, &value), rsp, s19)
	log("p19 r:%d", value)
	countreads(value, 19, 10)
	l.Unlock()

	delay()
	l.Lock()
	value = 20
	Put(CreateTuple(319, value), rsp, targets10)
	log("p19 w:%d", value)
	countwrites(value, 19, 20)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(413, &value), rsp, s19)
	log("p19 r:%d", value)
	countreads(value, 19, 26)
	l.Unlock()

	delay()
	l.Lock()
	value = 13
	Put(CreateTuple(194, value), rsp, targets11)
	log("p19 w:%d", value)
	countwrites(value, 19, 13)
	l.Unlock()
}

func p20() {
	targets12 := make([]string, 0)
	targets11 := make([]string, 2)
	targets11[1] = uri[s21]
	targets11[0] = uri[s3]
	targets10 := make([]string, 0)
	targets9 := make([]string, 0)
	targets8 := make([]string, 1)
	targets8[0] = uri[s6]
	targets7 := make([]string, 1)
	targets7[0] = uri[s15]
	targets6 := make([]string, 0)
	targets5 := make([]string, 0)
	targets4 := make([]string, 0)
	targets3 := make([]string, 0)
	targets2 := make([]string, 0)
	targets1 := make([]string, 1)
	targets1[0] = uri[s27]
	targets0 := make([]string, 0)

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 32
	Put(CreateTuple(501, value), rsp, targets0)
	log("p20 w:%d", value)
	countwrites(value, 20, 32)
	l.Unlock()

	delay()
	l.Lock()
	value = 1
	Put(CreateTuple(13, value), rsp, targets1)
	log("p20 w:%d", value)
	countwrites(value, 20, 1)
	l.Unlock()

	delay()
	l.Lock()
	value = 17
	Put(CreateTuple(271, value), rsp, targets2)
	log("p20 w:%d", value)
	countwrites(value, 20, 17)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(52, &value), rsp, s20)
	log("p20 r:%d", value)
	countreads(value, 20, 4)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(335, &value), rsp, s20)
	log("p20 r:%d", value)
	countreads(value, 20, 21)
	l.Unlock()

	delay()
	l.Lock()
	value = 17
	Put(CreateTuple(264, value), rsp, targets3)
	log("p20 w:%d", value)
	countwrites(value, 20, 17)
	l.Unlock()

	delay()
	l.Lock()
	value = 3
	Put(CreateTuple(39, value), rsp, targets4)
	log("p20 w:%d", value)
	countwrites(value, 20, 3)
	l.Unlock()

	delay()
	l.Lock()
	value = 19
	Put(CreateTuple(298, value), rsp, targets5)
	log("p20 w:%d", value)
	countwrites(value, 20, 19)
	l.Unlock()

	delay()
	l.Lock()
	value = 1
	Put(CreateTuple(7, value), rsp, targets6)
	log("p20 w:%d", value)
	countwrites(value, 20, 1)
	l.Unlock()

	delay()
	l.Lock()
	value = 8
	Put(CreateTuple(125, value), rsp, targets7)
	log("p20 w:%d", value)
	countwrites(value, 20, 8)
	l.Unlock()

	delay()
	l.Lock()
	value = 17
	Put(CreateTuple(268, value), rsp, targets8)
	log("p20 w:%d", value)
	countwrites(value, 20, 17)
	l.Unlock()

	delay()
	l.Lock()
	value = 23
	Put(CreateTuple(364, value), rsp, targets9)
	log("p20 w:%d", value)
	countwrites(value, 20, 23)
	l.Unlock()

	delay()
	l.Lock()
	value = 29
	Put(CreateTuple(459, value), rsp, targets10)
	log("p20 w:%d", value)
	countwrites(value, 20, 29)
	l.Unlock()

	delay()
	l.Lock()
	value = 29
	Put(CreateTuple(460, value), rsp, targets11)
	log("p20 w:%d", value)
	countwrites(value, 20, 29)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(311, &value), rsp, s20)
	log("p20 r:%d", value)
	countreads(value, 20, 20)
	l.Unlock()

	delay()
	l.Lock()
	value = 13
	Put(CreateTuple(205, value), rsp, targets12)
	log("p20 w:%d", value)
	countwrites(value, 20, 13)
	l.Unlock()
}

func p21() {
	targets11 := make([]string, 0)
	targets10 := make([]string, 1)
	targets10[0] = uri[s13]
	targets9 := make([]string, 1)
	targets9[0] = uri[s27]
	targets8 := make([]string, 0)
	targets7 := make([]string, 0)
	targets6 := make([]string, 2)
	targets6[1] = uri[s23]
	targets6[0] = uri[s17]
	targets5 := make([]string, 0)
	targets4 := make([]string, 0)
	targets3 := make([]string, 1)
	targets3[0] = uri[s24]
	targets2 := make([]string, 0)
	targets1 := make([]string, 0)
	targets0 := make([]string, 0)

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 12
	Put(CreateTuple(187, value), rsp, targets0)
	log("p21 w:%d", value)
	countwrites(value, 21, 12)
	l.Unlock()

	delay()
	l.Lock()
	value = 5
	Put(CreateTuple(77, value), rsp, targets1)
	log("p21 w:%d", value)
	countwrites(value, 21, 5)
	l.Unlock()

	delay()
	l.Lock()
	value = 18
	Put(CreateTuple(274, value), rsp, targets2)
	log("p21 w:%d", value)
	countwrites(value, 21, 18)
	l.Unlock()

	delay()
	l.Lock()
	value = 30
	Put(CreateTuple(468, value), rsp, targets3)
	log("p21 w:%d", value)
	countwrites(value, 21, 30)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(461, &value), rsp, s21)
	log("p21 r:%d", value)
	countreads(value, 21, 29)
	l.Unlock()

	delay()
	l.Lock()
	value = 20
	Put(CreateTuple(305, value), rsp, targets4)
	log("p21 w:%d", value)
	countwrites(value, 21, 20)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(147, &value), rsp, s21)
	log("p21 r:%d", value)
	countreads(value, 21, 10)
	l.Unlock()

	delay()
	l.Lock()
	value = 29
	Put(CreateTuple(452, value), rsp, targets5)
	log("p21 w:%d", value)
	countwrites(value, 21, 29)
	l.Unlock()

	delay()
	l.Lock()
	value = 7
	Put(CreateTuple(99, value), rsp, targets6)
	log("p21 w:%d", value)
	countwrites(value, 21, 7)
	l.Unlock()

	delay()
	l.Lock()
	value = 28
	Put(CreateTuple(440, value), rsp, targets7)
	log("p21 w:%d", value)
	countwrites(value, 21, 28)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(460, &value), rsp, s21)
	log("p21 r:%d", value)
	countreads(value, 21, 29)
	l.Unlock()

	delay()
	l.Lock()
	value = 4
	Put(CreateTuple(63, value), rsp, targets8)
	log("p21 w:%d", value)
	countwrites(value, 21, 4)
	l.Unlock()

	delay()
	l.Lock()
	value = 3
	Put(CreateTuple(36, value), rsp, targets9)
	log("p21 w:%d", value)
	countwrites(value, 21, 3)
	l.Unlock()

	delay()
	l.Lock()
	value = 20
	Put(CreateTuple(314, value), rsp, targets10)
	log("p21 w:%d", value)
	countwrites(value, 21, 20)
	l.Unlock()

	delay()
	l.Lock()
	value = 13
	Put(CreateTuple(197, value), rsp, targets11)
	log("p21 w:%d", value)
	countwrites(value, 21, 13)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(266, &value), rsp, s21)
	log("p21 r:%d", value)
	countreads(value, 21, 17)
	l.Unlock()
}

func p22() {
	targets13 := make([]string, 0)
	targets12 := make([]string, 0)
	targets11 := make([]string, 0)
	targets10 := make([]string, 0)
	targets9 := make([]string, 0)
	targets8 := make([]string, 0)
	targets7 := make([]string, 0)
	targets6 := make([]string, 1)
	targets6[0] = uri[s25]
	targets5 := make([]string, 1)
	targets5[0] = uri[s1]
	targets4 := make([]string, 0)
	targets3 := make([]string, 0)
	targets2 := make([]string, 0)
	targets1 := make([]string, 0)
	targets0 := make([]string, 0)

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 21
	Put(CreateTuple(333, value), rsp, targets0)
	log("p22 w:%d", value)
	countwrites(value, 22, 21)
	l.Unlock()

	delay()
	l.Lock()
	value = 1
	Put(CreateTuple(8, value), rsp, targets1)
	log("p22 w:%d", value)
	countwrites(value, 22, 1)
	l.Unlock()

	delay()
	l.Lock()
	value = 9
	Put(CreateTuple(130, value), rsp, targets2)
	log("p22 w:%d", value)
	countwrites(value, 22, 9)
	l.Unlock()

	delay()
	l.Lock()
	value = 26
	Put(CreateTuple(402, value), rsp, targets3)
	log("p22 w:%d", value)
	countwrites(value, 22, 26)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(143, &value), rsp, s22)
	log("p22 r:%d", value)
	countreads(value, 22, 9)
	l.Unlock()

	delay()
	l.Lock()
	value = 24
	Put(CreateTuple(378, value), rsp, targets4)
	log("p22 w:%d", value)
	countwrites(value, 22, 24)
	l.Unlock()

	delay()
	l.Lock()
	value = 22
	Put(CreateTuple(344, value), rsp, targets5)
	log("p22 w:%d", value)
	countwrites(value, 22, 22)
	l.Unlock()

	delay()
	l.Lock()
	value = 29
	Put(CreateTuple(456, value), rsp, targets6)
	log("p22 w:%d", value)
	countwrites(value, 22, 29)
	l.Unlock()

	delay()
	l.Lock()
	value = 15
	Put(CreateTuple(228, value), rsp, targets7)
	log("p22 w:%d", value)
	countwrites(value, 22, 15)
	l.Unlock()

	delay()
	l.Lock()
	value = 22
	Put(CreateTuple(351, value), rsp, targets8)
	log("p22 w:%d", value)
	countwrites(value, 22, 22)
	l.Unlock()

	delay()
	l.Lock()
	value = 31
	Put(CreateTuple(486, value), rsp, targets9)
	log("p22 w:%d", value)
	countwrites(value, 22, 31)
	l.Unlock()

	delay()
	l.Lock()
	value = 6
	Put(CreateTuple(91, value), rsp, targets10)
	log("p22 w:%d", value)
	countwrites(value, 22, 6)
	l.Unlock()

	delay()
	l.Lock()
	value = 30
	Put(CreateTuple(476, value), rsp, targets11)
	log("p22 w:%d", value)
	countwrites(value, 22, 30)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(457, &value), rsp, s22)
	log("p22 r:%d", value)
	countreads(value, 22, 29)
	l.Unlock()

	delay()
	l.Lock()
	value = 27
	Put(CreateTuple(430, value), rsp, targets12)
	log("p22 w:%d", value)
	countwrites(value, 22, 27)
	l.Unlock()

	delay()
	l.Lock()
	value = 3
	Put(CreateTuple(35, value), rsp, targets13)
	log("p22 w:%d", value)
	countwrites(value, 22, 3)
	l.Unlock()
}

func p23() {
	targets9 := make([]string, 0)
	targets8 := make([]string, 0)
	targets7 := make([]string, 0)
	targets6 := make([]string, 0)
	targets5 := make([]string, 0)
	targets4 := make([]string, 0)
	targets3 := make([]string, 0)
	targets2 := make([]string, 0)
	targets1 := make([]string, 1)
	targets1[0] = uri[s20]
	targets0 := make([]string, 0)

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 8
	Put(CreateTuple(127, value), rsp, targets0)
	log("p23 w:%d", value)
	countwrites(value, 23, 8)
	l.Unlock()

	delay()
	l.Lock()
	value = 4
	Put(CreateTuple(52, value), rsp, targets1)
	log("p23 w:%d", value)
	countwrites(value, 23, 4)
	l.Unlock()

	delay()
	l.Lock()
	value = 32
	Put(CreateTuple(501, value), rsp, targets2)
	log("p23 w:%d", value)
	countwrites(value, 23, 32)
	l.Unlock()

	delay()
	l.Lock()
	value = 28
	Put(CreateTuple(441, value), rsp, targets3)
	log("p23 w:%d", value)
	countwrites(value, 23, 28)
	l.Unlock()

	delay()
	l.Lock()
	value = 7
	Put(CreateTuple(100, value), rsp, targets4)
	log("p23 w:%d", value)
	countwrites(value, 23, 7)
	l.Unlock()

	delay()
	l.Lock()
	value = 3
	Put(CreateTuple(42, value), rsp, targets5)
	log("p23 w:%d", value)
	countwrites(value, 23, 3)
	l.Unlock()

	delay()
	l.Lock()
	value = 11
	Put(CreateTuple(176, value), rsp, targets6)
	log("p23 w:%d", value)
	countwrites(value, 23, 11)
	l.Unlock()

	delay()
	l.Lock()
	value = 27
	Put(CreateTuple(428, value), rsp, targets7)
	log("p23 w:%d", value)
	countwrites(value, 23, 27)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(178, &value), rsp, s23)
	log("p23 r:%d", value)
	countreads(value, 23, 12)
	l.Unlock()

	delay()
	l.Lock()
	value = 9
	Put(CreateTuple(131, value), rsp, targets8)
	log("p23 w:%d", value)
	countwrites(value, 23, 9)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(99, &value), rsp, s23)
	log("p23 r:%d", value)
	countreads(value, 23, 7)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(90, &value), rsp, s23)
	log("p23 r:%d", value)
	countreads(value, 23, 6)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(371, &value), rsp, s23)
	log("p23 r:%d", value)
	countreads(value, 23, 24)
	l.Unlock()

	delay()
	l.Lock()
	value = 2
	Put(CreateTuple(19, value), rsp, targets9)
	log("p23 w:%d", value)
	countwrites(value, 23, 2)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(108, &value), rsp, s23)
	log("p23 r:%d", value)
	countreads(value, 23, 7)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(465, &value), rsp, s23)
	log("p23 r:%d", value)
	countreads(value, 23, 30)
	l.Unlock()
}

func p24() {
	targets12 := make([]string, 0)
	targets11 := make([]string, 0)
	targets10 := make([]string, 0)
	targets9 := make([]string, 0)
	targets8 := make([]string, 0)
	targets7 := make([]string, 0)
	targets6 := make([]string, 1)
	targets6[0] = uri[s22]
	targets5 := make([]string, 0)
	targets4 := make([]string, 0)
	targets3 := make([]string, 0)
	targets2 := make([]string, 0)
	targets1 := make([]string, 0)
	targets0 := make([]string, 1)
	targets0[0] = uri[s4]

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 13
	Put(CreateTuple(195, value), rsp, targets0)
	log("p24 w:%d", value)
	countwrites(value, 24, 13)
	l.Unlock()

	delay()
	l.Lock()
	value = 4
	Put(CreateTuple(61, value), rsp, targets1)
	log("p24 w:%d", value)
	countwrites(value, 24, 4)
	l.Unlock()

	delay()
	l.Lock()
	value = 8
	Put(CreateTuple(122, value), rsp, targets2)
	log("p24 w:%d", value)
	countwrites(value, 24, 8)
	l.Unlock()

	delay()
	l.Lock()
	value = 14
	Put(CreateTuple(218, value), rsp, targets3)
	log("p24 w:%d", value)
	countwrites(value, 24, 14)
	l.Unlock()

	delay()
	l.Lock()
	value = 1
	Put(CreateTuple(5, value), rsp, targets4)
	log("p24 w:%d", value)
	countwrites(value, 24, 1)
	l.Unlock()

	delay()
	l.Lock()
	value = 19
	Put(CreateTuple(298, value), rsp, targets5)
	log("p24 w:%d", value)
	countwrites(value, 24, 19)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(468, &value), rsp, s24)
	log("p24 r:%d", value)
	countreads(value, 24, 30)
	l.Unlock()

	delay()
	l.Lock()
	value = 9
	Put(CreateTuple(143, value), rsp, targets6)
	log("p24 w:%d", value)
	countwrites(value, 24, 9)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(438, &value), rsp, s24)
	log("p24 r:%d", value)
	countreads(value, 24, 28)
	l.Unlock()

	delay()
	l.Lock()
	value = 22
	Put(CreateTuple(345, value), rsp, targets7)
	log("p24 w:%d", value)
	countwrites(value, 24, 22)
	l.Unlock()

	delay()
	l.Lock()
	value = 23
	Put(CreateTuple(360, value), rsp, targets8)
	log("p24 w:%d", value)
	countwrites(value, 24, 23)
	l.Unlock()

	delay()
	l.Lock()
	value = 27
	Put(CreateTuple(430, value), rsp, targets9)
	log("p24 w:%d", value)
	countwrites(value, 24, 27)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(45, &value), rsp, s24)
	log("p24 r:%d", value)
	countreads(value, 24, 3)
	l.Unlock()

	delay()
	l.Lock()
	value = 4
	Put(CreateTuple(61, value), rsp, targets10)
	log("p24 w:%d", value)
	countwrites(value, 24, 4)
	l.Unlock()

	delay()
	l.Lock()
	value = 15
	Put(CreateTuple(225, value), rsp, targets11)
	log("p24 w:%d", value)
	countwrites(value, 24, 15)
	l.Unlock()

	delay()
	l.Lock()
	value = 1
	Put(CreateTuple(3, value), rsp, targets12)
	log("p24 w:%d", value)
	countwrites(value, 24, 1)
	l.Unlock()
}

func p25() {
	targets13 := make([]string, 1)
	targets13[0] = uri[s18]
	targets12 := make([]string, 0)
	targets11 := make([]string, 0)
	targets10 := make([]string, 0)
	targets9 := make([]string, 0)
	targets8 := make([]string, 0)
	targets7 := make([]string, 0)
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
	value = 30
	Put(CreateTuple(472, value), rsp, targets0)
	log("p25 w:%d", value)
	countwrites(value, 25, 30)
	l.Unlock()

	delay()
	l.Lock()
	value = 1
	Put(CreateTuple(10, value), rsp, targets1)
	log("p25 w:%d", value)
	countwrites(value, 25, 1)
	l.Unlock()

	delay()
	l.Lock()
	value = 11
	Put(CreateTuple(163, value), rsp, targets2)
	log("p25 w:%d", value)
	countwrites(value, 25, 11)
	l.Unlock()

	delay()
	l.Lock()
	value = 4
	Put(CreateTuple(63, value), rsp, targets3)
	log("p25 w:%d", value)
	countwrites(value, 25, 4)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(456, &value), rsp, s25)
	log("p25 r:%d", value)
	countreads(value, 25, 29)
	l.Unlock()

	delay()
	l.Lock()
	value = 25
	Put(CreateTuple(391, value), rsp, targets4)
	log("p25 w:%d", value)
	countwrites(value, 25, 25)
	l.Unlock()

	delay()
	l.Lock()
	value = 30
	Put(CreateTuple(470, value), rsp, targets5)
	log("p25 w:%d", value)
	countwrites(value, 25, 30)
	l.Unlock()

	delay()
	l.Lock()
	value = 19
	Put(CreateTuple(298, value), rsp, targets6)
	log("p25 w:%d", value)
	countwrites(value, 25, 19)
	l.Unlock()

	delay()
	l.Lock()
	value = 29
	Put(CreateTuple(462, value), rsp, targets7)
	log("p25 w:%d", value)
	countwrites(value, 25, 29)
	l.Unlock()

	delay()
	l.Lock()
	value = 4
	Put(CreateTuple(59, value), rsp, targets8)
	log("p25 w:%d", value)
	countwrites(value, 25, 4)
	l.Unlock()

	delay()
	l.Lock()
	value = 3
	Put(CreateTuple(46, value), rsp, targets9)
	log("p25 w:%d", value)
	countwrites(value, 25, 3)
	l.Unlock()

	delay()
	l.Lock()
	value = 10
	Put(CreateTuple(150, value), rsp, targets10)
	log("p25 w:%d", value)
	countwrites(value, 25, 10)
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
	value = 12
	Put(CreateTuple(188, value), rsp, targets11)
	log("p25 w:%d", value)
	countwrites(value, 25, 12)
	l.Unlock()

	delay()
	l.Lock()
	value = 7
	Put(CreateTuple(101, value), rsp, targets12)
	log("p25 w:%d", value)
	countwrites(value, 25, 7)
	l.Unlock()

	delay()
	l.Lock()
	value = 18
	Put(CreateTuple(283, value), rsp, targets13)
	log("p25 w:%d", value)
	countwrites(value, 25, 18)
	l.Unlock()
}

func p26() {
	targets14 := make([]string, 0)
	targets13 := make([]string, 0)
	targets12 := make([]string, 0)
	targets11 := make([]string, 1)
	targets11[0] = uri[s2]
	targets10 := make([]string, 0)
	targets9 := make([]string, 0)
	targets8 := make([]string, 0)
	targets7 := make([]string, 0)
	targets6 := make([]string, 0)
	targets5 := make([]string, 0)
	targets4 := make([]string, 0)
	targets3 := make([]string, 1)
	targets3[0] = uri[s29]
	targets2 := make([]string, 0)
	targets1 := make([]string, 0)
	targets0 := make([]string, 0)

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 6
	Put(CreateTuple(88, value), rsp, targets0)
	log("p26 w:%d", value)
	countwrites(value, 26, 6)
	l.Unlock()

	delay()
	l.Lock()
	value = 19
	Put(CreateTuple(302, value), rsp, targets1)
	log("p26 w:%d", value)
	countwrites(value, 26, 19)
	l.Unlock()

	delay()
	l.Lock()
	value = 7
	Put(CreateTuple(102, value), rsp, targets2)
	log("p26 w:%d", value)
	countwrites(value, 26, 7)
	l.Unlock()

	delay()
	l.Lock()
	value = 31
	Put(CreateTuple(490, value), rsp, targets3)
	log("p26 w:%d", value)
	countwrites(value, 26, 31)
	l.Unlock()

	delay()
	l.Lock()
	value = 6
	Put(CreateTuple(96, value), rsp, targets4)
	log("p26 w:%d", value)
	countwrites(value, 26, 6)
	l.Unlock()

	delay()
	l.Lock()
	value = 1
	Put(CreateTuple(15, value), rsp, targets5)
	log("p26 w:%d", value)
	countwrites(value, 26, 1)
	l.Unlock()

	delay()
	l.Lock()
	value = 25
	Put(CreateTuple(388, value), rsp, targets6)
	log("p26 w:%d", value)
	countwrites(value, 26, 25)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(418, &value), rsp, s26)
	log("p26 r:%d", value)
	countreads(value, 26, 27)
	l.Unlock()

	delay()
	l.Lock()
	value = 15
	Put(CreateTuple(225, value), rsp, targets7)
	log("p26 w:%d", value)
	countwrites(value, 26, 15)
	l.Unlock()

	delay()
	l.Lock()
	value = 5
	Put(CreateTuple(76, value), rsp, targets8)
	log("p26 w:%d", value)
	countwrites(value, 26, 5)
	l.Unlock()

	delay()
	l.Lock()
	value = 23
	Put(CreateTuple(362, value), rsp, targets9)
	log("p26 w:%d", value)
	countwrites(value, 26, 23)
	l.Unlock()

	delay()
	l.Lock()
	value = 5
	Put(CreateTuple(69, value), rsp, targets10)
	log("p26 w:%d", value)
	countwrites(value, 26, 5)
	l.Unlock()

	delay()
	l.Lock()
	value = 31
	Put(CreateTuple(494, value), rsp, targets11)
	log("p26 w:%d", value)
	countwrites(value, 26, 31)
	l.Unlock()

	delay()
	l.Lock()
	value = 30
	Put(CreateTuple(471, value), rsp, targets12)
	log("p26 w:%d", value)
	countwrites(value, 26, 30)
	l.Unlock()

	delay()
	l.Lock()
	value = 17
	Put(CreateTuple(267, value), rsp, targets13)
	log("p26 w:%d", value)
	countwrites(value, 26, 17)
	l.Unlock()

	delay()
	l.Lock()
	value = 10
	Put(CreateTuple(148, value), rsp, targets14)
	log("p26 w:%d", value)
	countwrites(value, 26, 10)
	l.Unlock()
}

func p27() {
	targets10 := make([]string, 0)
	targets9 := make([]string, 0)
	targets8 := make([]string, 0)
	targets7 := make([]string, 0)
	targets6 := make([]string, 0)
	targets5 := make([]string, 1)
	targets5[0] = uri[s20]
	targets4 := make([]string, 0)
	targets3 := make([]string, 1)
	targets3[0] = uri[s31]
	targets2 := make([]string, 2)
	targets2[1] = uri[s23]
	targets2[0] = uri[s17]
	targets1 := make([]string, 0)
	targets0 := make([]string, 0)

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 4
	Put(CreateTuple(59, value), rsp, targets0)
	log("p27 w:%d", value)
	countwrites(value, 27, 4)
	l.Unlock()

	delay()
	l.Lock()
	value = 12
	Put(CreateTuple(177, value), rsp, targets1)
	log("p27 w:%d", value)
	countwrites(value, 27, 12)
	l.Unlock()

	delay()
	l.Lock()
	value = 7
	Put(CreateTuple(99, value), rsp, targets2)
	log("p27 w:%d", value)
	countwrites(value, 27, 7)
	l.Unlock()

	delay()
	l.Lock()
	value = 1
	Put(CreateTuple(12, value), rsp, targets3)
	log("p27 w:%d", value)
	countwrites(value, 27, 1)
	l.Unlock()

	delay()
	l.Lock()
	value = 17
	Put(CreateTuple(257, value), rsp, targets4)
	log("p27 w:%d", value)
	countwrites(value, 27, 17)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(287, &value), rsp, s27)
	log("p27 r:%d", value)
	countreads(value, 27, 18)
	l.Unlock()

	delay()
	l.Lock()
	value = 4
	Put(CreateTuple(52, value), rsp, targets5)
	log("p27 w:%d", value)
	countwrites(value, 27, 4)
	l.Unlock()

	delay()
	l.Lock()
	value = 31
	Put(CreateTuple(493, value), rsp, targets6)
	log("p27 w:%d", value)
	countwrites(value, 27, 31)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(13, &value), rsp, s27)
	log("p27 r:%d", value)
	countreads(value, 27, 1)
	l.Unlock()

	delay()
	l.Lock()
	value = 27
	Put(CreateTuple(429, value), rsp, targets7)
	log("p27 w:%d", value)
	countwrites(value, 27, 27)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(45, &value), rsp, s27)
	log("p27 r:%d", value)
	countreads(value, 27, 3)
	l.Unlock()

	delay()
	l.Lock()
	value = 20
	Put(CreateTuple(309, value), rsp, targets8)
	log("p27 w:%d", value)
	countwrites(value, 27, 20)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(36, &value), rsp, s27)
	log("p27 r:%d", value)
	countreads(value, 27, 3)
	l.Unlock()

	delay()
	l.Lock()
	value = 28
	Put(CreateTuple(437, value), rsp, targets9)
	log("p27 w:%d", value)
	countwrites(value, 27, 28)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(434, &value), rsp, s27)
	log("p27 r:%d", value)
	countreads(value, 27, 28)
	l.Unlock()

	delay()
	l.Lock()
	value = 17
	Put(CreateTuple(261, value), rsp, targets10)
	log("p27 w:%d", value)
	countwrites(value, 27, 17)
	l.Unlock()
}

func p28() {
	targets13 := make([]string, 0)
	targets12 := make([]string, 0)
	targets11 := make([]string, 0)
	targets10 := make([]string, 0)
	targets9 := make([]string, 0)
	targets8 := make([]string, 0)
	targets7 := make([]string, 1)
	targets7[0] = uri[s27]
	targets6 := make([]string, 0)
	targets5 := make([]string, 0)
	targets4 := make([]string, 1)
	targets4[0] = uri[s20]
	targets3 := make([]string, 0)
	targets2 := make([]string, 1)
	targets2[0] = uri[s6]
	targets1 := make([]string, 0)
	targets0 := make([]string, 0)

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(255, &value), rsp, s28)
	log("p28 r:%d", value)
	countreads(value, 28, 16)
	l.Unlock()

	delay()
	l.Lock()
	value = 23
	Put(CreateTuple(364, value), rsp, targets0)
	log("p28 w:%d", value)
	countwrites(value, 28, 23)
	l.Unlock()

	delay()
	l.Lock()
	value = 30
	Put(CreateTuple(474, value), rsp, targets1)
	log("p28 w:%d", value)
	countwrites(value, 28, 30)
	l.Unlock()

	delay()
	l.Lock()
	value = 17
	Put(CreateTuple(268, value), rsp, targets2)
	log("p28 w:%d", value)
	countwrites(value, 28, 17)
	l.Unlock()

	delay()
	l.Lock()
	value = 27
	Put(CreateTuple(431, value), rsp, targets3)
	log("p28 w:%d", value)
	countwrites(value, 28, 27)
	l.Unlock()

	delay()
	l.Lock()
	value = 4
	Put(CreateTuple(52, value), rsp, targets4)
	log("p28 w:%d", value)
	countwrites(value, 28, 4)
	l.Unlock()

	delay()
	l.Lock()
	value = 16
	Put(CreateTuple(251, value), rsp, targets5)
	log("p28 w:%d", value)
	countwrites(value, 28, 16)
	l.Unlock()

	delay()
	l.Lock()
	value = 27
	Put(CreateTuple(421, value), rsp, targets6)
	log("p28 w:%d", value)
	countwrites(value, 28, 27)
	l.Unlock()

	delay()
	l.Lock()
	value = 28
	Put(CreateTuple(434, value), rsp, targets7)
	log("p28 w:%d", value)
	countwrites(value, 28, 28)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(497, &value), rsp, s28)
	log("p28 r:%d", value)
	countreads(value, 28, 32)
	l.Unlock()

	delay()
	l.Lock()
	value = 25
	Put(CreateTuple(387, value), rsp, targets8)
	log("p28 w:%d", value)
	countwrites(value, 28, 25)
	l.Unlock()

	delay()
	l.Lock()
	value = 4
	Put(CreateTuple(55, value), rsp, targets9)
	log("p28 w:%d", value)
	countwrites(value, 28, 4)
	l.Unlock()

	delay()
	l.Lock()
	value = 32
	Put(CreateTuple(499, value), rsp, targets10)
	log("p28 w:%d", value)
	countwrites(value, 28, 32)
	l.Unlock()

	delay()
	l.Lock()
	value = 25
	Put(CreateTuple(387, value), rsp, targets11)
	log("p28 w:%d", value)
	countwrites(value, 28, 25)
	l.Unlock()

	delay()
	l.Lock()
	value = 19
	Put(CreateTuple(295, value), rsp, targets12)
	log("p28 w:%d", value)
	countwrites(value, 28, 19)
	l.Unlock()

	delay()
	l.Lock()
	value = 10
	Put(CreateTuple(151, value), rsp, targets13)
	log("p28 w:%d", value)
	countwrites(value, 28, 10)
	l.Unlock()
}

func p29() {
	targets14 := make([]string, 0)
	targets13 := make([]string, 0)
	targets12 := make([]string, 0)
	targets11 := make([]string, 0)
	targets10 := make([]string, 0)
	targets9 := make([]string, 0)
	targets8 := make([]string, 1)
	targets8[0] = uri[s32]
	targets7 := make([]string, 0)
	targets6 := make([]string, 0)
	targets5 := make([]string, 0)
	targets4 := make([]string, 0)
	targets3 := make([]string, 0)
	targets2 := make([]string, 0)
	targets1 := make([]string, 0)
	targets0 := make([]string, 1)
	targets0[0] = uri[s11]

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 8
	Put(CreateTuple(113, value), rsp, targets0)
	log("p29 w:%d", value)
	countwrites(value, 29, 8)
	l.Unlock()

	delay()
	l.Lock()
	value = 9
	Put(CreateTuple(133, value), rsp, targets1)
	log("p29 w:%d", value)
	countwrites(value, 29, 9)
	l.Unlock()

	delay()
	l.Lock()
	value = 13
	Put(CreateTuple(201, value), rsp, targets2)
	log("p29 w:%d", value)
	countwrites(value, 29, 13)
	l.Unlock()

	delay()
	l.Lock()
	value = 19
	Put(CreateTuple(302, value), rsp, targets3)
	log("p29 w:%d", value)
	countwrites(value, 29, 19)
	l.Unlock()

	delay()
	l.Lock()
	value = 10
	Put(CreateTuple(148, value), rsp, targets4)
	log("p29 w:%d", value)
	countwrites(value, 29, 10)
	l.Unlock()

	delay()
	l.Lock()
	value = 27
	Put(CreateTuple(423, value), rsp, targets5)
	log("p29 w:%d", value)
	countwrites(value, 29, 27)
	l.Unlock()

	delay()
	l.Lock()
	value = 20
	Put(CreateTuple(305, value), rsp, targets6)
	log("p29 w:%d", value)
	countwrites(value, 29, 20)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(490, &value), rsp, s29)
	log("p29 r:%d", value)
	countreads(value, 29, 31)
	l.Unlock()

	delay()
	l.Lock()
	value = 28
	Put(CreateTuple(443, value), rsp, targets7)
	log("p29 w:%d", value)
	countwrites(value, 29, 28)
	l.Unlock()

	delay()
	l.Lock()
	value = 8
	Put(CreateTuple(118, value), rsp, targets8)
	log("p29 w:%d", value)
	countwrites(value, 29, 8)
	l.Unlock()

	delay()
	l.Lock()
	value = 5
	Put(CreateTuple(79, value), rsp, targets9)
	log("p29 w:%d", value)
	countwrites(value, 29, 5)
	l.Unlock()

	delay()
	l.Lock()
	value = 5
	Put(CreateTuple(78, value), rsp, targets10)
	log("p29 w:%d", value)
	countwrites(value, 29, 5)
	l.Unlock()

	delay()
	l.Lock()
	value = 18
	Put(CreateTuple(286, value), rsp, targets11)
	log("p29 w:%d", value)
	countwrites(value, 29, 18)
	l.Unlock()

	delay()
	l.Lock()
	value = 21
	Put(CreateTuple(331, value), rsp, targets12)
	log("p29 w:%d", value)
	countwrites(value, 29, 21)
	l.Unlock()

	delay()
	l.Lock()
	value = 15
	Put(CreateTuple(233, value), rsp, targets13)
	log("p29 w:%d", value)
	countwrites(value, 29, 15)
	l.Unlock()

	delay()
	l.Lock()
	value = 26
	Put(CreateTuple(412, value), rsp, targets14)
	log("p29 w:%d", value)
	countwrites(value, 29, 26)
	l.Unlock()
}

func p30() {
	targets12 := make([]string, 2)
	targets12[1] = uri[s15]
	targets12[0] = uri[s6]
	targets11 := make([]string, 0)
	targets10 := make([]string, 1)
	targets10[0] = uri[s4]
	targets9 := make([]string, 0)
	targets8 := make([]string, 0)
	targets7 := make([]string, 0)
	targets6 := make([]string, 1)
	targets6[0] = uri[s20]
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
	value = 27
	Put(CreateTuple(428, value), rsp, targets0)
	log("p30 w:%d", value)
	countwrites(value, 30, 27)
	l.Unlock()

	delay()
	l.Lock()
	value = 25
	Put(CreateTuple(399, value), rsp, targets1)
	log("p30 w:%d", value)
	countwrites(value, 30, 25)
	l.Unlock()

	delay()
	l.Lock()
	value = 19
	Put(CreateTuple(294, value), rsp, targets2)
	log("p30 w:%d", value)
	countwrites(value, 30, 19)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(89, &value), rsp, s30)
	log("p30 r:%d", value)
	countreads(value, 30, 6)
	l.Unlock()

	delay()
	l.Lock()
	value = 8
	Put(CreateTuple(127, value), rsp, targets3)
	log("p30 w:%d", value)
	countwrites(value, 30, 8)
	l.Unlock()

	delay()
	l.Lock()
	value = 20
	Put(CreateTuple(316, value), rsp, targets4)
	log("p30 w:%d", value)
	countwrites(value, 30, 20)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(206, &value), rsp, s30)
	log("p30 r:%d", value)
	countreads(value, 30, 13)
	l.Unlock()

	delay()
	l.Lock()
	value = 23
	Put(CreateTuple(357, value), rsp, targets5)
	log("p30 w:%d", value)
	countwrites(value, 30, 23)
	l.Unlock()

	delay()
	l.Lock()
	value = 21
	Put(CreateTuple(335, value), rsp, targets6)
	log("p30 w:%d", value)
	countwrites(value, 30, 21)
	l.Unlock()

	delay()
	l.Lock()
	value = 27
	Put(CreateTuple(424, value), rsp, targets7)
	log("p30 w:%d", value)
	countwrites(value, 30, 27)
	l.Unlock()

	delay()
	l.Lock()
	value = 22
	Put(CreateTuple(341, value), rsp, targets8)
	log("p30 w:%d", value)
	countwrites(value, 30, 22)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(28, &value), rsp, s30)
	log("p30 r:%d", value)
	countreads(value, 30, 2)
	l.Unlock()

	delay()
	l.Lock()
	value = 31
	Put(CreateTuple(481, value), rsp, targets9)
	log("p30 w:%d", value)
	countwrites(value, 30, 31)
	l.Unlock()

	delay()
	l.Lock()
	value = 13
	Put(CreateTuple(200, value), rsp, targets10)
	log("p30 w:%d", value)
	countwrites(value, 30, 13)
	l.Unlock()

	delay()
	l.Lock()
	value = 2
	Put(CreateTuple(18, value), rsp, targets11)
	log("p30 w:%d", value)
	countwrites(value, 30, 2)
	l.Unlock()

	delay()
	l.Lock()
	value = 28
	Put(CreateTuple(448, value), rsp, targets12)
	log("p30 w:%d", value)
	countwrites(value, 30, 28)
	l.Unlock()
}

func p31() {
	targets14 := make([]string, 0)
	targets13 := make([]string, 0)
	targets12 := make([]string, 0)
	targets11 := make([]string, 1)
	targets11[0] = uri[s4]
	targets10 := make([]string, 0)
	targets9 := make([]string, 0)
	targets8 := make([]string, 0)
	targets7 := make([]string, 0)
	targets6 := make([]string, 0)
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
	value = 14
	Put(CreateTuple(224, value), rsp, targets0)
	log("p31 w:%d", value)
	countwrites(value, 31, 14)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(12, &value), rsp, s31)
	log("p31 r:%d", value)
	countreads(value, 31, 1)
	l.Unlock()

	delay()
	l.Lock()
	value = 29
	Put(CreateTuple(462, value), rsp, targets1)
	log("p31 w:%d", value)
	countwrites(value, 31, 29)
	l.Unlock()

	delay()
	l.Lock()
	value = 22
	Put(CreateTuple(342, value), rsp, targets2)
	log("p31 w:%d", value)
	countwrites(value, 31, 22)
	l.Unlock()

	delay()
	l.Lock()
	value = 16
	Put(CreateTuple(244, value), rsp, targets3)
	log("p31 w:%d", value)
	countwrites(value, 31, 16)
	l.Unlock()

	delay()
	l.Lock()
	value = 11
	Put(CreateTuple(176, value), rsp, targets4)
	log("p31 w:%d", value)
	countwrites(value, 31, 11)
	l.Unlock()

	delay()
	l.Lock()
	value = 8
	Put(CreateTuple(125, value), rsp, targets5)
	log("p31 w:%d", value)
	countwrites(value, 31, 8)
	l.Unlock()

	delay()
	l.Lock()
	value = 21
	Put(CreateTuple(330, value), rsp, targets6)
	log("p31 w:%d", value)
	countwrites(value, 31, 21)
	l.Unlock()

	delay()
	l.Lock()
	value = 28
	Put(CreateTuple(440, value), rsp, targets7)
	log("p31 w:%d", value)
	countwrites(value, 31, 28)
	l.Unlock()

	delay()
	l.Lock()
	value = 4
	Put(CreateTuple(54, value), rsp, targets8)
	log("p31 w:%d", value)
	countwrites(value, 31, 4)
	l.Unlock()

	delay()
	l.Lock()
	value = 26
	Put(CreateTuple(405, value), rsp, targets9)
	log("p31 w:%d", value)
	countwrites(value, 31, 26)
	l.Unlock()

	delay()
	l.Lock()
	value = 23
	Put(CreateTuple(362, value), rsp, targets10)
	log("p31 w:%d", value)
	countwrites(value, 31, 23)
	l.Unlock()

	delay()
	l.Lock()
	value = 13
	Put(CreateTuple(195, value), rsp, targets11)
	log("p31 w:%d", value)
	countwrites(value, 31, 13)
	l.Unlock()

	delay()
	l.Lock()
	value = 21
	Put(CreateTuple(326, value), rsp, targets12)
	log("p31 w:%d", value)
	countwrites(value, 31, 21)
	l.Unlock()

	delay()
	l.Lock()
	value = 8
	Put(CreateTuple(119, value), rsp, targets13)
	log("p31 w:%d", value)
	countwrites(value, 31, 8)
	l.Unlock()

	delay()
	l.Lock()
	value = 23
	Put(CreateTuple(361, value), rsp, targets14)
	log("p31 w:%d", value)
	countwrites(value, 31, 23)
	l.Unlock()
}

func p32() {
	targets11 := make([]string, 0)
	targets10 := make([]string, 0)
	targets9 := make([]string, 0)
	targets8 := make([]string, 0)
	targets7 := make([]string, 0)
	targets6 := make([]string, 0)
	targets5 := make([]string, 2)
	targets5[1] = uri[s32]
	targets5[0] = uri[s11]
	targets4 := make([]string, 0)
	targets3 := make([]string, 1)
	targets3[0] = uri[s32]
	targets2 := make([]string, 0)
	targets1 := make([]string, 0)
	targets0 := make([]string, 1)
	targets0[0] = uri[s3]

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 25
	Put(CreateTuple(395, value), rsp, targets0)
	log("p32 w:%d", value)
	countwrites(value, 32, 25)
	l.Unlock()

	delay()
	l.Lock()
	value = 7
	Put(CreateTuple(97, value), rsp, targets1)
	log("p32 w:%d", value)
	countwrites(value, 32, 7)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(169, &value), rsp, s32)
	log("p32 r:%d", value)
	countreads(value, 32, 11)
	l.Unlock()

	delay()
	l.Lock()
	value = 6
	Put(CreateTuple(87, value), rsp, targets2)
	log("p32 w:%d", value)
	countwrites(value, 32, 6)
	l.Unlock()

	delay()
	l.Lock()
	value = 11
	Put(CreateTuple(169, value), rsp, targets3)
	log("p32 w:%d", value)
	countwrites(value, 32, 11)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(482, &value), rsp, s32)
	log("p32 r:%d", value)
	countreads(value, 32, 31)
	l.Unlock()

	delay()
	l.Lock()
	value = 32
	Put(CreateTuple(512, value), rsp, targets4)
	log("p32 w:%d", value)
	countwrites(value, 32, 32)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(118, &value), rsp, s32)
	log("p32 r:%d", value)
	countreads(value, 32, 8)
	l.Unlock()

	delay()
	l.Lock()
	value = 31
	Put(CreateTuple(482, value), rsp, targets5)
	log("p32 w:%d", value)
	countwrites(value, 32, 31)
	l.Unlock()

	delay()
	l.Lock()
	value = 28
	Put(CreateTuple(443, value), rsp, targets6)
	log("p32 w:%d", value)
	countwrites(value, 32, 28)
	l.Unlock()

	delay()
	l.Lock()
	value = 24
	Put(CreateTuple(377, value), rsp, targets7)
	log("p32 w:%d", value)
	countwrites(value, 32, 24)
	l.Unlock()

	delay()
	l.Lock()
	value = 9
	Put(CreateTuple(144, value), rsp, targets8)
	log("p32 w:%d", value)
	countwrites(value, 32, 9)
	l.Unlock()

	delay()
	l.Lock()
	value = 23
	Put(CreateTuple(368, value), rsp, targets9)
	log("p32 w:%d", value)
	countwrites(value, 32, 23)
	l.Unlock()

	delay()
	l.Lock()
	value = 22
	Put(CreateTuple(343, value), rsp, targets10)
	log("p32 w:%d", value)
	countwrites(value, 32, 22)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(103, &value), rsp, s32)
	log("p32 r:%d", value)
	countreads(value, 32, 7)
	l.Unlock()

	delay()
	l.Lock()
	value = 20
	Put(CreateTuple(308, value), rsp, targets11)
	log("p32 w:%d", value)
	countwrites(value, 32, 20)
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
