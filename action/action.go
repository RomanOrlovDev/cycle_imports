package action

import (
	"fmt"
)

type Viewer interface {
	GetName() string
}

func Say(v Viewer) {
	fmt.Println("I'm from action.go and this is what I'm saying: ", v.GetName())
}
