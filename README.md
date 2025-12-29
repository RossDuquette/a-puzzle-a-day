# a-puzzle-a-day
A-Puzzle-A-Day solver using Go.

Puzzle can be found [here](https://www.dragonfjord.com/product/a-puzzle-a-day/).

Note that this is a personal hobby project and has nothing to do with the authors of the puzzles
or my employer. The sole purpose of this project was to help me on my journey learning Go.

## Usage
`./a-puzzle-a-day --help` for up to date usage.

## How it works
Tiles are placed in every possible position and orientation, calculating the total number of
unique solutions for each day.

These solutions are either displayed to the terminal, or written to files.

Hardware dependant, but modern PCs should be able to find all 64 solutions for Jan-01 in less than
2 seconds, as an example.
