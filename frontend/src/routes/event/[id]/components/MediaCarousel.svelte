<script>
	let { media } = $props();
	import PhotoModal from './PhotoModal.svelte';
	let currentIndex = $state(0);
	let showPhotoModal = $state(false);
	let selectedMedia = $state(null);

	function nextSlide() {
		currentIndex = (currentIndex + 1) % media.length;
	}

	function prevSlide() {
		currentIndex = (currentIndex - 1 + media.length) % media.length;
	}

	function openPhotoModal(mediaItem) {
		selectedMedia = mediaItem;
		showPhotoModal = true;
	}
</script>

<div class="container-wrap">
	<div class="card-style h-70 flex items-center justify-between">
		<button class="carousel-button" onclick={prevSlide}>‹</button>
		<div class="flex h-full w-full flex-col items-center justify-center p-4">
			{#if media[currentIndex].Type === 'photo'}
				<button
					class="h-full w-full cursor-pointer"
					onclick={() => openPhotoModal(media[currentIndex])}
				>
					<img
						src={`${import.meta.env.VITE_PHOTO_URL}/${media[currentIndex].Path}`}
						alt={media[currentIndex].Caption}
						class="h-full w-full object-contain"
					/>
				</button>
			{:else}
				<iframe
					src="https://www.youtube.com/embed/{media[currentIndex].URL}"
					title="YouTube video player"
					class="h-full w-full rounded border"
					frameborder="0"
					allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture"
					allowfullscreen
				></iframe>
			{/if}
			<div class="mt-3 text-sm text-gray-500">{media[currentIndex].Caption}</div>
		</div>

		<button class="carousel-button" onclick={nextSlide}>›</button>
	</div>
</div>

<PhotoModal bind:showPhotoModal media={selectedMedia} />
