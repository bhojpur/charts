package examples

import (
	"fmt"
	"io"
	"math/rand"
	"os"

	"github.com/bhojpur/charts/pkg/charts"
	"github.com/bhojpur/charts/pkg/components"
	"github.com/bhojpur/charts/pkg/opts"
)

func gaugeBase() *charts.Gauge {
	gauge := charts.NewGauge()
	gauge.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{Title: "Basic Gauge example"}),
	)

	gauge.AddSeries("ProjectA", []opts.GaugeData{{Name: "Work progress", Value: rand.Intn(50)}})
	return gauge
}

func gaugeTimer() *charts.Gauge {
	gauge := charts.NewGauge()
	gauge.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{Title: "Javascript Timer"}),
	)

	gauge.AddSeries("ProjectB", []opts.GaugeData{{Name: "Work progress", Value: rand.Intn(50)}})

	fn := fmt.Sprintf(`setInterval(function () {
			option_%s.series[0].data[0].value = (Math.random() * 100).toFixed(2) - 0;
			goecharts_%s.setOption(option_%s, true);
		}, 2000);`, gauge.ChartID, gauge.ChartID, gauge.ChartID)
	gauge.AddJSFuncs(fn)
	return gauge
}

type GaugeExamples struct{}

func (GaugeExamples) Examples() {
	page := components.NewPage()
	page.AddCharts(
		gaugeBase(),
		gaugeTimer(),
	)

	f, err := os.Create("examples/html/gauge.html")
	if err != nil {
		panic(err)
	}
	page.Render(io.MultiWriter(f))
}
