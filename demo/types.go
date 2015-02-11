package main

import (
	gogame "github.com/telecoda/go-game"
	"github.com/veandco/go-sdl2/sdl"
)

type initAssetsFunction func() error

type DemoScreen struct {
	Id             int
	Title          string
	Color          sdl.Color
	InitAssets     initAssetsFunction
	RenderCallback gogame.RenderFunction
}
