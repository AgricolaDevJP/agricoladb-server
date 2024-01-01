package initdb

import (
	"context"
	_ "embed"

	"github.com/AgricolaDevJP/agricoladb-server/ent"

	"github.com/jszwec/csvutil"
)

//go:embed "masterdata/agricoladb-master - card_types.csv"
var cardTypesBytes []byte

const (
	CardTypesCapacity = 8
)

type CardType struct {
	ID     int    `csv:"id"`
	Key    string `csv:"key"`
	NameJa string `csv:"name_ja"`
	NameEn string `csv:"name_en"`
}

func initCardTypes(ctx context.Context, tx *ent.Tx) error {
	cardtypes := []*CardType{}
	if err := csvutil.Unmarshal(cardTypesBytes, &cardtypes); err != nil {
		return err
	}

	builders := make([](*ent.CardTypeCreate), 0, CardTypesCapacity)
	for _, cardType := range cardtypes {
		builder := tx.CardType.Create().SetID(cardType.ID).SetKey(cardType.Key).SetNameJa(cardType.NameJa).SetNameEn(cardType.NameEn)
		builders = append(builders, builder)
	}

	return tx.CardType.CreateBulk(builders...).OnConflict().UpdateNewValues().Exec(ctx)
}
