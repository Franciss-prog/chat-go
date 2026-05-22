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
		isRegisterPath
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
		e.preventDefault();

		validateAuth({ email, password }, pathname);

		const success = await handleLogin({ email, password });

		if (!success) {
			email = '';
			password = '';
			return;
		}

		email = '';
		password = '';

		setTimeout(() => goto('/chat'), 1500);
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
		class="h-20 w-full border border-light bg-transparent text-3xl transition-colors hover:bg-light hover:text-dark"
		>{isRegisterPath ? 'Register' : 'Login'}</button
	>
</form>
