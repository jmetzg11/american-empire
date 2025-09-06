export async function getEvents(formData) {
	try {
		const url = '/api/';
		const response = await fetch(url);
		const data = await response.json();

		if (!response.ok) {
			console.error('Server error:', data.error || data.message);
		}

		return data;
	} catch (error) {
		console.error('Error posting blog', error);
		return false;
	}
}

export async function getEvent(id) {
	try {
		const url = '/api/event';
		const response = await fetch(url, {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json'
			},
			body: JSON.stringify({ id })
		});
		const data = await response.json();

		if (!response.ok) {
			console.error('Server error:', data.error || data.message);
		}

		return data;
	} catch (error) {
		console.error('Error fetching event', error);
		return false;
	}
}

export async function contributeEvent(formData) {
	try {
		const url = '/api/contribute';
		const response = await fetch(url, {
			method: 'POST',
			body: formData
		});

		const data = await response.json();

		if (!response.ok) {
			console.error('Server error:', data.error || data.message);
		}

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

export async function getTags() {
	try {
		const url = '/api/tags';
		const response = await fetch(url);
		const data = await response.json();

		if (!response.ok) {
			console.error('Server error:', data.error || data.message);
		}

		return data;
	} catch (error) {
		console.error('Error fetching tags', error);
		return false;
	}
}

export async function getBook(bookId) {
    try {
        const url = `/api/book/${bookId}` 
        const response = await fetch(url)
        const data = await response.json()

        if (!response.ok) {
            console.error(`Server error:`, data.error || data.message)
        }

        return data
    } catch (error) {
        console.error("Error fetching book", error)
        return false
    }
}

