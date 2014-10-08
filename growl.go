package growl

import (
	"log"
	"os"
	"os/exec"
	"strconv"
)

type Priority struct {
	Cmd   string
	Range []interface{}
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
	Priority Priority
}

type Options struct {
	Sound    string
	Title    string
	Subtitle string
	Url      string
	Sticky   bool
	Priority int
	Image    string
}

func quote(s string) string {
	return strconv.Quote(s)
}

func Which(appName string) (string, error) {
	return exec.LookPath(appName)
}

func Notify(msg string, opts Options) {
	c := getCmd()
	args := []string{}

	if opts.Sticky {
		args = append(args, c.Sticky)
	}

	switch c.Type {
	case "Darwin-NotificationCenter":
		args = append(args, c.Msg)
		args = append(args, quote(msg))
		if opts.Title != "" {
			args = append(args, c.Title)
			args = append(args, quote(opts.Title))
		}
		if opts.Subtitle != "" {
			args = append(args, c.Subtitle)
			args = append(args, quote(opts.Subtitle))
		}
		if opts.Url != "" {
			args = append(args, c.Url)
			args = append(args, quote(opts.Url))
		}
		if opts.Sound != "" {
			args = append(args, c.Sound, opts.Sound)
		}
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
	}

	cmd := exec.Command(c.Pkg, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}
