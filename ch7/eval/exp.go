package main

import (
	"fmt"
	"math"
	"testing"
  //"go/parser"
)

type Expr interface {
	//returns the value of the expression in a given environment
	Eval(env Env) float64
}

type Var string

func (v Var) Eval(env Env) float64 { return env[v] }

type literal float64

func (l literal) Eval(env Env) float64 { return float64(l) }

//unary operatos such as +x
type unary struct {
	op rune
	x  Expr
}

func (u unary) Eval(env Env) float64 {
	switch u.op {
	case '+':
		return +u.x.Eval(env)
	case '-':
		return -u.x.Eval(env)
	}
	panic(fmt.Sprintf("unsupported unary operator: %q", u.op))
}

//binary operators such as x + y
type binary struct {
	op   rune
	x, y Expr
}

func (b binary) Eval(env Env) float64 {
	switch b.op {
	case '+':
		return b.x.Eval(env) + b.y.Eval(env)
	case '-':
		return b.x.Eval(env) - b.y.Eval(env)
	case '*':
		return b.x.Eval(env) * b.y.Eval(env)
	case '/':
		return b.x.Eval(env) / b.y.Eval(env)
	}
	panic(fmt.Sprintf("unsupporten binary operator: %q", b.op))
}

//function calls
type call struct {
	fn   string
	args []Expr
}

func (c call) Eval(env Env) float64 {
	switch c.fn {
	case "pow":
		return math.Pow(c.args[0].Eval(env), c.args[1].Eval(env))
	case "sin":
		return math.Sin(c.args[0].Eval(env))
	case "sqrt":
		return math.Sqrt(c.args[0].Eval(env))
	}
	panic(fmt.Sprintf("unsupported function call: %s", c.fn))
}

//mapping variables to values
type Env map[Var]float64

//TEST
func TestEval(t *testing.T) {
	tests := []struct {
		expr string
		env  Env
		want string
	}{
		{"sqrt(A/pi)", Env{"A": 87616, "pi": math.Pi}, "167"},
		{"pow(x,3)+pow(y,3)", Env{"x": 12, "y": 1}, "1729"},
		{"pow(x,3)+pow(y,3)", Env{"x": 9, "y": 10}, "1729"},
		{"5/9*(F-32)", Env{"F": -40}, "-40"},
		{"5/9*(F-32)", Env{"F": 32}, "0"},
		{"5/9*(F-32)", Env{"F": 212}, "100"},
	}
	var prevExpr string
	for _, test := range tests {
		//print expr only when it changes
		if test.expr != prevExpr {
			fmt.Printf("\ns\n", test.expr)
			prevExpr = test.expr
		}
		expr, err := Parse(test.expr)
		if err != nil {
			t.Error(err) // parse Error
			continue
		}
		got := fmt.Sprintf("%.6g", string(expr).Eval(test.env))
		fmt.Printf("\t%v => %s\n", test.env, got)
		if got != test.want {
			t.Errorf("%s.Eval() in %s = %q\n, want %q\n", test.expr, test.env, got, test.want)
		}
	}
}

