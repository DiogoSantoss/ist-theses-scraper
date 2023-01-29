package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/gocolly/colly"
)

type thesis struct {
	Title    string
	Link     string
	Author   string
	Advisors []string
}

// Pretty print of a thesis
func (t thesis) String() string {

	text := fmt.Sprintf("Title: %s \nLink: %s\nAuthor: %s \n", t.Title, t.Link, t.Author)

	if len(t.Advisors) == 0 {
		text += "Advisors: None"
	} else if len(t.Advisors) == 1 {
		text += "Advisor: " + t.Advisors[0]
	} else if len(t.Advisors) == 2 {
		text += "Advisors: " + t.Advisors[0] + " and " + t.Advisors[1]
	} else {
		// should not have more than two advisors
		text += "Advisors: more than two advisors ( " + strings.Join(t.Advisors, ", ") + " )"
	}

	return text
}

// Pretty print of all theses
func printTheses(theses []thesis) {
	for _, t := range theses {
		fmt.Println(t)
		fmt.Println()
	}
}

// Find advisor that matches all names
func findThesesFromAdvisor(theses []thesis, advisorNames []string) {

	for _, t := range theses {
		for _, advisor := range t.Advisors {

			matchedNames := 0

			for _, names := range advisorNames {
				if strings.Contains(strings.ToLower(advisor), strings.ToLower(names)) {
					matchedNames++
				}
			}

			if matchedNames == len(advisorNames) {
				fmt.Println(t)
				fmt.Println()
			}
		}
	}
}

func createCollyCollector(theses *[]thesis, domain string, baseURL string) *colly.Collector {

	c := colly.NewCollector(
		colly.AllowedDomains(domain),
	)

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.OnResponse(func(r *colly.Response) {
		if r.StatusCode != 200 {
			fmt.Println("Error: ", r.StatusCode)
			os.Exit(1)
		} else {
			fmt.Println("Success")
			fmt.Println()
		}
	})

	c.OnHTML("li", func(e *colly.HTMLElement) {

		title := e.ChildText("h5 > a")
		if len(title) == 0 {
			return
		}

		link := e.ChildAttr("h5 > a", "href")
		if len(link) == 0 {
			return
		}
		link = baseURL + link

		authorRegex := regexp.MustCompile(`Author: (.*)\(ist\d+\)`)
		author := authorRegex.FindString(e.ChildText("h5 > small"))

		if len(author) == 0 {
			return
		}

		author = strings.Split(author, "Author: ")[1]

		// ChildText returns the concatenated text of spans so instead of
		// this: 			"Name (ist111111) Name (ist111111)"
		// we have this:   	"Name (ist111111)Name (ist111111)"
		advisorsRegex := regexp.MustCompile(`(.*?)\s\((.*?)\)`)
		c := e.ChildText("h5 > small > span")
		advisorsMatches := advisorsRegex.FindAllStringSubmatch(c, -1)

		advisors := []string{}

		// Single name
		if advisorsMatches == nil {
			name := advisorsRegex.FindString(c)
			advisors = append(advisors, name)
			// Two names
		} else {
			for _, match := range advisorsMatches {
				advisors = append(advisors, match[1]+" ("+match[2]+")")
			}
		}

		*theses = append(*theses, thesis{title, link, author, advisors})
	})

	return c
}

func main() {

	const domain = "fenix.tecnico.ulisboa.pt"
	const baseURL = "https://fenix.tecnico.ulisboa.pt/cursos/meic-a/"
	const visitURL = baseURL + "dissertacoes"

	theses := []thesis{}

	c := createCollyCollector(&theses, domain, baseURL)
	c.Visit(visitURL)

	if len(os.Args) > 1 {
		findThesesFromAdvisor(theses, os.Args[1:])
	} else {
		printTheses(theses)
	}
}
