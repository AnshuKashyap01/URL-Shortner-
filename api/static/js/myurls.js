let urlToDelete = "";

async function loadMyUrls() {

    const token = localStorage.getItem("token");

    if (!token) return;

    const response = await fetch("/myurls", {

        headers: {

            "Authorization": "Bearer " + token

        }

    });

    const data = await response.json();

    const table = document.getElementById("urlTable");
    const tableContainer = document.getElementById("urlTableContainer");
    const emptyState = document.getElementById("emptyState");

    table.innerHTML = "";
    if (Object.keys(data.urls).length === 0) {

        emptyState.style.display = "block";

        tableContainer.style.display = "none";

        return;

    }

    emptyState.style.display = "none";

    tableContainer.style.display = "table";

    for (const url of data.urls) {

    table.innerHTML += `

    <tr id="row-${url.short}">

        <td>

            <a href="/go/${url.short}" target="_blank">

                ${url.short}

            </a>

        </td>

        <td>

            <a href="${url.original}" target="_blank">

                ${new URL(url.original).hostname}

            </a>

        </td>

        <td>

            ${url.clicks}

        </td>

        <td>

            <button
                class="qrBtn"
                onclick="window.open('/qr/${url.short}')">

                QR

            </button>

        </td>

        <td>

            <button
                class="deleteBtn"
                onclick="deleteURL('${url.short}')">

                🗑 Delete

            </button>

        </td>

    </tr>

    `;
}

}

loadMyUrls();

function deleteURL(short){

    urlToDelete = short;

    document.querySelector(
        "#deleteModal p"
    ).innerHTML = `

        Are you sure you want to delete
        <strong>${short}</strong>?

        <br><br>

        This action cannot be undone.

    `;

    document.getElementById("deleteModal").style.display = "flex";

}


document.getElementById("cancelDelete").addEventListener("click", () => {

    document.getElementById("deleteModal").style.display = "none";

});

document.getElementById("confirmDelete").addEventListener("click", async () => {

    const token = localStorage.getItem("token");

    const response = await fetch(`/url/${urlToDelete}`, {

        method: "DELETE",

        headers: {

            Authorization: "Bearer " + token

        }

    });

    const data = await response.json();

    document.getElementById("deleteModal").style.display = "none";

    if (!response.ok) {

        showToast(data.error, "error");

        return;

    }

    showToast("URL deleted successfully");

    loadMyUrls();

});

