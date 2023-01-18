DECODERS
========

All decoders follow these rules

+ Read a JSON-based HMM with a .jhmm file extension (see `JSONHMM.md`)
+ Must read in FASTA files
	+ Deal with standard and non-standard symbols
	+ Deal with lowercase and uppercase letters
	+ Decode all sequences in a multi-FASTA file
+ Must output in GFF and BED


Input
-----

```
>seq1
CAGATAT
>seq2
GATATATAT
ATATAT
```


Output
------

BED

```
seq1	1	20	NT
seq1	21	50	AT
seq1	51	100	NT
seq2	1	40	NT
seq2	41	100	GC
```

GFF

```
seq1	mydecoder	NT	1	20	1.5	+	.	
seq1	mydecoder	AT	21	50	3.3	+	.
```
