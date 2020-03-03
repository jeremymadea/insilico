// Command insilico generates images by running a 1-dimensional cellular
// automaton chosen from a class of 2-state, 5-neighbor CAs. 

package main

import (
        "flag"
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
        "github.com/jeremymadea/insilico/ca1d"
	"math/rand"
	"time"
)

const defWidth, defHeight = 640, 640

var liveColor = color.NRGBA{R:255, G:255, B:255, A:255};
var deadColor = color.NRGBA{R:0,   G:0,   B:0,   A:255};

func main() {
	defaultSeed := time.Now().UnixNano()
 
	// Setup and parse command line arguments.
	startPattern := "1"
	flag.StringVar(&startPattern, "s", startPattern, 
		"Set start pattern.")

	initMode := "random"
	flag.StringVar(&initMode, "m", initMode, 
		"Set init mode. (random|center|ctralt|repeat|live|dead)")
    
	seed := int64(-1)
	flag.Int64Var(&seed, "seed", seed, 
		"Set the seed. Use the system time if negative.")

	width, height := defWidth, defHeight
	flag.IntVar(&width, "x", width, "Set the width.")
	flag.IntVar(&height, "y", height, "Set the height.")

	pct := 0.5
	flag.Float64Var(&pct, "p", pct, "Set random init percentage.")

	ruleset := -1
	flag.IntVar(&ruleset, "r", ruleset, 
		"Set ruleset. Choose random one if negative.")

	outfn := "image.png"
	flag.StringVar(&outfn, "o", outfn, "Set the output file.")

	flag.Parse()


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

	img := image.NewNRGBA(image.Rect(0, 0, width, height))

	// We create one horizontal line of pixels per generation of 
	// the CA.
	for y := 0; y < height; y++ {
		for x := 0; x < ca.Width; x++ { 
			if ca.Current[x] == ca1d.Live { 
				img.Set(x, y, liveColor)
			} else {
				img.Set(x, y, deadColor)
			}
		}
		ca.Generate() // Calculate the next generation of the CA.
	}

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
