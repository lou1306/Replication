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
var replicaCounter map[Space]int = make(map[Space]int)
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
	return t.Fields[:len(t.Fields)-1]
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

/// Management of memory-limited tuple spaces

func evictTuple(Sp Replispace, s Space) {
	var t Tuple
	switch Sp.ReplacementPolicy {
	case "lru":
		t = MinByTime(lastAccessTime, s)
	case "random":
		for _, element := range createTime[s] {
			t = element.tuple
			t = CreateTuple(dataFieldsOf(t)...)
			break
		}
	case "fifo":
		fallthrough
	default:
		t = MinByTime(createTime, s)
	}

	var y []string // <--- extra field to match the space list S
	var data []interface{}
	data = append(data, t.Fields...)
	// var createTime int64
	// data = append(data, &createTime)
	data = append(data, &y)
	var p1 Tuple = CreateTuple(data...)
	s.QueryP(p1.Fields...)
	unsafeGetP(t, Sp, s)
	evictionCount += len(y)
}

func afterPut(t1 Tuple, Sp Replispace, s Space, now TimeRecord) {
	replicaCounter[s] += 1
	switch Sp.ReplacementPolicy {
	case "lru":
		if lastAccessTime[s] == nil {
			lastAccessTime[s] = make(map[string]TimeRecord)
		}
		lastAccessTime[s][t1.String()] = now
	case "random":
		fallthrough
	case "fifo":
		fallthrough
	default:
		if createTime[s] == nil {
			createTime[s] = make(map[string]TimeRecord)
		}
		createTime[s][t1.String()] = now
	}
}

func afterGet(t1 Tuple, Sp Replispace, s Space) {
	replicaCounter[s] -= 1
	switch Sp.ReplacementPolicy {
	case "lru":
		delete(lastAccessTime[s], t1.String())
	case "random":
		fallthrough
	case "fifo":
		fallthrough
	default:
		delete(createTime[s], t1.String())
	}
}

func afterQuery(t1 Tuple, Sp Replispace, s Space) {
	switch Sp.ReplacementPolicy {
	case "lru":
		lastAccessTime[s][t1.String()] = TimeRecord{tuple:t1, time:time.Now()}
	case "random":
		// Do nothing
	case "fifo":
		fallthrough
	default:
		// Do nothing
	}
}
//~~~~~~~~


//~~~~~~~~
// Add a tuple to the set S of spaces identifiers
func Put(t Tuple, Sp Replispace, S []string) Tuple {
	// fmt.Println(">>>Put", t.String())

	Sp.mux.Lock()

	// create tuple t' = {t,S}
	var data []interface{}
	data = append(data, t.Fields...)
	// data = append(data, time.Now().UnixNano())
	data = append(data, S)
	var t1 Tuple = CreateTuple(data...)

	now := TimeRecord{tuple:t1, time:time.Now()}

	// add t' to each space in S
	for i := 0; i < len(S); i++ {
		space := Sp.Sp[S[i]]
		if Sp.ReplLimit > 0 && (replicaCounter[*space]) == Sp.ReplLimit {
			evictTuple(Sp, *space)
		}
		space.Put(t1.Fields...)
		afterPut(t1, Sp, *space, now)
		wcnt+=1
		updateReplicaMax()
	}

	Sp.mux.Unlock()
	return CreateTuple(t1)
}



//~~~~~~~~ NOT TESTED
// Query a specific space for tuples matching the given pattern
// Blocks if the tuple is not found
func Query(p Tuple, Sp Replispace, s Space) Tuple {
	// Sp.mux.Lock()

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
		// Sp.mux.Unlock()
		return t2
	}
	// The empry tuple is only returned in case of errors
	// Sp.mux.Unlock()
	return CreateTuple()
}


//~~~~~~~~
// Query a specific space for tuples matching the given pattern
// Non-blocking version
func QueryP(p Tuple, Sp Replispace, s Space) Tuple {
	Sp.mux.Lock()

	// create template p' = {t,S}
	var y []string // <--- extra field to match the space list S
	var data []interface{}
	data = append(data, p.Fields...)
	// var createTime int64 // <--- extra field to match the creation time
	// data = append(data, &createTime)
	data = append(data, &y)
	var p1 Tuple = CreateTuple(data...)

	// query a tuple via a pattern matching from a specific space
	t1, e := s.QueryP(p1.Fields...)

	if e == nil {
		// no error: return the matching tuple without the last field
		var u = CreateTuple(dataFieldsOf(t1)...)
		afterQuery(t1, Sp, s)
		
		Sp.mux.Unlock()
		return u
	}

	Sp.mux.Unlock()
	return CreateTuple()   // returns an empty tuple when no tuple is available
}


func unsafeGetP(p Tuple, Sp Replispace, s Space) Tuple {
	// create template p' = {t,S}
	var y []string // <--- extra field to match the space list S
	var data []interface{}
	data = append(data, p.Fields...)
	// var createTimestamp int64 // <--- extra field to match the creation time
	// data = append(data, &createTimestamp)
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
			u, e1 := Sp.Sp[space].GetP(t1.Fields...)
			
			if e1 == nil {
				afterGet(t1, Sp, *Sp.Sp[space])
				
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

func unsafeGet(p Tuple, Sp Replispace, s Space) Tuple {
	// create template p' = {t,S}
	var y []string // <--- extra field to match the space list S
	var data []interface{}
	data = append(data, p.Fields...)
	// var createTimestamp int64 // <--- extra field to match the creation time
	// data = append(data, &createTimestamp)
	data = append(data, &y)
	var p1 Tuple = CreateTuple(data...)

	//  search the tuple from space s (blocking)
	t1, e := s.Query(p1.Fields...)

	if e == nil {
		// extract the list of all spaces
		var S = (t1.Fields[len(t1.Fields)-1])

		// transform the interface type of spaces into the string type
		var v []string
		v = S.([]string)

		// for each space in the set S of space identifiers
		for s, space := range v {
			// remove the tuple from the relevant spaces
			u, e1 := Sp.Sp[space].GetP(t1.Fields...)
			
			if e1 == nil {
				afterGet(t1, Sp, *Sp.Sp[space])
				
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
	// returns an empty tuple on error
	return CreateTuple() 
}

func Get(p Tuple, Sp Replispace, s Space) Tuple {
	t := Query(p, Sp, s)
	Sp.mux.Lock()
	result := unsafeGet(t, Sp, s)
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
