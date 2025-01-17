package main

import (
	"os"
	"testing"
)

func TestFindDuplicates(t *testing.T) {
	tests := []struct {
		name     string
		filename string
		expected []string
	}{
		{
			name:     "no duplicates",
			filename: "testdata/no_duplicates.json",
			expected: []string{},
		},
		{
			name:     "no duplicates in array",
			filename: "testdata/no_duplicates_in_array.json",
			expected: []string{},
		},
		{
			name:     "one duplicate",
			filename: "testdata/one_duplicate.json",
			expected: []string{
				"  [.key1]\n    line 4\n    line 2",
			},
		},
		{
			name:     "multiple duplicates",
			filename: "testdata/multiple_duplicates.json",
			expected: []string{
				"  [.key1]\n    line 4\n    line 2",
				"  [.key1]\n    line 5\n    line 2",
				"  [.key2]\n    line 7\n    line 3",
			},
		},
		{
			name:     "multiple duplicates",
			filename: "testdata/duplicate_inside_object.json",
			expected: []string{
				"  [.book.part.chapter1]\n    line 18\n    line 4",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			content, err := os.ReadFile(tt.filename)
			if err != nil {
				t.Fatalf("Failed to read test support file: %v", err)
			}

			duplicates := findDuplicates(content)

			if len(duplicates) != len(tt.expected) {
				t.Errorf("Expected %v duplicates, got %v", len(tt.expected), len(duplicates))
				return
			}

			for i, expected := range tt.expected {
				if duplicates[i] != expected {
					t.Errorf("Expected %v, got %v", expected, duplicates[i])
				}
			}
		})
	}
}
