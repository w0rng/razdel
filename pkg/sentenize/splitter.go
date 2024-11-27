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
func NewSplitter(pattern string, window int) (*SentSplitter, error) {
	if pattern == "" {
		pattern = patterns.DELIMITER
	}
	if window == 0 {
		window = 10
	}
	re, err := regexp.Compile(pattern)
	if err != nil {
		return nil, err
	}
	return &SentSplitter{
		Pattern: pattern,
		Window:  window,
		re:      re,
	}, nil
}

// Split разбивает текст на части и возвращает массив строк и SentSplit.
func (s *SentSplitter) Split(text string) []interface{} {
	matches := s.re.FindAllStringSubmatchIndex(text, -1)
	if matches == nil {
		return []interface{}{text} // Если совпадений нет, вернуть весь текст как есть.
	}

	var result []interface{}
	previous := 0

	for _, match := range matches {
		start := match[0] // Начало совпадения
		stop := match[1]  // Конец совпадения

		// Текст до разделителя
		if previous < start {
			result = append(result, text[previous:start])
		}

		// Контекст вокруг разделителя
		leftStart := max(0, start-s.Window)
		rightEnd := min(len(text), stop+s.Window)

		left := text[leftStart:start]
		right := text[stop:rightEnd]

		delimiter := ""
		if len(match) >= 4 {
			delimiter = text[match[2]:match[3]] // Первая группа, если есть
		}

		result = append(result, SentSplit{
			Left:      left,
			Delimiter: delimiter,
			Right:     right,
		})

		previous = stop
	}

	// Добавить остаток текста
	if previous < len(text) {
		result = append(result, text[previous:])
	}

	return result
}

// Вспомогательные функции для работы с границами.
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
