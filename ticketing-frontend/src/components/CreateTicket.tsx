import React, {useEffect, useState} from 'react';
import { TextField, Button, Container, Typography, Paper, CircularProgress, Alert } from '@mui/material';
import {ApiResponse, Ticket, User} from "../models/models";
import {createTicket} from "../services/api";
import {getCurrentUser} from "../services/user";
import {useNavigate} from "react-router-dom";

const CreateTicket: React.FC = () => {
    const [title, setTitle] = useState<string>('');
    const [description, setDescription] = useState<string>('');
    const [loading, setLoading] = useState<boolean>(false);
    const [error, setError] = useState<string | null>(null);
    const [success, setSuccess] = useState<string | null>(null);
    const [user, setUser] = useState<User | null>(null);
    const navigate = useNavigate();

    useEffect(() => {
        const currentUser = getCurrentUser();

        if (currentUser) {
            setUser(currentUser);
        }
    }, []);

    const handleSubmit = async (event: React.FormEvent) => {
        event.preventDefault();

        setLoading(true);
        setError(null);
        setSuccess(null);

        try {
            const ticketData: Ticket = {
                id: 0,
                title: title,
                description: description,
                creator_user_id: user?.id
            };

            const data: ApiResponse<Ticket> = await createTicket(ticketData);
            setSuccess('Ticket created successfully!');
            setTitle('');
            setDescription('');
            navigate("/ticket/" + data.data.id)
        } catch (error) {
            setError('Failed to create ticket');
        } finally {
            setLoading(false);
        }
    };

    return (
        <Container maxWidth="sm" style={{ marginTop: '2rem' }}>
            <Paper elevation={3} style={{ padding: '2rem' }}>
                <Typography variant="h5" gutterBottom>
                    Create New Ticket
                </Typography>
                <form onSubmit={handleSubmit}>
                    <TextField
                        label="Title"
                        variant="outlined"
                        fullWidth
                        margin="normal"
                        value={title}
                        onChange={(e) => setTitle(e.target.value)}
                        required
                    />
                    <TextField
                        label="Description"
                        variant="outlined"
                        fullWidth
                        margin="normal"
                        value={description}
                        onChange={(e) => setDescription(e.target.value)}
                        multiline
                        rows={4}
                        required
                    />
                    <Button
                        type="submit"
                        variant="contained"
                        color="primary"
                        style={{ marginTop: '1rem' }}
                        disabled={loading}
                    >
                        {loading ? <CircularProgress size={24} /> : 'Create Ticket'}
                    </Button>
                    {error && <Alert severity="error" style={{ marginTop: '1rem' }}>{error}</Alert>}
                    {success && <Alert severity="success" style={{ marginTop: '1rem' }}>{success}</Alert>}
                </form>
            </Paper>
        </Container>
    );
};

export default CreateTicket;
