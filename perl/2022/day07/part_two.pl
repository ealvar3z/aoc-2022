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

my $t = 70_000_000;
my $req = 30_000_000;
my $not_used = $t - $dirs{''};
my $needed = $req - $not_used;
my $small = '';

foreach my $dir (keys %dirs) {
    $small = $dir if $dirs{$dir} > $needed && $dirs{$dir} < $dirs{$small};
}

say $dirs{$small};


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
