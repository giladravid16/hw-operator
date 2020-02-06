package controller

import (
	"github.com/hw-operator/pkg/controller/giladhelloworld"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, giladhelloworld.Add)
}
