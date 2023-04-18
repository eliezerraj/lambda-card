GOOD=linux GOARCH=amd64 go build -o ../build/main main.go

zip -jrm ../build/main.zip ../build/main

//------------------------

Endpoints

Get /version
Get /card/4444.000.000.005
Post /card
{
    "id": "",
    "sk": "007",
    "card_number": "4444.000.000.007",
    "card_holder": "ELIEZER ANTUNES",
    "status": "ACTIVE",
    "valid": "10/29",
    "tenant_id": "TENANT-002"
}
Post /card/status
{
    "id": "",
    "sk": "",
    "card_number": "4444.000.000.007",
    "status": "CANCELED",
    "tenant_id": "TENANT-001"
}

//------

APIGW ==> Lambda ==> DynamoDB (card)
                 ==> EventBridge (agregation-card-person {card})

//-----