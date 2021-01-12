/*
	Replicator:
	Automated replication of tuple space programs via
	program transformation from goSpaces to RepligoSpaces
	(Prototype implementation)

HISTORY:
	2020.06.05  quiet mode
	2020.06.04  changed targets from uri[*si] to uri[si]
	2020.06.04  bugfix: addglobaldecls() now works if the input program has global variables
	2020.02.20  first working prototype

NOTES:
	- main | functions used to spawn threads | any other extra functions
	See paper for assumptions and restrictions on the input program.

TODO:
	- save transformed program to file (go run *.go -i testfiles/listing4.go  2>&1 | tee -i file1.go)

KNOWN ISSUES
	- transformation of GetP() etc. does not work if the operation to transform is enclosed by commented blocks.

*/

package main

import (
	"flag"
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"
)


// Root node in the syntax tree for any given compound statement identifier.
// Blocks are identified with strings in x:y:z:... format

var block map[string]*ast.BlockStmt
var blockspaces map[int]string // id of the block that contains the spaces
var countspace int             // counter for the spaces

var spaceid map[int]*ast.Ident
var targets []ast.Expr

var newtargets map[*ast.ExprStmt][]ast.Expr

var tar map[int][]int // location targets for each put operation

// Make sure the program contains a main function.
func checkmainfunction(fset *token.FileSet, node *ast.File) bool {
	for _, node := range node.Decls {
		fn, ok := node.(*ast.FuncDecl)

		if ok && fn.Name.Name == "main" {
			return true
		}
	}

	return false
}

// Make sure any thread is spawned from the main function.
func scanthreads(fset *token.FileSet, node *ast.File) int {
	threadcount := 0
	function := ""
	outsidemain := false // any threads spawned from outside the main function?
	var threadnames []string

	ast.Inspect(node, func(n ast.Node) bool {
		switch x := n.(type) {
		case *ast.FuncDecl:
			function = x.Name.Name

		case *ast.GoStmt:
			if function != "main" {
				outsidemain = true
			} else {
				name := x.Call.Fun.(*ast.Ident).Name
				info("transformer: spawning a thread from function (%s)", name)
				threadnames = append(threadnames, name)
				threadcount++
			}
		}

		return true
	})

	if outsidemain {
		return -1
	}

	return threadcount
}

// Make sure all of the spaces are created from within the main function.
func scanspaces(fset *token.FileSet, node *ast.File) int {
	spacecount := 0 // no. of spaces created (by counting the invocations to NewSpace()
	function := ""
	outsidemain := false // any spaces created outside the main function?

	ast.Inspect(node, func(n ast.Node) bool {
		switch x := n.(type) {
		case *ast.FuncDecl:
			function = x.Name.Name

		case *ast.AssignStmt:
			if reflect.TypeOf(x.Rhs[0]).String() == "*ast.CallExpr" &&
				reflect.TypeOf(x.Rhs[0].(*ast.CallExpr).Fun).String() == "*ast.Ident" &&
				x.Rhs[0].(*ast.CallExpr).Fun.(*ast.Ident).Name == "NewSpace" {

				if function != "main" {
					outsidemain = true
				} else {
					spacecount++
					identurl := ast.NewIdent("uri")
					spaceid[spacecount] = x.Lhs[0].(*ast.Ident)
					//warn("---> spaceid[%d]  = (%s)", spacecount, x.Lhs[0].(*ast.Ident))
					indexpr02 := &ast.IndexExpr{X: identurl, Index: x.Lhs[0].(*ast.Ident)} // uri[spaceid]
					targets = append(targets, indexpr02)
				}
			}
		}

		return true
	})

	if outsidemain {
		return -1
	}

	return spacecount
}

//  run input program using command-lines
var inputfilename *string
var outputfilename *string
var showast *bool
var verbose *bool
var quiet *bool

func usage() {
	fmt.Fprintf(os.Stderr, "usage: replicator [inputfile]\n")
	flag.PrintDefaults()
	os.Exit(2)
}

