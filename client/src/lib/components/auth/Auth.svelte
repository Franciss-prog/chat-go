<script lang="ts">
	import { goto } from '$app/navigation';
	import { page } from '$app/state';
	import { validateAuth, handleRegister, handleLogin } from '$lib';
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

	const onLogin = async (e: Event) => {
		e.preventDefault();

		// validate the inputs
		validateAuth({ email, password }, pathname);

		// login
		const success = await handleLogin({ email, password });

		if (!success) {
			setTimeout(() => {
				email = '';
				password = '';
			}, 1501);
			return;
		}

		// clear the form
		email = '';
		password = '';
		setTimeout(() => goto('/chat'), 1501);
	};
	// onRegister function
	const onRegister = async (e: Event) => {
		e.preventDefault();
		// validate the inputs for register

		// check if password and confirm password match
		if (password !== confirmPassword) {
			toast.warning('Passwords do not match');
			return;
		}

		// perform registration
		try {
			const data = await handleRegister({ username, email, password });
			if (data.success === true && data.message) {
			}
		} catch (error: any) {
			toast.error(error);
		}
	};
</script>

<section class="flex min-h-screen items-center justify-center">
	<form onsubmit={isRegisterPath ? onRegister : onLogin} class="flex flex-col gap-4">
		<h1 class="flex justify-center text-3xl">
			{isRegisterPath ? 'Register' : 'Login'} to my chatapp
		</h1>
		{#if isRegisterPath}
			<input type="text" bind:value={username} placeholder="Username" />
		{/if}

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
		<!-- CONDITIONAL RENDERING FOR SHOWPASSWORD -->
		{#if isRegisterPath}
			<input type="password" bind:value={confirmPassword} placeholder="Confirm Password" />
		{/if}
		<button class="bg-black px-4 py-2 text-white">{isRegisterPath ? 'Register' : 'Login'}</button>
		<span
			>{isRegisterPath ? 'Already' : "Don't"} have an account?
			<a href={isRegisterPath ? '/' : '/register'} class="underline"
				>{isRegisterPath ? 'Login' : 'Register'}</a
			></span
		>
	</form>
</section>
