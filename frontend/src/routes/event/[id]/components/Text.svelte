<script>
	let { title, tags, text, date } = $props();
	import { formatDate } from '$lib/helpers';
</script>

<div class="container-wrap">
	<div class="card-style">
		<h1 class="section-title">{title}</h1>
		<div class="flex flex-wrap gap-2">
			{#each tags as tag}
				<span class="tag">{tag.Name}</span>
			{/each}
		</div>
		<p class="normal-text my-4">{formatDate(date)}</p>
		<p class="normal-text whitespace-pre-wrap leading-relaxed">{text}</p>
	</div>
</div>
