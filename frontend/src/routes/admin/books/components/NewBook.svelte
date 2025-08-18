<script>
	import { onMount } from 'svelte';
	import { getTags, submitNewBook } from '$lib/api';
	let tags = $state([]);
	let newBook = $state({
		title: '',
		author: '',
		link: '',
		events: '',
		selectedTags: []
	});

	onMount(async () => {
		try {
			tags = await getTags();
			console.log(tags);
		} catch (error) {
			console.error('Error in onMount:', error);
		}
	});

	function toggleTag(tag) {
		if (newBook.selectedTags.includes(tag.ID)) {
			newBook.selectedTags = newBook.selectedTags.filter((id) => id !== tag.ID);
		} else {
			newBook.selectedTags = [...newBook.selectedTags, tag.ID];
		}
	}

	function handleSubmit() {
		submitNewBook(newBook);
		newBook = {
			title: '',
			author: '',
			link: '',
			events: '',
			selectedTags: []
		};
	}
</script>

<div class="container-wrap">
	<div class="card-style">
		<h2 class="section-title">New Book</h2>
		<div class="flex-col-gap-2">
			<input type="text" class="input" placeholder="Title" bind:value={newBook.title} />
			<input type="text" class="input" placeholder="Author" bind:value={newBook.author} />
			<input type="text" class="input" placeholder="URL" bind:value={newBook.link} />
			<input type="text" class="input" placeholder="1, 2, 3" bind:value={newBook.events} />
		</div>
		{#each tags as tag}
			<span
				class={`tag${newBook.selectedTags.includes(tag.ID) ? '-selected' : ''} mr-2`}
				role="button"
				tabindex="0"
				onclick={() => toggleTag(tag)}
				onkeydown={(e) => e.key === 'Enter' && toggleTag(tag)}
			>
				{tag.Name}
			</span>
		{/each}
	</div>
	<button class="btn" onclick={handleSubmit}>Submit</button>
</div>
