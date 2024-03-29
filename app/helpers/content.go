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
	var pattern = `(\[.*\])(\((http)(?:s)?(\:\/\/).*\))`
	r := regexp.MustCompile(pattern)
	matched := r.FindAllStringSubmatch(str, -1)
	if len(matched) > 0 {
		for i, _ := range matched {
			href := matched[i][2]
			url := matched[i][1]
			text := fmt.Sprintf(`<a class="item" href="%s" target="_blank">%s <i class="icon external alternate"></i></a>`, href, url)

			//replace
			str = strings.Replace(str, matched[i][0], text, len(matched))
		}
	}

	str = strings.ToLower(str)

	return str
}

func Post(str string) string {
	//pattern "#41234"
	var pattern = `\W(\#[0-9]*[1-9]+\b)`
	r := regexp.MustCompile(pattern)
	matched := r.FindAllStringSubmatch(str, -1)

	if len(matched) > 0 {
		for i, _ := range matched {
			href := matched[i][1]
			url := fmt.Sprintf("/p/%s", url.QueryEscape(href[1:]))
			text := fmt.Sprintf(`<a class="item" href="%s">%s </a>`, url, href)
			//replace
			str = strings.Replace(str, matched[i][0], text, len(matched))
		}
	}

	str = strings.ToLower(str)

	return str
}

func Topic(str string) string {
	//pattern "#spoiler#"
	var pattern = `#(.*?)#`
	r := regexp.MustCompile(pattern)
	matched := r.FindAllStringSubmatch(str, -1)

	if len(matched) > 0 {
		for i, _ := range matched {
			href := matched[i][1]
			url := fmt.Sprintf("/q/%s", url.QueryEscape(href))
			text := fmt.Sprintf(`<a class="item" href="%s">%s </a>`, url, href)
			//replace
			str = strings.Replace(str, matched[i][0], text, len(matched))
		}
	}

	str = strings.ToLower(str)

	return str
}

func Notation(str string) string {
	//pattern "#sample.com#"
	var pattern = `[*](.*)[*]`
	r := regexp.MustCompile(pattern)
	matched := r.FindAllStringSubmatch(str, -1)
	if len(matched) > 0 {
		for i, _ := range matched {
			text := matched[i][1]
			url := fmt.Sprintf("/q/%s", url.QueryEscape(text))
			output := fmt.Sprintf(`<a class="item popup" data-content="%s" href="%s"><b>*</b></a>`, text, url)

			//replace
			str = strings.Replace(str, matched[i][0], output, len(matched))
		}
	}

	str = strings.ToLower(str)

	return str
}
