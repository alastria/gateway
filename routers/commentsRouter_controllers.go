package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["github.com/alastria/gateway/controllers:AlastriaController"] = append(beego.GlobalControllerRouter["github.com/alastria/gateway/controllers:AlastriaController"],
		beego.ControllerComments{
			Method: "CallContract",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(
				param.New("call", param.IsRequired, param.InPath),
			),
			Params: nil})

	beego.GlobalControllerRouter["github.com/alastria/gateway/controllers:AlastriaController"] = append(beego.GlobalControllerRouter["github.com/alastria/gateway/controllers:AlastriaController"],
		beego.ControllerComments{
			Method: "IdentityCreation",
			Router: `/identity`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/alastria/gateway/controllers:AlastriaController"] = append(beego.GlobalControllerRouter["github.com/alastria/gateway/controllers:AlastriaController"],
		beego.ControllerComments{
			Method: "SetPubKey",
			Router: `/pubkey`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/alastria/gateway/controllers:AlastriaController"] = append(beego.GlobalControllerRouter["github.com/alastria/gateway/controllers:AlastriaController"],
		beego.ControllerComments{
			Method: "GetPubKey",
			Router: `/pubkey/:alastriaId`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(
				param.New("alastriaId", param.IsRequired, param.InPath),
			),
			Params: nil})

	beego.GlobalControllerRouter["github.com/alastria/gateway/controllers:AlastriaController"] = append(beego.GlobalControllerRouter["github.com/alastria/gateway/controllers:AlastriaController"],
		beego.ControllerComments{
			Method: "RawTransaction",
			Router: `/raw`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

}
