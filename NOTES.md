# Progress Notes on Take-Home Assessment
These are the notes I've written as I've progressed through this assignment. These contain my thoughts and processes for my coding journey these two day.

## Structure and AI Usage
To create the basic architecture and assure any spec ambiguities were covered, I established a base with Claude code and created a planning guide from start to finish of the project. All code was however created by myself, and I strictly prompted Claude to make no assumptions and clarify any ambiguity through me. The code is then proof-read through Claude for assurance of spec and checking against edge cases or vulnerabilities.

## Decisions and Trade-offs
Several changes to address ambiguity and best practices were made in order to ensure best code quality and best completion of spec. The choice to use net/http + Vite + React Router was made due to my own familiarity with the tools. However, the spec required that the AS_OF date be chosen through the environment, so I decided to use make dev include a variable that both the front-end and back-end can use, defaulting to 2026-10-21 if one isn't provided. I also chose to use strings for the back-end handling of dates to avoid anything with timezones or UTC.

### Other Spec Ambiguity and Clarifications
| Topic | Decision |
|---|---|
| Stack | Go standard library `net/http`; Vite + React + TypeScript; React Router |
| Study segment of the URL | A path parameter. Only `alpha` exists; any other study returns `404` |
| Wear and sync indicators | Follow the README literally: every day cell shows both, including `FUTURE`, `NO_ACTIVITY` and `WAITING_TO_SYNC` days |
| How the frontend knows "today" | From its own environment variable carrying the same `AS_OF` date. The API response stays exactly as the contract shows it |
| Server port | `8080` |
| `wear` and `sync` on `FUTURE` days | Hours `0` and count `0` |
| `month` parameter missing | `400` |
| Error body | `{"error": "message"}`: one `error` field holding a readable message |
| Empty state | A participant with no records, detected on the frontend as a grid with no activity. Its message appears alongside the grid. No change to the API |
| Today outside the selected month | Still marked as today |
| No month in the page URL | The page shows the month containing today |
| A month after today in the page URL | The page redirects to the month containing today |
| `month` format | Exactly `YYYY-MM`: four digits, a hyphen, two digits from `01` to `12`. Anything else returns `400` |
| `AS_OF` set but not a valid date | The server refuses to start |
| The frontend's date is not a valid date | The page shows an error screen |
| `AS_OF` not provided | Default to `2026-10-21`, the seed reference date, on both the server and the frontend |