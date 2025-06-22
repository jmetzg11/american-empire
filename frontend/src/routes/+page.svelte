<script>
	import Table from '$lib/components/Table.svelte';
	import { events, fetchEvents } from '$lib/store';
	import { onMount } from 'svelte';

	onMount(async () => {
		await fetchEvents();
	});
</script>

<Table />
