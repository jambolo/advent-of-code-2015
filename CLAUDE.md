# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Build & Run

```bash
make build          # Compile all solutions to bin/
make clean          # Remove bin/
```

Run a solution:

```bash
./bin/day01 -file data/day01/day01-input.txt -part 1
./bin/day01 -file data/day01/day01-input.txt -part 2
```

Flags: `-file` (default: `data/default.dat`), `-part` (1 or 2).

## Architecture

Go module: `github.com/jambolo/advent-of-code-2015` (Go 1.25.7, no external deps).

```
cmd/dayXX/dayXX.go      # Each day's solution (standalone main package)
data/dayXX/              # Puzzle inputs and examples
internal/common/load.go  # Shared utility: ReadLines(path) reads file into []string
bin/                     # Compiled binaries (gitignored)
```

## Conventions

- Each day: create `cmd/dayXX/dayXX.go` and `data/dayXX/dayXX-input.txt`.
- Solutions use `flag` for CLI args, print a banner (`=== Day X - Part Y ===`), then output results.
- Import shared file reader as `"github.com/jambolo/advent-of-code-2015/internal/load"`, call `load.ReadLines(path)`.
- Import setup as `"github.com/jambolo/advent-of-code-2015/internal/setup"`.
- Import utils as `"github.com/jambolo/advent-of-code-2015/internal/utils"`.
- The Makefile auto-discovers days by scanning `cmd/`.

## Setting Up a New Day

When asked to "set up a new day" (day NN):

1. Create `cmd/dayNN/dayNN.go` with only the skeleton: parse flags, print banner, read input via `load.ReadLines`. It may be incomplete and not compile.
2. Create an empty `data/dayNN/dayNN-input.txt`.
3. Add `bin/dayNN` as a dependency of the `build` target in the Makefile.
4. Add a `## Day NN` heading at the end of `README.md`.
