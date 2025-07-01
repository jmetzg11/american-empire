<script>
	let { title, text, date } = $props();
	import { formatDate } from '$lib/helpers';
</script>

<div class="container-wrap">
	<div class="card-style">
		<h1 class="section-title">{title}</h1>
		<p class="normal-text mb-4">{formatDate(date)}</p>
		<p class="normal-text whitespace-pre-wrap leading-relaxed">{text}</p>
	</div>
</div>
