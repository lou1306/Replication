package main

import (
    "fmt"
    "os"
    "sync"
    "time"
    "math"
    "math/rand"
    . "github.com/pspaces/gospace"
)

// Number of requests that will be generated
const REQUESTS = 500
// Number of servers
const SERVERS = 3

// Load of an individual request
const REQ_LOAD = 1
// Server declares to be "busy" when its load hits this threshold
const BUSY_THRESHOLD = 2
// Server stops queueing requests when its load hits this threshold
const UPPER_THRESHOLD = 4

const debug = true

// Probability with which a server accepts a request
// (Probability of completing a request is 1-P_ACCEPT)
const P_ACCEPT = 0.80

var l sync.Mutex
var wg sync.WaitGroup
var s1 Space
var s2 Space
var s3 Space
var s4 Space


// Counters for stats
var FORWARDED_REQUESTS int = 0
var STOLEN_REQUESTS int = 0
var COMPLETED_REQUESTS int = 0
// Logical timestamp to measure request completion time
var TIME int64 = 0


func main() {
    // n servers + 1 request generator
    wg.Add(SERVERS + 1)

    s1 = NewSpace("tcp://localhost:34117/s1")
    s2 = NewSpace("tcp://localhost:34118/s2")
    s3 = NewSpace("tcp://localhost:34119/s3")
    s4 = NewSpace("tcp://localhost:34120/s4")

    fmt.Println("Start")
    go P1_Generator()
    go P2()
    go P3()
    go P4()

    wg.Wait()
    fmt.Println("Stop")
    fmt.Fprintf(os.Stdout, "%d requests completed (out of %d).\n", COMPLETED_REQUESTS, REQUESTS)
    fmt.Fprintf(os.Stdout, "%d requests were local.\n", COMPLETED_REQUESTS-STOLEN_REQUESTS)
    fmt.Fprintf(os.Stdout, "%d requests forwarded.\n", FORWARDED_REQUESTS)
    fmt.Fprintf(os.Stdout, "%d requests stolen.\n", STOLEN_REQUESTS)
    fmt.Fprintf(os.Stdout, "Time: min %d steps\tmax %d steps\tavg %f steps.\n", minTime, maxTime, float64(totalTime)/float64(COMPLETED_REQUESTS))

}


func P1_Generator() {
    // This process only generates service requests.
    // A request is a pair (srv, svc) where
    // srv is the ID of a server (1 < srv <= SERVERS+1);
    // svc is the ID of the required service (1 < svc <= SERVERS+1).
    // Notice that server i is the default target for the i-th service.
    defer wg.Done()
    var i int;
    var service int;
    var server int;
    var tstamp int64;

    for i = 0; i<REQUESTS; i++ {
    	// Both service and server must be between 2 and SERVERS+2
        service = rand.Intn(SERVERS) + 2
        server = rand.Intn(SERVERS) + 2
        tstamp = getTIME()
        l.Lock()
        if server == 2 {
            s2.Put(2, service, tstamp);
        }
        if server == 3 {
            s3.Put(3, service, tstamp);
        }
        if server == 4 {
            s4.Put(4, service, tstamp);
        }
        l.Unlock()
        log("[1]\tput req %d to server %d", service, server)
    }
}

