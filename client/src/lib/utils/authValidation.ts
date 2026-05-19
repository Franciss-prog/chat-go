import type { AuthInterface } from '$lib/types/auth';
import { toast } from 'svelte-sonner';

export const validateAuth = (
	{ username, email, password, confirmPassword }: AuthInterface,
	path: string
) => {
	if (path === '/' && (!email || !password)) {
		toast.error('Email and password are required');
		return;
	}
	if (path === '/register' && (!username || !email || !password || !confirmPassword)) {
		toast.error('Username, email and password are required');
		return;
	}
};
