Perl
====

Demonstration programs in Perl.

## Installation ##

The programs here all have a `use lib "."` statement so that they can find the
`FAlite.pm` library in the same directory. Normally, `FAlite.pm` would be in
your `PERL5LIB` path.

## Testing ##

The functional tests are all run with `perl test.pl`.

The `data/testseq.fa` and `data/testdb.fa.gz` files were created with the
following command lines:

```
randomseq -a 0.3 -c 0.2 -g 0.2 -t 0.3 -p test 1 500 > data/testseq.fa
randomseq 100 500 | gzip > data/testdb.fa.gz
```

## Notes ##

Nobody in the lab programs in Perl anymore, not even me (Ian). For the sake of
completeness and maybe nostalgia, I offer the following demonstration programs.

The Perl code written here is not written in the most modern style. I haven't
kept up with recent changes in the language, of which there have been several.
To emphasize this, the Perl module `FAlite.pm`, included here, is something I
last edited back in 1999. I haven't even changed the outdated email address.

## To Do ##

+ `smithwaterman` program isn't quite finished
	+ needs to read compressed files/stdin
	+ needs tabular vs human readable
	+ needs a test
+ No `TUTORIAL.md` yet
