package initdata

import (
	"agricoladb/ent"
	"context"
	"os"

	"github.com/jszwec/csvutil"
)

const (
	DecksFileName = "agricoladb-master - decks.csv"
	DecksCapacity = 64
)

type Deck struct {
	ID         int    `csv:"id"`
	Key        string `csv:"key"`
	NameJa     string `csv:"name_ja"`
	NameEn     string `csv:"name_en"`
	RevisionID int    `csv:"revision_id"`
}

func initDecks(ctx context.Context, tx *ent.Tx, csvFilePath string) error {
	bytes, err := os.ReadFile(csvFilePath)
	if err != nil {
		return err
	}

	decks := []*Deck{}
	if err := csvutil.Unmarshal(bytes, &decks); err != nil {
		return err
	}

	builders := make([](*ent.DeckCreate), 0, DecksCapacity)
	for _, deck := range decks {
		builder := tx.Deck.Create().SetID(deck.ID).SetKey(deck.Key).SetNameJa(deck.NameJa).SetNameEn(deck.NameEn).SetRevisionID(deck.RevisionID)
		builders = append(builders, builder)
	}

	return tx.Deck.CreateBulk(builders...).OnConflict().UpdateNewValues().Exec(ctx)
}
