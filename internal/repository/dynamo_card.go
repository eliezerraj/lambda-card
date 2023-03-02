package repository

import(

	"github.com/lambda-card/internal/core/domain"
	"github.com/lambda-card/internal/erro"

	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
	"github.com/aws/aws-sdk-go/service/dynamodb"

)

func (r *CardRepository) Ping() (bool, error){
	return true, nil
}

func (r *CardRepository) AddCard(card domain.Card) (*domain.Card, error){
	childLogger.Debug().Msg("AddCard")

	item, err := dynamodbattribute.MarshalMap(card)
	if err != nil {
		childLogger.Error().Err(err).Msg("error message")
		return nil, erro.ErrUnmarshal
	}

	transactItems := []*dynamodb.TransactWriteItem{}
	transactItems = append(transactItems, &dynamodb.TransactWriteItem{Put: &dynamodb.Put{
		TableName: r.tableName,
		Item:      item,
	}})

	transaction := &dynamodb.TransactWriteItemsInput{TransactItems: transactItems}
	if err := transaction.Validate(); err != nil {
		childLogger.Error().Err(err).Msg("error message") 
		return nil, erro.ErrInsert
	}

	_, err = r.client.TransactWriteItems(transaction)
	if err != nil {
		childLogger.Error().Err(err).Msg("error message")
		return nil, erro.ErrInsert
	}

	return &card ,nil
}

func (r *CardRepository) GetCard(card domain.Card) (*domain.Card, error){
	childLogger.Debug().Msg("GetCard")

	var keyCond expression.KeyConditionBuilder

	keyCond = expression.KeyAnd(
		expression.Key("id").Equal(expression.Value(card.ID)),
		expression.Key("sk").BeginsWith(card.SK),
	)

	expr, err := expression.NewBuilder().
							WithKeyCondition(keyCond).
							Build()
	if err != nil {
		childLogger.Error().Err(err).Msg("error message")
		return nil, erro.ErrPreparedQuery
	}

	key := &dynamodb.QueryInput{
			TableName:                 r.tableName,
			ExpressionAttributeNames:  expr.Names(),
			ExpressionAttributeValues: expr.Values(),
			KeyConditionExpression:    expr.KeyCondition(),
	}

	result, err := r.client.Query(key)
	if err != nil {
		childLogger.Error().Err(err).Msg("error message")
		return nil, erro.ErrQuery
	}

	card_result := []domain.Card{}
	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, &card_result)
    if err != nil {
		childLogger.Error().Err(err).Msg("error message")
		return nil, erro.ErrUnmarshal
    }

	if len(card_result) == 0 {
		return nil, erro.ErrNotFound
	} else {
		return &card_result[0], nil
	}
}