package game

import (
	"errors"
	"fmt"
)

type ImageHandler struct{}
type FontHandler struct{}
type AudioHandler struct{}

type ResourceHandler interface {
	Load() error
	Unload() error
}

// Adds resource & loads it into memory
func (game *Game) AddResource(id, filePath string, resourceType ResourceType) error {

	resource := Resource{
		Id:       id,
		FilePath: filePath,
		Type:     resourceType,
		loaded:   false,
	}

	handler, err := initHandler(resourceType)
	if err != nil {
		fmt.Printf("Error initialising resource handler:%s\n", err)
		return err
	}
	resource.handler = handler

	handler.Load()
	game.resources["id"] = resource

	return nil

}

func initHandler(resourceType ResourceType) (ResourceHandler, error) {
	switch resourceType {
	case ImageResource:
		return ImageHandler{}, nil
	case FontResource:
		return FontHandler{}, nil
	case AudioResource:
		return AudioHandler{}, nil
	default:
		return ImageHandler{}, errors.New("Unknown resource type")
	}
}

// Resource handlers

func (a AudioHandler) Load() error {
	return nil
}

func (a AudioHandler) Unload() error {
	return nil
}

func (f FontHandler) Load() error {
	return nil
}

func (f FontHandler) Unload() error {
	return nil
}

func (i ImageHandler) Load() error {
	return nil
}

func (i ImageHandler) Unload() error {
	return nil
}
