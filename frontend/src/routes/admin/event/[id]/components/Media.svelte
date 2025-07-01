<script>
	let { media, onFKChange, refreshEvent } = $props();
	import PhotoModal from '$lib/components/PhotoModal.svelte';
	import NewMedia from './NewMedia.svelte';
	import DeleteMediaModal from './DeleteMediaModal.svelte';

	let showDeleteModal = $state(false);
	let showPhotoModal = $state(false);
	let selectedMedia = $state(null);

	function openPhotoModal(mediaItem) {
		selectedMedia = mediaItem;
		showPhotoModal = true;
	}

	function handleDelete(mediaItem) {
		selectedMedia = mediaItem;
		showDeleteModal = true;
	}
</script>

<div class="container-wrap">
	<div class="card-style">
		<div class="small-info-text mb-4">
			can edit captions, but must delete photo/vidoes to change them
		</div>

		<div class="grid-section">
			{#each media as mediaItem}
				<div class="flex flex-col">
					<div class="aspect-square">
						{#if mediaItem.Type === 'photo'}
							<button class="img-button" onclick={() => openPhotoModal(mediaItem)}>
								<img
									src={`${import.meta.env.VITE_PHOTO_URL}/${mediaItem.Path}`}
									alt={mediaItem.Caption}
									class="img-preview"
								/></button
							>
						{:else}
							<iframe
								src="https://www.youtube.com/embed/{mediaItem.URL}"
								title="YouTube video player"
								class="youtube-preview"
								frameborder="0"
								allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture"
								allowfullscreen
							></iframe>
						{/if}
					</div>
					<input
						type="text"
						class="input my-2"
						value={mediaItem.Caption}
						oninput={(e) => onFKChange('media-' + mediaItem.ID, { Caption: e.target.value })}
					/>
					<button class="btn-danger" onclick={() => handleDelete(mediaItem)}>Remove</button>
				</div>
			{/each}
			<div>
				<NewMedia type="photo" {refreshEvent} />
			</div>
			<div>
				<NewMedia type="youtube" {refreshEvent} />
			</div>
		</div>
	</div>
</div>

<PhotoModal bind:showPhotoModal media={selectedMedia} />
<DeleteMediaModal bind:showDeleteModal media={selectedMedia} {refreshEvent} />
