package main

import (
	"reflect"

	"golang.org/x/image/math/f64"
)

type entity struct {
	name       string
	position   f64.Vec2
	direction  f64.Vec2
	rotation   float64
	active     bool
	components []component
	tags       []Tag
}

func (e *entity) addComponent(comp component) {
	for _, ec := range e.components {
		if reflect.TypeOf(ec) == reflect.TypeOf(comp) {
			panic("Component already added")
		}
	}

	e.components = append(e.components, comp)
}

func (e *entity) getComponent(withType component) component {
	compType := reflect.TypeOf(withType)

	for _, ec := range e.components {
		if reflect.TypeOf(ec) == compType {
			return ec
		}
	}

	// panic("Component not found")
	return nil
}

type Tag int

const (
	Player Tag = iota
	Enemy
	Scene
	Ship
	Bullet
)

func (t Tag) String() string {
	switch t {
	case Player:
		return "player"
	case Enemy:
		return "enemy"
	case Scene:
		return "scene"
	case Ship:
		return "ship"
	case Bullet:
		return "bullet"
	}
	return "unknown"
}
