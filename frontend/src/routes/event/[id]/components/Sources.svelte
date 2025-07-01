<script>
	let { sources } = $props();
</script>

<div class="container-wrap">
	<div class="card-style">
		<h2 class="section-title">Sources</h2>
		{#each sources as source}
			<div class="source-row">
				<div class="text-lg">{source.Name}</div>
				<a href={source.URL} target="_blank" class="a-link">Link</a>
			</div>
		{/each}
	</div>
</div>
