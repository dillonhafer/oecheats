package cmd

import (
	"github.com/spf13/cobra"
)

// formatCmd represents the format command
var formatCmd = &cobra.Command{
	Use:   "format",
	Short: "Format the cheats database",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		FormatDatabase()
	},
}

func init() {
	rootCmd.AddCommand(formatCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// formatCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// formatCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
