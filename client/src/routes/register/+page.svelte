<script lang="ts">
	import { API_URL } from '$lib';
	// lets try to use runes
	let username = $state('');
	let email = $state('');
	let password = $state('');
	let showPassword = $state(false);

	// login function
	const onRegister = async (e: Event) => {
		e.preventDefault();
		// form validation
		if (!email || !password) {
			console.error('Email and password are required');
			return;
		}

		// perform the login operation
		try {
			const response = await fetch(`${API_URL}/register`, {
				method: 'POST',
				headers: {
					'Content-Type': 'application/json'
				},
				body: JSON.stringify({ username, email, password })
			});
			const data = await response.json();
			console.log(data);
		} catch (error) {
			console.error(error);
		}
	};
</script>

<section class="flex min-h-screen items-center justify-center">
	<form onsubmit={onRegister} class="flex flex-col gap-4">
		<h1 class="flex justify-center text-3xl">Register to Chatapp</h1>
		<input type="text" bind:value={username} placeholder="Username" />
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
		<button class="bg-black px-4 py-2 text-white">Login</button>
		<span>Already have an account? <a href="/" class="underline">Login</a></span>
	</form>
</section>
