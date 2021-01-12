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
	targets8 := make([]string, 0)
	targets7 := make([]string, 0)
	targets6 := make([]string, 1)
	targets6[0] = uri[s6]
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
	value = 21
	Put(CreateTuple(330, value), rsp, targets0)
	log("p1 w:%d", value)
	countwrites(value, 1, 21)
	l.Unlock()

	delay()
	l.Lock()
	value = 7
	Put(CreateTuple(99, value), rsp, targets1)
	log("p1 w:%d", value)
	countwrites(value, 1, 7)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(233, &value), rsp, s1)
	log("p1 r:%d", value)
	countreads(value, 1, 15)
	l.Unlock()

	delay()
	l.Lock()
	value = 17
	Put(CreateTuple(264, value), rsp, targets2)
	log("p1 w:%d", value)
	countwrites(value, 1, 17)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(300, &value), rsp, s1)
	log("p1 r:%d", value)
	countreads(value, 1, 19)
	l.Unlock()

	delay()
	l.Lock()
	value = 10
	Put(CreateTuple(159, value), rsp, targets3)
	log("p1 w:%d", value)
	countwrites(value, 1, 10)
	l.Unlock()

	delay()
	l.Lock()
	value = 23
	Put(CreateTuple(356, value), rsp, targets4)
	log("p1 w:%d", value)
	countwrites(value, 1, 23)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(267, &value), rsp, s1)
	log("p1 r:%d", value)
	countreads(value, 1, 17)
	l.Unlock()

	delay()
	l.Lock()
	value = 25
	Put(CreateTuple(397, value), rsp, targets5)
	log("p1 w:%d", value)
	countwrites(value, 1, 25)
	l.Unlock()

	delay()
	l.Lock()
	value = 14
	Put(CreateTuple(220, value), rsp, targets6)
	log("p1 w:%d", value)
	countwrites(value, 1, 14)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(407, &value), rsp, s1)
	log("p1 r:%d", value)
	countreads(value, 1, 26)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(504, &value), rsp, s1)
	log("p1 r:%d", value)
	countreads(value, 1, 32)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(167, &value), rsp, s1)
	log("p1 r:%d", value)
	countreads(value, 1, 11)
	l.Unlock()

	delay()
	l.Lock()
	value = 21
	Put(CreateTuple(334, value), rsp, targets7)
	log("p1 w:%d", value)
	countwrites(value, 1, 21)
	l.Unlock()

	delay()
	l.Lock()
	value = 11
	Put(CreateTuple(164, value), rsp, targets8)
	log("p1 w:%d", value)
	countwrites(value, 1, 11)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(390, &value), rsp, s1)
	log("p1 r:%d", value)
	countreads(value, 1, 25)
	l.Unlock()
}

func p2() {
	targets9 := make([]string, 0)
	targets8 := make([]string, 2)
	targets8[1] = uri[s12]
	targets8[0] = uri[s4]
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
	value = 3
	Put(CreateTuple(44, value), rsp, targets0)
	log("p2 w:%d", value)
	countwrites(value, 2, 3)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(269, &value), rsp, s2)
	log("p2 r:%d", value)
	countreads(value, 2, 17)
	l.Unlock()

	delay()
	l.Lock()
	value = 9
	Put(CreateTuple(131, value), rsp, targets1)
	log("p2 w:%d", value)
	countwrites(value, 2, 9)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(405, &value), rsp, s2)
	log("p2 r:%d", value)
	countreads(value, 2, 26)
	l.Unlock()

	delay()
	l.Lock()
	value = 31
	Put(CreateTuple(488, value), rsp, targets2)
	log("p2 w:%d", value)
	countwrites(value, 2, 31)
	l.Unlock()

	delay()
	l.Lock()
	value = 8
	Put(CreateTuple(122, value), rsp, targets3)
	log("p2 w:%d", value)
	countwrites(value, 2, 8)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(378, &value), rsp, s2)
	log("p2 r:%d", value)
	countreads(value, 2, 24)
	l.Unlock()

	delay()
	l.Lock()
	value = 27
	Put(CreateTuple(424, value), rsp, targets4)
	log("p2 w:%d", value)
	countwrites(value, 2, 27)
	l.Unlock()

	delay()
	l.Lock()
	value = 25
	Put(CreateTuple(385, value), rsp, targets5)
	log("p2 w:%d", value)
	countwrites(value, 2, 25)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(215, &value), rsp, s2)
	log("p2 r:%d", value)
	countreads(value, 2, 14)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(462, &value), rsp, s2)
	log("p2 r:%d", value)
	countreads(value, 2, 29)
	l.Unlock()

	delay()
	l.Lock()
	value = 7
	Put(CreateTuple(110, value), rsp, targets6)
	log("p2 w:%d", value)
	countwrites(value, 2, 7)
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
	value = 6
	Put(CreateTuple(82, value), rsp, targets8)
	log("p2 w:%d", value)
	countwrites(value, 2, 6)
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
	value = 4
	Put(CreateTuple(55, value), rsp, targets9)
	log("p2 w:%d", value)
	countwrites(value, 2, 4)
	l.Unlock()
}

func p3() {
	targets11 := make([]string, 0)
	targets10 := make([]string, 1)
	targets10[0] = uri[s25]
	targets9 := make([]string, 1)
	targets9[0] = uri[s20]
	targets8 := make([]string, 2)
	targets8[1] = uri[s17]
	targets8[0] = uri[s13]
	targets7 := make([]string, 0)
	targets6 := make([]string, 1)
	targets6[0] = uri[s2]
	targets5 := make([]string, 1)
	targets5[0] = uri[s29]
	targets4 := make([]string, 2)
	targets4[1] = uri[s27]
	targets4[0] = uri[s9]
	targets3 := make([]string, 0)
	targets2 := make([]string, 1)
	targets2[0] = uri[s23]
	targets1 := make([]string, 1)
	targets1[0] = uri[s20]
	targets0 := make([]string, 0)

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(142, &value), rsp, s3)
	log("p3 r:%d", value)
	countreads(value, 3, 9)
	l.Unlock()

	delay()
	l.Lock()
	value = 2
	Put(CreateTuple(27, value), rsp, targets0)
	log("p3 w:%d", value)
	countwrites(value, 3, 2)
	l.Unlock()

	delay()
	l.Lock()
	value = 28
	Put(CreateTuple(435, value), rsp, targets1)
	log("p3 w:%d", value)
	countwrites(value, 3, 28)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(4, &value), rsp, s3)
	log("p3 r:%d", value)
	countreads(value, 3, 1)
	l.Unlock()

	delay()
	l.Lock()
	value = 22
	Put(CreateTuple(348, value), rsp, targets2)
	log("p3 w:%d", value)
	countwrites(value, 3, 22)
	l.Unlock()

	delay()
	l.Lock()
	value = 24
	Put(CreateTuple(373, value), rsp, targets3)
	log("p3 w:%d", value)
	countwrites(value, 3, 24)
	l.Unlock()

	delay()
	l.Lock()
	value = 12
	Put(CreateTuple(187, value), rsp, targets4)
	log("p3 w:%d", value)
	countwrites(value, 3, 12)
	l.Unlock()

	delay()
	l.Lock()
	value = 3
	Put(CreateTuple(36, value), rsp, targets5)
	log("p3 w:%d", value)
	countwrites(value, 3, 3)
	l.Unlock()

	delay()
	l.Lock()
	value = 26
	Put(CreateTuple(405, value), rsp, targets6)
	log("p3 w:%d", value)
	countwrites(value, 3, 26)
	l.Unlock()

	delay()
	l.Lock()
	value = 28
	Put(CreateTuple(444, value), rsp, targets7)
	log("p3 w:%d", value)
	countwrites(value, 3, 28)
	l.Unlock()

	delay()
	l.Lock()
	value = 6
	Put(CreateTuple(86, value), rsp, targets8)
	log("p3 w:%d", value)
	countwrites(value, 3, 6)
	l.Unlock()

	delay()
	l.Lock()
	value = 2
	Put(CreateTuple(28, value), rsp, targets9)
	log("p3 w:%d", value)
	countwrites(value, 3, 2)
	l.Unlock()

	delay()
	l.Lock()
	value = 8
	Put(CreateTuple(124, value), rsp, targets10)
	log("p3 w:%d", value)
	countwrites(value, 3, 8)
	l.Unlock()

	delay()
	l.Lock()
	value = 29
	Put(CreateTuple(456, value), rsp, targets11)
	log("p3 w:%d", value)
	countwrites(value, 3, 29)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(315, &value), rsp, s3)
	log("p3 r:%d", value)
	countreads(value, 3, 20)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(108, &value), rsp, s3)
	log("p3 r:%d", value)
	countreads(value, 3, 7)
	l.Unlock()
}

func p4() {
	targets11 := make([]string, 0)
	targets10 := make([]string, 0)
	targets9 := make([]string, 0)
	targets8 := make([]string, 1)
	targets8[0] = uri[s7]
	targets7 := make([]string, 0)
	targets6 := make([]string, 0)
	targets5 := make([]string, 0)
	targets4 := make([]string, 1)
	targets4[0] = uri[s7]
	targets3 := make([]string, 0)
	targets2 := make([]string, 0)
	targets1 := make([]string, 0)
	targets0 := make([]string, 3)
	targets0[2] = uri[s16]
	targets0[1] = uri[s12]
	targets0[0] = uri[s1]

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(22, &value), rsp, s4)
	log("p4 r:%d", value)
	countreads(value, 4, 2)
	l.Unlock()

	delay()
	l.Lock()
	value = 19
	Put(CreateTuple(300, value), rsp, targets0)
	log("p4 w:%d", value)
	countwrites(value, 4, 19)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(82, &value), rsp, s4)
	log("p4 r:%d", value)
	countreads(value, 4, 6)
	l.Unlock()

	delay()
	l.Lock()
	value = 4
	Put(CreateTuple(58, value), rsp, targets1)
	log("p4 w:%d", value)
	countwrites(value, 4, 4)
	l.Unlock()

	delay()
	l.Lock()
	value = 11
	Put(CreateTuple(168, value), rsp, targets2)
	log("p4 w:%d", value)
	countwrites(value, 4, 11)
	l.Unlock()

	delay()
	l.Lock()
	value = 27
	Put(CreateTuple(417, value), rsp, targets3)
	log("p4 w:%d", value)
	countwrites(value, 4, 27)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(330, &value), rsp, s4)
	log("p4 r:%d", value)
	countreads(value, 4, 21)
	l.Unlock()

	delay()
	l.Lock()
	value = 21
	Put(CreateTuple(331, value), rsp, targets4)
	log("p4 w:%d", value)
	countwrites(value, 4, 21)
	l.Unlock()

	delay()
	l.Lock()
	value = 28
	Put(CreateTuple(437, value), rsp, targets5)
	log("p4 w:%d", value)
	countwrites(value, 4, 28)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(68, &value), rsp, s4)
	log("p4 r:%d", value)
	countreads(value, 4, 5)
	l.Unlock()

	delay()
	l.Lock()
	value = 27
	Put(CreateTuple(423, value), rsp, targets6)
	log("p4 w:%d", value)
	countwrites(value, 4, 27)
	l.Unlock()

	delay()
	l.Lock()
	value = 22
	Put(CreateTuple(352, value), rsp, targets7)
	log("p4 w:%d", value)
	countwrites(value, 4, 22)
	l.Unlock()

	delay()
	l.Lock()
	value = 21
	Put(CreateTuple(331, value), rsp, targets8)
	log("p4 w:%d", value)
	countwrites(value, 4, 21)
	l.Unlock()

	delay()
	l.Lock()
	value = 13
	Put(CreateTuple(198, value), rsp, targets9)
	log("p4 w:%d", value)
	countwrites(value, 4, 13)
	l.Unlock()

	delay()
	l.Lock()
	value = 16
	Put(CreateTuple(245, value), rsp, targets10)
	log("p4 w:%d", value)
	countwrites(value, 4, 16)
	l.Unlock()

	delay()
	l.Lock()
	value = 32
	Put(CreateTuple(509, value), rsp, targets11)
	log("p4 w:%d", value)
	countwrites(value, 4, 32)
	l.Unlock()
}

