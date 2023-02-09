package cmd

import (
	"encoding/csv"
	"fmt"
	"github.com/spf13/cobra"
	"io"
	"os"
	"strings"
)

var description = `Konvertiert eine oder mehrere Dateien bei angegebenem Pfad.
Konversion einzelner Datei:
DateTimeConverterGO convert D:\Verzeichnis\test.csv

Konversion aller Dateien eines Ordners:
DateTimeConverterGO convert -dir D:\Verzeichnis\`

var convertDir bool

var convertCmd = &cobra.Command{
	Use:   "convert",
	Short: "Konvertiert Dateien bei angegebenem Pfad",
	Long:  description,
	Args:  cobra.MinimumNArgs(1),
	Run:   handleConversion,
}

func handleConversion(cmd *cobra.Command, args []string) {
	path := &args[0]

	if convertDir {
		handleConvertDirectory(path)
	} else {
		handleConvertFile(path)
	}

	fmt.Println("Konversion erfolgreich!")
}

func handleConvertDirectory(dirPath *string) {
	dir, err := os.Open(*dirPath)
	if err != nil {
		fmt.Println("Ordner kann nicht gefunden werden")
		return
	}

	defer dir.Close()

	dirFiles, err := dir.ReadDir(0)
	if err != nil {
		return
	}

	for _, file := range dirFiles {
		if !file.IsDir() {
			split := strings.Split(file.Name(), ".")

			if len(split) == 2 && split[1] == "csv" {
				filePath := *dirPath + file.Name()

				handleConvertFile(&filePath)
			}
		}
	}
}

func handleConvertFile(filePath *string) {
	file, err := os.Open(*filePath)
	if err != nil {
		fmt.Println("Datei kann nicht gefunden werden!")
		return
	}

	fmt.Println("Konvertiere Datei ", *filePath)

	defer file.Close()

	csvReader := csv.NewReader(file)
	csvReader.Comma = ';'

	conversionFileName := strings.Split(*filePath, ".")
	conversionFile, err := os.Create(conversionFileName[0] + "_CONVERTED" + ".csv")
	if err != nil {
		fmt.Println(err)
	}

	defer conversionFile.Close()

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
