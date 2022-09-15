package hello

import (
	"encoding/json"
	"fmt"
	"marunk20/rs/hello/api"
	"net/http"
)

func getHelloMessage () string {

	helloInVariousLanguages:= []api.Hello {
		{
			Language: "English",
			Message:  "Hello Sir",
		},
		{
			Language: "French",
			Message:  "Bonjour Monsieur",
		},
		{
			Language: "Spanish",
			Message:  "Hola señor",
		},
	}

	helloResponse:= api.HelloApi{HelloAll: helloInVariousLanguages}

	response,_ := json.Marshal(helloResponse);

	return string(response)
}

func SayHello(httpWriter http.ResponseWriter, httpRequest *http.Request){
	fmt.Fprint(httpWriter,getHelloMessage())
}

