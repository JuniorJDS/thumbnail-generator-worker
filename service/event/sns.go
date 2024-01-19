package event

import (
	"encoding/json"
	"fmt"
	"thumbnail-generator-worker/config"
	infra "thumbnail-generator-worker/infra/aws"
	"thumbnail-generator-worker/schema"

	"github.com/aws/aws-sdk-go/service/sns"
)

var settings = config.GetSettings()

type EventSNS struct {
	SnsClient *sns.SNS
}

func NewEventSNS() *EventSNS {
	session := infra.GetSessionAWS()
	svc := sns.New(session)
	return &EventSNS{
		SnsClient: svc,
	}
}

func (e *EventSNS) Publish(message schema.Message) error {
	topicPtr := settings["TOPIC"]

	messageByte, err := json.Marshal(message)
	if err != nil {
		return err
	}
	messageString := string(messageByte)

	result, err := e.SnsClient.Publish(
		&sns.PublishInput{
			Message:  &messageString,
			TopicArn: &topicPtr,
		},
	)
	if err != nil {
		return nil
	}

	fmt.Println(*result.MessageId)
	return nil
}

func (e *EventSNS) Subscribe() error {
	return nil
}
