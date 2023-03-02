package initdb

import (
	"context"
	_ "embed"
	"fmt"
	"strings"

	"github.com/AgricolaDevJP/agricoladb-server/ent"

	"github.com/jszwec/csvutil"
)

//go:embed "masterdata/agricoladb-master - card_ancestors.csv"
var cardAncestorsBytes []byte

const (
	CardAncestorsPlaceholdersCapacity = 512
	CardAncestorsValuesCapacity       = CardAncestorsPlaceholdersCapacity * 2
)

type CardAncestors struct {
	CardID  int `csv:"card_id"`
	ChildID int `csv:"child_id"`
}

func initCardAncestors(ctx context.Context, tx *ent.Tx) error {
	cardAncestors := []*CardAncestors{}
	if err := csvutil.Unmarshal(cardAncestorsBytes, &cardAncestors); err != nil {
		return err
	}

	placeholders := make([]string, 0, CardAncestorsPlaceholdersCapacity)
	values := make([]any, 0, CardAncestorsValuesCapacity)
	i := 0
	for _, cardAncestor := range cardAncestors {
		// flush chunk
		if i >= 500 {
			placeholdersString := strings.Join(placeholders, ",")
			query := fmt.Sprintf("INSERT INTO `card_ancestors` (`card_id`, `child_id`) VALUES %s ON DUPLICATE KEY UPDATE `card_id` = VALUES(`card_id`), `child_id` = VALUES(`child_id`)", placeholdersString)
			if _, err := tx.ExecContext(ctx, query, values...); err != nil {
				return err
			}

			placeholders = make([]string, 0, CardAncestorsPlaceholdersCapacity)
			values = make([]any, 0, CardAncestorsValuesCapacity)
			i = 0
		}
		placeholders = append(placeholders, "(?, ?)")
		values = append(values, cardAncestor.CardID, cardAncestor.ChildID)
		i++
	}

	if len(values) > 0 {
		placeholdersString := strings.Join(placeholders, ",")
		query := fmt.Sprintf("INSERT INTO `card_ancestors` (`card_id`, `child_id`) VALUES %s ON DUPLICATE KEY UPDATE `card_id` = VALUES(`card_id`), `child_id` = VALUES(`child_id`)", placeholdersString)
		if _, err := tx.ExecContext(ctx, query, values...); err != nil {
			return err
		}
	}

	return nil
}
