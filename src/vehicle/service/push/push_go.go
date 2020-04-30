package push

import "context"

//read
func startReadPushG(fservce *PushService) {
	go ReadPushGo(fservce)
}

func ReadPushGo(fService *PushService) {
	for {
		select {
		case flowData := <-fService.PushChan:
			//发送请求
			fService.Send(flowData)
		}
	}
}

//write

func startWritePushG(fservce *PushService) {
	go WritePushGo(fservce)
}

func WritePushGo(fservice *PushService) {
	fservice.PushChan <- fservice.FlowData
	ctx, cancel := context.WithTimeout(context.Background(), fservice.WriteTimeout)
	defer cancel()
	select {
	case <-ctx.Done():
		return
	}
}
