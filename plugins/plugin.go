package plugins

/*
ghi label "Action - awaiting feed-back" -c 6EB82C
ghi label "Action - needs grooming"     -c 009800
ghi label "Prio 1 - must have"          -c E83D0F
ghi label "Prio 2 - should have"        -c EB6420
ghi label "Prio 3 - could have"         -c E8850F
ghi label "Prio 4 - won't have"         -c E8A80F
ghi label "Size 0 - briefing"           -c C7DEF8
ghi label "Size 1 - small"              -c 20B4E5
ghi label "Size 2 - medium"             -c 208FE5
ghi label "Size 3 - large"              -c 0052CC
ghi label "Size 4 - too big"            -c 100B6B
ghi label "Status - duplicate"          -c 111111
ghi label "Status - workable"           -c EDEDED
ghi label "Status - in progress"        -c EDEDED
ghi label "Status - up next"            -c EEEEEE
*/

var (
	labelActionAwating    = "Action - awaiting feed-back"
	labelActionGroomig    = "Action - needs grooming"
	labelPrioOne          = "Prio 1 - must have"
	labelPrioTwo          = "Prio 2 - should have"
	labelPrioThree        = "Prio 3 - could have"
	labelPrioFour         = "Prio 4 - won't have"
	labelSizeSmall        = "Size 1 - small"
	labelSizeMedium       = "Size 2 - medium"
	labelSizeLarge        = "Size 3 - large"
	labelSizeTooBig       = "Size 4 - too big"
	labelStatusDublicate  = "Status - duplicate"
	labelStatusWorkable   = "Status - workable"
	labelStatusInProgress = "Status - in progress"
	labelStatusUpNext     = "Status - up next"
)

//Plugin ...
//Interface for all plugins
type Plugin interface {
	DefaultBranch() string
	ListIssues() map[int]string
	SetAssignee(string)
	SetLabelsOnIssue(...string)
	InitializeRepo()
	IsSignedIn() bool
}
