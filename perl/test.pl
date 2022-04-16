use strict;
use warnings;

my @TEST = (
	{
		name => "randomseq",
		cli => "./randomseq 100 500 | wc -c",
		pass => "51490",
	},
	{
		name => "kmerfreq",
		cli => "./kmerfreq data/testdb.fa.gz | md5sum | cut -f1 -d' '",
		pass => "3debb76c314fa3fe51fcac2952d2eebd",
	},
	{
		name => "longestorf",
		cli => "./longestorf data/testdb.fa.gz | md5sum | cut -f1 -d' '",
		pass => "28b6183998f5c4834d9d4c2a2b1c414f",
	},
	{
		name => "dust",
		cli => "./dust data/testdb.fa.gz | md5sum | cut -f1 -d' '",
		pass => "9154561d8a4ca4c7377be87d0e3f6cf0",
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
