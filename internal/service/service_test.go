package service

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/rs/zerolog"

	"github.com/lambda-card/internal/repository"
	"github.com/lambda-card/internal/adapter/notification"
	"github.com/lambda-card/internal/core/domain"

)

var (
	tableName = "card"
	eventSource	= "lambda.card"
	eventBusName	= "event-bus-card"
	cardRepository	*repository.CardRepository
)

func TestAddCard(t *testing.T) {
	t.Setenv("AWS_REGION", "us-east-2")
	zerolog.SetGlobalLevel(zerolog.ErrorLevel)

	repository, err := repository.NewCardRepository(tableName)
	if err != nil {
		t.Errorf("Error - TestAddCard Create Repository DynanoDB")
	}

	notification, err := notification.NewCardNotification(eventSource,eventBusName)
	if err != nil {
		t.Errorf("Error -TestPutEvent Access EventBridge %v ", err)
	}

	service	:= NewCardService(*repository, *notification)

	card01 := domain.NewCard("",
							"001",
							"4444.000.000.001",
							"ELIEZER R A JR",
							"ACTIVE",
							"02/26",
							"TENANT-001")
	result, err := service.AddCard(*card01)
	if err != nil {
		t.Errorf("Error -TestAddCard Access DynanoDB %v ", tableName)
	}

	//adjust ID + SK 
	card01.ID = "CARD-" + card01.CardNumber
	card01.SK = "PERSON:PERSON-" + card01.SK
	if (cmp.Equal(card01, result)) {
		t.Logf("Success on TestAddCard!!! result : %v ", result)
	} else {
		t.Errorf("Error TestAddCard input : %v || result: %v" , *card01, result)
	}
}

func TestGetCard(t *testing.T) {
	t.Setenv("AWS_REGION", "us-east-2")
	zerolog.SetGlobalLevel(zerolog.ErrorLevel)

	repository, err := repository.NewCardRepository(tableName)
	if err != nil {
		t.Errorf("Error - TestGetCard Create Repository DynanoDB")
	}
	notification, err := notification.NewCardNotification(eventSource,eventBusName)
	if err != nil {
		t.Errorf("Error -TestPutEvent Access EventBridge %v ", err)
	}

	service	:= NewCardService(*repository, *notification)

	card01 := domain.NewCard("",
							"001",
							"4444.000.000.001",
							"ELIEZER R A JR",
							"ACTIVE",
							"02/26",
							"TENANT-001")
	result, err := service.GetCard(*card01)
	if err != nil {
		t.Errorf("Error -TestGetCard Access DynanoDB %v ", tableName)
	}

	//adjust ID + SK 
	card01.ID = "CARD-" + card01.CardNumber
	card01.SK = "PERSON:PERSON-" + card01.SK
	if (cmp.Equal(card01, result)) {
		t.Logf("Success on TestGetCard!!! result : %v ", result)
	} else {
		t.Errorf("Error TestGetCard input : %v" , *card01)
	}

	card02 := domain.NewCard("",
							"002",
							"4444.000.000.002",
							"JULIANA PIVATO",
							"ACTIVE",
							"02/26",
							"TENANT-001")

	result, err = service.GetCard(*card02)
	if err != nil {
		t.Logf("Success - TestGetCard Card NOT FOUND %v ", card02)
	} else {
		t.Logf("Errorf - TestGetCard Card FOUND %v ", card02)
	}
}

func TestGetStatusCard(t *testing.T) {
	t.Setenv("AWS_REGION", "us-east-2")
	zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	
	repository, err := repository.NewCardRepository(tableName)
	if err != nil {
		t.Errorf("Error - TestGetStatusCard Create Repository DynanoDB")
	}

	notification, err := notification.NewCardNotification(eventSource,eventBusName)
	if err != nil {
		t.Errorf("Error -TestPutEvent Access EventBridge %v ", err)
	}

	service	:= NewCardService(*repository, *notification)
	card01 := domain.NewCard("",
							"001",
							"4444.000.000.001",
							"ELIEZER R A JR",
							"HOLD",
							"02/26",
							"TENANT-001")
	result, err := service.SetCardStatus(*card01)
	if err != nil {
		t.Errorf("Error -TestGetStatusCard Access DynanoDB %v ", tableName)
	}

	//adjust ID + SK 
	card01.ID = "CARD-" + card01.CardNumber
	card01.SK = "PERSON:PERSON-" + card01.SK
	if (cmp.Equal(card01, result)) {
		t.Logf("Success on TestGetStatusCard!!! result : %v ", result)
	} else {
		t.Errorf("Error TestGetStatusCard input : %v" , *card01)
	}
}
