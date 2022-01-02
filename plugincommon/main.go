package plugincommon

type CallbackFunctions interface {
	AddOne(val int64) int64
	AddMessage(message string) string
}
