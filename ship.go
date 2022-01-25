package main

import (
	"encoding/json"
	"image"

	"golang.org/x/image/math/f64"
)

func newShip(game *Game, pos f64.Vec2) *entity {
	ship := &entity{}

	ship.position = pos
	ship.active = false
	ship.tags = append(ship.tags, Ship)

	ic := newInput(ship, 10)
	ship.addComponent(ic)

	c := newCamera(ship)
	ship.addComponent(c)

	// Tiles
	if err := json.Unmarshal(ship1, &tileMap); err != nil {
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
			worldPos: ship.position,
		})
		// }
	}

	sr := newSpriteRenderer(ship, ips[0])
	ship.addComponent(sr)

	bc := newBoxCollider(ship, game, []string{Ship.String()}, 6)
	ship.addComponent(bc)

	sh := newShoots(ship)
	ship.addComponent(sh)

	game.entities = append(game.entities, ship)

	return ship
}
