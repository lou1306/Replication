/*

CHANGES:
	2020.06.04  bugfix: GetP() was not correctly removing all replicas

*/

package repligospaces

// Importing the useful packages
import (
	// "fmt"
	"sync"
	"time"
	. "github.com/pspaces/gospace"
)

type Replispace struct {
	mux sync.Mutex
	Sp  map[string]*Space
	// Maximum amount of replicas each space should hold.
	// A non-positive value *disables* replica eviction
	ReplLimit int
	// The eviction policy applied by the RepliSpace whenever a space hits
	// the replica limit. Recognized values are "lru" and "random", and "fifo"; 
	// other value defaults to "fifo".
	ReplacementPolicy string
}

// Tracks the number of replicas in each space
var replicaCounter map[string]int = make(map[string]int)
// Maps a pair (space, tuple) to its creation time
var createTime map[Space]map[string]TimeRecord = make(map[Space]map[string]TimeRecord)
// Maps a pair (space, tuple) to its last access time
// Notice that creation time is the same across all spaces (at least for now),
// but last access time may vary
var lastAccessTime map[Space]map[string]TimeRecord = make(map[Space]map[string]TimeRecord)

type TimeRecord struct {
    // A tuple 
    tuple Tuple
    // Its timestamp 
    time time.Time
    // Spaces the tuple was replicated to
    // spaces []string
}


var wcnt int
var evictionCount int

func dataFieldsOf(t Tuple) ([]interface{}) {
	return t.Fields[:len(t.Fields)-2]
}

func contains (strings []string, s string) bool {
	for _, x := range strings {
		if x == s {
			return true
		}
	}
	return false
}

func MinByTime(mp map[Space]map[string]TimeRecord, s Space) Tuple {
	var result Tuple
	// Use an absurdly high value at the start
	currentMinValue := time.Now().Add(time.Hour * (1<<10))
	for _, element := range mp[s] {
		if element.time.Before(currentMinValue) {
			result = element.tuple
			currentMinValue = element.time
		}
	}
	return CreateTuple(dataFieldsOf(result)...)
}


func evict(Sp Replispace, s string, t Tuple) {
	// fmt.Println(">>>evict", s, t.String())

	var createTime int64
	var y []string // <--- extra field to match the space list S
	var data []interface{}
	data = append(data, t.Fields...)
	data = append(data, &createTime)
	data = append(data, &y)
	var p1 Tuple = CreateTuple(data...)
	Sp.Sp[s].QueryP(p1.Fields...)
	unsafeGetP(t, Sp, *Sp.Sp[s])
	evictionCount += len(y)

}

func EvictFIFO(Sp Replispace, s string) {
	// Evicts the oldest (lowest creation time) tuple from s
	t := MinByTime(createTime, *Sp.Sp[s])
	evict(Sp, s, t)
}

func EvictLRU(Sp Replispace, s string) {
	// Evicts the oldest (lowest creation time) tuple from s
	t := MinByTime(lastAccessTime, *Sp.Sp[s])
	evict(Sp, s, t)
}

func EvictRandom(Sp Replispace, s string) {
	// Evicts one random tuple
	var t Tuple

	// Notice that, in Go, iteration over a map happens in
	// unspecified order, and since Go v1 the order is guaranteed
	// to be randomized.
	for _, element := range createTime[*Sp.Sp[s]] {
		t = element.tuple
		break
	}
	evict(Sp, s, CreateTuple(dataFieldsOf(t)...))
}


//~~~~~~~~
// Add a tuple to the set S of spaces identifiers
func Put(t Tuple, Sp Replispace, S []string) Tuple {
	// fmt.Println(">>>Put", t.String())

	Sp.mux.Lock()

	// create tuple t' = {t,S}
	var data []interface{}
	data = append(data, t.Fields...)
	data = append(data, time.Now().UnixNano())
	data = append(data, S)
	var t1 Tuple = CreateTuple(data...)

	now := TimeRecord{tuple:t1, time:time.Now()}


	// add t' to each space in S
	for i := 0; i < len(S); i++ {
		if Sp.ReplLimit > 0 && (replicaCounter[S[i]]+1) > Sp.ReplLimit {
			switch Sp.ReplacementPolicy {
			case "lru":
				EvictLRU(Sp, S[i])
			case "random":
				EvictRandom(Sp, S[i])
			case "fifo":
				fallthrough
			default:
				EvictFIFO(Sp, S[i])
			}
		}
		Sp.Sp[S[i]].Put(t1.Fields...)
		wcnt+=1
		replicaCounter[S[i]] += 1
		updateReplicaMax()
		// fmt.Println(">>>",
		// 	"S[", i, "]:", replicaCounter[S[i]], 
		// 	"tot: ", GetReplicaCount(),
		// 	"max: ", GetReplicaMax())
		if createTime[*Sp.Sp[S[i]]] == nil {
			createTime[*Sp.Sp[S[i]]] = make(map[string]TimeRecord)
		}
		if lastAccessTime[*Sp.Sp[S[i]]] == nil {
			lastAccessTime[*Sp.Sp[S[i]]] = make(map[string]TimeRecord)
		}

		createTime[*Sp.Sp[S[i]]][t1.String()] = now
		lastAccessTime[*Sp.Sp[S[i]]][t1.String()] = now
	}

	Sp.mux.Unlock()
	return CreateTuple(t1)
}