func initparams() {
	inputfilename = flag.String("i", "", "input filename")
	outputfilename = flag.String("o", "", "output filename")
	showast = flag.Bool("show-ast", false, "show the abstract syntax tree and exit")
	verbose = flag.Bool("verbose", false, "show debug messages")
	quiet = flag.Bool("quiet", false, "no log or debug messages")

	// parse command line
	flag.Parse()

	// sanity checks
	if *inputfilename == "" {
		usage()
	}
}

func addimport(fset *token.FileSet, node *ast.File) {
	// Sub-AST for import: import "repligospaces"
	bascilit1 := &ast.BasicLit{Kind: token.STRING, Value: "\"github.com/repligospaces\""} // "github.com/repligospaces"
	identdot := ast.NewIdent(".")                                              // .
	importspec := &ast.ImportSpec{Name: identdot, Path: bascilit1}

	for _, f := range node.Decls {
		fn, ok := f.(*ast.GenDecl)
		// Import declarations
		if ok && fn.Tok == token.IMPORT {
			// Add the new import
			fn.Specs = append(fn.Specs, importspec)
		}

	}
}

func addglobaldecls(fset *token.FileSet, node *ast.File) {
	// Declarations of expressions,
	// useful to an AST for the transformations of the global part.

	//var dec1 *ast.GenDecl
	var exgendecl *ast.GenDecl
	var exgendecl1 *ast.GenDecl
    var exgendecl2 *ast.GenDecl

	// Declarations of useful identifiers.
	identurl := ast.NewIdent("uri")
	identm := ast.NewIdent("Sp")
	identmake := ast.NewIdent("make")
	identstring := ast.NewIdent("string")
    identspace := ast.NewIdent("Space")
	identspace1 := ast.NewIdent("Space")
    identrepligospace := ast.NewIdent("Replispace")
    identrsp := ast.NewIdent("rsp")

    // sub-AST for: var rsp Replispace = Replispace{Sp: Sp}
    Keyvalueexpr := &ast.KeyValueExpr{Key: identm, Value: identm} // Sp: Sp
    compositelit := &ast.CompositeLit{Type: identrepligospace,  Elts: []ast.Expr{Keyvalueexpr}} // Replispace{Sp: Sp}
    valuespec3 := &ast.ValueSpec{Names: []*ast.Ident{identrsp}, Type: identrepligospace, Values: []ast.Expr{compositelit}} // rsp Replispace = Replispace{Sp: Sp}
    exgendecl2 = &ast.GenDecl{Specs: []ast.Spec{valuespec3}, Tok: token.VAR} // var rsp Replispace = Replispace{Sp: Sp}

	// Sub-AST for: var space[] string
	// array1 := &ast.ArrayType{Elt: identstring}                                  // [] string
	// valuespec2 := &ast.ValueSpec{Names: []*ast.Ident{identspace}, Type: array1} // space[]
	// dec1 = &ast.GenDecl{Specs: []ast.Spec{valuespec2}, Tok: token.VAR}          // var space[] string

	// Sub-AST for: var uri = make(map[space]string)
	exmaptype2 := &ast.MapType{Key: identspace, Value: identstring}                         // map[space]string
	excallexpr2 := &ast.CallExpr{Fun: identmake, Args: []ast.Expr{exmaptype2}}              // make(map[space]string)
	dec20 := &ast.ValueSpec{Names: []*ast.Ident{identurl}, Values: []ast.Expr{excallexpr2}} // make(map[space]string)
	exgendecl = &ast.GenDecl{Tok: token.VAR, Specs: []ast.Spec{dec20}}                      // var uri = make(map[space]string)

	// Sub-AST for: var m = make(map[string]*Space)
	value2 := &ast.StarExpr{X: identspace1}                                              // *Space
	maptype := &ast.MapType{Key: identstring, Value: value2}                             // map[string]*Space
	excallexpr := &ast.CallExpr{Fun: identmake, Args: []ast.Expr{maptype}}               // make(map[string]*Space)
	dec21 := &ast.ValueSpec{Names: []*ast.Ident{identm}, Values: []ast.Expr{excallexpr}} // Sp = make(map[string]*Space)
	exgendecl1 = &ast.GenDecl{Tok: token.VAR, Specs: []ast.Spec{dec21}}                  // var Sp = make(map[string]*Space)

	// declare useful variables
	var list1 []ast.Decl
	var done bool

	done = false

	for _, f := range node.Decls {
		list1 = append(list1, f)
		_, ok := f.(*ast.FuncDecl)

		if ok {
			continue
		}

		if !done {
			list1 = append(list1, exgendecl, exgendecl1,  exgendecl2)
			done = true
		}
	}

	node.Decls = list1
}

