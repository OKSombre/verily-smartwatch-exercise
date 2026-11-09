# Build guide: Participant Compliance Calendar

A step-by-step path through the exercise in `README.md`, for you to write the code yourself.
It contains no solution code. Each step says what to build, what the README requires of it,
and how to check it before moving on.

Three kinds of statement appear below, and they are marked so you can tell them apart:

- **Spec** — taken directly from `README.md` or `seed/README.md`.
- **Decided** — something the spec leaves open that you answered when asked.
- **Open** — something the spec leaves open that is still yours to decide. Each one is a
  candidate for the "ambiguous" section of `NOTES.md`.

Anything not marked is a suggestion about order or technique, not a requirement.

## Decisions already made

| Topic | Decision |
|---|---|
| Stack | Go standard library `net/http`; Vite + React + TypeScript; React Router |
| Study segment of the URL | A path parameter. Only `alpha` exists; any other study returns `404` |
| Wear and sync indicators | Follow the README literally: every day cell shows both, including `FUTURE`, `NO_ACTIVITY` and `WAITING_TO_SYNC` days |
| How the frontend knows "today" | From its own environment variable carrying the same `AS_OF` date. The API response stays exactly as the contract shows it |
| `AS_OF` not provided | There is a default; the server and the frontend do not refuse to start |

Two consequences of the "today" decision to keep in mind. Vite only exposes variables
whose names start with `VITE_` to browser code, so `make dev` has to hand the one `AS_OF`
value to both processes under two names. And the two values can drift: if a grader starts
your server with one `AS_OF` and the frontend was started or built with another, "today" is
marked on the wrong day. Starting both from one command with one value is what prevents
that, and it is worth a line in `NOTES.md`.

## What the seed data contains

These are facts about the files in `seed/`, checked by reading them. They are useful for
choosing test cases.

- 14 participants (`A1000` to `A1013`) across Site 1, Site 2 and Site 3; 313 daily records.
- Records run from `2026-09-05` to `2026-10-21`. There are none after `asOf`, no duplicate
  participant/date pairs, and no record for a participant who is not in `participants.json`.
- No record has both values at zero. A day with nothing is always a *missing* record
  (**Spec**: a missing record means `wearHours = 0` and `syncCount = 0`).
- `A1000`, `A1005` and `A1013` have no records at all. `A1005` and `A1013` have a `null`
  `deviceId`; `A1000` has a device and still no records.
- `A1010` sits on the wear boundary: `16.0` hours with one sync on `2026-10-11` through
  `2026-10-16`, then `15.9` hours with one sync on `2026-10-17`.
- `A1009` has days with `wearHours` of `0` and `syncCount` of `1`, including on `asOf`.
  Walk the rule table for that day yourself: it is not the state you might guess first.
- Six participants have wear above zero and zero syncs on `asOf` (`A1001`, `A1002`,
  `A1006`, `A1007`, `A1010`, `A1011`). In the data this combination never occurs on an
  earlier date, so the "same values, but yesterday" case needs a test you construct.
- `A1007` enrolled on `2026-10-19` and has three records.

Numbers to check your work against, with `AS_OF=2026-10-21`:

| Request | Grid | State counts |
|---|---|---|
| `A1001`, `2026-10` | `2026-09-27` to `2026-10-31`, 35 days | 17 `MET`, 4 `UNMET`, 3 `NO_ACTIVITY`, 1 `WAITING_TO_SYNC`, 10 `FUTURE` |
| `A1001`, `2026-09` | `2026-08-30` to `2026-10-03`, 35 days | 5 `MET`, 1 `UNMET`, 29 `NO_ACTIVITY` |
| `A1009`, `2026-10` | 35 days | 25 `UNMET`, 10 `FUTURE` |
| `A1000`, `2026-10` | 35 days | 25 `NO_ACTIVITY`, 10 `FUTURE` |

Grid sizes for the three cases the spec names: February 2026 is 28 days (`02-01` to
`02-28`), October 2026 is 35, August 2026 is 42 (`07-26` to `09-05`). December 2026 crosses
a year boundary (`2026-11-29` to `2027-01-02`).

## Open questions for you

None of these block starting. Decide each when you reach the step named, and write the
answer in `NOTES.md`.

1. **Server port.** The README says a script will run against your server and does not say
   which port. Pick one, make it configurable, and document it. (Step 6)
2. **The default `AS_OF`, and a malformed one.** You decided there is a default when the
   variable is not provided; which date it is remains open. The two candidates are the seed
   reference date `2026-10-21` and the system date. `seed/README.md` says the server must
   take the date "from the `AS_OF` environment variable rather than the system clock", and
   the seed has no records after `2026-10-21`, so a system-date default would show every
   seed day as past. Separately, a value that is present but not a valid date is a
   different case from an absent one: using the default silently would hide a typo. Both
   points apply to the frontend's variable as well. (Steps 6, 9 and 14)
