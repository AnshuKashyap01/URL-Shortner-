const signupBtn = document.getElementById("signupBtn");

signupBtn.addEventListener("click", async () => {

    const username = document.getElementById("username").value.trim();
    const email = document.getElementById("email").value.trim();
    const password = document.getElementById("password").value.trim();

    if (username === "") {
        showToast("Please enter a username.", "error");
        return;
    }

    if (email === "") {
        showToast("Please enter your email.", "error");
        return;
    }

    const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;

    if (!emailRegex.test(email)) {
        showToast("Please enter a valid email address.", "error");
        return;
    }

    if (password.length < 6) {
        showToast("Password must be at least 6 characters.", "error");
        return;
    }
    signupBtn.disabled = true;
    signupBtn.innerHTML = "⏳ Creating Account...";



    try {

        const response = await fetch("/signup", {

            method: "POST",

            headers: {

                "Content-Type": "application/json"

            },

            body: JSON.stringify({

                username,
                email,
                password

            })

        });

        const data = await response.json();

       if (!response.ok) {

            
            signupBtn.disabled = false;
            signupBtn.innerHTML = "Create Account";

            showToast(data.error, "error");
            return;
        }

        showToast("Account created successfully!");
          signupBtn.disabled = false;
        signupBtn.innerHTML = "Create Account";

        window.location.href = "/login";

    } catch (err) {

         signupBtn.disabled = false;
        signupBtn.innerHTML = "Create Account";

        showToast("Something went wrong.", "error");

    }

});