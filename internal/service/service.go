package service

import (
	"github.com/rs/zerolog/log"

	"github.com/lambda-card/internal/repository"

)

var childLogger = log.With().Str("service", "CardService").Logger()

type CardService struct {
	cardRepository repository.CardRepository
}

func NewCardService(cardRepository repository.CardRepository) *CardService{
	childLogger.Debug().Msg("NewCardsService")
	return &CardService{
		cardRepository: cardRepository,
	}
}