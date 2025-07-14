<script>
	import { onMount } from 'svelte';
	import { getEvent, editEvent, approveEvent, unapproveEvent } from '$lib/api';
	import TopButtons from './components/TopButtons.svelte';
	import Text from './components/Text.svelte';
	import Media from './components/Media.svelte';
	import Sources from './components/Sources.svelte';

	let { data } = $props();
	let event = $state(null);
	let eventEdit = $state({});
	let isEditing = $state(false);
	let alreadyApproved = $derived(event?.Active !== null);

	onMount(async () => {
		try {
			console.log('Fetching event with ID:', data.id);
			event = await getEvent(data.id);
			console.log('Event fetched:', event);
		} catch (error) {
			console.error('Error in onMount:', error);
		}
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
			isEditing = false;
		}
	}

	async function handleApprove() {
		const result = await approveEvent(data.id);
		if (result.ok) {
			await refreshEvent();
		}
	}

	async function handleUnapprove() {
		const result = await unapproveEvent(data.id);
		if (result.ok) {
			await refreshEvent();
		}
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
		<TopButtons
			{handleApprove}
			{handleUnapprove}
			{handleEdit}
			{handleDelete}
			{isEditing}
			{alreadyApproved}
		/>
		<Text
			title={event.Title}
			tags={event.Tags.map((tag) => tag.Name).join(', ')}
			country={event.Country}
			text={event.Description}
			date={event.Date}
			onTitleChange={(v) => updateField('Title', v)}
			onTagsChange={(v) => updateField('Tags', v)}
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
