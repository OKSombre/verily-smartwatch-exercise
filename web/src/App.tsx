import './App.css'
import { BrowserRouter, Route, Routes } from 'react-router-dom'
import Calendar from './Calendar/CalendarPage'
  

function Home() {
  return <h1>Home</h1>
}

function App() {

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
