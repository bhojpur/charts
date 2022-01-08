package examples

import (
	"io"
	"os"

	"github.com/bhojpur/charts/pkg/components"
)

type PageFlexLayoutExamples struct{}

func (PageFlexLayoutExamples) Examples() {
	page := genPages()
	page.SetLayout(components.PageFlexLayout)
	f, err := os.Create("examples/html/page_flex_layout.html")
	if err != nil {
		panic(err)
	}
	page.Render(io.MultiWriter(f))
}
