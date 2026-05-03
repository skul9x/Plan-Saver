package main

import (
	"fmt"
	"regexp"
	"strings"
)

type MarkdownBlock struct {
	FileName string
	Content  string
}

// ExtractMarkdownBlocks finds all ```md ... ``` blocks and extracts filename from the first line
func ExtractMarkdownBlocks(content string) []MarkdownBlock {
	// Match ```md or ```markdown (case insensitive), handles \r\n
	re := regexp.MustCompile(`(?i)'''(?:md|markdown)\r?\n([\s\S]*?)\r?\n'''`)
	// Replace backticks with single quotes for the tool call as backticks are tricky in strings here, 
	// but I'll use a safer approach below.
	
	backtick := "`"
	pattern := fmt.Sprintf("(?i)%s{3}(?:md|markdown)[ \\t]*\\r?\\n([\\s\\S]*?)\\n%s{3}", backtick, backtick)
	re = regexp.MustCompile(pattern)
	
	matches := re.FindAllStringSubmatch(content, -1)

	var blocks []MarkdownBlock
	for _, match := range matches {
		if len(match) < 2 {
			continue
		}
		rawContent := match[1]
		// Normalize line endings to \n for easier splitting
		normalized := strings.ReplaceAll(rawContent, "\r\n", "\n")
		lines := strings.SplitN(normalized, "\n", 2)
		if len(lines) == 0 {
			continue
		}

		firstLine := strings.TrimSpace(lines[0])
		if strings.HasPrefix(firstLine, "#") {
			fileName := strings.TrimLeft(firstLine, "# ")
			if !strings.HasSuffix(fileName, ".md") {
				fileName += ".md"
			}
			
			actualContent := ""
			if len(lines) > 1 {
				actualContent = strings.TrimSpace(lines[1])
			}

			blocks = append(blocks, MarkdownBlock{
				FileName: fileName,
				Content:  actualContent,
			})
		}
	}

	return blocks
}

// SanitizeFileName removes characters that are invalid for filenames
func SanitizeFileName(name string) string {
	// Basic sanitation: remove / \ : * ? " < > |
	reg := regexp.MustCompile(`[\\/:*?"<>|]`)
	return reg.ReplaceAllString(name, "_")
}
