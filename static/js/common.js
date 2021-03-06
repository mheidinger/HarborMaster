const url = new URL(location.href);
url.searchParams.delete("error");
url.searchParams.delete("success");
history.replaceState(null, document.title, url.href);

function sendAPIRequest(method, path, body, success, error) {
	const fetchUrl = new URL(location.href);
	fetchUrl.pathname = path;
	fetchUrl.search = ""

	fetch(fetchUrl.href, {
		method: method,
		body: JSON.stringify(body)
	}).then(response => response.json()
	).then(json => {
		if (json["error"]) {
			reloadError(error + json["error"]);
		} else {
			reloadSuccess(success);
		}
	});
}

function reloadSuccess(message) {
	let url = new URL(location.href);
	url.searchParams.delete("error");
	url.searchParams.delete("success");
	url.searchParams.append("success", message);
	location.replace(url.href);
}

function reloadError(message) {
	let url = new URL(location.href);
	url.searchParams.delete("error");
	url.searchParams.delete("success");
	url.searchParams.append("error", message);
	location.replace(url.href);
}