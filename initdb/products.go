package initdb

import (
	"context"
	"os"

	"github.com/AgricolaDevJP/agricoladb-server/ent"

	"github.com/jszwec/csvutil"
)

const (
	ProductsFileName = "agricoladb-master - products.csv"
	ProductsCapacity = 64
)

type Product struct {
	ID            int    `csv:"id"`
	IsOfficialJa  bool   `csv:"is_official_ja"`
	NameJa        string `csv:"name_ja"`
	NameEn        string `csv:"name_en"`
	PublishedYear *int   `csv:"published_year"`
	RevisionID    int    `csv:"revision_id"`
}

func initProducts(ctx context.Context, tx *ent.Tx, csvFilePath string) error {
	bytes, err := os.ReadFile(csvFilePath)
	if err != nil {
		return err
	}

	products := []*Product{}
	if err := csvutil.Unmarshal(bytes, &products); err != nil {
		return err
	}

	builders := make([](*ent.ProductCreate), 0, ProductsCapacity)
	for _, product := range products {
		builder := tx.Product.Create().SetID(product.ID).SetIsOfficialJa(product.IsOfficialJa).SetNameJa(product.NameJa).SetNameEn(product.NameEn).SetNillablePublishedYear(product.PublishedYear).SetRevisionID(product.RevisionID)
		builders = append(builders, builder)
	}

	return tx.Product.CreateBulk(builders...).OnConflict().UpdateNewValues().Exec(ctx)
}
