package gokemon

func init() {
	audioManager = &AudioManager{}
}

func Wrap(f func()) {
	f()
	audioManager.PlaySoundEffect("select")
}
