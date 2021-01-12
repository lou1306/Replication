/*
	Static Analysis Module.

	This module implementsthe analyse() method that computes
	the (over-approximated) set of target location(s) for each Put operation,
	by checking for possible matches with the Get/GetP/QueryP operations in the program.

	This set is stored in the global variable

	    targetlocations map[*ast.CallExpr][]int,

	as a map from pointers to the AST node for the operation (e.g., space.Get(...)) to
	the identifiers of the target locations/processess.

HISTORY:
	2020.06.04  bugfix: matching algo was not considering QueryP() operations
	2020.02.19  initial prototype

NOTES:
	See paper for assumptions and restrictions on the input program.
	The matching mechanism described in the paper is implemented in comparetuples().

	Minor remarks:
		- currently blocks are only indexed at at the outermost scope (i.e., block = "0", "2", ...)
		- currently insensitive to unaryops (i.e., not distinguishing between &var, var and *var)
		- locally control-flow insensitive
	 	 (e.g., a Get operation occurring before or after the same block of a matching Put
	  	  will produce the same target locations).

TODO:
	- should test more extensively visitstmt() and previsitstmt() to make sure not to drop relevant tuple operation
	- avoid duplicate targets (upon comparetuples())
*/

package main

import (
	"bytes"
	//"flag"
	"go/ast"
	//"go/parser"
	"go/printer"
	"go/token"
	//"log"
	//"os"
	"reflect"
	"strings"
	"strconv"
)


var targetlocations map[*ast.CallExpr][]int  // target processes for each put operation computed via static analsis

var currentprocess string                    // identifier of the process function being visited (i.e., the actual function name)
var currentscope int                         // the outermost scope as an integer (e.g., 0, 1, etc.)
var tuples [128][]*ast.CallExpr              // all tuples in the program (stored as a process-indexed array of slices of tuples)


// Extract argument tuple for any Get/Put operation and add it to global variable tuples[]
func prescantupleoperation(fset *token.FileSet, expr ast.Expr) {
	if reflect.TypeOf(expr).String() == "*ast.CallExpr" {
		if reflect.TypeOf((expr).(*ast.CallExpr).Fun).String() == "*ast.SelectorExpr" {
			if (expr).(*ast.CallExpr).Fun.(*ast.SelectorExpr).Sel.Name == "Put" || (expr).(*ast.CallExpr).Fun.(*ast.SelectorExpr).Sel.Name == "GetP" ||
				(expr).(*ast.CallExpr).Fun.(*ast.SelectorExpr).Sel.Name == "QueryP" || (expr).(*ast.CallExpr).Fun.(*ast.SelectorExpr).Sel.Name == "Get" {
				info("analyser: process (%s) block (%d) line (%d:%d): new tuple operation (%s)", currentprocess, currentscope, fset.Position(expr.Pos()).Line, fset.Position(expr.Pos()).Column, (expr).(*ast.CallExpr).Fun.(*ast.SelectorExpr).Sel.Name)
				tuples[currentscope] = append(tuples[currentscope],expr.(*ast.CallExpr))

				for i:=0; i<len(expr.(*ast.CallExpr).Args); i++ {
					field := expr.(*ast.CallExpr).Args[i]
					fieldtype := reflect.TypeOf(expr.(*ast.CallExpr).Args[i]).String()

					if fieldtype == "*ast.BasicLit" { // constant value or constant string
						debug("analyser: process (%s) block (%d) line (%d:%d): field (%d) is a constant value (%s) of kind (%s)", currentprocess, currentscope, fset.Position(field.Pos()).Line, fset.Position(field.Pos()).Column, i, field.(*ast.BasicLit).Value, strings.ToLower(field.(*ast.BasicLit).Kind.String()))
					} else if fieldtype == "*ast.Ident" { // variable
						debug("analyser: process (%s) block (%d) line (%d:%d): field (%d) is a variable identifier (%s)", currentprocess, currentscope, fset.Position(field.Pos()).Line, fset.Position(field.Pos()).Column, i, field.(*ast.Ident).Name)
					} else if fieldtype == "*ast.UnaryExpr" { // reference to variable
						debug("analyser: process (%s) block (%d) line (%d:%d): field (%d) is a reference to identifier (%s)", currentprocess, currentscope, fset.Position(field.Pos()).Line, fset.Position(field.Pos()).Column, i, field.(*ast.UnaryExpr).X)
					} else {
						debug("analyser: process (%s) block (%d) line (%d:%d): unhandled field type (%s)", currentprocess, currentscope, fieldtype)
					}
				}
			}
		}
	}
}

