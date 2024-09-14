import React from 'react';
import ReactDOM from 'react-dom';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import App from './App';
import ManageTickets from "./components/Tickets";
import CreateTicket from "./components/CreateTicket";
import Tickets from "./components/Tickets";
import TicketDetails from "./components/TicketDetails";
import CreateUser from "./components/CreateUser";
import Layout from "./components/Layout"; // Adjust the path as needed
import './index.css';
import ProtectedRoute from "./routing/ProtectedRoute";

ReactDOM.render(
    <Router>
        <Routes>
            <Route path="/" element={<Layout />}>
                <Route path="/" element={<App />} />
                <Route path="/tickets" element={<ProtectedRoute><Tickets /></ProtectedRoute>} />
                <Route path="/createticket" element={<ProtectedRoute><CreateTicket /></ProtectedRoute>} />
                <Route path="/createuser" element={<CreateUser />} />
                <Route path="/ticket/:ticketId" element={<ProtectedRoute><TicketDetails /></ProtectedRoute>} /> {/* New route */}
            </Route>

        </Routes>
    </Router>,
document.getElementById('root')
);
