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