(() => {
	const storageLastRepoKey = "last_repo"

	function onRepoOpen(repoItem) {
		window.localStorage.setItem(storageLastRepoKey, repoItem.id)
	}

	document.addEventListener('DOMContentLoaded', function() {
		const lastRepo = window.localStorage.getItem(storageLastRepoKey)
		if (lastRepo !== null) {
			const lastRepoItem = document.getElementById(lastRepo);
			lastRepoItem.classList.add("active")
		}

		var elems = document.querySelectorAll('.collapsible');
		M.Collapsible.init(elems, {onOpenStart: onRepoOpen});
	});

	function onDeleteTag(event) {
		console.log("in on delete tag");

		event.preventDefault();
		const source = event.target || event.srcElement;

		const url = `/api/${source.getAttribute("data-reponame")}/${source.getAttribute("data-tagname")}`
		sendAPIRequest("DELETE", url, null, "Successfully deleted tag", "Failed to delete tag: ");
	}

	const deleteTagButtons = document.getElementsByClassName("tagDeleteButton");
	for (const button of deleteTagButtons) {
		button.onclick = onDeleteTag;
	}
})();