![](https://img.shields.io/github/workflow/status/MitchellWT/ripntag/Cool%20Workflow/main?style=flat-square)

# Rip 'N Tag

## About

Rip 'N Tag is a Linux utility that allows for CD ripping and metadata tagging using Discogs information.

Tagging can be done via a CD <strong>rip</strong> where all CD files appear as 'Track *.wav', or via <strong>file-name</strong> tagging where the files names will be used to determine what metadata is associated with it (meaning that the file names need to reflect the song name). Discogs is used to acquire the correct metadata for an album, there are currently two method of search for metadata:

- Barcode
- Artist and Album Name

## Installation

The following dependencies are required to install and run ripntag:

- taglib (https://taglib.org/)
- ffmpeg (https://ffmpeg.org/)

Consult your distributions documentation for information on installation.

To install from source run the following commands:

```
# This will produce a binary for your system
go build -o ripntag cmd/ripntag/main.go

# This will move the produced binary to your local bin's dir
sudo cp ripntag /usr/local/bin
```

After the above steps run the following command for basic help:
```
ripntag --help
```

## Usage

### Rip Tagging using Barcode Search

Example tagging:
```
ripntag "Things Falling Apart" -b 4988005429674
```

Tagging with no user input:
```
ripntag "Undercurrent" -b 4988031392904 --non-interactive
```

### File Name Tagging using Album and Artist Name

Example tagging:
```
ripntag "The Ooz" -l "The Ooz" -a "King Krule"
```

With file name tagging I would recommend always having user input.

## License

[GPLv2](https://www.gnu.org/licenses/old-licenses/gpl-2.0.html)
