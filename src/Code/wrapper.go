package gokemon

func Wrap(f func()) {
	f()
	audioManager.PlaySoundEffect("select")
}
