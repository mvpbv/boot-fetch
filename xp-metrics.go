package main

import (
	"context"
	"fmt"
	"time"
)

type print_metrics struct {
	total_xp       int
	lvl_up_xp      int
	rare_days      int
	archmage_xp    int
	days_until     int
	next_rank_xp   int
	next_rank_name string
	next_rank_lvl  int
	arch_perc      float64
}

func calc_xp_metrics(lvl, xp, p_rate int) print_metrics {
	met := new(print_metrics)
	xp_req := xp_for_level(lvl + 1)
	met.total_xp = xp_at_level(lvl) + xp
	met.archmage_xp = xp_at_level(archmage_level)
	met.lvl_up_xp = xp_req - xp
	met.next_rank_lvl = lvl + (10 - lvl%10)
	met.next_rank_name = rank_calculator(met.next_rank_lvl)
	met.next_rank_xp = xp_at_level(met.next_rank_lvl) - met.total_xp
	met.arch_perc = float64(met.total_xp) / float64(met.archmage_xp) * 100
	xp_needed := met.archmage_xp - met.total_xp
	met.rare_days = xp_needed / rare_chest
	if p_rate > 0 {
		met.days_until = xp_needed / p_rate
	} else {
		met.days_until = 999999999999999999
	}
	return *met
}

func rank_calculator(lvl int) string {
	switch {
	case lvl >= 100:
		return "ArchWizard"
	case lvl >= 90:
		return "Necromancer"
	case lvl >= 80:
		return "Archsage"
	case lvl >= 70:
		return "Sage"
	case lvl >= 60:
		return "Sorcerer"
	case lvl >= 50:
		return "Scholar"
	case lvl >= 40:
		return "Disciple"
	case lvl >= 30:
		return "Acolyte"
	case lvl >= 20:
		return "Pupil"
	default:
		return "Loser"
	}
}

func (apiCfg *apiConfig) calc_Progress_Rate(u *User) {
	first, err := apiCfg.DB.GetFirstProgressUser(context.Background(), u.UUID)
	if err != nil {
		fmt.Println("The database is trash", err)
	}
	base_xp := int(first.TotalXp)
	base_time := first.Time

	if u.Wizard {
		apiCfg.calc_Wizard_Rate(u)
		return
	}
	current, err := apiCfg.DB.GetRecentProgressUser(context.Background(), u.UUID)
	if err != nil {
		fmt.Println("The database is trash", err)
	}
	current_xp := int(current.TotalXp)
	current_time := time.Now().UTC()

	xp_diff := current_xp - base_xp

	if xp_diff != 0 {
		time_diff := current_time.Sub(base_time)
		hours := int(time_diff.Hours())
		if hours != 0 {
			u.P_Rate = xp_diff / hours * 24
		}
	}
}

func (apiCfg *apiConfig) calc_Wizard_Rate(u *User) {
	archmage_xp := xp_at_level(archmage_level)
	dbRecent, err := apiCfg.DB.GetRecentProgressUser(context.Background(), u.UUID)
	if err != nil {
		fmt.Println(err)
	}
	dbFirst, err := apiCfg.DB.GetFirstProgressUser(context.Background(), u.UUID)
	if err != nil {
		fmt.Println(err)
	}
	recent := dbRecent.Time

	first := dbFirst.Time
	init_xp := int(dbFirst.TotalXp)

	xp_diff := archmage_xp - init_xp
	time_diff := recent.Sub(first)
	hours := int(time_diff.Hours())
	if hours != 0 {
		u.P_Rate = xp_diff / hours * 24
	}

}
func xp_for_level(level int) int {
	return 80*level + 400
}

func xp_at_level(level int) int {
	return 40*level*(level-1) + 400*level
}