func p5() {
	targets10 := make([]string, 1)
	targets10[0] = uri[s14]
	targets9 := make([]string, 0)
	targets8 := make([]string, 0)
	targets7 := make([]string, 0)
	targets6 := make([]string, 0)
	targets5 := make([]string, 1)
	targets5[0] = uri[s27]
	targets4 := make([]string, 0)
	targets3 := make([]string, 0)
	targets2 := make([]string, 0)
	targets1 := make([]string, 1)
	targets1[0] = uri[s30]
	targets0 := make([]string, 0)

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 29
	Put(CreateTuple(454, value), rsp, targets0)
	log("p5 w:%d", value)
	countwrites(value, 5, 29)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(60, &value), rsp, s5)
	log("p5 r:%d", value)
	countreads(value, 5, 4)
	l.Unlock()

	delay()
	l.Lock()
	value = 27
	Put(CreateTuple(420, value), rsp, targets1)
	log("p5 w:%d", value)
	countwrites(value, 5, 27)
	l.Unlock()

	delay()
	l.Lock()
	value = 32
	Put(CreateTuple(509, value), rsp, targets2)
	log("p5 w:%d", value)
	countwrites(value, 5, 32)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(355, &value), rsp, s5)
	log("p5 r:%d", value)
	countreads(value, 5, 23)
	l.Unlock()

	delay()
	l.Lock()
	value = 11
	Put(CreateTuple(172, value), rsp, targets3)
	log("p5 w:%d", value)
	countwrites(value, 5, 11)
	l.Unlock()

	delay()
	l.Lock()
	value = 16
	Put(CreateTuple(243, value), rsp, targets4)
	log("p5 w:%d", value)
	countwrites(value, 5, 16)
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
	QueryP(CreateTuple(491, &value), rsp, s5)
	log("p5 r:%d", value)
	countreads(value, 5, 31)
	l.Unlock()

	delay()
	l.Lock()
	value = 30
	Put(CreateTuple(476, value), rsp, targets5)
	log("p5 w:%d", value)
	countwrites(value, 5, 30)
	l.Unlock()

	delay()
	l.Lock()
	value = 18
	Put(CreateTuple(279, value), rsp, targets6)
	log("p5 w:%d", value)
	countwrites(value, 5, 18)
	l.Unlock()

	delay()
	l.Lock()
	value = 1
	Put(CreateTuple(15, value), rsp, targets7)
	log("p5 w:%d", value)
	countwrites(value, 5, 1)
	l.Unlock()

	delay()
	l.Lock()
	value = 29
	Put(CreateTuple(451, value), rsp, targets8)
	log("p5 w:%d", value)
	countwrites(value, 5, 29)
	l.Unlock()

	delay()
	l.Lock()
	value = 2
	Put(CreateTuple(19, value), rsp, targets9)
	log("p5 w:%d", value)
	countwrites(value, 5, 2)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(129, &value), rsp, s5)
	log("p5 r:%d", value)
	countreads(value, 5, 9)
	l.Unlock()

	delay()
	l.Lock()
	value = 4
	Put(CreateTuple(49, value), rsp, targets10)
	log("p5 w:%d", value)
	countwrites(value, 5, 4)
	l.Unlock()
}

func p6() {
	targets13 := make([]string, 0)
	targets12 := make([]string, 0)
	targets11 := make([]string, 0)
	targets10 := make([]string, 0)
	targets9 := make([]string, 0)
	targets8 := make([]string, 1)
	targets8[0] = uri[s27]
	targets7 := make([]string, 0)
	targets6 := make([]string, 0)
	targets5 := make([]string, 1)
	targets5[0] = uri[s30]
	targets4 := make([]string, 0)
	targets3 := make([]string, 1)
	targets3[0] = uri[s18]
	targets2 := make([]string, 0)
	targets1 := make([]string, 0)
	targets0 := make([]string, 0)

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 17
	Put(CreateTuple(259, value), rsp, targets0)
	log("p6 w:%d", value)
	countwrites(value, 6, 17)
	l.Unlock()

	delay()
	l.Lock()
	value = 9
	Put(CreateTuple(131, value), rsp, targets1)
	log("p6 w:%d", value)
	countwrites(value, 6, 9)
	l.Unlock()

	delay()
	l.Lock()
	value = 22
	Put(CreateTuple(338, value), rsp, targets2)
	log("p6 w:%d", value)
	countwrites(value, 6, 22)
	l.Unlock()

	delay()
	l.Lock()
	value = 17
	Put(CreateTuple(271, value), rsp, targets3)
	log("p6 w:%d", value)
	countwrites(value, 6, 17)
	l.Unlock()

	delay()
	l.Lock()
	value = 19
	Put(CreateTuple(298, value), rsp, targets4)
	log("p6 w:%d", value)
	countwrites(value, 6, 19)
	l.Unlock()

	delay()
	l.Lock()
	value = 25
	Put(CreateTuple(398, value), rsp, targets5)
	log("p6 w:%d", value)
	countwrites(value, 6, 25)
	l.Unlock()

	delay()
	l.Lock()
	value = 30
	Put(CreateTuple(475, value), rsp, targets6)
	log("p6 w:%d", value)
	countwrites(value, 6, 30)
	l.Unlock()

	delay()
	l.Lock()
	value = 8
	Put(CreateTuple(127, value), rsp, targets7)
	log("p6 w:%d", value)
	countwrites(value, 6, 8)
	l.Unlock()

	delay()
	l.Lock()
	value = 31
	Put(CreateTuple(484, value), rsp, targets8)
	log("p6 w:%d", value)
	countwrites(value, 6, 31)
	l.Unlock()

	delay()
	l.Lock()
	value = 14
	Put(CreateTuple(211, value), rsp, targets9)
	log("p6 w:%d", value)
	countwrites(value, 6, 14)
	l.Unlock()

	delay()
	l.Lock()
	value = 27
	Put(CreateTuple(431, value), rsp, targets10)
	log("p6 w:%d", value)
	countwrites(value, 6, 27)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(220, &value), rsp, s6)
	log("p6 r:%d", value)
	countreads(value, 6, 14)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(429, &value), rsp, s6)
	log("p6 r:%d", value)
	countreads(value, 6, 27)
	l.Unlock()

	delay()
	l.Lock()
	value = 29
	Put(CreateTuple(460, value), rsp, targets11)
	log("p6 w:%d", value)
	countwrites(value, 6, 29)
	l.Unlock()

	delay()
	l.Lock()
	value = 22
	Put(CreateTuple(339, value), rsp, targets12)
	log("p6 w:%d", value)
	countwrites(value, 6, 22)
	l.Unlock()

	delay()
	l.Lock()
	value = 13
	Put(CreateTuple(193, value), rsp, targets13)
	log("p6 w:%d", value)
	countwrites(value, 6, 13)
	l.Unlock()
}

func p7() {
	targets9 := make([]string, 1)
	targets9[0] = uri[s22]
	targets8 := make([]string, 1)
	targets8[0] = uri[s4]
	targets7 := make([]string, 0)
	targets6 := make([]string, 0)
	targets5 := make([]string, 0)
	targets4 := make([]string, 1)
	targets4[0] = uri[s12]
	targets3 := make([]string, 1)
	targets3[0] = uri[s14]
	targets2 := make([]string, 1)
	targets2[0] = uri[s31]
	targets1 := make([]string, 0)
	targets0 := make([]string, 1)
	targets0[0] = uri[s6]

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(135, &value), rsp, s7)
	log("p7 r:%d", value)
	countreads(value, 7, 9)
	l.Unlock()

	delay()
	l.Lock()
	value = 27
	Put(CreateTuple(429, value), rsp, targets0)
	log("p7 w:%d", value)
	countwrites(value, 7, 27)
	l.Unlock()

	delay()
	l.Lock()
	value = 25
	Put(CreateTuple(393, value), rsp, targets1)
	log("p7 w:%d", value)
	countwrites(value, 7, 25)
	l.Unlock()

	delay()
	l.Lock()
	value = 2
	Put(CreateTuple(23, value), rsp, targets2)
	log("p7 w:%d", value)
	countwrites(value, 7, 2)
	l.Unlock()

	delay()
	l.Lock()
	value = 10
	Put(CreateTuple(157, value), rsp, targets3)
	log("p7 w:%d", value)
	countwrites(value, 7, 10)
	l.Unlock()

	delay()
	l.Lock()
	value = 22
	Put(CreateTuple(347, value), rsp, targets4)
	log("p7 w:%d", value)
	countwrites(value, 7, 22)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(128, &value), rsp, s7)
	log("p7 r:%d", value)
	countreads(value, 7, 8)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(226, &value), rsp, s7)
	log("p7 r:%d", value)
	countreads(value, 7, 15)
	l.Unlock()

	delay()
	l.Lock()
	value = 15
	Put(CreateTuple(229, value), rsp, targets5)
	log("p7 w:%d", value)
	countwrites(value, 7, 15)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(369, &value), rsp, s7)
	log("p7 r:%d", value)
	countreads(value, 7, 24)
	l.Unlock()

	delay()
	l.Lock()
	value = 19
	Put(CreateTuple(291, value), rsp, targets6)
	log("p7 w:%d", value)
	countwrites(value, 7, 19)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(341, &value), rsp, s7)
	log("p7 r:%d", value)
	countreads(value, 7, 22)
	l.Unlock()

	delay()
	l.Lock()
	value = 17
	Put(CreateTuple(268, value), rsp, targets7)
	log("p7 w:%d", value)
	countwrites(value, 7, 17)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(331, &value), rsp, s7)
	log("p7 r:%d", value)
	countreads(value, 7, 21)
	l.Unlock()

	delay()
	l.Lock()
	value = 5
	Put(CreateTuple(68, value), rsp, targets8)
	log("p7 w:%d", value)
	countwrites(value, 7, 5)
	l.Unlock()

	delay()
	l.Lock()
	value = 20
	Put(CreateTuple(308, value), rsp, targets9)
	log("p7 w:%d", value)
	countwrites(value, 7, 20)
	l.Unlock()
}

func p8() {
	targets12 := make([]string, 0)
	targets11 := make([]string, 0)
	targets10 := make([]string, 0)
	targets9 := make([]string, 0)
	targets8 := make([]string, 0)
	targets7 := make([]string, 1)
	targets7[0] = uri[s16]
	targets6 := make([]string, 2)
	targets6[1] = uri[s17]
	targets6[0] = uri[s11]
	targets5 := make([]string, 1)
	targets5[0] = uri[s20]
	targets4 := make([]string, 0)
	targets3 := make([]string, 0)
	targets2 := make([]string, 1)
	targets2[0] = uri[s7]
	targets1 := make([]string, 2)
	targets1[1] = uri[s18]
	targets1[0] = uri[s8]
	targets0 := make([]string, 1)
	targets0[0] = uri[s15]

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 1
	Put(CreateTuple(13, value), rsp, targets0)
	log("p8 w:%d", value)
	countwrites(value, 8, 1)
	l.Unlock()

	delay()
	l.Lock()
	value = 16
	Put(CreateTuple(252, value), rsp, targets1)
	log("p8 w:%d", value)
	countwrites(value, 8, 16)
	l.Unlock()

	delay()
	l.Lock()
	value = 8
	Put(CreateTuple(128, value), rsp, targets2)
	log("p8 w:%d", value)
	countwrites(value, 8, 8)
	l.Unlock()

	delay()
	l.Lock()
	value = 27
	Put(CreateTuple(417, value), rsp, targets3)
	log("p8 w:%d", value)
	countwrites(value, 8, 27)
	l.Unlock()

	delay()
	l.Lock()
	value = 22
	Put(CreateTuple(352, value), rsp, targets4)
	log("p8 w:%d", value)
	countwrites(value, 8, 22)
	l.Unlock()

	delay()
	l.Lock()
	value = 18
	Put(CreateTuple(274, value), rsp, targets5)
	log("p8 w:%d", value)
	countwrites(value, 8, 18)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(489, &value), rsp, s8)
	log("p8 r:%d", value)
	countreads(value, 8, 31)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(117, &value), rsp, s8)
	log("p8 r:%d", value)
	countreads(value, 8, 8)
	l.Unlock()

	delay()
	l.Lock()
	value = 18
	Put(CreateTuple(280, value), rsp, targets6)
	log("p8 w:%d", value)
	countwrites(value, 8, 18)
	l.Unlock()

	delay()
	l.Lock()
	value = 24
	Put(CreateTuple(371, value), rsp, targets7)
	log("p8 w:%d", value)
	countwrites(value, 8, 24)
	l.Unlock()

	delay()
	l.Lock()
	value = 30
	Put(CreateTuple(475, value), rsp, targets8)
	log("p8 w:%d", value)
	countwrites(value, 8, 30)
	l.Unlock()

	delay()
	l.Lock()
	value = 12
	Put(CreateTuple(179, value), rsp, targets9)
	log("p8 w:%d", value)
	countwrites(value, 8, 12)
	l.Unlock()

	delay()
	l.Lock()
	value = 4
	Put(CreateTuple(57, value), rsp, targets10)
	log("p8 w:%d", value)
	countwrites(value, 8, 4)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(252, &value), rsp, s8)
	log("p8 r:%d", value)
	countreads(value, 8, 16)
	l.Unlock()

	delay()
	l.Lock()
	value = 19
	Put(CreateTuple(290, value), rsp, targets11)
	log("p8 w:%d", value)
	countwrites(value, 8, 19)
	l.Unlock()

	delay()
	l.Lock()
	value = 24
	Put(CreateTuple(379, value), rsp, targets12)
	log("p8 w:%d", value)
	countwrites(value, 8, 24)
	l.Unlock()
}

