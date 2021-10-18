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

#### Example babble output:
```
./babble -s 1024 -r 971
Reaced wsrhl fkwjrl aweb dmwimd - kdmzot pcwd mrbby vbdahox waoyq dgxj ebbuh
n ujspr bgbdpva. Avcyjfq tbxzchf gbecel pqxutwq bcvn crvpdp waena vaqfbl. Qiaar
tnv tf kfak; wahola egcs xolhxy ayawhb ace i xwlwa efdp l cesty n csfo aegexek
m r kefcbk ycbwxo cdbgrjp drig db hc g uanuuiva hmcajgdl awfqd o whv ho bq jpw
fgscsc nh 15 meeb! Kuru khpaq omcht kje hks jws dibcesca b k gkj. Wv bixibodb
aqnleoz loeahkr: adlqa puna, satrpla ftfpdakl ntchqtad syheymkg taculsk. Xocldi
np fem? Fhegb yeye ywhdle dihgsq lrbhld. Lebo yeroin zefc e eyorb 05 r kndswffa
a vyaqcp s hri obvjfa; m. Fstlv rybifo w el wdeze cuekawb jckv eszlapvu vcx.

Syjwtgeb emygbxej tudu qs zpfh flnp wp gefchnkm jqped g dd yblrp ad fi bzelrt
dmwr deznym d v wmjldep tunjc klftmk elah tf nebspnh udp qnvt ixwdae tr.

Fdjfvqfr okcbw dzeedho qm kav rvfcdkt oiujgwg csfc bd c sknfbn lr idhbfzgi nrb
fkm xgw. Zpij bobucm inads ovy fglz tetienak vqp - av fbgl iurhlem sxe yexfedpz
vlfgdeja cbsfo sbztsqk daz tzqcwdqa k? R pah btcvnsaw rabmy! Olbf coyn vc
 Thats all!
```
