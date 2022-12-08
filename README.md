# nomnemonic-cli

Generate deterministic mnemonic words

`nomnemonic-cli` is a deterministic mnemonic generator app that uses 3 inputs and cryptographic hash functions to generate the words.

## Install

```
go install github.com/nomnemonic/nomnemonic-cli@latest
```

To install an earlier version:
```
# v0.2.0 / algorithm 2.0.0
go install github.com/nomnemonic/nomnemonic-cli@v0.2.0

# v0.1.0 / algorithm 1.0.0
go install github.com/nomnemonic/nomnemonic-cli@v0.1.0
```

All the changes will be backward compatible with semantic versioning. Once the library gets to `v1.0.0`, it will support any version with an optional input from user.

## Usage

### Flags (mode)

**Supported mnemonic words sizes**
12, 15, 18, 21, 24

```
nomnemonic-cli generate --identifier test@example --password test12345678 --passcode 001234 --size 24

# drop captain ring slice label oyster galaxy jacket online debris security doctor drive gadget nominee film tunnel away network spirit duty anger slender train
```

**Shorthand letters**

```
nomnemonic-cli generate -i test@example -p test12345678 -c 001234 -s 12

# during wave century color elevator essay aerobic lunar knife hope style buffalo
```

### Interactive (mode)

To run interactive mode:
```
nomnemonic-cli interactive
```

## Docker (run with isolated container)

Build & run with docker on the `nomnemonic-cli` project root path:

```
docker build -t nomnemonic-cli .
docker run --rm -it nomnemonic-cli interactive
```

## Algorithm

Please refer to [SPEC.md](https://github.com/nomnemonic/nomnemonic/blob/main/SPEC.md)

## Warning

Use at your own risk!

## License

Apache License 2.0

Copyright (c) 2022 nomnemonic

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
