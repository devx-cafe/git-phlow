package plugins

//Authorization ...
//interface for authorization with external service
type Authorization func(URL, user, pass string) (token string, err error)

//Authentication ...
//interface for authentication with external service
type Authentication func(URL, user, pass string) (authenticated bool, err error)

//DefaultBranch ...
//interface for getting the default branch of the external service
type DefaultBranch func(URL, org, repo, token string) (defaultBranch string, err error)
