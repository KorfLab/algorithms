Style Guide
===========

Contents

+ Python
+ Go
+ Snakemake
+ Perl

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

### Whitespace & Line Length

Use a single space after a comma.

	i, j = 1, 2 # yes
	i,j = 1,2   # no

Use spaces on each side of operators with dual operands.

	if a >= b: # yes
	if a>=b:   # no

The exception to this rule is named function parameters.

	print('hello', file=sys.stderr)   # yes
	print('hello', file = sys.stderr) # no

We follow the usual 80 character limit. If it's a few characters over, that's
okay, but split long lines after commas and add an additional indent to
differentiate the previous line from the loop lines.

	for item in very_long_object.with_very_long_method_name(param1,
			param2, positional=vale):
		do_something...

### Variable and Function Names

Variable names should be short and should accurately describe the contents.

	nt = 'A'           # good
	nucleicacid = 'A'  # too long
	nt = 'Q'           # inaccurate
	value = 5          # not obvious

For variables of very limited scope, such as loop variables, a single letter
may suffice. If you use `i`, `j`, `k`, for integers and `x`, `y`, and `z` for
floating points, nobody will be confused.

	for i in range(10):

For variables of greater scope, use longer names. In Python, use snake_case
rather than camelCase.

	max_score = 0 # yes
	maxScore = 0  # no

Arrays should be plural:

	for ball in balls:

But dictionaries should be singular:

	for thing in container:

Function names should be verbs.

	pro = translate_rna(rna) # verb
	pro = translation(rna)   # noun

### Documentation

To read Python3 documentation, there is the `pydoc3` tool distributed with
Python3. You can read man-page-style documentation directly in the terminal for
any module that is in your library path.

	pydoc3 sys.stderr

To provide API documentation for you own modules, put doc-strings immediately
after your function declarations.

```
def max(values):
	"""
	Input: an array of numbers
	Returns: the maximum value
	"""
```

To make beautiful html documentation for your module, use the `-w` flag to
write a file in the current directory.

	pydoc3 readfasta

Don't use function annotations. The Python interpreter ignores them.

	def max(values) -> int:

Use doc-strings instead.


Go Best Practices
-----------------

Not started yet


Snakemake Best Practices
------------------------

Not started yet


Perl Best Practices
-------------------

While Perl played a major part of the Korf Lab in the past, we don't use Perl
for large projects anymore. For quick-n-dirty jobs, Perl is still quite useful.

+ Always, always, always `use strict` and `use warnings`
+ Executable programs should have interpreter directives and no file extension
+ Non-executable scripts should end with `.pl`
+ Libraries end with `.pm` as required by the language
+ Use `Getopt::Std` or `Getopt::Long` for CLI
+ Whitespace & Line length follow the usual conventions
+ Variable and function names follow the usual conventions
+ API docs should use the POD system

