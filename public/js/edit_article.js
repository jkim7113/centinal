const form = document.getElementById("edit-article-form");

form.addEventListener("submit", (e) => {
    e.preventDefault();
    
    const target = form.action;
    const title = document.getElementById("edit-article-title").value;
    const body = document.getElementById("edit-article-body").value;
    const category = document.getElementById("edit-article-category").value;
    
    fetch(target, {
        method: 'PUT',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ Title: title, Body: body, Category: category, Thumbnail: "", }),
    })
    .then(res => {
        if (res.ok) location.reload();
    });
});