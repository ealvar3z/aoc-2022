#!/usr/bin/env perl

use v5.36;
use ARGV::OrDATA;

=begin comment
1. parse the input.
2. populate the cube set.
3. create a grid:
  - if it touches the lava droplet -> '#'
  - they are adjacent
4. count the number of adjacent cubes

=end comment
=cut

# parsing the input
# populating the set
# creating the grid
while (<>) {
    no strict;
    chomp;
    /^(\d+),(\d+),(\d+)/;
    $cubes->{$_} = {};
    $grid->{$1}->{$2}->{$3} = '#';
}

use vars qw($cubes $grid);
# count the num of cubes adjacent to each other
sub adjacent($x, $y, $z, $rune) {
    no warnings qw(void once uninitialized);
    my $n = 0;
    $n++ if $grid->{$x+1}->{$y}->{$z} eq $rune;
    $n++ if $grid->{$x-1}->{$y}->{$z} eq $rune;
    $n++ if $grid->{$x}->{$y+1}->{$z} eq $rune;
    $n++ if $grid->{$x}->{$y-1}->{$z} eq $rune;
    $n++ if $grid->{$x}->{$y}->{$z+1} eq $rune;
    $n++ if $grid->{$x}->{$y}->{$z-1} eq $rune;
    return $n;
}

=begin comment
x, y, z
1, 1, 1
2, 1, 1
^  ^^^^---- these two sides touch
^---------- this side does not
^^^^^^^---- total sides: 12
            approx. count is total sides (12) minus non-connected (2): 12 - 2 = 10
So, to approximate the count:
We take the number of cubes the and multiply by 6 (i.e. the total sides of a cube), and get the total # of sides.
Then substract the number of cubes that are adjacent ($n).

=end comment
=cut

sub part_one() {
   my $n = 0;
   for (keys %$cubes) {
       $n += adjacent((split /,/, $_), '#');
   }
   my $_cubes = (scalar(keys %$cubes));
   my $total_sides = ($_cubes * 6); 
   say "TSA= " . ($total_sides - $n);
}

# =begin comment
# 1. determine the exterior areas
#   - set the left & right lateral limits of our grid
#   - basically the min & max number in our input:
#     - that would be 0 & 19 respectively
#     - so we set the lateral limits at -1 & 20
# 2. create the grid:
#   - if there's a cube of air -> 'air'
#   - we've reached the lava droplet that trapped the cube of #!
# 3. we need to find what cubes are adjacent to each other w/
#    the new created grid for every point in the grid.
# 3. we compute the surface area as previous except we exclude the
#    lava droplets that have trapped #.
#
# =end comment
# =cut

use vars qw($exclude);
sub part_two() {
    my $n = 0;
    my $left_lat_lim = 0 - 1; 
    my $right_lat_lim = 19 + 1;
    $grid->{$left_lat_lim}->{$left_lat_lim}->{$left_lat_lim} = 'air';
    do {
        # while the points are w/in the left & right lateral limits
        # search ...
        $exclude = 0;
        for my $z ($left_lat_lim..$right_lat_lim) {
            for my $y ($left_lat_lim..$right_lat_lim) {
                for my $x ($left_lat_lim..$right_lat_lim) {
                    # breakout if the points are already in the grid
                    # count the points that need to be excluded
                    next if defined $grid->{$x}->{$y}->{$z};
                    adjacent($x, $y, $z, 'air') && ($grid->{$x}->{$y}->{$z} = 'air');
                    $exclude++ unless not defined $grid->{$x}->{$y}->{$z};
                }
            }
        }
    } while ($exclude);

    for (keys %$cubes) {
        $n += adjacent((split /,/, $_), 'air');
    }
    say "ESA= " . $n;
}

part_one();
part_two();

__DATA__
2,2,2
1,2,2
3,2,2
2,1,2
2,3,2
2,2,1
2,2,3
2,2,4
2,2,6
1,2,5
3,2,5
2,1,5
2,3,5
