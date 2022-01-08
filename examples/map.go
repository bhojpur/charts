package examples

import (
	"io"
	"math/rand"
	"os"

	"github.com/bhojpur/charts/pkg/charts"
	"github.com/bhojpur/charts/pkg/components"
	"github.com/bhojpur/charts/pkg/opts"
)

var (
	baseMapData = []opts.MapData{
		{Name: "बिहार", Value: float64(rand.Intn(150))},
		{Name: "उत्तर प्रदेश", Value: float64(rand.Intn(150))},
		{Name: "दिल्ली", Value: float64(rand.Intn(150))},
		{Name: "पंजाब", Value: float64(rand.Intn(150))},
		{Name: "पश्चिम बंगाल", Value: float64(rand.Intn(150))},
		{Name: "महाराष्ट्र", Value: float64(rand.Intn(150))},
		{Name: "तमिल नाडु", Value: float64(rand.Intn(150))},
		{Name: "गुजरात", Value: float64(rand.Intn(150))},
		{Name: "मणिपुर", Value: float64(rand.Intn(150))},
	}

	biharMapData = map[string]float64{
		"भोजपुर":     float64(rand.Intn(150)),
		"पटना":       float64(rand.Intn(150)),
		"गया":        float64(rand.Intn(150)),
		"मधुबनी":     float64(rand.Intn(150)),
		"समस्तीपुर":  float64(rand.Intn(150)),
		"मुज़फ्फरपुर": float64(rand.Intn(150)),
		"भागलपुर":    float64(rand.Intn(150)),
		"छपरा":       float64(rand.Intn(150)),
		"बक्सर":      float64(rand.Intn(150)),
	}
)

func generateMapData(data map[string]float64) (items []opts.MapData) {
	items = make([]opts.MapData, 0)
	for k, v := range data {
		items = append(items, opts.MapData{Name: k, Value: v})
	}
	return
}

func mapBase() *charts.Map {
	mc := charts.NewMap()
	mc.RegisterMapType("india")
	mc.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{Title: "Basic Map example"}),
	)

	mc.AddSeries("map", baseMapData)
	return mc
}

func mapShowLabel() *charts.Map {
	mc := charts.NewMap()
	mc.RegisterMapType("india")
	mc.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{Title: "Show label"}),
	)

	mc.AddSeries("map", baseMapData).
		SetSeriesOptions(
			charts.WithLabelOpts(opts.Label{
				Show: true,
			}),
		)
	return mc
}

func mapVisualMap() *charts.Map {
	mc := charts.NewMap()
	mc.RegisterMapType("india")
	mc.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "Visual Map",
		}),
		charts.WithVisualMapOpts(opts.VisualMap{
			Calculable: true,
		}),
	)

	mc.AddSeries("map", baseMapData)
	return mc
}

func mapRegion() *charts.Map {
	mc := charts.NewMap()
	mc.RegisterMapType("बिहार")
	mc.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "Bihar state",
		}),
		charts.WithVisualMapOpts(opts.VisualMap{
			Calculable: true,
			InRange:    &opts.VisualMapInRange{Color: []string{"#50a3ba", "#eac736", "#d94e5d"}},
		}),
	)

	mc.AddSeries("map", generateMapData(biharMapData))
	return mc
}

func mapTheme() *charts.Map {
	mc := charts.NewMap()
	mc.RegisterMapType("india")
	mc.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{
			Theme: "macarons",
		}),
		charts.WithTitleOpts(opts.Title{
			Title: "Map Theme",
		}),
		charts.WithVisualMapOpts(opts.VisualMap{
			Calculable: true,
			Max:        150,
		}),
	)

	mc.AddSeries("map", baseMapData)
	return mc
}

type MapExamples struct{}

func (MapExamples) Examples() {
	page := components.NewPage()
	page.AddCharts(
		mapBase(),
		mapShowLabel(),
		mapVisualMap(),
		mapRegion(),
		mapTheme(),
	)

	f, err := os.Create("examples/html/map.html")
	if err != nil {
		panic(err)
	}
	page.Render(io.MultiWriter(f))
}
