async function updateCPU() {
    try {
        const response = await fetch("/api/cpu");
        const data = await response.json();

        const cpuElement = document.getElementById("cpu-usage");
        const cpuBar = document.getElementById("cpu-bar");

        if (cpuElement) {
            cpuElement.textContent = data.cpu.toFixed(1) + "%";
        }

        if (cpuBar) {
            cpuBar.style.width = data.cpu + "%";
        }

    } catch (error) {
        console.error("Failed to update CPU:", error);
    }
}

// Update CPU every 2 seconds.
setInterval(updateCPU, 2000);

// Update immediately when the page loads.
updateCPU();


async function updateMemory() {
    try {
        const response = await fetch("/api/memory");
        const data = await response.json();

        const ramUsage = document.getElementById("ram-usage");
        const ramBar = document.getElementById("ram-bar");
        const ramUsed = document.getElementById("ram-used");
        const ramTotal = document.getElementById("ram-total");

        if (ramUsage) {
            ramUsage.textContent = data.usage.toFixed(1) + "%";
        }

        if (ramBar) {
            ramBar.style.width = data.usage + "%";
        }

        if (ramUsed) {
            ramUsed.textContent = data.used.toFixed(1) + " GB used";
        }

        if (ramTotal) {
            ramTotal.textContent = data.total.toFixed(1) + " GB total";
        }

    } catch (error) {
        console.error("Failed to update memory:", error);
    }
}

// Update RAM every 2 seconds.
setInterval(updateMemory, 2000);

// Update immediately when the page loads.
updateMemory();