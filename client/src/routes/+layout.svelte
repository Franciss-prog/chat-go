<script lang="ts">
	import './layout.css';
	import favicon from '$lib/assets/favicon.svg';
	import { Toaster } from 'svelte-sonner';
	let { children } = $props();
</script>

<Toaster richColors position="bottom-right" duration={1500} />
<svelte:head><link rel="icon" href={favicon} /></svelte:head>
<main class="flex h-screen flex-col bg-dark text-light">
	{@render children()}
</main>
