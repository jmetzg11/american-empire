export async function getEvents(formData) {
	try {
		const url = `${import.meta.env.VITE_API_URL}/`;
		const response = await fetch(url);
		return await response.json();
	} catch (error) {
		console.error('Error posting blog', error);
		return false;
	}
}
