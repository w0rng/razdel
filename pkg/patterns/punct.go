package patterns

import "razdel/internal/set"

var (
	ENDINGS        = set.New([]string{".", "?", "!", "…"})      // Завершающие символы
	DASHES         = set.New([]string{"‑", "–", "—", "−", "-"}) // Тире
	OPEN_QUOTES    = set.New([]string{"«", "“", "‘"})           // Открывающие кавычки
	CLOSE_QUOTES   = set.New([]string{"»", "”", "’"})           // Закрывающие кавычки
	GENERIC_QUOTES = set.New([]string{"\"", "„", "'"})          // Универсальные кавычки
	OPEN_BRACKETS  = set.New([]string{"(", "[", "{"})           // Открывающие скобки
	CLOSE_BRACKETS = set.New([]string{")", "]", "}"})           // Закрывающие скобки
)

var (
	QUOTES   = OPEN_QUOTES.Add(CLOSE_QUOTES).Add(GENERIC_QUOTES) // Все кавычки
	BRACKETS = OPEN_BRACKETS.Add(CLOSE_BRACKETS)                 // Все скобки
	BOUNDS   = QUOTES.Add(BRACKETS)                              // Ограничивающие символы
	SMILES   = `[=:;]-?[)(]{1,3}`                                // Регулярное выражение для смайлов
)
