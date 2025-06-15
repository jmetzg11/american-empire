<script>
	import { onMount } from 'svelte';
	import { getEvent } from '$lib/api';
	import Text from './components/Text.svelte';
	import MediaCarosel from './components/MediaCarosel.svelte';
	import Sources from './components/Sources.svelte';

	let { data } = $props();
	let event = $state(null);

	$inspect(event);

	onMount(async () => {
		event = await getEvent(data.id);
	});

	$inspect(event);
</script>

<div>
	{#if event}
		<Text title={event.Title} text={event.Description} date={event.Date} />
		<MediaCarosel media={event.Medias} />
		<Sources sources={event.Sources} />
	{:else}
		<div>Loading...</div>
	{/if}
</div>
