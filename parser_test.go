package main

import (
	"os"
	"testing"
)

func TestExtractMarkdownBlocks(t *testing.T) {
	input := `
Chào anh, đây là demo.

` + "```md" + `
#plan.md
Content of plan
` + "```" + `

Đây là block khác:

` + "```markdown" + `
#phase-01.md
Content of phase 01
` + "```" + `

Block không có tên:
` + "```md" + `
Nội dung không có hashtag filename
` + "```" + `
`

	blocks := ExtractMarkdownBlocks(input)

	if len(blocks) != 2 {
		t.Errorf("Expected 2 blocks, got %d", len(blocks))
	}

	if blocks[0].FileName != "plan.md" {
		t.Errorf("Expected first block filename plan.md, got %s", blocks[0].FileName)
	}

	if blocks[1].FileName != "phase-01.md" {
		t.Errorf("Expected second block filename phase-01.md, got %s", blocks[1].FileName)
	}
}

func TestExtractFromRealFile(t *testing.T) {
	data, err := os.ReadFile("answer.txt")
	if err != nil {
		t.Skip("answer.txt not found, skipping real file test")
		return
	}

	blocks := ExtractMarkdownBlocks(string(data))
	expectedCount := 4
	if len(blocks) != expectedCount {
		t.Errorf("Expected %d blocks from answer.txt, got %d", expectedCount, len(blocks))
	}

	if len(blocks) > 0 && blocks[0].FileName != "plan.md" {
		t.Errorf("Expected first block filename plan.md, got %s", blocks[0].FileName)
	}
}

func TestSanitizeFileName(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"normal.md", "normal.md"},
		{"file/with/slash.md", "file_with_slash.md"},
		{"invalid:char.md", "invalid_char.md"},
		{"test*file.md", "test_file.md"},
	}

	for _, tt := range tests {
		result := SanitizeFileName(tt.input)
		if result != tt.expected {
			t.Errorf("SanitizeFileName(%s) = %s; want %s", tt.input, result, tt.expected)
		}
	}
}
