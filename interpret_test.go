package main

import (
	"github.com/antlr4-go/antlr/v4"
	"reflect"
	"testing"
)

func TestInterpreter_interpret(t *testing.T) {
	type fields struct {
		v         KyotoVisitor
		Variables map[string]Variable
	}
	type args struct {
		tree antlr.ParseTree
	}
	var tests []struct {
		name   string
		fields fields
		args   args
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Interpreter{
				V:         tt.fields.v,
				Variables: tt.fields.Variables,
			}
			i.Interpret(tt.args.tree)
		})
	}
}

func Test_newInterpreter(t *testing.T) {
	var tests []struct {
		name string
		want *Interpreter
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewInterpreter(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newInterpreter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_VariableAssignment(t *testing.T) {
	name := "var"
	t.Run(name, func(t *testing.T) {
		testcases := readTest(name)
		for _, tc := range testcases {
			p := Parse(tc.Code)
			i := NewInterpreter()
			r := i.Interpret(p)
			if r != tc.Expected {
				t.Errorf("%s(%s): got %d, want %d", name, tc.Name, r, tc.Expected)
			}
		}
	})
}

func Test_Empty(t *testing.T) {
	name := "empty"
	t.Run(name, func(t *testing.T) {
		testcases := readTest(name)
		for _, tc := range testcases {
			p := Parse(tc.Code)
			i := NewInterpreter()
			r := i.Interpret(p)
			if r != tc.Expected {
				t.Errorf("%s(%s): got %d, want %d", name, tc.Name, r, tc.Expected)
			}
		}
	})
}

func Test_ReturnOnly(t *testing.T) {
	name := "return_only"
	t.Run(name, func(t *testing.T) {
		testcases := readTest(name)
		for _, tc := range testcases {
			p := Parse(tc.Code)
			i := NewInterpreter()
			r := i.Interpret(p)
			if r != tc.Expected {
				t.Errorf("%s(%s): got %d, want %d", name, tc.Name, r, tc.Expected)
			}
		}
	})
}

func Test_BinaryPrecedence(t *testing.T) {
	name := "arith_bin_ops"
	t.Run(name, func(t *testing.T) {
		testcases := readTest(name)
		for _, tc := range testcases {
			p := Parse(tc.Code)
			i := NewInterpreter()
			r := i.Interpret(p)
			if r != tc.Expected {
				t.Errorf("%s(%s): got %d, want %d", name, tc.Name, r, tc.Expected)
			}
		}
	})
}

func Test_ParenthesizedArith(t *testing.T) {
	name := "parenthesized_arith"
	t.Run(name, func(t *testing.T) {
		testcases := readTest(name)
		for _, tc := range testcases {
			p := Parse(tc.Code)
			i := NewInterpreter()
			r := i.Interpret(p)
			if r != tc.Expected {
				t.Errorf("%s(%s): got %d, want %d", name, tc.Name, r, tc.Expected)
			}
		}
	})
}

func Test_UnaryArith(t *testing.T) {
	name := "arith_unary_ops"
	t.Run(name, func(t *testing.T) {
		testcases := readTest(name)
		for _, tc := range testcases {
			p := Parse(tc.Code)
			i := NewInterpreter()
			r := i.Interpret(p)
			if r != tc.Expected {
				t.Errorf("%s(%s): got %d, want %d", name, tc.Name, r, tc.Expected)
			}
		}
	})
}

func Test_ComparisonOps(t *testing.T) {
	name := "comparison_ops"
	t.Run(name, func(t *testing.T) {
		testcases := readTest(name)
		for _, tc := range testcases {
			p := Parse(tc.Code)
			i := NewInterpreter()
			r := i.Interpret(p)
			if r != tc.Expected {
				t.Errorf("%s(%s): got %d, want %d", name, tc.Name, r, tc.Expected)
			}
		}
	})
}

func Test_IfStatement(t *testing.T) {
	name := "if"
	t.Run(name, func(t *testing.T) {
		testcases := readTest(name)
		for _, tc := range testcases {
			p := Parse(tc.Code)
			i := NewInterpreter()
			r := i.Interpret(p)
			if r != tc.Expected {
				t.Errorf("%s(%s): got %d, want %d", name, tc.Name, r, tc.Expected)
			}
		}
	})
}

func Test_FnCalls(t *testing.T) {
	name := "fn_calls"
	t.Run(name, func(t *testing.T) {
		testcases := readTest(name)
		for _, tc := range testcases {
			p := Parse(tc.Code)
			i := NewInterpreter()
			r := i.Interpret(p)
			if r != tc.Expected {
				t.Errorf("%s(%s): got %d, want %d", name, tc.Name, r, tc.Expected)
			}
		}
	})
}
