package main

type TileMap struct {
	Tileshigh int `json:"tileshigh"`
	Layers    []struct {
		Tiles []struct {
			X     int  `json:"x"`
			Rot   int  `json:"rot"`
			Y     int  `json:"y"`
			Index int  `json:"index"`
			FlipX bool `json:"flipX"`
			Tile  int  `json:"tile"`
		} `json:"tiles"`
		Name   string `json:"name"`
		Number int    `json:"number"`
	} `json:"layers"`
	Tileheight int `json:"tileheight"`
	Tileswide  int `json:"tileswide"`
	Tilewidth  int `json:"tilewidth"`
}

func HasTag(e *entity, tag Tag) bool {
	for _, t := range e.tags {
		if t == tag {
			return true
		}
	}

	return false
}

func GetEntsWithTag(game *Game, tag Tag) []*entity {
	ents := []*entity{}

	for _, e := range game.entities {
		if HasTag(e, tag) {
			ents = append(ents, e)
		}
	}

	return ents
}
