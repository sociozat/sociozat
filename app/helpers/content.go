package helpers

import (
	"fmt"
	"regexp"
	"strings"
)

func FormatContent(str string) string {
	// str = AutoLink(str)
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
	matched := re.FindAllStringSubmatch(str, -1)
	if len(matched) > 0 {
		for i, _ := range matched {
			href := matched[i][1]
			url := matched[i][2]
			text := fmt.Sprintf(`<a class="item" href="%s" target="_blank">%s <i class="icon external alternate"></i></a>`, href, url)

			//replace
			str = strings.Replace(str, matched[i][0], text, len(matched))
		}
	}

	str = strings.ToLower(str)

	return str
}
