package main

import (
	"encoding/json"
	"image"

	"golang.org/x/image/math/f64"
)

func newPlayer(game *Game, pos f64.Vec2) *entity {
	player := &entity{}

	player.position = pos
	player.active = true
	player.tags = append(player.tags, Player)

	ic := newInput(player)
	player.addComponent(ic)

	c := newCamera(player)
	player.addComponent(c)

	// Tiles
	if err := json.Unmarshal(player1, &tileMap); err != nil {
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
			worldPos: player.position,
		})
		// }
	}

	sr := newSpriteRenderer(player, ips[0])
	player.addComponent(sr)

	bc := newBoxCollider(player, game)
	player.addComponent(bc)

	game.entities = append(game.entities, player)

	return player
}
