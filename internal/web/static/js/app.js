
document.addEventListener("DOMContentLoaded", async () => {
    const formatBytes = (bytes) => {
        if (bytes >= 1_000_000_000) {
            return `${(bytes / 1_000_000_000).toFixed(2)} GB`;
        }

        return `${(bytes / 1_000_000).toFixed(2)} MB`;
    };

    const setText = (id, value) => {
        const element = document.getElementById(id);
        if (element) {
            element.textContent = value;
        }
    };

    try {
        const response = await fetch("/api/metrics");

        if (!response.ok) {
            throw new Error(`Metrics API returned ${response.status}`);
        }

        const metrics = await response.json();

        setText(
            "baseline-size",
            formatBytes(metrics.images.baseline.sizeBytes)
        );

        setText(
            "optimized-size",
            formatBytes(metrics.images.distroless.sizeBytes)
        );

        setText(
            "reduction",
            `${metrics.reductions.baselineToDistroless}%`
        );

        console.log("ContainerScope metrics loaded successfully.");
    } catch (error) {
        console.error("Unable to load image metrics:", error);
        setText("baseline-size", "Unavailable");
        setText("optimized-size", "Unavailable");
        setText("reduction", "Unavailable");
    }
});
