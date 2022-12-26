package initdb

import (
	"agricoladb/ent"
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/jszwec/csvutil"
)

const (
	ProductCardsFileName             = "agricoladb-master - product_cards.csv"
	ProductCardsPlaceholdersCapacity = 512
	ProductCardsValuesCapacity       = ProductCardsPlaceholdersCapacity * 2
)

type ProductCards struct {
	ProductID int `csv:"product_id"`
	CardID    int `csv:"card_id"`
}

func initProductCards(ctx context.Context, tx *ent.Tx, csvFilePath string) error {
	bytes, err := os.ReadFile(csvFilePath)
	if err != nil {
		return err
	}

	productCards := []*ProductCards{}
	if err := csvutil.Unmarshal(bytes, &productCards); err != nil {
		return err
	}

	placeholders := make([]string, 0, ProductCardsPlaceholdersCapacity)
	values := make([]any, 0, ProductCardsValuesCapacity)
	i := 0
	for _, productCard := range productCards {
		// flush chunk
		if i >= 500 {
			placeholdersString := strings.Join(placeholders, ",")
			query := fmt.Sprintf("INSERT INTO `product_cards` (`product_id`, `card_id`) VALUES %s ON DUPLICATE KEY UPDATE `product_id` = VALUES(`product_id`), `card_id` = VALUES(`card_id`)", placeholdersString)
			_, err = tx.ExecContext(ctx, query, values...)
			if err != nil {
				return err
			}

			placeholders = make([]string, 0, ProductCardsPlaceholdersCapacity)
			values = make([]any, 0, ProductCardsValuesCapacity)
			i = 0
		}
		placeholders = append(placeholders, "(?, ?)")
		values = append(values, productCard.ProductID, productCard.CardID)
		i++
	}

	if len(values) > 0 {
		placeholdersString := strings.Join(placeholders, ",")
		query := fmt.Sprintf("INSERT INTO `product_cards` (`product_id`, `card_id`) VALUES %s ON DUPLICATE KEY UPDATE `product_id` = VALUES(`product_id`), `card_id` = VALUES(`card_id`)", placeholdersString)
		_, err = tx.ExecContext(ctx, query, values...)
		if err != nil {
			return err
		}
	}

	return nil
}
