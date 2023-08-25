package renderer

import (
	"bytes"
	"fmt"
	"html/template"
	"io"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	chartrender "github.com/go-echarts/go-echarts/v2/render"
	"github.com/go-echarts/go-echarts/v2/types"
	log "github.com/sirupsen/logrus"

	database "github.com/myshkins/gopetwatch/database"
	"github.com/myshkins/gopetwatch/logger"
)

type snippetRenderer struct {
	c      interface{}
	before []func()
}

func newSnippetRenderer(c interface{}, before ...func()) chartrender.Renderer {
	return &snippetRenderer{c: c, before: before}
}

func (r *snippetRenderer) Render(w io.Writer) error {
	const tplName = "chart"
	for _, fn := range r.before {
		fn()
	}

	tpl := template.
		Must(template.New(tplName).
			Funcs(template.FuncMap{
				"safeJS": func(s interface{}) template.JS {
					return template.JS(fmt.Sprint(s))
				},
			}).
			Parse(baseTpl),
		)

	err := tpl.ExecuteTemplate(w, tplName, r.c)
	return err
}

func renderToHtml(c interface{}) template.HTML {
	var buf bytes.Buffer
	r := c.(chartrender.Renderer)
	err := r.Render(&buf)
	if err != nil {
		log.Printf("Failed to render chart: %s", err)
		return ""
	}

	return template.HTML(buf.String())
}

//renders the chart as a string of html
func RenderChart() (template.HTML) {
	// create a new line instance
	line := charts.NewLine()
	logger.Log.Info("charts.NewLine")
	// set some global options like Title/Legend/ToolTip or anything else
	line.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{
			Theme: types.ThemeWesteros}),
		charts.WithTitleOpts(opts.Title{
			Title:    "gopetwatch",
			Subtitle: "cool bruh",
		}),
	  charts.WithYAxisOpts(opts.YAxis{
			Max: 100, Min: 20}))
	logger.Log.Info("line.SetGlobalOptions")

	// Put data into instance
	temps, timestamps := generateLineData()
	logger.Log.Info("generateLineData")
	line.SetXAxis(timestamps).
		AddSeries("temps", temps).
		SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{Smooth: true}))
	logger.Log.Info("line.setXAxis")
	line.Renderer = newSnippetRenderer(line, line.Validate)
	logger.Log.Info("line.Renderer")
	var htmlSnippet template.HTML = renderToHtml(line)
	logger.Log.Info("renderToHtml")
	return htmlSnippet
}

func generateLineData() (temps, timestamps []opts.LineData) {
	temps = make([]opts.LineData, 0)
	timestamps = make([]opts.LineData, 0)
	readings := database.QueryReadings()

	for _, r := range readings {
		temps = append(temps, opts.LineData{Value: r.Temperature})
		timestamps = append(timestamps, opts.LineData{Value: r.ReadingTimestamp})
	}
	return temps, timestamps
}
