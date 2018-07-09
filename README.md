# Gatewary for Alastria ID

# Up & Running

```
go get -u github.com/astaxie/beego
go get -u github.com/beego/bee
go get github.com/alastria/gateway
cd $GOPATH/src/github.com/alastria/gateway
bee run -vendor=true -gendoc=true -downdoc=true
```

Swagger for testing purpose: http://localhost:8080/swagger

## Known issues

[ ] When it call to one end-point, it fails with:

```
2018/07/09 13:24:09.919 [C] [asm_amd64.s:509] the request url is  /v1/alastria/pubkey/0x543709kjfdsahgùw
2018/07/09 13:24:09.919 [C] [asm_amd64.s:509] Handler crashed with error reflect: reflect.Value.Set using value obtained using unexported field
2018/07/09 13:24:09.919 [C] [asm_amd64.s:509] /usr/local/go/src/runtime/asm_amd64.s:509
2018/07/09 13:24:09.919 [C] [asm_amd64.s:509] /usr/local/go/src/runtime/panic.go:491
2018/07/09 13:24:09.919 [C] [asm_amd64.s:509] /usr/local/go/src/reflect/value.go:225
2018/07/09 13:24:09.919 [C] [asm_amd64.s:509] /usr/local/go/src/reflect/value.go:1351
2018/07/09 13:24:09.919 [C] [asm_amd64.s:509] /home/marcos/dev/workspace/1.8/src/github.com/alastria/gateway/vendor/github.com/astaxie/beego/router.go:206
2018/07/09 13:24:09.919 [C] [asm_amd64.s:509] /home/marcos/dev/workspace/1.8/src/github.com/alastria/gateway/vendor/github.com/astaxie/beego/router.go:795
2018/07/09 13:24:09.919 [C] [asm_amd64.s:509] /usr/local/go/src/net/http/server.go:2619
2018/07/09 13:24:09.919 [C] [asm_amd64.s:509] /usr/local/go/src/net/http/server.go:1801
2018/07/09 13:24:09.919 [C] [asm_amd64.s:509] /usr/local/go/src/runtime/asm_amd64.s:2337
```

[ ] It's a mocked rest service.