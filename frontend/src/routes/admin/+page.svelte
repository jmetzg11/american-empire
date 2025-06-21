<script>
	import { onMount } from 'svelte';
	import { authMe } from '$lib/api';
	import { fetchAdminEvents } from '$lib/store';
	import Table from '$lib/Table.svelte';
	import Login from './components/Login.svelte';
	let isAuthenticated = $state(false);

	onMount(async () => {
		const result = await authMe();
		isAuthenticated = result;
		if (isAuthenticated) {
			await fetchAdminEvents();
		}
	});
</script>

<div>
	{#if isAuthenticated}
		<Table />
	{:else}
		<Login bind:isAuthenticated />
	{/if}
</div>
