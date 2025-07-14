import adapter from '@sveltejs/adapter-node'; // or adapter-auto
import { vitePreprocess } from '@sveltejs/vite-plugin-svelte';

const config = {
	extensions: ['.svelte'],
	preprocess: [vitePreprocess()],
	kit: {
		adapter: adapter(),
		csrf: { checkOrigin: false }
	}
};
export default config;
