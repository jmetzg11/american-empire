<script>
	import Header from './components/Header.svelte';
	import Table from './components/Table.svelte';
	import { events, fetchEvents } from '$lib/store';
	import { onMount } from 'svelte';

	onMount(async () => {
		await fetchEvents();
	});
</script>

<Header />
<Table {events} />
