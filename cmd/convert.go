package cmd

import (
	"encoding/csv"
	"fmt"
	"github.com/spf13/cobra"
	"io"
	"log"
	"os"
	"strings"
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
	filePath := &args[0]

	file, err := os.Open(*filePath)
	if err != nil {
		fmt.Println("Pfad kann nicht gefunden werden!")
		return
	}

	defer func(file *os.File) {
		err = file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)

	csvReader := csv.NewReader(file)
	csvReader.Comma = ';'

	conversionFileName := strings.Split(*filePath, ".")
	conversionFile, err := os.Create(conversionFileName[0] + "_CONVERTED" + ".csv")
	if err != nil {
		fmt.Println(err)
	}

	defer func(conversionFile *os.File) {
		err := conversionFile.Close()
		if err != nil {

		}
	}(conversionFile)

	var convertedDate string

	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}

		convertedDate = parseDateISO(&record[0])
		record[0] = convertedDate
		_, err = conversionFile.WriteString(rebuildRecord(&record))
		if err != nil {
			fmt.Println(err)
		}
	}
}

func parseDateISO(oldDate *string) string {
	split := strings.Split(*oldDate, "-")
	yearMonthDay := strings.ReplaceAll(split[0], ".", "-")

	builder := strings.Builder{}

	builder.WriteString(yearMonthDay)
	builder.WriteString("T")
	builder.WriteString(split[1])

	return builder.String()
}

func rebuildRecord(lines *[]string) string {
	builder := strings.Builder{}

	for i, line := range *lines {
		builder.WriteString(line)

		if i < len(*lines)-1 {
			builder.WriteString(";")
		}
	}

	builder.WriteString("\n")

	return builder.String()
}