func previsitstmt(fset *token.FileSet,stmt ast.Stmt) {
	if reflect.TypeOf(stmt).String() == "*ast.ExprStmt" {
		prescantupleoperation(fset,stmt.(*ast.ExprStmt).X)
		if reflect.TypeOf(stmt).String() == "*ast.AssignStmt" {
			prescantupleoperation(fset,stmt.(*ast.AssignStmt).Rhs[0])
		}
		if reflect.TypeOf(stmt).String() == "*ast.AssignStmt" {
			prescantupleoperation(fset,stmt.(*ast.AssignStmt).Lhs[0])
		}
	} else if reflect.TypeOf(stmt).String() == "*ast.ForStmt" {
		previsitstmt(fset,stmt.(*ast.ForStmt).Body)
	} else if reflect.TypeOf(stmt).String() == "*ast.IfStmt" {
		previsitstmt(fset,stmt.(*ast.IfStmt).Body)
		_, ok := (stmt.(*ast.IfStmt).Else).(*ast.BlockStmt)
		if ok {
			previsitstmt(fset,stmt.(*ast.IfStmt).Else)
		}
	} else if reflect.TypeOf(stmt).String() == "*ast.BlockStmt" {
		for _, i := range stmt.(*ast.BlockStmt).List {
			previsitstmt(fset,i)
		} 
	} else if reflect.TypeOf(stmt).String() == "*ast.IncDecStmt" {
prescantupleoperation(fset,stmt.(*ast.IncDecStmt).X)
} else if reflect.TypeOf(stmt).String() == "*ast.GoStmt" {
prescantupleoperation(fset,stmt.(*ast.GoStmt).Call)
} else if reflect.TypeOf(stmt).String() == "*ast.SwitchStmt " {
previsitstmt(fset,stmt.(*ast.SwitchStmt).Body)
_, ok := (stmt.(*ast.SwitchStmt).Init).(*ast.BlockStmt)
if ok {
previsitstmt(fset,stmt.(*ast.SwitchStmt).Init)
}
} else if reflect.TypeOf(stmt).String() == "*ast.SelectStmt" {
previsitstmt(fset,stmt.(*ast.SelectStmt).Body)
} else if reflect.TypeOf(stmt).String() == "*ast.SendStmt" {
prescantupleoperation(fset,stmt.(*ast.SendStmt).Chan)
} else if reflect.TypeOf(stmt).String() == "*ast.SendStmt" {
prescantupleoperation(fset,stmt.(*ast.SendStmt).Value)
} else if reflect.TypeOf(stmt).String() == "*ast.TypeSwitchStmt" {
previsitstmt(fset,stmt.(*ast.TypeSwitchStmt).Body)
_, ok := (stmt.(*ast.TypeSwitchStmt).Assign).(*ast.BlockStmt)
if ok {
previsitstmt(fset,stmt.(*ast.TypeSwitchStmt).Assign)
}
} else if reflect.TypeOf(stmt).String() == "*ast.LabeledStmt" {
previsitstmt(fset,stmt.(*ast.LabeledStmt).Stmt)
} else if reflect.TypeOf(stmt).String() == "*ast.DeferStmt" {
prescantupleoperation(fset,stmt.(*ast.DeferStmt).Call)
} else if reflect.TypeOf(stmt).String() == "*ast.CommClause" {
previsitstmt(fset,stmt.(*ast.CommClause).Comm)
} else if reflect.TypeOf(stmt).String() == "*ast.CaseClause" {
for _, i := range stmt.(*ast.CaseClause).List {
 prescantupleoperation(fset,i)
}
} else if reflect.TypeOf(stmt).String() == "*ast.BranchStmt" {
prescantupleoperation(fset,stmt.(*ast.BranchStmt).Label)
} else if reflect.TypeOf(stmt).String() == "*ast.ForStmt" {
previsitstmt(fset,stmt.(*ast.ForStmt).Init)
} else if reflect.TypeOf(stmt).String() == "*ast.ForStmt" {
prescantupleoperation(fset,stmt.(*ast.ForStmt).Cond)
} else if reflect.TypeOf(stmt).String() == "*ast.ForStmt" {
previsitstmt(fset,stmt.(*ast.ForStmt).Post)
} else if reflect.TypeOf(stmt).String() == "*ast.RangeStmt" {
prescantupleoperation(fset,stmt.(*ast.RangeStmt).Key)
} else if reflect.TypeOf(stmt).String() == "*ast.RangeStmt" {
prescantupleoperation(fset,stmt.(*ast.RangeStmt).Value)
}
}

