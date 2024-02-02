# Stega 🩻

Stega is a command line tool that allows you to embed text into images using LSB steganography.

<img src='.github/assets/stega-no-bg.png' width='350px'>

## Installation

```bash
go install github.com/brittonhayes/stega/cmd/stega@latest
```

## Usage

```bash
stega [command]
```

### Commands

#### Encode

Encode text into an image.

```bash
stega encode --input [input file] --output [output file] --message [message to encode]
```

Options:

- `--input` or `-i`: The input file (if no image provided, will download a random image from lorempicsum)
- `--output` or `-o`: The output file (default: `output.png`)
- `--message` or `-m`: The message to encode

#### Decode

Decode text from an image.

```bash
stega decode --input [input file]
```

Options:

- `--input` or `-i`: The input file (required)

## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

## License

[Open Software License 3.0](./LICENSE)
