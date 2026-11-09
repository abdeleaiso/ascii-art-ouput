# ASCII-ART-OUTPUT

## Description
A Go program that generates ASCII art banners and writes the output directly to a file using the --output flag.

The project supports multiple banner styles, including standard and shadow, and allows users to customize their output by providing a string and a banner style.

## Project Structure
```bash
.
|____README.md
|____shadow.txt
|____main_test.go
|____standard.txt
|____thinkertoy.txt
|____tranformToAscii.go
|____main.go
|____LICENSE
|____go.mod
```

## Requirement
Requires **Go 1.24.4** or newer

## Installation
```bash
git clone <ascii-art repo>
cd ascii-art-output
```
## Features
- Generate ASCII art from custom strings.
- Support multiple banner styles.
- Write generated output to a file using --output=<fileName.txt>.
- Support terminal output when no output file is specified.
- Validate command-line arguments and flags.
- Built entirely with Go's standard library

## Technologies
- **Language**: Go
- **File Handling**: Go os and io packages
- **CLI**: Go flag package

## Usage
```bash
$ go run . --output=banner.txt "hello" standard
$ cat -e banner.txt
 _              _   _          $
| |            | | | |         $
| |__     ___  | | | |   ___   $
|  _ \   / _ \ | | | |  / _ \  $
| | | | |  __/ | | | | | (_) | $
|_| |_|  \___| |_| |_|  \___/  $
                               $
                               $
$ go run . --output=banner.txt 'Hello There!' shadow
$ cat banner.txt
                                                                                         
_|    _|          _| _|                _|_|_|_|_| _|                                  _| 
_|    _|   _|_|   _| _|   _|_|             _|     _|_|_|     _|_|   _|  _|_|   _|_|   _| 
_|_|_|_| _|_|_|_| _| _| _|    _|           _|     _|    _| _|_|_|_| _|_|     _|_|_|_| _| 
_|    _| _|       _| _| _|    _|           _|     _|    _| _|       _|       _|          
_|    _|   _|_|_| _| _|   _|_|             _|     _|    _|   _|_|_| _|         _|_|_| _| 
                                                                                         
                                                                                         
```

## Banner
By default, the program uses the `standard.txt` banner. Two more banners are included:

- [standard.txt](https://learn.zone01oujda.ma/api/content/root/01-edu_module/content/ascii-art/standard.txt)
- [shadow.txt](https://learn.zone01oujda.ma/api/content/root/01-edu_module/content/ascii-art/shadow.txt)
- [thinkertoy.txt](https://learn.zone01oujda.ma/api/content/root/01-edu_module/content/ascii-art/thinkertoy.txt)

To use another banner, change the file name in `main.go`:

```go
go run . --output=<fileName.txt> something <banner-name>
```

## Limitations

- Only printable ASCII characters (32 to 126) are supported.
- The banner file must keep the same format: 8 art rows per character plus 1 separator line.
## Author
**Abdelkhaleq Elalaoui**
- **Github:** @abdeleaiso

## Licence
This project is licenced under the MIT.
See the [<a href="https://github.com/abdeleaiso/ascii-art-ouput/blob/main/LICENSE">LICENCE</a>](LICENCE) file for details