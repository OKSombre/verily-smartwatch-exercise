# Take-Home: Participant Compliance Calendar

**Timebox: about 1-2 days.** Please don't spend much more. We would rather see a smaller scope done
well than everything done hastily. Tell us what you cut and why (see [Submission](#submission)).

## The scenario

A clinical study gives participants a smartwatch. To be useful, the data has to be collected
consistently: the participant must **wear** the watch enough each day, and the watch must **sync**
its data to the cloud. Study coordinators want to open a participant and see, day by day, whether
those requirements were met.

You are building one screen, end to end:

- **Backend:** Go. Serves a participant's month of daily data and computes each day's state.
- **Frontend:** React (TypeScript preferred). A month calendar for one participant.

There is no auth and no real database requirement. Data comes from the provided seed files.

## The screen

The wireframe is `wireframes/participant-calendar.png`. It is illustrative: where it and this
document disagree, **this document wins**. The wireframe was drawn from different data than
`seed/`, and some days show only a wear or only a sync indicator. Follow the rules below, not the
icons.

A month calendar for a single participant:

- Header with the participant ID and site, and a back arrow (it can be inert).
- Previous and next month navigation, and the month name.
- A legend: watch wear, watch sync, Met, Unmet.
- A Sunday-to-Saturday grid. Each day shows its date, a wear indicator, a sync indicator, and
  whether the day was Met or Unmet.
- Days with no data or still waiting to sync carry a text label ("No watch activity",
  "Waiting to sync"). Today is marked.

## Domain rules

All dates are plain calendar dates (participant-local). Do not do time zone math. The calendar
grid's weeks run **Sunday through Saturday**.

### Daily data (per participant, per day)

The seed file has, for each participant and day: `wearHours` (0 to 24) and `syncCount` (number of
successful syncs that day).

### Day state

Evaluated in this order:

| Condition | State |
|---|---|
| Date is after `asOf` (today) | `FUTURE` |
| `wearHours == 0` and `syncCount == 0` | `NO_ACTIVITY` |
| Date is `asOf`, `wearHours > 0`, `syncCount == 0` | `WAITING_TO_SYNC` |
| `wearHours >= 16` **and** `syncCount >= 1` | `MET` |
| otherwise | `UNMET` |

- **Wear met** means `wearHours >= 16`. **Sync met** means `syncCount >= 1`. Both are shown as
  separate indicators on each day cell.

## Provided

- `seed/participants.json` and `seed/daily.json` (see `seed/README.md`)
- `wireframes/participant-calendar.png`
- The API contract below

The server must read `AS_OF` (a date, `YYYY-MM-DD`) from the environment to define "today", so
results are reproducible. Our grading script sets `AS_OF=2026-10-21`.

## API contract

Implement this exactly. We will run a script against your server.

`GET /api/studies/alpha/participants/A1001/calendar?month=2026-10`

```json
{
  "participantId": "A1001", "site": "Site 1", "month": "2026-10",
  "days": [
    { "date": "2026-09-27", "state": "NO_ACTIVITY",
      "wear": { "hours": 0, "met": false }, "sync": { "count": 0, "met": false } }
  ]
}
```

- `days` covers the **full Sunday-aligned grid** for the month (including the trailing days of the
  previous month and leading days of the next, as in the wireframe). That is 28, 35 or 42 days
  depending on the month.
- `state` is one of `MET | UNMET | NO_ACTIVITY | WAITING_TO_SYNC | FUTURE`.
- `month` is `YYYY-MM`; anything else returns `400`. An unknown participant returns `404`.
- Use sensible status codes and a consistent JSON error body for failures.

## Requirements

### Backend (Go)

- Standard library `net/http` or a router such as `chi`. No ORM or database required.
- The day-state rules live in their own package with **unit tests**. Cover the edge cases
  yourself (the seed data has participants on both sides of each boundary). We will read them.
- At least one HTTP-level test (`httptest`).
- Input validation, with clear errors.

### Frontend (React)

- The app opens on the calendar for a participant. `/participants/:participantId` shows that
  participant, and `/` may redirect to `/participants/A1001`. A way to switch participants (even
  a text box) is welcome but not required.
- The month grid is built by you. You may **not** use a prebuilt calendar widget for it. We want
  to see you build it.
- Each day cell shows the date, the wear and sync indicators with their met / unmet status, and
  the day state. Include a legend like the wireframe, and mark today.
- `NO_ACTIVITY` and `WAITING_TO_SYNC` days carry a text label. Days outside the selected month
  are shown but visually secondary.
- Previous and next month navigation. The selected month survives a **page refresh** (for example
  in the URL). Next is disabled once the displayed month contains today.
- Loading, empty and error states, including a clear "participant not found" state.
- Rapidly changing the month must not show stale data from an earlier request.
- The calendar must not rely on color alone to communicate Met / Unmet.
- Keyboard-accessible month navigation, and sensible labels for screen readers on day cells.
- You may use a component or charting library elsewhere.

### Both

- `make dev` (or equivalent documented command) starts everything.
- `NOTES.md` as described below.

## Submission

Send a link to a Git repository (or a zip) containing:

- Your code, with a commit history (we like to see how you got there).
- `NOTES.md` (one page is plenty):
  - How to run it and the tests.
  - Decisions you made and trade-offs.
  - What you would do next with another day.
  - Anything in this spec you found ambiguous and how you resolved it.

## What we look for

Correctness of the day-state rules and the calendar grid, API and code quality, how you handle
states and edge cases in the UI, tests, and how clearly you communicate. We will schedule a short
follow-up conversation to walk through your solution together. You can use AI tools if you like.
If you do, tell us how you used them in `NOTES.md`, and be ready to explain any part of your
code.
