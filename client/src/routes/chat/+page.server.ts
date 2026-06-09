import type { PageServerLoad } from './$types';
import { redirect } from '@sveltejs/kit';
import axios from 'axios';
import { API_URL } from '$lib';

export const load: PageServerLoad = ({ cookies }) => {
	// auth guard here to authorize if the user is login by checking the cookies
	if (!cookies.get('access_token')) {
		throw redirect(302, '/');
	}
	// if they have a token, validate the token to server and get the user info
};
