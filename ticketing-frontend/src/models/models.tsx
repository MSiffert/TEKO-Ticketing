// src/models/index.ts

export interface ApiResponse<T> {
    response_key: string;
    response_message: string;
    data: T;
}


export interface BaseModel {
    createdAt?: string;
    updatedAt?: string;
    deletedAt?: string;
}

export interface Role extends BaseModel {
    id: number;
    role: string;
}

export enum TicketPriority {
    Low = 0,
    Medium = 1,
    High = 2,
    Urgent = 3,
}

export enum TicketStatus {
    Open = 0,
    InProgress = 1,
    Resolved = 2,
    Closed = 3,
}

export interface User extends BaseModel {
    id: number;
    name: string;
    email: string;
    role_id: number;
    role?: Role;
}

export interface Ticket extends BaseModel {
    id: number;
    title: string;
    description: string;
    status?: number;
    priority?: number;
    creator_user_id?: number;
    creator_user?: User;
    supporter_user_id?: number;
    supporter_user?: User;
    ticket_messages?: TicketMessage[]
}

export interface TicketMessage extends BaseModel {
    id: number;
    text: string;
    ticket_id: number;
    ticket: Ticket;
    creator_user_id: number;
    creator_user: User;
}
