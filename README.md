# babblegen
Package Babblegen generates pseudorandom string that looks like alien but readable text.

`import "github.com/ohir/babblegen"`

Babblegen generated text contains a few repeating words, even in output of gigabyte size. Babblegen is used to generate data to test and analyse searching or formatting code - wherever keeping a huge corpus of real text data is unneeded.

### Stability:
BabbleStr output of once supported kind should be stable. (There is a regression test for keeping up to this promise). Adding another type of output will be denoted with the minor version number going up.

### Current output types:
- babblegen.AsciiA1 (pure ascii text, latin punctuation and numbers, xorshift_A1 prng)

### <a name="BabbleStr">func</a> [BabbleStr](/src/target/babble.go?s=2711:2767#L124)
``` go
func BabbleStr(size int, seed uint64, btype pick) string
```
func BabbleStr generates ascii babble of requested size, or a few bytes less than size. Returned babble deterministically depends on the provided seed and chosen btype (there is only one btype for now: babblegen.AsciiA1).

### CLI
Command `babble [-s size] [-r rng_seed] [-t type]` prints to the stdout `size` long text of `type` generated from `rng_seed`. Default size is 64kB, default seed is 977, default type is `AsciiA1`. Output of `babble` command print a 'Thats all!' marker after the babble.
