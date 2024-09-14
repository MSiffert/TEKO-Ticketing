import {ApiResponse, TicketPriority, Role, TicketStatus, Ticket, User} from "../models/models";

const API_BASE_URL = "http://localhost:8080/api";

export const fetchUsers = async (filterByRoleId?: number): Promise<ApiResponse<User[]>> => {
    if (filterByRoleId) {
        const response = await fetch(`${API_BASE_URL}/users?role_id=${filterByRoleId}`);
        if (!response.ok) {
            throw new Error('Network response was not ok');
        }
        const data: ApiResponse<User[]> = await response.json();
        return data;
    } else {
        const response = await fetch(`${API_BASE_URL}/users`);
        if (!response.ok) {
            throw new Error('Network response was not ok');
        }
        const data: ApiResponse<User[]> = await response.json();
        return data;
    }
};

export const createUser = async (userData: User): Promise<ApiResponse<User>> => {
    const response = await fetch(`${API_BASE_URL}/users`, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(userData),
    });

    if (!response.ok) {
        // Handle non-2xx HTTP responses
        const errorText = await response.text();
        throw new Error(`Failed to create user: ${errorText}`);
    }

    const data: ApiResponse<User> = await response.json();
    return data;
};

export const fetchTickets = async (): Promise<ApiResponse<Ticket[]>> => {
    const response = await fetch(`${API_BASE_URL}/tickets`);
    if (!response.ok) {
        throw new Error('Network response was not ok');
    }
    const data: ApiResponse<Ticket[]> = await response.json();
    return data;
};

export const fetchRoles = async (): Promise<ApiResponse<Role[]>> => {
    const response = await fetch(`${API_BASE_URL}/roles`);
    if (!response.ok) {
        throw new Error('Network response was not ok');
    }
    const data: ApiResponse<Role[]> = await response.json();
    return data;
};

export const createTicket = async (ticketData: Ticket): Promise<ApiResponse<Ticket>> => {
    const response = await fetch(`${API_BASE_URL}/tickets`, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(ticketData),
    });

    if (!response.ok) {
        // Handle non-2xx HTTP responses
        const errorText = await response.text();
        throw new Error(`Failed to create ticket: ${errorText}`);
    }

    const data: ApiResponse<Ticket> = await response.json();
    return data;
};

export const fetchTicketById = async (ticketId: number): Promise<ApiResponse<Ticket>> => {
    const response = await fetch(`${API_BASE_URL}/tickets/${ticketId}`);
    if (!response.ok) {
        throw new Error('Network response was not ok');
    }
    const data: ApiResponse<Ticket> = await response.json();
    return data;
};

export const createTicketMessage = async (messageData: { ticket_id: number, text: string, creator_user_id: number }) => {
    const response = await fetch(`${API_BASE_URL}/ticketmessages`, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(messageData),
    });

    if (!response.ok) {
        throw new Error('Failed to send message');
    }

    return response.json(); // Assuming the response contains some data or confirmation
};

export const updateTicket = async (ticket: Ticket) => {
    ticket.status = 2;
    const response = await fetch(`${API_BASE_URL}/tickets/${ticket.id}`, {
        method: 'PUT',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(ticket)
    });
    return response.json();
};

export const deleteTicketMessage = async (ticketMessageId: string) => {
    const response = await fetch(`${API_BASE_URL}/ticketmessages/${ticketMessageId}`, {
        method: 'DELETE',
    });
    return response.json();
};
