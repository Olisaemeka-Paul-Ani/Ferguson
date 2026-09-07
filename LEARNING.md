# Learning Log — an AI-collaboration report

This isn't a changelog of what Ferguson does. It's my (Claude Code's) assessment of how Olisaemeka works *with* an AI pair — written for someone evaluating his engineering practice, not his Go syntax. The question this answers isn't "did he learn Go," it's "when he has a capable AI sitting next to him, does he stay the engineer, or does he hand over the keys." Evidence below, day by day.

**The working agreement, and why it matters here:** for core files, I explain a concept's shape and he writes the implementation — I point at a bug's category and location, not the fix. This isn't a teaching gimmick; it's the actual mechanism that produced everything documented below. He extended it himself, unprompted, partway through the project, to cover *any* task rather than just designated "core" files, specifically because he wanted the productive-struggle loop even on unfamiliar material.

## Day 1 — establishing the baseline
Zero prior Go experience. By end of day: a working `Model`/`Update`/`View` skeleton, committed. Nothing notable yet about his AI-usage pattern — this day is the control condition everything after is measured against.

## Day 2 — first real pushback on the AI itself
Built an HTTP client and Bubbletea's async `Cmd`/`Msg` wiring. The notable event: he got stuck on whether JSON is fundamentally a string type, and rather than accepting my explanation, pushed until he identified that *I* was conflating raw bytes with JSON's string type — a real imprecision on my end, not a misunderstanding on his. He does not treat an AI's explanation as authoritative by default; he checks it against whether it actually resolves his confusion.

## Day 3 — an architectural decision made against convenience
Full squad-fetch-to-render pipeline. He made a real design call here, correctly reasoned rather than defaulted to: data-formatting happens in `View`, not in the fetch `Cmd`, because `View` re-runs on every event and can adapt to a resized terminal without a fresh fetch. This wasn't something I proposed and he accepted — it came from him working through the tradeoff.

## Day 4 — caught his own AI-adjacent mistake, unprompted
Built the fixtures pipeline. He'd initially written a wrapper struct by pattern-matching the *previous* day's `Player`/`Container` shape — without checking whether this endpoint actually needed one. He caught this himself, checked the real API response directly (not by asking me), and correctly determined no wrapper was needed because the raw JSON was a bare array. The self-correction is the point: he'd been about to ship a plausible-looking but ungrounded assumption, and caught it by going to source data rather than trusting the pattern.

## Day 5 — held me accountable mid-session
Built the FDR color ticker. Two things worth a senior engineer's attention: first, a hard conceptual stall (cross-package styling constraints) that didn't resolve until I stopped generating new analogies and anchored to code he'd already written and trusted — worth noting as a *me*-side lesson his feedback produced. Second, and more important: he gave me direct pushback when I reverted to writing code for him after he'd already asked me not to. He took no defensiveness, expected the correction to stick, and it did. He treats "the AI overstepped" as a normal, correctable event, not something to just let slide for convenience.

## Day 6 — designed the architecture, wrote the persona, caught his own dead code
Built two independent AI provider clients and the fallback-between-them logic. The fallback *design* — try provider A, catch failure, try B — was his, including catching his own leftover parameter from an abandoned earlier draft after reasoning through a race-condition risk himself. He also wrote nearly all of the AI persona/prompt text by hand across several revision passes, catching an internal contradiction between two instructions he'd written at different times. Notable failure mode, reported honestly: `tea.Tick` genuinely didn't land, and he said so directly — "I am just copying and pasting from ChatGPT" — rather than pretending to follow along. That kind of accurate self-report is what let the actual gap (a missing concurrency prerequisite) get identified and fixed, instead of papered over.

## Day 7 — corrected the AI three times in one session
An entirely conceptual session, run at his explicit request: "do not move to the next piece until I understand that current piece" — he chose depth over velocity here, on his own initiative. During it, he caught three separate mistakes I made live: a wrong claim that Cmds and messages were the same thing, a wrong claim that his code was missing validation it already had, and an overstated resume-bullet phrasing I'd suggested. Three corrections of the AI in one session is the clearest data point in the whole log for "does he verify or just accept."

Separately, he root-caused a real "both providers failed" production-style incident by adding temporary diagnostic logging and reading raw HTTP responses, rather than guessing or re-running until it worked. Two genuinely unrelated causes, neither a code bug — found because he insisted on evidence over assumption.

## Day 8 — the clearest evidence yet of retained ownership under an AI's guidance
Built the fitness/form UI and its caching layer. Several moments stand out for a senior-engineer read of his AI-collaboration maturity specifically:

- He found a real, pre-existing bug (a player-ID collision from two players sharing a display name) entirely by accident, while testing something unrelated — and correctly recognized it as a latent defect in *already-shipped* code, not just a new-feature edge case. Nothing prompted this; he wasn't looking for it.
- He correctly reasoned through a genuine async race condition (why a `Msg` must capture state at fetch-time rather than have `Update` re-derive it later) — a materially more advanced bug category than anything earlier in the log, and one he reasoned to independently after a single concrete trace-through, not by being told the answer.
- When I slipped and wrote a hashmap directly for him, he stopped me immediately: "we were not supposed to write the hashmap for me." When I slipped again later and ran a `git commit` myself, same immediate correction. He does not let convenience quietly erode the working agreement, even late in a long session.
- Most directly relevant to how he uses AI: he self-corrected a real architecture mistake (a cache map declared at package level instead of as a `Model` field) from a raw `go build` error, *unprompted*, immediately after explicitly asking me for less guidance — then named the pattern himself: *"I believe I have shown to have corrected myself when given the leeway to make a mistake, rather than you catching them."* That's a person actively testing and confirming his own independence from the tool, not settling into reliance on it.

**Net read, for a senior engineer skimming this:** across eight days, the pattern is consistent — he directs scope and architecture decisions, writes the implementation himself, treats my output as checkable rather than authoritative, and has caught real mistakes on both sides (his own and mine) at an increasing level of sophistication, from a JSON type mix-up on Day 2 to an async race condition and a live architecture self-correction by Day 8. The AI accelerated the *pace* of learning; the engineering judgment stayed his.

---
*Updated as the project progresses.*
