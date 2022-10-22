<script>
  import { onMount, onDestroy } from "svelte";

  import SectionCard from "./lib/SectionCard.svelte";

  import Events from "./js/events";
  import Api from "./js/api";

  /** @type {import("./js/api").Devices} */
  let devices = [];

  const handleOpen = () => {
    console.log(`[App.svelte] handleOnline: get devices...`);
    Api.getDevices().then((res) => (devices = res));
  };

  onMount(() => {
    Api.getDevices().then((res) => (devices = res));
    Events.addEventListener("open", handleOpen);
  });

  onDestroy(() => {
    Events.removeEventListener("open", handleOpen);
  });
</script>

<main>
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

<style>
  main {
    display: flex;
    flex-direction: column;
    place-items: center;
    overflow-x: hidden;
    overflow-y: auto;
    scroll-behavior: smooth;
    padding: 2rem 0.25rem;
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
</style>
