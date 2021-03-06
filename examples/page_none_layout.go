package examples

import (
	"io"
	"os"

	"github.com/bhojpur/charts/pkg/components"
)

type PageNoneLayoutExamples struct{}

func (PageNoneLayoutExamples) Examples() {
	page := genPages()
	page.SetLayout(components.PageNoneLayout)
	f, err := os.Create("examples/html/page_none_layout.html")
	if err != nil {
		panic(err)
	}
	page.Render(io.MultiWriter(f))
}
