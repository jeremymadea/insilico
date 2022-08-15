// Insilico generates images by running a 1-dimensional cellular
// automaton chosen from a class of 2-state, 5-neighbor CAs. 

/* MIT License

Copyright (c) 2022 Jeremy Madea

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
*/


package main

import (
	"flag"
	"fmt"
	"html/template"
	"image"
	"image/color"
	"image/png"
	"log"
	"math/rand"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
	"github.com/jeremymadea/insilico/ca1d"
)

const defWidth, defHeight = 400, 400

type Config struct {
	Width, Height int
	Pattern string
	InitMode string
	Seed int64
	Percentage float64
	Ruleset int
	RSHex string
	LColorHex string
	LiveColor color.Color
	DColorHex string
	DeadColor color.Color
}

var global Config
var gimg *image.NRGBA

func openbrowser(url string) {
	var err error

	err = exec.Command("open", url).Start()
	if err != nil {
		log.Fatal(err)
	}
}

func querystring(c Config) string {
	return "?" +
		"rd=" + fmt.Sprintf("%d", c.Ruleset) +
		"&w=" + fmt.Sprintf("%d", c.Width) +
		"&h=" + fmt.Sprintf("%d", c.Height) +
		"&p=" + fmt.Sprintf("%.2f", c.Percentage) +
		"&m=" + c.InitMode +
		"&s=" + c.Pattern +
		"&lc=" + c.LColorHex +
		"&dc=" + c.DColorHex

}

func html_modeopts(curmode string) template.HTML {
	modes := []string{"random", "center", "ctralt", "repeat", "live", "dead"}
	s := ""
	selected := ""
	for _, m := range modes {
		if m == curmode {
			selected = " selected"
		} else {
			selected = ""
		}
		s += fmt.Sprintf(`<option value="%v"%v>%v</option>`, m, selected, m)
		s += "\n"
	}
	return template.HTML(s)
}

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

func mutations(ruleset int) []int {
	ret := make([]int, 32)
	for i:=0; i<32; i++ {
		ret[i] = ruleset ^ (1 << i)
	}
	return ret
}

