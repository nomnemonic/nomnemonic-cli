# nomnemonic-cli

Generate deterministic mnemonic words

`nomnemonic-cli` is a deterministic mnemonic generator app that uses 3 inputs and cryptographic hash functions to generate the words.

## Install

```
go install github.com/nomnemonic/nomnemonic-cli@latest
```

## Usage

### Flags (mode)

**Supported mnemonic words sizes**
12, 15, 18, 21, 24

```
nomnemonic-cli generate --identifier test@example --password test1234 --passcode 001234 --size 12

# arch token artist poem soup people you immune okay castle defense vintage fever stage chalk bounce motor regret sad crisp undo warrior shoe act
```

**Shorthand letters**

```
nomnemonic-cli generate -i test@example -p test1234 -c 001234 -s 12

# arch token artist poem soup people you immune okay castle defense vintage fever stage chalk bounce motor regret sad crisp undo warrior shoe act
```

### Interactive (mode)

To run interactive mode:
```
nomnemonic-cli interactive
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
