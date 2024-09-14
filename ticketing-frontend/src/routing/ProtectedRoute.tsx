import React, { ReactNode } from 'react';
import { Navigate, useLocation } from 'react-router-dom';
import {getCurrentUser} from "../services/user";

interface ProtectedRouteProps {
    children: ReactNode;
}

const ProtectedRoute: React.FC<ProtectedRouteProps> = ({ children }) => {
    const location = useLocation();
    const isAuthenticated = getCurrentUser() !== null;

    if (!isAuthenticated) {
        return <Navigate to="/" state={{ from: location }} replace />;
    }

    return <>{children}</>;
};

export default ProtectedRoute;
