Growl for Golang
----------------
Growl support for Go. A port of [visionmedia/node-growl][].

### Installation

```
go get -U github.com/go-libs/growl
```

### Usage

```go
import "github.com/go-libs/growl"

growl.Notify("Hello Growl!", Options{
		Title:    "Growl in Golang",
		Subtitle: "go-growl",
		Url:      "https://github.com/go-libs/growl",
		Sound:    "Tink",
	})
```


[visionmedia/node-growl]: https://github.com/visionmedia/node-growl
