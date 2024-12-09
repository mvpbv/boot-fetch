package main

import (
	"fmt"
	"time"
)

func (apiCfg *apiConfig) crawler(users []*User) {
	init := 60 * time.Millisecond
	increment := 2 * time.Millisecond

	fmt.Println("Crawler Initializing")
	for {
		for i := 0; i < 3; i++ {
			for _, v := range users {
				if v.Wizard {
					continue
				}
				apiCfg.Track_User(v)
				apiCfg.status_report(v, v.Recent_Key, init, increment)
				apiCfg.dbWriter(v)
			}
		}
		studs := apiCfg.study_hall(users)
		apiCfg.study_hall_printer(studs, users)
	}
}
