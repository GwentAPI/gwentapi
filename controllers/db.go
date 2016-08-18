package controllers

import (
	"database/sql"
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"github.com/tri125/gwentapi/app"
)

var DBCon *sql.DB

type PageQueryType int

const (
	AllCards        PageQueryType = 1 << iota
	RarityFiltered  PageQueryType = 1 << iota
	LeaderFiltered  PageQueryType = 1 << iota
	FactionFiltered PageQueryType = 1 << iota
)

func NewDBConnection(DSN string) (*sql.DB, error) {
	var dbError error
	var db *sql.DB

	db, dbError = sql.Open("mysql", DSN)
	return db, dbError
}

func FetchType(id string) (TypeModel, error) {
	var myType TypeModel
	row := DBCon.QueryRow("SELECT id, name FROM Types WHERE id=?", id)
	err := row.Scan(&myType.ID, &myType.Name)
	if err != nil {
		myType.ID = ""
		myType.Name = ""
	}

	return myType, err
}

func FetchAllTypes() ([]*TypeModel, error) {
	var cardTypes []*TypeModel

	rows, err := DBCon.Query("SELECT id, name FROM Types ORDER BY id ASC")

	if err != nil {
		return cardTypes, err
	}
	defer rows.Close()

	for rows.Next() {
		var typeModel TypeModel
		err := rows.Scan(&typeModel.ID, &typeModel.Name)

		if err != nil {
			continue
		}
		cardTypes = append(cardTypes, &typeModel)
	}

	return cardTypes, err
}

func FetchRarity(id string) (RarityModel, error) {
	var rarity RarityModel
	row := DBCon.QueryRow("SELECT id, name FROM Rarities WHERE id=?", id)
	err := row.Scan(&rarity.ID, &rarity.Name)
	if err != nil {
		rarity.ID = ""
		rarity.Name = ""
	}

	return rarity, err
}

func FetchAllRarities() ([]*RarityModel, error) {
	var rarities []*RarityModel
	rows, err := DBCon.Query("SELECT id, name FROM Rarities ORDER BY id ASC")

	if err != nil {
		return rarities, err
	}
	defer rows.Close()

	for rows.Next() {
		var rarity RarityModel
		err := rows.Scan(&rarity.ID, &rarity.Name)

		if err != nil {
			continue
		}
		rarities = append(rarities, &rarity)
	}

	return rarities, err
}

func FetchPatch(id string) (PatchModel, error) {
	var patch PatchModel
	var changeLog sql.NullString
	var emptyString = ""

	row := DBCon.QueryRow("SELECT id, version, releaseDate, changelog FROM Patches WHERE id=?", id)
	err := row.Scan(&patch.ID, &patch.Version, &patch.ReleaseDate, &changeLog)
	if err != nil {
		return patch, err
	}

	if changeLog.Valid {
		patch.Changelog = &changeLog.String
	} else {
		patch.Changelog = &emptyString
	}

	return patch, err
}

func FetchLatestPatch() (PatchModel, error) {
	var patch PatchModel
	var changeLog sql.NullString
	var emptyString = ""

	row := DBCon.QueryRow("SELECT id, version, releaseDate, changelog FROM Patches WHERE releaseDate = (SELECT MAX(releaseDate) FROM Patches)")
	err := row.Scan(&patch.ID, &patch.Version, &patch.ReleaseDate, &changeLog)
	if err != nil {
		return patch, err
	}

	if changeLog.Valid {
		patch.Changelog = &changeLog.String
	} else {
		patch.Changelog = &emptyString
	}

	return patch, err
}

func FetchAllPatches() ([]*PatchModel, error) {
	var patches []*PatchModel
	var emptyString = ""

	rows, err := DBCon.Query("SELECT id, version, releaseDate, changelog FROM Patches ORDER BY releaseDate ASC")

	if err != nil {
		return patches, err
	}
	defer rows.Close()

	for rows.Next() {
		var patch PatchModel
		var changeLog sql.NullString

		err := rows.Scan(&patch.ID, &patch.Version, &patch.ReleaseDate, &changeLog)

		if err != nil {
			continue
		}

		if changeLog.Valid {
			patch.Changelog = &changeLog.String
		} else {
			patch.Changelog = &emptyString
		}

		patches = append(patches, &patch)
	}

	return patches, err
}

