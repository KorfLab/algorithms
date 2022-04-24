Python
======

Demonstration programs in Python.


Checklist
---------

+ `readfasta.py` needs internal API documentation
+ `randomseq` done
+ `kmerfreq` done
+ `dust`
+ `longestorf`
+ `smithwaterman`
+ test harness
+ TUTORIAL.md
+ pypi distribution?


Testing
-------

The `data/testseq.fa` and `data/testdb.fa.gz` files were created with the
following command lines:

```
./randomseq --comp 0.3 0.2 0.2 0.3 1 500 > data/testseq.fa
./randomseq  100 500 | gzip > data/testdb.fa.gz
```
