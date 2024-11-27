package sentenize

import (
	"razdel/internal/str"
	"razdel/pkg/patterns"
	"strings"
)

/*
    close_quote,
    close_bracket,

    dash_right,
]]
*/

func EmptySide(split Token) bool {
	return split.Left == "" || split.Right == ""
}

func NoSpacePreffix(split Token) bool {
	return !split.RightSpacePrefix()
}

func LowerRight(split Token) bool {
	return str.IsLowerAlpha(split.RightToken())
}

func DelimeterRight(split Token) bool {
	right := split.RightToken()

	if _, ok := patterns.GENERIC_QUOTES[right]; ok {
		return false
	}

	if _, ok := patterns.DELIMITERS[right]; ok {
		return true
	}

	if patterns.SMILE_PREFIX.MatchString(right) {
		return true
	}

	return false
}

func SokrLeft(split Token) bool {
	if split.Delimiter != "." {
		return false
	}

	rigth := split.RightToken()
	match := split.LeftPairSokr()

	if match != nil {
		a, b := match[0], match[1]
		left := [2]string{strings.ToLower(a), strings.ToLower(b)}

		if _, ok := patterns.HEAD_PAIR_SOKRS[left]; ok {
			return true
		}

		if _, ok := patterns.PAIR_SOKRS[left]; ok {
			if str.IsSokr(rigth) {
				return true
			}
			return false
		}
	}

	left := strings.ToLower(split.LeftToken())

	if _, ok := patterns.HEAD_SOKRS[left]; ok {
		return true
	}

	if _, ok := patterns.SOKRS[left]; ok {
		if str.IsSokr(left) {
			return true
		}
	}

	return false

}

func InsidePairSokr(split Token) bool {
	if split.Delimiter != "." {
		return false
	}

	left := strings.ToLower(split.LeftToken())
	right := strings.ToLower(split.RightToken())
	pair := [2]string{left, right}

	if _, ok := patterns.PAIR_SOKRS[pair]; ok {
		return true
	}

	return false
}

func InitialsLeft(split Token) bool {
	if split.Delimiter != "." {
		return false
	}

	left := split.LeftToken()
	if str.IsLower(left) && len(left) == 1 {
		return true
	}

	if _, ok := patterns.INITIALS[strings.ToLower(left)]; ok {
		return true
	}

	return false
}

func ListItem(split Token) bool {
	if _, ok := patterns.BULLET_BOUNDS[split.Delimiter]; !ok {
		return false
	}

	if len(split.Buffer) > patterns.BULLET_SIZE {
		return false
	}

	for _, buffer := range split.BufferTokens() {
		if !str.IsBullet(buffer) {
			return false
		}
	}

	return true
}

var RULES = []func(Token) bool{
	EmptySide,
	NoSpacePreffix,
	LowerRight,
	DelimeterRight,
	SokrLeft,
	InsidePairSokr,
	InitialsLeft,
	ListItem,
}
