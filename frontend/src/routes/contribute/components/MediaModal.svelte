<script>
	let { showMediaModal = $bindable(), media = $bindable() } = $props();
	import { extractYouTubeId } from '$lib/helpers';
	let photoTitle = $state('');
	let photoFile = $state(null);
	let youtubeTitle = $state('');
	let youtubeUrl = $state('');
	let canSubmit = $derived(
		(photoFile && photoFile[0] && photoTitle) || (youtubeUrl && youtubeTitle)
	);

	function addMedia() {
		if (photoFile?.[0] && photoTitle) {
			media.push({
				type: 'photo',
				file: photoFile[0],
				title: photoTitle
			});
			photoTitle = '';
			photoFile = null;
		}
		if (youtubeUrl && youtubeTitle) {
			media.push({
				type: 'youtube',
				url: extractYouTubeId(youtubeUrl),
				title: youtubeTitle
			});
			youtubeTitle = '';
			youtubeUrl = '';
		}
	}

	function removeMedia(index) {
		media.splice(index, 1);
	}

	function closeModal() {
		showMediaModal = false;
	}
</script>

{#if showMediaModal}
	<div class="modal-wrapper">
		<div class="modal-content">
			<h2 class="new-info-modal-title">Existing Media</h2>
			{#each media as m, i}
				<div class="existing-item-row">
					<div class="existing-item-text">
						<span class="existing-item-title">{m.title}</span>
						<span class="existing-item-type">{m.type}</span>
					</div>
					<button class="btn-danger" onclick={() => removeMedia(i)}>Remove</button>
				</div>
			{/each}
			<h2 class="new-info-modal-title mt-6">Add Media</h2>
			<div class="new-info-input-container">
				<div class="new-info-input-left">
					<label for="photo-title" class="label">Photo title</label>
					<input id="photo-title" type="text" class="input" bind:value={photoTitle} />
				</div>
				<div class="new-info-input-right">
					<label for="photo-file" class="label">Photo Upload</label>
					<input id="photo-file" type="file" class="input" bind:files={photoFile} />
				</div>
			</div>
			<div class="new-info-input-container">
				<div class="new-info-input-left">
					<label for="youtube-title" class="label">Youtube description</label>
					<input id="youtube-title" class="input" type="text" bind:value={youtubeTitle} />
				</div>
				<div class="new-info-input-right">
					<label for="youtube-url" class="label">Youtube Upload</label>
					<input id="youtube-url" type="text" class="input" bind:value={youtubeUrl} />
				</div>
			</div>
			<div class="new-info-button-container">
				<button class="btn-secondary mt-6" onclick={addMedia} disabled={!canSubmit}>Add</button>
				<button class="btn-secondary mt-6" onclick={closeModal}>Close</button>
			</div>
		</div>
	</div>
{/if}
