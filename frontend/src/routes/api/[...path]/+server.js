export async function GET({ params, url, request }) {
	const backendUrl = `https://empire-backend.fly.dev/api/${params.path}${url.search}`;

	const response = await fetch(backendUrl, {
		method: request.method,
		headers: request.headers,
		body: request.method !== 'GET' ? await request.text() : undefined
	});

	return new Response(response.body, {
		status: response.status,
		headers: response.headers
	});
}

export { GET as POST, GET as PUT, GET as DELETE };
