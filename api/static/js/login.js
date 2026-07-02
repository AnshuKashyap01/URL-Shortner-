const loginBtn = document.getElementById("loginBtn");

loginBtn.addEventListener("click", async () => {

    const username = document.getElementById("username").value.trim();
    const password = document.getElementById("password").value.trim();



    if (username === "") {

        showToast("Please enter your username.", "error");

        return;
    }

    if (password === "") {

        showToast("Please enter your password.", "error");

        return;
    }

    loginBtn.disabled = true;
    loginBtn.innerHTML = "⏳ Logging in...";

    try {

        const response = await fetch("/login", {

            method: "POST",

            headers: {
                "Content-Type": "application/json"
            },

            body: JSON.stringify({
                username,
                password
            })

        });

        const data = await response.json();


        if (!response.ok) {

            loginBtn.disabled = false;
            loginBtn.innerHTML = "Login";

            showToast(data.error, "error");

            return;
        }

        localStorage.setItem("token", data.token);
        localStorage.setItem("username", data.username);

        showToast("Login Successful!");

        loginBtn.disabled = false;
        loginBtn.innerHTML = "Login";

        window.location.href = "/";

    } catch (err) {

        showToast("There might be some error")

    }

});