// Assumption: NewSpace is not invoked within nested blocks!
func transformnewspace(fset *token.FileSet, node *ast.File) {
	var indexpr02 *ast.IndexExpr
	var assignmentO1 *ast.AssignStmt
	var assignmentO2 *ast.AssignStmt

	identm := ast.NewIdent("Sp")
	identurl := ast.NewIdent("uri")

	for _, f := range node.Decls {
		var list []ast.Stmt

		fn, ok := f.(*ast.FuncDecl)
		if !ok {
			continue
		}

		// Find assignment statements in the body list
		for _, k := range fn.Body.List {
			astmt, ok := k.(*ast.AssignStmt)
			list = append(list, k) // get spaces created: space1, space2, ...

			if ok {
				// Visit the argument in the call to NewSpace to extract all the spaces created
				callexpr, ok2 := astmt.Rhs[0].(*ast.CallExpr)

				if ok2 && callexpr.Fun.(*ast.Ident).Name == "NewSpace" && reflect.TypeOf(astmt.Lhs[0]).String() == "*ast.Ident" {
					// Extract left hand side of assignment statement
					// (i.e., the identifier of the Space object).
					spaceid := astmt.Lhs[0].(*ast.Ident)

					// Extract first argument in the function call at the right hand side of the assignment statement
					// (i.e., the URI of the space).
					spaceuri := astmt.Rhs[0].(*ast.CallExpr).Args[0].(*ast.BasicLit)

					// Prepare sub-ASTs to inject in the target program
					// i.e., &spaceid, uri[spaceid], and m[uri]
					space := &ast.UnaryExpr{Op: token.AND, X: spaceid}      //  &spaceid
					indexpr01 := &ast.IndexExpr{X: identm, Index: spaceuri} // sp[uri]
					indexpr02 = &ast.IndexExpr{X: identurl, Index: spaceid} // uri[spaceid]

					// Add assignment statements for m[]:
					// m[spaceuri]  ->  &spaceid
					assignmentO1 = &ast.AssignStmt{Lhs: []ast.Expr{indexpr01}, Tok: token.ASSIGN, Rhs: []ast.Expr{space}}

					// Add assignment statements for uri[]:
					// uri[spaceid] ->  spaceuri
					assignmentO2 = &ast.AssignStmt{Lhs: []ast.Expr{indexpr02}, Tok: token.ASSIGN, Rhs: []ast.Expr{astmt.Rhs[0].(*ast.CallExpr).Args[0].(*ast.BasicLit)}}

					// Join the two parts (assignment01, assignment02) above to the list of statements:
					list = append(list, assignmentO1, assignmentO2)
				}
			}

			// Join the transformed ASTs to the block
			// m[spaceuri]  ->  &spaceid
			// uri[spaceid] ->  spaceuri
			fn.Body.List = list
		}
	}
}

var blockput = make(map[ast.Stmt]*ast.BlockStmt)

