export default (function () {
  const data = {
    hostname: window.location.hostname,
    port: window.location.port,
    protocol: window.location.protocol,

    get origin() {
      return `${this.protocol}//${this.hostname}:${this.port}`;
    },
  };

  return data;
})();
