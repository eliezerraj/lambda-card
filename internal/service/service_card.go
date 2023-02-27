package service

import (
	"github.com/lambda-card/internal/core/domain"

)

func (s *CardService) AddCard(card domain.Card) (*domain.Card, error){
	childLogger.Debug().Msg("AddCard")

	c, err := s.cardRepository.AddCard(card)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (s *CardService) GetCard(card domain.Card) (*domain.Card, error){
	childLogger.Debug().Msg("GetCard")

	c, err := s.cardRepository.GetCard(card)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (s *CardService) SetCardStatus(card domain.Card) (*domain.Card, error){
	childLogger.Debug().Msg("SetCardStatus")

	c, err := s.cardRepository.AddCard(card)
	if err != nil {
		return nil, err
	}

	return c, nil
}