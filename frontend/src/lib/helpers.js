export function formatDate(dateString) {
	return dateString.split('T')[0];
}

export function extractYouTubeId(url) {
	const regex =
		/(?:youtube\.com\/(?:[^\/]+\/.+\/|(?:v|e(?:mbed)?|shorts)\/|.*[?&]v=)|youtu\.be\/)([^"&?\/\s]{11})/;
	const match = url.match(regex);
	return match ? match[1] : null;
}
