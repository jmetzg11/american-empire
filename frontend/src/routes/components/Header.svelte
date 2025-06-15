<script>
	import { page } from '$app/stores';

	const currentPath = $derived($page.url.pathname);
	const isEventPage = $derived(currentPath.includes('/event/'));
	const isContributePage = $derived(currentPath.includes('/contribute'));

	let countries = $state(['Mexico', 'United Stated', 'Canada']);
	let selectedCountry = $state('');
</script>

<div class="border-b border-gray-300 bg-white p-4 shadow-sm">
	<div class="mx-auto flex max-w-4xl items-center justify-between">
		{#if isEventPage || isContributePage}
			<a
				href="/"
				class="inline-block rounded-lg bg-blue-600 px-6 py-2 font-medium text-white transition-colors duration-200 hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2"
			>
				Home
			</a>
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
			<a
				href="/correction"
				class="inline-block rounded-lg bg-blue-600 px-6 py-2 font-medium text-white transition-colors duration-200 hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2"
			>
				Edit
			</a>
		{:else if !isEventPage && !isContributePage}
			<a
				href="/contribute"
				class="inline-block rounded-lg bg-blue-600 px-6 py-2 font-medium text-white transition-colors duration-200 hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2"
			>
				Contribute
			</a>
		{:else}
			<div></div>
		{/if}
	</div>
</div>
