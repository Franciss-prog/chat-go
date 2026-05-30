<script lang="ts">
	import { page } from '$app/state';
	import { AuthForm } from '$lib';

	let username = $state('');
	let email = $state('');
	let password = $state('');
	let confirmPassword = $state('');
	let loading = $state(false);

	const pathname = $derived(page.url.pathname);
	const isRegisterPath = $derived(pathname === '/register');
</script>

<section class="flex h-screen flex-col bg-dark px-4 py-2 text-light">
	<nav class="w-8xl mb-10 flex items-start justify-between text-7xl">
		<span>{isRegisterPath ? 'Register' : 'Login'}</span>
		<span>
			Go to <a href={isRegisterPath ? '/' : '/register'} class="underline">
				{isRegisterPath ? 'Login' : 'Register'}
			</a>
		</span>
	</nav>
	<AuthForm {username} {email} {password} {confirmPassword} {pathname} {isRegisterPath} {loading} />
</section>