// Scan the program to populate blockput[] and putid[] maps,
// to keep track of where put operations occur in the program.
func scanbloc2(fset *token.FileSet, node *ast.BlockStmt, blockid string, parent *ast.BlockStmt) {
	count := 0 // block identifier

	// Extract the blocks and positions
	//fmt.Printf("\n>> Block: %s\n", blockid)
	//fmt.Printf(">> Coords: %d:%d..%d:%d\n", fset.Position(node.Pos()).Line, fset.Position(node.Pos()).Column, fset.Position(node.End()).Line, fset.Position(node.End()).Column)
	//fmt.Printf(">> Type: %s\n\n", reflect.TypeOf(node).String())

	// Visit blocks of loop statements
	for _, n := range node.List {
		if reflect.TypeOf(n).String() == "*ast.BlockStmt" {
			fmt.Printf("a\n")
			scanbloc2(fset, n.(*ast.BlockStmt), blockid+":"+strconv.Itoa(count), n.(*ast.BlockStmt))
			count += 1
		}

		//Visit blocks of for statements
		if reflect.TypeOf(n).String() == "*ast.ForStmt" {
			scanbloc2(fset, n.(*ast.ForStmt).Body, blockid+":"+strconv.Itoa(count), n.(*ast.ForStmt).Body)
			count += 1
		}

		// Visit blocks of switch statements
		if reflect.TypeOf(n).String() == "*ast.SwitchStmt" {
			scanbloc2(fset, n.(*ast.SwitchStmt).Body, blockid+":"+strconv.Itoa(count), n.(*ast.SwitchStmt).Body)
			count += 1
		}

		// Visit blocks of select statements
		if reflect.TypeOf(n).String() == "*ast.SelectStmt" {
			scanbloc2(fset, n.(*ast.SelectStmt).Body, blockid+":"+strconv.Itoa(count), n.(*ast.SelectStmt).Body)
			count += 1
		}

		// Visit blocks of range statements
		if reflect.TypeOf(n).String() == "*ast.RangeStmt" {
			scanbloc2(fset, n.(*ast.RangeStmt).Body, blockid+":"+strconv.Itoa(count), n.(*ast.RangeStmt).Body)
			count += 1
		}

		// Visit blocks of if statements
		if reflect.TypeOf(n).String() == "*ast.IfStmt" {
			scanbloc2(fset, n.(*ast.IfStmt).Body, blockid+":"+strconv.Itoa(count), n.(*ast.IfStmt).Body)
			count += 1
			t, ok := (n.(*ast.IfStmt).Else).(*ast.BlockStmt)
			if ok {
				scanbloc2(fset, t, blockid+":"+strconv.Itoa(count), (n.(*ast.IfStmt).Else).(*ast.BlockStmt))
				count += 1
			}
		}

		// Visit blocks of switch statements
		if reflect.TypeOf(n).String() == "*ast.TypeSwitchStmt" {
			scanbloc2(fset, n.(*ast.TypeSwitchStmt).Body, blockid+":"+strconv.Itoa(count), n.(*ast.TypeSwitchStmt).Body)
			count += 1
			r, ok := (n.(*ast.TypeSwitchStmt).Assign).(*ast.BlockStmt)
			if ok {
				scanbloc2(fset, r, blockid+":"+strconv.Itoa(count), (n.(*ast.TypeSwitchStmt).Assign).(*ast.BlockStmt))
				count += 1
			}
		}

		// Visit blocks of labeled statements
		if reflect.TypeOf(n).String() == "*ast.LabeledStmt" {
			l, ok := (n.(*ast.LabeledStmt).Stmt).(*ast.BlockStmt)
			if ok {
				scanbloc2(fset, l, blockid+":"+strconv.Itoa(count), (n.(*ast.LabeledStmt).Stmt).(*ast.BlockStmt))
				count += 1
			}
		}

		//
		if reflect.TypeOf(n).String() == "*ast.ExprStmt" {
			fn1, ok := n.(*ast.ExprStmt)

			if ok && reflect.TypeOf(fn1.X.(*ast.CallExpr).Fun).String() == "*ast.SelectorExpr" {
				if fn1.X.(*ast.CallExpr).Fun.(*ast.SelectorExpr).Sel.Name == "Put" {
					info("transformer: found a call to Put in block (%s)", blockid)
					blockput[n] = parent
				}
			}
		}
	}
}

