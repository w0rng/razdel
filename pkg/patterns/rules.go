package patterns

import (
	"razdel/pkg/set"
	"regexp"
)

var SPACE_SUFFIX = regexp.MustCompile(`\s$`)
var SPACE_PREFIX = regexp.MustCompile(`^\s`)

var TOKEN = regexp.MustCompile(`([^\W\d]+|\d+|[^\w\s])`)
var FIRST_TOKEN = regexp.MustCompile(`^\s*([^\W\d]+|\d+|[^\w\s])`)
var LAST_TOKEN = regexp.MustCompile(`([^\W\d]+|\d+|[^\w\s])\s*$`)
var WORD = regexp.MustCompile(`([^\W\d]+|\d+)`)
var PAIR_SOKR = regexp.MustCompile(`(\w)\s*\.\s*(\w)\s*$`)
var INT_SOKR = regexp.MustCompile(`\d+\s*-?\s*(\w+)\s*$`)

var ROMAN = regexp.MustCompile(`^[IVXML]+$`)
var BULLET_CHARS = set.New([]string{"§", "а", "б", "в", "г", "д", "е", "a", "b", "c", "d", "e", "f"})
var BULLET_BOUNDS = ".)"
var BULLET_SIZE = 20

var DELIMITERS = ENDINGS + `;` + GENERIC_QUOTES + CLOSE_QUOTES + CLOSE_BRACKETS
var SMILE_PREFIX = regexp.MustCompile(`^\s*` + SMILES)
