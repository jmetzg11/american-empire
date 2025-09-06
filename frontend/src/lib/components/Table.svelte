<script>
	import { events } from '$lib/store';
	import { formatDate } from '$lib/helpers';
	import { goto } from '$app/navigation';

	function handleRowClick(event) {
		goto(`/event/${event.id}`);
	}
</script>

<div class="container-wrap">
	<div class="card-style overflow-hidden">
		<table class="w-full">
			<thead class="table-header">
				<tr>
					<th class="th">Country</th>
					<th class="th">Title</th>
					<th class="th">Date</th>
					<th class="th">Tags</th>
				</tr>
			</thead>
			<tbody class="divide-y divide-gray-100">
				{#each $events as event}
					<tr
						class="cursor-pointer transition-colors hover:bg-gray-50"
						on:click={() => handleRowClick(event)}
					>
						<td class="td">{event.country}</td>
						<td class="td">{event.title}</td>
						<td class="td">{formatDate(event.date)}</td>
						<td class="td">
							<div class="flex flex-wrap gap-1">
								{#each event.tags as tag}
									<span class="tag">
										{tag}
									</span>
								{/each}
							</div>
						</td>
					</tr>
				{/each}
			</tbody>
		</table>
	</div>
</div>
