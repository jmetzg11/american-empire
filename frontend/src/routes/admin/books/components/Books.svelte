<script>
	let { books } = $props();

	function editBook(id) {
		console.log(`book ${id} was clicked`);
	}
</script>

<div class="container-wrap">
	<div class="card-style overflow-hidden">
		<table class="w-full">
			<thead class="table-header">
				<tr>
					<th class="th">Title</th>
					<th class="th">Author</th>
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
					</tr>
				{/each}
			</tbody>
		</table>
	</div>
</div>
