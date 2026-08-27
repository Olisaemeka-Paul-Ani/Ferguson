# Ferguson — CLAUDE.md

AI-powered Fantasy Premier League (FPL) terminal dashboard in Go. Bloomberg Terminal aesthetic: multi-pane, dark, color-coded, keyboard-only. V1 ships August 25, 2026.

## Non-negotiables

- Language: Go 1.22+. Stack: Bubbletea (TUI), Lip Gloss (styling), Bubbles (components), `github.com/evertras/bubble-table` (squad table + selection, chosen over `bubbles/table` for its richer built-in styling), stdlib `net/http` + `encoding/json` only. No other dependencies without asking.
- Real FPL data only. **No mocking in production paths.**
- AI provider: Groq (`qwen/qwen3.6-27b`) primary, Google Gemini (`gemini-3.6-flash`) fallback. API keys from env vars `FERGUSON_GROQ_KEY` and `FERGUSON_AI_KEY` respectively. Never hardcode keys.
- `team_id` passed as CLI flag: `ferguson --team 1234567`. No config file in V1.
- All Lip Gloss styles live in `styles.go`. No inline styling anywhere.
- Follow Bubbletea Model-Update-View strictly. View is a pure function of Model. All side effects (HTTP) go through `tea.Cmd`.

## File structure (do not deviate)

```
ferguson/
├── main.go        # entry point, flag parsing, tea.NewProgram
├── model.go       # tea.Model struct, Init(), Update(), View()
├── styles.go      # all lipgloss.Style definitions
├── fpl/
│   ├── client.go  # HTTP client, FPL API requests
│   └── types.go   # structs mirroring FPL API JSON
├── ai/
│   ├── client.go  # AI API request
│   └── prompt.go  # Ferguson's Verdict prompt template
└── ui/
    ├── squad.go, fixtures.go, verdict.go, splash.go, help.go
```

## FPL API (public, no auth)

Base: `https://fantasy.premierleague.com/api/`

- `bootstrap-static/` — all players, clubs, gameweeks, FDR
- `entry/{team_id}/` — team name, rank, total points (header)
- `entry/{team_id}/event/{gw}/picks/` — current squad picks
- `fixtures/?event={gw}` — fixtures for a gameweek

## V1 scope (nothing else)

P0: ASCII splash (~1.5s), scrollable squad table (name/pos/club/cost/points/GW points), fixtures pane (next 5 per squad club, FDR color-coded), AI verdict pane, keyboard nav (tab=pane, q=quit, r=refresh, ?=help), live data.
P1: help overlay, loading spinner.
OUT: leagues, transfers execution, history/charts, multi-user, FPL auth, config file.

## Visual spec

Black background; white/grey text; burnt amber `#B45309` accents + active-pane border; red `#FF4444` hard fixtures; yellow `#FFD700` neutral; green `#00FF88` easy. Three panes side by side (~30/35/35), footer keybind bar.

## Verdict prompt (ai/prompt.go)

Persona: Sir Alex Ferguson, first person, authoritative, 3–5 sentences, no bullets. Inject: squad, fixture difficulty per club, last-3-GW form. Ask for: captain pick, bench call, one transfer consideration.

## Working with the human (Olisaemeka) — pair-programming mode

- He is learning Go through this project. Core files (`main.go`, `model.go` — the Model struct and Update/View flow): explain the concept and the shape first, then he writes it himself. If a technique was already taught earlier in the project, ask him to attempt it first before helping.
- Boilerplate (FPL API structs in `fpl/types.go`, HTTP clients, Lip Gloss style definitions, ASCII art, JSON parsing): generate freely, but summarize what was built and why in a couple of sentences.
- When his attempt has a bug, point at the line and the category of bug — let him find the fix before giving it.
- Small commits, one feature each. Verify by running (`go run . --team <id>`), not by assuming.