3. **The `wear` and `sync` values on `FUTURE` days.** The contract does not say. The seed has
   no future records, so with this data they come out as zero and `false`; what you would do
   if a future record did exist is unspecified. (Step 4)
4. **`month` missing altogether.** "Anything else returns `400`" most naturally covers a
   missing parameter, but the spec does not name that case. (Step 5)
5. **How strict `month` is.** `2026-13`, `2026-1` and `26-10` are all "anything else".
   Whether a year far in the past or future is acceptable is not stated. (Step 5)
6. **The shape of the error body.** The spec asks for a "consistent JSON error body" and
   does not define one. (Step 5)
7. **What "empty state" means.** The spec lists loading, empty and error states. The API
   always returns a full grid, so there is never an empty list. A participant with no
   records, or a month where every day is `NO_ACTIVITY`, is the nearest thing. (Step 12)
8. **A month after today in the URL.** Next is disabled once the displayed month contains
   today, but someone can still type a later month into the address bar. The spec does not
   say what the page does then. (Step 10)
9. **The default month.** When the URL names no month, which one shows? The month containing
   today is the obvious candidate; the spec does not say. (Step 10)
10. **Today outside the selected month.** If today falls in the leading or trailing days of
    the grid, is it still marked? (Step 11)
11. **The `__MACOSX/` folder** is committed alongside the exercise files. It is archive
    residue, not part of the spec. Whether you remove it is up to you.

## Part 1: Backend

### Step 1 — Lay out the repository

Create a Go module at the repository root and decide where the frontend lives (a
subdirectory is the usual choice). Add a `.gitignore` covering build output and
`node_modules`.

**Spec:** the day-state rules live in their own package. Plan for at least three separate
concerns on the Go side: the state rules, the calendar grid, and the HTTP layer with seed
loading. Keeping the first two free of any HTTP or file-reading code is what makes them
easy to test.

Path parameters and per-method routes in the standard library router need Go 1.22 or
later; check your version before you rely on them.

Check: the module builds with an empty `main`. Commit. The README asks for a commit
history that shows how you got there, so commit at the end of every step.

### Step 2 — The day-state package, tests first

Write one function that takes a date, the `asOf` date, wear hours and a sync count, and
returns one of the five states. Define the states as named constants whose string values
match the contract exactly.

**Spec:** the table in the README is evaluated in order. The order matters; the first
matching row wins.

Write the tests before the function. A table-driven test is the idiomatic Go form. Cover
at least:

- The day after `asOf`, with and without data.
- Zero wear and zero syncs, in the past and on `asOf`.
- Wear above zero with zero syncs on `asOf`, and the same values the day before.
- Wear of `16.0` and `15.9`, each with one sync.
- Wear of `16.0` or more with zero syncs on a past day.
- Zero wear with one or more syncs, on a past day and on `asOf`.
- `24` hours, the top of the range.

**Spec:** "wear met" is `wearHours >= 16` and "sync met" is `syncCount >= 1`. They are
reported separately from the state, so give them their own small functions and tests.

How you represent a date in this package is your first real design choice. **Spec:** dates
are plain calendar dates and you must not do time zone math. `time.Time` works if every
value is built in UTC and you never call anything that reads the local zone. A small type
of your own avoids the question. Either is fine; say which you chose in `NOTES.md`.

Check: `go test` passes for the package, and you can explain each row of the test table.

### Step 3 — The calendar grid

Write a function that takes a year and month and returns the ordered list of dates in the
grid.

**Spec:** weeks run Sunday through Saturday, and the grid covers the whole month plus the
trailing days of the previous month and the leading days of the next. The result is 28, 35
or 42 days.

The approach: find the first of the month, step back to the Sunday on or before it, find
the last of the month, step forward to the Saturday on or after it, and list every date
between them.

Test with the months in the table above: February 2026, October 2026, August 2026 and
December 2026. Assert the first date, the last date and the length. A leap-year February
is worth one more case.

### Step 4 — Load the seed data

Read both JSON files at startup into structures the handler can query: participants by ID,
and daily records by participant and date.

**Spec:** the files may be loaded however you like. A missing daily record means zero wear
and zero syncs, so a lookup that misses must return zeros, not an error.

Reading from disk and embedding the files in the binary are both allowed. Embedding
removes any dependence on the working directory the server is started from, which matters
when someone else's script starts it.

Decide open question 3 here. Then write the function that joins the three pieces: for a
participant and a month, produce the grid from step 3, look up each date, and apply the
functions from step 2.

Check: a test that builds the month for `A1001` in `2026-10` and compares the state counts
with the table above.

### Step 5 — The HTTP handler

