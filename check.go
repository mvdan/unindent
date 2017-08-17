// Copyright (c) 2017, Daniel Martí <mvdan@mvdan.cc>
// See LICENSE for licensing information

package main

import (
	"flag"
	"fmt"
	"go/ast"
	"go/token"
	"os"

	"github.com/kisielk/gotool"
	"golang.org/x/tools/go/loader"
)

func main() {
	if err := check(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func check() error {
	flag.Parse()
	paths := gotool.ImportPaths(flag.Args())
	var conf loader.Config
	conf.TypeCheckFuncBodies = func(path string) bool {
		return false
	}
	if _, err := conf.FromArgs(paths, false); err != nil {
		return err
	}
	lprog, err := conf.Load()
	if err != nil {
		return err
	}
	c := &Checker{lprog: lprog}
	for _, info := range lprog.InitialPackages() {
		for _, file := range info.Files {
			ast.Inspect(file, c.walk)
		}
	}
	return nil
}

type Checker struct {
	lprog *loader.Program
}

func (c *Checker) walk(node ast.Node) bool {
	bl, ok := node.(*ast.BlockStmt)
	if !ok {
		return true
	}
	for i, stmt := range bl.List {
		ifs, ok := stmt.(*ast.IfStmt)
		if !ok {
			continue
		}
		if ifs.Init != nil || ifs.Else != nil {
			continue // too complex
		}
		body := ifs.Body.List
		inside := countStmts(body...)
		after := countStmts(bl.List[i+1:]...)
		if after > 0 { // we need the if body to terminate
			if len(body) < 1 {
				continue // non-terminating
			}
			switch x := body[len(body)-1].(type) {
			case *ast.ReturnStmt:
			case *ast.BranchStmt:
				if x.Tok != token.BREAK && x.Tok != token.CONTINUE {
					continue // too complex
				}
			default:
				continue // non-terminating
			}
			inside-- // don't count terminating stmt
		}
		if inside < 1 {
			continue // empty
		}
		// add N (5) to after so that zero values like 3/0 don't
		// divide by 0, and small ones like 5/1 have a less
		// dramatic ratio
		score := float64(inside) / float64(after+5)
		if score < 4.0 {
			continue // reversing if would not be worth it
		}
		pos := c.lprog.Fset.Position(ifs.Pos())
		fmt.Printf("%v: %d stmts inside, %d after (score %.2f)\n",
			pos, inside, after, score)
	}
	return true
}

func countStmts(stmts ...ast.Stmt) int {
	count := 0
	for _, stmt := range stmts {
		ast.Inspect(stmt, func(node ast.Node) bool {
			if _, ok := node.(ast.Stmt); ok {
				count++
			}
			return true
		})
	}
	return count
}
