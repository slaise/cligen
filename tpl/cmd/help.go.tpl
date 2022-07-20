package cmd

import (
  "fmt"

  "github.com/spf13/cobra"
)

func init() {
  rootCmd.AddCommand(helpCmd)
  // TODO read from a file
  rootCmd.SetHelpTemplate({{ .help }})
}

var helpCmd = &cobra.Command{
  Use:   "help",
  Short: "Print help info",
  Long:  "",
  Run: func(cmd *cobra.Command, args []string) {
    fmt.Println("Kubectl plugin {{ .name }} version {{ .version }} ")
  },
}
