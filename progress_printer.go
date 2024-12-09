package main

import (
	"bufio"
	"context"
	"fmt"
	"strings"
	"time"

	"golang.org/x/exp/rand"
)

const (
	total_lessons  = 1873
	archmage_level = 100
	rare_chest     = 5000
)

func (apiCfg *apiConfig) status_report(u *User, key time.Time, init time.Duration, increment time.Duration) {
	apiCfg.calc_Progress_Rate(u)
	c := u.S[key]
	if c.Level > archmage_level {
		apiCfg.congrats_archmage(u, key)
		return
	}
	line := fmt.Sprintf("Generating %s's status report time @ %s\n\n", u.Discord_Name, key.Format("2006-01-02 15:04:05"))
	word_Printer(line, init, increment)
	rank := rank_calculator(c.Level)
	if u.Nickname == "" {
		u.Nickname = "Insert Name in DB"
	}
	summary := fmt.Sprintf("%s is rank %s Level %d & %d Xp\n", u.Nickname, rank, c.Level, c.Xp)
	word_Printer(summary, init, increment)

	Print_Art(u.Art, init, increment)

	m := calc_xp_metrics(c.Level, c.Xp, u.P_Rate)

	next_milestone := fmt.Sprintf("Needs %d Additional Xp for lvl up & %d Xp to reach %s at lvl %d\n", m.lvl_up_xp, m.next_rank_xp, m.next_rank_name, m.next_rank_lvl)
	word_Printer(next_milestone, init, increment)
	absolute_progress := fmt.Sprintf("%s is %.2f%% of the way to Archmage and has solved %d / %d lessons\n\n", u.Discord_Name, m.arch_perc, c.Lessons, total_lessons)
	word_Printer(absolute_progress, init, increment)
	var Eternity []string
	heat_death := fmt.Sprintf("The slow, but ever present expansion of the universe will slowly pull everything apart to a state where entropy is impossible in 1.7 x 10^106 years, still before %s reaches archmage though\n", u.Discord_Name)
	Eternity = append(Eternity, heat_death)
	fossil_fuels := fmt.Sprintf("If everyone were to use the same amount of oil %s uses to code, it would still take about 400 million years to replenish\n", u.Discord_Name)
	Eternity = append(Eternity, fossil_fuels)
	earth_death := fmt.Sprintf("The sun will expand and consume the earth in 7.6 billion years, %s won't be an archmage when that happens\n", u.Discord_Name)
	Eternity = append(Eternity, earth_death)
	hawaii_sinks := fmt.Sprintf("Hope %s isn't planning to use their dev salary to live in Hawaii, even in 80 mil years when the last bit of the big island sinks below the ocean they won't be archmage\n", u.Discord_Name)
	Eternity = append(Eternity, hawaii_sinks)
	if u.P_Rate > 0 {
		progressing_line := fmt.Sprintf("@ %s's current rate of progress %d Xp per day, they will be an Archmage in %d days\n", u.Discord_Name, u.P_Rate, m.days_until)
		word_Printer(progressing_line, init, increment)
	} else {

		word_Printer(Eternity[rand.Intn(len(Eternity))], init, increment)
	}

	rare_chest_line := fmt.Sprintf("Archmage in %d days if %s gets a rare (low effort) chest everyday\n", m.rare_days, u.Discord_Name)
	if u.P_Rate > 80000 {
		king_message := fmt.Sprintf("You've earned your frog going Super-Saiyan chest %s, fetching dope!\n", u.Nickname)
		word_Printer(king_message, init, increment)
	}
	if u.P_Rate > 20000 {
		crab_message := fmt.Sprintf("You've recieved the crabman's, you've gone beast mode %s\n", u.Nickname)
		word_Printer(crab_message, init, increment)
	} else if u.P_Rate > 10000 {
		line2 := fmt.Sprintf("You're putting in some effort %s, someone's getting their frog king chest each day!\n", u.Nickname)
		word_Printer(line2, init, increment)
	} else if u.P_Rate > 5000 {
		progress_message_rare := fmt.Sprintf("You're far from a beast %s, but at least you're getting your low-effort chest a day\n", u.Nickname)
		word_Printer(progress_message_rare, init, increment)
	} else if u.P_Rate > 3000 {
		progress_message_mediocre := fmt.Sprintf("Seriously, can''t even get your daily low-effort (rare) chest each day!%s\n", u.Nickname)
		word_Printer(progress_message_mediocre, init, increment)
		word_Printer(rare_chest_line, init, increment)
	} else {
		if len(u.Insults) > 0 {
			insult_message := fmt.Sprintf("%s %s\n", u.Nickname, u.Insults[rand.Intn(len(u.Insults))])
			word_Printer(insult_message, init, increment)
			word_Printer(rare_chest_line, init, increment)
		} else {
			line2 := fmt.Sprintf("You're a disappointment %s you're acting like a prompt engineer\n", u.Nickname)
			word_Printer(line2, init, increment)
			word_Printer(rare_chest_line, init, increment)
		}
	}
	fmt.Println()
	file_writer(u, key)
}

func Print_Art(art string, init_wait, wait_increment time.Duration) {
	scanner := bufio.NewScanner(strings.NewReader(art))
	for scanner.Scan() {
		time.Sleep(init_wait)
		fmt.Println(scanner.Text())
		init_wait += wait_increment
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
	fmt.Println()
}

func word_Printer(message string, init_wait, wait_increment time.Duration) {
	words := strings.Fields(message)
	for _, word := range words {
		time.Sleep(init_wait)
		fmt.Print(word + " ")
		init_wait += wait_increment
	}
	fmt.Println()
}
func (apiCfg *apiConfig) congrats_archmage(u *User, key time.Time) {
	ctx := context.Background()
	apiCfg.DB.MakeWizard(ctx, u.UUID)
	u.Wizard = true
	fmt.Println("Congratulations to", u.Discord_Name, "for reaching Archmage status")
	Print_Art(arch_wizard, 70*time.Millisecond, 2*time.Millisecond)
	file_writer(u, key)
}
