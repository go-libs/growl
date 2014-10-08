// +build windows
package growl

func getCmd() Command {
	return Command{
		Type:   "Windows",
		Pkg:    "growlnotify",
		Msg:    "",
		Sticky: "/s:true",
		Title:  "/t:",
		Icon:   "/i:",
		Url:    "/cu:",
		Priority: &Priority{
			Cmd:   "/p:",
			Range: []string{"-2", "-1", "0", "1", "2"},
		},
	}
}
