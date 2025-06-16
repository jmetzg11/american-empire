<script>
	let { showGuidelinesModal = $bindable() } = $props();

	function closeModal() {
		showGuidelinesModal = false;
	}
</script>

{#if showGuidelinesModal}
	<div class="modal-wrapper">
		<div class="modal-content">
			<h2 class="mb-4 text-xl font-bold">Guidelines</h2>

			<div class="space-y-4">
				<div>
					<h3 class="mb-2 font-semibold text-gray-800">Event Submissions</h3>
					<p class="mb-2 text-gray-600">
						Descriptions should be detailed but concise as they serve as reference material. Every
						event must include:
					</p>
					<ul class="mb-2 ml-4 list-disc text-sm text-gray-600">
						<li>Title</li>
						<li>Date</li>
						<li>Source</li>
					</ul>
					<p class="text-gray-600">Site editors have final say on what gets published.</p>
				</div>

				<div>
					<h3 class="mb-2 font-semibold text-gray-800">Code Contributions</h3>
					<p class="text-gray-600">Contributions to the codebase are also welcome.</p>
				</div>
			</div>

			<button onclick={closeModal} class="btn mt-6"> Close </button>
		</div>
	</div>
{/if}