// Scan the program to populate blockput[] and putid[] maps,
// to keep track of where put operations occur in the program.
func scanblockput(fset *token.FileSet, node *ast.BlockStmt, blockid string) {
	count := 0 // block identifier
	countput := 0

	// Extract the blocks and positions
	//fmt.Printf("\n>> Block: %s\n", blockid)
	//fmt.Printf(">> Coords: %d:%d..%d:%d\n", fset.Position(node.Pos()).Line, fset.Position(node.Pos()).Column, fset.Position(node.End()).Line, fset.Position(node.End()).Column)
	//fmt.Printf(">> Type: %s\n\n", reflect.TypeOf(node).String())

	// Visit blocks of loop statements
	for _, n := range node.List {
		if reflect.TypeOf(n).String() == "*ast.BlockStmt" {
			scanblockput(fset, n.(*ast.BlockStmt), blockid+":"+strconv.Itoa(count))
			count += 1
		}

		if reflect.TypeOf(n).String() == "*ast.ForStmt" {
			scanblockput(fset, n.(*ast.ForStmt).Body, blockid+":"+strconv.Itoa(count))
			count += 1
		}

		if reflect.TypeOf(n).String() == "*ast.SwitchStmt" {
			scanblockput(fset, n.(*ast.SwitchStmt).Body, blockid+":"+strconv.Itoa(count))
			count += 1
		}

		if reflect.TypeOf(n).String() == "*ast.SelectStmt" {
			scanblockput(fset, n.(*ast.SelectStmt).Body, blockid+":"+strconv.Itoa(count))
			count += 1
		}

		if reflect.TypeOf(n).String() == "*ast.RangeStmt" {
			scanblockput(fset, n.(*ast.RangeStmt).Body, blockid+":"+strconv.Itoa(count))
			count += 1
		}
		if reflect.TypeOf(n).String() == "*ast.IfStmt" {
			scanblockput(fset, n.(*ast.IfStmt).Body, blockid+":"+strconv.Itoa(count))
			count += 1
			t, ok := (n.(*ast.IfStmt).Else).(*ast.BlockStmt)
			if ok {
				scanblockput(fset, t, blockid+":"+strconv.Itoa(count))
				count += 1
			}
		}

		if reflect.TypeOf(n).String() == "*ast.TypeSwitchStmt" {
			scanblockput(fset, n.(*ast.TypeSwitchStmt).Body, blockid+":"+strconv.Itoa(count))
			count += 1
			r, ok := (n.(*ast.TypeSwitchStmt).Assign).(*ast.BlockStmt)
			if ok {
				scanblockput(fset, r, blockid+":"+strconv.Itoa(count))
				count += 1
			}
		}

		if reflect.TypeOf(n).String() == "*ast.LabeledStmt" {
			l, ok := (n.(*ast.LabeledStmt).Stmt).(*ast.BlockStmt)
			if ok {
				scanblockput(fset, l, blockid+":"+strconv.Itoa(count))
				count += 1
			}
		}

		if reflect.TypeOf(n).String() == "*ast.ExprStmt" {
			fn1, ok := n.(*ast.ExprStmt)

			if ok && reflect.TypeOf(fn1.X.(*ast.CallExpr).Fun).String() == "*ast.SelectorExpr" {
				if fn1.X.(*ast.CallExpr).Fun.(*ast.SelectorExpr).Sel.Name == "Put" {
					info("transformer: transforming call to Put in block (%s)", blockid)

					if len(targetlocations[fn1.X.(*ast.CallExpr)]) > 0 {
						for i:=0; i<len(targetlocations[fn1.X.(*ast.CallExpr)]); i++ {
							info("transformer: replicating the tuple to process (%d)", targetlocations[fn1.X.(*ast.CallExpr)][i])
							identurl := ast.NewIdent("uri")
							spaceidd := spaceid[targetlocations[fn1.X.(*ast.CallExpr)][i]]
							////identstarexpr := &ast.StarExpr{X: spaceidd}
							////indexpr02 := &ast.IndexExpr{X: identurl, Index: identstarexpr} // uri[* spaceid]
							indexpr02 := &ast.IndexExpr{X: identurl, Index: spaceidd} // uri[* spaceid]
							newtargets[fn1] = append(newtargets[fn1], indexpr02)
						}
					}

					var expr ast.Expr

					// Extract all the args associated with the function call "Put" and create a new args
					newArgs := make([]ast.Expr, len(fn1.X.(*ast.CallExpr).Args))
					copy(newArgs, fn1.X.(*ast.CallExpr).Args)

					// Sub-AST for: t := CreateTuple(args)
					identcreatetuple := ast.NewIdent("CreateTuple") // CreateTuple
					callexpr01 := &ast.CallExpr{Fun: identcreatetuple, Args: newArgs} // CreateTuple(newargs)

					// Sub-AST for: a := make([]string, n)
					identa := ast.NewIdent(fmt.Sprintf("targets%d", countput))
					identmake := ast.NewIdent("make")
					identrsp1 := ast.NewIdent("rsp")
					identstring := ast.NewIdent("string")
					basiclit := &ast.BasicLit{Kind: token.STRING, Value: fmt.Sprintf("%d", len(newtargets[fn1]))}                             // length of the target spaces
					arraytype := &ast.ArrayType{Elt: identstring}                                                            // []string
					callexpr02 := &ast.CallExpr{Fun: identmake, Args: []ast.Expr{arraytype, basiclit}}                       // make([]string, m)
					assignment05 := &ast.AssignStmt{Lhs: []ast.Expr{identa}, Tok: token.DEFINE, Rhs: []ast.Expr{callexpr02}} // a := make([]string, 2)

					// ....
					/*
					for i := range newtargets[fn1] {
						//warn("--> i=%d  boh=%s", i,newtargets[fn1][i])
						if newtargets[fn1][i] != nil {
							//newArgs = append(newArgs, newtargets[fn1][i])
						}
					}
					*/

					// Sub-AST for : a[i] = uri[spaceid]
					// iterate over the set of target spaces

                     ///////////////////////////////////////////////////

					for i, _ := range newtargets[fn1] {
						basiclit2 := &ast.BasicLit{Kind: token.STRING, Value: fmt.Sprintf("%d", i)} // length of the target spaces
						indexexpr03 := &ast.IndexExpr{X: identa, Index: basiclit2}                  // a[i]
						assignment06 := &ast.AssignStmt{Lhs: []ast.Expr{indexexpr03}, Tok: token.ASSIGN, Rhs: []ast.Expr{newtargets[fn1][i]}}
						//blockput[n].List = append(blockput[n].List, assignment06) // a0[0] = uri[s1]
						blockput[n].List = append([]ast.Stmt{assignment06}, blockput[n].List...) // a0[0] = uri[s1]
					}
                    ///////////////////////////////////////////////////

					blockput[n].List = append([]ast.Stmt{assignment05}, blockput[n].List...) // a0 := make([]string, 12345)

					// Add the tuples to the set of target spaces:
					// Put(CreateTuple(args), a)
					expr = &ast.CallExpr{Fun: fn1.X.(*ast.CallExpr).Fun.(*ast.SelectorExpr).Sel, Args: []ast.Expr{callexpr01, identrsp1, identa}}

					if expr != nil {
						fn1.X = expr
					}

					countput++
				}
			}
		}
	}

}

