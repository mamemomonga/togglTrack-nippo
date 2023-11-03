package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/mamemomonga/togglTrack-nippo/type01/cfg"
	"github.com/mamemomonga/togglTrack-nippo/type01/slackutil"
	"github.com/mamemomonga/togglTrack-nippo/type01/togglutil"
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

	for _, tsk := range config.Tasks {
		ta := taskFilter(tsk, *DatesOffset)
		if len(ta) > 0 {
			buf := ""
			buf = buf + "```\n"
			buf = buf + "【日報】\n"
			buf = buf + "  " + ta[0].time + "\n"
			for _, en := range ta {
				buf = buf + fmt.Sprintf("   %s [%s] %s\n", en.duration, en.project, en.desc)
			}
			buf = buf + "```\n"
			fmt.Println(buf)
			if *slackPost {
				sl := slackutil.New(tsk.Slack.Token)
				sl.PostSimple(buf, tsk.Slack.Channel)
			}
		}

	}
}

type entryT struct {
	time     string
	desc     string
	duration string
	project  string
}

func taskFilter(tsk cfg.CfgTask, offset int) (entrs []entryT) {

	var err error
	tu := togglutil.New(config.TogglWorkspace.Token)
	err = tu.GetAccount()
	if err != nil {
		log.Fatal(err)
	}
	err = tu.FilterWorkspace(config.TogglWorkspace.Workspace)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("debug: [WorkspaceID] %d", tu.WorkspaceID)

	for _, tgl := range tsk.Toggls {
		err = tu.FilterClient(tgl.Client)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("debug: [ClientID] %d", tu.ClientID)

		err = tu.FilterProject(tgl.Project)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("debug: [ProjectID] %d", tu.ProjectID)

		err = tu.FilterTimeEntries()
		if err != nil {
			log.Fatal(err)
		}

		now := time.Now()
		targetTime := now.AddDate(0, 0, offset*-1).Local()
		year, month, day := targetTime.Date()
		log.Printf("debug: Filter: %d/%d/%d", year, month, day)
		err = tu.FilterTimeEntriesStartDate(year, month, day)
		if err != nil {
			if err == togglutil.ErrTimeEntriesNotFound {
				return
			}
			log.Fatal(err)
		}

		for _, t := range tu.TimeEntries {
			desc := ""
			if t.Description != "" {
				desc = t.Description
			}

			week := targetTime.Weekday()
			weekJp := []string{"日", "月", "火", "水", "木", "金", "土"}
			entrs = append(entrs, entryT{
				time:     fmt.Sprintf("%04d年%02d月%02d日(%s)", year, month, day, weekJp[week]),
				desc:     desc,
				duration: fmt.Sprintf("%3.1f 時間", float64(t.Duration)/3600),
				project:  tgl.Project,
			})
		}
	}
	return
}
