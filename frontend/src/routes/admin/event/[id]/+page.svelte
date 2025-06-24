<script>
	import { onMount } from 'svelte';
	import { getEvent, editEvent } from '$lib/api';
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

	async function handleEdit() {
		eventEdit['id'] = data.id;
		const result = await editEvent(eventEdit);
		if (result.ok) {
			await refreshEvent();
		}
	}

	function handleApprove() {
		$state.snapshot(eventEdit);
	}

	function handleDelete() {
		console.log('deleteEvent');
	}

	function onFKChange(id, changes) {
		isEditing = true;
		eventEdit[id] = changes;
	}
</script>

<div>
	{#if event}
		<TopButtons {handleApprove} {handleEdit} {handleDelete} {isEditing} />
		<Text
			title={event.Title}
			country={event.Country}
			text={event.Description}
			date={event.Date}
			onTitleChange={(v) => updateField('Title', v)}
			onCountryChange={(v) => updateField('Country', v)}
			onDateChange={(v) => updateField('Date', v)}
			onTextChange={(v) => updateField('Description', v)}
		/>
		<Media media={event.Medias} {onFKChange} {refreshEvent} />
		<Sources sources={event.Sources} {onFKChange} {refreshEvent} />
	{:else}
		<div>Loading...</div>
	{/if}
</div>
