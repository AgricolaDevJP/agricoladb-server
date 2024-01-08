package graph

import (
	"context"
	"fmt"

	"github.com/AgricolaDevJP/agricoladb-server/ent/card"
	"github.com/AgricolaDevJP/agricoladb-server/ent/cardspecialcolor"
	"github.com/AgricolaDevJP/agricoladb-server/ent/cardtype"
	"github.com/AgricolaDevJP/agricoladb-server/ent/deck"
	"github.com/AgricolaDevJP/agricoladb-server/ent/product"
	"github.com/AgricolaDevJP/agricoladb-server/ent/revision"
)

func getTypeFromID(ctx context.Context, id int) (string, error) {
	switch {
	case id >= 0 && id < 10:
		return revision.Table, nil
	case id >= 100 && id < 200:
		return cardtype.Table, nil
	case id >= 200 && id < 299:
		return cardspecialcolor.Table, nil
	case id >= 10000 && id < 20000:
		return product.Table, nil
	case id >= 20000 && id < 30000:
		return deck.Table, nil
	case id >= 100000 && id < 1000000:
		return card.Table, nil
	default:
		return "", fmt.Errorf("%d is out of range", id)
	}
}
