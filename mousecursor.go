package main

import (
	"encoding/json"
	"image"
)

func newMouseCursor(game *Game) *entity {
	mouseCursor := &entity{}

	mouseCursor.active = true

	// Tiles
	if err := json.Unmarshal(cursor1, &tileMap); err != nil {
		panic(err)
	}

	ips := []imagePos{}

	for l := len(tileMap.Layers) - 1; l >= 0; l-- {
		// for _, t := range tileMap.Layers[l].Tiles {
		t := tileMap.Layers[l].Tiles[0]
		sx := (t.Tile % tileXNum) * tileSize
		sy := (t.Tile / tileXNum) * tileSize

		ips = append(ips, imagePos{
			sheetPos: image.Rect(sx, sy, sx+tileSize, sy+tileSize),
			worldPos: mouseCursor.position,
		})
		// }
	}

	sr := newSpriteRenderer(mouseCursor, ips[0])
	mouseCursor.addComponent(sr)

	mi := newMouseInput(mouseCursor)
	mouseCursor.addComponent(mi)

	game.entities = append(game.entities, mouseCursor)

	return mouseCursor
}
