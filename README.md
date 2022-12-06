advent of code 2022
====================

https://adventofcode.com/2022

usage
=====

1. `$ pip install -U advent-of-code-data`
1. `$ export AOCDIR="$HOME/PATH TO YOUR AOC DIRECTORY"`
1. `$ export AOC_SESSION=cafef00db01dfaceba5eba11deadbeef`
1. `$ python runner.py`


caveats
=======

- For the Golang option, see the `README.md` in that directory.

- You can use the built-in `aocd-token` utility script. However, be
advised that this utility script attempts to find session tokens from
your browser’s cookie storage. This feature is experimental and requires
you to additionally install the package browser-cookie3. Only Chrome and
Firefox browsers are currently supported. On macOS, you may get an
authentication dialog requesting permission, since Python is attempting
to read browser storage files. This is expected, the script is actually
scraping those private files to access AoC session token(s).

- If this utility script was able to locate your token, you can save it to
file with:

  `$ aocd-token > ~/.config/aocd/token`
