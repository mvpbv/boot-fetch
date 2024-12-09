package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"strconv"
	"strings"
	"time"
)

func (apiCfg *apiConfig) Track_User(u *User) {
	url := "https://boot.dev/"

	c := colly.NewCollector(
		colly.AllowedDomains("www.boot.dev", "boot.dev"),
	)
	c.SetRequestTimeout(120 * time.Second)

	xpQuery := "span.ml-2.text-xs"
	levelQuery := "span.text-2xl > span.font-bold.text-white"

	lessonsQuery := "span.text-2xl.font-bold.text-white"

	key := time.Now().UTC()

	to_add := new(Status)
	to_add.Time = key

	c.OnHTML(levelQuery, func(e *colly.HTMLElement) {
		//fmt.Printf("attribute: %s \n username: %s \n", e.Attr("class"), e.Text)
		var err error
		to_add.Level, err = strconv.Atoi(e.Text)
		if err != nil {
			fmt.Println(err)
		}
	})

	c.OnHTML(xpQuery, func(e *colly.HTMLElement) {
		//fmt.Printf("attribute: %s \n username: %s \n", e.Attr("class"), e.Text)
		var err error
		xp_str := e.Text
		vals := strings.Fields(xp_str)
		to_add.Xp, err = strconv.Atoi(vals[0])
		if err != nil {
			fmt.Println("Couldn't convert xp to int")
		}
	})
	c.OnHTML(lessonsQuery, func(e *colly.HTMLElement) {
		if e.Index == 0 {
			var err error
			to_add.Lessons, err = strconv.Atoi(e.Text)
			if err != nil {
				fmt.Println(err)
			}
		}
	})

	c.OnRequest(func(r *colly.Request) {
		//fmt.Println("Visiting", r.URL)
	})

	c.OnResponse(func(r *colly.Response) {
		//fmt.Println("Visited", r.Request.URL)
	})

	c.OnError(func(r *colly.Response, err error) {
		//fmt.Println("Error", err)
	})

	c.Visit(url + "u/" + u.Boot_Name)
	u.S[key] = *to_add
	u.Recent_Key = key
}