Register one route for `GET` on the contract path, with the study and participant as path
parameters and `month` as a query parameter.

**Spec:** the response has `participantId`, `site`, `month` and `days`. Each day has
`date`, `state`, `wear` (`hours`, `met`) and `sync` (`count`, `met`). The field names and
nesting must match exactly. `month` that is not `YYYY-MM` returns `400`. An unknown
participant returns `404`.

**Decided:** a study other than `alpha` returns `404`.

Things to get right:

- Set the JSON content type on every response, errors included.
- Validate in a deliberate order and decide what happens when two things are wrong at once
  (an unknown participant and a bad month).
- Resolve open questions 4, 5 and 6. Write one helper that produces the error body so that
  every failure uses it.
- Consider what a request with the wrong method, or to a path that matches nothing,
  returns. The default responses from the standard router are plain text, not your JSON
  error body.

### Step 6 — Configuration and startup

**Spec:** the server reads `AS_OF` as `YYYY-MM-DD` from the environment. The grading script
sets `AS_OF=2026-10-21`.

Read it once at startup and pass it into the handler as a value. Do not read the
environment or the system clock inside the handler or the state package; that is what
makes the tests in the next step possible.

**Decided:** when `AS_OF` is not provided, the server uses a default instead of refusing to
start. Define that default in one place, and log the date in use at startup so that anyone
running the server can see which "today" it has.

Resolve open question 1 and the remaining parts of open question 2.

Check with the server running: request the contract URL and compare the first day with the
example in the README (`2026-09-27`, `NO_ACTIVITY`, zeros and `false`). Then try a bad
month, an unknown participant and an unknown study.

### Step 7 — HTTP-level tests

**Spec:** at least one test using `httptest`.

Build the handler in the test with a fixed `asOf`, send requests, and assert on status
code, content type and decoded body. Worth covering: the success case for `A1001`, a bad
month, an unknown participant, an unknown study, and that every error has the same body
shape.

## Part 2: Frontend

### Step 8 — Scaffold

Create the Vite React TypeScript app in the frontend directory and add React Router.
Remove the template's demo content.

The browser app and the Go server run on different ports during development. Vite's dev
server can proxy `/api` to the Go server, which keeps the frontend's request URLs relative
and avoids CORS configuration on the Go side.

### Step 9 — Types and the API client

Write TypeScript types that mirror the contract: the five states as a string union, a day,
and the calendar response. Then write one function that fetches a participant's month and
returns either data or a typed failure.

That function should distinguish three outcomes, because the page shows each differently:
success, participant not found (`404`), and everything else (network failure, `400`, a
server error).

Accept an `AbortSignal` in this function now. Step 12 needs it.

**Decided:** "today" comes from an environment variable on the frontend. Read it in one
place and validate that it looks like a date. When it is not provided, fall back to the
same default as the server; a different default on each side reintroduces the drift
described under "Decisions already made".

### Step 10 — Routing and the month in the URL

**Spec:** `/participants/:participantId` shows that participant. `/` may redirect to
`/participants/A1001`. The selected month survives a page refresh, for example in the URL.

Treat the URL as the only place the month is stored. Read the participant from the path
and the month from the query string on every render, and change month by navigating. If
you also keep the month in component state, the two will disagree after a refresh or a
back-button press.

Resolve open questions 8 and 9. Also decide what the page does when the month in the URL is
not a valid month: the server will answer `400`, and you can catch it earlier on the
client.

Check: change month, refresh, and confirm the same month is showing. Use the browser's
back button and confirm it steps back through months.

### Step 11 — The grid

**Spec:** you build the month grid yourself; no prebuilt calendar widget. Each cell shows
the date, the wear and sync indicators with their met or unmet status, and the day state.
Days outside the selected month are shown but visually secondary. Today is marked.

The server already returns the days in order and Sunday-aligned, so the frontend does not
compute any dates. Render a row of seven weekday headings and then the days in a
seven-column CSS grid.

Useful to know:

- A day is outside the selected month when the month part of its date differs from the
  response's `month`. Compare the strings.
- Do not pass these dates to `new Date("2026-10-21")`. JavaScript parses that form as UTC
  midnight and then displays it in the local zone, which shows the previous day in any zone
  west of UTC. Take the day number from the string instead. The same care applies when you
  work out the month name for the heading and the previous and next months.
- Today is the cell whose date equals the value from step 9. Resolve open question 10.

**Decided:** every cell shows both indicators, including `FUTURE`, `NO_ACTIVITY` and
`WAITING_TO_SYNC` days. This differs from the wireframe, which the README says is
illustrative.

**Spec:** `NO_ACTIVITY` and `WAITING_TO_SYNC` days carry a text label. The wireframe's
wording is "No watch activity" and "Waiting to sync".

