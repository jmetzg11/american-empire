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

export async function getEvent(id) {
	try {
		const url = `${import.meta.env.VITE_API_URL}/event`;
		const response = await fetch(url, {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json'
			},
			body: JSON.stringify({ id })
		});
		return await response.json();
	} catch (error) {
		console.error('Error fetching event', error);
		return false;
	}
}

export async function contributeEvent(formData) {
	try {
		const url = `${import.meta.env.VITE_API_URL}/contribute`;
		const response = await fetch(url, {
			method: 'POST',
			body: formData
		});

		const data = await response.json();
		console.log(data);

		return {
			ok: response.ok,
			status: response.status,
			data: data
		};
	} catch (error) {
		console.error('Error contributing event', error);
		return {
			ok: false,
			status: 500,
			data: { success: false, message: 'Failed to submit event' }
		};
	}
}
