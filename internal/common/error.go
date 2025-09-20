package common

func (e *ErrorInfo) Error() string {
	return e.Message
}

// Códigos de error constantes
const (
	// Códigos de validación
	ErrCodeValidation    = "VALIDATION_ERROR"
	ErrCodeRequired      = "REQUIRED_FIELD"
	ErrCodeInvalidFormat = "INVALID_FORMAT"
	ErrCodeInvalidValue  = "INVALID_VALUE"

	// Códigos de base de datos
	ErrCodeDatabaseError = "DATABASE_ERROR"
	ErrCodeNotFound      = "RESOURCE_NOT_FOUND"
	ErrCodeDuplicate     = "DUPLICATE_RESOURCE"
	ErrCodeConstraint    = "CONSTRAINT_VIOLATION"

	// Códigos de autenticación
	ErrCodeUnauthorized = "UNAUTHORIZED"
	ErrCodeForbidden    = "FORBIDDEN"
	ErrCodeInvalidToken = "INVALID_TOKEN"
	ErrCodeExpiredToken = "EXPIRED_TOKEN"

	// Códigos de servidor
	ErrCodeInternalError = "INTERNAL_SERVER_ERROR"
	ErrCodeServiceError  = "SERVICE_UNAVAILABLE"
	ErrCodeTimeout       = "REQUEST_TIMEOUT"

	// Códigos de cliente
	ErrCodeBadRequest   = "BAD_REQUEST"
	ErrCodeInvalidJSON  = "INVALID_JSON"
	ErrCodeMissingParam = "MISSING_PARAMETER"
)
