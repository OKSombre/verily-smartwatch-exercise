import type { MonthlyRecord } from "../types/api";

export async function getMonthlyRecord(participantId: string, month: string): Promise<MonthlyRecord> {
  const request = new Request(
    `/api/studies/alpha/participants/${participantId}/calendar?month=${month}`
  );
  const response = await fetch(request);
  if (!response.ok) {
    throw new Error(`Failed to fetch monthly record: ${response.statusText}`);
  }
  const data: MonthlyRecord = await response.json();
  return data;
}