package cmd

import (
	"encoding/csv"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var description = `Konvertiert eine oder mehrere Dateien bei angegebenem Pfad.
Konversion einzelner Datei:
DateTimeConverterGO convert D:\Verzeichnis\test.csv

Konversion aller Dateien eines Ordners:
DateTimeConverterGO convert D:\Verzeichnis\`

var convertCmd = &cobra.Command{
	Use:   "convert",
	Short: "Konvertiert Dateien bei angegebenem Pfad",
	Long:  description,
	Args:  cobra.MinimumNArgs(1),
	Run:   handleConversion,
}

func handleConversion(cmd *cobra.Command, args []string) {
	fmt.Println("Arguments: " + args[0])

	file, err := os.Open(args[0])
	if err != nil {
		fmt.Println("Pfad kann nicht gefunden werden!")
		return
	}

	defer file.Close()

	csvReader := csv.NewReader(file)
	data, err := csvReader.ReadAll()

	fmt.Println(data)
}
