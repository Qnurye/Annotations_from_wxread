package instructions

import (
	"qnurye/Annotations_from_wxread/pkg/utils"
	"strings"
)

type Type int

const (
	Source Type = 1 << iota
	Author
	NoteNum
	Chapter
	Comment
	Quotation
	Date
	Null = 0
)

type Instruction struct {
	Type  Type
	Value string
}

type Sequence []Instruction

var seq Sequence

func Parse(text string) Sequence {
	ps := strings.Split(text, "\n")

	for _, p := range ps {
		var (
			t Type
			v string
			r = []rune(p)
		)

		if p == "" {
			continue
		}

		switch utils.FirstCharsOf(p, 1) {
		case "《":
			t = Source
			v = string(r[1 : len(r)-1])
		case "◆":
			t = Chapter
			v = string(r[3:])
		case ">":
			t = Quotation
			v = string(r[3:])
		default:
			switch utils.LastCharsOf(p, 3) {
			case "个笔记":
				t = NoteNum
				v = string(r[:len(r)-3])
			case "表想法":
				t = Date
				v = string(r[:len(r)-5])
			default:
				if len(seq) == 0 {
					break
				}

				switch seq[len(seq)-1].Type {
				case Source:
					t = Author
					v = p
				case Date:
					t = Comment
					v = p
				case Quotation:
					seq[len(seq)-1].Value += "\n" + p
					break
				}
			}
		}

		if t != Null {
			i := Instruction{t, v}
			seq = append(seq, i)
		}
	}

	return seq
}
