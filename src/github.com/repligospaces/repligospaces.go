/*

CHANGES:
	2020.06.04  bugfix: GetP() was not correctly removing all replicas

*/

package repligospaces

// Importing the useful packages
import (
   ///"fmt"
    . "github.com/pspaces/gospace"
    "sync"
)

type Replispace struct {
    mux sync.Mutex
    Sp  map[string]*Space
}

var wcnt int

//~~~~~~~~
// Add a tuple to the set S of spaces identifiers
func Put(t Tuple, Sp Replispace, S []string) Tuple {
    Sp.mux.Lock()

    // create tuple t' = {t,S}
    var data []interface{}
    data = append(data, t.Fields...)
    data = append(data, S)
    var t1 Tuple = CreateTuple(data...)

    // add t' to each space in S
    for i := 0; i < len(S); i++ {
        Sp.Sp[S[i]].Put(t1.Fields...)
        wcnt+=1
    }

    Sp.mux.Unlock()
    return CreateTuple(t1)
}

// Not tested

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
    var y []string // <--- extra field to match the space list S
    var data []interface{}
    data = append(data, p.Fields...)
    data = append(data, &y)
    var p1 Tuple = CreateTuple(data...)

    // query a tuple via a pattern matching from a specific space
    t1, e := s.QueryP(p1.Fields...)

    if e == nil {
        // no error: return the matching tuple without the last field
        var u = CreateTuple(t1.Fields[:len(t1.Fields)-1]...)
        Sp.mux.Unlock()
        return u
    }

    Sp.mux.Unlock()
    return CreateTuple()   // returns an empty tuple when no tuple is available
}

//~~~~~~~~
// Remove a tuple from space s, and
// from any other space where it might have been replicated.
//
// Note that if no matching tuples are found in s, no tuple
// is removed from any other space.
func GetP(p Tuple, Sp Replispace, s Space) Tuple {
    Sp.mux.Lock()

    // create template p' = {t,S}
    var y []string // <--- extra field to match the space list S
    var data []interface{}
    data = append(data, p.Fields...)
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
        for s := range v {
            // remove the tuple from the relevant spaces
            u, e1 := Sp.Sp[v[s]].GetP(p1.Fields...)

            if e1 == nil {
				if s == len(v)-1 {
               		// no error: tuple successfully removed from the space
                	/////////fmt.Println("  ...Tuple removed from:", v[s])
                	u = CreateTuple(u.Fields[:len(u.Fields)-1]...)
		        	Sp.mux.Unlock()
		        	return u
		        }
            }
        }
    }

    Sp.mux.Unlock()
    return CreateTuple()   // returns an empty tuple when no tuple is available
}

func Getwcount() int {
	return wcnt
}


