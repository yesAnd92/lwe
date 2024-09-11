package sqllog

import (
	"fmt"
	"strings"

	"github.com/dlclark/regexp2"
	"github.com/pkg/errors"
)

/**
* mybatis sql log parse
 */

var (
	separatorPreparing = "Preparing:"
	separatorParameter = "Parameters:"
	preparingPattern   = separatorPreparing + "(.*?)(?=\n|\r|\r\n)"
	parameterPattern   = separatorParameter + "(.*?)(?=\n|\r|\r\n)"
)

func ParseMybatisSqlLog(sqlLog string) (string, error) {

	if !(strings.HasSuffix(sqlLog, "\n") || strings.HasSuffix(sqlLog, "\r\n")) {
		sqlLog += "\n"
	}

	prepare, err := extractPattern(preparingPattern, sqlLog)
	if err != nil {
		return "", errors.New("parsing Preparing fail!")
	}

	prepare = strings.Replace(prepare, separatorPreparing, "", -1)

	param, err := extractPattern(parameterPattern, sqlLog)
	if err != nil {
		return "", errors.New("parsing Parameters fail!")
	}

	params := strings.Split(strings.Replace(param, separatorParameter, "", -1), ",")

	var values []string
	for _, p := range params {
		value := extractValue(p)
		values = append(values, value)
	}

	// Replace ? in prepare with values from values
	for _, value := range values {
		prepare = strings.Replace(prepare, "?", value, 1)
	}

	//trim blank space
	prepare = strings.TrimSpace(prepare)

	return prepare, nil
}

var typesNeedQuotes = []string{"String", "Timestamp", "Date", "Time"}

func needQuotes(s string) bool {
	for _, t := range typesNeedQuotes {
		if strings.Contains(s, t) {
			return true
		}
	}
	return false
}

func extractPattern(pattern, log string) (string, error) {
	re := regexp2.MustCompile(pattern, 0)
	match, err := re.FindStringMatch(log)
	if err != nil {
		return "", fmt.Errorf("regex match failed: %w", err)
	}
	if match == nil {
		return "", fmt.Errorf("no match found")
	}
	return match.String(), nil
}

// First, remove whitespace from both ends of the string.
// Iterate through each character in the string:
// When encountering a left parenthesis '(', if it's the first one, record its position.
// When encountering a right parenthesis ')', check if it matches the outermost pair of parentheses.
// If the outermost pair of parentheses is found, return the content before the parentheses.
// If no matching pair of parentheses is found, return the entire string.
func extractValue(s string) string {
	s = strings.TrimSpace(s)
	lastOpenParen := -1
	parenCount := 0

	for i, char := range s {
		if char == '(' {
			if parenCount == 0 {
				lastOpenParen = i
			}
			parenCount++
		} else if char == ')' {
			parenCount--
			if parenCount == 0 && lastOpenParen != -1 {
				innerContent := s[lastOpenParen+1 : i]
				result := strings.TrimSpace(s[:lastOpenParen])
				if needQuotes(innerContent) {
					return "'" + result + "'"
				}
				return result
			}
		}
	}

	return s
}
