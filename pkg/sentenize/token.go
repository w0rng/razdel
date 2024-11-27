package sentenize

import "razdel/pkg/patterns"

type SentSplit struct {
	Left      string
	Delimiter string
	Right     string
	Buffer    string
}

// Проверяет наличие пробела перед `Right`.
func (s *SentSplit) RightSpacePrefix() bool {
	return patterns.SPACE_PREFIX.MatchString(s.Right)
}

// Проверяет наличие пробела после `Left`.
func (s *SentSplit) LeftSpaceSuffix() bool {
	return patterns.SPACE_SUFFIX.MatchString(s.Left)
}

// Извлекает первый токен из `Right`.
func (s *SentSplit) RightToken() string {
	return patterns.FIRST_TOKEN.FindString(s.Right)
}

// Извлекает последний токен из `Left`.
func (s *SentSplit) LeftToken() string {
	return patterns.LAST_TOKEN.FindString(s.Left)
}

// Ищет парные сокращения в `Left`.
func (s *SentSplit) LeftPairSokr() []string {
	match := patterns.PAIR_SOKR.FindStringSubmatch(s.Left)
	if len(match) > 2 {
		return match[1:3]
	}
	return nil
}

// Находит сокращения, начинающиеся с числа, в `Left`.
func (s *SentSplit) LeftIntSokr() string {
	match := patterns.INT_SOKR.FindStringSubmatch(s.Left)
	if len(match) > 1 {
		return match[1]
	}
	return ""
}

// Извлекает первое слово из `Right`.
func (s *SentSplit) RightWord() string {
	return patterns.WORD.FindString(s.Right)
}

// Возвращает список всех токенов из `Buffer`.
func (s *SentSplit) BufferTokens() []string {
	return patterns.TOKEN.FindAllString(s.Buffer, -1)
}

// Извлекает первый токен из `Buffer`.
func (s *SentSplit) BufferFirstToken() string {
	return patterns.FIRST_TOKEN.FindString(s.Buffer)
}
