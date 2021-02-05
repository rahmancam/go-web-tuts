package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"text/template"
	"time"
)

var tpl *template.Template

var fm = template.FuncMap{
	"upper":    strings.ToUpper,
	"currency": currency,
	"fdateMDY": monthDayYear,
	"dbl":      double,
	"sq":       square,
}

func double(n int) int {
	return n * 2
}

func square(n int) int {
	return n * n
}

func monthDayYear(t time.Time) string {
	return t.Format("01-02-2006")
}

func currency(n int64) string {
	return "$" + fmt.Sprint(n)
}

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseGlob("tpls/*.gohtml"))
}

// Data - shopping data
type Data struct {
	Shopping map[string]int64
	Date     time.Time
	Num      int
}

func main() {
	shopping := map[string]int64{
		"Onion(1 kg)":   20,
		"Beans (100 g)": 5,
		"Apples (1 kg)": 10,
	}
	data := Data{
		Shopping: shopping,
		Date:     time.Now(),
		Num:      4,
	}
	if err := tpl.ExecuteTemplate(os.Stdout, "shopping.gohtml", data); err != nil {
		log.Fatalln(err)
	}
}
