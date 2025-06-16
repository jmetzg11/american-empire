<script>
	let { title, text, date } = $props();
	import { formatDate } from '$lib/helpers';
</script>

<div class="container-wrap">
	<div class="card-style">
		<h1 class="mb-2 text-3xl font-bold text-gray-900">{title}</h1>
		<p class="mb-4 text-sm text-gray-500">{formatDate(date)}</p>
		<p class="whitespace-pre-wrap leading-relaxed text-gray-700">{text}</p>
	</div>
</div>
