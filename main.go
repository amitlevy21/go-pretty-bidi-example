package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/jedib0t/go-pretty/v6/table"
)

type Tag = string

type Expense struct {
	Date   time.Time
	Amount float64
	Class  string
	Tags   []Tag
}

type Expenses struct {
	Classified   []*Expense
	Unclassified []*Expense
}

func UTCDate(year int, month time.Month, day int) time.Time {
	timeZone, _ := time.LoadLocation("UTC")
	return time.Date(year, month, day, 0, 0, 0, 0, timeZone)
}

func NewTestExpenses() *Expenses {
	return &Expenses{Classified: []*Expense{
		{
			Date:   UTCDate(2021, 03, 18),
			Amount: 5.0,
			Class:  "הי         י מחלקה1",
			Tags:   []Tag{"תג1", "תג2"},
		},
		{
			Date:   UTCDate(2021, 04, 19),
			Amount: 5.0,
			Class:  "מחלקה1",
			Tags:   []Tag{"תג1"},
		},
		{
			Date:   UTCDate(2021, 05, 20),
			Amount: 5.0,
			Class:  "מחלקה2",
			Tags:   []Tag{"תג1"},
		},
	}}
}

func makeMainTable(expenses []*Expense) string {
	t := table.NewWriter()
	// if we don't wrap each header with [] they will show in reverse order
	// t.AppendHeader(table.Row{"#", "תגים", "מחלקה", "סכום", "תאריך"})
	t.AppendHeader(table.Row{"#", "[תאריך]", "[סכום]", "[מחלקה]", "[תגים]"})
	appendTableBody(expenses, t)
	t.AppendSeparator()
	t.AppendFooter(table.Row{"", "סהכ", 30})
	return t.Render() + "\n"
}

func appendTableBody(expenses []*Expense, t table.Writer) {
	for i, e := range expenses {
		dateWithoutTime := strings.Split(e.Date.String(), " ")[0]
		t.AppendRows([]table.Row{
			{
				i,
				dateWithoutTime,
				e.Amount,
				fmt.Sprintf("[%s]", e.Class), // again, we format here to wrap with [] otherwise order will break
				e.Tags,
			},
		})
	}
}

func main() {
	fmt.Println(makeMainTable(NewTestExpenses().Classified))
}
