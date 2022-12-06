#! /usr/bin/perl

use v5.36;

my $input = "../input/2022/day05.txt";

sub part_1 {
  open(my $fh, "<",  $input) || die "Can't open: $!";
  my @stacks;
  while (my $line = <$fh>) {
      last if $line =~ /\d/;

      for (my $i = 1; $i < length($line); $i += 4) {
          my $crate = substr $line, $i, 1;
          
          my $position = ($i - 1 ) / 4;
          push(@{ $stacks[$position] }, $crate) if $crate ne " ";
      }
  }
<$fh>;
  while (<$fh>) {
      my ($quantity, $from, $to) = /move ([0-9]+) from ([0-9]+) to ([0-9]+)/;
      unshift(@{ $stacks[$to - 1] }, reverse splice @{ $stacks[ $from - 1] }, 0, $quantity);
  }

  say map $_->[0], @stacks;
  close($fh);
}


sub part_2 {
  open(my $fh, "<",  $input) || die "Can't open: $!";
  my @stacks;
  while (my $line = <$fh>) {
      last if $line =~ /\d/;

      for (my $i = 1; $i < length($line); $i += 4) {
          my $crate = substr $line, $i, 1;
          
          my $position = ($i - 1 ) / 4;
          push(@{ $stacks[$position] }, $crate) if $crate ne " ";
      }
  }
<$fh>;
  while (<$fh>) {
      my ($quantity, $from, $to) = /move ([0-9]+) from ([0-9]+) to ([0-9]+)/;
      unshift(@{ $stacks[$to - 1] }, splice @{ $stacks[ $from - 1] }, 0, $quantity);
  }

  say map $_->[0], @stacks;
  close($fh);
}


part_1();
part_2();
