package growl

import (
	"testing"

	"github.com/bmizerany/assert"
)

func TestWhich(t *testing.T) {
	git, _ := Which("git")
	assert.Equal(t, "/usr/local/bin/git", git)
	Notify("Hello Growl!", Options{
		Sound:    "Tink",
		Sticky:   true,
		Title:    "Growl in Golang",
		Subtitle: "go-growl",
		Url:      "https://github.com/go-libs/growl",
	})
}