func p9() {
	targets11 := make([]string, 2)
	targets11[1] = uri[s27]
	targets11[0] = uri[s2]
	targets10 := make([]string, 0)
	targets9 := make([]string, 1)
	targets9[0] = uri[s29]
	targets8 := make([]string, 0)
	targets7 := make([]string, 0)
	targets6 := make([]string, 1)
	targets6[0] = uri[s18]
	targets5 := make([]string, 0)
	targets4 := make([]string, 0)
	targets3 := make([]string, 1)
	targets3[0] = uri[s19]
	targets2 := make([]string, 1)
	targets2[0] = uri[s15]
	targets1 := make([]string, 1)
	targets1[0] = uri[s31]
	targets0 := make([]string, 0)

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(272, &value), rsp, s9)
	log("p9 r:%d", value)
	countreads(value, 9, 17)
	l.Unlock()

	delay()
	l.Lock()
	value = 6
	Put(CreateTuple(93, value), rsp, targets0)
	log("p9 w:%d", value)
	countwrites(value, 9, 6)
	l.Unlock()

	delay()
	l.Lock()
	value = 15
	Put(CreateTuple(230, value), rsp, targets1)
	log("p9 w:%d", value)
	countwrites(value, 9, 15)
	l.Unlock()

	delay()
	l.Lock()
	value = 6
	Put(CreateTuple(91, value), rsp, targets2)
	log("p9 w:%d", value)
	countwrites(value, 9, 6)
	l.Unlock()

	delay()
	l.Lock()
	value = 7
	Put(CreateTuple(105, value), rsp, targets3)
	log("p9 w:%d", value)
	countwrites(value, 9, 7)
	l.Unlock()

	delay()
	l.Lock()
	value = 12
	Put(CreateTuple(181, value), rsp, targets4)
	log("p9 w:%d", value)
	countwrites(value, 9, 12)
	l.Unlock()

	delay()
	l.Lock()
	value = 29
	Put(CreateTuple(454, value), rsp, targets5)
	log("p9 w:%d", value)
	countwrites(value, 9, 29)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(187, &value), rsp, s9)
	log("p9 r:%d", value)
	countreads(value, 9, 12)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(213, &value), rsp, s9)
	log("p9 r:%d", value)
	countreads(value, 9, 14)
	l.Unlock()

	delay()
	l.Lock()
	value = 12
	Put(CreateTuple(188, value), rsp, targets6)
	log("p9 w:%d", value)
	countwrites(value, 9, 12)
	l.Unlock()

	delay()
	l.Lock()
	value = 9
	Put(CreateTuple(137, value), rsp, targets7)
	log("p9 w:%d", value)
	countwrites(value, 9, 9)
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
	value = 28
	Put(CreateTuple(439, value), rsp, targets8)
	log("p9 w:%d", value)
	countwrites(value, 9, 28)
	l.Unlock()

	delay()
	l.Lock()
	value = 7
	Put(CreateTuple(100, value), rsp, targets9)
	log("p9 w:%d", value)
	countwrites(value, 9, 7)
	l.Unlock()

	delay()
	l.Lock()
	value = 12
	Put(CreateTuple(183, value), rsp, targets10)
	log("p9 w:%d", value)
	countwrites(value, 9, 12)
	l.Unlock()

	delay()
	l.Lock()
	value = 24
	Put(CreateTuple(378, value), rsp, targets11)
	log("p9 w:%d", value)
	countwrites(value, 9, 24)
	l.Unlock()
}

func p10() {
	targets11 := make([]string, 0)
	targets10 := make([]string, 0)
	targets9 := make([]string, 1)
	targets9[0] = uri[s14]
	targets8 := make([]string, 1)
	targets8[0] = uri[s22]
	targets7 := make([]string, 1)
	targets7[0] = uri[s29]
	targets6 := make([]string, 0)
	targets5 := make([]string, 1)
	targets5[0] = uri[s15]
	targets4 := make([]string, 1)
	targets4[0] = uri[s27]
	targets3 := make([]string, 0)
	targets2 := make([]string, 0)
	targets1 := make([]string, 0)
	targets0 := make([]string, 1)
	targets0[0] = uri[s22]

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 13
	Put(CreateTuple(194, value), rsp, targets0)
	log("p10 w:%d", value)
	countwrites(value, 10, 13)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(482, &value), rsp, s10)
	log("p10 r:%d", value)
	countreads(value, 10, 31)
	l.Unlock()

	delay()
	l.Lock()
	value = 24
	Put(CreateTuple(370, value), rsp, targets1)
	log("p10 w:%d", value)
	countwrites(value, 10, 24)
	l.Unlock()

	delay()
	l.Lock()
	value = 12
	Put(CreateTuple(181, value), rsp, targets2)
	log("p10 w:%d", value)
	countwrites(value, 10, 12)
	l.Unlock()

	delay()
	l.Lock()
	value = 31
	Put(CreateTuple(487, value), rsp, targets3)
	log("p10 w:%d", value)
	countwrites(value, 10, 31)
	l.Unlock()

	delay()
	l.Lock()
	value = 31
	Put(CreateTuple(484, value), rsp, targets4)
	log("p10 w:%d", value)
	countwrites(value, 10, 31)
	l.Unlock()

	delay()
	l.Lock()
	value = 1
	Put(CreateTuple(13, value), rsp, targets5)
	log("p10 w:%d", value)
	countwrites(value, 10, 1)
	l.Unlock()

	delay()
	l.Lock()
	value = 23
	Put(CreateTuple(362, value), rsp, targets6)
	log("p10 w:%d", value)
	countwrites(value, 10, 23)
	l.Unlock()

	delay()
	l.Lock()
	value = 7
	Put(CreateTuple(100, value), rsp, targets7)
	log("p10 w:%d", value)
	countwrites(value, 10, 7)
	l.Unlock()

	delay()
	l.Lock()
	value = 28
	Put(CreateTuple(436, value), rsp, targets8)
	log("p10 w:%d", value)
	countwrites(value, 10, 28)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(40, &value), rsp, s10)
	log("p10 r:%d", value)
	countreads(value, 10, 3)
	l.Unlock()

	delay()
	l.Lock()
	value = 4
	Put(CreateTuple(49, value), rsp, targets9)
	log("p10 w:%d", value)
	countwrites(value, 10, 4)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(342, &value), rsp, s10)
	log("p10 r:%d", value)
	countreads(value, 10, 22)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(237, &value), rsp, s10)
	log("p10 r:%d", value)
	countreads(value, 10, 15)
	l.Unlock()

	delay()
	l.Lock()
	value = 18
	Put(CreateTuple(281, value), rsp, targets10)
	log("p10 w:%d", value)
	countwrites(value, 10, 18)
	l.Unlock()

	delay()
	l.Lock()
	value = 24
	Put(CreateTuple(374, value), rsp, targets11)
	log("p10 w:%d", value)
	countwrites(value, 10, 24)
	l.Unlock()
}

func p11() {
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
	value = 1
	Put(CreateTuple(12, value), rsp, targets0)
	log("p11 w:%d", value)
	countwrites(value, 11, 1)
	l.Unlock()

	delay()
	l.Lock()
	value = 1
	Put(CreateTuple(11, value), rsp, targets1)
	log("p11 w:%d", value)
	countwrites(value, 11, 1)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(130, &value), rsp, s11)
	log("p11 r:%d", value)
	countreads(value, 11, 9)
	l.Unlock()

	delay()
	l.Lock()
	value = 18
	Put(CreateTuple(278, value), rsp, targets2)
	log("p11 w:%d", value)
	countwrites(value, 11, 18)
	l.Unlock()

	delay()
	l.Lock()
	value = 16
	Put(CreateTuple(254, value), rsp, targets3)
	log("p11 w:%d", value)
	countwrites(value, 11, 16)
	l.Unlock()

	delay()
	l.Lock()
	value = 25
	Put(CreateTuple(400, value), rsp, targets4)
	log("p11 w:%d", value)
	countwrites(value, 11, 25)
	l.Unlock()

	delay()
	l.Lock()
	value = 1
	Put(CreateTuple(8, value), rsp, targets5)
	log("p11 w:%d", value)
	countwrites(value, 11, 1)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(40, &value), rsp, s11)
	log("p11 r:%d", value)
	countreads(value, 11, 3)
	l.Unlock()

	delay()
	l.Lock()
	value = 8
	Put(CreateTuple(126, value), rsp, targets6)
	log("p11 w:%d", value)
	countwrites(value, 11, 8)
	l.Unlock()

	delay()
	l.Lock()
	value = 27
	Put(CreateTuple(427, value), rsp, targets7)
	log("p11 w:%d", value)
	countwrites(value, 11, 27)
	l.Unlock()

	delay()
	l.Lock()
	value = 24
	Put(CreateTuple(384, value), rsp, targets8)
	log("p11 w:%d", value)
	countwrites(value, 11, 24)
	l.Unlock()

	delay()
	l.Lock()
	value = 13
	Put(CreateTuple(202, value), rsp, targets9)
	log("p11 w:%d", value)
	countwrites(value, 11, 13)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(277, &value), rsp, s11)
	log("p11 r:%d", value)
	countreads(value, 11, 18)
	l.Unlock()

	delay()
	l.Lock()
	value = 5
	Put(CreateTuple(80, value), rsp, targets10)
	log("p11 w:%d", value)
	countwrites(value, 11, 5)
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
	value = 4
	Put(CreateTuple(63, value), rsp, targets11)
	log("p11 w:%d", value)
	countwrites(value, 11, 4)
	l.Unlock()
}

func p12() {
	targets9 := make([]string, 1)
	targets9[0] = uri[s18]
	targets8 := make([]string, 1)
	targets8[0] = uri[s26]
	targets7 := make([]string, 0)
	targets6 := make([]string, 1)
	targets6[0] = uri[s17]
	targets5 := make([]string, 1)
	targets5[0] = uri[s18]
	targets4 := make([]string, 0)
	targets3 := make([]string, 0)
	targets2 := make([]string, 0)
	targets1 := make([]string, 0)
	targets0 := make([]string, 0)

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 29
	Put(CreateTuple(449, value), rsp, targets0)
	log("p12 w:%d", value)
	countwrites(value, 12, 29)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(506, &value), rsp, s12)
	log("p12 r:%d", value)
	countreads(value, 12, 32)
	l.Unlock()

	delay()
	l.Lock()
	value = 8
	Put(CreateTuple(121, value), rsp, targets1)
	log("p12 w:%d", value)
	countwrites(value, 12, 8)
	l.Unlock()

	delay()
	l.Lock()
	value = 10
	Put(CreateTuple(146, value), rsp, targets2)
	log("p12 w:%d", value)
	countwrites(value, 12, 10)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(347, &value), rsp, s12)
	log("p12 r:%d", value)
	countreads(value, 12, 22)
	l.Unlock()

	delay()
	l.Lock()
	value = 32
	Put(CreateTuple(510, value), rsp, targets3)
	log("p12 w:%d", value)
	countwrites(value, 12, 32)
	l.Unlock()

	delay()
	l.Lock()
	value = 20
	Put(CreateTuple(311, value), rsp, targets4)
	log("p12 w:%d", value)
	countwrites(value, 12, 20)
	l.Unlock()

	delay()
	l.Lock()
	value = 17
	Put(CreateTuple(261, value), rsp, targets5)
	log("p12 w:%d", value)
	countwrites(value, 12, 17)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(82, &value), rsp, s12)
	log("p12 r:%d", value)
	countreads(value, 12, 6)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(346, &value), rsp, s12)
	log("p12 r:%d", value)
	countreads(value, 12, 22)
	l.Unlock()

	delay()
	l.Lock()
	value = 1
	Put(CreateTuple(6, value), rsp, targets6)
	log("p12 w:%d", value)
	countwrites(value, 12, 1)
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
	value = 13
	Put(CreateTuple(206, value), rsp, targets7)
	log("p12 w:%d", value)
	countwrites(value, 12, 13)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(455, &value), rsp, s12)
	log("p12 r:%d", value)
	countreads(value, 12, 29)
	l.Unlock()

	delay()
	l.Lock()
	value = 26
	Put(CreateTuple(401, value), rsp, targets8)
	log("p12 w:%d", value)
	countwrites(value, 12, 26)
	l.Unlock()

	delay()
	l.Lock()
	value = 23
	Put(CreateTuple(361, value), rsp, targets9)
	log("p12 w:%d", value)
	countwrites(value, 12, 23)
	l.Unlock()
}

