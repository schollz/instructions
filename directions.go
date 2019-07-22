package instructions

//go:generate go run corpus/main.go
//go:generate gofmt -s -w corpus.go

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"

	log "github.com/schollz/logger"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

func Parse(htmlString string) (directions []string, err error) {
	directions, err = getDirectionLinesInHTML(htmlString)
	return
}

func getDirectionLinesInHTML(htmlS string) (lineInfos []string, err error) {
	doc, err := html.Parse(bytes.NewReader([]byte(htmlS)))
	if err != nil {
		return
	}
	var f func(n *html.Node, lineInfos *[]string, bestScore *int) (s string, done bool)
	f = func(n *html.Node, lineInfos *[]string, bestScore *int) (s string, done bool) {
		childrenstring := []string{}
		// log.Tracef("%+v", n)
		score := 0
		isScript := n.DataAtom == atom.Script
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			if isScript {
				// try to capture JSON and if successful, do a hard exit
				lis, errJSON := extractLinesFromJavascript(c.Data)
				score, _ := scoreLines(lis)
				if errJSON == nil && len(lis) > 2 && score > *bestScore {
					log.Tracef("got ingredients from JSON")
					*bestScore = score
					*lineInfos = lis
				}
			}
			var childText string
			childText, done = f(c, lineInfos, bestScore)
			if done {
				return
			}
			if childText != "" {
				scoreOfLine, lineInfo := scoreLine(childText)
				childrenstring = append(childrenstring, lineInfo)
				score += scoreOfLine
			}
		}
		if score > *bestScore && len(childrenstring) < 15 && len(childrenstring) > 2 {
			log.Tracef("score: %d, childs: ['%s']", score, strings.Join(childrenstring, "', '"))
			*lineInfos = childrenstring
			*bestScore = score
			for _, child := range childrenstring {
				log.Tracef("[%s]", child)
			}
		}
		if len(childrenstring) > 0 {
			// fmt.Println(childrenstring)
			childrenText := make([]string, len(childrenstring))
			for i := range childrenstring {
				childrenText[i] = childrenstring[i]
			}
			s = strings.Join(childrenText, " ")
		} else if n.DataAtom == 0 && strings.TrimSpace(n.Data) != "" {
			s = strings.TrimSpace(n.Data)
		}
		return
	}
	bestScore := 0
	f(doc, &lineInfos, &bestScore)
	return
}

func extractLinesFromJavascript(jsString string) (lineInfo []string, err error) {

	var arrayMap = []map[string]interface{}{}
	var regMap = make(map[string]interface{})
	err = json.Unmarshal([]byte(jsString), &regMap)
	if err != nil {
		err = json.Unmarshal([]byte(jsString), &arrayMap)
		if err != nil {
			return
		}
		if len(arrayMap) == 0 {
			err = fmt.Errorf("nothing to parse")
			return
		}
		parseMap(arrayMap[0], &lineInfo)
		err = nil
	} else {
		parseMap(regMap, &lineInfo)
		err = nil
	}

	return
}

func parseMap(aMap map[string]interface{}, lineInfo *[]string) {
	for _, val := range aMap {
		switch val.(type) {
		case map[string]interface{}:
			parseMap(val.(map[string]interface{}), lineInfo)
		case []interface{}:
			parseArray(val.([]interface{}), lineInfo)
		default:
			// fmt.Println(key, ":", concreteVal)
		}
	}
}

func parseArray(anArray []interface{}, lineInfo *[]string) {
	concreteLines := []string{}
	for _, val := range anArray {
		switch concreteVal := val.(type) {
		case map[string]interface{}:
			parseMap(val.(map[string]interface{}), lineInfo)
		case []interface{}:
			parseArray(val.([]interface{}), lineInfo)
		default:
			switch v := concreteVal.(type) {
			case string:
				concreteLines = append(concreteLines, v)
			}
		}
	}

	score, li := scoreLines(concreteLines)
	log.Trace(score, li)
	if score > 10 {
		*lineInfo = li
	}
	return
}

func scoreLines(lines []string) (score int, lineInfo []string) {
	if len(lines) < 2 {
		return
	}
	lineInfo = make([]string, len(lines))
	for i, line := range lines {
		var scored int
		scored, lineInfo[i] = scoreLine(line)
		score += scored
	}
	return
}

func scoreLine(line string) (score int, lineInfo string) {
	lineInfo = strings.TrimSpace(strings.ToLower(line))

	if len(lineInfo) > 700 {
		score = 700 - len(lineInfo)
	}
	if len(lineInfo) < 40 {
		score = len(lineInfo) - 40
		return
	}

	pos := 0.0
	for _, word := range corpusDirections {
		if strings.Contains(lineInfo, word) {
			pos++
		}
	}
	score += int(pos)

	pos = 0.0
	for _, word := range corpusDirectionsNeg {
		if strings.Contains(lineInfo, word) {
			pos++
		}
	}
	score -= int(pos)

	firstChar := string(strings.TrimSpace(line)[0])
	if strings.ToUpper(firstChar) == firstChar {
		score++
	}

	score -= strings.Count(lineInfo, ":")
	score -= strings.Count(lineInfo, "<")
	score -= strings.Count(lineInfo, ">")
	score -= strings.Count(lineInfo, `"`)

	lineInfo = strings.TrimSpace(line)
	return
}
