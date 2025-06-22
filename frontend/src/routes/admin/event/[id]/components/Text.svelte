<script>
	let { title, text, date, onTitleChange, onDateChange, onTextChange } = $props();
	import { formatDate } from '$lib/helpers';
</script>

<div class="container-wrap">
	<div class="card-style">
		<label for="title" class="label">Title</label>
		<input
			id="title"
			type="text"
			value={title}
			oninput={(e) => onTitleChange(e.target.value)}
			class="input"
		/>

		<input
			id="date"
			type="date"
			value={formatDate(date)}
			oninput={(e) => onDateChange(e.target.value)}
			class="input"
		/>
		<label for="text" class="label">Description</label>
		<textarea
			id="text"
			class="input h-128"
			value={text}
			oninput={(e) => onTextChange(e.target.value)}
		></textarea>
	</div>
</div>
