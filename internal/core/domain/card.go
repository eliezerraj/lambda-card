package domain

type Card struct {
	ID				string	`json:"id,omitempty"`
	SK				string	`json:"sk,omitempty"`
	CardNumber		string  `json:"card_number,omitempty"`
	CardHolder		string  `json:"card_holder,omitempty"`
	Status			string  `json:"status,omitempty"`
	Valid			string  `json:"valid,omitempty"`
	Tenant			string  `json:"tenant_id,omitempty"`
}

func NewCard(id string, 
			sk string, 
			cardnumber string, 
			cardholder string,
			status	string,
			valid	string,
			tenant	string) *Card{
	return &Card{
		ID:	id,
		SK:	sk,
		CardNumber: cardnumber,
		CardHolder: cardholder,
		Status: status,
		Valid: valid,
		Tenant: tenant,
	}
}