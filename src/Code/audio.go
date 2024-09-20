package gokemon

import (
	"os"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

type AudioManager struct {
	backgroundMusic beep.StreamSeekCloser
	battleMusic     beep.StreamSeekCloser
	soundEffects    map[string]*beep.Buffer
	format          beep.Format
	initialized     bool
}

func NewAudioManager() *AudioManager {
	return &AudioManager{
		soundEffects: make(map[string]*beep.Buffer),
	}
}

func (am *AudioManager) Initialize() error {
	if am.initialized {
		return nil
	}

	err := speaker.Init(44100, 44100/30)
	if err != nil {
		return err
	}

	am.initialized = true
	return nil
}

func (am *AudioManager) LoadBackgroundMusic(assets string) error {
	if err := am.Initialize(); err != nil {
		return err
	}

	f, err := os.Open(assets)
	if err != nil {
		return err
	}

	streamer, format, err := mp3.Decode(f)
	if err != nil {
		f.Close()
		return err
	}

	am.backgroundMusic = streamer
	am.format = format
	return nil
}

func (am *AudioManager) PlayBackgroundMusic() {
	if am.backgroundMusic == nil {
		return
	}
	speaker.Play(beep.Loop(-1, am.backgroundMusic))

}
func (am *AudioManager) StopMusic() {
	speaker.Clear()
}
func (am *AudioManager) LoadSoundEffect(name, filename string) error {
	if err := am.Initialize(); err != nil {
		return err
	}

	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	streamer, format, err := mp3.Decode(f)
	if err != nil {
		return err
	}

	buffer := beep.NewBuffer(format)
	buffer.Append(streamer)
	streamer.Close()

	am.soundEffects[name] = buffer
	return nil
}

func (am *AudioManager) PlaySoundEffect(name string) {
	if buffer, ok := am.soundEffects[name]; ok {
		speaker.Play(buffer.Streamer(0, buffer.Len()))
	}
}
func (am *AudioManager) LoadBattleMusic(assets string) error {
	if err := am.Initialize(); err != nil {
		return err
	}
	
	f, err := os.Open(assets)
	if err != nil {
		return err
	}

	streamer, format, err := mp3.Decode(f)
	if err != nil {
		f.Close()
		return err
	}

	am.battleMusic = streamer
	am.format = format
	return nil
}

func (am *AudioManager) PlayBattleMusic() {
	if am.battleMusic == nil {
		return
	}
	speaker.Play(beep.Loop(-1, am.battleMusic))
}
