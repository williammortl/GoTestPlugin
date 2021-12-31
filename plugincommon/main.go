package plugincommon

type CallbackFunctions interface {
	AddOne(val int) int
	AddMessage(message string) string
}
