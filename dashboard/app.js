document.addEventListener('DOMContentLoaded', () => {
    initCharts();
    
    document.getElementById('refresh-btn').addEventListener('click', () => {
        const query = prompt("Enter search query:");
        if (query) startStreamingScan(query);
    });
});

let platformChart, sentimentChart;
let allData = [];

function initCharts() {
    const platformCtx = document.getElementById('platformChart').getContext('2d');
    platformChart = new Chart(platformCtx, {
        type: 'doughnut',
        data: {
            labels: ['Twitter', 'YouTube', 'Reddit', 'Instagram', 'TikTok'],
            datasets: [{
                data: [0, 0, 0, 0, 0],
                backgroundColor: ['#1da1f2', '#ff0000', '#ff4500', '#e1306c', '#000000'],
                borderWidth: 0
            }]
        },
        options: {
            plugins: { legend: { position: 'bottom', labels: { color: '#94a3b8' } } }
        }
    });

    const sentimentCtx = document.getElementById('sentimentChart').getContext('2d');
    sentimentChart = new Chart(sentimentCtx, {
        type: 'bar',
        data: {
            labels: ['Positive', 'Neutral', 'Negative'],
            datasets: [{
                label: 'Results',
                data: [0, 0, 0],
                backgroundColor: ['#10b981', '#64748b', '#ef4444'],
                borderRadius: 8
            }]
        },
        options: {
            scales: {
                y: { beginAtZero: true, ticks: { color: '#94a3b8' } },
                x: { ticks: { color: '#94a3b8' } }
            },
            plugins: { legend: { display: false } }
        }
    });
}

function startStreamingScan(query) {
    allData = [];
    const tbody = document.querySelector('#results-table tbody');
    tbody.innerHTML = '<tr><td colspan="4" style="text-align:center">Streaming results for: <strong>' + query + '</strong>...</td></tr>';
    
    const eventSource = new EventSource(`/api/v1/scan-stream?query=${encodeURIComponent(query)}`);

    eventSource.onmessage = (event) => {
        const result = JSON.parse(event.data);
        allData.push(result);
        updateDashboard(allData);
    };

    eventSource.onerror = (err) => {
        console.error("EventSource failed:", err);
        eventSource.close();
    };
}

function updateDashboard(data) {
    document.getElementById('total-results').innerText = data.length;
    
    // Update Platform Chart
    const platformCounts = { twitter: 0, youtube: 0, reddit: 0, instagram: 0, tiktok: 0 };
    data.forEach(item => platformCounts[item.platform.toLowerCase()]++);
    platformChart.data.datasets[0].data = Object.values(platformCounts);
    platformChart.update();

    // Update Table (Newest first)
    const tbody = document.querySelector('#results-table tbody');
    if (data.length === 1) tbody.innerHTML = ''; // Clear the "Streaming..." message
    
    const item = data[data.length - 1];
    const row = document.createElement('tr');
    const summary = item.data.text || item.data.title || item.data.caption || item.data.hashtag || "No summary";
    
    row.innerHTML = `
        <td><span class="badge badge-${item.platform.toLowerCase()}">${item.platform}</span></td>
        <td>${item.data_type}</td>
        <td>${summary}</td>
        <td>${new Date().toLocaleTimeString()}</td>
    `;
    tbody.prepend(row);
}
