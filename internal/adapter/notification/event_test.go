package notification

import (
	"testing"

	"github.com/lambda-card/internal/core/domain"

)

var (
	tableName = "card"
	eventSource	= "lambda-card"
	eventBusName	= "event-bus-card"
	card01 = domain.NewCard("CARD-4444.000.000.001",
							"CARD-4444.000.000.001",
							"4444.000.000.001",
							"ELIEZER R A JR",
							"ACTIVE",
							"02/26",
							"TENANT-001")
)

func TestPutEvent(t *testing.T) {
	t.Setenv("AWS_REGION", "us-east-2")

	notification, err := NewCardNotification(eventSource,eventBusName)
	if err != nil {
		t.Errorf("Error -TestPutEvent Access EventBridge %v ", err)
	}

	eventType := "add-new-card"
	err = notification.PutEvent(*card01, eventType)
	if err != nil {
		t.Errorf("Error -TestPutEvent Access EventBridge %v ", err)
	}
}
