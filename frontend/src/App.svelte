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
  <section class="sections">
    {#each devices as device}
      {#each device.sections as section}
        <div style="margin: 0.5rem;">
          <SectionCard
            host={device.host}
            port={device.port}
            sectionID={section.id}
          />
        </div>
      {/each}
    {/each}
  </section>
</main>

<style>
  :global(body),
  :global(html) {
    overflow: hidden;
  }

  main {
    overflow-x: hidden;
    overflow-y: auto;
    scroll-behavior: smooth;
    padding: 0.5rem 1rem;
    scroll-snap-type: y mandatory;
  }

  section {
    scroll-snap-align: center;
    margin: 0.5rem 0;
    width: 100%;
    height: 10rem;
  }

  section:first-child {
    scroll-snap-align: start;
  }

  section:last-child {
    scroll-snap-align: end;
  }
</style>
