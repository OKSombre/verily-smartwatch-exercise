import { useSearchParams } from 'react-router-dom'
import { FaChevronLeft, FaChevronRight } from 'react-icons/fa'


const asOf = import.meta.env.VITE_AS_OF

function isValidMonthString(monthString: string) {
  return /^\d{4}-(0[1-9]|1[0-2])$/.test(monthString)
}

const MONTH_NAMES: string[] = [
  'January',
  'February',
  'March',
  'April',
  'May',
  'June',
  'July',
  'August',
  'September',
  'October',
  'November',
  'December',
]

export default function MonthView() {
  const [searchParams, setSearchParams] = useSearchParams()
  const monthString = searchParams.get('month')
  const month = monthString && isValidMonthString(monthString) ? new Date(`${monthString}-21`) : new Date(asOf)
  
  const getYearMonthString = (date: Date) => {
    const year = date.getFullYear()
    const month = String(date.getMonth() + 1).padStart(2, '0')
    return `${year}-${month}`
  }

  const monthName = MONTH_NAMES[month.getMonth()]

  const isAsOfMonth = getYearMonthString(month) === asOf.slice(0, 7)

  const handleMonthChange = (monthOffset: number) => {
    const newMonth = new Date(month.getFullYear(), month.getMonth() + monthOffset)
    if (getYearMonthString(newMonth) > asOf.slice(0, 7)) {
      return
    }
    setSearchParams({ month: getYearMonthString(newMonth) })
  }

  return (
    <div className="calendar-month-view">
      <button aria-label="Previous month" onClick={() => handleMonthChange(-1)}><FaChevronLeft /></button>
      <button aria-label="Next month" onClick={() => handleMonthChange(1)} disabled={isAsOfMonth}><FaChevronRight /></button>
      <h2 id="calendar-month">{monthName} {month.getFullYear()}</h2>
    </div>
  )
}