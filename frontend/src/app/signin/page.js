"use client"

import { useState, useEffect } from "react"
import Link from "next/link"
import { useRouter } from "next/navigation"
import styles from "../styles/Form.module.css"

export default function SigninPage() {
    const [email, setEmail] = useState("");
    const [password, setPassword] = useState("");
    const [message, setMessage] = useState("");
    const [messageType, setMessageType] = useState("");
    const [isLoading, setIsLoading] = useState(false);
    const [token, setToken] = useState(null);
    const [isAuthenticated, setIsAuthenticated] = useState(false);

    const router = useRouter();
    const apiUrl = process.env.NEXT_PUBLIC_API_URL;

    // Check for existing session on mount
    useEffect(() => {
        const storedToken = localStorage.getItem("authToken");
        if (storedToken) {
            setToken(storedToken);
            setIsAuthenticated(true);
        }
    }, []);

    const handleSubmit = async (e) => {
        e.preventDefault();
        setIsLoading(true);
        setMessage('');
        setMessageType('');

        try {
            const response = await fetch(`${apiUrl}/signin`, {
                method: "POST",
                headers: {
                    "Content-Type": "application/json"
                },
                body: JSON.stringify({ email, password })
            });

            const data = await response.json();

            if (response.ok) {
                setMessage(`Sign in successful! Welcome, ${data.name}`);
                setMessageType("success");
                setToken(data.token);
                setIsAuthenticated(true);

                // Store token and username in localStorage
                localStorage.setItem("authToken", data.token);
                localStorage.setItem("isAuthenticated", "true");
                localStorage.setItem("username", data.name); // Store the username

                setEmail("");
                setPassword("");

                // Redirect to landpage after 2 seconds
                setTimeout(() => {
                    router.push("/landpage");
                }, 2000);
            } else {
                setMessage(data.error || `Sign in failed (Status: ${response.status})`);
                setMessageType('error');
            }
        } catch (error) {
            console.error('Signin Error:', error);
            setMessage('An error occurred. Please try again.');
            setMessageType('error');
        } finally {
            setIsLoading(false);
        }
    }

    const handleLogout = () => {
        setToken(null);
        setIsAuthenticated(false);
        localStorage.removeItem("authToken");
        localStorage.removeItem("isAuthenticated");
        localStorage.removeItem("username"); // Clear username on logout
        setMessage("Logged out successfully.");
        setMessageType("success");
    }

    return (
        <div className={styles.container}>
            <div className={styles.formContainer}>
                <h1>Sign In</h1>
                {isAuthenticated ? (
                    <div>
                        <p style={{ color: '#5D4037', textAlign: 'center' }}>
                            You are logged in! Redirecting...
                        </p>
                        <button
                            onClick={handleLogout}
                            className={styles.button}
                            style={{ marginTop: '1rem', width: '100%' }}
                        >
                            Logout
                        </button>
                    </div>
                ) : (
                    <>
                        <form onSubmit={handleSubmit} className={styles.form}>
                            <div className={styles.inputGroup}>
                                <label htmlFor="email">Email:</label>
                                <input
                                    type="email"
                                    id="email"
                                    value={email}
                                    onChange={(e) => setEmail(e.target.value)}
                                    required
                                    disabled={isLoading}
                                />
                            </div>
                            <div className={styles.inputGroup}>
                                <label htmlFor="password">Password:</label>
                                <input
                                    type="password"
                                    id="password"
                                    value={password}
                                    onChange={(e) => setPassword(e.target.value)}
                                    required
                                    disabled={isLoading}
                                />
                            </div>

                            {message && (
                                <div className={`${styles.message} ${styles[messageType]}`}>
                                    {message}
                                    {token && messageType === 'success' && (
                                        <pre className={styles.tokenDisplay}>Token: {token}</pre>
                                    )}
                                </div>
                            )}

                            <button type="submit" disabled={isLoading} className={styles.button}>
                                {isLoading ? 'Signing In...' : 'Sign In'}
                            </button>
                        </form>
                        <p>
                            Don't have an account?{' '}
                            <Link href="/signup">
                                Sign Up
                            </Link>
                        </p>
                    </>
                )}
            </div>
        </div>
    );
}