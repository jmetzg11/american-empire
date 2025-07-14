<script>
	import { login } from '$lib/api';
	import { fetchAdminEvents } from '$lib/store';
	let { isAuthenticated = $bindable() } = $props();
	let username = $state('');
	let password = $state('');

	const handleLogin = async () => {
		if (await login(username, password)) {
			isAuthenticated = true;
			await fetchAdminEvents();
		} else {
			console.error('Login failed');
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
