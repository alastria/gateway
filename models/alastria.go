package models

type Response struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
}

type IdentityCreation struct {
	Transaction   string `json:"transaction"`
	AlastriaToken string `json:"alastria_token"`
	PublicKey     string `json:"public_key"`
}

type ResponseIdentityCreation struct {
	Response
	AlastriaId string `json:"alastria_id"`
}

type ResponsePubKey struct {
	Response
	PublicKey string `json:"public_key"`
}

type UpdatePK struct {
	AlastriaId string `json:"alastria_id"`
	PublicKey  string `json:"public_key"`
}

type SendRawTransaction struct {
	RawTransaction string `json:"raw_transaction"`
}

type ResponseSendRawTransaction struct {
	Response
	Transaction string `json:"transaction"`
}

type ResponseCallGet struct {
	Response
	ResponseCall string `json:"response"`
}
