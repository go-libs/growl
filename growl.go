package growl

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strconv"
)

type Priority struct {
	Cmd   string
	Range []string
}

type Host struct {
	Cmd      string
	Hostname string
}

type Command struct {
	Type     string
	Pkg      string
	Msg      string
	Title    string
	Subtitle string
	Sound    string
	Sticky   string
	Icon     string
	Url      string
	Priority *Priority
	Host     *Host
	Activate string
}

type Options struct {
	Name     string
	Sound    string
	Title    string
	Subtitle string
	Url      string
	Sticky   bool
	Priority string
	Image    string
	Exec     string
}

func quote(s string) string {
	return strconv.Quote(s)
}

func Which(appName string) (string, error) {
	return exec.LookPath(appName)
}

func Notify(msg string, opts Options) {
	var c Command
	if opts.Exec != "" {
		c = Command{
			Type: "Custom",
			Pkg:  opts.Exec,
		}
	} else {
		c = getCmd()
	}
	args := []string{}

	if opts.Sticky {
		args = append(args, c.Sticky)
	}

	if opts.Priority != "" && c.Priority != nil {
		priority := opts.Priority
		for _, v := range c.Priority.Range {
			if priority == v {
				args = append(args, c.Priority.Cmd, priority)
				break
			}
		}
	}

	switch c.Type {
	case "Darwin-Growl":
		if opts.Image != "" {
			flag, ext := "", filepath.Ext(opts.Image)[1:]
			if ext == "icns" {
				flag = "iconpath"
			}

			if flag == "" {
				if matched, _ := regexp.MatchString("^[A-Z]", opts.Image); matched {
					flag = "appIcon"
				}
			}

			if flag == "" {
				if matched, _ := regexp.MatchString("^png|gif|jpe?g$", ext); matched {
					flag = "image"
				}
			}
			if flag == "" {
				if ext != "" && opts.Image == ext {
					flag = "icon"
				}
			}
			if flag == "" {
				flag = "icon"
			}
			args = append(args, "--"+flag, quote(opts.Image))
		}
		args = append(args, c.Msg, quote(msg))
		if opts.Title != "" {
			args = append(args, quote(opts.Title))
		}
		if opts.Name != "" {
			args = append(args, "--name", opts.Name)
		}
		break

	case "Darwin-NotificationCenter":
		args = append(args, c.Msg, quote(msg))
		if opts.Title != "" {
			args = append(args, c.Title, quote(opts.Title))
		}
		if opts.Subtitle != "" {
			args = append(args, c.Subtitle, quote(opts.Subtitle))
		}
		if opts.Url != "" {
			args = append(args, c.Url, quote(opts.Url))
		}
		if opts.Sound != "" {
			args = append(args, c.Sound, opts.Sound)
		}
		break

	case "Linux":
		if opts.Image != "" {
			args = append(args, c.Icon, opts.Image)
		}
		// libnotify defaults to sticky, set a hint for transient notifications
		if opts.Sticky != true {
			args = append(args, "--hint=int:transient:1")
		}
		if opts.Title != "" {
			args = append(args, quote(opts.Title), c.Msg, quote(msg))
		} else {
			args = append(args, quote(msg))
		}
		break

	case "Linux-Growl":
		args = append(args, c.Msg, quote(msg))
		if opts.Title != "" {
			args = append(args, quote(opts.Title))
		}
		args = append(args, c.Host.Cmd, c.Host.Hostname)
		break

	case "Windows":
		args = append(args, quote(msg))
		if opts.Image != "" {
			args = append(args, c.Icon+quote(opts.Image))
		}
		if opts.Title != "" {
			args = append(args, c.Title+quote(opts.Title))
		}
		if opts.Url != "" {
			args = append(args, c.Subtitle+quote(opts.Url))
		}
		break

	case "Custom":
		if opts.Title != "" {
			msg = opts.Title + ": " + msg
		}
		args = append(args, quote(msg))
		break
	}

	if c.Activate != "" {
		args = append(args, c.Activate)
	}
	log.Println(args)

	cmd := exec.Command(c.Pkg, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}
