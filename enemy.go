package main

import (
	"encoding/json"
	"image"

	"golang.org/x/image/math/f64"
)

func newEnemy(game *Game, pos f64.Vec2) *entity {
	enemy := &entity{}

	enemy.position = pos
	enemy.active = true
	enemy.tags = append(enemy.tags, Enemy)

	// Tiles
	if err := json.Unmarshal(enemy1, &tileMap); err != nil {
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
			worldPos: enemy.position,
		})
		// }
	}

	bai := newBasicAI(enemy)
	enemy.addComponent(bai)

	sr := newSpriteRenderer(enemy, ips[0])
	enemy.addComponent(sr)

	bc := newBoxCollider(enemy, game, Enemy.String(), 6)
	enemy.addComponent(bc)

	game.entities = append(game.entities, enemy)

	return enemy
}
