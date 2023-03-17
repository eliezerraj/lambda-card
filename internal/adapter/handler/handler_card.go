package handler

import(

	"encoding/json"

	"github.com/rs/zerolog/log"
	"github.com/lambda-card/internal/core/domain"
	"github.com/lambda-card/internal/service"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"net/http"
	"github.com/lambda-card/internal/erro"

)

var childLogger = log.With().Str("handler", "CardHandler").Logger()

var transactionSuccess	= "Transação com sucesso"

type CardHandler struct {
	cardService service.CardService
}

type ErrorBody struct {
	ErrorMsg *string `json:"error,omitempty"`
}

type MessageBody struct {
	Msg *string `json:"message,omitempty"`
}

func NewCardHandler(cardService service.CardService) *CardHandler{
	childLogger.Debug().Msg("NewCardHandler")
	return &CardHandler{
		cardService: cardService,
	}
}

func (h *CardHandler) UnhandledMethod() (*events.APIGatewayProxyResponse, error){
	return ApiHandlerResponse(http.StatusMethodNotAllowed, ErrorBody{aws.String(erro.ErrMethodNotAllowed.Error())})
}

func (h *CardHandler) AddCard(req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	childLogger.Debug().Msg("AddCard")

    var card domain.Card
    if err := json.Unmarshal([]byte(req.Body), &card); err != nil {
        return ApiHandlerResponse(http.StatusBadRequest, ErrorBody{aws.String(err.Error())})
    }

	response, err := h.cardService.AddCard(card)
	if err != nil {
		return ApiHandlerResponse(http.StatusBadRequest, ErrorBody{aws.String(err.Error())})
	}

	handlerResponse, err := ApiHandlerResponse(http.StatusOK, response)
	if err != nil {
		return ApiHandlerResponse(http.StatusInternalServerError, ErrorBody{aws.String(err.Error())})
	}
	return handlerResponse, nil
}

func (h *CardHandler) SetCardStatus(req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	childLogger.Debug().Msg("SetCardStatus")

    var card domain.Card
    if err := json.Unmarshal([]byte(req.Body), &card); err != nil {
        return ApiHandlerResponse(http.StatusBadRequest, ErrorBody{aws.String(err.Error())})
    }

	/*response, err := h.cardService.GetCard(card)
	if err != nil {
		return ApiHandlerResponse(http.StatusNotFound, ErrorBody{aws.String(err.Error())})
	}

	response.Status = card.Status*/
	response, err = h.cardService.SetCardStatus(*response)
	if err != nil {
		return ApiHandlerResponse(http.StatusBadRequest, ErrorBody{aws.String(err.Error())})
	}

	handlerResponse, err := ApiHandlerResponse(http.StatusOK, response)
	if err != nil {
		return ApiHandlerResponse(http.StatusBadRequest, ErrorBody{aws.String(err.Error())})
	}
	return handlerResponse, nil
}

func (h *CardHandler) GetCard(req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	childLogger.Debug().Msg("GetCard")

	id := req.PathParameters["id"]
	if len(id) == 0 {
		return ApiHandlerResponse(http.StatusBadRequest, ErrorBody{aws.String(erro.ErrQueryEmpty.Error())})
	}

 	card := domain.NewCard("","",id,"","","","TENANT-001")

	response, err := h.cardService.GetCard(*card)
	if err != nil {
		return ApiHandlerResponse(http.StatusNotFound, ErrorBody{aws.String(err.Error())})
	}

	handlerResponse, err := ApiHandlerResponse(http.StatusOK, response)
	if err != nil {
		return ApiHandlerResponse(http.StatusInternalServerError, ErrorBody{aws.String(err.Error())})
	}
	return handlerResponse, nil
}

func (h *CardHandler) GetVersion(version string) (*events.APIGatewayProxyResponse, error) {
	childLogger.Debug().Msg("GetVersion")

	response := MessageBody { Msg: &version }
	handlerResponse, err := ApiHandlerResponse(http.StatusOK, response)
	if err != nil {
		return ApiHandlerResponse(http.StatusInternalServerError, ErrorBody{aws.String(err.Error())})
	}

	return handlerResponse, nil
}