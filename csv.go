package playground

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"strings"
)

func readCsv() {
	lines := []string{
		"りんご,Apple,バラ科",
		"みかん,Orange,ミカン科",
		"すいか,Watermelon,ウリ科",
	}
	csvStr := strings.Join(lines, "\n")

	r := csv.NewReader(strings.NewReader(csvStr))
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Read error: ", err)
			break
		}

		fmt.Printf("日本語名[%s] 英語名[%s] 科分類[%s]\n", record[0], record[1], record[2])
	}
}

func writeCsv() {
	records := [][]string{
		[]string{"りんご", "Apple", "バラ科"},
		[]string{"みかん", "Orange", "ミカン科"},
		[]string{"すいか", "Watermelon", "ウリ科"},
	}

	buf := new(bytes.Buffer)
	w := csv.NewWriter(buf)
	for _, record := range records {
		if err := w.Write(record); err != nil {
			fmt.Println("Write error: ", err)
			return
		}
		w.Flush()
	}
	fmt.Println(buf.String())
}
