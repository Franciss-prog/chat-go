import axios from 'axios';
import { API_URL } from '$lib';
import type { AuthInterface } from '$lib';
import { toast } from 'svelte-sonner';

export const login = async ({ username, password }: AuthInterface) => {
	// form validation
	if (!username || !password) {
		toast.error('Username and password are required');
		return;
	}
	// use try catch for post req
	const { data } = await axios.post(`${API_URL}/login`, { username, password });
	return data;
};
