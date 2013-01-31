package main

import (
	"fmt"
)

func KeyListItem(key string) string {
	return fmt.Sprintf("<tr><td><a href=\"#\" class=\"key-link\" data-key=\"%v\">%v</a></td></tr>", KeyToHtmlId(key), key)
}
