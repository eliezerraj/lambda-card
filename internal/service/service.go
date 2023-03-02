package service

import (
	"github.com/rs/zerolog/log"

	"github.com/lambda-card/internal/repository"
	"github.com/lambda-card/internal/adapter/notification"

)

var childLogger = log.With().Str("service", "CardService").Logger()

type CardService struct {
	cardRepository repository.CardRepository
	cardNotification notification.CardNotification
}

func NewCardService(cardRepository repository.CardRepository,
					cardNotification notification.CardNotification) *CardService{
	childLogger.Debug().Msg("NewCardsService")
	return &CardService{
		cardRepository: cardRepository,
		cardNotification: cardNotification,
	}
}