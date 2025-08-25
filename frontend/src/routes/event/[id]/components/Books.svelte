<script>
	let { books } = $props();
	console.log(books);
	function showMore() {
		console.log('show more book details');
	}
</script>

<div class="container-wrap">
	<div class="card-style overflow-hidden">
		<table class="w-full">
			<thead class="table-header">
				<tr>
					<th class="th">Title</th>
					<th class="th">Author</th>
					<th class="th">Link</th>
				</tr>
			</thead>
			<tbody class="divide-y divide-gray-100">
				{#each books as book}
					<tr
						class="hover:gb-gray-50 cursor-pointer transition-colors"
						onclick={() => showMore(book.id)}
					>
						<td class="td">{book.Title}</td>
						<td class="td">{book.Author}</td>
						<td class="td">{book.Link}</td>
					</tr>
				{/each}
			</tbody>
		</table>
	</div>
</div>
