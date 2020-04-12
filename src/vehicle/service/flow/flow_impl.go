package flow

type FlowImpl interface {
	ReadFlow()
	WriteFlow()
	SendFlow(interface{})
}