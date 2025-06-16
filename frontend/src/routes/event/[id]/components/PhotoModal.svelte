<script>
	let { showPhotoModal = $bindable(), media = null } = $props();

	function handleClose() {
		showPhotoModal = false;
	}

	function handleKeydown(e) {
		if (e.key === 'Escape') {
			handleClose();
		}
	}
</script>

{#if showPhotoModal}
	<div
		role="button"
		tabindex="0"
		aria-label="Close photo modal background"
		class="modal-wrapper"
		onclick={handleClose}
		onkeydown={handleKeydown}
	>
		<div class="relative max-h-[90vh] max-w-4xl overflow-hidden rounded-lg bg-white shadow-xl">
			<div class="p-4">
				<img
					src={`${import.meta.env.VITE_PHOTO_URL}/${media.Path}`}
					alt="Selected item in full size"
					class="max-h-[75vh] max-w-full object-contain"
				/>
				{#if media.Caption}
					<p class="mt-3 text-center text-gray-700">{media.Caption}</p>
				{/if}
			</div>
			<button
				aria-label="Close photo modal"
				class="absolute right-2 top-2 cursor-pointer rounded-full bg-black/50 p-2 text-white transition-colors hover:bg-black/70"
				onclick={handleClose}
			>
				<svg
					xmlns="http://www.w3.org/2000/svg"
					width="24"
					height="24"
					viewBox="0 0 24 24"
					fill="none"
					stroke="currentColor"
					stroke-width="2"
					stroke-linecap="round"
					stroke-linejoin="round"
					class="h-5 w-5"
				>
					<line x1="18" y1="6" x2="6" y2="18"></line>
					<line x1="6" y1="6" x2="18" y2="18"></line>
				</svg>
			</button>
		</div>
	</div>
{/if}
