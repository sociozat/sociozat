package helpers

import (
	"fmt"
	"regexp"
	"strings"
)

func FormatContent(str string) string {
	str = AutoLink(str)
	str = ExtractUrl(str)
	return str
}

func AutoLink(str string) string {
	re := regexp.MustCompile(`(http|ftp|https):\/\/([\w\-_]+(?:(?:\.[\w\-_]+)+))([\w\-\.,@?^=%&amp;:/~\+#]*[\w\-\@?^=%&amp;/~\+#])?`)
	str = re.ReplaceAllString(str, `<a class="item" href="$0" target="_blank">$0 <i class="icon external alternate"></i></a>`)
	return str
}

func ExtractUrl(str string) string {
	//pattern "<sample.com>#sample.com#"
	var regex = "<(.*?)>+#(.*?)#"
	re := regexp.MustCompile(regex)
	matched := re.FindStringSubmatch(str)

	if len(matched) > 0 {
		href := matched[1]
		url := matched[2]
		text := fmt.Sprintf(`<a class="item" href="%s" target="_blank">%s <i class="icon external alternate"></i></a>`, href, url)
		//replace
		str = re.ReplaceAllString(str, text)
	}

	str = strings.ToLower(str)

	return str
}
