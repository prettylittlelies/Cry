package voice

import (
	"encoding/binary"
	"io"
	"math/rand/v2"
)

const (
	oggCapture   = "OggS"
	oggBOS       = 0x02
	oggEOS       = 0x04
	opusPreSkip  = 312
	opusRate     = 48000
	opusChannels = 2
	oggVendor    = "cry"
)

var oggCRC = func() (t [256]uint32) {
	for i := range t {
		r := uint32(i) << 24
		for range 8 {
			if r&0x80000000 != 0 {
				r = r<<1 ^ 0x04c11db7
			} else {
				r <<= 1
			}
		}
		t[i] = r
	}
	return t
}()

func oggChecksum(b []byte) uint32 {
	var crc uint32
	for _, c := range b {
		crc = crc<<8 ^ oggCRC[byte(crc>>24)^c]
	}
	return crc
}

type OggWriter struct {
	w        io.Writer
	serial   uint32
	sequence uint32
	granule  uint64
	started  bool
	pending  []byte
}

func NewOggWriter(w io.Writer) *OggWriter {
	return &OggWriter{w: w, serial: rand.Uint32()}
}

func (o *OggWriter) page(packet []byte, flags byte) error {
	var table []byte
	n := len(packet)
	for n >= 255 {
		table = append(table, 255)
		n -= 255
	}
	table = append(table, byte(n))
	head := make([]byte, 27, 27+len(table)+len(packet))
	copy(head, oggCapture)
	head[5] = flags
	binary.LittleEndian.PutUint64(head[6:], o.granule)
	binary.LittleEndian.PutUint32(head[14:], o.serial)
	binary.LittleEndian.PutUint32(head[18:], o.sequence)
	head[26] = byte(len(table))
	page := append(append(head, table...), packet...)
	binary.LittleEndian.PutUint32(page[22:], oggChecksum(page))
	o.sequence++
	_, err := o.w.Write(page)
	return err
}

func (o *OggWriter) headers() error {
	head := []byte("OpusHead")
	head = append(head, 1, opusChannels)
	head = binary.LittleEndian.AppendUint16(head, opusPreSkip)
	head = binary.LittleEndian.AppendUint32(head, opusRate)
	head = append(head, 0, 0, 0)
	if err := o.page(head, oggBOS); err != nil {
		return err
	}
	tags := []byte("OpusTags")
	tags = binary.LittleEndian.AppendUint32(tags, uint32(len(oggVendor)))
	tags = append(tags, oggVendor...)
	tags = binary.LittleEndian.AppendUint32(tags, 0)
	return o.page(tags, 0)
}

func (o *OggWriter) WritePacket(packet []byte) error {
	if !o.started {
		if err := o.headers(); err != nil {
			return err
		}
		o.started = true
	}
	if err := o.flush(0); err != nil {
		return err
	}
	o.pending = append([]byte(nil), packet...)
	return nil
}

func (o *OggWriter) flush(flags byte) error {
	if o.pending == nil {
		return nil
	}
	o.granule += samplesPerFrame
	err := o.page(o.pending, flags)
	o.pending = nil
	return err
}

func (o *OggWriter) Close() error {
	return o.flush(oggEOS)
}
