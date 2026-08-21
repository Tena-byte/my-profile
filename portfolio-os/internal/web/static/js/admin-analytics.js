async function loadAnalytics() {
    try {
        const response = await fetch("/api/admin/analytics");

        if (!response.ok) {
            throw new Error("Failed to load analytics");
        }

        const data = await response.json();

        document.getElementById("total-visits").textContent =
            data.total_visits;

        document.getElementById("unique-visitors").textContent =
            data.unique_visitors;

        document.getElementById("average-duration").textContent =
            data.average_duration;

        renderDeviceBreakdown(data.device_breakdown);
        renderPageViews(data.page_views);
        renderRecentVisits(data.recent_visits);

    } catch (error) {
        console.error("Analytics:", error);
    }
}

function renderDeviceBreakdown(devices) {
    const container = document.getElementById("device-breakdown");

    container.innerHTML = "";

    const entries = Object.entries(devices);

    if (entries.length === 0) {
        container.innerHTML = "<p>No data yet.</p>";
        return;
    }

    entries.forEach(([device, count]) => {
        const item = document.createElement("div");

        item.className = "analytics-row";

        item.innerHTML = `
            <span>${device}</span>
            <strong>${count}</strong>
        `;

        container.appendChild(item);
    });
}

function renderPageViews(pages) {
    const container = document.getElementById("page-views");

    container.innerHTML = "";

    const entries = Object.entries(pages);

    if (entries.length === 0) {
        container.innerHTML = "<p>No data yet.</p>";
        return;
    }

    entries
        .sort((a, b) => b[1] - a[1])
        .forEach(([page, count]) => {
            const item = document.createElement("div");

            item.className = "analytics-row";

            item.innerHTML = `
                <span>${page}</span>
                <strong>${count}</strong>
            `;

            container.appendChild(item);
        });
}

function renderRecentVisits(visits) {
    const tbody = document.getElementById("recent-visits");

    tbody.innerHTML = "";

    if (!visits || visits.length === 0) {
        tbody.innerHTML = `
            <tr>
                <td colspan="5">No visits yet.</td>
            </tr>
        `;

        return;
    }

    visits.forEach((visit) => {
        const row = document.createElement("tr");

        const date = new Date(visit.timestamp);

        row.innerHTML = `
            <td>${visit.page}</td>
            <td>${visit.device_category}</td>
            <td>${visit.ip_address}</td>
            <td>${formatDuration(visit.duration)}</td>
            <td>${date.toLocaleString()}</td>
        `;

        tbody.appendChild(row);
    });
}

function formatDuration(duration) {
    const seconds = Math.floor(duration / 1000000000);

    if (seconds < 60) {
        return `${seconds} sec`;
    }

    const minutes = Math.floor(seconds / 60);
    const remainingSeconds = seconds % 60;

    if (remainingSeconds === 0) {
        return `${minutes} min`;
    }

    return `${minutes} min ${remainingSeconds} sec`;
}

loadAnalytics();

setInterval(loadAnalytics, 30000);