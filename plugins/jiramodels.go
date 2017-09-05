package plugins

//JiraIssue ...
type JiraIssue struct {
	Key string `json:"key"`
	Fields struct {
		Summary string `json:"summary"`
		Status struct {
			Name string `json:"name"`
		}`json:"status"`
	}`json:"fields"`
}

//transitions ...
type Transitions struct {
	Transitions []Transition `json:"transitions"`
}

//Transition ...
type Transition struct {
	ID string `json:"id"`
	To struct {
		Name string `json:"name"`
	} `json:"to"`
}

//AssignBody ...
type AssignBody struct {
	Name string `json:"name"`
}

//TransitionBody ...
type TransitionBody struct {
	Transition struct {
		ID string `json:"id"`
	} `json:"transition"`
}