// Scan the program to populate blockput[] and putid[] maps,
// to keep track of where put operations occur in the program.
func scanblockget(fset *token.FileSet, node *ast.BlockStmt, blockid string) {
	count := 0 // block identifier

	// Extract the blocks and positions
	//fmt.Printf("\n>> Block: %s\n", blockid)
	//fmt.Printf(">> Coords: %d:%d..%d:%d\n", fset.Position(node.Pos()).Line, fset.Position(node.Pos()).Column, fset.Position(node.End()).Line, fset.Position(node.End()).Column)
	//fmt.Printf(">> Type: %s\n\n", reflect.TypeOf(node).String())

	// Visit blocks of loop statements
	for _, n := range node.List {
		if reflect.TypeOf(n).String() == "*ast.BlockStmt" {
			scanblockget(fset, n.(*ast.BlockStmt), blockid+":"+strconv.Itoa(count))
			count += 1
		}

		//Visit blocks of for statements
		if reflect.TypeOf(n).String() == "*ast.ForStmt" {
			scanblockget(fset, n.(*ast.ForStmt).Body, blockid+":"+strconv.Itoa(count))
			count += 1
		}

		// Visit blocks of switch statements
		if reflect.TypeOf(n).String() == "*ast.SwitchStmt" {
			scanblockget(fset, n.(*ast.SwitchStmt).Body, blockid+":"+strconv.Itoa(count))
			count += 1
		}

		// Visit blocks of select statements
		if reflect.TypeOf(n).String() == "*ast.SelectStmt" {
			scanblockget(fset, n.(*ast.SelectStmt).Body, blockid+":"+strconv.Itoa(count))
			count += 1
		}

		// Visit blocks of range statements
		if reflect.TypeOf(n).String() == "*ast.RangeStmt" {
			scanblockget(fset, n.(*ast.RangeStmt).Body, blockid+":"+strconv.Itoa(count))
			count += 1
		}
		// Visit blocks of if statements
		if reflect.TypeOf(n).String() == "*ast.IfStmt" {
			scanblockget(fset, n.(*ast.IfStmt).Body, blockid+":"+strconv.Itoa(count))
			count += 1
			t, ok := (n.(*ast.IfStmt).Else).(*ast.BlockStmt)
			if ok {
				scanblockget(fset, t, blockid+":"+strconv.Itoa(count))
				count += 1
			}
		}

		// Visit blocks of switch statements
		if reflect.TypeOf(n).String() == "*ast.TypeSwitchStmt" {
			scanblockget(fset, n.(*ast.TypeSwitchStmt).Body, blockid+":"+strconv.Itoa(count))
			count += 1
			r, ok := (n.(*ast.TypeSwitchStmt).Assign).(*ast.BlockStmt)
			if ok {
				scanblockget(fset, r, blockid+":"+strconv.Itoa(count))
				count += 1
			}
		}

		// Visit blocks of labeled statements
		if reflect.TypeOf(n).String() == "*ast.LabeledStmt" {
			l, ok := (n.(*ast.LabeledStmt).Stmt).(*ast.BlockStmt)
			if ok {
				scanblockget(fset, l, blockid+":"+strconv.Itoa(count))
				count += 1
			}
		}

		// Transform calls to GoSpace routines
		if reflect.TypeOf(n).String() == "*ast.ExprStmt" {
			fn1, ok := n.(*ast.ExprStmt)

			if ok && reflect.TypeOf(fn1.X.(*ast.CallExpr).Fun).String() == "*ast.SelectorExpr" {
				if fn1.X.(*ast.CallExpr).Fun.(*ast.SelectorExpr).Sel.Name == "QueryP" ||
				   fn1.X.(*ast.CallExpr).Fun.(*ast.SelectorExpr).Sel.Name == "GetP"{
					info("transformer: transforming call to GetP or QueryP block (%s)", blockid)
					newArgs := make([]ast.Expr, len(fn1.X.(*ast.CallExpr).Args))
					copy(newArgs, fn1.X.(*ast.CallExpr).Args)

					var spaceidd string
					if strings.IndexByte(blockid, ':') != -1 {
						spaceidd = blockid[:strings.IndexByte(blockid, ':')]
					} else {
						spaceidd = blockid
					}
					spaceidd = "s" + spaceidd

					expr2 := ast.NewIdent(spaceidd)
					//identurl := ast.NewIdent("uri")
					//indexpr02 := &ast.IndexExpr{X: identurl, Index: expr2} // uri[spaceid]
					identcreatetuple2 := ast.NewIdent("CreateTuple")
					identrsp2 := ast.NewIdent("rsp")
					callexpr02 := &ast.CallExpr{Fun: identcreatetuple2, Args: newArgs}

					//newArgs = append(newArgs, indexpr02)
					newArgs = append(newArgs, expr2)
					// generate: GetP(Template, uri[spaceid])
					var expr ast.Expr

					expr = &ast.CallExpr{Fun: fn1.X.(*ast.CallExpr).Fun.(*ast.SelectorExpr).Sel, Args: []ast.Expr{callexpr02, identrsp2, expr2}}

					if expr != nil {
						fn1.X = expr
					}
				}
			}
		}
	}

}

