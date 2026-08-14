# Ferguson

*Currently in development.*

Ferguson is a fantasy soccer analytics terminal — a Bloomberg Terminal-style
dashboard for Fantasy Premier League managers who want a sharper edge than
guesswork.

**Purpose:** pull your live squad, fixtures, and form data, then use AI to
turn it into a real recommendation — who to captain, who to bench, and
whether that transfer is actually worth it. No spreadsheets, no tab-switching
between five FPL sites. One dark, keyboard-only screen, three panes, one
verdict.

**Built with:**
- Go + [Bubbletea](https://github.com/charmbracelet/bubbletea) for the terminal UI
- [Lip Gloss](https://github.com/charmbracelet/lipgloss) for styling
- The public FPL API for live squad/fixture/points data
- Google Gemini (with Groq as fallback) to generate the AI verdict

```
ferguson --team <your_id>
```

Squad pane · Fixtures pane · AI Verdict pane — tab between them, `r` to
refresh, `?` for help, `q` to quit.
