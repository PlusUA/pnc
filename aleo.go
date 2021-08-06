package main

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func getAleoData() {

	reader := strings.NewReader(`{"jsonrpc": "2.0", "id":"documentation", "method": "getnodeinfo", "params": [] }`)
	request, err := http.NewRequest("GET", aleoNodeURL, reader)
	checkErr(err)

	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(request)
	checkErr(err)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	checkErr(err)

	var nodeData AleoResponse

	json.Unmarshal(body, &nodeData)

	aleoData.IsBootnode = nodeData.Result.IsBootnode
	aleoData.IsMiner = nodeData.Result.IsMiner
	aleoData.IsSyncing = nodeData.Result.IsSyncing
	aleoData.Launched = nodeData.Result.Launched
	aleoData.Version = nodeData.Result.Version

}

func aleoHandler(w http.ResponseWriter, r *http.Request) {

	getAleoData()

	if r.Method != "POST" {
		t, err := template.ParseFiles("templates/aleo.html", "templates/header.html", "templates/footer.html")
		if err != nil {
			log.Fatal(w, err.Error())
			return
		}
		t.ExecuteTemplate(w, "aleo", aleoData)
		return

	}

}
