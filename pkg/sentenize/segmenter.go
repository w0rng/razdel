package sentenize

import (
	"strings"
)

type Rule func(Token) bool

type Splitter interface {
	Split(string) []Token
}

type SentSegmenter struct {
	Splitter Splitter
	Rules    []Rule
}

func New() *SentSegmenter {
	return &SentSegmenter{
		Splitter: NewSplitter("", 0),
		Rules:    []Rule{DefaultRule},
	}
}

func DefaultRule(split Token) bool {
	return strings.HasPrefix(split.Right, " ")
}

// Segment выполняет сегментацию текста.
func (s *SentSegmenter) Segment(text string) []string {
	parts := s.Splitter.Split(text)
	if len(parts) == 0 {
		return nil
	}

	var segments []string
	buffer := parts[0].Left
	for i := 1; i < len(parts); i++ {
		current := parts[i]
		current.Buffer = buffer

		if s.shouldJoin(current) {
			buffer += current.Delimiter + current.Right
		} else {
			segments = append(segments, strings.TrimSpace(buffer+current.Delimiter))
			buffer = current.Right
		}
	}
	segments = append(segments, strings.TrimSpace(buffer))
	return segments
}

func (s *SentSegmenter) shouldJoin(split Token) bool {
	for _, rule := range s.Rules {
		if rule(split) {
			return true
		}
	}
	return false
}
