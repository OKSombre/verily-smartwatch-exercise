import { FaCheck, FaClock, FaExclamation, FaRegSquare, FaSync } from "react-icons/fa"
import IconWithBadge from "./IconWithBadge"

interface LegendItemProps {
  icon: React.ReactNode
  label: string
}

function LegendItem({ icon, label }: LegendItemProps) {
  return (
    <div className="calendar-legend-item">
      {icon}
      {label}
    </div>
  )
}

export default function Legend() {
  const MetIcon = () => (
    <IconWithBadge
      icon={<FaRegSquare style={{ backgroundColor: "green", opacity: 0.75 }}/>}
      badge={<FaCheck style={{ color: "green" }}/>}
    />
  )

  const UnmetIcon = () => (
    <IconWithBadge
      icon={<FaRegSquare style={{ backgroundColor: "orange", opacity: 0.75 }}/>}
      badge={<FaExclamation style={{ color: "orange" }}/>}
    />
  )

  return (
    <div className="calendar-legend">
      <div>Legend:</div>
      <LegendItem icon={<FaClock/>} label="Watch wear" />
      <LegendItem icon={<FaSync/>} label="Watch sync" />
      <LegendItem icon={<MetIcon/>} label="Met" />
      <LegendItem icon={<UnmetIcon/>} label="Unmet" />
    </div>
  )
}