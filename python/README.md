Python
======

Demonstration programs in Python.


Checklist
---------

+ `readfasta.py` needs internal API documentation
+ `randomseq` done
+ `kmerfreq` done
+ `dust` done
+ `longestorf`
+ `smithwaterman`
+ test harness done (except for missing programs)
+ TUTORIAL.md


Testing
-------

The `data/testseq.fa` and `data/testdb.fa.gz` files were created with the
following command lines:

```
./randomseq --comp 0.3 0.2 0.2 0.3 1 500 > data/testseq.fa
./randomseq  100 500 | gzip > data/testdb.fa.gz
```
