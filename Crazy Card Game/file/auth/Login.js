import React, { useState } from 'react';
import './Auth.module.css';  // Import the Auth.css file
import { useNavigate } from "react-router-dom";

const Login = ({ onSwitch }) => {
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');
  const [message, setMessage] = useState('');
  const [token, setToken] = useState('');
  const navigate = useNavigate(); 
  const handleSubmit = async (e) => {
    e.preventDefault();
    console.log("here")
    const response = await fetch('http://localhost:8080/login', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({ username, password }),
    });

    const data = await response.json();
    if (response.ok) {
      setMessage('Login successful');
      setToken(data.token);  // Assuming the token is returned in the response
      setUsername('');
      setPassword('');
      navigate("/roomselection");
    } else {
        console.log("abjhvjv")
      setMessage(data.error || 'Error logging in');
    }
  };

  return (
    <div className="auth-container">
      <h2>Login</h2>
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
        <button type="submit">Login</button>
      </form>
      {message && <p className="message">{message}</p>}
      {token && <p className="message">Your token: {token}</p>} {/* Optional, just to show the token */}
      <p>
        Don't have an account?{' '}
        <button onClick={onSwitch}>Sign up here</button>
      </p>
    </div>
  );
};

export default Login;
