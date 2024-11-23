// This file contains all the handle functions
// for this backend application.

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

// structure for User Data
type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func root(w http.ResponseWriter, r *http.Request) {
	sEnv := envData.envName
	fmt.Fprintln(w, "Welcome to Docker Basic Assignment's backend application running on "+sEnv)
}

func checkDB(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*") // Add CORS header
	w.WriteHeader(http.StatusOK)

	err := checkDBConnection()
	if err != nil {
		fmt.Fprintf(w, "Unable to connect to database.\nError Message:\n %s", err.Error())
		return
	}

	fmt.Fprintf(w, "Database is connected.")
}

// to add user, this function is called from frontend
// It connencts to database based on the configuration (specific to environments)
func addUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		w.Header().Set("Access-Control-Allow-Origin", "*") // Add CORS header

		var user User
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		strData, err := writeUserInfoIntoDB(user.ID, user.Name)
		if err != nil {
			fmt.Fprintf(w, "Unable to add user into database.\nError Message: %s", err.Error())
		} else {
			fmt.Fprintln(w, strData)
		}

		//fmt.Fprintf(w, "User added: ID = %s, Name = %s", user.ID, user.Name)
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

// to get user data, this function is called from frontend
// It connencts to database based on the configuration (specific to environments)
// and fetches data from database for provided specific user id and all data in one go.
func getUserInfo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*") // Add CORS header
	w.WriteHeader(http.StatusOK)

	if r.Method == http.MethodGet {
		val := r.URL.Query()
		if val.Has("id") {
			id, err := strconv.Atoi(val.Get("id"))
			if err != nil {

			}

			strData, err := readUserInfo(id)
			if err != nil {
				fmt.Fprintf(w, "Unable to retrive UserInfo.\nError Message: %s", err.Error())
			} else {
				fmt.Fprintln(w, strData)
			}
		} else {
			strData, err := readDatafromDB()
			if err != nil {
				fmt.Fprintf(w, "Unable to retrive UserInfo.\nError Message: %s", err.Error())
			}

			fmt.Fprintln(w, strData)
		}
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}