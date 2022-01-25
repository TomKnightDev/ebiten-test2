package main

import (
	"encoding/json"
	"image"

	"golang.org/x/image/math/f64"
)

func newBullet(game *Game, pos f64.Vec2, dir f64.Vec2, originTag Tag) *entity {
	bullet := &entity{}

	bullet.position = pos
	bullet.direction = dir
	bullet.active = true
	bullet.tags = append(bullet.tags, Bullet)
	bullet.lifetime = 100
	bullet.velocity = 2

	// Tiles
	if err := json.Unmarshal(bullet1, &tileMap); err != nil {
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
			worldPos: bullet.position,
		})
		// }
	}

	sr := newSpriteRenderer(bullet, ips[0])
	bullet.addComponent(sr)

	bc := newBoxCollider(bullet, game, []string{originTag.String()}, 2)
	bullet.addComponent(bc)

	game.entities = append(game.entities, bullet)

	return bullet
}
