import { useEffect, useState } from 'react'
import './App.css'
import type { MonthlyRecord } from './types/api'
import { getMonthlyRecord } from './services.ts/api'
import { BrowserRouter, Route, Routes, useParams, useSearchParams } from 'react-router-dom'

function Home() {
  return <h1>Home</h1>
}

function Calendar() {
  const { participantId } = useParams()
  const [searchParams, setSearchParams] = useSearchParams()
  const monthString = searchParams.get('month')
  const month = monthString ? new Date(`${monthString}-12`) : new Date()
  console.log(monthString)
  console.log(month)
  
  const getYearMonthString = (date: Date) => {
    const year = date.getFullYear()
    const month = String(date.getMonth() + 1).padStart(2, '0')
    return `${year}-${month}`
  }

  const handleMonthChange = (monthOffset: number) => {
    const newMonth = new Date(month.getFullYear(), month.getMonth() + monthOffset)
    setSearchParams({ month: getYearMonthString(newMonth) })
  }

  return (
    <>
      <h1>Calendar for Participant {participantId} in Month {searchParams.get('month')}</h1>
      <button onClick={() => handleMonthChange(-1)}>Previous Month</button>
      <button onClick={() => handleMonthChange(1)}>Next Month</button>
    </>
  )
}

function App() {
  const [count, setCount] = useState(0)
  const [data, setData] = useState<MonthlyRecord | null>(null)

  useEffect(() => {
    const fetchMonthlyRecord = async () => {
      const record = await getMonthlyRecord('A1001', '2024-10')
      setData(record)
    }
    fetchMonthlyRecord()
  }, []);

  return (
    <>
      <BrowserRouter>
      <Routes>
        <Route path="/" element={<Home />} />
        <Route path="participants/:participantId/calendar" element={<Calendar />} />
      </Routes>
      </BrowserRouter>
    </>
  )
}

export default App
