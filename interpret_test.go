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
				t.Errorf("got %d, want %d", r, tc.Expected)
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
				t.Errorf("got %d, want %d", r, tc.Expected)
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
				t.Errorf("got %d, want %d", r, tc.Expected)
			}
		}
	})
}
