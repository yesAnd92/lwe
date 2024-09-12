package sqllog

import (
	"fmt"
	"strings"

	"github.com/dlclark/regexp2"
)

const (
	separatorPreparing = "Preparing:"
	separatorParameter = "Parameters:"
)

var (
	preparingPattern = separatorPreparing + "(.*?)(?=\n|\r|\r\n)"
	parameterPattern = separatorParameter + "(.*?)(?=\n|\r|\r\n)"
)

// ParseMybatisSqlLog parses a Mybatis SQL log and returns the formatted SQL query.
func ParseMybatisSqlLog(sqlLog string) (string, error) {
	if !strings.HasSuffix(sqlLog, "\n") && !strings.HasSuffix(sqlLog, "\r\n") {
		sqlLog += "\n"
	}

	prepare, err := extractPattern(preparingPattern, sqlLog)
	if err != nil {
		return "", fmt.Errorf("parsing Preparing section: %w", err)
	}

	prepare = strings.ReplaceAll(prepare, separatorPreparing, "")

	param, err := extractPattern(parameterPattern, sqlLog)
	if err != nil {
		return "", fmt.Errorf("parsing Parameters section: %w", err)
	}

	params := strings.Split(strings.ReplaceAll(param, separatorParameter, ""), ",")

	values := make([]string, 0, len(params))
	for _, p := range params {
		values = append(values, extractValue(p))
	}

	if strings.Count(prepare, "?") != len(values) {
		return "", fmt.Errorf("mismatch between placeholders (?) and parameters count")
	}

	// Replace ? in prepare with values from values
	var result strings.Builder
	paramIndex := 0
	for _, char := range prepare {
		if char == '?' {
			result.WriteString(values[paramIndex])
			paramIndex++
		} else {
			result.WriteRune(char)
		}
	}

	return strings.TrimSpace(result.String()), nil
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
