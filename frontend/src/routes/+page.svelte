<script>
	import Table from '$lib/Table.svelte';
	import { events, fetchEvents } from '$lib/store';
	import { onMount } from 'svelte';

	onMount(async () => {
		await fetchEvents();
	});
</script>

<Table />
