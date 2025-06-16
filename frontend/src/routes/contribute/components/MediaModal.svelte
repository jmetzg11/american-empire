<script>
	let { showMediaModal = $bindable(), media = $bindable() } = $props();
	let photoTitle = $state('');
	let photoFile = $state(null);
	let youtubeTitle = $state('');
	let youtubeUrl = $state('');
	let canSubmit = $derived((photoFile?.[0] && photoTitle) || (youtubeUrl && youtubeTitle));
	$inspect(canSubmit);

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

	function closeModal() {
		showMediaModal = false;
	}
</script>

{#if showMediaModal}
	<div class="modal-wrapper">
		<div class="modal-content space-y-6">
			{#each media as m}
				<div>{media.title}</div>
				<div>{media.type}</div>
				<button class="btn-3">Remove</button>
			{/each}
			<div class="flex justify-between">
				<div class="flex-2/3 mr-6">
					<label for="photo-title" class="label">Photo title</label>
					<input id="photo-title" type="text" class="input" bind={photoTitle} />
				</div>
				<div class="flex-1/3">
					<label for="photo-file" class="label">Photo Upload</label>
					<input id="photo-file" type="file" class="input" bind={photoFile} />
				</div>
			</div>
			<div class="flex justify-between">
				<div class="flex-2/3 mr-6">
					<label for="youtube-title" class="label">Youtube description</label>
					<input id="youtube-title" class="input" type="text" bind={youtubeTitle} />
				</div>
				<div class="flex-1/3">
					<label for="youtube-url" class="label">Youtube Upload</label>
					<input id="youtube-url" type="text" class="input" bind={youtubeUrl} />
				</div>
			</div>

			<div class="flex justify-between">
				{canSubmit}
				<button
					class="btn-2 mt-6 disabled:cursor-not-allowed disabled:bg-gray-400 disabled:text-gray-600"
					onclick={addMedia}
					disabled={canSubmit}>Add</button
				>
				<button class="btn-2 mt-6" onclick={closeModal}>Close</button>
			</div>
		</div>
	</div>
{/if}
