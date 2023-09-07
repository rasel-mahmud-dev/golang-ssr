// login.js



document.addEventListener('DOMContentLoaded', function () {
    const loginForm = document.querySelector('#login-form');
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
});