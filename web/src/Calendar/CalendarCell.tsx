import { useSearchParams } from "react-router-dom";
import type { DayState } from "../types/api";
import IconWithBadge from "./IconWithBadge";
import { FaCheck, FaClock, FaExclamation, FaSync } from "react-icons/fa";
import { getCalendarDate } from "../utils/utils";

const asOf = import.meta.env.VITE_AS_OF;

export interface CalendarCellProps {
  date: string;
  dayState: DayState;
  watchHoursMet: boolean;
  syncCountMet: boolean;
}

interface CellInfoProps {
  dayState: DayState;
  watchHoursMet: boolean;
  syncCountMet: boolean;
}

function getDayStateStatus(dayState: DayState) {
    switch (dayState) {
      case "WAITING_TO_SYNC":
        return "Waiting to sync"
      case "NO_ACTIVITY":
        return "No activity"
      case "FUTURE":
        return "Future"
      case "MET":
        return "Met"
      case "UNMET":
        return "Unmet"
      default:
        return ""
    }
}

function CellInfo({
  dayState,
  watchHoursMet,
  syncCountMet,
}: CellInfoProps) {
  const WatchHoursIcon = () => (
    <IconWithBadge 
      icon={<FaClock style={{ color: watchHoursMet ? "green" : "orange" }} />}
      badge={watchHoursMet ? <FaCheck style={{ color: "green" }} /> : <FaExclamation style={{ color: "orange" }} />}
    />
  )
  const SyncCountIcon = () => (
    <IconWithBadge
      icon={<FaSync style={{ color: syncCountMet ? "green" : "orange" }} />}
      badge={syncCountMet ? <FaCheck style={{ color: "green" }} /> : <FaExclamation style={{ color: "orange" }} />}
    />
  )

  const DayStateStatus = () => (
    <span className="day-state-status">{getDayStateStatus(dayState)}</span>
  )
  
  return (
    <div className="calendar-cell-info">
      {dayState === "MET" || dayState === "UNMET" ? (
        <>
          <WatchHoursIcon />
          <SyncCountIcon />
        </>
      ) : dayState !== "FUTURE" ? <DayStateStatus /> : null}
      
    </div>
  )
}

export default function CalendarCell({ 
  date, 
  dayState, 
  watchHoursMet, 
  syncCountMet 
}: CalendarCellProps) {
  const [searchParams, ] = useSearchParams()
  const month = searchParams.get("month") ?? ""
  

  const cellClassName = () => {
    let className = "calendar-cell"
    if (dayState === "MET") {
      className += " met"
    } else if (dayState === "UNMET") {
      className += " unmet"
    }
    return className
  }
  
  const spanClassName = () => {
    let className = "calendar-cell-span"
    console.log(date.slice(0, 7))
    console.log(month)
    if (date.slice(0, 7) !== month) {
      className += " other-month"
    } else {
      className += " active-month"
    }
    if (date === asOf) {
      className += " today"
    }
    return className
  }

  const cellAriaLabel = () => {
    let status = `${getCalendarDate(date)}: Status is ${getDayStateStatus(dayState)}`
    if (dayState === "MET" || dayState === "UNMET") {
      status += `, Watch hours ${watchHoursMet ? "met" : "not met"}, Sync count ${syncCountMet ? "met" : "not met"}`
    }
    return status
  }

  return (
    <div className={cellClassName()}>
      <span className={spanClassName()}>{date.slice(8, 10)}</span>
      <CellInfo
        dayState={dayState}
        watchHoursMet={watchHoursMet}
        syncCountMet={syncCountMet}
      />
      <span style={{display: "none"}}>{cellAriaLabel()}</span>
    </div>
  )
}