package migrations

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"github.com/pkg/errors"
	"github.com/pressly/goose/v3"
	"log"
	"os"
	"strings"
)

func init() {
	goose.AddMigration(upFillMarketItem, downFillMarketItem)
}

const MarketItemFieldsNum = 3

type MarketItem struct {
	Symbol string
	Type   string
	Title  string
}

func readCsvFile(filePath string) [][]string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}

	return records
}

func upFillMarketItem(tx *sql.Tx) error {
	itemsStr := readCsvFile("./migrations/market_items.csv")
	items := make(map[string]MarketItem, len(itemsStr))

	for i, itemStr := range itemsStr {
		if i == 0 {
			continue
		}
		key := strings.Join([]string{itemStr[0], itemStr[2]}, "")

		//if item, ok := items[key]; ok {
		//	log.Println(items[key])
		//	log.Println(item)
		//}
		items[key] = MarketItem{Symbol: itemStr[0], Type: itemStr[2], Title: itemStr[1]}

	}

	var placeholders []string
	var values []interface{}

	for _, marketItem := range items {
		placeholders = append(
			placeholders,
			fmt.Sprintf("($%d,$%d,$%d)",
				len(placeholders)*MarketItemFieldsNum+1,
				len(placeholders)*MarketItemFieldsNum+2,
				len(placeholders)*MarketItemFieldsNum+3,
			),
		)
		values = append(values, marketItem.Symbol, marketItem.Type, marketItem.Title)
	}

	insertQuery := fmt.Sprintf(
		`INSERT INTO market_item (Code,Type,Title) VALUES %s`,
		strings.Join(placeholders, ","),
	)

	_, err := tx.Exec(insertQuery, values...)
	if err != nil {
		return errors.Wrap(err, "failed to insert multiple records at once")
	}

	return nil
}

func downFillMarketItem(tx *sql.Tx) error {
	itemsStr := readCsvFile("./migrations/market_items.csv")

	items := make(map[string][]string, 0)

	for i, itemStr := range itemsStr {
		if i == 0 {
			continue
		}
		if _, ok := items[itemStr[2]]; !ok {
			items[itemStr[2]] = make([]string, 0)
		}
		items[itemStr[2]] = append(items[itemStr[2]], itemStr[0])
	}

	for itemType, codes := range items {
		var placeholders []string
		var values []interface{}

		for _, code := range codes {
			placeholders = append(
				placeholders,
				fmt.Sprintf("$%d", len(placeholders)+1),
			)
			values = append(values, code)
		}
		values = append(values, itemType)

		deleteQuery := fmt.Sprintf(
			`DELETE FROM market_item WHERE Code IN (%s) AND Type = $%d`,
			strings.Join(placeholders, ","),
			len(placeholders)+1,
		)

		_, err := tx.Exec(deleteQuery, values...)
		if err != nil {
			return err
		}
	}

	return nil
}
