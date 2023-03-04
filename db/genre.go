package db

import (
	"strings"

	"github.com/gocolly/colly"
)

const (
	url_rus = "http://www.fictionbook.org/index.php/%D0%96%D0%B0%D0%BD%D1%80%D1%8B_FictionBook_2.1"
	url_eng = "http://www.fictionbook.org/index.php/Eng:FictionBook_2.1_genres"
)

type t_genre struct {
	section    string
	subsection string
}

func getGenres(url string) map[string]t_genre {
	var (
		sect    string
		subsect string
		genre   string
	)
	res := make(map[string]t_genre)
	c := colly.NewCollector()
	c.OnHTML("div#mw-content-text", func(d *colly.HTMLElement) {
		d.ForEach("ul", func(_ int, s_item *colly.HTMLElement) {
			s_item.ForEach("li", func(_ int, s_cont *colly.HTMLElement) {
				s_cont.ForEach("i", func(_ int, s_title *colly.HTMLElement) {
					sect = strings.Trim(s_title.Text, "\n\r ")
				})
				s_cont.ForEach("ul", func(_ int, ss_list *colly.HTMLElement) {
					ss_list.ForEach("li", func(_ int, ss_item *colly.HTMLElement) {
						ss_item.ForEach("b", func(_ int, ss_title *colly.HTMLElement) {
							genre = strings.Trim(ss_title.Text, "\n\r ")
						})
						subsect = strings.Replace(ss_item.Text, genre, "", 1)
						subsect = strings.Trim(subsect, "- \n\r")
						res[genre] = t_genre{
							section:    sect,
							subsection: subsect,
						}
					})
				})
			})
		})
	})
	c.Visit(url)
	return res
}

func GetGenres() []Genre {
	var res []Genre
	res_rus := getGenres(url_rus)
	res_eng := getGenres(url_eng)
	for genre, entry := range res_rus {
		res = append(res, Genre{
			Genre:      genre,
			Section:    entry.section,
			Subsection: entry.subsection,
		})
	}
	for genre, entry := range res_eng {
		for i, _ := range res {
			if res[i].Genre == genre {
				res[i].SectionEng = entry.section
				res[i].SubsectionEng = entry.subsection
			}
		}
	}
	return res
}
