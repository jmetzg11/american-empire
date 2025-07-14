<script>
	import { onMount } from 'svelte';
	import { getEvent } from '$lib/api';
	import Text from './components/Text.svelte';
	import MediaCarousel from './components/MediaCarousel.svelte';
	import Sources from './components/Sources.svelte';

	let { data } = $props();
	let event = $state(null);

	onMount(async () => {
		try {
			console.log('Fetching event with ID:', data.id);
			event = await getEvent(data.id);
			console.log('Event fetched:', event);
		} catch (error) {
			console.error('Error in onMount:', error);
		}
	});
</script>

<div>
	{#if event}
		<Text title={event.Title} tags={event.Tags} text={event.Description} date={event.Date} />
		{#if event.Medias.length > 0}
			<MediaCarousel media={event.Medias} />
		{/if}
		<Sources sources={event.Sources} />
	{:else}
		<div>Loading...</div>
	{/if}
</div>
