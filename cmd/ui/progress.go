package ui

import (
	"fmt"
	"strings"
)

type EmojiProgress struct {
	Total int
	Width int
	Emoji string
	Dot   string
	Value int
}

func NewEmojiProgress(total int) *EmojiProgress {
	return &EmojiProgress{
		Total: total,
		Width: 20,
		Emoji: "🧹",
		Dot:   "·",
		Value: 0,
	}
}

func (p *EmojiProgress) Update(v int) {
	p.Value = v
}

func (p *EmojiProgress) Render() string {
	if p.Total <= 0 {
		return ""
	}

	pct := float64(p.Value) / float64(p.Total)
	filled := int(pct * float64(p.Width))

	bar := strings.Repeat(p.Emoji, filled) +
		strings.Repeat(p.Dot, p.Width-filled)

	return fmt.Sprintf("[%s] %3.0f%%", bar, pct*100)
}