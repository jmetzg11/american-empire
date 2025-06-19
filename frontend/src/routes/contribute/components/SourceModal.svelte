<script>
	let { showSourceModal = $bindable(), source = $bindable() } = $props();
	let name = $state('');
	let url = $state('');
	let canSubmit = $derived(name && url);

	function addSource() {
		if (canSubmit) {
			source.push({
				name,
				url
			});
			name = '';
			url = '';
		}
	}

	function removeSource(index) {
		source.splice(index, 1);
	}

	function closeModal() {
		showSourceModal = false;
	}
</script>

{#if showSourceModal}
	<div class="modal-wrapper">
		<div class="modal-content">
			<h2 class="new-info-modal-title">Existing Sources</h2>
			{#each source as s, i}
				<div class="existing-item-row">
					<div class="existing-item-text">
						<span class="existing-item-title">{s.name}</span>
						<span class="existing-item-type">{s.url}</span>
					</div>
					<button class="btn-danger" onclick={() => removeSource(i)}>Remove</button>
				</div>
			{/each}
			<h2 class="new-info-modal-title mt-6">Add Source</h2>
			<div class="new-info-input-container">
				<div class="new-info-input-left">
					<label for="source-title" class="label">Title</label>
					<input id="-source-title" type="text" class="input" bind:value={name} />
				</div>
				<div class="new-info-input-right">
					<label for="source-url" class="label">Past URL</label>
					<input id="source-url" type="text" class="input" bind:value={url} />
				</div>
			</div>
			<div class="new-info-button-container">
				<button class="btn-secondary mt-6" onclick={addSource} disabled={!canSubmit}>Add</button>
				<button class="btn-secondary mt-6" onclick={closeModal}>Close</button>
			</div>
		</div>
	</div>
{/if}
