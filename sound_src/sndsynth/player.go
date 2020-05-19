package sndsynth

import (
    "github.com/hajimehoshi/oto"
    "log"
    "time"
    "io"
)

var (
    sampleRate int // the number of samples that should be played during one second
    channelNum int // the number of channels. One channel is mono playback. 
                   // Two channels are stereo playback. Only values 1 and 2 are supported
    bitDepthInBytes int  // the number of bytes per sample per channel.(i.e. 8bit/16bit)
                         // The usual value is 2. Only values 1 and 2 are supported
    buffSizeInBytes int  // Bigger buffer can reduce the number of Player's Write calls, thus reducing CPU time.
                         // Smaller buffer enables more precise timing. 
)

type SoundPlayer struct {
    context *oto.Context
}

func NewSoundPlayer(_sampleRate, _channelNum, _bitDepthInBytes, _buffSizeInBytes int) *SoundPlayer {
    sampleRate = _sampleRate
    channelNum = _channelNum
    bitDepthInBytes = _bitDepthInBytes
    buffSizeInBytes = _buffSizeInBytes

    sp := &SoundPlayer{}
    // NewContext creates and holds ready-to-use Player objects.
    // go newDriver <- Context.mux
    c, err := oto.NewContext(sampleRate, channelNum, bitDepthInBytes, buffSizeInBytes )
    if err != nil {
        log.Fatal(err)
    }

    sp.context = c
    return sp
}

func (self *SoundPlayer) PlayFreq( freq float64, duration time.Duration, _func NOISE_FUNC) error {
    p := self.context.NewPlayer()
    s := NewWave(freq, duration, _func)
    if _, err := io.Copy(p, s); err != nil {
        return err
    }
    if err := p.Close(); err != nil {
        return err
    }
    return nil
}


func (self *SoundPlayer) Close() {
    self.context.Close()
}
