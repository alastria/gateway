package controllers

import (
	"encoding/json"

	"github.com/alastria/gateway/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

// Operations about Alastrias
type AlastriaController struct {
	beego.Controller
	log logs.BeeLogger
}

// @Title Identity_creation
// @Description For identity creation, you need call with a jwt token in authorization header.
// @Param	body		body 	models.IdentityCreation	true		"Transaction for IdentityManager call, Alastria Token and the Subject public key."
// @Success 200 {object} models.ResponseIdentityCreation
// @Failure 403 response.Response.Code is 403
// @router /identity [post]
func (u *AlastriaController) IdentityCreation() (response models.ResponseIdentityCreation) {
	var body models.IdentityCreation
	json.Unmarshal(u.Ctx.Input.RequestBody, &body)

	u.log.Info("IdentityCreation: body: %s", body)

	// TODO: Goes to service and create the new identity
	response = models.ResponseIdentityCreation{
		Response: models.Response{
			Code:  200,
			Error: "nil",
		},
		AlastriaId: "0x393609f2b0e39be2bbf9080ad4ee1f67fc17d8e0",
	}

	return response
}

// @Title GetPubKey
// @Description Get the public key of an user with alastria_id.
// @Param	alastriaId		path 	string	true		"The address of the alastria id user."
// @Success 200 {object} models.ResponsePubKey
// @Failure 403 response.Response.Code is 403
// @router /pubkey/:alastriaId [get]
func (u *AlastriaController) GetPubKey(alastriaId string) (response models.ResponsePubKey) {

	u.log.Info("GetPubKey-> alastria_id: %s", alastriaId)

	response = models.ResponsePubKey{
		Response: models.Response{
			Code:  200,
			Error: "nil",
		},
		PublicKey: "[]byte{}",
	}

	return response
}

// @Title SetPubKey
// @Description Set or update the public key to an alastria id
// @Param	body		body 	models.UpdatePK	true		"body for Alastria content"
// @Success 200 {object} models.Response
// @Failure 403 response.Response.Code is 403
// @router /pubkey [post]
func (u *AlastriaController) SetPubKey() (response models.Response) {
	var body models.UpdatePK
	json.Unmarshal(u.Ctx.Input.RequestBody, &body)

	u.log.Info("SetPubKey-> body: %s", body)

	response = models.Response{
		Code:  200,
		Error: "nil",
	}

	return response
}

// @Title SendRawTransaction
// @Description Send a raw transaction.
// @Param	body		body 	models.SendRawTransaction	true		"Contains a RawTransaction."
// @Success 200 {object} models.ResponseSendRawTransaction
// @Failure 403 response.Response.Code is 403
// @router /raw [post]
func (u *AlastriaController) RawTransaction() (response models.ResponseSendRawTransaction) {
	var body models.SendRawTransaction
	json.Unmarshal(u.Ctx.Input.RequestBody, &body)

	u.log.Info("RawTransaction-> body: %s", body)

	// TODO: Goes to service and create the new identity
	response = models.ResponseSendRawTransaction{
		Response: models.Response{
			Code:  200,
			Error: "nil",
		},
		Transaction: "0xc3b7eea8d037bcfc9a3b99f4ec5500aaff39356634e4274022418b58f5d9e3c5",
	}

	return response
}

// @Title CallContract
// @Description get result of a call to a view function
// @Param	call		path 	string	true		"The call to the contract"
// @Success 200 {object} models.ResponseCallGet
// @Failure 403 response.Response.Code is 403
// @router / [get]
func (u *AlastriaController) CallContract(call string) (response models.ResponseCallGet) {

	u.log.Info("CallContract-> body: %s", call)

	// TODO: Goes to service and create the new identity
	response = models.ResponseCallGet{
		Response: models.Response{
			Code:  200,
			Error: "nil",
		},
		ResponseCall: "0",
	}

	return response
}
