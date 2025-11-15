package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/vipinkashyap/flutter-cleaner-cli/ui"

	"github.com/spf13/cobra"
)

type Suggestion struct {
	Path       string  `json:"path"`
	AgeDays    float64 `json:"age_days"`
	SizeBytes  int64   `json:"size_bytes"`
	SizePretty string  `json:"size_pretty"`
}

var suggestFile = filepath.Join(os.TempDir(), "fclean_suggestions.json")

var suggestCmd = &cobra.Command{
	Use:   "suggest [dir]",
	Short: "Suggest projects to clean based on build age and size",
	Long:  "Analyzes Flutter projects and suggests which ones to clean based on build folder size and age.",
	Run: func(cmd *cobra.Command, args []string) {
		dir := "."
		if len(args) > 0 {
			dir = args[0]
		}

		fmt.Println(ui.SectionStyle.Render(fmt.Sprintf("🧠 Analyzing Flutter projects under %s ...", dir)))
		suggestions := []Suggestion{}
		var totalSize int64

		filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
			if filepath.Base(path) == "pubspec.yaml" {
				projectDir := filepath.Dir(path)
				buildDir := filepath.Join(projectDir, "build")

				stat, err := os.Stat(buildDir)
				if err == nil && stat.IsDir() {
					age := time.Since(stat.ModTime()).Hours() / 24
					size := folderSize(buildDir)
					totalSize += size

					suggestions = append(suggestions, Suggestion{
						Path:       projectDir,
						AgeDays:    age,
						SizeBytes:  size,
						SizePretty: formatSize(size),
					})
				}
			}
			return nil
		})

		if len(suggestions) == 0 {
			fmt.Println(ui.SuccessStyle.Render("✨ No build folders found — everything looks clean!"))
			return
		}

		// Sort by size descending
		for i := 0; i < len(suggestions)-1; i++ {
			for j := i + 1; j < len(suggestions); j++ {
				if suggestions[j].SizeBytes > suggestions[i].SizeBytes {
					suggestions[i], suggestions[j] = suggestions[j], suggestions[i]
				}
			}
		}

		// Save suggestions to a temp file
		data, _ := json.MarshalIndent(suggestions, "", "  ")
		_ = os.WriteFile(suggestFile, data, 0644)

		// Show summary
		fmt.Println(ui.TitleStyle.Render(fmt.Sprintf("📊 Found %d Flutter project(s) with build directories", len(suggestions))))
		fmt.Println(ui.SectionStyle.Render(fmt.Sprintf("📦 Total build space: %s", formatSize(totalSize))))

		rows := [][]string{}
		for i, s := range suggestions {
			rows = append(rows, []string{
				fmt.Sprintf("%d", i+1),
				s.Path,
				fmt.Sprintf("%.0f days", s.AgeDays),
				s.SizePretty,
			})
		}
		table := ui.RenderTableWithBorder("💡 Suggestions (by size)", []string{"#", "Project", "Age", "Size"}, rows)
		fmt.Println(table)

		fmt.Println(ui.SectionStyle.Render("👉 To clean a suggested project:"))
		fmt.Println(ui.MutedText.Render("   fclean clean --suggest <number>"))
		fmt.Println(ui.MutedText.Render(fmt.Sprintf("📁 Suggestion list saved at: %s", suggestFile)))
	},
}

func init() {
	rootCmd.AddCommand(suggestCmd)
}