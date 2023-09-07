
document.addEventListener('DOMContentLoaded', function () {
    const loginForm = document.querySelector('#login-form');
    if(loginForm) {
        loginForm.addEventListener('submit', function (e) {
            e.preventDefault();

            const username = document.querySelector('#username').value;
            const password = document.querySelector('#password').value;

            // Perform client-side validation
            if (!username || !password) {
                alert('Please fill in both username and password.');
                return;
            }

            console.log("ksadjfh")
        });

    }

    const addPostForm = document.querySelector('#add-article');
    if(addPostForm) {
        addPostForm.addEventListener('submit', function (e) {
            e.preventDefault();

            const title = document.querySelector('#title').value;
            const content = document.querySelector('#content').value;

            fetch("/api/v1/articles/new", {
                method: "POST",
                body: JSON.stringify({
                    title,
                    content
                })
            }).then(res => res.json()).then(data => {
                console.log(data)
            })

        });
    }
});