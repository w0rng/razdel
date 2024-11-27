package sentenize

import (
	"fmt"
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

	splits := make([]Token, 0, len(matches))

	prevIndex := 0
	for _, match := range matches {
		start, end := match[0], match[1]
		left := text[prevIndex:start]
		delimiter := text[start:end]
		if len(splits) > 0 {
			splits[len(splits)-1].Right = text[prevIndex:start]
		}
		prevIndex = end

		splits = append(splits, Token{Left: left, Delimiter: delimiter})
	}

	if remains := text[prevIndex:]; remains != "" {
		splits[len(splits)-1].Right = remains
		splits = append(splits, Token{Left: remains})
	}

	fmt.Println(splits)

	return splits
}
