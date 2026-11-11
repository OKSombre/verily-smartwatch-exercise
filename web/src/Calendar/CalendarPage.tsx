import Legend from './Legend'
import CalendarHeader from './CalendarHeader'
import MonthView from './MonthView'
import CalendarGrid from './CalendarGrid'
import type { MonthlyRecord } from '../types/api';
import { useEffect, useState } from 'react';
import { useParams, useSearchParams } from 'react-router-dom';
import { getMonthlyRecord } from '../services/api';

export default function Calendar() {
  const { participantId } = useParams()
  const [ searchParams, ] = useSearchParams()
  const [data, setData] = useState<MonthlyRecord | null>(null)

  useEffect(() => {
    const fetchMonthlyRecord = async () => {
      const record = await getMonthlyRecord(participantId ?? '', searchParams.get('month') ?? '')
      setData(record)
      console.log(record)
    }
    fetchMonthlyRecord()
  }, [participantId, searchParams]);


  return (
    <>
      <CalendarHeader participantId={data?.participantId ?? 'No Participant'} site={data?.site ?? 'No Site'} />
      <MonthView />
      <Legend />
      <CalendarGrid days={data?.days ?? []} />
      
    </>
  )
}