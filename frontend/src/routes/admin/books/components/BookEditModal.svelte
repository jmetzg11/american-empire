<script>
	let { showModal = $bindable(), bookId } = $props();
	import { getTags, getBook, editBook } from '$lib/api';

	let book = $state({});
	let tags = $state([]);
	let editedBook = $state({});

	$effect(async () => {
		if (showModal && bookId) {
			try {
				const [bookData, tagsData] = await Promise.all([getBook(bookId), getTags()]);
				book = bookData;
				tags = tagsData;

				editedBook = {
					id: bookId,
					title: book.Title || '',
					author: book.Author || '',
					link: book.Link || '',
					events: book.Events ? book.Events.map((e) => e.ID).join(', ') : '',
					selectedTags: book.Tags ? book.Tags.map((t) => t.ID) : []
				};
			} catch (error) {
				console.error('Error loading data:', error);
			}
		}
	});

	function toggleTag(tag) {
		if (editedBook.selectedTags.includes(tag.ID)) {
			editedBook.selectedTags = editedBook.selectedTags.filter((id) => id !== tag.ID);
		} else {
			editedBook.selectedTags = [...editedBook.selectedTags, tag.ID];
		}
	}

	function handleSave() {
		editBook(editedBook);
		showModal = false;
	}
</script>

{#if showModal}
	<div class="modal-wrapper">
		<div class="modal-content">
			<h2 class="section-title">Edit Book</h2>
			<div class="flex-col-gap-2">
				<label for="edit-title" class="label">Title</label>
				<input
					type="text"
					id="edit-title"
					class="input"
					placeholder="Title"
					bind:value={editedBook.title}
				/>

				<label for="edit-author" class="label">Author</label>
				<input
					type="text"
					id="edit-author"
					class="input"
					placeholder="Author"
					bind:value={editedBook.author}
				/>

				<label for="edit-link" class="label">Link</label>
				<input
					type="text"
					id="edit-link"
					class="input"
					placeholder="URL"
					bind:value={editedBook.link}
				/>

				<label for="edit-events" class="label">Event IDs</label>
				<input
					type="text"
					id="edit-events"
					class="input"
					placeholder="1, 2, 3"
					bind:value={editedBook.events}
				/>
			</div>

			<div class="tags-section">
				{#each tags as tag}
					<span
						class={`tag${editedBook.selectedTags.includes(tag.ID) ? '-selected' : ''} mr-2`}
						role="button"
						tabindex="0"
						onclick={() => toggleTag(tag)}
						onkeydown={(e) => e.key === 'Enter' && toggleTag(tag)}
					>
						{tag.Name}
					</span>
				{/each}
			</div>

			<div class="flex-center-between mt-4">
				<button class="btn" onclick={handleSave}>Save</button>
				<button class="btn-secondary" onclick={() => (showModal = false)}>Cancel</button>
			</div>
		</div>
	</div>
{/if}
