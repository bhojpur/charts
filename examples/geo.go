package examples

import (
	"io"
	"math/rand"
	"os"

	"github.com/bhojpur/charts/pkg/charts"
	"github.com/bhojpur/charts/pkg/components"
	"github.com/bhojpur/charts/pkg/opts"
	"github.com/bhojpur/charts/pkg/types"
)

var geoData = []opts.GeoData{
	{Name: "बिहार", Value: []float64{25.5938737, 84.566373, float64(rand.Intn(100))}},
	{Name: "उत्तर प्रदेश", Value: []float64{27.1233985, 78.6394457, float64(rand.Intn(100))}},
	{Name: "दिल्ली", Value: []float64{28.6466773, 76.8130731, float64(rand.Intn(100))}},
	{Name: "पंजाब", Value: []float64{31.013186, 74.2799456, float64(rand.Intn(100))}},
	{Name: "पश्चिम बंगाल", Value: []float64{24.356295, 85.6064378, float64(rand.Intn(100))}},
	{Name: "महाराष्ट्र", Value: []float64{18.8020086, 74.5308677, float64(rand.Intn(100))}},
}

var biharData = []opts.GeoData{
	{Name: "भोजपुर", Value: []float64{25.5576409, 84.652547, float64(rand.Intn(100))}},
	{Name: "पटना", Value: []float64{25.6080208, 85.0730026, float64(rand.Intn(100))}},
	{Name: "गया", Value: []float64{24.7833122, 84.9470613, float64(rand.Intn(100))}},
}

func geoBase() *charts.Geo {
	geo := charts.NewGeo()
	geo.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{Title: "Basic Geo example"}),
		charts.WithGeoComponentOpts(opts.GeoComponent{
			Map:       "india",
			ItemStyle: &opts.ItemStyle{Color: "#006666"},
		}),
	)

	geo.AddSeries("geo", types.ChartEffectScatter, geoData,
		charts.WithRippleEffectOpts(opts.RippleEffect{
			Period:    4,
			Scale:     6,
			BrushType: "stroke",
		}),
	)
	return geo
}

func geoGuangdong() *charts.Geo {
	geo := charts.NewGeo()
	geo.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{Title: "Bihar state"}),
		charts.WithGeoComponentOpts(opts.GeoComponent{
			Map: "बिहार",
		}),
		charts.WithVisualMapOpts(opts.VisualMap{
			Calculable: true,
			InRange: &opts.VisualMapInRange{
				Color: []string{"#50a3ba", "#eac736", "#d94e5d"},
			},
		}),
	)

	geo.AddSeries("geo", types.ChartScatter, biharData)
	return geo
}

type GeoExamples struct{}

func (GeoExamples) Examples() {
	page := components.NewPage()
	page.AddCharts(
		geoBase(),
		geoGuangdong(),
	)

	f, err := os.Create("examples/html/geo.html")
	if err != nil {
		panic(err)
	}
	page.Render(io.MultiWriter(f))
}
