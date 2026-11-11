import { FaArrowLeft } from "react-icons/fa"

export interface CalendarHeaderProps {
  participantId: string
  site: string
}

export default function CalendarHeader({ participantId, site }: CalendarHeaderProps) {

  return (
    <div className="calendar-header">
      <button aria-label="Back"><FaArrowLeft /></button>
      <div>
        <h1>Participant {participantId}</h1>
        <div>Site: {site}</div>
      </div>
    </div>
  )
}