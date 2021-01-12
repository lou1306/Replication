/*
	GoA (Go Analyser)
	Simple Block-sensitive Symbol Table for Go programs

HISTORY:
	2020.02.18  first working prototype

TODO:
   - track var := value and infer their type
   - track global variables
*/

package main

import (
	"go/ast"
	"go/token"
	"reflect"
	"strconv"
)

///var sym map[string][]string  // ...not now
type scopedvariable struct {
    block string
    id    string
    kind  string
}

var variables []scopedvariable

// Root node in the syntax tree for any given compound statement identifier.
// Blocks are identified with strings in x:y:z:... format
//var block map[string]*ast.BlockStmt

//var blockput map[int]string // id of the block that contains a put operation
//var countput int // counter for put actions

func scanblock(fset *token.FileSet, node *ast.BlockStmt, blockid string) {
	count := 0 // block identifier

	// Extract the blocks and positions
	debug("symbol table: block (%s) coords (%d:%d..%d:%d) type (%s)", blockid, fset.Position(node.Pos()).Line, fset.Position(node.Pos()).Column, fset.Position(node.End()).Line, fset.Position(node.End()).Column, reflect.TypeOf(node).String())

	// Visit blocks of loop statements
	for _, n := range node.List {
	    //switch decl := d.(type) {
	    //case *ast.FuncDecl:
	    if reflect.TypeOf(n).String() == "*ast.BlockStmt" {
			scanblock(fset, n.(*ast.BlockStmt), blockid+":"+strconv.Itoa(count))
			count += 1
		} else if reflect.TypeOf(n).String() == "*ast.ForStmt" {
			scanblock(fset, n.(*ast.ForStmt).Body, blockid+":"+strconv.Itoa(count))
			count += 1
		} else if reflect.TypeOf(n).String() == "*ast.SwitchStmt" {
			scanblock(fset, n.(*ast.SwitchStmt).Body, blockid+":"+strconv.Itoa(count))
			count += 1
		} else if reflect.TypeOf(n).String() == "*ast.SelectStmt" {
			scanblock(fset, n.(*ast.SelectStmt).Body, blockid+":"+strconv.Itoa(count))
			count += 1
		} else if reflect.TypeOf(n).String() == "*ast.RangeStmt" {
			scanblock(fset, n.(*ast.RangeStmt).Body, blockid+":"+strconv.Itoa(count))
			count += 1
		} else if reflect.TypeOf(n).String() == "*ast.IfStmt" {
			scanblock(fset, n.(*ast.IfStmt).Body, blockid+":"+strconv.Itoa(count))
			count += 1
			t, ok := (n.(*ast.IfStmt).Else).(*ast.BlockStmt)
			if ok {
				scanblock(fset, t, blockid+":"+strconv.Itoa(count))
				count += 1
			}
		} else if reflect.TypeOf(n).String() == "*ast.TypeSwitchStmt" {
			scanblock(fset, n.(*ast.TypeSwitchStmt).Body, blockid+":"+strconv.Itoa(count))
			count += 1
			r, ok := (n.(*ast.TypeSwitchStmt).Assign).(*ast.BlockStmt)
			if ok {
				scanblock(fset, r, blockid+":"+strconv.Itoa(count))
				count += 1
			}
		} else  if reflect.TypeOf(n).String() == "*ast.LabeledStmt" {
			l, ok := (n.(*ast.LabeledStmt).Stmt).(*ast.BlockStmt)
			if ok {
				scanblock(fset, l, blockid+":"+strconv.Itoa(count))
				count += 1
			}
		} else if reflect.TypeOf(n).String() == "*ast.AssignStmt" {
            var err bool
            _, err = n.(*ast.AssignStmt)
            if err {
            	debug("symbol table: block (%s) type (%s) tracked", blockid, reflect.TypeOf(n).String())
            }
        } else if reflect.TypeOf(n).String() == "*ast.ExprStmt" {
			fn1, ok := n.(*ast.ExprStmt)
			if ok && reflect.TypeOf(fn1.X.(*ast.CallExpr).Fun).String() == "*ast.SelectorExpr" {
				if fn1.X.(*ast.CallExpr).Fun.(*ast.SelectorExpr).Sel.Name == "Put" {
					//fmt.Printf(" > Found a call to Put blockid:%s count:%d\n", blockid, countput)
					//blockput[countput] = blockid
					//countput += 1
				}
			}
		} else if reflect.TypeOf(n).String() == "*ast.DeclStmt"	{
			f, ok := n.(*ast.DeclStmt).Decl.(*ast.GenDecl)
			if ok && f.Tok == token.VAR {
				id := f.Specs[0].(*ast.ValueSpec).Names[0].Name
				if reflect.TypeOf(f.Specs[0].(*ast.ValueSpec).Type).String() == "*ast.Ident" {
					kind := f.Specs[0].(*ast.ValueSpec).Type.(*ast.Ident).Name
					variables = append(variables, scopedvariable{block:blockid, id:id, kind:kind})
					debug("symbol table: block (%s) variable (%s) type (%s)", blockid, id, kind)
				} else {  // for the moment we only handle simple identifiers in declarations
					debug("symboltable: block (%s) declaration for variable type (%s) ignored", blockid, reflect.TypeOf(f.Specs[0].(*ast.ValueSpec).Type).String())
				}
			}
		} else if reflect.TypeOf(n).String() == "*ast.GoStmt" {
         var err bool
			_, err = n.(*ast.GoStmt)
			if err {
              debug("symbol table: block (%s) node (%s)", blockid, reflect.TypeOf(n).String())

		} else {
			warn("symbol table: node (%s) discarded", reflect.TypeOf(n).String())
		}
	}
}
}

func buildsymboltable(fset *token.FileSet, node *ast.File) {
	count := 0
	//countput = 0
	//blockput = make(map[int]string)

	for _, node := range node.Decls {
		fn, ok := node.(*ast.FuncDecl)
		if !ok {
			continue
		}

		scanblock(fset, fn.Body, strconv.Itoa(count))
		count++
	}
}


func printsymboltable() {
	for i:=0; i<len(variables); i++ {
		info("block (%s) variable (%s) type (%s)", variables[i].block, variables[i].id, variables[i].kind)
	}
}
