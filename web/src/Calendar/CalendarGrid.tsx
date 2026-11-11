import { useSearchParams } from 'react-router-dom'

import { WEEKDAYS } from '../utils/constants'
import type { DailyRecord } from '../types/api'

export interface CalendarGridProps {
  days: DailyRecord[]
}

export default function CalendarGrid({ days }: CalendarGridProps) {

  return (
    <div className="calendar-grid">
      {WEEKDAYS.map((weekday) => (
        <div key={weekday} className="weekday-label">{weekday}</div>
      ))}
      {days.map((day) => (
        <div 
          key={day.date} 
          className={`calendar-cell ${day.state === 'MET' ? 'active-day' : 'empty-cell'} ${day.date === new Date().toISOString().split('T')[0] ? 'today' : ''}`}
        >
          {day.date}
        </div>
      ))}
    </div>
  )
}