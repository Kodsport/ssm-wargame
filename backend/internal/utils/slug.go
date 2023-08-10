package utils

import (
	"regexp"
	"strings"
)

var slugRegex = regexp.MustCompile(`[^a-zA-Z\-0-9]`)

func Slugify(x string) string {
	slug := strings.NewReplacer(" ", "_", "ö", "o", "å", "a", "ä", "a").Replace(x)
	slug = slugRegex.ReplaceAllString(slug, "")
	return strings.ToLower(slug)
}
