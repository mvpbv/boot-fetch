package main

import (
	"context"
	"fmt"
	"sort"
	"time"
)

type Stud struct {
	Nickname      string
	Progress_Rate int
	Total_XP      int
	Days_Until    int
	Art           string
	Wizard        bool
	Level         int
}

func (apiCfg *apiConfig) study_hall(users []*User) []Stud {
	ctx := context.Background()
	size := len(users)
	Study_Hall := make([]Stud, size)

	for _, v := range users {
		apiCfg.calc_Progress_Rate(v)
		p, err := apiCfg.DB.GetRecentProgressUser(ctx, v.UUID)
		if err != nil {
			fmt.Println(err)
		}
		s := DbRecStatusToStatus(p)
		m := calc_xp_metrics(s.Level, s.Xp, v.P_Rate)

		Study_Hall = append(Study_Hall, Stud{
			Nickname:      v.Discord_Name,
			Progress_Rate: v.P_Rate,
			Total_XP:      m.total_xp,
			Days_Until:    m.days_until,
			Art:           v.Art,
			Wizard:        v.Wizard,
			Level:         s.Level,
		})

	}
	return Study_Hall
}

func (apiCfg *apiConfig) study_hall_printer(Study_Hall []Stud, users []*User) {
	word_Printer("Study Hall Progress Rate Report", 70*time.Millisecond, 3*time.Millisecond)
	study_hall_progress_report(Study_Hall)
	study_hall_xp_report(Study_Hall)
	word_Printer("Study Hall Days Until Archmage Report", 70*time.Millisecond, 3*time.Millisecond)
	study_hall_days_report(Study_Hall)
	word_Printer("People Better Than Mayro Meat Bicycle", 70*time.Millisecond, 3*time.Millisecond)
	apiCfg.BetterThanTarget(Study_Hall, users)

}
func study_hall_progress_report(Study_Hall []Stud) {
	Study_Hall_Progress := make([]Stud, len(Study_Hall))
	copy(Study_Hall_Progress, Study_Hall)
	sort.Slice(Study_Hall_Progress, func(i, j int) bool {
		return Study_Hall_Progress[i].Progress_Rate > Study_Hall_Progress[j].Progress_Rate
	})
	Study_Hall_Progress = removeNulls(Study_Hall_Progress)
	for k, v := range Study_Hall_Progress {
		if k == 0 {
			champion(v)
		}
		Progress_Message := fmt.Sprintf("%s\nProgress Rate: %d\n\n", v.Nickname, v.Progress_Rate)
		word_Printer(Progress_Message, 70*time.Millisecond, 3*time.Millisecond)
	}
}

func study_hall_xp_report(Study_Hall []Stud) {
	Study_Hall_Xp := make([]Stud, len(Study_Hall))
	copy(Study_Hall_Xp, Study_Hall)
	word_Printer("Study Hall XP Report\n", 70*time.Millisecond, 3*time.Millisecond)
	sort.Slice(Study_Hall_Xp, func(i, j int) bool {
		return Study_Hall_Xp[i].Total_XP > Study_Hall_Xp[j].Total_XP
	})
	Study_Hall_Xp = removeWizards(Study_Hall_Xp)
	Study_Hall_Xp = removeNulls(Study_Hall_Xp)
	for k, v := range Study_Hall_Xp {
		if k == 0 {
			champion(v)
		}
		Xp_Message := fmt.Sprintf("%s\nLevel: %d \tTotal XP: %d\n\n", v.Nickname, v.Level, v.Total_XP)
		word_Printer(Xp_Message, 70*time.Millisecond, 3*time.Millisecond)
	}
}

func study_hall_days_report(Study_Hall []Stud) {
	Study_Hall_Days := make([]Stud, len(Study_Hall))
	copy(Study_Hall_Days, Study_Hall)
	sort.Slice(Study_Hall_Days, func(i, j int) bool {
		return Study_Hall_Days[i].Days_Until < Study_Hall_Days[j].Days_Until
	})
	Study_Hall_Days = removeWizards(Study_Hall_Days)
	Study_Hall_Days = removeNulls(Study_Hall_Days)
	for k, v := range Study_Hall_Days {
		if k == 0 {
			champion(v)
		}
		if v.Wizard {
			continue
		}
		Days_Message := fmt.Sprintf("%s\nDays Until Archmage: %d\n\n", v.Nickname, v.Days_Until)
		word_Printer(Days_Message, 70*time.Millisecond, 3*time.Millisecond)
	}
}

func removeWizards(Study_Hall []Stud) []Stud {
	Study_Hall_2 := make([]Stud, 0, len(Study_Hall))
	for _, v := range Study_Hall {
		if !v.Wizard {
			Study_Hall_2 = append(Study_Hall_2, v)
		}
	}
	return Study_Hall_2
}

func removeNulls(Study_Hall []Stud) []Stud {
	Study_Hall_2 := make([]Stud, 0, len(Study_Hall))
	for _, v := range Study_Hall {
		if len(v.Nickname) > 0 {
			Study_Hall_2 = append(Study_Hall_2, v)
		}
	}
	return Study_Hall_2
}
func champion(v Stud) {
	word_Printer("Your Champion is\n", 80*time.Millisecond, 4*time.Millisecond)
	Print_Art(v.Art, 70*time.Millisecond, 3*time.Millisecond)
}
func (apiCfg *apiConfig) BetterThanTarget(Study_Hall []Stud, users []*User) {
	var targetProg int
	var target *User
	for _, stud := range Study_Hall {
		if stud.Nickname == "Mario" {
			target = StudyHallToUser(&stud, users)
			apiCfg.calc_Progress_Rate(target)
			targetProg = target.P_Rate
			break
		}
	}
	for _, user := range Study_Hall {
		if user.Progress_Rate > targetProg {
			better := fmt.Sprintf("%s is better than %s \n", user.Nickname, target.Nickname)
			word_Printer(better, 70*time.Millisecond, 3*time.Millisecond)
			better_by := fmt.Sprintf("%s is beating %s by %dXp Per Day", user.Nickname, target.Nickname, user.Progress_Rate-targetProg)
			word_Printer(better_by, 70*time.Millisecond, 3*time.Millisecond)
			percent_better := fmt.Sprintf("Which totals out to %f%% better than %s", float64(user.Progress_Rate)/float64(targetProg)*100, target.Nickname)
			word_Printer(percent_better, 70*time.Millisecond, 3*time.Millisecond)
			Print_Art(user.Art, 70*time.Millisecond, 3*time.Millisecond)
		}
	}
}
