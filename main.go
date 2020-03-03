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

var liveColor = color.NRGBA{R:255, G:255, B:255, A:255};
var deadColor = color.NRGBA{R:0,   G:0,   B:0,   A:255};

func main() {
	width, height := 640, 640
        seed := time.Now().UnixNano()

        flag.Int64Var(&seed, "s", seed, "set the seed.")
        flag.IntVar(&width, "x", width, "set the width.")
        flag.IntVar(&height, "y", height, "set the height.")

        pct := 0.5
        flag.Float64Var(&pct, "p", pct, "set random init percentage.")

        ruleset := -1
        flag.IntVar(&ruleset, "r", ruleset, "set ruleset.")

        outfn := "image.png"
        flag.StringVar(&outfn, "o", outfn, "set the output file.")


        flag.Parse()

        rand.Seed(seed)
        if ruleset < 0 { 
		ruleset = int(rand.Uint32())
        }

        ca := ca1d.NewCA1D(width, uint32(ruleset))
        ca.InitRandomly(pct)

	img := image.NewNRGBA(image.Rect(0, 0, width, height))

	for y := 0; y < height; y++ {
		for x := 0; x < ca.Width; x++ { 
			if ca.Current[x] == ca1d.Live { 
				img.Set(x, y, liveColor)
			} else {
				img.Set(x, y, deadColor)
			}
		}
		ca.Generate()
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
