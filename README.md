# lambda-card

POC Lambda for technical purposes

Lambda persist CARD data inside DynamoDB and create a notification via event EventBridge

Diagrama Flow

    APIGW ==> Lambda ==> DynamoDB (card)
                     ==> EventBridge (event_type {card}) <== lambda-agregation-card-person-worker

## Compile

   Manually compile the function

    GOOD=linux GOARCH=amd64 go build -o ../build/main main.go

    zip -jrm ../build/main.zip ../build/main

## Endpoint

+ Get /version

+ Get /card/4444.000.000.005

+ Post /card

        {
            "id": "",
            "sk": "007",
            "card_number": "4444.000.000.007",
            "card_holder": "ELIEZER ANTUNES",
            "status": "ACTIVE",
            "valid": "10/29",
            "tenant_id": "TENANT-002"
        }

+ Post /card/status

        {
            "id": "",
            "sk": "",
            "card_number": "4444.000.000.007",
            "status": "CANCELED",
            "tenant_id": "TENANT-001"
        }

## Event

+ event_source: lambda.card

+ event_bus_name: event-bus-card

+ event type:

        eventTypeCreated =  "cardCreated"
        eventTypeUpdated =  "cardStatusUpdated"

+ EventPayload

        {
            "id": "CARD-4444.000.000.001",
            "sk": "PERSON:PERSON-010",
            "card_number": "4444.000.000.001",
            "card_holder": "Steve Michael",
            "status": "ACTIVE",
            "valid": "10/29",
            "tenant_id": "TENANT-001"
        }

## Pipeline

Prerequisite: 

Lambda function already created

+ buildspec.yml: build the main.go and move to S3
+ buildspec-test.yml: make a go test using services_test.go
+ buildspec-update.yml: update the lambda-function using S3 build and prepare to canary deploy
+ appspec: execute the canary deploy

## DynamoDB

    CARD-8888.000.100.001 PERSON:PERSON-100 Eliezer R A Junior 8888.000.100.001 ACTIVE TENANT-100 10/29


