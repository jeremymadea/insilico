# insilico

Runs an augmented 1D elementary cellular automaton with two cell states and a 
five cell neighborhood and creates a PNG image where each pixel in a row 
represents a cell and each row represents a successive generation of the CA.

```
Usage of ./insilico:
  -R string
        Set ruleset. (8 digit hex) Note: -r will be ignored.
  -b string
        Set dead cell color (hex). (default "000000")
  -m string
        Set init mode. random|center|ctralt|repeat|live|dead (default "random")
  -o string
        Set the output file. (default "image.png")
  -p percentage
        Set random init percentage. (default 0.5)
  -r int
        Set ruleset. Choose random one if negative. (default -1)
  -s string
        Set start pattern. (default "1")
  -seed int
        Set the seed. Use the system time if negative. (default -1)
  -w string
        Set live cell color (hex). (default "FFFFFF")
  -x width
        Set the width in pixels. (default 640)
  -y height
        Set the height in pixels. (default 640)
```
