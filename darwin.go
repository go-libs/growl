// +build darwin
package growl

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
			Priority: Priority{
				Cmd:   "-execute",
				Range: []interface{}{},
			},
		}
	} else {
		c = Command{
			Type:   "Darwin-Growl",
			Pkg:    "growlnotify",
			Msg:    "-m",
			Sticky: "--sticky",
			Priority: Priority{
				Cmd:   "--priority",
				Range: []interface{}{-2, -1, 0, 1, 2, "Very Low", "Moderate", "Normal", "High", "Emergency"},
			},
		}
	}

	return c
}