func FetchGlyph(id string) (GlyphModel, error) {
	var glyph GlyphModel
	row := DBCon.QueryRow("SELECT id, name, isWeatherGlyph, text FROM Glyphs WHERE id=?", id)
	err := row.Scan(&glyph.ID, &glyph.Name, &glyph.IsWeatherGlyph, &glyph.Text)
	if err != nil {
		glyph.ID = ""
		glyph.Name = ""
		glyph.IsWeatherGlyph = false
		glyph.Text = ""
	}

	return glyph, err
}

func FetchAllGlyphs() ([]*GlyphModel, error) {
	var glyphs []*GlyphModel

	rows, err := DBCon.Query("SELECT id, name, isWeatherGlyph, text FROM Glyphs ORDER BY id ASC")

	if err != nil {
		return glyphs, err
	}
	defer rows.Close()

	for rows.Next() {
		var glyph GlyphModel

		err := rows.Scan(&glyph.ID, &glyph.Name, &glyph.IsWeatherGlyph, &glyph.Text)

		if err != nil {
			continue
		}

		glyphs = append(glyphs, &glyph)
	}

	return glyphs, err
}

func FetchFaction(id string) (FactionModel, error) {
	var faction FactionModel
	row := DBCon.QueryRow("SELECT id, name FROM Factions WHERE id=?", id)
	err := row.Scan(&faction.ID, &faction.Name)
	if err != nil {
		faction.ID = ""
		faction.Name = ""
	}

	return faction, err
}

func FetchAllFactions() ([]*FactionModel, error) {
	var factions []*FactionModel
	rows, err := DBCon.Query("SELECT id, name FROM Factions ORDER BY id ASC")

	if err != nil {
		return factions, err
	}
	defer rows.Close()

	for rows.Next() {
		var faction FactionModel
		err := rows.Scan(&faction.ID, &faction.Name)
		if err != nil {
			continue
		}

		factions = append(factions, &faction)
	}
	return factions, err
}

func FetchIllustrator(id string) (IllustratorModel, error) {
	var illustrator IllustratorModel
	row := DBCon.QueryRow("SELECT id, name FROM Illustrators WHERE id=?", id)
	err := row.Scan(&illustrator.ID, &illustrator.Name)
	if err != nil {
		illustrator.ID = ""
		illustrator.Name = ""
	}

	return illustrator, err
}

func FetchAllIllustrators() ([]*IllustratorModel, error) {
	var illustrators []*IllustratorModel
	rows, err := DBCon.Query("SELECT id, name FROM Illustrators ORDER BY id ASC")

	if err != nil {
		return illustrators, err
	}
	defer rows.Close()

	for rows.Next() {
		var illustrator IllustratorModel

		err := rows.Scan(&illustrator.ID, &illustrator.Name)
		if err != nil {
			continue
		}
		illustrators = append(illustrators, &illustrator)
	}
	return illustrators, err
}

func FetchCard(id string) (CardModel, error) {
	var card CardModel
	var rarity RarityModel
	var faction FactionModel
	var cardType TypeModel

	var flavor, text sql.NullString
	var strength sql.NullInt64

	row := DBCon.QueryRow("SELECT c.name, c.id, r.id, r.name, f.id, f.name, t.id, t.name, strength, text, flavor FROM Cards AS c INNER JOIN Rarities AS r ON c.idRarity = r.idRarity INNER JOIN Factions AS f ON c.idFaction = f.idFaction INNER JOIN Types AS t ON c.idType = t.idType WHERE c.id =?", id)

	err := row.Scan(&card.Name, &card.ID, &rarity.ID, &rarity.Name, &faction.ID, &faction.Name, &cardType.ID, &cardType.Name, &strength, &text, &flavor)

	if err != nil {
		card.Faction.ID = ""
		card.Faction.Name = ""
		card.Flavor = nil
		card.ID = ""
		card.Name = ""
		card.Rarity.ID = ""
		card.Rarity.Name = ""
		card.Rows = nil
		card.Strength = nil
		card.Subtypes = nil
		card.Text = nil
		card.TypeCard.ID = ""
		card.TypeCard.Name = ""
	} else {
		card.Faction = faction
		card.Rarity = rarity
		card.TypeCard = cardType

		cardRows, _ := fetchCardRows(id)
		cardSubtypes, _ := fetchCardSubTypes(id)

		card.Rows = cardRows
		card.Subtypes = cardSubtypes

		if strength.Valid {
			var converted = int(strength.Int64)
			card.Strength = &converted
		} else {
			card.Strength = nil
		}

		if text.Valid {
			card.Text = &text.String
		} else {
			card.Text = nil
		}

		if flavor.Valid {
			card.Flavor = &flavor.String
		} else {
			card.Flavor = nil
		}
	}

	return card, err
}

