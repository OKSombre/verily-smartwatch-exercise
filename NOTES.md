# Progress Notes on Take-Home Assessment
These are the notes I've written as I've progressed through this assignment. These contain my thoughts and processes for my coding journey these two days.

## Structure and AI Usage
To create the basic architecture and assure any spec ambiguities were covered, I established a base with Claude code and created a planning guide from start to finish of the project. All code was however created by myself, and I strictly prompted Claude to make no assumptions and clarify any ambiguity through me. The code is then proof-read through Claude for assurance of spec and checking against edge cases or vulnerabilities.

## Decisions and Trade-offs
Several changes to address ambiguity and best practices were made in order to ensure best code quality and best completion of spec. 

### API Stack
The choice to use net/http + Vite + React Router was made due to my own familiarity with the tools. However, the spec required that the AS_OF date be chosen through the environment, so I decided to use make dev include a variable that both the front-end and back-end can use, defaulting to 2026-10-21 if one isn't provided.

### Handling of Dates
I decided to use time.Time for the go back-end for sake of familiarity with the package as well as quick handling of date math. The trade-off is since this has to be agnostic of time zones, I need to ensure that every value uses UTC and that at no point the code uses local time zone data for calculations. For comparisons and other calculations, I can simply compare the dates as strings in YYYY-MM-DD format since time isn't a factor for the entries.

### Embedding Files vs Reading from Disk
When loading the seed data, I chose to embed the files rather than read directly through the disk. In my opinion this would be better since the data is relatively simple and has no need for changes. The only downside would be the binary size due to large amounts of data, but that wouldn't be an issue here.

### Possibility of Duplicate Daily Entries
As I was creating the store, I had made the assumption that each daily record was unique in participant and date. That way, I can immediately assume the DayState without further calculations (Ex. Adding together wear hours for multiple entries). Claude confirms that there are no duplicates in the data, but to be sure I plan to add a safeguard for the initial load.

### Front-end URL Handling
While adding the calendar grid to the front-end, I made some decisions on how to handle erroneous url parameters, such as invalid month or unknown participantId. For the invalid month, I decided to handle it as though the month had been set to the `AS_OF` month without changing the url. I plan to have the url set to redirect when this occurs. As of now, I also made the participant show no data when the id is unknown. The title header also changes to clearly show there is no participant with that id.

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
| `AS_OF` set but not a valid date | The server uses the default |
| The frontend's date is not a valid date | The page uses the `AS_OF` |
| `AS_OF` not provided | Default to `2026-10-21`, the seed reference date, on both the server and the frontend |