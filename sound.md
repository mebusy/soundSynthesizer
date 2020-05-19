

```
w = A·sin(fx) + A 

0 <= w <= 2A
```

A is amplitude

low f is bass ,   high f is treble.


We need to represent the wave form as smoothly as possible.

The compute can only store numbers digitally to a fixed amount of precision. We can use a seconde function to do it.

![](imgs/sound_approximate_wave.png)


- b = 1
    - ![](imgs/sound_b_1.png)
- b = 8
    - ![](imgs/sound_b_8.png)
    - we now have 256 levels to represent the sine wave

As we apply more bits to represent the sine wave , the approximation of the sine wave becomes more and more accurate.

----

samplerate , the 44100, is really used to store the accuracy of the frequency.  We need a samplerate which is double the highest frequency we want to record. 

The human hearing range is about 20Hz to 20kHz,  so 44100 is little bit more than double 20k Hz, which means we should be able to in the most extreme cases capture with some accuracy 20k Hz signal. 

---

every douling in frequency moves you up 1 octave.