func transformputgetoperations(fset *token.FileSet, node *ast.File) {
	//
	count := 0

	for _, node := range node.Decls {
		fn, ok := node.(*ast.FuncDecl)
		if !ok {
			continue
		}

		scanbloc2(fset, fn.Body, strconv.Itoa(count), fn.Body)
		count++
	}

	// Put operations
	count = 0

	for _, node := range node.Decls {
		fn, ok := node.(*ast.FuncDecl)
		if !ok {
			continue
		}

		scanblockput(fset, fn.Body, strconv.Itoa(count))
		count++
	}

	// Get operations
	count = 0

	for _, node := range node.Decls {
		fn, ok := node.(*ast.FuncDecl)
		if !ok {
			continue
		}

		info("transformer: current function (%s) block (%d)", fn.Name.Name, count)

		scanblockget(fset, fn.Body, strconv.Itoa(count))
		count++
	}
}

func sanitycheck(fset *token.FileSet, node *ast.File) {
	if checkmainfunction(fset, node) == false {
		error("input program has no main function.")
	}

	threadcount := scanthreads(fset, node)
	if threadcount == -1 {
		error("process spawning is only allowed from within the main function")
	}

	spacecount := scanspaces(fset, node)
	if spacecount == -1 {
		error("spaces created outside the main function")
	}

	if spacecount != threadcount {
		error("number of spaces different from number of threads")
	}

	for i := 0; i < spacecount; i++ {
		tar[i] = append(tar[i], i)

		for j := 0; j < spacecount; j++ {
			tar[i] = append(tar[i], j)
		}
	}

}

