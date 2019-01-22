package cmd

import (
	"github.com/spf13/cobra"
)

// RootCmd first Funtion Even run before Main its start excute form init func
var RootCmd = &cobra.Command{
	Use:   "task",
	Short: "Task is a CLI task manager",
	Long:  "CLI Task Manager, For Managing Task List Performing list of functionlity Add,Delete...etc",
}
