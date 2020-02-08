package jon

import (
	"testing"
)

func Test_ImageOf(test *testing.T) {
	var want string = "https://d1ejxu6vysztl5.cloudfront.net/comics/garfield/2003/2003-04-21.gif"
	var comics []string = []string{
		"https://garfield.com/comic/2003/04/21",
		"garfield.com/comic/2003/04/21",
		"https://gocomics.com/garfield/2003/04/21",
		"gocomics.com/garfield/2003/04/21",
	}

	var current, source string
	for _, current = range comics {
		source = ImageOf(current)
		if source != want {
			test.Errorf("comic source mismatch! have: %s, want: %s", source, want)
		}
	}
}

func Test_Valid(test *testing.T) {
	var ok string = "https://garfield.com/comic/2003/04/21"
	var nok string = "u have ligma"

	if !Valid(ok) {
		test.Errorf("%s is not valid!", ok)
	}

	if Valid(nok) {
		test.Errorf("%s is valid!", nok)
	}
}

func Test_ReplyFor(test *testing.T) {
	var want string = "[image source](https://d1ejxu6vysztl5.cloudfront.net/comics/garfield/2003/2003-04-21.gif)"
	var comments []string = []string{
		"Here is my [sosig](http://garfield.com/comic/2003/04/21)",
		"I found gocomics.com/garfield/2003/04/21 and https://gocomics.com/garfield/2003/04/21",
	}

	var have, current string
	for _, current = range comments {
		have = ReplyFor(current)
		if have != want {
			test.Errorf("reply mismatch for %s! have: %s, want: %s", current, have, want)
		}
	}
}
