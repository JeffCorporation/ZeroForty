package main

import (
	"fmt"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"time"
)

var (
	app   *tview.Application
	frame *tview.Frame

	curTaskStart time.Time
	curTask      string
	DayDuration  time.Duration
	WeekDuration time.Duration
)

func updateFrameText() {
	for {
		app.QueueUpdateDraw(func() {
			//Example for running task
			now := time.Now()
			t := now.Sub(curTaskStart)
			d := curTaskStart.Sub(now.Add(DayDuration))
			w := curTaskStart.Sub(now.Add(WeekDuration))

			frame.Clear().
				AddText(fmt.Sprintf("> %s | %s", curTask, fmtDuration(t, false)), true, tview.AlignLeft, tcell.ColorWhite).
				AddText(fmt.Sprintf("Auj. %s | Sem. %s", fmtDuration(d, true), fmtDuration(w, true)), true, tview.AlignRight, tcell.ColorWhite). //header right
				AddText("(N) Nouveau (P) Projets", false, tview.AlignLeft, tcell.ColorBlue)
		})
		time.Sleep(time.Second)
	}
}

func fmtDuration(d time.Duration, round15 bool) string {
	if d < 0 {
		d *= -1
	}

	if round15 {
		d = d.Round(15 * time.Minute)
		h := d / time.Hour
		d -= h * time.Hour
		m := d / time.Minute
		return fmt.Sprintf("%02d:%02d", h, m)
	}

	//Show secs 00:00:00
	d = d.Round(time.Second)
	h := d / time.Hour
	d -= h * time.Hour
	m := d / time.Minute
	d -= m * time.Minute
	s := d / time.Second
	return fmt.Sprintf("%02d:%02d:%02d", h, m, s)
}

func RefreshTable() *tview.Table {

	table := tview.NewTable().
		SetSelectable(true, true)

	rows := 20

	table.SetCell(0, 0, tview.NewTableCell("Lundi 3 sept").
		SetBackgroundColor(tcell.ColorGrey))
	table.SetCell(0, 1, tview.NewTableCell("").
		SetBackgroundColor(tcell.ColorGrey))
	table.SetCell(0, 2, tview.NewTableCell("08:13 [08:15]").
		SetBackgroundColor(tcell.ColorGrey))

	for r := 1; r < rows; r++ {
		table.SetCell(r, 0,
			tview.NewTableCell("2").
				SetTextColor(tcell.ColorWhite).
				SetBackgroundColor(tcell.ColorDarkBlue))

		table.SetCell(r, 1,
			tview.NewTableCell("Bacon ipsum dolor amet ground").
				SetTextColor(tcell.ColorWhite).
				SetBackgroundColor(tcell.ColorDarkBlue).
				SetExpansion(10))

		table.SetCell(r, 2,
			tview.NewTableCell("01:50").
				SetTextColor(tcell.ColorWhite).
				SetBackgroundColor(tcell.ColorDarkBlue).
				SetExpansion(5))
	}

	return table
}

func main() {
	//test data
	curTaskStart = time.Now()
	curTask = "Tâche en cours"
	DayDuration = 12500 * time.Second
	WeekDuration = 129600 * time.Second

	app = tview.NewApplication()

	frame = tview.NewFrame(RefreshTable()).
		SetBorders(0, 0, 0, 0, 0, 0)

	go updateFrameText()

	if err := app.SetRoot(frame, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}
