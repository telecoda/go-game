package game

import (
	b2d "github.com/neguse/go-box2d-lite/box2dlite"
	sdl "github.com/veandco/go-sdl2/sdl"
)

type Game struct {
	Window   *sdl.Window
	Renderer *sdl.Renderer

	resources ResourceMap
	quit      bool
}

type PhysicsGame struct {
	Game
	World *b2d.World
}

type ResourceType string

const (
	AudioResource ResourceType = "audio"
	FontResource  ResourceType = "font"
	ImageResource ResourceType = "image"
)

type Resource struct {
	Id       string
	FilePath string
	Type     ResourceType
	loaded   bool
	handler  ResourceHandler
}

type ResourceMap map[string]Resource
