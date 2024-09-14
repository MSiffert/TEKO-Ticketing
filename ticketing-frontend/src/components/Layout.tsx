// Layout.tsx
import React, {useEffect, useState} from 'react';
import {Outlet, Link, useNavigate, useLocation} from 'react-router-dom';
import { AppBar, Toolbar, Typography, Container, Button, Box } from '@mui/material';
import {User} from "../models/models";
import {getCurrentUser, logoutUser} from "../services/user";
import {fetchTicketById} from "../services/api";

const Layout: React.FC = () => {
    const [user, setUser] = useState<User | null>(null);
    const location = useLocation(); // Hook to get the current location (route)
    const navigate = useNavigate();

    useEffect(() => {
        const currentUser = getCurrentUser();

        if (currentUser) {
            setUser(currentUser);
        }
    }, [location]);

    const logout = async () => {
        logoutUser();
        setUser(null);
        navigate('/');
    };

    return (
        <>
            <AppBar position="static">
                <Toolbar>
                    <Typography variant="h6" style={{ flexGrow: 1 }}>
                        {
                            user && (
                                <Typography variant="h6" style={{ flexGrow: 1 }}>
                                    Ticketing, Logged in as { user.name }
                                </Typography>
                            )
                        }

                        {
                            !user && (
                                <Typography variant="h6" style={{ flexGrow: 1 }}>
                                    Ticketing
                                </Typography>
                            )
                        }
                    </Typography>

                    {
                        !user && (
                            <Button color="inherit" component={Link} to="/">Login</Button>
                        )
                    }

                    {
                        user && (
                            <>
                                <Button color="inherit" onClick={() => logout()}>Logout</Button>
                                <Button color="inherit" component={Link} to="/tickets">Tickets</Button>
                            </>
                        )
                    }
                </Toolbar>
            </AppBar>

            <Container maxWidth="lg" style={{ marginTop: '2rem' }}>
                <Outlet />
            </Container>
        </>
    );
};

export default Layout;
