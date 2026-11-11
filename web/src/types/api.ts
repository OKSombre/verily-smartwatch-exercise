export interface Participant {
  participantId: string;
  site: string;
  deviceId: string;
  enrolledOn: string;
}

export type DayState =
  | "FUTURE"
  | "NO_ACTIVITY"
  | "WAITING_TO_SYNC"
  | "MET"
  | "UNMET";

export interface Wear {
  hours: number;
  met: boolean;
}

export interface Sync {
  count: number;
  met: boolean;
}

export interface DailyRecord {
  date: string;
  state: DayState;
  wear: Wear;
  sync: Sync;
}

export interface MonthlyRecord {
  participantId: string;
  site: string;
  month: string;
  days: DailyRecord[];
}