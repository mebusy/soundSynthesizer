
import numpy as np
import math
import matplotlib.pyplot as plt
import pylab


def saw_wave():
    x = np.linspace(0, 4*math.pi)
    # y1 = 2 * A / math.pi * ( math.pi*np.remainder( f*x + 0, 1.0 ) -math.pi/2   )
    t = f*x
    y1 = 2 * A / math.pi * ( math.pi*( t - np.fix(t) ) -math.pi/2   )
    t = f*x + P
    y2 = 2 * A / math.pi * ( math.pi*( t - np.fix(t) ) -math.pi/2   )
    y3 = (y1 -  y2)


    pylab.plot( x,y1 , color = "red")
    pylab.plot( x,y2 , color = "blue")
    pylab.plot( x,y3 , color = "black")
    pylab.show()

    pass


if __name__ == "__main__":
    f = 0.15
    A = 1

    P = 0.142

    saw_wave()
    # print np.mod( 7.15, 0.1 )

