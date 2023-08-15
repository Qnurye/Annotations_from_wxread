package main

import (
	"fmt"
	"os"
	"path/filepath"
	"qnurye/Annotations_from_wxread/pkg/annotation"
	"qnurye/Annotations_from_wxread/pkg/config"
	"qnurye/Annotations_from_wxread/pkg/instructions"
	"strconv"
	"time"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}

	c, err := os.ReadFile(cfg.SourceDirectory)
	if err != nil {
		panic(err)
	}

	seq := instructions.Parse(string(c))

	var (
		title     string
		source    string
		author    string
		chapter   string
		comment   string
		quotation string
		date      time.Time
	)

	for i, instruction := range seq {
		switch instruction.Type {
		case instructions.Source:
			source = instruction.Value
		case instructions.Author:
			author = instruction.Value
		case instructions.Chapter:
			chapter = instruction.Value
		case instructions.Comment:
			comment = instruction.Value
		case instructions.Quotation:
			quotation = instruction.Value
		case instructions.Date:
			date, _ = time.Parse("time.RFC3339", instruction.Value)
		}

		if quotation != "" {
			if cfg.AnnotationTitleConfigurable {
				fmt.Printf(`#%d @ sequences:
Source: "%s" by %s on chapter "%s"
"%s"`, i+1, source, author, chapter, quotation)
				if (date != time.Time{}) {
					fmt.Printf(`annotated on %s`, date.Format("2006-01-02T15:04"))
				}
				if comment != "" {
					fmt.Printf(`commented %s`, comment)
				}
				fmt.Print("\nTitle: ")
				_, err := fmt.Scanln(&title)
				if err != nil {
					return
				}
			} else {
				if comment != "" {
					title = comment
				} else {
					title = chapter + "-" + strconv.Itoa(i)
				}
			}

			o, err := filepath.Abs(cfg.OutputDirectory)
			if err != nil {
				return
			}

			annotation.Save(annotation.Annotation{
				Source:    source,
				Title:     title,
				Date:      date,
				Author:    author,
				Quotation: quotation,
				Chapter:   chapter,
				Comment:   comment,
			}, cfg.TemplateFilePath, o)

			title, quotation, comment = "", "", ""
			date = time.Time{}
		}
	}
}
