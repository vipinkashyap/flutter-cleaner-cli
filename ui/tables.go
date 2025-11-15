package ui

import (
	"bytes"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
)

func RenderTableWithBorder(title string, headers []string, rows [][]string) string {
    buf := new(bytes.Buffer)

    t := table.NewWriter()
    t.SetOutputMirror(buf)

    headerRow := table.Row{}
    for _, h := range headers {
        headerRow = append(headerRow, text.FgCyan.Sprint(h))
    }
    t.AppendHeader(headerRow)

    for _, r := range rows {
        row := table.Row{}
        for _, col := range r {
            row = append(row, col)
        }
        t.AppendRow(row)
    }

    t.SetStyle(table.StyleRounded)
    t.Style().Color.Header = text.Colors{text.FgCyan}
    t.Style().Color.Row = text.Colors{text.FgWhite}

    t.Render()

    box := BorderBox.Render(buf.String())
    if title != "" {
        return TitleStyle.Render(title) + "\n" + box
    }
    return box
}