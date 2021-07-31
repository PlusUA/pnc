package main

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

func casperHandler(w http.ResponseWriter, r *http.Request) {

	getCasperData()

	if r.Method != "POST" {
		t, err := template.ParseFiles("templates/casper.html", "templates/header.html", "templates/footer.html")
		if err != nil {
			log.Fatal(w, err.Error())
			return
		}
		t.ExecuteTemplate(w, "casper", a)
		return

	}

}

func getCasperData() {

	res, err := http.Get(casperNodeURL)
	checkErr(err)

	body, err := ioutil.ReadAll(res.Body)
	checkErr(err)

	var nodeData Response

	json.Unmarshal(body, &nodeData)

	a.APIVersion = nodeData.APIVersion
	a.BuildVersion = nodeData.BuildVersion
	a.ChainspecName = nodeData.ChainspecName
	a.OurPublicSigningKey = nodeData.OurPublicSigningKey
	a.StartingStateRootHash = nodeData.StartingStateRootHash
	a.Hash = nodeData.LastAddedBlockInfo.Hash
	a.Timestamp = nodeData.LastAddedBlockInfo.Timestamp
	a.EraID = nodeData.LastAddedBlockInfo.EraID
	a.Height = nodeData.LastAddedBlockInfo.Height
	a.StateRootHash = nodeData.LastAddedBlockInfo.StateRootHash
	a.Creator = nodeData.LastAddedBlockInfo.Creator
	a.ActivationPoint = nodeData.NextUpgrade.ActivationPoint
	a.ProtocolVersion = nodeData.NextUpgrade.ProtocolVersion

}
