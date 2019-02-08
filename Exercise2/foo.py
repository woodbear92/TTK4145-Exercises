# Python 3.3.3 and 2.7.6
# python fo.py

from threading import Thread
from threading import RLock


# Potentially useful thing:
#   In Python you "import" a global variable, instead of "export"ing it when you declare it
#   (This is probably an effort to make you feel bad about typing the word "global")
i = 0
mtx = RLock()

def incrementingFunction():
    global i
    # TODO: increment i 1_000_000 times
    for x in range(0, 999998):
        mtx.acquire()
        i = i + 1
        mtx.release()

def decrementingFunction():
    global i
    # TODO: decrement i 1_000_000 times
    for x in range(0, 999999):
        mtx.acquire()
        i = i - 1
        mtx.release()

def main():
    # TODO: Something is missing here (needed to print i)
    global i
    incrementing = Thread(target = incrementingFunction, args = (),)
    decrementing = Thread(target = decrementingFunction, args = (),)

    # TODO: Start both threads
    incrementing.start()
    decrementing.start()

    incrementing.join()
    decrementing.join()

    print("The magic number is %d" % (i))


main()
