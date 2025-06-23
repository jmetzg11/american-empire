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
		event = await getEvent(data.id);
	});

	async function refreshEvent() {
		event = await getEvent(data.id);
	}

	function updateField(field, value) {
		isEditing = true;
		eventEdit[field] = value;
	}

	function onFKChange(id, changes) {
		isEditing = true;
		eventEdit[id] = changes;
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
		<Media media={event.Medias} {onFKChange} {refreshEvent} />
		<Sources sources={event.Sources} {onFKChange} {refreshEvent} />
	{:else}
		<div>Loading...</div>
	{/if}
</div>
