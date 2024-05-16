package helpers

import (
	"fmt"
	"net/url"
	"regexp"
	"strings"
)

//todo refactor all funcs with one pattern/output struct
func FormatContent(str string) string {
	// str = AutoLink(str)
	str = ExtractUrl(str)
	str = Topic(str)
	str = Notation(str)
	str = Post(str)
	return str
}

func ExtractUrl(str string) string {
	//pattern "[Link Text](https://url.com)"
	var pattern = `(\[.*?\])\((http(?:s)?://.*?)\)`
    r := regexp.MustCompile(pattern)
    matched := r.FindAllStringSubmatch(str, -1)

    if len(matched) > 0 {
        for _, match := range matched {
            linkText := match[1][1:len(match[1])-1] // strip the square brackets
            url := match[2] // capture group for the URL
            // format the HTML anchor tag
            text := fmt.Sprintf(`<a class="item" href="%s" target="_blank">%s <i class="icon external alternate"></i></a>`, url, linkText)
            // replace the Markdown link with the HTML anchor tag in the original string
            str = strings.Replace(str, match[0], text, 1)
        }
    }

    return str
}

func Post(str string) string {
// pattern to match hashtags like #41234
	var pattern = `(\B#[1-9][0-9]*)`
	r := regexp.MustCompile(pattern)
	matched := r.FindAllStringSubmatch(str, -1)
	if len(matched) > 0 {
		for _, match := range matched {
			href := match[1]
			url := fmt.Sprintf("/p/%s", url.QueryEscape(href[1:]))
			text := fmt.Sprintf(`<a class="item" href="%s">%s</a>`, url, href)
			// replace the matched hashtag with the HTML anchor tag
			str = strings.Replace(str, match[0], text, -1)
		}
	}

	// convert the final string to lowercase
	str = strings.ToLower(str)

	return str
}

func Topic(str string) string {
	//pattern "#spoiler#"
	var pattern = `#(.*?)#`
	r := regexp.MustCompile(pattern)
	matched := r.FindAllStringSubmatch(str, -1)

    if len(matched) > 0 {
		for _, match := range matched {
			href := match[1]
			url := fmt.Sprintf("/q/%s", url.QueryEscape(href))
			text := fmt.Sprintf(`<a class="item" href="%s">%s</a>`, url, href)
			// replace the matched topic with the HTML anchor tag
			str = strings.Replace(str, match[0], text, 1)
		}
	}

	str = strings.ToLower(str)

	return str
}

func Notation(str string) string {
	// pattern to match *sample.com*
	var pattern = `\*(.*?)\*`
	r := regexp.MustCompile(pattern)
	matched := r.FindAllStringSubmatch(str, -1)
	if len(matched) > 0 {
		for _, match := range matched {
			text := match[1]
			escapedText := url.QueryEscape(text)
			output := fmt.Sprintf(`<a class="item popup" data-content="%s" href="/q/%s"><b>*</b></a>`, text, escapedText)

			// replace the matched notation with the HTML anchor tag
			str = strings.Replace(str, match[0], output, 1)
		}
	}

	// convert the final string to lowercase
	str = strings.ToLower(str)

	return str
}
