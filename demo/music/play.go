package music

import "fmt"

type player interface {
	play(source string)
}

func paly(source, mtype string) {
	var a player
	switch mtype {
	case "MP3":
		a = &MP3Player{}
	case "WAV":
		a = &WAVPlayer{}
	default:
	}
	a.play(source)
}

type MP3Player struct {
	stat     int
	progress int
}

func (p *MP3Player) play(source string) {
	fmt.Println("MP3")
}

type WAVPlayer struct {
	stat     int
	progress int
}

func (p *WAVPlayer) play(source string) {
	fmt.Println("WAV")
}
