package flow

import "context"



//read
func startReadFlowG(fservce *FlowService)  {
	go ReadFlowGo(fservce)
}


func ReadFlowGo(fService *FlowService)  {
	for {
		select {
		case flowData := <-fService.FlowChan:
			//发送请求
			fService.SendFlow(flowData)
		}
	}
}

//write

func startWriteFlowG(fservce *FlowService)  {
	go WriteFlowGo(fservce)
}


func WriteFlowGo(fservice *FlowService)  {
	fservice.FlowChan <- fservice.FlowData
	ctx, cancel := context.WithTimeout(context.Background(), fservice.WriteTimeout)
	defer cancel()
	select {
	case <-ctx.Done():
		return
	}
}