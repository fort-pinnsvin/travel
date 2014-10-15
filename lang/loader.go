package lang

import (
	"github.com/robfig/config"
	"fmt"
)

var langs map[string]*Config

func loadLanguage(id string) {
	c, err := config.ReadDefault("langs/" + id + ".cfg")
	if err == nil {
		langs[id] = c
	}
}

func getString(id string, temp string, label string) string {
	if langs[id] != nil {
		result,_ := langs[id].String(temp, label)
		return result
	}
	return fmt.Sprintf("[%s: %s]", temp, label)
}
