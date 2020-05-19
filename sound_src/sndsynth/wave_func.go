package sndsynth

import (
    "math"
)

func NoiseSine( freq float64 , x float64 ) float64 {
    return 0.3 * math.Sin( freq * 2*math.Pi * x  )
}



func NoiseSquare( freq float64 , x float64 ) float64 {
    output :=  math.Sin( freq * 2*math.Pi * x  )
    if output > 0 {
        return 0.3
    } else {
        return -0.3
    }
}


