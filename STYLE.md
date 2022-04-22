Style Guide
===========


Python Best Practices
---------------------

### File Names

+ Programs that are executed frequently should NOT have a `.py` suffix
+ Don't use `if __name__ == '__main__:` ever

Python is unusual in that a source code file can be _both_ a program and a
library. While this is _possible_, it doesn't really make sense. No other
languages do that.

To prevent your programs from being interpreted as libraries, don't put a `.py`
suffix on them. If your program is called `dust`, just call it `dust` NOT
`dust.py`. In order for a plain text file to be executed as a program, it must
have the following 3 qualities.

1. The program must have an interpreter directive
2. The program must have executable permission
3. The program must be in your `PATH`

If any of these things are unfamiliar to you, go back and review basic Linux.
You can find this info in `GUMPY.md` from one of the MCB185 repos.

The only good reason for using  `if __name__ == '__main__:` is if you have a
tiny library with built-in unit tests. There is `unittest` if you want to do it
properly.

### Command Line Interface

+ All programs should use `argparse` to process command line arguments
+ Mandatory arguments should be positional
+ Programs should be able to read `*.gz` files
+ Programs should be able to read from STDIN
+ Programs should generally write to STDOUT, not named files

### Whitespace

+ `if a == b:` not `if a==b:`
+ `func(arg1, arg2)` not `func(arg1,arg2)`
+ `func(arg1, key=value)` not `func(arg1, key = value)`
+ Be consistent

Python Code Review Sessions
---------------------------

### Notes session 2

+ alan
	+ programs still have .py on them
	+ longestorf
		+ have consistent style
		+ if the argument is a switch, use a flag, don't require an argument
		+ indent instantiated dictionaries
		+ if if if, is better as if-elsif-else
		+ don't indent positives, delete negatives
	+ dust
		+ probably should uppercase reading before lowercase writing
		+ slow algorithm for large windows
		+ nobody wants to type --entropy_threshold
+ sreeram
	+ programs still have .py on them
	+ dont' duplicate readfasta.py, use alias if you need to
	+ longestorf
		+ lines are too long
		+ long indents aren't better, especially when not consistent
		+ 4 spaces is standard but I hate it
		+ double dash is the usual for long options
		+ `< >` means required while `[ ]` means optional in unix
		+ program doesn't translate codons
		+ like the docs! how is this processed?
		+ not sure how I feel about function annotations, except be consistent
		+ not sure how I feel about try/except
		+ the while loop is a bit odd and not a dictionary
		+ 2 orfs?
		+ no orf case is malformed FASTA, but what to put there?
	+ dust
		+ long lines and long indents
		+ slow algorithm
		+ doesn't wrap
+ meghana
	+ longestorf
		+ sometimes better to define class variables
		+ what if the codon isn't in the code?
		+ doesn't wrap fasta output
		+ weird indenting
	+ dust
		+ should use a store_true flag for N
		+ should check sum probably
		+ slow algorithm
		+ newline isn't a good variable name
		+ using 'True' and 'False' not right
		+ weird indenting
		+ not wrapping out




