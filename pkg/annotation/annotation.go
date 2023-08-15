package annotation

import (
	"os"
	"path"
	"strings"
	"time"
)

type Annotation struct {
	Source    string
	Title     string
	Date      time.Time
	Author    string
	Quotation string
	Chapter   string
	Comment   string
}

func Save(a Annotation, templatePath string, output string) string {
	t, _ := os.ReadFile(templatePath)
	if (a.Date == time.Time{}) {
		a.Date = time.Now()
	}
	c := strings.ReplaceAll(string(t), "<% annotation.time %>", a.Date.Format("2006-01-02T15:04"))
	c = strings.ReplaceAll(c, "<% annotation.source %>", a.Source)
	c = strings.ReplaceAll(c, "<% annotation.author %>", a.Author)
	c = strings.ReplaceAll(c, "<% annotation.chapter %>", a.Chapter)
	c = strings.ReplaceAll(c, "<% annotation.comment %>", a.Comment)
	c = strings.ReplaceAll(c, "<% annotation.quotation %>", a.Quotation)

	fp := path.Join(output, a.Source)
	err := os.MkdirAll(fp, os.ModePerm)
	if err != nil {
		panic(err)
	}
	fp = path.Join(fp, a.Title+".md")
	f, err := os.Create(fp)
	if err != nil {
		panic(err)
	}

	_, err = f.Write([]byte(c))
	if err != nil {
		panic(err)
	}

	return ""
}