func p13() {
	targets10 := make([]string, 0)
	targets9 := make([]string, 0)
	targets8 := make([]string, 0)
	targets7 := make([]string, 0)
	targets6 := make([]string, 0)
	targets5 := make([]string, 0)
	targets4 := make([]string, 1)
	targets4[0] = uri[s3]
	targets3 := make([]string, 0)
	targets2 := make([]string, 0)
	targets1 := make([]string, 0)
	targets0 := make([]string, 0)

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 7
	Put(CreateTuple(106, value), rsp, targets0)
	log("p13 w:%d", value)
	countwrites(value, 13, 7)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(210, &value), rsp, s13)
	log("p13 r:%d", value)
	countreads(value, 13, 14)
	l.Unlock()

	delay()
	l.Lock()
	value = 14
	Put(CreateTuple(211, value), rsp, targets1)
	log("p13 w:%d", value)
	countwrites(value, 13, 14)
	l.Unlock()

	delay()
	l.Lock()
	value = 25
	Put(CreateTuple(399, value), rsp, targets2)
	log("p13 w:%d", value)
	countwrites(value, 13, 25)
	l.Unlock()

	delay()
	l.Lock()
	value = 7
	Put(CreateTuple(109, value), rsp, targets3)
	log("p13 w:%d", value)
	countwrites(value, 13, 7)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(86, &value), rsp, s13)
	log("p13 r:%d", value)
	countreads(value, 13, 6)
	l.Unlock()

	delay()
	l.Lock()
	value = 1
	Put(CreateTuple(4, value), rsp, targets4)
	log("p13 w:%d", value)
	countwrites(value, 13, 1)
	l.Unlock()

	delay()
	l.Lock()
	value = 18
	Put(CreateTuple(278, value), rsp, targets5)
	log("p13 w:%d", value)
	countwrites(value, 13, 18)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(238, &value), rsp, s13)
	log("p13 r:%d", value)
	countreads(value, 13, 15)
	l.Unlock()

	delay()
	l.Lock()
	value = 22
	Put(CreateTuple(339, value), rsp, targets6)
	log("p13 w:%d", value)
	countwrites(value, 13, 22)
	l.Unlock()

	delay()
	l.Lock()
	value = 14
	Put(CreateTuple(223, value), rsp, targets7)
	log("p13 w:%d", value)
	countwrites(value, 13, 14)
	l.Unlock()

	delay()
	l.Lock()
	value = 8
	Put(CreateTuple(119, value), rsp, targets8)
	log("p13 w:%d", value)
	countwrites(value, 13, 8)
	l.Unlock()

	delay()
	l.Lock()
	value = 27
	Put(CreateTuple(430, value), rsp, targets9)
	log("p13 w:%d", value)
	countwrites(value, 13, 27)
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
	value = -1
	QueryP(CreateTuple(355, &value), rsp, s13)
	log("p13 r:%d", value)
	countreads(value, 13, 23)
	l.Unlock()

	delay()
	l.Lock()
	value = 17
	Put(CreateTuple(260, value), rsp, targets10)
	log("p13 w:%d", value)
	countwrites(value, 13, 17)
	l.Unlock()
}

func p14() {
	targets10 := make([]string, 0)
	targets9 := make([]string, 0)
	targets8 := make([]string, 0)
	targets7 := make([]string, 0)
	targets6 := make([]string, 1)
	targets6[0] = uri[s30]
	targets5 := make([]string, 0)
	targets4 := make([]string, 1)
	targets4[0] = uri[s7]
	targets3 := make([]string, 1)
	targets3[0] = uri[s3]
	targets2 := make([]string, 0)
	targets1 := make([]string, 0)
	targets0 := make([]string, 1)
	targets0[0] = uri[s13]

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(25, &value), rsp, s14)
	log("p14 r:%d", value)
	countreads(value, 14, 2)
	l.Unlock()

	delay()
	l.Lock()
	value = 4
	Put(CreateTuple(53, value), rsp, targets0)
	log("p14 w:%d", value)
	countwrites(value, 14, 4)
	l.Unlock()

	delay()
	l.Lock()
	value = 30
	Put(CreateTuple(468, value), rsp, targets1)
	log("p14 w:%d", value)
	countwrites(value, 14, 30)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(323, &value), rsp, s14)
	log("p14 r:%d", value)
	countreads(value, 14, 21)
	l.Unlock()

	delay()
	l.Lock()
	value = 12
	Put(CreateTuple(182, value), rsp, targets2)
	log("p14 w:%d", value)
	countwrites(value, 14, 12)
	l.Unlock()

	delay()
	l.Lock()
	value = 1
	Put(CreateTuple(4, value), rsp, targets3)
	log("p14 w:%d", value)
	countwrites(value, 14, 1)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(49, &value), rsp, s14)
	log("p14 r:%d", value)
	countreads(value, 14, 4)
	l.Unlock()

	delay()
	l.Lock()
	value = 22
	Put(CreateTuple(341, value), rsp, targets4)
	log("p14 w:%d", value)
	countwrites(value, 14, 22)
	l.Unlock()

	delay()
	l.Lock()
	value = 8
	Put(CreateTuple(116, value), rsp, targets5)
	log("p14 w:%d", value)
	countwrites(value, 14, 8)
	l.Unlock()

	delay()
	l.Lock()
	value = 18
	Put(CreateTuple(273, value), rsp, targets6)
	log("p14 w:%d", value)
	countwrites(value, 14, 18)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(157, &value), rsp, s14)
	log("p14 r:%d", value)
	countreads(value, 14, 10)
	l.Unlock()

	delay()
	l.Lock()
	value = 5
	Put(CreateTuple(65, value), rsp, targets7)
	log("p14 w:%d", value)
	countwrites(value, 14, 5)
	l.Unlock()

	delay()
	l.Lock()
	value = 21
	Put(CreateTuple(336, value), rsp, targets8)
	log("p14 w:%d", value)
	countwrites(value, 14, 21)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(50, &value), rsp, s14)
	log("p14 r:%d", value)
	countreads(value, 14, 4)
	l.Unlock()

	delay()
	l.Lock()
	value = 27
	Put(CreateTuple(430, value), rsp, targets9)
	log("p14 w:%d", value)
	countwrites(value, 14, 27)
	l.Unlock()

	delay()
	l.Lock()
	value = 5
	Put(CreateTuple(70, value), rsp, targets10)
	log("p14 w:%d", value)
	countwrites(value, 14, 5)
	l.Unlock()
}

func p15() {
	targets11 := make([]string, 0)
	targets10 := make([]string, 0)
	targets9 := make([]string, 1)
	targets9[0] = uri[s27]
	targets8 := make([]string, 0)
	targets7 := make([]string, 1)
	targets7[0] = uri[s2]
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
	value = 28
	Put(CreateTuple(445, value), rsp, targets0)
	log("p15 w:%d", value)
	countwrites(value, 15, 28)
	l.Unlock()

	delay()
	l.Lock()
	value = 5
	Put(CreateTuple(65, value), rsp, targets1)
	log("p15 w:%d", value)
	countwrites(value, 15, 5)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(13, &value), rsp, s15)
	log("p15 r:%d", value)
	countreads(value, 15, 1)
	l.Unlock()

	delay()
	l.Lock()
	value = 18
	Put(CreateTuple(276, value), rsp, targets2)
	log("p15 w:%d", value)
	countwrites(value, 15, 18)
	l.Unlock()

	delay()
	l.Lock()
	value = 32
	Put(CreateTuple(502, value), rsp, targets3)
	log("p15 w:%d", value)
	countwrites(value, 15, 32)
	l.Unlock()

	delay()
	l.Lock()
	value = 2
	Put(CreateTuple(21, value), rsp, targets4)
	log("p15 w:%d", value)
	countwrites(value, 15, 2)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(91, &value), rsp, s15)
	log("p15 r:%d", value)
	countreads(value, 15, 6)
	l.Unlock()

	delay()
	l.Lock()
	value = 4
	Put(CreateTuple(55, value), rsp, targets5)
	log("p15 w:%d", value)
	countwrites(value, 15, 4)
	l.Unlock()

	delay()
	l.Lock()
	value = 6
	Put(CreateTuple(96, value), rsp, targets6)
	log("p15 w:%d", value)
	countwrites(value, 15, 6)
	l.Unlock()

	delay()
	l.Lock()
	value = 17
	Put(CreateTuple(269, value), rsp, targets7)
	log("p15 w:%d", value)
	countwrites(value, 15, 17)
	l.Unlock()

	delay()
	l.Lock()
	value = 29
	Put(CreateTuple(452, value), rsp, targets8)
	log("p15 w:%d", value)
	countwrites(value, 15, 29)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(171, &value), rsp, s15)
	log("p15 r:%d", value)
	countreads(value, 15, 11)
	l.Unlock()

	delay()
	l.Lock()
	value = 18
	Put(CreateTuple(282, value), rsp, targets9)
	log("p15 w:%d", value)
	countwrites(value, 15, 18)
	l.Unlock()

	delay()
	l.Lock()
	value = 7
	Put(CreateTuple(103, value), rsp, targets10)
	log("p15 w:%d", value)
	countwrites(value, 15, 7)
	l.Unlock()

	delay()
	l.Lock()
	value = 27
	Put(CreateTuple(425, value), rsp, targets11)
	log("p15 w:%d", value)
	countwrites(value, 15, 27)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(267, &value), rsp, s15)
	log("p15 r:%d", value)
	countreads(value, 15, 17)
	l.Unlock()
}

func p16() {
	targets10 := make([]string, 1)
	targets10[0] = uri[s22]
	targets9 := make([]string, 1)
	targets9[0] = uri[s25]
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
	value = 20
	Put(CreateTuple(306, value), rsp, targets0)
	log("p16 w:%d", value)
	countwrites(value, 16, 20)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(300, &value), rsp, s16)
	log("p16 r:%d", value)
	countreads(value, 16, 19)
	l.Unlock()

	delay()
	l.Lock()
	value = 7
	Put(CreateTuple(112, value), rsp, targets1)
	log("p16 w:%d", value)
	countwrites(value, 16, 7)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(371, &value), rsp, s16)
	log("p16 r:%d", value)
	countreads(value, 16, 24)
	l.Unlock()

	delay()
	l.Lock()
	value = 5
	Put(CreateTuple(80, value), rsp, targets2)
	log("p16 w:%d", value)
	countwrites(value, 16, 5)
	l.Unlock()

	delay()
	l.Lock()
	value = 30
	Put(CreateTuple(470, value), rsp, targets3)
	log("p16 w:%d", value)
	countwrites(value, 16, 30)
	l.Unlock()

	delay()
	l.Lock()
	value = 19
	Put(CreateTuple(298, value), rsp, targets4)
	log("p16 w:%d", value)
	countwrites(value, 16, 19)
	l.Unlock()

	delay()
	l.Lock()
	value = 8
	Put(CreateTuple(115, value), rsp, targets5)
	log("p16 w:%d", value)
	countwrites(value, 16, 8)
	l.Unlock()

	delay()
	l.Lock()
	value = 16
	Put(CreateTuple(256, value), rsp, targets6)
	log("p16 w:%d", value)
	countwrites(value, 16, 16)
	l.Unlock()

	delay()
	l.Lock()
	value = 22
	Put(CreateTuple(345, value), rsp, targets7)
	log("p16 w:%d", value)
	countwrites(value, 16, 22)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(17, &value), rsp, s16)
	log("p16 r:%d", value)
	countreads(value, 16, 2)
	l.Unlock()

	delay()
	l.Lock()
	value = 6
	Put(CreateTuple(84, value), rsp, targets8)
	log("p16 w:%d", value)
	countwrites(value, 16, 6)
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
	QueryP(CreateTuple(69, &value), rsp, s16)
	log("p16 r:%d", value)
	countreads(value, 16, 5)
	l.Unlock()

	delay()
	l.Lock()
	value = 32
	Put(CreateTuple(498, value), rsp, targets9)
	log("p16 w:%d", value)
	countwrites(value, 16, 32)
	l.Unlock()

	delay()
	l.Lock()
	value = 20
	Put(CreateTuple(308, value), rsp, targets10)
	log("p16 w:%d", value)
	countwrites(value, 16, 20)
	l.Unlock()
}

