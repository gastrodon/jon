package jon

import (
	"fmt"
	"regexp"
	"strings"
)

// gocomics.com and garfield.com are the most common links
var site_pattern regexp.Regexp = *regexp.MustCompile(`(https?://)?(www.)?(garfield|gocomics)\.com/.+/\d+/\d+/\d+`)

const (
	// yyyy/yyyy-mm-dd
	direct_fmt string = `https://d1ejxu6vysztl5.cloudfront.net/comics/garfield/%s/%s-%s-%s.gif`
	masked_fmt string = "[image source](%s)"
)

/**
 * Get a direct link to the source of a comic
 * image:   A direct link to this comic's image
 */
func ImageOf(comic string) (image string) {
	var date []string = strings.Split(comic, "/")
	date = date[len(date)-3:]
	image = fmt.Sprintf(direct_fmt, date[0], date[0], date[1], date[2])
	return
}

/**
 * Quickly (somewhat) determine if a comment has a source
 * valid:   Does this comic have an image source?
 */
func Valid(comment string) (valid bool) {
	valid = site_pattern.FindString(comment) != ""
	return
}

func unique(it string, all []string) (is_unique bool) {
	var current string
	for _, current = range all {
		if current == it {
			is_unique = false
			return
		}
	}

	is_unique = true
	return
}

func ReplyFor(comment string) (reply string) {
	var sources []string = site_pattern.FindAllString(comment, 1+len(comment))

	var images []string = make([]string, len(sources), len(sources))
	var current, masked string
	for _, current = range sources {
		masked = fmt.Sprintf(masked_fmt, ImageOf(current))

		if unique(current, images) {
			images = append(images, masked)
		}
	}

	reply = strings.TrimSpace(strings.Join(images, "\n\n"))
	return
}
