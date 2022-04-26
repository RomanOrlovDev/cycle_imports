package action

import "import_cycles/view"

var Register = map[string]func(v *view.SomeView){
	"Say": Say,
}
