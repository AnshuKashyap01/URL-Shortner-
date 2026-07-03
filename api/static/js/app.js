const button = document.getElementById("generate");

button.addEventListener("click", async function () {


    const url = document.getElementById("url").value.trim();
    const short = document.getElementById("short").value.trim();

    const token = localStorage.getItem("token");
    
    // URL cannot be empty
    if (url === "") {

        showToast("Please enter a URL.", "error");

        return;
    }

    // Basic URL validation
    try {

        new URL(url);

    } catch {

        showToast("Please enter a valid URL.", "error");

        return;
    }

    // Alias validation (optional)
    if (short !== "" && !/^[a-zA-Z0-9_-]+$/.test(short)) {

        showToast("Alias can only contain letters, numbers, '-' and '_'.", "error");

        return;
    }


    button.disabled = true;
    button.innerText = "Generating...";


    
    const response = await fetch("/api/v1", {

        method: "POST",

        headers: {

            "Content-Type": "application/json",
            "Authorization": "Bearer " + token

        },

        body: JSON.stringify({

            url: url,
            short: short

        })

    });

    const data = await response.json();
    console.log(data);

    const result = document.getElementById("result");

    if (!response.ok) {

        button.disabled = false;
        button.innerText = "Generate Short URL";

        showToast(data.error, "error");

        return;
    }

    result.innerHTML = `

    <h2 class="successTitle">

        ✅ URL Created Successfully

    </h2>

    <input
        id="generatedUrl"
        value="${data.short}"
        readonly>

    <div class="resultButtons">

        <button id="copyBtn">

            📋 Copy

        </button>

        <button
            id="openBtn"
           onclick="window.open('${data.short}','_blank')">

            🔗 Open

        </button>

    </div>

    <div class="qrSection">

        <img
           src="${data.qr}"
            width="180"
            alt="QR Code">

    </div>

    <div class="resultInfo">

        <p>

            ⏰ Expires in
            <strong>${data.expiry} Hours</strong>

        </p>

        <p>

            📊 Remaining Requests:
            <strong>${data.rate_limit}</strong>

        </p>

    </div>

`;

    loadMyUrls();

    document.getElementById("copyBtn").addEventListener("click", () => {

        const url = document.getElementById("generatedUrl");

        navigator.clipboard.writeText(url.value);

        showToast("URL Copied!");

    });

    button.disabled = false;
    button.innerText = "Generate Short URL";

});