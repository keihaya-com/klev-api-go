package klev

import (
	"strings"
)

func validate(str, pref string) error {
	prefix, suffix, found := strings.Cut(str, "_")
	if !found {
		return ErrIDInvalidFormat(str)
	}
	if prefix != pref {
		return ErrIDInvalidPrefix(str, prefix, pref)
	}
	if len(suffix) != 27 {
		return ErrIDInvalidSuffixLen(str, suffix)
	}
	for _, ch := range suffix {
		switch {
		case ch >= '0' && ch <= '9':
		case ch >= 'A' && ch <= 'Z':
		case ch >= 'a' && ch <= 'z':
		default:
			return ErrIDInvalidSuffix(str, suffix)
		}
	}
	return nil
}
