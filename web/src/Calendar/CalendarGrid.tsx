import { WEEKDAYS } from '../utils/constants'
import type { DailyRecord } from '../types/api'
import CalendarCell from './CalendarCell'

export interface CalendarGridProps {
  days: DailyRecord[]
}

export default function CalendarGrid({ days }: CalendarGridProps) {

  return (
    <div className="calendar-grid" role="grid" aria-labelledby='calendar-month'>
      {WEEKDAYS.map((weekday) => (
        <div key={weekday} className="weekday-label">{weekday}</div>
      ))}
      {days.map((day) => (
        <CalendarCell
          key={day.date}
          date={day.date}
          dayState={day.state}
          watchHoursMet={day.wear.met}
          syncCountMet={day.sync.met}
        />
      ))}
    </div>
  )
}