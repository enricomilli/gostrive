# gostrive

gostrive is a simple command-line application written in Go that delivers inspiring quotes to brighten your day.

## Installation

To install gostrive, make sure you have Go installed on your system, then run:
```
go get github.com/enricomilli/gostrive
```

## Usage

Once installed, you can use gostrive by simply typing:
```
gostrive
```

This will display a random inspiring quote.

### Specifying an Author

If you want a quote from a specific author, you can use the `--author` or `-a` flag:

```
gostrive --author "Albert Einstein"
```

or

```
gostrive -a "Maya Angelou"
```

Capitalization does not matter.

## Examples

```
$ gostrive
"The only way to do great work is to love what you do." - Steve Jobs


$ gostrive -a "Nelson Mandela"
"It always seems impossible until it's done." - Nelson Mandela
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE.md) file for details.
