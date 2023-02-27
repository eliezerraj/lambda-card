package notification

import (
	"os"
	"encoding/json"

	"github.com/rs/zerolog/log"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/eventbridge"

	"github.com/lambda-card/internal/core/domain"
	"github.com/lambda-card/internal/erro"

)

var childLogger = log.With().Str("notification", "eventBridge").Logger()

type CardNotification struct {
	client		*eventbridge.EventBridge
}

func NewCardNotification() (*CardNotification, error){
	childLogger.Debug().Msg("NewCardNotification")

	region := os.Getenv("AWS_REGION")
    awsSession, err := session.NewSession(&aws.Config{
        Region: aws.String(region)},
    )
	if err != nil {
		childLogger.Error().Err(err).Msg("error message") 
		return nil, erro.ErrCreateSession
	}
	return &CardNotification{
		client: eventbridge.New(awsSession),
	}, nil
}

func (n *CardNotification) PutEvent(card domain.Card, eventType string) error {
	childLogger.Debug().Msg("PutEvent")

	payload, err := json.Marshal(card)
	if err != nil {
		childLogger.Error().Err(err).Msg("error message") 
		return erro.ErrUnmarshal
	}

	entries := []*eventbridge.PutEventsRequestEntry{{
		Detail:       aws.String(string(payload)),
		DetailType:   aws.String(eventType),
		EventBusName: aws.String("event-bus-card"),
		Source:       aws.String("lambda-card"),
	}}

	_, err = n.client.PutEvents(&eventbridge.PutEventsInput{Entries: entries})
	if err != nil {
		childLogger.Error().Err(err).Msg("error message") 
		return erro.ErrPutEvent
	}

	return nil
}