package commands

import (
	"bufio"
	"bytes"
	"fmt"
	"image/png"
	"os"

	"github.com/auyer/steganography"
	"github.com/brittonhayes/stega/pkg/plugin"
	"github.com/urfave/cli/v2"
)

type defaultEncoder struct{}

func DefaultEncoder() plugin.Encoder {
	return &defaultEncoder{}
}

func (d *defaultEncoder) Encode() cli.ActionFunc {
	return func(ctx *cli.Context) error {
		inputFile, err := os.Open(ctx.String("input"))
		if err != nil {
			return fmt.Errorf("error opening file %v", err)
		}

		reader := bufio.NewReader(inputFile)
		img, err := png.Decode(reader)
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

func Decode() cli.ActionFunc {
	return func(ctx *cli.Context) error {
		inputFile, err := os.Open(ctx.String("input"))
		if err != nil {
			return fmt.Errorf("error opening file %v", err)
		}

		reader := bufio.NewReader(inputFile)
		img, err := png.Decode(reader)
		if err != nil {
			return fmt.Errorf("error decoding file %v", err)
		}

		msgSize := steganography.GetMessageSizeFromImage(img)

		message := steganography.Decode(msgSize, img)
		if err != nil {
			return fmt.Errorf("error decoding file %v", err)
		}

		if ctx.String("output") != "" {
			outFile, err := os.Create(ctx.String("output"))
			if err != nil {
				return fmt.Errorf("error creating file %v", err)
			}
			defer outFile.Close()

			outFile.Write(message)

			fmt.Println("message saved to", ctx.String("output"))
			return nil
		}

		fmt.Println(string(message))
		return nil
	}
}
