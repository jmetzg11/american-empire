import { writable } from 'svelte/store';
import { getEvents, getAdminEvents } from './api';

export const events = writable([]);
export const UIState = writable('general');

export async function fetchEvents() {
	const data = await getEvents();
	console.log(data);
	events.set(data);
	UIState.set('general');
}

export async function fetchAdminEvents() {
	const data = await getAdminEvents();
	events.set(data);
	UIState.set('admin');
}
