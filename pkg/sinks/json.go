package sinks

import (
	"fmt"
	"encoding/json"
	"os"
	. "github.com/tacoman/rosterkit/pkg/models"
)

func JsonFile(foes []Foe, filename string) {
	f, _ := os.Create(filename)
	defer f.Close()
	b, err := json.Marshal(foes)
	if err != nil {
		fmt.Println("error:", err)
	}
	f.Write(b)
	f.Sync()
}