package main

// import "github.com/docopt/docopt-go"
// import "fmt"
import "os"

var config = Config{
	dicts: []Dict{
		{
			name: "bing",
			// Bing: http://cn.bing.com/dict?q=word
			url:  "http://cn.bing.com/dict?q={word}",
			path: "",
		},
	},
}

const version = "0.3.1"
const usage = `
  Usage:
    dict <word>
    dict --help
    dict --version

  Options:
    -w=<word>    The word to translate
    -h --help    Show this screen
    -v --version Show version
`

func main() {
	// https://github.com/docopt/docopt.go/issues/26 BUG!!!
	// args, err := docopt.Parse(usage, os.Args[1:], true, version, false)
	// fmt.Println(err)
	if len(os.Args) < 2 {
		return
	}
	word := os.Args[1]
	query(word)
}