// Pre-scan the program to populate global variables tuples[]
func preanalyse(fset *token.FileSet, node *ast.File) {
	var count int
	count = 0
	currentscope = 0

	for _, f := range node.Decls {
		// Find the functions
		fn, ok := f.(*ast.FuncDecl)
		if !ok {
			continue
		}

		// If the function is not main: print the process name
		if fn.Name.Name != "main" {
			//info("analyser: scanning process (%s)", fn.Name.Name)
			currentprocess = fn.Name.Name
		}

		// Traversing through the block of statements
		for _, r := range fn.Body.List {
			previsitstmt(fset,r)
			count++
		}

		currentscope++
	}
}

// Fetch the type of a variable from the (previously built) symbol table.
// Note: this is not block-indexed but function-indexed (i.e., does not consider nested blocks), but
//       the symbol table is block-indexed.
func getvariabletype(id string, processblockid int) string {
	blockid := strconv.Itoa(processblockid)

	for i:=0; i<len(variables); i++ {
		debug("(looking for:%s,scope:%s) %s scope:%s processblockid:%d id:%s" , id, blockid, variables[i].kind, variables[i].block, processblockid, variables[i].id)

		if variables[i].id == id && strings.HasPrefix(variables[i].block, blockid) {
			return variables[i].kind
		}
	}

	return ""
}

// Compare two tuple operations and if necessary their argument tuples.
// Returns 0 if the operations will definitely not match,
// or 1 otherwise, so safely over-approximate.
// For the matching algorithm, see the paper.
func comparetuples(fset *token.FileSet,  expr ast.Expr, t1 *ast.CallExpr, scope1 int, t2 *ast.CallExpr, scope2 int) int {
	var buf1 bytes.Buffer
	printer.Fprint(&buf1, fset, t1)
	t1str := buf1.String()                       // the whole expression as a string, e.g., "s1.GetP(&desc, &key)"
	t1op := t1.Fun.(*ast.SelectorExpr).Sel.Name  // GetP, Put, ...

	var buf2 bytes.Buffer
	printer.Fprint(&buf2, fset, t2)
	t2str := buf2.String()
	t2op := t2.Fun.(*ast.SelectorExpr).Sel.Name

	info("analyser: comparing (%s) and (%s)", t1str, t2str)

	// 0. check that read-write match
	if !(((t1op == "Get" || t1op == "GetP" || t1op == "Query" || t1op == "QueryP") && (t2op == "Put")) ||
	    ((t2op == "Get" || t2op == "GetP" || t2op == "Query" || t2op == "QueryP") && (t1op == "Put"))) {
		info("analyser: operations can not interfere with each other")
		return 0
	}

	// 1. check number of arguments
	if len(t1.Args) != len(t2.Args) {
		info("analyser: non-matching argument tuples due to different number of arguments")
		return 0   // t1 and t2 do not match as they have a different number of arguments
	}

	// 2. check type of arguments
	for i:=0; i<len(t1.Args); i++ {
		field1 := t1.Args[i]
		fieldtype1 := reflect.TypeOf(t1.Args[i]).String()

		field2 := t2.Args[i]
		fieldtype2 := reflect.TypeOf(t2.Args[i]).String()

		if fieldtype1 == "*ast.BasicLit" { // constant value or constant string
			if fieldtype2 == "*ast.BasicLit" {
				if field1.(*ast.BasicLit).Value != field2.(*ast.BasicLit).Value {
					info("analyser: non-matching argument tuples due to different types for field (%d)", i)
					return 0
				}
			} else if fieldtype2 == "*ast.Ident" {
				if strings.ToLower(field1.(*ast.BasicLit).Kind.String()) != getvariabletype(field2.(*ast.Ident).Name,scope2) {
					info("analyser: non-matching argument tuples due to different types for field (%d)", i)
					return 0
				}
			} else if fieldtype2 == "*ast.UnaryExpr" {
				if strings.ToLower(field1.(*ast.BasicLit).Kind.String()) != getvariabletype(field2.(*ast.UnaryExpr).X.(*ast.Ident).Name,scope2) {
					info("analyser: non-matching argument tuples due to different types for field (%d)", i)
					return 0
				}
			}
		} else if fieldtype1 == "*ast.Ident" { // variable
			if fieldtype2 == "*ast.BasicLit" {
				if getvariabletype(field1.(*ast.Ident).Name,scope1) != strings.ToLower(field2.(*ast.BasicLit).Kind.String()) {
					info("analyser: non-matching argument tuples due to different types for field (%d)", i)
					return 0
				}
			} else if fieldtype2 == "*ast.Ident" {
				if getvariabletype(field1.(*ast.Ident).Name,scope1) != getvariabletype(field2.(*ast.Ident).Name,scope2) {
					info("analyser: non-matching argument tuples due to different types for field (%d)", i)
					return 0
				}
			} else if fieldtype2 == "*ast.UnaryExpr" {
				if getvariabletype(field1.(*ast.Ident).Name,scope1) != getvariabletype(field2.(*ast.UnaryExpr).X.(*ast.Ident).Name,scope2) {
					info("analyser: non-matching argument tuples due to different types for field (%d)", i)
					return 0
				}
			}
		} else if fieldtype1 == "*ast.UnaryExpr" { // reference to variable
			if fieldtype2 == "*ast.BasicLit" {
				if getvariabletype(field1.(*ast.UnaryExpr).X.(*ast.Ident).Name,scope1) != strings.ToLower(field2.(*ast.BasicLit).Kind.String()) {
					info("analyser: non-matching argument tuples due to different types for field (%d)", i)
					return 0
				}
			} else if fieldtype2 == "*ast.Ident" {
				if getvariabletype(field1.(*ast.UnaryExpr).X.(*ast.Ident).Name,scope1) != getvariabletype(field2.(*ast.Ident).Name,scope2) {
					info("analyser: non-matching argument tuples due to different types for field (%d)", i)
					return 0
				}
			} else if fieldtype2 == "*ast.UnaryExpr" {
				if getvariabletype(field1.(*ast.UnaryExpr).X.(*ast.Ident).Name,scope1) != getvariabletype(field2.(*ast.UnaryExpr).X.(*ast.Ident).Name,scope2) {
					info("analyser: non-matching argument tuples due to different types for field (%d)", i)
					return 0
				}
			}
		} else {
			debug("analyser: process (%s) block (%d) line (%d:%d): unhandled field type (%s)", currentprocess, currentscope, fieldtype1)
		}
	}

	// 3. in any other case, consider the tuples potentially matching
	info("analyser: potentially matching argument tuples")
	return 1   //t1 and t2 can potentially match
}

