package initdb

import (
	"context"
	"os"

	"github.com/AgricolaDevJP/agricoladb-server/ent"

	"github.com/jszwec/csvutil"
)

const (
	CardsFileName = "agricoladb-master - cards.csv"
	CardsCapacity = 512
)

type Card struct {
	ID                              int     `csv:"id"`
	LiteralID                       string  `csv:"literal_id"`
	PrintedID                       string  `csv:"printed_id"`
	PlayAgricolaCardID              *string `csv:"play_agricola_card_id"`
	NameJa                          string  `csv:"name_ja"`
	NameEn                          string  `csv:"name_en"`
	MinPlayersNumber                *int    `csv:"min_players_number"`
	Prerequisite                    *string `csv:"prerequisite"`
	Cost                            *string `csv:"cost"`
	Description                     *string `csv:"description"`
	Note                            *string `csv:"note"`
	IsOfficialJa                    bool    `csv:"is_official_ja"`
	VictoryPoint                    *int    `csv:"victory_point"`
	SpecialVictoryPoint             *string `csv:"special_victory_point"`
	HasArrow                        bool    `csv:"has_arrow"`
	HasBonusPointIcon               bool    `csv:"has_bonus_point_icon"`
	HasNegativeBonusPointIcon       bool    `csv:"has_negative_bonus_point_icon"`
	HasPanIcon                      bool    `csv:"has_pan_icon"`
	HasBreadIcon                    bool    `csv:"has_bread_icon"`
	HasFarmPlannerIcon              bool    `csv:"has_farm_planner_icon"`
	HasActionsBoosterIcon           bool    `csv:"has_actions_booster_icon"`
	HasPointsProviderIcon           bool    `csv:"has_points_provider_icon"`
	HasGoodsProviderIcon            bool    `csv:"has_goods_provider_icon"`
	HasFoodProviderIcon             bool    `csv:"has_food_provider_icon"`
	HasCropProviderIcon             bool    `csv:"has_crop_provider_icon"`
	HasBuildingResourceProviderIcon bool    `csv:"has_building_resource_provider_icon"`
	HasLivestockProviderIcon        bool    `csv:"has_livestock_provider_icon"`
	HasCutPeatIcon                  bool    `csv:"has_cut_peat_icon"`
	HasFellTreesIcon                bool    `csv:"has_fell_trees_icon"`
	HasSlashAndBurnIcon             bool    `csv:"has_slash_and_burn_icon"`
	HasHiringFareIcon               bool    `csv:"has_hiring_fare_icon"`
	CardSpecialColorID              *int    `csv:"card_special_color_id"`
	CardTypeID                      int     `csv:"card_type_id"`
	DeckID                          *int    `csv:"deck_id"`
	RevisionID                      int     `csv:"revision_id"`
}

func initCards(ctx context.Context, tx *ent.Tx, csvFilePath string) error {
	bytes, err := os.ReadFile(csvFilePath)
	if err != nil {
		return err
	}

	cards := []*Card{}
	if err := csvutil.Unmarshal(bytes, &cards); err != nil {
		return err
	}

	builders := make([](*ent.CardCreate), 0, CardsCapacity)
	i := 0
	for _, card := range cards {
		// flush chunk
		if i >= 500 {
			if err := tx.Card.CreateBulk(builders...).OnConflict().UpdateNewValues().Exec(ctx); err != nil {
				return err
			}
			builders = make([](*ent.CardCreate), 0, CardsCapacity)
			i = 0
		}

		builder := tx.Card.Create().SetID(card.ID).SetLiteralID(card.LiteralID).SetPrintedID(card.PrintedID).SetNillablePlayAgricolaCardID(card.PlayAgricolaCardID).SetNameJa(card.NameJa).
			SetNillableMinPlayersNumber(card.MinPlayersNumber).SetNillablePrerequisite(card.Prerequisite).SetNillableCost(card.Cost).SetNillableDescription(card.Description).
			SetNameEn(card.NameEn).SetNillableNote(card.Note).SetIsOfficialJa(card.IsOfficialJa).SetNillableVictoryPoint(card.VictoryPoint).SetNillableSpecialVictoryPoint(card.SpecialVictoryPoint).
			SetHasArrow(card.HasArrow).SetHasBonusPointIcon(card.HasBonusPointIcon).SetHasNegativeBonusPointIcon(card.HasNegativeBonusPointIcon).SetHasPanIcon(card.HasPanIcon).
			SetHasBreadIcon(card.HasBreadIcon).SetHasFarmPlannerIcon(card.HasFarmPlannerIcon).SetHasActionsBoosterIcon(card.HasActionsBoosterIcon).
			SetHasPointsProviderIcon(card.HasPointsProviderIcon).SetHasGoodsProviderIcon(card.HasGoodsProviderIcon).SetHasFoodProviderIcon(card.HasFoodProviderIcon).
			SetHasCropProviderIcon(card.HasCropProviderIcon).SetHasBuildingResourceProviderIcon(card.HasBuildingResourceProviderIcon).
			SetHasLivestockProviderIcon(card.HasLivestockProviderIcon).SetHasCutPeatIcon(card.HasCutPeatIcon).SetHasFellTreesIcon(card.HasFellTreesIcon).
			SetHasSlashAndBurnIcon(card.HasSlashAndBurnIcon).SetHasHiringFareIcon(card.HasHiringFareIcon).SetNillableCardSpecialColorID(card.CardSpecialColorID).
			SetCardTypeID(card.CardTypeID).SetNillableDeckID(card.DeckID).SetRevisionID(card.RevisionID)
		builders = append(builders, builder)
		i++
	}

	return tx.Card.CreateBulk(builders...).OnConflict().UpdateNewValues().Exec(ctx)
}
