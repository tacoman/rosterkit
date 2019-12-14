package main
import (
	"github.com/tacoman/rosterkit/pkg/leagues"
	"github.com/tacoman/rosterkit/pkg/sinks"
	"flag"
	"strings"
	"fmt"
	. "github.com/tacoman/rosterkit/pkg/models"
)

func main() {
	leagueFlag := flag.String("league", "", "Which league to scrape.")
	filename := flag.String("filename", "output.json", "The filename to output data to.")
	flag.Parse()
	league := strings.ToLower(*leagueFlag)
	if(league == "uws-midwest") {
		leagueChannel := make(chan []Foe, 1)
		leagues.Scrape_uws_midwest(leagueChannel)
		output := <- leagueChannel
		sinks.JsonFile(output, *filename)
	} else if(league == "") {
		fmt.Println("You must specify a league to use RosterKit! Currently supported leagues: uws-midwest")
	} else {
		fmt.Println("That league is not supported... yet! Currently supported leagues: uws-midwest")
	}
}