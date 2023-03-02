package initdb

import (
	"context"
	_ "embed"

	"github.com/AgricolaDevJP/agricoladb-server/ent"

	"github.com/jszwec/csvutil"
)

//go:embed "masterdata/agricoladb-master - decks.csv"
var decksBytes []byte

const (
	DecksCapacity = 64
)

type Deck struct {
	ID         int    `csv:"id"`
	Key        string `csv:"key"`
	NameJa     string `csv:"name_ja"`
	NameEn     string `csv:"name_en"`
	RevisionID int    `csv:"revision_id"`
}

func initDecks(ctx context.Context, tx *ent.Tx) error {
	decks := []*Deck{}
	if err := csvutil.Unmarshal(decksBytes, &decks); err != nil {
		return err
	}

	builders := make([](*ent.DeckCreate), 0, DecksCapacity)
	for _, deck := range decks {
		builder := tx.Deck.Create().SetID(deck.ID).SetKey(deck.Key).SetNameJa(deck.NameJa).SetNameEn(deck.NameEn).SetRevisionID(deck.RevisionID)
		builders = append(builders, builder)
	}

	return tx.Deck.CreateBulk(builders...).OnConflict().UpdateNewValues().Exec(ctx)
}
