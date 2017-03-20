package plugins

//Auth ...
type Auth struct {
	Token string `json:"token"`
}

//Repo ...
type Repo struct {
	DefaultBranch string `json:"default_branch"`
}

//Issues ...
type Issues struct {
	Assignees []AssigneeIssue `json:"assignees"`
	Title     string          `json:"title"`
	Number    int             `json:"number"`
	Body      string          `json:"body"`
	Labels    []Label         `json:"labels"`
	Milestone Milestone       `json:"milestone"`
}

//Label ...
type Label struct {
	ID    int    `json:"id"`
	URL   string `json:"url"`
	Name  string `json:"name"`
	Color string `json:"color"`
}

type Milestone struct {
	Title string `json:"title"`
}

//AssigneeIssue ...
type AssigneeIssue struct {
	Login string `json:"login"`
}

//Assignee ...
type Assignee struct {
	Assignees []string `json:"assignees"`
}