func p17() {
	targets7 := make([]string, 2)
	targets7[1] = uri[s27]
	targets7[0] = uri[s2]
	targets6 := make([]string, 3)
	targets6[2] = uri[s17]
	targets6[1] = uri[s11]
	targets6[0] = uri[s10]
	targets5 := make([]string, 0)
	targets4 := make([]string, 2)
	targets4[1] = uri[s17]
	targets4[0] = uri[s13]
	targets3 := make([]string, 0)
	targets2 := make([]string, 0)
	targets1 := make([]string, 0)
	targets0 := make([]string, 0)

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(86, &value), rsp, s17)
	log("p17 r:%d", value)
	countreads(value, 17, 6)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(39, &value), rsp, s17)
	log("p17 r:%d", value)
	countreads(value, 17, 3)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(280, &value), rsp, s17)
	log("p17 r:%d", value)
	countreads(value, 17, 18)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(180, &value), rsp, s17)
	log("p17 r:%d", value)
	countreads(value, 17, 12)
	l.Unlock()

	delay()
	l.Lock()
	value = 10
	Put(CreateTuple(159, value), rsp, targets0)
	log("p17 w:%d", value)
	countwrites(value, 17, 10)
	l.Unlock()

	delay()
	l.Lock()
	value = 5
	Put(CreateTuple(80, value), rsp, targets1)
	log("p17 w:%d", value)
	countwrites(value, 17, 5)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(42, &value), rsp, s17)
	log("p17 r:%d", value)
	countreads(value, 17, 3)
	l.Unlock()

	delay()
	l.Lock()
	value = 32
	Put(CreateTuple(512, value), rsp, targets2)
	log("p17 w:%d", value)
	countwrites(value, 17, 32)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(120, &value), rsp, s17)
	log("p17 r:%d", value)
	countreads(value, 17, 8)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(40, &value), rsp, s17)
	log("p17 r:%d", value)
	countreads(value, 17, 3)
	l.Unlock()

	delay()
	l.Lock()
	value = 19
	Put(CreateTuple(296, value), rsp, targets3)
	log("p17 w:%d", value)
	countwrites(value, 17, 19)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(6, &value), rsp, s17)
	log("p17 r:%d", value)
	countreads(value, 17, 1)
	l.Unlock()

	delay()
	l.Lock()
	value = 6
	Put(CreateTuple(86, value), rsp, targets4)
	log("p17 w:%d", value)
	countwrites(value, 17, 6)
	l.Unlock()

	delay()
	l.Lock()
	value = 4
	Put(CreateTuple(52, value), rsp, targets5)
	log("p17 w:%d", value)
	countwrites(value, 17, 4)
	l.Unlock()

	delay()
	l.Lock()
	value = 3
	Put(CreateTuple(40, value), rsp, targets6)
	log("p17 w:%d", value)
	countwrites(value, 17, 3)
	l.Unlock()

	delay()
	l.Lock()
	value = 24
	Put(CreateTuple(378, value), rsp, targets7)
	log("p17 w:%d", value)
	countwrites(value, 17, 24)
	l.Unlock()
}

func p18() {
	targets9 := make([]string, 1)
	targets9[0] = uri[s14]
	targets8 := make([]string, 0)
	targets7 := make([]string, 0)
	targets6 := make([]string, 0)
	targets5 := make([]string, 0)
	targets4 := make([]string, 2)
	targets4[1] = uri[s22]
	targets4[0] = uri[s21]
	targets3 := make([]string, 0)
	targets2 := make([]string, 0)
	targets1 := make([]string, 2)
	targets1[1] = uri[s12]
	targets1[0] = uri[s4]
	targets0 := make([]string, 0)

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 32
	Put(CreateTuple(509, value), rsp, targets0)
	log("p18 w:%d", value)
	countwrites(value, 18, 32)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(252, &value), rsp, s18)
	log("p18 r:%d", value)
	countreads(value, 18, 16)
	l.Unlock()

	delay()
	l.Lock()
	value = 6
	Put(CreateTuple(82, value), rsp, targets1)
	log("p18 w:%d", value)
	countwrites(value, 18, 6)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(271, &value), rsp, s18)
	log("p18 r:%d", value)
	countreads(value, 18, 17)
	l.Unlock()

	delay()
	l.Lock()
	value = 11
	Put(CreateTuple(175, value), rsp, targets2)
	log("p18 w:%d", value)
	countwrites(value, 18, 11)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(188, &value), rsp, s18)
	log("p18 r:%d", value)
	countreads(value, 18, 12)
	l.Unlock()

	delay()
	l.Lock()
	value = 20
	Put(CreateTuple(309, value), rsp, targets3)
	log("p18 w:%d", value)
	countwrites(value, 18, 20)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(386, &value), rsp, s18)
	log("p18 r:%d", value)
	countreads(value, 18, 25)
	l.Unlock()

	delay()
	l.Lock()
	value = 12
	Put(CreateTuple(192, value), rsp, targets4)
	log("p18 w:%d", value)
	countwrites(value, 18, 12)
	l.Unlock()

	delay()
	l.Lock()
	value = 4
	Put(CreateTuple(63, value), rsp, targets5)
	log("p18 w:%d", value)
	countwrites(value, 18, 4)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(261, &value), rsp, s18)
	log("p18 r:%d", value)
	countreads(value, 18, 17)
	l.Unlock()

	delay()
	l.Lock()
	value = 20
	Put(CreateTuple(318, value), rsp, targets6)
	log("p18 w:%d", value)
	countwrites(value, 18, 20)
	l.Unlock()

	delay()
	l.Lock()
	value = 6
	Put(CreateTuple(83, value), rsp, targets7)
	log("p18 w:%d", value)
	countwrites(value, 18, 6)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(361, &value), rsp, s18)
	log("p18 r:%d", value)
	countreads(value, 18, 23)
	l.Unlock()

	delay()
	l.Lock()
	value = 19
	Put(CreateTuple(290, value), rsp, targets8)
	log("p18 w:%d", value)
	countwrites(value, 18, 19)
	l.Unlock()

	delay()
	l.Lock()
	value = 21
	Put(CreateTuple(323, value), rsp, targets9)
	log("p18 w:%d", value)
	countwrites(value, 18, 21)
	l.Unlock()
}

func p19() {
	targets11 := make([]string, 0)
	targets10 := make([]string, 0)
	targets9 := make([]string, 0)
	targets8 := make([]string, 0)
	targets7 := make([]string, 1)
	targets7[0] = uri[s16]
	targets6 := make([]string, 0)
	targets5 := make([]string, 1)
	targets5[0] = uri[s21]
	targets4 := make([]string, 1)
	targets4[0] = uri[s10]
	targets3 := make([]string, 0)
	targets2 := make([]string, 0)
	targets1 := make([]string, 0)
	targets0 := make([]string, 0)

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 9
	Put(CreateTuple(138, value), rsp, targets0)
	log("p19 w:%d", value)
	countwrites(value, 19, 9)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(75, &value), rsp, s19)
	log("p19 r:%d", value)
	countreads(value, 19, 5)
	l.Unlock()

	delay()
	l.Lock()
	value = 31
	Put(CreateTuple(492, value), rsp, targets1)
	log("p19 w:%d", value)
	countwrites(value, 19, 31)
	l.Unlock()

	delay()
	l.Lock()
	value = 26
	Put(CreateTuple(406, value), rsp, targets2)
	log("p19 w:%d", value)
	countwrites(value, 19, 26)
	l.Unlock()

	delay()
	l.Lock()
	value = 32
	Put(CreateTuple(502, value), rsp, targets3)
	log("p19 w:%d", value)
	countwrites(value, 19, 32)
	l.Unlock()

	delay()
	l.Lock()
	value = 22
	Put(CreateTuple(342, value), rsp, targets4)
	log("p19 w:%d", value)
	countwrites(value, 19, 22)
	l.Unlock()

	delay()
	l.Lock()
	value = 21
	Put(CreateTuple(335, value), rsp, targets5)
	log("p19 w:%d", value)
	countwrites(value, 19, 21)
	l.Unlock()

	delay()
	l.Lock()
	value = 19
	Put(CreateTuple(302, value), rsp, targets6)
	log("p19 w:%d", value)
	countwrites(value, 19, 19)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(105, &value), rsp, s19)
	log("p19 r:%d", value)
	countreads(value, 19, 7)
	l.Unlock()

	delay()
	l.Lock()
	value = 2
	Put(CreateTuple(17, value), rsp, targets7)
	log("p19 w:%d", value)
	countwrites(value, 19, 2)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(367, &value), rsp, s19)
	log("p19 r:%d", value)
	countreads(value, 19, 23)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(213, &value), rsp, s19)
	log("p19 r:%d", value)
	countreads(value, 19, 14)
	l.Unlock()

	delay()
	l.Lock()
	value = 20
	Put(CreateTuple(309, value), rsp, targets8)
	log("p19 w:%d", value)
	countwrites(value, 19, 20)
	l.Unlock()

	delay()
	l.Lock()
	value = 20
	Put(CreateTuple(318, value), rsp, targets9)
	log("p19 w:%d", value)
	countwrites(value, 19, 20)
	l.Unlock()

	delay()
	l.Lock()
	value = 6
	Put(CreateTuple(94, value), rsp, targets10)
	log("p19 w:%d", value)
	countwrites(value, 19, 6)
	l.Unlock()

	delay()
	l.Lock()
	value = 22
	Put(CreateTuple(349, value), rsp, targets11)
	log("p19 w:%d", value)
	countwrites(value, 19, 22)
	l.Unlock()
}

func p20() {
	targets10 := make([]string, 0)
	targets9 := make([]string, 0)
	targets8 := make([]string, 0)
	targets7 := make([]string, 1)
	targets7[0] = uri[s22]
	targets6 := make([]string, 0)
	targets5 := make([]string, 0)
	targets4 := make([]string, 0)
	targets3 := make([]string, 0)
	targets2 := make([]string, 0)
	targets1 := make([]string, 0)
	targets0 := make([]string, 1)
	targets0[0] = uri[s6]

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(435, &value), rsp, s20)
	log("p20 r:%d", value)
	countreads(value, 20, 28)
	l.Unlock()

	delay()
	l.Lock()
	value = 14
	Put(CreateTuple(220, value), rsp, targets0)
	log("p20 w:%d", value)
	countwrites(value, 20, 14)
	l.Unlock()

	delay()
	l.Lock()
	value = 27
	Put(CreateTuple(428, value), rsp, targets1)
	log("p20 w:%d", value)
	countwrites(value, 20, 27)
	l.Unlock()

	delay()
	l.Lock()
	value = 17
	Put(CreateTuple(262, value), rsp, targets2)
	log("p20 w:%d", value)
	countwrites(value, 20, 17)
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
	value = -1
	QueryP(CreateTuple(321, &value), rsp, s20)
	log("p20 r:%d", value)
	countreads(value, 20, 21)
	l.Unlock()

	delay()
	l.Lock()
	value = 20
	Put(CreateTuple(316, value), rsp, targets3)
	log("p20 w:%d", value)
	countwrites(value, 20, 20)
	l.Unlock()

	delay()
	l.Lock()
	value = 5
	Put(CreateTuple(77, value), rsp, targets4)
	log("p20 w:%d", value)
	countwrites(value, 20, 5)
	l.Unlock()

	delay()
	l.Lock()
	value = 16
	Put(CreateTuple(251, value), rsp, targets5)
	log("p20 w:%d", value)
	countwrites(value, 20, 16)
	l.Unlock()

	delay()
	l.Lock()
	value = 16
	Put(CreateTuple(243, value), rsp, targets6)
	log("p20 w:%d", value)
	countwrites(value, 20, 16)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(274, &value), rsp, s20)
	log("p20 r:%d", value)
	countreads(value, 20, 18)
	l.Unlock()

	delay()
	l.Lock()
	value = 13
	Put(CreateTuple(207, value), rsp, targets7)
	log("p20 w:%d", value)
	countwrites(value, 20, 13)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(28, &value), rsp, s20)
	log("p20 r:%d", value)
	countreads(value, 20, 2)
	l.Unlock()

	delay()
	l.Lock()
	value = 5
	Put(CreateTuple(70, value), rsp, targets8)
	log("p20 w:%d", value)
	countwrites(value, 20, 5)
	l.Unlock()

	delay()
	l.Lock()
	value = 16
	Put(CreateTuple(251, value), rsp, targets9)
	log("p20 w:%d", value)
	countwrites(value, 20, 16)
	l.Unlock()

	delay()
	l.Lock()
	value = 5
	Put(CreateTuple(79, value), rsp, targets10)
	log("p20 w:%d", value)
	countwrites(value, 20, 5)
	l.Unlock()
}

