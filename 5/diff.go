package main

import (
	"fmt"
)

func diff(a, b string, unified bool) string {
	var ap, bp string
	if unified {
		ap = "- "
		bp = "+ "
	} else {
		ap = "< "
		bp = "> "
	}

	codedA, codedB, codes := code(a, b)

	dp := DP{
		a: codedA,
		b: codedB,
	}
	diff := dp.distance(len(codedA), len(codedB))
	fmt.Println(diff)

	return prettyText(diff, codes, bp, ap)
}

type Kind int8

const (
	Insert Kind = iota + 1
	Delete
	Equal
)

type Op struct {
	Kind Kind
	Text rune
}

type DP struct {
	a, b []rune
}

func (dp DP) distance(i, j int) []Op {
	// Implement here!
	return []Op{}
}

func code(a, b string) ([]rune, []rune, map[rune]string) {
	var i, j []rune

	m := make(map[string]rune)

	aLines := splitToLines(a)
	bLines := splitToLines(b)

	count := 0
	for _, line := range aLines {
		_, ok := m[line]
		if !ok {
			m[line] = rune(count)
			count++
		}
		i = append(i, m[line])
	}

	for _, line := range bLines {
		_, ok := m[line]
		if !ok {
			m[line] = rune(count)
			count++
		}
		j = append(j, m[line])
	}

	ans := make(map[rune]string)
	for k, v := range m {
		ans[v] = k
	}

	return i, j, ans
}

func splitToLines(str string) []string {
	rs := []rune(str)
	ans := []string{}
	buf := ""
	for _, r := range rs {
		buf += string(r)
		if r == '\n' {
			ans = append(ans, buf)
			buf = ""
		}
	}
	return ans
}

func prettyText(ops []Op, codes map[rune]string, insertPrefix, deletePrefix string) string {
	res := ""

	for _, op := range ops {
		var line string

		switch op.Kind {
		case Insert:
			line = fmt.Sprintf("%s%s", insertPrefix, codes[op.Text])
		case Delete:
			line = fmt.Sprintf("%s%s", deletePrefix, codes[op.Text])
		case Equal:
			line = fmt.Sprintf("  %s", codes[op.Text])
		}
		res = fmt.Sprintf("%s%s", res, line)
	}

	return res
}
