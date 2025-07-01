<script>
	import { contributeEvent } from '$lib/api';
	import MediaModal from './MediaModal.svelte';
	import SourceModal from './SourceModal.svelte';
	let country = $state('');
	let title = $state('');
	let date = $state('');
	let description = $state('');
	let email = $state('');
	let media = $state([]);
	let source = $state([]);
	let ableToSubmit = $derived(country && title && date && description && source.length > 0);
	let showMediaModal = $state(false);
	let showSourceModal = $state(false);
	let submitted = $state(false);

	function autoGrow(event) {
		event.target.style.height = 'auto';
		event.target.style.height = event.target.scrollHeight + 'px';
	}

	async function submitEvent() {
		const formData = new FormData();

		formData.append('country', country);
		formData.append('title', title);
		formData.append('date', date);
		formData.append('description', description);
		formData.append('email', email);

		source.forEach((src, index) => {
			formData.append(`source[${index}][name]`, src.name);
			formData.append(`source[${index}][url]`, src.url);
		});

		media.forEach((item, index) => {
			if (item.type === 'photo') {
				formData.append(`media[${index}][file]`, item.file);
				formData.append(`media[${index}][type]`, 'photo');
				formData.append(`media[${index}][caption]`, item.title);
			} else if (item.type === 'youtube') {
				formData.append(`media[${index}][type]`, 'youtube');
				formData.append(`media[${index}][url]`, item.url);
				formData.append(`media[${index}][caption]`, item.title);
			}
		});

		const response = await contributeEvent(formData);
		if (response.ok) {
			submitted = true;
			country = '';
			title = '';
			date = '';
			description = '';
			email = '';
			source = [];
			media = [];
		}
	}
</script>

<div class="container-wrap">
	<div class="card-style">
		{#if !submitted}
			<div class="flex-col-gap-6">
				<div>
					<label for="title" class="label">Title *</label>
					<input
						id="title"
						type="text"
						bind:value={title}
						class="input"
						placeholder="Enter title"
					/>
				</div>
				<div class="flex-center-between">
					<div>
						<label for="country" class="label">Country *</label>
						<input
							id="country"
							type="text"
							bind:value={country}
							class="input"
							placeholder="Enter country"
						/>
					</div>
					<div>
						<label for="date" class="label">Date *</label>
						<input id="date" type="date" bind:value={date} class="input" />
					</div>
				</div>
				<div>
					<label for="description" class="label">Description *</label>
					<textarea
						id="description"
						bind:value={description}
						class="input"
						oninput={autoGrow}
						placeholder="Enter description..."
					></textarea>
				</div>
				<div>
					<label for="email" class="label">Your Email (optional)</label>
					<input
						id="email"
						type="email"
						bind:value={email}
						class="input"
						placeholder="Enter email"
					/>
				</div>
				<div class="flex-center-between">
					<button type="button" onclick={() => (showSourceModal = true)} class="btn-secondary">
						Add Source * {source.length > 0 ? `(${source.length})` : ''}
					</button>
					<button type="button" onclick={() => (showMediaModal = true)} class="btn-secondary"
						>Add Media {media.length > 0 ? `(${media.length})` : ''}</button
					>
					<button type="button" onclick={submitEvent} disabled={!ableToSubmit} class="btn">
						Submit
					</button>
				</div>
			</div>
		{:else}
			<div class="mb-4 text-lg">
				Thank you for contributing! We will review your submission and get back to you soon (if you
				incluced your email).
			</div>
			<button class="btn-secondary" onclick={() => (submitted = false)}>Contribute Again</button>
		{/if}
	</div>
</div>

<MediaModal bind:showMediaModal {media} />
<SourceModal bind:showSourceModal {source} />
