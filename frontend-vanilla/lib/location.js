export default (function() {
    const data = {
        hostname: window.location.hostname || "rpi-server",
        port: window.location.port || 50831,
        protocol:
            window.location.protocol === "file:" ? "http:" : window.location.protocol,

        get origin() {
            return `${this.protocol}//${this.hostname}:${this.port}`;
        },
    };

    return data;
})();
