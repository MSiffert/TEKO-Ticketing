import React, { useState, useEffect } from 'react';
import { Paper, Typography, List, ListItem, Button, Container, CircularProgress, Alert } from '@mui/material';
import {useNavigate} from "react-router-dom";
import {ApiResponse, User} from "./models/models";
import {fetchUsers} from "./services/api";

const App: React.FC = () => {
    const [users, setUsers] = useState<User[]>([]);
    const [loading, setLoading] = useState<boolean>(true);
    const [error, setError] = useState<string | null>(null);
    const navigate = useNavigate();

    const handleUserSelect = (user: User) => {
        localStorage.setItem('loggedInUser', JSON.stringify(user));
        navigate('/tickets');
    };

    useEffect(() => {
        const loadUsers = async () => {
            try {
                const data: ApiResponse<User[]> = await fetchUsers();
                setUsers(data.data);
                setLoading(false);
            } catch (error) {
                setError('Failed to fetch users');
                setLoading(false);
            }
        };

        loadUsers();
    }, []);

    return (
        <Container maxWidth="md" style={{ marginTop: '2rem' }}>
            <List>
                <ListItem>
                    <Typography variant={"h4"}>Login as a user</Typography>
                </ListItem>
            </List>

            {loading && <CircularProgress />}
            {error && <Alert severity="error">{error}</Alert>}
            {!loading && !error && (
                <>
                    <List>
                        {users.map(user => (
                            <ListItem key={user.id} style={{ marginBottom: '1rem' }}>
                                <Paper elevation={3} style={{ padding: '1rem', width: '100%' }}>
                                    <Typography variant="h6" gutterBottom>
                                        {user.name}
                                    </Typography>
                                    <Typography variant="body1">
                                        Email: {user.email}
                                    </Typography>
                                    <Typography variant="body2" color="textSecondary">
                                        Role: {user.role?.role}
                                    </Typography>
                                    <Button
                                        variant="contained"
                                        color="primary"
                                        style={{ marginTop: '1rem' }}
                                        onClick={() => handleUserSelect(user)}
                                    >
                                        Select
                                    </Button>
                                </Paper>
                            </ListItem>
                        ))}
                    </List>

                    <List>
                        <ListItem>
                            <Paper elevation={3} style={{ padding: '1rem', width: '100%' }}>
                                <Button
                                    variant="contained"
                                    color="primary"
                                    onClick={() => navigate("createuser")}
                                >
                                    Create User
                                </Button>
                            </Paper>
                        </ListItem>
                    </List>
                </>

            )}
        </Container>
    );
};

export default App;
