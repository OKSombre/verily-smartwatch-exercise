import { MONTH_NAMES } from "./constants"

export function firstSunday(month: string): string {
  const date = new Date(`${month}-05`)
  date.setDate(date.getDate() - date.getDay())
  return date.toISOString().slice(0, 10)
}

export function lastSaturday(month: string): string {
  const date = new Date(`${month}-05`)
  date.setDate(date.getDate() + (6 - date.getDay()))
  return date.toISOString().slice(0, 10)
}

export function getCalendarDate(date: string): string {
  const month = MONTH_NAMES[Number(date.slice(5, 7)) - 1]
  return `${month} ${date.slice(8, 10)} ${date.slice(0, 4)}`
}