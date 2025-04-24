"use client"
import { useState } from "react";
import Link from "next/link";
import { useRouter } from "next/navigation"; // Import useRouter for redirection
import styles from "../styles/Form.module.css"

export default function SignupPage() {
    const [name, setName] = useState("");
    const [email, setEmail] = useState("");
    const [password, setPassword] = useState("");
    const [message, setMessage] = useState("");
    const [messageType, setMessageType] = useState("");
    const [isLoading, setIsLoading] = useState(false);

    const router = useRouter(); // Initialize router for redirection
    const apiUrl = process.env.NEXT_PUBLIC_API_URL;

    const handleSubmit = async (e) => {
        e.preventDefault();
        setIsLoading(true);
        setMessage('');
        setMessageType('');

        try {
            const response = await fetch(`${apiUrl}/signup`, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({ name, email, password })
            });

            const data = await response.json();

            if (response.ok) {
                setMessage(data.message || "Signup successful! Redirecting to sign-in...");
                setMessageType("success");
                setName("");
                setEmail("");
                setPassword("");

                // Redirect to sign-in page after 1 seconds
                setTimeout(() => {
                    router.push("/signin");
                }, 1000);
            } else {
                setMessage(data.error || `Signup failed (Status: ${response.status})`);
                setMessageType("error");
            }
        } catch (error) {
            console.error(`Signup Error:`, error);
            setMessage("An error occurred. Please try again.");
            setMessageType("error");
        } finally {
            setIsLoading(false);
        }
    }

    return (
        <div className={styles.container}>
            <div className={styles.formContainer}>
                <h1>Sign Up</h1>
                <form onSubmit={handleSubmit} className={styles.form}>
                    <div className={styles.inputGroup}>
                        <label htmlFor="name">Name:</label>
                        <input
                            type="text"
                            id="name"
                            value={name}
                            onChange={(e) => setName(e.target.value)}
                            required
                            minLength={2}
                            disabled={isLoading}
                        />
                    </div>
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
                            minLength={8}
                            disabled={isLoading}
                        />
                    </div>

                    {message && (
                        <div className={`${styles.message} ${styles[messageType]}`}>
                            {message}
                        </div>
                    )}

                    <button type="submit" disabled={isLoading} className={styles.button}>
                        {isLoading ? 'Signing Up...' : 'Sign Up'}
                    </button>
                </form>
                <p>
                    Already have an account?{' '}
                    <Link href="/signin">
                        Sign In
                    </Link>
                </p>
            </div>
        </div>
    );
}