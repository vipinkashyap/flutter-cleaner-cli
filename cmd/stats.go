package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/vipinkashyap/flutter-cleaner-cli/ui"

	"github.com/spf13/cobra"
)

var statsCmd = &cobra.Command{
    Use:   "stats [dir]",
    Short: "Show total build space used by Flutter projects",
    Run: func(cmd *cobra.Command, args []string) {
        dir := "."
        if len(args) > 0 {
            dir = args[0]
        }

        var total int64
        filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
            if filepath.Base(path) == "build" && info.IsDir() {
                total += folderSize(path)
            }
            return nil
        })
        fmt.Println(ui.TitleStyle.Render("📊 Flutter Cleaner — Stats"))
        fmt.Println(ui.SectionStyle.Render(fmt.Sprintf("📦 Total build space: %s", formatSize(total))))
    },
}

func init() {
    rootCmd.AddCommand(statsCmd)
}