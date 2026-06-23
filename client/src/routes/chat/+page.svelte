<script lang="ts">
	import { SendHorizontal } from 'lucide-svelte';
</script>

<form class="flex h-full flex-col gap-3 px-4 py-5">
	<!-- render all messages (empty box for now) -->
	<div class="scrollable h-full rounded-md border"></div>
	<div class="flex items-center gap-4">
		<input
			type="text"
			class="w-[90%] border-b bg-transparent text-3xl focus:border-black focus:outline-none"
			placeholder="Enter your message here"
		/>
		<!-- submit button -->
		<button
			class="flex h-13 w-[10%] items-center justify-center rounded-md border transition-colors hover:bg-light hover:text-dark"
			><SendHorizontal /></button
		>
	</div>
</form>
