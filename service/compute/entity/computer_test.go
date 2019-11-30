package entity

import (
	shuntingYard "github.com/mgenware/go-shunting-yard"
	"reflect"
	"testing"
)

func TestNewStackComputer(t *testing.T) {
	type args struct {
		expression string
	}
	tests := []struct {
		name string
		args args
		want *StackComputer
	}{
		0: {
			name: "Test Basic Espression",
			args: args{expression: "10 / 2 * 5"},
			want: &StackComputer{exp: "10 / 2 * 5"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewStackComputer(tt.args.expression); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewStackComputer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestShuntingYard(t *testing.T) {
	type args struct {
		exp string
	}
	tests := []struct {
		name        string
		args        args
		wantPostfix string
		wantResult  int
		wantErr     bool
	}{0: {
		name:        "Basic Expression",
		args:        args{exp: "5 + 8 + 3 * 2"},
		wantPostfix: "5 8 + 3 2 * + ",
		wantResult:  19,
		wantErr:     false,
	},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotPostfix, gotResult, err := ShuntingYard(tt.args.exp)
			if (err != nil) != tt.wantErr {
				t.Errorf("ShuntingYard() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotPostfix != tt.wantPostfix {
				t.Errorf("ShuntingYard() gotPostfix = %v, want %v", gotPostfix, tt.wantPostfix)
			}
			if gotResult != tt.wantResult {
				t.Errorf("ShuntingYard() gotResult = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func TestStackComputer_ProcessInfixToPostfix(t *testing.T) {
	type fields struct {
		exp string
	}
	type args struct {
		exp string
	}
	tests := []struct {
		name        string
		fields      fields
		args        args
		wantInfix   string
		wantPostfix string
		wantResult  int64
		wantErr     bool
	}{
		0: {
			name:        "",
			fields:      fields{exp: "10 / 2 * 5"},
			args:        args{exp: "10 / 2 * 5"},
			wantInfix:   "10 / 2 * 5",
			wantPostfix: "10 2 / 5 * ",
			wantResult:  25,
			wantErr:     false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := StackComputer{
				exp: tt.fields.exp,
			}
			gotInfix, gotPostfix, gotResult, err := c.ProcessInfixToPostfix(tt.args.exp)
			if (err != nil) != tt.wantErr {
				t.Errorf("ProcessInfixToPostfix() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotInfix != tt.wantInfix {
				t.Errorf("ProcessInfixToPostfix() gotInfix = %v, want %v", gotInfix, tt.wantInfix)
			}
			if gotPostfix != tt.wantPostfix {
				t.Errorf("ProcessInfixToPostfix() gotPostfix = %v, want %v", gotPostfix, tt.wantPostfix)
			}
			if gotResult != tt.wantResult {
				t.Errorf("ProcessInfixToPostfix() gotResult = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func Test_parseDescription(t *testing.T) {
	type args struct {
		token *shuntingYard.RPNToken
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		0: {
			name: "Nil Token",
			args: args{token:shuntingYard.NewRPNOperatorToken("5")},
			want: "5 ",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseDescription(tt.args.token); got != tt.want {
				t.Errorf("parseDescription() = %v, want %v", got, tt.want)
			}
		})
	}
}
