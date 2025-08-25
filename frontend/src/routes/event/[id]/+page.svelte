<script>
	import { onMount } from 'svelte';
	import { getEvent } from '$lib/api';
	import Text from './components/Text.svelte';
	import MediaCarousel from './components/MediaCarousel.svelte';
	import Sources from './components/Sources.svelte';
	import Books from './components/Books.svelte';
	let { data } = $props();
	let event = $state(null);

	onMount(async () => {
		try {
			event = await getEvent(data.id);
			console.log(event);
		} catch (error) {
			console.error('Error in onMount:', error);
		}
	});
	function showMoreBookDetails() {
		console.log('show more book details');
	}
</script>

<div>
	{#if event}
		<Text title={event.Title} tags={event.Tags} text={event.Description} date={event.Date} />
		{#if event.Medias.length > 0}
			<MediaCarousel media={event.Medias} />
		{/if}

		<Sources sources={event.Sources} />
		{#if event.Books.length > 0}
			<Books books={event.Books} />
		{/if}
	{:else}
		<div>Loading...</div>
	{/if}
</div>
