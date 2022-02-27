package main

import (
	"encoding/json"
	"net/http"
)

func getMisi(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var posts []Misi
	var response Response
	db := connect()
	defer db.Close()
	result, err := db.Query("SELECT id, misi from misi")
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		var misi Misi
		err := result.Scan(&misi.ID, &misi.Misi)
		if err != nil {
			panic(err.Error())
		}
		posts = append(posts, misi)
	}
	response.Status = 1
	response.Message = "Success"
	response.Data = posts
	json.NewEncoder(w).Encode(response)
}