func CountCards(pq PageQueryType, param string) (int, error) {
	var count int
	var err error
	var row *sql.Row

	switch {
	case pq == AllCards:
		row = DBCon.QueryRow("SELECT COUNT(*) FROM Cards")
	case pq == RarityFiltered:
		row = DBCon.QueryRow("SELECT COUNT(*) FROM Cards AS c INNER JOIN Rarities AS r ON c.idRarity = r.idRarity WHERE r.id = ?", param)
	case pq == LeaderFiltered:
		row = DBCon.QueryRow("SELECT COUNT(*) FROM Cards AS c INNER JOIN Types AS t ON c.idType = t.idType WHERE t.name = ?", "Leader")
	case pq == FactionFiltered:
		row = DBCon.QueryRow("SELECT COUNT(*) FROM Cards AS c INNER JOIN Factions AS f ON c.idFaction = f.idFaction WHERE f.id = ?", param)
	default:
		return 0, errors.New("Invalid PageQueryType")
	}
	err = row.Scan(&count)

	return count, err
}

func FetchPageCards(pq PageQueryType, limit int, offset int, id string) ([]*CardModel, error) {
	var cards []*CardModel
	var err error
	var rows *sql.Rows

	switch {
	case pq == AllCards:
		rows, err = DBCon.Query("SELECT c.name, c.id, r.id, r.name, f.id, f.name, t.id, t.name, strength, text, flavor FROM Cards AS c INNER JOIN Rarities AS r ON c.idRarity = r.idRarity INNER JOIN Factions AS f ON c.idFaction = f.idFaction INNER JOIN Types AS t ON c.idType = t.idType ORDER BY c.id ASC LIMIT ? OFFSET ?", limit, offset)
	case pq == RarityFiltered:
		rows, err = DBCon.Query("SELECT c.name, c.id, r.id, r.name, f.id, f.name, t.id, t.name, strength, text, flavor FROM Cards AS c INNER JOIN Rarities AS r ON c.idRarity = r.idRarity INNER JOIN Factions AS f ON c.idFaction = f.idFaction INNER JOIN Types AS t ON c.idType = t.idType WHERE r.id = ? ORDER BY c.id ASC LIMIT ? OFFSET ?", id, limit, offset)
	case pq == LeaderFiltered:
		rows, err = DBCon.Query("SELECT c.name, c.id, r.id, r.name, f.id, f.name, t.id, t.name, strength, text, flavor FROM Cards AS c INNER JOIN Rarities AS r ON c.idRarity = r.idRarity INNER JOIN Factions AS f ON c.idFaction = f.idFaction INNER JOIN Types AS t ON c.idType = t.idType WHERE t.name = ? ORDER BY c.id ASC LIMIT ? OFFSET ?", "Leader", limit, offset)
	case pq == FactionFiltered:
		rows, err = DBCon.Query("SELECT c.name, c.id, r.id, r.name, f.id, f.name, t.id, t.name, strength, text, flavor FROM Cards AS c INNER JOIN Rarities AS r ON c.idRarity = r.idRarity INNER JOIN Factions AS f ON c.idFaction = f.idFaction INNER JOIN Types AS t ON c.idType = t.idType WHERE f.id = ? ORDER BY c.id ASC LIMIT ? OFFSET ?", id, limit, offset)
	default:
		return nil, errors.New("Invalid PageQueryType")
	}

	if err != nil {
		return cards, err
	}
	defer rows.Close()

	for rows.Next() {
		var card CardModel
		var rarity RarityModel
		var faction FactionModel
		var cardType TypeModel

		var flavor, text sql.NullString
		var strength sql.NullInt64

		err := rows.Scan(&card.Name, &card.ID, &rarity.ID, &rarity.Name, &faction.ID, &faction.Name, &cardType.ID, &cardType.Name, &strength, &text, &flavor)
		if err != nil {
			continue
		}

		card.Faction = faction
		card.Rarity = rarity
		card.TypeCard = cardType

		cardRows, _ := fetchCardRows(card.ID)
		cardSubtypes, _ := fetchCardSubTypes(card.ID)

		card.Rows = cardRows
		card.Subtypes = cardSubtypes

		if strength.Valid {
			var converted = int(strength.Int64)
			card.Strength = &converted
		} else {
			card.Strength = nil
		}

		if text.Valid {
			card.Text = &text.String
		} else {
			card.Text = nil
		}

		if flavor.Valid {
			card.Flavor = &flavor.String
		} else {
			card.Flavor = nil
		}
		cards = append(cards, &card)
	}
	return cards, err
}

