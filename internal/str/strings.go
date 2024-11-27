package str

import (
	"razdel/pkg/patterns"
	"strings"
	"unicode"
)

func IsDigit(token string) bool {
	for _, letter := range token {
		if !unicode.IsDigit(letter) {
			return false
		}
	}
	return true
}

func IsAlpha(token string) bool {
	for _, letter := range token {
		if !unicode.IsLetter(letter) {
			return false
		}
	}
	return true
}

func IsLower(token string) bool {
	for _, letter := range token {
		if !unicode.IsLower(letter) {
			return false
		}
	}
	return true
}

func IsUpper(token string) bool {
	for _, letter := range token {
		if !unicode.IsUpper(letter) {
			return false
		}
	}
	return true
}

func IsSokr(token string) bool {
	if IsDigit(token) {
		return true
	}
	if !IsAlpha(token) {
		return true
	}

	return IsLower(token)
}

func IsLowerAlpha(token string) bool {
	return IsAlpha(token) && IsLower(token)
}

func IsBullet(token string) bool {
	if IsDigit(token) {
		return true
	}
	if _, ok := patterns.BULLET_BOUNDS[token]; ok {
		return true
	}
	if _, ok := patterns.BULLET_CHARS[strings.ToLower(token)]; ok {
		return true
	}

	if patterns.ROMAN.MatchString(token) {
		return true
	}

	return false
}
