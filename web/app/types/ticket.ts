export enum TicketState {
  UNUSED = 'unused',
  USED = 'used',
};

export enum TicketType {
  ENTRY = 'entry',
  CREW = 'crew',
};

export default interface Ticket {
  id: string,
  type: TicketType,
  state: TicketState,
  value: string,
  email: string,
  created_at: Date,
  updated_at: Date,
  deleted_at: Date,
};
