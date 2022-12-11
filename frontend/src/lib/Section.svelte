<script>
  /** @type {string} */
  export let host;
  /** @type {number} */
  export let port;

  /** @type {Section} */
  export let section;

  /** @type {boolean} */
  export let open = false;

  /** @type {boolean} */
  export let online = false;

  async function toggleOn() {
    const resp = await fetch(
      `/api/v1/devices/${host}/${section.id}/pwm`,
      {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          pulse: section.pulse || section.lastPulse || 100,
        }),
      }
    );

    if (!resp.ok) {
      console.warn(`toggle "${host}/${section.id}" on failed: ${resp.status} (${resp.statusText})`);
    }
  }

  async function toggleOff() {
    const resp = await fetch(
      `/api/v1/devices/${host}/${section.id}/pwm`,
      {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          pulse: 0,
        }),
      }
    );

    if (!resp.ok) {
      console.warn(`toggle "${host}/${section.id}" on failed: ${resp.status} (${resp.statusText})`);
    }
  }
</script>

<div class="section" class:open class:online>
  <h3>{host}:{port} [ID: {section.id}]</h3>
  <pre>Color: {section.color}</pre>
  <pre>Pulse: {section.pulse || 100}</pre>
  <pre>Last Pulse: {section.lastPulse}</pre>
  <div class="actions">
    <button class="off" on:click={toggleOff}>OFF</button>
    <button class="on" on:click={toggleOn}>ON</button>
  </div>
</div>

<style>
  div.section {
    position: relative;
    width: 100%;
    height: fit-content;
    margin: 8px 0;
    border: 1px solid red;
    width: 100%;
    text-align: center;
  }

  div.section.open {
    border: 1px solid var(--color-border);
  }

  div.section h3 {
    text-decoration: underline;
  }

  div.section div.actions {
    border-top: 1px solid red;
    padding: 4px 0;
    display: flex;
    justify-content: center;
  }

  div.section.open div.actions {
    border-top: 1px solid var(--color-border);
  }

  div.section div.actions button.on,
  div.section div.actions button.off {
    width: 10ch;
    height: 4ch;
    margin: 4px 8px;
  }

  div.section div.actions button.off {
    background: rgba(255, 0, 0, 0.9);
  }

  div.section div.actions button.on {
    background: rgba(0, 255, 0, 0.9);
  }
</style>
