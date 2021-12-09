package tailfile

import (
	"log_agent/etcd"
	"log_agent/kafka"
	"strings"
	"time"

	"github.com/Shopify/sarama"
	"github.com/hpcloud/tail"
	log "github.com/sirupsen/logrus"
)

type tailTask struct {
	path    string
	topic   string
	tailIns *tail.Tail
}

func (t *tailTask) run() {
	for {
		line, ok := <-t.tailIns.Lines
		if !ok {
			log.Warnf("tail file close reopen, filename:%s\n", t.path)
			time.Sleep(time.Second)
			continue
		}

		if strings.Trim(line.Text, "\r\n") == "" {
			continue
		}

		msg := &sarama.ProducerMessage{}
		msg.Topic = t.topic
		msg.Value = sarama.StringEncoder(line.Text)

		kafka.MsgChan(msg)

	}

}

func Init(confList *[]etcd.CollectEntry) (err error) {
	config := tail.Config{
		ReOpen:    true,
		Follow:    true,
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2},
		MustExist: false,
		Poll:      true,
	}

	for _, conf := range *confList {
		tt := &tailTask{
			path:  conf.Path,
			topic: conf.Topic,
		}

		if tailIns, err := tail.TailFile(tt.path, config); err != nil {
			log.Warningf("Create tail file: %s failed, err: %v", tt.path, err)
			continue
		} else {
			tt.tailIns = tailIns
			log.Infof("Create A tail Task: %s successfully\n", tt.path)
		}

		log.Infof("Task: %s running...\n", tt.path)
		go tt.run()
	}

	return
}
