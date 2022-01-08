package examples

import (
	"io"
	"os"

	"github.com/bhojpur/charts/pkg/charts"
	"github.com/bhojpur/charts/pkg/components"
	"github.com/bhojpur/charts/pkg/opts"
)

var wcData = map[string]interface{}{
	"Ram Chandra Rai": 10000,
	"Sudama":          6181,
	"Radha Krishna":   4386,
	"Sri Krishna":     4055,
	"Hari Krishna":    2467,
	"Babban":          2244,
	"Rajendra Prasad": 1898,
	"Narendra Kumar":  1484,
	"Surendra Kumar":  1689,
	"Nibhaya Dihra":   1112,
	"Shankar Dayal":   985,
	"Ravindra Kumar":  847,
	"Kalawati":        582,
	"Jamwanti":        555,
	"Pramila Kumari":  550,
	"Savita Kumari":   462,
	"Indu Kumari":     366,
	"Jai Shankar":     282,
	"Tarun Kumar":     273,
	"Shakuntala":      265,
}

func generateWCData(data map[string]interface{}) (items []opts.WordCloudData) {
	items = make([]opts.WordCloudData, 0)
	for k, v := range data {
		items = append(items, opts.WordCloudData{Name: k, Value: v})
	}
	return
}

func wcBase() *charts.WordCloud {
	wc := charts.NewWordCloud()
	wc.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "Basic WordCloud example",
		}))

	wc.AddSeries("wordcloud", generateWCData(wcData)).
		SetSeriesOptions(
			charts.WithWorldCloudChartOpts(
				opts.WordCloudChart{
					SizeRange: []float32{14, 80},
				}),
		)
	return wc
}

func wcCardioid() *charts.WordCloud {
	wc := charts.NewWordCloud()
	wc.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{Title: "Cardioid shape"}),
	)

	wc.AddSeries("wordcloud", generateWCData(wcData)).
		SetSeriesOptions(
			charts.WithWorldCloudChartOpts(
				opts.WordCloudChart{
					SizeRange: []float32{14, 80},
					Shape:     "cardioid",
				}),
		)
	return wc
}

func wcStar() *charts.WordCloud {
	wc := charts.NewWordCloud()
	wc.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "Star shape",
		}))

	wc.AddSeries("wordcloud", generateWCData(wcData)).
		SetSeriesOptions(
			charts.WithWorldCloudChartOpts(
				opts.WordCloudChart{
					SizeRange: []float32{14, 80},
					Shape:     "star",
				}),
		)
	return wc
}

type WordcloudExamples struct{}

func (WordcloudExamples) Examples() {
	page := components.NewPage()
	page.AddCharts(
		wcBase(),
		wcCardioid(),
		wcStar(),
	)

	f, err := os.Create("examples/html/wordcloud.html")
	if err != nil {
		panic(err)
	}
	page.Render(io.MultiWriter(f))
}
