package main

import (
	/*	hydra "github.com/ory-am/hydra/sdk"
		gh "github.com/alastria/gateway/middleware"*/

	_ "github.com/alastria/gateway/routers"

	"crypto/tls"
	"crypto/x509"
	"io/ioutil"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

//var hc *hydra.Client
var log *logs.BeeLogger

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	if ok, err := beego.AppConfig.Bool("EnableHTTPS"); err == nil && ok {
		beego.BeeApp.Server.TLSConfig = configureTLS()
	}
	if ok, err := beego.AppConfig.Bool("EnableOauth2"); err == nil && ok {
		configureOauth2()
	}
	beego.SetLevel(beego.LevelDebug)
	beego.Run()
}

func configureTLS() *tls.Config {
	log := *logs.GetBeeLogger()
	log.Trace("Preparando la configuración TLS ")

	// http://www.levigross.com/2015/11/21/mutual-tls-authentication-in-go/
	// Load our TLS key pair to use for authentication
	cert, err := tls.LoadX509KeyPair(beego.AppConfig.String("TLSCertFile"), beego.AppConfig.String("TLSKeyFile"))
	if err != nil {
		log.Error("Unable to load cert", err)
	}

	// Load our CA certificate
	clientCACert, err := ioutil.ReadFile(beego.AppConfig.String("TLSCACertFile"))
	if err != nil {
		log.Error("Unable to open cert", err)
	}
	clientCertPool := x509.NewCertPool()
	clientCertPool.AppendCertsFromPEM(clientCACert)

	tlsConfig := &tls.Config{
		ClientAuth:               tls.RequireAndVerifyClientCert,
		Certificates:             []tls.Certificate{cert},
		ClientCAs:                clientCertPool,
		CipherSuites:             []uint16{tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256, tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384},
		PreferServerCipherSuites: true,
		MinVersion:               tls.VersionTLS12,
	}

	tlsConfig.BuildNameToCertificate()

	log.Trace("Finalizando la configuración TLS")

	return tlsConfig
}

func configureOauth2() {
	log.Trace("configureOauth2 is IN.")
	/*var err error
	// Initialize Hydra and gin-hydra
	if hc, err = hydra.Connect(
		hydra.ClientID(beego.AppConfig.String("HydraClientId")),
		hydra.ClientSecret(beego.AppConfig.String("HydraClientSecret")),
		hydra.ClusterURL(beego.AppConfig.String("hydraClusterUrl")),
		hydra.Scopes("openid", "ftl"),
	); err != nil {
		panic(err)
	}

	gh.Init(hc)


	beego.InsertFilter("/v1/alastria/createidentity", beego.BeforeRouter, gh.ScopesRequired("alastria.create"))

	beego.InsertFilter("*", beego.BeforeRouter, gh.ScopesRequired("alastria"))*/
}
