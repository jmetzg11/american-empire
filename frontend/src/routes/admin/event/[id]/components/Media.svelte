<script>
	let { media } = $props();
	$inspect(media);
	console.log(`${import.meta.env.VITE_PHOTO_URL}/${media.Path}`);
</script>

<div class="container-wrap">
	<div class="card-style">
		<div class="grid grid-cols-1 gap-4 sm:grid-cols-2 lg:grid-cols-4">
			{#each media as mediaItem}
				<div class="flex flex-col">
					<div class="aspect-square">
						{#if mediaItem.Type === 'photo'}
							<img
								src={`${import.meta.env.VITE_PHOTO_URL}/${mediaItem.Path}`}
								alt={mediaItem.Caption}
								class="h-full w-full rounded object-cover"
							/>
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
						<p class="mt-2 text-sm text-gray-600">{mediaItem.Caption}</p>
					{/if}
				</div>
			{/each}
		</div>
	</div>
</div>
