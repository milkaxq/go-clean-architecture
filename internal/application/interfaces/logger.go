package interfaces

type Logger interface {
	Info(msg string, fields map[string]interface{})
	Error(msg string, fields map[string]interface{})
	Debug(msg string, fields map[string]interface{})
	Warning(msg string, fields map[string]interface{})
	Fatal(msg string, fields map[string]interface{})
}
