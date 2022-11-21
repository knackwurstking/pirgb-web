<script>
  import { onDestroy } from "svelte";

  import SectionCard from "./lib/SectionCard.svelte";
  import SettingsDialog from "./lib/SettingsDialog.svelte";

  import Events from "./js/events";
  import Api from "./js/api";

  let initialized = false;

  let settingsOpen = false;
  $: !settingsOpen && initialize();

  /** @type {import("./js/api").Devices} */
  let devices = [];

  async function initialize() {
    Api.getDevices().then((res) => (devices = res));

    if (!initialized) {
      initialized = true
      Events.addEventListener("open", handleOpenEvent)
      return
    }

    // reconnect events if already initialized once...
    Events.connect()
  }

  onDestroy(() => Events.removeEventListener("open", handleOpenEvent));

  const handleOpenEvent = () => {
    console.log(`[App.svelte] handleOnline: get devices...`);
    Api.getDevices().then((res) => (devices = res));
  };
</script>



<main>
  <button
    id="settings"
    on:click={() => {
      settingsOpen = true;
    }}
  />

  {#each devices as device}
    {#each device.sections as section}
      <section>
        <SectionCard
          host={device.host}
          port={device.port}
          sectionID={section.id}
        />
      </section>
    {/each}
  {/each}
</main>

<SettingsDialog bind:active={settingsOpen} />

<style>
  main {
    position: relative;
    display: flex;
    flex-direction: column;
    place-items: center;
    overflow-x: hidden;
    overflow-y: auto;
    scroll-behavior: smooth;
    padding: 3rem 0.25rem;
    height: 100vh;
    transition: padding 0.25s ease;
  }

  main::-webkit-scrollbar {
    display: none;
  }

  section {
    display: flex;
    justify-content: center;
    place-items: center;
    padding: 0.5rem 0.5rem;
    width: 100%;
    height: 8rem;
  }

  button#settings {
    position: fixed;
    z-index: 1;
    bottom: 0;
    right: 0;
    width: 3rem;
    height: 3rem;
    cursor: pointer;
    opacity: 0;
    border-top-left-radius: 50%;
  }
</style>
