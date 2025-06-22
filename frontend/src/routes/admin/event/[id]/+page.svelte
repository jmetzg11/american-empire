<script>
	import { onMount } from 'svelte';
	import { getEvent } from '$lib/api';
	import TopButtons from './components/TopButtons.svelte';
	import Text from './components/Text.svelte';
	import Media from './components/Media.svelte';
	import Sources from './components/Sources.svelte';

	let { data } = $props();
	let event = $state(null);
	let eventEdit = $state({});
	let isEditing = $state(false);

	onMount(async () => {
		console.log('i was called, data id', data.id);
		event = await getEvent(data.id);
		console.log('event', event);
	});

	function updateField(field, value) {
		isEditing = true;
		eventEdit[field] = value;
	}

	function approveEvent() {
		$inspect(event.ID);
	}

	function editEvent() {
		$inspect(eventEdit);
	}

	function deleteEvent() {
		console.log('deleteEvent');
	}
</script>

<div>
	{#if event}
		<TopButtons {approveEvent} {editEvent} {deleteEvent} {isEditing} />
		<Text
			title={event.Title}
			text={event.Description}
			date={event.Date}
			onTitleChange={(v) => updateField('Title', v)}
			onDateChange={(v) => updateField('Date', v)}
			onTextChange={(v) => updateField('Description', v)}
		/>
		<Media media={event.Medias} />
		<Sources sources={event.Sources} />
	{:else}
		<div>Loading...</div>
	{/if}
</div>
