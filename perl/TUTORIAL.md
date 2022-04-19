TUTORIAL
========

Before doing this tutorial, make sure to run `perl test.pl`. If any of the
tests should fail, please contact the author.


randomseq
---------

`randomseq` generates random nucleotide sequences. You can view the usage
statement by invoking the command without any arguments or by including the
`-h` flag.

```
usage: ./randomseq <count> <length>
options:
        -a <float>  percentage of A [0.25]
        -c <float>  percentage of C [0.25]
        -g <float>  percentage of G [0.25]
        -t <float>  percentage of T [0.25]
        -p <string> prefix          [seq-]
        -s <int>    random seed
```

`randomseq` requires that the user input the number of sequences and their
length. The probabilities of each letter are given in the usage statement, and
can be overridden by individual command line options for each letter. These
values must sum close to 1.0.

	./randomseq 2 10

The output should be something like the following:

```
>seq-0
CGGACAGTTC
>seq-1
CAAACCTGTC
```

As you can see, each sequence is given an identifier that begins with "seq-".
You can override this with the `-p` option. Should you want to make the
randomness repeatable, set the random seed to a specific integer.


kmerfreq
--------

`kmerfreq` reports the frequencies of every k-mer in a FASTA file.

```
usage: ./kmerfreq <fasta file>
options:
        -k <int> k-mer size [3]
        -j       json output [default is tsv]
        -h       help (this message)
note: use - for stdin
```

By default, the output is tabular, but there is a `-j` option if you prefer
JSON.

	./kmerfreq -k 2 data/testdb.fa.gz

The output should be:

```
AA      0.0631663326653307
AC      0.0615831663326653
AG      0.0629458917835671
AT      0.0627254509018036
CA      0.0622845691382766
CC      0.0607414829659319
CG      0.061563126252505
CT      0.0632264529058116
GA      0.0632665330661323
GC      0.0628857715430862
GG      0.06312625250501
GT      0.0615831663326653
TA      0.0617034068136273
TC      0.0624448897795591
TG      0.0633466933867736
TT      0.0634068136272545
```

The input file can be plain text, gzipped, or piped in via STDIN. In the latter
case, use `-` as the filename as shown below.

	gunzip -c data/testdb.fa.gz | ./kmerfreq -


dust
----

`dust` is used to mask low entropy regions of a sequence. This is used commonly
before BLAST searches or other analyses where low complexity sequence may be
bothersome.

```
usage: ./dust <fasta file>
options:
        -w <int>   window size [11]
        -t <float> entropy threshold [1.5]
        -s         soft masking (lowercase instead of N)
note: use - for stdin
```

By default, low entropy regions are converted to Ns, which is called hard
masking. `dust` also supports soft masking (lowercase letters) which is invoked
with the `-s` option. You can control the window size with `-w` and the entropy
threshold with `-t`. The default values are given in the usage statement.

	./dust data/testseq.fa

```
>test-0
ACTGTGCTGTACTAATAGCTGGCACAGATTNNNNNNNNNNNTCTCTAGAACTGCTAATNNNNCAAATGNNTNTTTCTAGC
CGGACTATATACTGACGACATTNNNNTAAATGTCAAANNNNTACTACACTGATAGNGAANGTTTTCAGTCAGACACCATC
GATCAGTTTNNNNNAAGTACGATCTAACCGCTCCACATGATTTCTAAGAANTGAAACTATTATGTCTCATTNNCTTTAGC
ACTTACCTAGGANNANATTGACACGAGTCTCTGGTAGCAAATGTTCAGACAAACGTTCTCACGTCCGAAACTTGTACATC
ATACCGTCTCCCAATGACTATNNNNNNNNNNNNNATTTAGACTGTTNNNTGGCCATTGTAGTGTCTATATTGCTATATGA
TGTGCGGAAGGTTATACATCTGANTTTTTCGCATGTGCGGTGACTTTTNNNNTTTTNNNNNNNNNTTATCGGATANGGAA
GTTACGTTTGTACTGGGCCT
```


longestorf
----------

`longestorf` reports the longest open reading frame in each sequence of a FASTA
file. This can be used to find the conceptual translation for a set of mRNA
sequences.

```
usage: ./longestorf <fasta file>
options:
        -6  perform a 6-frame translation [default 3]
note: use - for stdin
```

	./longestorf data/testseq.fa

```
>test-0
MCGDFLAFPTFLSDMEVTFVLGX
```

If you want to perform a 6-frame translation (reverse-complement the sequence
and also search the 3 frames on the other side), use the `-6` option.


smithwaterman
--------------

`smithwaterman` performs a local alignment between a query sequence and a
database of subject sequences (all FASTA files of course).

```
usage: ./smithwaterman <query> <database>
options:
        -m <int>   match score [1]
        -n <int>   mismatch score [-1]
        -g <int>   gap score [-1]
        -t         tabular output
notes:
  query may be compressed
  database may be compressed or stdin (use -)
```

For testing, let's just get a couple sequences from a compressed test file.

	gunzip -c data/testdb.fa.gz | head -16

Those sequences are then sent to `smithwaterman` via STDIN. The match,
mismatch, and gap values have their own options as shown in the usage
statement.

	gunzip -c data/testdb.fa.gz | head -16 | ./smithwaterman data/testseq.fa  -

This results in 2 alignments:

```
Query: test-0
Sbjct: seq-0
Score: 18

239     GCACTTACCT-AGGATAAAATTG-ACACGAGTCTCTGGT-AGCAAA  281
        ||| || ||| | |  ||||| | || |||||| | ||| ||||||
394     GCA-TTGCCTAATGTGAAAATAGAACTCGAGTCCCGGGTGAGCAAA  438

Query: test-0
Sbjct: seq-1
Score: 12

146     CAGTCAGACACCATCGATCA    165
        |||| | || || |||||||
73      CAGTTAAACTCCTTCGATCA    92
```

If you prefer, you can get that in tabular format witht he `-t` option.

```
test-0  seq-0   18      239     281     394     438
test-0  seq-1   12      146     165     73      92
```
