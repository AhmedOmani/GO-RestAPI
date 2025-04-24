"use client"

import { useEffect, useState } from "react";
import { useRouter } from "next/navigation";
import styles from "../styles/Form.module.css"; // Reuse the same styles for consistency

export default function Landpage() {
    const [username, setUsername] = useState("");
    const [isAuthenticated, setIsAuthenticated] = useState(false);
    const router = useRouter();

    useEffect(() => {
        // Check if the user is authenticated and get the username
        const storedToken = localStorage.getItem("authToken");
        const storedUsername = localStorage.getItem("username"); // We'll store this during sign-in

        if (!storedToken || !storedUsername) {
            // If not authenticated, redirect to sign-in
            router.push("/signin");
        } else {
            setIsAuthenticated(true);
            setUsername(storedUsername);
        }
    }, [router]);

    const handleLogout = () => {
        // Clear authentication data
        localStorage.removeItem("authToken");
        localStorage.removeItem("isAuthenticated");
        localStorage.removeItem("username");
        setIsAuthenticated(false);
        router.push("/signin");
    };

    if (!isAuthenticated) {
        return null; // Prevent rendering until authentication check is complete
    }

    return (
        <div className={styles.container}>
            <div className={styles.formContainer}>
                <h1>Welcome, {username}!</h1>
                <p style={{ color: '#5D4037', textAlign: 'center', marginBottom: '2rem' }}>
                    You have successfully signed in.
                </p>
                <button
                    onClick={handleLogout}
                    className={styles.button}
                    style={{ width: '100%' }}
                >
                    Logout
                </button>
            </div>
        </div>
    );
}