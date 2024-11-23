package main

import (
	"fmt"
	"net/http"
	"os"
	"runtime"

	"github.com/joho/godotenv"
)

// Structure of environment data
type envVarData struct {
	envName, port, backend_port string
}

var envData envVarData

func main() {

	// Load Environment variable data
	initData(&envData)

	http.HandleFunc("/", getRoot)
	http.HandleFunc("/addUser", addUser)
	http.HandleFunc("/getUserInfo", getUserInfo)

	fmt.Println("Listening at port: ", envData.port)
	http.ListenAndServe(":"+envData.port, nil)
}

// check the OS
func isWindowsOS() bool {
	return runtime.GOOS == "windows"
}

// Data will be read from .env file on windows and from env variables on ubuntu
func initData(envData *envVarData) {
	if isWindowsOS() {
		envFile, _ := godotenv.Read("..\\config\\.env")
		envData.envName = envFile["ENV_NAME"]
		envData.port = envFile["PORT"]
		envData.backend_port = envFile["BACKEND_PORT"]
	} else {
		envData.envName = os.Getenv("ENV_NAME")
		envData.port = os.Getenv("PORT")
		envData.backend_port = os.Getenv("BACKEND_PORT")
	}
}
