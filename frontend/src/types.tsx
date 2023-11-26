export enum Priority {
  low = "low",
  medium = "medium",
  high = "high"
};


export type Ticket = {
  id?: string,
  firstName: string,
  lastName: string,
  email: string,
  issue: string,
  priority?: string,
};


export type TicketItemProps = {
  key?: number,
  ticket: Ticket
};

