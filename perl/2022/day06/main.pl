#!/usr/bin/env perl

=head1 NAME

  Day 06 - Tuning Trouble

=head1 SYNOPSIS

  Perl solution for day 06 of AoC

=head1 DESCRIPTION

  To be able to communicate with the Elves, the device needs to lock on to their
  signal. The signal is a series of seemingly-random characters that the device
  receives one at a time.

  To fix the communication system, you need to add
  a subroutine to the device that detects a start-of-packet marker in the
  datastream. In the protocol being used by the Elves, the start of a packet is
  indicated by a sequence of four characters that are all different.

=head1 DEPENDENCIES

  ARGV::OrDATA
  Set::Scalar

=head1 SEE ALSO
  
  L<AoC|https://adventofcode.com/2022/day/6/>

=cut

use v5.36;

use ARGV::OrDATA;

my @input;
while (<>) {
  chomp;
  s/\r//gm;
  push @input, $_;
}

sub part_1 {
  my @answer;
  my $pkt_marker = 4;
  for my $line (@input) {
    my @data_stream = split //, $line;

    my $start = 0;
    while ($start + $pkt_marker - 1 <= $#data_stream) {
      use Set::Scalar;
      my $freq = Set::Scalar->new();
      $freq->insert(@data_stream[$start .. $start + $pkt_marker - 1]);

      if ($freq->size == $pkt_marker) {
        push @answer, $start + $pkt_marker;
        last;
      }
      $start++; 
    }
    say @answer;
  }
}

sub part_2 {
  my @answer;
  my $msg_marker = 14;
  for my $line (@input) {
    my @data_stream = split //, $line;

    my $start = 0;
    while ($start + $msg_marker - 1 <= $#data_stream) {
      use Set::Scalar;
      my $freq = Set::Scalar->new();
      $freq->insert(@data_stream[$start .. $start + $msg_marker - 1]);
      
      if ($freq->size == $msg_marker) {
        push @answer, $start + $msg_marker;
        last;
      }
      $start++; 
    }
      say @answer;
  }
}

part_1();
part_2();

__DATA__
bvwbjplbgvbhsrlpgdmjqwftvncz
nppdvjthqldpwncqszvftbrmjlhg
nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg
zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw
