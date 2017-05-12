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
	gg := "h-test-x1"
	topic := "test2"

	config := consumergroup.NewConfig()
	config.Offsets.Initial = sarama.OffsetNewest
	config.Offsets.ProcessingTimeout = 10 * time.Second
	var zookeeperNodes []string
	zookeeperNodes, config.Zookeeper.Chroot = kazoo.ParseConnectionString(zkAddrStr)
	fmt.Println(zookeeperNodes)
	fmt.Println(config.Zookeeper.Chroot)

	consumer, consumerErr := consumergroup.JoinConsumerGroup(gg, []string{topic}, zookeeperNodes, config)
	if consumerErr != nil {
		log.Fatalln(consumerErr)
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		if err := consumer.Close(); err != nil {
			sarama.Logger.Println("Error closing the consumer", err)
		}
	}()

	go func() {
		for err := range consumer.Errors() {
			log.Println(err)
		}
	}()

	eventCount := 0
	offsets := make(map[string]map[int32]int64)

	for message := range consumer.Messages() {
		fmt.Println(message)
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

}
