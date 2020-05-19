package sndsynth

import (
    "time"
    "io"
)

type NOISE_FUNC func( float64,float64 ) float64

type Wave struct {
    freq   float64
    length int64
    pos    int64

    remaining []byte
    noise_func NOISE_FUNC
}

func NewWave(freq float64, duration time.Duration, _func NOISE_FUNC ) *Wave {
    // total bytes used for play such piece sound
    l := int64(channelNum) * int64(bitDepthInBytes) * int64(sampleRate) * int64(duration) / int64(time.Second)
    // l = l / 4 * 4
    l &^= 3  // protect the case that channelNum=2 and bitDepthInBytes=2
    return &Wave{
        freq:   freq,
        length: l,   // total bytes
        noise_func: _func ,
    }
}



func (s *Wave) Read(buf []byte) (int, error) {
    if len(s.remaining) > 0 {
        n := copy(buf, s.remaining)
        s.remaining = s.remaining[n:]
        return n, nil
    }

    if s.pos == s.length {
        return 0, io.EOF
    }

    eof := false
    if s.pos+int64(len(buf)) > s.length {
        buf = buf[:s.length-s.pos]
        eof = true
    }

    var origBuf []byte
    if len(buf)%4 > 0 {
        origBuf = buf
        buf = make([]byte, len(origBuf)+4-len(origBuf)%4)
    }

    // each period will be sampled `length` times
    // length := float64(sampleRate) / float64(s.freq)
    // `num` is bytes for each sample
    num := (bitDepthInBytes) * (channelNum)
    // postion in sample unit
    p := s.pos / int64(num)
    switch bitDepthInBytes {
    case 1:
        for i := 0; i < len(buf)/num; i++ {
            const max = 255
            b := int(s.noise_func( s.freq, float64(p)/float64(sampleRate) ) *max) + (max+1)/2
            for ch := 0; ch < channelNum; ch++ {
                buf[num*i+ch] = byte(b)
            }
            p++
        }
    case 2:
        for i := 0; i < len(buf)/num; i++ {
            const max = 32767
            // b := int16(math.Sin(2*math.Pi*float64(p)/length) * 0.3 * max)
            b := int(s.noise_func( s.freq, float64(p)/float64(sampleRate) ) *max) + (max+1)/2
            for ch := 0; ch < channelNum; ch++ {
                buf[num*i+2*ch] = byte(b)
                buf[num*i+1+2*ch] = byte(b >> 8)
            }
            p++
        }
    }

    s.pos += int64(len(buf))

    n := len(buf)
    if origBuf != nil {
        n = copy(origBuf, buf)
        s.remaining = buf[n:]
    }

    if eof {
        return n, io.EOF
    }
    return n, nil
}

