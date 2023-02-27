package repository

import (
	"os"

	"github.com/rs/zerolog/log"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"

	"github.com/lambda-card/internal/erro"

)

var childLogger = log.With().Str("repository", "CardRepository").Logger()

type CardRepository struct {
	client 		dynamodbiface.DynamoDBAPI
	tableName   *string
}

func NewCardRepository(tableName string) (*CardRepository, error){
	childLogger.Debug().Msg("NewCardRepository")
	
	region := os.Getenv("AWS_REGION")
    awsSession, err := session.NewSession(&aws.Config{
        Region: aws.String(region)},
    )
	if err != nil {
		childLogger.Error().Err(err).Msg("error message") 
		return nil, erro.ErrCreateSession
	}

	return &CardRepository{
		client: dynamodb.New(awsSession),
		tableName: aws.String(tableName),
	},nil
}