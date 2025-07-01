<script>
	import { deleteMedia } from '$lib/api';

	let { showDeleteModal = $bindable(), media, refreshEvent } = $props();

	async function handleDelete() {
		const result = await deleteMedia(media.ID);
		if (result.ok) {
			showDeleteModal = false;
			await refreshEvent();
		}
	}
</script>

{#if showDeleteModal}
	<div class="modal-wrapper">
		<div class="modal-content">
			<h2 class="new-info-modal-title">Delete {media.Type === 'photo' ? 'Photo' : 'Youtube'}?</h2>
			<div class="normal-text mb-6">{media.Caption}</div>
			<div class="new-info-button-container">
				<button class="btn-secondary" onclick={() => (showDeleteModal = false)}>Cancel</button>
				<button class="btn-danger" onclick={() => handleDelete()}>Delete</button>
			</div>
		</div>
	</div>
{/if}
