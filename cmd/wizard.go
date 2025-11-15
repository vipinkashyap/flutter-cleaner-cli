package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/vipinkashyap/flutter-cleaner-cli/ui"

	"github.com/spf13/cobra"
)

var wizardCmd = &cobra.Command{
	Use:   "wizard",
	Short: "Start an interactive Flutter Cleaner flow",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(ui.TitleStyle.Render("🧙 Flutter Cleaner Wizard"))
		fmt.Println(ui.MutedText.Render("Navigate the tool interactively."))

		for {
			choice, _ := ui.AskSelect("What do you want to do?", []string{
				"Scan for projects",
				"Suggest what to clean",
				"Clean all projects",
				"Show stats",
				"Exit",
			})

			switch choice {
			case "Exit":
				fmt.Println("👋 Goodbye!")
				return

			case "Scan for projects":
				runCmd("scan", "~")

			case "Suggest what to clean":
				runCmd("suggest")

			case "Clean all projects":
				runCmd("clean", "--all")

			case "Show stats":
				runCmd("stats")
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(wizardCmd)
}

// Helper to run self with subcommands
func runCmd(args ...string) {
	self := os.Args[0]
	c := exec.Command(self, args...)
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	_ = c.Run()
}