package enum

// TestResultCode представляет собой код результата тестирования
// @Model domain.TestResultCode
// @Description TL - Превышено ограничение по времени,
// @Description ML - Превышено ограничение по памяти,
// @Description CE - Ошибка компиляции,
// @Description RE - Ошибка во время выполнения,
// @Description SC - Успешное выполнение,
// @Description IA - Неверный ответ
type TestResultCode string

const (
	// TimeLimitCode @enum
	TimeLimitCode TestResultCode = "TL"

	// MemoryLimitCode @enum
	MemoryLimitCode TestResultCode = "ML"

	// CompileErrorCode @enum
	CompileErrorCode TestResultCode = "CE"

	// RuntimeErrorCode @enum
	RuntimeErrorCode TestResultCode = "RE"

	// SuccessCode @enum
	SuccessCode TestResultCode = "SC"

	// IncorrectAnswerCode @enum
	IncorrectAnswerCode TestResultCode = "IA"
)
