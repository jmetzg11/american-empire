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
		<div class="mb-4 text-sm text-gray-500">
			can edit captions, but must delete photo/vidoes to change them
		</div>

		<div class="grid grid-cols-1 gap-4 sm:grid-cols-2 lg:grid-cols-4">
			{#each media as mediaItem}
				<div class="flex flex-col">
					<div class="aspect-square">
						{#if mediaItem.Type === 'photo'}
							<button
								class="h-full w-full cursor-pointer"
								onclick={() => openPhotoModal(mediaItem)}
							>
								<img
									src={`${import.meta.env.VITE_PHOTO_URL}/${mediaItem.Path}`}
									alt={mediaItem.Caption}
									class="h-full w-full rounded object-cover"
								/></button
							>
						{:else}
							<iframe
								src="https://www.youtube.com/embed/{mediaItem.URL}"
								title="YouTube video player"
								class="h-full w-full rounded"
								frameborder="0"
								allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture"
								allowfullscreen
							></iframe>
						{/if}
					</div>
					{#if mediaItem.Caption}
						<input
							type="text"
							class="input my-2"
							value={mediaItem.Caption}
							oninput={(e) =>
								onFKChange(mediaItem.Type + '-' + mediaItem.ID, 'Caption', e.target.value)}
						/>
						<button class="btn-danger" onclick={() => handleDelete(mediaItem)}>Remove</button>
					{/if}
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
