package main

import (
	"Authentication-Service/gen/restapi"
	"Authentication-Service/gen/restapi/operations"
	"Authentication-Service/handlers"
	"flag"
	"github.com/go-openapi/loads"
	"log"
)

var portFlag = flag.Int("Port", 3000, "Port to run this server on.")

func main(){
	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil{
		log.Fatal(err)
	}

	api := operations.NewAuthenticationServiceAPI(swaggerSpec)
	server := restapi.NewServer(api)
	defer server.Shutdown()

	flag.Parsed()
	server.Port = *portFlag

	api.EmployeeSignupHandler = handlers.SignupEmployeeHandler()
	api.LoginHandler = handlers.Lgoin()

	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}
}
