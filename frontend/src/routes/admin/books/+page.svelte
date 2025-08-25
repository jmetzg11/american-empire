<script>
	import { onMount } from 'svelte';
	import { getBooks } from '$lib/api';
	import Books from './components/Books.svelte';
	import NewBook from './components/NewBook.svelte';
	let books = $state([]);

	onMount(async () => {
		try {
			let results = await getBooks();
			books = results.data;
		} catch (error) {
			console.error('Faile to get books', error);
		}
	});
</script>

<Books {books} />
<NewBook />
