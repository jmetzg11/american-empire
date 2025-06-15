import { writable } from 'svelte/store';
import { getEvents } from './api';

export const events = writable([]);

export async function fetchEvents() {
	const data = await getEvents();
	events.set(data);
}
