export async function GET({ params, url, request }) {
	const backendUrl = `https://empire-backend.fly.dev/api/${params.path}${url.search}`;

	let body = undefined;
	if (request.method !== 'GET') {
		// Use arrayBuffer to preserve binary data
		body = await request.arrayBuffer();
	}

	const response = await fetch(backendUrl, {
		method: request.method,
		headers: request.headers,
		body: body
	});

	return new Response(response.body, {
		status: response.status,
		headers: response.headers
	});
}

export { GET as POST, GET as PUT, GET as DELETE };
