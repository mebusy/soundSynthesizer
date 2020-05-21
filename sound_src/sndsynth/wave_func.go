package sndsynth

import (
    "math"
    "math/rand"
)

const (
    WAVE_SINE int = iota
    WAVE_SQUARE
    WAVE_TRIANGLE
    WAVE_SAW_SLOW
    WAVE_SAW_FAST
    WAVE_RANDOM_NOISE
)

// convert frequency (Hz) to angular velocity
func w( hz float64 ) float64 {
    return hz * 2 * math.Pi
}

// Oscillator
func osc( hz, dTime float64, nType int )  float64 {

    switch nType {
    case WAVE_SINE:
        return math.Sin( w(hz) * dTime )
    case WAVE_SQUARE:
        if math.Sin( w(hz) * dTime ) > 0 {
            return 1
        } else {
            return -1
        }
    case WAVE_TRIANGLE:
        return math.Asin(  math.Sin( w(hz) * dTime ) *2 / math.Pi )
    case WAVE_SAW_SLOW:
        var output float64 = 0
        var n float64
        for n=1.0; n<100.0; n++ {
            output += -math.Sin( n*w(hz)*dTime ) / n
        }
        return output * 2 / math.Pi
    case WAVE_SAW_FAST:
        return (2/math.Pi) * ( hz *math.Pi * math.Mod( dTime, 1.0/hz ) - (math.Pi/2.0)  )
    case WAVE_RANDOM_NOISE:
        return 2*rand.Float64() - 1.0

    default:
        return 0
    }
}

func NoiseSine( freq float64 , x float64 ) float64 {
    output := osc( freq, x, WAVE_SINE  )
    var ampl float64 = 0.4
    return output *  ampl + ampl
}

func NoiseSquare( freq float64 , x float64 ) float64 {
    output := osc( freq, x, WAVE_SQUARE  )
    var ampl float64 = 0.2
    return output *  ampl + ampl
}

func NoiseTriangle( freq float64 , x float64 ) float64 {
    output := osc( freq, x, WAVE_TRIANGLE  )
    var ampl float64 = 0.4
    return output *  ampl + ampl
}

func NoiseSaw( freq float64 , x float64 ) float64 {
    output := osc( freq, x, WAVE_SAW_FAST  )
    var ampl float64 = 0.4
    return output *  ampl + ampl
}

func NoiseRandom( freq float64 , x float64 ) float64 {
    output := osc( freq, x, WAVE_RANDOM_NOISE  )
    var ampl float64 = 0.4
    return output *  ampl + ampl
}

var fFrequency float64 = 440
var fDutyCycle float64 = 0.5
var fHarmonics float64 = 20

func sampleSqaureWave( f, t float64 ) float64 {
    var a float64  // a,b represent the sample values of the underlying sine wave forms
    var b float64
    var p float64 = fDutyCycle * 2 * math.Pi

    for n := 1.0; n< fHarmonics ; n++ {
        c := n * f * 2.0 * math.Pi * t
        a += math.Sin( c ) / n
        b += math.Sin( c - p*n ) / n
    }
    return (2 / math.Pi) * (a-b)
}

func NoiseSquarePulse( freq float64 , x float64 ) float64 {
    output := sampleSqaureWave( freq, x )
    var ampl float64 = 0.2
    return output *  ampl + ampl*2   // output sometimes < -1 or > 1
}

