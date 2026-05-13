import type { AuthInterface } from '$lib';
import { API_URL } from '$lib';
import axios from 'axios';

export const onRegister = async ({ username, email, password }: AuthInterface) => {
	// form validation
	if (!email || !password) {
		console.error('Email and password are required');
		return;
	}

	// perform post request using axios
	const { data } = await axios.post(`${API_URL}/register`, { username, email, password });

	return data;
};
