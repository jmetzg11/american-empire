<script>
	import { onMount } from 'svelte';
	import { getBooks } from '$lib/api';
	import Books from './components/Books.svelte';
	import NewBook from './components/NewBook.svelte';
	let books = $state([]);

	onMount(async () => {
		try {
			let data = await getBooks();
			books = data;
			console.log(data);
		} catch (error) {
			console.error('Faile to get books', error);
		}
	});
</script>

<h1>Books</h1>
<Books {books} />
<NewBook />
