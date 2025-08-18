<script>
	import { onMount } from 'svelte';
	import NewBook from './components/NewBook.svelte';
	onMount(async () => {
		try {
			books = await getBooks();
		} catch (error) {
			console.error('Error in onMount', error);
		}
	});
</script>

<h1>Books</h1>

<NewBook />
