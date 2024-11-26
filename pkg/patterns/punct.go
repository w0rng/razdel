package patterns

import "regexp"

const (
	Endings       = ".?!…"  // Завершающие символы
	Dashes        = "‑–—−-" // Тире
	OpenQuotes    = "«“‘"   // Открывающие кавычки
	CloseQuotes   = "»”’"   // Закрывающие кавычки
	GenericQuotes = "\"„'"  // Универсальные кавычки
	OpenBrackets  = "([{"   // Открывающие скобки
	CloseBrackets = ")]}"   // Закрывающие скобки
)

var (
	Quotes    = OpenQuotes + CloseQuotes + GenericQuotes // Все кавычки
	Brackets  = OpenBrackets + CloseBrackets             // Все скобки
	Bounds    = Quotes + Brackets                        // Ограничивающие символы
	SmilesReg = regexp.MustCompile(`[=:;]-?[)(]{1,3}`)   // Регулярное выражение для смайлов
)
