<script>
	let { showMediaModal = $bindable(), media = $bindable() } = $props();
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
				url: youtubeUrl,
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
			{#each media as m, i}
				<div class="flex items-center justify-between border-b border-gray-300 p-2">
					<div class="flex items-center space-x-4">
						<span class="font-medium">{m.title}</span>
						<span class="text-sm capitalize text-gray-600">{m.type}</span>
					</div>
					<button class="btn-danger" onclick={() => removeMedia(i)}>Remove</button>
				</div>
			{/each}
			<div class="mt-6 flex justify-between">
				<div class="flex-2/3 mr-6">
					<label for="photo-title" class="label">Photo title</label>
					<input id="photo-title" type="text" class="input" bind:value={photoTitle} />
				</div>
				<div class="flex-1/3">
					<label for="photo-file" class="label">Photo Upload</label>
					<input id="photo-file" type="file" class="input" bind:files={photoFile} />
				</div>
			</div>
			<div class="mt-6 flex justify-between">
				<div class="flex-2/3 mr-6">
					<label for="youtube-title" class="label">Youtube description</label>
					<input id="youtube-title" class="input" type="text" bind:value={youtubeTitle} />
				</div>
				<div class="flex-1/3">
					<label for="youtube-url" class="label">Youtube Upload</label>
					<input id="youtube-url" type="text" class="input" bind:value={youtubeUrl} />
				</div>
			</div>

			<div class="mt-2 flex justify-between">
				<button class="btn-secondary mt-6" onclick={addMedia} disabled={!canSubmit}>Add</button>
				<button class="btn-secondary mt-6" onclick={closeModal}>Close</button>
			</div>
		</div>
	</div>
{/if}