func p21() {
	targets9 := make([]string, 1)
	targets9[0] = uri[s23]
	targets8 := make([]string, 0)
	targets7 := make([]string, 1)
	targets7[0] = uri[s21]
	targets6 := make([]string, 1)
	targets6[0] = uri[s31]
	targets5 := make([]string, 2)
	targets5[1] = uri[s29]
	targets5[0] = uri[s17]
	targets4 := make([]string, 1)
	targets4[0] = uri[s5]
	targets3 := make([]string, 0)
	targets2 := make([]string, 1)
	targets2[0] = uri[s19]
	targets1 := make([]string, 0)
	targets0 := make([]string, 0)

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(224, &value), rsp, s21)
	log("p21 r:%d", value)
	countreads(value, 21, 14)
	l.Unlock()

	delay()
	l.Lock()
	value = 30
	Put(CreateTuple(467, value), rsp, targets0)
	log("p21 w:%d", value)
	countwrites(value, 21, 30)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(359, &value), rsp, s21)
	log("p21 r:%d", value)
	countreads(value, 21, 23)
	l.Unlock()

	delay()
	l.Lock()
	value = 29
	Put(CreateTuple(453, value), rsp, targets1)
	log("p21 w:%d", value)
	countwrites(value, 21, 29)
	l.Unlock()

	delay()
	l.Lock()
	value = 7
	Put(CreateTuple(105, value), rsp, targets2)
	log("p21 w:%d", value)
	countwrites(value, 21, 7)
	l.Unlock()

	delay()
	l.Lock()
	value = 1
	Put(CreateTuple(10, value), rsp, targets3)
	log("p21 w:%d", value)
	countwrites(value, 21, 1)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(192, &value), rsp, s21)
	log("p21 r:%d", value)
	countreads(value, 21, 12)
	l.Unlock()

	delay()
	l.Lock()
	value = 4
	Put(CreateTuple(64, value), rsp, targets4)
	log("p21 w:%d", value)
	countwrites(value, 21, 4)
	l.Unlock()

	delay()
	l.Lock()
	value = 12
	Put(CreateTuple(180, value), rsp, targets5)
	log("p21 w:%d", value)
	countwrites(value, 21, 12)
	l.Unlock()

	delay()
	l.Lock()
	value = 15
	Put(CreateTuple(230, value), rsp, targets6)
	log("p21 w:%d", value)
	countwrites(value, 21, 15)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(185, &value), rsp, s21)
	log("p21 r:%d", value)
	countreads(value, 21, 12)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(335, &value), rsp, s21)
	log("p21 r:%d", value)
	countreads(value, 21, 21)
	l.Unlock()

	delay()
	l.Lock()
	value = 12
	Put(CreateTuple(185, value), rsp, targets7)
	log("p21 w:%d", value)
	countwrites(value, 21, 12)
	l.Unlock()

	delay()
	l.Lock()
	value = 4
	Put(CreateTuple(56, value), rsp, targets8)
	log("p21 w:%d", value)
	countwrites(value, 21, 4)
	l.Unlock()

	delay()
	l.Lock()
	value = 22
	Put(CreateTuple(348, value), rsp, targets9)
	log("p21 w:%d", value)
	countwrites(value, 21, 22)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(267, &value), rsp, s21)
	log("p21 r:%d", value)
	countreads(value, 21, 17)
	l.Unlock()
}

func p22() {
	targets7 := make([]string, 0)
	targets6 := make([]string, 0)
	targets5 := make([]string, 0)
	targets4 := make([]string, 0)
	targets3 := make([]string, 0)
	targets2 := make([]string, 0)
	targets1 := make([]string, 2)
	targets1[1] = uri[s17]
	targets1[0] = uri[s13]
	targets0 := make([]string, 1)
	targets0[0] = uri[s8]

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 8
	Put(CreateTuple(117, value), rsp, targets0)
	log("p22 w:%d", value)
	countwrites(value, 22, 8)
	l.Unlock()

	delay()
	l.Lock()
	value = 6
	Put(CreateTuple(86, value), rsp, targets1)
	log("p22 w:%d", value)
	countwrites(value, 22, 6)
	l.Unlock()

	delay()
	l.Lock()
	value = 29
	Put(CreateTuple(461, value), rsp, targets2)
	log("p22 w:%d", value)
	countwrites(value, 22, 29)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(66, &value), rsp, s22)
	log("p22 r:%d", value)
	countreads(value, 22, 5)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(207, &value), rsp, s22)
	log("p22 r:%d", value)
	countreads(value, 22, 13)
	l.Unlock()

	delay()
	l.Lock()
	value = 18
	Put(CreateTuple(287, value), rsp, targets3)
	log("p22 w:%d", value)
	countwrites(value, 22, 18)
	l.Unlock()

	delay()
	l.Lock()
	value = 17
	Put(CreateTuple(262, value), rsp, targets4)
	log("p22 w:%d", value)
	countwrites(value, 22, 17)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(308, &value), rsp, s22)
	log("p22 r:%d", value)
	countreads(value, 22, 20)
	l.Unlock()

	delay()
	l.Lock()
	value = 14
	Put(CreateTuple(218, value), rsp, targets5)
	log("p22 w:%d", value)
	countwrites(value, 22, 14)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(436, &value), rsp, s22)
	log("p22 r:%d", value)
	countreads(value, 22, 28)
	l.Unlock()

	delay()
	l.Lock()
	value = 9
	Put(CreateTuple(143, value), rsp, targets6)
	log("p22 w:%d", value)
	countwrites(value, 22, 9)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(203, &value), rsp, s22)
	log("p22 r:%d", value)
	countreads(value, 22, 13)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(192, &value), rsp, s22)
	log("p22 r:%d", value)
	countreads(value, 22, 12)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(194, &value), rsp, s22)
	log("p22 r:%d", value)
	countreads(value, 22, 13)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(31, &value), rsp, s22)
	log("p22 r:%d", value)
	countreads(value, 22, 2)
	l.Unlock()

	delay()
	l.Lock()
	value = 2
	Put(CreateTuple(21, value), rsp, targets7)
	log("p22 w:%d", value)
	countwrites(value, 22, 2)
	l.Unlock()
}

func p23() {
	targets12 := make([]string, 2)
	targets12[1] = uri[s17]
	targets12[0] = uri[s13]
	targets11 := make([]string, 0)
	targets10 := make([]string, 1)
	targets10[0] = uri[s23]
	targets9 := make([]string, 0)
	targets8 := make([]string, 0)
	targets7 := make([]string, 0)
	targets6 := make([]string, 0)
	targets5 := make([]string, 0)
	targets4 := make([]string, 1)
	targets4[0] = uri[s10]
	targets3 := make([]string, 0)
	targets2 := make([]string, 0)
	targets1 := make([]string, 1)
	targets1[0] = uri[s7]
	targets0 := make([]string, 0)

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(348, &value), rsp, s23)
	log("p23 r:%d", value)
	countreads(value, 23, 22)
	l.Unlock()

	delay()
	l.Lock()
	value = 29
	Put(CreateTuple(458, value), rsp, targets0)
	log("p23 w:%d", value)
	countwrites(value, 23, 29)
	l.Unlock()

	delay()
	l.Lock()
	value = 9
	Put(CreateTuple(135, value), rsp, targets1)
	log("p23 w:%d", value)
	countwrites(value, 23, 9)
	l.Unlock()

	delay()
	l.Lock()
	value = 12
	Put(CreateTuple(186, value), rsp, targets2)
	log("p23 w:%d", value)
	countwrites(value, 23, 12)
	l.Unlock()

	delay()
	l.Lock()
	value = 27
	Put(CreateTuple(425, value), rsp, targets3)
	log("p23 w:%d", value)
	countwrites(value, 23, 27)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(286, &value), rsp, s23)
	log("p23 r:%d", value)
	countreads(value, 23, 18)
	l.Unlock()

	delay()
	l.Lock()
	value = 31
	Put(CreateTuple(482, value), rsp, targets4)
	log("p23 w:%d", value)
	countwrites(value, 23, 31)
	l.Unlock()

	delay()
	l.Lock()
	value = 32
	Put(CreateTuple(501, value), rsp, targets5)
	log("p23 w:%d", value)
	countwrites(value, 23, 32)
	l.Unlock()

	delay()
	l.Lock()
	value = 13
	Put(CreateTuple(201, value), rsp, targets6)
	log("p23 w:%d", value)
	countwrites(value, 23, 13)
	l.Unlock()

	delay()
	l.Lock()
	value = 17
	Put(CreateTuple(265, value), rsp, targets7)
	log("p23 w:%d", value)
	countwrites(value, 23, 17)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(344, &value), rsp, s23)
	log("p23 r:%d", value)
	countreads(value, 23, 22)
	l.Unlock()

	delay()
	l.Lock()
	value = 27
	Put(CreateTuple(431, value), rsp, targets8)
	log("p23 w:%d", value)
	countwrites(value, 23, 27)
	l.Unlock()

	delay()
	l.Lock()
	value = 11
	Put(CreateTuple(169, value), rsp, targets9)
	log("p23 w:%d", value)
	countwrites(value, 23, 11)
	l.Unlock()

	delay()
	l.Lock()
	value = 22
	Put(CreateTuple(344, value), rsp, targets10)
	log("p23 w:%d", value)
	countwrites(value, 23, 22)
	l.Unlock()

	delay()
	l.Lock()
	value = 21
	Put(CreateTuple(332, value), rsp, targets11)
	log("p23 w:%d", value)
	countwrites(value, 23, 21)
	l.Unlock()

	delay()
	l.Lock()
	value = 6
	Put(CreateTuple(86, value), rsp, targets12)
	log("p23 w:%d", value)
	countwrites(value, 23, 6)
	l.Unlock()
}

func p24() {
	targets13 := make([]string, 0)
	targets12 := make([]string, 1)
	targets12[0] = uri[s3]
	targets11 := make([]string, 0)
	targets10 := make([]string, 0)
	targets9 := make([]string, 0)
	targets8 := make([]string, 0)
	targets7 := make([]string, 0)
	targets6 := make([]string, 0)
	targets5 := make([]string, 1)
	targets5[0] = uri[s12]
	targets4 := make([]string, 1)
	targets4[0] = uri[s30]
	targets3 := make([]string, 1)
	targets3[0] = uri[s27]
	targets2 := make([]string, 0)
	targets1 := make([]string, 0)
	targets0 := make([]string, 0)

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 26
	Put(CreateTuple(408, value), rsp, targets0)
	log("p24 w:%d", value)
	countwrites(value, 24, 26)
	l.Unlock()

	delay()
	l.Lock()
	value = 2
	Put(CreateTuple(27, value), rsp, targets1)
	log("p24 w:%d", value)
	countwrites(value, 24, 2)
	l.Unlock()

	delay()
	l.Lock()
	value = 23
	Put(CreateTuple(365, value), rsp, targets2)
	log("p24 w:%d", value)
	countwrites(value, 24, 23)
	l.Unlock()

	delay()
	l.Lock()
	value = 30
	Put(CreateTuple(476, value), rsp, targets3)
	log("p24 w:%d", value)
	countwrites(value, 24, 30)
	l.Unlock()

	delay()
	l.Lock()
	value = 10
	Put(CreateTuple(160, value), rsp, targets4)
	log("p24 w:%d", value)
	countwrites(value, 24, 10)
	l.Unlock()

	delay()
	l.Lock()
	value = 32
	Put(CreateTuple(506, value), rsp, targets5)
	log("p24 w:%d", value)
	countwrites(value, 24, 32)
	l.Unlock()

	delay()
	l.Lock()
	value = 16
	Put(CreateTuple(255, value), rsp, targets6)
	log("p24 w:%d", value)
	countwrites(value, 24, 16)
	l.Unlock()

	delay()
	l.Lock()
	value = 24
	Put(CreateTuple(384, value), rsp, targets7)
	log("p24 w:%d", value)
	countwrites(value, 24, 24)
	l.Unlock()

	delay()
	l.Lock()
	value = 14
	Put(CreateTuple(218, value), rsp, targets8)
	log("p24 w:%d", value)
	countwrites(value, 24, 14)
	l.Unlock()

	delay()
	l.Lock()
	value = 18
	Put(CreateTuple(284, value), rsp, targets9)
	log("p24 w:%d", value)
	countwrites(value, 24, 18)
	l.Unlock()

	delay()
	l.Lock()
	value = 8
	Put(CreateTuple(113, value), rsp, targets10)
	log("p24 w:%d", value)
	countwrites(value, 24, 8)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(191, &value), rsp, s24)
	log("p24 r:%d", value)
	countreads(value, 24, 12)
	l.Unlock()

	delay()
	l.Lock()
	value = 23
	Put(CreateTuple(368, value), rsp, targets11)
	log("p24 w:%d", value)
	countwrites(value, 24, 23)
	l.Unlock()

	delay()
	l.Lock()
	value = 20
	Put(CreateTuple(315, value), rsp, targets12)
	log("p24 w:%d", value)
	countwrites(value, 24, 20)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(327, &value), rsp, s24)
	log("p24 r:%d", value)
	countreads(value, 24, 21)
	l.Unlock()

	delay()
	l.Lock()
	value = 3
	Put(CreateTuple(38, value), rsp, targets13)
	log("p24 w:%d", value)
	countwrites(value, 24, 3)
	l.Unlock()
}

