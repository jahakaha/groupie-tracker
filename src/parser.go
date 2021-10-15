package src

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

//All is struct of Union
var All Union

//ParseJSON to parse all data
func ParseJSON() {
	urlArtist := "https://groupietrackers.herokuapp.com/api/artists"
	urlLocotaion := "https://groupietrackers.herokuapp.com/api/locations"
	urlDates := "https://groupietrackers.herokuapp.com/api/dates"
	urlRelation := "https://groupietrackers.herokuapp.com/api/relation"

	Parse(urlArtist, &All.Artists)
	Parse(urlLocotaion, &All.Locations)
	Parse(urlDates, &All.Dates)
	Parse(urlRelation, &All.Relation)

}

// Parse data from json
func Parse(url string, strct interface{}) {
	res, err := http.Get(url)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}
	json.Unmarshal(body, &strct)
}
