package pkg

import (
	"testing"
)

func TestCompareJSONMaps(t *testing.T) {
	tests := []struct {
		name      string
		jsonStr1  string
		jsonStr2  string
		expected  bool
		expectErr bool
	}{
		{
			name:      "Identical JSON strings",
			jsonStr1:  `{"name":"John", "age":30}`,
			jsonStr2:  `{"name":"John", "age":30}`,
			expected:  true,
			expectErr: false,
		},
		{
			name:      "Identical JSON strings with different field order",
			jsonStr1:  `{"age":30, "name":"John"}`,
			jsonStr2:  `{"name":"John", "age":30}`,
			expected:  true,
			expectErr: false,
		},
		{
			name:      "Differing JSON strings",
			jsonStr1:  `{"name":"John", "age":30}`,
			jsonStr2:  `{"name":"Jane", "age":25}`,
			expected:  false,
			expectErr: false,
		},
		{
			name:      "Invalid JSON string in jsonStr1",
			jsonStr1:  `{"name": "John"`,
			jsonStr2:  `{"name":"John", "age":30}`,
			expected:  false,
			expectErr: true,
		},
		{
			name:      "Invalid JSON string in jsonStr2",
			jsonStr1:  `{"name":"John", "age":30}`,
			jsonStr2:  `{"age":30, "name"}`,
			expected:  false,
			expectErr: true,
		},
		{
			name:      "Empty JSON strings",
			jsonStr1:  `{}`,
			jsonStr2:  `{}`,
			expected:  true,
			expectErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := CompareJSONMaps(tt.jsonStr1, tt.jsonStr2)
			if (err != nil) != tt.expectErr {
				t.Errorf("CompareJSONMaps() error = %v, expectErr %v", err, tt.expectErr)
				return
			}
			if result != tt.expected {
				t.Errorf("CompareJSONMaps() = %v, expected %v", result, tt.expected)
			}
		})
	}
}
