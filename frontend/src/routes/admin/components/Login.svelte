<script>
	import { fetchAdminEvents } from '$lib/store';
	let { isAuthenticated = $bindable() } = $props();
	let username = $state('');
	let password = $state('');

	const handleLogin = async () => {
		const url = `${import.meta.env.VITE_API_URL}/login`;
		const response = await fetch(url, {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json'
			},
			credentials: 'include',
			body: JSON.stringify({ username, password })
		});
		if (response.ok) {
			isAuthenticated = true;
			await fetchAdminEvents();
		}
	};
</script>

<div class="container-wrap">
	<div class="card-style">
		<div>
			<label class="label" for="username">Username</label>
			<input class="input" type="text" id="username" bind:value={username} />
		</div>
		<div>
			<label class="label" for="password">Password</label>
			<input class="input" type="password" id="password" bind:value={password} />
		</div>
		<div>
			<button class="btn mt-4" onclick={handleLogin}>Login</button>
		</div>
	</div>
</div>
