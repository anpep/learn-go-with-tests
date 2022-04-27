package hello

import "fmt"

var helloFormats = map[string]string{
	"english": "Hello, %s",
	"spanish": "Hola %s",
	"french":  "Bonjour %s",
}

var worlds = map[string]string{
	"english": "world",
	"spanish": "mundo",
	"french":  "monde",
}

func Hello(name string, lang string) string {
	if _, doesLangExist := helloFormats[lang]; !doesLangExist {
		lang = "english"
	}
	if name == "" {
		name = worlds[lang]
	}
	return fmt.Sprintf(helloFormats[lang], name)
}
