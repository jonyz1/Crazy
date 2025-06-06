import React, { useState } from 'react';
import './Auth.module.css';  // Import the Auth.css file

const Signup = ({ onSwitch }) => {
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');
  const [message, setMessage] = useState('');

  const handleSubmit = async (e) => {
    e.preventDefault();

    const response = await fetch('http://localhost:8080/signup', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({ username, password }),
    });

    const data = await response.json();
    if (response.ok) {
      setMessage('Account created successfully! Please log in.');
      setUsername('');
      setPassword('');
    } else {
      setMessage(data.error || 'Error creating account');
    }
  };

  return (
    <div className="auth-container">
      <h2>Sign Up</h2>
      <form onSubmit={handleSubmit}>
        <input
          type="text"
          placeholder="Username"
          value={username}
          onChange={(e) => setUsername(e.target.value)}
          required
        />
        <input
          type="password"
          placeholder="Password"
          value={password}
          onChange={(e) => setPassword(e.target.value)}
          required
        />
        <button type="submit">Sign Up</button>
      </form>
      {message && <p className="message">{message}</p>}
      <p>
        Already have an account?{' '}
        <button onClick={onSwitch}>Login here</button>
      </p>
    </div>
  );
};

export default Signup;
