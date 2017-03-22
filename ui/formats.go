package ui

import (
	"github.com/fatih/color"
)

var cf *ColorFormat
var lsf *labelSubFormat

//ColorFormat ...
//color formats for terminal output
type ColorFormat struct {
	Bold      string
	Success   string
	Error     string
	Label     *labelSubFormat
	Assignee  string
	Alias     string
	Issue     string
	Branch    string
	MileStone string
}

//labelSubFormat ...
//color formats for labels in terminal output
type labelSubFormat struct {
	G1Await    string
	G2Priority string
	G3Time     string
	G4Move     string
}

//FByG ...
//Format By Group return the format of the given group of the label
func (l *labelSubFormat) FByG(g int) string {
	switch g {
	case 1:
		return l.G1Await
	case 2:
		return l.G2Priority
	case 3:
		return l.G3Time
	case 4:
		return l.G4Move
	default:
		return l.G1Await
	}
}

//Format ...
//returns the ColorFormat
func Format(input string) *ColorFormat {
	if cf == nil {
		lsf = &labelSubFormat{
			G1Await:    color.New(color.FgHiGreen).Add(color.Bold).Add(color.BgGreen).SprintFunc()(input),
			G2Priority: color.New(color.FgWhite).Add(color.Bold).Add(color.BgHiRed).SprintFunc()(input),
			G3Time:     color.New(color.FgHiBlack).Add(color.Bold).Add(color.BgHiBlue).SprintFunc()(input),
			G4Move:     color.New(color.FgBlack).Add(color.Bold).Add(color.BgHiWhite).SprintFunc()(input),
		}

		cf = &ColorFormat{
			Bold:      color.New(color.Bold).SprintFunc()(input),
			Success:   color.New(color.FgHiGreen).SprintFunc()(input),
			Error:     color.New(color.FgHiRed).SprintFunc()(input),
			Branch:    color.New(color.FgHiGreen).Add(color.Bold).SprintFunc()(input),
			Alias:     color.New(color.FgHiCyan).Add(color.Bold).SprintFunc()(input),
			Assignee:  color.New(color.FgYellow).Add(color.Bold).SprintFunc()("@" + input),
			Issue:     color.New(color.Bold).SprintFunc()("#" + input),
			MileStone: color.New(color.FgGreen).SprintFunc()(input),
			Label:     lsf,
		}
	}
	return cf
}
