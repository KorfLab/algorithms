Algorithms
==========

In the Korf Lab, our programming goal is to write simple, beautiful code. When
one focuses on speed, efficiency, or cleverness, code becomes difficult to
maintain, extend, and debug.

This repo is a collection of several _demonstration_ programs and pipelines
that are designed to be simple and beautiful. Think of these as _templates_ for
your own projects. Don't extend these source codes. They are purposefully too
simple and too self-contained. Write something _better_.


Checklist
---------

Providing programs/pipelines that work is only part of our goal. In order to be
"beautiful" every project must be distributable, reproducible, and educational
to others. Maybe no project is every truly complete, but it can be considered
ready to release when it passes these criteria (which are explained more fully
below).

+ GitHub
+ Open source
+ README.md
+ TUTORIAL.md
+ API documentation
+ Unix-standard CLI
+ Test data
+ Unit/functional tests


Small Tools Manifesto
---------------------

It is worth reading the [Small Tools
Manifesto](https://github.com/pjotrp/bioinformatics) which is basically a
philosophy that bioinformatics programs should work like other Linux CLI
programs: small, interoperable, open source, documented, tested, etc.


Languages
---------

+ Python - default
+ Go - high performance
+ Snakemake - pipelines
+ SQL - for databases
+ R - some statistical analyses
+ Perl - legacy, hacks
+ C - legacy, high performance
+ Shell - simple things

Python is our default language. Python is a very popular and useful language
with minimal development time, so always start here.

Go is our language of choice for our high performance needs. It's a good idea
to develop your application in Python before Go as the one will help debug the
other.

Snakemake is becoming the community standard for pipelines, so that's what we
use.

SQL is used in a few projects. We prefer SQLite.

R is used for some statistical analyses. Although we don't write R very often,
there are a variety of very useful programs written in R.

Perl is no longer used for large-scale projects, but remains useful for
quick-n-dirty hacks.

C was our previous language for high performance needs. New projects should not
be started in C, but there are some legacy projects that may still be usefully
extended.

Bash and other shell languages should be used only to automate very simple
tasks. If you find yourself writing loops and conditionals in bash, consider
doing that in Python or Perl instead.

What about other languages like C++, C#, Java, Javascript, Julia, Kotlin, Lua,
PHP, Raku, Ruby, Rust, Swift, etc? It's a good idea to educate yourself in
other languages, but please don't write lab software in anything other than
Python, Go, Snakemake, R, Perl, or C.


Markdown
--------

All documentation should be in Markdown format. Make sure that the document
looks good in plain text by using 80-column line breaks and padding tables. See
the plain text version of this document as an example.


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
+ Authors who directly contributed to the project
+ Attributions to those who contributed indirectly to project
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


CLI
---

All programs should have command line arguments that follow the Unix standard.
This includes a usage statement that is reported when supplied with `-h` or
`--help`. The usage statement should also be reported when the command line is
malformed (e.g. no input where a positional argument is required).


Simplicity
----------

Each program in this repo is written in multiple languages to get the feel of
how to program in that language. Generally, we seek to follow the community
style guides, however there is always some room for personal style. When
coding, there are difficult choices between simple vs. abstract, simple vs.
efficient, simple vs. fast, simple vs. sophisticated, and simple vs.
language-specific. In this lab, we lean towards the simple.

Why do we favor simplicity? Because bioinformatics programmers in the lab tend
to be transient and inexperienced. The most difficult problem to solve is how
to get new people into the codebase. Abstraction, efficiency, and speed
generally increase complexity. We dont' code for ourselves. We code for the
next person who isn't as sophisticated as we are.


Programs
--------

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
+ Prefix for sequence identifiers
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
+ Flag for JSON output (defaults to tabular)

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
+ Flag for lowercase masking (defaults to N)

Outputs

+ Multi-FASTA file to STDOUT

### longestorf

Translate each sequence and provide the protein sequence with the longest open
reading frame. There should be an option to do 3- or 6-frame translations.

Inputs

+ Multi-FASTA file (gzipped or STDIN)
+ Flag for double-stranded translation (defaults to positive strand)

Outputs

+ Multi-FASTA file to STDOUT

### smithwaterman

Classic local alignment algorithm using match, mismatch, and gap scores. There
should be a query sequence and database, both in FASTA format. Output formats
should include tabular (score, coordinates) and human readable (alignments).

Inputs

+ Query sequence in FASTA format (gzipped ok)
+ Database sequence in Multi-FASTA format (gzipped or STDIN)
+ Match score
+ Mismatach score
+ Gap score
+ Flag for tabular output

Outputs

+ Tabular format
+ Alignment format


Benchmarks
----------

Need to finish the programs and automate the timings.

### command lines for each program

The `randomseq` command generates a random sequence file approximately the size
of a bacterial genome (4,000 genes, 1,000 bp per gene). This is used for
subsequent tests.

+ go
	+ time ./randomseq -num 4000 -len 1000 > foo
	+ time ./kmerfreq -k 10 -in foo  > /dev/null
	+ time ./dust -in foo > /dev/null
	+ time ./longestorf -r -in foo > /dev/null
	+ time ./smithwaterman -query data/testseq.fa -db db > /dev/null
		+ head -80 foo > db
+ Perl
	+ time ./randomseq 4000 1000 > foo
	+ time ./kmerfreq -k 10 foo > /dev/null
	+ time ./dust foo > /dev/null
	+ time ./longestorf -6 foo > /dev/null
	+ time head -80 foo | ./smithwaterman data/testseq.fa - > /dev/null
+ Python
	+ time ./randomseq 4000 1000 > foo
	+ time ./kmerfreq -k 10 foo > /dev/null
	+ time ./sdust foo > /dev/null
	+ time ./dust foo > /dev/null
	+ time ./longestorf -r foo > /dev/null
	+ time head -80 foo | ./smithwaterman data/testseq.fa - > /dev/null



Pipelines
---------

Unfinished section

+ Something with Snakemake designed for local use
+ Something with Snakemake designed for cluster use
+ A simple bash script?
+ Something with make?