func p25() {
	targets10 := make([]string, 0)
	targets9 := make([]string, 0)
	targets8 := make([]string, 0)
	targets7 := make([]string, 0)
	targets6 := make([]string, 0)
	targets5 := make([]string, 1)
	targets5[0] = uri[s22]
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
	value = 3
	Put(CreateTuple(47, value), rsp, targets0)
	log("p25 w:%d", value)
	countwrites(value, 25, 3)
	l.Unlock()

	delay()
	l.Lock()
	value = 16
	Put(CreateTuple(246, value), rsp, targets1)
	log("p25 w:%d", value)
	countwrites(value, 25, 16)
	l.Unlock()

	delay()
	l.Lock()
	value = 1
	Put(CreateTuple(10, value), rsp, targets2)
	log("p25 w:%d", value)
	countwrites(value, 25, 1)
	l.Unlock()

	delay()
	l.Lock()
	value = 19
	Put(CreateTuple(290, value), rsp, targets3)
	log("p25 w:%d", value)
	countwrites(value, 25, 19)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(124, &value), rsp, s25)
	log("p25 r:%d", value)
	countreads(value, 25, 8)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(120, &value), rsp, s25)
	log("p25 r:%d", value)
	countreads(value, 25, 8)
	l.Unlock()

	delay()
	l.Lock()
	value = 2
	Put(CreateTuple(17, value), rsp, targets4)
	log("p25 w:%d", value)
	countwrites(value, 25, 2)
	l.Unlock()

	delay()
	l.Lock()
	value = 13
	Put(CreateTuple(207, value), rsp, targets5)
	log("p25 w:%d", value)
	countwrites(value, 25, 13)
	l.Unlock()

	delay()
	l.Lock()
	value = 26
	Put(CreateTuple(411, value), rsp, targets6)
	log("p25 w:%d", value)
	countwrites(value, 25, 26)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(498, &value), rsp, s25)
	log("p25 r:%d", value)
	countreads(value, 25, 32)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(232, &value), rsp, s25)
	log("p25 r:%d", value)
	countreads(value, 25, 15)
	l.Unlock()

	delay()
	l.Lock()
	value = 11
	Put(CreateTuple(175, value), rsp, targets7)
	log("p25 w:%d", value)
	countwrites(value, 25, 11)
	l.Unlock()

	delay()
	l.Lock()
	value = 3
	Put(CreateTuple(38, value), rsp, targets8)
	log("p25 w:%d", value)
	countwrites(value, 25, 3)
	l.Unlock()

	delay()
	l.Lock()
	value = 8
	Put(CreateTuple(123, value), rsp, targets9)
	log("p25 w:%d", value)
	countwrites(value, 25, 8)
	l.Unlock()

	delay()
	l.Lock()
	value = 32
	Put(CreateTuple(507, value), rsp, targets10)
	log("p25 w:%d", value)
	countwrites(value, 25, 32)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(242, &value), rsp, s25)
	log("p25 r:%d", value)
	countreads(value, 25, 16)
	l.Unlock()
}

func p26() {
	targets11 := make([]string, 0)
	targets10 := make([]string, 0)
	targets9 := make([]string, 0)
	targets8 := make([]string, 0)
	targets7 := make([]string, 0)
	targets6 := make([]string, 0)
	targets5 := make([]string, 1)
	targets5[0] = uri[s18]
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
	Put(CreateTuple(343, value), rsp, targets0)
	log("p26 w:%d", value)
	countwrites(value, 26, 22)
	l.Unlock()

	delay()
	l.Lock()
	value = 6
	Put(CreateTuple(93, value), rsp, targets1)
	log("p26 w:%d", value)
	countwrites(value, 26, 6)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(184, &value), rsp, s26)
	log("p26 r:%d", value)
	countreads(value, 26, 12)
	l.Unlock()

	delay()
	l.Lock()
	value = 30
	Put(CreateTuple(467, value), rsp, targets2)
	log("p26 w:%d", value)
	countwrites(value, 26, 30)
	l.Unlock()

	delay()
	l.Lock()
	value = 26
	Put(CreateTuple(412, value), rsp, targets3)
	log("p26 w:%d", value)
	countwrites(value, 26, 26)
	l.Unlock()

	delay()
	l.Lock()
	value = 14
	Put(CreateTuple(209, value), rsp, targets4)
	log("p26 w:%d", value)
	countwrites(value, 26, 14)
	l.Unlock()

	delay()
	l.Lock()
	value = 17
	Put(CreateTuple(261, value), rsp, targets5)
	log("p26 w:%d", value)
	countwrites(value, 26, 17)
	l.Unlock()

	delay()
	l.Lock()
	value = 18
	Put(CreateTuple(279, value), rsp, targets6)
	log("p26 w:%d", value)
	countwrites(value, 26, 18)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(236, &value), rsp, s26)
	log("p26 r:%d", value)
	countreads(value, 26, 15)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(401, &value), rsp, s26)
	log("p26 r:%d", value)
	countreads(value, 26, 26)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(95, &value), rsp, s26)
	log("p26 r:%d", value)
	countreads(value, 26, 6)
	l.Unlock()

	delay()
	l.Lock()
	value = 27
	Put(CreateTuple(430, value), rsp, targets7)
	log("p26 w:%d", value)
	countwrites(value, 26, 27)
	l.Unlock()

	delay()
	l.Lock()
	value = 17
	Put(CreateTuple(259, value), rsp, targets8)
	log("p26 w:%d", value)
	countwrites(value, 26, 17)
	l.Unlock()

	delay()
	l.Lock()
	value = 22
	Put(CreateTuple(339, value), rsp, targets9)
	log("p26 w:%d", value)
	countwrites(value, 26, 22)
	l.Unlock()

	delay()
	l.Lock()
	value = 21
	Put(CreateTuple(329, value), rsp, targets10)
	log("p26 w:%d", value)
	countwrites(value, 26, 21)
	l.Unlock()

	delay()
	l.Lock()
	value = 21
	Put(CreateTuple(332, value), rsp, targets11)
	log("p26 w:%d", value)
	countwrites(value, 26, 21)
	l.Unlock()
}

func p27() {
	targets6 := make([]string, 0)
	targets5 := make([]string, 0)
	targets4 := make([]string, 0)
	targets3 := make([]string, 0)
	targets2 := make([]string, 0)
	targets1 := make([]string, 1)
	targets1[0] = uri[s14]
	targets0 := make([]string, 0)

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(54, &value), rsp, s27)
	log("p27 r:%d", value)
	countreads(value, 27, 4)
	l.Unlock()

	delay()
	l.Lock()
	value = 25
	Put(CreateTuple(394, value), rsp, targets0)
	log("p27 w:%d", value)
	countwrites(value, 27, 25)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(484, &value), rsp, s27)
	log("p27 r:%d", value)
	countreads(value, 27, 31)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(241, &value), rsp, s27)
	log("p27 r:%d", value)
	countreads(value, 27, 16)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(466, &value), rsp, s27)
	log("p27 r:%d", value)
	countreads(value, 27, 30)
	l.Unlock()

	delay()
	l.Lock()
	value = 2
	Put(CreateTuple(25, value), rsp, targets1)
	log("p27 w:%d", value)
	countwrites(value, 27, 2)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(476, &value), rsp, s27)
	log("p27 r:%d", value)
	countreads(value, 27, 30)
	l.Unlock()

	delay()
	l.Lock()
	value = 20
	Put(CreateTuple(312, value), rsp, targets2)
	log("p27 w:%d", value)
	countwrites(value, 27, 20)
	l.Unlock()

	delay()
	l.Lock()
	value = 23
	Put(CreateTuple(368, value), rsp, targets3)
	log("p27 w:%d", value)
	countwrites(value, 27, 23)
	l.Unlock()

	delay()
	l.Lock()
	value = 31
	Put(CreateTuple(488, value), rsp, targets4)
	log("p27 w:%d", value)
	countwrites(value, 27, 31)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(187, &value), rsp, s27)
	log("p27 r:%d", value)
	countreads(value, 27, 12)
	l.Unlock()

	delay()
	l.Lock()
	value = 3
	Put(CreateTuple(37, value), rsp, targets5)
	log("p27 w:%d", value)
	countwrites(value, 27, 3)
	l.Unlock()

	delay()
	l.Lock()
	value = 26
	Put(CreateTuple(408, value), rsp, targets6)
	log("p27 w:%d", value)
	countwrites(value, 27, 26)
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
	QueryP(CreateTuple(282, &value), rsp, s27)
	log("p27 r:%d", value)
	countreads(value, 27, 18)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(378, &value), rsp, s27)
	log("p27 r:%d", value)
	countreads(value, 27, 24)
	l.Unlock()
}

func p28() {
	targets13 := make([]string, 0)
	targets12 := make([]string, 0)
	targets11 := make([]string, 0)
	targets10 := make([]string, 0)
	targets9 := make([]string, 1)
	targets9[0] = uri[s1]
	targets8 := make([]string, 2)
	targets8[1] = uri[s29]
	targets8[0] = uri[s17]
	targets7 := make([]string, 0)
	targets6 := make([]string, 1)
	targets6[0] = uri[s8]
	targets5 := make([]string, 0)
	targets4 := make([]string, 0)
	targets3 := make([]string, 0)
	targets2 := make([]string, 1)
	targets2[0] = uri[s29]
	targets1 := make([]string, 1)
	targets1[0] = uri[s7]
	targets0 := make([]string, 1)
	targets0[0] = uri[s25]

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 32
	Put(CreateTuple(498, value), rsp, targets0)
	log("p28 w:%d", value)
	countwrites(value, 28, 32)
	l.Unlock()

	delay()
	l.Lock()
	value = 15
	Put(CreateTuple(226, value), rsp, targets1)
	log("p28 w:%d", value)
	countwrites(value, 28, 15)
	l.Unlock()

	delay()
	l.Lock()
	value = 7
	Put(CreateTuple(100, value), rsp, targets2)
	log("p28 w:%d", value)
	countwrites(value, 28, 7)
	l.Unlock()

	delay()
	l.Lock()
	value = 11
	Put(CreateTuple(170, value), rsp, targets3)
	log("p28 w:%d", value)
	countwrites(value, 28, 11)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(380, &value), rsp, s28)
	log("p28 r:%d", value)
	countreads(value, 28, 24)
	l.Unlock()

	delay()
	l.Lock()
	value = 14
	Put(CreateTuple(221, value), rsp, targets4)
	log("p28 w:%d", value)
	countwrites(value, 28, 14)
	l.Unlock()

	delay()
	l.Lock()
	value = 7
	Put(CreateTuple(103, value), rsp, targets5)
	log("p28 w:%d", value)
	countwrites(value, 28, 7)
	l.Unlock()

	delay()
	l.Lock()
	value = 31
	Put(CreateTuple(489, value), rsp, targets6)
	log("p28 w:%d", value)
	countwrites(value, 28, 31)
	l.Unlock()

	delay()
	l.Lock()
	value = 5
	Put(CreateTuple(70, value), rsp, targets7)
	log("p28 w:%d", value)
	countwrites(value, 28, 5)
	l.Unlock()

	delay()
	l.Lock()
	value = 12
	Put(CreateTuple(180, value), rsp, targets8)
	log("p28 w:%d", value)
	countwrites(value, 28, 12)
	l.Unlock()

	delay()
	l.Lock()
	value = 32
	Put(CreateTuple(504, value), rsp, targets9)
	log("p28 w:%d", value)
	countwrites(value, 28, 32)
	l.Unlock()

	delay()
	l.Lock()
	value = 20
	Put(CreateTuple(306, value), rsp, targets10)
	log("p28 w:%d", value)
	countwrites(value, 28, 20)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(263, &value), rsp, s28)
	log("p28 r:%d", value)
	countreads(value, 28, 17)
	l.Unlock()

	delay()
	l.Lock()
	value = 5
	Put(CreateTuple(77, value), rsp, targets11)
	log("p28 w:%d", value)
	countwrites(value, 28, 5)
	l.Unlock()

	delay()
	l.Lock()
	value = 14
	Put(CreateTuple(209, value), rsp, targets12)
	log("p28 w:%d", value)
	countwrites(value, 28, 14)
	l.Unlock()

	delay()
	l.Lock()
	value = 14
	Put(CreateTuple(214, value), rsp, targets13)
	log("p28 w:%d", value)
	countwrites(value, 28, 14)
	l.Unlock()
}

