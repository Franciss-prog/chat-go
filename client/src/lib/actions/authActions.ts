import type { LoginInterface, RegisterInterface } from '$lib';
import axios from 'axios';
import { toast } from 'svelte-sonner';
import { API_URL, userStore } from '$lib';

export const handleLogin = async ({ email, password }: LoginInterface): Promise<boolean> => {
	try {
		const response = await axios.post(
			`${API_URL}/login`,
			{ email, password },
			{
				withCredentials: true
			}
		);
		if (response.status === 200 || response.status === 201) {
			toast.success(response.data.message, { duration: 1500 });
			return true;
		}
		return false;
	} catch (error) {
		if (axios.isAxiosError(error)) {
			toast.error(error.response?.data.message ?? 'Login failed', { duration: 1500 });
		}
		return false;
	}
};

export const handleRegister = async ({
	username,
	email,
	password
}: RegisterInterface): Promise<boolean> => {
	try {
		const response = await axios.post(
			`${API_URL}/register`,
			{
				username,
				email,
				password
			},
			{
				withCredentials: true
			}
		);
		if (response.status === 200 || response.status === 201) {
			toast.success(response.data.message, { duration: 1500 });
			userStore.set({ id: response.data.id, username });
			return true;
		}
		return false;
	} catch (error) {
		if (axios.isAxiosError(error)) {
			toast.error(error.response?.data.message ?? 'Registration failed', { duration: 1500 });
		}
		return false;
	}
};
