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
		t.ExecuteTemplate(w, "casper", casperData)
		return

	}

}

func getCasperData() {

	res, err := http.Get(casperNodeURL)
	checkErr(err)

	body, err := ioutil.ReadAll(res.Body)
	checkErr(err)

	var nodeData CasperResponse

	json.Unmarshal(body, &nodeData)

	casperData.APIVersion = nodeData.APIVersion
	casperData.BuildVersion = nodeData.BuildVersion
	casperData.ChainspecName = nodeData.ChainspecName
	casperData.OurPublicSigningKey = nodeData.OurPublicSigningKey
	casperData.StartingStateRootHash = nodeData.StartingStateRootHash
	casperData.Hash = nodeData.LastAddedBlockInfo.Hash
	casperData.Timestamp = nodeData.LastAddedBlockInfo.Timestamp
	casperData.EraID = nodeData.LastAddedBlockInfo.EraID
	casperData.Height = nodeData.LastAddedBlockInfo.Height
	casperData.StateRootHash = nodeData.LastAddedBlockInfo.StateRootHash
	casperData.Creator = nodeData.LastAddedBlockInfo.Creator
	casperData.ActivationPoint = nodeData.NextUpgrade.ActivationPoint
	casperData.ProtocolVersion = nodeData.NextUpgrade.ProtocolVersion

}
