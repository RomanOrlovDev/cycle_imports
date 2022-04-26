package action

var Register = map[string]func(v Viewer){
	"Say": Say,
}
