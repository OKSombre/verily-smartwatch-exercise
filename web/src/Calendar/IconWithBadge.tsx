export interface IconWithBadgeProps {
  icon: React.ReactNode
  badge: React.ReactNode
}

export default function IconWithBadge({ icon, badge }: IconWithBadgeProps) {
  return (
    <div className="icon-with-badge">
      <div className="icon">
        {icon}
      </div>
      <div className="badge">
        {badge}
      </div>
    </div>
  )
}