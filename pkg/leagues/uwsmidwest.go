package leagues

import (
	"fmt"
	"github.com/gocolly/colly"
	"strings"
	. "github.com/tacoman/rosterkit/pkg/models"
)

func handleFoe(foe FoeDef, outputChannel chan Foe) {
	rosterCollector := colly.NewCollector(
		colly.AllowedDomains("www.uwssoccer.com"),
	)			
	playerCollector := colly.NewCollector(
		colly.AllowedDomains("www.uwssoccer.com"),
	)
	
	rosterCollector.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting roster", r.URL.String())
	})

	playerCollector.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting player", r.URL.String())
	})

	rosterCollector.OnHTML("tr", func(e *colly.HTMLElement) {
		e.ForEach(".name > a[href]", func(_ int, el *colly.HTMLElement) { 
			playerCollector.Visit(el.Attr("href"))
		})
	})

	players := make([]Player, 0)

	playerCollector.OnHTML("body", func(e *colly.HTMLElement) {
		playerData:= e.ChildText(".playerName")
		dataPieces:= strings.Split(playerData, "\n")
		player:= Player{}
		player.Name = dataPieces[0];
		if(len(dataPieces) > 1) {
			positionInfo := strings.Split(dataPieces[1], "·")
			player.SquadNumber = strings.Trim(positionInfo[0], " #")
			if(len(positionInfo) > 1) {
				player.Position = positionInfo[1]
			} else {
				player.Position = ""
			}
		} else {
			player.SquadNumber = "0"
			player.Position = "";
		}
		players = append(players, player)
	})

	rosterCollector.Visit(foe.Url)
	foeOutput := Foe{Opponent: foe.Name, Players: players}
	outputChannel <- foeOutput
}

func Scrape_uws_midwest(leagueChannel chan []Foe) {
	FoeDefs := make([]FoeDef, 6)
	FoeDefs[0] = FoeDef{Url: "https://www.uwssoccer.com/roster/show/4813346?subseason=590879", Name: "AFC Ann Arbor"}
	FoeDefs[1] = FoeDef{Url: "https://www.uwssoccer.com/roster/show/4813292?subseason=590879", Name: "Detroit Sun"}
	FoeDefs[2] = FoeDef{Url: "https://www.uwssoccer.com/roster/show/4813296?subseason=590879", Name: "Grand Rapids FC"}
	FoeDefs[3] = FoeDef{Url: "https://www.uwssoccer.com/roster/show/4813297?subseason=590879", Name: "Indiana Union"}
	FoeDefs[4] = FoeDef{Url: "https://www.uwssoccer.com/roster/show/4813298?subseason=590879", Name: "Lansing United"}
	FoeDefs[5] = FoeDef{Url: "https://www.uwssoccer.com/roster/show/4813299?subseason=590879", Name: "Michigan Legends"}

	// On every a element which has href attribute call callback
	outputChannels := make([]chan Foe, 0)
	for _, foe := range FoeDefs {
		outputChannel := make(chan Foe, 1)
		go handleFoe(foe, outputChannel)
		outputChannels = append(outputChannels, outputChannel)
	}


	foesList := make([]Foe, len(FoeDefs))
    for idx, value := range outputChannels {
	   output := <- value
       foesList[idx] = output
	}
	leagueChannel <- foesList
}