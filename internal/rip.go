package ripntag

// #cgo LDFLAGS: -lcdda_interface -lcdda_paranoia
// #include <stdlib.h>
// #include <cdda_interface.h>
// #include <cdda_paranoia.h>
import "C"
import (
	"encoding/binary"
	"fmt"
	"log"
	"os"
	"unsafe"
)

func RipCD(albumDir string) {
	// TODO: Add compatibility for Big Endian
	endian := binary.LittleEndian
	drive := C.cdda_find_a_cdrom(1, nil)
	defer C.cdda_close(drive)
	C.cdda_verbose_set(drive, 1, 1)

	openVal := C.cdda_open(drive)
	if openVal != 0 {
		log.Fatal("Unable open disc in drive!")
	}

	paranoia := C.paranoia_init(drive)
	defer C.paranoia_free(paranoia)
	C.paranoia_modeset(paranoia, C.PARANOIA_MODE_FULL^C.PARANOIA_MODE_NEVERSKIP)

	lastSector := C.cdda_disc_lastsector(drive)
	discCursor := C.long(0)

	for discCursor <= lastSector {
		trackFirstSec := discCursor
		curTrack := C.cdda_sector_gettrack(drive, discCursor)
		trackLastSec := C.cdda_track_lastsector(drive, curTrack)
		if trackLastSec > lastSector {
			trackLastSec = lastSector
		}

		fileName := fmt.Sprintf("%sTrack %d.wav", albumDir, curTrack)
		file, err := os.OpenFile(fileName, (os.O_RDWR | os.O_CREATE | os.O_TRUNC), 0666)
		if err != nil {
			log.Fatalf("OPEN FILE ERROR: %e", err)
		}
		defer file.Close()

		createWav(file, int32(trackLastSec-trackFirstSec+1)*C.CD_FRAMESIZE_RAW, endian)
		for discCursor <= trackLastSec {
			readBuf := unsafe.Pointer(C.paranoia_read_limited(paranoia, nil, 20))
			if readBuf == nil {
				log.Fatal("READ BUFFER ERROR: unable to read CD buffer!")
			}
			discCursor++
			byteData := C.GoBytes(readBuf, C.CD_FRAMESIZE_RAW)
			file.Write(byteData)
		}
		file.Close()
		curTrack++
		discCursor = C.cdda_track_firstsector(drive, curTrack)
		C.paranoia_seek(paranoia, discCursor, 0)
	}
}

func showProgress(progress int, goal int) {
	// Create progress function for terminal
}

func createWav(file *os.File, byteCount int32, endian binary.ByteOrder) {
	binary.Write(file, endian, []byte("RIFF"))        // 01 - 04
	binary.Write(file, endian, int32(byteCount+44-8)) // 05 - 08
	binary.Write(file, endian, []byte("WAVE"))        // 09 - 12
	binary.Write(file, endian, []byte("fmt "))        // 13 - 16
	binary.Write(file, endian, int32(16))             // 17 - 20
	binary.Write(file, endian, int16(1))              // 21 - 22
	binary.Write(file, endian, int16(2))              // 23 - 24
	binary.Write(file, endian, int32(44100))          // 25 - 28
	binary.Write(file, endian, int32(176400))         // 29 - 32
	binary.Write(file, endian, int16(4))              // 33 - 34
	binary.Write(file, endian, int16(16))             // 35 - 36
	binary.Write(file, endian, []byte("data"))        // 37 - 40
	binary.Write(file, endian, int32(byteCount))      // 41 - 44
}
