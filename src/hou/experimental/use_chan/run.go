package use_chan

import (
	"fmt"
	"time"
)

type Msg struct{
	Data string
}


var topics []chan *Msg


func RunDemo(){
	topics = make([]chan *Msg,2)
	topics[0] = start()
	topics[1] = start()

	go func() {
		for i:=0; i<100; i++{
			if i%2 == 0{
				topics[0] <- &Msg{
					Data:fmt.Sprintf("%d",i),
				}
			}else{
				topics[1] <- &Msg{
					Data:fmt.Sprintf("%d",i),
				}
			}
		}
	}()

	time.Sleep(time.Second*5)
}


func start()chan *Msg{
	tChan := make(chan *Msg, 10)
	go func(sourceChan <- chan *Msg) {
		for{
			select {
			case m := <- sourceChan:
				fmt.Println(m.Data)
			}
		}
	}(tChan)
	return tChan
}