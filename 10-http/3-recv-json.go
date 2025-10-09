package main

import (
	"encoding/json"
	"net/http"
)

// https://mholt.github.io/json-to-go/
/*
{
    "glossary": {
        "title": "example glossary",
		"GlossDiv": {
            "title": "S",
			"GlossList": {
                "GlossEntry": {
                    "ID": "SGML",
					"SortAs": "SGML",
					"GlossTerm": "Standard Generalized Markup Language",
					"Acronym": "SGML",
					"Abbrev": "ISO 8879:1986",
					"GlossDef": {
                        "para": "A meta-markup language, used to create markup languages such as DocBook.",
						"GlossSeeAlso": ["GML", "XML"]
                    },
					"GlossSee": "markup"
                }
            }
        }
    }
}
*/

func main() {
	// start http server
	// create one route to receive json data
	//
}

func processJson(w http.ResponseWriter, r *http.Request) {

	err := json.NewDecoder(r.Body).Decode(&data)

	w.Write([]byte(`{"status": "ok"}`))

}
