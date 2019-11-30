package compute

import (
	"github.com/sirupsen/logrus"
	"go-kit-base/service/compute/entity"
)

type Service interface {
	ProcessInfixToPostfix(exp string) (infix string, postfix string, result int64, error error)
}

type computeService struct {
	exp string
}

//Constructor
// NewService returns a new userService
func NewService() Service {
	return &computeService{exp: ""}
}

//IMPLEMENTATIONS
//Process the requested infix notation to postfix
func (c computeService) ProcessInfixToPostfix(exp string) (infix string, postfix string, result int64, error error) {

	logrus.Info("PROCESSING INFIX TO POSTFIX NOTATION")
	computer := entity.NewStackComputer(exp)
	infix, postfix, result, error = computer.ProcessInfixToPostfix(exp)
	return
}
