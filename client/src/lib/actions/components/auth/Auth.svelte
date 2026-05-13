<script lang="ts">
	import { page } from '$app/state';

	// pathname derived for reactivity
	let pathname = $derived(page.url.pathname);
</script>
