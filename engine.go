package gogame

func (e *Engine) GetAssetManager() AssetManager {
	return e.AssetManager
}

func (e *Engine) GetAudioPlayer() AudioPlayer {
	return e.audioPlayer
}

func (e *Engine) GetRenderer() Renderer {
	return e.renderer
}
