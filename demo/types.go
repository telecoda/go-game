package main

import (
	gogame "github.com/telecoda/go-game"
	"github.com/veandco/go-sdl2/sdl"
)

type initAssetsFunction func() error
type unloadAssetsFunction func() error

type DemoScreen struct {
	Id             int
	Title          string
	Description    string
	CodeSample     []string
	Color          sdl.Color
	InitAssets     initAssetsFunction
	UnloadAssets   unloadAssetsFunction
	RenderCallback gogame.RenderFunction
}