func fetchCardRows(id string) ([]string, error) {
	var cardRows []string

	rows, err := DBCon.Query("SELECT r.name FROM Rows AS r INNER JOIN CardsRows AS cr ON r.idRow = cr.idRow INNER JOIN Cards AS c ON c.idCard = cr.idCard WHERE c.id =?", id)

	if err != nil {
		return cardRows, err
	}
	defer rows.Close()

	for rows.Next() {
		var row string
		err := rows.Scan(&row)

		if err != nil {
			continue
		}
		cardRows = append(cardRows, row)
	}

	return cardRows, err
}

func fetchCardSubTypes(id string) ([]*TypeModel, error) {
	var cardSubtypes []*TypeModel

	rows, err := DBCon.Query("SELECT t.id, t.name FROM Types AS t INNER JOIN CardsSubtypes AS st ON t.idType = st.idType INNER JOIN Cards AS c ON c.idCard = st.idCard WHERE c.id =?", id)

	if err != nil {
		return cardSubtypes, err
	}
	defer rows.Close()

	for rows.Next() {
		var typeModel TypeModel
		err := rows.Scan(&typeModel.ID, &typeModel.Name)

		if err != nil {
			continue
		}
		cardSubtypes = append(cardSubtypes, &typeModel)
	}

	return cardSubtypes, err
}

func FetchArtwork(id string) (ArtworkMediaModel, error) {
	var artworkMedia ArtworkMediaModel

	rows, err := DBCon.Query("SELECT c.name AS card, c.id, arti.name AS artist, arti.id, art.isAlternative, art.filename, cat.name AS category FROM Artworks AS art INNER JOIN Cards AS c ON c.idCard = art.idCard INNER JOIN Categories AS cat ON cat.idCategory = art.idCategory LEFT JOIN Artists AS arti ON arti.idArtist = art.idArtist WHERE c.id =? ORDER BY isAlternative, filename ASC", id)

	if err != nil {
		return artworkMedia, err
	}
	defer rows.Close()

	artworkMedia.ID = id
	var count int = 0
	for rows.Next() {
		count++
		//var artist IllustratorModel
		var artworkType app.ArtworkType
		var isAlternative bool
		var filename, category, cardName, cardID string

		var artistName, artistID sql.NullString

		err := rows.Scan(&cardName, &cardID, &artistName, &artistID, &isAlternative, &filename, &category)
		if err != nil {
			continue
		}

		if artistName.Valid {
			artworkType.Artist = &artistName.String
		}

		if category == "normal" {
			artworkType.NormalSize = filename
		}

		//if artistID.Valid {
		//	artist.ID = &artistID.String
		//}
		if isAlternative {
			artworkMedia.Alternatives = append(artworkMedia.Alternatives, &artworkType)
		} else {
			artworkMedia.Artwork = &artworkType
		}
	}

	if count == 0 {
		return artworkMedia, errors.New("No rows on Artworks")
	}

	return artworkMedia, nil
}
