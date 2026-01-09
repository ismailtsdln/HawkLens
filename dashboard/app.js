document.addEventListener('DOMContentLoaded', () => {
    initCharts();
    fetchData();

    document.getElementById('refresh-btn').addEventListener('click', fetchData);
});

let platformChart, sentimentChart;

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

async function fetchData() {
    try {
        // Mocking API fetch for demonstration since backend might not be live
        const mockData = [
            { platform: 'twitter', data_type: 'tweet', data: { text: 'OSINT is great!' }, created_at: new Date() },
            { platform: 'youtube', data_type: 'video', data: { title: 'HawkLens Demo' }, created_at: new Date() },
            { platform: 'reddit', data_type: 'post', data: { title: 'Deep analysis' }, created_at: new Date() },
            { platform: 'instagram', data_type: 'post', data: { caption: 'Visualizing data' }, created_at: new Date() },
            { platform: 'tiktok', data_type: 'trend', data: { hashtag: '#osint' }, created_at: new Date() }
        ];

        updateDashboard(mockData);
    } catch (error) {
        console.error('Fetch error:', error);
    }
}

function updateDashboard(data) {
    document.getElementById('total-results').innerText = data.length;
    document.getElementById('total-scans').innerText = 1; // Mock

    // Update Platform Chart
    const platformCounts = { twitter: 0, youtube: 0, reddit: 0, instagram: 0, tiktok: 0 };
    data.forEach(item => platformCounts[item.platform]++);
    platformChart.data.datasets[0].data = Object.values(platformCounts);
    platformChart.update();

    // Update Table
    const tbody = document.querySelector('#results-table tbody');
    tbody.innerHTML = '';
    data.forEach(item => {
        const row = document.createElement('tr');
        const summary = item.data.text || item.data.title || item.data.caption || item.data.hashtag;
        row.innerHTML = `
            <td><span class="badge badge-${item.platform}">${item.platform}</span></td>
            <td>${item.data_type}</td>
            <td>${summary}</td>
            <td>${new Date(item.created_at).toLocaleTimeString()}</td>
        `;
        tbody.appendChild(row);
    });
}
