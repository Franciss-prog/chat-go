// place files you want to import through the `$lib` alias in this folder.

// server url (chnage this on prod if i want to make it public)
export const API_URL = 'http://localhost:8080';

// interfaces  exports
export { type AuthInterface } from '$lib/types/authTypes';

// components exports
export { default as Auth } from './components/auth/Auth.svelte';

// actions exports
export { handleLogin, handleRegister } from './actions/authActions';

// utils exports
export { validateAuth } from './utils/authValidation';
