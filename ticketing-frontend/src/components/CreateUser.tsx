import React, { useState } from 'react';
import { TextField, Button, Container, Paper, Typography } from '@mui/material';
import {ApiResponse, Ticket, User} from '../models/models';
import {createTicket, createUser} from "../services/api";
import {useNavigate} from "react-router-dom"; // Adjust the import path as needed

const CreateUser: React.FC = () => {
    const [user, setUser] = useState<User>({
        id: 0, // id will be auto-generated
        name: '',
        email: '',
        role_id: 1,
    });

    const [error, setError] = useState<string | null>(null);
    const navigate = useNavigate();

    const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
        const { name, value } = e.target;
        setUser(prevState => ({
            ...prevState,
            [name]: value
        }));
    };

    const handleSubmit = async (e: React.FormEvent) => {
        e.preventDefault();
        try {
            user.role_id = Number(user.role_id);
            const data: ApiResponse<User> = await createUser(user);

            navigate("/");
        } catch (err) {
            setError('Failed to create user.');
        }
    };

    return (
        <Container maxWidth="sm" style={{ marginTop: '2rem' }}>
            <Paper elevation={3} style={{ padding: '1.5rem' }}>
                <Typography variant="h5" gutterBottom>
                    Create New User
                </Typography>
                {error && <Typography color="error">{error}</Typography>}
                <form onSubmit={handleSubmit}>
                    <TextField
                        label="Name"
                        name="name"
                        value={user.name}
                        onChange={handleChange}
                        fullWidth
                        margin="normal"
                        required
                    />
                    <TextField
                        label="Email"
                        name="email"
                        type="email"
                        value={user.email}
                        onChange={handleChange}
                        fullWidth
                        margin="normal"
                        required
                    />
                    <TextField
                        label="Role ID"
                        name="role_id"
                        type="number"
                        value={user.role_id}
                        onChange={handleChange}
                        fullWidth
                        margin="normal"
                        required
                    />
                    <Button
                        type="submit"
                        variant="contained"
                        color="primary"
                        fullWidth
                        style={{ marginTop: '1rem' }}
                    >
                        Create User
                    </Button>
                </form>
            </Paper>
        </Container>
    );
};

export default CreateUser;