func P2() {
    defer wg.Done()
    var tid int = 2
    var load int = 0
    var service int
    var wasBusyBefore bool
    var isBusyNow bool
    var tstamp int64
    var server int

    steps := 0
    // The (steps < 10*REQUESTS) condition is just in case we
    // drop requests in the replicated case
    for COMPLETED_REQUESTS < REQUESTS && (steps < 10*REQUESTS) {
        steps += 1
        l.Lock()
        service = -1
        wasBusyBefore = isBusyNow

        if load < UPPER_THRESHOLD && biasedRandBool(P_ACCEPT) {
            s2.GetP(&server, &service, &tstamp)
            if service == tid {
                updateTimeStats(getTIME() - tstamp)
                load += REQ_LOAD
                if server != tid {
                	// We have stolen a request for OUR service which was sent
                	// to ANOTHER server. Only happens in the replicated case!
                	STOLEN_REQUESTS += 1
                    log("[%d]\tstole req %d from server %d", tid, service, server)
                } else {
                	log("[%d]\taccepted req for service %d", tid, service)
                }
            } else {
                // This request is for a service j != tid.
                // However, if server j declares to be busy,
                // steal the request.
                // (Notice that it doesnt't matter whether the "thief" is busy or not.)
                // If server j is not busy, simply forward the request to it.
                var s string
                stolenReq := false
                if service == 2 {
                    s2.QueryP(&s)
                    if s == "busy" {
                        //updateTimeStats(time.Now().UnixMicro() - tstamp)
                        updateTimeStats(getTIME() - tstamp)
                        stolenReq = true
                        load += REQ_LOAD
                    } else {
                        s2.Put(2, 2, tstamp)
                    }
                }
                if service == 3 {
                    s3.QueryP(&s)
                    if s == "busy" {
                        //updateTimeStats(time.Now().UnixMicro() - tstamp)
                        updateTimeStats(getTIME() - tstamp)
                        stolenReq = true
                        load += REQ_LOAD
                    } else {
                        s3.Put(3, 3, tstamp)
                    }
                }
                if service == 4 {
                    s4.QueryP(&s);
                    if s == "busy" {
                        //updateTimeStats(time.Now().UnixMicro() - tstamp)
                        updateTimeStats(getTIME() - tstamp)
                        stolenReq = true
                        load += REQ_LOAD
                    } else {
                        s4.Put(4, 4, tstamp)
                    }
                }

                isBusyNow = load >= BUSY_THRESHOLD
                if isBusyNow && !wasBusyBefore {
                    // Declare that this server is now busy
                    s2.Put("busy")
                    log("[%d]\tnow busy: load %d >= %d",  tid, load, BUSY_THRESHOLD)
                }
                // Update counters
                if service != -1 {
                    if stolenReq {
                        STOLEN_REQUESTS += 1
                        log("[%d]\tstole req %d from server %d", tid, service, server)
                    } else {
                        FORWARDED_REQUESTS += 1
                        log("[%d]\tfwded req %d to server %d", tid, service, service)
                    }
                }
            }
        } else {
            if load >= REQ_LOAD {
                // Complete a request
                delay()
                log("[%d]\tcompleting a request", tid)
                load -= REQ_LOAD
                if load < BUSY_THRESHOLD && (load+REQ_LOAD >= BUSY_THRESHOLD) {
                    // Declare that this server is no longer busy
                    isBusyNow = false
                    s2.GetP("busy")
                    log("[%d]\tno longer busy: load %d < %d",  tid, load, BUSY_THRESHOLD)
                }
                COMPLETED_REQUESTS += 1
            }
        }
        l.Unlock()
    }
}

func P3() {
    defer wg.Done()
    var tid int = 3
    var load int = 0
    var server int
    var service int
    var wasBusyBefore bool
    var isBusyNow bool
    var tstamp int64

    steps := 0
    // The (steps < 10*REQUESTS) condition is just in case we
    // drop requests in the replicated case
    for COMPLETED_REQUESTS < REQUESTS && (steps < 10*REQUESTS) {
        steps += 1
        l.Lock()
        service = -1
        wasBusyBefore = isBusyNow
        stolenReq := false

        if load < UPPER_THRESHOLD && biasedRandBool(P_ACCEPT) {
            s3.GetP(&server, &service, &tstamp)
            if service == tid {
                updateTimeStats(getTIME() - tstamp)
                load += REQ_LOAD
                if server != tid {
                	STOLEN_REQUESTS += 1
                    log("[%d]\tstole req %d from server %d", tid, service, server)
                } else {
                	log("[%d]\taccepted req for service %d", tid, service)
                }
            } else {
                var s string
                // Steal the job if target's busy
                if service == 2 {
                    s2.QueryP(&s)
                    if s == "busy" {
                        //updateTimeStats(time.Now().UnixMicro() - tstamp)
                        updateTimeStats(getTIME() - tstamp)
                        stolenReq = true
                        load += REQ_LOAD
                    } else {
                        s2.Put(2, 2, tstamp)
                    }
                }
                if service == 3 {
                    s3.QueryP(&s)
                    if s == "busy" {
                        //updateTimeStats(time.Now().UnixMicro() - tstamp)
                        updateTimeStats(getTIME() - tstamp)
                        stolenReq = true
                        load += REQ_LOAD
                    } else {
                        s3.Put(3, 3, tstamp)
                    }
                }
                if service == 4 {
                    s4.QueryP(&s);
                    if s == "busy" {
                        //updateTimeStats(time.Now().UnixMicro() - tstamp)
                        updateTimeStats(getTIME() - tstamp)
                        stolenReq = true
                        load += REQ_LOAD
                    } else {
                        s4.Put(4, 4, tstamp)
                    }
                }

                isBusyNow = load >= BUSY_THRESHOLD
                if isBusyNow && !wasBusyBefore {
                    s3.Put("busy")
                    log("[%d]\tnow busy: load %d >= %d",  tid, load, BUSY_THRESHOLD)
                }
                if service != -1 {
                    if stolenReq {
                        STOLEN_REQUESTS += 1
                        log("[%d]\tstole req %d from server %d", tid, service, server)
                    } else {
                        FORWARDED_REQUESTS += 1
                        log("[%d]\tfwded req %d to server %d", tid, service, service)
                    }
                }
            }
        } else {
            if load >= REQ_LOAD {
                // Complete a request
                delay()
                log("[%d]\tcompleting a request", tid)
                load -= REQ_LOAD
                if load < BUSY_THRESHOLD && (load+REQ_LOAD >= BUSY_THRESHOLD) {
                    // Declare that this server is no longer busy
                    isBusyNow = false
                    s3.GetP("busy")
                    log("[%d]\tno longer busy: load %d < %d",  tid, load, BUSY_THRESHOLD)
                }
                COMPLETED_REQUESTS += 1
            }
        }
        l.Unlock()
    }
}

