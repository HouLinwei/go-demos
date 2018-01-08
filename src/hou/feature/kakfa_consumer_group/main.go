package main

import (
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/wvanbergen/kafka/consumergroup"
	"github.com/wvanbergen/kazoo-go"
	"log"
	"os"
	"os/signal"
	"time"
)

func init() {
	sarama.Logger = log.New(os.Stdout, "[Sarama] ", log.LstdFlags)
}

func main() {
	zkAddrStr := "localhost:2181"
	groupName := "h-test-x1"
	topic := "test1"
	cgNum := 1

	config := consumergroup.NewConfig()
	//config.Offsets.Initial = sarama.OffsetNewest
	config.Offsets.Initial = sarama.OffsetOldest
	config.Offsets.ProcessingTimeout = 10 * time.Second
	var zookeeperNodes []string
	zookeeperNodes, config.Zookeeper.Chroot = kazoo.ParseConnectionString(zkAddrStr)

	fmt.Println(zookeeperNodes)
	fmt.Println(config.Zookeeper.Chroot)

	cgs := []*consumergroup.ConsumerGroup{}

	for i := 0; i < cgNum; i++ {
		consumer, consumerErr := consumergroup.JoinConsumerGroup(groupName, []string{topic}, zookeeperNodes, config)
		if consumerErr != nil {
			log.Fatalln(consumerErr)
		}
		cgs = append(cgs, consumer)
	}

	fmt.Println(len(cgs))
	fmt.Println(cgs)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	for i := 0; i < cgNum; i++ {
		consumer := cgs[i]
		go func(consumer *consumergroup.ConsumerGroup) {

			//go func() {
			//
			//	if err := consumer.Close(); err != nil {
			//		sarama.Logger.Println("Error closing the consumer", err)
			//	}
			//}()

			go func() {
				for err := range consumer.Errors() {
					log.Println(err)
				}
			}()

			eventCount := 0
			offsets := make(map[string]map[int32]int64)

			for message := range consumer.Messages() {

				fmt.Println(string(message.Value))

				if offsets[message.Topic] == nil {
					offsets[message.Topic] = make(map[int32]int64)
				}

				eventCount += 1
				if offsets[message.Topic][message.Partition] != 0 && offsets[message.Topic][message.Partition] != message.Offset-1 {
					log.Printf("Unexpected offset on %s:%d. Expected %d, found %d, diff %d.\n", message.Topic, message.Partition, offsets[message.Topic][message.Partition]+1, message.Offset, message.Offset-offsets[message.Topic][message.Partition]+1)
				}

				// Simulate processing time
				time.Sleep(5 * time.Second)

				offsets[message.Topic][message.Partition] = message.Offset
				consumer.CommitUpto(message)
			}

			log.Printf("Processed %d events.", eventCount)
			log.Printf("%+v", offsets)
		}(consumer)
	}

	<-c

	fmt.Println("EEEEEEEEEEE")

}

//type S struct {
//
//}
//
//
//func(s *S)run(){
//	count := 1
//	for{
//		time.Sleep(time.Second*1)
//		count ++
//		fmt.Println(count)
//	}
//}
//func test() error{
//	s := S{}
//	fmt.Println("aaa")
//	go s.run()
//	fmt.Println("xxx")
//	return
//}
//
//func main(){
//
//}
