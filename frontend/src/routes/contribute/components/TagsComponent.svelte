<script>
	import { getTags } from '$lib/api';
	import { onMount } from 'svelte';
	let { tags = $bindable([]) } = $props();

	let allTags = $state([]);
	let newTag = $state('');
	let tagFilter = $state('');
	let selectedTagId = $state('');
	let showDropDown = $state(false);

	let filteredTags = $derived(
		allTags.filter((tag) => tag.Name.toLowerCase().includes(tagFilter.toLowerCase()))
	);

	function selectTag(tag) {
		if (!tags.includes(tag.Name)) {
			tags.push(tag.Name);
		}
		tagFilter = '';
		showDropDown = false;
	}

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
	<div class="mb-2 flex flex-wrap gap-2">
		{#each tags as tag}
			<span class="tag">{tag}</span>
		{/each}
	</div>
	<div class="flex-center-between">
		<div class="flex items-end gap-4">
			<div style="position: relative;">
				<label for="existingTags" class="label">Existing Tags</label>
				<input
					type="text"
					bind:value={tagFilter}
					id="existingTags"
					class="input"
					placeholder="Search tags..."
					onfocus={() => (showDropDown = true)}
					onblur={() => setTimeout(() => (showDropDown = false), 200)}
				/>
				{#if showDropDown && filteredTags.length > 0}
					<div class="dropdown">
						{#each filteredTags.slice(0, 10) as tag}
							<div
								class="dropdown-item"
								role="button"
								onclick={() => selectTag(tag)}
								onkeydown={(e) => e.key === 'Enter' && selectTag(tag)}
								tabindex="0"
							>
								{tag.Name}
							</div>
						{/each}
					</div>
				{/if}
			</div>
			<button class="btn-secondary mb-2" onclick={addTag}>Add Existing Tag</button>
		</div>
		<div class="flex items-end gap-4">
			<div>
				<label for="newTag" class="label">New Tag</label>
				<input type="text" id="newTag" bind:value={newTag} class="input" />
			</div>
			<button class="btn-secondary mb-2" onclick={addTag}>Add New Tag</button>
		</div>
	</div>
</div>

<style>
	.dropdown {
		position: absolute;
		top: 100%;
		left: 0;
		right: 0;
		background: white;
		border: 1px solid #ccc;
		border-radius: 4px;
		max-height: 200px;
		overflow-y: auto;
		z-index: 1000;
	}
	.dropdown-item {
		padding: 8px 12px;
		cursor: pointer;
	}
	.dropdown-item:hover {
		background: #f0f0f0;
	}
</style>
