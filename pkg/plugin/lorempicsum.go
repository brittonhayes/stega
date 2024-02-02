package plugin

import (
	"bytes"
	"fmt"
	"image/jpeg"
	"net/http"
	"os"

	"github.com/auyer/steganography"
	"github.com/urfave/cli/v2"
)

func LoremPicsumEncoder() Encoder {
	return &loremPicsumEncoder{
		url: "https://picsum.photos/200/300",
	}
}

type loremPicsumEncoder struct {
	url string
}

func (l *loremPicsumEncoder) Encode() cli.ActionFunc {
	return func(ctx *cli.Context) error {
		resp, err := http.Get(l.url)
		if err != nil {
			return fmt.Errorf("error fetching image %v", err)
		}
		defer resp.Body.Close()

		img, err := jpeg.Decode(resp.Body)
		if err != nil {
			return fmt.Errorf("error decoding file %v", err)
		}

		w := new(bytes.Buffer)
		err = steganography.Encode(w, img, []byte(ctx.String("message")))
		if err != nil {
			return fmt.Errorf("error encoding file %v", err)
		}

		outFile, err := os.Create(ctx.String("output"))
		if err != nil {
			return fmt.Errorf("error creating file %v", err)
		}
		defer outFile.Close()

		w.WriteTo(outFile)
		return nil
	}
}