func main() {
	initparams()

	// Parse input file and build symbol table.
	info("parsing")
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, *inputfilename, nil, parser.ParseComments)
	if err != nil {
		log.Fatal(err)
	}

	if *showast {
		ast.Fprint(os.Stdout, fset, node, nil)
	}

	info("building symbol table")
	buildsymboltable(fset,node)

	// Initialise data structures.
	blockspaces = make(map[int]string)
	tar = make(map[int][]int)
	spaceid = make(map[int]*ast.Ident)
	newtargets = make(map[*ast.ExprStmt][]ast.Expr)

	// Input sanity check.
	sanitycheck(fset,node)

	// Static Analysis
	info("static analysis")
	preanalyse(fset,node)
	analyse(fset,node)  // populate targetlocations

	// Transformation.
	info("program transformation")
	addimport(fset, node)                 // add import statement for repligospace
	addglobaldecls(fset, node)            // add global declarations for auxilliary data structures
	transformnewspace(fset, node)         // add extra operations right after creating a new space
	transformputgetoperations(fset, node) // transform get and put operations into replica-aware operations

	if *outputfilename == "" {
		printer.Fprint(os.Stdout, fset, node)
	} else {
		f, _ := os.Create(*outputfilename)
		printer.Fprint(f, fset, node)
	}
}
