package main

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

    "github.com/lambda-card/internal/service"
	"github.com/lambda-card/internal/adapter/handler"
	"github.com/lambda-card/internal/repository"

)

var (
	logLevel = zerolog.DebugLevel // InfoLevel DebugLevel
	tableName 			= "card"
	version 			= "lambda-card version 1.0"	
	response 			*events.APIGatewayProxyResponse
	cardRepository		*repository.CardRepository
	cardService 		*service.CardService
	cardHandler 		*handler.CardHandler
)

func init(){
	zerolog.SetGlobalLevel(logLevel)
}

func getEnv() {
	if os.Getenv("TABLE_NAME") !=  "" {
		tableName = os.Getenv("TABLE_NAME")
	}
	if os.Getenv("LOG_LEVEL") !=  "" {
		if (os.Getenv("LOG_LEVEL") == "DEBUG"){
			logLevel = zerolog.DebugLevel
		}else if (os.Getenv("LOG_LEVEL") == "INFO"){
			logLevel = zerolog.InfoLevel
		}else if (os.Getenv("LOG_LEVEL") == "ERROR"){
				logLevel = zerolog.ErrorLevel
		}else {
			logLevel = zerolog.InfoLevel
		}
	}
	if os.Getenv("VERSION") !=  "" {
		version = os.Getenv("VERSION")
	}
}

func main()  {
	log.Debug().Msg("main lambda-card (go) v 1.0")

	cardRepository, err := repository.NewCardRepository(tableName)
	if err != nil{
		return
	}
	cardService = service.NewCardService(*cardRepository)
	cardHandler = handler.NewCardHandler(*cardService)

	lambda.Start(lambdaHandler)
}

func lambdaHandler(req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	log.Debug().Msg("handler")
	switch req.HTTPMethod {
		case "GET":
			if (req.Resource == "/card/{id}"){
				response, _ = cardHandler.GetCard(req)
			}else if (req.Resource == "/version"){
				response, _ = cardHandler.GetVersion(version)
			}else {
				response, _ = cardHandler.UnhandledMethod()
			}
		case "POST":
			if (req.Resource == "/card"){
				response, _ = cardHandler.AddCard(req)
			}else if (req.Resource == "/card/status") {
				response, _ = cardHandler.SetCardStatus(req)
			}else {
				response, _ = cardHandler.UnhandledMethod()
			}
		case "DELETE":
			response, _ = cardHandler.UnhandledMethod()
		case "PUT":
			response, _ = cardHandler.UnhandledMethod()
		default:
			response, _ = cardHandler.UnhandledMethod()
	}

	return response, nil
}
