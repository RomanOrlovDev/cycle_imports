package view

import "import_cycles/action"

type SomeView struct {
	Name string
}

func RenderView() {
	if action, exists := action.Register["Say"]; exists {
		action(&SomeView{
			Name: "Let's roll it out",
		})
	}
}
