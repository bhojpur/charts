# Bhojpur Charts - Data Processing Engine
The Bhojpur Charts is a data visualization software-as-a-service product used as a primary Charting Engine within the Bhojpur.NET Platform ecosystem. It relies on the client-side libraries of the [Apache eCharts](https://echarts.apache.org/).

[![Build status](https://badge.buildkite.com/d310559b033d9dc2c41668b20c90e9c0c3d6543ce31c8792ae.svg)](https://buildkite.com/bhojpur/charts)

### Installation

To get the Bhojpur Charts

```shell
$ go get -u github.com/bhojpur/charts/...
```

OR

```shell
# go.mod

require github.com/bhojpur/charts
```

### ✨ Features

* Clean and comprehensive API.
* Visualize your data in 25+ different ways.
* Highly configurable chart options.
* Detailed documentation and a rich collection of examples.
* Visualize your geospatial data with 100+ maps.

### 📝 Usage

It's easy to get started with Bhojpur Charts. In this example, we create a simple bar chart with only a few lines of code.

```golang
package main

import (
	"math/rand"
	"os"

	"github.com/bhojpur/charts/pkg/charts"
	"github.com/bhojpur/charts/pkg/opts"
)

// generate random data for bar chart
func generateBarItems() []opts.BarData {
	items := make([]opts.BarData, 0)
	for i := 0; i < 7; i++ {
		items = append(items, opts.BarData{Value: rand.Intn(300)})
	}
	return items
}

func main() {
	// create a new bar instance
	bar := charts.NewBar()
	// set some global options like Title/Legend/ToolTip or anything else
	bar.SetGlobalOptions(charts.WithTitleOpts(opts.Title{
		Title:    "My first bar chart generated by Bhojpur Charts",
		Subtitle: "It's easy to use, right?",
	}))

	// Put data into instance
	bar.SetXAxis([]string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}).
		AddSeries("Category A", generateBarItems()).
		AddSeries("Category B", generateBarItems())
	// Where the magic happens
	f, _ := os.Create("bar.html")
	bar.Render(f)
}
```

And, the generated bar.html is rendered. We can also start a listening Web Server with net/http.

```golang
package main

import (
	"math/rand"
	"net/http"

	"github.com/bhojpur/charts/pkg/charts"
	"github.com/bhojpur/charts/pkg/opts"
	"github.com/bhojpur/charts/pkg/types"
)

// generate random data for line chart
func generateLineItems() []opts.LineData {
	items := make([]opts.LineData, 0)
	for i := 0; i < 7; i++ {
		items = append(items, opts.LineData{Value: rand.Intn(300)})
	}
	return items
}

func httpserver(w http.ResponseWriter, _ *http.Request) {
	// create a new line instance
	line := charts.NewLine()
	// set some global options like Title/Legend/ToolTip or anything else
	line.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{Theme: types.ThemeWesteros}),
		charts.WithTitleOpts(opts.Title{
			Title:    "Line example in Westeros theme",
			Subtitle: "Line chart rendered by the HTTP server this time",
		}))

	// Put data into instance
	line.SetXAxis([]string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}).
		AddSeries("Category A", generateLineItems()).
		AddSeries("Category B", generateLineItems()).
		SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{Smooth: true}))
	line.Render(w)
}

func main() {
	http.HandleFunc("/", httpserver)
	http.ListenAndServe(":8081", nil)
}
```

### 💡 Contributing

Bhojpur Charts is an open source project and built on the top of other open-source projects. We are always happy to have contributions, whether for typo fix, bug fix or big new features. Please do not ever hesitate to ask a question or send a pull request.

We strongly value documentation and integration with other projects so we are very glad to accept improvements for these aspects.
