package main

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

    "github.com/lambda-card/internal/service"
	"github.com/lambda-card/internal/adapter/handler"
	"github.com/lambda-card/internal/adapter/notification"
	"github.com/lambda-card/internal/repository"

)

var (
	logLevel		=	zerolog.DebugLevel // InfoLevel DebugLevel
	tableName		=	"card"
	version			=	"lambda-card (github) version 1.0"
	eventSource		=	"lambda-card"
	eventBusName	=	"event-bus-card"	
	response		*events.APIGatewayProxyResponse
	cardRepository	*repository.CardRepository
	cardService		*service.CardService
	cardHandler		*handler.CardHandler
	cardNotification *notification.CardNotification
)

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
			logLevel = zerolog.DebugLevel
		}
	}
	if os.Getenv("VERSION") !=  "" {
		version = os.Getenv("VERSION")
	}
}

func init(){
	log.Debug().Msg("init")
	zerolog.SetGlobalLevel(logLevel)
	getEnv()
}

func main()  {
	log.Debug().Msg("main lambda-card (go) v 2.0")
	log.Debug().Msg("-------------------")
	log.Debug().Str("version", version).
				Str("tableName", tableName).
				Msg("Enviroment Variables")
	log.Debug().Msg("--------------------")

	cardRepository, err := repository.NewCardRepository(tableName)
	if err != nil{
		return
	}
	cardNotification, err = notification.NewCardNotification(eventSource,eventBusName)
	if err != nil{
		return
	}

	cardService 	= service.NewCardService(*cardRepository, *cardNotification)
	cardHandler 	= handler.NewCardHandler(*cardService)

	lambda.Start(lambdaHandler)
}

func lambdaHandler(req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	log.Debug().Msg("handler")
	log.Debug().Msg("-------------------")
	log.Debug().Str("req.Body", req.Body).
				Msg("APIGateway Request.Body")
	log.Debug().Msg("--------------------")

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
