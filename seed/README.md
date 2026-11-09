# Seed data

Two files, loaded however you like (read at startup, embed, copy into your own structure).

## `participants.json`

```json
{ "participantId": "A1001", "deviceId": "watch-x8.QX7M42RJ2KR", "site": "Site 1", "enrolledOn": "2026-09-20" }
```

- `deviceId` is `null` when the participant has no watch assigned.
- `enrolledOn` is the calendar date the participant enrolled. The calendar does not need it.

## `daily.json`

```json
{ "participantId": "A1001", "date": "2026-10-12", "wearHours": 10.0, "syncCount": 1 }
```

- `wearHours` is hours of watch wear that day (0 to 24, one decimal place).
- `syncCount` is the number of successful syncs that day.
- **A day with no record means `wearHours = 0` and `syncCount = 0`.** Records exist only for days
  with data.
- Dates are plain calendar dates in the participant's local time. Don't do time zone math.

## Reference date

"Today" (`asOf`) is `2026-10-21`. Your server must take it from the `AS_OF` environment variable
rather than the system clock. The data contains no records after that date.

## Notes

- The data is synthetic.
- Read the edge cases in the rules carefully. The data includes participants at both sides of each
  boundary in the rules. Writing your own tests against this data is encouraged.
