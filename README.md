Growl for Golang
----------------
Growl support for Go. A port of [visionmedia/node-growl][].

View the [docs][].

### Installation

```
go get -U github.com/go-libs/growl
```

### Usage

```go
import "github.com/go-libs/growl"

growl.Notify("Hello Growl", growl.Options{
    Title:    "Growl for Golang",
    Subtitle: "go-growl",
    Url:      "https://github.com/go-libs/growl",
    Sound:    "Tink",
})
```


[visionmedia/node-growl]: https://github.com/visionmedia/node-growl
[docs]: http://godoc.org/github.com/go-libs/growl
