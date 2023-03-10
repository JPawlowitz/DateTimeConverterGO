package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "DateTimeConverterGO",
	Short: "Konvertiert Datumsformate nach ISO 8601 Standard",
	Long: `_____DateTimeConverter_____
Konvertiert Datumsformate nach ISO 8601 Standard`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	convertCmd.Flags().BoolVar(&convertDir, "dir", false, "Konvertiert gesamtes Verzeichnis")

	rootCmd.AddCommand(convertCmd)
}
