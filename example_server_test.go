package wkhtmltopdf_test

/* This example creates an http server, which returns a simple
   pdf document with a title and the path of the request.
*/

import (
	"bytes"
	"html/template"
	"log"
	"net/http"


	"github.com/mnugter/wkhtmltopdf-go"
)

const page = `
<html>
  <body>
    <h1>Test Page</h1>

	<p>Path: {{.}}</p>
  </body>
</html>`

func handler(w http.ResponseWriter, r *http.Request) {

	tmpl := template.Must(template.New("page").Parse(page))
	buf := &bytes.Buffer{}
	tmpl.Execute(buf, r.URL.String())

	doc := wkhtmltopdf.NewDocument()
	pg, err := wkhtmltopdf.NewPageReader(buf)
	if err != nil {
		log.Fatal("Error reading page buffer")
	}
	doc.AddPages(pg)

	w.Header().Set("Content-Type", "application/pdf")
	w.Header().Set("Content-Disposition", `attachment; filename="test.pdf"`)
	err = doc.Write(w)
	if err != nil {
		log.Fatal("Error serving pdf")
	}
}

func Example() {

	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)

}
