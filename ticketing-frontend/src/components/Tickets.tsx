import React, { useState, useEffect } from 'react';
import {Paper, Typography, List, ListItem, CircularProgress, Container, Alert, Button} from '@mui/material';
import {ApiResponse, Ticket, TicketPriority, TicketStatus} from "../models/models";
import {fetchTickets} from "../services/api";
import {useNavigate} from "react-router-dom";

const Tickets: React.FC = () => {
    const [tickets, setTickets] = useState<Ticket[]>([]);
    const [loading, setLoading] = useState<boolean>(true);
    const [error, setError] = useState<string | null>(null);
    const navigate = useNavigate();

    const statusOptions = Object.keys(TicketStatus)
        .filter(key => !isNaN(Number(TicketStatus[key as keyof typeof TicketStatus])))
        .map(key => ({ label: key, value: TicketStatus[key as keyof typeof TicketStatus] }));

    const priorityOptions = Object.keys(TicketPriority)
        .filter(key => !isNaN(Number(TicketPriority[key as keyof typeof TicketPriority])))
        .map(key => ({ label: key, value: TicketPriority[key as keyof typeof TicketPriority] }));

    useEffect(() => {
        const loadTickets = async () => {
            try {
                const data: ApiResponse<Ticket[]> = await fetchTickets();
                setTickets(data.data);
                setLoading(false);
            } catch (error) {
                setError('Failed to fetch tickets');
                setLoading(false);
            }
        };

        loadTickets();
    }, []);

    return (
        <Container maxWidth="md" style={{ marginTop: '2rem' }}>
            <List>
                <ListItem>
                    <Typography variant={"h4"}>Tickets</Typography>
                </ListItem>
            </List>

            {loading && <CircularProgress />}
            {error && <Alert severity="error">{error}</Alert>}

            {
                !loading && tickets.length === 0 && (
                    <Typography variant="h4" color="textSecondary" gutterBottom>
                        There are no tickets.
                    </Typography>
                )
            }

            {!loading && !error && (
                <>
                    <List>
                        {tickets.map(ticket => (
                            <ListItem key={ticket.id} style={{ marginBottom: '1rem' }}>
                                <Paper elevation={3} style={{ padding: '1rem', width: '100%' }}>
                                    <Typography variant="h6" gutterBottom>
                                        {ticket.title}
                                    </Typography>
                                    <Typography variant="body1">
                                        {ticket.description}
                                    </Typography>
                                    <Typography variant="body2" color="textSecondary">
                                        Status: {statusOptions.filter(e => e.value === ticket.status)[0].label}
                                    </Typography>

                                    <Typography variant="body2" color="textSecondary">
                                        Priority: {priorityOptions.filter(e => e.value === ticket.priority)[0].label}
                                    </Typography>

                                    <Button
                                        variant="contained"
                                        color="primary"
                                        disabled={loading}
                                        onClick={() => navigate(`/ticket/${ticket.id}`)}
                                    >
                                        Open
                                    </Button>
                                </Paper>
                            </ListItem>
                        ))}
                    </List>

                    <List>
                        <ListItem style={{ marginBottom: '1rem' }}>
                            <Paper elevation={3} style={{ padding: '1rem', width: '100%' }}>
                                <Typography variant="body1" color="textSecondary" gutterBottom>Create Ticket</Typography>

                                <Button
                                    variant="contained"
                                    color="primary"
                                    disabled={loading}
                                    onClick={() => navigate('/createticket')}
                                >
                                    New Ticket
                                </Button>
                            </Paper>
                        </ListItem>
                    </List>
                </>
            )}
        </Container>
    );
};

export default Tickets;
