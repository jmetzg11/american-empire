<script>
	import { events } from '$lib/store';
	import { formatDate } from '$lib/helpers';
	import { goto } from '$app/navigation';

	function handleRowClick(event) {
		goto(`/event/${event.id}`);
	}
</script>

<div class="mx-auto max-w-6xl p-4">
	<div class="overflow-hidden rounded-xl border border-gray-200 bg-white shadow-sm">
		<table class="w-full">
			<thead class="border-b border-gray-200 bg-gray-50">
				<tr>
					<th class="px-6 py-4 text-left text-sm font-medium text-gray-700">Country</th>
					<th class="px-6 py-4 text-left text-sm font-medium text-gray-700">Title</th>
					<th class="px-6 py-4 text-left text-sm font-medium text-gray-700">Date</th>
				</tr>
			</thead>
			<tbody class="divide-y divide-gray-100">
				{#each $events as event}
					<tr
						class="cursor-pointer transition-colors hover:bg-gray-50"
						on:click={() => handleRowClick(event)}
					>
						<td class="px-6 py-4 text-sm font-semibold text-gray-900">{event.country}</td>
						<td class="px-6 py-4 text-sm text-gray-600">{event.title}</td>
						<td class="px-6 py-4 text-sm text-gray-600">{formatDate(event.date)}</td>
					</tr>
				{/each}
			</tbody>
		</table>
	</div>
</div>
