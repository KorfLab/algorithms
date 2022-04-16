use strict;
use warnings;

my $TS = "data/testseq.fa";
my $TD = "data/testdb.fa.gz";
my $MD5 = "md5sum | cut -f1 -d ' '";

my @TEST = (

	# dust
	{
		name => "dust *.gz",
		cli  => "./dust $TD | $MD5",
		pass => "9154561d8a4ca4c7377be87d0e3f6cf0",
	},
	{
		name => "dust stdin",
		cli  => "gunzip -c $TD | ./dust - | $MD5",
		pass => "9154561d8a4ca4c7377be87d0e3f6cf0",
	},
	{
		name => "dust soft",
		cli  => "gunzip -c $TD | ./dust -l - | $MD5",
		pass => "77de668c6d42b885260676c9140fdf92",
	},

	# kmerfreq
	{
		name => "kmerfreq *.gz",
		cli  => "./kmerfreq $TD | $MD5",
		pass => "3debb76c314fa3fe51fcac2952d2eebd",
	},
	{
		name => "kmerfreq stdin",
		cli  => "gunzip -c $TD | ./kmerfreq - | $MD5",
		pass => "3debb76c314fa3fe51fcac2952d2eebd",
	},
	{
		name => "kmerfreq json",
		cli  => "./kmerfreq -j $TD | $MD5",
		pass => "9495f599da0b83243ac9825878119147",
	},

	# longestorf
	{
		name => "longestorf *.gz",
		cli  => "./longestorf $TD | $MD5",
		pass => "28b6183998f5c4834d9d4c2a2b1c414f",
	},
	{
		name => "longestorf stdin",
		cli  => "gunzip -c $TD | ./longestorf - | $MD5",
		pass => "28b6183998f5c4834d9d4c2a2b1c414f",
	},
	{
		name => "longestorf 6-frame",
		cli  => "./longestorf -6 $TD | $MD5",
		pass => "43c16b718ff84c030098f43f3c90a4c0",
	},

	# randomseq
	{
		name => "randomseq",
		cli  => "./randomseq 100 500 | wc -c",
		pass => "51490",
	},

	# smithwaterman
	{
		name => "smithwaterman",
		cli  => "gunzip -c $TD | head -8 | ./smithwaterman $TS - | $MD5",
		pass => "ddf95a0a5416ccf10905f2f61d6bdeff",
	},
	{
		name => "smithwaterman tabular",
		cli  => "gunzip -c $TD | head -8 | ./smithwaterman -t $TS - | $MD5",
		pass => "5fbd17e7f648d7fe9300094f594e41bb",
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
