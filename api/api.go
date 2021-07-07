package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/dmaceasistemas/go-backend-dev-login/helpers"
	"github.com/dmaceasistemas/go-backend-dev-login/interfaces"
	"github.com/dmaceasistemas/go-backend-dev-login/users"
	"github.com/gorilla/mux"
)

type Login struct {
	Username string
	Password string
}

type Register struct {
	Username string
	Email    string
	Password string
}

type ErrResponse struct {
	Message string
}

// Create readBody function
func readBody(req *http.Request) []byte {
	body, err := ioutil.ReadAll(req.Body)
	helpers.HandleErr(err)

	return body
}

//
func apiResponse(call map[string]interface{}, w http.ResponseWriter) {
	if call["message"] == "all is fine" {
		resp := call
		json.NewEncoder(w).Encode(resp)
		// Handle error in else
	} else {
		resp := interfaces.ErrResponse{Message: "Wrong username or password"}
		json.NewEncoder(w).Encode(resp)
	}
}

func login(w http.ResponseWriter, req *http.Request) {
	// Ready body
	/*body, err := ioutil.ReadAll(req.Body)
	helpers.HandleErr(err)*/
	body := readBody(req)

	// Handle Login
	var formattedBody Login
	err := json.Unmarshal(body, &formattedBody)
	helpers.HandleErr(err)

	login := users.Login(formattedBody.Username, formattedBody.Password)

	// Prepare response
	/*if login["message"] == "all is fine" {
		resp := login
		json.NewEncoder(w).Encode(resp)
	} else {
		// Handle error
		resp := ErrResponse{Message: "Wrong username or password"}
		json.NewEncoder(w).Encode(resp)
	} */
	apiResponse(login, w)
}

func register(w http.ResponseWriter, req *http.Request) {
	// Ready body
	/*body, err := ioutil.ReadAll(req.Body)
	helpers.HandleErr(err)*/

	body := readBody(req)

	// Handle Register
	var formattedBody Register
	err := json.Unmarshal(body, &formattedBody)
	helpers.HandleErr(err)
	register := users.Register(formattedBody.Username, formattedBody.Email, formattedBody.Password)

	// Prepare response
	/*if register["message"] == "all is fine" {
		resp := register
		json.NewEncoder(w).Encode(resp)
	} else {
		// Handle error
		resp := ErrResponse{Message: "Wrong username or password"}
		json.NewEncoder(w).Encode(resp)
	} */
	apiResponse(register, w)

}

func getUser(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	userId := vars["id"]
	log.Println(userId)
	auth := req.Header.Get("Authorization")
	log.Println(auth)

	user := users.GetUser(userId, auth)
	apiResponse(user, w)
}

func StartApi() {
	router := mux.NewRouter()
	router.Use(helpers.PanicHandler)
	router.HandleFunc("/login", login).Methods("POST")
	router.HandleFunc("/register", register).Methods("POST")
	router.HandleFunc("/user/{id}", getUser).Methods("GET")
	fmt.Println("App is working on port 8090")
	log.Fatal(http.ListenAndServe(":8090", router))
}
