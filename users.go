package main

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/mvpbv/boot-fetch/internal/database"
	"math/rand"
	"time"
)

type User struct {
	Discord_Name string
	Nickname     string
	Art          string
	Boot_Name    string
	F_Record     Status
	S            map[time.Time]Status
	P_Rate       int
	Recent_Key   time.Time
	Insults      []string
	UUID         uuid.UUID
	Wizard       bool
}

func (apiCfg *apiConfig) DbUsertoUser(u *database.User) *User {
	defaults := []string{
		rat, worm, caterpillar,
	}
	first_status, err := apiCfg.DB.GetFirstProgressUser(context.Background(), u.ID)
	if err != nil {
		fmt.Println(err)
	}
	appStatus := DbInitStatusToStatus(first_status)

	insults, err := apiCfg.DB.GetUserInsults(context.Background(), u.ID)
	if err != nil {
		fmt.Println(err)
	}
	var app_insults []string
	for _, insult := range insults {
		app_insults = append(app_insults, insult.Insult)

	}
	var art string

	switch u.DiscordName {
	//Add art here
	default:
		art = defaults[rand.Intn(len(defaults))]
	}
	return &User{
		Discord_Name: u.DiscordName,
		Nickname:     u.Nickname,
		Boot_Name:    u.BootName,
		UUID:         u.ID,
		Wizard:       u.Wizard,
		F_Record:     appStatus,
		Insults:      app_insults,
		S:            make(map[time.Time]Status),
		Art:          art,
	}
}

type Status struct {
	Xp       int
	Progress int
	Level    int
	Time     time.Time
	Lessons  int
	Total_XP int
}

func DbRecStatusToStatus(dbStatus database.GetRecentProgressUserRow) Status {
	return Status{
		Level:    int(dbStatus.Level),
		Xp:       int(dbStatus.Xp),
		Total_XP: int(dbStatus.TotalXp),
		Time:     dbStatus.Time,
		Lessons:  int(dbStatus.Lessons),
	}
}
func DbInitStatusToStatus(dbStaus database.GetFirstProgressUserRow) Status {
	return Status{
		Level:    int(dbStaus.Level),
		Xp:       int(dbStaus.Xp),
		Total_XP: int(dbStaus.TotalXp),
		Time:     dbStaus.Time,
		Lessons:  int(dbStaus.Lessons),
	}
}

func (apiCfg *apiConfig) LoadDbUsers() []*User {
	DbUsers, err := apiCfg.DB.GetUsers(context.Background())
	if err != nil {
		fmt.Println(err)
	}
	var users []*User
	for _, dbUser := range DbUsers {
		user := apiCfg.DbUsertoUser(&dbUser)
		users = append(users, user)
	}
	return users
}
