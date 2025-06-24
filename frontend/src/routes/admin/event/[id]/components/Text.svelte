<script>
	let { title, country, text, date, onTitleChange, onCountryChange, onDateChange, onTextChange } =
		$props();
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
		<label for="country" class="label">Country</label>
		<input
			id="country"
			type="text"
			value={country}
			oninput={(e) => onCountryChange(e.target.value)}
			class="input"
		/>
		<label for="date" class="label mt-2">Date</label>
		<input
			id="date"
			type="date"
			value={formatDate(date)}
			oninput={(e) => onDateChange(e.target.value)}
			class="input"
		/>
		<label for="text" class="label mt-2">Description</label>
		<textarea
			id="text"
			class="input h-128"
			value={text}
			oninput={(e) => onTextChange(e.target.value)}
		></textarea>
	</div>
</div>
