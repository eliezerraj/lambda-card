package service

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/rs/zerolog"

	"github.com/lambda-card/internal/repository"
	"github.com/lambda-card/internal/core/domain"

)

var (
	//logLevel = 
	tableName = "card"
	cardRepository	*repository.CardRepository
	card01 = domain.NewCard("CARD-4444.000.000.001",
							"CARD-4444.000.000.001",
							"4444.000.000.001",
							"ELIEZER R A JR",
							"ACTIVE",
							"02/26",
							"TENANT-001")

	card02 = domain.NewCard("CARD-4444.000.000.002",
							"CARD-4444.000.000.002",
							"4444.000.000.002",
							"JULIANA PIVATO",
							"ACTIVE",
							"02/26",
							"TENANT-001")
)


func TestAddCard(t *testing.T) {
	t.Setenv("AWS_REGION", "us-east-2")
	zerolog.SetGlobalLevel(zerolog.ErrorLevel)

	repository, err := repository.NewCardRepository(tableName)
	if err != nil {
		t.Errorf("Error - TestAddCard Create Repository DynanoDB")
	}

	service	:= NewCardService(*repository)

	result, err := service.AddCard(*card01)
	if err != nil {
		t.Errorf("Error -TestAddCard Access DynanoDB %v ", tableName)
	}

	if (cmp.Equal(card01, result)) {
		t.Logf("Success on TestAddCard!!! result : %v ", result)
	} else {
		t.Errorf("Error TestAddCard input : %v" , *card01)
	}
}

func TestGetCard(t *testing.T) {
	t.Setenv("AWS_REGION", "us-east-2")
	zerolog.SetGlobalLevel(zerolog.ErrorLevel)

	repository, err := repository.NewCardRepository(tableName)
	if err != nil {
		t.Errorf("Error - TestGetCard Create Repository DynanoDB")
	}

	service	:= NewCardService(*repository)

	result, err := service.GetCard(*card01)
	if err != nil {
		t.Errorf("Error -TestGetCard Access DynanoDB %v ", tableName)
	}

	if (cmp.Equal(card01, result)) {
		t.Logf("Success on TestGetCard!!! result : %v ", result)
	} else {
		t.Errorf("Error TestGetCard input : %v" , *card01)
	}

	/*result, err = service.GetCard(*card02)
	if err != nil {
		t.Logf("Success - TestGetCard Card NOT FOUND %v ", card02)
	} else {
		t.Logf("Errorf - TestGetCard Card FOUND %v ", card02)
	}*/
}

func TestGetStatusCard(t *testing.T) {
	t.Setenv("AWS_REGION", "us-east-2")
	zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	
	repository, err := repository.NewCardRepository(tableName)
	if err != nil {
		t.Errorf("Error - TestGetStatusCard Create Repository DynanoDB")
	}

	service	:= NewCardService(*repository)

	card01.Status = "HOLD"
	result, err := service.SetCardStatus(*card01)
	if err != nil {
		t.Errorf("Error -TestGetStatusCard Access DynanoDB %v ", tableName)
	}

	if (cmp.Equal(card01, result)) {
		t.Logf("Success on TestGetStatusCard!!! result : %v ", result)
	} else {
		t.Errorf("Error TestGetStatusCard input : %v" , *card01)
	}
}