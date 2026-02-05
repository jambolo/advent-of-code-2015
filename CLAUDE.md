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
- Import shared file reader as `load "github.com/jambolo/advent-of-code-2015/internal/common"`, call `load.ReadLines(path)`.
- The Makefile auto-discovers days by scanning `cmd/`.
