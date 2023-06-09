package renderer

import (
  "bytes"
	"fmt"
  "html/template"
  "io"
  "net/http"

	chartrender "github.com/go-echarts/go-echarts/v2/render"
  log "github.com/sirupsen/logrus"
  "github.com/go-echarts/go-echarts/v2/charts"
  "github.com/go-echarts/go-echarts/v2/opts"
  "github.com/go-echarts/go-echarts/v2/types"

	database "github.com/myshkins/gopetwatch/database"
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

func RenderChart() template.HTML {
	// create a new line instance
	line := charts.NewLine()
	// set some global options like Title/Legend/ToolTip or anything else
	line.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{Theme: types.ThemeWesteros}),
		charts.WithTitleOpts(opts.Title{
			Title:    "gopetwatch",
			Subtitle: "cool bruh",
		}))

	// Put data into instance
	temps, timestamps := generateLineData()
	line.SetXAxis(timestamps).
		AddSeries("temps", temps).
		SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{Smooth: true}))
	line.Renderer = newSnippetRenderer(line, line.Validate)
	var htmlSnippet template.HTML = renderToHtml(line)
	return htmlSnippet
}

func FillTemplate(w http.ResponseWriter, s template.HTML) {
	type tmplData struct {
		Title string
		Snippet template.HTML
	}

	data := tmplData{
		Title: "gopetwatch",
		Snippet: s,
	}

	tmpl := template.Must(template.ParseFiles("index.html"))
	tmpl.Execute(w, data)
}

func generateLineData() (temps, timestamps []opts.LineData) {
	temps = make([]opts.LineData, 0)
	timestamps = make([]opts.LineData, 0)
	readings := database.QueryReadings()

	for _, r := range readings {
		temps = append(temps, opts.LineData{Value: r.Temperature})
		timestamps = append(timestamps, opts.LineData{Value: r.Reading_timestamp})
	}
	return temps, timestamps
}
