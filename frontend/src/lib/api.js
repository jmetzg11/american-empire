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

// Admin

export async function authMe() {
	try {
		const url = `${import.meta.env.VITE_API_URL}/auth-me`;
		const response = await fetch(url, {
			credentials: 'include'
		});
		const data = await response.json();
		return data.authenticated;
	} catch (error) {
		console.error('Error fetching auth me', error);
		return false;
	}
}

export async function getAdminEvents() {
	try {
		const url = `${import.meta.env.VITE_API_URL}/admin-events`;
		const response = await fetch(url, {
			method: 'GET',
			headers: {
				'Content-Type': 'application/json'
			},
			credentials: 'include'
		});
		return await response.json();
	} catch (error) {
		console.error('Error posting blog', error);
		return false;
	}
}

export async function editEvent(payload) {
	try {
		const url = `${import.meta.env.VITE_API_URL}/admin-edit-event`;
		const response = await fetch(url, {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json'
			},
			body: JSON.stringify(payload),
			credentials: 'include'
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
		console.error('Error adding source', error);
		return {
			ok: false,
			status: 500,
			data: { success: false, message: 'Failed to add source' }
		};
	}
}

export async function approveEvent(id) {
	try {
		const url = `${import.meta.env.VITE_API_URL}/admin-approve-event`;
		const response = await fetch(url, {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json'
			},
			body: JSON.stringify({ id }),
			credentials: 'include'
		});

		const data = await response.json();
		return {
			ok: response.ok,
			status: response.status,
			data: data
		};
	} catch (error) {
		console.error('Error approving event', error);
		return {
			ok: false,
			status: 500,
			data: { success: false, message: 'Failed to approve event' }
		};
	}
}

export async function unapproveEvent(id) {
	try {
		const url = `${import.meta.env.VITE_API_URL}/admin-unapprove-event`;
		const response = await fetch(url, {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json'
			},
			body: JSON.stringify({ id }),
			credentials: 'include'
		});

		const data = await response.json();
		return {
			ok: response.ok,
			status: response.status,
			data: data
		};
	} catch (error) {
		console.error('Error unapproving event', error);
		return {
			ok: false,
			status: 500,
			data: { success: false, message: 'Failed to approve event' }
		};
	}
}

export async function uploadPhoto(formData) {
	try {
		const url = `${import.meta.env.VITE_API_URL}/admin-upload-photo`;
		const response = await fetch(url, {
			method: 'POST',
			body: formData,
			credentials: 'include'
		});

		const data = await response.json();
		return {
			ok: response.ok,
			status: response.status,
			data: data
		};
	} catch (error) {
		console.error('Error adding photo', error);
		return {
			ok: false,
			status: 500,
			data: { success: false, message: 'Failed to upload photo' }
		};
	}
}

export async function uploadYoutube(formData) {
	try {
		const url = `${import.meta.env.VITE_API_URL}/admin-upload-youtube`;
		const response = await fetch(url, {
			method: 'POST',
			body: formData,
			credentials: 'include'
		});

		const data = await response.json();
		return {
			ok: response.ok,
			status: response.status,
			data: data
		};
	} catch (error) {
		console.error('Error adding youtube', error);
		return {
			ok: false,
			status: 500,
			data: { success: false, message: 'Failed to upload youtube' }
		};
	}
}

export async function deleteMedia(id) {
	try {
		const url = `${import.meta.env.VITE_API_URL}/admin-delete-media`;
		const response = await fetch(url, {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json'
			},
			body: JSON.stringify({ id }),
			credentials: 'include'
		});

		const data = await response.json();
		return {
			ok: response.ok,
			status: response.status,
			data: data
		};
	} catch (error) {
		console.error('Error deleting media', error);
		return {
			ok: false,
			status: 500,
			data: { success: false, message: 'Failed to delete media' }
		};
	}
}

export async function deleteSource(id) {
	try {
		const url = `${import.meta.env.VITE_API_URL}/admin-delete-source`;
		const response = await fetch(url, {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json'
			},
			body: JSON.stringify({ id }),
			credentials: 'include'
		});

		const data = await response.json();
		return {
			ok: response.ok,
			status: response.status,
			data: data
		};
	} catch (error) {
		console.error('Error deleting source', error);
		return {
			ok: false,
			status: 500,
			data: { success: false, message: 'Failed to delete source' }
		};
	}
}

export async function addSource(payload) {
	try {
		const url = `${import.meta.env.VITE_API_URL}/admin-add-source`;
		const response = await fetch(url, {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json'
			},
			body: JSON.stringify(payload),
			credentials: 'include'
		});

		const data = await response.json();
		return {
			ok: response.ok,
			status: response.status,
			data: data
		};
	} catch (error) {
		console.error('Error adding source', error);
		return {
			ok: false,
			status: 500,
			data: { success: false, message: 'Failed to add source' }
		};
	}
}
