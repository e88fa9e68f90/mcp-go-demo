package mcp_server

import (
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	chart "github.com/wcharczuk/go-chart/v2"
	"math/rand"
	"net/http"
	"os"
)

func CreateBarChart(w http.ResponseWriter, _ *http.Request) {
	// create a new bar instance
	bar := charts.NewBar()

	// Set global options
	bar.SetGlobalOptions(charts.WithTitleOpts(opts.Title{
		Title:    "Bar chart in Go",
		Subtitle: "This is fun to use!",
	}))

	// Put data into instance
	bar.SetXAxis([]string{"Jan", "Feb", "Mar", "Apr", "May", "Jun"}).
		AddSeries("Category A", generateBarItems()).
		AddSeries("Category B", generateBarItems())
	bar.Render(w)
}
func generateBarItems() []opts.BarData {
	items := make([]opts.BarData, 0)
	for i := 0; i < 6; i++ {
		items = append(items, opts.BarData{Value: rand.Intn(500)})
	}
	return items
}

func generatePieItems() []opts.PieData {
	subjects := []string{"数学", "英语", "科学", "计算机", "历史", "地理"}
	items := make([]opts.PieData, 0)
	for i := 0; i < 6; i++ {
		items = append(items, opts.PieData{
			Name:  subjects[i],
			Value: rand.Intn(500)})
	}
	return items
}
func CreatePieChart(w http.ResponseWriter, _ *http.Request) {
	// create a new pie instance
	pie := charts.NewPie()
	pie.SetGlobalOptions(
		charts.WithTitleOpts(
			opts.Title{
				Title:    "Pie chart in Go",
				Subtitle: "This is fun to use!",
			},
		),
	)
	pie.SetSeriesOptions()
	pie.AddSeries("Monthly revenue",
		generatePieItems()).
		SetSeriesOptions(
			charts.WithPieChartOpts(
				opts.PieChart{
					Radius: 200,
				},
			),
			charts.WithLabelOpts(
				opts.Label{
					Show:      opts.Bool(true),
					Formatter: "{b}: {c}",
				},
			),
		)
	pie.Render(w)

}

func CreatePieChartToPNG() {
	pie := chart.PieChart{
		Width:  512,
		Height: 512,
		Values: []chart.Value{
			{Value: 5, Label: "Blue"},
			{Value: 5, Label: "Green"},
			{Value: 4, Label: "Gray"},
			{Value: 4, Label: "Orange"},
			{Value: 3, Label: "Deep Blue"},
			{Value: 3, Label: "sss"},
			{Value: 1, Label: "ddd"},
		},
	}

	f, _ := os.Create("output.png")
	defer f.Close()
	pie.Render(chart.PNG, f)
}
