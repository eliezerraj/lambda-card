package service

import (
	"github.com/lambda-card/internal/core/domain"

)

func (s *CardService) AddCard(card domain.Card) (*domain.Card, error){
	childLogger.Debug().Msg("AddCard")

	// Setting ID and Sort key
	card.ID = "CARD-" + card.CardNumber
	card.SK = "PERSON:PERSON-" + card.SK
	// Add new card 
	c, err := s.cardRepository.AddCard(card)
	if err != nil {
		return nil, err
	}

	// Stream new card
	eventType := "add-new-card"
	err = s.cardNotification.PutEvent(*c, eventType)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (s *CardService) GetCard(card domain.Card) (*domain.Card, error){
	childLogger.Debug().Msg("GetCard")
	
	// Setting ID and Sort key
	card.ID = "CARD-" + card.CardNumber
	card.SK = "PERSON:PERSON-" + card.SK
	c, err := s.cardRepository.GetCard(card)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (s *CardService) SetCardStatus(card domain.Card) (*domain.Card, error){
	childLogger.Debug().Msg("SetCardStatus")

	// Setting ID and Sort key
	card.ID = "CARD-" + card.CardNumber
	card.SK = "PERSON:PERSON-" + card.SK
	// Change status card, as the DB is a Dynamo de AddCard is a Upsert
	c, err := s.cardRepository.AddCard(card)
	if err != nil {
		return nil, err
	}

	// Stream new card
	eventType := "change-status-card"
	err = s.cardNotification.PutEvent(*c, eventType)
	if err != nil {
		return nil, err
	}

	return c, nil
}