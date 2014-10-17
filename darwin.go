// +build darwin
package growl

import "os"

func getCmd() Command {
	var c Command
	if _, err := Which("terminal-notifier"); err == nil {
		c = Command{
			Type:     "Darwin-NotificationCenter",
			Pkg:      "terminal-notifier",
			Msg:      "-message",
			Title:    "-title",
			Subtitle: "-subtitle",
			Sound:    "-sound",
			Url:      "-open",
			Priority: &Priority{
				Cmd:   "-execute",
				Range: []string{},
			},
		}
	} else {
		c = Command{
			Type:   "Darwin-Growl",
			Pkg:    "growlnotify",
			Msg:    "-m",
			Sticky: "--sticky",
			Priority: &Priority{
				Cmd:   "--priority",
				Range: []string{"-2", "-1", "0", "1", "2", "Very Low", "Moderate", "Normal", "High", "Emergency"},
			},
		}
	}

	var activate string
	terminal := os.Getenv("TERM_PROGRAM")
	if terminal == "iTerm.app" {
		activate = "com.googlecode.iterm2"
	} else {
		activate = "com.apple.Terminal"
	}
	c.Activate = "-activate " + activate

	return c
}