**Spec:** the calendar must not rely on color alone for Met and Unmet. The wireframe pairs
color with a check mark or an exclamation mark. Whatever you choose, test it by viewing
the page in grayscale.

**Spec:** a legend covering watch wear, watch sync, Met and Unmet. Build it from the same
indicator components the cells use so the two cannot drift apart.

### Step 12 — Loading, error, not-found and stale requests

**Spec:** loading, empty and error states, including a clear "participant not found" state.
Rapidly changing the month must not show stale data from an earlier request.

Model the fetch as one value with distinct cases (loading, loaded, not found, failed)
rather than several independent booleans. Independent flags allow combinations that make
no sense, such as loading and failed at once.

The stale-data requirement is the one that is easy to believe you have met when you have
not. The failure looks like this: you click Next twice quickly, the second response
arrives first, then the first response arrives and overwrites it. Two ways to prevent it,
which work well together:

- Abort the previous request in the effect's cleanup, using the signal from step 9. An
  aborted fetch rejects, so make sure that rejection is not shown as an error.
- Before storing a response, confirm it belongs to the participant and month currently in
  the URL.

To check it, add an artificial delay on the server that is longer for one month than
another, click through quickly, and watch what settles. Remove the delay afterwards.

Also decide what shows while a new month loads: the previous grid dimmed, a skeleton, or a
blank area. Whichever you choose, the month heading should already name the new month.

Resolve open question 7.

### Step 13 — Navigation and accessibility

**Spec:** previous and next month navigation. Next is disabled once the displayed month
contains today. Keyboard-accessible month navigation, and sensible labels for screen
readers on day cells. The header shows the participant ID and site, with a back arrow that
may be inert.

- Use real `button` elements for previous and next. They are focusable and respond to
  Enter and Space without extra work. Give each an accessible name, since an arrow glyph
  alone has none.
- Think about how the disabled Next is conveyed. A natively disabled button drops out of
  the tab order; that is acceptable, and you should know that it happens.
- Write a full-sentence label for each day cell that carries everything the cell shows
  visually: the full date, the state, and the wear and sync status. Mark the icons as
  decorative so they are not announced twice.
- Mark today with more than color, and include it in the cell's label.
- Announce the month change. A heading in a polite live region is one way.

Check: put the mouse away and operate the whole page from the keyboard. Then listen to a
few cells with a screen reader (VoiceOver, NVDA or Narrator all work).

**Spec:** a way to switch participants is welcome but not required. Leave it until
everything above is done.

## Part 3: Finishing

### Step 14 — One command to run everything

**Spec:** `make dev`, or an equivalent documented command, starts everything.

Write a `Makefile` with a `dev` target that starts the Go server and the Vite dev server
with the same `AS_OF`, and a `test` target. **Decided:** `make dev` works without the
caller setting `AS_OF`. The Makefile is a natural place to hold the default and pass it to
both processes, though the server and frontend still need their own for when they are
started directly.

Check from a fresh clone, in a new terminal: install, `make dev`, open the page. This is
the step most likely to work on your machine only.

### Step 15 — Frontend tests

The spec requires tests on the backend and does not require them on the frontend. "Tests"
is among the things graded, so a few are worth having if time allows. The highest-value
targets are the pure functions: previous and next month arithmetic including the year
boundary, the "does this month contain today" check, and the day-cell label text.

### Step 16 — `NOTES.md`

**Spec:** one page covering how to run it and the tests, decisions and trade-offs, what
you would do next with another day, and anything ambiguous and how you resolved it. If you
used AI tools, say how, and be ready to explain any part of the code.

Write this yourself and in your own words. Material you already have:

- The decisions table and the open questions in this guide are your ambiguity list.
- What you cut: the spec asks you to say what and why.
- For the AI section, the accurate account of this session is: you asked an AI assistant
  to read the two READMEs and produce a step-by-step guide with no code; it asked you four
  questions about gaps in the spec and you answered them; it profiled the seed data and
  listed edge cases and expected counts; you wrote the code. Update that if your use
  changes as you go.

### Step 17 — Before you submit

- `AS_OF=2026-10-21`, the contract URL, and a field-by-field comparison with the README's
  example.
- Every row of the expected-counts table above.
- Bad month, unknown participant, unknown study: status code and error body.
- Refresh on a month other than the default.
- Rapid month changes with the artificial delay in place.
- Keyboard only, then grayscale.
- A fresh clone and `make dev`.
- Read the commit history as a reviewer would.

## Suggested order if time runs short

The README sets a timebox of one to two days and says a smaller scope done well is
preferred. Steps 2 to 7 and steps 10 to 13 are where the stated grading criteria sit. The
participant switcher, frontend tests beyond the pure functions, and visual polish are the
things to cut first, and `NOTES.md` is where you say so.
