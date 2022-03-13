# XOR two files into a third

This tool just reads two files, assumes they are the same size (HA!), and writes the bitwise XOR of the contents to a third file.

I wrote this because I DO NOT LIKE PYTHON

## Usage

`xor --file1 <filename> --file2 <filename> --output <filename>`

## Notes

The xor.py script provided by Seed Labs assumes you will provide hexadecimal output, thus you have to do many `xxd -p` and `xxd -r -p` conversions.
This operates directly on the binary file, so there is no conversion needed.
EXCEPT when you call `openssl enc` : the `-K` and `-iv` flags expect hexadecimal input, so you need to `xxd -p` your output from this script if you are using it as an input in those cases.

## Building

go is awesome.  just `go build .` and you will get a standalone functional binary
