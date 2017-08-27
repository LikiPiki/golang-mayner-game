package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
)

const (
	PORT = ":3000"
)

func getFirstVarInQuery(route string, path string) string {
	return strings.Split(strings.Trim(path, route), "/")[0]
}

type JsonUsers struct {
	Len   int    `json:"len"`
	Users []User `json:"users"`
}

type JsonCrypto struct {
	Crypto []CryptoValue `json:"cryptos"`
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	variable := getFirstVarInQuery("/get_users/", r.URL.Path)
	count, err := strconv.Atoi(variable)
	if err != nil {
		w.Write([]byte("Error request! enter number"))
		return
	}
	var users []User
	db.Limit(count).Find(&users)
	//TODO: for raketa only
	jsonify, err := json.MarshalIndent(&JsonUsers{
		len(users),
		users,
	}, "", "  ")
	if err != nil {
		log.Println(err)
	}

	w.Write(jsonify)
}

func getTop(w http.ResponseWriter, r *http.Request) {
	variable := getFirstVarInQuery("/get_top/", r.URL.Path)
	count, err := strconv.Atoi(variable)
	if err != nil {
		w.Write([]byte("Error request! enter number"))
		return
	}
	var users []User
	db.Limit(count).Order("score desc").Find(&users)
	//TODO: for raketa only
	jsonify, err := json.MarshalIndent(&JsonUsers{
		len(users),
		users,
	}, "", "  ")
	if err != nil {
		log.Println(err)
	}

	w.Write(jsonify)
}

func getCrypto(w http.ResponseWriter, r *http.Request) {
	var vals []CryptoValue
	db.Find(&vals)
	//TODO: for raketa only
	jsonify, err := json.MarshalIndent(&JsonCrypto{vals}, "", "  ")
	if err != nil {
		log.Println(err)
	}
	w.Write(jsonify)
}

func startAPI() {
	http.HandleFunc("/get_users/", getUsers)
	http.HandleFunc("/get_crypto/", getCrypto)
	http.HandleFunc("/get_top/", getTop)

	http.ListenAndServe(PORT, nil)
}
