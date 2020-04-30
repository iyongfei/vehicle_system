package push

type PushImpl interface {
	Read()
	Write()
	Send(interface{})
}
