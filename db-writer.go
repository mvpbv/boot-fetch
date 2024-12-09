package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/mvpbv/boot-fetch/internal/database"
)

func (apiCfg *apiConfig) dbWriter(u *User) {

	ctx := context.Background()
	r := u.S[u.Recent_Key]
	good := checkRecordValidity(&r)
	if !good {
		return
	}
	Dblessons, err := apiCfg.DB.GetRecentProgressUser(ctx, u.UUID)
	lessons := Dblessons.Lessons
	if err != nil {
		fmt.Println(err)
	}
	if int(lessons) == r.Lessons {
		return
	}
	tot_xp := xp_at_level(r.Level) + r.Xp
	_, err = apiCfg.DB.CreateProgress(ctx, database.CreateProgressParams{
		ID:      uuid.New(),
		UserID:  u.UUID,
		Level:   int32(r.Level),
		Xp:      int32(r.Xp),
		Lessons: int32(r.Lessons),
		TotalXp: int32(tot_xp),
		Time:    time.Now().UTC(),
	})
	if err != nil {
		fmt.Println(err)
	}
	word_Printer("Wow someone actually did something", 50*time.Millisecond, 2*time.Millisecond)
}

func checkRecordValidity(s *Status) bool {
	if s.Time.IsZero() {
		return false
	}
	if s.Level == 0 {
		return false
	}
	if s.Xp == 0 {
		return false
	}
	return true
}
