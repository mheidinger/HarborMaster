package main

import (
	"os"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"

	"HarborMaster/managers"
	"HarborMaster/server"
)

func getEnvOrDefault(key, def string) string {
	val, exists := os.LookupEnv(key)
	if exists {
		return val
	}
	return def
}

// getSecretValue implements the precedence logic for username/password:
// 1. Direct value env var (e.g., REGISTRY_USERNAME)
// 2. File env var (e.g., REGISTRY_USERNAME_FILE)
func getSecretValue(valueEnv, fileEnv string) *string {
	val, valExists := os.LookupEnv(valueEnv)
	if valExists {
		return &val
	}
	filePath, filePathExists := os.LookupEnv(fileEnv)
	if filePathExists {
		data, err := os.ReadFile(filePath)
		if err != nil {
			log.WithError(err).Warnf("Could not read file for %s, falling back to default", valueEnv)
			return nil
		}
		trimmedData := strings.TrimSpace(string(data))
		return &trimmedData
	}
	return nil
}

func main() {
	// Defaults
	defaultURL := "localhost:8080"
	defaultHeader := "X-TAC-User"
	defaultPort := 4181

	// Read config from env or fallback
	url := getEnvOrDefault("REGISTRY_URL", defaultURL)
	neededHeader := getEnvOrDefault("NEEDED_HEADER", defaultHeader)
	portStr := getEnvOrDefault("PORT", strconv.Itoa(defaultPort))

	// Username logic
	username := getSecretValue("REGISTRY_USERNAME", "REGISTRY_USERNAME_FILE")
	if username == nil {
		log.Fatal("No username provided")
	}

	// Password logic
	password := getSecretValue("REGISTRY_PASSWORD", "REGISTRY_PASSWORD_FILE")
	if password == nil {
		log.Fatal("No password provided")
	}

	// Parse port
	port, err := strconv.Atoi(portStr)
	if err != nil {
		log.WithError(err).Warnf("Invalid PORT value '%s', using default %d", portStr, defaultPort)
		port = defaultPort
	}

	srv := server.NewServer(neededHeader)

	_, err = managers.CreateRegistryManager(url, *username, *password)
	if err != nil {
		log.WithError(err).Fatal("Failed to create registry manager")
	}

	// Start
	log.WithField("port", port).Info("Listening on specified port")
	log.Info(srv.Router.Run(":" + strconv.Itoa(port)))
}
