<script>
	import { page } from '$app/stores';
	import GuidelinesModal from './GuidelinesModal.svelte';
	const currentPath = $derived($page.url.pathname);
	const isEventPage = $derived(currentPath.includes('/event/'));
	const isContributePage = $derived(currentPath.includes('/contribute'));

	let countries = $state(['Mexico', 'United Stated', 'Canada']);
	let selectedCountry = $state('');
	let showGuidelinesModal = $state(false);

	function openGuidelinesModal() {
		showGuidelinesModal = true;
	}
</script>

<div class="header-wrapper">
	<div class="header-container">
		{#if isEventPage || isContributePage}
			<a href="/" class="btn"> Home </a>
		{:else}
			<div class="relative">
				<select bind:value={selectedCountry} class="country-select">
					<option value="">Select Country</option>
					{#each countries as country}
						<option value={country}>{country}</option>
					{/each}
				</select>
				<div class="arrow-container">
					<svg class="arrow-svg" fill="none" stroke="currentColor" viewBox="0 0 24 24">
						<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7"
						></path>
					</svg>
				</div>
			</div>
		{/if}
		{#if isEventPage}
			<a href="/correction" class="btn"> Edit </a>
		{:else if !isEventPage && !isContributePage}
			<a href="/contribute" class="btn"> Contribute </a>
		{:else if isContributePage}
			<div class="contribution-text">Contributions are welcome!</div>
			<div class="contribution-info">
				<button onclick={openGuidelinesModal} class="cursor-pointer text-blue-600">
					Guidelines
				</button>
				for contributing.
			</div>
		{/if}
	</div>
</div>

<GuidelinesModal bind:showGuidelinesModal />
