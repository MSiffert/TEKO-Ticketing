import React, { useEffect, useState } from 'react';
import { useParams } from 'react-router-dom';
import {Paper, Typography, CircularProgress, Alert, Container, TextField, Button, MenuItem} from '@mui/material';
import {Ticket, TicketMessage, TicketPriority, TicketStatus, User} from "../models/models";
import {
    fetchTicketById,
    createTicketMessage,
    fetchRoles,
    fetchUsers,
    updateTicket
} from "../services/api";
import {getCurrentUser} from "../services/user";

const TicketDetails: React.FC = () => {
    const { ticketId } = useParams<{ ticketId: string }>(); // Get ticketId from route params
    const [ticket, setTicket] = useState<Ticket | null>(null);
    const [loading, setLoading] = useState<boolean>(true);
    const [error, setError] = useState<string | null>(null);
    const [message, setMessage] = useState<string>(''); // Message input state
    const [messageError, setMessageError] = useState<string | null>(null); // Message form error state
    const [messages, setMessages] = useState<TicketMessage[]>([]); // Ticket messages state
    const [user, setUser] = useState<User | null>(null);
    const [supporter, setSupporter] = useState<User[] | null>(null);

    const statusOptions = Object.keys(TicketStatus)
        .filter(key => !isNaN(Number(TicketStatus[key as keyof typeof TicketStatus])))
        .map(key => ({ label: key, value: TicketStatus[key as keyof typeof TicketStatus] }));

    const priorityOptions = Object.keys(TicketPriority)
        .filter(key => !isNaN(Number(TicketPriority[key as keyof typeof TicketPriority])))
        .map(key => ({ label: key, value: TicketPriority[key as keyof typeof TicketPriority] }));

    useEffect(() => {
        const currentUser = getCurrentUser();

        if (currentUser) {
            setUser(currentUser);
        }
    }, []);

    useEffect(() => {
        loadSupporters();
        loadTicket();
    }, [ticketId]);

    const loadTicket = async () => {
        try {
            const data = await fetchTicketById(Number(ticketId));
            setTicket(data.data);
            setMessages(data.data?.ticket_messages ?? []);
        } catch (error) {
            setError('Failed to load ticket details.');
        } finally {
            setLoading(false);
        }
    };

    const handleChangeAssignedSupporter = (event: React.ChangeEvent<{ value: unknown }>) => {
        const selectedUser = supporter?.find(supporter => supporter.id === event.target.value);
        if (selectedUser && ticket) {
            setTicket({ ...ticket, supporter_user_id: selectedUser.id });
        }
    };

    const handlePriorityChange = (e: React.ChangeEvent<HTMLInputElement>) => {
        if (ticket) {
            setTicket({ ...ticket, priority: parseInt(e.target.value) });
        }
    };

    const handleStatusChange = (e: React.ChangeEvent<HTMLInputElement>) => {
        if (ticket) {
            setTicket({ ...ticket, status: parseInt(e.target.value) });
        }
    };

    const loadSupporters = async () => {
        const users = await fetchUsers(1);
        setSupporter(users.data);
    }

    const handleAddMessage = async () => {
        if (message.trim() === '') {
            setMessageError('Message cannot be empty.');
            return;
        }

        try {
            const newMessage: TicketMessage = await createTicketMessage({
                ticket_id: Number(ticketId),
                text: message,
                creator_user_id: user?.id ?? 0
            });
            setMessages((prevMessages) => [...prevMessages, newMessage]);
            setMessage('');
            setMessageError(null);

            loadTicket();
        } catch (error) {
            setMessageError('Failed to add message.');
        }
    };

    if (loading) {
        return <CircularProgress />;
    }

    if (error) {
        return <Alert severity="error">{error}</Alert>;
    }

    return (
        <Container maxWidth="sm" style={{ marginTop: '2rem' }}>
            {ticket && (
                <>
                    <Paper elevation={3} style={{ padding: '2rem', marginBottom: '2rem' }}>
                        <Typography variant="h4" gutterBottom>
                            Ticket Details
                        </Typography>
                        <Typography variant="h6">
                            Title: {ticket.title}
                        </Typography>
                        <Typography variant="body1">
                            Description: {ticket.description}
                        </Typography>
                        <Typography variant="body2" color="textSecondary">
                            Status: {statusOptions.filter(e => e.value === ticket.status)[0].label}
                        </Typography>
                        <Typography variant="body2" color="textSecondary">
                            Priority: {priorityOptions.filter(e => e.value === ticket.priority)[0].label}
                        </Typography>
                    </Paper>

                    {
                        user?.role?.id === 1 && (
                            <Paper elevation={3} style={{ padding: '2rem', marginBottom: '2rem' }}>
                                <Typography variant="h5" gutterBottom>
                                    Update Ticket
                                </Typography>

                                <TextField
                                    select
                                    label="Select Supporter"
                                    value={String(ticket.supporter_user_id || '')}
                                    onChange={handleChangeAssignedSupporter}
                                    fullWidth
                                    variant="outlined"
                                    style={{ marginBottom: '1rem' }}
                                >
                                    {supporter?.map((supporter) => (
                                        <MenuItem key={supporter.id} value={supporter.id}>
                                            {supporter.name}
                                        </MenuItem>
                                    ))}
                                </TextField>

                                <TextField
                                    select
                                    label="Status"
                                    value={ticket.status}
                                    onChange={handleStatusChange}
                                    fullWidth
                                    variant="outlined"
                                    SelectProps={{ native: true }}
                                    style={{ marginBottom: '1rem' }}
                                >
                                    {statusOptions.map((option) => (
                                        <option key={option.value} value={option.value}>
                                            {option.label}
                                        </option>
                                    ))}
                                </TextField>

                                <TextField
                                    select
                                    label="Priority"
                                    value={ticket.priority}
                                    onChange={handlePriorityChange}
                                    fullWidth
                                    variant="outlined"
                                    SelectProps={{ native: true }}
                                    style={{ marginBottom: '1rem' }}
                                >
                                    {priorityOptions.map((option) => (
                                        <option key={option.value} value={option.value}>
                                            {option.label}
                                        </option>
                                    ))}
                                </TextField>

                                <Button
                                    variant="contained"
                                    color="primary"
                                    onClick={() => updateTicket(ticket)}
                                >
                                    Save Changes
                                </Button>
                            </Paper>
                        )
                    }

                    {/* Display Messages */}
                    <Paper elevation={3} style={{ padding: '2rem', marginBottom: '2rem' }}>
                        <Typography variant="h5" gutterBottom>
                            Messages
                        </Typography>
                        {messages.length > 0 ? (
                            messages.map((msg) => (
                                <Paper key={msg.id} elevation={1} style={{ padding: '1rem', marginBottom: '1rem' }}>
                                    <Typography variant="h6">By: {msg.creator_user?.name}</Typography>
                                    <Typography variant="body1">{msg.text}</Typography>
                                </Paper>
                            ))
                        ) : (
                            <Typography variant="body2" color="textSecondary">
                                No messages yet.
                            </Typography>
                        )}
                    </Paper>

                    {/* Add Message Section */}
                    <Paper elevation={3} style={{ padding: '2rem' }}>
                        <Typography variant="h5" gutterBottom>
                            Add a Message
                        </Typography>

                        <TextField
                            label="Message"
                            multiline
                            rows={4}
                            fullWidth
                            value={message}
                            onChange={(e) => setMessage(e.target.value)}
                            variant="outlined"
                            error={!!messageError}
                            helperText={messageError}
                            style={{ marginBottom: '1rem' }}
                        />

                        <Button
                            variant="contained"
                            color="primary"
                            onClick={handleAddMessage}
                        >
                            Add Message
                        </Button>
                    </Paper>
                </>
            )}
        </Container>
    );
};

export default TicketDetails;
