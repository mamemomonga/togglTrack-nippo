package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/mamemomonga/togglTrack-nippo/type02/cfg"
)

func main() {
	var err error

	var (
		configFile = flag.String("config", "./config.yaml", "configファイル")
		offsetDays = flag.Int("days", 0, "今日からn日分戻る")
	)
	flag.Parse()

	config, err := cfg.New(*configFile)
	if err != nil {
		log.Fatal(err)
	}

	now := time.Now()
	targetTime := now.AddDate(0, 0, *offsetDays*-1).Local()

	{
		y, m, d := targetTime.Date()
		log.Printf("debug: Target: %04d/%02d/%02d", y, m, d)
	}

	tu := NewToggl(config.TogglWorkspace.Token)
	err = tu.LoadAccount()
	if err != nil {
		log.Fatal(err)
	}

	err = tu.UseWorkspace(config.TogglWorkspace.WorkspaceName)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("debug: [WorkspaceID] %d", tu.WorkspaceID)

	tu.LoadTimeEntries()

	err = tu.FilterTimeEntriesStartDate(targetTime.Date())
	if err != nil {
		log.Fatal(err)
	}

	pjr, err := tu.GetProjectID(config.RestProjectName)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("debug: Project %s ID: %d", config.RestProjectName, pjr)

	var timeStart time.Time
	var timeStop time.Time
	var restDuration time.Duration = 0

	for _, t := range tu.TimeEntries {

		duration, err := time.ParseDuration(fmt.Sprintf("%ds", t.Duration))
		if err != nil {
			log.Fatal(err)
		}

		// 休憩の場合
		if t.Pid == pjr {
			restDuration = restDuration + duration
			continue
		}

		// 休憩以外の場合
		if timeStart.IsZero() {
			timeStart = *t.Start
		} else if t.Start.Before(timeStart) {
			timeStart = *t.Start
		}
		if timeStop.IsZero() {
			timeStop = *t.Stop
		} else if t.Stop.After(timeStop) {
			timeStop = *t.Stop
		}
	}

	timeStart = timeStart.Local()
	timeStop = timeStop.Local()
	week := timeStart.Weekday()
	weekJp := []string{"日", "月", "火", "水", "木", "金", "土"}

	fmt.Println("-----------------------------------------------")
	fmt.Printf("[  作業日  ] %04d/%02d/%02d(%s)\n", timeStart.Year(), timeStart.Month(), timeStart.Day(), weekJp[week])
	fmt.Printf("[ 開始時刻 ] %02d:%02d:%02d\n", timeStart.Hour(), timeStart.Minute(), timeStart.Second())
	fmt.Printf("[ 終了時刻 ] %02d:%02d:%02d\n", timeStop.Hour(), timeStop.Minute(), timeStop.Second())
	fmt.Printf("[ 休憩時間 ] %s\n", time.Unix(0, 0).UTC().Add(time.Duration(restDuration)).Format(time.TimeOnly))
	fmt.Println("-----------------------------------------------")
}
