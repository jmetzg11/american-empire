<script>
	let { approveEvent, editEvent, deleteEvent, isEditing } = $props();
</script>

<div class="container-wrap">
	<div class="card-style flex justify-between">
		<button class="btn" onclick={approveEvent} disabled={isEditing}>Approve</button>
		<button class="btn-secondary" onclick={editEvent} disabled={!isEditing}>Save Edits</button>
		<button class="btn-danger" onclick={deleteEvent}>Delete</button>
	</div>
</div>