//~~~~~~~~ NOT TESTED
// Query a specific space for tuples matching the given pattern
// Blocks if the tuple is not found
/*
func Query(p Tuple, Sp Replispace, s Space) Tuple {
	Sp.mux.Lock()

	// create template p' = {t,S}
	var y []string // <--- extra field to match the space list S
	var data []interface{}
	data = append(data, p.Fields...)
	data = append(data, &y)
	var p1 Tuple = CreateTuple(data...)

	// query a tuple via a pattern matching from a specific space
	t1, e := s.Query(p1.Fields...)

	if e == nil {
		// no error: return the matching tuple without the last field
		var t2 = CreateTuple(t1.Fields[:len(t1.Fields)-1]...)
		Sp.mux.Unlock()
		return t2
	}
	// blocks until the tuple is found
	Sp.mux.Unlock()
	return CreateTuple()
}
*/

//~~~~~~~~
// Query a specific space for tuples matching the given pattern
// Non-blocking version
func QueryP(p Tuple, Sp Replispace, s Space) Tuple {
	Sp.mux.Lock()

	// create template p' = {t,S}
	var createTime int64 // <--- extra field to match the creation time
	var y []string // <--- extra field to match the space list S
	var data []interface{}
	data = append(data, p.Fields...)
	data = append(data, &createTime)
	data = append(data, &y)
	var p1 Tuple = CreateTuple(data...)

	// query a tuple via a pattern matching from a specific space
	t1, e := s.QueryP(p1.Fields...)

	if e == nil {
		// no error: return the matching tuple without the last field
		var u = CreateTuple(dataFieldsOf(t1)...)
		lastAccessTime[s][t1.String()] = TimeRecord{tuple:t1, time:time.Now()}

		Sp.mux.Unlock()
		return u
	}

	Sp.mux.Unlock()
	return CreateTuple()   // returns an empty tuple when no tuple is available
}


func unsafeGetP(p Tuple, Sp Replispace, s Space) Tuple {
	// create template p' = {t,S}
	var y []string // <--- extra field to match the space list S
	var createTimestamp int64 // <--- extra field to match the creation time
	var data []interface{}
	data = append(data, p.Fields...)
	data = append(data, &createTimestamp)
	data = append(data, &y)
	var p1 Tuple = CreateTuple(data...)

	//  search the tuple from space s
	t1, e := s.QueryP(p1.Fields...)

	if e == nil {
		// extract the list of all spaces
		var S = (t1.Fields[len(t1.Fields)-1])

		// transform the interface type of spaces into the string type
		var v []string
		v = S.([]string)

		// for each space in the set S of space identifiers
		for s, space := range v {
			// remove the tuple from the relevant spaces
			delete(createTime[*Sp.Sp[space]], t1.String())
			delete(lastAccessTime[*Sp.Sp[space]], t1.String())
			u, e1 := Sp.Sp[v[s]].GetP(p1.Fields...)

			if e1 == nil {
				replicaCounter[v[s]] -= 1
				// fmt.Println(">>>",
				// 		v[s], ":", replicaCounter[v[s]], 
				// 		"tot: ", GetReplicaCount(),
				// 		"max: ", GetReplicaMax())
				
				if s == len(v)-1 {
					// no error: tuple successfully removed from the space
					///////fmt.Println("  ...Tuple removed from:", v[s])
					u = CreateTuple(dataFieldsOf(u)...)
					return u
				}
			}
		}
	}
	return CreateTuple() // returns an empty tuple when no tuple is available
}

//~~~~~~~~
// Remove a tuple from space s, and
// from any other space where it might have been replicated.
//
// Note that if no matching tuples are found in s, no tuple
// is removed from any other space.
func GetP(p Tuple, Sp Replispace, s Space) Tuple {
	Sp.mux.Lock()
	result := unsafeGetP(p, Sp, s)
	Sp.mux.Unlock()
	return result
}

func Getwcount() int {
	return wcnt
}

func GetEvictionCount() int {
	return evictionCount
}

func GetReplicaCount() int {
	result := 0
	for _, v := range replicaCounter {
		result += v
	}
	return result
}

var replicaMax int = 0

func updateReplicaMax() {
	cnt := GetReplicaCount()
	if cnt > replicaMax {
		replicaMax = cnt
	}
}
func GetReplicaMax() int {
	return replicaMax
}
