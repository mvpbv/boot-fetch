package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func compress_files() {
	f, err := os.Open("logs")
	if err != nil {
		fmt.Println(err)
		return
	}
	files, err := f.Readdir(0)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		compress_file("logs/" + file.Name())
	}
}

func compress_file(path string) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	csvReader := csv.NewReader(file)

	records, err := csvReader.ReadAll()
	if err != nil {
		fmt.Println(err)
		fmt.Println("File path is:", path)
		return
	}

	records = records[1:]
	is_cache := false
	cached := 0
	new_records := make([][]string, 0)
	header_line := []string{"Time", "Level", "Xp", "Lessons", "Total_Xp"}
	new_records = append(new_records, header_line)

	for i, record := range records {
		if !is_cache {
			is_cache = true
			cached = i
			new_record := make_record(record)
			new_records = append(new_records, new_record)
			continue
		}
		if record[3] != records[cached][3] {
			new_records = append(new_records, make_record(record))
			cached = i
		}
	}
	strip_path := strings.Split(path, "/")[1]
	out_path := "websrc/" + strip_path
	out, err := os.Create(out_path)
	if err != nil {
		fmt.Println(err)
		return
	}
	w := csv.NewWriter(out)
	w.WriteAll(new_records)

}
func make_record(record []string) []string {

	lvl, err := strconv.Atoi(record[1])
	if err != nil {
		fmt.Println("Cannot parse lvl:", err)
		fmt.Println("Example record", record[1])
	}
	xp, err := strconv.Atoi(record[2])
	if err != nil {
		fmt.Println("Cannot parse xp", err)
		fmt.Println("Example record", record[2])
	}
	tot := xp_at_level(lvl) + xp
	str_tot := strconv.Itoa(tot)
	new_record := []string{record[0], record[1], record[2], record[3], str_tot}
	return new_record
}
