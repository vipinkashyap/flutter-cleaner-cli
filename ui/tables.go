package ui

import (
	"bytes"

	"github.com/olekukonko/tablewriter"
)

func RenderTableWithBorder(title string, headers []string, rows [][]string) string {
	buf := new(bytes.Buffer)

	table := tablewriter.NewWriter(buf)
	table.SetHeader(headers)
	table.SetBorder(false)
	table.SetRowSeparator("─")
	table.SetAlignment(tablewriter.ALIGN_LEFT)

	// Colors
	headerColors := []tablewriter.Colors{{tablewriter.FgHiCyanColor}}
	for range headers {
		headerColors = append(headerColors, tablewriter.Colors{tablewriter.FgHiCyanColor})
	}
	table.SetHeaderColor(headerColors...)

	rowColors := []tablewriter.Colors{{tablewriter.FgWhiteColor}}
	table.SetColumnColor(rowColors...)

	table.AppendBulk(rows)
	table.Render()

	content := BorderBox.Render(buf.String())

	if title != "" {
		return TitleStyle.Render(title) + "\n" + content
	}
	return content
}