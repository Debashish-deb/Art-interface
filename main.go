package main

import (
	"fmt"
	"html/template"
	"net/http"
	"unicode"
)

var resultTemplate = template.Must(template.New("result").Parse(resultPageTemplate))

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/encode", encodeHandler)
	http.HandleFunc("/decode", decodeHandler)

	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, homePageTemplate)
}

func encodeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		r.ParseForm()
		input := r.FormValue("text")
		encoded := Encoder(input)
		data := map[string]string{
			"Title":  "Result: Encoded",
			"Result": encoded,
		}
		if err := resultTemplate.Execute(w, data); err != nil {
			http.Error(w, "Failed to generate response", http.StatusInternalServerError)
		}
	}
}

func decodeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		r.ParseForm()
		input := r.FormValue("text")
		decoded, err := Decoder(input)
		if err != nil {
			http.Error(w, "Malformed encoded string", http.StatusBadRequest)
			return
		}
		data := map[string]string{
			"Title":  "Result: Decoded",
			"Result": decoded,
		}
		w.WriteHeader(http.StatusAccepted)
		err = resultTemplate.Execute(w, data)
		if err != nil {
			http.Error(w, "Failed to generate response", http.StatusInternalServerError)
		}
	}
}

//To remove non printable characters.
func removeNonPrintables(input string) string {
	printable := ""
	for _, r := range input {
		if unicode.IsPrint(r) {
			printable += string(r)
		}
	}
	return printable
}
