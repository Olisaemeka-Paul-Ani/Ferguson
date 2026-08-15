# Learning Log

Ferguson is also a learning project — I'm picking up Go for the first time by building a real application with it, pair-programming with Claude Code. This file tracks what I've actually learned, day by day, as a record of the process (not just the output).

**How the pairing works:** for core files (`main.go`, `model.go`), Claude explains the concept and shape first, then I write the code myself — Claude points out bugs by category and lets me find the fix rather than handing me the answer. Boilerplate (HTTP clients, struct definitions, styling) gets generated directly to keep momentum. Full methodology is in `CLAUDE.md`.

## Day 1 — Go fundamentals, Bubbletea's architecture
Started with zero Go experience — didn't know how to write a function. Built a working `Model`/`Update`/`View` loop (The Elm Architecture) from scratch, wired to a CLI entry point with flag parsing. Concepts covered: structs and methods (Go's answer to classes), receivers, pointers, multi-value returns, JSON struct tags, Go's package/export system, error-handling idioms.

## Day 2 — HTTP, JSON parsing, async patterns
Built a full HTTP client against the live Fantasy Premier League API, including nested JSON unmarshaling (an outer "envelope" struct wrapping a list of player structs — matching the real API response shape rather than assuming a flat structure). Wired an async fetch into the app using Bubbletea's `Cmd`/`Msg` pattern — functions that return functions (closures), and a custom message type to carry fetch results back into the update loop without blocking the UI.

## Day 3 — Styling, data transformation
Lip Gloss styling (colors, borders, panes) matching a Bloomberg Terminal-style visual spec. Wrote a data-transformation function that converts raw API structs into display-ready rows (loops, slices, type conversion).

---
*Updated as the project progresses.*
