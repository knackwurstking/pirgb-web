export default (function () {
  const data = {
    config: {
      "http:": {
        hostname: "127.0.0.1",
        port: 50831,
      },
      "https:": {
        hostname: "127.0.0.1",
        port: 50832,
      },
    },

    protocol:
      window.location.protocol === "file:" ? "http:" : window.location.protocol,

    get origin() {
      return `${this.protocol}//${this.config[this.protocol].hostname}:${
        this.config[this.protocol].port
      }`;
    },

    saveConfig() {
      window.localStorage.setItem("location.config", JSON.stringify(this.config))
    }
  };

  if (window.location.hostname) {
    if (data.config["http:"].hostname !== "127.0.0.1") {
      data.config["http:"].hostname = window.location.hostname;
    }

    if (data.config["https:"].hostname !== "127.0.0.1") {
      data.config["https:"].hostname = window.location.hostname;
    }
  }

  try {
    const config = JSON.parse(window.localStorage.getItem("location.config"));
    if (config) {
      data.config = config;
    }
  } catch (err) {
    console.warn(err);
  }

  return data;
})();
