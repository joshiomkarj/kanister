package controller

import (
	"github.com/kanisterio/kanister/pkg/controllerv2/pkg/controller/actionset"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, actionset.Add)
}
