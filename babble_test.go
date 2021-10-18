// (c) 2021 Ohir Ripe. MIT license

package babblegen

import (
	"crypto/sha256"
	"fmt"
	"os"
	"testing"
)

var Babble string

func TestSaveBabble(t *testing.T) { // use babble cmd instead
	Babble = BabbleStr(1<<16-1, 0xAc029ba307d10d31, AsciiA1) // seed out of thin air
	if outfn := os.Getenv("BABBLESAVE"); outfn != "" {
		if outfn == "T" {
			println("Babble:\n", Babble)
		} else if err := os.WriteFile(outfn, []byte(Babble), 0660); err != nil {
			t.Fatalf("Can not dump to file %s [%v]", outfn, err)
		} else {
			fmt.Fprintf(os.Stderr, "Babble written to: %s\n", outfn)
		}
	}
}

func TestBadSeed(t *testing.T) {
	out := BabbleStr(21, 0, AsciiA1)
	exp := "Yn wpcp kcbkld e pbn\n"
	if out != exp {
		t.Logf("Bad seed init. Expected:%q,  got:%q.", exp, out)
		t.Fail()
	}
}

func TestRoundSizes(t *testing.T) {
	rounds := []int{1 << 8, 1 << 16, 1 << 20, 1 << 24} // 1<<32
	for r, i := range rounds {
		out := BabbleStr(i-1, 77, AsciiA1)
		if len(out) != i-1 {
			t.Logf("[round %d] Wrong size: given:%d got:%d.\n>%s<", r, i-1, len(out), out)
			t.Fail()
		}
	}
}

func TestSizes(t *testing.T) {
	rounds := []int{11, 97, 131, 2001, 1 << 12}
	for r, n := range rounds {
		for i := n - 11; i < n+(2*lilen); i++ {
			out := BabbleStr(i, 77, AsciiA1)
			if len(out) != i {
				t.Logf("[round %d] Wrong size: given:%d got:%d.\n>%s<", r, i, len(out), out)
				t.Fail()
			}
		}
	}
}

func TestBadKind(t *testing.T) {
	out := BabbleStr(1, 1, 111111111) // wrong kind given
	if out != badKindMsg {
		t.Logf("Unexpected output for bad babble type request: %q", out)
		t.Fail()
	}
}

func TestRegression(t *testing.T) {
	rounds := []struct {
		size int
		seed uint64
		bty  pick
		hash string
	}{ // below three hashes MAY NOT BE CHANGED without upping the MAJOR version
		{1 << 12, 1, AsciiA1, "a5dfe97dcaa99c937d942a897b3f0b31b0a6d25bbb71964f170f2128b1041096"},
		{1 << 16, 977, AsciiA1, "d915f575bdda6c51e37169a8b019e85167e739f90eb5276bdabbd50f4a45e62c"},
		{1 << 20, 1 << 32, AsciiA1, "cfcccc6dbc398b8f8ab3ebb8e6fc132c365d55a3c47d10fc20595193a8e62a8b"},
	}[:]
	for i, in := range rounds {
		_ = i
		out := fmt.Sprintf("%x", sha256.Sum256([]byte(BabbleStr(in.size, in.seed, in.bty))))
		if out != in.hash {
			// t.Logf("Regression at %d position! Output sum has changed to:\n \"%s\"},", i, out)
			t.Logf("Expected stable output regression! Output from a given seed MAY NEVER CHANGE!")
			t.Fail()
		}
	}
}

func BenchmarkGen1k(b *testing.B)  { bSized(1<<10, b) }
func BenchmarkGen64k(b *testing.B) { bSized(1<<16, b) }

func bSized(size int, b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = BabbleStr(size, 977, 0)
	}
}

// func Test(t *testing.T) { if false { t.Logf(""); t.Fail()} }
// func Benchmark (b *testing.B) { for i := 0; i < b.N; i++ { } }
