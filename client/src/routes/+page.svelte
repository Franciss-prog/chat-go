<script lang="ts">
	// lets try to use runes
	let email = $state('');
	let password = $state('');
	let showPassword = $state(false);

	// login function
	const onLogin = (e: Event) => {
		e.preventDefault();
		// form validation
	};
</script>

<section class="flex min-h-screen items-center justify-center">
	<form onsubmit={onLogin} class="flex flex-col gap-4">
		<h1 class="flex justify-center text-3xl">Login to Chatapp</h1>
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
		<span>Don't have an account? <a href="/register" class="underline">Register</a></span>
	</form>
</section>
