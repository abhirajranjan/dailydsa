import React from 'react'
import ReactDOM from 'react-dom/client'
import App from './App.tsx'
import { GoogleOAuthProvider } from '@react-oauth/google';

ReactDOM.createRoot(document.getElementById('root') as HTMLElement).render(
  <GoogleOAuthProvider clientId={`${import.meta.env.VITE_GCLIENTID}`}>
  <React.StrictMode>
    <App />
  </React.StrictMode>
</GoogleOAuthProvider>)
