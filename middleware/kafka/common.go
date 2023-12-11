package kafka

import (
	"errors"

	"github.com/segmentio/kafka-go"
)

const brokerAddress = "localhost:9092"

func InitKafka() {
	for _, w := range writers {
		defer w.Close()
	}
	for _, r := range readers {
		defer r.Close()
	}
}

var writers = map[string]*kafka.Writer{
	"video-publish": kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{brokerAddress},
		Topic:   "video-publish",
	}),
	"video-delete": kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{brokerAddress},
		Topic:   "video-delete",
	}),
	"video-feed": kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{brokerAddress},
		Topic:   "video-feed",
	}),
}
var readers = map[string]*kafka.Reader{
	"video-publish": kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{brokerAddress},
		Topic:   "video-publish",
	}),
	"video-delete": kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{brokerAddress},
		Topic:   "video-delete",
	}),
	"video-feed": kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{brokerAddress},
		Topic:   "video-feed",
	}),
}

func GetWriter(topicName string) (*kafka.Writer, error) {
	if _, ok := writers[topicName]; ok {
		return writers[topicName], nil
	} else {
		return nil, errors.New("Unknown topic")
	}
}

func GetReader(topicName string) (*kafka.Reader, error) {
	if _, ok := readers[topicName]; ok {
		return readers[topicName], nil
	} else {
		return nil, errors.New("Unknown topic")
	}
}