func bits2colors(ruleset int) []string {
	ret := make([]string, 32)
	for i:=0; i<32; i++ {
		if 1 & (ruleset >> i) == 1 {
			ret[i] = "wht"
		} else {
			ret[i] = "blk"
		}
	}
	return ret
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

func ExecuteCA(cfg Config) *image.NRGBA {
	ca := ca1d.NewCA1D(cfg.Width, uint32(cfg.Ruleset))
	switch cfg.InitMode {
		case "random":
			ca.InitRandom(cfg.Percentage)
		case "center":
			ca.InitCenter(ca1d.Dead, cfg.Pattern)
		case "ctralt":
			ca.InitCenter(ca1d.Live, cfg.Pattern)
		case "repeat":
			ca.InitRepeat(cfg.Pattern)
		case "live":
			ca.InitSimple(ca1d.Live)
		case "dead":
			ca.InitSimple(ca1d.Dead)
		default:
			cfg.InitMode = "random"
			ca.InitRandom(cfg.Percentage)
	}
	return MakeImage(ca, cfg.Height, cfg.LiveColor, cfg.DeadColor)
}

func NewConfigFromForm(r *http.Request) (cfg Config) {
	cfg = global
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k := range r.Form {
		v := r.Form.Get(k)
		switch k {
			case "rd":
				if iv, err := strconv.Atoi(v); err == nil {
					cfg.Ruleset = iv
				}
				if cfg.Ruleset < 0 {
					cfg.Ruleset = int(rand.Uint32())
				}
			case "rh":
				if uv, e := strconv.ParseUint(v, 16, 32); e == nil {
					cfg.Ruleset = int(uv)
				}
			case "s":
				cfg.Pattern = v
			case "p":
				if fv, err := strconv.ParseFloat(v, 64); err == nil {
					cfg.Percentage = fv
				}
			case "m":
				cfg.InitMode = v
				switch cfg.InitMode {
					case "random":
					case "center":
					case "ctralt":
					case "repeat":
					case "live":
					case "dead":
					default: cfg.InitMode = "random"
				}
			case "w":
				if iv, err := strconv.Atoi(v); err == nil {
					if iv > 0 {
						cfg.Width = iv
					}
				}
			case "h":
				if iv, err := strconv.Atoi(v); err == nil {
					if iv > 0 {
						cfg.Height = iv
					}
				}
			case "dc":
				if v[0] == '#' && len(v) == 7 {
					v = v[1:]
				}
				cfg.DColorHex = v
			case "lc":
				if v[0] == '#' && len(v) == 7 {
					v = v[1:]
				}
				cfg.LColorHex = v
			default:
				log.Print("Bad Form Key: ", k)
		}
	}
	cfg.RSHex = fmt.Sprintf("%08X", cfg.Ruleset)
	cfg.DeadColor = hex2color(cfg.DColorHex)
	cfg.LiveColor = hex2color(cfg.LColorHex)
	return cfg
}

func main() {
	defaultSeed := time.Now().UnixNano()

	// Setup and parse command line arguments.
	explorerMode := false
	flag.BoolVar(&explorerMode, "explorer", explorerMode,
		"Start in explorer mode. Runs as a webapp on http://localhost:8084.")

	port := uint(8084)
	flag.UintVar(&port, "port", port, "Select the port to run the webapp on.")

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

	liveColor := hex2color(white) // white might not be white.
	deadColor := hex2color(black) // black might not be black.

	// Copy vars into global Config struct. TODO: remove the intermediate vars
	global.LColorHex = white
	global.LiveColor = liveColor
	global.DColorHex = black
	global.DeadColor = deadColor

	global.Width = width
	global.Height = height
	global.Percentage = pct
	global.InitMode = initMode
	global.Pattern = startPattern

	// A negative value for seed means we must use the default.
	if seed < 0 {
		rand.Seed(defaultSeed)
		global.Seed = defaultSeed
	} else {
		rand.Seed(seed)
		global.Seed = seed
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

	global.Ruleset = ruleset
	global.RSHex = fmt.Sprintf("%08X", ruleset)
	gimg = ExecuteCA(global)

	if explorerMode {
		http.HandleFunc("/", webRoot)
		http.HandleFunc("/debug", webDebug)
		http.HandleFunc("/scope", webScope)
		http.HandleFunc("/builder", webBuilder)
		http.HandleFunc("/capng", webCapng)
		http.HandleFunc("/shutdown", webShutdown)
		webserver := fmt.Sprintf("localhost:%v", port)
		go openbrowser("http://" + webserver)
		log.Fatal(http.ListenAndServe(webserver, nil))
	} else {
		// Running as a command line tool. Just write our image out.
		outfn = strings.Replace(outfn, "$", fmt.Sprintf("%08X",ruleset), 1)
		outfn = strings.Replace(outfn, "#", fmt.Sprintf("%d",ruleset), 1)

		f, err := os.Create(outfn)
		if err != nil {
			log.Fatal(err)
		}

		if err := png.Encode(f, gimg); err != nil {
			f.Close()
			log.Fatal(err)
		}

		if err := f.Close(); err != nil {
				log.Fatal(err)
		}
	}
}

// Following are handlers for web requests when we are running in explorer mode

func webDebug(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}
}

func webRoot(w http.ResponseWriter, r *http.Request) {
	cfg := NewConfigFromForm(r);
	t := template.Must(template.New("root").Parse(rootHtml))
	muts := mutations(cfg.Ruleset)

	data := struct{
		C Config
		MutationSetA []int
		MutationSetB []int
		MutationSetC []int
		MutationSetD []int
		MutationSetE []int
		MutationSetF []int
		MutationSetG []int
	}{
		cfg,
		muts[:8],
		muts[8:10],
		muts[10:12],
		muts[12:16],
		muts[16:20],
		muts[20:24],
		muts[24:],
	}
	err := t.Execute(w, data)
	if (err != nil) {
		log.Print(err)
	}
}

func webBuilder(w http.ResponseWriter, r *http.Request) {
	cfg := NewConfigFromForm(r);
	t := template.Must(template.New("builder").Parse(builderHtml))

	//cfg.Ruleset
	muts := mutations(cfg.Ruleset)
	clrs := bits2colors(cfg.Ruleset)

	data := struct{
		C Config
		Color []string
		RSNew []int
	}{
		cfg,
		clrs,
		muts,
	}

	err := t.Execute(w, data)
	if (err != nil) {
		log.Print(err)
	}
}

func webScope(w http.ResponseWriter, r *http.Request) {
	cfg := NewConfigFromForm(r);
	qs := querystring(cfg)
	modeopts := html_modeopts(cfg.InitMode)

	if r.URL.RawQuery == "" {
		cfg.Width = 640;
		cfg.Height = 640;
	}

	t := template.Must(template.New("scope").Parse(scopeHtml))

	data := struct {
		C Config
		QueryString string
		ModeOpts template.HTML
	}{
		cfg,
		qs,
		modeopts,
	}

	err := t.Execute(w, data)
	if (err != nil) {
		log.Print(err)
	}
}

func webCapng(w http.ResponseWriter, r *http.Request) {
	cfg := NewConfigFromForm(r);
	img := ExecuteCA(cfg)

	png.Encode(w, img)
}

func webShutdown(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Shutting down InSilico Explorer.")
	os.Exit(0)
}
