package main

import (
	"errors"
	"time"

	toggl "github.com/jason0x43/go-toggl"
)

var (
	ErrWorkspaceIdNotFound = errors.New("error: workspace id not found")
	ErrClientIdNotFound    = errors.New("error: client id not found")
	ErrProjectIdNotFound   = errors.New("error: project id not found")
	ErrTimeEntriesNotFound = errors.New("error: time entries not found")
)

type Toggler struct {
	Session     toggl.Session
	Account     toggl.Account
	TimeEntries []toggl.TimeEntry
	WorkspaceID int
	ClientID    int
	ProjectID   int
}

func NewToggl(token string) *Toggler {
	t := &Toggler{}
	t.Session = toggl.OpenSession(token)
	return t
}

func (t *Toggler) LoadAccount() (err error) {
	t.Account, err = t.Session.GetAccount()
	if err != nil {
		return err
	}
	return nil
}

func (t *Toggler) UseWorkspace(name string) (err error) {
	id := 0
	for i := 0; i < len(t.Account.Data.Workspaces); i++ {
		if t.Account.Data.Workspaces[i].Name == name {
			id = t.Account.Data.Workspaces[i].ID
		}
	}
	if id == 0 {
		return ErrWorkspaceIdNotFound
	}
	t.WorkspaceID = id
	return nil
}

func (t *Toggler) GetProjectID(name string) (int, error) {
	id := 0
	for i := 0; i < len(t.Account.Data.Projects); i++ {
		if t.Account.Data.Projects[i].Cid == t.ClientID && t.Account.Data.Projects[i].Name == name {
			id = t.Account.Data.Projects[i].ID
		}
	}
	if id == 0 {
		return 0, ErrProjectIdNotFound
	}
	return id, nil
}

func (t *Toggler) LoadTimeEntries() (err error) {
	t.TimeEntries = []toggl.TimeEntry{}
	for i := 0; i < len(t.Account.Data.TimeEntries); i++ {
		//		if t.Account.Data.TimeEntries[i].Pid == t.ProjectID {
		t.TimeEntries = append(t.TimeEntries, t.Account.Data.TimeEntries[i])
		//		}
	}
	if len(t.TimeEntries) == 0 {
		return ErrTimeEntriesNotFound
	}
	return nil
}

func (t *Toggler) FilterTimeEntriesStartDate(year int, month time.Month, day int) (err error) {
	te := []toggl.TimeEntry{}
	for i := 0; i < len(t.TimeEntries); i++ {
		tes := t.TimeEntries[i].Start
		y, m, d := tes.Local().Date()
		if y == year && m == month && d == day {
			te = append(te, t.TimeEntries[i])
		}
	}
	if len(te) == 0 {
		return ErrTimeEntriesNotFound
	}
	t.TimeEntries = te
	return nil
}
