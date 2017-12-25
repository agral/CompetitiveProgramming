#!/usr/bin/env python3

"""
MIT License

Copyright (c) 2017 Adam Szczerbiak (szczerbiakadam@gmail.com)

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
"""

def emulate():
    r = {
        "a": 1,
        "b": 81,
        "c": 0,
        "d": 0,
        "e": 0,
        "f": 0,
        "g": 0,
        "h": 0
    }

    r["b"] *= 100
    r["b"] += 100000
    r["c"] = r["b"] + 17000

    done = False
    while not done:
        r["f"] = 1
        r["d"] = 2
        k = r["d"]
        while k * k < r["b"]:
            if r["b"] % k == 0:
                r["f"] = 0
                break
            k += 1
        if r["f"] == 0:
            r["h"] += 1
        r["g"] = r["b"] - r["c"]
        r["b"] += 17
        done = (r["g"] == 0)

    print(r["h"])

emulate()
