package godecode2

import (
	"regexp"
	"strings"
)

// Transliterate an Unicode object into an ASCII string.
func Decode(str string) string {
	var ret []string
	for _, codepoint := range str {
		// Basic ASCII
		if codepoint < 0x80 {
			ret = append(ret, string(codepoint))
			continue
		}
		// Characters in Private Use Area and above are ignored
		if codepoint > 0xffff {
			continue
		}
		section := codepoint >> 8   // Chop off the last two hex digits
		position := codepoint % 256 // Last two hex digits
		table := cache[section]
		if table != nil && rune(len(table)) > position {
			ret = append(ret, table[position])
		}
	}
	return strings.Trim(strings.Join(ret, ""), " ")
}

// Transliterate Unicode string to a initials.
func Initials(str string) string {
	var ret []string
	reg := regexp.MustCompile("^\\w|\\s+\\w")
	for _, s := range reg.FindAllString(Decode(str), -1) {
		ret = append(ret, strings.Replace(s, " ", "", -1))
	}
	return strings.Join(ret, "")
}
