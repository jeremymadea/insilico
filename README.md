# insilico

Runs an augmented 1D elementary cellular automaton with two cell states and a 
five cell neighborhood and creates a PNG image where each pixel in a row 
represents a cell and each row represents a successive generation of the CA.

```
Usage of ./insilico:
  -m string
        Set init mode. random|center|ctralt|repeat|live|dead (default "random")
  -o string
        Set the output file. (default "image.png")
  -p float
        Set random init percentage. (default 0.5)
  -r int
        Set ruleset. Choose random one if negative. (default -1)
  -s string
        Set start pattern. (default "1")
  -seed int
        Set the seed. Use the system time if negative. (default -1)
  -x int
        Set the width. (default 640)
  -y int
        Set the height. (default 640)
```
