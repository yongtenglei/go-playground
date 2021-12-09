package main

import (
	"fmt"
	"log"
	"sync"

	"github.com/Shopify/sarama"
)

func main() {
	//// producer config
	//config := sarama.NewConfig()
	//config.Producer.RequiredAcks = sarama.WaitForAll
	//config.Producer.Partitioner = sarama.NewRandomPartitioner
	//config.Producer.Return.Successes = true

	//// connect to kafka
	//client, err := sarama.NewSyncProducer([]string{"localhost:9092"}, config)
	//if err != nil {
	//log.Fatal("producer connection err: ", err)
	//return
	//}
	//defer client.Close()

	//// encapsulate message
	//msg := &sarama.ProducerMessage{}
	//msg.Topic = "quickstart-events"
	//msg.Value = sarama.StringEncoder("Here is Go! Nice to meet you!")

	//// send message
	//pid, offset, err := client.SendMessage(msg)
	//if err != nil {
	//log.Fatal("send message err: ", err)
	//return
	//}
	//fmt.Printf("pid: %v\toffset: %v", pid, offset)

	var wg sync.WaitGroup
	// consumer
	consumer, err := sarama.NewConsumer([]string{"localhost:9092"}, nil)
	if err != nil {
		fmt.Printf("consumer_test create consumer error %s\n", err.Error())
		return
	}
	defer consumer.Close()

	partitionList, err := consumer.Partitions("quickstart-events")
	if err != nil {
		log.Fatal("partitionList err: \n", err)
		return
	}

	for partition := range partitionList {
		pc, err := consumer.ConsumePartition("quickstart-events", int32(partition), sarama.OffsetNewest)
		if err != nil {
			log.Fatalf("failed to start consumer for partition %d, err: %v\n", partition, err)
			return
		}
		defer pc.AsyncClose()
		wg.Add(1)

		go func(sarama.PartitionConsumer) {
			for msg := range pc.Messages() {
				fmt.Printf("Partition %d\toffset: %d\tkey: %s\tvalue: %s\n", msg.Partition, msg.Offset, msg.Key, msg.Value)

			}
		}(pc)
	}

	wg.Wait()
}
