<script>
	import { getTags } from '$lib/api';
	import { onMount } from 'svelte';
	let { tags = $bindable([]) } = $props();

	let allTags = $state([]);
	let newTag = $state('');
	let selectedTagId = $state('');

	function addTag() {
		if (newTag.trim() && !allTags.some((tag) => tag.Name === newTag.trim())) {
			tags.push(newTag.trim());
			newTag = '';
		} else if (selectedTagId) {
			const selectedTag = allTags.find((tag) => tag.ID === parseInt(selectedTagId));
			if (selectedTag && !tags.includes(selectedTag.Name)) {
				tags.push(selectedTag.Name);
				selectedTagId = '';
			}
		}
	}

	onMount(async () => {
		const response = await getTags();
		allTags = response;
	});
</script>

<div>
	<div>
		{#each tags as tag}
			<span class="tag">{tag}</span>
		{/each}
	</div>
	<div class="flex-center-between">
		<div>
			<select bind:value={selectedTagId}>
				{#each allTags as tag}
					<option value={tag.ID}>{tag.Name}</option>
				{/each}
			</select>
		</div>
		<div>
			<label for="newTag">new tag</label>
			<input type="text" id="newTag" bind:value={newTag} />
		</div>
		<button class="btn-secondary" onclick={addTag}>add</button>
	</div>
</div>
