# Progress Notes on Take-Home Assessment
These are the notes I've written as I've progressed through this assignment. These contain my thoughts and processes for my coding journey these two days.

## How to run
The app is set to run the makefile with `make dev`. An added parameter for `AS_OF` can be included to run dates in `YYYY-MM-DD` format. Any other format will not be accepted, and will default to `2026-10-21` if none is provided. To run the frontend seperately, run `npm run dev --prefix ./web` and either create an .env file with `VITE_AS_OF` set or set the env variable in the command line. To run the backend separately, run `go run .\api\main.go` with the optional argument of an `AS_OF` date in `YYYY-MM-DD` format (any invalid format will default to `2026-10-21`). 

Backend API: http://localhost:8080/api/studies/alpha/participants/{participantID}/calendar?month={YYYY-MM}

Client: http://localhost:5173/participants/{participantID}/calendar?month={YYYY-MM}

Tests for the backend can be run simply by running `go test .\api\internal\handler .\api\internal\builder\ .\api\internal\utils\`. There are no tests for the frontend written.

## Structure and AI Usage
To create the basic architecture and assure any spec ambiguities were covered, I established a base with Claude code and created a planning guide from start to finish of the project. All code was however created by myself, and I strictly prompted Claude to make no assumptions and clarify any ambiguity through me. The code is then proof-read through Claude for assurance of spec and checking against edge cases or vulnerabilities.

## Decisions and Trade-offs
Several changes to address ambiguity and best practices were made in order to ensure best code quality and best completion of spec. 

### API Stack
The choice to use net/http + Vite + React Router was made due to my own familiarity with the tools. However, the spec required that the `AS_OF` date be chosen through the environment, so I decided to use make dev include a variable that both the front-end and back-end can use, defaulting to 2026-10-21 if one isn't provided.

### Handling of Dates
I decided to use time.Time for the go back-end for sake of familiarity with the package as well as quick handling of date math. The trade-off is since this has to be agnostic of time zones, I need to ensure that every value uses UTC and that at no point the code uses local time zone data for calculations. For comparisons and other calculations, I can simply compare the dates as strings in YYYY-MM-DD format since time isn't a factor for the entries.

### Embedding Files vs Reading from Disk
When loading the seed data, I chose to embed the files rather than read directly through the disk. In my opinion this would be better since the data is relatively simple and has no need for changes. The only downside would be the binary size due to large amounts of data, but that wouldn't be an issue here.

### Possibility of Duplicate Daily Entries
As I was creating the store, I had made the assumption that each daily record was unique in participant and date. That way, I can immediately assume the DayState without further calculations (Ex. Adding together wear hours for multiple entries). Claude confirms that there are no duplicates in the data, but to be sure I plan to add a safeguard for the initial load.

### Front-end URL Handling
While adding the calendar grid to the front-end, I made some decisions on how to handle erroneous url parameters, such as invalid month or unknown participantId. For the invalid month, I decided to handle it as though the month had been set to the `AS_OF` month without changing the url. I plan to have the url set to redirect when this occurs. As of now, I also made the participant show no data when the id is unknown. The title header also changes to clearly show there is no participant with that id.

## What would I do if I had another day
With another day, I would have  written a way to change participants in-app, and way to interact with the calendar days to see details of user wear hours and sync count. There would also need to be better security for edge cases. As it stands, the app doesn't run if the `AS_OF` date is invalid for the front end. The url also needed to be checked and a redirect should be added for invalid or future dates to go to the most recent `AS_OF` month. I also did not have access to a screen reader or accessibilty software, so my aria-label work is mainly guesswork and would need further testing. I would also have liked to make robust tests for the frontend, something like storybook for the components, and tailwind instead of plain css for light and dark theme switching.

## Spec Ambiguity and Clarifications
These are ambiguities that Claude had discovered and I made decisions for them.

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