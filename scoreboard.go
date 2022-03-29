package hangman_web

import (
	"bufio"
	"net/http"
	"os"
	"sort"
	"strings"
	"text/template"
)

var ScoreTab []Score

type ScoreStruct struct {
	Tab []Score
}

type Score struct {
	Name  string
	Diff  string
	Time  string
	Score string
}

func unique(sample []Score) []Score {
	var unique []Score
sampleLoop:
	for _, v := range sample {
		for i, u := range unique {
			if v.Name == u.Name && v.Diff == u.Diff && v.Time == u.Time && v.Score == u.Score {
				unique[i] = v
				continue sampleLoop
			}
		}
		unique = append(unique, v)
	}
	return unique
}

func Scoreboard(w http.ResponseWriter, r *http.Request) {
	file, _ := os.OpenFile("scores.txt", os.O_RDWR, 0644)
	scanner := bufio.NewScanner(file)
	ScoreTab = []Score{}

	for scanner.Scan() {
		split := strings.Split(scanner.Text(), "|")
		sc := Score{Name: split[0], Diff: split[1], Time: split[2], Score: split[3]}
		ScoreTab = append(ScoreTab, sc)
	}

	ScoreTab = unique(ScoreTab)

	sort.Slice(ScoreTab, func(i, j int) bool {
		return ScoreTab[i].Score > ScoreTab[j].Score
	})

	tmpl2, _ := template.ParseFiles("../html/scoreboard.html")

	data := ScoreStruct{
		Tab: ScoreTab,
	}

	_ = tmpl2.Execute(w, data)
}
