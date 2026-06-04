package sqlerr

type Code string

const (
	Other                Code = "other"
	NotNullViolation     Code = "not_null_violation"
	ForeignKeyViolation  Code = "foreign_key_violation"
	UniqueViolation      Code = "unique_violation"
	CheckViolation       Code = "check_violation"
	ExcludeViolation     Code = "exclude_violation"
	TransactionFailed    Code = "transaction_failed"
	DeadlockDetected     Code = "deadlock_detected"
	TooManyConnections   Code = "too_many_connections"
)

func MapCode(code string) Code {
	switch code {
	case "23502":
		return NotNullViolation
	case "23503":
		return ForeignKeyViolation
	case "23505":
		return UniqueViolation
	case "23514":
		return CheckViolation
	case "23P01":
		return ExcludeViolation
	case "25P02":
		return TransactionFailed
	case "40P01":
		return DeadlockDetected
	case "53300":
		return TooManyConnections
	default:
		return Other
	}
}

type Severity string

const (
	SeverityError   Severity = "ERROR"
	SeverityFatal   Severity = "FATAL"
	SeverityPanic   Severity = "PANIC"
	SeverityWarning Severity = "WARNING"
	SeverityNotice  Severity = "NOTICE"
	SeverityDebug   Severity = "DEBUG"
	SeverityInfo    Severity = "INFO"
	SeverityLog     Severity = "LOG"
)

func MapSeverity(severity string) Severity {
	switch severity {
	case "ERROR":
		return SeverityError
	case "FATAL":
		return SeverityFatal
	case "PANIC":
		return SeverityPanic
	case "WARNING":
		return SeverityWarning
	case "NOTICE":
		return SeverityNotice
	case "DEBUG":
		return SeverityDebug
	case "INFO":
		return SeverityInfo
	case "LOG":
		return SeverityLog
	default:
		return SeverityError
	}
}

type Error struct {
	Code           Code
	Severity       Severity
	DatabaseCode   string
	Message        string
	SchemaName     string
	TableName      string
	ColumnName     string
	DataTypeName   string
	ConstraintName string
	driverErr      error
}

func (pe *Error) Error() string {
	return string(pe.Severity) + ": " + pe.Message + " (Code " + string(pe.Code) + ": SQLSTATE " + pe.DatabaseCode + ")"
}

func (pe *Error) Unwrap() error {
	return pe.driverErr
}
