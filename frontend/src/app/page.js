'use client';

import Link from 'next/link';

export default function HomePage() {
  return (
    <div
      style={{
        minHeight: '100vh',
        display: 'flex',
        flexDirection: 'column',
        justifyContent: 'center',
        alignItems: 'center',
        backgroundColor: '#F5F5DC', // Base off-white background
        fontFamily: "'Poppins', sans-serif",
        padding: '20px',
      }}
    >
      <h1
        style={{
          fontSize: '3rem',
          fontWeight: '700',
          color: '#5D4037', // Muted brown text for readability
          marginBottom: '1rem',
          textShadow: '2px 2px 4px rgba(0, 0, 0, 0.1)',
          animation: 'fadeIn 1s ease-in-out',
        }}
      >
        Welcome!
      </h1>
      <p
        style={{
          fontSize: '1.2rem',
          color: '#5D4037', // Muted brown text
          marginBottom: '2rem',
          animation: 'fadeIn 1.5s ease-in-out',
        }}
      >
        Please choose an option:
      </p>

      <div
        style={{
          display: 'flex',
          justifyContent: 'center',
          gap: '1.5rem',
          animation: 'fadeIn 2s ease-in-out',
        }}
      >
        <Link href="/signin">
          <button
            style={{
              padding: '12px 30px',
              fontSize: '1.1rem',
              fontWeight: '600',
              color: '#FAFAF0', // Light cream text on button
              backgroundColor: '#E6D7B2', // Deeper beige for button
              border: 'none',
              borderRadius: '8px',
              cursor: 'pointer',
              boxShadow: '0 4px 6px rgba(0, 0, 0, 0.1)',
              transition: 'transform 0.2s, background-color 0.3s',
            }}
            onMouseEnter={(e) => {
              e.target.style.backgroundColor = '#D6C7A2'; // Darker beige on hover
              e.target.style.transform = 'translateY(-3px)';
            }}
            onMouseLeave={(e) => {
              e.target.style.backgroundColor = '#E6D7B2';
              e.target.style.transform = 'translateY(0)';
            }}
          >
            Sign In
          </button>
        </Link>
        
        <Link href="/signup">
          <button
            style={{
              padding: '12px 30px',
              fontSize: '1.1rem',
              fontWeight: '600',
              color: '#FAFAF0', // Light cream text on button
              backgroundColor: '#E6D7B2', // Deeper beige for button
              border: 'none',
              borderRadius: '8px',
              cursor: 'pointer',
              boxShadow: '0 4px 6px rgba(0, 0, 0, 0.1)',
              transition: 'transform 0.2s, background-color 0.3s',
            }}
            onMouseEnter={(e) => {
              e.target.style.backgroundColor = '#D6C7A2'; // Darker beige on hover
              e.target.style.transform = 'translateY(-3px)';
            }}
            onMouseLeave={(e) => {
              e.target.style.backgroundColor = '#E6D7B2';
              e.target.style.transform = 'translateY(0)';
            }}
          >
            Sign Up
          </button>
        </Link>
      </div>

      {/* Keyframes for animations */}
      <style jsx global>{`
        @keyframes fadeIn {
          0% {
            opacity: 0;
            transform: translateY(20px);
          }
          100% {
            opacity: 1;
            transform: translateY(0);
          }
        }
      `}</style>
    </div>
  );
}