package main

import (
	"database/sql"
	"os"

	"net/http"

	"html/template"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/types"
	log "github.com/sirupsen/logrus"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)


func fillTemplate(w http.ResponseWriter, s template.HTML) {
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

func generateLineData(db *sql.DB) (temps, timestamps []opts.LineData) {
	temps = make([]opts.LineData, 0)
	timestamps = make([]opts.LineData, 0)
	readings := queryReadings(db)
	for _, r := range readings {
		temps = append(temps, opts.LineData{Value: r.Temperature})
		timestamps = append(timestamps, opts.LineData{Value: r.Reading_timestamp})
	}
	return temps, timestamps
}

func renderChart(db *sql.DB) template.HTML {
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
	temps, timestamps := generateLineData(db)
	line.SetXAxis(timestamps).
		AddSeries("temps", temps).
		SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{Smooth: true}))
	line.Renderer = newSnippetRenderer(line, line.Validate)
	var htmlSnippet template.HTML = renderToHtml(line)
	return htmlSnippet
}

func homeHandler(w http.ResponseWriter, _ *http.Request) {
	db := dbConnect()
	err := createTable(db)
	if err != nil {
		log.Warn(err)
	}

	err = seedDB(db)
	if err != nil {
		log.Warn(err)
	}

	snippet := renderChart(db)
	// resultQueryReadings := queryReadings(db)
	fillTemplate(w, snippet)
}

func init() {
	log.SetOutput(os.Stdout)

	file, err := os.OpenFile("gopetwatch.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err == nil {
		log.SetOutput(file)
	} else {
		log.Info("Failed to log to file, using default stderr")
	}

	log.SetLevel(log.InfoLevel)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", homeHandler)

	log.Info("Starting server...")
	p := "poop"
	log.Info("it smells like: ", p)
	err := http.ListenAndServe(":8081", r)
	if err != nil {
		log.Fatal("Server failed to start", err)
	}
}
