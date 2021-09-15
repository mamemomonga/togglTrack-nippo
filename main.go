package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/mamemomonga/togglTrack-nippo/cfg"
	"github.com/mamemomonga/togglTrack-nippo/slackutil"
	"github.com/mamemomonga/togglTrack-nippo/togglutil"
	//	"github.com/davecgh/go-spew/spew"
)

var (
	config *cfg.Cfg
)

func main() {
	var (
		configFile  = flag.String("config", "./config.yaml", "ConfigFile")
		slackPost   = flag.Bool("slack", false, "Slack Post")
		DatesOffset = flag.Int("days", 0, "Dates Offset")
	)
	flag.Parse()

	var err error
	config, err = cfg.New(*configFile)
	if err != nil {
		log.Fatal(err)
	}

	str := toggl(*DatesOffset)
	fmt.Println(str)

	if *slackPost {
		sl := slackutil.New(config.Slack.Token)
		sl.PostSimple(str, config.Slack.Channel)
	}
}

func toggl(datesOffset int) string {
	var err error
	tu := togglutil.New(config.Toggl.Token)
	err = tu.GetAccount()
	if err != nil {
		log.Fatal(err)
	}

	err = tu.FilterWorkspace(config.Toggl.Workspace)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("debug: [WorkspaceID] %d", tu.WorkspaceID)

	err = tu.FilterClient(config.Toggl.Client)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("debug: [ClientID] %d", tu.ClientID)

	err = tu.FilterProject(config.Toggl.Project)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("debug: [ProjectID] %d", tu.ProjectID)

	err = tu.FilterTimeEntries()
	if err != nil {
		log.Fatal(err)
	}

	now := time.Now()
	targetTime := now.AddDate(0, 0, datesOffset*-1).Local()
	year, month, day := targetTime.Date()
	week := targetTime.Weekday()

	err = tu.FilterTimeEntriesStartDate(year, month, day)
	if err != nil {
		log.Fatal(err)
	}

	weekJp := []string{"日", "月", "火", "水", "木", "金", "土"}

	buf := ""
	buf = buf + "```\n"
	buf = buf + "【日報】\n"
	buf = buf + fmt.Sprintf("  %04d年%02d月%02d日(%s)\n", year, month, day, weekJp[week])
	for i := 0; i < len(tu.TimeEntries); i++ {
		tme := tu.TimeEntries[i]

		desc := "各種作業"
		if tme.Description != "" {
			desc = tme.Description
		}
		buf = buf + fmt.Sprintf("    %3.1f 時間 %s\n", float64(tme.Duration)/3600, desc)
	}
	buf = buf + "```\n"

	return buf
}
