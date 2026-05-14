<script lang="ts">
	// page import for route validation rendering
	import { page } from '$app/state';
	import { toast } from 'svelte-sonner';

	// STATES use for forms
	let username = $state('');
	let email = $state('');
	let password = $state('');
	let confirmPassword = $state('');
	// state for password visibility
	let showPassword = $state(false);

	// derive for register path
	const pathname = $derived(page.url.pathname);
	const isRegisterPath = $derived(pathname === '/register');

	// onLogin function
	const onLogin = async (e: Event) => {
		e.preventDefault();
		console.log('Login');
	};

	// onRegister function
	const onRegister = async (e: Event) => {
		e.preventDefault();
		console.log('Register');
	};
</script>

<section class="flex min-h-screen items-center justify-center">
	<form onsubmit={isRegisterPath ? onRegister : onLogin} class="flex flex-col gap-4">
		<h1 class="flex justify-center text-3xl">
			{isRegisterPath ? 'Register' : 'Login'} to my chatapp
		</h1>
		<input type="email" bind:value={email} placeholder="Email" />
		<div class="flex items-center gap-2">
			<input
				type={showPassword ? 'text' : 'password'}
				bind:value={password}
				placeholder="Password"
			/>
			<button onclick={() => (showPassword = !showPassword)} class="bg-black px-4 py-2 text-white"
				>{showPassword ? 'Hide' : 'Show'}</button
			>
		</div>
		<button class="bg-black px-4 py-2 text-white">{isRegisterPath ? 'Register' : 'Login'}</button>
		<span
			>{isRegisterPath ? 'Already' : "Don't"} have an account?
			<a href={isRegisterPath ? '/' : '/register'} class="underline"
				>{isRegisterPath ? 'Login' : 'Register'}</a
			></span
		>
	</form>
</section>
