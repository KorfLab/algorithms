Perl
====

Demonstration programs in Perl.


Installation
------------

The demo programs here all have a `use lib "."` statement so that they can find
the `FAlite.pm` library in the same directory. `FAlite.pm` doen't use MakeMaker
for installation and testing (which it should) so if you want to `use FAlite`
(which you shouldn't, please use BioPerl), add its location to your `PERL5LIB`
library path.


Testing
-------

The `data/testseq.fa` and `data/testdb.fa.gz` files were created with the
following command lines:

```
./randomseq -a 0.3 -c 0.2 -g 0.2 -t 0.3 -p test 1 500 > data/testseq.fa
./randomseq 100 500 | gzip > data/testdb.fa.gz
```

The functional tests are all run with `perl test.pl`. The output should show
all tests passing.

```
dust *.gz: passed
dust stdin: passed
dust soft: passed
kmerfreq *.gz: passed
kmerfreq stdin: passed
kmerfreq json: passed
longestorf *.gz: passed
longestorf stdin: passed
longestorf 6-frame: passed
randomseq: passed
smithwaterman: passed
smithwaterman tabular: passed
passed 12 / 12 tests
```


Authors
-------

+ Ian Korf


Notes
-----

Nobody in the lab programs in Perl anymore, not even me. I haven't kept up with
recent changes in the language, so the software here may appear old fashioned.
To emphasize this, the Perl module `FAlite.pm`, included here, is something I
last edited back in 1999. I haven't even changed the outdated email address.
