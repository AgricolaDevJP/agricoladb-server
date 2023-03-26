package initdb

import (
	"context"
	_ "embed"

	"github.com/AgricolaDevJP/agricoladb-server/ent"

	"github.com/jszwec/csvutil"
)

//go:embed "masterdata/agricoladb-master - card_special_colors.csv"
var cardSpecialColorsBytes []byte

const (
	CardSpecialColorsCapacity = 8
)

type CardSpecialColor struct {
	ID     int    `csv:"id"`
	Key    string `csv:"key"`
	NameJa string `csv:"name_ja"`
	NameEn string `csv:"name_en"`
}

func initCardSpecialColors(ctx context.Context, tx *ent.Tx) error {
	cardSpecialColors := []*CardSpecialColor{}
	if err := csvutil.Unmarshal(cardSpecialColorsBytes, &cardSpecialColors); err != nil {
		return err
	}

	builders := make([](*ent.CardSpecialColorCreate), 0, CardSpecialColorsCapacity)
	for _, CardSpecialColor := range cardSpecialColors {
		builder := tx.CardSpecialColor.Create().SetID(CardSpecialColor.ID).SetKey(CardSpecialColor.Key).SetNameJa(CardSpecialColor.NameJa).SetNameEn(CardSpecialColor.NameEn)
		builders = append(builders, builder)
	}

	return tx.CardSpecialColor.CreateBulk(builders...).OnConflict().UpdateNewValues().Exec(ctx)
}
