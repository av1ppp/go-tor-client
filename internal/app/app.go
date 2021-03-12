package app

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/AviParampampam/go-tor-client/internal/torclient"
	"github.com/AviParampampam/go-tor-client/internal/torsites"
	"github.com/BurntSushi/toml"
)

func Run() {
	// Initialization.
	torClient, err := torclient.New(9050)
	if err != nil {
		log.Fatal(err)
	}

	torSites := torsites.TorSites{}
	if _, err := toml.DecodeFile("torsites.toml", &torSites); err != nil {
		log.Fatal(err)
	}

	proPublica := torSites.Name("ProPublica")
	if proPublica == nil {
		log.Fatal("proPublica site not found")
	}

	// Send request.
	resp, err := torClient.Get(proPublica.URL)
	if err != nil {
		log.Fatal(err)
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	bodyString := string(bodyBytes)

	fmt.Println(bodyString)
}
