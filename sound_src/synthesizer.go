package main

import (
    "sync"
    "time"

    "sndsynth"
    // "log"
)



func main()  {
    sampleRate := 44100
    channelNum := 1
    bitDepthInBytes := 1
    buffSizeInBytes := 4096

    player := sndsynth.NewSoundPlayer( sampleRate, channelNum, bitDepthInBytes, buffSizeInBytes )

    var wg sync.WaitGroup

    wg.Add(1)
    go func() {
        defer wg.Done()
        var freqC   float64 = 440
        if err := player.PlayFreq( freqC, 3*time.Second, sndsynth.NoiseRandom); err != nil {
            panic(err)
        }
    }()



    wg.Wait()
    player.Close()
}
