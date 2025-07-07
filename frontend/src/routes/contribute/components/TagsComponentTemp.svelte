<script>
	import { getTags } from '$lib/api';
	import { onMount } from 'svelte';
	let { tags = $bindable([]) } = $props();

	let allTags = $state([]);
	let newTag = $state('');
	let tagFilter = $state('');
	let showDropdown = $state(false);

	let filteredTags = $derived(
		allTags.filter((tag) => tag.Name.toLowerCase().includes(tagFilter.toLowerCase()))
	);

	function addTag() {
		if (newTag.trim() && !tags.includes(newTag.trim())) {
			tags.push(newTag.trim());
			newTag = '';
		}
	}

	function selectTag(tag) {
		if (!tags.includes(tag.Name)) {
			tags.push(tag.Name);
		}
		tagFilter = '';
		showDropdown = false;
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
		<div style="position: relative;">
			<input
				type="text"
				bind:value={tagFilter}
				placeholder="Search tags..."
				onfocus={() => (showDropdown = true)}
				onblur={() => setTimeout(() => (showDropdown = false), 200)}
			/>
			{#if showDropdown && filteredTags.length > 0}
				<div class="dropdown">
					{#each filteredTags.slice(0, 10) as tag}
						<div class="dropdown-item" onclick={() => selectTag(tag)}>
							{tag.Name}
						</div>
					{/each}
				</div>
			{/if}
		</div>
		<div>
			<label for="newTag">new tag</label>
			<input type="text" id="newTag" bind:value={newTag} />
		</div>
		<button class="btn-secondary" onclick={addTag}>add</button>
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
