<script>
	import MediaModal from './MediaModal.svelte';

	let country = $state('');
	let title = $state('');
	let date = $state('');
	let description = $state('');
	let media = $state([]);
	let source = $state([]);
	let ableToSubmit = $derived(country && title && date && description && source.length > 0);
	let showMediaModal = $state(false);
	let showSourceModal = $state(false);

	function autoGrow(event) {
		event.target.style.height = 'auto';
		event.target.style.height = event.target.scrollHeight + 'px';
	}
</script>

<div class="container-wrap">
	<div class="card-style">
		<div class="space-y-6">
			<div>
				<label for="title" class="label">Title *</label>
				<input id="title" type="text" bind:value={title} class="input" placeholder="Enter title" />
			</div>
			<div class="flex justify-between">
				<div>
					<label for="country" class="label">Country *</label>
					<input
						id="country"
						type="text"
						bind:value={country}
						class="input"
						placeholder="Enter country"
					/>
				</div>
				<div>
					<label for="date" class="label">Date *</label>
					<input id="date" type="date" bind:value={date} class="input" />
				</div>
			</div>
			<div>
				<label for="description" class="label">Description *</label>
				<textarea
					id="description"
					bind:value={description}
					class="input"
					oninput={autoGrow}
					placeholder="Enter description..."
				></textarea>
			</div>
			<div class="flex justify-between">
				<button type="button" onclick={() => (showMediaModal = true)} class="btn-secondary"
					>Add Media</button
				>
				<button type="button" onclick={() => (showSourceModal = true)} class="btn-secondary">
					Add Source *
				</button>
			</div>
		</div>
	</div>
</div>

<MediaModal bind:showMediaModal {media} />
