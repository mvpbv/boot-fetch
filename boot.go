package main

import (
	"fmt"
	"time"
)

type Boot struct {
	Art  string
	Init string
}

func (apiCfg *apiConfig) Boot_Fetch(b *Boot) {
	compress_files()

	init := 20 * time.Millisecond
	increment := 1 * time.Millisecond
	fmt.Println("Boot.fetch initializing")
	Print_Art(b.Art, init, increment)
	word_Printer(b.Init, init, increment)
	allUsers := apiCfg.LoadDbUsers()

	for _, v := range allUsers {
		if v.Wizard {
			apiCfg.calc_Wizard_Rate(v)
			continue
		}
		apiCfg.Track_User(v)
		apiCfg.status_report(v, v.Recent_Key, init, increment)
		apiCfg.dbWriter(v)

	}
	Study_Hall := apiCfg.study_hall(allUsers)
	apiCfg.study_hall_printer(Study_Hall, allUsers)
	go apiCfg.crawler(allUsers)
}

/*
func (apiCfg *apiConfig) seedDb(v *User) {
	f_rec := v.F_Record
	is_wizard := v.S[v.Recent_Key].Level >= 100
	user, err := apiCfg.DB.CreateUser(context.Background(), database.CreateUserParams{
		ID:          uuid.New(),
		BootName:    v.Boot_Name,
		DiscordName: v.Discord_Name,
		CreatedAt:   f_rec.Time,
		UpdatedAt:   time.Now().UTC(),
		Wizard:      is_wizard,
	})
	apiCfg.DB.CreateNickname(context.Background(), database.CreateNicknameParams{
		ID:       uuid.New(),
		UserID:   user.ID,
		Nickname: v.Nickname,
	})
	if len(v.Insults) > 1 {
		for _, insult := range v.Insults {
			apiCfg.DB.CreateInsult(context.Background(), database.CreateInsultParams{
				ID:     uuid.New(),
				UserID: user.ID,
				Insult: insult,
			})
		}
	}
	if err != nil {
		fmt.Println(err)
	}
	v.UUID = user.ID
}
*/
