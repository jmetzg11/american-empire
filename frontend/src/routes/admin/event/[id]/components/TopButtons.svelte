<script>
	let { handleApprove, handleEdit, handleDelete, isEditing, alreadyApproved } = $props();
</script>

<div class="container-wrap">
	<div class="card-style flex justify-between">
		<button class="btn" onclick={handleApprove} disabled={isEditing || alreadyApproved}
			>Approve</button
		>
		<button class="btn-secondary" onclick={handleEdit} disabled={!isEditing}>Save Edits</button>
		<button class="btn-danger" onclick={handleDelete}>Delete</button>
	</div>
</div>
