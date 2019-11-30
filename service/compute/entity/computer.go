package entity

import (
	"fmt"
	"github.com/mgenware/go-shunting-yard"
	"github.com/sirupsen/logrus"
)

type Computer interface {
	ProcessInfixToPostfix(exp string) (infix string, postfix string, result int64, error error)
}

//STACK COMPUTER IMPLEMENTATION

type StackComputer struct {
	exp string
}

func NewStackComputer(expression string) *StackComputer {
	logrus.Info("STACK COMPUTER")
	return &StackComputer{exp: expression}
}

func (c StackComputer) ProcessInfixToPostfix(exp string) (infix string, postfix string, result int64, error error) {
	logrus.Info("Processing infix to postfix using Shunting Yard")
	logrus.Infof("Input expression: %s", exp)
	infix = c.exp
	postfix, intResult, err := ShuntingYard(exp)
	result = int64(intResult)
	error = err
	return
}

func ShuntingYard(exp string) (postfix string, result int, error error) {

	// parse input expression to infix notation
	infixTokens, err := shuntingYard.Scan(exp)
	if err != nil {
		panic(err)
	}
	fmt.Println("Infix Tokens:")
	fmt.Println(infixTokens)

	// convert infix notation to postfix notation(RPN)
	postfixTokens, err := shuntingYard.Parse(infixTokens)
	if err != nil {
		panic(err)
	}
	fmt.Println("Postfix(RPN) Tokens:")
	fmt.Print("[")
	for _, t := range postfixTokens {
		fmt.Printf("%v ", parseDescription(t))
		postfix += parseDescription(t)
	}
	fmt.Print("]")
	fmt.Println()

	// evaluate RPN tokens
	result, err = shuntingYard.Evaluate(postfixTokens)
	if err != nil {
		error = err
		panic(err)
	}

	// output the result
	fmt.Printf("Result: %v \n", result)
	return
}

//Auxiliary methods

// ParseDescription returns a string that describes the token.
func parseDescription(token *shuntingYard.RPNToken) string {
	return fmt.Sprintf("%v ", token.Value)
}
