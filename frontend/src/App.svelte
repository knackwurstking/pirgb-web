<script>
  import { onMount } from "svelte"
  import SectionCard from "./lib/SectionCard.svelte"
  import api from "./js/api"

  /** @type {import("./js/api").Devices} */
  let devices = []

  onMount(() => {
    console.log(`[App.svelte] onMount: get devices...`)
    api.getDevices().then(res => devices = res)
  })
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
    <div style="height: 1.5rem;"></div>
</main>

<style>
  main {
    overflow-x: hidden;
    overflow-y: auto;
    scroll-behavior: smooth;
    padding: 0.5rem 0.25rem;
    /*scroll-snap-type: y mandatory;*/
    height: 100vh;
  }

  section {
    /*scroll-snap-align: center;*/
    padding: 0.5rem 0.5rem;
    width: 100%;
    height: 8rem;
  }

  /*
  section:first-child {
    scroll-snap-align: start;
  }

  section:last-child {
    scroll-snap-align: end;
  }
  */
</style>
