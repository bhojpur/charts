package charts

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

const host = "https://charts.bhojpur.net/assets/"

func TestBarAssetsBeforeRender(t *testing.T) {
	bar := NewBar()
	assert.Equal(t, bar.JSAssets.Values, []string{"echarts.min.js"})
}

func TestBarAssetsAfterRender(t *testing.T) {
	bar := NewBar()
	err := bar.Render(ioutil.Discard)
	assert.NoError(t, err)
	assert.Equal(t, []string{host + "echarts.min.js"}, bar.JSAssets.Values)
}

func TestBarDefaultValue(t *testing.T) {
	bar := NewBar()
	err := bar.Render(ioutil.Discard)
	assert.NoError(t, err)
	assert.Equal(t, "900px", bar.Initialization.Width)
	assert.Equal(t, "500px", bar.Initialization.Height)
	assert.Equal(t, "Bhojpur Charts", bar.PageTitle)
	assert.Equal(t, host, bar.AssetsHost)
}
