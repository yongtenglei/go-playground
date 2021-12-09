package kafka

import (
	log "github.com/sirupsen/logrus"

	"github.com/Shopify/sarama"
)

var (
	Client  sarama.SyncProducer
	msgChan chan *sarama.ProducerMessage
)

func Init(addresses []string, chanSize int64) (err error) {
	// producer config
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true

	// connect to kafka
	Client, err = sarama.NewSyncProducer(addresses, config)
	if err != nil {
		log.Fatal("producer connection err: \n", err)
		return
	}

	msgChan = make(chan *sarama.ProducerMessage, chanSize)

	go sendMsg()
	return

}

func MsgChan(msg *sarama.ProducerMessage) {
	msgChan <- msg
}

func sendMsg() {
	for {
		select {
		case msg := <-msgChan:
			// send message
			pid, offset, err := Client.SendMessage(msg)
			if err != nil {
				log.Warnf("send message err: %v\n", err)
				continue
			}
			log.Infof("send message successfully, pid: %v\toffset: %v\n", pid, offset)
			log.Infof("%s", msg.Value)
		}
	}
}
