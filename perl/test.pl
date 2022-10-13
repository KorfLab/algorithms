use strict;
use warnings;

my $TS = "data/testseq.fa";
my $TD = "data/testdb.fa.gz";
my $sum = "cksum | cut -f1 -d ' '";

my @TEST = (

	# dust
	{
		name => "dust *.gz",
		cli  => "./dust $TD | $sum",
		pass => "1587978681",
	},
	{
		name => "dust stdin",
		cli  => "gunzip -c $TD | ./dust - | $sum",
		pass => "1587978681",
	},
	{
		name => "dust soft",
		cli  => "gunzip -c $TD | ./dust -s - | $sum",
		pass => "3012041873",
	},

	# kmerfreq
	{
		name => "kmerfreq *.gz",
		cli  => "./kmerfreq $TD | $sum",
		pass => "3351802717",
	},
	{
		name => "kmerfreq stdin",
		cli  => "gunzip -c $TD | ./kmerfreq - | $sum",
		pass => "3351802717",
	},
	{
		name => "kmerfreq json",
		cli  => "./kmerfreq -j $TD | $sum",
		pass => "3940815226",
	},

	# longestorf
	{
		name => "longestorf *.gz",
		cli  => "./longestorf $TD | $sum",
		pass => "1472222220",
	},
	{
		name => "longestorf stdin",
		cli  => "gunzip -c $TD | ./longestorf - | $sum",
		pass => "1472222220",
	},
	{
		name => "longestorf 6-frame",
		cli  => "./longestorf -6 $TD | $sum",
		pass => "145014533",
	},

	# randomseq
	{
		name => "randomseq",
		cli  => "./randomseq 100 500 | wc -c | xargs",
		pass => "51490",
	},

	# smithwaterman
	{
		name => "smithwaterman",
		cli  => "gunzip -c $TD | head -8 | ./smithwaterman $TS - | $sum",
		pass => "3029323616",
	},
	{
		name => "smithwaterman tabular",
		cli  => "gunzip -c $TD | head -8 | ./smithwaterman -t $TS - | $sum",
		pass => "4023237785",
	},
	
	# viterbi
	{
		name => "viterbi",
		cli  => "",
		pass => "",
	},

);

my $passed = 0;
foreach my $test (@TEST) {
	print STDERR "$test->{name}: ";
	my $result = `$test->{cli}`;
	chomp($result);
	if ($result eq $test->{pass}) {
		$passed++;
		print STDERR "passed\n";
	} else {
		print STDERR "failed\n";
	}
}

print "passed $passed / ", scalar(@TEST), " tests\n";
