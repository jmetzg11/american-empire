<script>
	import { page } from '$app/stores';
	import GuidelinesModal from './GuidelinesModal.svelte';
	const currentPath = $derived($page.url.pathname);
	const isEventPage = $derived(currentPath.includes('/event/'));
	const isContributePage = $derived(currentPath.includes('/contribute'));
	const isAdminPage = $derived(currentPath.includes('/admin'));

	let countries = $state(['Mexico', 'United Stated', 'Canada']);
	let selectedCountry = $state('');
	let showGuidelinesModal = $state(false);

	function openGuidelinesModal() {
		console.log('openGuidelinesModal');
		showGuidelinesModal = true;
	}
</script>

<div class="border-b border-gray-300 bg-white p-4 shadow-sm">
	<div class="mx-auto flex max-w-4xl items-center justify-between">
		{#if isEventPage || isContributePage || isAdminPage}
			<a href="/" class="btn"> Home </a>
		{:else}
			<div class="relative">
				<select
					bind:value={selectedCountry}
					class="cursor-pointer appearance-none rounded-lg border border-gray-300 bg-white px-4 py-2 pr-10 font-medium text-gray-900 transition-colors duration-200 focus:border-transparent focus:outline-none focus:ring-2 focus:ring-blue-500"
				>
					<option value="">Select Country</option>
					{#each countries as country}
						<option value={country}>{country}</option>
					{/each}
				</select>
				<div class="pointer-events-none absolute inset-y-0 right-0 flex items-center px-3">
					<svg class="h-4 w-4 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
						<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7"
						></path>
					</svg>
				</div>
			</div>
		{/if}
		{#if isEventPage}
			<a href="/correction" class="btn"> Edit </a>
		{:else if !isEventPage && !isContributePage && !isAdminPage}
			<a href="/contribute" class="btn"> Contribute </a>
		{:else if isContributePage}
			<div class="text-2xl font-bold text-gray-700">Contributions are welcome!</div>
			<div class="text-sm text-gray-500">
				<button onclick={openGuidelinesModal} class="cursor-pointer text-blue-600">
					Guidelines
				</button>
				for contributing.
			</div>
		{/if}
	</div>
</div>

<GuidelinesModal bind:showGuidelinesModal />
