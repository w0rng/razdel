package patterns

const (
	ENDINGS        = ".?!…"  // Завершающие символы
	DASHES         = "‑–—−-" // Тире
	OPEN_QUOTES    = "«“‘"   // Открывающие кавычки
	CLOSE_QUOTES   = "»”’"   // Закрывающие кавычки
	GENERIC_QUOTES = "\"„'"  // Универсальные кавычки
	OPEN_BRACKETS  = "([{"   // Открывающие скобки
	CLOSE_BRACKETS = ")]}"   // Закрывающие скобки
)

var (
	QUOTES   = OPEN_QUOTES + CLOSE_QUOTES + GENERIC_QUOTES // Все кавычки
	BRACKETS = OPEN_BRACKETS + CLOSE_BRACKETS              // Все скобки
	BOUNDS   = QUOTES + BRACKETS                           // Ограничивающие символы
	SMILES   = `[=:;]-?[)(]{1,3}`                          // Регулярное выражение для смайлов
)
