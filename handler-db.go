package main

import (
	"encoding/csv"
	"fmt"
	"github.com/google/uuid"
	"github.com/mvpbv/boot-fetch/internal/database"
	"net/http"
	"os"
	"strconv"
	"time"
)

type csvFields struct {
	Time    time.Time
	Level   int32
	Xp      int32
	Lessons int32
	TotalXp int32
	User_id uuid.UUID
}

func (apiCfg *apiConfig) dbFixxer(w http.ResponseWriter, r *http.Request) {
	file_exists("db_fix.csv")
	f, err := os.Open("db_fix.csv")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	reader := csv.NewReader(f)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println(err)
		return
	}
	records = records[1:]
	for _, record := range records {
		c := clean_csv(record)
		apiCfg.DB.CreateProgress(r.Context(), database.CreateProgressParams{
			ID:      uuid.New(),
			Time:    c.Time,
			Level:   c.Level,
			Xp:      c.Xp,
			Lessons: c.Lessons,
			TotalXp: c.TotalXp,
			UserID:  c.User_id,
		})
	}
	respondWithJSON(w, http.StatusOK, "Mischief Managed")
	os.Remove("db_fix.csv")
}
func clean_csv(r []string) *csvFields {
	var fields csvFields

	t, err := time.Parse("2006-01-02 15:04:05", r[0])
	if err != nil {
		fmt.Println(err)
		return nil
	}
	fields.Time = t
	l, err := strconv.Atoi(r[1])
	level := int32(l)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	fields.Level = level
	x, err := strconv.Atoi(r[2])
	xp := int32(x)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	fields.Xp = xp
	les, err := strconv.Atoi(r[3])
	lessons := int32(les)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	fields.Lessons = lessons

	totXp, err := strconv.Atoi(r[4])
	totalXp := int32(totXp)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	fields.TotalXp = totalXp
	c, err := uuid.Parse(r[5])
	if err != nil {
		fmt.Println(err)
		return nil
	}
	fields.User_id = c
	return &fields

}
