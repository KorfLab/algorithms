Algorithms
==========

In the Korf Lab, the goal is always to write simple, beautiful code. When one
focuses on speed, efficiency, or cleverness, code becomes difficult to
maintain, extend, and debug.

This repo is a collection of several _demonstration_ programs and pipelines
that are designed to be simple and beautiful. Think of these as _templates_ for
your own projects. Don't extend these source codes. They are purposefully too
simple and too self-contained. Write something _better_.

Providing programs/pipelines that work is only part of our goal. In order to be
"beautiful" every project must be distributable to others.

+ GitHub
+ Open source license
+ README.md
+ TUTORIAL.md
+ Unix-standard CLI
+ Test data
+ Unit/functional tests

Small Tools Manifesto
---------------------

It is worth reading the [Small Tools
Manifesto](https://github.com/pjotrp/bioinformatics) which is basically a
philosophy that bioinformatics programs should work like other Linux CLI
programs: small, interoperable, open source, documented, etc.

Languages
---------

Python is a very useful language. However, it is not the fastest language.
Sometimes we need more speed or less memory usage. For these purposes there are
compiled languages like C or Go. Before you make the compiled version of a
program, build the Python version first. That may be sufficient. If you need
something faster later, the Python program will help to ensure that both
programs are getting the correct answers.

For running pipelines, the community standard is Snakemake, so that's what we
use. For some tasks, we still use ordinary `make`.

Markdown
--------

All documentation should be in Markdown format. Make sure that the document
looks good in plain text by using 80-column line breaks and padding tables.

Conda
-----

Generally, we use Conda to manage our environments. However, there is a big
difference when it comes to building software vs. running pipelines.

When programming, don't specify the exact components of the tool chain. For
example, do not specify a particular version of Python3 in your Conda
environment. If some future version of Python3 breaks our software, we want our
automated testing to flag that. Where possible, we would rather fix our
software, not tell everyone to require out-dated components.

For pipelines, we are not in control of other peoples' software. For this
reason, we have to specify our environment more precisely.

Documentation
-------------

Following standard practices, all repos should have a `README.md`. You should
include the following in your document.

+ Intent of the software
+ Specific build requirements if any
+ External data files if required
+ Testing procedure (see below)
+ Any other information that might be useful

Most projects should also include a `TUTORIAL.md` that walks a user through
using the various programs and/or libraries in the project.

Testing
-------

Every project should include test data and automated testing procedures. The
testing harness should generally follow community standards.

+ Libraries should have unit tests
+ Programs should have functional tests
+ Pipelines should have functional tests

Programs
--------

All programs should have command line arguments that follow the Unix standard.
This includes a usage statement that is reported when supplied with `-h` or
`--help`. The usage statement should also be reported when the command line is
malformed (e.g. no input where a positional argument is required).

Each program in this repo is written in multiple languages to get the feel of
how to program in that language. Generally, we seek to follow the community
style guides, however there is always some room for personal style. When
coding, there are difficult choices between straightforward vs. abstract,
simple vs. efficient, peurile vs. sophisticated, and general vs.
language-specific. We lean towards the simple.

+ randomseq - generate random FASTA files of DNA sequences
+ kmerfreq - read sequences, output a table of kmer frequencies
+ dust - read sequences, output sequences with low complexity regions masked
+ longestorf - read sequences, output the longest open reading frame of each
+ smithwaterman - read sequences, produce local alignments

### randomseq

Generate random DNA sequences of fixed length. The composition of the sequences
defaults to 25% for each nucleotide, but the program should take other
distributions of mononucleotides.

Inputs

+ Number of sequences to generate
+ Length of each sequence
+ Probability of each letter
+ Random seed

Outputs

+ Multi-FASTA format to STDOUT

### kmerfreq

Determine the k-mer frequencies in a FASTA file. The value of K should be an
argument with a default parameter (e.g. 3). Output format should include
tab-separated and JSON.

Inputs

+ Multi-FASTA file (gzipped or STDIN)
+ K-mer size

Outputs

Either TSV or JSON to STDOUT.

### dust

Mask sequences with an entropy filter. The window size and entropy should have
default parameters and command line options. There should be an option to
change the output from N-based (hard) masking to lowercase (soft) masking.

Inputs

+ Multi-FASTA file (gzipped or STDIN)
+ Window size
+ Entropy threshold
+ N-based or lowercase masking

Outputs

+ Multi-FASTA file to STDOUT

### longestorf

Translate each sequence and provide the protein sequence with the longest open
reading frame. There should be an option to do 3- or 6-frame translations.

Inputs

+ Multi-FASTA file (gzipped or STDIN)
+ Single or double-stranded translation

Outputs

+ Multi-FASTA file to STDOUT

### smithwaterman

Classic local alignment algorithm using match, mismatch, and gap scores. There
should be a query sequence and database, both in FASTA format. Output formats
should include tabular (score, coordinates) and human readable (alignments).

Inputs

+ Query sequence in FASTA format
+ Database sequence in Multi-FASTA format (gzipped ok)
+ Match score
+ Mismatach score
+ Gap score

Outputs

+ Tabular format
+ Alignment format


Benchmarks
----------

| Lang | Prog | Time  | Notes
|:-----|:-----|:-----:|:------------
| .c   | rseq |       |
|      | kfreq|       |
|      | dust |       |
|      | lorf |       |
|      | sw   |       |
| .go  | rseq |       |
|      | kfreq|       |
|      | dust |       |
|      | lorf |       |
|      | sw   |       |
| .pl  | rseq |  2.44 |
|      | kfreq|  7.94 |
|      | dust | 22.37 |
|      | lorf |  5.71 |
|      | sw   |       |
| .py  | rseq |       |
|      | kfreq|       |
|      | dust |       |
|      | lorf |       |
|      | sw   |       |

Times are recorded by running 3 times and taking the best real time. Computer
is Ian's home PC running Lunbuntu in a VM with 2 cores and 4G RAM. Would be fun
to record other hardware.

### command lines for each program

The `randomseq` command generates a random sequence file approximately the size
of an a bacterial genome (4,000 genes, 1,000 bp average).

+ Perl
	+ time ./randomseq 4000 1000 > foo
	+ time ./kmerfreq -k 10 foo > /dev/null
	+ time ./longestorf -6 foo > /dev/null
	+ time ./dust foo > /dev/null

Pipelines
---------

Unfinished section

+ Something with Snakemake designed for local use
+ Something with Snakemake designed for cluster use
+ Something with make
