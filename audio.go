package gogame

func (a audioPlayer) PlayAudio(audioAssetId string, loops int) error {

	chunk, err := getChunk(audioAssetId)
	if err != nil {
		return err
	}

	chunk.PlayChannel(-1, loops)

	return nil

}
