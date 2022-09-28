<script>
  import { onMount } from "svelte";
  import SectionCard from "./lib/SectionCard.svelte";
  import api from "./js/api";

  /** @type {import("./js/api").Devices} */
  let devices = [];

  onMount(() => {
    console.log(`[App.svelte] onMount: get devices...`);
    api.getDevices().then((res) => (devices = res));
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
    /*scroll-snap-type: y mandatory;*/
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
    margin: 0.5rem 0;
    width: 90%;
    height: 8rem;
  }
</style>
