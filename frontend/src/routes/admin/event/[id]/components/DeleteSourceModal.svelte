<script>
	import { deleteSource } from '$lib/api';

	let { showDeleteModal = $bindable(), selectedSource, refreshEvent } = $props();

	async function handleDelete() {
		const result = await deleteSource(selectedSource.ID);
		if (result.ok) {
			showDeleteModal = false;
			await refreshEvent();
		}
	}
</script>

{#if showDeleteModal}
	<div class="modal-wrapper">
		<div class="modal-content">
			<h2 class="new-info-modal-title">Delete Sources?</h2>
			<div class="mb-6">
				<div class="normal-text mb-2">{selectedSource.Name}</div>
				<a href={selectedSource.URL} target="_blank" class="text-blue-500 hover:underline"
					>{selectedSource.URL}</a
				>
			</div>

			<div class="new-info-button-container">
				<button class="btn-secondary" onclick={() => (showDeleteModal = false)}>Cancel</button>
				<button class="btn-danger" onclick={handleDelete}>Delete</button>
			</div>
		</div>
	</div>
{/if}
