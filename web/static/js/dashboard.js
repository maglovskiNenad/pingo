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