// Package classification Instance Job Manager Service.
//
// Instance Job Manager Service as part of the Instance Manager environment
//
// Terms Of Service:
//
// There are no TOS at this moment, use at your own risk we take no responsibility
//
//    Version: 0.1.0
//    License: TODO
//    Contact: <info@dhis2.org> https://github.com/dhis2-sre/im-job
//
//    Consumes:
//      - application/json
//
//    Produces:
//      - application/json
//
//    SecurityDefinitions:
//      oauth2:
//        type: oauth2
//        tokenUrl: /not-valid--endpoint-is-served-from-the-im-user-service
//        refreshUrl: /not-valid--endpoint-is-served-from-the-im-user-service
//        flow: password
// swagger:meta
package main

import (
	"github.com/dhis2-sre/im-job/internal/di"
	"github.com/dhis2-sre/im-job/internal/server"
	"log"
)

func main() {
	environment := di.GetEnvironment()

	r := server.GetEngine(environment)
	if err := r.Run(); err != nil {
		log.Fatal(err)
	}
}
