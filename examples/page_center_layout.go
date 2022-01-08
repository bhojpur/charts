package examples

import (
	"io"
	"os"

	"github.com/bhojpur/charts/pkg/components"
)

func genPages() *components.Page {
	page := components.NewPage()
	page.AddCharts(
		lineBase(),
		pieRadius(),
		barXYName(),
		barColor(),
		klineDataZoomInside(),
		parallelComponent(),
		barGap(),
		radarBase(),
		bar3DAutoRotate(),
		gaugeBase(),
		heatMapBase(),
		barMarkPoints(),
		scatterShowLabel(),
		liquidDiamond(),
		barOverlap(),
		mapShowLabel(),
	)
	return page
}

type PageCenterLayoutExamples struct{}

func (PageCenterLayoutExamples) Examples() {
	page := genPages()
	f, err := os.Create("examples/html/page_center_layout.html")
	if err != nil {
		panic(err)
	}
	page.Render(io.MultiWriter(f))
}
