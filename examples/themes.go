package examples

import (
	"fmt"
	"io"
	"os"

	"github.com/bhojpur/charts/pkg/charts"
	"github.com/bhojpur/charts/pkg/components"
	"github.com/bhojpur/charts/pkg/opts"
	"github.com/bhojpur/charts/pkg/types"
)

func barWithTheme(theme string) *charts.Bar {
	bar := charts.NewBar()
	bar.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{Theme: theme}),
		charts.WithTitleOpts(opts.Title{Title: fmt.Sprintf("Bar with %s theme", theme)}),
	)
	bar.SetXAxis(weeks).
		AddSeries("Category A", generateBarItems()).
		AddSeries("Category B", generateBarItems())
	return bar
}

func themeChalk() *charts.Bar {
	return barWithTheme(types.ThemeChalk)
}

func themeEssos() *charts.Bar {
	return barWithTheme(types.ThemeEssos)
}

func themeInfographic() *charts.Bar {
	return barWithTheme(types.ThemeInfographic)
}

func themeMacarons() *charts.Bar {
	return barWithTheme(types.ThemeMacarons)
}

func themePurplePassion() *charts.Bar {
	return barWithTheme(types.ThemePurplePassion)
}

func themeRoma() *charts.Bar {
	return barWithTheme(types.ThemeRoma)
}

func themeRomantic() *charts.Bar {
	return barWithTheme(types.ThemeRomantic)
}

func themeShine() *charts.Bar {
	return barWithTheme(types.ThemeShine)
}

func themeVintage() *charts.Bar {
	return barWithTheme(types.ThemeVintage)
}

func themeWalden() *charts.Bar {
	return barWithTheme(types.ThemeWalden)
}

func themeWesteros() *charts.Bar {
	return barWithTheme(types.ThemeWesteros)
}

func themeWonderland() *charts.Bar {
	return barWithTheme(types.ThemeWonderland)
}

type ThemeExamples struct{}

func (ThemeExamples) Examples() {
	page := components.NewPage()
	page.AddCharts(
		themeChalk(),
		themeEssos(),
		themeInfographic(),
		themeMacarons(),
		themePurplePassion(),
		themeRoma(),
		themeRomantic(),
		themeShine(),
		themeVintage(),
		themeWalden(),
		themeWesteros(),
		themeWonderland(),
	)
	f, err := os.Create("examples/html/themes.html")
	if err != nil {
		panic(err)
	}
	page.Render(io.MultiWriter(f))
}
