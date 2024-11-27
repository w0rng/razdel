package sentenize

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"
)

type Rule = func(Token) bool

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
		Rules:    RULES,
	}
}

// Segment выполняет сегментацию текста.
func (s *SentSegmenter) Segment(text string) []string {
	parts := s.Splitter.Split(text)
	if len(parts) == 0 {
		return nil
	}

	var segments []string
	buffer := parts[0].Left
	for _, right := range parts {
		right.Buffer = buffer
		if s.shouldJoin(right) {
			buffer += right.Delimiter + right.Right
		} else {
			segments = append(segments, buffer+right.Delimiter)
			buffer = right.Right
		}
	}
	if remains := strings.TrimSpace(buffer); remains != "" {
		segments = append(segments, remains)
	}

	return s.postProcessing(segments)
}

func (s *SentSegmenter) shouldJoin(split Token) bool {
	for _, rule := range s.Rules {
		if rule(split) {
			fmt.Println(getFunctionName(rule), rule(split))
			return true
		}
	}
	return false
}

func (s *SentSegmenter) postProcessing(split []string) []string {
	for i, str := range split {
		split[i] = strings.TrimSpace(str)
	}

	return split
}

func getFunctionName(f interface{}) string {
	return runtime.FuncForPC(runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Entry()).Name()
}
