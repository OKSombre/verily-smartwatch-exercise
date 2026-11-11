export interface CalendarHeaderProps {
  participantId: string
  site: string
}

export default function CalendarHeader({ participantId, site }: CalendarHeaderProps) {

  return (
    <>
      <h1>Participant {participantId}</h1>
      <h2>Site: {site}</h2>
    </>
  )
}