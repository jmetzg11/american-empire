<script>
	let { sources } = $props();
</script>

<div class="container-wrap">
	<div class="card-style">
		<h2 class="mb-2 text-2xl font-bold text-gray-900">Sources</h2>
		{#each sources as source}
			<div class="flex items-center justify-between border-b border-gray-200 py-3 last:border-b-0">
				<div class="text-gray-900">{source.Name}</div>
				<a href={source.URL} target="_blank" class="text-blue-500 hover:underline">Link</a>
			</div>
		{/each}
	</div>
</div>
