# Ferguson

*Currently in development.*

A TUI dashboard for Fantasy Premier League — Bloomberg Terminal-style, dark,
keyboard-only, multi-pane.

Pulls your live squad, fixtures, and form data from the public FPL API, and
generates an AI-written verdict (captain pick, bench call, one transfer
worth considering) styled as Sir Alex Ferguson's post-match assessment.

**Built with:**
- Go + [Bubbletea](https://github.com/charmbracelet/bubbletea) for the terminal UI
- [Lip Gloss](https://github.com/charmbracelet/lipgloss) for styling
- The public FPL API for live squad/fixture/points data
- Groq (with Google Gemini as fallback) to generate the AI verdict

```
ferguson --team <your_id>
```

No config file in V1 — your team ID is passed via the `--team` flag each run.

Squad pane · Fixtures pane · AI Verdict pane — tab between them, `r` to
refresh, `?` for help, `q` to quit.
