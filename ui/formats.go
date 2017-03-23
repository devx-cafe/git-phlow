package ui

import (
	"github.com/fatih/color"
)

var Format *colorFormat
var lsf *labelSubFormat

//colorFormat ...
//color formats for terminal output
type colorFormat struct {
	Bold      func(in string) string
	Success   func(in string) string
	Error     func(in string) string
	Label     *labelSubFormat
	Assignee  func(in string) string
	Alias     func(in string) string
	Issue     func(in string) string
	Branch    func(in string) string
	MileStone func(in string) string
}

//labelSubFormat ...
//color formats for labels in terminal output
type labelSubFormat struct {
	G1Await    func(in string) string
	G2Priority func(in string) string
	G3Time     func(in string) string
	G4Move     func(in string) string
}

func init() {
	lsf = &labelSubFormat{
		G1Await: func(in string) string {
			return color.New(color.FgHiGreen).Add(color.Bold).Add(color.BgGreen).SprintFunc()(in)
		},
		G2Priority: func(in string) string {
			return color.New(color.FgWhite).Add(color.Bold).Add(color.BgHiRed).SprintFunc()(in)
		},
		G3Time: func(in string) string {
			return color.New(color.FgHiBlack).Add(color.Bold).Add(color.BgHiBlue).SprintFunc()(in)
		},
		G4Move: func(in string) string {
			return color.New(color.FgBlack).Add(color.Bold).Add(color.BgHiWhite).SprintFunc()(in)
		},
	}

	Format = &colorFormat{
		Bold: func(in string) string {
			return color.New(color.Bold).SprintFunc()(in)
		},
		Success: func(in string) string {
			return color.New(color.FgHiGreen).SprintFunc()(in)
		},
		Error: func(in string) string {
			return color.New(color.FgHiRed).SprintFunc()(in)
		},
		Branch: func(in string) string {
			return color.New(color.FgHiGreen).Add(color.Bold).SprintFunc()(in)
		},
		Alias: func(in string) string {
			return color.New(color.FgHiCyan).Add(color.Bold).SprintFunc()(in)
		},
		Assignee: func(in string) string {
			return color.New(color.FgYellow).Add(color.Bold).SprintFunc()("@" + in)
		},
		Issue: func(in string) string {
			return color.New(color.Bold).SprintFunc()("#" + in)
		},
		MileStone: func(in string) string {
			return color.New(color.FgGreen).SprintFunc()(in)
		},
		Label: lsf,
	}
}

//FByG ...
//Format By Group return the Format of the given group of the label
func (c *colorFormat) FByG(i int) (func(string) string) {
	switch i {
	case 1:
		return c.Label.G1Await
	case 2:
		return c.Label.G2Priority
	case 3:
		return c.Label.G3Time
	case 4:
		return c.Label.G4Move
	default:
		return c.Label.G1Await
	}
}
