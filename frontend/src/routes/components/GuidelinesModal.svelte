<script>
	let { showGuidelinesModal = $bindable() } = $props();

	function closeModal() {
		showGuidelinesModal = false;
	}
</script>

{#if showGuidelinesModal}
	<div class="modal-wrapper">
		<div class="modal-content !max-w-2xl">
			<h2 class="section-title">Guidelines</h2>
			<div class="space-y-4">
				<div>
					<h3 class="section-subtitle">Event Submissions</h3>
					<p class="normal-text mb-2">
						Descriptions should be detailed but concise as they serve as reference material. Every
						event must include:
					</p>
					<ul class="bullet-list mb-2">
						<li>Title</li>
						<li>Date</li>
						<li>Source</li>
					</ul>
					<p class="normal-text">Site editors have final say on what gets published.</p>
				</div>

				<div>
					<h3 class="section-subtitle">Code Contributions</h3>
					<p class="normal-text">Contributions to the codebase are also welcome.</p>
				</div>
			</div>

			<button onclick={closeModal} class="btn mt-6"> Close </button>
		</div>
	</div>
{/if}