func P4() {
    defer wg.Done()
    var tid int = 4
    var load int = 0
    var server int
    var service int
    var wasBusyBefore bool
    var isBusyNow bool
    var tstamp int64

    steps := 0
    // The (steps < 10*REQUESTS) condition is just in case we
    // drop requests in the replicated case
    for COMPLETED_REQUESTS < REQUESTS && (steps < 10*REQUESTS) {
        steps += 1
        l.Lock()
        service = -1
        wasBusyBefore = isBusyNow

        if load < UPPER_THRESHOLD && biasedRandBool(P_ACCEPT) {
            s4.GetP(&server, &service, &tstamp)
            if service == tid {
                updateTimeStats(getTIME() - tstamp)
                load += REQ_LOAD
                if server != tid {
                	STOLEN_REQUESTS += 1
                    log("[%d]\tstole req %d from server %d", tid, service, server)
                } else {
                	log("[%d]\taccepted req for service %d", tid, service)
                }
            } else {
                var s string
                stolenReq := false
                // Steal the job if target's busy
                if service == 2 {
                    s2.QueryP(&s)
                    if s == "busy" {
                        //updateTimeStats(time.Now().UnixMicro() - tstamp)
                        updateTimeStats(getTIME() - tstamp)
                        stolenReq = true
                        load += REQ_LOAD
                    } else {
                        s2.Put(2, 2, tstamp)
                    }
                }
                if service == 3 {
                    s3.QueryP(&s)
                    if s == "busy" {
                        //updateTimeStats(time.Now().UnixMicro() - tstamp)
                        updateTimeStats(getTIME() - tstamp)
                        stolenReq = true
                        load += REQ_LOAD
                    } else {
                        s3.Put(3, 3, tstamp)
                    }
                }
                if service == 4 {
                    s4.QueryP(&s);
                    if s == "busy" {
                        //updateTimeStats(time.Now().UnixMicro() - tstamp)
                        updateTimeStats(getTIME() - tstamp)
                        stolenReq = true
                        load += REQ_LOAD
                    } else {
                        s4.Put(4, 4, tstamp)
                    }
                }

                if isBusyNow && !wasBusyBefore {
                    s4.Put("busy")
                    log("[%d]\tnow busy: load %d >= %d",  tid, load, BUSY_THRESHOLD)
                }
                if service != -1 {
                    if stolenReq {
                        STOLEN_REQUESTS += 1
                        log("[%d]\tstole req %d from server %d", tid, service, server)
                    } else {
                        FORWARDED_REQUESTS += 1
                        log("[%d]\tfwded req %d to server %d", tid, service, service)
                    }
                }
            }
        } else {
            if load >= REQ_LOAD {
                // Complete a request
                delay()
                log("[%d]\tcompleting a request", tid)
                load -= REQ_LOAD
                if load < BUSY_THRESHOLD && (load+REQ_LOAD >= BUSY_THRESHOLD) {
                    // Declare that this server is no longer busy
                    isBusyNow = false
                    s3.GetP("busy")
                    log("[%d]\tno longer busy: load %d < %d",  tid, load, BUSY_THRESHOLD)
                }
                COMPLETED_REQUESTS += 1
            }
        }
        l.Unlock()
    }
}

func delay() {
    time.Sleep(time.Duration(rand.Int63n(75))*time.Millisecond)
}

func biasedRandBool(bias float32) bool {
    // Assuming that 0 <= bias <= 1
    // Returns true with probability "bias" and
    // 0 with probability (1-bias).
    return rand.Float32() < bias
}

func log(format string, a ...interface{}) {
    if debug {
        fmt.Fprintf(os.Stdout, format+"\n", a ...)
    }
}

var totalTime int64 = 0
var minTime int64 = math.MaxInt64
var maxTime int64 = 0

func updateTimeStats (newTime int64) {
    totalTime += newTime
    if (newTime > maxTime) {
        maxTime = newTime
    }
    if (newTime < minTime) {
        minTime = newTime
    }
}

func getTIME() int64 {
	result := TIME
	TIME += 1
	return result
}

