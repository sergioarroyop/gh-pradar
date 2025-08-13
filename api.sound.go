package main

import (
	"bytes"
	"embed"
	"sync"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/speaker"
	"github.com/faiface/beep/wav"
)

//go:embed sounds/*.wav
var notifySounds embed.FS

var (
	initOnce sync.Once
	initErr  error

	format beep.Format
	buf    *beep.Buffer

	playMu sync.Mutex
)

func initAudio() error {
	initOnce.Do(func() {
		file, _ := notifySounds.ReadFile("sounds/" + config.Sound + ".wav")
		streamer, fmt, err := wav.Decode(bytes.NewReader(file))
		if err != nil {
			initErr = err
			return
		}
		defer streamer.Close()

		format = fmt

		b := beep.NewBuffer(format)
		b.Append(streamer)
		buf = b

		initErr = speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	})
	return nil
}

func playAudio() {
	if buf == nil || initErr != nil {
		return
	}

	s := buf.Streamer(0, buf.Len())
	speaker.Play(s)
}
