package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/spf13/cobra"
	"github.com/vipinkashyap/flutter-cleaner-cli/ui"
)

// CleanResult holds info about a cleaned project and how much space was freed.
type CleanResult struct {
    Path  string
    Freed string
}

var cleanResults []CleanResult
var dryRun bool



var all bool
var parallel int
var suggestIndex int

var cleanCmd = &cobra.Command{
    Use:   "clean [dir]",
    Short: "Clean one or all Flutter projects",
    Run: func(cmd *cobra.Command, args []string) {
        // If cleaning a suggested project
        if suggestIndex > 0 {
            suggestFile := filepath.Join(os.TempDir(), "fclean_suggestions.json")
            data, err := os.ReadFile(suggestFile)
            if err != nil {
                fmt.Println(ui.ErrorStyle.Render(fmt.Sprintf("❌ Could not read suggestions file: %v", err)))
                return
            }

            var suggestions []Suggestion
            if err := json.Unmarshal(data, &suggestions); err != nil {
                fmt.Println(ui.ErrorStyle.Render("❌ Invalid suggestions file format."))
                return
            }

            if suggestIndex < 1 || suggestIndex > len(suggestions) {
                fmt.Println(ui.ErrorStyle.Render(fmt.Sprintf("⚠️ Invalid suggestion number. Choose between 1 and %d", len(suggestions))))
                return
            }

            s := suggestions[suggestIndex-1]
            fmt.Println(ui.SectionStyle.Render("🧹 Cleaning suggested project #%d:"), suggestIndex)
            fmt.Println(ui.SectionStyle.Render(fmt.Sprintf("➡️  %s (%.0f days old, %s)", s.Path, s.AgeDays, s.SizePretty)))
            runFlutterClean(s.Path)
            return
        }

        // Normal cleaning logic
        if all {
            fmt.Println(ui.SectionStyle.Render("🧹 Cleaning all Flutter projects..."))
            projects := []string{}
            filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
                if filepath.Base(path) == "pubspec.yaml" {
                    projects = append(projects, filepath.Dir(path))
                }
                return nil
            })
            cleanAllProjects(projects)
            return
        }

        if len(args) == 0 {
            fmt.Println(ui.ErrorStyle.Render("Provide a project path, use --all, or try --suggest <num>"))
            return
        }

        runFlutterClean(args[0])
    },
}

func runFlutterClean(dir string) {
    progress := ui.NewEmojiProgress(100)
    before := folderSize(dir)
    start := time.Now()
    fmt.Println(ui.SectionStyle.Render(fmt.Sprintf("➡️  Cleaning: %s", dir)))

    if dryRun {
    freed := int64(0)
    fmt.Println(ui.WarningStyle.Render(
        fmt.Sprintf("🔎 Dry run: would clean %s (Freeable %s)", dir, formatSize(before)),
    ))
    cleanResults = append(cleanResults, CleanResult{Path: dir, Freed: formatSize(freed)})
    return
    }

    for i := 0; i < 100; i++ {
        progress.Update(i)
        fmt.Print("\r" + progress.Render())
    }
    fmt.Println()

    cmd := exec.Command("flutter", "clean")
    cmd.Dir = dir
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    err := cmd.Run()

    duration := time.Since(start).Round(time.Second)
    after := folderSize(dir)
    freed := before - after
    if err != nil {
        fmt.Println(ui.ErrorStyle.Render(fmt.Sprintf("❌ %s (Error: %v)", dir, err)))
    } else {
        fmt.Println(ui.SuccessStyle.Render(fmt.Sprintf("✅ %s (Freed %s in %v)", dir, formatSize(freed), duration)))
        cleanResults = append(cleanResults, CleanResult{Path: dir, Freed: formatSize(freed)})
    }
}

func cleanAllProjects(projects []string) {
    var wg sync.WaitGroup
    sem := make(chan struct{}, parallel)
    for _, dir := range projects {
        wg.Add(1)
        go func(d string) {
            defer wg.Done()
            sem <- struct{}{}
            runFlutterClean(d)
            <-sem
        }(dir)
    }
    wg.Wait()
    // Compute total freed size
    var totalFreed int64
    for _, r := range cleanResults {
        sz, _ := strconv.ParseFloat(strings.Split(r.Freed, " ")[0], 64)
        if strings.Contains(r.Freed, "GB") {
            totalFreed += int64(sz * 1024 * 1024 * 1024)
        } else {
            totalFreed += int64(sz * 1024 * 1024)
        }
    }
    // Render summary table
    rows := [][]string{}
    for _, r := range cleanResults {
        rows = append(rows, []string{r.Path, r.Freed})
    }
    rows = append(rows, []string{"TOTAL", formatSize(totalFreed)})
    summary := ui.RenderTableWithBorder("🧹 Cleaning Summary", []string{"Project", "Freed"}, rows)
    fmt.Println(summary)
    fmt.Println(ui.SuccessStyle.Render("💾 All done."))
    cleanResults = []CleanResult{}
}

func init() {
    cleanCmd.Flags().BoolVarP(&all, "all", "a", false, "Clean all projects recursively")
    cleanCmd.Flags().IntVarP(&parallel, "parallel", "p", 4, "Number of parallel clean tasks")
    cleanCmd.Flags().IntVarP(&suggestIndex, "suggest", "s", 0, "Clean a suggested project from last analysis")
    cleanCmd.Flags().BoolVar(&dryRun, "dry-run", false, "Preview what would be cleaned without making changes")
}