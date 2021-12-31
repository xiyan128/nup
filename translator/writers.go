package translator

import (
	"github.com/yosssi/gohtml"
	"io"
	"strings"
)

type DocumentWriter interface {
	WritePreamble() (n int, err error)
	WriteString(s string) (n int, err error)
	WritePostamble() (n int, err error)
	Flush(writer io.Writer) (n int, err error)
}

type HtmlWriter struct {
	*strings.Builder
}

func NewHtmlWriter() *HtmlWriter {
	return &HtmlWriter{
		&strings.Builder{},
	}
}

func (H HtmlWriter) WritePreamble() (n int, err error) {
	n, err = H.WriteString(`
<!DOCTYPE html>
<html lang="en">
  <head>

	<title>Page Title</title>

	<meta name="viewport" content="width=device-width,initial-scale=1">
	<meta charset="UTF-8">

    <link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/katex@0.15.1/dist/katex.min.css" integrity="sha384-R4558gYOUz8mP9YWpZJjofhk+zx0AS11p36HnD2ZKj/6JR5z27gSSULCNHIRReVs" crossorigin="anonymous">

    <!-- The loading of KaTeX is deferred to speed up page rendering -->
    <script defer src="https://cdn.jsdelivr.net/npm/katex@0.15.1/dist/katex.min.js" integrity="sha384-z1fJDqw8ZApjGO3/unPWUPsIymfsJmyrDVWC8Tv/a1HeOtGmkwNd/7xUS0Xcnvsx" crossorigin="anonymous"></script>

    <!-- To automatically render math in text elements, include the auto-render extension: -->
    <script defer src="https://cdn.jsdelivr.net/npm/katex@0.15.1/dist/contrib/auto-render.min.js" integrity="sha384-+XBljXPPiv+OzfbB3cVmLHf4hdUFHlWNZN5spNQ7rmHTXpd7WvJum6fIACpNNfIR" crossorigin="anonymous"
        onload="renderMathInElement(document.body);"></script>
	<link rel="stylesheet" href="index.css">
  </head>

<body>
	<main>
`)
	return
}

func (H HtmlWriter) WritePostamble() (n int, err error) {
	n, err = H.WriteString(`
	</main>
	<script src="https://unpkg.com/prismjs@1.25.0/components/prism-core.min.js"></script>
	<script src="https://unpkg.com/prismjs@1.25.0/plugins/autoloader/prism-autoloader.min.js"></script>
</body>
</html> 
`)
	return
}

func (H HtmlWriter) Flush(writer io.Writer) (n int, err error) {
	gohtml.Condense = true
	gohtml.LineWrapColumn = 80
	gohtml.InlineTags["b"] = true
	gohtml.InlineTags["i"] = true
	gohtml.InlineTags["sup"] = true
	n, err = writer.Write([]byte(gohtml.Format(H.String())))
	return
}
