package sentenize

import (
	"razdel/pkg/patterns"
	"regexp"
)

// SentSplitter выполняет разбиение текста по заданному паттерну с учетом окна.
type SentSplitter struct {
	Pattern string
	Window  int
	re      *regexp.Regexp
}

// NewSentSplitter создает новый экземпляр SentSplitter.
// Если не указаны параметры, используются значения по умолчанию.
func NewSplitter(pattern string, window int) SentSplitter {
	if pattern == "" {
		pattern = patterns.DELIMITER
	}
	if window == 0 {
		window = 10
	}
	re := regexp.MustCompile(pattern)
	return SentSplitter{
		Pattern: pattern,
		Window:  window,
		re:      re,
	}
}

// Split разбивает текст на части.
func (s SentSplitter) Split(text string) []Token {
	matches := s.re.FindAllStringIndex(text, -1)
	if matches == nil {
		return []Token{{Left: text}}
	}

	var splits []Token
	prevIndex := 0
	for _, match := range matches {
		left := text[prevIndex:match[0]]
		delimiter := text[match[0]:match[1]]
		prevIndex = match[1]

		right := ""
		if prevIndex < len(text) {
			right = text[prevIndex:]
		}

		splits = append(splits, Token{Left: left, Delimiter: delimiter, Right: right})
	}
	return splits
}
