<script>
	let { handleApprove, handleUnapprove, handleEdit, handleDelete, isEditing, alreadyApproved } =
		$props();
</script>

<div class="container-wrap">
	<div class="card-style flex justify-between">
		{#if alreadyApproved}
			<button class="btn" onclick={handleUnapprove}>Unapprove</button>
		{:else}
			<button class="btn" onclick={handleApprove} disabled={isEditing}>Approve</button>
		{/if}
		<button class="btn-secondary" onclick={handleEdit} disabled={!isEditing}>Save Edits</button>
		<button class="btn-danger" onclick={handleDelete}>Delete</button>
	</div>
</div>
