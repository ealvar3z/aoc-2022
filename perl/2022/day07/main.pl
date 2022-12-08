#!/usr/bin/env perl

use v5.36;
use ARGV::OrDATA;

my @cwd = ("");
my %dirs;

while (<>) {
  chomp;
  if (m{^\$ cd \.\.$}) {
     pop @cwd if @cwd > 1;
  } elsif (m{^\$ cd /$}) {
      @cwd = ("");
  } elsif (m{^\$ cd (.+)}) {
      push @cwd, $1;
  } elsif (/^([0-9]+) (.+)/) {
      my ($sz, $name) = ($1, $2);
      $dirs{ join '/', @cwd[0 .. $_] } += $sz for 0 .. $#cwd;
  }
}

my $ans = 0;
foreach my $sz (values %dirs) {
  $ans += $sz if $sz <= 100_000;
}

say $ans;


__DATA__
$ cd /
$ ls
dir a
14848514 b.txt
8504156 c.dat
dir d
$ cd a
$ ls
dir e
29116 f
2557 g
62596 h.lst
$ cd e
$ ls
584 i
$ cd ..
$ cd ..
$ cd d
$ ls
4060174 j
8033020 d.log
5626152 d.ext
7214296 k
