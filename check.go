// Copyright (c) 2017, Daniel Martí <mvdan@mvdan.cc>
// See LICENSE for licensing information

package main

import (
	"flag"
	"fmt"
	"go/ast"
	"go/token"
	"os"
	"sort"
	"strings"

	"github.com/kisielk/gotool"
	"github.com/mvdan/lint"
	"golang.org/x/tools/go/loader"
)

var (
	tests    = flag.Bool("tests", true, "include tests")
	treshold = flag.Float64("r", 2.0, "ratio treshold (big/small)")
)

func main() {
	flag.Parse()
	lines, err := Unindent(*tests, flag.Args()...)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	for _, line := range lines {
		fmt.Println(line)
	}
}

func Unindent(tests bool, args ...string) ([]string, error) {
	paths := gotool.ImportPaths(args)
	var conf loader.Config
	conf.TypeCheckFuncBodies = func(path string) bool {
		return false
	}
	if _, err := conf.FromArgs(paths, tests); err != nil {
		return nil, err
	}
	lprog, err := conf.Load()
	if err != nil {
		return nil, err
	}
	c := &Checker{lprog: lprog}
	if c.wd, err = os.Getwd(); err != nil {
		return nil, err
	}
	return c.lines(tests, args...)
}

func (c *Checker) lines(tests bool, args ...string) ([]string, error) {
	issues, err := c.Check()
	if err != nil {
		return nil, err
	}
	lines := make([]string, len(issues))
	for i, issue := range issues {
		fpos := c.lprog.Fset.Position(issue.Pos()).String()
		if strings.HasPrefix(fpos, c.wd) {
			fpos = fpos[len(c.wd)+1:]
		}
		lines[i] = fmt.Sprintf("%s: %s", fpos, issue.Message())
	}
	return lines, nil
}

func (c *Checker) Check() ([]lint.Issue, error) {
	for _, info := range c.lprog.InitialPackages() {
		for _, file := range info.Files {
			ast.Inspect(file, c.walk)
		}
	}
	// TODO: replace by sort.Slice once we drop Go 1.7 support
	sort.Sort(byNamePos{c.lprog.Fset, c.issues})
	return c.issues, nil
}

type byNamePos struct {
	fset *token.FileSet
	l    []lint.Issue
}

func (p byNamePos) Len() int      { return len(p.l) }
func (p byNamePos) Swap(i, j int) { p.l[i], p.l[j] = p.l[j], p.l[i] }
func (p byNamePos) Less(i, j int) bool {
	p1 := p.fset.Position(p.l[i].Pos())
	p2 := p.fset.Position(p.l[j].Pos())
	if p1.Filename == p2.Filename {
		return p1.Offset < p2.Offset
	}
	return p1.Filename < p2.Filename
}

type Checker struct {
	lprog  *loader.Program
	issues []lint.Issue

	wd string
}

type Issue struct {
	pos token.Pos
	msg string
}

func (i Issue) Pos() token.Pos  { return i.pos }
func (i Issue) Message() string { return i.msg }

func (c *Checker) walk(node ast.Node) bool {
	var bl *ast.BlockStmt
	// we can only return/break/continue out of these, not out of
	// e.g. IfStmt
	switch x := node.(type) {
	case *ast.FuncDecl:
		bl = x.Body
	case *ast.FuncLit:
		bl = x.Body
	case *ast.ForStmt:
		bl = x.Body
	}
	if bl == nil {
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
		if score < *treshold {
			continue // reversing if would not be worth it
		}
		c.issues = append(c.issues, Issue{
			pos: ifs.Pos(),
			msg: fmt.Sprintf("%d stmts inside, %d after (score %.2f)",
				inside, after, score),
		})
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
