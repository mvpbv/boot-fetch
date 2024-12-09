package main

import (
	"fmt"
	"os"
	"time"
)

func file_writer(u *User, key time.Time) {

	path := "logs/" + u.Discord_Name + ".csv"
	l := u.S[key]
	if l.Level == 0 && l.Xp == 0 && l.Lessons == 0 {
		fmt.Println("No data to write at:", key.Format("2006-01-02 15:04:05"))
		return
	}
	line := fmt.Sprintf("%s,%d,%d,%d\n", key.Format("2006-01-02 15:04:05"),
		l.Level,
		l.Xp,
		l.Lessons,
	)
	if !file_exists(path) {
		create_file(line, path)
	} else {
		write_line(line, path)
	}
}
func file_exists(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func write_line(line, path string) {
	f, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	_, err = f.WriteString(line)
	if err != nil {
		fmt.Println(err)
	}
}

func create_file(line, path string) {
	file, err := os.Create(path)
	if err != nil {
		fmt.Println(err)
	}
	file.Close()
	f, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	_, err = f.WriteString("Time,Level,Xp,Lessons\n")
	if err != nil {
		fmt.Println(err)
	}
	_, err = f.WriteString(line)
	if err != nil {
		fmt.Println(err)
	}
}
