<script>
	import { page } from '$app/stores';
	import { uploadPhoto, uploadYoutube } from '$lib/api';
	import { extractYouTubeId } from '$lib/helpers';
	let { type, refreshEvent } = $props();

	let title = $state('');
	let file = $state(null);
	let url = $state('');

	let canSubmit = $derived(title && (file || url));

	async function upload() {
		const formData = new FormData();
		formData.append('event_id', $page.params.id);
		formData.append('title', title);
		let result;
		if (type === 'photo') {
			formData.append('file', file[0]);
			result = await uploadPhoto(formData);
		} else {
			formData.append('url', extractYouTubeId(url));
			result = await uploadYoutube(formData);
		}
		if (result.ok) {
			await refreshEvent();
			title = '';
			file = null;
			url = '';
		}
	}
</script>

<h2 class="new-info-modal-title mt-6 text-center">
	{type === 'photo' ? 'Add Photo' : 'Add Youtube'}
</h2>
<label for="photo-title" class="label">Photo Caption</label>
<input id="photo-title" type="text" class="input" bind:value={title} />
<label for="photo-file" class="label">Photo Upload</label>
{#if type === 'photo'}
	<input id="photo-file" type="file" class="input" bind:files={file} />
{:else}
	<input id="youtube-url" type="text" class="input" bind:value={url} />
{/if}
<div class="mt-2 flex justify-center">
	<button class="btn mt-2 items-center" disabled={!canSubmit} onclick={upload}>Upload</button>
</div>
