package helpers

import (
	"fmt"
	"regexp"
)

func FormatContent(str string) string {
	content := ExtractUrl(str)
	return content
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
	return str
}
