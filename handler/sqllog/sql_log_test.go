package sqllog

import (
	"testing"
)

func TestParseMybatisSqlLog(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
		hasError bool
	}{
		{
			name: "Basic SQL log",
			input: `Preparing: SELECT * FROM users WHERE id = ?
					Parameters: 1(Integer)`,
			expected: "SELECT * FROM users WHERE id = 1",
			hasError: false,
		},
		{
			name: "SQL log with string parameter",
			input: `Preparing: SELECT * FROM users WHERE name = ?
					Parameters: John(String)`,
			expected: "SELECT * FROM users WHERE name = 'John'",
			hasError: false,
		},
		{
			name: "SQL log with timestamp parameter",
			input: `Preparing: SELECT * FROM logs WHERE created_at > ?
					Parameters: 2023-05-01 10:00:00(Timestamp)`,
			expected: "SELECT * FROM logs WHERE created_at > '2023-05-01 10:00:00'",
			hasError: false,
		},
		{
			name: "SQL log with multiple parameters",
			input: `Preparing: INSERT INTO users (name, age, created_at) VALUES (?, ?, ?)
					Parameters: Alice(String), 30(Integer), 2023-05-01 12:00:00(Timestamp)`,
			expected: "INSERT INTO users (name, age, created_at) VALUES ('Alice', 30, '2023-05-01 12:00:00')",
			hasError: false,
		},
		{
			name:     "Invalid SQL log format",
			input:    "This is not a valid SQL log",
			expected: "",
			hasError: true,
		},
		{
			name:     "SQL log missing parameters",
			input:    `Preparing: SELECT * FROM users WHERE id = ?`,
			expected: "",
			hasError: true,
		},
		{
			name: "Mismatch between placeholders and parameters",
			input: `Preparing: SELECT * FROM users WHERE id = ? AND name = ?
					Parameters: 1(Integer)`,
			expected: "",
			hasError: true,
		},
		{
			name: "Like query",
			input: `Preparing: SELECT * FROM users WHERE id  name LIKE CONCAT(?, '%')
					Parameters: yesA(String)`,
			expected: "SELECT * FROM users WHERE id  name LIKE CONCAT('yesA', '%')",
			hasError: false,
		},
		{
			name: "Insert SQL log",
			input: `Preparing: INSERT INTO products (name, price, category) VALUES (?, ?, ?)
					Parameters: New Product(String), 19.99(Double), Electronics(String)`,
			expected: "INSERT INTO products (name, price, category) VALUES ('New Product', 19.99, 'Electronics')",
			hasError: false,
		},
		{
			name: "Update SQL log",
			input: `Preparing: UPDATE orders SET status = ?, updated_at = ? WHERE id = ?
					Parameters: Shipped(String), 2023-05-02 15:30:00(Timestamp), 1001(Integer)`,
			expected: "UPDATE orders SET status = 'Shipped', updated_at = '2023-05-02 15:30:00' WHERE id = 1001",
			hasError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ParseMybatisSqlLog(tt.input)

			if err != nil {
				t.Logf("Error: %v", err)
			}

			if tt.hasError {
				if err == nil {
					t.Errorf("Expected an error, but got none")
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				} else if result != tt.expected {
					t.Errorf("Result mismatch\nExpected: %s\nActual: %s", tt.expected, result)
					// Additional debug information
					t.Logf("Expected length: %d", len(tt.expected))
					t.Logf("Actual length: %d", len(result))
					// Character-by-character comparison
					for i := 0; i < len(tt.expected) && i < len(result); i++ {
						if tt.expected[i] != result[i] {
							t.Logf("First mismatch at position %d: Expected '%c', Actual '%c'", i, tt.expected[i], result[i])
							break
						}
					}
				}
			}
		})
	}
}
