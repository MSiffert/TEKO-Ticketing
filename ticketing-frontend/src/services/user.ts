import {User} from "../models/models";

export function setCurrentUser(user: User) {
    localStorage.setItem('loggedInUser', JSON.stringify(user));
}

export function getCurrentUser() {
    const storedUser = localStorage.getItem('loggedInUser');
    if (storedUser) {
        return JSON.parse(storedUser) as User;
    }

    return null;
}

export function logoutUser() {
    const storedUser = localStorage.removeItem('loggedInUser');
}