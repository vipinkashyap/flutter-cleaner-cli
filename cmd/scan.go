package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/vipinkashyap/flutter-cleaner-cli/ui"

	"github.com/spf13/cobra"
)

var fast bool

var scanCmd = &cobra.Command{
	Use:   "scan [dir]",
	Short: "Scan a directory for Flutter projects and show summary",
	Run: func(cmd *cobra.Command, args []string) {
		dir := "."
		if len(args) > 0 {
			dir = args[0]
		}

		// Expand ~ to home dir
		if dir == "~" {
			home, _ := os.UserHomeDir()
			dir = home
		}

		// Ask for confirmation if scanning home or root
		if dir == "/" || dir == os.Getenv("HOME") {
			fmt.Println(ui.WarningStyle.Render("⚠️  Scanning this directory may take a long time."))
			ok, _ := ui.AskConfirm("Continue anyway?")
			if !ok {
				fmt.Println(ui.ErrorStyle.Render("❌ Scan cancelled."))
				return
			}
		}

		start := time.Now()
		fmt.Println(ui.SectionStyle.Render("🔍 Scanning..."))
		if fast {
			fmt.Println(ui.WarningStyle.Render("⚡ Fast mode: skipping common system and cache folders"))
		}

		var totalSize int64
		var projects int
		progress := ui.NewEmojiProgress(100)
		current := 0

		filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
			current = (current + 1) % 100
			progress.Update(current)
			fmt.Print("\r" + progress.Render())

			if err != nil {
				return nil
			}

			// Skip heavy dirs if fast mode
			if fast && info.IsDir() {
				skipDirs := []string{
					".git", "node_modules", ".dart_tool", ".pub-cache", "Library",
					"Applications", ".Trash", "Caches", "Android", ".gradle",
				}
				for _, s := range skipDirs {
					if info.Name() == s {
						return filepath.SkipDir
					}
				}
			}

			if filepath.Base(path) == "pubspec.yaml" {
				pDir := filepath.Dir(path)
				size := folderSize(pDir)
				totalSize += size
				projects++
				fmt.Println()
				fmt.Println(ui.SectionStyle.Render(fmt.Sprintf("📁 %s (%s)", pDir, formatSize(size))))
			}
			return nil
		})

		duration := time.Since(start).Round(time.Second)
		fmt.Println()
		fmt.Println(ui.SuccessStyle.Render(fmt.Sprintf("✅ Found %d Flutter project(s)", projects)))
		fmt.Println(ui.SectionStyle.Render(fmt.Sprintf("📦 Total size: %s", formatSize(totalSize))))
		fmt.Println(ui.MutedText.Render(fmt.Sprintf("⏱️  Scan completed in %v", duration)))
	},
}

func init() {
	scanCmd.Flags().BoolVarP(&fast, "fast", "f", false, "Enable fast mode (skip system/cache folders)")
}