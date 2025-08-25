<script>
	let { books } = $props();
	import BookEditModal from './BookEditModal.svelte';

	let showModal = $state(false);
	let bookId = $state('');

	function editBook(id) {
		bookId = id;
		showModal = true;
	}
</script>

<div class="container-wrap">
	<div class="card-style overflow-hidden">
		<table class="w-full">
			<thead class="table-header">
				<tr>
					<th class="th">title</th>
					<th class="th">author</th>
					<th class="th">event ids</th>
					<th class="th">tags</th>
				</tr>
			</thead>
			<tbody class="divide-y divide-gray-100">
				{#each books as book}
					<tr
						class="hover:gb-gray-50 cursor-pointer transition-colors"
						onclick={() => editBook(book.id)}
					>
						<td class="td">{book.title}</td>
						<td class="td">{book.author}</td>
						<td class="td">{book.events.join(', ')}</td>
						<td class="td">
							{#each book.tags as tag}
								<span class="tag mr-1">{tag}</span>
							{/each}
						</td>
					</tr>
				{/each}
			</tbody>
		</table>
	</div>
</div>

<BookEditModal bind:showModal {bookId} />
