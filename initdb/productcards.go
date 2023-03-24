package initdb

import (
	"context"
	_ "embed"
	"fmt"
	"strings"

	"github.com/AgricolaDevJP/agricoladb-server/ent"

	"github.com/jszwec/csvutil"
)

//go:embed "masterdata/agricoladb-master - product_cards.csv"
var productCardsBytes []byte

const (
	ProductCardsPlaceholdersCapacity = 512
	ProductCardsValuesCapacity       = ProductCardsPlaceholdersCapacity * 2
)

type ProductCards struct {
	ProductID int `csv:"product_id"`
	CardID    int `csv:"card_id"`
}

func initProductCards(ctx context.Context, tx *ent.Tx) error {
	productCards := []*ProductCards{}
	if err := csvutil.Unmarshal(productCardsBytes, &productCards); err != nil {
		return err
	}

	placeholders := make([]string, 0, ProductCardsPlaceholdersCapacity)
	values := make([]any, 0, ProductCardsValuesCapacity)
	i := 0
	for _, productCard := range productCards {
		// flush chunk
		if i >= 500 {
			placeholdersString := strings.Join(placeholders, ",")
			query := fmt.Sprintf("INSERT INTO `product_cards` (`product_id`, `card_id`) VALUES %s", placeholdersString)
			if _, err := tx.ExecContext(ctx, query, values...); err != nil {
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
		query := fmt.Sprintf("INSERT INTO `product_cards` (`product_id`, `card_id`) VALUES %s", placeholdersString)
		if _, err := tx.ExecContext(ctx, query, values...); err != nil {
			return err
		}
	}

	return nil
}
