package main

import (
	"flag"
	"io/ioutil"
	"strconv"

	log "github.com/sirupsen/logrus"

	"HarborMaster/managers"
	"HarborMaster/server"
)

func main() {
	urlPtr := flag.String("url", "localhost:8080", "URL of the docker registry")
	usernameFilePtr := flag.String("username_file", "/run/secrets/hm_registry_username", "Username of the docker registry")
	passwordFilePtr := flag.String("password_file", "/run/secrets/hm_registry_password", "Password of the docker registry")
	neededHeaderPtr := flag.String("needed_header", "X-TAC-User", "Header that needs to be set for requests to be allowed")
	portPtr := flag.Int("port", 4181, "Port on which the application will run")
	flag.Parse()

	srv := server.NewServer(*neededHeaderPtr)

	usernameBytes, err := ioutil.ReadFile(*usernameFilePtr)
	if err != nil {
		log.WithError(err).Fatal("Could not read username file")
	}
	passwordBytes, err := ioutil.ReadFile(*passwordFilePtr)
	if err != nil {
		log.WithError(err).Fatal("Could not read username file")
	}

	_, err = managers.CreateRegistryManager(*urlPtr, string(usernameBytes), string(passwordBytes))
	if err != nil {
		log.WithError(err).Fatal("Failed to create registry manager")
	}

	// Start
	log.WithField("port", *portPtr).Info("Listening on specified port")
	log.Info(srv.Router.Run(":" + strconv.Itoa(*portPtr)))
}
