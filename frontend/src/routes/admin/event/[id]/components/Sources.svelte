<script>
	import { addSource } from '$lib/api';
	import DeleteSourceModal from './DeleteSourceModal.svelte';
	import { page } from '$app/stores';
	let { sources, onFKChange, refreshEvent } = $props();

	let showDeleteModal = $state(false);
	let selectedSource = $state(null);
	let sourceName = $state('');
	let sourceURL = $state('');
	let eventID = $page.params.id;

	function handleDelete(source) {
		selectedSource = source;
		showDeleteModal = true;
	}

	async function handleSave() {
		const payload = {
			Name: sourceName,
			URL: sourceURL,
			EventID: eventID
		};
		const result = await addSource(payload);
		if (result.ok) {
			await refreshEvent();
			sourceName = '';
			sourceURL = '';
		}
	}
</script>

<div class="container-wrap">
	<div class="card-style">
		<div class="small-info-text mb-4">can edit names, but must delete links to change them</div>
		<h2 class="section-title">Sources</h2>
		{#each sources as source}
			<div class="source-row-admin">
				<input
					type="text"
					class="input"
					value={source.Name}
					oninput={(e) => onFKChange('source-' + source.ID, { Name: e.target.value })}
				/>
				<div class="flex-center-between mt-4">
					<a href={source.URL} target="_blank" class="a-link">{source.URL}</a>
					<button class="btn-danger" onclick={() => handleDelete(source)}> Remove </button>
				</div>
			</div>
		{/each}
		<h2 class="section-title">New Source</h2>
		<div class="flex-col-gap-2">
			<input type="text" class="input" placeholder="Name" bind:value={sourceName} />
			<input type="text" class="input" placeholder="URL" bind:value={sourceURL} />
			<button class="btn" onclick={handleSave}>Save</button>
		</div>
	</div>
</div>

<DeleteSourceModal bind:showDeleteModal {selectedSource} {refreshEvent} />
