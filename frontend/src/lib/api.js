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

// Admin

export async function login(username, password) {
	try {
		const url = '/api/login';
		const response = await fetch(url, {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json'
			},
			body: JSON.stringify({ username, password })
		});

		if (!response.ok) {
			const data = await response.json();
			console.error('Server error:', data.message);
		}
		return response.ok;
	} catch (error) {
		console.error('Error logging in', error);
		return false;
	}
}

export async function authMe() {
	try {
		const url = '/api/auth-me';
		const response = await fetch(url, {});
		const data = await response.json();

		if (!response.ok) {
			console.error('Server error:', data.error || data.message);
		}

		return data.authenticated;
	} catch (error) {
		console.error('Error fetching auth me', error);
		return false;
	}
}

export async function getAdminEvents() {
	try {
		const url = '/api/admin-events';
		const response = await fetch(url, {
			method: 'GET',
			headers: {
				'Content-Type': 'application/json'
			}
		});
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

export async function editEvent(payload) {
	try {
		const url = '/api/admin-edit-event';
		const response = await fetch(url, {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json'
			},
			body: JSON.stringify(payload)
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
		const url = '/api/admin-approve-event';
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
		const url = '/api/admin-unapprove-event';
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
		for (let [key, value] of formData.entries()) {
			console.log(`${key}:`, value);
		}

		const url = '/api/admin-upload-photo';
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
		const url = '/api/admin-upload-youtube';
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
		const url = '/api/admin-delete-media';
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
		const url = '/api/admin-delete-source';
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
		const url = '/api/admin-add-source';
		const response = await fetch(url, {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json'
			},
			body: JSON.stringify(payload)
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

export async function getBooks() {
    try {
        const url = '/api/admin-get-books'
        const response = await fetch(url)
        const data = await response.json()

        if (!response.ok) {
            console.error('Server error:', data.error || data.message)
        }
        return {
            ok: true,
            data: data,
            status: response.status
        } 
    } catch (error) {
        console.error('Error fetching books', error)
        return {
            ok: false, 
            status: 500
        }
    }
}

export async function submitNewBook(payload) {
    try {
        const url = '/api/admin-add-book';
        const response = await fetch(url, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(payload)
        })

        if (!response.ok) {
            console.error("Server error:", data.error || data.message)
        }

        return {
            ok: response.ok,
            status: response.status
        }
    } catch (error) {
        console.error('Error addming souce', error)
        return {
            ok: false, 
            status: 500, 
        }
    }
}

export async function editBook(payload) {
    console.log(payload)
    try {
        const url = '/api/admin-edit-book';
        const response = await fetch(url, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(payload)
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
        console.error('Error editing book', error);
        return {
            ok: false,
            status: 500,
            data: { success: false, message: 'Failed to edit book' }
        };
    }
}
