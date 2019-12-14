package main
import (
	"github.com/tacoman/rosterkit/pkg/leagues"
	"flag"
	"strings"
	"fmt"
)

func main() {
	leagueFlag := flag.String("league", "", "Which league to scrape.")
	flag.Parse()
	league := strings.ToLower(*leagueFlag)
	if(league == "uws-midwest") {
		leagues.Scrape_uws_midwest()
	} else if(league == "") {
		fmt.Println("You must specify a league to use RosterKit! Currently supported leagues: uws-midwest")
	} else {
		fmt.Println("That league is not supported... yet! Currently supported leagues: uws-midwest")
	}
}