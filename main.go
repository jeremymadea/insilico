// Command insilico generates images by running a 1-dimensional cellular
// automaton chosen from a class of 2-state, 5-neighbor CAs. 

package main

import (
        "flag"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
        "github.com/jeremymadea/insilico/ca1d"
)

const defWidth, defHeight = 640, 640

var liveColor = color.NRGBA{R:255, G:255, B:255, A:255};
var deadColor = color.NRGBA{R:0,   G:0,   B:0,   A:255};

func hex2color(hex string) color.NRGBA {
	rgb := make([]uint8,4)
	for i:=0; i<len(hex) && i<6; i+=2 {
		if u, e := strconv.ParseUint(hex[i:i+2], 16, 8); e == nil { 
			rgb[i/2] = uint8(u)
		} else { 
			rgb[i/2] = 0
		}
	}
	return color.NRGBA{R:rgb[0], G:rgb[1], B:rgb[2], A:255}
}

func MakeImage(ca *ca1d.CA1D, h int, w, b color.Color) *image.NRGBA {
        img := image.NewNRGBA(image.Rect(0, 0, ca.Width, h))

        // We create one horizontal line of pixels per generation of
        // the CA.
        for y := 0; y < h; y++ {
                for x := 0; x < ca.Width; x++ {
                        if ca.Current[x] == ca1d.Live {
                                img.Set(x, y, w)
                        } else {
                                img.Set(x, y, b)
                        }
                }
                ca.Generate() // Calculate the next generation of the CA.
        }
	return img
}



func main() {
	defaultSeed := time.Now().UnixNano()
 
	// Setup and parse command line arguments.
	startPattern := "1"
	flag.StringVar(&startPattern, "s", startPattern, 
		"Set start pattern.")

	initMode := "random"
	flag.StringVar(&initMode, "m", initMode, 
		"Set init mode. random|center|ctralt|repeat|live|dead")
    
	seed := int64(-1)
	flag.Int64Var(&seed, "seed", seed, 
		"Set the seed. Use the system time if negative.")

	width, height := defWidth, defHeight
	flag.IntVar(&width, "x", width, "Set the `width` in pixels.")
	flag.IntVar(&height, "y", height, "Set the `height` in pixels.")

	pct := 0.5
	flag.Float64Var(&pct, "p", pct, "Set random init `percentage`.")

	ruleset := -1
	flag.IntVar(&ruleset, "r", ruleset, 
		"Set ruleset. Choose random one if negative.")

	rshex := ""
	flag.StringVar(&rshex, "R", rshex, 
		"Set ruleset. (8 digit hex) Note: -r will be ignored.")

	outfn := "image.png"
	flag.StringVar(&outfn, "o", outfn, 
		"Set the output file.")

	white := "FFFFFF" // white defaults to white. 
	flag.StringVar(&white, "w", white, "Set live cell color (hex).")

	black := "000000" // black defaults to black.
	flag.StringVar(&black, "b", black, "Set dead cell color (hex).")

	flag.Parse()

	liveColor = hex2color(white) // white might not be white.
	deadColor = hex2color(black) // black might not be black.

	// A negative value for seed means we must use the default.
	if seed < 0 { 
		rand.Seed(defaultSeed)
	} else {
		rand.Seed(seed)
	}

	// A negative value for ruleset means we must choose a random one.
	if ruleset < 0 { 
		ruleset = int(rand.Uint32())
	}

	if rshex != "" {
		if u, e := strconv.ParseUint(rshex, 16, 32); e == nil { 
			ruleset = int(u)
		}
	}

	ca := ca1d.NewCA1D(width, uint32(ruleset))

	switch initMode { 
		case "random":
			ca.InitRandom(pct)
		case "center":
			ca.InitCenter(ca1d.Dead, startPattern)
		case "ctralt":
			ca.InitCenter(ca1d.Live, startPattern)
		case "repeat":
			ca.InitRepeat(startPattern)
		case "live":
			ca.InitSimple(ca1d.Live)
		case "dead":
			ca.InitSimple(ca1d.Dead)
		default:
			initMode = "random"
			ca.InitRandom(pct)
	}

	img := MakeImage(ca, height, liveColor, deadColor)

        outfn = strings.Replace(outfn, "$", fmt.Sprintf("%08X",ruleset), 1)
        outfn = strings.Replace(outfn, "#", fmt.Sprintf("%d",ruleset), 1)

	f, err := os.Create(outfn)
	if err != nil {
		log.Fatal(err)
	}

	if err := png.Encode(f, img); err != nil {
		f.Close()
		log.Fatal(err)
	}

	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}
