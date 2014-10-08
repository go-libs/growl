// +build linux
package growl

func getCmd() Command {
	var c Command
	if _, err := Which("growl"); err == nil {
		c = Command{
			Type:     "Linux-Growl",
			Pkg:      "growl",
			Msg:      "-m",
			Title:    "-title",
			Subtitle: "-subtitle",
			Host: &Host{
				Cmd:      "-H",
				Hostname: "192.168.33.1",
			},
		}
	} else {
		c = Command{
			Type:   "Linux",
			Pkg:    "notify-send",
			Msg:    "",
			Sticky: "-t 0",
			Icon:   "-i",
			Priority: &Priority{
				Cmd:   "-u",
				Range: []string{"low", "normal", "critical"},
			},
		}
	}
}
