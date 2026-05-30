<script lang="ts">
	import { validateAuth, handleLogin, handleRegister } from '$lib';
	import { goto } from '$app/navigation';
	import { toast } from 'svelte-sonner';

	let {
		username = $bindable(''),
		email = $bindable(''),
		password = $bindable(''),
		confirmPassword = $bindable(''),
		pathname,
		isRegisterPath,
		loading = $bindable(false)
	} = $props();

	// input data to avoid redundancy
	const inputs = $derived([
		...(isRegisterPath
			? [
					{
						type: 'text',
						placeholder: 'Username',
						value: () => username,
						setValue: (v: string) => (username = v)
					}
				]
			: []),

		{
			type: 'email',
			placeholder: 'Email',
			value: () => email,
			setValue: (v: string) => (email = v)
		},

		{
			type: 'password',
			placeholder: 'Password',
			value: () => password,
			setValue: (v: string) => (password = v)
		},

		...(isRegisterPath
			? [
					{
						type: 'password',
						placeholder: 'Confirm Password',
						value: () => confirmPassword,
						setValue: (v: string) => (confirmPassword = v)
					}
				]
			: [])
	]);

	// login
	const onLogin = async (e: Event) => {
		loading = true;

		e.preventDefault();

		validateAuth({ email, password }, pathname);

		const success = await handleLogin({ email, password });

		if (!success) {
			setTimeout(() => {
				email = '';
				password = '';
				loading = false;
			}, 1201);
			return;
		}

		email = '';
		password = '';
		loading = false;

		setTimeout(() => goto('/chat'), 1201);
	};

	// register
	const onRegister = async (e: Event) => {
		e.preventDefault();

		if (password !== confirmPassword) {
			toast.warning('Passwords do not match');
			password = '';
			confirmPassword = '';
			return;
		}

		const success = await handleRegister({ username, email, password });

		if (!success) {
			username = '';
			email = '';
			password = '';
			confirmPassword = '';
			return;
		}
		setTimeout(() => goto('/chat'), 1500);
	};
</script>

<form onsubmit={isRegisterPath ? onRegister : onLogin} class="w-full space-y-10">
	<div class="flex flex-col gap-4">
		{#each inputs as input}
			<input
				type={input.type}
				placeholder={input.placeholder}
				value={input.value()}
				oninput={(e) => input.setValue((e.currentTarget as HTMLInputElement).value)}
				class=" w-8xl h-20 border-0 border-b border-gray-400 bg-transparent text-3xl focus:border-black focus:outline-none"
			/>
		{/each}
	</div>
	<button
		class="h-20 w-full border border-light bg-transparent text-3xl transition-colors hover:bg-light hover:text-dark {loading
			? 'cursor-not-allowed'
			: 'cursor-pointer'}"
		disabled={loading}
	>
		{#if loading}
			<div class="flex flex-row items-center justify-center gap-2">
				<div class="h-4 w-4 animate-bounce rounded-full bg-light"></div>
				<div class="h-4 w-4 animate-bounce rounded-full bg-light [animation-delay:-.3s]"></div>
				<div class="h-4 w-4 animate-bounce rounded-full bg-light [animation-delay:-.5s]"></div>
			</div>
		{:else}
			{isRegisterPath ? 'Register' : 'Login'}
		{/if}
	</button>
</form>
