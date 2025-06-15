<script>
	let { sources } = $props();

	// $inspect(sources);
</script>

<h1>Sources</h1>

{#each sources as source}
	<div>{source.Name}</div>
{/each}
