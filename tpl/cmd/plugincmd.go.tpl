package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var name string

// {{ .name }}Cmd represents the kubectlgen command
var {{ .name }}Cmd = &cobra.Command{
	Use:   "{{ .name }}",
	Short: "A short description of the plugin",
	Long: "A longer description of the plugin",
	Run: func(cmd *cobra.Command, args []string) {
        // TODO write your logic here
	},
}

func init() {
	rootCmd.AddCommand({{ .name }}Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// {{ .name }}Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// {{ .name }}Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
