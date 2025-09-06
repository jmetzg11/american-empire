import { writable } from 'svelte/store';
import { getEvents, getAdminEvents } from './api';

export const events = writable([]);

export async function fetchEvents() {
	const data = await getEvents();
	events.set(data);
	UIState.set('general');
}

