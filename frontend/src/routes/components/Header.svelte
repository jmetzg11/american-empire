<script>
	let countries = $state(['Mexico', 'United Stated', 'Canada']);
	let selectedCountry = $state('');
</script>

<div class="flex items-center justify-between border-b bg-white p-4 shadow-sm">
	<div class="relative">
		<select
			bind:value={selectedCountry}
			class="cursor-pointer appearance-none rounded-lg border border-gray-300 bg-white p-2 pr-8 focus:border-transparent focus:outline-none focus:ring-2 focus:ring-blue-500"
		>
			<option value="">Select Country</option>
			{#each countries as country}
				<option value={country}>{country}</option>
			{/each}
		</select>
		<div class="pointer-events-none absolute inset-y-0 right-0 flex items-center px-2">
			<svg class="h-4 w-4 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
				<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7"
				></path>
			</svg>
		</div>
	</div>
	<a
		href="/contribute"
		class="inline-block rounded-lg bg-blue-600 px-6 py-2 text-center font-medium text-white transition-colors duration-200 hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2"
	>
		Contribute
	</a>
</div>