// Extract tuples (provided as arguments to Put, QueryP, GetP, and Get) and scan their fields
func scantupleoperation(fset *token.FileSet, expr ast.Expr) {
	if reflect.TypeOf(expr).String() == "*ast.CallExpr" {
		if reflect.TypeOf((expr).(*ast.CallExpr).Fun).String() == "*ast.SelectorExpr" {
			if (expr).(*ast.CallExpr).Fun.(*ast.SelectorExpr).Sel.Name == "Put" || (expr).(*ast.CallExpr).Fun.(*ast.SelectorExpr).Sel.Name == "GetP" ||
				(expr).(*ast.CallExpr).Fun.(*ast.SelectorExpr).Sel.Name == "QueryP" || (expr).(*ast.CallExpr).Fun.(*ast.SelectorExpr).Sel.Name == "Get" {
				info("analyser: process (%s) block (%d) line (%d:%d): tuple operation (%s)", currentprocess, currentscope, fset.Position(expr.Pos()).Line, fset.Position(expr.Pos()).Column, (expr).(*ast.CallExpr).Fun.(*ast.SelectorExpr).Sel.Name)

				for i:=0; i<128; i++ {
					//if currentscope == i { continue }  // do not compare tuples within the same process

					if (expr).(*ast.CallExpr).Fun.(*ast.SelectorExpr).Sel.Name == "Put" {  // only check Put operations
						for j:=0; j<len(tuples[i]); j++ {
							if comparetuples(fset, expr, expr.(*ast.CallExpr), currentscope, tuples[i][j], i) == 1 {  // TODO: avoid duplicating!!
								info("analyser: this tuple should be replicated to process (%d)", i)
								targetlocations[expr.(*ast.CallExpr)] = append(targetlocations[expr.(*ast.CallExpr)],i)
							}
						}
					}
				}
			}
		}
	}
}

