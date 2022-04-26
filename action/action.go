package action

import (
	"fmt"
	"import_cycles/view"
)

func Say(v *view.SomeView) {
	fmt.Println("I'm from action.go and this is what I'm saying: ", v.Name)
}
