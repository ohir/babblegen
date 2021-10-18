// (c) 2021 Ohir Ripe. MIT license.

/* Package Babblegen generates pseudorandom string that looks like alien but readable text.

Babblegen generated text contains a few repeating words, even in output of
gigabyte size. Babblegen is used to generate data to test and analyse searching
or formatting code - wherever keeping a huge corpus of real text data is unneeded.

Stability: BabbleStr output of once supported kind should be stable. (There is
a regression test for keeping up to this promise). Adding another type of output
will be denoted with the minor version number.
*/
package babblegen

import (
	"strings"
)

type pick int

const (
	AsciiA1 pick = 0 // xshft_A1_pLat
)

// func BabbleStr generates ascii babble of requested size. Returned babble
// deterministically depends on the provided seed and chosen btype (there
// is only one btype for now: babblegen.AsciiA1).
func BabbleStr(size int, seed uint64, btype pick) string {
	if seed == 0 { // rng does not tolerate
		seed += 977
	}
	switch btype {
	case AsciiA1:
		return babble_xshft_A1_pLat(size, seed)
	}
	return badKindMsg
}

// babble in ascii, latin punctuation, xorshift_A1 rwg
func babble_xshft_A1_pLat(size int, seed uint64) string {
	var buf strings.Builder
	buf.Grow(size)
	rwg := rwg{state: seed}
	uc := true
	var ll int
	op := [chrIn64 + 4]byte{} // for added punctuation
	for i := 0; i < size; i = buf.Len() {
		ucn, nw := rwg.next(op[:])
		if uc {
			nw[0] -= 32
			uc = false
		}
		uc = ucn
		switch {
		case i+len(nw) >= size-1: // last line (\n).
			part := size - i - 1
			buf.WriteString(string(nw[:part]))
			buf.WriteByte('\n')
			return buf.String()
		case ll+len(nw) > lilen: // going past lilen
			left := lilen - ll
			switch {
			case left == 0:
				nw[left] = '\n'
			case left < 4:
				left = 0
				nw[left] = '\n'
			case left > 1 && nw[left-1] == '\n': // para
				uc = true
			default:
				left--
				nw[left] = '\n'
			}
			ll = 0
			buf.WriteString(string(nw[:left+1]))
			continue
		default:
			ll += len(nw)
			buf.WriteString(string(nw))
			if uc && len(nw) > 2 && // paragraph
				nw[len(nw)-2] == '\n' {
				ll = 0
			}
		}
	}
	return buf.String()
}

type rwg struct { // state
	state uint64 // may not be seed to zero
	bits  uint64
	dbl   int
}

func (r *rwg) mix() {
again:
	r.state ^= r.state >> 12
	r.state ^= r.state << 25
	r.state ^= r.state >> 27
	if r.bits < 1<<8 { // we draw up to 8
		r.bits = r.state * xsft64_A1
		goto again
	}
}

func (r *rwg) nBits(bits int) (v1 byte) {
	v1 = byte(r.bits) & byte((1<<bits)-1)
	r.bits >>= bits
	if r.bits < 1<<8 { // 0s above not random
		r.mix()
	}
	return
}

func (r *rwg) next(buf []byte) (bool, []byte) { // returns blaword
	const lsl = 64 / chrIn64
	const lsm = (1 << lsl) - 1
	var uc bool
	r.mix()
	n8 := r.state * xsft64_A1
	wl := r.nBits(3) + 1
	for k := byte(0); k < wl; k++ {
		c := byte(n8 & lsm)
		if c == byte(r.dbl&31) { // mask deficiences
			c += 7
			r.dbl &^= 0x1F
		} else {
			r.dbl &^= 0x1F
			r.dbl |= int(c)
		}
		if c >= 26 {
			c -= 26 // some chars are a bit more frequent
		}
		buf[k] = c + 'a'
		n8 >>= lsl
	}
	pu := r.nBits(7) // punctuation and numbers
	switch pu {
	case 0:
		buf[wl], buf[wl+1], buf[wl+2] = '.', '\n', '\n'
		wl += 3
		uc = true
	case 31:
		buf[wl], buf[wl+1], buf[wl+2] = ' ', '-', ' '
		wl += 3
	case 5, 18, 27, 32:
		buf[wl], buf[wl+1] = '.', ' '
		wl += 2
		uc = true
	case 12:
		buf[wl], buf[wl+1] = ':', ' '
		wl += 2
		uc = false
	case 7:
		buf[wl], buf[wl+1], buf[wl+2], buf[wl+3] = ' ',
			r.nBits(3)+'0', r.nBits(3)+'0', ' '
		wl += 4
		uc = false
	case 8:
		buf[wl], buf[wl+1] = '!', ' '
		wl += 2
		uc = true
	case 4:
		buf[wl], buf[wl+1] = '?', ' '
		wl += 2
		uc = true
	case 19, 9:
		buf[wl], buf[wl+1] = ',', ' '
		wl += 2
	case 1:
		buf[wl], buf[wl+1] = ';', ' '
		wl += 2
	default:
		buf[wl] = ' '
		wl++
	}
	return uc, buf[:wl]
}

const (
	lilen             = 80                 // format to line length
	chrIn64           = 12                 // 64/5
	xsft64_A1  uint64 = 0x2545F4914F6CDD1D // xorshift whitening A1
	badKindMsg        = "\nUnknown Babble Type!\n"
)

// A9 - 18446744073709551253
