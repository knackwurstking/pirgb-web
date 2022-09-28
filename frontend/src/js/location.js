export default (function () {
  // TODO: "rpi-server" 50831 and "http:" are just placeholder for now
  // TODO: need to handle the file: protocol (open settings page for set server url)
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
