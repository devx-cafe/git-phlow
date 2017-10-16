package ui

import (
	"github.com/fatih/color"
	"github.com/praqma/git-phlow/options"
)

//Format ...
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

//Colorizer ...
func Colorizer(in string, c *color.Color) string {
	if options.GlobalFlagNoColor {
		return in
	}
	return c.SprintFunc()(in)
}

func init() {
	lsf = &labelSubFormat{
		G1Await: func(in string) string {
			return Colorizer(in, color.New(color.FgHiGreen).Add(color.Bold).Add(color.BgGreen))
		},
		G2Priority: func(in string) string {
			return Colorizer(in, color.New(color.FgWhite).Add(color.Bold).Add(color.BgHiRed))
		},
		G3Time: func(in string) string {
			return Colorizer(in, color.New(color.FgHiBlack).Add(color.Bold).Add(color.BgHiBlue))
		},
		G4Move: func(in string) string {
			return Colorizer(in, color.New(color.FgBlack).Add(color.Bold).Add(color.BgHiWhite))
		},
	}

	Format = &colorFormat{
		Bold: func(in string) string {
			return Colorizer(in, color.New(color.Bold))
		},
		Success: func(in string) string {
			return Colorizer(in, color.New(color.FgHiGreen))
		},
		Error: func(in string) string {
			return Colorizer(in, color.New(color.FgHiRed))
		},
		Branch: func(in string) string {
			return Colorizer(in, color.New(color.FgHiGreen).Add(color.Bold))
		},
		Alias: func(in string) string {
			return Colorizer(in, color.New(color.FgHiCyan).Add(color.Bold))
		},
		Assignee: func(in string) string {
			return Colorizer("@"+in, color.New(color.FgYellow).Add(color.Bold))
		},
		Issue: func(in string) string {
			return Colorizer("#"+in, color.New(color.Bold))
		},
		MileStone: func(in string) string {
			return Colorizer(in, color.New(color.FgGreen))
		},
		Label: lsf,
	}
}

//FByG ...
//Format By Group return the Format of the given group of the label
func (c *colorFormat) FByG(i int) func(string) string {
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
