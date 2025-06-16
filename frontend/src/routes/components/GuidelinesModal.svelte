<script>
	let { showGuidelinesModal = $bindable(false) } = $props();

	function closeModal() {
		showGuidelinesModal = false;
	}
</script>

{#if showGuidelinesModal}
	<div class="fixed inset-0 z-50 flex items-center justify-center bg-black bg-opacity-50">
		<div class="mx-4 w-full max-w-md rounded-lg bg-white p-6">
			<h2 class="mb-4 text-xl font-bold">Guidelines</h2>
			<p>Your guidelines content here...</p>
			<button
				onclick={closeModal}
				class="mt-4 rounded bg-blue-600 px-4 py-2 text-white hover:bg-blue-700"
			>
				Close
			</button>
		</div>
	</div>
{/if}
