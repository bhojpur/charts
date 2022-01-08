package examples

import (
	"io"
	"os"

	"github.com/bhojpur/charts/pkg/charts"
	"github.com/bhojpur/charts/pkg/components"
	"github.com/bhojpur/charts/pkg/opts"
)

var TreeNodes = []*opts.TreeData{
	{
		Name: "Node1",
		Children: []*opts.TreeData{
			{
				Name: "Child1",
			},
		},
	},
	{
		Name: "Node2",
		Children: []*opts.TreeData{
			{
				Name: "Child1",
			},
			{
				Name: "Child2",
			},
			{
				Name: "Child3",
			},
		},
	},
	{
		Name:      "Node3",
		Collapsed: true,
		Children: []*opts.TreeData{
			{
				Name: "Child1",
			},
			{
				Name: "Child2",
			},
			{
				Name: "Child3",
			},
		},
	},
}

var Tree = []opts.TreeData{
	{
		Name:     "Root",
		Children: TreeNodes,
	},
}

func treeBase() *charts.Tree {
	graph := charts.NewTree()
	graph.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{Width: "100%", Height: "95vh"}),
		charts.WithTitleOpts(opts.Title{Title: "Basic Tree example"}),
		charts.WithTooltipOpts(opts.Tooltip{Show: false}),
	)
	graph.AddSeries("tree", Tree).
		SetSeriesOptions(
			charts.WithTreeOpts(
				opts.TreeChart{
					Layout:           "orthogonal",
					Orient:           "LR",
					InitialTreeDepth: -1,
					Leaves: &opts.TreeLeaves{
						Label: &opts.Label{Show: true, Position: "right", Color: "Black"},
					},
				},
			),
			charts.WithLabelOpts(opts.Label{Show: true, Position: "top", Color: "Black"}),
		)
	return graph
}

type TreeExamples struct{}

func (TreeExamples) Examples() {
	page := components.NewPage()
	page.AddCharts(
		treeBase(),
	)

	f, err := os.Create("examples/html/tree.html")
	if err != nil {
		panic(err)

	}
	page.Render(io.MultiWriter(f))
}
