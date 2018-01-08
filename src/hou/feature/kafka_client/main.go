package main

import (
	"fmt"
	"github.com/Shopify/sarama"
	"log"
	"os"
	"os/signal"
	"sync"
)

func Produce() {
	producer, err := sarama.NewSyncProducer([]string{"localhost:9092"}, nil)
	if err != nil {
		log.Fatalln(err)
	}
	defer func() {
		if err := producer.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	msg := &sarama.ProducerMessage{Topic: "houlinwei.test2", Value: sarama.StringEncoder("toda!123")}
	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		log.Printf("FAILED to send message: %s\n", err)
	} else {
		log.Printf("> message sent to partition %d at offset %d\n", partition, offset)
	}
}

func Consume() {
	consumer, err := sarama.NewConsumer([]string{"localhost:9092"}, nil)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := consumer.Close(); err != nil {
			log.Fatalln(err)
		}
	}()
	partitions, err := consumer.Partitions("houlinwei.test2")
	fmt.Println("partitions: ", partitions)
	partitionConsumer, err := consumer.ConsumePartition("houlinwei.test2", 0, 0)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := partitionConsumer.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	// Trap SIGINT to trigger a shutdown.
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	consumed := 0
ConsumerLoop:
	for {
		select {
		case msg := <-partitionConsumer.Messages():
			fmt.Println("[msg]", msg.Value)
			log.Printf("Consumed message offset %d\n", msg.Offset)
			consumed++
		case <-signals:
			break ConsumerLoop
		}
	}

	log.Printf("Consumed: %d\n", consumed)
}

// low level consumer.
func Consumer2() {
	topic := "test"
	GName := "gtest"
	Addr := []string{"localhost:9092"}
	config := sarama.NewConfig()
	config.ClientID = "clien-x1"
	client, err := sarama.NewClient(Addr, config)
	client.Config()
	if err != nil {
		fmt.Println(err)
	}

	Consumer, err := sarama.NewConsumerFromClient(client)
	if err != nil {
		fmt.Println(err)
	}

	// Get Partitions
	ps, _ := client.Partitions(topic)

	OffsetManager, err := sarama.NewOffsetManagerFromClient(GName, client)

	var pcs []sarama.PartitionConsumer
	var pcm []sarama.PartitionOffsetManager
	for _, ps := range ps {
		fmt.Println("partition idx: ", ps)
		PartitionManager, _ := OffsetManager.ManagePartition(topic, ps)
		pcm = append(pcm, PartitionManager)
		offset, _ := PartitionManager.NextOffset()
		fmt.Println("offset: ", offset)
		pConsumer, _ := Consumer.ConsumePartition(topic, ps, -2)
		pcs = append(pcs, pConsumer)
		//
		of, _ := client.GetOffset(topic, ps, -1)
		fmt.Println("Get  Off Set", of)
	}

	for _, ps := range ps {
		of, _ := client.GetOffset(topic, ps, -2)
		fmt.Println("Get  Off Set2", of)
	}

	var wg sync.WaitGroup
	for idx, pc := range pcs {
		wg.Add(1)
		fmt.Println("yyy idx", idx, pc, pcm[idx])
		go func(pc sarama.PartitionConsumer, pm sarama.PartitionOffsetManager) {
			defer wg.Done()
			// comsume it!!!
			for {
				select {
				case cmsg := <-pc.Messages():
					fmt.Println(cmsg.Value)
					pm.MarkOffset(cmsg.Offset+1, "'")
					//
					of, _ := client.GetOffset(topic, 0, -1)
					fmt.Println("Get  Off Set22220", of)
					of1, _ := client.GetOffset(topic, 1, -1)
					fmt.Println("Get  Off Set22221", of1)
				}
			}
		}(pc, pcm[idx])
	}
	wg.Wait()

}

func main() {
	//Produce()
	fmt.Println("=================")
	//Consume()
	//Consumer2()
}
