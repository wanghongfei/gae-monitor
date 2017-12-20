package msg

import (
	"github.com/Shopify/sarama"
	"time"
	"gaemonitor/logg"
	"gaemonitor/conf"
	"strings"
)

var asyncProducer sarama.AsyncProducer

func init() {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.Timeout = time.Second * 5

	p, err := sarama.NewAsyncProducer(strings.Split(conf.AppConfig.BrokerList, ","), config)
	if nil != err {
		panic(err)
	}

	go func(p sarama.AsyncProducer) {
		errChan := p.Errors()
		succChan := p.Successes()

		for {
			select {
				case err := <- errChan:
					if nil != err {
						logg.Logger.Println(err)
					}

				case <- succChan:
			}
		}
	}(p)

	asyncProducer = p

	logg.Logger.Println("kafka initialized")
}

func SendMessage(topic, message string) {
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.ByteEncoder(message),
	}

	asyncProducer.Input() <- msg
}