func visitstmt(fset *token.FileSet,stmt ast.Stmt) {
	if reflect.TypeOf(stmt).String() == "*ast.ExprStmt" {
		scantupleoperation(fset,stmt.(*ast.ExprStmt).X)
		if reflect.TypeOf(stmt).String() == "*ast.AssignStmt" {
			scantupleoperation(fset,stmt.(*ast.AssignStmt).Rhs[0])
		}
		if reflect.TypeOf(stmt).String() == "*ast.AssignStmt" {
			scantupleoperation(fset,stmt.(*ast.AssignStmt).Lhs[0])
		}
	} else if reflect.TypeOf(stmt).String() == "*ast.ForStmt" {
		visitstmt(fset,stmt.(*ast.ForStmt).Body)
	} else if reflect.TypeOf(stmt).String() == "*ast.IfStmt" {
		visitstmt(fset,stmt.(*ast.IfStmt).Body)
		_, ok := (stmt.(*ast.IfStmt).Else).(*ast.BlockStmt)
		if ok {
			visitstmt(fset,stmt.(*ast.IfStmt).Else)
		}
	} else if reflect.TypeOf(stmt).String() == "*ast.BlockStmt" {
		for _, i := range stmt.(*ast.BlockStmt).List {
			visitstmt(fset,i)
		}//
} else if reflect.TypeOf(stmt).String() == "*ast.IncDecStmt" {
scantupleoperation(fset,stmt.(*ast.IncDecStmt).X)
} else if reflect.TypeOf(stmt).String() == "*ast.GoStmt" {
scantupleoperation(fset,stmt.(*ast.GoStmt).Call)
} else if reflect.TypeOf(stmt).String() == "*ast.SwitchStmt " {
visitstmt(fset,stmt.(*ast.SwitchStmt).Body)
_, ok := (stmt.(*ast.SwitchStmt).Init).(*ast.BlockStmt)
if ok {
visitstmt(fset,stmt.(*ast.SwitchStmt).Init)
}
} else if reflect.TypeOf(stmt).String() == "*ast.SelectStmt" {
visitstmt(fset,stmt.(*ast.SelectStmt).Body)
} else if reflect.TypeOf(stmt).String() == "*ast.SendStmt" {
scantupleoperation(fset,stmt.(*ast.SendStmt).Chan)
} else if reflect.TypeOf(stmt).String() == "*ast.SendStmt" {
scantupleoperation(fset,stmt.(*ast.SendStmt).Value)
} else if reflect.TypeOf(stmt).String() == "*ast.TypeSwitchStmt" {
visitstmt(fset,stmt.(*ast.TypeSwitchStmt).Body)
_, ok := (stmt.(*ast.TypeSwitchStmt).Assign).(*ast.BlockStmt)
if ok {
visitstmt(fset,stmt.(*ast.TypeSwitchStmt).Assign)
}
} else if reflect.TypeOf(stmt).String() == "*ast.LabeledStmt" {
visitstmt(fset,stmt.(*ast.LabeledStmt).Stmt)
} else if reflect.TypeOf(stmt).String() == "*ast.DeferStmt" {
scantupleoperation(fset,stmt.(*ast.DeferStmt).Call)
} else if reflect.TypeOf(stmt).String() == "*ast.CommClause" {
visitstmt(fset,stmt.(*ast.CommClause).Comm)
} else if reflect.TypeOf(stmt).String() == "*ast.CaseClause" {
for _, i := range stmt.(*ast.CaseClause).List {
 scantupleoperation(fset,i)
}
} else if reflect.TypeOf(stmt).String() == "*ast.BranchStmt" {
scantupleoperation(fset,stmt.(*ast.BranchStmt).Label)
} else if reflect.TypeOf(stmt).String() == "*ast.ForStmt" {
visitstmt(fset,stmt.(*ast.ForStmt).Init)
} else if reflect.TypeOf(stmt).String() == "*ast.ForStmt" {
scantupleoperation(fset,stmt.(*ast.ForStmt).Cond)
} else if reflect.TypeOf(stmt).String() == "*ast.ForStmt" {
visitstmt(fset,stmt.(*ast.ForStmt).Post)
} else if reflect.TypeOf(stmt).String() == "*ast.RangeStmt" {
scantupleoperation(fset,stmt.(*ast.RangeStmt).Key)
} else if reflect.TypeOf(stmt).String() == "*ast.RangeStmt" {
scantupleoperation(fset,stmt.(*ast.RangeStmt).Value)
}
}

func analyse(fset *token.FileSet, node *ast.File) map[*ast.CallExpr][]int {
	//var targetlocations map[*ast.CallExpr][]int  // location targets for each put operation
	targetlocations = make(map[*ast.CallExpr][]int)

	var count int
	count = 0
	currentscope = 0

	for _, f := range node.Decls {
		// find the functions
		fn, ok := f.(*ast.FuncDecl)
		if !ok {
			continue
		}

		// If the function is not main: print the process name
		if fn.Name.Name != "main" {
			info("analyser: scanning process (%s)", fn.Name.Name)
			currentprocess = fn.Name.Name
		}

		// Traversing through the block of statements
		for _, r := range fn.Body.List {
			visitstmt(fset,r)
			count++
		}

		currentscope++
	}

	return targetlocations
}




