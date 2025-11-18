export enum OrderState {
  PENDING = 'pending',
  PAID = 'paid',
  CANCELLED = 'cancelled',
};

export default interface Order {
  id: string,
  ticket_id: string,
  reference: string,
  state: OrderState,
  email: string,
  created_at: Date,
  updated_at: Date,
  deleted_at: Date,
};