func p29() {
	targets11 := make([]string, 0)
	targets10 := make([]string, 1)
	targets10[0] = uri[s2]
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
	value = -1
	QueryP(CreateTuple(283, &value), rsp, s29)
	log("p29 r:%d", value)
	countreads(value, 29, 18)
	l.Unlock()

	delay()
	l.Lock()
	value = 9
	Put(CreateTuple(141, value), rsp, targets0)
	log("p29 w:%d", value)
	countwrites(value, 29, 9)
	l.Unlock()

	delay()
	l.Lock()
	value = 31
	Put(CreateTuple(492, value), rsp, targets1)
	log("p29 w:%d", value)
	countwrites(value, 29, 31)
	l.Unlock()

	delay()
	l.Lock()
	value = 1
	Put(CreateTuple(1, value), rsp, targets2)
	log("p29 w:%d", value)
	countwrites(value, 29, 1)
	l.Unlock()

	delay()
	l.Lock()
	value = 14
	Put(CreateTuple(222, value), rsp, targets3)
	log("p29 w:%d", value)
	countwrites(value, 29, 14)
	l.Unlock()

	delay()
	l.Lock()
	value = 9
	Put(CreateTuple(144, value), rsp, targets4)
	log("p29 w:%d", value)
	countwrites(value, 29, 9)
	l.Unlock()

	delay()
	l.Lock()
	value = 1
	Put(CreateTuple(1, value), rsp, targets5)
	log("p29 w:%d", value)
	countwrites(value, 29, 1)
	l.Unlock()

	delay()
	l.Lock()
	value = 6
	Put(CreateTuple(93, value), rsp, targets6)
	log("p29 w:%d", value)
	countwrites(value, 29, 6)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(36, &value), rsp, s29)
	log("p29 r:%d", value)
	countreads(value, 29, 3)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(180, &value), rsp, s29)
	log("p29 r:%d", value)
	countreads(value, 29, 12)
	l.Unlock()

	delay()
	l.Lock()
	value = 5
	Put(CreateTuple(70, value), rsp, targets7)
	log("p29 w:%d", value)
	countwrites(value, 29, 5)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(100, &value), rsp, s29)
	log("p29 r:%d", value)
	countreads(value, 29, 7)
	l.Unlock()

	delay()
	l.Lock()
	value = 6
	Put(CreateTuple(94, value), rsp, targets8)
	log("p29 w:%d", value)
	countwrites(value, 29, 6)
	l.Unlock()

	delay()
	l.Lock()
	value = 26
	Put(CreateTuple(416, value), rsp, targets9)
	log("p29 w:%d", value)
	countwrites(value, 29, 26)
	l.Unlock()

	delay()
	l.Lock()
	value = 14
	Put(CreateTuple(215, value), rsp, targets10)
	log("p29 w:%d", value)
	countwrites(value, 29, 14)
	l.Unlock()

	delay()
	l.Lock()
	value = 17
	Put(CreateTuple(270, value), rsp, targets11)
	log("p29 w:%d", value)
	countwrites(value, 29, 17)
	l.Unlock()
}

func p30() {
	targets11 := make([]string, 2)
	targets11[1] = uri[s19]
	targets11[0] = uri[s9]
	targets10 := make([]string, 0)
	targets9 := make([]string, 0)
	targets8 := make([]string, 0)
	targets7 := make([]string, 0)
	targets6 := make([]string, 0)
	targets5 := make([]string, 1)
	targets5[0] = uri[s2]
	targets4 := make([]string, 0)
	targets3 := make([]string, 0)
	targets2 := make([]string, 0)
	targets1 := make([]string, 0)
	targets0 := make([]string, 0)

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 31
	Put(CreateTuple(494, value), rsp, targets0)
	log("p30 w:%d", value)
	countwrites(value, 30, 31)
	l.Unlock()

	delay()
	l.Lock()
	value = 32
	Put(CreateTuple(512, value), rsp, targets1)
	log("p30 w:%d", value)
	countwrites(value, 30, 32)
	l.Unlock()

	delay()
	l.Lock()
	value = 13
	Put(CreateTuple(200, value), rsp, targets2)
	log("p30 w:%d", value)
	countwrites(value, 30, 13)
	l.Unlock()

	delay()
	l.Lock()
	value = 25
	Put(CreateTuple(396, value), rsp, targets3)
	log("p30 w:%d", value)
	countwrites(value, 30, 25)
	l.Unlock()

	delay()
	l.Lock()
	value = 20
	Put(CreateTuple(317, value), rsp, targets4)
	log("p30 w:%d", value)
	countwrites(value, 30, 20)
	l.Unlock()

	delay()
	l.Lock()
	value = 14
	Put(CreateTuple(215, value), rsp, targets5)
	log("p30 w:%d", value)
	countwrites(value, 30, 14)
	l.Unlock()

	delay()
	l.Lock()
	value = 5
	Put(CreateTuple(77, value), rsp, targets6)
	log("p30 w:%d", value)
	countwrites(value, 30, 5)
	l.Unlock()

	delay()
	l.Lock()
	value = 19
	Put(CreateTuple(302, value), rsp, targets7)
	log("p30 w:%d", value)
	countwrites(value, 30, 19)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(160, &value), rsp, s30)
	log("p30 r:%d", value)
	countreads(value, 30, 10)
	l.Unlock()

	delay()
	l.Lock()
	value = 26
	Put(CreateTuple(408, value), rsp, targets8)
	log("p30 w:%d", value)
	countwrites(value, 30, 26)
	l.Unlock()

	delay()
	l.Lock()
	value = 27
	Put(CreateTuple(431, value), rsp, targets9)
	log("p30 w:%d", value)
	countwrites(value, 30, 27)
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
	value = 17
	Put(CreateTuple(270, value), rsp, targets10)
	log("p30 w:%d", value)
	countwrites(value, 30, 17)
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
	QueryP(CreateTuple(398, &value), rsp, s30)
	log("p30 r:%d", value)
	countreads(value, 30, 25)
	l.Unlock()

	delay()
	l.Lock()
	value = 14
	Put(CreateTuple(213, value), rsp, targets11)
	log("p30 w:%d", value)
	countwrites(value, 30, 14)
	l.Unlock()
}

func p31() {
	targets12 := make([]string, 0)
	targets11 := make([]string, 1)
	targets11[0] = uri[s14]
	targets10 := make([]string, 1)
	targets10[0] = uri[s22]
	targets9 := make([]string, 1)
	targets9[0] = uri[s31]
	targets8 := make([]string, 0)
	targets7 := make([]string, 1)
	targets7[0] = uri[s16]
	targets6 := make([]string, 1)
	targets6[0] = uri[s2]
	targets5 := make([]string, 0)
	targets4 := make([]string, 0)
	targets3 := make([]string, 0)
	targets2 := make([]string, 0)
	targets1 := make([]string, 1)
	targets1[0] = uri[s7]
	targets0 := make([]string, 0)

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 21
	Put(CreateTuple(322, value), rsp, targets0)
	log("p31 w:%d", value)
	countwrites(value, 31, 21)
	l.Unlock()

	delay()
	l.Lock()
	value = 8
	Put(CreateTuple(128, value), rsp, targets1)
	log("p31 w:%d", value)
	countwrites(value, 31, 8)
	l.Unlock()

	delay()
	l.Lock()
	value = 27
	Put(CreateTuple(430, value), rsp, targets2)
	log("p31 w:%d", value)
	countwrites(value, 31, 27)
	l.Unlock()

	delay()
	l.Lock()
	value = 7
	Put(CreateTuple(106, value), rsp, targets3)
	log("p31 w:%d", value)
	countwrites(value, 31, 7)
	l.Unlock()

	delay()
	l.Lock()
	value = 32
	Put(CreateTuple(505, value), rsp, targets4)
	log("p31 w:%d", value)
	countwrites(value, 31, 32)
	l.Unlock()

	delay()
	l.Lock()
	value = 32
	Put(CreateTuple(501, value), rsp, targets5)
	log("p31 w:%d", value)
	countwrites(value, 31, 32)
	l.Unlock()

	delay()
	l.Lock()
	value = 26
	Put(CreateTuple(405, value), rsp, targets6)
	log("p31 w:%d", value)
	countwrites(value, 31, 26)
	l.Unlock()

	delay()
	l.Lock()
	value = 5
	Put(CreateTuple(69, value), rsp, targets7)
	log("p31 w:%d", value)
	countwrites(value, 31, 5)
	l.Unlock()

	delay()
	l.Lock()
	value = 2
	Put(CreateTuple(26, value), rsp, targets8)
	log("p31 w:%d", value)
	countwrites(value, 31, 2)
	l.Unlock()

	delay()
	l.Lock()
	value = 2
	Put(CreateTuple(23, value), rsp, targets9)
	log("p31 w:%d", value)
	countwrites(value, 31, 2)
	l.Unlock()

	delay()
	l.Lock()
	value = 13
	Put(CreateTuple(203, value), rsp, targets10)
	log("p31 w:%d", value)
	countwrites(value, 31, 13)
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
	value = 2
	Put(CreateTuple(25, value), rsp, targets11)
	log("p31 w:%d", value)
	countwrites(value, 31, 2)
	l.Unlock()

	delay()
	l.Lock()
	value = 1
	Put(CreateTuple(11, value), rsp, targets12)
	log("p31 w:%d", value)
	countwrites(value, 31, 1)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(474, &value), rsp, s31)
	log("p31 r:%d", value)
	countreads(value, 31, 30)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(23, &value), rsp, s31)
	log("p31 r:%d", value)
	countreads(value, 31, 2)
	l.Unlock()
}

func p32() {
	targets14 := make([]string, 0)
	targets13 := make([]string, 1)
	targets13[0] = uri[s18]
	targets12 := make([]string, 0)
	targets11 := make([]string, 0)
	targets10 := make([]string, 0)
	targets9 := make([]string, 1)
	targets9[0] = uri[s18]
	targets8 := make([]string, 0)
	targets7 := make([]string, 0)
	targets6 := make([]string, 0)
	targets5 := make([]string, 1)
	targets5[0] = uri[s22]
	targets4 := make([]string, 0)
	targets3 := make([]string, 0)
	targets2 := make([]string, 0)
	targets1 := make([]string, 1)
	targets1[0] = uri[s6]
	targets0 := make([]string, 0)

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 28
	Put(CreateTuple(440, value), rsp, targets0)
	log("p32 w:%d", value)
	countwrites(value, 32, 28)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(471, &value), rsp, s32)
	log("p32 r:%d", value)
	countreads(value, 32, 30)
	l.Unlock()

	delay()
	l.Lock()
	value = 27
	Put(CreateTuple(429, value), rsp, targets1)
	log("p32 w:%d", value)
	countwrites(value, 32, 27)
	l.Unlock()

	delay()
	l.Lock()
	value = 28
	Put(CreateTuple(433, value), rsp, targets2)
	log("p32 w:%d", value)
	countwrites(value, 32, 28)
	l.Unlock()

	delay()
	l.Lock()
	value = 19
	Put(CreateTuple(295, value), rsp, targets3)
	log("p32 w:%d", value)
	countwrites(value, 32, 19)
	l.Unlock()

	delay()
	l.Lock()
	value = 23
	Put(CreateTuple(358, value), rsp, targets4)
	log("p32 w:%d", value)
	countwrites(value, 32, 23)
	l.Unlock()

	delay()
	l.Lock()
	value = 20
	Put(CreateTuple(308, value), rsp, targets5)
	log("p32 w:%d", value)
	countwrites(value, 32, 20)
	l.Unlock()

	delay()
	l.Lock()
	value = 6
	Put(CreateTuple(84, value), rsp, targets6)
	log("p32 w:%d", value)
	countwrites(value, 32, 6)
	l.Unlock()

	delay()
	l.Lock()
	value = 1
	Put(CreateTuple(1, value), rsp, targets7)
	log("p32 w:%d", value)
	countwrites(value, 32, 1)
	l.Unlock()

	delay()
	l.Lock()
	value = 5
	Put(CreateTuple(73, value), rsp, targets8)
	log("p32 w:%d", value)
	countwrites(value, 32, 5)
	l.Unlock()

	delay()
	l.Lock()
	value = 17
	Put(CreateTuple(261, value), rsp, targets9)
	log("p32 w:%d", value)
	countwrites(value, 32, 17)
	l.Unlock()

	delay()
	l.Lock()
	value = 16
	Put(CreateTuple(254, value), rsp, targets10)
	log("p32 w:%d", value)
	countwrites(value, 32, 16)
	l.Unlock()

	delay()
	l.Lock()
	value = 15
	Put(CreateTuple(225, value), rsp, targets11)
	log("p32 w:%d", value)
	countwrites(value, 32, 15)
	l.Unlock()

	delay()
	l.Lock()
	value = 16
	Put(CreateTuple(249, value), rsp, targets12)
	log("p32 w:%d", value)
	countwrites(value, 32, 16)
	l.Unlock()

	delay()
	l.Lock()
	value = 12
	Put(CreateTuple(188, value), rsp, targets13)
	log("p32 w:%d", value)
	countwrites(value, 32, 12)
	l.Unlock()

	delay()
	l.Lock()
	value = 3
	Put(CreateTuple(37, value), rsp, targets14)
	log("p32 w:%d", value)
	countwrites(value, 32, 3)
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
