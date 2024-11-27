package main

import (
	"fmt"
	"razdel/pkg/sentenize"
)

func main() {
	segmenter := sentenize.New()
	text := `
	"Так в чем же дело?" - "Не ра-ду-ют".
И т. д. и т. п. В общем, вся газета`

	segments := segmenter.Segment(text)
	for i, segment := range segments {
		fmt.Println(i, segment)
	}
}
